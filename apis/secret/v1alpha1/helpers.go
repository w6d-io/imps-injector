/*
Copyright 2020 WILDCARD SA.

Licensed under the WILDCARD SA License, Version 1.0 (the "License");
WILDCARD SA is register in french corporation.
You may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.w6d.io/licenses/LICENSE-1.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is prohibited.
Created on 12/02/2022
*/

package v1alpha1

import (
	"context"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"

	corev1 "k8s.io/api/core/v1"

	"github.com/w6d-io/imps-injector/pkg/toolx"
	"github.com/w6d-io/x/logx"
)

func (s *ServiceAccountSelector) OpIn(m map[string]string) bool {
	for _, value := range s.Values {
		if v, ok := m[s.Key]; ok && v == value {
			return true
		}
	}
	return false
}

func (s *ServiceAccountSelector) OpExists(m map[string]string) bool {
	_, ok := m[s.Key]
	return ok
}

func (s *ServiceAccountSelector) OpContains(m map[string]string) bool {
	for _, value := range s.Values {
		if v, ok := m[s.Key]; ok && strings.Contains(v, value) {
			return true
		}
	}
	return false
}

func (s *ServiceAccountSelector) OpStartWith(m map[string]string) bool {
	for _, value := range s.Values {
		if v, ok := m[s.Key]; ok && strings.HasPrefix(v, value) {
			return true
		}
	}
	return false
}

func (s *ServiceAccountSelector) OpNotIn(m map[string]string) bool {
	return !s.OpIn(m)
}

func (s *ServiceAccountSelector) OpDoesNotExist(m map[string]string) bool {
	return !s.OpExists(m)
}

func (s *ServiceAccountSelector) OpDoesNotContain(m map[string]string) bool {
	return !s.OpContains(m)
}

func (s *ServiceAccountSelector) OpDoesNotStartWith(m map[string]string) bool {
	return !s.OpStartWith(m)
}

func (in ImagePullSecretInjectorSpec) Match(ctx context.Context, r client.Client, s *corev1.ServiceAccount) *corev1.ServiceAccountList {
	return in.Matches(ctx, r, &corev1.ServiceAccountList{Items: []corev1.ServiceAccount{
		*s,
	}})
}

type op func(map[string]string) bool
type ops []op

func (in ImagePullSecretInjectorSpec) Matches(ctx context.Context, r client.Client, l *corev1.ServiceAccountList) *corev1.ServiceAccountList {
	log := logx.WithName(ctx, "Match")
	sal := &serviceAccountList{
		r:    r,
		imps: in,
		list: &corev1.ServiceAccountList{
			TypeMeta: l.TypeMeta,
			ListMeta: l.ListMeta,
		},
	}

	var (
		labelProcess      = getOps(in.LabelSelector)
		annotationProcess = getOps(in.AnnotationSelector)
		//nsLabelProcess      = getOps(in.NamespaceLabelSelector)
		//nsAnnotationProcess = getOps(in.NamespaceAnnotationSelector)
	)
	for _, sa := range l.Items {
		if labelProcess.exec(sa.Labels) || annotationProcess.exec(sa.Annotations) ||
			toolx.InArray(sa.Namespace, in.Namespaces) {
			log.V(3).Info("add sa in list", "name", sa.Name, "namespace", sa.Namespace)
			sal.list.Items = append(sal.list.Items, sa)
		}
	}
	return sal.list
}

func (ps ops) exec(value map[string]string) (res bool) {
	for _, p := range ps {
		res = res || p(value)
	}
	return
}

func getOps(in []ServiceAccountSelector) ops {
	var ps ops
	for _, i := range in {
		switch i.Operator {
		case OpIn:
			ps = append(ps, i.OpIn)
		case OpExists:
			ps = append(ps, i.OpExists)
		case OpContains:
			ps = append(ps, i.OpContains)
		case OpStartWith:
			ps = append(ps, i.OpStartWith)
		case OpNotIn:
			ps = append(ps, i.OpNotIn)
		case OpDoesNotExist:
			ps = append(ps, i.OpDoesNotExist)
		case OpDoesNotContain:
			ps = append(ps, i.OpDoesNotContain)
		case OpDoesNotStartWith:
			ps = append(ps, i.OpDoesNotStartWith)
		}
	}
	return ps
}
