// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kinesis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Kinesis Analytics Application resource. Kinesis Analytics is a managed service that
// allows processing and analyzing streaming data using standard SQL.
//
// For more details, see the [Amazon Kinesis Analytics Documentation](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/what-is.html).
//
// > **Note:** To manage Amazon Kinesis Data Analytics for Apache Flink applications, use the `kinesisanalyticsv2.Application` resource.
//
// ## Example Usage
// ### Kinesis Stream Input
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testStream, err := kinesis.NewStream(ctx, "testStream", &kinesis.StreamArgs{
//				ShardCount: pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = kinesis.NewAnalyticsApplication(ctx, "testApplication", &kinesis.AnalyticsApplicationArgs{
//				Inputs: &kinesis.AnalyticsApplicationInputsArgs{
//					NamePrefix: pulumi.String("test_prefix"),
//					KinesisStream: &kinesis.AnalyticsApplicationInputsKinesisStreamArgs{
//						ResourceArn: testStream.Arn,
//						RoleArn:     pulumi.Any(aws_iam_role.Test.Arn),
//					},
//					Parallelism: &kinesis.AnalyticsApplicationInputsParallelismArgs{
//						Count: pulumi.Int(1),
//					},
//					Schema: &kinesis.AnalyticsApplicationInputsSchemaArgs{
//						RecordColumns: kinesis.AnalyticsApplicationInputsSchemaRecordColumnArray{
//							&kinesis.AnalyticsApplicationInputsSchemaRecordColumnArgs{
//								Mapping: pulumi.String("$.test"),
//								Name:    pulumi.String("test"),
//								SqlType: pulumi.String("VARCHAR(8)"),
//							},
//						},
//						RecordEncoding: pulumi.String("UTF-8"),
//						RecordFormat: &kinesis.AnalyticsApplicationInputsSchemaRecordFormatArgs{
//							MappingParameters: &kinesis.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersArgs{
//								Json: &kinesis.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersJsonArgs{
//									RecordRowPath: pulumi.String("$"),
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
// ### Starting An Application
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", nil)
//			if err != nil {
//				return err
//			}
//			exampleLogStream, err := cloudwatch.NewLogStream(ctx, "exampleLogStream", &cloudwatch.LogStreamArgs{
//				LogGroupName: exampleLogGroup.Name,
//			})
//			if err != nil {
//				return err
//			}
//			exampleStream, err := kinesis.NewStream(ctx, "exampleStream", &kinesis.StreamArgs{
//				ShardCount: pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			exampleFirehoseDeliveryStream, err := kinesis.NewFirehoseDeliveryStream(ctx, "exampleFirehoseDeliveryStream", &kinesis.FirehoseDeliveryStreamArgs{
//				Destination: pulumi.String("extended_s3"),
//				ExtendedS3Configuration: &kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationArgs{
//					BucketArn: pulumi.Any(aws_s3_bucket.Example.Arn),
//					RoleArn:   pulumi.Any(aws_iam_role.Example.Arn),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = kinesis.NewAnalyticsApplication(ctx, "test", &kinesis.AnalyticsApplicationArgs{
//				CloudwatchLoggingOptions: &kinesis.AnalyticsApplicationCloudwatchLoggingOptionsArgs{
//					LogStreamArn: exampleLogStream.Arn,
//					RoleArn:      pulumi.Any(aws_iam_role.Example.Arn),
//				},
//				Inputs: &kinesis.AnalyticsApplicationInputsArgs{
//					NamePrefix: pulumi.String("example_prefix"),
//					Schema: &kinesis.AnalyticsApplicationInputsSchemaArgs{
//						RecordColumns: kinesis.AnalyticsApplicationInputsSchemaRecordColumnArray{
//							&kinesis.AnalyticsApplicationInputsSchemaRecordColumnArgs{
//								Name:    pulumi.String("COLUMN_1"),
//								SqlType: pulumi.String("INTEGER"),
//							},
//						},
//						RecordFormat: &kinesis.AnalyticsApplicationInputsSchemaRecordFormatArgs{
//							MappingParameters: &kinesis.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersArgs{
//								Csv: &kinesis.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsvArgs{
//									RecordColumnDelimiter: pulumi.String(","),
//									RecordRowDelimiter:    pulumi.String("|"),
//								},
//							},
//						},
//					},
//					KinesisStream: &kinesis.AnalyticsApplicationInputsKinesisStreamArgs{
//						ResourceArn: exampleStream.Arn,
//						RoleArn:     pulumi.Any(aws_iam_role.Example.Arn),
//					},
//					StartingPositionConfigurations: kinesis.AnalyticsApplicationInputsStartingPositionConfigurationArray{
//						&kinesis.AnalyticsApplicationInputsStartingPositionConfigurationArgs{
//							StartingPosition: pulumi.String("NOW"),
//						},
//					},
//				},
//				Outputs: kinesis.AnalyticsApplicationOutputTypeArray{
//					&kinesis.AnalyticsApplicationOutputTypeArgs{
//						Name: pulumi.String("OUTPUT_1"),
//						Schema: &kinesis.AnalyticsApplicationOutputSchemaArgs{
//							RecordFormatType: pulumi.String("CSV"),
//						},
//						KinesisFirehose: &kinesis.AnalyticsApplicationOutputKinesisFirehoseArgs{
//							ResourceArn: exampleFirehoseDeliveryStream.Arn,
//							RoleArn:     pulumi.Any(aws_iam_role.Example.Arn),
//						},
//					},
//				},
//				StartApplication: pulumi.Bool(true),
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
// Using `pulumi import`, import Kinesis Analytics Application using ARN. For example:
//
// ```sh
//
//	$ pulumi import aws:kinesis/analyticsApplication:AnalyticsApplication example arn:aws:kinesisanalytics:us-west-2:1234567890:application/example
//
// ```
type AnalyticsApplication struct {
	pulumi.CustomResourceState

	// The ARN of the Kinesis Analytics Appliation.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The CloudWatch log stream options to monitor application errors.
	// See CloudWatch Logging Options below for more details.
	CloudwatchLoggingOptions AnalyticsApplicationCloudwatchLoggingOptionsPtrOutput `pulumi:"cloudwatchLoggingOptions"`
	// SQL Code to transform input data, and generate output.
	Code pulumi.StringPtrOutput `pulumi:"code"`
	// The Timestamp when the application version was created.
	CreateTimestamp pulumi.StringOutput `pulumi:"createTimestamp"`
	// Description of the application.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Input configuration of the application. See Inputs below for more details.
	Inputs AnalyticsApplicationInputsPtrOutput `pulumi:"inputs"`
	// The Timestamp when the application was last updated.
	LastUpdateTimestamp pulumi.StringOutput `pulumi:"lastUpdateTimestamp"`
	// Name of the Kinesis Analytics Application.
	Name pulumi.StringOutput `pulumi:"name"`
	// Output destination configuration of the application. See Outputs below for more details.
	Outputs AnalyticsApplicationOutputTypeArrayOutput `pulumi:"outputs"`
	// An S3 Reference Data Source for the application.
	// See Reference Data Sources below for more details.
	ReferenceDataSources AnalyticsApplicationReferenceDataSourcesPtrOutput `pulumi:"referenceDataSources"`
	// Whether to start or stop the Kinesis Analytics Application. To start an application, an input with a defined `startingPosition` must be configured.
	// To modify an application's starting position, first stop the application by setting `startApplication = false`, then update `startingPosition` and set `startApplication = true`.
	StartApplication pulumi.BoolPtrOutput `pulumi:"startApplication"`
	// The Status of the application.
	Status pulumi.StringOutput `pulumi:"status"`
	// Key-value map of tags for the Kinesis Analytics Application. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The Version of the application.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewAnalyticsApplication registers a new resource with the given unique name, arguments, and options.
func NewAnalyticsApplication(ctx *pulumi.Context,
	name string, args *AnalyticsApplicationArgs, opts ...pulumi.ResourceOption) (*AnalyticsApplication, error) {
	if args == nil {
		args = &AnalyticsApplicationArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AnalyticsApplication
	err := ctx.RegisterResource("aws:kinesis/analyticsApplication:AnalyticsApplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnalyticsApplication gets an existing AnalyticsApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalyticsApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalyticsApplicationState, opts ...pulumi.ResourceOption) (*AnalyticsApplication, error) {
	var resource AnalyticsApplication
	err := ctx.ReadResource("aws:kinesis/analyticsApplication:AnalyticsApplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnalyticsApplication resources.
type analyticsApplicationState struct {
	// The ARN of the Kinesis Analytics Appliation.
	Arn *string `pulumi:"arn"`
	// The CloudWatch log stream options to monitor application errors.
	// See CloudWatch Logging Options below for more details.
	CloudwatchLoggingOptions *AnalyticsApplicationCloudwatchLoggingOptions `pulumi:"cloudwatchLoggingOptions"`
	// SQL Code to transform input data, and generate output.
	Code *string `pulumi:"code"`
	// The Timestamp when the application version was created.
	CreateTimestamp *string `pulumi:"createTimestamp"`
	// Description of the application.
	Description *string `pulumi:"description"`
	// Input configuration of the application. See Inputs below for more details.
	Inputs *AnalyticsApplicationInputs `pulumi:"inputs"`
	// The Timestamp when the application was last updated.
	LastUpdateTimestamp *string `pulumi:"lastUpdateTimestamp"`
	// Name of the Kinesis Analytics Application.
	Name *string `pulumi:"name"`
	// Output destination configuration of the application. See Outputs below for more details.
	Outputs []AnalyticsApplicationOutputType `pulumi:"outputs"`
	// An S3 Reference Data Source for the application.
	// See Reference Data Sources below for more details.
	ReferenceDataSources *AnalyticsApplicationReferenceDataSources `pulumi:"referenceDataSources"`
	// Whether to start or stop the Kinesis Analytics Application. To start an application, an input with a defined `startingPosition` must be configured.
	// To modify an application's starting position, first stop the application by setting `startApplication = false`, then update `startingPosition` and set `startApplication = true`.
	StartApplication *bool `pulumi:"startApplication"`
	// The Status of the application.
	Status *string `pulumi:"status"`
	// Key-value map of tags for the Kinesis Analytics Application. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The Version of the application.
	Version *int `pulumi:"version"`
}

type AnalyticsApplicationState struct {
	// The ARN of the Kinesis Analytics Appliation.
	Arn pulumi.StringPtrInput
	// The CloudWatch log stream options to monitor application errors.
	// See CloudWatch Logging Options below for more details.
	CloudwatchLoggingOptions AnalyticsApplicationCloudwatchLoggingOptionsPtrInput
	// SQL Code to transform input data, and generate output.
	Code pulumi.StringPtrInput
	// The Timestamp when the application version was created.
	CreateTimestamp pulumi.StringPtrInput
	// Description of the application.
	Description pulumi.StringPtrInput
	// Input configuration of the application. See Inputs below for more details.
	Inputs AnalyticsApplicationInputsPtrInput
	// The Timestamp when the application was last updated.
	LastUpdateTimestamp pulumi.StringPtrInput
	// Name of the Kinesis Analytics Application.
	Name pulumi.StringPtrInput
	// Output destination configuration of the application. See Outputs below for more details.
	Outputs AnalyticsApplicationOutputTypeArrayInput
	// An S3 Reference Data Source for the application.
	// See Reference Data Sources below for more details.
	ReferenceDataSources AnalyticsApplicationReferenceDataSourcesPtrInput
	// Whether to start or stop the Kinesis Analytics Application. To start an application, an input with a defined `startingPosition` must be configured.
	// To modify an application's starting position, first stop the application by setting `startApplication = false`, then update `startingPosition` and set `startApplication = true`.
	StartApplication pulumi.BoolPtrInput
	// The Status of the application.
	Status pulumi.StringPtrInput
	// Key-value map of tags for the Kinesis Analytics Application. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// The Version of the application.
	Version pulumi.IntPtrInput
}

func (AnalyticsApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsApplicationState)(nil)).Elem()
}

type analyticsApplicationArgs struct {
	// The CloudWatch log stream options to monitor application errors.
	// See CloudWatch Logging Options below for more details.
	CloudwatchLoggingOptions *AnalyticsApplicationCloudwatchLoggingOptions `pulumi:"cloudwatchLoggingOptions"`
	// SQL Code to transform input data, and generate output.
	Code *string `pulumi:"code"`
	// Description of the application.
	Description *string `pulumi:"description"`
	// Input configuration of the application. See Inputs below for more details.
	Inputs *AnalyticsApplicationInputs `pulumi:"inputs"`
	// Name of the Kinesis Analytics Application.
	Name *string `pulumi:"name"`
	// Output destination configuration of the application. See Outputs below for more details.
	Outputs []AnalyticsApplicationOutputType `pulumi:"outputs"`
	// An S3 Reference Data Source for the application.
	// See Reference Data Sources below for more details.
	ReferenceDataSources *AnalyticsApplicationReferenceDataSources `pulumi:"referenceDataSources"`
	// Whether to start or stop the Kinesis Analytics Application. To start an application, an input with a defined `startingPosition` must be configured.
	// To modify an application's starting position, first stop the application by setting `startApplication = false`, then update `startingPosition` and set `startApplication = true`.
	StartApplication *bool `pulumi:"startApplication"`
	// Key-value map of tags for the Kinesis Analytics Application. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AnalyticsApplication resource.
type AnalyticsApplicationArgs struct {
	// The CloudWatch log stream options to monitor application errors.
	// See CloudWatch Logging Options below for more details.
	CloudwatchLoggingOptions AnalyticsApplicationCloudwatchLoggingOptionsPtrInput
	// SQL Code to transform input data, and generate output.
	Code pulumi.StringPtrInput
	// Description of the application.
	Description pulumi.StringPtrInput
	// Input configuration of the application. See Inputs below for more details.
	Inputs AnalyticsApplicationInputsPtrInput
	// Name of the Kinesis Analytics Application.
	Name pulumi.StringPtrInput
	// Output destination configuration of the application. See Outputs below for more details.
	Outputs AnalyticsApplicationOutputTypeArrayInput
	// An S3 Reference Data Source for the application.
	// See Reference Data Sources below for more details.
	ReferenceDataSources AnalyticsApplicationReferenceDataSourcesPtrInput
	// Whether to start or stop the Kinesis Analytics Application. To start an application, an input with a defined `startingPosition` must be configured.
	// To modify an application's starting position, first stop the application by setting `startApplication = false`, then update `startingPosition` and set `startApplication = true`.
	StartApplication pulumi.BoolPtrInput
	// Key-value map of tags for the Kinesis Analytics Application. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (AnalyticsApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsApplicationArgs)(nil)).Elem()
}

type AnalyticsApplicationInput interface {
	pulumi.Input

	ToAnalyticsApplicationOutput() AnalyticsApplicationOutput
	ToAnalyticsApplicationOutputWithContext(ctx context.Context) AnalyticsApplicationOutput
}

func (*AnalyticsApplication) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticsApplication)(nil)).Elem()
}

func (i *AnalyticsApplication) ToAnalyticsApplicationOutput() AnalyticsApplicationOutput {
	return i.ToAnalyticsApplicationOutputWithContext(context.Background())
}

func (i *AnalyticsApplication) ToAnalyticsApplicationOutputWithContext(ctx context.Context) AnalyticsApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsApplicationOutput)
}

// AnalyticsApplicationArrayInput is an input type that accepts AnalyticsApplicationArray and AnalyticsApplicationArrayOutput values.
// You can construct a concrete instance of `AnalyticsApplicationArrayInput` via:
//
//	AnalyticsApplicationArray{ AnalyticsApplicationArgs{...} }
type AnalyticsApplicationArrayInput interface {
	pulumi.Input

	ToAnalyticsApplicationArrayOutput() AnalyticsApplicationArrayOutput
	ToAnalyticsApplicationArrayOutputWithContext(context.Context) AnalyticsApplicationArrayOutput
}

type AnalyticsApplicationArray []AnalyticsApplicationInput

func (AnalyticsApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AnalyticsApplication)(nil)).Elem()
}

