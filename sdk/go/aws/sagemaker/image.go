// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Sagemaker Image resource.
//
// ## Example Usage
// ### Basic usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/sagemaker"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sagemaker.NewImage(ctx, "example", &sagemaker.ImageArgs{
// 			ImageName: pulumi.String("example"),
// 			RoleArn:   pulumi.Any(aws_iam_role.Test.Arn),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Sagemaker Code Images can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:sagemaker/image:Image test_image my-code-repo
// ```
type Image struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) assigned by AWS to this Image.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the image.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the image. When the image is added to a domain (must be unique to the domain).
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The name of the image. Must be unique to your account.
	ImageName pulumi.StringOutput `pulumi:"imageName"`
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImageName == nil {
		return nil, errors.New("invalid value for required argument 'ImageName'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource Image
	err := ctx.RegisterResource("aws:sagemaker/image:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("aws:sagemaker/image:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this Image.
	Arn *string `pulumi:"arn"`
	// The description of the image.
	Description *string `pulumi:"description"`
	// The display name of the image. When the image is added to a domain (must be unique to the domain).
	DisplayName *string `pulumi:"displayName"`
	// The name of the image. Must be unique to your account.
	ImageName *string `pulumi:"imageName"`
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.
	RoleArn *string `pulumi:"roleArn"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ImageState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this Image.
	Arn pulumi.StringPtrInput
	// The description of the image.
	Description pulumi.StringPtrInput
	// The display name of the image. When the image is added to a domain (must be unique to the domain).
	DisplayName pulumi.StringPtrInput
	// The name of the image. Must be unique to your account.
	ImageName pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.
	RoleArn pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	// The description of the image.
	Description *string `pulumi:"description"`
	// The display name of the image. When the image is added to a domain (must be unique to the domain).
	DisplayName *string `pulumi:"displayName"`
	// The name of the image. Must be unique to your account.
	ImageName string `pulumi:"imageName"`
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.
	RoleArn string `pulumi:"roleArn"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// The description of the image.
	Description pulumi.StringPtrInput
	// The display name of the image. When the image is added to a domain (must be unique to the domain).
	DisplayName pulumi.StringPtrInput
	// The name of the image. Must be unique to your account.
	ImageName pulumi.StringInput
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.
	RoleArn pulumi.StringInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

type ImageInput interface {
	pulumi.Input

	ToImageOutput() ImageOutput
	ToImageOutputWithContext(ctx context.Context) ImageOutput
}

func (*Image) ElementType() reflect.Type {
	return reflect.TypeOf((*Image)(nil))
}

func (i *Image) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i *Image) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

type ImageOutput struct {
	*pulumi.OutputState
}

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Image)(nil))
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ImageOutput{})
}
