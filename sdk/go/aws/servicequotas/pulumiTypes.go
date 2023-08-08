// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicequotas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type ServiceQuotaUsageMetric struct {
	// The metric dimensions.
	MetricDimensions []ServiceQuotaUsageMetricMetricDimension `pulumi:"metricDimensions"`
	// The name of the metric.
	MetricName *string `pulumi:"metricName"`
	// The namespace of the metric.
	MetricNamespace *string `pulumi:"metricNamespace"`
	// The metric statistic that AWS recommend you use when determining quota usage.
	MetricStatisticRecommendation *string `pulumi:"metricStatisticRecommendation"`
}

// ServiceQuotaUsageMetricInput is an input type that accepts ServiceQuotaUsageMetricArgs and ServiceQuotaUsageMetricOutput values.
// You can construct a concrete instance of `ServiceQuotaUsageMetricInput` via:
//
//	ServiceQuotaUsageMetricArgs{...}
type ServiceQuotaUsageMetricInput interface {
	pulumi.Input

	ToServiceQuotaUsageMetricOutput() ServiceQuotaUsageMetricOutput
	ToServiceQuotaUsageMetricOutputWithContext(context.Context) ServiceQuotaUsageMetricOutput
}

type ServiceQuotaUsageMetricArgs struct {
	// The metric dimensions.
	MetricDimensions ServiceQuotaUsageMetricMetricDimensionArrayInput `pulumi:"metricDimensions"`
	// The name of the metric.
	MetricName pulumi.StringPtrInput `pulumi:"metricName"`
	// The namespace of the metric.
	MetricNamespace pulumi.StringPtrInput `pulumi:"metricNamespace"`
	// The metric statistic that AWS recommend you use when determining quota usage.
	MetricStatisticRecommendation pulumi.StringPtrInput `pulumi:"metricStatisticRecommendation"`
}

func (ServiceQuotaUsageMetricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceQuotaUsageMetric)(nil)).Elem()
}

func (i ServiceQuotaUsageMetricArgs) ToServiceQuotaUsageMetricOutput() ServiceQuotaUsageMetricOutput {
	return i.ToServiceQuotaUsageMetricOutputWithContext(context.Background())
}

func (i ServiceQuotaUsageMetricArgs) ToServiceQuotaUsageMetricOutputWithContext(ctx context.Context) ServiceQuotaUsageMetricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceQuotaUsageMetricOutput)
}

// ServiceQuotaUsageMetricArrayInput is an input type that accepts ServiceQuotaUsageMetricArray and ServiceQuotaUsageMetricArrayOutput values.
// You can construct a concrete instance of `ServiceQuotaUsageMetricArrayInput` via:
//
//	ServiceQuotaUsageMetricArray{ ServiceQuotaUsageMetricArgs{...} }
type ServiceQuotaUsageMetricArrayInput interface {
	pulumi.Input

	ToServiceQuotaUsageMetricArrayOutput() ServiceQuotaUsageMetricArrayOutput
	ToServiceQuotaUsageMetricArrayOutputWithContext(context.Context) ServiceQuotaUsageMetricArrayOutput
}

type ServiceQuotaUsageMetricArray []ServiceQuotaUsageMetricInput

func (ServiceQuotaUsageMetricArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceQuotaUsageMetric)(nil)).Elem()
}

func (i ServiceQuotaUsageMetricArray) ToServiceQuotaUsageMetricArrayOutput() ServiceQuotaUsageMetricArrayOutput {
	return i.ToServiceQuotaUsageMetricArrayOutputWithContext(context.Background())
}

func (i ServiceQuotaUsageMetricArray) ToServiceQuotaUsageMetricArrayOutputWithContext(ctx context.Context) ServiceQuotaUsageMetricArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceQuotaUsageMetricArrayOutput)
}

type ServiceQuotaUsageMetricOutput struct{ *pulumi.OutputState }

func (ServiceQuotaUsageMetricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceQuotaUsageMetric)(nil)).Elem()
}