func (i AnalyticsApplicationArray) ToAnalyticsApplicationArrayOutput() AnalyticsApplicationArrayOutput {
	return i.ToAnalyticsApplicationArrayOutputWithContext(context.Background())
}

func (i AnalyticsApplicationArray) ToAnalyticsApplicationArrayOutputWithContext(ctx context.Context) AnalyticsApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsApplicationArrayOutput)
}

// AnalyticsApplicationMapInput is an input type that accepts AnalyticsApplicationMap and AnalyticsApplicationMapOutput values.
// You can construct a concrete instance of `AnalyticsApplicationMapInput` via:
//
//	AnalyticsApplicationMap{ "key": AnalyticsApplicationArgs{...} }
type AnalyticsApplicationMapInput interface {
	pulumi.Input

	ToAnalyticsApplicationMapOutput() AnalyticsApplicationMapOutput
	ToAnalyticsApplicationMapOutputWithContext(context.Context) AnalyticsApplicationMapOutput
}

type AnalyticsApplicationMap map[string]AnalyticsApplicationInput

func (AnalyticsApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AnalyticsApplication)(nil)).Elem()
}

func (i AnalyticsApplicationMap) ToAnalyticsApplicationMapOutput() AnalyticsApplicationMapOutput {
	return i.ToAnalyticsApplicationMapOutputWithContext(context.Background())
}

