// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListAccountShareSynchronizations(ctx *pulumi.Context, args *ListAccountShareSynchronizationsArgs, opts ...pulumi.InvokeOption) (*ListAccountShareSynchronizationsResult, error) {
	var rv ListAccountShareSynchronizationsResult
	err := ctx.Invoke("azurerm:datashare:listAccountShareSynchronizations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAccountShareSynchronizationsArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share.
	ShareName string `pulumi:"shareName"`
}

// List response for get ShareSynchronization.
type ListAccountShareSynchronizationsResult struct {
	// The Url of next result page.
	NextLink *string `pulumi:"nextLink"`
	// Collection of items of type DataTransferObjects.
	Value []ShareSynchronizationResponse `pulumi:"value"`
}