func (o ServiceQuotaUsageMetricOutput) ToServiceQuotaUsageMetricOutput() ServiceQuotaUsageMetricOutput {
	return o
}

func (o ServiceQuotaUsageMetricOutput) ToServiceQuotaUsageMetricOutputWithContext(ctx context.Context) ServiceQuotaUsageMetricOutput {
	return o
}

// The metric dimensions.
func (o ServiceQuotaUsageMetricOutput) MetricDimensions() ServiceQuotaUsageMetricMetricDimensionArrayOutput {
	return o.ApplyT(func(v ServiceQuotaUsageMetric) []ServiceQuotaUsageMetricMetricDimension { return v.MetricDimensions }).(ServiceQuotaUsageMetricMetricDimensionArrayOutput)
}

// The name of the metric.
func (o ServiceQuotaUsageMetricOutput) MetricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceQuotaUsageMetric) *string { return v.MetricName }).(pulumi.StringPtrOutput)
}

// The namespace of the metric.
func (o ServiceQuotaUsageMetricOutput) MetricNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceQuotaUsageMetric) *string { return v.MetricNamespace }).(pulumi.StringPtrOutput)
}

// The metric statistic that AWS recommend you use when determining quota usage.
func (o ServiceQuotaUsageMetricOutput) MetricStatisticRecommendation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceQuotaUsageMetric) *string { return v.MetricStatisticRecommendation }).(pulumi.StringPtrOutput)
}

type ServiceQuotaUsageMetricArrayOutput struct{ *pulumi.OutputState }

func (ServiceQuotaUsageMetricArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceQuotaUsageMetric)(nil)).Elem()
}

func (o ServiceQuotaUsageMetricArrayOutput) ToServiceQuotaUsageMetricArrayOutput() ServiceQuotaUsageMetricArrayOutput {
	return o
}

func (o ServiceQuotaUsageMetricArrayOutput) ToServiceQuotaUsageMetricArrayOutputWithContext(ctx context.Context) ServiceQuotaUsageMetricArrayOutput {
	return o
}

func (o ServiceQuotaUsageMetricArrayOutput) Index(i pulumi.IntInput) ServiceQuotaUsageMetricOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceQuotaUsageMetric {
		return vs[0].([]ServiceQuotaUsageMetric)[vs[1].(int)]
	}).(ServiceQuotaUsageMetricOutput)
}

type ServiceQuotaUsageMetricMetricDimension struct {
	Class    *string `pulumi:"class"`
	Resource *string `pulumi:"resource"`
	Service  *string `pulumi:"service"`
	Type     *string `pulumi:"type"`
}

// ServiceQuotaUsageMetricMetricDimensionInput is an input type that accepts ServiceQuotaUsageMetricMetricDimensionArgs and ServiceQuotaUsageMetricMetricDimensionOutput values.
// You can construct a concrete instance of `ServiceQuotaUsageMetricMetricDimensionInput` via:
//
//	ServiceQuotaUsageMetricMetricDimensionArgs{...}
type ServiceQuotaUsageMetricMetricDimensionInput interface {
	pulumi.Input

	ToServiceQuotaUsageMetricMetricDimensionOutput() ServiceQuotaUsageMetricMetricDimensionOutput
	ToServiceQuotaUsageMetricMetricDimensionOutputWithContext(context.Context) ServiceQuotaUsageMetricMetricDimensionOutput
}

type ServiceQuotaUsageMetricMetricDimensionArgs struct {
	Class    pulumi.StringPtrInput `pulumi:"class"`
	Resource pulumi.StringPtrInput `pulumi:"resource"`
	Service  pulumi.StringPtrInput `pulumi:"service"`
	Type     pulumi.StringPtrInput `pulumi:"type"`
}

func (ServiceQuotaUsageMetricMetricDimensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceQuotaUsageMetricMetricDimension)(nil)).Elem()
}

