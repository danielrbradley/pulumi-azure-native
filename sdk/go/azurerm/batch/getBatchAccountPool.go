// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupBatchAccountPool(ctx *pulumi.Context, args *LookupBatchAccountPoolArgs, opts ...pulumi.InvokeOption) (*LookupBatchAccountPoolResult, error) {
	var rv LookupBatchAccountPoolResult
	err := ctx.Invoke("azurerm:batch:getBatchAccountPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBatchAccountPoolArgs struct {
	// The name of the Batch account.
	AccountName string `pulumi:"accountName"`
	// The pool name. This must be unique within the account.
	Name string `pulumi:"name"`
	// The name of the resource group that contains the Batch account.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Contains information about a pool.
type LookupBatchAccountPoolResult struct {
	// The ETag of the resource, used for concurrency statements.
	Etag string `pulumi:"etag"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The properties associated with the pool.
	Properties PoolPropertiesResponse `pulumi:"properties"`
	// The type of the resource.
	Type string `pulumi:"type"`
}