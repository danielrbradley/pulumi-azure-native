// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150501

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupWorkbook(ctx *pulumi.Context, args *LookupWorkbookArgs, opts ...pulumi.InvokeOption) (*LookupWorkbookResult, error) {
	var rv LookupWorkbookResult
	err := ctx.Invoke("azurerm:insights/v20150501:getWorkbook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkbookArgs struct {
	// The name of the Application Insights component resource.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Application Insights workbook definition.
type LookupWorkbookResult struct {
	// The kind of workbook. Choices are user and shared.
	Kind *string `pulumi:"kind"`
	// Resource location
	Location *string `pulumi:"location"`
	// Azure resource name
	Name string `pulumi:"name"`
	// Metadata describing a web test for an Azure resource.
	Properties WorkbookPropertiesResponse `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type
	Type string `pulumi:"type"`
}