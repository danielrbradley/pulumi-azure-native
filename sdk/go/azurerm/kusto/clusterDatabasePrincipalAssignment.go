// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Class representing a database principal assignment.
type ClusterDatabasePrincipalAssignment struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The database principal.
	Properties DatabasePrincipalPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewClusterDatabasePrincipalAssignment registers a new resource with the given unique name, arguments, and options.
func NewClusterDatabasePrincipalAssignment(ctx *pulumi.Context,
	name string, args *ClusterDatabasePrincipalAssignmentArgs, opts ...pulumi.ResourceOption) (*ClusterDatabasePrincipalAssignment, error) {
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.DatabaseName == nil {
		return nil, errors.New("missing required argument 'DatabaseName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ClusterDatabasePrincipalAssignmentArgs{}
	}
	var resource ClusterDatabasePrincipalAssignment
	err := ctx.RegisterResource("azurerm:kusto:ClusterDatabasePrincipalAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterDatabasePrincipalAssignment gets an existing ClusterDatabasePrincipalAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterDatabasePrincipalAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterDatabasePrincipalAssignmentState, opts ...pulumi.ResourceOption) (*ClusterDatabasePrincipalAssignment, error) {
	var resource ClusterDatabasePrincipalAssignment
	err := ctx.ReadResource("azurerm:kusto:ClusterDatabasePrincipalAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterDatabasePrincipalAssignment resources.
type clusterDatabasePrincipalAssignmentState struct {
	// The name of the resource
	Name *string `pulumi:"name"`
	// The database principal.
	Properties *DatabasePrincipalPropertiesResponse `pulumi:"properties"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type ClusterDatabasePrincipalAssignmentState struct {
	// The name of the resource
	Name pulumi.StringPtrInput
	// The database principal.
	Properties DatabasePrincipalPropertiesResponsePtrInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (ClusterDatabasePrincipalAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterDatabasePrincipalAssignmentState)(nil)).Elem()
}

type clusterDatabasePrincipalAssignmentArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the database in the Kusto cluster.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the Kusto principalAssignment.
	Name string `pulumi:"name"`
	// The database principal.
	Properties *DatabasePrincipalProperties `pulumi:"properties"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ClusterDatabasePrincipalAssignment resource.
type ClusterDatabasePrincipalAssignmentArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput
	// The name of the database in the Kusto cluster.
	DatabaseName pulumi.StringInput
	// The name of the Kusto principalAssignment.
	Name pulumi.StringInput
	// The database principal.
	Properties DatabasePrincipalPropertiesPtrInput
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput
}

func (ClusterDatabasePrincipalAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterDatabasePrincipalAssignmentArgs)(nil)).Elem()
}