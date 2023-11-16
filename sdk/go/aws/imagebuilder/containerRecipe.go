// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Image Builder Container Recipe.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/imagebuilder"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := imagebuilder.NewContainerRecipe(ctx, "example", &imagebuilder.ContainerRecipeArgs{
//				Version:       pulumi.String("1.0.0"),
//				ContainerType: pulumi.String("DOCKER"),
//				ParentImage:   pulumi.String("arn:aws:imagebuilder:eu-central-1:aws:image/amazon-linux-x86-latest/x.x.x"),
//				TargetRepository: &imagebuilder.ContainerRecipeTargetRepositoryArgs{
//					RepositoryName: pulumi.Any(aws_ecr_repository.Example.Name),
//					Service:        pulumi.String("ECR"),
//				},
//				Components: imagebuilder.ContainerRecipeComponentArray{
//					&imagebuilder.ContainerRecipeComponentArgs{
//						ComponentArn: pulumi.Any(aws_imagebuilder_component.Example.Arn),
//						Parameters: imagebuilder.ContainerRecipeComponentParameterArray{
//							&imagebuilder.ContainerRecipeComponentParameterArgs{
//								Name:  pulumi.String("Parameter1"),
//								Value: pulumi.String("Value1"),
//							},
//							&imagebuilder.ContainerRecipeComponentParameterArgs{
//								Name:  pulumi.String("Parameter2"),
//								Value: pulumi.String("Value2"),
//							},
//						},
//					},
//				},
//				DockerfileTemplateData: pulumi.String("FROM {{{ imagebuilder:parentImage }}}\n{{{ imagebuilder:environments }}}\n{{{ imagebuilder:components }}}\n"),
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
// Using `pulumi import`, import `aws_imagebuilder_container_recipe` resources using the Amazon Resource Name (ARN). For example:
//
// ```sh
//
//	$ pulumi import aws:imagebuilder/containerRecipe:ContainerRecipe example arn:aws:imagebuilder:us-east-1:123456789012:container-recipe/example/1.0.0
//
// ```
type ContainerRecipe struct {
	pulumi.CustomResourceState

	// (Required) Amazon Resource Name (ARN) of the container recipe.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Ordered configuration block(s) with components for the container recipe. Detailed below.
	Components ContainerRecipeComponentArrayOutput `pulumi:"components"`
	// The type of the container to create. Valid values: `DOCKER`.
	ContainerType pulumi.StringOutput `pulumi:"containerType"`
	// Date the container recipe was created.
	DateCreated pulumi.StringOutput `pulumi:"dateCreated"`
	// The description of the container recipe.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Dockerfile template used to build the image as an inline data blob.
	DockerfileTemplateData pulumi.StringOutput `pulumi:"dockerfileTemplateData"`
	// The Amazon S3 URI for the Dockerfile that will be used to build the container image.
	DockerfileTemplateUri pulumi.StringPtrOutput `pulumi:"dockerfileTemplateUri"`
	// Whether to encrypt the volume. Defaults to unset, which is the value inherited from the parent image.
	Encrypted pulumi.BoolOutput `pulumi:"encrypted"`
	// Configuration block used to configure an instance for building and testing container images. Detailed below.
	InstanceConfiguration ContainerRecipeInstanceConfigurationPtrOutput `pulumi:"instanceConfiguration"`
	// The KMS key used to encrypt the container image.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// The name of the container recipe.
	Name pulumi.StringOutput `pulumi:"name"`
	// Owner of the container recipe.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// The base image for the container recipe.
	ParentImage pulumi.StringOutput `pulumi:"parentImage"`
	// Platform of the container recipe.
	Platform pulumi.StringOutput `pulumi:"platform"`
	// Specifies the operating system platform when you use a custom base image.
	PlatformOverride pulumi.StringPtrOutput `pulumi:"platformOverride"`
	// Key-value map of resource tags for the container recipe. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The destination repository for the container image. Detailed below.
	TargetRepository ContainerRecipeTargetRepositoryOutput `pulumi:"targetRepository"`
	// Version of the container recipe.
	//
	// The following attributes are optional:
	Version pulumi.StringOutput `pulumi:"version"`
	// The working directory to be used during build and test workflows.
	WorkingDirectory pulumi.StringPtrOutput `pulumi:"workingDirectory"`
}

