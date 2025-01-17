// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Route Filter Rule Resource
//
// Deprecated: Version v20170901 will be removed in the next major version of the provider. Upgrade to version v20180501 or later.
func LookupRouteFilterRule(ctx *pulumi.Context, args *LookupRouteFilterRuleArgs, opts ...pulumi.InvokeOption) (*LookupRouteFilterRuleResult, error) {
	var rv LookupRouteFilterRuleResult
	err := ctx.Invoke("azure-native:network/v20170901:getRouteFilterRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteFilterRuleArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the route filter.
	RouteFilterName string `pulumi:"routeFilterName"`
	// The name of the rule.
	RuleName string `pulumi:"ruleName"`
}

// Route Filter Rule Resource
type LookupRouteFilterRuleResult struct {
	// The access type of the rule. Valid values are: 'Allow', 'Deny'
	Access string `pulumi:"access"`
	// The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020']
	Communities []string `pulumi:"communities"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', 'Succeeded' and 'Failed'.
	ProvisioningState string `pulumi:"provisioningState"`
	// The rule type of the rule. Valid value is: 'Community'
	RouteFilterRuleType string `pulumi:"routeFilterRuleType"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

func LookupRouteFilterRuleOutput(ctx *pulumi.Context, args LookupRouteFilterRuleOutputArgs, opts ...pulumi.InvokeOption) LookupRouteFilterRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteFilterRuleResult, error) {
			args := v.(LookupRouteFilterRuleArgs)
			r, err := LookupRouteFilterRule(ctx, &args, opts...)
			var s LookupRouteFilterRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteFilterRuleResultOutput)
}

type LookupRouteFilterRuleOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the route filter.
	RouteFilterName pulumi.StringInput `pulumi:"routeFilterName"`
	// The name of the rule.
	RuleName pulumi.StringInput `pulumi:"ruleName"`
}

func (LookupRouteFilterRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteFilterRuleArgs)(nil)).Elem()
}

// Route Filter Rule Resource
type LookupRouteFilterRuleResultOutput struct{ *pulumi.OutputState }

func (LookupRouteFilterRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteFilterRuleResult)(nil)).Elem()
}

func (o LookupRouteFilterRuleResultOutput) ToLookupRouteFilterRuleResultOutput() LookupRouteFilterRuleResultOutput {
	return o
}

func (o LookupRouteFilterRuleResultOutput) ToLookupRouteFilterRuleResultOutputWithContext(ctx context.Context) LookupRouteFilterRuleResultOutput {
	return o
}

// The access type of the rule. Valid values are: 'Allow', 'Deny'
func (o LookupRouteFilterRuleResultOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterRuleResult) string { return v.Access }).(pulumi.StringOutput)
}

// The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020']
func (o LookupRouteFilterRuleResultOutput) Communities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRouteFilterRuleResult) []string { return v.Communities }).(pulumi.StringArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupRouteFilterRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupRouteFilterRuleResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteFilterRuleResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupRouteFilterRuleResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteFilterRuleResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupRouteFilterRuleResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteFilterRuleResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', 'Succeeded' and 'Failed'.
func (o LookupRouteFilterRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The rule type of the rule. Valid value is: 'Community'
func (o LookupRouteFilterRuleResultOutput) RouteFilterRuleType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterRuleResult) string { return v.RouteFilterRuleType }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupRouteFilterRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRouteFilterRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteFilterRuleResultOutput{})
}
