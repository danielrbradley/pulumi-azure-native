// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20171111preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a Blueprint artifact.
//
// Deprecated: Version v20171111preview will be removed in the next major version of the provider. Upgrade to version v20181101preview or later.
func LookupArtifact(ctx *pulumi.Context, args *LookupArtifactArgs, opts ...pulumi.InvokeOption) (*LookupArtifactResult, error) {
	var rv LookupArtifactResult
	err := ctx.Invoke("azure-native:blueprint/v20171111preview:getArtifact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArtifactArgs struct {
	// name of the artifact.
	ArtifactName string `pulumi:"artifactName"`
	// name of the blueprint.
	BlueprintName string `pulumi:"blueprintName"`
	// ManagementGroup where blueprint stores.
	ManagementGroupName string `pulumi:"managementGroupName"`
}

// Represents a Blueprint artifact.
type LookupArtifactResult struct {
	// String Id used to locate any resource on Azure.
	Id string `pulumi:"id"`
	// Specifies the kind of Blueprint artifact.
	Kind string `pulumi:"kind"`
	// Name of this resource.
	Name string `pulumi:"name"`
	// Type of this resource.
	Type string `pulumi:"type"`
}

func LookupArtifactOutput(ctx *pulumi.Context, args LookupArtifactOutputArgs, opts ...pulumi.InvokeOption) LookupArtifactResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupArtifactResult, error) {
			args := v.(LookupArtifactArgs)
			r, err := LookupArtifact(ctx, &args, opts...)
			var s LookupArtifactResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupArtifactResultOutput)
}

type LookupArtifactOutputArgs struct {
	// name of the artifact.
	ArtifactName pulumi.StringInput `pulumi:"artifactName"`
	// name of the blueprint.
	BlueprintName pulumi.StringInput `pulumi:"blueprintName"`
	// ManagementGroup where blueprint stores.
	ManagementGroupName pulumi.StringInput `pulumi:"managementGroupName"`
}

func (LookupArtifactOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArtifactArgs)(nil)).Elem()
}

// Represents a Blueprint artifact.
type LookupArtifactResultOutput struct{ *pulumi.OutputState }

func (LookupArtifactResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArtifactResult)(nil)).Elem()
}

func (o LookupArtifactResultOutput) ToLookupArtifactResultOutput() LookupArtifactResultOutput {
	return o
}

func (o LookupArtifactResultOutput) ToLookupArtifactResultOutputWithContext(ctx context.Context) LookupArtifactResultOutput {
	return o
}

// String Id used to locate any resource on Azure.
func (o LookupArtifactResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specifies the kind of Blueprint artifact.
func (o LookupArtifactResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of this resource.
func (o LookupArtifactResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactResult) string { return v.Name }).(pulumi.StringOutput)
}

// Type of this resource.
func (o LookupArtifactResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupArtifactResultOutput{})
}