// NewContainerRecipe registers a new resource with the given unique name, arguments, and options.
func NewContainerRecipe(ctx *pulumi.Context,
	name string, args *ContainerRecipeArgs, opts ...pulumi.ResourceOption) (*ContainerRecipe, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Components == nil {
		return nil, errors.New("invalid value for required argument 'Components'")
	}
	if args.ContainerType == nil {
		return nil, errors.New("invalid value for required argument 'ContainerType'")
	}
	if args.ParentImage == nil {
		return nil, errors.New("invalid value for required argument 'ParentImage'")
	}
	if args.TargetRepository == nil {
		return nil, errors.New("invalid value for required argument 'TargetRepository'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContainerRecipe
	err := ctx.RegisterResource("aws:imagebuilder/containerRecipe:ContainerRecipe", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerRecipe gets an existing ContainerRecipe resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerRecipe(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerRecipeState, opts ...pulumi.ResourceOption) (*ContainerRecipe, error) {
	var resource ContainerRecipe
	err := ctx.ReadResource("aws:imagebuilder/containerRecipe:ContainerRecipe", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerRecipe resources.
type containerRecipeState struct {
	// (Required) Amazon Resource Name (ARN) of the container recipe.
	Arn *string `pulumi:"arn"`
	// Ordered configuration block(s) with components for the container recipe. Detailed below.
	Components []ContainerRecipeComponent `pulumi:"components"`
	// The type of the container to create. Valid values: `DOCKER`.
	ContainerType *string `pulumi:"containerType"`
	// Date the container recipe was created.
	DateCreated *string `pulumi:"dateCreated"`
	// The description of the container recipe.
	Description *string `pulumi:"description"`
	// The Dockerfile template used to build the image as an inline data blob.
	DockerfileTemplateData *string `pulumi:"dockerfileTemplateData"`
	// The Amazon S3 URI for the Dockerfile that will be used to build the container image.
	DockerfileTemplateUri *string `pulumi:"dockerfileTemplateUri"`
	// Whether to encrypt the volume. Defaults to unset, which is the value inherited from the parent image.
	Encrypted *bool `pulumi:"encrypted"`
	// Configuration block used to configure an instance for building and testing container images. Detailed below.
	InstanceConfiguration *ContainerRecipeInstanceConfiguration `pulumi:"instanceConfiguration"`
	// The KMS key used to encrypt the container image.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The name of the container recipe.
	Name *string `pulumi:"name"`
	// Owner of the container recipe.
	Owner *string `pulumi:"owner"`
	// The base image for the container recipe.
	ParentImage *string `pulumi:"parentImage"`
	// Platform of the container recipe.
	Platform *string `pulumi:"platform"`
	// Specifies the operating system platform when you use a custom base image.
	PlatformOverride *string `pulumi:"platformOverride"`
	// Key-value map of resource tags for the container recipe. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The destination repository for the container image. Detailed below.
	TargetRepository *ContainerRecipeTargetRepository `pulumi:"targetRepository"`
	// Version of the container recipe.
	//
	// The following attributes are optional:
	Version *string `pulumi:"version"`
	// The working directory to be used during build and test workflows.
	WorkingDirectory *string `pulumi:"workingDirectory"`
}

type ContainerRecipeState struct {
	// (Required) Amazon Resource Name (ARN) of the container recipe.
	Arn pulumi.StringPtrInput
	// Ordered configuration block(s) with components for the container recipe. Detailed below.
	Components ContainerRecipeComponentArrayInput
	// The type of the container to create. Valid values: `DOCKER`.
	ContainerType pulumi.StringPtrInput
	// Date the container recipe was created.
	DateCreated pulumi.StringPtrInput
	// The description of the container recipe.
	Description pulumi.StringPtrInput
	// The Dockerfile template used to build the image as an inline data blob.
	DockerfileTemplateData pulumi.StringPtrInput
	// The Amazon S3 URI for the Dockerfile that will be used to build the container image.
	DockerfileTemplateUri pulumi.StringPtrInput
	// Whether to encrypt the volume. Defaults to unset, which is the value inherited from the parent image.
	Encrypted pulumi.BoolPtrInput
	// Configuration block used to configure an instance for building and testing container images. Detailed below.
	InstanceConfiguration ContainerRecipeInstanceConfigurationPtrInput
	// The KMS key used to encrypt the container image.
	KmsKeyId pulumi.StringPtrInput
	// The name of the container recipe.
	Name pulumi.StringPtrInput
	// Owner of the container recipe.
	Owner pulumi.StringPtrInput
	// The base image for the container recipe.
	ParentImage pulumi.StringPtrInput
	// Platform of the container recipe.
	Platform pulumi.StringPtrInput
	// Specifies the operating system platform when you use a custom base image.
	PlatformOverride pulumi.StringPtrInput
	// Key-value map of resource tags for the container recipe. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// The destination repository for the container image. Detailed below.
	TargetRepository ContainerRecipeTargetRepositoryPtrInput
	// Version of the container recipe.
	//
	// The following attributes are optional:
	Version pulumi.StringPtrInput
	// The working directory to be used during build and test workflows.
	WorkingDirectory pulumi.StringPtrInput
}

func (ContainerRecipeState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRecipeState)(nil)).Elem()
}

type containerRecipeArgs struct {
	// Ordered configuration block(s) with components for the container recipe. Detailed below.
	Components []ContainerRecipeComponent `pulumi:"components"`
	// The type of the container to create. Valid values: `DOCKER`.
	ContainerType string `pulumi:"containerType"`
	// The description of the container recipe.
	Description *string `pulumi:"description"`
	// The Dockerfile template used to build the image as an inline data blob.
	DockerfileTemplateData *string `pulumi:"dockerfileTemplateData"`
	// The Amazon S3 URI for the Dockerfile that will be used to build the container image.
	DockerfileTemplateUri *string `pulumi:"dockerfileTemplateUri"`
	// Configuration block used to configure an instance for building and testing container images. Detailed below.
	InstanceConfiguration *ContainerRecipeInstanceConfiguration `pulumi:"instanceConfiguration"`
	// The KMS key used to encrypt the container image.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The name of the container recipe.
	Name *string `pulumi:"name"`
	// The base image for the container recipe.
	ParentImage string `pulumi:"parentImage"`
	// Specifies the operating system platform when you use a custom base image.
	PlatformOverride *string `pulumi:"platformOverride"`
	// Key-value map of resource tags for the container recipe. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The destination repository for the container image. Detailed below.
	TargetRepository ContainerRecipeTargetRepository `pulumi:"targetRepository"`
	// Version of the container recipe.
	//
	// The following attributes are optional:
	Version string `pulumi:"version"`
	// The working directory to be used during build and test workflows.
	WorkingDirectory *string `pulumi:"workingDirectory"`
}

// The set of arguments for constructing a ContainerRecipe resource.
type ContainerRecipeArgs struct {
	// Ordered configuration block(s) with components for the container recipe. Detailed below.
	Components ContainerRecipeComponentArrayInput
	// The type of the container to create. Valid values: `DOCKER`.
	ContainerType pulumi.StringInput
	// The description of the container recipe.
	Description pulumi.StringPtrInput
	// The Dockerfile template used to build the image as an inline data blob.
	DockerfileTemplateData pulumi.StringPtrInput
	// The Amazon S3 URI for the Dockerfile that will be used to build the container image.
	DockerfileTemplateUri pulumi.StringPtrInput
	// Configuration block used to configure an instance for building and testing container images. Detailed below.
	InstanceConfiguration ContainerRecipeInstanceConfigurationPtrInput
	// The KMS key used to encrypt the container image.
	KmsKeyId pulumi.StringPtrInput
	// The name of the container recipe.
	Name pulumi.StringPtrInput
	// The base image for the container recipe.
	ParentImage pulumi.StringInput
	// Specifies the operating system platform when you use a custom base image.
	PlatformOverride pulumi.StringPtrInput
	// Key-value map of resource tags for the container recipe. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The destination repository for the container image. Detailed below.
	TargetRepository ContainerRecipeTargetRepositoryInput
	// Version of the container recipe.
	//
	// The following attributes are optional:
	Version pulumi.StringInput
	// The working directory to be used during build and test workflows.
	WorkingDirectory pulumi.StringPtrInput
}

func (ContainerRecipeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRecipeArgs)(nil)).Elem()
}

type ContainerRecipeInput interface {
	pulumi.Input

	ToContainerRecipeOutput() ContainerRecipeOutput
	ToContainerRecipeOutputWithContext(ctx context.Context) ContainerRecipeOutput
}

func (*ContainerRecipe) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRecipe)(nil)).Elem()
}

