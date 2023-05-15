// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a SageMaker Space resource.
//
// ## Example Usage
// ### Basic usage
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
//			_, err := sagemaker.NewSpace(ctx, "example", &sagemaker.SpaceArgs{
//				DomainId:  pulumi.Any(aws_sagemaker_domain.Test.Id),
//				SpaceName: pulumi.String("example"),
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
// SageMaker Spaces can be imported using the `id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:sagemaker/space:Space test_space arn:aws:sagemaker:us-west-2:123456789012:space/domain-id/space-name
//
// ```
type Space struct {
	pulumi.CustomResourceState

	// The space's Amazon Resource Name (ARN).
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ID of the associated Domain.
	DomainId pulumi.StringOutput `pulumi:"domainId"`
	// The ID of the space's profile in the Amazon Elastic File System volume.
	HomeEfsFileSystemUid pulumi.StringOutput `pulumi:"homeEfsFileSystemUid"`
	// The name of the space.
	SpaceName pulumi.StringOutput `pulumi:"spaceName"`
	// A collection of space settings. See Space Settings below.
	SpaceSettings SpaceSpaceSettingsPtrOutput `pulumi:"spaceSettings"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewSpace registers a new resource with the given unique name, arguments, and options.
func NewSpace(ctx *pulumi.Context,
	name string, args *SpaceArgs, opts ...pulumi.ResourceOption) (*Space, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainId == nil {
		return nil, errors.New("invalid value for required argument 'DomainId'")
	}
	if args.SpaceName == nil {
		return nil, errors.New("invalid value for required argument 'SpaceName'")
	}
	var resource Space
	err := ctx.RegisterResource("aws:sagemaker/space:Space", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpace gets an existing Space resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpaceState, opts ...pulumi.ResourceOption) (*Space, error) {
	var resource Space
	err := ctx.ReadResource("aws:sagemaker/space:Space", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Space resources.
type spaceState struct {
	// The space's Amazon Resource Name (ARN).
	Arn *string `pulumi:"arn"`
	// The ID of the associated Domain.
	DomainId *string `pulumi:"domainId"`
	// The ID of the space's profile in the Amazon Elastic File System volume.
	HomeEfsFileSystemUid *string `pulumi:"homeEfsFileSystemUid"`
	// The name of the space.
	SpaceName *string `pulumi:"spaceName"`
	// A collection of space settings. See Space Settings below.
	SpaceSettings *SpaceSpaceSettings `pulumi:"spaceSettings"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type SpaceState struct {
	// The space's Amazon Resource Name (ARN).
	Arn pulumi.StringPtrInput
	// The ID of the associated Domain.
	DomainId pulumi.StringPtrInput
	// The ID of the space's profile in the Amazon Elastic File System volume.
	HomeEfsFileSystemUid pulumi.StringPtrInput
	// The name of the space.
	SpaceName pulumi.StringPtrInput
	// A collection of space settings. See Space Settings below.
	SpaceSettings SpaceSpaceSettingsPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (SpaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*spaceState)(nil)).Elem()
}

