// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20151201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An application package which represents a particular version of an application.
type ApplicationPackageResponse struct {
	// The format of the application package, if the package is active.
	Format *string `pulumi:"format"`
	// The ID of the application.
	Id *string `pulumi:"id"`
	// The time at which the package was last activated, if the package is active.
	LastActivationTime *string `pulumi:"lastActivationTime"`
	// The current state of the application package.
	State *string `pulumi:"state"`
	// The storage URL at which the application package is stored.
	StorageUrl *string `pulumi:"storageUrl"`
	// The UTC time at which the storage URL will expire.
	StorageUrlExpiry *string `pulumi:"storageUrlExpiry"`
	// The version of the application package.
	Version *string `pulumi:"version"`
}

// An application package which represents a particular version of an application.
type ApplicationPackageResponseOutput struct{ *pulumi.OutputState }

func (ApplicationPackageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackageResponse)(nil)).Elem()
}

func (o ApplicationPackageResponseOutput) ToApplicationPackageResponseOutput() ApplicationPackageResponseOutput {
	return o
}

func (o ApplicationPackageResponseOutput) ToApplicationPackageResponseOutputWithContext(ctx context.Context) ApplicationPackageResponseOutput {
	return o
}

// The format of the application package, if the package is active.
func (o ApplicationPackageResponseOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) *string { return v.Format }).(pulumi.StringPtrOutput)
}

// The ID of the application.
func (o ApplicationPackageResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The time at which the package was last activated, if the package is active.
func (o ApplicationPackageResponseOutput) LastActivationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) *string { return v.LastActivationTime }).(pulumi.StringPtrOutput)
}

// The current state of the application package.
func (o ApplicationPackageResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

// The storage URL at which the application package is stored.
func (o ApplicationPackageResponseOutput) StorageUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) *string { return v.StorageUrl }).(pulumi.StringPtrOutput)
}

// The UTC time at which the storage URL will expire.
func (o ApplicationPackageResponseOutput) StorageUrlExpiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) *string { return v.StorageUrlExpiry }).(pulumi.StringPtrOutput)
}

// The version of the application package.
func (o ApplicationPackageResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ApplicationPackageResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationPackageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationPackageResponse)(nil)).Elem()
}

func (o ApplicationPackageResponseArrayOutput) ToApplicationPackageResponseArrayOutput() ApplicationPackageResponseArrayOutput {
	return o
}

func (o ApplicationPackageResponseArrayOutput) ToApplicationPackageResponseArrayOutputWithContext(ctx context.Context) ApplicationPackageResponseArrayOutput {
	return o
}

func (o ApplicationPackageResponseArrayOutput) Index(i pulumi.IntInput) ApplicationPackageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationPackageResponse {
		return vs[0].([]ApplicationPackageResponse)[vs[1].(int)]
	}).(ApplicationPackageResponseOutput)
}

// The properties related to auto storage account.
type AutoStorageBaseProperties struct {
	// The resource ID of the storage account to be used for auto storage account.
	StorageAccountId string `pulumi:"storageAccountId"`
}

// AutoStorageBasePropertiesInput is an input type that accepts AutoStorageBasePropertiesArgs and AutoStorageBasePropertiesOutput values.
// You can construct a concrete instance of `AutoStorageBasePropertiesInput` via:
//
//          AutoStorageBasePropertiesArgs{...}
type AutoStorageBasePropertiesInput interface {
	pulumi.Input

	ToAutoStorageBasePropertiesOutput() AutoStorageBasePropertiesOutput
	ToAutoStorageBasePropertiesOutputWithContext(context.Context) AutoStorageBasePropertiesOutput
}

// The properties related to auto storage account.
type AutoStorageBasePropertiesArgs struct {
	// The resource ID of the storage account to be used for auto storage account.
	StorageAccountId pulumi.StringInput `pulumi:"storageAccountId"`
}

func (AutoStorageBasePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoStorageBaseProperties)(nil)).Elem()
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesOutput() AutoStorageBasePropertiesOutput {
	return i.ToAutoStorageBasePropertiesOutputWithContext(context.Background())
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesOutputWithContext(ctx context.Context) AutoStorageBasePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoStorageBasePropertiesOutput)
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return i.ToAutoStorageBasePropertiesPtrOutputWithContext(context.Background())
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoStorageBasePropertiesOutput).ToAutoStorageBasePropertiesPtrOutputWithContext(ctx)
}

// AutoStorageBasePropertiesPtrInput is an input type that accepts AutoStorageBasePropertiesArgs, AutoStorageBasePropertiesPtr and AutoStorageBasePropertiesPtrOutput values.
// You can construct a concrete instance of `AutoStorageBasePropertiesPtrInput` via:
//
//          AutoStorageBasePropertiesArgs{...}
//
//  or:
//
//          nil
type AutoStorageBasePropertiesPtrInput interface {
	pulumi.Input

	ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput
	ToAutoStorageBasePropertiesPtrOutputWithContext(context.Context) AutoStorageBasePropertiesPtrOutput
}

type autoStorageBasePropertiesPtrType AutoStorageBasePropertiesArgs

