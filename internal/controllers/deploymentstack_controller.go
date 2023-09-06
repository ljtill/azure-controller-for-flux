/*
MIT License

Copyright (c) 2023 Lyon Till

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	azurev1alpha1 "github.com/ljtill/azure-controller-for-flux/api/v1alpha1"
	"github.com/ljtill/azure-controller-for-flux/internal/predicates"
)

// DeploymentStackReconciler reconciles a DeploymentStack object
type DeploymentStackReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=azure.contrib.fluxcd.io,resources=deploymentstacks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=azure.contrib.fluxcd.io,resources=deploymentstacks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=azure.contrib.fluxcd.io,resources=deploymentstacks/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DeploymentStack object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.15.0/pkg/reconcile
func (r *DeploymentStackReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	log := ctrl.LoggerFrom(ctx)
	log.Info("Reconciling DeploymentStack")

	var dplStack azurev1alpha1.DeploymentStack

	// NOTE: Get Azure object
	err := r.Get(ctx, req.NamespacedName, &dplStack)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
// Watches for Source Controller and DeploymentStack events
// TODO: Add OCI Repository support
func (r *DeploymentStackReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.DeploymentStack{}, builder.WithPredicates(predicates.GitRepositoryRevisionChangePredicate{})).
		Complete(r)
}