type spaceArgs struct {
	// The ID of the associated Domain.
	DomainId string `pulumi:"domainId"`
	// The name of the space.
	SpaceName string `pulumi:"spaceName"`
	// A collection of space settings. See Space Settings below.
	SpaceSettings *SpaceSpaceSettings `pulumi:"spaceSettings"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Space resource.
type SpaceArgs struct {
	// The ID of the associated Domain.
	DomainId pulumi.StringInput
	// The name of the space.
	SpaceName pulumi.StringInput
	// A collection of space settings. See Space Settings below.
	SpaceSettings SpaceSpaceSettingsPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (SpaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spaceArgs)(nil)).Elem()
}

type SpaceInput interface {
	pulumi.Input

	ToSpaceOutput() SpaceOutput
	ToSpaceOutputWithContext(ctx context.Context) SpaceOutput
}

func (*Space) ElementType() reflect.Type {
	return reflect.TypeOf((**Space)(nil)).Elem()
}

func (i *Space) ToSpaceOutput() SpaceOutput {
	return i.ToSpaceOutputWithContext(context.Background())
}

func (i *Space) ToSpaceOutputWithContext(ctx context.Context) SpaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpaceOutput)
}

// SpaceArrayInput is an input type that accepts SpaceArray and SpaceArrayOutput values.
// You can construct a concrete instance of `SpaceArrayInput` via:
//
//	SpaceArray{ SpaceArgs{...} }
type SpaceArrayInput interface {
	pulumi.Input

	ToSpaceArrayOutput() SpaceArrayOutput
	ToSpaceArrayOutputWithContext(context.Context) SpaceArrayOutput
}

type SpaceArray []SpaceInput

func (SpaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Space)(nil)).Elem()
}

func (i SpaceArray) ToSpaceArrayOutput() SpaceArrayOutput {
	return i.ToSpaceArrayOutputWithContext(context.Background())
}

func (i SpaceArray) ToSpaceArrayOutputWithContext(ctx context.Context) SpaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpaceArrayOutput)
}

// SpaceMapInput is an input type that accepts SpaceMap and SpaceMapOutput values.
// You can construct a concrete instance of `SpaceMapInput` via:
//
//	SpaceMap{ "key": SpaceArgs{...} }
type SpaceMapInput interface {
	pulumi.Input

	ToSpaceMapOutput() SpaceMapOutput
	ToSpaceMapOutputWithContext(context.Context) SpaceMapOutput
}

type SpaceMap map[string]SpaceInput

func (SpaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Space)(nil)).Elem()
}

func (i SpaceMap) ToSpaceMapOutput() SpaceMapOutput {
	return i.ToSpaceMapOutputWithContext(context.Background())
}

func (i SpaceMap) ToSpaceMapOutputWithContext(ctx context.Context) SpaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpaceMapOutput)
}

type SpaceOutput struct{ *pulumi.OutputState }

func (SpaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Space)(nil)).Elem()
}

func (o SpaceOutput) ToSpaceOutput() SpaceOutput {
	return o
}

func (o SpaceOutput) ToSpaceOutputWithContext(ctx context.Context) SpaceOutput {
	return o
}

// The space's Amazon Resource Name (ARN).
func (o SpaceOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Space) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The ID of the associated Domain.
func (o SpaceOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v *Space) pulumi.StringOutput { return v.DomainId }).(pulumi.StringOutput)
}

// The ID of the space's profile in the Amazon Elastic File System volume.
func (o SpaceOutput) HomeEfsFileSystemUid() pulumi.StringOutput {
	return o.ApplyT(func(v *Space) pulumi.StringOutput { return v.HomeEfsFileSystemUid }).(pulumi.StringOutput)
}

// The name of the space.
func (o SpaceOutput) SpaceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Space) pulumi.StringOutput { return v.SpaceName }).(pulumi.StringOutput)
}

// A collection of space settings. See Space Settings below.
func (o SpaceOutput) SpaceSettings() SpaceSpaceSettingsPtrOutput {
	return o.ApplyT(func(v *Space) SpaceSpaceSettingsPtrOutput { return v.SpaceSettings }).(SpaceSpaceSettingsPtrOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o SpaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Space) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o SpaceOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Space) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type SpaceArrayOutput struct{ *pulumi.OutputState }

func (SpaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Space)(nil)).Elem()
}

func (o SpaceArrayOutput) ToSpaceArrayOutput() SpaceArrayOutput {
	return o
}

func (o SpaceArrayOutput) ToSpaceArrayOutputWithContext(ctx context.Context) SpaceArrayOutput {
	return o
}

func (o SpaceArrayOutput) Index(i pulumi.IntInput) SpaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Space {
		return vs[0].([]*Space)[vs[1].(int)]
	}).(SpaceOutput)
}

type SpaceMapOutput struct{ *pulumi.OutputState }

func (SpaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Space)(nil)).Elem()
}

func (o SpaceMapOutput) ToSpaceMapOutput() SpaceMapOutput {
	return o
}

func (o SpaceMapOutput) ToSpaceMapOutputWithContext(ctx context.Context) SpaceMapOutput {
	return o
}

func (o SpaceMapOutput) MapIndex(k pulumi.StringInput) SpaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Space {
		return vs[0].(map[string]*Space)[vs[1].(string)]
	}).(SpaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpaceInput)(nil)).Elem(), &Space{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpaceArrayInput)(nil)).Elem(), SpaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpaceMapInput)(nil)).Elem(), SpaceMap{})
	pulumi.RegisterOutputType(SpaceOutput{})
	pulumi.RegisterOutputType(SpaceArrayOutput{})
	pulumi.RegisterOutputType(SpaceMapOutput{})
}
