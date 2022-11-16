// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an S3 bucket Object Lock configuration resource. For more information about Object Locking, go to [Using S3 Object Lock](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock.html) in the Amazon S3 User Guide.
//
// > **NOTE:** This resource **does not enable** Object Lock for **new** buckets. It configures a default retention period for objects placed in the specified bucket.
// Thus, to **enable** Object Lock for a **new** bucket, see the Using object lock configuration section in  the `s3.BucketV2` resource or the Object Lock configuration for a new bucket example below.
// If you want to **enable** Object Lock for an **existing** bucket, contact AWS Support and see the Object Lock configuration for an existing bucket example below.
//
// ## Example Usage
// ### Object Lock configuration for a new bucket
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", &s3.BucketV2Args{
//				ObjectLockEnabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketObjectLockConfigurationV2(ctx, "exampleBucketObjectLockConfigurationV2", &s3.BucketObjectLockConfigurationV2Args{
//				Bucket: exampleBucketV2.Bucket,
//				Rule: &s3.BucketObjectLockConfigurationV2RuleArgs{
//					DefaultRetention: &s3.BucketObjectLockConfigurationV2RuleDefaultRetentionArgs{
//						Mode: pulumi.String("COMPLIANCE"),
//						Days: pulumi.Int(5),
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
// ### Object Lock configuration for an existing bucket
//
// This is a multistep process that requires AWS Support intervention.
//
//  1. Enable versioning on your S3 bucket, if you have not already done so.
//     Doing so will generate an "Object Lock token" in the back-end.
//
// <!-- markdownlint-disable MD029 -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
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
//			_, err = s3.NewBucketVersioningV2(ctx, "exampleBucketVersioningV2", &s3.BucketVersioningV2Args{
//				Bucket: exampleBucketV2.Bucket,
//				VersioningConfiguration: &s3.BucketVersioningV2VersioningConfigurationArgs{
//					Status: pulumi.String("Enabled"),
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
// <!-- markdownlint-disable MD029 -->
//
//  2. Contact AWS Support to provide you with the "Object Lock token" for the specified bucket and use the token (or token ID) within your new `s3.BucketObjectLockConfigurationV2` resource.
//     Notice the `objectLockEnabled` argument does not need to be specified as it defaults to `Enabled`.
//
// <!-- markdownlint-disable MD029 -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := s3.NewBucketObjectLockConfigurationV2(ctx, "example", &s3.BucketObjectLockConfigurationV2Args{
//				Bucket: pulumi.Any(aws_s3_bucket.Example.Bucket),
//				Rule: &s3.BucketObjectLockConfigurationV2RuleArgs{
//					DefaultRetention: &s3.BucketObjectLockConfigurationV2RuleDefaultRetentionArgs{
//						Mode: pulumi.String("COMPLIANCE"),
//						Days: pulumi.Int(5),
//					},
//				},
//				Token: pulumi.String("NG2MKsfoLqV3A+aquXneSG4LOu/ekrlXkRXwIPFVfERT7XOPos+/k444d7RIH0E3W3p5QU6ml2exS2F/eYCFmMWHJ3hFZGk6al1sIJkmNhUMYmsv0jYVQyTTZNLM+DnfooA6SATt39mM1VW1yJh4E+XljMlWzaBwHKbss3/EjlGDjOmVhaSs4Z6427mMCaFD0RLwsYY7zX49gEc31YfOMJGxbXCXSeyNwAhhM/A8UH7gQf38RmjHjjAFbbbLtl8arsxTPW8F1IYohqwmKIr9DnotLLj8Tg44U2SPwujVaqmlKKP9s41rfgb4UbIm7khSafDBng0LGfxC4pMlT9Ny2w=="),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!-- markdownlint-disable MD029 -->
//
// ## Import
//
// S3 bucket Object Lock configuration can be imported in one of two ways. If the owner (account ID) of the source bucket is the same account used to configure the AWS Provider, the S3 bucket Object Lock configuration resource should be imported using the `bucket` e.g.,
//
// ```sh
//
//	$ pulumi import aws:s3/bucketObjectLockConfigurationV2:BucketObjectLockConfigurationV2 example bucket-name
//
// ```
//
//	If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, the S3 bucket Object Lock configuration resource should be imported using the `bucket` and `expected_bucket_owner` separated by a comma (`,`) e.g.,
//
// ```sh
//
//	$ pulumi import aws:s3/bucketObjectLockConfigurationV2:BucketObjectLockConfigurationV2 example bucket-name,123456789012
//
// ```
type BucketObjectLockConfigurationV2 struct {
	pulumi.CustomResourceState

	// The name of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The account ID of the expected bucket owner.
	ExpectedBucketOwner pulumi.StringPtrOutput `pulumi:"expectedBucketOwner"`
	// Indicates whether this bucket has an Object Lock configuration enabled. Defaults to `Enabled`. Valid values: `Enabled`.
	ObjectLockEnabled pulumi.StringPtrOutput `pulumi:"objectLockEnabled"`
	// Configuration block for specifying the Object Lock rule for the specified object detailed below.
	Rule BucketObjectLockConfigurationV2RulePtrOutput `pulumi:"rule"`
	// A token to allow Object Lock to be enabled for an existing bucket. You must contact AWS support for the bucket's "Object Lock token".
	// The token is generated in the back-end when [versioning](https://docs.aws.amazon.com/AmazonS3/latest/userguide/manage-versioning-examples.html) is enabled on a bucket. For more details on versioning, see the `s3.BucketVersioningV2` resource.
	Token pulumi.StringPtrOutput `pulumi:"token"`
}

