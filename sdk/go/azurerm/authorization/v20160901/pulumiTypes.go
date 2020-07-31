// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The lock information.
type ManagementLockAtResourceGroupLevelType struct {
	// The name of the lock.
	Name string `pulumi:"name"`
	// The properties of the lock.
	Properties ManagementLockPropertiesResponse `pulumi:"properties"`
	// The resource type of the lock - Microsoft.Authorization/locks.
	Type string `pulumi:"type"`
}

// ManagementLockAtResourceGroupLevelTypeInput is an input type that accepts ManagementLockAtResourceGroupLevelTypeArgs and ManagementLockAtResourceGroupLevelTypeOutput values.
// You can construct a concrete instance of `ManagementLockAtResourceGroupLevelTypeInput` via:
//
//          ManagementLockAtResourceGroupLevelTypeArgs{...}
type ManagementLockAtResourceGroupLevelTypeInput interface {
	pulumi.Input

	ToManagementLockAtResourceGroupLevelTypeOutput() ManagementLockAtResourceGroupLevelTypeOutput
	ToManagementLockAtResourceGroupLevelTypeOutputWithContext(context.Context) ManagementLockAtResourceGroupLevelTypeOutput
}

// The lock information.
type ManagementLockAtResourceGroupLevelTypeArgs struct {
	// The name of the lock.
	Name pulumi.StringInput `pulumi:"name"`
	// The properties of the lock.
	Properties ManagementLockPropertiesResponseInput `pulumi:"properties"`
	// The resource type of the lock - Microsoft.Authorization/locks.
	Type pulumi.StringInput `pulumi:"type"`
}

func (ManagementLockAtResourceGroupLevelTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockAtResourceGroupLevelType)(nil)).Elem()
}

func (i ManagementLockAtResourceGroupLevelTypeArgs) ToManagementLockAtResourceGroupLevelTypeOutput() ManagementLockAtResourceGroupLevelTypeOutput {
	return i.ToManagementLockAtResourceGroupLevelTypeOutputWithContext(context.Background())
}

func (i ManagementLockAtResourceGroupLevelTypeArgs) ToManagementLockAtResourceGroupLevelTypeOutputWithContext(ctx context.Context) ManagementLockAtResourceGroupLevelTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockAtResourceGroupLevelTypeOutput)
}

// The lock information.
type ManagementLockAtResourceGroupLevelTypeOutput struct{ *pulumi.OutputState }

func (ManagementLockAtResourceGroupLevelTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockAtResourceGroupLevelType)(nil)).Elem()
}

func (o ManagementLockAtResourceGroupLevelTypeOutput) ToManagementLockAtResourceGroupLevelTypeOutput() ManagementLockAtResourceGroupLevelTypeOutput {
	return o
}

func (o ManagementLockAtResourceGroupLevelTypeOutput) ToManagementLockAtResourceGroupLevelTypeOutputWithContext(ctx context.Context) ManagementLockAtResourceGroupLevelTypeOutput {
	return o
}

// The name of the lock.
func (o ManagementLockAtResourceGroupLevelTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementLockAtResourceGroupLevelType) string { return v.Name }).(pulumi.StringOutput)
}

// The properties of the lock.
func (o ManagementLockAtResourceGroupLevelTypeOutput) Properties() ManagementLockPropertiesResponseOutput {
	return o.ApplyT(func(v ManagementLockAtResourceGroupLevelType) ManagementLockPropertiesResponse { return v.Properties }).(ManagementLockPropertiesResponseOutput)
}

// The resource type of the lock - Microsoft.Authorization/locks.
func (o ManagementLockAtResourceGroupLevelTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementLockAtResourceGroupLevelType) string { return v.Type }).(pulumi.StringOutput)
}

// The lock information.
type ManagementLockAtResourceLevelType struct {
	// The name of the lock.
	Name string `pulumi:"name"`
	// The properties of the lock.
	Properties ManagementLockPropertiesResponse `pulumi:"properties"`
	// The resource type of the lock - Microsoft.Authorization/locks.
	Type string `pulumi:"type"`
}

// ManagementLockAtResourceLevelTypeInput is an input type that accepts ManagementLockAtResourceLevelTypeArgs and ManagementLockAtResourceLevelTypeOutput values.
// You can construct a concrete instance of `ManagementLockAtResourceLevelTypeInput` via:
//
//          ManagementLockAtResourceLevelTypeArgs{...}
type ManagementLockAtResourceLevelTypeInput interface {
	pulumi.Input

	ToManagementLockAtResourceLevelTypeOutput() ManagementLockAtResourceLevelTypeOutput
	ToManagementLockAtResourceLevelTypeOutputWithContext(context.Context) ManagementLockAtResourceLevelTypeOutput
}

