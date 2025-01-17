// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20171230preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A budget resource.
//
// Deprecated: Version v20171230preview will be removed in the next major version of the provider. Upgrade to version v20191001 or later.
func LookupBudget(ctx *pulumi.Context, args *LookupBudgetArgs, opts ...pulumi.InvokeOption) (*LookupBudgetResult, error) {
	var rv LookupBudgetResult
	err := ctx.Invoke("azure-native:consumption/v20171230preview:getBudget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBudgetArgs struct {
	// Budget name.
	Name string `pulumi:"name"`
}

// A budget resource.
type LookupBudgetResult struct {
	// The total amount of cost to track with the budget
	Amount float64 `pulumi:"amount"`
	// The category of the budget, whether the budget tracks cost or something else.
	Category string `pulumi:"category"`
	// The current amount of cost which is being tracked for a budget.
	CurrentSpend CurrentSpendResponse `pulumi:"currentSpend"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag *string `pulumi:"eTag"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Dictionary of notifications associated with the budget. Budget can have up to five notifications.
	Notifications map[string]NotificationResponse `pulumi:"notifications"`
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain.
	TimeGrain string `pulumi:"timeGrain"`
	// Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than three months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
	TimePeriod BudgetTimePeriodResponse `pulumi:"timePeriod"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupBudgetOutput(ctx *pulumi.Context, args LookupBudgetOutputArgs, opts ...pulumi.InvokeOption) LookupBudgetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBudgetResult, error) {
			args := v.(LookupBudgetArgs)
			r, err := LookupBudget(ctx, &args, opts...)
			var s LookupBudgetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBudgetResultOutput)
}

type LookupBudgetOutputArgs struct {
	// Budget name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupBudgetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBudgetArgs)(nil)).Elem()
}

// A budget resource.
type LookupBudgetResultOutput struct{ *pulumi.OutputState }

func (LookupBudgetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBudgetResult)(nil)).Elem()
}

func (o LookupBudgetResultOutput) ToLookupBudgetResultOutput() LookupBudgetResultOutput {
	return o
}

func (o LookupBudgetResultOutput) ToLookupBudgetResultOutputWithContext(ctx context.Context) LookupBudgetResultOutput {
	return o
}

// The total amount of cost to track with the budget
func (o LookupBudgetResultOutput) Amount() pulumi.Float64Output {
	return o.ApplyT(func(v LookupBudgetResult) float64 { return v.Amount }).(pulumi.Float64Output)
}

// The category of the budget, whether the budget tracks cost or something else.
func (o LookupBudgetResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.Category }).(pulumi.StringOutput)
}

// The current amount of cost which is being tracked for a budget.
func (o LookupBudgetResultOutput) CurrentSpend() CurrentSpendResponseOutput {
	return o.ApplyT(func(v LookupBudgetResult) CurrentSpendResponse { return v.CurrentSpend }).(CurrentSpendResponseOutput)
}

// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
func (o LookupBudgetResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBudgetResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupBudgetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupBudgetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.Name }).(pulumi.StringOutput)
}

// Dictionary of notifications associated with the budget. Budget can have up to five notifications.
func (o LookupBudgetResultOutput) Notifications() NotificationResponseMapOutput {
	return o.ApplyT(func(v LookupBudgetResult) map[string]NotificationResponse { return v.Notifications }).(NotificationResponseMapOutput)
}

// The time covered by a budget. Tracking of the amount will be reset based on the time grain.
func (o LookupBudgetResultOutput) TimeGrain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.TimeGrain }).(pulumi.StringOutput)
}

// Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than three months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
func (o LookupBudgetResultOutput) TimePeriod() BudgetTimePeriodResponseOutput {
	return o.ApplyT(func(v LookupBudgetResult) BudgetTimePeriodResponse { return v.TimePeriod }).(BudgetTimePeriodResponseOutput)
}

// Resource type.
func (o LookupBudgetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBudgetResultOutput{})
}
