// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Policy for sharing applications on this compute instance among users of parent workspace. If Personal, only the creator can access applications on this compute instance. When Shared, any workspace user can access applications on this instance depending on his/her assigned role.
type ApplicationSharingPolicy string

const (
	ApplicationSharingPolicyPersonal = ApplicationSharingPolicy("Personal")
	ApplicationSharingPolicyShared   = ApplicationSharingPolicy("Shared")
)

// Logging level for batch inference operation.
type BatchLoggingLevel string

const (
	BatchLoggingLevelInfo    = BatchLoggingLevel("Info")
	BatchLoggingLevelWarning = BatchLoggingLevel("Warning")
	BatchLoggingLevelDebug   = BatchLoggingLevel("Debug")
)

// Indicates how the output will be organized.
type BatchOutputAction string

const (
	BatchOutputActionSummaryOnly = BatchOutputAction("SummaryOnly")
	BatchOutputActionAppendRow   = BatchOutputAction("AppendRow")
)

// Intended usage of the cluster
type ClusterPurpose string

const (
	ClusterPurposeFastProd  = ClusterPurpose("FastProd")
	ClusterPurposeDenseProd = ClusterPurpose("DenseProd")
	ClusterPurposeDevTest   = ClusterPurpose("DevTest")
)

// The Compute Instance Authorization type. Available values are personal (default).
type ComputeInstanceAuthorizationType string

const (
	ComputeInstanceAuthorizationTypePersonal = ComputeInstanceAuthorizationType("personal")
)

// The compute power action.
type ComputePowerAction string

const (
	ComputePowerActionStart = ComputePowerAction("Start")
	ComputePowerActionStop  = ComputePowerAction("Stop")
)

// The type of compute
type ComputeType string

const (
	ComputeTypeAKS               = ComputeType("AKS")
	ComputeTypeAmlCompute        = ComputeType("AmlCompute")
	ComputeTypeComputeInstance   = ComputeType("ComputeInstance")
	ComputeTypeDataFactory       = ComputeType("DataFactory")
	ComputeTypeVirtualMachine    = ComputeType("VirtualMachine")
	ComputeTypeHDInsight         = ComputeType("HDInsight")
	ComputeTypeDatabricks        = ComputeType("Databricks")
	ComputeTypeDataLakeAnalytics = ComputeType("DataLakeAnalytics")
	ComputeTypeSynapseSpark      = ComputeType("SynapseSpark")
)

// The type of container to retrieve logs from.
type ContainerType string

const (
	ContainerTypeStorageInitializer = ContainerType("StorageInitializer")
	ContainerTypeInferenceServer    = ContainerType("InferenceServer")
)

// [Required] Storage type backing the datastore.
type ContentsType string

const (
	ContentsTypeAzureBlob         = ContentsType("AzureBlob")
	ContentsTypeAzureDataLakeGen1 = ContentsType("AzureDataLakeGen1")
	ContentsTypeAzureDataLakeGen2 = ContentsType("AzureDataLakeGen2")
	ContentsTypeAzureFile         = ContentsType("AzureFile")
	ContentsTypeAzureMySql        = ContentsType("AzureMySql")
	ContentsTypeAzurePostgreSql   = ContentsType("AzurePostgreSql")
	ContentsTypeAzureSqlDatabase  = ContentsType("AzureSqlDatabase")
	ContentsTypeGlusterFs         = ContentsType("GlusterFs")
)

// [Required] Credential type used to authentication with storage.
type CredentialsType string

const (
	CredentialsTypeAccountKey       = CredentialsType("AccountKey")
	CredentialsTypeCertificate      = CredentialsType("Certificate")
	CredentialsTypeNone             = CredentialsType("None")
	CredentialsTypeSas              = CredentialsType("Sas")
	CredentialsTypeServicePrincipal = CredentialsType("ServicePrincipal")
	CredentialsTypeSqlAdmin         = CredentialsType("SqlAdmin")
)

// Mechanism for data movement to datastore.
type DataBindingMode string

const (
	DataBindingModeMount          = DataBindingMode("Mount")
	DataBindingModeDownload       = DataBindingMode("Download")
	DataBindingModeUpload         = DataBindingMode("Upload")
	DataBindingModeReadOnlyMount  = DataBindingMode("ReadOnlyMount")
	DataBindingModeReadWriteMount = DataBindingMode("ReadWriteMount")
	DataBindingModeDirect         = DataBindingMode("Direct")
	DataBindingModeEvalMount      = DataBindingMode("EvalMount")
	DataBindingModeEvalDownload   = DataBindingMode("EvalDownload")
)

// The Format of dataset.
type DatasetType string

const (
	DatasetTypeSimple   = DatasetType("Simple")
	DatasetTypeDataflow = DatasetType("Dataflow")
)

type DaysOfWeek string

