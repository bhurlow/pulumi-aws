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

// Provides an [S3 Intelligent-Tiering](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intelligent-tiering.html) configuration resource.
//
// ## Example Usage
// ### Add intelligent tiering configuration for entire S3 bucket
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
//			example, err := s3.NewBucketV2(ctx, "example", nil)
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketIntelligentTieringConfiguration(ctx, "example-entire-bucket", &s3.BucketIntelligentTieringConfigurationArgs{
//				Bucket: example.ID(),
//				Tierings: s3.BucketIntelligentTieringConfigurationTieringArray{
//					&s3.BucketIntelligentTieringConfigurationTieringArgs{
//						AccessTier: pulumi.String("DEEP_ARCHIVE_ACCESS"),
//						Days:       pulumi.Int(180),
//					},
//					&s3.BucketIntelligentTieringConfigurationTieringArgs{
//						AccessTier: pulumi.String("ARCHIVE_ACCESS"),
//						Days:       pulumi.Int(125),
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
// ### Add intelligent tiering configuration with S3 object filter
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
//			example, err := s3.NewBucketV2(ctx, "example", nil)
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketIntelligentTieringConfiguration(ctx, "example-filtered", &s3.BucketIntelligentTieringConfigurationArgs{
//				Bucket: example.ID(),
//				Status: pulumi.String("Disabled"),
//				Filter: &s3.BucketIntelligentTieringConfigurationFilterArgs{
//					Prefix: pulumi.String("documents/"),
//					Tags: pulumi.StringMap{
//						"priority": pulumi.String("high"),
//						"class":    pulumi.String("blue"),
//					},
//				},
//				Tierings: s3.BucketIntelligentTieringConfigurationTieringArray{
//					&s3.BucketIntelligentTieringConfigurationTieringArgs{
//						AccessTier: pulumi.String("ARCHIVE_ACCESS"),
//						Days:       pulumi.Int(125),
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
// Using `pulumi import`, import S3 bucket intelligent tiering configurations using `bucket:name`. For example:
//
// ```sh
//
//	$ pulumi import aws:s3/bucketIntelligentTieringConfiguration:BucketIntelligentTieringConfiguration my-bucket-entire-bucket my-bucket:EntireBucket
//
// ```
type BucketIntelligentTieringConfiguration struct {
	pulumi.CustomResourceState

	// Name of the bucket this intelligent tiering configuration is associated with.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Bucket filter. The configuration only includes objects that meet the filter's criteria (documented below).
	Filter BucketIntelligentTieringConfigurationFilterPtrOutput `pulumi:"filter"`
	// Unique name used to identify the S3 Intelligent-Tiering configuration for the bucket.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the status of the configuration. Valid values: `Enabled`, `Disabled`.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// S3 Intelligent-Tiering storage class tiers of the configuration (documented below).
	Tierings BucketIntelligentTieringConfigurationTieringArrayOutput `pulumi:"tierings"`
}

// NewBucketIntelligentTieringConfiguration registers a new resource with the given unique name, arguments, and options.
func NewBucketIntelligentTieringConfiguration(ctx *pulumi.Context,
	name string, args *BucketIntelligentTieringConfigurationArgs, opts ...pulumi.ResourceOption) (*BucketIntelligentTieringConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Tierings == nil {
		return nil, errors.New("invalid value for required argument 'Tierings'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BucketIntelligentTieringConfiguration
	err := ctx.RegisterResource("aws:s3/bucketIntelligentTieringConfiguration:BucketIntelligentTieringConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketIntelligentTieringConfiguration gets an existing BucketIntelligentTieringConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketIntelligentTieringConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketIntelligentTieringConfigurationState, opts ...pulumi.ResourceOption) (*BucketIntelligentTieringConfiguration, error) {
	var resource BucketIntelligentTieringConfiguration
	err := ctx.ReadResource("aws:s3/bucketIntelligentTieringConfiguration:BucketIntelligentTieringConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketIntelligentTieringConfiguration resources.
type bucketIntelligentTieringConfigurationState struct {
	// Name of the bucket this intelligent tiering configuration is associated with.
	Bucket *string `pulumi:"bucket"`
	// Bucket filter. The configuration only includes objects that meet the filter's criteria (documented below).
	Filter *BucketIntelligentTieringConfigurationFilter `pulumi:"filter"`
	// Unique name used to identify the S3 Intelligent-Tiering configuration for the bucket.
	Name *string `pulumi:"name"`
	// Specifies the status of the configuration. Valid values: `Enabled`, `Disabled`.
	Status *string `pulumi:"status"`
	// S3 Intelligent-Tiering storage class tiers of the configuration (documented below).
	Tierings []BucketIntelligentTieringConfigurationTiering `pulumi:"tierings"`
}

type BucketIntelligentTieringConfigurationState struct {
	// Name of the bucket this intelligent tiering configuration is associated with.
	Bucket pulumi.StringPtrInput
	// Bucket filter. The configuration only includes objects that meet the filter's criteria (documented below).
	Filter BucketIntelligentTieringConfigurationFilterPtrInput
	// Unique name used to identify the S3 Intelligent-Tiering configuration for the bucket.
	Name pulumi.StringPtrInput
	// Specifies the status of the configuration. Valid values: `Enabled`, `Disabled`.
	Status pulumi.StringPtrInput
	// S3 Intelligent-Tiering storage class tiers of the configuration (documented below).
	Tierings BucketIntelligentTieringConfigurationTieringArrayInput
}

func (BucketIntelligentTieringConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketIntelligentTieringConfigurationState)(nil)).Elem()
}

type bucketIntelligentTieringConfigurationArgs struct {
	// Name of the bucket this intelligent tiering configuration is associated with.
	Bucket string `pulumi:"bucket"`
	// Bucket filter. The configuration only includes objects that meet the filter's criteria (documented below).
	Filter *BucketIntelligentTieringConfigurationFilter `pulumi:"filter"`
	// Unique name used to identify the S3 Intelligent-Tiering configuration for the bucket.
	Name *string `pulumi:"name"`
	// Specifies the status of the configuration. Valid values: `Enabled`, `Disabled`.
	Status *string `pulumi:"status"`
	// S3 Intelligent-Tiering storage class tiers of the configuration (documented below).
	Tierings []BucketIntelligentTieringConfigurationTiering `pulumi:"tierings"`
}

// The set of arguments for constructing a BucketIntelligentTieringConfiguration resource.
type BucketIntelligentTieringConfigurationArgs struct {
	// Name of the bucket this intelligent tiering configuration is associated with.
	Bucket pulumi.StringInput
	// Bucket filter. The configuration only includes objects that meet the filter's criteria (documented below).
	Filter BucketIntelligentTieringConfigurationFilterPtrInput
	// Unique name used to identify the S3 Intelligent-Tiering configuration for the bucket.
	Name pulumi.StringPtrInput
	// Specifies the status of the configuration. Valid values: `Enabled`, `Disabled`.
	Status pulumi.StringPtrInput
	// S3 Intelligent-Tiering storage class tiers of the configuration (documented below).
	Tierings BucketIntelligentTieringConfigurationTieringArrayInput
}

func (BucketIntelligentTieringConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketIntelligentTieringConfigurationArgs)(nil)).Elem()
}

type BucketIntelligentTieringConfigurationInput interface {
	pulumi.Input

	ToBucketIntelligentTieringConfigurationOutput() BucketIntelligentTieringConfigurationOutput
	ToBucketIntelligentTieringConfigurationOutputWithContext(ctx context.Context) BucketIntelligentTieringConfigurationOutput
}

func (*BucketIntelligentTieringConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketIntelligentTieringConfiguration)(nil)).Elem()
}

func (i *BucketIntelligentTieringConfiguration) ToBucketIntelligentTieringConfigurationOutput() BucketIntelligentTieringConfigurationOutput {
	return i.ToBucketIntelligentTieringConfigurationOutputWithContext(context.Background())
}

func (i *BucketIntelligentTieringConfiguration) ToBucketIntelligentTieringConfigurationOutputWithContext(ctx context.Context) BucketIntelligentTieringConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketIntelligentTieringConfigurationOutput)
}

// BucketIntelligentTieringConfigurationArrayInput is an input type that accepts BucketIntelligentTieringConfigurationArray and BucketIntelligentTieringConfigurationArrayOutput values.
// You can construct a concrete instance of `BucketIntelligentTieringConfigurationArrayInput` via:
//
//	BucketIntelligentTieringConfigurationArray{ BucketIntelligentTieringConfigurationArgs{...} }
type BucketIntelligentTieringConfigurationArrayInput interface {
	pulumi.Input

	ToBucketIntelligentTieringConfigurationArrayOutput() BucketIntelligentTieringConfigurationArrayOutput
	ToBucketIntelligentTieringConfigurationArrayOutputWithContext(context.Context) BucketIntelligentTieringConfigurationArrayOutput
}

type BucketIntelligentTieringConfigurationArray []BucketIntelligentTieringConfigurationInput

func (BucketIntelligentTieringConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketIntelligentTieringConfiguration)(nil)).Elem()
}

