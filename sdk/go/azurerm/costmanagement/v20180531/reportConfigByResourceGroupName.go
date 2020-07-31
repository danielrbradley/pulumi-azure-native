// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180531

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A report config resource.
type ReportConfigByResourceGroupName struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the report config.
	Properties ReportConfigPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReportConfigByResourceGroupName registers a new resource with the given unique name, arguments, and options.
func NewReportConfigByResourceGroupName(ctx *pulumi.Context,
	name string, args *ReportConfigByResourceGroupNameArgs, opts ...pulumi.ResourceOption) (*ReportConfigByResourceGroupName, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ReportConfigByResourceGroupNameArgs{}
	}
	var resource ReportConfigByResourceGroupName
	err := ctx.RegisterResource("azurerm:costmanagement/v20180531:ReportConfigByResourceGroupName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReportConfigByResourceGroupName gets an existing ReportConfigByResourceGroupName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReportConfigByResourceGroupName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReportConfigByResourceGroupNameState, opts ...pulumi.ResourceOption) (*ReportConfigByResourceGroupName, error) {
	var resource ReportConfigByResourceGroupName
	err := ctx.ReadResource("azurerm:costmanagement/v20180531:ReportConfigByResourceGroupName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReportConfigByResourceGroupName resources.
type reportConfigByResourceGroupNameState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// The properties of the report config.
	Properties *ReportConfigPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type ReportConfigByResourceGroupNameState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// The properties of the report config.
	Properties ReportConfigPropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (ReportConfigByResourceGroupNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*reportConfigByResourceGroupNameState)(nil)).Elem()
}

type reportConfigByResourceGroupNameArgs struct {
	// Report Config Name.
	Name string `pulumi:"name"`
	// The properties of the report config.
	Properties *ReportConfigProperties `pulumi:"properties"`
	// Azure Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ReportConfigByResourceGroupName resource.
type ReportConfigByResourceGroupNameArgs struct {
	// Report Config Name.
	Name pulumi.StringInput
	// The properties of the report config.
	Properties ReportConfigPropertiesPtrInput
	// Azure Resource Group Name.
	ResourceGroupName pulumi.StringInput
}

func (ReportConfigByResourceGroupNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reportConfigByResourceGroupNameArgs)(nil)).Elem()
}