func (i *ContainerRecipe) ToContainerRecipeOutput() ContainerRecipeOutput {
	return i.ToContainerRecipeOutputWithContext(context.Background())
}

func (i *ContainerRecipe) ToContainerRecipeOutputWithContext(ctx context.Context) ContainerRecipeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRecipeOutput)
}

// ContainerRecipeArrayInput is an input type that accepts ContainerRecipeArray and ContainerRecipeArrayOutput values.
// You can construct a concrete instance of `ContainerRecipeArrayInput` via:
//
//	ContainerRecipeArray{ ContainerRecipeArgs{...} }
type ContainerRecipeArrayInput interface {
	pulumi.Input

	ToContainerRecipeArrayOutput() ContainerRecipeArrayOutput
	ToContainerRecipeArrayOutputWithContext(context.Context) ContainerRecipeArrayOutput
}

type ContainerRecipeArray []ContainerRecipeInput

func (ContainerRecipeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerRecipe)(nil)).Elem()
}

func (i ContainerRecipeArray) ToContainerRecipeArrayOutput() ContainerRecipeArrayOutput {
	return i.ToContainerRecipeArrayOutputWithContext(context.Background())
}

func (i ContainerRecipeArray) ToContainerRecipeArrayOutputWithContext(ctx context.Context) ContainerRecipeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRecipeArrayOutput)
}

