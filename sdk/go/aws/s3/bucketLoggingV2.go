// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an S3 bucket (server access) logging resource. For more information, see [Logging requests using server access logging](https://docs.aws.amazon.com/AmazonS3/latest/userguide/ServerLogs.html)
// in the AWS S3 User Guide.
//
// > **Note:** Amazon S3 supports server access logging, AWS CloudTrail, or a combination of both. Refer to the [Logging options for Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/userguide/logging-with-S3.html)
// to decide which method meets your requirements.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
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
//			_, err = s3.NewBucketAclV2(ctx, "exampleBucketAclV2", &s3.BucketAclV2Args{
//				Bucket: exampleBucketV2.ID(),
//				Acl:    pulumi.String("private"),
//			})
//			if err != nil {
//				return err
//			}
//			logBucket, err := s3.NewBucketV2(ctx, "logBucket", nil)
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketAclV2(ctx, "logBucketAcl", &s3.BucketAclV2Args{
//				Bucket: logBucket.ID(),
//				Acl:    pulumi.String("log-delivery-write"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketLoggingV2(ctx, "exampleBucketLoggingV2", &s3.BucketLoggingV2Args{
//				Bucket:       exampleBucketV2.ID(),
//				TargetBucket: logBucket.ID(),
//				TargetPrefix: pulumi.String("log/"),
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
// If the owner (account ID) of the source bucket is the same account used to configure the AWS Provider, import using the `bucket`terraform import {
//
//	to = aws_s3_bucket_logging.example
//
//	id = "bucket-name" } If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, import using the `bucket` and `expected_bucket_owner` separated by a comma (`,`)terraform import {
//
//	to = aws_s3_bucket_logging.example
//
//	id = "bucket-name,123456789012" } **Using `pulumi import` to import** S3 bucket logging using the `bucket` or using the `bucket` and `expected_bucket_owner` separated by a comma (`,`). For exampleIf the owner (account ID) of the source bucket is the same account used to configure the AWS Provider, import using the `bucket`console % pulumi import aws_s3_bucket_logging.example bucket-name If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, import using the `bucket` and `expected_bucket_owner` separated by a comma (`,`)console % pulumi import aws_s3_bucket_logging.example bucket-name,123456789012
type BucketLoggingV2 struct {
	pulumi.CustomResourceState

	// Name of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Account ID of the expected bucket owner.
	ExpectedBucketOwner pulumi.StringPtrOutput `pulumi:"expectedBucketOwner"`
	// Name of the bucket where you want Amazon S3 to store server access logs.
	TargetBucket pulumi.StringOutput `pulumi:"targetBucket"`
	// Set of configuration blocks with information for granting permissions. See below.
	TargetGrants BucketLoggingV2TargetGrantArrayOutput `pulumi:"targetGrants"`
	// Prefix for all log object keys.
	TargetPrefix pulumi.StringOutput `pulumi:"targetPrefix"`
}

// NewBucketLoggingV2 registers a new resource with the given unique name, arguments, and options.
func NewBucketLoggingV2(ctx *pulumi.Context,
	name string, args *BucketLoggingV2Args, opts ...pulumi.ResourceOption) (*BucketLoggingV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.TargetBucket == nil {
		return nil, errors.New("invalid value for required argument 'TargetBucket'")
	}
	if args.TargetPrefix == nil {
		return nil, errors.New("invalid value for required argument 'TargetPrefix'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BucketLoggingV2
	err := ctx.RegisterResource("aws:s3/bucketLoggingV2:BucketLoggingV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketLoggingV2 gets an existing BucketLoggingV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketLoggingV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketLoggingV2State, opts ...pulumi.ResourceOption) (*BucketLoggingV2, error) {
	var resource BucketLoggingV2
	err := ctx.ReadResource("aws:s3/bucketLoggingV2:BucketLoggingV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketLoggingV2 resources.
type bucketLoggingV2State struct {
	// Name of the bucket.
	Bucket *string `pulumi:"bucket"`
	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `pulumi:"expectedBucketOwner"`
	// Name of the bucket where you want Amazon S3 to store server access logs.
	TargetBucket *string `pulumi:"targetBucket"`
	// Set of configuration blocks with information for granting permissions. See below.
	TargetGrants []BucketLoggingV2TargetGrant `pulumi:"targetGrants"`
	// Prefix for all log object keys.
	TargetPrefix *string `pulumi:"targetPrefix"`
}

type BucketLoggingV2State struct {
	// Name of the bucket.
	Bucket pulumi.StringPtrInput
	// Account ID of the expected bucket owner.
	ExpectedBucketOwner pulumi.StringPtrInput
	// Name of the bucket where you want Amazon S3 to store server access logs.
	TargetBucket pulumi.StringPtrInput
	// Set of configuration blocks with information for granting permissions. See below.
	TargetGrants BucketLoggingV2TargetGrantArrayInput
	// Prefix for all log object keys.
	TargetPrefix pulumi.StringPtrInput
}

func (BucketLoggingV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketLoggingV2State)(nil)).Elem()
}

type bucketLoggingV2Args struct {
	// Name of the bucket.
	Bucket string `pulumi:"bucket"`
	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `pulumi:"expectedBucketOwner"`
	// Name of the bucket where you want Amazon S3 to store server access logs.
	TargetBucket string `pulumi:"targetBucket"`
	// Set of configuration blocks with information for granting permissions. See below.
	TargetGrants []BucketLoggingV2TargetGrant `pulumi:"targetGrants"`
	// Prefix for all log object keys.
	TargetPrefix string `pulumi:"targetPrefix"`
}

// The set of arguments for constructing a BucketLoggingV2 resource.
type BucketLoggingV2Args struct {
	// Name of the bucket.
	Bucket pulumi.StringInput
	// Account ID of the expected bucket owner.
	ExpectedBucketOwner pulumi.StringPtrInput
	// Name of the bucket where you want Amazon S3 to store server access logs.
	TargetBucket pulumi.StringInput
	// Set of configuration blocks with information for granting permissions. See below.
	TargetGrants BucketLoggingV2TargetGrantArrayInput
	// Prefix for all log object keys.
	TargetPrefix pulumi.StringInput
}

func (BucketLoggingV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketLoggingV2Args)(nil)).Elem()
}

type BucketLoggingV2Input interface {
	pulumi.Input

	ToBucketLoggingV2Output() BucketLoggingV2Output
	ToBucketLoggingV2OutputWithContext(ctx context.Context) BucketLoggingV2Output
}

func (*BucketLoggingV2) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketLoggingV2)(nil)).Elem()
}

