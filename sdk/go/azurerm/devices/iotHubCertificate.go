// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devices

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The X509 Certificate.
type IotHubCertificate struct {
	pulumi.CustomResourceState

	// The entity tag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the certificate.
	Name pulumi.StringOutput `pulumi:"name"`
	// The description of an X509 CA Certificate.
	Properties CertificatePropertiesResponseOutput `pulumi:"properties"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIotHubCertificate registers a new resource with the given unique name, arguments, and options.
func NewIotHubCertificate(ctx *pulumi.Context,
	name string, args *IotHubCertificateArgs, opts ...pulumi.ResourceOption) (*IotHubCertificate, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &IotHubCertificateArgs{}
	}
	var resource IotHubCertificate
	err := ctx.RegisterResource("azurerm:devices:IotHubCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIotHubCertificate gets an existing IotHubCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIotHubCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotHubCertificateState, opts ...pulumi.ResourceOption) (*IotHubCertificate, error) {
	var resource IotHubCertificate
	err := ctx.ReadResource("azurerm:devices:IotHubCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IotHubCertificate resources.
type iotHubCertificateState struct {
	// The entity tag.
	Etag *string `pulumi:"etag"`
	// The name of the certificate.
	Name *string `pulumi:"name"`
	// The description of an X509 CA Certificate.
	Properties *CertificatePropertiesResponse `pulumi:"properties"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type IotHubCertificateState struct {
	// The entity tag.
	Etag pulumi.StringPtrInput
	// The name of the certificate.
	Name pulumi.StringPtrInput
	// The description of an X509 CA Certificate.
	Properties CertificatePropertiesResponsePtrInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (IotHubCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubCertificateState)(nil)).Elem()
}

type iotHubCertificateArgs struct {
	// base-64 representation of the X509 leaf certificate .cer file or just .pem file content.
	Certificate *string `pulumi:"certificate"`
	// The name of the certificate
	Name string `pulumi:"name"`
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the IoT hub.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a IotHubCertificate resource.
type IotHubCertificateArgs struct {
	// base-64 representation of the X509 leaf certificate .cer file or just .pem file content.
	Certificate pulumi.StringPtrInput
	// The name of the certificate
	Name pulumi.StringInput
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName pulumi.StringInput
	// The name of the IoT hub.
	ResourceName pulumi.StringInput
}

func (IotHubCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubCertificateArgs)(nil)).Elem()
}