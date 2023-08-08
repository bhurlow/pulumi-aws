// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3control

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage an S3 Object Lambda Access Point.
// An Object Lambda access point is associated with exactly one standard access point and thus one Amazon S3 bucket.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3control"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
//			if err != nil {
//				return err
//			}
//			exampleAccessPoint, err := s3.NewAccessPoint(ctx, "exampleAccessPoint", &s3.AccessPointArgs{
//				Bucket: exampleBucketV2.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = s3control.NewObjectLambdaAccessPoint(ctx, "exampleObjectLambdaAccessPoint", &s3control.ObjectLambdaAccessPointArgs{
//				Configuration: &s3control.ObjectLambdaAccessPointConfigurationArgs{
//					SupportingAccessPoint: exampleAccessPoint.Arn,
//					TransformationConfigurations: s3control.ObjectLambdaAccessPointConfigurationTransformationConfigurationArray{
//						&s3control.ObjectLambdaAccessPointConfigurationTransformationConfigurationArgs{
//							Actions: pulumi.StringArray{
//								pulumi.String("GetObject"),
//							},
//							ContentTransformation: &s3control.ObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformationArgs{
//								AwsLambda: &s3control.ObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformationAwsLambdaArgs{
//									FunctionArn: pulumi.Any(aws_lambda_function.Example.Arn),
//								},
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
// terraform import {
//
//	to = aws_s3control_object_lambda_access_point.example
//
//	id = "123456789012:example" } Using `pulumi import`, import Object Lambda Access Points using the `account_id` and `name`, separated by a colon (`:`). For exampleconsole % pulumi import aws_s3control_object_lambda_access_point.example 123456789012:example
type ObjectLambdaAccessPoint struct {
	pulumi.CustomResourceState

	// The AWS account ID for the owner of the bucket for which you want to create an Object Lambda Access Point. Defaults to automatically determined account ID of the AWS provider.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// Amazon Resource Name (ARN) of the Object Lambda Access Point.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A configuration block containing details about the Object Lambda Access Point. See Configuration below for more details.
	Configuration ObjectLambdaAccessPointConfigurationOutput `pulumi:"configuration"`
	// The name for this Object Lambda Access Point.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewObjectLambdaAccessPoint registers a new resource with the given unique name, arguments, and options.
func NewObjectLambdaAccessPoint(ctx *pulumi.Context,
	name string, args *ObjectLambdaAccessPointArgs, opts ...pulumi.ResourceOption) (*ObjectLambdaAccessPoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Configuration == nil {
		return nil, errors.New("invalid value for required argument 'Configuration'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ObjectLambdaAccessPoint
	err := ctx.RegisterResource("aws:s3control/objectLambdaAccessPoint:ObjectLambdaAccessPoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectLambdaAccessPoint gets an existing ObjectLambdaAccessPoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectLambdaAccessPoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectLambdaAccessPointState, opts ...pulumi.ResourceOption) (*ObjectLambdaAccessPoint, error) {
	var resource ObjectLambdaAccessPoint
	err := ctx.ReadResource("aws:s3control/objectLambdaAccessPoint:ObjectLambdaAccessPoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectLambdaAccessPoint resources.
type objectLambdaAccessPointState struct {
	// The AWS account ID for the owner of the bucket for which you want to create an Object Lambda Access Point. Defaults to automatically determined account ID of the AWS provider.
	AccountId *string `pulumi:"accountId"`
	// Amazon Resource Name (ARN) of the Object Lambda Access Point.
	Arn *string `pulumi:"arn"`
	// A configuration block containing details about the Object Lambda Access Point. See Configuration below for more details.
	Configuration *ObjectLambdaAccessPointConfiguration `pulumi:"configuration"`
	// The name for this Object Lambda Access Point.
	Name *string `pulumi:"name"`
}

type ObjectLambdaAccessPointState struct {
	// The AWS account ID for the owner of the bucket for which you want to create an Object Lambda Access Point. Defaults to automatically determined account ID of the AWS provider.
	AccountId pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the Object Lambda Access Point.
	Arn pulumi.StringPtrInput
	// A configuration block containing details about the Object Lambda Access Point. See Configuration below for more details.
	Configuration ObjectLambdaAccessPointConfigurationPtrInput
	// The name for this Object Lambda Access Point.
	Name pulumi.StringPtrInput
}

func (ObjectLambdaAccessPointState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectLambdaAccessPointState)(nil)).Elem()
}

type objectLambdaAccessPointArgs struct {
	// The AWS account ID for the owner of the bucket for which you want to create an Object Lambda Access Point. Defaults to automatically determined account ID of the AWS provider.
	AccountId *string `pulumi:"accountId"`
	// A configuration block containing details about the Object Lambda Access Point. See Configuration below for more details.
	Configuration ObjectLambdaAccessPointConfiguration `pulumi:"configuration"`
	// The name for this Object Lambda Access Point.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ObjectLambdaAccessPoint resource.
type ObjectLambdaAccessPointArgs struct {
	// The AWS account ID for the owner of the bucket for which you want to create an Object Lambda Access Point. Defaults to automatically determined account ID of the AWS provider.
	AccountId pulumi.StringPtrInput
	// A configuration block containing details about the Object Lambda Access Point. See Configuration below for more details.
	Configuration ObjectLambdaAccessPointConfigurationInput
	// The name for this Object Lambda Access Point.
	Name pulumi.StringPtrInput
}

func (ObjectLambdaAccessPointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectLambdaAccessPointArgs)(nil)).Elem()
}

type ObjectLambdaAccessPointInput interface {
	pulumi.Input

	ToObjectLambdaAccessPointOutput() ObjectLambdaAccessPointOutput
	ToObjectLambdaAccessPointOutputWithContext(ctx context.Context) ObjectLambdaAccessPointOutput
}

func (*ObjectLambdaAccessPoint) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectLambdaAccessPoint)(nil)).Elem()
}

