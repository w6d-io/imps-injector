# Build the manager binary
FROM golang:1.17 as builder
ARG GOVERSION=1.17
ARG VCS_REF
ARG BUILD_DATE
ARG VERSION
ENV GO111MODULE="on" \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY apis/ apis/
COPY controllers/ controllers/
COPY internal/ internal/
COPY pkg/ pkg/

# Build
RUN go build \
    -ldflags="-X 'github.com/w6d-io/imps-injector/internal/config.Version=${VERSION}' -X 'github.com/w6d-io/imps-injector/internal/config.Revision=${VCS_REF}' -X 'github.com/w6d-io/imps-injector/internal/config.GoVersion=go${GOVERSION}' -X 'github.com/w6d-io/imps-injector/internal/config.Built=${BUILD_DATE}'" \
    -a -o imps-injector main.go

# Use distroless as minimal base image to package the imps-injector binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/base:nonroot
ARG VCS_REF
ARG BUILD_DATE
ARG VERSION
ARG USER_EMAIL="david.alexandre@w6d.io"
ARG USER_NAME="David ALEXANDRE"
LABEL maintainer="${USER_NAME} <${USER_EMAIL}>" \
        io.w6d.vcs-ref=$VCS_REF       \
        io.w6d.build-date=$BUILD_DATE \
        io.w6d.version=$VERSION

WORKDIR /
COPY --from=builder /workspace/imps-injector .
USER 65532:65532

ENTRYPOINT ["/imps-injector"]