// The lock information.
type ManagementLockAtResourceLevelTypeArgs struct {
	// The name of the lock.
	Name pulumi.StringInput `pulumi:"name"`
	// The properties of the lock.
	Properties ManagementLockPropertiesResponseInput `pulumi:"properties"`
	// The resource type of the lock - Microsoft.Authorization/locks.
	Type pulumi.StringInput `pulumi:"type"`
}

func (ManagementLockAtResourceLevelTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockAtResourceLevelType)(nil)).Elem()
}

func (i ManagementLockAtResourceLevelTypeArgs) ToManagementLockAtResourceLevelTypeOutput() ManagementLockAtResourceLevelTypeOutput {
	return i.ToManagementLockAtResourceLevelTypeOutputWithContext(context.Background())
}

func (i ManagementLockAtResourceLevelTypeArgs) ToManagementLockAtResourceLevelTypeOutputWithContext(ctx context.Context) ManagementLockAtResourceLevelTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockAtResourceLevelTypeOutput)
}

// The lock information.
type ManagementLockAtResourceLevelTypeOutput struct{ *pulumi.OutputState }

func (ManagementLockAtResourceLevelTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockAtResourceLevelType)(nil)).Elem()
}

func (o ManagementLockAtResourceLevelTypeOutput) ToManagementLockAtResourceLevelTypeOutput() ManagementLockAtResourceLevelTypeOutput {
	return o
}

func (o ManagementLockAtResourceLevelTypeOutput) ToManagementLockAtResourceLevelTypeOutputWithContext(ctx context.Context) ManagementLockAtResourceLevelTypeOutput {
	return o
}

// The name of the lock.
func (o ManagementLockAtResourceLevelTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementLockAtResourceLevelType) string { return v.Name }).(pulumi.StringOutput)
}

// The properties of the lock.
func (o ManagementLockAtResourceLevelTypeOutput) Properties() ManagementLockPropertiesResponseOutput {
	return o.ApplyT(func(v ManagementLockAtResourceLevelType) ManagementLockPropertiesResponse { return v.Properties }).(ManagementLockPropertiesResponseOutput)
}

// The resource type of the lock - Microsoft.Authorization/locks.
func (o ManagementLockAtResourceLevelTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementLockAtResourceLevelType) string { return v.Type }).(pulumi.StringOutput)
}

// The lock information.
type ManagementLockAtSubscriptionLevelType struct {
	// The name of the lock.
	Name string `pulumi:"name"`
	// The properties of the lock.
	Properties ManagementLockPropertiesResponse `pulumi:"properties"`
	// The resource type of the lock - Microsoft.Authorization/locks.
	Type string `pulumi:"type"`
}

// ManagementLockAtSubscriptionLevelTypeInput is an input type that accepts ManagementLockAtSubscriptionLevelTypeArgs and ManagementLockAtSubscriptionLevelTypeOutput values.
// You can construct a concrete instance of `ManagementLockAtSubscriptionLevelTypeInput` via:
//
//          ManagementLockAtSubscriptionLevelTypeArgs{...}
type ManagementLockAtSubscriptionLevelTypeInput interface {
	pulumi.Input

	ToManagementLockAtSubscriptionLevelTypeOutput() ManagementLockAtSubscriptionLevelTypeOutput
	ToManagementLockAtSubscriptionLevelTypeOutputWithContext(context.Context) ManagementLockAtSubscriptionLevelTypeOutput
}

// The lock information.
type ManagementLockAtSubscriptionLevelTypeArgs struct {
	// The name of the lock.
	Name pulumi.StringInput `pulumi:"name"`
	// The properties of the lock.
	Properties ManagementLockPropertiesResponseInput `pulumi:"properties"`
	// The resource type of the lock - Microsoft.Authorization/locks.
	Type pulumi.StringInput `pulumi:"type"`
}

func (ManagementLockAtSubscriptionLevelTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockAtSubscriptionLevelType)(nil)).Elem()
}

func (i ManagementLockAtSubscriptionLevelTypeArgs) ToManagementLockAtSubscriptionLevelTypeOutput() ManagementLockAtSubscriptionLevelTypeOutput {
	return i.ToManagementLockAtSubscriptionLevelTypeOutputWithContext(context.Background())
}