func (i ServiceQuotaUsageMetricMetricDimensionArgs) ToServiceQuotaUsageMetricMetricDimensionOutput() ServiceQuotaUsageMetricMetricDimensionOutput {
	return i.ToServiceQuotaUsageMetricMetricDimensionOutputWithContext(context.Background())
}

func (i ServiceQuotaUsageMetricMetricDimensionArgs) ToServiceQuotaUsageMetricMetricDimensionOutputWithContext(ctx context.Context) ServiceQuotaUsageMetricMetricDimensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceQuotaUsageMetricMetricDimensionOutput)
}

// ServiceQuotaUsageMetricMetricDimensionArrayInput is an input type that accepts ServiceQuotaUsageMetricMetricDimensionArray and ServiceQuotaUsageMetricMetricDimensionArrayOutput values.
// You can construct a concrete instance of `ServiceQuotaUsageMetricMetricDimensionArrayInput` via:
//
//	ServiceQuotaUsageMetricMetricDimensionArray{ ServiceQuotaUsageMetricMetricDimensionArgs{...} }
type ServiceQuotaUsageMetricMetricDimensionArrayInput interface {
	pulumi.Input

	ToServiceQuotaUsageMetricMetricDimensionArrayOutput() ServiceQuotaUsageMetricMetricDimensionArrayOutput
	ToServiceQuotaUsageMetricMetricDimensionArrayOutputWithContext(context.Context) ServiceQuotaUsageMetricMetricDimensionArrayOutput
}

type ServiceQuotaUsageMetricMetricDimensionArray []ServiceQuotaUsageMetricMetricDimensionInput

func (ServiceQuotaUsageMetricMetricDimensionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceQuotaUsageMetricMetricDimension)(nil)).Elem()
}

func (i ServiceQuotaUsageMetricMetricDimensionArray) ToServiceQuotaUsageMetricMetricDimensionArrayOutput() ServiceQuotaUsageMetricMetricDimensionArrayOutput {
	return i.ToServiceQuotaUsageMetricMetricDimensionArrayOutputWithContext(context.Background())
}

func (i ServiceQuotaUsageMetricMetricDimensionArray) ToServiceQuotaUsageMetricMetricDimensionArrayOutputWithContext(ctx context.Context) ServiceQuotaUsageMetricMetricDimensionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceQuotaUsageMetricMetricDimensionArrayOutput)
}

type ServiceQuotaUsageMetricMetricDimensionOutput struct{ *pulumi.OutputState }

func (ServiceQuotaUsageMetricMetricDimensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceQuotaUsageMetricMetricDimension)(nil)).Elem()
}

func (o ServiceQuotaUsageMetricMetricDimensionOutput) ToServiceQuotaUsageMetricMetricDimensionOutput() ServiceQuotaUsageMetricMetricDimensionOutput {
	return o
}

func (o ServiceQuotaUsageMetricMetricDimensionOutput) ToServiceQuotaUsageMetricMetricDimensionOutputWithContext(ctx context.Context) ServiceQuotaUsageMetricMetricDimensionOutput {
	return o
}

func (o ServiceQuotaUsageMetricMetricDimensionOutput) Class() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceQuotaUsageMetricMetricDimension) *string { return v.Class }).(pulumi.StringPtrOutput)
}

func (o ServiceQuotaUsageMetricMetricDimensionOutput) Resource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceQuotaUsageMetricMetricDimension) *string { return v.Resource }).(pulumi.StringPtrOutput)
}

func (o ServiceQuotaUsageMetricMetricDimensionOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceQuotaUsageMetricMetricDimension) *string { return v.Service }).(pulumi.StringPtrOutput)
}

func (o ServiceQuotaUsageMetricMetricDimensionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceQuotaUsageMetricMetricDimension) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ServiceQuotaUsageMetricMetricDimensionArrayOutput struct{ *pulumi.OutputState }

func (ServiceQuotaUsageMetricMetricDimensionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceQuotaUsageMetricMetricDimension)(nil)).Elem()
}

