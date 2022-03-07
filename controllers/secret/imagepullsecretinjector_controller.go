/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package secret

import (
    "context"
    "fmt"
    "k8s.io/apimachinery/pkg/api/meta"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/utils/pointer"
    "sigs.k8s.io/controller-runtime/pkg/controller"

    "github.com/google/uuid"
    corev1 "k8s.io/api/core/v1"
    apierrors "k8s.io/apimachinery/pkg/api/errors"
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/apimachinery/pkg/types"
    "k8s.io/client-go/util/retry"
    ctrl "sigs.k8s.io/controller-runtime"
    "sigs.k8s.io/controller-runtime/pkg/client"
    "sigs.k8s.io/controller-runtime/pkg/handler"
    "sigs.k8s.io/controller-runtime/pkg/reconcile"
    "sigs.k8s.io/controller-runtime/pkg/source"

    impsi "github.com/w6d-io/imps-injector/apis/secret/v1alpha1"
    "github.com/w6d-io/imps-injector/pkg/toolx"
    "github.com/w6d-io/x/logx"
)

// ImagePullSecretInjectorReconciler reconciles a ImagePullSecretInjector object
type ImagePullSecretInjectorReconciler struct {
    client.Client
    Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=secret.w6d.io,resources=imagepullsecretinjectors,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=secret.w6d.io,resources=imagepullsecretinjectors/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=secret.w6d.io,resources=imagepullsecretinjectors/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ImagePullSecretInjector object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *ImagePullSecretInjectorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    correlationID := uuid.New().String()
    ctx = context.WithValue(ctx, logx.CorrelationID, correlationID)
    ctx = context.WithValue(ctx, toolx.ContextKeyImps, req.NamespacedName.String())
    log := logx.WithName(ctx, "Reconcile").WithValues("impsi", req.NamespacedName, "correlationId", correlationID)

    ii := new(impsi.ImagePullSecretInjector)
    if err := r.Get(ctx, req.NamespacedName, ii); err != nil {
        if apierrors.IsNotFound(err) {
            log.Info("ImagePullSecretInjector does not exist")
            return ctrl.Result{}, nil
        }
        log.Error(err, "get ImagePullSecretInjector failed")
        return ctrl.Result{}, err
    }
    sas := &corev1.ServiceAccountList{}
    err := r.List(ctx, sas)
    if err != nil {
        if apierrors.IsNotFound(err) {
            log.Info("ImagePullSecretInjector does not exist")
            return ctrl.Result{}, nil
        }
        log.Error(err, "list service account failed")
        return ctrl.Result{}, err
    }
    saList := ii.Spec.Matches(ctx, r.Client, sas)
    if len(saList.Items) == 0 {
        log.Info("nothing to do")
        //err := r.UpdateStatus(ctx, ii, 0, 0, impsi.StatusSynchronized, "nothing to do")
        //if err != nil {
        //	log.Error(err, "update status failed")
        //}
        return ctrl.Result{Requeue: false}, nil
    }
    var done int32
    //reQueued := false
    for _, sa := range saList.Items {
        if err := r.putSecret(ctx, ii, sa); err != nil {
            log.Error(err, "put secret failed")
            //reQueued = true
            if err != nil {
                log.Error(err, "update status failed. will re-queued")
            }
        }
        done++
    }

    var total = int32(len(saList.Items))
    var state impsi.Status
    switch {
    case done < total:
        state = impsi.StatusPartiallySynchronized
    case done == 0:
        state = impsi.StatusDisSynchronized
    default:
        state = impsi.StatusSynchronized
    }
    if err := r.UpdateStatus(ctx, ii, done, total, state, ""); err != nil {
        log.Error(err, "update status failed")
    }
    log.Info("reconciled")
    return ctrl.Result{Requeue: false}, nil
}

