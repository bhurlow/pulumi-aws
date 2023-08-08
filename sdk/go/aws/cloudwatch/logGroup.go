// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CloudWatch Log Group resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudwatch.NewLogGroup(ctx, "yada", &cloudwatch.LogGroupArgs{
//				Tags: pulumi.StringMap{
//					"Application": pulumi.String("serviceA"),
//					"Environment": pulumi.String("production"),
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
// terraform import {
//
//	to = aws_cloudwatch_log_group.test_group
//
//	id = "yada" } Using `pulumi import`, import Cloudwatch Log Groups using the `name`. For exampleconsole % pulumi import aws_cloudwatch_log_group.test_group yada
type LogGroup struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) specifying the log group. Any `:*` suffix added by the API, denoting all CloudWatch Log Streams under the CloudWatch Log Group, is removed for greater compatibility with other AWS services that do not accept the suffix.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ARN of the KMS Key to use when encrypting log data. Please note, after the AWS KMS CMK is disassociated from the log group,
	// AWS CloudWatch Logs stops encrypting newly ingested data for the log group. All previously ingested data remains encrypted, and AWS CloudWatch Logs requires
	// permissions for the CMK whenever the encrypted data is requested.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// The name of the log group. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// Specifies the number of days
	// you want to retain log events in the specified log group.  Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1096, 1827, 2192, 2557, 2922, 3288, 3653, and 0.
	// If you select 0, the events in the log group are always retained and never expire.
	RetentionInDays pulumi.IntPtrOutput `pulumi:"retentionInDays"`
	// Set to true if you do not wish the log group (and any logs it may contain) to be deleted at destroy time, and instead just remove the log group from the state.
	SkipDestroy pulumi.BoolPtrOutput `pulumi:"skipDestroy"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewLogGroup registers a new resource with the given unique name, arguments, and options.
func NewLogGroup(ctx *pulumi.Context,
	name string, args *LogGroupArgs, opts ...pulumi.ResourceOption) (*LogGroup, error) {
	if args == nil {
		args = &LogGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LogGroup
	err := ctx.RegisterResource("aws:cloudwatch/logGroup:LogGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogGroup gets an existing LogGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogGroupState, opts ...pulumi.ResourceOption) (*LogGroup, error) {
	var resource LogGroup
	err := ctx.ReadResource("aws:cloudwatch/logGroup:LogGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogGroup resources.
type logGroupState struct {
	// The Amazon Resource Name (ARN) specifying the log group. Any `:*` suffix added by the API, denoting all CloudWatch Log Streams under the CloudWatch Log Group, is removed for greater compatibility with other AWS services that do not accept the suffix.
	Arn *string `pulumi:"arn"`
	// The ARN of the KMS Key to use when encrypting log data. Please note, after the AWS KMS CMK is disassociated from the log group,
	// AWS CloudWatch Logs stops encrypting newly ingested data for the log group. All previously ingested data remains encrypted, and AWS CloudWatch Logs requires
	// permissions for the CMK whenever the encrypted data is requested.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The name of the log group. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Specifies the number of days
	// you want to retain log events in the specified log group.  Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1096, 1827, 2192, 2557, 2922, 3288, 3653, and 0.
	// If you select 0, the events in the log group are always retained and never expire.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// Set to true if you do not wish the log group (and any logs it may contain) to be deleted at destroy time, and instead just remove the log group from the state.
	SkipDestroy *bool `pulumi:"skipDestroy"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type LogGroupState struct {
	// The Amazon Resource Name (ARN) specifying the log group. Any `:*` suffix added by the API, denoting all CloudWatch Log Streams under the CloudWatch Log Group, is removed for greater compatibility with other AWS services that do not accept the suffix.
	Arn pulumi.StringPtrInput
	// The ARN of the KMS Key to use when encrypting log data. Please note, after the AWS KMS CMK is disassociated from the log group,
	// AWS CloudWatch Logs stops encrypting newly ingested data for the log group. All previously ingested data remains encrypted, and AWS CloudWatch Logs requires
	// permissions for the CMK whenever the encrypted data is requested.
	KmsKeyId pulumi.StringPtrInput
	// The name of the log group. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Specifies the number of days
	// you want to retain log events in the specified log group.  Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1096, 1827, 2192, 2557, 2922, 3288, 3653, and 0.
	// If you select 0, the events in the log group are always retained and never expire.
	RetentionInDays pulumi.IntPtrInput
	// Set to true if you do not wish the log group (and any logs it may contain) to be deleted at destroy time, and instead just remove the log group from the state.
	SkipDestroy pulumi.BoolPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (LogGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*logGroupState)(nil)).Elem()
}