// ContainerRecipeMapInput is an input type that accepts ContainerRecipeMap and ContainerRecipeMapOutput values.
// You can construct a concrete instance of `ContainerRecipeMapInput` via:
//
//	ContainerRecipeMap{ "key": ContainerRecipeArgs{...} }
type ContainerRecipeMapInput interface {
	pulumi.Input

	ToContainerRecipeMapOutput() ContainerRecipeMapOutput
	ToContainerRecipeMapOutputWithContext(context.Context) ContainerRecipeMapOutput
}

type ContainerRecipeMap map[string]ContainerRecipeInput

func (ContainerRecipeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerRecipe)(nil)).Elem()
}

func (i ContainerRecipeMap) ToContainerRecipeMapOutput() ContainerRecipeMapOutput {
	return i.ToContainerRecipeMapOutputWithContext(context.Background())
}

func (i ContainerRecipeMap) ToContainerRecipeMapOutputWithContext(ctx context.Context) ContainerRecipeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRecipeMapOutput)
}

type ContainerRecipeOutput struct{ *pulumi.OutputState }

func (ContainerRecipeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRecipe)(nil)).Elem()
}

func (o ContainerRecipeOutput) ToContainerRecipeOutput() ContainerRecipeOutput {
	return o
}

func (o ContainerRecipeOutput) ToContainerRecipeOutputWithContext(ctx context.Context) ContainerRecipeOutput {
	return o
}

// (Required) Amazon Resource Name (ARN) of the container recipe.
func (o ContainerRecipeOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRecipe) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Ordered configuration block(s) with components for the container recipe. Detailed below.
func (o ContainerRecipeOutput) Components() ContainerRecipeComponentArrayOutput {
	return o.ApplyT(func(v *ContainerRecipe) ContainerRecipeComponentArrayOutput { return v.Components }).(ContainerRecipeComponentArrayOutput)
}

// The type of the container to create. Valid values: `DOCKER`.
func (o ContainerRecipeOutput) ContainerType() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRecipe) pulumi.StringOutput { return v.ContainerType }).(pulumi.StringOutput)
}

// Date the container recipe was created.
func (o ContainerRecipeOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRecipe) pulumi.StringOutput { return v.DateCreated }).(pulumi.StringOutput)
}

// The description of the container recipe.
func (o ContainerRecipeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRecipe) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Dockerfile template used to build the image as an inline data blob.
func (o ContainerRecipeOutput) DockerfileTemplateData() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRecipe) pulumi.StringOutput { return v.DockerfileTemplateData }).(pulumi.StringOutput)
}

// The Amazon S3 URI for the Dockerfile that will be used to build the container image.
func (o ContainerRecipeOutput) DockerfileTemplateUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRecipe) pulumi.StringPtrOutput { return v.DockerfileTemplateUri }).(pulumi.StringPtrOutput)
}