func (i AnalyticsApplicationMap) ToAnalyticsApplicationMapOutputWithContext(ctx context.Context) AnalyticsApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsApplicationMapOutput)
}

type AnalyticsApplicationOutput struct{ *pulumi.OutputState }

func (AnalyticsApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticsApplication)(nil)).Elem()
}

func (o AnalyticsApplicationOutput) ToAnalyticsApplicationOutput() AnalyticsApplicationOutput {
	return o
}

func (o AnalyticsApplicationOutput) ToAnalyticsApplicationOutputWithContext(ctx context.Context) AnalyticsApplicationOutput {
	return o
}

// The ARN of the Kinesis Analytics Appliation.
func (o AnalyticsApplicationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *AnalyticsApplication) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The CloudWatch log stream options to monitor application errors.
// See CloudWatch Logging Options below for more details.
func (o AnalyticsApplicationOutput) CloudwatchLoggingOptions() AnalyticsApplicationCloudwatchLoggingOptionsPtrOutput {
	return o.ApplyT(func(v *AnalyticsApplication) AnalyticsApplicationCloudwatchLoggingOptionsPtrOutput {
		return v.CloudwatchLoggingOptions
	}).(AnalyticsApplicationCloudwatchLoggingOptionsPtrOutput)
}

// SQL Code to transform input data, and generate output.
func (o AnalyticsApplicationOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnalyticsApplication) pulumi.StringPtrOutput { return v.Code }).(pulumi.StringPtrOutput)
}