func (i ManagementLockAtSubscriptionLevelTypeArgs) ToManagementLockAtSubscriptionLevelTypeOutputWithContext(ctx context.Context) ManagementLockAtSubscriptionLevelTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockAtSubscriptionLevelTypeOutput)
}

// The lock information.
type ManagementLockAtSubscriptionLevelTypeOutput struct{ *pulumi.OutputState }

func (ManagementLockAtSubscriptionLevelTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockAtSubscriptionLevelType)(nil)).Elem()
}

func (o ManagementLockAtSubscriptionLevelTypeOutput) ToManagementLockAtSubscriptionLevelTypeOutput() ManagementLockAtSubscriptionLevelTypeOutput {
	return o
}

func (o ManagementLockAtSubscriptionLevelTypeOutput) ToManagementLockAtSubscriptionLevelTypeOutputWithContext(ctx context.Context) ManagementLockAtSubscriptionLevelTypeOutput {
	return o
}

// The name of the lock.
func (o ManagementLockAtSubscriptionLevelTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementLockAtSubscriptionLevelType) string { return v.Name }).(pulumi.StringOutput)
}

// The properties of the lock.
func (o ManagementLockAtSubscriptionLevelTypeOutput) Properties() ManagementLockPropertiesResponseOutput {
	return o.ApplyT(func(v ManagementLockAtSubscriptionLevelType) ManagementLockPropertiesResponse { return v.Properties }).(ManagementLockPropertiesResponseOutput)
}

// The resource type of the lock - Microsoft.Authorization/locks.
func (o ManagementLockAtSubscriptionLevelTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementLockAtSubscriptionLevelType) string { return v.Type }).(pulumi.StringOutput)
}

// The lock information.
type ManagementLockByScopeType struct {
	// The name of the lock.
	Name string `pulumi:"name"`
	// The properties of the lock.
	Properties ManagementLockPropertiesResponse `pulumi:"properties"`
	// The resource type of the lock - Microsoft.Authorization/locks.
	Type string `pulumi:"type"`
}

// ManagementLockByScopeTypeInput is an input type that accepts ManagementLockByScopeTypeArgs and ManagementLockByScopeTypeOutput values.
// You can construct a concrete instance of `ManagementLockByScopeTypeInput` via:
//
//          ManagementLockByScopeTypeArgs{...}
type ManagementLockByScopeTypeInput interface {
	pulumi.Input

	ToManagementLockByScopeTypeOutput() ManagementLockByScopeTypeOutput
	ToManagementLockByScopeTypeOutputWithContext(context.Context) ManagementLockByScopeTypeOutput
}

// The lock information.
type ManagementLockByScopeTypeArgs struct {
	// The name of the lock.
	Name pulumi.StringInput `pulumi:"name"`
	// The properties of the lock.
	Properties ManagementLockPropertiesResponseInput `pulumi:"properties"`
	// The resource type of the lock - Microsoft.Authorization/locks.
	Type pulumi.StringInput `pulumi:"type"`
}

func (ManagementLockByScopeTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockByScopeType)(nil)).Elem()
}

func (i ManagementLockByScopeTypeArgs) ToManagementLockByScopeTypeOutput() ManagementLockByScopeTypeOutput {
	return i.ToManagementLockByScopeTypeOutputWithContext(context.Background())
}

func (i ManagementLockByScopeTypeArgs) ToManagementLockByScopeTypeOutputWithContext(ctx context.Context) ManagementLockByScopeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockByScopeTypeOutput)
}

// The lock information.
type ManagementLockByScopeTypeOutput struct{ *pulumi.OutputState }

func (ManagementLockByScopeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockByScopeType)(nil)).Elem()
}

func (o ManagementLockByScopeTypeOutput) ToManagementLockByScopeTypeOutput() ManagementLockByScopeTypeOutput {
	return o
}

func (o ManagementLockByScopeTypeOutput) ToManagementLockByScopeTypeOutputWithContext(ctx context.Context) ManagementLockByScopeTypeOutput {
	return o
}

// The name of the lock.
func (o ManagementLockByScopeTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementLockByScopeType) string { return v.Name }).(pulumi.StringOutput)
}

// The properties of the lock.
func (o ManagementLockByScopeTypeOutput) Properties() ManagementLockPropertiesResponseOutput {
	return o.ApplyT(func(v ManagementLockByScopeType) ManagementLockPropertiesResponse { return v.Properties }).(ManagementLockPropertiesResponseOutput)
}

