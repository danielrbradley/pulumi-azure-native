// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180531

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupReportConfig(ctx *pulumi.Context, args *LookupReportConfigArgs, opts ...pulumi.InvokeOption) (*LookupReportConfigResult, error) {
	var rv LookupReportConfigResult
	err := ctx.Invoke("azurerm:costmanagement/v20180531:getReportConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReportConfigArgs struct {
	// Report Config Name.
	Name string `pulumi:"name"`
}

// A report config resource.
type LookupReportConfigResult struct {
	// Resource name.
	Name string `pulumi:"name"`
	// The properties of the report config.
	Properties ReportConfigPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}