func (o ServiceQuotaUsageMetricMetricDimensionArrayOutput) ToServiceQuotaUsageMetricMetricDimensionArrayOutput() ServiceQuotaUsageMetricMetricDimensionArrayOutput {
	return o
}

func (o ServiceQuotaUsageMetricMetricDimensionArrayOutput) ToServiceQuotaUsageMetricMetricDimensionArrayOutputWithContext(ctx context.Context) ServiceQuotaUsageMetricMetricDimensionArrayOutput {
	return o
}

func (o ServiceQuotaUsageMetricMetricDimensionArrayOutput) Index(i pulumi.IntInput) ServiceQuotaUsageMetricMetricDimensionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceQuotaUsageMetricMetricDimension {
		return vs[0].([]ServiceQuotaUsageMetricMetricDimension)[vs[1].(int)]
	}).(ServiceQuotaUsageMetricMetricDimensionOutput)
}

type GetServiceQuotaUsageMetric struct {
	// The metric dimensions.
	MetricDimensions []GetServiceQuotaUsageMetricMetricDimension `pulumi:"metricDimensions"`
	// The name of the metric.
	MetricName string `pulumi:"metricName"`
	// The namespace of the metric.
	MetricNamespace string `pulumi:"metricNamespace"`
	// The metric statistic that AWS recommend you use when determining quota usage.
	MetricStatisticRecommendation string `pulumi:"metricStatisticRecommendation"`
}

// GetServiceQuotaUsageMetricInput is an input type that accepts GetServiceQuotaUsageMetricArgs and GetServiceQuotaUsageMetricOutput values.
// You can construct a concrete instance of `GetServiceQuotaUsageMetricInput` via:
//
//	GetServiceQuotaUsageMetricArgs{...}
type GetServiceQuotaUsageMetricInput interface {
	pulumi.Input

	ToGetServiceQuotaUsageMetricOutput() GetServiceQuotaUsageMetricOutput
	ToGetServiceQuotaUsageMetricOutputWithContext(context.Context) GetServiceQuotaUsageMetricOutput
}

type GetServiceQuotaUsageMetricArgs struct {
	// The metric dimensions.
	MetricDimensions GetServiceQuotaUsageMetricMetricDimensionArrayInput `pulumi:"metricDimensions"`
	// The name of the metric.
	MetricName pulumi.StringInput `pulumi:"metricName"`
	// The namespace of the metric.
	MetricNamespace pulumi.StringInput `pulumi:"metricNamespace"`
	// The metric statistic that AWS recommend you use when determining quota usage.
	MetricStatisticRecommendation pulumi.StringInput `pulumi:"metricStatisticRecommendation"`
}

func (GetServiceQuotaUsageMetricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceQuotaUsageMetric)(nil)).Elem()
}

func (i GetServiceQuotaUsageMetricArgs) ToGetServiceQuotaUsageMetricOutput() GetServiceQuotaUsageMetricOutput {
	return i.ToGetServiceQuotaUsageMetricOutputWithContext(context.Background())
}

func (i GetServiceQuotaUsageMetricArgs) ToGetServiceQuotaUsageMetricOutputWithContext(ctx context.Context) GetServiceQuotaUsageMetricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetServiceQuotaUsageMetricOutput)
}

// GetServiceQuotaUsageMetricArrayInput is an input type that accepts GetServiceQuotaUsageMetricArray and GetServiceQuotaUsageMetricArrayOutput values.
// You can construct a concrete instance of `GetServiceQuotaUsageMetricArrayInput` via:
//
//	GetServiceQuotaUsageMetricArray{ GetServiceQuotaUsageMetricArgs{...} }
type GetServiceQuotaUsageMetricArrayInput interface {
	pulumi.Input

	ToGetServiceQuotaUsageMetricArrayOutput() GetServiceQuotaUsageMetricArrayOutput
	ToGetServiceQuotaUsageMetricArrayOutputWithContext(context.Context) GetServiceQuotaUsageMetricArrayOutput
}

