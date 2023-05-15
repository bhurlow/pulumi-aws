// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a SageMaker monitoring schedule resource.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/sagemaker"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sagemaker.NewMonitoringSchedule(ctx, "test", &sagemaker.MonitoringScheduleArgs{
//				MonitoringScheduleConfig: &sagemaker.MonitoringScheduleMonitoringScheduleConfigArgs{
//					MonitoringJobDefinitionName: pulumi.Any(aws_sagemaker_data_quality_job_definition.Test.Name),
//					MonitoringType:              pulumi.String("DataQuality"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Monitoring schedules can be imported using the `name`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:sagemaker/monitoringSchedule:MonitoringSchedule test_monitoring_schedule monitoring-schedule-foo
//
// ```
type MonitoringSchedule struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) assigned by AWS to this monitoring schedule.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The configuration object that specifies the monitoring schedule and defines the monitoring job. Fields are documented below.
	MonitoringScheduleConfig MonitoringScheduleMonitoringScheduleConfigOutput `pulumi:"monitoringScheduleConfig"`
	// The name of the monitoring schedule. The name must be unique within an AWS Region within an AWS account. If omitted, the provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewMonitoringSchedule registers a new resource with the given unique name, arguments, and options.
func NewMonitoringSchedule(ctx *pulumi.Context,
	name string, args *MonitoringScheduleArgs, opts ...pulumi.ResourceOption) (*MonitoringSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MonitoringScheduleConfig == nil {
		return nil, errors.New("invalid value for required argument 'MonitoringScheduleConfig'")
	}
	var resource MonitoringSchedule
	err := ctx.RegisterResource("aws:sagemaker/monitoringSchedule:MonitoringSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMonitoringSchedule gets an existing MonitoringSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitoringSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitoringScheduleState, opts ...pulumi.ResourceOption) (*MonitoringSchedule, error) {
	var resource MonitoringSchedule
	err := ctx.ReadResource("aws:sagemaker/monitoringSchedule:MonitoringSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MonitoringSchedule resources.
type monitoringScheduleState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this monitoring schedule.
	Arn *string `pulumi:"arn"`
	// The configuration object that specifies the monitoring schedule and defines the monitoring job. Fields are documented below.
	MonitoringScheduleConfig *MonitoringScheduleMonitoringScheduleConfig `pulumi:"monitoringScheduleConfig"`
	// The name of the monitoring schedule. The name must be unique within an AWS Region within an AWS account. If omitted, the provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type MonitoringScheduleState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this monitoring schedule.
	Arn pulumi.StringPtrInput
	// The configuration object that specifies the monitoring schedule and defines the monitoring job. Fields are documented below.
	MonitoringScheduleConfig MonitoringScheduleMonitoringScheduleConfigPtrInput
	// The name of the monitoring schedule. The name must be unique within an AWS Region within an AWS account. If omitted, the provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (MonitoringScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringScheduleState)(nil)).Elem()
}

type monitoringScheduleArgs struct {
	// The configuration object that specifies the monitoring schedule and defines the monitoring job. Fields are documented below.
	MonitoringScheduleConfig MonitoringScheduleMonitoringScheduleConfig `pulumi:"monitoringScheduleConfig"`
	// The name of the monitoring schedule. The name must be unique within an AWS Region within an AWS account. If omitted, the provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MonitoringSchedule resource.
type MonitoringScheduleArgs struct {
	// The configuration object that specifies the monitoring schedule and defines the monitoring job. Fields are documented below.
	MonitoringScheduleConfig MonitoringScheduleMonitoringScheduleConfigInput
	// The name of the monitoring schedule. The name must be unique within an AWS Region within an AWS account. If omitted, the provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (MonitoringScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringScheduleArgs)(nil)).Elem()
}

type MonitoringScheduleInput interface {
	pulumi.Input

	ToMonitoringScheduleOutput() MonitoringScheduleOutput
	ToMonitoringScheduleOutputWithContext(ctx context.Context) MonitoringScheduleOutput
}

func (*MonitoringSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringSchedule)(nil)).Elem()
}

func (i *MonitoringSchedule) ToMonitoringScheduleOutput() MonitoringScheduleOutput {
	return i.ToMonitoringScheduleOutputWithContext(context.Background())
}

func (i *MonitoringSchedule) ToMonitoringScheduleOutputWithContext(ctx context.Context) MonitoringScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringScheduleOutput)
}

// MonitoringScheduleArrayInput is an input type that accepts MonitoringScheduleArray and MonitoringScheduleArrayOutput values.
// You can construct a concrete instance of `MonitoringScheduleArrayInput` via:
//
//	MonitoringScheduleArray{ MonitoringScheduleArgs{...} }
type MonitoringScheduleArrayInput interface {
	pulumi.Input

	ToMonitoringScheduleArrayOutput() MonitoringScheduleArrayOutput
	ToMonitoringScheduleArrayOutputWithContext(context.Context) MonitoringScheduleArrayOutput
}

