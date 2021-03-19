// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3control

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to manage an S3 Control Bucket Lifecycle Configuration.
//
// > **NOTE:** Each S3 Control Bucket can only have one Lifecycle Configuration. Using multiple of this resource against the same S3 Control Bucket will result in perpetual differences each provider run.
//
// > This functionality is for managing [S3 on Outposts](https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html). To manage S3 Bucket Lifecycle Configurations in an AWS Partition, see the `s3.Bucket` resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3control"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := s3control.NewBucketLifecycleConfiguration(ctx, "example", &s3control.BucketLifecycleConfigurationArgs{
// 			Bucket: pulumi.Any(aws_s3control_bucket.Example.Arn),
// 			Rules: s3control.BucketLifecycleConfigurationRuleArray{
// 				&s3control.BucketLifecycleConfigurationRuleArgs{
// 					Expiration: &s3control.BucketLifecycleConfigurationRuleExpirationArgs{
// 						Days: pulumi.Int(365),
// 					},
// 					Filter: &s3control.BucketLifecycleConfigurationRuleFilterArgs{
// 						Prefix: pulumi.String("logs/"),
// 					},
// 					Id: pulumi.String("logs"),
// 				},
// 				&s3control.BucketLifecycleConfigurationRuleArgs{
// 					Expiration: &s3control.BucketLifecycleConfigurationRuleExpirationArgs{
// 						Days: pulumi.Int(7),
// 					},
// 					Filter: &s3control.BucketLifecycleConfigurationRuleFilterArgs{
// 						Prefix: pulumi.String("temp/"),
// 					},
// 					Id: pulumi.String("temp"),
// 				},
// 			},
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
// S3 Control Bucket Lifecycle Configurations can be imported using the Amazon Resource Name (ARN), e.g.
//
// ```sh
//  $ pulumi import aws:s3control/bucketLifecycleConfiguration:BucketLifecycleConfiguration example arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-12345678/bucket/example
// ```
type BucketLifecycleConfiguration struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Configuration block(s) containing lifecycle rules for the bucket.
	Rules BucketLifecycleConfigurationRuleArrayOutput `pulumi:"rules"`
}

// NewBucketLifecycleConfiguration registers a new resource with the given unique name, arguments, and options.
func NewBucketLifecycleConfiguration(ctx *pulumi.Context,
	name string, args *BucketLifecycleConfigurationArgs, opts ...pulumi.ResourceOption) (*BucketLifecycleConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	var resource BucketLifecycleConfiguration
	err := ctx.RegisterResource("aws:s3control/bucketLifecycleConfiguration:BucketLifecycleConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketLifecycleConfiguration gets an existing BucketLifecycleConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketLifecycleConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketLifecycleConfigurationState, opts ...pulumi.ResourceOption) (*BucketLifecycleConfiguration, error) {
	var resource BucketLifecycleConfiguration
	err := ctx.ReadResource("aws:s3control/bucketLifecycleConfiguration:BucketLifecycleConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketLifecycleConfiguration resources.
type bucketLifecycleConfigurationState struct {
	// Amazon Resource Name (ARN) of the bucket.
	Bucket *string `pulumi:"bucket"`
	// Configuration block(s) containing lifecycle rules for the bucket.
	Rules []BucketLifecycleConfigurationRule `pulumi:"rules"`
}

type BucketLifecycleConfigurationState struct {
	// Amazon Resource Name (ARN) of the bucket.
	Bucket pulumi.StringPtrInput
	// Configuration block(s) containing lifecycle rules for the bucket.
	Rules BucketLifecycleConfigurationRuleArrayInput
}

func (BucketLifecycleConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketLifecycleConfigurationState)(nil)).Elem()
}

type bucketLifecycleConfigurationArgs struct {
	// Amazon Resource Name (ARN) of the bucket.
	Bucket string `pulumi:"bucket"`
	// Configuration block(s) containing lifecycle rules for the bucket.
	Rules []BucketLifecycleConfigurationRule `pulumi:"rules"`
}

// The set of arguments for constructing a BucketLifecycleConfiguration resource.
type BucketLifecycleConfigurationArgs struct {
	// Amazon Resource Name (ARN) of the bucket.
	Bucket pulumi.StringInput
	// Configuration block(s) containing lifecycle rules for the bucket.
	Rules BucketLifecycleConfigurationRuleArrayInput
}

func (BucketLifecycleConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketLifecycleConfigurationArgs)(nil)).Elem()
}

