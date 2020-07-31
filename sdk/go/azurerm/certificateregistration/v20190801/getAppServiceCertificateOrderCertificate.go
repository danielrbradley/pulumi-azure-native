// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupAppServiceCertificateOrderCertificate(ctx *pulumi.Context, args *LookupAppServiceCertificateOrderCertificateArgs, opts ...pulumi.InvokeOption) (*LookupAppServiceCertificateOrderCertificateResult, error) {
	var rv LookupAppServiceCertificateOrderCertificateResult
	err := ctx.Invoke("azurerm:certificateregistration/v20190801:getAppServiceCertificateOrderCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppServiceCertificateOrderCertificateArgs struct {
	// Name of the certificate order.
	CertificateOrderName string `pulumi:"certificateOrderName"`
	// Name of the certificate.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Key Vault container ARM resource for a certificate that is purchased through Azure.
type LookupAppServiceCertificateOrderCertificateResult struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location string `pulumi:"location"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Core resource properties
	Properties AppServiceCertificateResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}