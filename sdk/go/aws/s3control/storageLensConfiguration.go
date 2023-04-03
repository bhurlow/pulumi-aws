// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3control

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage an S3 Storage Lens configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3control"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			current, err := aws.GetCallerIdentity(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = s3control.NewStorageLensConfiguration(ctx, "example", &s3control.StorageLensConfigurationArgs{
//				ConfigId: pulumi.String("example-1"),
//				StorageLensConfiguration: &s3control.StorageLensConfigurationStorageLensConfigurationArgs{
//					Enabled: pulumi.Bool(true),
//					AccountLevel: &s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelArgs{
//						ActivityMetrics: &s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelActivityMetricsArgs{
//							Enabled: pulumi.Bool(true),
//						},
//						BucketLevel: &s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelArgs{
//							ActivityMetrics: &s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelActivityMetricsArgs{
//								Enabled: pulumi.Bool(true),
//							},
//						},
//					},
//					DataExport: &s3control.StorageLensConfigurationStorageLensConfigurationDataExportArgs{
//						CloudWatchMetrics: &s3control.StorageLensConfigurationStorageLensConfigurationDataExportCloudWatchMetricsArgs{
//							Enabled: pulumi.Bool(true),
//						},
//						S3BucketDestination: &s3control.StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationArgs{
//							AccountId:           *pulumi.String(current.AccountId),
//							Arn:                 pulumi.Any(aws_s3_bucket.Target.Arn),
//							Format:              pulumi.String("CSV"),
//							OutputSchemaVersion: pulumi.String("V_1"),
//							Encryption: &s3control.StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationEncryptionArgs{
//								SseS3s: s3control.StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationEncryptionSseS3Array{
//									nil,
//								},
//							},
//						},
//					},
//					Exclude: &s3control.StorageLensConfigurationStorageLensConfigurationExcludeArgs{
//						Buckets: pulumi.StringArray{
//							aws_s3_bucket.B1.Arn,
//							aws_s3_bucket.B2.Arn,
//						},
//						Regions: pulumi.StringArray{
//							pulumi.String("us-east-2"),
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
// S3 Storage Lens configurations can be imported using the `account_id` and `config_id`, separated by a colon (`:`), e.g.
//
// ```sh
//
//	$ pulumi import aws:s3control/storageLensConfiguration:StorageLensConfiguration example 123456789012:example-1
//
// ```
type StorageLensConfiguration struct {
	pulumi.CustomResourceState

	// The AWS account ID for the S3 Storage Lens configuration. Defaults to automatically determined account ID of the AWS provider.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The Amazon Resource Name (ARN) of the Amazon Web Services organization.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ID of the S3 Storage Lens configuration.
	ConfigId pulumi.StringOutput `pulumi:"configId"`
	// The S3 Storage Lens configuration. See Storage Lens Configuration below for more details.
	StorageLensConfiguration StorageLensConfigurationStorageLensConfigurationOutput `pulumi:"storageLensConfiguration"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewStorageLensConfiguration registers a new resource with the given unique name, arguments, and options.
func NewStorageLensConfiguration(ctx *pulumi.Context,
	name string, args *StorageLensConfigurationArgs, opts ...pulumi.ResourceOption) (*StorageLensConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.StorageLensConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'StorageLensConfiguration'")
	}
	var resource StorageLensConfiguration
	err := ctx.RegisterResource("aws:s3control/storageLensConfiguration:StorageLensConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageLensConfiguration gets an existing StorageLensConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageLensConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageLensConfigurationState, opts ...pulumi.ResourceOption) (*StorageLensConfiguration, error) {
	var resource StorageLensConfiguration
	err := ctx.ReadResource("aws:s3control/storageLensConfiguration:StorageLensConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageLensConfiguration resources.
type storageLensConfigurationState struct {
	// The AWS account ID for the S3 Storage Lens configuration. Defaults to automatically determined account ID of the AWS provider.
	AccountId *string `pulumi:"accountId"`
	// The Amazon Resource Name (ARN) of the Amazon Web Services organization.
	Arn *string `pulumi:"arn"`
	// The ID of the S3 Storage Lens configuration.
	ConfigId *string `pulumi:"configId"`
	// The S3 Storage Lens configuration. See Storage Lens Configuration below for more details.
	StorageLensConfiguration *StorageLensConfigurationStorageLensConfiguration `pulumi:"storageLensConfiguration"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type StorageLensConfigurationState struct {
	// The AWS account ID for the S3 Storage Lens configuration. Defaults to automatically determined account ID of the AWS provider.
	AccountId pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the Amazon Web Services organization.
	Arn pulumi.StringPtrInput
	// The ID of the S3 Storage Lens configuration.
	ConfigId pulumi.StringPtrInput
	// The S3 Storage Lens configuration. See Storage Lens Configuration below for more details.
	StorageLensConfiguration StorageLensConfigurationStorageLensConfigurationPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (StorageLensConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageLensConfigurationState)(nil)).Elem()
}