// The resource type of the lock - Microsoft.Authorization/locks.
func (o ManagementLockByScopeTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementLockByScopeType) string { return v.Type }).(pulumi.StringOutput)
}

// Lock owner properties.
type ManagementLockOwner struct {
	// The application ID of the lock owner.
	ApplicationId *string `pulumi:"applicationId"`
}

// ManagementLockOwnerInput is an input type that accepts ManagementLockOwnerArgs and ManagementLockOwnerOutput values.
// You can construct a concrete instance of `ManagementLockOwnerInput` via:
//
//          ManagementLockOwnerArgs{...}
type ManagementLockOwnerInput interface {
	pulumi.Input

	ToManagementLockOwnerOutput() ManagementLockOwnerOutput
	ToManagementLockOwnerOutputWithContext(context.Context) ManagementLockOwnerOutput
}

// Lock owner properties.
type ManagementLockOwnerArgs struct {
	// The application ID of the lock owner.
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
}

func (ManagementLockOwnerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockOwner)(nil)).Elem()
}

func (i ManagementLockOwnerArgs) ToManagementLockOwnerOutput() ManagementLockOwnerOutput {
	return i.ToManagementLockOwnerOutputWithContext(context.Background())
}

func (i ManagementLockOwnerArgs) ToManagementLockOwnerOutputWithContext(ctx context.Context) ManagementLockOwnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockOwnerOutput)
}

// ManagementLockOwnerArrayInput is an input type that accepts ManagementLockOwnerArray and ManagementLockOwnerArrayOutput values.
// You can construct a concrete instance of `ManagementLockOwnerArrayInput` via:
//
//          ManagementLockOwnerArray{ ManagementLockOwnerArgs{...} }
type ManagementLockOwnerArrayInput interface {
	pulumi.Input

	ToManagementLockOwnerArrayOutput() ManagementLockOwnerArrayOutput
	ToManagementLockOwnerArrayOutputWithContext(context.Context) ManagementLockOwnerArrayOutput
}

type ManagementLockOwnerArray []ManagementLockOwnerInput

func (ManagementLockOwnerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementLockOwner)(nil)).Elem()
}

func (i ManagementLockOwnerArray) ToManagementLockOwnerArrayOutput() ManagementLockOwnerArrayOutput {
	return i.ToManagementLockOwnerArrayOutputWithContext(context.Background())
}

func (i ManagementLockOwnerArray) ToManagementLockOwnerArrayOutputWithContext(ctx context.Context) ManagementLockOwnerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockOwnerArrayOutput)
}

// Lock owner properties.
type ManagementLockOwnerOutput struct{ *pulumi.OutputState }

func (ManagementLockOwnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockOwner)(nil)).Elem()
}

func (o ManagementLockOwnerOutput) ToManagementLockOwnerOutput() ManagementLockOwnerOutput {
	return o
}

func (o ManagementLockOwnerOutput) ToManagementLockOwnerOutputWithContext(ctx context.Context) ManagementLockOwnerOutput {
	return o
}

// The application ID of the lock owner.
func (o ManagementLockOwnerOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementLockOwner) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

type ManagementLockOwnerArrayOutput struct{ *pulumi.OutputState }

func (ManagementLockOwnerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementLockOwner)(nil)).Elem()
}

func (o ManagementLockOwnerArrayOutput) ToManagementLockOwnerArrayOutput() ManagementLockOwnerArrayOutput {
	return o
}

func (o ManagementLockOwnerArrayOutput) ToManagementLockOwnerArrayOutputWithContext(ctx context.Context) ManagementLockOwnerArrayOutput {
	return o
}

func (o ManagementLockOwnerArrayOutput) Index(i pulumi.IntInput) ManagementLockOwnerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementLockOwner {
		return vs[0].([]ManagementLockOwner)[vs[1].(int)]
	}).(ManagementLockOwnerOutput)
}

// Lock owner properties.
type ManagementLockOwnerResponse struct {
	// The application ID of the lock owner.
	ApplicationId *string `pulumi:"applicationId"`
}

// ManagementLockOwnerResponseInput is an input type that accepts ManagementLockOwnerResponseArgs and ManagementLockOwnerResponseOutput values.
// You can construct a concrete instance of `ManagementLockOwnerResponseInput` via:
//
//          ManagementLockOwnerResponseArgs{...}
type ManagementLockOwnerResponseInput interface {
	pulumi.Input

	ToManagementLockOwnerResponseOutput() ManagementLockOwnerResponseOutput
	ToManagementLockOwnerResponseOutputWithContext(context.Context) ManagementLockOwnerResponseOutput
}