// NewBucketObjectLockConfigurationV2 registers a new resource with the given unique name, arguments, and options.
func NewBucketObjectLockConfigurationV2(ctx *pulumi.Context,
	name string, args *BucketObjectLockConfigurationV2Args, opts ...pulumi.ResourceOption) (*BucketObjectLockConfigurationV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	var resource BucketObjectLockConfigurationV2
	err := ctx.RegisterResource("aws:s3/bucketObjectLockConfigurationV2:BucketObjectLockConfigurationV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketObjectLockConfigurationV2 gets an existing BucketObjectLockConfigurationV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketObjectLockConfigurationV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketObjectLockConfigurationV2State, opts ...pulumi.ResourceOption) (*BucketObjectLockConfigurationV2, error) {
	var resource BucketObjectLockConfigurationV2
	err := ctx.ReadResource("aws:s3/bucketObjectLockConfigurationV2:BucketObjectLockConfigurationV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketObjectLockConfigurationV2 resources.
type bucketObjectLockConfigurationV2State struct {
	// The name of the bucket.
	Bucket *string `pulumi:"bucket"`
	// The account ID of the expected bucket owner.
	ExpectedBucketOwner *string `pulumi:"expectedBucketOwner"`
	// Indicates whether this bucket has an Object Lock configuration enabled. Defaults to `Enabled`. Valid values: `Enabled`.
	ObjectLockEnabled *string `pulumi:"objectLockEnabled"`
	// Configuration block for specifying the Object Lock rule for the specified object detailed below.
	Rule *BucketObjectLockConfigurationV2Rule `pulumi:"rule"`
	// A token to allow Object Lock to be enabled for an existing bucket. You must contact AWS support for the bucket's "Object Lock token".
	// The token is generated in the back-end when [versioning](https://docs.aws.amazon.com/AmazonS3/latest/userguide/manage-versioning-examples.html) is enabled on a bucket. For more details on versioning, see the `s3.BucketVersioningV2` resource.
	Token *string `pulumi:"token"`
}

type BucketObjectLockConfigurationV2State struct {
	// The name of the bucket.
	Bucket pulumi.StringPtrInput
	// The account ID of the expected bucket owner.
	ExpectedBucketOwner pulumi.StringPtrInput
	// Indicates whether this bucket has an Object Lock configuration enabled. Defaults to `Enabled`. Valid values: `Enabled`.
	ObjectLockEnabled pulumi.StringPtrInput
	// Configuration block for specifying the Object Lock rule for the specified object detailed below.
	Rule BucketObjectLockConfigurationV2RulePtrInput
	// A token to allow Object Lock to be enabled for an existing bucket. You must contact AWS support for the bucket's "Object Lock token".
	// The token is generated in the back-end when [versioning](https://docs.aws.amazon.com/AmazonS3/latest/userguide/manage-versioning-examples.html) is enabled on a bucket. For more details on versioning, see the `s3.BucketVersioningV2` resource.
	Token pulumi.StringPtrInput
}

func (BucketObjectLockConfigurationV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketObjectLockConfigurationV2State)(nil)).Elem()
}

type bucketObjectLockConfigurationV2Args struct {
	// The name of the bucket.
	Bucket string `pulumi:"bucket"`
	// The account ID of the expected bucket owner.
	ExpectedBucketOwner *string `pulumi:"expectedBucketOwner"`
	// Indicates whether this bucket has an Object Lock configuration enabled. Defaults to `Enabled`. Valid values: `Enabled`.
	ObjectLockEnabled *string `pulumi:"objectLockEnabled"`
	// Configuration block for specifying the Object Lock rule for the specified object detailed below.
	Rule *BucketObjectLockConfigurationV2Rule `pulumi:"rule"`
	// A token to allow Object Lock to be enabled for an existing bucket. You must contact AWS support for the bucket's "Object Lock token".
	// The token is generated in the back-end when [versioning](https://docs.aws.amazon.com/AmazonS3/latest/userguide/manage-versioning-examples.html) is enabled on a bucket. For more details on versioning, see the `s3.BucketVersioningV2` resource.
	Token *string `pulumi:"token"`
}

// The set of arguments for constructing a BucketObjectLockConfigurationV2 resource.
type BucketObjectLockConfigurationV2Args struct {
	// The name of the bucket.
	Bucket pulumi.StringInput
	// The account ID of the expected bucket owner.
	ExpectedBucketOwner pulumi.StringPtrInput
	// Indicates whether this bucket has an Object Lock configuration enabled. Defaults to `Enabled`. Valid values: `Enabled`.
	ObjectLockEnabled pulumi.StringPtrInput
	// Configuration block for specifying the Object Lock rule for the specified object detailed below.
	Rule BucketObjectLockConfigurationV2RulePtrInput
	// A token to allow Object Lock to be enabled for an existing bucket. You must contact AWS support for the bucket's "Object Lock token".
	// The token is generated in the back-end when [versioning](https://docs.aws.amazon.com/AmazonS3/latest/userguide/manage-versioning-examples.html) is enabled on a bucket. For more details on versioning, see the `s3.BucketVersioningV2` resource.
	Token pulumi.StringPtrInput
}

func (BucketObjectLockConfigurationV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketObjectLockConfigurationV2Args)(nil)).Elem()
}

