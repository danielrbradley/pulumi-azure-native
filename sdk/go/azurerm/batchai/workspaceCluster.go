// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package batchai

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Information about a Cluster.
type WorkspaceCluster struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties associated with the Cluster.
	Properties ClusterPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkspaceCluster registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceCluster(ctx *pulumi.Context,
	name string, args *WorkspaceClusterArgs, opts ...pulumi.ResourceOption) (*WorkspaceCluster, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &WorkspaceClusterArgs{}
	}
	var resource WorkspaceCluster
	err := ctx.RegisterResource("azurerm:batchai:WorkspaceCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceCluster gets an existing WorkspaceCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceClusterState, opts ...pulumi.ResourceOption) (*WorkspaceCluster, error) {
	var resource WorkspaceCluster
	err := ctx.ReadResource("azurerm:batchai:WorkspaceCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceCluster resources.
type workspaceClusterState struct {
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The properties associated with the Cluster.
	Properties *ClusterPropertiesResponse `pulumi:"properties"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type WorkspaceClusterState struct {
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The properties associated with the Cluster.
	Properties ClusterPropertiesResponsePtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (WorkspaceClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceClusterState)(nil)).Elem()
}

type workspaceClusterArgs struct {
	// The name of the cluster within the specified resource group. Cluster names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
	Name string `pulumi:"name"`
	// The properties of the Cluster.
	Properties *ClusterBaseProperties `pulumi:"properties"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace. Workspace names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a WorkspaceCluster resource.
type WorkspaceClusterArgs struct {
	// The name of the cluster within the specified resource group. Cluster names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
	Name pulumi.StringInput
	// The properties of the Cluster.
	Properties ClusterBasePropertiesPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace. Workspace names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
	WorkspaceName pulumi.StringInput
}

func (WorkspaceClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceClusterArgs)(nil)).Elem()
}