func (i *BucketLoggingV2) ToBucketLoggingV2Output() BucketLoggingV2Output {
	return i.ToBucketLoggingV2OutputWithContext(context.Background())
}

func (i *BucketLoggingV2) ToBucketLoggingV2OutputWithContext(ctx context.Context) BucketLoggingV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLoggingV2Output)
}

// BucketLoggingV2ArrayInput is an input type that accepts BucketLoggingV2Array and BucketLoggingV2ArrayOutput values.
// You can construct a concrete instance of `BucketLoggingV2ArrayInput` via:
//
//	BucketLoggingV2Array{ BucketLoggingV2Args{...} }
type BucketLoggingV2ArrayInput interface {
	pulumi.Input

	ToBucketLoggingV2ArrayOutput() BucketLoggingV2ArrayOutput
	ToBucketLoggingV2ArrayOutputWithContext(context.Context) BucketLoggingV2ArrayOutput
}

type BucketLoggingV2Array []BucketLoggingV2Input

func (BucketLoggingV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketLoggingV2)(nil)).Elem()
}

func (i BucketLoggingV2Array) ToBucketLoggingV2ArrayOutput() BucketLoggingV2ArrayOutput {
	return i.ToBucketLoggingV2ArrayOutputWithContext(context.Background())
}

func (i BucketLoggingV2Array) ToBucketLoggingV2ArrayOutputWithContext(ctx context.Context) BucketLoggingV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLoggingV2ArrayOutput)
}

// BucketLoggingV2MapInput is an input type that accepts BucketLoggingV2Map and BucketLoggingV2MapOutput values.
// You can construct a concrete instance of `BucketLoggingV2MapInput` via:
//
//	BucketLoggingV2Map{ "key": BucketLoggingV2Args{...} }
type BucketLoggingV2MapInput interface {
	pulumi.Input

	ToBucketLoggingV2MapOutput() BucketLoggingV2MapOutput
	ToBucketLoggingV2MapOutputWithContext(context.Context) BucketLoggingV2MapOutput
}

