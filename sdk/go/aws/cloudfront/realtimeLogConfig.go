// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CloudFront real-time log configuration resource.
//
// ## Import
//
// CloudFront real-time log configurations can be imported using the ARN, e.g.,
//
// ```sh
//
//	$ pulumi import aws:cloudfront/realtimeLogConfig:RealtimeLogConfig example arn:aws:cloudfront::111122223333:realtime-log-config/ExampleNameForRealtimeLogConfig
//
// ```
type RealtimeLogConfig struct {
	pulumi.CustomResourceState

	// The ARN (Amazon Resource Name) of the CloudFront real-time log configuration.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Amazon Kinesis data streams where real-time log data is sent.
	Endpoint RealtimeLogConfigEndpointOutput `pulumi:"endpoint"`
	// The fields that are included in each real-time log record. See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) for supported values.
	Fields pulumi.StringArrayOutput `pulumi:"fields"`
	// The unique name to identify this real-time log configuration.
	Name pulumi.StringOutput `pulumi:"name"`
	// The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between `1` and `100`, inclusive.
	SamplingRate pulumi.IntOutput `pulumi:"samplingRate"`
}

// NewRealtimeLogConfig registers a new resource with the given unique name, arguments, and options.
func NewRealtimeLogConfig(ctx *pulumi.Context,
	name string, args *RealtimeLogConfigArgs, opts ...pulumi.ResourceOption) (*RealtimeLogConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Endpoint == nil {
		return nil, errors.New("invalid value for required argument 'Endpoint'")
	}
	if args.Fields == nil {
		return nil, errors.New("invalid value for required argument 'Fields'")
	}
	if args.SamplingRate == nil {
		return nil, errors.New("invalid value for required argument 'SamplingRate'")
	}
	var resource RealtimeLogConfig
	err := ctx.RegisterResource("aws:cloudfront/realtimeLogConfig:RealtimeLogConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRealtimeLogConfig gets an existing RealtimeLogConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRealtimeLogConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RealtimeLogConfigState, opts ...pulumi.ResourceOption) (*RealtimeLogConfig, error) {
	var resource RealtimeLogConfig
	err := ctx.ReadResource("aws:cloudfront/realtimeLogConfig:RealtimeLogConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RealtimeLogConfig resources.
type realtimeLogConfigState struct {
	// The ARN (Amazon Resource Name) of the CloudFront real-time log configuration.
	Arn *string `pulumi:"arn"`
	// The Amazon Kinesis data streams where real-time log data is sent.
	Endpoint *RealtimeLogConfigEndpoint `pulumi:"endpoint"`
	// The fields that are included in each real-time log record. See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) for supported values.
	Fields []string `pulumi:"fields"`
	// The unique name to identify this real-time log configuration.
	Name *string `pulumi:"name"`
	// The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between `1` and `100`, inclusive.
	SamplingRate *int `pulumi:"samplingRate"`
}

type RealtimeLogConfigState struct {
	// The ARN (Amazon Resource Name) of the CloudFront real-time log configuration.
	Arn pulumi.StringPtrInput
	// The Amazon Kinesis data streams where real-time log data is sent.
	Endpoint RealtimeLogConfigEndpointPtrInput
	// The fields that are included in each real-time log record. See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) for supported values.
	Fields pulumi.StringArrayInput
	// The unique name to identify this real-time log configuration.
	Name pulumi.StringPtrInput
	// The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between `1` and `100`, inclusive.
	SamplingRate pulumi.IntPtrInput
}

func (RealtimeLogConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*realtimeLogConfigState)(nil)).Elem()
}

type realtimeLogConfigArgs struct {
	// The Amazon Kinesis data streams where real-time log data is sent.
	Endpoint RealtimeLogConfigEndpoint `pulumi:"endpoint"`
	// The fields that are included in each real-time log record. See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) for supported values.
	Fields []string `pulumi:"fields"`
	// The unique name to identify this real-time log configuration.
	Name *string `pulumi:"name"`
	// The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between `1` and `100`, inclusive.
	SamplingRate int `pulumi:"samplingRate"`
}

// The set of arguments for constructing a RealtimeLogConfig resource.
type RealtimeLogConfigArgs struct {
	// The Amazon Kinesis data streams where real-time log data is sent.
	Endpoint RealtimeLogConfigEndpointInput
	// The fields that are included in each real-time log record. See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) for supported values.
	Fields pulumi.StringArrayInput
	// The unique name to identify this real-time log configuration.
	Name pulumi.StringPtrInput
	// The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between `1` and `100`, inclusive.
	SamplingRate pulumi.IntInput
}

func (RealtimeLogConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*realtimeLogConfigArgs)(nil)).Elem()
}

type RealtimeLogConfigInput interface {
	pulumi.Input

	ToRealtimeLogConfigOutput() RealtimeLogConfigOutput
	ToRealtimeLogConfigOutputWithContext(ctx context.Context) RealtimeLogConfigOutput
}

func (*RealtimeLogConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**RealtimeLogConfig)(nil)).Elem()
}

func (i *RealtimeLogConfig) ToRealtimeLogConfigOutput() RealtimeLogConfigOutput {
	return i.ToRealtimeLogConfigOutputWithContext(context.Background())
}

func (i *RealtimeLogConfig) ToRealtimeLogConfigOutputWithContext(ctx context.Context) RealtimeLogConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealtimeLogConfigOutput)
}

