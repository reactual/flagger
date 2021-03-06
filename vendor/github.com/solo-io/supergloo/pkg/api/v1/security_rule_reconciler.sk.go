// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionSecurityRuleFunc func(original, desired *SecurityRule) (bool, error)

type SecurityRuleReconciler interface {
	Reconcile(namespace string, desiredResources SecurityRuleList, transition TransitionSecurityRuleFunc, opts clients.ListOpts) error
}

func securityRulesToResources(list SecurityRuleList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, securityRule := range list {
		resourceList = append(resourceList, securityRule)
	}
	return resourceList
}

func NewSecurityRuleReconciler(client SecurityRuleClient) SecurityRuleReconciler {
	return &securityRuleReconciler{
		base: reconcile.NewReconciler(client.BaseClient()),
	}
}

type securityRuleReconciler struct {
	base reconcile.Reconciler
}

func (r *securityRuleReconciler) Reconcile(namespace string, desiredResources SecurityRuleList, transition TransitionSecurityRuleFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "securityRule_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*SecurityRule), desired.(*SecurityRule))
		}
	}
	return r.base.Reconcile(namespace, securityRulesToResources(desiredResources), transitionResources, opts)
}