type GetServiceQuotaUsageMetricArray []GetServiceQuotaUsageMetricInput

func (GetServiceQuotaUsageMetricArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetServiceQuotaUsageMetric)(nil)).Elem()
}

func (i GetServiceQuotaUsageMetricArray) ToGetServiceQuotaUsageMetricArrayOutput() GetServiceQuotaUsageMetricArrayOutput {
	return i.ToGetServiceQuotaUsageMetricArrayOutputWithContext(context.Background())
}

func (i GetServiceQuotaUsageMetricArray) ToGetServiceQuotaUsageMetricArrayOutputWithContext(ctx context.Context) GetServiceQuotaUsageMetricArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetServiceQuotaUsageMetricArrayOutput)
}

type GetServiceQuotaUsageMetricOutput struct{ *pulumi.OutputState }

func (GetServiceQuotaUsageMetricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceQuotaUsageMetric)(nil)).Elem()
}

func (o GetServiceQuotaUsageMetricOutput) ToGetServiceQuotaUsageMetricOutput() GetServiceQuotaUsageMetricOutput {
	return o
}

func (o GetServiceQuotaUsageMetricOutput) ToGetServiceQuotaUsageMetricOutputWithContext(ctx context.Context) GetServiceQuotaUsageMetricOutput {
	return o
}

// The metric dimensions.
func (o GetServiceQuotaUsageMetricOutput) MetricDimensions() GetServiceQuotaUsageMetricMetricDimensionArrayOutput {
	return o.ApplyT(func(v GetServiceQuotaUsageMetric) []GetServiceQuotaUsageMetricMetricDimension {
		return v.MetricDimensions
	}).(GetServiceQuotaUsageMetricMetricDimensionArrayOutput)
}

// The name of the metric.
func (o GetServiceQuotaUsageMetricOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceQuotaUsageMetric) string { return v.MetricName }).(pulumi.StringOutput)
}

// The namespace of the metric.
func (o GetServiceQuotaUsageMetricOutput) MetricNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceQuotaUsageMetric) string { return v.MetricNamespace }).(pulumi.StringOutput)
}

// The metric statistic that AWS recommend you use when determining quota usage.
func (o GetServiceQuotaUsageMetricOutput) MetricStatisticRecommendation() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceQuotaUsageMetric) string { return v.MetricStatisticRecommendation }).(pulumi.StringOutput)
}

type GetServiceQuotaUsageMetricArrayOutput struct{ *pulumi.OutputState }

func (GetServiceQuotaUsageMetricArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetServiceQuotaUsageMetric)(nil)).Elem()
}

func (o GetServiceQuotaUsageMetricArrayOutput) ToGetServiceQuotaUsageMetricArrayOutput() GetServiceQuotaUsageMetricArrayOutput {
	return o
}

func (o GetServiceQuotaUsageMetricArrayOutput) ToGetServiceQuotaUsageMetricArrayOutputWithContext(ctx context.Context) GetServiceQuotaUsageMetricArrayOutput {
	return o
}

func (o GetServiceQuotaUsageMetricArrayOutput) Index(i pulumi.IntInput) GetServiceQuotaUsageMetricOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetServiceQuotaUsageMetric {
		return vs[0].([]GetServiceQuotaUsageMetric)[vs[1].(int)]
	}).(GetServiceQuotaUsageMetricOutput)
}

type GetServiceQuotaUsageMetricMetricDimension struct {
	Class    string `pulumi:"class"`
	Resource string `pulumi:"resource"`
	Service  string `pulumi:"service"`
	Type     string `pulumi:"type"`
}

