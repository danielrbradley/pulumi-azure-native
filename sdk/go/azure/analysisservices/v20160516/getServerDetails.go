// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160516

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents an instance of an Analysis Services resource.
//
// Deprecated: Version v20160516 will be removed in the next major version of the provider. Upgrade to version v20170801 or later.
func LookupServerDetails(ctx *pulumi.Context, args *LookupServerDetailsArgs, opts ...pulumi.InvokeOption) (*LookupServerDetailsResult, error) {
	var rv LookupServerDetailsResult
	err := ctx.Invoke("azure-native:analysisservices/v20160516:getServerDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupServerDetailsArgs struct {
	// The name of the Azure Resource group of which a given Analysis Services server is part. This name must be at least 1 character in length, and no more than 90.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
	ServerName string `pulumi:"serverName"`
}

// Represents an instance of an Analysis Services resource.
type LookupServerDetailsResult struct {
	// A collection of AS server administrators
	AsAdministrators *ServerAdministratorsResponse `pulumi:"asAdministrators"`
	// The container URI of backup blob.
	BackupBlobContainerUri *string `pulumi:"backupBlobContainerUri"`
	// An identifier that represents the Analysis Services resource.
	Id string `pulumi:"id"`
	// Location of the Analysis Services resource.
	Location string `pulumi:"location"`
	// The managed mode of the server (0 = not managed, 1 = managed).
	ManagedMode *int `pulumi:"managedMode"`
	// The name of the Analysis Services resource.
	Name string `pulumi:"name"`
	// The current deployment state of Analysis Services resource. The provisioningState is to indicate states for resource provisioning.
	ProvisioningState string `pulumi:"provisioningState"`
	// The full name of the Analysis Services resource.
	ServerFullName string `pulumi:"serverFullName"`
	// The server monitor mode for AS server
	ServerMonitorMode *int `pulumi:"serverMonitorMode"`
	// The SKU of the Analysis Services resource.
	Sku ResourceSkuResponse `pulumi:"sku"`
	// The current state of Analysis Services resource. The state is to indicate more states outside of resource provisioning.
	State string `pulumi:"state"`
	// Key-value pairs of additional resource provisioning properties.
	Tags map[string]string `pulumi:"tags"`
	// The type of the Analysis Services resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupServerDetailsResult
func (val *LookupServerDetailsResult) Defaults() *LookupServerDetailsResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ManagedMode) {
		managedMode_ := 1
		tmp.ManagedMode = &managedMode_
	}
	if isZero(tmp.ServerMonitorMode) {
		serverMonitorMode_ := 1
		tmp.ServerMonitorMode = &serverMonitorMode_
	}
	tmp.Sku = *tmp.Sku.Defaults()

	return &tmp
}

func LookupServerDetailsOutput(ctx *pulumi.Context, args LookupServerDetailsOutputArgs, opts ...pulumi.InvokeOption) LookupServerDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerDetailsResult, error) {
			args := v.(LookupServerDetailsArgs)
			r, err := LookupServerDetails(ctx, &args, opts...)
			var s LookupServerDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerDetailsResultOutput)
}

type LookupServerDetailsOutputArgs struct {
	// The name of the Azure Resource group of which a given Analysis Services server is part. This name must be at least 1 character in length, and no more than 90.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerDetailsArgs)(nil)).Elem()
}

// Represents an instance of an Analysis Services resource.
type LookupServerDetailsResultOutput struct{ *pulumi.OutputState }

func (LookupServerDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerDetailsResult)(nil)).Elem()
}

func (o LookupServerDetailsResultOutput) ToLookupServerDetailsResultOutput() LookupServerDetailsResultOutput {
	return o
}

func (o LookupServerDetailsResultOutput) ToLookupServerDetailsResultOutputWithContext(ctx context.Context) LookupServerDetailsResultOutput {
	return o
}

// A collection of AS server administrators
func (o LookupServerDetailsResultOutput) AsAdministrators() ServerAdministratorsResponsePtrOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) *ServerAdministratorsResponse { return v.AsAdministrators }).(ServerAdministratorsResponsePtrOutput)
}

// The container URI of backup blob.
func (o LookupServerDetailsResultOutput) BackupBlobContainerUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) *string { return v.BackupBlobContainerUri }).(pulumi.StringPtrOutput)
}

// An identifier that represents the Analysis Services resource.
func (o LookupServerDetailsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Location of the Analysis Services resource.
func (o LookupServerDetailsResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) string { return v.Location }).(pulumi.StringOutput)
}

// The managed mode of the server (0 = not managed, 1 = managed).
func (o LookupServerDetailsResultOutput) ManagedMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) *int { return v.ManagedMode }).(pulumi.IntPtrOutput)
}

// The name of the Analysis Services resource.
func (o LookupServerDetailsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current deployment state of Analysis Services resource. The provisioningState is to indicate states for resource provisioning.
func (o LookupServerDetailsResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The full name of the Analysis Services resource.
func (o LookupServerDetailsResultOutput) ServerFullName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) string { return v.ServerFullName }).(pulumi.StringOutput)
}

// The server monitor mode for AS server
func (o LookupServerDetailsResultOutput) ServerMonitorMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) *int { return v.ServerMonitorMode }).(pulumi.IntPtrOutput)
}

// The SKU of the Analysis Services resource.
func (o LookupServerDetailsResultOutput) Sku() ResourceSkuResponseOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) ResourceSkuResponse { return v.Sku }).(ResourceSkuResponseOutput)
}

// The current state of Analysis Services resource. The state is to indicate more states outside of resource provisioning.
func (o LookupServerDetailsResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) string { return v.State }).(pulumi.StringOutput)
}

// Key-value pairs of additional resource provisioning properties.
func (o LookupServerDetailsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the Analysis Services resource.
func (o LookupServerDetailsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDetailsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerDetailsResultOutput{})
}