type logGroupArgs struct {
	// The ARN of the KMS Key to use when encrypting log data. Please note, after the AWS KMS CMK is disassociated from the log group,
	// AWS CloudWatch Logs stops encrypting newly ingested data for the log group. All previously ingested data remains encrypted, and AWS CloudWatch Logs requires
	// permissions for the CMK whenever the encrypted data is requested.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The name of the log group. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Specifies the number of days
	// you want to retain log events in the specified log group.  Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1096, 1827, 2192, 2557, 2922, 3288, 3653, and 0.
	// If you select 0, the events in the log group are always retained and never expire.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// Set to true if you do not wish the log group (and any logs it may contain) to be deleted at destroy time, and instead just remove the log group from the state.
	SkipDestroy *bool `pulumi:"skipDestroy"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LogGroup resource.
type LogGroupArgs struct {
	// The ARN of the KMS Key to use when encrypting log data. Please note, after the AWS KMS CMK is disassociated from the log group,
	// AWS CloudWatch Logs stops encrypting newly ingested data for the log group. All previously ingested data remains encrypted, and AWS CloudWatch Logs requires
	// permissions for the CMK whenever the encrypted data is requested.
	KmsKeyId pulumi.StringPtrInput
	// The name of the log group. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Specifies the number of days
	// you want to retain log events in the specified log group.  Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1096, 1827, 2192, 2557, 2922, 3288, 3653, and 0.
	// If you select 0, the events in the log group are always retained and never expire.
	RetentionInDays pulumi.IntPtrInput
	// Set to true if you do not wish the log group (and any logs it may contain) to be deleted at destroy time, and instead just remove the log group from the state.
	SkipDestroy pulumi.BoolPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (LogGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logGroupArgs)(nil)).Elem()
}

type LogGroupInput interface {
	pulumi.Input

	ToLogGroupOutput() LogGroupOutput
	ToLogGroupOutputWithContext(ctx context.Context) LogGroupOutput
}

func (*LogGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**LogGroup)(nil)).Elem()
}

func (i *LogGroup) ToLogGroupOutput() LogGroupOutput {
	return i.ToLogGroupOutputWithContext(context.Background())
}

func (i *LogGroup) ToLogGroupOutputWithContext(ctx context.Context) LogGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogGroupOutput)
}

// LogGroupArrayInput is an input type that accepts LogGroupArray and LogGroupArrayOutput values.
// You can construct a concrete instance of `LogGroupArrayInput` via:
//
//	LogGroupArray{ LogGroupArgs{...} }
type LogGroupArrayInput interface {
	pulumi.Input

	ToLogGroupArrayOutput() LogGroupArrayOutput
	ToLogGroupArrayOutputWithContext(context.Context) LogGroupArrayOutput
}

type LogGroupArray []LogGroupInput

func (LogGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogGroup)(nil)).Elem()
}

func (i LogGroupArray) ToLogGroupArrayOutput() LogGroupArrayOutput {
	return i.ToLogGroupArrayOutputWithContext(context.Background())
}

func (i LogGroupArray) ToLogGroupArrayOutputWithContext(ctx context.Context) LogGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogGroupArrayOutput)
}

// LogGroupMapInput is an input type that accepts LogGroupMap and LogGroupMapOutput values.
// You can construct a concrete instance of `LogGroupMapInput` via:
//
//	LogGroupMap{ "key": LogGroupArgs{...} }
type LogGroupMapInput interface {
	pulumi.Input

	ToLogGroupMapOutput() LogGroupMapOutput
	ToLogGroupMapOutputWithContext(context.Context) LogGroupMapOutput
}

type LogGroupMap map[string]LogGroupInput

func (LogGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogGroup)(nil)).Elem()
}

func (i LogGroupMap) ToLogGroupMapOutput() LogGroupMapOutput {
	return i.ToLogGroupMapOutputWithContext(context.Background())
}

func (i LogGroupMap) ToLogGroupMapOutputWithContext(ctx context.Context) LogGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogGroupMapOutput)
}

type LogGroupOutput struct{ *pulumi.OutputState }

func (LogGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogGroup)(nil)).Elem()
}

func (o LogGroupOutput) ToLogGroupOutput() LogGroupOutput {
	return o
}

func (o LogGroupOutput) ToLogGroupOutputWithContext(ctx context.Context) LogGroupOutput {
	return o
}

// The Amazon Resource Name (ARN) specifying the log group. Any `:*` suffix added by the API, denoting all CloudWatch Log Streams under the CloudWatch Log Group, is removed for greater compatibility with other AWS services that do not accept the suffix.
func (o LogGroupOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *LogGroup) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The ARN of the KMS Key to use when encrypting log data. Please note, after the AWS KMS CMK is disassociated from the log group,
// AWS CloudWatch Logs stops encrypting newly ingested data for the log group. All previously ingested data remains encrypted, and AWS CloudWatch Logs requires
// permissions for the CMK whenever the encrypted data is requested.
func (o LogGroupOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogGroup) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// The name of the log group. If omitted, this provider will assign a random, unique name.
func (o LogGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LogGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
func (o LogGroupOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *LogGroup) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

// Specifies the number of days
// you want to retain log events in the specified log group.  Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1096, 1827, 2192, 2557, 2922, 3288, 3653, and 0.
// If you select 0, the events in the log group are always retained and never expire.
func (o LogGroupOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LogGroup) pulumi.IntPtrOutput { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

// Set to true if you do not wish the log group (and any logs it may contain) to be deleted at destroy time, and instead just remove the log group from the state.
func (o LogGroupOutput) SkipDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LogGroup) pulumi.BoolPtrOutput { return v.SkipDestroy }).(pulumi.BoolPtrOutput)
}

// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o LogGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LogGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o LogGroupOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LogGroup) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type LogGroupArrayOutput struct{ *pulumi.OutputState }

func (LogGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogGroup)(nil)).Elem()
}

func (o LogGroupArrayOutput) ToLogGroupArrayOutput() LogGroupArrayOutput {
	return o
}

func (o LogGroupArrayOutput) ToLogGroupArrayOutputWithContext(ctx context.Context) LogGroupArrayOutput {
	return o
}

func (o LogGroupArrayOutput) Index(i pulumi.IntInput) LogGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogGroup {
		return vs[0].([]*LogGroup)[vs[1].(int)]
	}).(LogGroupOutput)
}

type LogGroupMapOutput struct{ *pulumi.OutputState }

func (LogGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogGroup)(nil)).Elem()
}

func (o LogGroupMapOutput) ToLogGroupMapOutput() LogGroupMapOutput {
	return o
}

func (o LogGroupMapOutput) ToLogGroupMapOutputWithContext(ctx context.Context) LogGroupMapOutput {
	return o
}

func (o LogGroupMapOutput) MapIndex(k pulumi.StringInput) LogGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogGroup {
		return vs[0].(map[string]*LogGroup)[vs[1].(string)]
	}).(LogGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogGroupInput)(nil)).Elem(), &LogGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogGroupArrayInput)(nil)).Elem(), LogGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogGroupMapInput)(nil)).Elem(), LogGroupMap{})
	pulumi.RegisterOutputType(LogGroupOutput{})
	pulumi.RegisterOutputType(LogGroupArrayOutput{})
	pulumi.RegisterOutputType(LogGroupMapOutput{})
}
