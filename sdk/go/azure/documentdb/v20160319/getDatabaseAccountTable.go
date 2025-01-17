// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160319

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Cosmos DB Table.
//
// Deprecated: Version v20160319 will be removed in the next major version of the provider. Upgrade to version v20210301preview or later.
func LookupDatabaseAccountTable(ctx *pulumi.Context, args *LookupDatabaseAccountTableArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountTableResult, error) {
	var rv LookupDatabaseAccountTableResult
	err := ctx.Invoke("azure-native:documentdb/v20160319:getDatabaseAccountTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseAccountTableArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Cosmos DB table name.
	TableName string `pulumi:"tableName"`
}

// An Azure Cosmos DB Table.
type LookupDatabaseAccountTableResult struct {
	// The unique resource identifier of the database account.
	Id string `pulumi:"id"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the database account.
	Name string `pulumi:"name"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type string `pulumi:"type"`
}

func LookupDatabaseAccountTableOutput(ctx *pulumi.Context, args LookupDatabaseAccountTableOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseAccountTableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseAccountTableResult, error) {
			args := v.(LookupDatabaseAccountTableArgs)
			r, err := LookupDatabaseAccountTable(ctx, &args, opts...)
			var s LookupDatabaseAccountTableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseAccountTableResultOutput)
}

type LookupDatabaseAccountTableOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Cosmos DB table name.
	TableName pulumi.StringInput `pulumi:"tableName"`
}

func (LookupDatabaseAccountTableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountTableArgs)(nil)).Elem()
}

// An Azure Cosmos DB Table.
type LookupDatabaseAccountTableResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseAccountTableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountTableResult)(nil)).Elem()
}

func (o LookupDatabaseAccountTableResultOutput) ToLookupDatabaseAccountTableResultOutput() LookupDatabaseAccountTableResultOutput {
	return o
}

func (o LookupDatabaseAccountTableResultOutput) ToLookupDatabaseAccountTableResultOutputWithContext(ctx context.Context) LookupDatabaseAccountTableResultOutput {
	return o
}

// The unique resource identifier of the database account.
func (o LookupDatabaseAccountTableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountTableResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource group to which the resource belongs.
func (o LookupDatabaseAccountTableResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountTableResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the database account.
func (o LookupDatabaseAccountTableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountTableResult) string { return v.Name }).(pulumi.StringOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o LookupDatabaseAccountTableResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseAccountTableResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o LookupDatabaseAccountTableResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountTableResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseAccountTableResultOutput{})
}
