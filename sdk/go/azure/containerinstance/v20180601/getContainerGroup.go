// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A container group.
//
// Deprecated: Version v20180601 will be removed in the next major version of the provider. Upgrade to version v20210301 or later.
func LookupContainerGroup(ctx *pulumi.Context, args *LookupContainerGroupArgs, opts ...pulumi.InvokeOption) (*LookupContainerGroupResult, error) {
	var rv LookupContainerGroupResult
	err := ctx.Invoke("azure-native:containerinstance/v20180601:getContainerGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContainerGroupArgs struct {
	// The name of the container group.
	ContainerGroupName string `pulumi:"containerGroupName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A container group.
type LookupContainerGroupResult struct {
	// The containers within the container group.
	Containers []ContainerResponse `pulumi:"containers"`
	// The diagnostic information for a container group.
	Diagnostics *ContainerGroupDiagnosticsResponse `pulumi:"diagnostics"`
	// The resource id.
	Id string `pulumi:"id"`
	// The image registry credentials by which the container group is created from.
	ImageRegistryCredentials []ImageRegistryCredentialResponse `pulumi:"imageRegistryCredentials"`
	// The instance view of the container group. Only valid in response.
	InstanceView ContainerGroupResponseInstanceView `pulumi:"instanceView"`
	// The IP address type of the container group.
	IpAddress *IpAddressResponse `pulumi:"ipAddress"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource name.
	Name string `pulumi:"name"`
	// The operating system type required by the containers in the container group.
	OsType string `pulumi:"osType"`
	// The provisioning state of the container group. This only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// Restart policy for all containers within the container group.
	// - `Always` Always restart
	// - `OnFailure` Restart on failure
	// - `Never` Never restart
	RestartPolicy *string `pulumi:"restartPolicy"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
	// The list of volumes that can be mounted by containers in this container group.
	Volumes []VolumeResponse `pulumi:"volumes"`
}

func LookupContainerGroupOutput(ctx *pulumi.Context, args LookupContainerGroupOutputArgs, opts ...pulumi.InvokeOption) LookupContainerGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerGroupResult, error) {
			args := v.(LookupContainerGroupArgs)
			r, err := LookupContainerGroup(ctx, &args, opts...)
			var s LookupContainerGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContainerGroupResultOutput)
}

type LookupContainerGroupOutputArgs struct {
	// The name of the container group.
	ContainerGroupName pulumi.StringInput `pulumi:"containerGroupName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupContainerGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerGroupArgs)(nil)).Elem()
}

// A container group.
type LookupContainerGroupResultOutput struct{ *pulumi.OutputState }

func (LookupContainerGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerGroupResult)(nil)).Elem()
}

func (o LookupContainerGroupResultOutput) ToLookupContainerGroupResultOutput() LookupContainerGroupResultOutput {
	return o
}

func (o LookupContainerGroupResultOutput) ToLookupContainerGroupResultOutputWithContext(ctx context.Context) LookupContainerGroupResultOutput {
	return o
}

// The containers within the container group.
func (o LookupContainerGroupResultOutput) Containers() ContainerResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []ContainerResponse { return v.Containers }).(ContainerResponseArrayOutput)
}

// The diagnostic information for a container group.
func (o LookupContainerGroupResultOutput) Diagnostics() ContainerGroupDiagnosticsResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *ContainerGroupDiagnosticsResponse { return v.Diagnostics }).(ContainerGroupDiagnosticsResponsePtrOutput)
}

// The resource id.
func (o LookupContainerGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The image registry credentials by which the container group is created from.
func (o LookupContainerGroupResultOutput) ImageRegistryCredentials() ImageRegistryCredentialResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []ImageRegistryCredentialResponse {
		return v.ImageRegistryCredentials
	}).(ImageRegistryCredentialResponseArrayOutput)
}

// The instance view of the container group. Only valid in response.
func (o LookupContainerGroupResultOutput) InstanceView() ContainerGroupResponseInstanceViewOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) ContainerGroupResponseInstanceView { return v.InstanceView }).(ContainerGroupResponseInstanceViewOutput)
}

// The IP address type of the container group.
func (o LookupContainerGroupResultOutput) IpAddress() IpAddressResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *IpAddressResponse { return v.IpAddress }).(IpAddressResponsePtrOutput)
}

// The resource location.
func (o LookupContainerGroupResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name.
func (o LookupContainerGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The operating system type required by the containers in the container group.
func (o LookupContainerGroupResultOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.OsType }).(pulumi.StringOutput)
}

// The provisioning state of the container group. This only appears in the response.
func (o LookupContainerGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Restart policy for all containers within the container group.
// - `Always` Always restart
// - `OnFailure` Restart on failure
// - `Never` Never restart
func (o LookupContainerGroupResultOutput) RestartPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) *string { return v.RestartPolicy }).(pulumi.StringPtrOutput)
}

// The resource tags.
func (o LookupContainerGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o LookupContainerGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

// The list of volumes that can be mounted by containers in this container group.
func (o LookupContainerGroupResultOutput) Volumes() VolumeResponseArrayOutput {
	return o.ApplyT(func(v LookupContainerGroupResult) []VolumeResponse { return v.Volumes }).(VolumeResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerGroupResultOutput{})
}
