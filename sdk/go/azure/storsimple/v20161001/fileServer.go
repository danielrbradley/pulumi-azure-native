// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20161001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The file server.
//
// Deprecated: Version v20161001 will be removed in the next major version of the provider. Upgrade to version v20170601 or later.
type FileServer struct {
	pulumi.CustomResourceState

	// The backup policy id.
	BackupScheduleGroupId pulumi.StringOutput `pulumi:"backupScheduleGroupId"`
	// The description of the file server
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Domain of the file server
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The storage domain id.
	StorageDomainId pulumi.StringOutput `pulumi:"storageDomainId"`
	// The type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFileServer registers a new resource with the given unique name, arguments, and options.
func NewFileServer(ctx *pulumi.Context,
	name string, args *FileServerArgs, opts ...pulumi.ResourceOption) (*FileServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupScheduleGroupId == nil {
		return nil, errors.New("invalid value for required argument 'BackupScheduleGroupId'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageDomainId == nil {
		return nil, errors.New("invalid value for required argument 'StorageDomainId'")
	}
	var resource FileServer
	err := ctx.RegisterResource("azure-native:storsimple/v20161001:FileServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFileServer gets an existing FileServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileServerState, opts ...pulumi.ResourceOption) (*FileServer, error) {
	var resource FileServer
	err := ctx.ReadResource("azure-native:storsimple/v20161001:FileServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FileServer resources.
type fileServerState struct {
}

type FileServerState struct {
}

func (FileServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileServerState)(nil)).Elem()
}

type fileServerArgs struct {
	// The backup policy id.
	BackupScheduleGroupId string `pulumi:"backupScheduleGroupId"`
	// The description of the file server
	Description *string `pulumi:"description"`
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// Domain of the file server
	DomainName string `pulumi:"domainName"`
	// The file server name.
	FileServerName *string `pulumi:"fileServerName"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The storage domain id.
	StorageDomainId string `pulumi:"storageDomainId"`
}

// The set of arguments for constructing a FileServer resource.
type FileServerArgs struct {
	// The backup policy id.
	BackupScheduleGroupId pulumi.StringInput
	// The description of the file server
	Description pulumi.StringPtrInput
	// The device name.
	DeviceName pulumi.StringInput
	// Domain of the file server
	DomainName pulumi.StringInput
	// The file server name.
	FileServerName pulumi.StringPtrInput
	// The manager name
	ManagerName pulumi.StringInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
	// The storage domain id.
	StorageDomainId pulumi.StringInput
}

func (FileServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileServerArgs)(nil)).Elem()
}

type FileServerInput interface {
	pulumi.Input

	ToFileServerOutput() FileServerOutput
	ToFileServerOutputWithContext(ctx context.Context) FileServerOutput
}

func (*FileServer) ElementType() reflect.Type {
	return reflect.TypeOf((**FileServer)(nil)).Elem()
}

func (i *FileServer) ToFileServerOutput() FileServerOutput {
	return i.ToFileServerOutputWithContext(context.Background())
}

func (i *FileServer) ToFileServerOutputWithContext(ctx context.Context) FileServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileServerOutput)
}

type FileServerOutput struct{ *pulumi.OutputState }

func (FileServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileServer)(nil)).Elem()
}

func (o FileServerOutput) ToFileServerOutput() FileServerOutput {
	return o
}

func (o FileServerOutput) ToFileServerOutputWithContext(ctx context.Context) FileServerOutput {
	return o
}

// The backup policy id.
func (o FileServerOutput) BackupScheduleGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *FileServer) pulumi.StringOutput { return v.BackupScheduleGroupId }).(pulumi.StringOutput)
}

// The description of the file server
func (o FileServerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileServer) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Domain of the file server
func (o FileServerOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *FileServer) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// The name.
func (o FileServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FileServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The storage domain id.
func (o FileServerOutput) StorageDomainId() pulumi.StringOutput {
	return o.ApplyT(func(v *FileServer) pulumi.StringOutput { return v.StorageDomainId }).(pulumi.StringOutput)
}

// The type.
func (o FileServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FileServer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FileServerOutput{})
}