const (
	DaysOfWeekSunday    = DaysOfWeek("Sunday")
	DaysOfWeekMonday    = DaysOfWeek("Monday")
	DaysOfWeekTuesday   = DaysOfWeek("Tuesday")
	DaysOfWeekWednesday = DaysOfWeek("Wednesday")
	DaysOfWeekThursday  = DaysOfWeek("Thursday")
	DaysOfWeekFriday    = DaysOfWeek("Friday")
	DaysOfWeekSaturday  = DaysOfWeek("Saturday")
)

// [Required] Specifies the type of distribution framework.
type DistributionType string

const (
	DistributionTypePyTorch    = DistributionType("PyTorch")
	DistributionTypeTensorFlow = DistributionType("TensorFlow")
	DistributionTypeMpi        = DistributionType("Mpi")
)

// [Required] Docker specification must be either Build or Image
type DockerSpecificationType string

const (
	DockerSpecificationTypeBuild = DockerSpecificationType("Build")
	DockerSpecificationTypeImage = DockerSpecificationType("Image")
)

// [Required] Name of policy configuration
type EarlyTerminationPolicyType string

const (
	EarlyTerminationPolicyTypeBandit              = EarlyTerminationPolicyType("Bandit")
	EarlyTerminationPolicyTypeMedianStopping      = EarlyTerminationPolicyType("MedianStopping")
	EarlyTerminationPolicyTypeTruncationSelection = EarlyTerminationPolicyType("TruncationSelection")
)

// Indicates whether or not the encryption is enabled for the workspace.
type EncryptionStatus string

const (
	EncryptionStatusEnabled  = EncryptionStatus("Enabled")
	EncryptionStatusDisabled = EncryptionStatus("Disabled")
)

// [Required] Inference endpoint authentication mode type
type EndpointAuthMode string

const (
	EndpointAuthModeAMLToken = EndpointAuthMode("AMLToken")
	EndpointAuthModeKey      = EndpointAuthMode("Key")
	EndpointAuthModeAADToken = EndpointAuthMode("AADToken")
)

// [Required] The compute type of the endpoint.
type EndpointComputeType string

const (
	EndpointComputeTypeManaged        = EndpointComputeType("Managed")
	EndpointComputeTypeK8S            = EndpointComputeType("K8S")
	EndpointComputeTypeAzureMLCompute = EndpointComputeType("AzureMLCompute")
)

// [Required] Defines supported metric goals for hyperparameter tuning
type Goal string

const (
	GoalMinimize = Goal("Minimize")
	GoalMaximize = Goal("Maximize")
)

// [Required] Specifies the type of identity framework.
type IdentityConfigurationType string

const (
	IdentityConfigurationTypeManaged  = IdentityConfigurationType("Managed")
	IdentityConfigurationTypeAMLToken = IdentityConfigurationType("AMLToken")
)

// Annotation type of image labeling job.
type ImageAnnotationType string

const (
	ImageAnnotationTypeClassification       = ImageAnnotationType("Classification")
	ImageAnnotationTypeBoundingBox          = ImageAnnotationType("BoundingBox")
	ImageAnnotationTypeInstanceSegmentation = ImageAnnotationType("InstanceSegmentation")
)

// [Required] Specifies the type of job. This field should always be set to "Labeling".
type JobType string

const (
	JobTypeCommand  = JobType("Command")
	JobTypeSweep    = JobType("Sweep")
	JobTypeLabeling = JobType("Labeling")
)

// Load Balancer Type
type LoadBalancerType string

const (
	LoadBalancerTypePublicIp             = LoadBalancerType("PublicIp")
	LoadBalancerTypeInternalLoadBalancer = LoadBalancerType("InternalLoadBalancer")
)

// [Required] Media type of the job.
type MediaType string

const (
	MediaTypeImage = MediaType("Image")
	MediaTypeText  = MediaType("Text")
)

// The OS type the Environment.
type OperatingSystemType string

const (
	OperatingSystemTypeLinux   = OperatingSystemType("Linux")
	OperatingSystemTypeWindows = OperatingSystemType("Windows")
)

// Type of the linked service.
type OriginType string

const (
	OriginTypeSynapse = OriginType("Synapse")
)

// Compute OS Type
type OsType string

const (
	OsTypeLinux   = OsType("Linux")
	OsTypeWindows = OsType("Windows")
)

// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending      = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved     = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected     = PrivateEndpointServiceConnectionStatus("Rejected")
	PrivateEndpointServiceConnectionStatusDisconnected = PrivateEndpointServiceConnectionStatus("Disconnected")
	PrivateEndpointServiceConnectionStatusTimeout      = PrivateEndpointServiceConnectionStatus("Timeout")
)

// The recurrence frequency.
type RecurrenceFrequency string

