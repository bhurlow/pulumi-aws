// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an S3 bucket accelerate configuration resource. See the [Requirements for using Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/userguide/transfer-acceleration.html#transfer-acceleration-requirements) for more details.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		mybucket, err := s3.NewBucketV2(ctx, "mybucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucketAccelerateConfigurationV2(ctx, "example", &s3.BucketAccelerateConfigurationV2Args{
// 			Bucket: mybucket.Bucket,
// 			Status: pulumi.String("Enabled"),
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
// S3 bucket accelerate configuration can be imported using the `bucket` e.g.,
//
// ```sh
//  $ pulumi import aws:s3/bucketAccelerateConfigurationV2:BucketAccelerateConfigurationV2 example bucket-name
// ```
//
//  In addition, S3 bucket accelerate configuration can be imported using the `bucket` and `expected_bucket_owner` separated by a comma (`,`) e.g.,
//
// ```sh
//  $ pulumi import aws:s3/bucketAccelerateConfigurationV2:BucketAccelerateConfigurationV2 example bucket-name,123456789012
// ```
type BucketAccelerateConfigurationV2 struct {
	pulumi.CustomResourceState

	// The name of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The account ID of the expected bucket owner.
	ExpectedBucketOwner pulumi.StringPtrOutput `pulumi:"expectedBucketOwner"`
	// The transfer acceleration state of the bucket. Valid values: `Enabled`, `Suspended`.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewBucketAccelerateConfigurationV2 registers a new resource with the given unique name, arguments, and options.
func NewBucketAccelerateConfigurationV2(ctx *pulumi.Context,
	name string, args *BucketAccelerateConfigurationV2Args, opts ...pulumi.ResourceOption) (*BucketAccelerateConfigurationV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	var resource BucketAccelerateConfigurationV2
	err := ctx.RegisterResource("aws:s3/bucketAccelerateConfigurationV2:BucketAccelerateConfigurationV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketAccelerateConfigurationV2 gets an existing BucketAccelerateConfigurationV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketAccelerateConfigurationV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketAccelerateConfigurationV2State, opts ...pulumi.ResourceOption) (*BucketAccelerateConfigurationV2, error) {
	var resource BucketAccelerateConfigurationV2
	err := ctx.ReadResource("aws:s3/bucketAccelerateConfigurationV2:BucketAccelerateConfigurationV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketAccelerateConfigurationV2 resources.
type bucketAccelerateConfigurationV2State struct {
	// The name of the bucket.
	Bucket *string `pulumi:"bucket"`
	// The account ID of the expected bucket owner.
	ExpectedBucketOwner *string `pulumi:"expectedBucketOwner"`
	// The transfer acceleration state of the bucket. Valid values: `Enabled`, `Suspended`.
	Status *string `pulumi:"status"`
}

type BucketAccelerateConfigurationV2State struct {
	// The name of the bucket.
	Bucket pulumi.StringPtrInput
	// The account ID of the expected bucket owner.
	ExpectedBucketOwner pulumi.StringPtrInput
	// The transfer acceleration state of the bucket. Valid values: `Enabled`, `Suspended`.
	Status pulumi.StringPtrInput
}

func (BucketAccelerateConfigurationV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketAccelerateConfigurationV2State)(nil)).Elem()
}

type bucketAccelerateConfigurationV2Args struct {
	// The name of the bucket.
	Bucket string `pulumi:"bucket"`
	// The account ID of the expected bucket owner.
	ExpectedBucketOwner *string `pulumi:"expectedBucketOwner"`
	// The transfer acceleration state of the bucket. Valid values: `Enabled`, `Suspended`.
	Status string `pulumi:"status"`
}

// The set of arguments for constructing a BucketAccelerateConfigurationV2 resource.
type BucketAccelerateConfigurationV2Args struct {
	// The name of the bucket.
	Bucket pulumi.StringInput
	// The account ID of the expected bucket owner.
	ExpectedBucketOwner pulumi.StringPtrInput
	// The transfer acceleration state of the bucket. Valid values: `Enabled`, `Suspended`.
	Status pulumi.StringInput
}

func (BucketAccelerateConfigurationV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketAccelerateConfigurationV2Args)(nil)).Elem()
}

type BucketAccelerateConfigurationV2Input interface {
	pulumi.Input

	ToBucketAccelerateConfigurationV2Output() BucketAccelerateConfigurationV2Output
	ToBucketAccelerateConfigurationV2OutputWithContext(ctx context.Context) BucketAccelerateConfigurationV2Output
}

func (*BucketAccelerateConfigurationV2) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketAccelerateConfigurationV2)(nil)).Elem()
}

func (i *BucketAccelerateConfigurationV2) ToBucketAccelerateConfigurationV2Output() BucketAccelerateConfigurationV2Output {
	return i.ToBucketAccelerateConfigurationV2OutputWithContext(context.Background())
}

func (i *BucketAccelerateConfigurationV2) ToBucketAccelerateConfigurationV2OutputWithContext(ctx context.Context) BucketAccelerateConfigurationV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(BucketAccelerateConfigurationV2Output)
}

// BucketAccelerateConfigurationV2ArrayInput is an input type that accepts BucketAccelerateConfigurationV2Array and BucketAccelerateConfigurationV2ArrayOutput values.
// You can construct a concrete instance of `BucketAccelerateConfigurationV2ArrayInput` via:
//
//          BucketAccelerateConfigurationV2Array{ BucketAccelerateConfigurationV2Args{...} }
type BucketAccelerateConfigurationV2ArrayInput interface {
	pulumi.Input

	ToBucketAccelerateConfigurationV2ArrayOutput() BucketAccelerateConfigurationV2ArrayOutput
	ToBucketAccelerateConfigurationV2ArrayOutputWithContext(context.Context) BucketAccelerateConfigurationV2ArrayOutput
}

