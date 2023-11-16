// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package evidently

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CloudWatch Evidently Segment resource.
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/evidently"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := evidently.NewSegment(ctx, "example", &evidently.SegmentArgs{
//				Pattern: pulumi.String("{\"Price\":[{\"numeric\":[\">\",10,\"<=\",20]}]}"),
//				Tags: pulumi.StringMap{
//					"Key1": pulumi.String("example Segment"),
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
// ### With JSON object in pattern
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/evidently"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := evidently.NewSegment(ctx, "example", &evidently.SegmentArgs{
//				Pattern: pulumi.String(`  {
//	    "Price": [
//	      {
//	        "numeric": [">",10,"<=",20]
//	      }
//	    ]
//	  }
//
// `),
//
//				Tags: pulumi.StringMap{
//					"Key1": pulumi.String("example Segment"),
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
// ### With Description
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/evidently"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := evidently.NewSegment(ctx, "example", &evidently.SegmentArgs{
//				Description: pulumi.String("example"),
//				Pattern:     pulumi.String("{\"Price\":[{\"numeric\":[\">\",10,\"<=\",20]}]}"),
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
// Using `pulumi import`, import CloudWatch Evidently Segment using the `arn`. For example:
//
// ```sh
//
//	$ pulumi import aws:evidently/segment:Segment example arn:aws:evidently:us-west-2:123456789012:segment/example
//
// ```
type Segment struct {
	pulumi.CustomResourceState

	// The ARN of the segment.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The date and time that the segment is created.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// Specifies the description of the segment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The number of experiments that this segment is used in. This count includes all current experiments, not just those that are currently running.
	ExperimentCount pulumi.IntOutput `pulumi:"experimentCount"`
	// The date and time that this segment was most recently updated.
	LastUpdatedTime pulumi.StringOutput `pulumi:"lastUpdatedTime"`
	// The number of launches that this segment is used in. This count includes all current launches, not just those that are currently running.
	LaunchCount pulumi.IntOutput `pulumi:"launchCount"`
	// A name for the segment.
	Name pulumi.StringOutput `pulumi:"name"`
	// The pattern to use for the segment. For more information about pattern syntax, see [Segment rule pattern syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments.html#CloudWatch-Evidently-segments-syntax.html).
	Pattern pulumi.StringOutput `pulumi:"pattern"`
	// Tags to apply to the segment. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewSegment registers a new resource with the given unique name, arguments, and options.
func NewSegment(ctx *pulumi.Context,
	name string, args *SegmentArgs, opts ...pulumi.ResourceOption) (*Segment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Pattern == nil {
		return nil, errors.New("invalid value for required argument 'Pattern'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Segment
	err := ctx.RegisterResource("aws:evidently/segment:Segment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSegment gets an existing Segment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSegment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SegmentState, opts ...pulumi.ResourceOption) (*Segment, error) {
	var resource Segment
	err := ctx.ReadResource("aws:evidently/segment:Segment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Segment resources.
type segmentState struct {
	// The ARN of the segment.
	Arn *string `pulumi:"arn"`
	// The date and time that the segment is created.
	CreatedTime *string `pulumi:"createdTime"`
	// Specifies the description of the segment.
	Description *string `pulumi:"description"`
	// The number of experiments that this segment is used in. This count includes all current experiments, not just those that are currently running.
	ExperimentCount *int `pulumi:"experimentCount"`
	// The date and time that this segment was most recently updated.
	LastUpdatedTime *string `pulumi:"lastUpdatedTime"`
	// The number of launches that this segment is used in. This count includes all current launches, not just those that are currently running.
	LaunchCount *int `pulumi:"launchCount"`
	// A name for the segment.
	Name *string `pulumi:"name"`
	// The pattern to use for the segment. For more information about pattern syntax, see [Segment rule pattern syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments.html#CloudWatch-Evidently-segments-syntax.html).
	Pattern *string `pulumi:"pattern"`
	// Tags to apply to the segment. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type SegmentState struct {
	// The ARN of the segment.
	Arn pulumi.StringPtrInput
	// The date and time that the segment is created.
	CreatedTime pulumi.StringPtrInput
	// Specifies the description of the segment.
	Description pulumi.StringPtrInput
	// The number of experiments that this segment is used in. This count includes all current experiments, not just those that are currently running.
	ExperimentCount pulumi.IntPtrInput
	// The date and time that this segment was most recently updated.
	LastUpdatedTime pulumi.StringPtrInput
	// The number of launches that this segment is used in. This count includes all current launches, not just those that are currently running.
	LaunchCount pulumi.IntPtrInput
	// A name for the segment.
	Name pulumi.StringPtrInput
	// The pattern to use for the segment. For more information about pattern syntax, see [Segment rule pattern syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments.html#CloudWatch-Evidently-segments-syntax.html).
	Pattern pulumi.StringPtrInput
	// Tags to apply to the segment. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (SegmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*segmentState)(nil)).Elem()
}

type segmentArgs struct {
	// Specifies the description of the segment.
	Description *string `pulumi:"description"`
	// A name for the segment.
	Name *string `pulumi:"name"`
	// The pattern to use for the segment. For more information about pattern syntax, see [Segment rule pattern syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments.html#CloudWatch-Evidently-segments-syntax.html).
	Pattern string `pulumi:"pattern"`
	// Tags to apply to the segment. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Segment resource.
type SegmentArgs struct {
	// Specifies the description of the segment.
	Description pulumi.StringPtrInput
	// A name for the segment.
	Name pulumi.StringPtrInput
	// The pattern to use for the segment. For more information about pattern syntax, see [Segment rule pattern syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments.html#CloudWatch-Evidently-segments-syntax.html).
	Pattern pulumi.StringInput
	// Tags to apply to the segment. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (SegmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*segmentArgs)(nil)).Elem()
}

type SegmentInput interface {
	pulumi.Input

	ToSegmentOutput() SegmentOutput
	ToSegmentOutputWithContext(ctx context.Context) SegmentOutput
}

func (*Segment) ElementType() reflect.Type {
	return reflect.TypeOf((**Segment)(nil)).Elem()
}

func (i *Segment) ToSegmentOutput() SegmentOutput {
	return i.ToSegmentOutputWithContext(context.Background())
}

func (i *Segment) ToSegmentOutputWithContext(ctx context.Context) SegmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SegmentOutput)
}

// SegmentArrayInput is an input type that accepts SegmentArray and SegmentArrayOutput values.
// You can construct a concrete instance of `SegmentArrayInput` via:
//
//	SegmentArray{ SegmentArgs{...} }
type SegmentArrayInput interface {
	pulumi.Input

	ToSegmentArrayOutput() SegmentArrayOutput
	ToSegmentArrayOutputWithContext(context.Context) SegmentArrayOutput
}

type SegmentArray []SegmentInput

func (SegmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Segment)(nil)).Elem()
}

func (i SegmentArray) ToSegmentArrayOutput() SegmentArrayOutput {
	return i.ToSegmentArrayOutputWithContext(context.Background())
}

func (i SegmentArray) ToSegmentArrayOutputWithContext(ctx context.Context) SegmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SegmentArrayOutput)
}

// SegmentMapInput is an input type that accepts SegmentMap and SegmentMapOutput values.
// You can construct a concrete instance of `SegmentMapInput` via:
//
//	SegmentMap{ "key": SegmentArgs{...} }
type SegmentMapInput interface {
	pulumi.Input

	ToSegmentMapOutput() SegmentMapOutput
	ToSegmentMapOutputWithContext(context.Context) SegmentMapOutput
}

type SegmentMap map[string]SegmentInput

func (SegmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Segment)(nil)).Elem()
}