type BucketObjectLockConfigurationV2Input interface {
	pulumi.Input

	ToBucketObjectLockConfigurationV2Output() BucketObjectLockConfigurationV2Output
	ToBucketObjectLockConfigurationV2OutputWithContext(ctx context.Context) BucketObjectLockConfigurationV2Output
}

func (*BucketObjectLockConfigurationV2) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketObjectLockConfigurationV2)(nil)).Elem()
}

func (i *BucketObjectLockConfigurationV2) ToBucketObjectLockConfigurationV2Output() BucketObjectLockConfigurationV2Output {
	return i.ToBucketObjectLockConfigurationV2OutputWithContext(context.Background())
}

func (i *BucketObjectLockConfigurationV2) ToBucketObjectLockConfigurationV2OutputWithContext(ctx context.Context) BucketObjectLockConfigurationV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(BucketObjectLockConfigurationV2Output)
}

// BucketObjectLockConfigurationV2ArrayInput is an input type that accepts BucketObjectLockConfigurationV2Array and BucketObjectLockConfigurationV2ArrayOutput values.
// You can construct a concrete instance of `BucketObjectLockConfigurationV2ArrayInput` via:
//
//	BucketObjectLockConfigurationV2Array{ BucketObjectLockConfigurationV2Args{...} }
type BucketObjectLockConfigurationV2ArrayInput interface {
	pulumi.Input

	ToBucketObjectLockConfigurationV2ArrayOutput() BucketObjectLockConfigurationV2ArrayOutput
	ToBucketObjectLockConfigurationV2ArrayOutputWithContext(context.Context) BucketObjectLockConfigurationV2ArrayOutput
}