// Lock owner properties.
type ManagementLockOwnerResponseArgs struct {
	// The application ID of the lock owner.
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
}

func (ManagementLockOwnerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockOwnerResponse)(nil)).Elem()
}

func (i ManagementLockOwnerResponseArgs) ToManagementLockOwnerResponseOutput() ManagementLockOwnerResponseOutput {
	return i.ToManagementLockOwnerResponseOutputWithContext(context.Background())
}

func (i ManagementLockOwnerResponseArgs) ToManagementLockOwnerResponseOutputWithContext(ctx context.Context) ManagementLockOwnerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockOwnerResponseOutput)
}

// ManagementLockOwnerResponseArrayInput is an input type that accepts ManagementLockOwnerResponseArray and ManagementLockOwnerResponseArrayOutput values.
// You can construct a concrete instance of `ManagementLockOwnerResponseArrayInput` via:
//
//          ManagementLockOwnerResponseArray{ ManagementLockOwnerResponseArgs{...} }
type ManagementLockOwnerResponseArrayInput interface {
	pulumi.Input

	ToManagementLockOwnerResponseArrayOutput() ManagementLockOwnerResponseArrayOutput
	ToManagementLockOwnerResponseArrayOutputWithContext(context.Context) ManagementLockOwnerResponseArrayOutput
}

type ManagementLockOwnerResponseArray []ManagementLockOwnerResponseInput

func (ManagementLockOwnerResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementLockOwnerResponse)(nil)).Elem()
}

func (i ManagementLockOwnerResponseArray) ToManagementLockOwnerResponseArrayOutput() ManagementLockOwnerResponseArrayOutput {
	return i.ToManagementLockOwnerResponseArrayOutputWithContext(context.Background())
}

func (i ManagementLockOwnerResponseArray) ToManagementLockOwnerResponseArrayOutputWithContext(ctx context.Context) ManagementLockOwnerResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockOwnerResponseArrayOutput)
}

// Lock owner properties.
type ManagementLockOwnerResponseOutput struct{ *pulumi.OutputState }

func (ManagementLockOwnerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockOwnerResponse)(nil)).Elem()
}

func (o ManagementLockOwnerResponseOutput) ToManagementLockOwnerResponseOutput() ManagementLockOwnerResponseOutput {
	return o
}

func (o ManagementLockOwnerResponseOutput) ToManagementLockOwnerResponseOutputWithContext(ctx context.Context) ManagementLockOwnerResponseOutput {
	return o
}

// The application ID of the lock owner.
func (o ManagementLockOwnerResponseOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementLockOwnerResponse) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

type ManagementLockOwnerResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagementLockOwnerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementLockOwnerResponse)(nil)).Elem()
}

func (o ManagementLockOwnerResponseArrayOutput) ToManagementLockOwnerResponseArrayOutput() ManagementLockOwnerResponseArrayOutput {
	return o
}

func (o ManagementLockOwnerResponseArrayOutput) ToManagementLockOwnerResponseArrayOutputWithContext(ctx context.Context) ManagementLockOwnerResponseArrayOutput {
	return o
}

func (o ManagementLockOwnerResponseArrayOutput) Index(i pulumi.IntInput) ManagementLockOwnerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementLockOwnerResponse {
		return vs[0].([]ManagementLockOwnerResponse)[vs[1].(int)]
	}).(ManagementLockOwnerResponseOutput)
}

// The lock properties.
type ManagementLockProperties struct {
	// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
	Level string `pulumi:"level"`
	// Notes about the lock. Maximum of 512 characters.
	Notes *string `pulumi:"notes"`
	// The owners of the lock.
	Owners []ManagementLockOwner `pulumi:"owners"`
}

// ManagementLockPropertiesInput is an input type that accepts ManagementLockPropertiesArgs and ManagementLockPropertiesOutput values.
// You can construct a concrete instance of `ManagementLockPropertiesInput` via:
//
//          ManagementLockPropertiesArgs{...}
type ManagementLockPropertiesInput interface {
	pulumi.Input

	ToManagementLockPropertiesOutput() ManagementLockPropertiesOutput
	ToManagementLockPropertiesOutputWithContext(context.Context) ManagementLockPropertiesOutput
}