// The Timestamp when the application version was created.
func (o AnalyticsApplicationOutput) CreateTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *AnalyticsApplication) pulumi.StringOutput { return v.CreateTimestamp }).(pulumi.StringOutput)
}

// Description of the application.
func (o AnalyticsApplicationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnalyticsApplication) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Input configuration of the application. See Inputs below for more details.
func (o AnalyticsApplicationOutput) Inputs() AnalyticsApplicationInputsPtrOutput {
	return o.ApplyT(func(v *AnalyticsApplication) AnalyticsApplicationInputsPtrOutput { return v.Inputs }).(AnalyticsApplicationInputsPtrOutput)
}

// The Timestamp when the application was last updated.
func (o AnalyticsApplicationOutput) LastUpdateTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *AnalyticsApplication) pulumi.StringOutput { return v.LastUpdateTimestamp }).(pulumi.StringOutput)
}

// Name of the Kinesis Analytics Application.
func (o AnalyticsApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AnalyticsApplication) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Output destination configuration of the application. See Outputs below for more details.
func (o AnalyticsApplicationOutput) Outputs() AnalyticsApplicationOutputTypeArrayOutput {
	return o.ApplyT(func(v *AnalyticsApplication) AnalyticsApplicationOutputTypeArrayOutput { return v.Outputs }).(AnalyticsApplicationOutputTypeArrayOutput)
}