// UpdateStatus set the status of tekton resource state
func (r *ImagePullSecretInjectorReconciler) UpdateStatus(ctx context.Context, ii *impsi.ImagePullSecretInjector, count, total int32, state impsi.Status, message string) error {
    log := logx.WithName(ctx, "Reconcile.UpdateStatus")
    var err error
    log.V(1).Info("update status")
    key := client.ObjectKeyFromObject(ii)
    obj := &impsi.ImagePullSecretInjector{}
    if err := r.Get(ctx, key, obj); err != nil {
        if !apierrors.IsNotFound(err) {
            return err
        }
    }
    err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
        obj.Status.Status = pointer.String(fmt.Sprintf("%d/%d", count, total))

        meta.SetStatusCondition(&obj.Status.Conditions, metav1.Condition{
            Type:    string(state),
            Status:  r.GetStatus(state),
            Reason:  string(state),
            Message: message,
        })
        if err := r.Status().Update(ctx, obj); err != nil {
            log.Error(err, "unable to update play status (retry)")
            return err
        }
        return nil
    })
    if err != nil {
        log.Error(err, "unable to update play status")
        return err
    }
    return nil
}

func (r *ImagePullSecretInjectorReconciler) GetStatus(status impsi.Status) metav1.ConditionStatus {
    switch status {
    case impsi.StatusDisSynchronized:
        return metav1.ConditionFalse
    case impsi.StatusSynchronized:
        return metav1.ConditionTrue
    default:
        return metav1.ConditionUnknown
    }
}

func (r *ImagePullSecretInjectorReconciler) putSecret(
    ctx context.Context,
    injector *impsi.ImagePullSecretInjector,
    account corev1.ServiceAccount,
) error {
    log := logx.WithName(ctx, "putSecret")
    nn := types.NamespacedName{Name: account.Name, Namespace: account.Namespace}
    log.V(1).Info("loop on secrets")
    for _, secret := range injector.Spec.Secrets {
        account.Annotations = mapAppend(account.Annotations, secret.Annotations)
        account.Labels = mapAppend(account.Labels, secret.Labels)
        if toolx.InArray(corev1.LocalObjectReference{Name: secret.Name}, account.ImagePullSecrets) {
            log.V(1).Info("secret already present", "secret", secret.Name, "sa", nn.String())
            continue
        }
        log.V(1).Info("add secret")
        account.ImagePullSecrets = append(account.ImagePullSecrets,
            corev1.LocalObjectReference{Name: secret.Name})
    }
    err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
        log.V(1).Info("update sa", "resource", nn.String())
        err := r.Update(ctx, &account)
        if err != nil {
            log.Error(err, "update failed", "retry", "true")
        }
        return err
    })
    return err
}

func mapAppend(target map[string]string, source map[string]string) map[string]string {
    m := target
    if m == nil {
        m = make(map[string]string)
    }
    for key, value := range source {
        m[key] = value
    }
    return m
}

// SetupWithManager sets up the controller with the Manager.
func (r *ImagePullSecretInjectorReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).
        For(&impsi.ImagePullSecretInjector{}).
        Watches(
            &source.Kind{Type: &corev1.ServiceAccount{}},
            handler.EnqueueRequestsFromMapFunc(func(object client.Object) []reconcile.Request {
                return r.impsiReferencingServiceAccount(object)
            }),
        ).
        WithOptions(controller.Options{
            MaxConcurrentReconciles: 10,
        }).
        Complete(r)
}

func (r *ImagePullSecretInjectorReconciler) impsiReferencingServiceAccount(obj client.Object) []ctrl.Request {
    correlationID := uuid.New().String()
    ctx := context.WithValue(context.Background(), logx.CorrelationID, correlationID)
    log := logx.WithName(ctx, "impsiReferencingServiceAccount").
        WithValues("name", obj.GetName(), "namespace", obj.GetNamespace(), "kind", obj.GetObjectKind())
    sa, ok := obj.(*corev1.ServiceAccount)
    if !ok {
        log.Info("object is not a service account")
    }
    impsiList := &impsi.ImagePullSecretInjectorList{}

    if err := r.Client.List(ctx, impsiList); err != nil {
        log.Error(err, "fail to list")
        return []ctrl.Request{}
    }
    var res []ctrl.Request
    for _, ii := range impsiList.Items {
        if len(ii.Spec.Match(ctx, r.Client, sa).Items) > 0 {
            res = append(res, ctrl.Request{
                NamespacedName: types.NamespacedName{
                    Name:      ii.Name,
                    Namespace: ii.Namespace,
                },
            })
        }
    }
    return res
}