type MonitoringScheduleArray []MonitoringScheduleInput

func (MonitoringScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MonitoringSchedule)(nil)).Elem()
}

func (i MonitoringScheduleArray) ToMonitoringScheduleArrayOutput() MonitoringScheduleArrayOutput {
	return i.ToMonitoringScheduleArrayOutputWithContext(context.Background())
}

func (i MonitoringScheduleArray) ToMonitoringScheduleArrayOutputWithContext(ctx context.Context) MonitoringScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringScheduleArrayOutput)
}

// MonitoringScheduleMapInput is an input type that accepts MonitoringScheduleMap and MonitoringScheduleMapOutput values.
// You can construct a concrete instance of `MonitoringScheduleMapInput` via:
//
//	MonitoringScheduleMap{ "key": MonitoringScheduleArgs{...} }
type MonitoringScheduleMapInput interface {
	pulumi.Input

	ToMonitoringScheduleMapOutput() MonitoringScheduleMapOutput
	ToMonitoringScheduleMapOutputWithContext(context.Context) MonitoringScheduleMapOutput
}

type MonitoringScheduleMap map[string]MonitoringScheduleInput

func (MonitoringScheduleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MonitoringSchedule)(nil)).Elem()
}

func (i MonitoringScheduleMap) ToMonitoringScheduleMapOutput() MonitoringScheduleMapOutput {
	return i.ToMonitoringScheduleMapOutputWithContext(context.Background())
}

func (i MonitoringScheduleMap) ToMonitoringScheduleMapOutputWithContext(ctx context.Context) MonitoringScheduleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringScheduleMapOutput)
}

type MonitoringScheduleOutput struct{ *pulumi.OutputState }

func (MonitoringScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringSchedule)(nil)).Elem()
}

func (o MonitoringScheduleOutput) ToMonitoringScheduleOutput() MonitoringScheduleOutput {
	return o
}

func (o MonitoringScheduleOutput) ToMonitoringScheduleOutputWithContext(ctx context.Context) MonitoringScheduleOutput {
	return o
}

// The Amazon Resource Name (ARN) assigned by AWS to this monitoring schedule.
func (o MonitoringScheduleOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringSchedule) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The configuration object that specifies the monitoring schedule and defines the monitoring job. Fields are documented below.
func (o MonitoringScheduleOutput) MonitoringScheduleConfig() MonitoringScheduleMonitoringScheduleConfigOutput {
	return o.ApplyT(func(v *MonitoringSchedule) MonitoringScheduleMonitoringScheduleConfigOutput {
		return v.MonitoringScheduleConfig
	}).(MonitoringScheduleMonitoringScheduleConfigOutput)
}

// The name of the monitoring schedule. The name must be unique within an AWS Region within an AWS account. If omitted, the provider will assign a random, unique name.
func (o MonitoringScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringSchedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o MonitoringScheduleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MonitoringSchedule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o MonitoringScheduleOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MonitoringSchedule) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type MonitoringScheduleArrayOutput struct{ *pulumi.OutputState }

func (MonitoringScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MonitoringSchedule)(nil)).Elem()
}

func (o MonitoringScheduleArrayOutput) ToMonitoringScheduleArrayOutput() MonitoringScheduleArrayOutput {
	return o
}

func (o MonitoringScheduleArrayOutput) ToMonitoringScheduleArrayOutputWithContext(ctx context.Context) MonitoringScheduleArrayOutput {
	return o
}

func (o MonitoringScheduleArrayOutput) Index(i pulumi.IntInput) MonitoringScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MonitoringSchedule {
		return vs[0].([]*MonitoringSchedule)[vs[1].(int)]
	}).(MonitoringScheduleOutput)
}

type MonitoringScheduleMapOutput struct{ *pulumi.OutputState }

func (MonitoringScheduleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MonitoringSchedule)(nil)).Elem()
}

func (o MonitoringScheduleMapOutput) ToMonitoringScheduleMapOutput() MonitoringScheduleMapOutput {
	return o
}

func (o MonitoringScheduleMapOutput) ToMonitoringScheduleMapOutputWithContext(ctx context.Context) MonitoringScheduleMapOutput {
	return o
}

func (o MonitoringScheduleMapOutput) MapIndex(k pulumi.StringInput) MonitoringScheduleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MonitoringSchedule {
		return vs[0].(map[string]*MonitoringSchedule)[vs[1].(string)]
	}).(MonitoringScheduleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MonitoringScheduleInput)(nil)).Elem(), &MonitoringSchedule{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitoringScheduleArrayInput)(nil)).Elem(), MonitoringScheduleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitoringScheduleMapInput)(nil)).Elem(), MonitoringScheduleMap{})
	pulumi.RegisterOutputType(MonitoringScheduleOutput{})
	pulumi.RegisterOutputType(MonitoringScheduleArrayOutput{})
	pulumi.RegisterOutputType(MonitoringScheduleMapOutput{})
}
