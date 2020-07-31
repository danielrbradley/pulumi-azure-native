// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupViewByScope(ctx *pulumi.Context, args *LookupViewByScopeArgs, opts ...pulumi.InvokeOption) (*LookupViewByScopeResult, error) {
	var rv LookupViewByScopeResult
	err := ctx.Invoke("azurerm:costmanagement/v20200601:getViewByScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupViewByScopeArgs struct {
	// View name
	Name string `pulumi:"name"`
	// The scope associated with view operations. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, 'providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for External Billing Account scope and 'providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for External Subscription scope.
	Scope string `pulumi:"scope"`
}

// States and configurations of Cost Analysis.
type LookupViewByScopeResult struct {
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag *string `pulumi:"eTag"`
	// Resource name.
	Name string `pulumi:"name"`
	// The properties of the view.
	Properties ViewPropertiesResponse `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}