// The lock properties.
type ManagementLockPropertiesArgs struct {
	// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
	Level pulumi.StringInput `pulumi:"level"`
	// Notes about the lock. Maximum of 512 characters.
	Notes pulumi.StringPtrInput `pulumi:"notes"`
	// The owners of the lock.
	Owners ManagementLockOwnerArrayInput `pulumi:"owners"`
}

func (ManagementLockPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockProperties)(nil)).Elem()
}

func (i ManagementLockPropertiesArgs) ToManagementLockPropertiesOutput() ManagementLockPropertiesOutput {
	return i.ToManagementLockPropertiesOutputWithContext(context.Background())
}

func (i ManagementLockPropertiesArgs) ToManagementLockPropertiesOutputWithContext(ctx context.Context) ManagementLockPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockPropertiesOutput)
}

func (i ManagementLockPropertiesArgs) ToManagementLockPropertiesPtrOutput() ManagementLockPropertiesPtrOutput {
	return i.ToManagementLockPropertiesPtrOutputWithContext(context.Background())
}

func (i ManagementLockPropertiesArgs) ToManagementLockPropertiesPtrOutputWithContext(ctx context.Context) ManagementLockPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockPropertiesOutput).ToManagementLockPropertiesPtrOutputWithContext(ctx)
}

// ManagementLockPropertiesPtrInput is an input type that accepts ManagementLockPropertiesArgs, ManagementLockPropertiesPtr and ManagementLockPropertiesPtrOutput values.
// You can construct a concrete instance of `ManagementLockPropertiesPtrInput` via:
//
//          ManagementLockPropertiesArgs{...}
//
//  or:
//
//          nil
type ManagementLockPropertiesPtrInput interface {
	pulumi.Input

	ToManagementLockPropertiesPtrOutput() ManagementLockPropertiesPtrOutput
	ToManagementLockPropertiesPtrOutputWithContext(context.Context) ManagementLockPropertiesPtrOutput
}

type managementLockPropertiesPtrType ManagementLockPropertiesArgs

func ManagementLockPropertiesPtr(v *ManagementLockPropertiesArgs) ManagementLockPropertiesPtrInput {
	return (*managementLockPropertiesPtrType)(v)
}

func (*managementLockPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementLockProperties)(nil)).Elem()
}

func (i *managementLockPropertiesPtrType) ToManagementLockPropertiesPtrOutput() ManagementLockPropertiesPtrOutput {
	return i.ToManagementLockPropertiesPtrOutputWithContext(context.Background())
}

func (i *managementLockPropertiesPtrType) ToManagementLockPropertiesPtrOutputWithContext(ctx context.Context) ManagementLockPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockPropertiesPtrOutput)
}

// The lock properties.
type ManagementLockPropertiesOutput struct{ *pulumi.OutputState }

func (ManagementLockPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockProperties)(nil)).Elem()
}

func (o ManagementLockPropertiesOutput) ToManagementLockPropertiesOutput() ManagementLockPropertiesOutput {
	return o
}

func (o ManagementLockPropertiesOutput) ToManagementLockPropertiesOutputWithContext(ctx context.Context) ManagementLockPropertiesOutput {
	return o
}

func (o ManagementLockPropertiesOutput) ToManagementLockPropertiesPtrOutput() ManagementLockPropertiesPtrOutput {
	return o.ToManagementLockPropertiesPtrOutputWithContext(context.Background())
}

func (o ManagementLockPropertiesOutput) ToManagementLockPropertiesPtrOutputWithContext(ctx context.Context) ManagementLockPropertiesPtrOutput {
	return o.ApplyT(func(v ManagementLockProperties) *ManagementLockProperties {
		return &v
	}).(ManagementLockPropertiesPtrOutput)
}

// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
func (o ManagementLockPropertiesOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementLockProperties) string { return v.Level }).(pulumi.StringOutput)
}

// Notes about the lock. Maximum of 512 characters.
func (o ManagementLockPropertiesOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementLockProperties) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

// The owners of the lock.
func (o ManagementLockPropertiesOutput) Owners() ManagementLockOwnerArrayOutput {
	return o.ApplyT(func(v ManagementLockProperties) []ManagementLockOwner { return v.Owners }).(ManagementLockOwnerArrayOutput)
}

type ManagementLockPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ManagementLockPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementLockProperties)(nil)).Elem()
}

func (o ManagementLockPropertiesPtrOutput) ToManagementLockPropertiesPtrOutput() ManagementLockPropertiesPtrOutput {
	return o
}