type BucketObjectLockConfigurationV2Array []BucketObjectLockConfigurationV2Input

func (BucketObjectLockConfigurationV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketObjectLockConfigurationV2)(nil)).Elem()
}

func (i BucketObjectLockConfigurationV2Array) ToBucketObjectLockConfigurationV2ArrayOutput() BucketObjectLockConfigurationV2ArrayOutput {
	return i.ToBucketObjectLockConfigurationV2ArrayOutputWithContext(context.Background())
}

func (i BucketObjectLockConfigurationV2Array) ToBucketObjectLockConfigurationV2ArrayOutputWithContext(ctx context.Context) BucketObjectLockConfigurationV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketObjectLockConfigurationV2ArrayOutput)
}

// BucketObjectLockConfigurationV2MapInput is an input type that accepts BucketObjectLockConfigurationV2Map and BucketObjectLockConfigurationV2MapOutput values.
// You can construct a concrete instance of `BucketObjectLockConfigurationV2MapInput` via:
//
//	BucketObjectLockConfigurationV2Map{ "key": BucketObjectLockConfigurationV2Args{...} }
type BucketObjectLockConfigurationV2MapInput interface {
	pulumi.Input

	ToBucketObjectLockConfigurationV2MapOutput() BucketObjectLockConfigurationV2MapOutput
	ToBucketObjectLockConfigurationV2MapOutputWithContext(context.Context) BucketObjectLockConfigurationV2MapOutput
}

type BucketObjectLockConfigurationV2Map map[string]BucketObjectLockConfigurationV2Input

func (BucketObjectLockConfigurationV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketObjectLockConfigurationV2)(nil)).Elem()
}

func (i BucketObjectLockConfigurationV2Map) ToBucketObjectLockConfigurationV2MapOutput() BucketObjectLockConfigurationV2MapOutput {
	return i.ToBucketObjectLockConfigurationV2MapOutputWithContext(context.Background())
}

func (i BucketObjectLockConfigurationV2Map) ToBucketObjectLockConfigurationV2MapOutputWithContext(ctx context.Context) BucketObjectLockConfigurationV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketObjectLockConfigurationV2MapOutput)
}

type BucketObjectLockConfigurationV2Output struct{ *pulumi.OutputState }

func (BucketObjectLockConfigurationV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketObjectLockConfigurationV2)(nil)).Elem()
}

func (o BucketObjectLockConfigurationV2Output) ToBucketObjectLockConfigurationV2Output() BucketObjectLockConfigurationV2Output {
	return o
}

func (o BucketObjectLockConfigurationV2Output) ToBucketObjectLockConfigurationV2OutputWithContext(ctx context.Context) BucketObjectLockConfigurationV2Output {
	return o
}

// The name of the bucket.
func (o BucketObjectLockConfigurationV2Output) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObjectLockConfigurationV2) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The account ID of the expected bucket owner.
func (o BucketObjectLockConfigurationV2Output) ExpectedBucketOwner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObjectLockConfigurationV2) pulumi.StringPtrOutput { return v.ExpectedBucketOwner }).(pulumi.StringPtrOutput)
}