type storageLensConfigurationArgs struct {
	// The AWS account ID for the S3 Storage Lens configuration. Defaults to automatically determined account ID of the AWS provider.
	AccountId *string `pulumi:"accountId"`
	// The ID of the S3 Storage Lens configuration.
	ConfigId string `pulumi:"configId"`
	// The S3 Storage Lens configuration. See Storage Lens Configuration below for more details.
	StorageLensConfiguration StorageLensConfigurationStorageLensConfiguration `pulumi:"storageLensConfiguration"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a StorageLensConfiguration resource.
type StorageLensConfigurationArgs struct {
	// The AWS account ID for the S3 Storage Lens configuration. Defaults to automatically determined account ID of the AWS provider.
	AccountId pulumi.StringPtrInput
	// The ID of the S3 Storage Lens configuration.
	ConfigId pulumi.StringInput
	// The S3 Storage Lens configuration. See Storage Lens Configuration below for more details.
	StorageLensConfiguration StorageLensConfigurationStorageLensConfigurationInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (StorageLensConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageLensConfigurationArgs)(nil)).Elem()
}

type StorageLensConfigurationInput interface {
	pulumi.Input

	ToStorageLensConfigurationOutput() StorageLensConfigurationOutput
	ToStorageLensConfigurationOutputWithContext(ctx context.Context) StorageLensConfigurationOutput
}

func (*StorageLensConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageLensConfiguration)(nil)).Elem()
}

func (i *StorageLensConfiguration) ToStorageLensConfigurationOutput() StorageLensConfigurationOutput {
	return i.ToStorageLensConfigurationOutputWithContext(context.Background())
}

func (i *StorageLensConfiguration) ToStorageLensConfigurationOutputWithContext(ctx context.Context) StorageLensConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageLensConfigurationOutput)
}

// StorageLensConfigurationArrayInput is an input type that accepts StorageLensConfigurationArray and StorageLensConfigurationArrayOutput values.
// You can construct a concrete instance of `StorageLensConfigurationArrayInput` via:
//
//	StorageLensConfigurationArray{ StorageLensConfigurationArgs{...} }
type StorageLensConfigurationArrayInput interface {
	pulumi.Input

	ToStorageLensConfigurationArrayOutput() StorageLensConfigurationArrayOutput
	ToStorageLensConfigurationArrayOutputWithContext(context.Context) StorageLensConfigurationArrayOutput
}

type StorageLensConfigurationArray []StorageLensConfigurationInput

func (StorageLensConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StorageLensConfiguration)(nil)).Elem()
}

func (i StorageLensConfigurationArray) ToStorageLensConfigurationArrayOutput() StorageLensConfigurationArrayOutput {
	return i.ToStorageLensConfigurationArrayOutputWithContext(context.Background())
}

func (i StorageLensConfigurationArray) ToStorageLensConfigurationArrayOutputWithContext(ctx context.Context) StorageLensConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageLensConfigurationArrayOutput)
}

// StorageLensConfigurationMapInput is an input type that accepts StorageLensConfigurationMap and StorageLensConfigurationMapOutput values.
// You can construct a concrete instance of `StorageLensConfigurationMapInput` via:
//
//	StorageLensConfigurationMap{ "key": StorageLensConfigurationArgs{...} }
type StorageLensConfigurationMapInput interface {
	pulumi.Input

	ToStorageLensConfigurationMapOutput() StorageLensConfigurationMapOutput
	ToStorageLensConfigurationMapOutputWithContext(context.Context) StorageLensConfigurationMapOutput
}

type StorageLensConfigurationMap map[string]StorageLensConfigurationInput

func (StorageLensConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StorageLensConfiguration)(nil)).Elem()
}

func (i StorageLensConfigurationMap) ToStorageLensConfigurationMapOutput() StorageLensConfigurationMapOutput {
	return i.ToStorageLensConfigurationMapOutputWithContext(context.Background())
}

func (i StorageLensConfigurationMap) ToStorageLensConfigurationMapOutputWithContext(ctx context.Context) StorageLensConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageLensConfigurationMapOutput)
}

type StorageLensConfigurationOutput struct{ *pulumi.OutputState }

func (StorageLensConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageLensConfiguration)(nil)).Elem()
}

func (o StorageLensConfigurationOutput) ToStorageLensConfigurationOutput() StorageLensConfigurationOutput {
	return o
}

func (o StorageLensConfigurationOutput) ToStorageLensConfigurationOutputWithContext(ctx context.Context) StorageLensConfigurationOutput {
	return o
}

// The AWS account ID for the S3 Storage Lens configuration. Defaults to automatically determined account ID of the AWS provider.
func (o StorageLensConfigurationOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageLensConfiguration) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the Amazon Web Services organization.
func (o StorageLensConfigurationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageLensConfiguration) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The ID of the S3 Storage Lens configuration.
func (o StorageLensConfigurationOutput) ConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageLensConfiguration) pulumi.StringOutput { return v.ConfigId }).(pulumi.StringOutput)
}

// The S3 Storage Lens configuration. See Storage Lens Configuration below for more details.
func (o StorageLensConfigurationOutput) StorageLensConfiguration() StorageLensConfigurationStorageLensConfigurationOutput {
	return o.ApplyT(func(v *StorageLensConfiguration) StorageLensConfigurationStorageLensConfigurationOutput {
		return v.StorageLensConfiguration
	}).(StorageLensConfigurationStorageLensConfigurationOutput)
}

// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o StorageLensConfigurationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StorageLensConfiguration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o StorageLensConfigurationOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StorageLensConfiguration) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type StorageLensConfigurationArrayOutput struct{ *pulumi.OutputState }

func (StorageLensConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StorageLensConfiguration)(nil)).Elem()
}

func (o StorageLensConfigurationArrayOutput) ToStorageLensConfigurationArrayOutput() StorageLensConfigurationArrayOutput {
	return o
}

func (o StorageLensConfigurationArrayOutput) ToStorageLensConfigurationArrayOutputWithContext(ctx context.Context) StorageLensConfigurationArrayOutput {
	return o
}

func (o StorageLensConfigurationArrayOutput) Index(i pulumi.IntInput) StorageLensConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StorageLensConfiguration {
		return vs[0].([]*StorageLensConfiguration)[vs[1].(int)]
	}).(StorageLensConfigurationOutput)
}

type StorageLensConfigurationMapOutput struct{ *pulumi.OutputState }

func (StorageLensConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StorageLensConfiguration)(nil)).Elem()
}

func (o StorageLensConfigurationMapOutput) ToStorageLensConfigurationMapOutput() StorageLensConfigurationMapOutput {
	return o
}

func (o StorageLensConfigurationMapOutput) ToStorageLensConfigurationMapOutputWithContext(ctx context.Context) StorageLensConfigurationMapOutput {
	return o
}

func (o StorageLensConfigurationMapOutput) MapIndex(k pulumi.StringInput) StorageLensConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StorageLensConfiguration {
		return vs[0].(map[string]*StorageLensConfiguration)[vs[1].(string)]
	}).(StorageLensConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageLensConfigurationInput)(nil)).Elem(), &StorageLensConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageLensConfigurationArrayInput)(nil)).Elem(), StorageLensConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageLensConfigurationMapInput)(nil)).Elem(), StorageLensConfigurationMap{})
	pulumi.RegisterOutputType(StorageLensConfigurationOutput{})
	pulumi.RegisterOutputType(StorageLensConfigurationArrayOutput{})
	pulumi.RegisterOutputType(StorageLensConfigurationMapOutput{})
}
