// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Image Builder Image Pipeline.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/imagebuilder"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := imagebuilder.NewImagePipeline(ctx, "example", &imagebuilder.ImagePipelineArgs{
//				ImageRecipeArn:                 pulumi.Any(aws_imagebuilder_image_recipe.Example.Arn),
//				InfrastructureConfigurationArn: pulumi.Any(aws_imagebuilder_infrastructure_configuration.Example.Arn),
//				Schedule: &imagebuilder.ImagePipelineScheduleArgs{
//					ScheduleExpression: pulumi.String("cron(0 0 * * ? *)"),
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
// `aws_imagebuilder_image_pipeline` resources can be imported using the Amazon Resource Name (ARN), e.g.,
//
// ```sh
//
//	$ pulumi import aws:imagebuilder/imagePipeline:ImagePipeline example arn:aws:imagebuilder:us-east-1:123456789012:image-pipeline/example
//
// ```
type ImagePipeline struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the image pipeline.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Amazon Resource Name (ARN) of the container recipe.
	ContainerRecipeArn pulumi.StringPtrOutput `pulumi:"containerRecipeArn"`
	// Date the image pipeline was created.
	DateCreated pulumi.StringOutput `pulumi:"dateCreated"`
	// Date the image pipeline was last run.
	DateLastRun pulumi.StringOutput `pulumi:"dateLastRun"`
	// Date the image pipeline will run next.
	DateNextRun pulumi.StringOutput `pulumi:"dateNextRun"`
	// Date the image pipeline was updated.
	DateUpdated pulumi.StringOutput `pulumi:"dateUpdated"`
	// Description of the image pipeline.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
	DistributionConfigurationArn pulumi.StringPtrOutput `pulumi:"distributionConfigurationArn"`
	// Whether additional information about the image being created is collected. Defaults to `true`.
	EnhancedImageMetadataEnabled pulumi.BoolPtrOutput `pulumi:"enhancedImageMetadataEnabled"`
	// Amazon Resource Name (ARN) of the image recipe.
	ImageRecipeArn pulumi.StringPtrOutput `pulumi:"imageRecipeArn"`
	// Configuration block with image tests configuration. Detailed below.
	ImageTestsConfiguration ImagePipelineImageTestsConfigurationOutput `pulumi:"imageTestsConfiguration"`
	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
	InfrastructureConfigurationArn pulumi.StringOutput `pulumi:"infrastructureConfigurationArn"`
	// Name of the image pipeline.
	Name pulumi.StringOutput `pulumi:"name"`
	// Platform of the image pipeline.
	Platform pulumi.StringOutput `pulumi:"platform"`
	// Configuration block with schedule settings. Detailed below.
	Schedule ImagePipelineSchedulePtrOutput `pulumi:"schedule"`
	// Status of the image pipeline. Valid values are `DISABLED` and `ENABLED`. Defaults to `ENABLED`.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Key-value map of resource tags for the image pipeline. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewImagePipeline registers a new resource with the given unique name, arguments, and options.
