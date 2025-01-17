// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A database blob auditing policy.
type DatabaseBlobAuditingPolicy struct {
	pulumi.CustomResourceState

	// Specifies the Actions and Actions-Groups to audit.
	AuditActionsAndGroups pulumi.StringArrayOutput `pulumi:"auditActionsAndGroups"`
	// Specifies whether storageAccountAccessKey value is the storage’s secondary key.
	IsStorageSecondaryKeyInUse pulumi.BoolPtrOutput `pulumi:"isStorageSecondaryKeyInUse"`
	// Resource kind.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the number of days to keep in the audit logs.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
	// Specifies the state of the policy. If state is Enabled, storageEndpoint and storageAccountAccessKey are required.
	State pulumi.StringOutput `pulumi:"state"`
	// Specifies the blob storage subscription Id.
	StorageAccountSubscriptionId pulumi.StringPtrOutput `pulumi:"storageAccountSubscriptionId"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). If state is Enabled, storageEndpoint is required.
	StorageEndpoint pulumi.StringPtrOutput `pulumi:"storageEndpoint"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatabaseBlobAuditingPolicy registers a new resource with the given unique name, arguments, and options.
func NewDatabaseBlobAuditingPolicy(ctx *pulumi.Context,
	name string, args *DatabaseBlobAuditingPolicyArgs, opts ...pulumi.ResourceOption) (*DatabaseBlobAuditingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:DatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:DatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:DatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:DatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:DatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:DatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:DatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:DatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:DatabaseBlobAuditingPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseBlobAuditingPolicy
	err := ctx.RegisterResource("azure-native:sql/v20150501preview:DatabaseBlobAuditingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseBlobAuditingPolicy gets an existing DatabaseBlobAuditingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseBlobAuditingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseBlobAuditingPolicyState, opts ...pulumi.ResourceOption) (*DatabaseBlobAuditingPolicy, error) {
	var resource DatabaseBlobAuditingPolicy
	err := ctx.ReadResource("azure-native:sql/v20150501preview:DatabaseBlobAuditingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseBlobAuditingPolicy resources.
type databaseBlobAuditingPolicyState struct {
}

type DatabaseBlobAuditingPolicyState struct {
}

func (DatabaseBlobAuditingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseBlobAuditingPolicyState)(nil)).Elem()
}

type databaseBlobAuditingPolicyArgs struct {
	// Specifies the Actions and Actions-Groups to audit.
	AuditActionsAndGroups []string `pulumi:"auditActionsAndGroups"`
	// The name of the blob auditing policy.
	BlobAuditingPolicyName *string `pulumi:"blobAuditingPolicyName"`
	// The name of the database for which the blob auditing policy will be defined.
	DatabaseName string `pulumi:"databaseName"`
	// Specifies whether storageAccountAccessKey value is the storage’s secondary key.
	IsStorageSecondaryKeyInUse *bool `pulumi:"isStorageSecondaryKeyInUse"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the number of days to keep in the audit logs.
	RetentionDays *int `pulumi:"retentionDays"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// Specifies the state of the policy. If state is Enabled, storageEndpoint and storageAccountAccessKey are required.
	State BlobAuditingPolicyState `pulumi:"state"`
	// Specifies the identifier key of the auditing storage account. If state is Enabled, storageAccountAccessKey is required.
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	// Specifies the blob storage subscription Id.
	StorageAccountSubscriptionId *string `pulumi:"storageAccountSubscriptionId"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). If state is Enabled, storageEndpoint is required.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
}

// The set of arguments for constructing a DatabaseBlobAuditingPolicy resource.
type DatabaseBlobAuditingPolicyArgs struct {
	// Specifies the Actions and Actions-Groups to audit.
	AuditActionsAndGroups pulumi.StringArrayInput
	// The name of the blob auditing policy.
	BlobAuditingPolicyName pulumi.StringPtrInput
	// The name of the database for which the blob auditing policy will be defined.
	DatabaseName pulumi.StringInput
	// Specifies whether storageAccountAccessKey value is the storage’s secondary key.
	IsStorageSecondaryKeyInUse pulumi.BoolPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Specifies the number of days to keep in the audit logs.
	RetentionDays pulumi.IntPtrInput
	// The name of the server.
	ServerName pulumi.StringInput
	// Specifies the state of the policy. If state is Enabled, storageEndpoint and storageAccountAccessKey are required.
	State BlobAuditingPolicyStateInput
	// Specifies the identifier key of the auditing storage account. If state is Enabled, storageAccountAccessKey is required.
	StorageAccountAccessKey pulumi.StringPtrInput
	// Specifies the blob storage subscription Id.
	StorageAccountSubscriptionId pulumi.StringPtrInput
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). If state is Enabled, storageEndpoint is required.
	StorageEndpoint pulumi.StringPtrInput
}

func (DatabaseBlobAuditingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseBlobAuditingPolicyArgs)(nil)).Elem()
}

type DatabaseBlobAuditingPolicyInput interface {
	pulumi.Input

	ToDatabaseBlobAuditingPolicyOutput() DatabaseBlobAuditingPolicyOutput
	ToDatabaseBlobAuditingPolicyOutputWithContext(ctx context.Context) DatabaseBlobAuditingPolicyOutput
}

func (*DatabaseBlobAuditingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseBlobAuditingPolicy)(nil)).Elem()
}

func (i *DatabaseBlobAuditingPolicy) ToDatabaseBlobAuditingPolicyOutput() DatabaseBlobAuditingPolicyOutput {
	return i.ToDatabaseBlobAuditingPolicyOutputWithContext(context.Background())
}

func (i *DatabaseBlobAuditingPolicy) ToDatabaseBlobAuditingPolicyOutputWithContext(ctx context.Context) DatabaseBlobAuditingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseBlobAuditingPolicyOutput)
}

type DatabaseBlobAuditingPolicyOutput struct{ *pulumi.OutputState }

func (DatabaseBlobAuditingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseBlobAuditingPolicy)(nil)).Elem()
}

func (o DatabaseBlobAuditingPolicyOutput) ToDatabaseBlobAuditingPolicyOutput() DatabaseBlobAuditingPolicyOutput {
	return o
}

func (o DatabaseBlobAuditingPolicyOutput) ToDatabaseBlobAuditingPolicyOutputWithContext(ctx context.Context) DatabaseBlobAuditingPolicyOutput {
	return o
}

// Specifies the Actions and Actions-Groups to audit.
func (o DatabaseBlobAuditingPolicyOutput) AuditActionsAndGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DatabaseBlobAuditingPolicy) pulumi.StringArrayOutput { return v.AuditActionsAndGroups }).(pulumi.StringArrayOutput)
}

// Specifies whether storageAccountAccessKey value is the storage’s secondary key.
func (o DatabaseBlobAuditingPolicyOutput) IsStorageSecondaryKeyInUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseBlobAuditingPolicy) pulumi.BoolPtrOutput { return v.IsStorageSecondaryKeyInUse }).(pulumi.BoolPtrOutput)
}

// Resource kind.
func (o DatabaseBlobAuditingPolicyOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseBlobAuditingPolicy) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource name.
func (o DatabaseBlobAuditingPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseBlobAuditingPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the number of days to keep in the audit logs.
func (o DatabaseBlobAuditingPolicyOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DatabaseBlobAuditingPolicy) pulumi.IntPtrOutput { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

// Specifies the state of the policy. If state is Enabled, storageEndpoint and storageAccountAccessKey are required.
func (o DatabaseBlobAuditingPolicyOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseBlobAuditingPolicy) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Specifies the blob storage subscription Id.
func (o DatabaseBlobAuditingPolicyOutput) StorageAccountSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseBlobAuditingPolicy) pulumi.StringPtrOutput { return v.StorageAccountSubscriptionId }).(pulumi.StringPtrOutput)
}

// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). If state is Enabled, storageEndpoint is required.
func (o DatabaseBlobAuditingPolicyOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseBlobAuditingPolicy) pulumi.StringPtrOutput { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o DatabaseBlobAuditingPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseBlobAuditingPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseBlobAuditingPolicyOutput{})
}