// RealtimeLogConfigArrayInput is an input type that accepts RealtimeLogConfigArray and RealtimeLogConfigArrayOutput values.
// You can construct a concrete instance of `RealtimeLogConfigArrayInput` via:
//
//	RealtimeLogConfigArray{ RealtimeLogConfigArgs{...} }
type RealtimeLogConfigArrayInput interface {
	pulumi.Input

	ToRealtimeLogConfigArrayOutput() RealtimeLogConfigArrayOutput
	ToRealtimeLogConfigArrayOutputWithContext(context.Context) RealtimeLogConfigArrayOutput
}

type RealtimeLogConfigArray []RealtimeLogConfigInput

func (RealtimeLogConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealtimeLogConfig)(nil)).Elem()
}

func (i RealtimeLogConfigArray) ToRealtimeLogConfigArrayOutput() RealtimeLogConfigArrayOutput {
	return i.ToRealtimeLogConfigArrayOutputWithContext(context.Background())
}

func (i RealtimeLogConfigArray) ToRealtimeLogConfigArrayOutputWithContext(ctx context.Context) RealtimeLogConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealtimeLogConfigArrayOutput)
}

// RealtimeLogConfigMapInput is an input type that accepts RealtimeLogConfigMap and RealtimeLogConfigMapOutput values.
// You can construct a concrete instance of `RealtimeLogConfigMapInput` via:
//
//	RealtimeLogConfigMap{ "key": RealtimeLogConfigArgs{...} }
type RealtimeLogConfigMapInput interface {
	pulumi.Input

	ToRealtimeLogConfigMapOutput() RealtimeLogConfigMapOutput
	ToRealtimeLogConfigMapOutputWithContext(context.Context) RealtimeLogConfigMapOutput
}

type RealtimeLogConfigMap map[string]RealtimeLogConfigInput

func (RealtimeLogConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealtimeLogConfig)(nil)).Elem()
}

func (i RealtimeLogConfigMap) ToRealtimeLogConfigMapOutput() RealtimeLogConfigMapOutput {
	return i.ToRealtimeLogConfigMapOutputWithContext(context.Background())
}

func (i RealtimeLogConfigMap) ToRealtimeLogConfigMapOutputWithContext(ctx context.Context) RealtimeLogConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealtimeLogConfigMapOutput)
}

type RealtimeLogConfigOutput struct{ *pulumi.OutputState }

func (RealtimeLogConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RealtimeLogConfig)(nil)).Elem()
}

func (o RealtimeLogConfigOutput) ToRealtimeLogConfigOutput() RealtimeLogConfigOutput {
	return o
}

func (o RealtimeLogConfigOutput) ToRealtimeLogConfigOutputWithContext(ctx context.Context) RealtimeLogConfigOutput {
	return o
}

// The ARN (Amazon Resource Name) of the CloudFront real-time log configuration.
func (o RealtimeLogConfigOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *RealtimeLogConfig) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The Amazon Kinesis data streams where real-time log data is sent.
func (o RealtimeLogConfigOutput) Endpoint() RealtimeLogConfigEndpointOutput {
	return o.ApplyT(func(v *RealtimeLogConfig) RealtimeLogConfigEndpointOutput { return v.Endpoint }).(RealtimeLogConfigEndpointOutput)
}

// The fields that are included in each real-time log record. See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) for supported values.
func (o RealtimeLogConfigOutput) Fields() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RealtimeLogConfig) pulumi.StringArrayOutput { return v.Fields }).(pulumi.StringArrayOutput)
}

// The unique name to identify this real-time log configuration.
func (o RealtimeLogConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RealtimeLogConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between `1` and `100`, inclusive.
func (o RealtimeLogConfigOutput) SamplingRate() pulumi.IntOutput {
	return o.ApplyT(func(v *RealtimeLogConfig) pulumi.IntOutput { return v.SamplingRate }).(pulumi.IntOutput)
}

type RealtimeLogConfigArrayOutput struct{ *pulumi.OutputState }

func (RealtimeLogConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealtimeLogConfig)(nil)).Elem()
}

func (o RealtimeLogConfigArrayOutput) ToRealtimeLogConfigArrayOutput() RealtimeLogConfigArrayOutput {
	return o
}

func (o RealtimeLogConfigArrayOutput) ToRealtimeLogConfigArrayOutputWithContext(ctx context.Context) RealtimeLogConfigArrayOutput {
	return o
}

func (o RealtimeLogConfigArrayOutput) Index(i pulumi.IntInput) RealtimeLogConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RealtimeLogConfig {
		return vs[0].([]*RealtimeLogConfig)[vs[1].(int)]
	}).(RealtimeLogConfigOutput)
}

type RealtimeLogConfigMapOutput struct{ *pulumi.OutputState }

func (RealtimeLogConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealtimeLogConfig)(nil)).Elem()
}

func (o RealtimeLogConfigMapOutput) ToRealtimeLogConfigMapOutput() RealtimeLogConfigMapOutput {
	return o
}

func (o RealtimeLogConfigMapOutput) ToRealtimeLogConfigMapOutputWithContext(ctx context.Context) RealtimeLogConfigMapOutput {
	return o
}

func (o RealtimeLogConfigMapOutput) MapIndex(k pulumi.StringInput) RealtimeLogConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RealtimeLogConfig {
		return vs[0].(map[string]*RealtimeLogConfig)[vs[1].(string)]
	}).(RealtimeLogConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RealtimeLogConfigInput)(nil)).Elem(), &RealtimeLogConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealtimeLogConfigArrayInput)(nil)).Elem(), RealtimeLogConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealtimeLogConfigMapInput)(nil)).Elem(), RealtimeLogConfigMap{})
	pulumi.RegisterOutputType(RealtimeLogConfigOutput{})
	pulumi.RegisterOutputType(RealtimeLogConfigArrayOutput{})
	pulumi.RegisterOutputType(RealtimeLogConfigMapOutput{})
}