func (i SegmentMap) ToSegmentMapOutput() SegmentMapOutput {
	return i.ToSegmentMapOutputWithContext(context.Background())
}

func (i SegmentMap) ToSegmentMapOutputWithContext(ctx context.Context) SegmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SegmentMapOutput)
}

type SegmentOutput struct{ *pulumi.OutputState }

func (SegmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Segment)(nil)).Elem()
}

func (o SegmentOutput) ToSegmentOutput() SegmentOutput {
	return o
}

func (o SegmentOutput) ToSegmentOutputWithContext(ctx context.Context) SegmentOutput {
	return o
}

// The ARN of the segment.
func (o SegmentOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Segment) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The date and time that the segment is created.
func (o SegmentOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Segment) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// Specifies the description of the segment.
func (o SegmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Segment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The number of experiments that this segment is used in. This count includes all current experiments, not just those that are currently running.
func (o SegmentOutput) ExperimentCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Segment) pulumi.IntOutput { return v.ExperimentCount }).(pulumi.IntOutput)
}

// The date and time that this segment was most recently updated.
func (o SegmentOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Segment) pulumi.StringOutput { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

// The number of launches that this segment is used in. This count includes all current launches, not just those that are currently running.
func (o SegmentOutput) LaunchCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Segment) pulumi.IntOutput { return v.LaunchCount }).(pulumi.IntOutput)
}

// A name for the segment.
func (o SegmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Segment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The pattern to use for the segment. For more information about pattern syntax, see [Segment rule pattern syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments.html#CloudWatch-Evidently-segments-syntax.html).
func (o SegmentOutput) Pattern() pulumi.StringOutput {
	return o.ApplyT(func(v *Segment) pulumi.StringOutput { return v.Pattern }).(pulumi.StringOutput)
}

// Tags to apply to the segment. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o SegmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Segment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o SegmentOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Segment) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type SegmentArrayOutput struct{ *pulumi.OutputState }

func (SegmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Segment)(nil)).Elem()
}

func (o SegmentArrayOutput) ToSegmentArrayOutput() SegmentArrayOutput {
	return o
}

func (o SegmentArrayOutput) ToSegmentArrayOutputWithContext(ctx context.Context) SegmentArrayOutput {
	return o
}

func (o SegmentArrayOutput) Index(i pulumi.IntInput) SegmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Segment {
		return vs[0].([]*Segment)[vs[1].(int)]
	}).(SegmentOutput)
}

type SegmentMapOutput struct{ *pulumi.OutputState }

func (SegmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Segment)(nil)).Elem()
}

func (o SegmentMapOutput) ToSegmentMapOutput() SegmentMapOutput {
	return o
}

func (o SegmentMapOutput) ToSegmentMapOutputWithContext(ctx context.Context) SegmentMapOutput {
	return o
}

func (o SegmentMapOutput) MapIndex(k pulumi.StringInput) SegmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Segment {
		return vs[0].(map[string]*Segment)[vs[1].(string)]
	}).(SegmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SegmentInput)(nil)).Elem(), &Segment{})
	pulumi.RegisterInputType(reflect.TypeOf((*SegmentArrayInput)(nil)).Elem(), SegmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SegmentMapInput)(nil)).Elem(), SegmentMap{})
	pulumi.RegisterOutputType(SegmentOutput{})
	pulumi.RegisterOutputType(SegmentArrayOutput{})
	pulumi.RegisterOutputType(SegmentMapOutput{})
}
