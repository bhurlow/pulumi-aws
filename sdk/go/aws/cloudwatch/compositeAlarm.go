// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CloudWatch Composite Alarm resource.
//
// > **NOTE:** An alarm (composite or metric) cannot be destroyed when there are other composite alarms depending on it. This can lead to a cyclical dependency on update, as the provider will unsuccessfully attempt to destroy alarms before updating the rule. Consider using `dependsOn`, references to alarm names, and two-stage updates.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudwatch"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudwatch.NewCompositeAlarm(ctx, "example", &cloudwatch.CompositeAlarmArgs{
//				AlarmDescription: pulumi.String("This is a composite alarm!"),
//				AlarmName:        pulumi.String("example-composite-alarm"),
//				AlarmActions:     pulumi.Any(aws_sns_topic.Example.Arn),
//				OkActions:        pulumi.Any(aws_sns_topic.Example.Arn),
//				AlarmRule:        pulumi.String(fmt.Sprintf("ALARM(%v) OR\nALARM(%v)\n", aws_cloudwatch_metric_alarm.Alpha.Alarm_name, aws_cloudwatch_metric_alarm.Bravo.Alarm_name)),
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
// Use the `alarm_name` to import a CloudWatch Composite Alarm. For example
//
// ```sh
//
//	$ pulumi import aws:cloudwatch/compositeAlarm:CompositeAlarm test my-alarm
//
// ```
type CompositeAlarm struct {
	pulumi.CustomResourceState

	// Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to `true`.
	ActionsEnabled pulumi.BoolPtrOutput `pulumi:"actionsEnabled"`
	// The set of actions to execute when this alarm transitions to the `ALARM` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	AlarmActions pulumi.StringArrayOutput `pulumi:"alarmActions"`
	// The description for the composite alarm.
	AlarmDescription pulumi.StringPtrOutput `pulumi:"alarmDescription"`
	// The name for the composite alarm. This name must be unique within the region.
	AlarmName pulumi.StringOutput `pulumi:"alarmName"`
	// An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see [Creating a Composite Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Composite_Alarm.html). The maximum length is 10240 characters.
	AlarmRule pulumi.StringOutput `pulumi:"alarmRule"`
	// The ARN of the composite alarm.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The set of actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	InsufficientDataActions pulumi.StringArrayOutput `pulumi:"insufficientDataActions"`
	// The set of actions to execute when this alarm transitions to an `OK` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	OkActions pulumi.StringArrayOutput `pulumi:"okActions"`
	// A map of tags to associate with the alarm. Up to 50 tags are allowed. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewCompositeAlarm registers a new resource with the given unique name, arguments, and options.
func NewCompositeAlarm(ctx *pulumi.Context,
	name string, args *CompositeAlarmArgs, opts ...pulumi.ResourceOption) (*CompositeAlarm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AlarmName == nil {
		return nil, errors.New("invalid value for required argument 'AlarmName'")
	}
	if args.AlarmRule == nil {
		return nil, errors.New("invalid value for required argument 'AlarmRule'")
	}
	var resource CompositeAlarm
	err := ctx.RegisterResource("aws:cloudwatch/compositeAlarm:CompositeAlarm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCompositeAlarm gets an existing CompositeAlarm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCompositeAlarm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CompositeAlarmState, opts ...pulumi.ResourceOption) (*CompositeAlarm, error) {
	var resource CompositeAlarm
	err := ctx.ReadResource("aws:cloudwatch/compositeAlarm:CompositeAlarm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CompositeAlarm resources.
type compositeAlarmState struct {
	// Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to `true`.
	ActionsEnabled *bool `pulumi:"actionsEnabled"`
	// The set of actions to execute when this alarm transitions to the `ALARM` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	AlarmActions []string `pulumi:"alarmActions"`
	// The description for the composite alarm.
	AlarmDescription *string `pulumi:"alarmDescription"`
	// The name for the composite alarm. This name must be unique within the region.
	AlarmName *string `pulumi:"alarmName"`
	// An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see [Creating a Composite Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Composite_Alarm.html). The maximum length is 10240 characters.
	AlarmRule *string `pulumi:"alarmRule"`
	// The ARN of the composite alarm.
	Arn *string `pulumi:"arn"`
	// The set of actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	InsufficientDataActions []string `pulumi:"insufficientDataActions"`
	// The set of actions to execute when this alarm transitions to an `OK` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	OkActions []string `pulumi:"okActions"`
	// A map of tags to associate with the alarm. Up to 50 tags are allowed. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type CompositeAlarmState struct {
	// Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to `true`.
	ActionsEnabled pulumi.BoolPtrInput
	// The set of actions to execute when this alarm transitions to the `ALARM` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	AlarmActions pulumi.StringArrayInput
	// The description for the composite alarm.
	AlarmDescription pulumi.StringPtrInput
	// The name for the composite alarm. This name must be unique within the region.
	AlarmName pulumi.StringPtrInput
	// An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see [Creating a Composite Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Composite_Alarm.html). The maximum length is 10240 characters.
	AlarmRule pulumi.StringPtrInput
	// The ARN of the composite alarm.
	Arn pulumi.StringPtrInput
	// The set of actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	InsufficientDataActions pulumi.StringArrayInput
	// The set of actions to execute when this alarm transitions to an `OK` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	OkActions pulumi.StringArrayInput
	// A map of tags to associate with the alarm. Up to 50 tags are allowed. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (CompositeAlarmState) ElementType() reflect.Type {
	return reflect.TypeOf((*compositeAlarmState)(nil)).Elem()
}

type compositeAlarmArgs struct {
	// Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to `true`.
	ActionsEnabled *bool `pulumi:"actionsEnabled"`
	// The set of actions to execute when this alarm transitions to the `ALARM` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	AlarmActions []string `pulumi:"alarmActions"`
	// The description for the composite alarm.
	AlarmDescription *string `pulumi:"alarmDescription"`
	// The name for the composite alarm. This name must be unique within the region.
	AlarmName string `pulumi:"alarmName"`
	// An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see [Creating a Composite Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Composite_Alarm.html). The maximum length is 10240 characters.
	AlarmRule string `pulumi:"alarmRule"`
	// The set of actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	InsufficientDataActions []string `pulumi:"insufficientDataActions"`
	// The set of actions to execute when this alarm transitions to an `OK` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	OkActions []string `pulumi:"okActions"`
	// A map of tags to associate with the alarm. Up to 50 tags are allowed. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CompositeAlarm resource.
type CompositeAlarmArgs struct {
	// Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to `true`.
	ActionsEnabled pulumi.BoolPtrInput
	// The set of actions to execute when this alarm transitions to the `ALARM` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	AlarmActions pulumi.StringArrayInput
	// The description for the composite alarm.
	AlarmDescription pulumi.StringPtrInput
	// The name for the composite alarm. This name must be unique within the region.
	AlarmName pulumi.StringInput
	// An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see [Creating a Composite Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Composite_Alarm.html). The maximum length is 10240 characters.
	AlarmRule pulumi.StringInput
	// The set of actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	InsufficientDataActions pulumi.StringArrayInput
	// The set of actions to execute when this alarm transitions to an `OK` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	OkActions pulumi.StringArrayInput
	// A map of tags to associate with the alarm. Up to 50 tags are allowed. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (CompositeAlarmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*compositeAlarmArgs)(nil)).Elem()
}

type CompositeAlarmInput interface {
	pulumi.Input

	ToCompositeAlarmOutput() CompositeAlarmOutput
	ToCompositeAlarmOutputWithContext(ctx context.Context) CompositeAlarmOutput
}

func (*CompositeAlarm) ElementType() reflect.Type {
	return reflect.TypeOf((**CompositeAlarm)(nil)).Elem()
}

func (i *CompositeAlarm) ToCompositeAlarmOutput() CompositeAlarmOutput {
	return i.ToCompositeAlarmOutputWithContext(context.Background())
}

func (i *CompositeAlarm) ToCompositeAlarmOutputWithContext(ctx context.Context) CompositeAlarmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompositeAlarmOutput)
}

// CompositeAlarmArrayInput is an input type that accepts CompositeAlarmArray and CompositeAlarmArrayOutput values.
// You can construct a concrete instance of `CompositeAlarmArrayInput` via:
//
//	CompositeAlarmArray{ CompositeAlarmArgs{...} }
type CompositeAlarmArrayInput interface {
	pulumi.Input

	ToCompositeAlarmArrayOutput() CompositeAlarmArrayOutput
	ToCompositeAlarmArrayOutputWithContext(context.Context) CompositeAlarmArrayOutput
}

type CompositeAlarmArray []CompositeAlarmInput

func (CompositeAlarmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CompositeAlarm)(nil)).Elem()
}

func (i CompositeAlarmArray) ToCompositeAlarmArrayOutput() CompositeAlarmArrayOutput {
	return i.ToCompositeAlarmArrayOutputWithContext(context.Background())
}

func (i CompositeAlarmArray) ToCompositeAlarmArrayOutputWithContext(ctx context.Context) CompositeAlarmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompositeAlarmArrayOutput)
}