func AutoStorageBasePropertiesPtr(v *AutoStorageBasePropertiesArgs) AutoStorageBasePropertiesPtrInput {
	return (*autoStorageBasePropertiesPtrType)(v)
}

func (*autoStorageBasePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoStorageBaseProperties)(nil)).Elem()
}

func (i *autoStorageBasePropertiesPtrType) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return i.ToAutoStorageBasePropertiesPtrOutputWithContext(context.Background())
}

func (i *autoStorageBasePropertiesPtrType) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoStorageBasePropertiesPtrOutput)
}

// The properties related to auto storage account.
type AutoStorageBasePropertiesOutput struct{ *pulumi.OutputState }

func (AutoStorageBasePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoStorageBaseProperties)(nil)).Elem()
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesOutput() AutoStorageBasePropertiesOutput {
	return o
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesOutputWithContext(ctx context.Context) AutoStorageBasePropertiesOutput {
	return o
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return o.ToAutoStorageBasePropertiesPtrOutputWithContext(context.Background())
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoStorageBaseProperties) *AutoStorageBaseProperties {
		return &v
	}).(AutoStorageBasePropertiesPtrOutput)
}

// The resource ID of the storage account to be used for auto storage account.
func (o AutoStorageBasePropertiesOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v AutoStorageBaseProperties) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

type AutoStorageBasePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AutoStorageBasePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoStorageBaseProperties)(nil)).Elem()
}

func (o AutoStorageBasePropertiesPtrOutput) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return o
}

func (o AutoStorageBasePropertiesPtrOutput) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return o
}

func (o AutoStorageBasePropertiesPtrOutput) Elem() AutoStorageBasePropertiesOutput {
	return o.ApplyT(func(v *AutoStorageBaseProperties) AutoStorageBaseProperties {
		if v != nil {
			return *v
		}
		var ret AutoStorageBaseProperties
		return ret
	}).(AutoStorageBasePropertiesOutput)
}

// The resource ID of the storage account to be used for auto storage account.
func (o AutoStorageBasePropertiesPtrOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoStorageBaseProperties) *string {
		if v == nil {
			return nil
		}
		return &v.StorageAccountId
	}).(pulumi.StringPtrOutput)
}

// Contains information about the auto storage account associated with a Batch account.
type AutoStoragePropertiesResponse struct {
	// The UTC time at which storage keys were last synchronized with the Batch account.
	LastKeySync string `pulumi:"lastKeySync"`
	// The resource ID of the storage account to be used for auto storage account.
	StorageAccountId string `pulumi:"storageAccountId"`
}

// Contains information about the auto storage account associated with a Batch account.
type AutoStoragePropertiesResponseOutput struct{ *pulumi.OutputState }

func (AutoStoragePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoStoragePropertiesResponse)(nil)).Elem()
}

func (o AutoStoragePropertiesResponseOutput) ToAutoStoragePropertiesResponseOutput() AutoStoragePropertiesResponseOutput {
	return o
}

func (o AutoStoragePropertiesResponseOutput) ToAutoStoragePropertiesResponseOutputWithContext(ctx context.Context) AutoStoragePropertiesResponseOutput {
	return o
}

// The UTC time at which storage keys were last synchronized with the Batch account.
func (o AutoStoragePropertiesResponseOutput) LastKeySync() pulumi.StringOutput {
	return o.ApplyT(func(v AutoStoragePropertiesResponse) string { return v.LastKeySync }).(pulumi.StringOutput)
}

// The resource ID of the storage account to be used for auto storage account.
func (o AutoStoragePropertiesResponseOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v AutoStoragePropertiesResponse) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

type AutoStoragePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoStoragePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoStoragePropertiesResponse)(nil)).Elem()
}

func (o AutoStoragePropertiesResponsePtrOutput) ToAutoStoragePropertiesResponsePtrOutput() AutoStoragePropertiesResponsePtrOutput {
	return o
}

func (o AutoStoragePropertiesResponsePtrOutput) ToAutoStoragePropertiesResponsePtrOutputWithContext(ctx context.Context) AutoStoragePropertiesResponsePtrOutput {
	return o
}

func (o AutoStoragePropertiesResponsePtrOutput) Elem() AutoStoragePropertiesResponseOutput {
	return o.ApplyT(func(v *AutoStoragePropertiesResponse) AutoStoragePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AutoStoragePropertiesResponse
		return ret
	}).(AutoStoragePropertiesResponseOutput)
}

// The UTC time at which storage keys were last synchronized with the Batch account.
func (o AutoStoragePropertiesResponsePtrOutput) LastKeySync() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoStoragePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastKeySync
	}).(pulumi.StringPtrOutput)
}

// The resource ID of the storage account to be used for auto storage account.
func (o AutoStoragePropertiesResponsePtrOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoStoragePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StorageAccountId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationPackageResponseOutput{})
	pulumi.RegisterOutputType(ApplicationPackageResponseArrayOutput{})
	pulumi.RegisterOutputType(AutoStorageBasePropertiesOutput{})
	pulumi.RegisterOutputType(AutoStorageBasePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AutoStoragePropertiesResponseOutput{})
	pulumi.RegisterOutputType(AutoStoragePropertiesResponsePtrOutput{})
}