// Whether to encrypt the volume. Defaults to unset, which is the value inherited from the parent image.
func (o ContainerRecipeOutput) Encrypted() pulumi.BoolOutput {
	return o.ApplyT(func(v *ContainerRecipe) pulumi.BoolOutput { return v.Encrypted }).(pulumi.BoolOutput)
}

// Configuration block used to configure an instance for building and testing container images. Detailed below.
func (o ContainerRecipeOutput) InstanceConfiguration() ContainerRecipeInstanceConfigurationPtrOutput {
	return o.ApplyT(func(v *ContainerRecipe) ContainerRecipeInstanceConfigurationPtrOutput { return v.InstanceConfiguration }).(ContainerRecipeInstanceConfigurationPtrOutput)
}

// The KMS key used to encrypt the container image.
func (o ContainerRecipeOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRecipe) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// The name of the container recipe.
func (o ContainerRecipeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRecipe) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Owner of the container recipe.
func (o ContainerRecipeOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRecipe) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// The base image for the container recipe.
func (o ContainerRecipeOutput) ParentImage() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRecipe) pulumi.StringOutput { return v.ParentImage }).(pulumi.StringOutput)
}

// Platform of the container recipe.
func (o ContainerRecipeOutput) Platform() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRecipe) pulumi.StringOutput { return v.Platform }).(pulumi.StringOutput)
}

// Specifies the operating system platform when you use a custom base image.
func (o ContainerRecipeOutput) PlatformOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRecipe) pulumi.StringPtrOutput { return v.PlatformOverride }).(pulumi.StringPtrOutput)
}

// Key-value map of resource tags for the container recipe. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ContainerRecipeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ContainerRecipe) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o ContainerRecipeOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ContainerRecipe) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The destination repository for the container image. Detailed below.
func (o ContainerRecipeOutput) TargetRepository() ContainerRecipeTargetRepositoryOutput {
	return o.ApplyT(func(v *ContainerRecipe) ContainerRecipeTargetRepositoryOutput { return v.TargetRepository }).(ContainerRecipeTargetRepositoryOutput)
}

// Version of the container recipe.
//
// The following attributes are optional:
func (o ContainerRecipeOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRecipe) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

// The working directory to be used during build and test workflows.
func (o ContainerRecipeOutput) WorkingDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRecipe) pulumi.StringPtrOutput { return v.WorkingDirectory }).(pulumi.StringPtrOutput)
}

type ContainerRecipeArrayOutput struct{ *pulumi.OutputState }

func (ContainerRecipeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerRecipe)(nil)).Elem()
}

func (o ContainerRecipeArrayOutput) ToContainerRecipeArrayOutput() ContainerRecipeArrayOutput {
	return o
}

func (o ContainerRecipeArrayOutput) ToContainerRecipeArrayOutputWithContext(ctx context.Context) ContainerRecipeArrayOutput {
	return o
}

func (o ContainerRecipeArrayOutput) Index(i pulumi.IntInput) ContainerRecipeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContainerRecipe {
		return vs[0].([]*ContainerRecipe)[vs[1].(int)]
	}).(ContainerRecipeOutput)
}

type ContainerRecipeMapOutput struct{ *pulumi.OutputState }

func (ContainerRecipeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerRecipe)(nil)).Elem()
}

func (o ContainerRecipeMapOutput) ToContainerRecipeMapOutput() ContainerRecipeMapOutput {
	return o
}

func (o ContainerRecipeMapOutput) ToContainerRecipeMapOutputWithContext(ctx context.Context) ContainerRecipeMapOutput {
	return o
}

func (o ContainerRecipeMapOutput) MapIndex(k pulumi.StringInput) ContainerRecipeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContainerRecipe {
		return vs[0].(map[string]*ContainerRecipe)[vs[1].(string)]
	}).(ContainerRecipeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRecipeInput)(nil)).Elem(), &ContainerRecipe{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRecipeArrayInput)(nil)).Elem(), ContainerRecipeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRecipeMapInput)(nil)).Elem(), ContainerRecipeMap{})
	pulumi.RegisterOutputType(ContainerRecipeOutput{})
	pulumi.RegisterOutputType(ContainerRecipeArrayOutput{})
	pulumi.RegisterOutputType(ContainerRecipeMapOutput{})
}