func (o ManagementLockPropertiesPtrOutput) ToManagementLockPropertiesPtrOutputWithContext(ctx context.Context) ManagementLockPropertiesPtrOutput {
	return o
}

func (o ManagementLockPropertiesPtrOutput) Elem() ManagementLockPropertiesOutput {
	return o.ApplyT(func(v *ManagementLockProperties) ManagementLockProperties { return *v }).(ManagementLockPropertiesOutput)
}

// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
func (o ManagementLockPropertiesPtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementLockProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Level
	}).(pulumi.StringPtrOutput)
}

// Notes about the lock. Maximum of 512 characters.
func (o ManagementLockPropertiesPtrOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementLockProperties) *string {
		if v == nil {
			return nil
		}
		return v.Notes
	}).(pulumi.StringPtrOutput)
}

// The owners of the lock.
func (o ManagementLockPropertiesPtrOutput) Owners() ManagementLockOwnerArrayOutput {
	return o.ApplyT(func(v *ManagementLockProperties) []ManagementLockOwner {
		if v == nil {
			return nil
		}
		return v.Owners
	}).(ManagementLockOwnerArrayOutput)
}

// The lock properties.
type ManagementLockPropertiesResponse struct {
	// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
	Level string `pulumi:"level"`
	// Notes about the lock. Maximum of 512 characters.
	Notes *string `pulumi:"notes"`
	// The owners of the lock.
	Owners []ManagementLockOwnerResponse `pulumi:"owners"`
}

// ManagementLockPropertiesResponseInput is an input type that accepts ManagementLockPropertiesResponseArgs and ManagementLockPropertiesResponseOutput values.
// You can construct a concrete instance of `ManagementLockPropertiesResponseInput` via:
//
//          ManagementLockPropertiesResponseArgs{...}
type ManagementLockPropertiesResponseInput interface {
	pulumi.Input

	ToManagementLockPropertiesResponseOutput() ManagementLockPropertiesResponseOutput
	ToManagementLockPropertiesResponseOutputWithContext(context.Context) ManagementLockPropertiesResponseOutput
}

// The lock properties.
type ManagementLockPropertiesResponseArgs struct {
	// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
	Level pulumi.StringInput `pulumi:"level"`
	// Notes about the lock. Maximum of 512 characters.
	Notes pulumi.StringPtrInput `pulumi:"notes"`
	// The owners of the lock.
	Owners ManagementLockOwnerResponseArrayInput `pulumi:"owners"`
}

func (ManagementLockPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockPropertiesResponse)(nil)).Elem()
}

func (i ManagementLockPropertiesResponseArgs) ToManagementLockPropertiesResponseOutput() ManagementLockPropertiesResponseOutput {
	return i.ToManagementLockPropertiesResponseOutputWithContext(context.Background())
}

func (i ManagementLockPropertiesResponseArgs) ToManagementLockPropertiesResponseOutputWithContext(ctx context.Context) ManagementLockPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockPropertiesResponseOutput)
}

func (i ManagementLockPropertiesResponseArgs) ToManagementLockPropertiesResponsePtrOutput() ManagementLockPropertiesResponsePtrOutput {
	return i.ToManagementLockPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ManagementLockPropertiesResponseArgs) ToManagementLockPropertiesResponsePtrOutputWithContext(ctx context.Context) ManagementLockPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockPropertiesResponseOutput).ToManagementLockPropertiesResponsePtrOutputWithContext(ctx)
}

// ManagementLockPropertiesResponsePtrInput is an input type that accepts ManagementLockPropertiesResponseArgs, ManagementLockPropertiesResponsePtr and ManagementLockPropertiesResponsePtrOutput values.
// You can construct a concrete instance of `ManagementLockPropertiesResponsePtrInput` via:
//
//          ManagementLockPropertiesResponseArgs{...}
//
//  or:
//
//          nil
type ManagementLockPropertiesResponsePtrInput interface {
	pulumi.Input

	ToManagementLockPropertiesResponsePtrOutput() ManagementLockPropertiesResponsePtrOutput
	ToManagementLockPropertiesResponsePtrOutputWithContext(context.Context) ManagementLockPropertiesResponsePtrOutput
}

type managementLockPropertiesResponsePtrType ManagementLockPropertiesResponseArgs

func ManagementLockPropertiesResponsePtr(v *ManagementLockPropertiesResponseArgs) ManagementLockPropertiesResponsePtrInput {
	return (*managementLockPropertiesResponsePtrType)(v)
}