func (i *ObjectLambdaAccessPoint) ToObjectLambdaAccessPointOutput() ObjectLambdaAccessPointOutput {
	return i.ToObjectLambdaAccessPointOutputWithContext(context.Background())
}

func (i *ObjectLambdaAccessPoint) ToObjectLambdaAccessPointOutputWithContext(ctx context.Context) ObjectLambdaAccessPointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectLambdaAccessPointOutput)
}

// ObjectLambdaAccessPointArrayInput is an input type that accepts ObjectLambdaAccessPointArray and ObjectLambdaAccessPointArrayOutput values.
// You can construct a concrete instance of `ObjectLambdaAccessPointArrayInput` via:
//
//	ObjectLambdaAccessPointArray{ ObjectLambdaAccessPointArgs{...} }
type ObjectLambdaAccessPointArrayInput interface {
	pulumi.Input

	ToObjectLambdaAccessPointArrayOutput() ObjectLambdaAccessPointArrayOutput
	ToObjectLambdaAccessPointArrayOutputWithContext(context.Context) ObjectLambdaAccessPointArrayOutput
}

type ObjectLambdaAccessPointArray []ObjectLambdaAccessPointInput

func (ObjectLambdaAccessPointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectLambdaAccessPoint)(nil)).Elem()
}

func (i ObjectLambdaAccessPointArray) ToObjectLambdaAccessPointArrayOutput() ObjectLambdaAccessPointArrayOutput {
	return i.ToObjectLambdaAccessPointArrayOutputWithContext(context.Background())
}

func (i ObjectLambdaAccessPointArray) ToObjectLambdaAccessPointArrayOutputWithContext(ctx context.Context) ObjectLambdaAccessPointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectLambdaAccessPointArrayOutput)
}

// ObjectLambdaAccessPointMapInput is an input type that accepts ObjectLambdaAccessPointMap and ObjectLambdaAccessPointMapOutput values.
// You can construct a concrete instance of `ObjectLambdaAccessPointMapInput` via:
//
//	ObjectLambdaAccessPointMap{ "key": ObjectLambdaAccessPointArgs{...} }
type ObjectLambdaAccessPointMapInput interface {
	pulumi.Input

	ToObjectLambdaAccessPointMapOutput() ObjectLambdaAccessPointMapOutput
	ToObjectLambdaAccessPointMapOutputWithContext(context.Context) ObjectLambdaAccessPointMapOutput
}

type ObjectLambdaAccessPointMap map[string]ObjectLambdaAccessPointInput

func (ObjectLambdaAccessPointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectLambdaAccessPoint)(nil)).Elem()
}