type BucketAccelerateConfigurationV2Array []BucketAccelerateConfigurationV2Input

func (BucketAccelerateConfigurationV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketAccelerateConfigurationV2)(nil)).Elem()
}

func (i BucketAccelerateConfigurationV2Array) ToBucketAccelerateConfigurationV2ArrayOutput() BucketAccelerateConfigurationV2ArrayOutput {
	return i.ToBucketAccelerateConfigurationV2ArrayOutputWithContext(context.Background())
}

func (i BucketAccelerateConfigurationV2Array) ToBucketAccelerateConfigurationV2ArrayOutputWithContext(ctx context.Context) BucketAccelerateConfigurationV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketAccelerateConfigurationV2ArrayOutput)
}

// BucketAccelerateConfigurationV2MapInput is an input type that accepts BucketAccelerateConfigurationV2Map and BucketAccelerateConfigurationV2MapOutput values.
// You can construct a concrete instance of `BucketAccelerateConfigurationV2MapInput` via:
//
//          BucketAccelerateConfigurationV2Map{ "key": BucketAccelerateConfigurationV2Args{...} }
type BucketAccelerateConfigurationV2MapInput interface {
	pulumi.Input

	ToBucketAccelerateConfigurationV2MapOutput() BucketAccelerateConfigurationV2MapOutput
	ToBucketAccelerateConfigurationV2MapOutputWithContext(context.Context) BucketAccelerateConfigurationV2MapOutput
}

type BucketAccelerateConfigurationV2Map map[string]BucketAccelerateConfigurationV2Input

func (BucketAccelerateConfigurationV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketAccelerateConfigurationV2)(nil)).Elem()
}

func (i BucketAccelerateConfigurationV2Map) ToBucketAccelerateConfigurationV2MapOutput() BucketAccelerateConfigurationV2MapOutput {
	return i.ToBucketAccelerateConfigurationV2MapOutputWithContext(context.Background())
}

func (i BucketAccelerateConfigurationV2Map) ToBucketAccelerateConfigurationV2MapOutputWithContext(ctx context.Context) BucketAccelerateConfigurationV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketAccelerateConfigurationV2MapOutput)
}

type BucketAccelerateConfigurationV2Output struct{ *pulumi.OutputState }

func (BucketAccelerateConfigurationV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketAccelerateConfigurationV2)(nil)).Elem()
}

func (o BucketAccelerateConfigurationV2Output) ToBucketAccelerateConfigurationV2Output() BucketAccelerateConfigurationV2Output {
	return o
}

func (o BucketAccelerateConfigurationV2Output) ToBucketAccelerateConfigurationV2OutputWithContext(ctx context.Context) BucketAccelerateConfigurationV2Output {
	return o
}

type BucketAccelerateConfigurationV2ArrayOutput struct{ *pulumi.OutputState }

func (BucketAccelerateConfigurationV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketAccelerateConfigurationV2)(nil)).Elem()
}

func (o BucketAccelerateConfigurationV2ArrayOutput) ToBucketAccelerateConfigurationV2ArrayOutput() BucketAccelerateConfigurationV2ArrayOutput {
	return o
}

func (o BucketAccelerateConfigurationV2ArrayOutput) ToBucketAccelerateConfigurationV2ArrayOutputWithContext(ctx context.Context) BucketAccelerateConfigurationV2ArrayOutput {
	return o
}

func (o BucketAccelerateConfigurationV2ArrayOutput) Index(i pulumi.IntInput) BucketAccelerateConfigurationV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketAccelerateConfigurationV2 {
		return vs[0].([]*BucketAccelerateConfigurationV2)[vs[1].(int)]
	}).(BucketAccelerateConfigurationV2Output)
}

type BucketAccelerateConfigurationV2MapOutput struct{ *pulumi.OutputState }

func (BucketAccelerateConfigurationV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketAccelerateConfigurationV2)(nil)).Elem()
}

func (o BucketAccelerateConfigurationV2MapOutput) ToBucketAccelerateConfigurationV2MapOutput() BucketAccelerateConfigurationV2MapOutput {
	return o
}

func (o BucketAccelerateConfigurationV2MapOutput) ToBucketAccelerateConfigurationV2MapOutputWithContext(ctx context.Context) BucketAccelerateConfigurationV2MapOutput {
	return o
}

func (o BucketAccelerateConfigurationV2MapOutput) MapIndex(k pulumi.StringInput) BucketAccelerateConfigurationV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketAccelerateConfigurationV2 {
		return vs[0].(map[string]*BucketAccelerateConfigurationV2)[vs[1].(string)]
	}).(BucketAccelerateConfigurationV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketAccelerateConfigurationV2Input)(nil)).Elem(), &BucketAccelerateConfigurationV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketAccelerateConfigurationV2ArrayInput)(nil)).Elem(), BucketAccelerateConfigurationV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketAccelerateConfigurationV2MapInput)(nil)).Elem(), BucketAccelerateConfigurationV2Map{})
	pulumi.RegisterOutputType(BucketAccelerateConfigurationV2Output{})
	pulumi.RegisterOutputType(BucketAccelerateConfigurationV2ArrayOutput{})
	pulumi.RegisterOutputType(BucketAccelerateConfigurationV2MapOutput{})
}