func (*managementLockPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementLockPropertiesResponse)(nil)).Elem()
}

func (i *managementLockPropertiesResponsePtrType) ToManagementLockPropertiesResponsePtrOutput() ManagementLockPropertiesResponsePtrOutput {
	return i.ToManagementLockPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *managementLockPropertiesResponsePtrType) ToManagementLockPropertiesResponsePtrOutputWithContext(ctx context.Context) ManagementLockPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockPropertiesResponsePtrOutput)
}

// The lock properties.
type ManagementLockPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ManagementLockPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockPropertiesResponse)(nil)).Elem()
}

func (o ManagementLockPropertiesResponseOutput) ToManagementLockPropertiesResponseOutput() ManagementLockPropertiesResponseOutput {
	return o
}

func (o ManagementLockPropertiesResponseOutput) ToManagementLockPropertiesResponseOutputWithContext(ctx context.Context) ManagementLockPropertiesResponseOutput {
	return o
}

func (o ManagementLockPropertiesResponseOutput) ToManagementLockPropertiesResponsePtrOutput() ManagementLockPropertiesResponsePtrOutput {
	return o.ToManagementLockPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ManagementLockPropertiesResponseOutput) ToManagementLockPropertiesResponsePtrOutputWithContext(ctx context.Context) ManagementLockPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ManagementLockPropertiesResponse) *ManagementLockPropertiesResponse {
		return &v
	}).(ManagementLockPropertiesResponsePtrOutput)
}

// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
func (o ManagementLockPropertiesResponseOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementLockPropertiesResponse) string { return v.Level }).(pulumi.StringOutput)
}

// Notes about the lock. Maximum of 512 characters.
func (o ManagementLockPropertiesResponseOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementLockPropertiesResponse) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

// The owners of the lock.
func (o ManagementLockPropertiesResponseOutput) Owners() ManagementLockOwnerResponseArrayOutput {
	return o.ApplyT(func(v ManagementLockPropertiesResponse) []ManagementLockOwnerResponse { return v.Owners }).(ManagementLockOwnerResponseArrayOutput)
}

type ManagementLockPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagementLockPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementLockPropertiesResponse)(nil)).Elem()
}

func (o ManagementLockPropertiesResponsePtrOutput) ToManagementLockPropertiesResponsePtrOutput() ManagementLockPropertiesResponsePtrOutput {
	return o
}

func (o ManagementLockPropertiesResponsePtrOutput) ToManagementLockPropertiesResponsePtrOutputWithContext(ctx context.Context) ManagementLockPropertiesResponsePtrOutput {
	return o
}

func (o ManagementLockPropertiesResponsePtrOutput) Elem() ManagementLockPropertiesResponseOutput {
	return o.ApplyT(func(v *ManagementLockPropertiesResponse) ManagementLockPropertiesResponse { return *v }).(ManagementLockPropertiesResponseOutput)
}

// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
func (o ManagementLockPropertiesResponsePtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementLockPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Level
	}).(pulumi.StringPtrOutput)
}

// Notes about the lock. Maximum of 512 characters.
func (o ManagementLockPropertiesResponsePtrOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementLockPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Notes
	}).(pulumi.StringPtrOutput)
}

// The owners of the lock.
func (o ManagementLockPropertiesResponsePtrOutput) Owners() ManagementLockOwnerResponseArrayOutput {
	return o.ApplyT(func(v *ManagementLockPropertiesResponse) []ManagementLockOwnerResponse {
		if v == nil {
			return nil
		}
		return v.Owners
	}).(ManagementLockOwnerResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagementLockAtResourceGroupLevelTypeOutput{})
	pulumi.RegisterOutputType(ManagementLockAtResourceLevelTypeOutput{})
	pulumi.RegisterOutputType(ManagementLockAtSubscriptionLevelTypeOutput{})
	pulumi.RegisterOutputType(ManagementLockByScopeTypeOutput{})
	pulumi.RegisterOutputType(ManagementLockOwnerOutput{})
	pulumi.RegisterOutputType(ManagementLockOwnerArrayOutput{})
	pulumi.RegisterOutputType(ManagementLockOwnerResponseOutput{})
	pulumi.RegisterOutputType(ManagementLockOwnerResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagementLockPropertiesOutput{})
	pulumi.RegisterOutputType(ManagementLockPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ManagementLockPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ManagementLockPropertiesResponsePtrOutput{})
}