// GetServiceQuotaUsageMetricMetricDimensionInput is an input type that accepts GetServiceQuotaUsageMetricMetricDimensionArgs and GetServiceQuotaUsageMetricMetricDimensionOutput values.
// You can construct a concrete instance of `GetServiceQuotaUsageMetricMetricDimensionInput` via:
//
//	GetServiceQuotaUsageMetricMetricDimensionArgs{...}
type GetServiceQuotaUsageMetricMetricDimensionInput interface {
	pulumi.Input

	ToGetServiceQuotaUsageMetricMetricDimensionOutput() GetServiceQuotaUsageMetricMetricDimensionOutput
	ToGetServiceQuotaUsageMetricMetricDimensionOutputWithContext(context.Context) GetServiceQuotaUsageMetricMetricDimensionOutput
}

type GetServiceQuotaUsageMetricMetricDimensionArgs struct {
	Class    pulumi.StringInput `pulumi:"class"`
	Resource pulumi.StringInput `pulumi:"resource"`
	Service  pulumi.StringInput `pulumi:"service"`
	Type     pulumi.StringInput `pulumi:"type"`
}

func (GetServiceQuotaUsageMetricMetricDimensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceQuotaUsageMetricMetricDimension)(nil)).Elem()
}

func (i GetServiceQuotaUsageMetricMetricDimensionArgs) ToGetServiceQuotaUsageMetricMetricDimensionOutput() GetServiceQuotaUsageMetricMetricDimensionOutput {
	return i.ToGetServiceQuotaUsageMetricMetricDimensionOutputWithContext(context.Background())
}

func (i GetServiceQuotaUsageMetricMetricDimensionArgs) ToGetServiceQuotaUsageMetricMetricDimensionOutputWithContext(ctx context.Context) GetServiceQuotaUsageMetricMetricDimensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetServiceQuotaUsageMetricMetricDimensionOutput)
}

// GetServiceQuotaUsageMetricMetricDimensionArrayInput is an input type that accepts GetServiceQuotaUsageMetricMetricDimensionArray and GetServiceQuotaUsageMetricMetricDimensionArrayOutput values.
// You can construct a concrete instance of `GetServiceQuotaUsageMetricMetricDimensionArrayInput` via:
//
//	GetServiceQuotaUsageMetricMetricDimensionArray{ GetServiceQuotaUsageMetricMetricDimensionArgs{...} }
type GetServiceQuotaUsageMetricMetricDimensionArrayInput interface {
	pulumi.Input

	ToGetServiceQuotaUsageMetricMetricDimensionArrayOutput() GetServiceQuotaUsageMetricMetricDimensionArrayOutput
	ToGetServiceQuotaUsageMetricMetricDimensionArrayOutputWithContext(context.Context) GetServiceQuotaUsageMetricMetricDimensionArrayOutput
}

type GetServiceQuotaUsageMetricMetricDimensionArray []GetServiceQuotaUsageMetricMetricDimensionInput

func (GetServiceQuotaUsageMetricMetricDimensionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetServiceQuotaUsageMetricMetricDimension)(nil)).Elem()
}

func (i GetServiceQuotaUsageMetricMetricDimensionArray) ToGetServiceQuotaUsageMetricMetricDimensionArrayOutput() GetServiceQuotaUsageMetricMetricDimensionArrayOutput {
	return i.ToGetServiceQuotaUsageMetricMetricDimensionArrayOutputWithContext(context.Background())
}

func (i GetServiceQuotaUsageMetricMetricDimensionArray) ToGetServiceQuotaUsageMetricMetricDimensionArrayOutputWithContext(ctx context.Context) GetServiceQuotaUsageMetricMetricDimensionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetServiceQuotaUsageMetricMetricDimensionArrayOutput)
}

type GetServiceQuotaUsageMetricMetricDimensionOutput struct{ *pulumi.OutputState }

func (GetServiceQuotaUsageMetricMetricDimensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceQuotaUsageMetricMetricDimension)(nil)).Elem()
}

func (o GetServiceQuotaUsageMetricMetricDimensionOutput) ToGetServiceQuotaUsageMetricMetricDimensionOutput() GetServiceQuotaUsageMetricMetricDimensionOutput {
	return o
}

