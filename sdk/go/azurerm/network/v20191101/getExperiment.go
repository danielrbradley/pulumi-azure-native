// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupExperiment(ctx *pulumi.Context, args *LookupExperimentArgs, opts ...pulumi.InvokeOption) (*LookupExperimentResult, error) {
	var rv LookupExperimentResult
	err := ctx.Invoke("azurerm:network/v20191101:getExperiment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExperimentArgs struct {
	// The Experiment identifier associated with the Experiment
	Name string `pulumi:"name"`
	// The Profile identifier associated with the Tenant and Partner
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Defines the properties of an Experiment
type LookupExperimentResult struct {
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The properties of an Experiment
	Properties ExperimentPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}