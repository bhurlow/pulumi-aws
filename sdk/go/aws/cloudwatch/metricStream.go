// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CloudWatch Metric Stream resource.
//
// ## Example Usage
// ### Filters
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			streamsAssumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Effect: pulumi.StringRef("Allow"),
//						Principals: []iam.GetPolicyDocumentStatementPrincipal{
//							{
//								Type: "Service",
//								Identifiers: []string{
//									"streams.metrics.cloudwatch.amazonaws.com",
//								},
//							},
//						},
//						Actions: []string{
//							"sts:AssumeRole",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			metricStreamToFirehoseRole, err := iam.NewRole(ctx, "metricStreamToFirehoseRole", &iam.RoleArgs{
//				AssumeRolePolicy: *pulumi.String(streamsAssumeRole.Json),
//			})
//			if err != nil {
//				return err
//			}
//			bucket, err := s3.NewBucketV2(ctx, "bucket", nil)
//			if err != nil {
//				return err
//			}
//			firehoseAssumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Effect: pulumi.StringRef("Allow"),
//						Principals: []iam.GetPolicyDocumentStatementPrincipal{
//							{
//								Type: "Service",
//								Identifiers: []string{
//									"firehose.amazonaws.com",
//								},
//							},
//						},
//						Actions: []string{
//							"sts:AssumeRole",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			firehoseToS3Role, err := iam.NewRole(ctx, "firehoseToS3Role", &iam.RoleArgs{
//				AssumeRolePolicy: *pulumi.String(firehoseAssumeRole.Json),
//			})
//			if err != nil {
//				return err
//			}
//			s3Stream, err := kinesis.NewFirehoseDeliveryStream(ctx, "s3Stream", &kinesis.FirehoseDeliveryStreamArgs{
//				Destination: pulumi.String("extended_s3"),
//				ExtendedS3Configuration: &kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationArgs{
//					RoleArn:   firehoseToS3Role.Arn,
//					BucketArn: bucket.Arn,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudwatch.NewMetricStream(ctx, "main", &cloudwatch.MetricStreamArgs{
//				RoleArn:      metricStreamToFirehoseRole.Arn,
//				FirehoseArn:  s3Stream.Arn,
//				OutputFormat: pulumi.String("json"),
//				IncludeFilters: cloudwatch.MetricStreamIncludeFilterArray{
//					&cloudwatch.MetricStreamIncludeFilterArgs{
//						Namespace: pulumi.String("AWS/EC2"),
//						MetricNames: pulumi.StringArray{
//							pulumi.String("CPUUtilization"),
//							pulumi.String("NetworkOut"),
//						},
//					},
//					&cloudwatch.MetricStreamIncludeFilterArgs{
//						Namespace:   pulumi.String("AWS/EBS"),
//						MetricNames: pulumi.StringArray{},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			metricStreamToFirehosePolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
//				Statements: iam.GetPolicyDocumentStatementArray{
//					&iam.GetPolicyDocumentStatementArgs{
//						Effect: pulumi.String("Allow"),
//						Actions: pulumi.StringArray{
//							pulumi.String("firehose:PutRecord"),
//							pulumi.String("firehose:PutRecordBatch"),
//						},
//						Resources: pulumi.StringArray{
//							s3Stream.Arn,
//						},
//					},
//				},
//			}, nil)
//			_, err = iam.NewRolePolicy(ctx, "metricStreamToFirehoseRolePolicy", &iam.RolePolicyArgs{
//				Role: metricStreamToFirehoseRole.ID(),
//				Policy: metricStreamToFirehosePolicyDocument.ApplyT(func(metricStreamToFirehosePolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
//					return &metricStreamToFirehosePolicyDocument.Json, nil
//				}).(pulumi.StringPtrOutput),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketAclV2(ctx, "bucketAcl", &s3.BucketAclV2Args{
//				Bucket: bucket.ID(),
//				Acl:    pulumi.String("private"),
//			})
//			if err != nil {
//				return err
//			}
//			firehoseToS3PolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
//				Statements: iam.GetPolicyDocumentStatementArray{
//					&iam.GetPolicyDocumentStatementArgs{
//						Effect: pulumi.String("Allow"),
//						Actions: pulumi.StringArray{
//							pulumi.String("s3:AbortMultipartUpload"),
//							pulumi.String("s3:GetBucketLocation"),
//							pulumi.String("s3:GetObject"),
//							pulumi.String("s3:ListBucket"),
//							pulumi.String("s3:ListBucketMultipartUploads"),
//							pulumi.String("s3:PutObject"),
//						},
//						Resources: pulumi.StringArray{
//							bucket.Arn,
//							bucket.Arn.ApplyT(func(arn string) (string, error) {
//								return fmt.Sprintf("%v/*", arn), nil
//							}).(pulumi.StringOutput),
//						},
//					},
//				},
//			}, nil)
//			_, err = iam.NewRolePolicy(ctx, "firehoseToS3RolePolicy", &iam.RolePolicyArgs{
//				Role: firehoseToS3Role.ID(),
//				Policy: firehoseToS3PolicyDocument.ApplyT(func(firehoseToS3PolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
//					return &firehoseToS3PolicyDocument.Json, nil
//				}).(pulumi.StringPtrOutput),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Additional Statistics
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudwatch.NewMetricStream(ctx, "main", &cloudwatch.MetricStreamArgs{
//				RoleArn:      pulumi.Any(aws_iam_role.Metric_stream_to_firehose.Arn),
//				FirehoseArn:  pulumi.Any(aws_kinesis_firehose_delivery_stream.S3_stream.Arn),
//				OutputFormat: pulumi.String("json"),
//				StatisticsConfigurations: cloudwatch.MetricStreamStatisticsConfigurationArray{
//					&cloudwatch.MetricStreamStatisticsConfigurationArgs{
//						AdditionalStatistics: pulumi.StringArray{
//							pulumi.String("p1"),
//							pulumi.String("tm99"),
//						},
//						IncludeMetrics: cloudwatch.MetricStreamStatisticsConfigurationIncludeMetricArray{
//							&cloudwatch.MetricStreamStatisticsConfigurationIncludeMetricArgs{
//								MetricName: pulumi.String("CPUUtilization"),
//								Namespace:  pulumi.String("AWS/EC2"),
//							},
//						},
//					},
//					&cloudwatch.MetricStreamStatisticsConfigurationArgs{
//						AdditionalStatistics: pulumi.StringArray{
//							pulumi.String("TS(50.5:)"),
//						},
//						IncludeMetrics: cloudwatch.MetricStreamStatisticsConfigurationIncludeMetricArray{
//							&cloudwatch.MetricStreamStatisticsConfigurationIncludeMetricArgs{
//								MetricName: pulumi.String("CPUUtilization"),
//								Namespace:  pulumi.String("AWS/EC2"),
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
//	to = aws_cloudwatch_metric_stream.sample
//
//	id = "sample-stream-name" } Using `pulumi import`, import CloudWatch metric streams using the `name`. For exampleconsole % pulumi import aws_cloudwatch_metric_stream.sample sample-stream-name
type MetricStream struct {
	pulumi.CustomResourceState

	// ARN of the metric stream.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the metric stream was created.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// List of exclusive metric filters. If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces and the conditional metric names that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is excluded. Conflicts with `includeFilter`.
	ExcludeFilters MetricStreamExcludeFilterArrayOutput `pulumi:"excludeFilters"`
	// ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
	FirehoseArn pulumi.StringOutput `pulumi:"firehoseArn"`
	// List of inclusive metric filters. If you specify this parameter, the stream sends only the conditional metric names from the metric namespaces that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is included. Conflicts with `excludeFilter`.
	IncludeFilters MetricStreamIncludeFilterArrayOutput `pulumi:"includeFilters"`
	// If you are creating a metric stream in a monitoring account, specify true to include metrics from source accounts that are linked to this monitoring account, in the metric stream. The default is false. For more information about linking accounts, see [CloudWatch cross-account observability](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html).
	IncludeLinkedAccountsMetrics pulumi.BoolPtrOutput `pulumi:"includeLinkedAccountsMetrics"`
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the metric stream was last updated.
	LastUpdateDate pulumi.StringOutput `pulumi:"lastUpdateDate"`
	// Friendly name of the metric stream. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// Output format for the stream. Possible values are `json` and `opentelemetry0.7`. For more information about output formats, see [Metric streams output formats](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html).
	//
	// The following arguments are optional:
	OutputFormat pulumi.StringOutput `pulumi:"outputFormat"`
	// ARN of the IAM role that this metric stream will use to access Amazon Kinesis Firehose resources. For more information about role permissions, see [Trust between CloudWatch and Kinesis Data Firehose](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-trustpolicy.html).
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// State of the metric stream. Possible values are `running` and `stopped`.
	State pulumi.StringOutput `pulumi:"state"`
	// For each entry in this array, you specify one or more metrics and the list of additional statistics to stream for those metrics. The additional statistics that you can stream depend on the stream's `outputFormat`. If the OutputFormat is `json`, you can stream any additional statistic that is supported by CloudWatch, listed in [CloudWatch statistics definitions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.html). If the OutputFormat is `opentelemetry0.7`, you can stream percentile statistics (p99 etc.). See details below.
	StatisticsConfigurations MetricStreamStatisticsConfigurationArrayOutput `pulumi:"statisticsConfigurations"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewMetricStream registers a new resource with the given unique name, arguments, and options.
func NewMetricStream(ctx *pulumi.Context,
	name string, args *MetricStreamArgs, opts ...pulumi.ResourceOption) (*MetricStream, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FirehoseArn == nil {
		return nil, errors.New("invalid value for required argument 'FirehoseArn'")
	}
	if args.OutputFormat == nil {
		return nil, errors.New("invalid value for required argument 'OutputFormat'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MetricStream
	err := ctx.RegisterResource("aws:cloudwatch/metricStream:MetricStream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetricStream gets an existing MetricStream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetricStream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetricStreamState, opts ...pulumi.ResourceOption) (*MetricStream, error) {
	var resource MetricStream
	err := ctx.ReadResource("aws:cloudwatch/metricStream:MetricStream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetricStream resources.
type metricStreamState struct {
	// ARN of the metric stream.
	Arn *string `pulumi:"arn"`
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the metric stream was created.
	CreationDate *string `pulumi:"creationDate"`
	// List of exclusive metric filters. If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces and the conditional metric names that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is excluded. Conflicts with `includeFilter`.
	ExcludeFilters []MetricStreamExcludeFilter `pulumi:"excludeFilters"`
	// ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
	FirehoseArn *string `pulumi:"firehoseArn"`
	// List of inclusive metric filters. If you specify this parameter, the stream sends only the conditional metric names from the metric namespaces that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is included. Conflicts with `excludeFilter`.
	IncludeFilters []MetricStreamIncludeFilter `pulumi:"includeFilters"`
	// If you are creating a metric stream in a monitoring account, specify true to include metrics from source accounts that are linked to this monitoring account, in the metric stream. The default is false. For more information about linking accounts, see [CloudWatch cross-account observability](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html).
	IncludeLinkedAccountsMetrics *bool `pulumi:"includeLinkedAccountsMetrics"`
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the metric stream was last updated.
	LastUpdateDate *string `pulumi:"lastUpdateDate"`
	// Friendly name of the metric stream. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name *string `pulumi:"name"`
	// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Output format for the stream. Possible values are `json` and `opentelemetry0.7`. For more information about output formats, see [Metric streams output formats](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html).
	//
	// The following arguments are optional:
	OutputFormat *string `pulumi:"outputFormat"`
	// ARN of the IAM role that this metric stream will use to access Amazon Kinesis Firehose resources. For more information about role permissions, see [Trust between CloudWatch and Kinesis Data Firehose](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-trustpolicy.html).
	RoleArn *string `pulumi:"roleArn"`
	// State of the metric stream. Possible values are `running` and `stopped`.
	State *string `pulumi:"state"`
	// For each entry in this array, you specify one or more metrics and the list of additional statistics to stream for those metrics. The additional statistics that you can stream depend on the stream's `outputFormat`. If the OutputFormat is `json`, you can stream any additional statistic that is supported by CloudWatch, listed in [CloudWatch statistics definitions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.html). If the OutputFormat is `opentelemetry0.7`, you can stream percentile statistics (p99 etc.). See details below.
	StatisticsConfigurations []MetricStreamStatisticsConfiguration `pulumi:"statisticsConfigurations"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type MetricStreamState struct {
	// ARN of the metric stream.
	Arn pulumi.StringPtrInput
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the metric stream was created.
	CreationDate pulumi.StringPtrInput
	// List of exclusive metric filters. If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces and the conditional metric names that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is excluded. Conflicts with `includeFilter`.
	ExcludeFilters MetricStreamExcludeFilterArrayInput
	// ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
	FirehoseArn pulumi.StringPtrInput
	// List of inclusive metric filters. If you specify this parameter, the stream sends only the conditional metric names from the metric namespaces that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is included. Conflicts with `excludeFilter`.
	IncludeFilters MetricStreamIncludeFilterArrayInput
	// If you are creating a metric stream in a monitoring account, specify true to include metrics from source accounts that are linked to this monitoring account, in the metric stream. The default is false. For more information about linking accounts, see [CloudWatch cross-account observability](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html).
	IncludeLinkedAccountsMetrics pulumi.BoolPtrInput
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the metric stream was last updated.
	LastUpdateDate pulumi.StringPtrInput
	// Friendly name of the metric stream. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name pulumi.StringPtrInput
	// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Output format for the stream. Possible values are `json` and `opentelemetry0.7`. For more information about output formats, see [Metric streams output formats](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html).
	//
	// The following arguments are optional:
	OutputFormat pulumi.StringPtrInput
	// ARN of the IAM role that this metric stream will use to access Amazon Kinesis Firehose resources. For more information about role permissions, see [Trust between CloudWatch and Kinesis Data Firehose](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-trustpolicy.html).
	RoleArn pulumi.StringPtrInput
	// State of the metric stream. Possible values are `running` and `stopped`.
	State pulumi.StringPtrInput
	// For each entry in this array, you specify one or more metrics and the list of additional statistics to stream for those metrics. The additional statistics that you can stream depend on the stream's `outputFormat`. If the OutputFormat is `json`, you can stream any additional statistic that is supported by CloudWatch, listed in [CloudWatch statistics definitions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.html). If the OutputFormat is `opentelemetry0.7`, you can stream percentile statistics (p99 etc.). See details below.
	StatisticsConfigurations MetricStreamStatisticsConfigurationArrayInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (MetricStreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*metricStreamState)(nil)).Elem()
}

type metricStreamArgs struct {
	// List of exclusive metric filters. If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces and the conditional metric names that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is excluded. Conflicts with `includeFilter`.
	ExcludeFilters []MetricStreamExcludeFilter `pulumi:"excludeFilters"`
	// ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
	FirehoseArn string `pulumi:"firehoseArn"`
	// List of inclusive metric filters. If you specify this parameter, the stream sends only the conditional metric names from the metric namespaces that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is included. Conflicts with `excludeFilter`.
	IncludeFilters []MetricStreamIncludeFilter `pulumi:"includeFilters"`
	// If you are creating a metric stream in a monitoring account, specify true to include metrics from source accounts that are linked to this monitoring account, in the metric stream. The default is false. For more information about linking accounts, see [CloudWatch cross-account observability](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html).
	IncludeLinkedAccountsMetrics *bool `pulumi:"includeLinkedAccountsMetrics"`
	// Friendly name of the metric stream. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name *string `pulumi:"name"`
	// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Output format for the stream. Possible values are `json` and `opentelemetry0.7`. For more information about output formats, see [Metric streams output formats](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html).
	//
	// The following arguments are optional:
	OutputFormat string `pulumi:"outputFormat"`
	// ARN of the IAM role that this metric stream will use to access Amazon Kinesis Firehose resources. For more information about role permissions, see [Trust between CloudWatch and Kinesis Data Firehose](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-trustpolicy.html).
	RoleArn string `pulumi:"roleArn"`
	// For each entry in this array, you specify one or more metrics and the list of additional statistics to stream for those metrics. The additional statistics that you can stream depend on the stream's `outputFormat`. If the OutputFormat is `json`, you can stream any additional statistic that is supported by CloudWatch, listed in [CloudWatch statistics definitions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.html). If the OutputFormat is `opentelemetry0.7`, you can stream percentile statistics (p99 etc.). See details below.
	StatisticsConfigurations []MetricStreamStatisticsConfiguration `pulumi:"statisticsConfigurations"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MetricStream resource.
type MetricStreamArgs struct {
	// List of exclusive metric filters. If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces and the conditional metric names that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is excluded. Conflicts with `includeFilter`.
	ExcludeFilters MetricStreamExcludeFilterArrayInput
	// ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
	FirehoseArn pulumi.StringInput
	// List of inclusive metric filters. If you specify this parameter, the stream sends only the conditional metric names from the metric namespaces that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is included. Conflicts with `excludeFilter`.
	IncludeFilters MetricStreamIncludeFilterArrayInput
	// If you are creating a metric stream in a monitoring account, specify true to include metrics from source accounts that are linked to this monitoring account, in the metric stream. The default is false. For more information about linking accounts, see [CloudWatch cross-account observability](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html).
	IncludeLinkedAccountsMetrics pulumi.BoolPtrInput
	// Friendly name of the metric stream. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name pulumi.StringPtrInput
	// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Output format for the stream. Possible values are `json` and `opentelemetry0.7`. For more information about output formats, see [Metric streams output formats](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html).
	//
	// The following arguments are optional:
	OutputFormat pulumi.StringInput
	// ARN of the IAM role that this metric stream will use to access Amazon Kinesis Firehose resources. For more information about role permissions, see [Trust between CloudWatch and Kinesis Data Firehose](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-trustpolicy.html).
	RoleArn pulumi.StringInput
	// For each entry in this array, you specify one or more metrics and the list of additional statistics to stream for those metrics. The additional statistics that you can stream depend on the stream's `outputFormat`. If the OutputFormat is `json`, you can stream any additional statistic that is supported by CloudWatch, listed in [CloudWatch statistics definitions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.html). If the OutputFormat is `opentelemetry0.7`, you can stream percentile statistics (p99 etc.). See details below.
	StatisticsConfigurations MetricStreamStatisticsConfigurationArrayInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (MetricStreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metricStreamArgs)(nil)).Elem()
}

type MetricStreamInput interface {
	pulumi.Input

	ToMetricStreamOutput() MetricStreamOutput
	ToMetricStreamOutputWithContext(ctx context.Context) MetricStreamOutput
}

func (*MetricStream) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricStream)(nil)).Elem()
}

func (i *MetricStream) ToMetricStreamOutput() MetricStreamOutput {
	return i.ToMetricStreamOutputWithContext(context.Background())
}

func (i *MetricStream) ToMetricStreamOutputWithContext(ctx context.Context) MetricStreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricStreamOutput)
}

// MetricStreamArrayInput is an input type that accepts MetricStreamArray and MetricStreamArrayOutput values.
// You can construct a concrete instance of `MetricStreamArrayInput` via:
//
//	MetricStreamArray{ MetricStreamArgs{...} }
type MetricStreamArrayInput interface {
	pulumi.Input

	ToMetricStreamArrayOutput() MetricStreamArrayOutput
	ToMetricStreamArrayOutputWithContext(context.Context) MetricStreamArrayOutput
}

type MetricStreamArray []MetricStreamInput

func (MetricStreamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetricStream)(nil)).Elem()
}

func (i MetricStreamArray) ToMetricStreamArrayOutput() MetricStreamArrayOutput {
	return i.ToMetricStreamArrayOutputWithContext(context.Background())
}

func (i MetricStreamArray) ToMetricStreamArrayOutputWithContext(ctx context.Context) MetricStreamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricStreamArrayOutput)
}

// MetricStreamMapInput is an input type that accepts MetricStreamMap and MetricStreamMapOutput values.
// You can construct a concrete instance of `MetricStreamMapInput` via:
//
//	MetricStreamMap{ "key": MetricStreamArgs{...} }
type MetricStreamMapInput interface {
	pulumi.Input

	ToMetricStreamMapOutput() MetricStreamMapOutput
	ToMetricStreamMapOutputWithContext(context.Context) MetricStreamMapOutput
}

type MetricStreamMap map[string]MetricStreamInput

func (MetricStreamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetricStream)(nil)).Elem()
}

func (i MetricStreamMap) ToMetricStreamMapOutput() MetricStreamMapOutput {
	return i.ToMetricStreamMapOutputWithContext(context.Background())
}

func (i MetricStreamMap) ToMetricStreamMapOutputWithContext(ctx context.Context) MetricStreamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricStreamMapOutput)
}

type MetricStreamOutput struct{ *pulumi.OutputState }

func (MetricStreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricStream)(nil)).Elem()
}

func (o MetricStreamOutput) ToMetricStreamOutput() MetricStreamOutput {
	return o
}

func (o MetricStreamOutput) ToMetricStreamOutputWithContext(ctx context.Context) MetricStreamOutput {
	return o
}

// ARN of the metric stream.
func (o MetricStreamOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricStream) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the metric stream was created.
func (o MetricStreamOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricStream) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// List of exclusive metric filters. If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces and the conditional metric names that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is excluded. Conflicts with `includeFilter`.
func (o MetricStreamOutput) ExcludeFilters() MetricStreamExcludeFilterArrayOutput {
	return o.ApplyT(func(v *MetricStream) MetricStreamExcludeFilterArrayOutput { return v.ExcludeFilters }).(MetricStreamExcludeFilterArrayOutput)
}

// ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
func (o MetricStreamOutput) FirehoseArn() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricStream) pulumi.StringOutput { return v.FirehoseArn }).(pulumi.StringOutput)
}

// List of inclusive metric filters. If you specify this parameter, the stream sends only the conditional metric names from the metric namespaces that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is included. Conflicts with `excludeFilter`.
func (o MetricStreamOutput) IncludeFilters() MetricStreamIncludeFilterArrayOutput {
	return o.ApplyT(func(v *MetricStream) MetricStreamIncludeFilterArrayOutput { return v.IncludeFilters }).(MetricStreamIncludeFilterArrayOutput)
}

// If you are creating a metric stream in a monitoring account, specify true to include metrics from source accounts that are linked to this monitoring account, in the metric stream. The default is false. For more information about linking accounts, see [CloudWatch cross-account observability](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html).
func (o MetricStreamOutput) IncludeLinkedAccountsMetrics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MetricStream) pulumi.BoolPtrOutput { return v.IncludeLinkedAccountsMetrics }).(pulumi.BoolPtrOutput)
}

// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the metric stream was last updated.
func (o MetricStreamOutput) LastUpdateDate() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricStream) pulumi.StringOutput { return v.LastUpdateDate }).(pulumi.StringOutput)
}

// Friendly name of the metric stream. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
func (o MetricStreamOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricStream) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
func (o MetricStreamOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricStream) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

// Output format for the stream. Possible values are `json` and `opentelemetry0.7`. For more information about output formats, see [Metric streams output formats](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html).
//
// The following arguments are optional:
func (o MetricStreamOutput) OutputFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricStream) pulumi.StringOutput { return v.OutputFormat }).(pulumi.StringOutput)
}

// ARN of the IAM role that this metric stream will use to access Amazon Kinesis Firehose resources. For more information about role permissions, see [Trust between CloudWatch and Kinesis Data Firehose](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-trustpolicy.html).
func (o MetricStreamOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricStream) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// State of the metric stream. Possible values are `running` and `stopped`.
func (o MetricStreamOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricStream) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// For each entry in this array, you specify one or more metrics and the list of additional statistics to stream for those metrics. The additional statistics that you can stream depend on the stream's `outputFormat`. If the OutputFormat is `json`, you can stream any additional statistic that is supported by CloudWatch, listed in [CloudWatch statistics definitions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.html). If the OutputFormat is `opentelemetry0.7`, you can stream percentile statistics (p99 etc.). See details below.
func (o MetricStreamOutput) StatisticsConfigurations() MetricStreamStatisticsConfigurationArrayOutput {
	return o.ApplyT(func(v *MetricStream) MetricStreamStatisticsConfigurationArrayOutput {
		return v.StatisticsConfigurations
	}).(MetricStreamStatisticsConfigurationArrayOutput)
}

// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o MetricStreamOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MetricStream) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o MetricStreamOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MetricStream) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type MetricStreamArrayOutput struct{ *pulumi.OutputState }

func (MetricStreamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetricStream)(nil)).Elem()
}

func (o MetricStreamArrayOutput) ToMetricStreamArrayOutput() MetricStreamArrayOutput {
	return o
}

func (o MetricStreamArrayOutput) ToMetricStreamArrayOutputWithContext(ctx context.Context) MetricStreamArrayOutput {
	return o
}

func (o MetricStreamArrayOutput) Index(i pulumi.IntInput) MetricStreamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MetricStream {
		return vs[0].([]*MetricStream)[vs[1].(int)]
	}).(MetricStreamOutput)
}

type MetricStreamMapOutput struct{ *pulumi.OutputState }

func (MetricStreamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetricStream)(nil)).Elem()
}

func (o MetricStreamMapOutput) ToMetricStreamMapOutput() MetricStreamMapOutput {
	return o
}

func (o MetricStreamMapOutput) ToMetricStreamMapOutputWithContext(ctx context.Context) MetricStreamMapOutput {
	return o
}

func (o MetricStreamMapOutput) MapIndex(k pulumi.StringInput) MetricStreamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MetricStream {
		return vs[0].(map[string]*MetricStream)[vs[1].(string)]
	}).(MetricStreamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetricStreamInput)(nil)).Elem(), &MetricStream{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricStreamArrayInput)(nil)).Elem(), MetricStreamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricStreamMapInput)(nil)).Elem(), MetricStreamMap{})
	pulumi.RegisterOutputType(MetricStreamOutput{})
	pulumi.RegisterOutputType(MetricStreamArrayOutput{})
	pulumi.RegisterOutputType(MetricStreamMapOutput{})
}