func (i BucketIntelligentTieringConfigurationArray) ToBucketIntelligentTieringConfigurationArrayOutput() BucketIntelligentTieringConfigurationArrayOutput {
	return i.ToBucketIntelligentTieringConfigurationArrayOutputWithContext(context.Background())
}

func (i BucketIntelligentTieringConfigurationArray) ToBucketIntelligentTieringConfigurationArrayOutputWithContext(ctx context.Context) BucketIntelligentTieringConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketIntelligentTieringConfigurationArrayOutput)
}

// BucketIntelligentTieringConfigurationMapInput is an input type that accepts BucketIntelligentTieringConfigurationMap and BucketIntelligentTieringConfigurationMapOutput values.
// You can construct a concrete instance of `BucketIntelligentTieringConfigurationMapInput` via:
//
//	BucketIntelligentTieringConfigurationMap{ "key": BucketIntelligentTieringConfigurationArgs{...} }
type BucketIntelligentTieringConfigurationMapInput interface {
	pulumi.Input

	ToBucketIntelligentTieringConfigurationMapOutput() BucketIntelligentTieringConfigurationMapOutput
	ToBucketIntelligentTieringConfigurationMapOutputWithContext(context.Context) BucketIntelligentTieringConfigurationMapOutput
}

type BucketIntelligentTieringConfigurationMap map[string]BucketIntelligentTieringConfigurationInput

