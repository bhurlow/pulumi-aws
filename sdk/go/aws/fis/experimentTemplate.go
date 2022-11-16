// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an FIS Experiment Template, which can be used to run an experiment.
// An experiment template contains one or more actions to run on specified targets during an experiment.
// It also contains the stop conditions that prevent the experiment from going out of bounds.
// See [Amazon Fault Injection Simulator](https://docs.aws.amazon.com/fis/index.html)
// for more information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/fis"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fis.NewExperimentTemplate(ctx, "example", &fis.ExperimentTemplateArgs{
//				Description: pulumi.String("example"),
//				RoleArn:     pulumi.Any(aws_iam_role.Example.Arn),
//				StopConditions: fis.ExperimentTemplateStopConditionArray{
//					&fis.ExperimentTemplateStopConditionArgs{
//						Source: pulumi.String("none"),
//					},
//				},
//				Actions: fis.ExperimentTemplateActionArray{
//					&fis.ExperimentTemplateActionArgs{
//						Name:     pulumi.String("example-action"),
//						ActionId: pulumi.String("aws:ec2:terminate-instances"),
//						Target: &fis.ExperimentTemplateActionTargetArgs{
//							Key:   pulumi.String("Instances"),
//							Value: pulumi.String("example-target"),
//						},
//					},
//				},
//				Targets: fis.ExperimentTemplateTargetArray{
//					&fis.ExperimentTemplateTargetArgs{
//						Name:          pulumi.String("example-target"),
//						ResourceType:  pulumi.String("aws:ec2:instance"),
//						SelectionMode: pulumi.String("COUNT(1)"),
//						ResourceTags: fis.ExperimentTemplateTargetResourceTagArray{
//							&fis.ExperimentTemplateTargetResourceTagArgs{
//								Key:   pulumi.String("env"),
//								Value: pulumi.String("example"),
//							},
//						},
//					},
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
// FIS Experiment Templates can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import aws:fis/experimentTemplate:ExperimentTemplate template EXT123AbCdEfGhIjK
//
// ```
type ExperimentTemplate struct {
	pulumi.CustomResourceState

	// Action to be performed during an experiment. See below.
	Actions ExperimentTemplateActionArrayOutput `pulumi:"actions"`
	// Description of the action.
	Description pulumi.StringOutput `pulumi:"description"`
	// ARN of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// When an ongoing experiment should be stopped. See below.
	StopConditions ExperimentTemplateStopConditionArrayOutput `pulumi:"stopConditions"`
	// Key-value mapping of tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    pulumi.StringMapOutput `pulumi:"tags"`
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Action's target, if applicable. See below.
	Targets ExperimentTemplateTargetArrayOutput `pulumi:"targets"`
}

