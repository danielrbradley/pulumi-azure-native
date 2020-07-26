// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupAccountShareDataSet(ctx *pulumi.Context, args *LookupAccountShareDataSetArgs, opts ...pulumi.InvokeOption) (*LookupAccountShareDataSetResult, error) {
	var rv LookupAccountShareDataSetResult
	err := ctx.Invoke("azurerm:datashare:getAccountShareDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountShareDataSetArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSet.
	Name string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share.
	ShareName string `pulumi:"shareName"`
}

// A DataSet data transfer object.
type LookupAccountShareDataSetResult struct {
	// Kind of data set.
	Kind string `pulumi:"kind"`
	// Name of the azure resource
	Name string `pulumi:"name"`
	// Type of the azure resource
	Type string `pulumi:"type"`
}