func (BucketIntelligentTieringConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketIntelligentTieringConfiguration)(nil)).Elem()
}

func (i BucketIntelligentTieringConfigurationMap) ToBucketIntelligentTieringConfigurationMapOutput() BucketIntelligentTieringConfigurationMapOutput {
	return i.ToBucketIntelligentTieringConfigurationMapOutputWithContext(context.Background())
}

func (i BucketIntelligentTieringConfigurationMap) ToBucketIntelligentTieringConfigurationMapOutputWithContext(ctx context.Context) BucketIntelligentTieringConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketIntelligentTieringConfigurationMapOutput)
}

type BucketIntelligentTieringConfigurationOutput struct{ *pulumi.OutputState }

func (BucketIntelligentTieringConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketIntelligentTieringConfiguration)(nil)).Elem()
}

func (o BucketIntelligentTieringConfigurationOutput) ToBucketIntelligentTieringConfigurationOutput() BucketIntelligentTieringConfigurationOutput {
	return o
}

func (o BucketIntelligentTieringConfigurationOutput) ToBucketIntelligentTieringConfigurationOutputWithContext(ctx context.Context) BucketIntelligentTieringConfigurationOutput {
	return o
}

// Name of the bucket this intelligent tiering configuration is associated with.
func (o BucketIntelligentTieringConfigurationOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketIntelligentTieringConfiguration) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// Bucket filter. The configuration only includes objects that meet the filter's criteria (documented below).
func (o BucketIntelligentTieringConfigurationOutput) Filter() BucketIntelligentTieringConfigurationFilterPtrOutput {
	return o.ApplyT(func(v *BucketIntelligentTieringConfiguration) BucketIntelligentTieringConfigurationFilterPtrOutput {
		return v.Filter
	}).(BucketIntelligentTieringConfigurationFilterPtrOutput)
}

