// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicequotas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve information about a Service Quota.
//
// > **NOTE:** Global quotas apply to all AWS regions, but can only be accessed in `us-east-1` in the Commercial partition or `us-gov-west-1` in the GovCloud partition. In other regions, the AWS API will return the error `The request failed because the specified service does not exist.`
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicequotas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := servicequotas.LookupServiceQuota(ctx, &servicequotas.LookupServiceQuotaArgs{
//				QuotaCode:   pulumi.StringRef("L-F678F1CE"),
//				ServiceCode: "vpc",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = servicequotas.LookupServiceQuota(ctx, &servicequotas.LookupServiceQuotaArgs{
//				QuotaName:   pulumi.StringRef("VPCs per Region"),
//				ServiceCode: "vpc",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupServiceQuota(ctx *pulumi.Context, args *LookupServiceQuotaArgs, opts ...pulumi.InvokeOption) (*LookupServiceQuotaResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceQuotaResult
	err := ctx.Invoke("aws:servicequotas/getServiceQuota:getServiceQuota", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceQuota.
type LookupServiceQuotaArgs struct {
	// Quota code within the service. When configured, the data source directly looks up the service quota. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html). One of `quotaCode` or `quotaName` must be specified.
	QuotaCode *string `pulumi:"quotaCode"`
	// Quota name within the service. When configured, the data source searches through all service quotas to find the matching quota name. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html). One of `quotaName` or `quotaCode` must be specified.
	QuotaName *string `pulumi:"quotaName"`
	// Service code for the quota. Available values can be found with the `servicequotas.getService` data source or [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
	ServiceCode string `pulumi:"serviceCode"`
}

// A collection of values returned by getServiceQuota.
type LookupServiceQuotaResult struct {
	// Whether the service quota is adjustable.
	Adjustable bool `pulumi:"adjustable"`
	// ARN of the service quota.
	Arn string `pulumi:"arn"`
	// Default value of the service quota.
	DefaultValue float64 `pulumi:"defaultValue"`
	// Whether the service quota is global for the AWS account.
	GlobalQuota bool `pulumi:"globalQuota"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	QuotaCode   string `pulumi:"quotaCode"`
	QuotaName   string `pulumi:"quotaName"`
	ServiceCode string `pulumi:"serviceCode"`
	// Name of the service.
	ServiceName string `pulumi:"serviceName"`
	// Information about the measurement.
	UsageMetrics []GetServiceQuotaUsageMetric `pulumi:"usageMetrics"`
	// Current value of the service quota.
	Value float64 `pulumi:"value"`
}

func LookupServiceQuotaOutput(ctx *pulumi.Context, args LookupServiceQuotaOutputArgs, opts ...pulumi.InvokeOption) LookupServiceQuotaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceQuotaResult, error) {
			args := v.(LookupServiceQuotaArgs)
			r, err := LookupServiceQuota(ctx, &args, opts...)
			var s LookupServiceQuotaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceQuotaResultOutput)
}

// A collection of arguments for invoking getServiceQuota.
type LookupServiceQuotaOutputArgs struct {
	// Quota code within the service. When configured, the data source directly looks up the service quota. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html). One of `quotaCode` or `quotaName` must be specified.
	QuotaCode pulumi.StringPtrInput `pulumi:"quotaCode"`
	// Quota name within the service. When configured, the data source searches through all service quotas to find the matching quota name. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html). One of `quotaName` or `quotaCode` must be specified.
	QuotaName pulumi.StringPtrInput `pulumi:"quotaName"`
	// Service code for the quota. Available values can be found with the `servicequotas.getService` data source or [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
	ServiceCode pulumi.StringInput `pulumi:"serviceCode"`
}

func (LookupServiceQuotaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceQuotaArgs)(nil)).Elem()
}

// A collection of values returned by getServiceQuota.
type LookupServiceQuotaResultOutput struct{ *pulumi.OutputState }

func (LookupServiceQuotaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceQuotaResult)(nil)).Elem()
}

func (o LookupServiceQuotaResultOutput) ToLookupServiceQuotaResultOutput() LookupServiceQuotaResultOutput {
	return o
}

func (o LookupServiceQuotaResultOutput) ToLookupServiceQuotaResultOutputWithContext(ctx context.Context) LookupServiceQuotaResultOutput {
	return o
}

// Whether the service quota is adjustable.
func (o LookupServiceQuotaResultOutput) Adjustable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServiceQuotaResult) bool { return v.Adjustable }).(pulumi.BoolOutput)
}

// ARN of the service quota.
func (o LookupServiceQuotaResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceQuotaResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Default value of the service quota.
func (o LookupServiceQuotaResultOutput) DefaultValue() pulumi.Float64Output {
	return o.ApplyT(func(v LookupServiceQuotaResult) float64 { return v.DefaultValue }).(pulumi.Float64Output)
}

// Whether the service quota is global for the AWS account.
func (o LookupServiceQuotaResultOutput) GlobalQuota() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServiceQuotaResult) bool { return v.GlobalQuota }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupServiceQuotaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceQuotaResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServiceQuotaResultOutput) QuotaCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceQuotaResult) string { return v.QuotaCode }).(pulumi.StringOutput)
}

func (o LookupServiceQuotaResultOutput) QuotaName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceQuotaResult) string { return v.QuotaName }).(pulumi.StringOutput)
}

func (o LookupServiceQuotaResultOutput) ServiceCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceQuotaResult) string { return v.ServiceCode }).(pulumi.StringOutput)
}

// Name of the service.
func (o LookupServiceQuotaResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceQuotaResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Information about the measurement.
func (o LookupServiceQuotaResultOutput) UsageMetrics() GetServiceQuotaUsageMetricArrayOutput {
	return o.ApplyT(func(v LookupServiceQuotaResult) []GetServiceQuotaUsageMetric { return v.UsageMetrics }).(GetServiceQuotaUsageMetricArrayOutput)
}

// Current value of the service quota.
func (o LookupServiceQuotaResultOutput) Value() pulumi.Float64Output {
	return o.ApplyT(func(v LookupServiceQuotaResult) float64 { return v.Value }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceQuotaResultOutput{})
}