type BucketLifecycleConfigurationInput interface {
	pulumi.Input

	ToBucketLifecycleConfigurationOutput() BucketLifecycleConfigurationOutput
	ToBucketLifecycleConfigurationOutputWithContext(ctx context.Context) BucketLifecycleConfigurationOutput
}

func (*BucketLifecycleConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketLifecycleConfiguration)(nil))
}

func (i *BucketLifecycleConfiguration) ToBucketLifecycleConfigurationOutput() BucketLifecycleConfigurationOutput {
	return i.ToBucketLifecycleConfigurationOutputWithContext(context.Background())
}

func (i *BucketLifecycleConfiguration) ToBucketLifecycleConfigurationOutputWithContext(ctx context.Context) BucketLifecycleConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLifecycleConfigurationOutput)
}

func (i *BucketLifecycleConfiguration) ToBucketLifecycleConfigurationPtrOutput() BucketLifecycleConfigurationPtrOutput {
	return i.ToBucketLifecycleConfigurationPtrOutputWithContext(context.Background())
}

func (i *BucketLifecycleConfiguration) ToBucketLifecycleConfigurationPtrOutputWithContext(ctx context.Context) BucketLifecycleConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLifecycleConfigurationPtrOutput)
}

type BucketLifecycleConfigurationPtrInput interface {
	pulumi.Input

	ToBucketLifecycleConfigurationPtrOutput() BucketLifecycleConfigurationPtrOutput
	ToBucketLifecycleConfigurationPtrOutputWithContext(ctx context.Context) BucketLifecycleConfigurationPtrOutput
}

type bucketLifecycleConfigurationPtrType BucketLifecycleConfigurationArgs

func (*bucketLifecycleConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketLifecycleConfiguration)(nil))
}

func (i *bucketLifecycleConfigurationPtrType) ToBucketLifecycleConfigurationPtrOutput() BucketLifecycleConfigurationPtrOutput {
	return i.ToBucketLifecycleConfigurationPtrOutputWithContext(context.Background())
}

func (i *bucketLifecycleConfigurationPtrType) ToBucketLifecycleConfigurationPtrOutputWithContext(ctx context.Context) BucketLifecycleConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLifecycleConfigurationPtrOutput)
}

// BucketLifecycleConfigurationArrayInput is an input type that accepts BucketLifecycleConfigurationArray and BucketLifecycleConfigurationArrayOutput values.
// You can construct a concrete instance of `BucketLifecycleConfigurationArrayInput` via:
//
//          BucketLifecycleConfigurationArray{ BucketLifecycleConfigurationArgs{...} }
type BucketLifecycleConfigurationArrayInput interface {
	pulumi.Input

	ToBucketLifecycleConfigurationArrayOutput() BucketLifecycleConfigurationArrayOutput
	ToBucketLifecycleConfigurationArrayOutputWithContext(context.Context) BucketLifecycleConfigurationArrayOutput
}

type BucketLifecycleConfigurationArray []BucketLifecycleConfigurationInput

func (BucketLifecycleConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*BucketLifecycleConfiguration)(nil))
}

func (i BucketLifecycleConfigurationArray) ToBucketLifecycleConfigurationArrayOutput() BucketLifecycleConfigurationArrayOutput {
	return i.ToBucketLifecycleConfigurationArrayOutputWithContext(context.Background())
}

func (i BucketLifecycleConfigurationArray) ToBucketLifecycleConfigurationArrayOutputWithContext(ctx context.Context) BucketLifecycleConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLifecycleConfigurationArrayOutput)
}

// BucketLifecycleConfigurationMapInput is an input type that accepts BucketLifecycleConfigurationMap and BucketLifecycleConfigurationMapOutput values.
// You can construct a concrete instance of `BucketLifecycleConfigurationMapInput` via:
//
//          BucketLifecycleConfigurationMap{ "key": BucketLifecycleConfigurationArgs{...} }
type BucketLifecycleConfigurationMapInput interface {
	pulumi.Input

	ToBucketLifecycleConfigurationMapOutput() BucketLifecycleConfigurationMapOutput
	ToBucketLifecycleConfigurationMapOutputWithContext(context.Context) BucketLifecycleConfigurationMapOutput
}

type BucketLifecycleConfigurationMap map[string]BucketLifecycleConfigurationInput