// Unique name used to identify the S3 Intelligent-Tiering configuration for the bucket.
func (o BucketIntelligentTieringConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketIntelligentTieringConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the status of the configuration. Valid values: `Enabled`, `Disabled`.
func (o BucketIntelligentTieringConfigurationOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketIntelligentTieringConfiguration) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// S3 Intelligent-Tiering storage class tiers of the configuration (documented below).
func (o BucketIntelligentTieringConfigurationOutput) Tierings() BucketIntelligentTieringConfigurationTieringArrayOutput {
	return o.ApplyT(func(v *BucketIntelligentTieringConfiguration) BucketIntelligentTieringConfigurationTieringArrayOutput {
		return v.Tierings
	}).(BucketIntelligentTieringConfigurationTieringArrayOutput)
}

type BucketIntelligentTieringConfigurationArrayOutput struct{ *pulumi.OutputState }

func (BucketIntelligentTieringConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketIntelligentTieringConfiguration)(nil)).Elem()
}

func (o BucketIntelligentTieringConfigurationArrayOutput) ToBucketIntelligentTieringConfigurationArrayOutput() BucketIntelligentTieringConfigurationArrayOutput {
	return o
}

func (o BucketIntelligentTieringConfigurationArrayOutput) ToBucketIntelligentTieringConfigurationArrayOutputWithContext(ctx context.Context) BucketIntelligentTieringConfigurationArrayOutput {
	return o
}

func (o BucketIntelligentTieringConfigurationArrayOutput) Index(i pulumi.IntInput) BucketIntelligentTieringConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketIntelligentTieringConfiguration {
		return vs[0].([]*BucketIntelligentTieringConfiguration)[vs[1].(int)]
	}).(BucketIntelligentTieringConfigurationOutput)
}

type BucketIntelligentTieringConfigurationMapOutput struct{ *pulumi.OutputState }

func (BucketIntelligentTieringConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketIntelligentTieringConfiguration)(nil)).Elem()
}

func (o BucketIntelligentTieringConfigurationMapOutput) ToBucketIntelligentTieringConfigurationMapOutput() BucketIntelligentTieringConfigurationMapOutput {
	return o
}

func (o BucketIntelligentTieringConfigurationMapOutput) ToBucketIntelligentTieringConfigurationMapOutputWithContext(ctx context.Context) BucketIntelligentTieringConfigurationMapOutput {
	return o
}

func (o BucketIntelligentTieringConfigurationMapOutput) MapIndex(k pulumi.StringInput) BucketIntelligentTieringConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketIntelligentTieringConfiguration {
		return vs[0].(map[string]*BucketIntelligentTieringConfiguration)[vs[1].(string)]
	}).(BucketIntelligentTieringConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketIntelligentTieringConfigurationInput)(nil)).Elem(), &BucketIntelligentTieringConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketIntelligentTieringConfigurationArrayInput)(nil)).Elem(), BucketIntelligentTieringConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketIntelligentTieringConfigurationMapInput)(nil)).Elem(), BucketIntelligentTieringConfigurationMap{})
	pulumi.RegisterOutputType(BucketIntelligentTieringConfigurationOutput{})
	pulumi.RegisterOutputType(BucketIntelligentTieringConfigurationArrayOutput{})
	pulumi.RegisterOutputType(BucketIntelligentTieringConfigurationMapOutput{})
}