// NewExperimentTemplate registers a new resource with the given unique name, arguments, and options.
func NewExperimentTemplate(ctx *pulumi.Context,
	name string, args *ExperimentTemplateArgs, opts ...pulumi.ResourceOption) (*ExperimentTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	if args.StopConditions == nil {
		return nil, errors.New("invalid value for required argument 'StopConditions'")
	}
	var resource ExperimentTemplate
	err := ctx.RegisterResource("aws:fis/experimentTemplate:ExperimentTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExperimentTemplate gets an existing ExperimentTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExperimentTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExperimentTemplateState, opts ...pulumi.ResourceOption) (*ExperimentTemplate, error) {
	var resource ExperimentTemplate
	err := ctx.ReadResource("aws:fis/experimentTemplate:ExperimentTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExperimentTemplate resources.
type experimentTemplateState struct {
	// Action to be performed during an experiment. See below.
	Actions []ExperimentTemplateAction `pulumi:"actions"`
	// Description of the action.
	Description *string `pulumi:"description"`
	// ARN of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
	RoleArn *string `pulumi:"roleArn"`
	// When an ongoing experiment should be stopped. See below.
	StopConditions []ExperimentTemplateStopCondition `pulumi:"stopConditions"`
	// Key-value mapping of tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    map[string]string `pulumi:"tags"`
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Action's target, if applicable. See below.
	Targets []ExperimentTemplateTarget `pulumi:"targets"`
}

type ExperimentTemplateState struct {
	// Action to be performed during an experiment. See below.
	Actions ExperimentTemplateActionArrayInput
	// Description of the action.
	Description pulumi.StringPtrInput
	// ARN of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
	RoleArn pulumi.StringPtrInput
	// When an ongoing experiment should be stopped. See below.
	StopConditions ExperimentTemplateStopConditionArrayInput
	// Key-value mapping of tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    pulumi.StringMapInput
	TagsAll pulumi.StringMapInput
	// Action's target, if applicable. See below.
	Targets ExperimentTemplateTargetArrayInput
}

func (ExperimentTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*experimentTemplateState)(nil)).Elem()
}

type experimentTemplateArgs struct {
	// Action to be performed during an experiment. See below.
	Actions []ExperimentTemplateAction `pulumi:"actions"`
	// Description of the action.
	Description string `pulumi:"description"`
	// ARN of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
	RoleArn string `pulumi:"roleArn"`
	// When an ongoing experiment should be stopped. See below.
	StopConditions []ExperimentTemplateStopCondition `pulumi:"stopConditions"`
	// Key-value mapping of tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Action's target, if applicable. See below.
	Targets []ExperimentTemplateTarget `pulumi:"targets"`
}

// The set of arguments for constructing a ExperimentTemplate resource.
type ExperimentTemplateArgs struct {
	// Action to be performed during an experiment. See below.
	Actions ExperimentTemplateActionArrayInput
	// Description of the action.
	Description pulumi.StringInput
	// ARN of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
	RoleArn pulumi.StringInput
	// When an ongoing experiment should be stopped. See below.
	StopConditions ExperimentTemplateStopConditionArrayInput
	// Key-value mapping of tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Action's target, if applicable. See below.
	Targets ExperimentTemplateTargetArrayInput
}

func (ExperimentTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*experimentTemplateArgs)(nil)).Elem()
}

type ExperimentTemplateInput interface {
	pulumi.Input

	ToExperimentTemplateOutput() ExperimentTemplateOutput
	ToExperimentTemplateOutputWithContext(ctx context.Context) ExperimentTemplateOutput
}

func (*ExperimentTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**ExperimentTemplate)(nil)).Elem()
}

func (i *ExperimentTemplate) ToExperimentTemplateOutput() ExperimentTemplateOutput {
	return i.ToExperimentTemplateOutputWithContext(context.Background())
}

func (i *ExperimentTemplate) ToExperimentTemplateOutputWithContext(ctx context.Context) ExperimentTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentTemplateOutput)
}

// ExperimentTemplateArrayInput is an input type that accepts ExperimentTemplateArray and ExperimentTemplateArrayOutput values.
// You can construct a concrete instance of `ExperimentTemplateArrayInput` via:
//
//	ExperimentTemplateArray{ ExperimentTemplateArgs{...} }
type ExperimentTemplateArrayInput interface {
	pulumi.Input

	ToExperimentTemplateArrayOutput() ExperimentTemplateArrayOutput
	ToExperimentTemplateArrayOutputWithContext(context.Context) ExperimentTemplateArrayOutput
}

type ExperimentTemplateArray []ExperimentTemplateInput

func (ExperimentTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExperimentTemplate)(nil)).Elem()
}

func (i ExperimentTemplateArray) ToExperimentTemplateArrayOutput() ExperimentTemplateArrayOutput {
	return i.ToExperimentTemplateArrayOutputWithContext(context.Background())
}

func (i ExperimentTemplateArray) ToExperimentTemplateArrayOutputWithContext(ctx context.Context) ExperimentTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentTemplateArrayOutput)
}

// ExperimentTemplateMapInput is an input type that accepts ExperimentTemplateMap and ExperimentTemplateMapOutput values.
// You can construct a concrete instance of `ExperimentTemplateMapInput` via:
//
//	ExperimentTemplateMap{ "key": ExperimentTemplateArgs{...} }
type ExperimentTemplateMapInput interface {
	pulumi.Input

	ToExperimentTemplateMapOutput() ExperimentTemplateMapOutput
	ToExperimentTemplateMapOutputWithContext(context.Context) ExperimentTemplateMapOutput
}

