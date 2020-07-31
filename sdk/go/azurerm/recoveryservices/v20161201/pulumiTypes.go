// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20161201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Base class for container with backup items. Containers with specific workloads are derived from this class.
type ProtectionContainerType struct {
	// Optional ETag.
	ETag *string `pulumi:"eTag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name associated with the resource.
	Name string `pulumi:"name"`
	// ProtectionContainerResource properties
	Properties ProtectionContainerResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type string `pulumi:"type"`
}

// ProtectionContainerTypeInput is an input type that accepts ProtectionContainerTypeArgs and ProtectionContainerTypeOutput values.
// You can construct a concrete instance of `ProtectionContainerTypeInput` via:
//
//          ProtectionContainerTypeArgs{...}
type ProtectionContainerTypeInput interface {
	pulumi.Input

	ToProtectionContainerTypeOutput() ProtectionContainerTypeOutput
	ToProtectionContainerTypeOutputWithContext(context.Context) ProtectionContainerTypeOutput
}

// Base class for container with backup items. Containers with specific workloads are derived from this class.
type ProtectionContainerTypeArgs struct {
	// Optional ETag.
	ETag pulumi.StringPtrInput `pulumi:"eTag"`
	// Resource location.
	Location pulumi.StringPtrInput `pulumi:"location"`
	// Resource name associated with the resource.
	Name pulumi.StringInput `pulumi:"name"`
	// ProtectionContainerResource properties
	Properties ProtectionContainerResponseInput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type pulumi.StringInput `pulumi:"type"`
}

func (ProtectionContainerTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionContainerType)(nil)).Elem()
}

func (i ProtectionContainerTypeArgs) ToProtectionContainerTypeOutput() ProtectionContainerTypeOutput {
	return i.ToProtectionContainerTypeOutputWithContext(context.Background())
}

func (i ProtectionContainerTypeArgs) ToProtectionContainerTypeOutputWithContext(ctx context.Context) ProtectionContainerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionContainerTypeOutput)
}

// Base class for container with backup items. Containers with specific workloads are derived from this class.
type ProtectionContainerTypeOutput struct{ *pulumi.OutputState }

func (ProtectionContainerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionContainerType)(nil)).Elem()
}

func (o ProtectionContainerTypeOutput) ToProtectionContainerTypeOutput() ProtectionContainerTypeOutput {
	return o
}

func (o ProtectionContainerTypeOutput) ToProtectionContainerTypeOutputWithContext(ctx context.Context) ProtectionContainerTypeOutput {
	return o
}

// Optional ETag.
func (o ProtectionContainerTypeOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerType) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o ProtectionContainerTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerType) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name associated with the resource.
func (o ProtectionContainerTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ProtectionContainerType) string { return v.Name }).(pulumi.StringOutput)
}

// ProtectionContainerResource properties
func (o ProtectionContainerTypeOutput) Properties() ProtectionContainerResponseOutput {
	return o.ApplyT(func(v ProtectionContainerType) ProtectionContainerResponse { return v.Properties }).(ProtectionContainerResponseOutput)
}

// Resource tags.
func (o ProtectionContainerTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ProtectionContainerType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
func (o ProtectionContainerTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ProtectionContainerType) string { return v.Type }).(pulumi.StringOutput)
}

// Base class for container with backup items. Containers with specific workloads are derived from this class.
type ProtectionContainerDefinition struct {
	// Type of backup management for the container.
	BackupManagementType *string `pulumi:"backupManagementType"`
	// Type of the container. The value of this property for: 1. Compute Azure VM is Microsoft.Compute/virtualMachines 2.
	// Classic Compute Azure VM is Microsoft.ClassicCompute/virtualMachines 3. Windows machines (like MAB, DPM etc) is
	// Windows 4. Azure SQL instance is AzureSqlContainer. 5. Storage containers is StorageContainer. 6. Azure workload
	// Backup is VMAppContainer
	ContainerType *string `pulumi:"containerType"`
	// Friendly name of the container.
	FriendlyName *string `pulumi:"friendlyName"`
	// Status of health of the container.
	HealthStatus *string `pulumi:"healthStatus"`
	// Status of registration of the container with the Recovery Services Vault.
	RegistrationStatus *string `pulumi:"registrationStatus"`
}

// ProtectionContainerDefinitionInput is an input type that accepts ProtectionContainerDefinitionArgs and ProtectionContainerDefinitionOutput values.
// You can construct a concrete instance of `ProtectionContainerDefinitionInput` via:
//
//          ProtectionContainerDefinitionArgs{...}
type ProtectionContainerDefinitionInput interface {
	pulumi.Input

	ToProtectionContainerDefinitionOutput() ProtectionContainerDefinitionOutput
	ToProtectionContainerDefinitionOutputWithContext(context.Context) ProtectionContainerDefinitionOutput
}

// Base class for container with backup items. Containers with specific workloads are derived from this class.
type ProtectionContainerDefinitionArgs struct {
	// Type of backup management for the container.
	BackupManagementType pulumi.StringPtrInput `pulumi:"backupManagementType"`
	// Type of the container. The value of this property for: 1. Compute Azure VM is Microsoft.Compute/virtualMachines 2.
	// Classic Compute Azure VM is Microsoft.ClassicCompute/virtualMachines 3. Windows machines (like MAB, DPM etc) is
	// Windows 4. Azure SQL instance is AzureSqlContainer. 5. Storage containers is StorageContainer. 6. Azure workload
	// Backup is VMAppContainer
	ContainerType pulumi.StringPtrInput `pulumi:"containerType"`
	// Friendly name of the container.
	FriendlyName pulumi.StringPtrInput `pulumi:"friendlyName"`
	// Status of health of the container.
	HealthStatus pulumi.StringPtrInput `pulumi:"healthStatus"`
	// Status of registration of the container with the Recovery Services Vault.
	RegistrationStatus pulumi.StringPtrInput `pulumi:"registrationStatus"`
}

func (ProtectionContainerDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionContainerDefinition)(nil)).Elem()
}

func (i ProtectionContainerDefinitionArgs) ToProtectionContainerDefinitionOutput() ProtectionContainerDefinitionOutput {
	return i.ToProtectionContainerDefinitionOutputWithContext(context.Background())
}

func (i ProtectionContainerDefinitionArgs) ToProtectionContainerDefinitionOutputWithContext(ctx context.Context) ProtectionContainerDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionContainerDefinitionOutput)
}

func (i ProtectionContainerDefinitionArgs) ToProtectionContainerDefinitionPtrOutput() ProtectionContainerDefinitionPtrOutput {
	return i.ToProtectionContainerDefinitionPtrOutputWithContext(context.Background())
}

func (i ProtectionContainerDefinitionArgs) ToProtectionContainerDefinitionPtrOutputWithContext(ctx context.Context) ProtectionContainerDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionContainerDefinitionOutput).ToProtectionContainerDefinitionPtrOutputWithContext(ctx)
}

// ProtectionContainerDefinitionPtrInput is an input type that accepts ProtectionContainerDefinitionArgs, ProtectionContainerDefinitionPtr and ProtectionContainerDefinitionPtrOutput values.
// You can construct a concrete instance of `ProtectionContainerDefinitionPtrInput` via:
//
//          ProtectionContainerDefinitionArgs{...}
//
//  or:
//
//          nil
type ProtectionContainerDefinitionPtrInput interface {
	pulumi.Input

	ToProtectionContainerDefinitionPtrOutput() ProtectionContainerDefinitionPtrOutput
	ToProtectionContainerDefinitionPtrOutputWithContext(context.Context) ProtectionContainerDefinitionPtrOutput
}

type protectionContainerDefinitionPtrType ProtectionContainerDefinitionArgs

func ProtectionContainerDefinitionPtr(v *ProtectionContainerDefinitionArgs) ProtectionContainerDefinitionPtrInput {
	return (*protectionContainerDefinitionPtrType)(v)
}

func (*protectionContainerDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionContainerDefinition)(nil)).Elem()
}

func (i *protectionContainerDefinitionPtrType) ToProtectionContainerDefinitionPtrOutput() ProtectionContainerDefinitionPtrOutput {
	return i.ToProtectionContainerDefinitionPtrOutputWithContext(context.Background())
}

func (i *protectionContainerDefinitionPtrType) ToProtectionContainerDefinitionPtrOutputWithContext(ctx context.Context) ProtectionContainerDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionContainerDefinitionPtrOutput)
}

// Base class for container with backup items. Containers with specific workloads are derived from this class.
type ProtectionContainerDefinitionOutput struct{ *pulumi.OutputState }

func (ProtectionContainerDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionContainerDefinition)(nil)).Elem()
}

func (o ProtectionContainerDefinitionOutput) ToProtectionContainerDefinitionOutput() ProtectionContainerDefinitionOutput {
	return o
}

func (o ProtectionContainerDefinitionOutput) ToProtectionContainerDefinitionOutputWithContext(ctx context.Context) ProtectionContainerDefinitionOutput {
	return o
}

func (o ProtectionContainerDefinitionOutput) ToProtectionContainerDefinitionPtrOutput() ProtectionContainerDefinitionPtrOutput {
	return o.ToProtectionContainerDefinitionPtrOutputWithContext(context.Background())
}

func (o ProtectionContainerDefinitionOutput) ToProtectionContainerDefinitionPtrOutputWithContext(ctx context.Context) ProtectionContainerDefinitionPtrOutput {
	return o.ApplyT(func(v ProtectionContainerDefinition) *ProtectionContainerDefinition {
		return &v
	}).(ProtectionContainerDefinitionPtrOutput)
}

// Type of backup management for the container.
func (o ProtectionContainerDefinitionOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerDefinition) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

// Type of the container. The value of this property for: 1. Compute Azure VM is Microsoft.Compute/virtualMachines 2.
// Classic Compute Azure VM is Microsoft.ClassicCompute/virtualMachines 3. Windows machines (like MAB, DPM etc) is
// Windows 4. Azure SQL instance is AzureSqlContainer. 5. Storage containers is StorageContainer. 6. Azure workload
// Backup is VMAppContainer
func (o ProtectionContainerDefinitionOutput) ContainerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerDefinition) *string { return v.ContainerType }).(pulumi.StringPtrOutput)
}

// Friendly name of the container.
func (o ProtectionContainerDefinitionOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerDefinition) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// Status of health of the container.
func (o ProtectionContainerDefinitionOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerDefinition) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

// Status of registration of the container with the Recovery Services Vault.
func (o ProtectionContainerDefinitionOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerDefinition) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

type ProtectionContainerDefinitionPtrOutput struct{ *pulumi.OutputState }

func (ProtectionContainerDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionContainerDefinition)(nil)).Elem()
}

func (o ProtectionContainerDefinitionPtrOutput) ToProtectionContainerDefinitionPtrOutput() ProtectionContainerDefinitionPtrOutput {
	return o
}

func (o ProtectionContainerDefinitionPtrOutput) ToProtectionContainerDefinitionPtrOutputWithContext(ctx context.Context) ProtectionContainerDefinitionPtrOutput {
	return o
}

func (o ProtectionContainerDefinitionPtrOutput) Elem() ProtectionContainerDefinitionOutput {
	return o.ApplyT(func(v *ProtectionContainerDefinition) ProtectionContainerDefinition { return *v }).(ProtectionContainerDefinitionOutput)
}

// Type of backup management for the container.
func (o ProtectionContainerDefinitionPtrOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerDefinition) *string {
		if v == nil {
			return nil
		}
		return v.BackupManagementType
	}).(pulumi.StringPtrOutput)
}

// Type of the container. The value of this property for: 1. Compute Azure VM is Microsoft.Compute/virtualMachines 2.
// Classic Compute Azure VM is Microsoft.ClassicCompute/virtualMachines 3. Windows machines (like MAB, DPM etc) is
// Windows 4. Azure SQL instance is AzureSqlContainer. 5. Storage containers is StorageContainer. 6. Azure workload
// Backup is VMAppContainer
func (o ProtectionContainerDefinitionPtrOutput) ContainerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerDefinition) *string {
		if v == nil {
			return nil
		}
		return v.ContainerType
	}).(pulumi.StringPtrOutput)
}

// Friendly name of the container.
func (o ProtectionContainerDefinitionPtrOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerDefinition) *string {
		if v == nil {
			return nil
		}
		return v.FriendlyName
	}).(pulumi.StringPtrOutput)
}

// Status of health of the container.
func (o ProtectionContainerDefinitionPtrOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerDefinition) *string {
		if v == nil {
			return nil
		}
		return v.HealthStatus
	}).(pulumi.StringPtrOutput)
}

// Status of registration of the container with the Recovery Services Vault.
func (o ProtectionContainerDefinitionPtrOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerDefinition) *string {
		if v == nil {
			return nil
		}
		return v.RegistrationStatus
	}).(pulumi.StringPtrOutput)
}

// Base class for container with backup items. Containers with specific workloads are derived from this class.
type ProtectionContainerResponse struct {
	// Type of backup management for the container.
	BackupManagementType *string `pulumi:"backupManagementType"`
	// Type of the container. The value of this property for: 1. Compute Azure VM is Microsoft.Compute/virtualMachines 2.
	// Classic Compute Azure VM is Microsoft.ClassicCompute/virtualMachines 3. Windows machines (like MAB, DPM etc) is
	// Windows 4. Azure SQL instance is AzureSqlContainer. 5. Storage containers is StorageContainer. 6. Azure workload
	// Backup is VMAppContainer
	ContainerType *string `pulumi:"containerType"`
	// Friendly name of the container.
	FriendlyName *string `pulumi:"friendlyName"`
	// Status of health of the container.
	HealthStatus *string `pulumi:"healthStatus"`
	// Status of registration of the container with the Recovery Services Vault.
	RegistrationStatus *string `pulumi:"registrationStatus"`
}

// ProtectionContainerResponseInput is an input type that accepts ProtectionContainerResponseArgs and ProtectionContainerResponseOutput values.
// You can construct a concrete instance of `ProtectionContainerResponseInput` via:
//
//          ProtectionContainerResponseArgs{...}
type ProtectionContainerResponseInput interface {
	pulumi.Input

	ToProtectionContainerResponseOutput() ProtectionContainerResponseOutput
	ToProtectionContainerResponseOutputWithContext(context.Context) ProtectionContainerResponseOutput
}

// Base class for container with backup items. Containers with specific workloads are derived from this class.
type ProtectionContainerResponseArgs struct {
	// Type of backup management for the container.
	BackupManagementType pulumi.StringPtrInput `pulumi:"backupManagementType"`
	// Type of the container. The value of this property for: 1. Compute Azure VM is Microsoft.Compute/virtualMachines 2.
	// Classic Compute Azure VM is Microsoft.ClassicCompute/virtualMachines 3. Windows machines (like MAB, DPM etc) is
	// Windows 4. Azure SQL instance is AzureSqlContainer. 5. Storage containers is StorageContainer. 6. Azure workload
	// Backup is VMAppContainer
	ContainerType pulumi.StringPtrInput `pulumi:"containerType"`
	// Friendly name of the container.
	FriendlyName pulumi.StringPtrInput `pulumi:"friendlyName"`
	// Status of health of the container.
	HealthStatus pulumi.StringPtrInput `pulumi:"healthStatus"`
	// Status of registration of the container with the Recovery Services Vault.
	RegistrationStatus pulumi.StringPtrInput `pulumi:"registrationStatus"`
}

func (ProtectionContainerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionContainerResponse)(nil)).Elem()
}

func (i ProtectionContainerResponseArgs) ToProtectionContainerResponseOutput() ProtectionContainerResponseOutput {
	return i.ToProtectionContainerResponseOutputWithContext(context.Background())
}

func (i ProtectionContainerResponseArgs) ToProtectionContainerResponseOutputWithContext(ctx context.Context) ProtectionContainerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionContainerResponseOutput)
}

func (i ProtectionContainerResponseArgs) ToProtectionContainerResponsePtrOutput() ProtectionContainerResponsePtrOutput {
	return i.ToProtectionContainerResponsePtrOutputWithContext(context.Background())
}

func (i ProtectionContainerResponseArgs) ToProtectionContainerResponsePtrOutputWithContext(ctx context.Context) ProtectionContainerResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionContainerResponseOutput).ToProtectionContainerResponsePtrOutputWithContext(ctx)
}

// ProtectionContainerResponsePtrInput is an input type that accepts ProtectionContainerResponseArgs, ProtectionContainerResponsePtr and ProtectionContainerResponsePtrOutput values.
// You can construct a concrete instance of `ProtectionContainerResponsePtrInput` via:
//
//          ProtectionContainerResponseArgs{...}
//
//  or:
//
//          nil
type ProtectionContainerResponsePtrInput interface {
	pulumi.Input

	ToProtectionContainerResponsePtrOutput() ProtectionContainerResponsePtrOutput
	ToProtectionContainerResponsePtrOutputWithContext(context.Context) ProtectionContainerResponsePtrOutput
}

type protectionContainerResponsePtrType ProtectionContainerResponseArgs

func ProtectionContainerResponsePtr(v *ProtectionContainerResponseArgs) ProtectionContainerResponsePtrInput {
	return (*protectionContainerResponsePtrType)(v)
}

func (*protectionContainerResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionContainerResponse)(nil)).Elem()
}

func (i *protectionContainerResponsePtrType) ToProtectionContainerResponsePtrOutput() ProtectionContainerResponsePtrOutput {
	return i.ToProtectionContainerResponsePtrOutputWithContext(context.Background())
}

func (i *protectionContainerResponsePtrType) ToProtectionContainerResponsePtrOutputWithContext(ctx context.Context) ProtectionContainerResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionContainerResponsePtrOutput)
}

// Base class for container with backup items. Containers with specific workloads are derived from this class.
type ProtectionContainerResponseOutput struct{ *pulumi.OutputState }

func (ProtectionContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionContainerResponse)(nil)).Elem()
}

func (o ProtectionContainerResponseOutput) ToProtectionContainerResponseOutput() ProtectionContainerResponseOutput {
	return o
}

func (o ProtectionContainerResponseOutput) ToProtectionContainerResponseOutputWithContext(ctx context.Context) ProtectionContainerResponseOutput {
	return o
}

func (o ProtectionContainerResponseOutput) ToProtectionContainerResponsePtrOutput() ProtectionContainerResponsePtrOutput {
	return o.ToProtectionContainerResponsePtrOutputWithContext(context.Background())
}

func (o ProtectionContainerResponseOutput) ToProtectionContainerResponsePtrOutputWithContext(ctx context.Context) ProtectionContainerResponsePtrOutput {
	return o.ApplyT(func(v ProtectionContainerResponse) *ProtectionContainerResponse {
		return &v
	}).(ProtectionContainerResponsePtrOutput)
}

// Type of backup management for the container.
func (o ProtectionContainerResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

// Type of the container. The value of this property for: 1. Compute Azure VM is Microsoft.Compute/virtualMachines 2.
// Classic Compute Azure VM is Microsoft.ClassicCompute/virtualMachines 3. Windows machines (like MAB, DPM etc) is
// Windows 4. Azure SQL instance is AzureSqlContainer. 5. Storage containers is StorageContainer. 6. Azure workload
// Backup is VMAppContainer
func (o ProtectionContainerResponseOutput) ContainerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerResponse) *string { return v.ContainerType }).(pulumi.StringPtrOutput)
}

// Friendly name of the container.
func (o ProtectionContainerResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// Status of health of the container.
func (o ProtectionContainerResponseOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerResponse) *string { return v.HealthStatus }).(pulumi.StringPtrOutput)
}

// Status of registration of the container with the Recovery Services Vault.
func (o ProtectionContainerResponseOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerResponse) *string { return v.RegistrationStatus }).(pulumi.StringPtrOutput)
}

type ProtectionContainerResponsePtrOutput struct{ *pulumi.OutputState }

func (ProtectionContainerResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionContainerResponse)(nil)).Elem()
}

func (o ProtectionContainerResponsePtrOutput) ToProtectionContainerResponsePtrOutput() ProtectionContainerResponsePtrOutput {
	return o
}

func (o ProtectionContainerResponsePtrOutput) ToProtectionContainerResponsePtrOutputWithContext(ctx context.Context) ProtectionContainerResponsePtrOutput {
	return o
}

func (o ProtectionContainerResponsePtrOutput) Elem() ProtectionContainerResponseOutput {
	return o.ApplyT(func(v *ProtectionContainerResponse) ProtectionContainerResponse { return *v }).(ProtectionContainerResponseOutput)
}

// Type of backup management for the container.
func (o ProtectionContainerResponsePtrOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerResponse) *string {
		if v == nil {
			return nil
		}
		return v.BackupManagementType
	}).(pulumi.StringPtrOutput)
}

// Type of the container. The value of this property for: 1. Compute Azure VM is Microsoft.Compute/virtualMachines 2.
// Classic Compute Azure VM is Microsoft.ClassicCompute/virtualMachines 3. Windows machines (like MAB, DPM etc) is
// Windows 4. Azure SQL instance is AzureSqlContainer. 5. Storage containers is StorageContainer. 6. Azure workload
// Backup is VMAppContainer
func (o ProtectionContainerResponsePtrOutput) ContainerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContainerType
	}).(pulumi.StringPtrOutput)
}

// Friendly name of the container.
func (o ProtectionContainerResponsePtrOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerResponse) *string {
		if v == nil {
			return nil
		}
		return v.FriendlyName
	}).(pulumi.StringPtrOutput)
}

// Status of health of the container.
func (o ProtectionContainerResponsePtrOutput) HealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerResponse) *string {
		if v == nil {
			return nil
		}
		return v.HealthStatus
	}).(pulumi.StringPtrOutput)
}

// Status of registration of the container with the Recovery Services Vault.
func (o ProtectionContainerResponsePtrOutput) RegistrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerResponse) *string {
		if v == nil {
			return nil
		}
		return v.RegistrationStatus
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ProtectionContainerTypeOutput{})
	pulumi.RegisterOutputType(ProtectionContainerDefinitionOutput{})
	pulumi.RegisterOutputType(ProtectionContainerDefinitionPtrOutput{})
	pulumi.RegisterOutputType(ProtectionContainerResponseOutput{})
	pulumi.RegisterOutputType(ProtectionContainerResponsePtrOutput{})
}