type BucketLoggingV2Map map[string]BucketLoggingV2Input

func (BucketLoggingV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketLoggingV2)(nil)).Elem()
}

func (i BucketLoggingV2Map) ToBucketLoggingV2MapOutput() BucketLoggingV2MapOutput {
	return i.ToBucketLoggingV2MapOutputWithContext(context.Background())
}

func (i BucketLoggingV2Map) ToBucketLoggingV2MapOutputWithContext(ctx context.Context) BucketLoggingV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLoggingV2MapOutput)
}

type BucketLoggingV2Output struct{ *pulumi.OutputState }

func (BucketLoggingV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketLoggingV2)(nil)).Elem()
}

func (o BucketLoggingV2Output) ToBucketLoggingV2Output() BucketLoggingV2Output {
	return o
}

func (o BucketLoggingV2Output) ToBucketLoggingV2OutputWithContext(ctx context.Context) BucketLoggingV2Output {
	return o
}

// Name of the bucket.
func (o BucketLoggingV2Output) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketLoggingV2) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// Account ID of the expected bucket owner.
func (o BucketLoggingV2Output) ExpectedBucketOwner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketLoggingV2) pulumi.StringPtrOutput { return v.ExpectedBucketOwner }).(pulumi.StringPtrOutput)
}

// Name of the bucket where you want Amazon S3 to store server access logs.
func (o BucketLoggingV2Output) TargetBucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketLoggingV2) pulumi.StringOutput { return v.TargetBucket }).(pulumi.StringOutput)
}

// Set of configuration blocks with information for granting permissions. See below.
func (o BucketLoggingV2Output) TargetGrants() BucketLoggingV2TargetGrantArrayOutput {
	return o.ApplyT(func(v *BucketLoggingV2) BucketLoggingV2TargetGrantArrayOutput { return v.TargetGrants }).(BucketLoggingV2TargetGrantArrayOutput)
}

// Prefix for all log object keys.
func (o BucketLoggingV2Output) TargetPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketLoggingV2) pulumi.StringOutput { return v.TargetPrefix }).(pulumi.StringOutput)
}

type BucketLoggingV2ArrayOutput struct{ *pulumi.OutputState }

func (BucketLoggingV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketLoggingV2)(nil)).Elem()
}

func (o BucketLoggingV2ArrayOutput) ToBucketLoggingV2ArrayOutput() BucketLoggingV2ArrayOutput {
	return o
}

func (o BucketLoggingV2ArrayOutput) ToBucketLoggingV2ArrayOutputWithContext(ctx context.Context) BucketLoggingV2ArrayOutput {
	return o
}

func (o BucketLoggingV2ArrayOutput) Index(i pulumi.IntInput) BucketLoggingV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketLoggingV2 {
		return vs[0].([]*BucketLoggingV2)[vs[1].(int)]
	}).(BucketLoggingV2Output)
}

type BucketLoggingV2MapOutput struct{ *pulumi.OutputState }

func (BucketLoggingV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketLoggingV2)(nil)).Elem()
}

func (o BucketLoggingV2MapOutput) ToBucketLoggingV2MapOutput() BucketLoggingV2MapOutput {
	return o
}

func (o BucketLoggingV2MapOutput) ToBucketLoggingV2MapOutputWithContext(ctx context.Context) BucketLoggingV2MapOutput {
	return o
}

func (o BucketLoggingV2MapOutput) MapIndex(k pulumi.StringInput) BucketLoggingV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketLoggingV2 {
		return vs[0].(map[string]*BucketLoggingV2)[vs[1].(string)]
	}).(BucketLoggingV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketLoggingV2Input)(nil)).Elem(), &BucketLoggingV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketLoggingV2ArrayInput)(nil)).Elem(), BucketLoggingV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketLoggingV2MapInput)(nil)).Elem(), BucketLoggingV2Map{})
	pulumi.RegisterOutputType(BucketLoggingV2Output{})
	pulumi.RegisterOutputType(BucketLoggingV2ArrayOutput{})
	pulumi.RegisterOutputType(BucketLoggingV2MapOutput{})
}
