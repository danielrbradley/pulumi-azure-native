// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160515

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A schedule.
//
// Deprecated: Version v20160515 will be removed in the next major version of the provider. Upgrade to version v20180915 or later.
func LookupSchedule(ctx *pulumi.Context, args *LookupScheduleArgs, opts ...pulumi.InvokeOption) (*LookupScheduleResult, error) {
	var rv LookupScheduleResult
	err := ctx.Invoke("azure-native:devtestlab/v20160515:getSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduleArgs struct {
	// Specify the $expand query. Example: 'properties($select=status)'
	Expand *string `pulumi:"expand"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The name of the schedule.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A schedule.
type LookupScheduleResult struct {
	// The creation date of the schedule.
	CreatedDate string `pulumi:"createdDate"`
	// If the schedule will occur once each day of the week, specify the daily recurrence.
	DailyRecurrence *DayDetailsResponse `pulumi:"dailyRecurrence"`
	// If the schedule will occur multiple times a day, specify the hourly recurrence.
	HourlyRecurrence *HourDetailsResponse `pulumi:"hourlyRecurrence"`
	// The identifier of the resource.
	Id string `pulumi:"id"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Notification settings.
	NotificationSettings *NotificationSettingsResponse `pulumi:"notificationSettings"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The status of the schedule (i.e. Enabled, Disabled)
	Status *string `pulumi:"status"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The resource ID to which the schedule belongs
	TargetResourceId *string `pulumi:"targetResourceId"`
	// The task type of the schedule (e.g. LabVmsShutdownTask, LabVmAutoStart).
	TaskType *string `pulumi:"taskType"`
	// The time zone ID (e.g. Pacific Standard time).
	TimeZoneId *string `pulumi:"timeZoneId"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
	// If the schedule will occur only some days of the week, specify the weekly recurrence.
	WeeklyRecurrence *WeekDetailsResponse `pulumi:"weeklyRecurrence"`
}

func LookupScheduleOutput(ctx *pulumi.Context, args LookupScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScheduleResult, error) {
			args := v.(LookupScheduleArgs)
			r, err := LookupSchedule(ctx, &args, opts...)
			var s LookupScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScheduleResultOutput)
}

type LookupScheduleOutputArgs struct {
	// Specify the $expand query. Example: 'properties($select=status)'
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the lab.
	LabName pulumi.StringInput `pulumi:"labName"`
	// The name of the schedule.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduleArgs)(nil)).Elem()
}

// A schedule.
type LookupScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduleResult)(nil)).Elem()
}

func (o LookupScheduleResultOutput) ToLookupScheduleResultOutput() LookupScheduleResultOutput {
	return o
}

func (o LookupScheduleResultOutput) ToLookupScheduleResultOutputWithContext(ctx context.Context) LookupScheduleResultOutput {
	return o
}

// The creation date of the schedule.
func (o LookupScheduleResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

// If the schedule will occur once each day of the week, specify the daily recurrence.
func (o LookupScheduleResultOutput) DailyRecurrence() DayDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *DayDetailsResponse { return v.DailyRecurrence }).(DayDetailsResponsePtrOutput)
}

// If the schedule will occur multiple times a day, specify the hourly recurrence.
func (o LookupScheduleResultOutput) HourlyRecurrence() HourDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *HourDetailsResponse { return v.HourlyRecurrence }).(HourDetailsResponsePtrOutput)
}

// The identifier of the resource.
func (o LookupScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource.
func (o LookupScheduleResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Notification settings.
func (o LookupScheduleResultOutput) NotificationSettings() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *NotificationSettingsResponse { return v.NotificationSettings }).(NotificationSettingsResponsePtrOutput)
}

// The provisioning status of the resource.
func (o LookupScheduleResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The status of the schedule (i.e. Enabled, Disabled)
func (o LookupScheduleResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The tags of the resource.
func (o LookupScheduleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupScheduleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource ID to which the schedule belongs
func (o LookupScheduleResultOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

// The task type of the schedule (e.g. LabVmsShutdownTask, LabVmAutoStart).
func (o LookupScheduleResultOutput) TaskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.TaskType }).(pulumi.StringPtrOutput)
}

// The time zone ID (e.g. Pacific Standard time).
func (o LookupScheduleResultOutput) TimeZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.TimeZoneId }).(pulumi.StringPtrOutput)
}

// The type of the resource.
func (o LookupScheduleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Type }).(pulumi.StringOutput)
}

// The unique immutable identifier of a resource (Guid).
func (o LookupScheduleResultOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

// If the schedule will occur only some days of the week, specify the weekly recurrence.
func (o LookupScheduleResultOutput) WeeklyRecurrence() WeekDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *WeekDetailsResponse { return v.WeeklyRecurrence }).(WeekDetailsResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduleResultOutput{})
}
