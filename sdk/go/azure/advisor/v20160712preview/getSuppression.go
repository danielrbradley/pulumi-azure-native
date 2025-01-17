// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160712preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The details of the snoozed or dismissed rule; for example, the duration, name, and GUID associated with the rule.
//
// Deprecated: Version v20160712preview will be removed in the next major version of the provider. Upgrade to version v20200101 or later.
func LookupSuppression(ctx *pulumi.Context, args *LookupSuppressionArgs, opts ...pulumi.InvokeOption) (*LookupSuppressionResult, error) {
	var rv LookupSuppressionResult
	err := ctx.Invoke("azure-native:advisor/v20160712preview:getSuppression", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSuppressionArgs struct {
	// The name of the suppression.
	Name string `pulumi:"name"`
	// The recommendation ID.
	RecommendationId string `pulumi:"recommendationId"`
	// The fully qualified Azure Resource Manager identifier of the resource to which the recommendation applies.
	ResourceUri string `pulumi:"resourceUri"`
}

// The details of the snoozed or dismissed rule; for example, the duration, name, and GUID associated with the rule.
type LookupSuppressionResult struct {
	// The resource ID.
	Id string `pulumi:"id"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The GUID of the suppression.
	SuppressionId *string `pulumi:"suppressionId"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The duration for which the suppression is valid.
	Ttl *string `pulumi:"ttl"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupSuppressionOutput(ctx *pulumi.Context, args LookupSuppressionOutputArgs, opts ...pulumi.InvokeOption) LookupSuppressionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSuppressionResult, error) {
			args := v.(LookupSuppressionArgs)
			r, err := LookupSuppression(ctx, &args, opts...)
			var s LookupSuppressionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSuppressionResultOutput)
}

type LookupSuppressionOutputArgs struct {
	// The name of the suppression.
	Name pulumi.StringInput `pulumi:"name"`
	// The recommendation ID.
	RecommendationId pulumi.StringInput `pulumi:"recommendationId"`
	// The fully qualified Azure Resource Manager identifier of the resource to which the recommendation applies.
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (LookupSuppressionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSuppressionArgs)(nil)).Elem()
}

// The details of the snoozed or dismissed rule; for example, the duration, name, and GUID associated with the rule.
type LookupSuppressionResultOutput struct{ *pulumi.OutputState }

func (LookupSuppressionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSuppressionResult)(nil)).Elem()
}

func (o LookupSuppressionResultOutput) ToLookupSuppressionResultOutput() LookupSuppressionResultOutput {
	return o
}

func (o LookupSuppressionResultOutput) ToLookupSuppressionResultOutputWithContext(ctx context.Context) LookupSuppressionResultOutput {
	return o
}

// The resource ID.
func (o LookupSuppressionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSuppressionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource. This cannot be changed after the resource is created.
func (o LookupSuppressionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSuppressionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupSuppressionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSuppressionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The GUID of the suppression.
func (o LookupSuppressionResultOutput) SuppressionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSuppressionResult) *string { return v.SuppressionId }).(pulumi.StringPtrOutput)
}

// The tags of the resource.
func (o LookupSuppressionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSuppressionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The duration for which the suppression is valid.
func (o LookupSuppressionResultOutput) Ttl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSuppressionResult) *string { return v.Ttl }).(pulumi.StringPtrOutput)
}

// The type of the resource.
func (o LookupSuppressionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSuppressionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSuppressionResultOutput{})
}