// CompositeAlarmMapInput is an input type that accepts CompositeAlarmMap and CompositeAlarmMapOutput values.
// You can construct a concrete instance of `CompositeAlarmMapInput` via:
//
//	CompositeAlarmMap{ "key": CompositeAlarmArgs{...} }
type CompositeAlarmMapInput interface {
	pulumi.Input

	ToCompositeAlarmMapOutput() CompositeAlarmMapOutput
	ToCompositeAlarmMapOutputWithContext(context.Context) CompositeAlarmMapOutput
}

type CompositeAlarmMap map[string]CompositeAlarmInput

func (CompositeAlarmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CompositeAlarm)(nil)).Elem()
}

func (i CompositeAlarmMap) ToCompositeAlarmMapOutput() CompositeAlarmMapOutput {
	return i.ToCompositeAlarmMapOutputWithContext(context.Background())
}

func (i CompositeAlarmMap) ToCompositeAlarmMapOutputWithContext(ctx context.Context) CompositeAlarmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompositeAlarmMapOutput)
}

type CompositeAlarmOutput struct{ *pulumi.OutputState }

func (CompositeAlarmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CompositeAlarm)(nil)).Elem()
}

func (o CompositeAlarmOutput) ToCompositeAlarmOutput() CompositeAlarmOutput {
	return o
}