func (i ObjectLambdaAccessPointMap) ToObjectLambdaAccessPointMapOutput() ObjectLambdaAccessPointMapOutput {
	return i.ToObjectLambdaAccessPointMapOutputWithContext(context.Background())
}

func (i ObjectLambdaAccessPointMap) ToObjectLambdaAccessPointMapOutputWithContext(ctx context.Context) ObjectLambdaAccessPointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectLambdaAccessPointMapOutput)
}

type ObjectLambdaAccessPointOutput struct{ *pulumi.OutputState }

func (ObjectLambdaAccessPointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectLambdaAccessPoint)(nil)).Elem()
}

func (o ObjectLambdaAccessPointOutput) ToObjectLambdaAccessPointOutput() ObjectLambdaAccessPointOutput {
	return o
}

func (o ObjectLambdaAccessPointOutput) ToObjectLambdaAccessPointOutputWithContext(ctx context.Context) ObjectLambdaAccessPointOutput {
	return o
}

// The AWS account ID for the owner of the bucket for which you want to create an Object Lambda Access Point. Defaults to automatically determined account ID of the AWS provider.
func (o ObjectLambdaAccessPointOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectLambdaAccessPoint) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// Amazon Resource Name (ARN) of the Object Lambda Access Point.
func (o ObjectLambdaAccessPointOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectLambdaAccessPoint) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// A configuration block containing details about the Object Lambda Access Point. See Configuration below for more details.
func (o ObjectLambdaAccessPointOutput) Configuration() ObjectLambdaAccessPointConfigurationOutput {
	return o.ApplyT(func(v *ObjectLambdaAccessPoint) ObjectLambdaAccessPointConfigurationOutput { return v.Configuration }).(ObjectLambdaAccessPointConfigurationOutput)
}

// The name for this Object Lambda Access Point.
func (o ObjectLambdaAccessPointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectLambdaAccessPoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type ObjectLambdaAccessPointArrayOutput struct{ *pulumi.OutputState }

func (ObjectLambdaAccessPointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectLambdaAccessPoint)(nil)).Elem()
}

func (o ObjectLambdaAccessPointArrayOutput) ToObjectLambdaAccessPointArrayOutput() ObjectLambdaAccessPointArrayOutput {
	return o
}

func (o ObjectLambdaAccessPointArrayOutput) ToObjectLambdaAccessPointArrayOutputWithContext(ctx context.Context) ObjectLambdaAccessPointArrayOutput {
	return o
}

func (o ObjectLambdaAccessPointArrayOutput) Index(i pulumi.IntInput) ObjectLambdaAccessPointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ObjectLambdaAccessPoint {
		return vs[0].([]*ObjectLambdaAccessPoint)[vs[1].(int)]
	}).(ObjectLambdaAccessPointOutput)
}

type ObjectLambdaAccessPointMapOutput struct{ *pulumi.OutputState }

func (ObjectLambdaAccessPointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectLambdaAccessPoint)(nil)).Elem()
}

func (o ObjectLambdaAccessPointMapOutput) ToObjectLambdaAccessPointMapOutput() ObjectLambdaAccessPointMapOutput {
	return o
}

func (o ObjectLambdaAccessPointMapOutput) ToObjectLambdaAccessPointMapOutputWithContext(ctx context.Context) ObjectLambdaAccessPointMapOutput {
	return o
}

func (o ObjectLambdaAccessPointMapOutput) MapIndex(k pulumi.StringInput) ObjectLambdaAccessPointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ObjectLambdaAccessPoint {
		return vs[0].(map[string]*ObjectLambdaAccessPoint)[vs[1].(string)]
	}).(ObjectLambdaAccessPointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectLambdaAccessPointInput)(nil)).Elem(), &ObjectLambdaAccessPoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectLambdaAccessPointArrayInput)(nil)).Elem(), ObjectLambdaAccessPointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectLambdaAccessPointMapInput)(nil)).Elem(), ObjectLambdaAccessPointMap{})
	pulumi.RegisterOutputType(ObjectLambdaAccessPointOutput{})
	pulumi.RegisterOutputType(ObjectLambdaAccessPointArrayOutput{})
	pulumi.RegisterOutputType(ObjectLambdaAccessPointMapOutput{})
}