const (
	RecurrenceFrequencyNotSpecified = RecurrenceFrequency("NotSpecified")
	RecurrenceFrequencySecond       = RecurrenceFrequency("Second")
	RecurrenceFrequencyMinute       = RecurrenceFrequency("Minute")
	RecurrenceFrequencyHour         = RecurrenceFrequency("Hour")
	RecurrenceFrequencyDay          = RecurrenceFrequency("Day")
	RecurrenceFrequencyWeek         = RecurrenceFrequency("Week")
	RecurrenceFrequencyMonth        = RecurrenceFrequency("Month")
	RecurrenceFrequencyYear         = RecurrenceFrequency("Year")
)

// [Required] Specifies the type of asset reference.
type ReferenceType string

const (
	ReferenceTypeId         = ReferenceType("Id")
	ReferenceTypeDataPath   = ReferenceType("DataPath")
	ReferenceTypeOutputPath = ReferenceType("OutputPath")
)

// State of the public SSH port. Possible values are: Disabled - Indicates that the public ssh port is closed on all nodes of the cluster. Enabled - Indicates that the public ssh port is open on all nodes of the cluster. NotSpecified - Indicates that the public ssh port is closed on all nodes of the cluster if VNet is defined, else is open all public nodes. It can be default only during cluster creation time, after creation it will be either enabled or disabled.
type RemoteLoginPortPublicAccess string

const (
	RemoteLoginPortPublicAccessEnabled      = RemoteLoginPortPublicAccess("Enabled")
	RemoteLoginPortPublicAccessDisabled     = RemoteLoginPortPublicAccess("Disabled")
	RemoteLoginPortPublicAccessNotSpecified = RemoteLoginPortPublicAccess("NotSpecified")
)

// Defines values for a ResourceIdentity's type.
type ResourceIdentityAssignment string

const (
	ResourceIdentityAssignmentSystemAssigned               = ResourceIdentityAssignment("SystemAssigned")
	ResourceIdentityAssignmentUserAssigned                 = ResourceIdentityAssignment("UserAssigned")
	ResourceIdentityAssignment_SystemAssigned_UserAssigned = ResourceIdentityAssignment("SystemAssigned,UserAssigned")
	ResourceIdentityAssignmentNone                         = ResourceIdentityAssignment("None")
)

// The identity type.
type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned,UserAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ResourceIdentityTypeInput is an input type that accepts ResourceIdentityTypeArgs and ResourceIdentityTypeOutput values.
// You can construct a concrete instance of `ResourceIdentityTypeInput` via:
//
//          ResourceIdentityTypeArgs{...}
type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

// [Required] Type of the hyperparameter sampling algorithms
type SamplingAlgorithm string

const (
	SamplingAlgorithmGrid     = SamplingAlgorithm("Grid")
	SamplingAlgorithmRandom   = SamplingAlgorithm("Random")
	SamplingAlgorithmBayesian = SamplingAlgorithm("Bayesian")
)

// [Required] Type of deployment scaling algorithm
type ScaleType string

const (
	ScaleTypeAuto   = ScaleType("Auto")
	ScaleTypeManual = ScaleType("Manual")
)

// The schedule status.
type ScheduleStatus string

const (
	ScheduleStatusEnabled  = ScheduleStatus("Enabled")
	ScheduleStatusDisabled = ScheduleStatus("Disabled")
)

// [Required] Credential type used to authentication with storage.
type SecretsType string

const (
	SecretsTypeAccountKey       = SecretsType("AccountKey")
	SecretsTypeCertificate      = SecretsType("Certificate")
	SecretsTypeNone             = SecretsType("None")
	SecretsTypeSas              = SecretsType("Sas")
	SecretsTypeServicePrincipal = SecretsType("ServicePrincipal")
	SecretsTypeSqlAdmin         = SecretsType("SqlAdmin")
)

// State of the public SSH port. Possible values are: Disabled - Indicates that the public ssh port is closed on this instance. Enabled - Indicates that the public ssh port is open and accessible according to the VNet/subnet policy if applicable.
type SshPublicAccess string

const (
	SshPublicAccessEnabled  = SshPublicAccess("Enabled")
	SshPublicAccessDisabled = SshPublicAccess("Disabled")
)

// Annotation type of text labeling job.
type TextAnnotationType string

const (
	TextAnnotationTypeClassification = TextAnnotationType("Classification")
)

// The schedule trigger type.
type TriggerType string

const (
	TriggerTypeRecurrence = TriggerType("Recurrence")
	TriggerTypeCron       = TriggerType("Cron")
)

// format for the workspace connection value
type ValueFormat string

const (
	ValueFormatJSON = ValueFormat("JSON")
)

// Virtual Machine priority
type VmPriority string

const (
	VmPriorityDedicated   = VmPriority("Dedicated")
	VmPriorityLowPriority = VmPriority("LowPriority")
)

func init() {
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}