func NewImagePipeline(ctx *pulumi.Context,
	name string, args *ImagePipelineArgs, opts ...pulumi.ResourceOption) (*ImagePipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InfrastructureConfigurationArn == nil {
		return nil, errors.New("invalid value for required argument 'InfrastructureConfigurationArn'")
	}
	var resource ImagePipeline
	err := ctx.RegisterResource("aws:imagebuilder/imagePipeline:ImagePipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImagePipeline gets an existing ImagePipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImagePipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImagePipelineState, opts ...pulumi.ResourceOption) (*ImagePipeline, error) {
	var resource ImagePipeline
	err := ctx.ReadResource("aws:imagebuilder/imagePipeline:ImagePipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImagePipeline resources.
type imagePipelineState struct {
	// Amazon Resource Name (ARN) of the image pipeline.
	Arn *string `pulumi:"arn"`
	// Amazon Resource Name (ARN) of the container recipe.
	ContainerRecipeArn *string `pulumi:"containerRecipeArn"`
	// Date the image pipeline was created.
	DateCreated *string `pulumi:"dateCreated"`
	// Date the image pipeline was last run.
	DateLastRun *string `pulumi:"dateLastRun"`
	// Date the image pipeline will run next.
	DateNextRun *string `pulumi:"dateNextRun"`
	// Date the image pipeline was updated.
	DateUpdated *string `pulumi:"dateUpdated"`
	// Description of the image pipeline.
	Description *string `pulumi:"description"`
	// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
	DistributionConfigurationArn *string `pulumi:"distributionConfigurationArn"`
	// Whether additional information about the image being created is collected. Defaults to `true`.
	EnhancedImageMetadataEnabled *bool `pulumi:"enhancedImageMetadataEnabled"`
	// Amazon Resource Name (ARN) of the image recipe.
	ImageRecipeArn *string `pulumi:"imageRecipeArn"`
	// Configuration block with image tests configuration. Detailed below.
	ImageTestsConfiguration *ImagePipelineImageTestsConfiguration `pulumi:"imageTestsConfiguration"`
	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
	InfrastructureConfigurationArn *string `pulumi:"infrastructureConfigurationArn"`
	// Name of the image pipeline.
	Name *string `pulumi:"name"`
	// Platform of the image pipeline.
	Platform *string `pulumi:"platform"`
	// Configuration block with schedule settings. Detailed below.
	Schedule *ImagePipelineSchedule `pulumi:"schedule"`
	// Status of the image pipeline. Valid values are `DISABLED` and `ENABLED`. Defaults to `ENABLED`.
	Status *string `pulumi:"status"`
	// Key-value map of resource tags for the image pipeline. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ImagePipelineState struct {
	// Amazon Resource Name (ARN) of the image pipeline.
	Arn pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the container recipe.
	ContainerRecipeArn pulumi.StringPtrInput
	// Date the image pipeline was created.
	DateCreated pulumi.StringPtrInput
	// Date the image pipeline was last run.
	DateLastRun pulumi.StringPtrInput
	// Date the image pipeline will run next.
	DateNextRun pulumi.StringPtrInput
	// Date the image pipeline was updated.
	DateUpdated pulumi.StringPtrInput
	// Description of the image pipeline.
	Description pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
	DistributionConfigurationArn pulumi.StringPtrInput
	// Whether additional information about the image being created is collected. Defaults to `true`.
	EnhancedImageMetadataEnabled pulumi.BoolPtrInput
	// Amazon Resource Name (ARN) of the image recipe.
	ImageRecipeArn pulumi.StringPtrInput
	// Configuration block with image tests configuration. Detailed below.
	ImageTestsConfiguration ImagePipelineImageTestsConfigurationPtrInput
	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
	InfrastructureConfigurationArn pulumi.StringPtrInput
	// Name of the image pipeline.
	Name pulumi.StringPtrInput
	// Platform of the image pipeline.
	Platform pulumi.StringPtrInput
	// Configuration block with schedule settings. Detailed below.
	Schedule ImagePipelineSchedulePtrInput
	// Status of the image pipeline. Valid values are `DISABLED` and `ENABLED`. Defaults to `ENABLED`.
	Status pulumi.StringPtrInput
	// Key-value map of resource tags for the image pipeline. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (ImagePipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*imagePipelineState)(nil)).Elem()
}

type imagePipelineArgs struct {
	// Amazon Resource Name (ARN) of the container recipe.
	ContainerRecipeArn *string `pulumi:"containerRecipeArn"`
	// Description of the image pipeline.
	Description *string `pulumi:"description"`
	// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
	DistributionConfigurationArn *string `pulumi:"distributionConfigurationArn"`
	// Whether additional information about the image being created is collected. Defaults to `true`.
	EnhancedImageMetadataEnabled *bool `pulumi:"enhancedImageMetadataEnabled"`
	// Amazon Resource Name (ARN) of the image recipe.
	ImageRecipeArn *string `pulumi:"imageRecipeArn"`
	// Configuration block with image tests configuration. Detailed below.
	ImageTestsConfiguration *ImagePipelineImageTestsConfiguration `pulumi:"imageTestsConfiguration"`
	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
	InfrastructureConfigurationArn string `pulumi:"infrastructureConfigurationArn"`
	// Name of the image pipeline.
	Name *string `pulumi:"name"`
	// Configuration block with schedule settings. Detailed below.
	Schedule *ImagePipelineSchedule `pulumi:"schedule"`
	// Status of the image pipeline. Valid values are `DISABLED` and `ENABLED`. Defaults to `ENABLED`.
	Status *string `pulumi:"status"`
	// Key-value map of resource tags for the image pipeline. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ImagePipeline resource.
type ImagePipelineArgs struct {
	// Amazon Resource Name (ARN) of the container recipe.
	ContainerRecipeArn pulumi.StringPtrInput
	// Description of the image pipeline.
	Description pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
	DistributionConfigurationArn pulumi.StringPtrInput
	// Whether additional information about the image being created is collected. Defaults to `true`.
	EnhancedImageMetadataEnabled pulumi.BoolPtrInput
	// Amazon Resource Name (ARN) of the image recipe.
	ImageRecipeArn pulumi.StringPtrInput
	// Configuration block with image tests configuration. Detailed below.
	ImageTestsConfiguration ImagePipelineImageTestsConfigurationPtrInput
	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
	InfrastructureConfigurationArn pulumi.StringInput
	// Name of the image pipeline.
	Name pulumi.StringPtrInput
	// Configuration block with schedule settings. Detailed below.
	Schedule ImagePipelineSchedulePtrInput
	// Status of the image pipeline. Valid values are `DISABLED` and `ENABLED`. Defaults to `ENABLED`.
	Status pulumi.StringPtrInput
	// Key-value map of resource tags for the image pipeline. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ImagePipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imagePipelineArgs)(nil)).Elem()
}

type ImagePipelineInput interface {
	pulumi.Input

	ToImagePipelineOutput() ImagePipelineOutput
	ToImagePipelineOutputWithContext(ctx context.Context) ImagePipelineOutput
}

func (*ImagePipeline) ElementType() reflect.Type {
	return reflect.TypeOf((**ImagePipeline)(nil)).Elem()
}

func (i *ImagePipeline) ToImagePipelineOutput() ImagePipelineOutput {
	return i.ToImagePipelineOutputWithContext(context.Background())
}

func (i *ImagePipeline) ToImagePipelineOutputWithContext(ctx context.Context) ImagePipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImagePipelineOutput)
}

// ImagePipelineArrayInput is an input type that accepts ImagePipelineArray and ImagePipelineArrayOutput values.
// You can construct a concrete instance of `ImagePipelineArrayInput` via:
//
//	ImagePipelineArray{ ImagePipelineArgs{...} }
type ImagePipelineArrayInput interface {
	pulumi.Input

	ToImagePipelineArrayOutput() ImagePipelineArrayOutput
	ToImagePipelineArrayOutputWithContext(context.Context) ImagePipelineArrayOutput
}

type ImagePipelineArray []ImagePipelineInput

func (ImagePipelineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImagePipeline)(nil)).Elem()
}