func (o CompositeAlarmOutput) ToCompositeAlarmOutputWithContext(ctx context.Context) CompositeAlarmOutput {
	return o
}

// Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to `true`.
func (o CompositeAlarmOutput) ActionsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CompositeAlarm) pulumi.BoolPtrOutput { return v.ActionsEnabled }).(pulumi.BoolPtrOutput)
}

// The set of actions to execute when this alarm transitions to the `ALARM` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
func (o CompositeAlarmOutput) AlarmActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CompositeAlarm) pulumi.StringArrayOutput { return v.AlarmActions }).(pulumi.StringArrayOutput)
}

// The description for the composite alarm.
func (o CompositeAlarmOutput) AlarmDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CompositeAlarm) pulumi.StringPtrOutput { return v.AlarmDescription }).(pulumi.StringPtrOutput)
}

// The name for the composite alarm. This name must be unique within the region.
func (o CompositeAlarmOutput) AlarmName() pulumi.StringOutput {
	return o.ApplyT(func(v *CompositeAlarm) pulumi.StringOutput { return v.AlarmName }).(pulumi.StringOutput)
}

// An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see [Creating a Composite Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Composite_Alarm.html). The maximum length is 10240 characters.
func (o CompositeAlarmOutput) AlarmRule() pulumi.StringOutput {
	return o.ApplyT(func(v *CompositeAlarm) pulumi.StringOutput { return v.AlarmRule }).(pulumi.StringOutput)
}

// The ARN of the composite alarm.
func (o CompositeAlarmOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *CompositeAlarm) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The set of actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
func (o CompositeAlarmOutput) InsufficientDataActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CompositeAlarm) pulumi.StringArrayOutput { return v.InsufficientDataActions }).(pulumi.StringArrayOutput)
}

// The set of actions to execute when this alarm transitions to an `OK` state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
func (o CompositeAlarmOutput) OkActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CompositeAlarm) pulumi.StringArrayOutput { return v.OkActions }).(pulumi.StringArrayOutput)
}

// A map of tags to associate with the alarm. Up to 50 tags are allowed. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o CompositeAlarmOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CompositeAlarm) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o CompositeAlarmOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CompositeAlarm) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type CompositeAlarmArrayOutput struct{ *pulumi.OutputState }

func (CompositeAlarmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CompositeAlarm)(nil)).Elem()
}

func (o CompositeAlarmArrayOutput) ToCompositeAlarmArrayOutput() CompositeAlarmArrayOutput {
	return o
}

func (o CompositeAlarmArrayOutput) ToCompositeAlarmArrayOutputWithContext(ctx context.Context) CompositeAlarmArrayOutput {
	return o
}

func (o CompositeAlarmArrayOutput) Index(i pulumi.IntInput) CompositeAlarmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CompositeAlarm {
		return vs[0].([]*CompositeAlarm)[vs[1].(int)]
	}).(CompositeAlarmOutput)
}

type CompositeAlarmMapOutput struct{ *pulumi.OutputState }

func (CompositeAlarmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CompositeAlarm)(nil)).Elem()
}

func (o CompositeAlarmMapOutput) ToCompositeAlarmMapOutput() CompositeAlarmMapOutput {
	return o
}

func (o CompositeAlarmMapOutput) ToCompositeAlarmMapOutputWithContext(ctx context.Context) CompositeAlarmMapOutput {
	return o
}

func (o CompositeAlarmMapOutput) MapIndex(k pulumi.StringInput) CompositeAlarmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CompositeAlarm {
		return vs[0].(map[string]*CompositeAlarm)[vs[1].(string)]
	}).(CompositeAlarmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CompositeAlarmInput)(nil)).Elem(), &CompositeAlarm{})
	pulumi.RegisterInputType(reflect.TypeOf((*CompositeAlarmArrayInput)(nil)).Elem(), CompositeAlarmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CompositeAlarmMapInput)(nil)).Elem(), CompositeAlarmMap{})
	pulumi.RegisterOutputType(CompositeAlarmOutput{})
	pulumi.RegisterOutputType(CompositeAlarmArrayOutput{})
	pulumi.RegisterOutputType(CompositeAlarmMapOutput{})
}
