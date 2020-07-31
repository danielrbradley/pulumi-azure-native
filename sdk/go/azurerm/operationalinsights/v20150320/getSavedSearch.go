// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150320

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupSavedSearch(ctx *pulumi.Context, args *LookupSavedSearchArgs, opts ...pulumi.InvokeOption) (*LookupSavedSearchResult, error) {
	var rv LookupSavedSearchResult
	err := ctx.Invoke("azurerm:operationalinsights/v20150320:getSavedSearch", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSavedSearchArgs struct {
	// The id of the saved search.
	Name string `pulumi:"name"`
	// The Resource Group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Log Analytics Workspace name.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Value object for saved search results.
type LookupSavedSearchResult struct {
	// The ETag of the saved search.
	ETag *string `pulumi:"eTag"`
	// The name of the saved search.
	Name string `pulumi:"name"`
	// The properties of the saved search.
	Properties SavedSearchPropertiesResponse `pulumi:"properties"`
	// The type of the saved search.
	Type string `pulumi:"type"`
}