func (i ImagePipelineArray) ToImagePipelineArrayOutput() ImagePipelineArrayOutput {
	return i.ToImagePipelineArrayOutputWithContext(context.Background())
}

func (i ImagePipelineArray) ToImagePipelineArrayOutputWithContext(ctx context.Context) ImagePipelineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImagePipelineArrayOutput)
}

// ImagePipelineMapInput is an input type that accepts ImagePipelineMap and ImagePipelineMapOutput values.
// You can construct a concrete instance of `ImagePipelineMapInput` via:
//
//	ImagePipelineMap{ "key": ImagePipelineArgs{...} }
type ImagePipelineMapInput interface {
	pulumi.Input

	ToImagePipelineMapOutput() ImagePipelineMapOutput
	ToImagePipelineMapOutputWithContext(context.Context) ImagePipelineMapOutput
}

type ImagePipelineMap map[string]ImagePipelineInput

func (ImagePipelineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImagePipeline)(nil)).Elem()
}

func (i ImagePipelineMap) ToImagePipelineMapOutput() ImagePipelineMapOutput {
	return i.ToImagePipelineMapOutputWithContext(context.Background())
}

func (i ImagePipelineMap) ToImagePipelineMapOutputWithContext(ctx context.Context) ImagePipelineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImagePipelineMapOutput)
}

type ImagePipelineOutput struct{ *pulumi.OutputState }

func (ImagePipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImagePipeline)(nil)).Elem()
}

func (o ImagePipelineOutput) ToImagePipelineOutput() ImagePipelineOutput {
	return o
}

func (o ImagePipelineOutput) ToImagePipelineOutputWithContext(ctx context.Context) ImagePipelineOutput {
	return o
}

// Amazon Resource Name (ARN) of the image pipeline.
func (o ImagePipelineOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagePipeline) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Amazon Resource Name (ARN) of the container recipe.
func (o ImagePipelineOutput) ContainerRecipeArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImagePipeline) pulumi.StringPtrOutput { return v.ContainerRecipeArn }).(pulumi.StringPtrOutput)
}

// Date the image pipeline was created.
func (o ImagePipelineOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagePipeline) pulumi.StringOutput { return v.DateCreated }).(pulumi.StringOutput)
}

// Date the image pipeline was last run.
func (o ImagePipelineOutput) DateLastRun() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagePipeline) pulumi.StringOutput { return v.DateLastRun }).(pulumi.StringOutput)
}

