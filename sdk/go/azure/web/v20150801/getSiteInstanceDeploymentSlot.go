// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents user credentials used for publishing activity
//
// Deprecated: Version v20150801 will be removed in the next major version of the provider. Upgrade to version v20150801preview or later.
func LookupSiteInstanceDeploymentSlot(ctx *pulumi.Context, args *LookupSiteInstanceDeploymentSlotArgs, opts ...pulumi.InvokeOption) (*LookupSiteInstanceDeploymentSlotResult, error) {
	var rv LookupSiteInstanceDeploymentSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteInstanceDeploymentSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteInstanceDeploymentSlotArgs struct {
	// Id of the deployment
	Id string `pulumi:"id"`
	// Id of web app instance
	InstanceId string `pulumi:"instanceId"`
	// Name of web app
	Name string `pulumi:"name"`
	// Name of resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of web app slot. If not specified then will default to production slot.
	Slot string `pulumi:"slot"`
}

// Represents user credentials used for publishing activity
type LookupSiteInstanceDeploymentSlotResult struct {
	// Active
	Active *bool `pulumi:"active"`
	// Author
	Author *string `pulumi:"author"`
	// AuthorEmail
	AuthorEmail *string `pulumi:"authorEmail"`
	// Deployer
	Deployer *string `pulumi:"deployer"`
	// Detail
	Details *string `pulumi:"details"`
	// EndTime
	EndTime *string `pulumi:"endTime"`
	// Resource Id
	Id *string `pulumi:"id"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location string `pulumi:"location"`
	// Message
	Message *string `pulumi:"message"`
	// Resource Name
	Name *string `pulumi:"name"`
	// StartTime
	StartTime *string `pulumi:"startTime"`
	// Status
	Status *int `pulumi:"status"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

func LookupSiteInstanceDeploymentSlotOutput(ctx *pulumi.Context, args LookupSiteInstanceDeploymentSlotOutputArgs, opts ...pulumi.InvokeOption) LookupSiteInstanceDeploymentSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteInstanceDeploymentSlotResult, error) {
			args := v.(LookupSiteInstanceDeploymentSlotArgs)
			r, err := LookupSiteInstanceDeploymentSlot(ctx, &args, opts...)
			var s LookupSiteInstanceDeploymentSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteInstanceDeploymentSlotResultOutput)
}

type LookupSiteInstanceDeploymentSlotOutputArgs struct {
	// Id of the deployment
	Id pulumi.StringInput `pulumi:"id"`
	// Id of web app instance
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Name of web app
	Name pulumi.StringInput `pulumi:"name"`
	// Name of resource group
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of web app slot. If not specified then will default to production slot.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (LookupSiteInstanceDeploymentSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteInstanceDeploymentSlotArgs)(nil)).Elem()
}

// Represents user credentials used for publishing activity
type LookupSiteInstanceDeploymentSlotResultOutput struct{ *pulumi.OutputState }

func (LookupSiteInstanceDeploymentSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteInstanceDeploymentSlotResult)(nil)).Elem()
}

func (o LookupSiteInstanceDeploymentSlotResultOutput) ToLookupSiteInstanceDeploymentSlotResultOutput() LookupSiteInstanceDeploymentSlotResultOutput {
	return o
}

func (o LookupSiteInstanceDeploymentSlotResultOutput) ToLookupSiteInstanceDeploymentSlotResultOutputWithContext(ctx context.Context) LookupSiteInstanceDeploymentSlotResultOutput {
	return o
}

// Active
func (o LookupSiteInstanceDeploymentSlotResultOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *bool { return v.Active }).(pulumi.BoolPtrOutput)
}

// Author
func (o LookupSiteInstanceDeploymentSlotResultOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.Author }).(pulumi.StringPtrOutput)
}

// AuthorEmail
func (o LookupSiteInstanceDeploymentSlotResultOutput) AuthorEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.AuthorEmail }).(pulumi.StringPtrOutput)
}

// Deployer
func (o LookupSiteInstanceDeploymentSlotResultOutput) Deployer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.Deployer }).(pulumi.StringPtrOutput)
}

// Detail
func (o LookupSiteInstanceDeploymentSlotResultOutput) Details() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.Details }).(pulumi.StringPtrOutput)
}

// EndTime
func (o LookupSiteInstanceDeploymentSlotResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupSiteInstanceDeploymentSlotResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Kind of resource
func (o LookupSiteInstanceDeploymentSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Location
func (o LookupSiteInstanceDeploymentSlotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) string { return v.Location }).(pulumi.StringOutput)
}

// Message
func (o LookupSiteInstanceDeploymentSlotResultOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.Message }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o LookupSiteInstanceDeploymentSlotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// StartTime
func (o LookupSiteInstanceDeploymentSlotResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

// Status
func (o LookupSiteInstanceDeploymentSlotResultOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *int { return v.Status }).(pulumi.IntPtrOutput)
}

// Resource tags
func (o LookupSiteInstanceDeploymentSlotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupSiteInstanceDeploymentSlotResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteInstanceDeploymentSlotResultOutput{})
}