type ExperimentTemplateMap map[string]ExperimentTemplateInput

func (ExperimentTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExperimentTemplate)(nil)).Elem()
}

func (i ExperimentTemplateMap) ToExperimentTemplateMapOutput() ExperimentTemplateMapOutput {
	return i.ToExperimentTemplateMapOutputWithContext(context.Background())
}

func (i ExperimentTemplateMap) ToExperimentTemplateMapOutputWithContext(ctx context.Context) ExperimentTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentTemplateMapOutput)
}

type ExperimentTemplateOutput struct{ *pulumi.OutputState }

func (ExperimentTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExperimentTemplate)(nil)).Elem()
}

func (o ExperimentTemplateOutput) ToExperimentTemplateOutput() ExperimentTemplateOutput {
	return o
}

func (o ExperimentTemplateOutput) ToExperimentTemplateOutputWithContext(ctx context.Context) ExperimentTemplateOutput {
	return o
}

// Action to be performed during an experiment. See below.
func (o ExperimentTemplateOutput) Actions() ExperimentTemplateActionArrayOutput {
	return o.ApplyT(func(v *ExperimentTemplate) ExperimentTemplateActionArrayOutput { return v.Actions }).(ExperimentTemplateActionArrayOutput)
}

// Description of the action.
func (o ExperimentTemplateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ExperimentTemplate) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// ARN of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
func (o ExperimentTemplateOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ExperimentTemplate) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// When an ongoing experiment should be stopped. See below.
func (o ExperimentTemplateOutput) StopConditions() ExperimentTemplateStopConditionArrayOutput {
	return o.ApplyT(func(v *ExperimentTemplate) ExperimentTemplateStopConditionArrayOutput { return v.StopConditions }).(ExperimentTemplateStopConditionArrayOutput)
}

// Key-value mapping of tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ExperimentTemplateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ExperimentTemplate) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ExperimentTemplateOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ExperimentTemplate) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Action's target, if applicable. See below.
func (o ExperimentTemplateOutput) Targets() ExperimentTemplateTargetArrayOutput {
	return o.ApplyT(func(v *ExperimentTemplate) ExperimentTemplateTargetArrayOutput { return v.Targets }).(ExperimentTemplateTargetArrayOutput)
}

type ExperimentTemplateArrayOutput struct{ *pulumi.OutputState }

func (ExperimentTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExperimentTemplate)(nil)).Elem()
}

func (o ExperimentTemplateArrayOutput) ToExperimentTemplateArrayOutput() ExperimentTemplateArrayOutput {
	return o
}

func (o ExperimentTemplateArrayOutput) ToExperimentTemplateArrayOutputWithContext(ctx context.Context) ExperimentTemplateArrayOutput {
	return o
}

func (o ExperimentTemplateArrayOutput) Index(i pulumi.IntInput) ExperimentTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExperimentTemplate {
		return vs[0].([]*ExperimentTemplate)[vs[1].(int)]
	}).(ExperimentTemplateOutput)
}

type ExperimentTemplateMapOutput struct{ *pulumi.OutputState }

func (ExperimentTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExperimentTemplate)(nil)).Elem()
}

func (o ExperimentTemplateMapOutput) ToExperimentTemplateMapOutput() ExperimentTemplateMapOutput {
	return o
}

func (o ExperimentTemplateMapOutput) ToExperimentTemplateMapOutputWithContext(ctx context.Context) ExperimentTemplateMapOutput {
	return o
}

func (o ExperimentTemplateMapOutput) MapIndex(k pulumi.StringInput) ExperimentTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExperimentTemplate {
		return vs[0].(map[string]*ExperimentTemplate)[vs[1].(string)]
	}).(ExperimentTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExperimentTemplateInput)(nil)).Elem(), &ExperimentTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExperimentTemplateArrayInput)(nil)).Elem(), ExperimentTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExperimentTemplateMapInput)(nil)).Elem(), ExperimentTemplateMap{})
	pulumi.RegisterOutputType(ExperimentTemplateOutput{})
	pulumi.RegisterOutputType(ExperimentTemplateArrayOutput{})
	pulumi.RegisterOutputType(ExperimentTemplateMapOutput{})
}