// Date the image pipeline will run next.
func (o ImagePipelineOutput) DateNextRun() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagePipeline) pulumi.StringOutput { return v.DateNextRun }).(pulumi.StringOutput)
}

// Date the image pipeline was updated.
func (o ImagePipelineOutput) DateUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagePipeline) pulumi.StringOutput { return v.DateUpdated }).(pulumi.StringOutput)
}

// Description of the image pipeline.
func (o ImagePipelineOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImagePipeline) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
func (o ImagePipelineOutput) DistributionConfigurationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImagePipeline) pulumi.StringPtrOutput { return v.DistributionConfigurationArn }).(pulumi.StringPtrOutput)
}

// Whether additional information about the image being created is collected. Defaults to `true`.
func (o ImagePipelineOutput) EnhancedImageMetadataEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ImagePipeline) pulumi.BoolPtrOutput { return v.EnhancedImageMetadataEnabled }).(pulumi.BoolPtrOutput)
}

// Amazon Resource Name (ARN) of the image recipe.
func (o ImagePipelineOutput) ImageRecipeArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImagePipeline) pulumi.StringPtrOutput { return v.ImageRecipeArn }).(pulumi.StringPtrOutput)
}

// Configuration block with image tests configuration. Detailed below.
func (o ImagePipelineOutput) ImageTestsConfiguration() ImagePipelineImageTestsConfigurationOutput {
	return o.ApplyT(func(v *ImagePipeline) ImagePipelineImageTestsConfigurationOutput { return v.ImageTestsConfiguration }).(ImagePipelineImageTestsConfigurationOutput)
}

// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
func (o ImagePipelineOutput) InfrastructureConfigurationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagePipeline) pulumi.StringOutput { return v.InfrastructureConfigurationArn }).(pulumi.StringOutput)
}

// Name of the image pipeline.
func (o ImagePipelineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagePipeline) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Platform of the image pipeline.
func (o ImagePipelineOutput) Platform() pulumi.StringOutput {
	return o.ApplyT(func(v *ImagePipeline) pulumi.StringOutput { return v.Platform }).(pulumi.StringOutput)
}

// Configuration block with schedule settings. Detailed below.
func (o ImagePipelineOutput) Schedule() ImagePipelineSchedulePtrOutput {
	return o.ApplyT(func(v *ImagePipeline) ImagePipelineSchedulePtrOutput { return v.Schedule }).(ImagePipelineSchedulePtrOutput)
}

// Status of the image pipeline. Valid values are `DISABLED` and `ENABLED`. Defaults to `ENABLED`.
func (o ImagePipelineOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImagePipeline) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// Key-value map of resource tags for the image pipeline. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ImagePipelineOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ImagePipeline) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o ImagePipelineOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ImagePipeline) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type ImagePipelineArrayOutput struct{ *pulumi.OutputState }

func (ImagePipelineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImagePipeline)(nil)).Elem()
}

func (o ImagePipelineArrayOutput) ToImagePipelineArrayOutput() ImagePipelineArrayOutput {
	return o
}

func (o ImagePipelineArrayOutput) ToImagePipelineArrayOutputWithContext(ctx context.Context) ImagePipelineArrayOutput {
	return o
}

func (o ImagePipelineArrayOutput) Index(i pulumi.IntInput) ImagePipelineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ImagePipeline {
		return vs[0].([]*ImagePipeline)[vs[1].(int)]
	}).(ImagePipelineOutput)
}

type ImagePipelineMapOutput struct{ *pulumi.OutputState }

func (ImagePipelineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImagePipeline)(nil)).Elem()
}

func (o ImagePipelineMapOutput) ToImagePipelineMapOutput() ImagePipelineMapOutput {
	return o
}

func (o ImagePipelineMapOutput) ToImagePipelineMapOutputWithContext(ctx context.Context) ImagePipelineMapOutput {
	return o
}

func (o ImagePipelineMapOutput) MapIndex(k pulumi.StringInput) ImagePipelineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ImagePipeline {
		return vs[0].(map[string]*ImagePipeline)[vs[1].(string)]
	}).(ImagePipelineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImagePipelineInput)(nil)).Elem(), &ImagePipeline{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImagePipelineArrayInput)(nil)).Elem(), ImagePipelineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImagePipelineMapInput)(nil)).Elem(), ImagePipelineMap{})
	pulumi.RegisterOutputType(ImagePipelineOutput{})
	pulumi.RegisterOutputType(ImagePipelineArrayOutput{})
	pulumi.RegisterOutputType(ImagePipelineMapOutput{})
}
