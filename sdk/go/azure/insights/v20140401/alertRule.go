// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20140401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The alert rule resource.
//
// Deprecated: Version v20140401 will be removed in the next major version of the provider. Upgrade to version v20150401 or later.
type AlertRule struct {
	pulumi.CustomResourceState

	// action that is performed when the alert rule becomes active, and when an alert condition is resolved.
	Action pulumi.AnyOutput `pulumi:"action"`
	// the array of actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
	Actions pulumi.ArrayOutput `pulumi:"actions"`
	// the condition that results in the alert rule being activated.
	Condition pulumi.AnyOutput `pulumi:"condition"`
	// the description of the alert rule that will be included in the alert email.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// the flag that indicates whether the alert rule is enabled.
	IsEnabled pulumi.BoolOutput `pulumi:"isEnabled"`
	// Last time the rule was updated in ISO8601 format.
	LastUpdatedTime pulumi.StringOutput `pulumi:"lastUpdatedTime"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// the provisioning state.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAlertRule registers a new resource with the given unique name, arguments, and options.
func NewAlertRule(ctx *pulumi.Context,
	name string, args *AlertRuleArgs, opts ...pulumi.ResourceOption) (*AlertRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Condition == nil {
		return nil, errors.New("invalid value for required argument 'Condition'")
	}
	if args.IsEnabled == nil {
		return nil, errors.New("invalid value for required argument 'IsEnabled'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20160301:AlertRule"),
		},
	})
	opts = append(opts, aliases)
	var resource AlertRule
	err := ctx.RegisterResource("azure-native:insights/v20140401:AlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertRule gets an existing AlertRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertRuleState, opts ...pulumi.ResourceOption) (*AlertRule, error) {
	var resource AlertRule
	err := ctx.ReadResource("azure-native:insights/v20140401:AlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertRule resources.
type alertRuleState struct {
}

type AlertRuleState struct {
}

func (AlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertRuleState)(nil)).Elem()
}

type alertRuleArgs struct {
	// action that is performed when the alert rule becomes active, and when an alert condition is resolved.
	Action interface{} `pulumi:"action"`
	// the array of actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
	Actions []interface{} `pulumi:"actions"`
	// the condition that results in the alert rule being activated.
	Condition interface{} `pulumi:"condition"`
	// the description of the alert rule that will be included in the alert email.
	Description *string `pulumi:"description"`
	// the flag that indicates whether the alert rule is enabled.
	IsEnabled bool `pulumi:"isEnabled"`
	// Resource location
	Location *string `pulumi:"location"`
	// the name of the alert rule.
	Name string `pulumi:"name"`
	// the provisioning state.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the rule.
	RuleName *string `pulumi:"ruleName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AlertRule resource.
type AlertRuleArgs struct {
	// action that is performed when the alert rule becomes active, and when an alert condition is resolved.
	Action pulumi.Input
	// the array of actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
	Actions pulumi.ArrayInput
	// the condition that results in the alert rule being activated.
	Condition pulumi.Input
	// the description of the alert rule that will be included in the alert email.
	Description pulumi.StringPtrInput
	// the flag that indicates whether the alert rule is enabled.
	IsEnabled pulumi.BoolInput
	// Resource location
	Location pulumi.StringPtrInput
	// the name of the alert rule.
	Name pulumi.StringInput
	// the provisioning state.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the rule.
	RuleName pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (AlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertRuleArgs)(nil)).Elem()
}

type AlertRuleInput interface {
	pulumi.Input

	ToAlertRuleOutput() AlertRuleOutput
	ToAlertRuleOutputWithContext(ctx context.Context) AlertRuleOutput
}

func (*AlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertRule)(nil)).Elem()
}

func (i *AlertRule) ToAlertRuleOutput() AlertRuleOutput {
	return i.ToAlertRuleOutputWithContext(context.Background())
}

func (i *AlertRule) ToAlertRuleOutputWithContext(ctx context.Context) AlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleOutput)
}

type AlertRuleOutput struct{ *pulumi.OutputState }

func (AlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertRule)(nil)).Elem()
}

func (o AlertRuleOutput) ToAlertRuleOutput() AlertRuleOutput {
	return o
}

func (o AlertRuleOutput) ToAlertRuleOutputWithContext(ctx context.Context) AlertRuleOutput {
	return o
}

// action that is performed when the alert rule becomes active, and when an alert condition is resolved.
func (o AlertRuleOutput) Action() pulumi.AnyOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.AnyOutput { return v.Action }).(pulumi.AnyOutput)
}

// the array of actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
func (o AlertRuleOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.ArrayOutput { return v.Actions }).(pulumi.ArrayOutput)
}

// the condition that results in the alert rule being activated.
func (o AlertRuleOutput) Condition() pulumi.AnyOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.AnyOutput { return v.Condition }).(pulumi.AnyOutput)
}

// the description of the alert rule that will be included in the alert email.
func (o AlertRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// the flag that indicates whether the alert rule is enabled.
func (o AlertRuleOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.BoolOutput { return v.IsEnabled }).(pulumi.BoolOutput)
}

// Last time the rule was updated in ISO8601 format.
func (o AlertRuleOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.StringOutput { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

// Resource location
func (o AlertRuleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name
func (o AlertRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// the provisioning state.
func (o AlertRuleOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Resource tags
func (o AlertRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type
func (o AlertRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AlertRuleOutput{})
}