// An S3 Reference Data Source for the application.
// See Reference Data Sources below for more details.
func (o AnalyticsApplicationOutput) ReferenceDataSources() AnalyticsApplicationReferenceDataSourcesPtrOutput {
	return o.ApplyT(func(v *AnalyticsApplication) AnalyticsApplicationReferenceDataSourcesPtrOutput {
		return v.ReferenceDataSources
	}).(AnalyticsApplicationReferenceDataSourcesPtrOutput)
}

// Whether to start or stop the Kinesis Analytics Application. To start an application, an input with a defined `startingPosition` must be configured.
// To modify an application's starting position, first stop the application by setting `startApplication = false`, then update `startingPosition` and set `startApplication = true`.
func (o AnalyticsApplicationOutput) StartApplication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AnalyticsApplication) pulumi.BoolPtrOutput { return v.StartApplication }).(pulumi.BoolPtrOutput)
}

// The Status of the application.
func (o AnalyticsApplicationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AnalyticsApplication) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Key-value map of tags for the Kinesis Analytics Application. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o AnalyticsApplicationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AnalyticsApplication) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o AnalyticsApplicationOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AnalyticsApplication) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The Version of the application.
func (o AnalyticsApplicationOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *AnalyticsApplication) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

type AnalyticsApplicationArrayOutput struct{ *pulumi.OutputState }