// Indicates whether this bucket has an Object Lock configuration enabled. Defaults to `Enabled`. Valid values: `Enabled`.
func (o BucketObjectLockConfigurationV2Output) ObjectLockEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObjectLockConfigurationV2) pulumi.StringPtrOutput { return v.ObjectLockEnabled }).(pulumi.StringPtrOutput)
}

// Configuration block for specifying the Object Lock rule for the specified object detailed below.
func (o BucketObjectLockConfigurationV2Output) Rule() BucketObjectLockConfigurationV2RulePtrOutput {
	return o.ApplyT(func(v *BucketObjectLockConfigurationV2) BucketObjectLockConfigurationV2RulePtrOutput { return v.Rule }).(BucketObjectLockConfigurationV2RulePtrOutput)
}

// A token to allow Object Lock to be enabled for an existing bucket. You must contact AWS support for the bucket's "Object Lock token".
// The token is generated in the back-end when [versioning](https://docs.aws.amazon.com/AmazonS3/latest/userguide/manage-versioning-examples.html) is enabled on a bucket. For more details on versioning, see the `s3.BucketVersioningV2` resource.
func (o BucketObjectLockConfigurationV2Output) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObjectLockConfigurationV2) pulumi.StringPtrOutput { return v.Token }).(pulumi.StringPtrOutput)
}

type BucketObjectLockConfigurationV2ArrayOutput struct{ *pulumi.OutputState }

func (BucketObjectLockConfigurationV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketObjectLockConfigurationV2)(nil)).Elem()
}

func (o BucketObjectLockConfigurationV2ArrayOutput) ToBucketObjectLockConfigurationV2ArrayOutput() BucketObjectLockConfigurationV2ArrayOutput {
	return o
}

func (o BucketObjectLockConfigurationV2ArrayOutput) ToBucketObjectLockConfigurationV2ArrayOutputWithContext(ctx context.Context) BucketObjectLockConfigurationV2ArrayOutput {
	return o
}

func (o BucketObjectLockConfigurationV2ArrayOutput) Index(i pulumi.IntInput) BucketObjectLockConfigurationV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketObjectLockConfigurationV2 {
		return vs[0].([]*BucketObjectLockConfigurationV2)[vs[1].(int)]
	}).(BucketObjectLockConfigurationV2Output)
}

type BucketObjectLockConfigurationV2MapOutput struct{ *pulumi.OutputState }

func (BucketObjectLockConfigurationV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketObjectLockConfigurationV2)(nil)).Elem()
}

func (o BucketObjectLockConfigurationV2MapOutput) ToBucketObjectLockConfigurationV2MapOutput() BucketObjectLockConfigurationV2MapOutput {
	return o
}

func (o BucketObjectLockConfigurationV2MapOutput) ToBucketObjectLockConfigurationV2MapOutputWithContext(ctx context.Context) BucketObjectLockConfigurationV2MapOutput {
	return o
}

func (o BucketObjectLockConfigurationV2MapOutput) MapIndex(k pulumi.StringInput) BucketObjectLockConfigurationV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketObjectLockConfigurationV2 {
		return vs[0].(map[string]*BucketObjectLockConfigurationV2)[vs[1].(string)]
	}).(BucketObjectLockConfigurationV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketObjectLockConfigurationV2Input)(nil)).Elem(), &BucketObjectLockConfigurationV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketObjectLockConfigurationV2ArrayInput)(nil)).Elem(), BucketObjectLockConfigurationV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketObjectLockConfigurationV2MapInput)(nil)).Elem(), BucketObjectLockConfigurationV2Map{})
	pulumi.RegisterOutputType(BucketObjectLockConfigurationV2Output{})
	pulumi.RegisterOutputType(BucketObjectLockConfigurationV2ArrayOutput{})
	pulumi.RegisterOutputType(BucketObjectLockConfigurationV2MapOutput{})
}
