// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A SQL server registration.
//
// Deprecated: Version v20170301preview will be removed in the next major version of the provider. Upgrade to version v20190724preview or later.
func LookupSqlServerRegistration(ctx *pulumi.Context, args *LookupSqlServerRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupSqlServerRegistrationResult, error) {
	var rv LookupSqlServerRegistrationResult
	err := ctx.Invoke("azure-native:azuredata/v20170301preview:getSqlServerRegistration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlServerRegistrationArgs struct {
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the SQL Server registration.
	SqlServerRegistrationName string `pulumi:"sqlServerRegistrationName"`
}

// A SQL server registration.
type LookupSqlServerRegistrationResult struct {
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Optional Properties as JSON string
	PropertyBag *string `pulumi:"propertyBag"`
	// Resource Group Name
	ResourceGroup *string `pulumi:"resourceGroup"`
	// Subscription Id
	SubscriptionId *string `pulumi:"subscriptionId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupSqlServerRegistrationOutput(ctx *pulumi.Context, args LookupSqlServerRegistrationOutputArgs, opts ...pulumi.InvokeOption) LookupSqlServerRegistrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlServerRegistrationResult, error) {
			args := v.(LookupSqlServerRegistrationArgs)
			r, err := LookupSqlServerRegistration(ctx, &args, opts...)
			var s LookupSqlServerRegistrationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlServerRegistrationResultOutput)
}

type LookupSqlServerRegistrationOutputArgs struct {
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the SQL Server registration.
	SqlServerRegistrationName pulumi.StringInput `pulumi:"sqlServerRegistrationName"`
}

func (LookupSqlServerRegistrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlServerRegistrationArgs)(nil)).Elem()
}

// A SQL server registration.
type LookupSqlServerRegistrationResultOutput struct{ *pulumi.OutputState }

func (LookupSqlServerRegistrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlServerRegistrationResult)(nil)).Elem()
}

func (o LookupSqlServerRegistrationResultOutput) ToLookupSqlServerRegistrationResultOutput() LookupSqlServerRegistrationResultOutput {
	return o
}

func (o LookupSqlServerRegistrationResultOutput) ToLookupSqlServerRegistrationResultOutputWithContext(ctx context.Context) LookupSqlServerRegistrationResultOutput {
	return o
}

// Resource ID.
func (o LookupSqlServerRegistrationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupSqlServerRegistrationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupSqlServerRegistrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional Properties as JSON string
func (o LookupSqlServerRegistrationResultOutput) PropertyBag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) *string { return v.PropertyBag }).(pulumi.StringPtrOutput)
}

// Resource Group Name
func (o LookupSqlServerRegistrationResultOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

// Subscription Id
func (o LookupSqlServerRegistrationResultOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupSqlServerRegistrationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupSqlServerRegistrationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlServerRegistrationResultOutput{})
}