func (o GetServiceQuotaUsageMetricMetricDimensionOutput) ToGetServiceQuotaUsageMetricMetricDimensionOutputWithContext(ctx context.Context) GetServiceQuotaUsageMetricMetricDimensionOutput {
	return o
}

func (o GetServiceQuotaUsageMetricMetricDimensionOutput) Class() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceQuotaUsageMetricMetricDimension) string { return v.Class }).(pulumi.StringOutput)
}

func (o GetServiceQuotaUsageMetricMetricDimensionOutput) Resource() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceQuotaUsageMetricMetricDimension) string { return v.Resource }).(pulumi.StringOutput)
}

func (o GetServiceQuotaUsageMetricMetricDimensionOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceQuotaUsageMetricMetricDimension) string { return v.Service }).(pulumi.StringOutput)
}

func (o GetServiceQuotaUsageMetricMetricDimensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceQuotaUsageMetricMetricDimension) string { return v.Type }).(pulumi.StringOutput)
}

type GetServiceQuotaUsageMetricMetricDimensionArrayOutput struct{ *pulumi.OutputState }

func (GetServiceQuotaUsageMetricMetricDimensionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetServiceQuotaUsageMetricMetricDimension)(nil)).Elem()
}

func (o GetServiceQuotaUsageMetricMetricDimensionArrayOutput) ToGetServiceQuotaUsageMetricMetricDimensionArrayOutput() GetServiceQuotaUsageMetricMetricDimensionArrayOutput {
	return o
}

func (o GetServiceQuotaUsageMetricMetricDimensionArrayOutput) ToGetServiceQuotaUsageMetricMetricDimensionArrayOutputWithContext(ctx context.Context) GetServiceQuotaUsageMetricMetricDimensionArrayOutput {
	return o
}

func (o GetServiceQuotaUsageMetricMetricDimensionArrayOutput) Index(i pulumi.IntInput) GetServiceQuotaUsageMetricMetricDimensionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetServiceQuotaUsageMetricMetricDimension {
		return vs[0].([]GetServiceQuotaUsageMetricMetricDimension)[vs[1].(int)]
	}).(GetServiceQuotaUsageMetricMetricDimensionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceQuotaUsageMetricInput)(nil)).Elem(), ServiceQuotaUsageMetricArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceQuotaUsageMetricArrayInput)(nil)).Elem(), ServiceQuotaUsageMetricArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceQuotaUsageMetricMetricDimensionInput)(nil)).Elem(), ServiceQuotaUsageMetricMetricDimensionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceQuotaUsageMetricMetricDimensionArrayInput)(nil)).Elem(), ServiceQuotaUsageMetricMetricDimensionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetServiceQuotaUsageMetricInput)(nil)).Elem(), GetServiceQuotaUsageMetricArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetServiceQuotaUsageMetricArrayInput)(nil)).Elem(), GetServiceQuotaUsageMetricArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetServiceQuotaUsageMetricMetricDimensionInput)(nil)).Elem(), GetServiceQuotaUsageMetricMetricDimensionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetServiceQuotaUsageMetricMetricDimensionArrayInput)(nil)).Elem(), GetServiceQuotaUsageMetricMetricDimensionArray{})
	pulumi.RegisterOutputType(ServiceQuotaUsageMetricOutput{})
	pulumi.RegisterOutputType(ServiceQuotaUsageMetricArrayOutput{})
	pulumi.RegisterOutputType(ServiceQuotaUsageMetricMetricDimensionOutput{})
	pulumi.RegisterOutputType(ServiceQuotaUsageMetricMetricDimensionArrayOutput{})
	pulumi.RegisterOutputType(GetServiceQuotaUsageMetricOutput{})
	pulumi.RegisterOutputType(GetServiceQuotaUsageMetricArrayOutput{})
	pulumi.RegisterOutputType(GetServiceQuotaUsageMetricMetricDimensionOutput{})
	pulumi.RegisterOutputType(GetServiceQuotaUsageMetricMetricDimensionArrayOutput{})
}
