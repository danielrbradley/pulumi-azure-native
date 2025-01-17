// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160330

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Network security rule
//
// Deprecated: Version v20160330 will be removed in the next major version of the provider. Upgrade to version v20180501 or later.
func LookupSecurityRule(ctx *pulumi.Context, args *LookupSecurityRuleArgs, opts ...pulumi.InvokeOption) (*LookupSecurityRuleResult, error) {
	var rv LookupSecurityRuleResult
	err := ctx.Invoke("azure-native:network/v20160330:getSecurityRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityRuleArgs struct {
	// The name of the network security group.
	NetworkSecurityGroupName string `pulumi:"networkSecurityGroupName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the security rule.
	SecurityRuleName string `pulumi:"securityRuleName"`
}

// Network security rule
type LookupSecurityRuleResult struct {
	// Gets or sets network traffic is allowed or denied. Possible values are 'Allow' and 'Deny'
	Access string `pulumi:"access"`
	// Gets or sets a description for this rule. Restricted to 140 chars.
	Description *string `pulumi:"description"`
	// Gets or sets destination address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	DestinationAddressPrefix string `pulumi:"destinationAddressPrefix"`
	// Gets or sets Destination Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	DestinationPortRange *string `pulumi:"destinationPortRange"`
	// Gets or sets the direction of the rule.InBound or Outbound. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
	Direction string `pulumi:"direction"`
	// A unique read-only string that changes whenever the resource is updated
	Etag *string `pulumi:"etag"`
	// Resource Id
	Id *string `pulumi:"id"`
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
	Name *string `pulumi:"name"`
	// Gets or sets the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority *int `pulumi:"priority"`
	// Gets or sets Network protocol this rule applies to. Can be Tcp, Udp or All(*).
	Protocol string `pulumi:"protocol"`
	// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState *string `pulumi:"provisioningState"`
	// Gets or sets source address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
	SourceAddressPrefix string `pulumi:"sourceAddressPrefix"`
	// Gets or sets Source Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	SourcePortRange *string `pulumi:"sourcePortRange"`
}

func LookupSecurityRuleOutput(ctx *pulumi.Context, args LookupSecurityRuleOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityRuleResult, error) {
			args := v.(LookupSecurityRuleArgs)
			r, err := LookupSecurityRule(ctx, &args, opts...)
			var s LookupSecurityRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityRuleResultOutput)
}

type LookupSecurityRuleOutputArgs struct {
	// The name of the network security group.
	NetworkSecurityGroupName pulumi.StringInput `pulumi:"networkSecurityGroupName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the security rule.
	SecurityRuleName pulumi.StringInput `pulumi:"securityRuleName"`
}

func (LookupSecurityRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityRuleArgs)(nil)).Elem()
}

// Network security rule
type LookupSecurityRuleResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityRuleResult)(nil)).Elem()
}

func (o LookupSecurityRuleResultOutput) ToLookupSecurityRuleResultOutput() LookupSecurityRuleResultOutput {
	return o
}

func (o LookupSecurityRuleResultOutput) ToLookupSecurityRuleResultOutputWithContext(ctx context.Context) LookupSecurityRuleResultOutput {
	return o
}

// Gets or sets network traffic is allowed or denied. Possible values are 'Allow' and 'Deny'
func (o LookupSecurityRuleResultOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) string { return v.Access }).(pulumi.StringOutput)
}

// Gets or sets a description for this rule. Restricted to 140 chars.
func (o LookupSecurityRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Gets or sets destination address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
func (o LookupSecurityRuleResultOutput) DestinationAddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) string { return v.DestinationAddressPrefix }).(pulumi.StringOutput)
}

// Gets or sets Destination Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
func (o LookupSecurityRuleResultOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
}

// Gets or sets the direction of the rule.InBound or Outbound. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
func (o LookupSecurityRuleResultOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) string { return v.Direction }).(pulumi.StringOutput)
}

// A unique read-only string that changes whenever the resource is updated
func (o LookupSecurityRuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupSecurityRuleResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
func (o LookupSecurityRuleResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Gets or sets the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
func (o LookupSecurityRuleResultOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

// Gets or sets Network protocol this rule applies to. Can be Tcp, Udp or All(*).
func (o LookupSecurityRuleResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) string { return v.Protocol }).(pulumi.StringOutput)
}

// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
func (o LookupSecurityRuleResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Gets or sets source address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
func (o LookupSecurityRuleResultOutput) SourceAddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) string { return v.SourceAddressPrefix }).(pulumi.StringOutput)
}

// Gets or sets Source Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
func (o LookupSecurityRuleResultOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.SourcePortRange }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityRuleResultOutput{})
}