func (AnalyticsApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AnalyticsApplication)(nil)).Elem()
}

func (o AnalyticsApplicationArrayOutput) ToAnalyticsApplicationArrayOutput() AnalyticsApplicationArrayOutput {
	return o
}

func (o AnalyticsApplicationArrayOutput) ToAnalyticsApplicationArrayOutputWithContext(ctx context.Context) AnalyticsApplicationArrayOutput {
	return o
}

func (o AnalyticsApplicationArrayOutput) Index(i pulumi.IntInput) AnalyticsApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AnalyticsApplication {
		return vs[0].([]*AnalyticsApplication)[vs[1].(int)]
	}).(AnalyticsApplicationOutput)
}

type AnalyticsApplicationMapOutput struct{ *pulumi.OutputState }

func (AnalyticsApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AnalyticsApplication)(nil)).Elem()
}

func (o AnalyticsApplicationMapOutput) ToAnalyticsApplicationMapOutput() AnalyticsApplicationMapOutput {
	return o
}

func (o AnalyticsApplicationMapOutput) ToAnalyticsApplicationMapOutputWithContext(ctx context.Context) AnalyticsApplicationMapOutput {
	return o
}

func (o AnalyticsApplicationMapOutput) MapIndex(k pulumi.StringInput) AnalyticsApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AnalyticsApplication {
		return vs[0].(map[string]*AnalyticsApplication)[vs[1].(string)]
	}).(AnalyticsApplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AnalyticsApplicationInput)(nil)).Elem(), &AnalyticsApplication{})
	pulumi.RegisterInputType(reflect.TypeOf((*AnalyticsApplicationArrayInput)(nil)).Elem(), AnalyticsApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AnalyticsApplicationMapInput)(nil)).Elem(), AnalyticsApplicationMap{})
	pulumi.RegisterOutputType(AnalyticsApplicationOutput{})
	pulumi.RegisterOutputType(AnalyticsApplicationArrayOutput{})
	pulumi.RegisterOutputType(AnalyticsApplicationMapOutput{})
}