func (BucketLifecycleConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*BucketLifecycleConfiguration)(nil))
}

func (i BucketLifecycleConfigurationMap) ToBucketLifecycleConfigurationMapOutput() BucketLifecycleConfigurationMapOutput {
	return i.ToBucketLifecycleConfigurationMapOutputWithContext(context.Background())
}

func (i BucketLifecycleConfigurationMap) ToBucketLifecycleConfigurationMapOutputWithContext(ctx context.Context) BucketLifecycleConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLifecycleConfigurationMapOutput)
}

type BucketLifecycleConfigurationOutput struct {
	*pulumi.OutputState
}

func (BucketLifecycleConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketLifecycleConfiguration)(nil))
}

func (o BucketLifecycleConfigurationOutput) ToBucketLifecycleConfigurationOutput() BucketLifecycleConfigurationOutput {
	return o
}

func (o BucketLifecycleConfigurationOutput) ToBucketLifecycleConfigurationOutputWithContext(ctx context.Context) BucketLifecycleConfigurationOutput {
	return o
}

func (o BucketLifecycleConfigurationOutput) ToBucketLifecycleConfigurationPtrOutput() BucketLifecycleConfigurationPtrOutput {
	return o.ToBucketLifecycleConfigurationPtrOutputWithContext(context.Background())
}

func (o BucketLifecycleConfigurationOutput) ToBucketLifecycleConfigurationPtrOutputWithContext(ctx context.Context) BucketLifecycleConfigurationPtrOutput {
	return o.ApplyT(func(v BucketLifecycleConfiguration) *BucketLifecycleConfiguration {
		return &v
	}).(BucketLifecycleConfigurationPtrOutput)
}

type BucketLifecycleConfigurationPtrOutput struct {
	*pulumi.OutputState
}

func (BucketLifecycleConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketLifecycleConfiguration)(nil))
}

func (o BucketLifecycleConfigurationPtrOutput) ToBucketLifecycleConfigurationPtrOutput() BucketLifecycleConfigurationPtrOutput {
	return o
}

func (o BucketLifecycleConfigurationPtrOutput) ToBucketLifecycleConfigurationPtrOutputWithContext(ctx context.Context) BucketLifecycleConfigurationPtrOutput {
	return o
}

type BucketLifecycleConfigurationArrayOutput struct{ *pulumi.OutputState }

func (BucketLifecycleConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BucketLifecycleConfiguration)(nil))
}

func (o BucketLifecycleConfigurationArrayOutput) ToBucketLifecycleConfigurationArrayOutput() BucketLifecycleConfigurationArrayOutput {
	return o
}

func (o BucketLifecycleConfigurationArrayOutput) ToBucketLifecycleConfigurationArrayOutputWithContext(ctx context.Context) BucketLifecycleConfigurationArrayOutput {
	return o
}

func (o BucketLifecycleConfigurationArrayOutput) Index(i pulumi.IntInput) BucketLifecycleConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BucketLifecycleConfiguration {
		return vs[0].([]BucketLifecycleConfiguration)[vs[1].(int)]
	}).(BucketLifecycleConfigurationOutput)
}

type BucketLifecycleConfigurationMapOutput struct{ *pulumi.OutputState }

func (BucketLifecycleConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BucketLifecycleConfiguration)(nil))
}

func (o BucketLifecycleConfigurationMapOutput) ToBucketLifecycleConfigurationMapOutput() BucketLifecycleConfigurationMapOutput {
	return o
}

func (o BucketLifecycleConfigurationMapOutput) ToBucketLifecycleConfigurationMapOutputWithContext(ctx context.Context) BucketLifecycleConfigurationMapOutput {
	return o
}

func (o BucketLifecycleConfigurationMapOutput) MapIndex(k pulumi.StringInput) BucketLifecycleConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) BucketLifecycleConfiguration {
		return vs[0].(map[string]BucketLifecycleConfiguration)[vs[1].(string)]
	}).(BucketLifecycleConfigurationOutput)
}

func init() {
	pulumi.RegisterOutputType(BucketLifecycleConfigurationOutput{})
	pulumi.RegisterOutputType(BucketLifecycleConfigurationPtrOutput{})
	pulumi.RegisterOutputType(BucketLifecycleConfigurationArrayOutput{})
	pulumi.RegisterOutputType(BucketLifecycleConfigurationMapOutput{})
}
