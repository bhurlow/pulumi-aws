// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPC/Subnet/ENI/Transit Gateway/Transit Gateway Attachment Flow Log to capture IP traffic for a specific network
// interface, subnet, or VPC. Logs are sent to a CloudWatch Log Group, a S3 Bucket, or Amazon Kinesis Data Firehose
//
// ## Example Usage
// ### CloudWatch Logging
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudwatch"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
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
//			assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Effect: pulumi.StringRef("Allow"),
//						Principals: []iam.GetPolicyDocumentStatementPrincipal{
//							{
//								Type: "Service",
//								Identifiers: []string{
//									"vpc-flow-logs.amazonaws.com",
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
//			exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
//				AssumeRolePolicy: *pulumi.String(assumeRole.Json),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewFlowLog(ctx, "exampleFlowLog", &ec2.FlowLogArgs{
//				IamRoleArn:     exampleRole.Arn,
//				LogDestination: exampleLogGroup.Arn,
//				TrafficType:    pulumi.String("ALL"),
//				VpcId:          pulumi.Any(aws_vpc.Example.Id),
//			})
//			if err != nil {
//				return err
//			}
//			examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Effect: pulumi.StringRef("Allow"),
//						Actions: []string{
//							"logs:CreateLogGroup",
//							"logs:CreateLogStream",
//							"logs:PutLogEvents",
//							"logs:DescribeLogGroups",
//							"logs:DescribeLogStreams",
//						},
//						Resources: []string{
//							"*",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = iam.NewRolePolicy(ctx, "exampleRolePolicy", &iam.RolePolicyArgs{
//				Role:   exampleRole.ID(),
//				Policy: *pulumi.String(examplePolicyDocument.Json),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### S3 Logging
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
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
//			_, err = ec2.NewFlowLog(ctx, "exampleFlowLog", &ec2.FlowLogArgs{
//				LogDestination:     exampleBucketV2.Arn,
//				LogDestinationType: pulumi.String("s3"),
//				TrafficType:        pulumi.String("ALL"),
//				VpcId:              pulumi.Any(aws_vpc.Example.Id),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### S3 Logging in Apache Parquet format with per-hour partitions
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
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
//			_, err = ec2.NewFlowLog(ctx, "exampleFlowLog", &ec2.FlowLogArgs{
//				LogDestination:     exampleBucketV2.Arn,
//				LogDestinationType: pulumi.String("s3"),
//				TrafficType:        pulumi.String("ALL"),
//				VpcId:              pulumi.Any(aws_vpc.Example.Id),
//				DestinationOptions: &ec2.FlowLogDestinationOptionsArgs{
//					FileFormat:       pulumi.String("parquet"),
//					PerHourPartition: pulumi.Bool(true),
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
// Flow Logs can be imported using the `id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:ec2/flowLog:FlowLog test_flow_log fl-1a2b3c4d
//
// ```
type FlowLog struct {
	pulumi.CustomResourceState

	// The ARN of the Flow Log.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
	DeliverCrossAccountRole pulumi.StringPtrOutput `pulumi:"deliverCrossAccountRole"`
	// Describes the destination options for a flow log. More details below.
	DestinationOptions FlowLogDestinationOptionsPtrOutput `pulumi:"destinationOptions"`
	// Elastic Network Interface ID to attach to
	EniId pulumi.StringPtrOutput `pulumi:"eniId"`
	// The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
	IamRoleArn pulumi.StringPtrOutput `pulumi:"iamRoleArn"`
	// The ARN of the logging destination. Either `logDestination` or `logGroupName` must be set.
	LogDestination pulumi.StringOutput `pulumi:"logDestination"`
	// The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`, `kinesis-data-firehose`. Default: `cloud-watch-logs`.
	LogDestinationType pulumi.StringPtrOutput `pulumi:"logDestinationType"`
	// The fields to include in the flow log record, in the order in which they should appear.
	LogFormat pulumi.StringOutput `pulumi:"logFormat"`
	// *Deprecated:* Use `logDestination` instead. The name of the CloudWatch log group. Either `logGroupName` or `logDestination` must be set.
	//
	// Deprecated: use 'log_destination' argument instead
	LogGroupName pulumi.StringOutput `pulumi:"logGroupName"`
	// The maximum interval of time
	// during which a flow of packets is captured and aggregated into a flow
	// log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
	// minutes). Default: `600`. When `transitGatewayId` or `transitGatewayAttachmentId` is specified, `maxAggregationInterval` *must* be 60 seconds (1 minute).
	MaxAggregationInterval pulumi.IntPtrOutput `pulumi:"maxAggregationInterval"`
	// Subnet ID to attach to
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
	TrafficType pulumi.StringPtrOutput `pulumi:"trafficType"`
	// Transit Gateway Attachment ID to attach to
	TransitGatewayAttachmentId pulumi.StringPtrOutput `pulumi:"transitGatewayAttachmentId"`
	// Transit Gateway ID to attach to
	TransitGatewayId pulumi.StringPtrOutput `pulumi:"transitGatewayId"`
	// VPC ID to attach to
	VpcId pulumi.StringPtrOutput `pulumi:"vpcId"`
}

// NewFlowLog registers a new resource with the given unique name, arguments, and options.
func NewFlowLog(ctx *pulumi.Context,
	name string, args *FlowLogArgs, opts ...pulumi.ResourceOption) (*FlowLog, error) {
	if args == nil {
		args = &FlowLogArgs{}
	}

	var resource FlowLog
	err := ctx.RegisterResource("aws:ec2/flowLog:FlowLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlowLog gets an existing FlowLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlowLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowLogState, opts ...pulumi.ResourceOption) (*FlowLog, error) {
	var resource FlowLog
	err := ctx.ReadResource("aws:ec2/flowLog:FlowLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlowLog resources.
type flowLogState struct {
	// The ARN of the Flow Log.
	Arn *string `pulumi:"arn"`
	// ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
	DeliverCrossAccountRole *string `pulumi:"deliverCrossAccountRole"`
	// Describes the destination options for a flow log. More details below.
	DestinationOptions *FlowLogDestinationOptions `pulumi:"destinationOptions"`
	// Elastic Network Interface ID to attach to
	EniId *string `pulumi:"eniId"`
	// The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
	IamRoleArn *string `pulumi:"iamRoleArn"`
	// The ARN of the logging destination. Either `logDestination` or `logGroupName` must be set.
	LogDestination *string `pulumi:"logDestination"`
	// The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`, `kinesis-data-firehose`. Default: `cloud-watch-logs`.
	LogDestinationType *string `pulumi:"logDestinationType"`
	// The fields to include in the flow log record, in the order in which they should appear.
	LogFormat *string `pulumi:"logFormat"`
	// *Deprecated:* Use `logDestination` instead. The name of the CloudWatch log group. Either `logGroupName` or `logDestination` must be set.
	//
	// Deprecated: use 'log_destination' argument instead
	LogGroupName *string `pulumi:"logGroupName"`
	// The maximum interval of time
	// during which a flow of packets is captured and aggregated into a flow
	// log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
	// minutes). Default: `600`. When `transitGatewayId` or `transitGatewayAttachmentId` is specified, `maxAggregationInterval` *must* be 60 seconds (1 minute).
	MaxAggregationInterval *int `pulumi:"maxAggregationInterval"`
	// Subnet ID to attach to
	SubnetId *string `pulumi:"subnetId"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
	TrafficType *string `pulumi:"trafficType"`
	// Transit Gateway Attachment ID to attach to
	TransitGatewayAttachmentId *string `pulumi:"transitGatewayAttachmentId"`
	// Transit Gateway ID to attach to
	TransitGatewayId *string `pulumi:"transitGatewayId"`
	// VPC ID to attach to
	VpcId *string `pulumi:"vpcId"`
}

type FlowLogState struct {
	// The ARN of the Flow Log.
	Arn pulumi.StringPtrInput
	// ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
	DeliverCrossAccountRole pulumi.StringPtrInput
	// Describes the destination options for a flow log. More details below.
	DestinationOptions FlowLogDestinationOptionsPtrInput
	// Elastic Network Interface ID to attach to
	EniId pulumi.StringPtrInput
	// The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
	IamRoleArn pulumi.StringPtrInput
	// The ARN of the logging destination. Either `logDestination` or `logGroupName` must be set.
	LogDestination pulumi.StringPtrInput
	// The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`, `kinesis-data-firehose`. Default: `cloud-watch-logs`.
	LogDestinationType pulumi.StringPtrInput
	// The fields to include in the flow log record, in the order in which they should appear.
	LogFormat pulumi.StringPtrInput
	// *Deprecated:* Use `logDestination` instead. The name of the CloudWatch log group. Either `logGroupName` or `logDestination` must be set.
	//
	// Deprecated: use 'log_destination' argument instead
	LogGroupName pulumi.StringPtrInput
	// The maximum interval of time
	// during which a flow of packets is captured and aggregated into a flow
	// log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
	// minutes). Default: `600`. When `transitGatewayId` or `transitGatewayAttachmentId` is specified, `maxAggregationInterval` *must* be 60 seconds (1 minute).
	MaxAggregationInterval pulumi.IntPtrInput
	// Subnet ID to attach to
	SubnetId pulumi.StringPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
	TrafficType pulumi.StringPtrInput
	// Transit Gateway Attachment ID to attach to
	TransitGatewayAttachmentId pulumi.StringPtrInput
	// Transit Gateway ID to attach to
	TransitGatewayId pulumi.StringPtrInput
	// VPC ID to attach to
	VpcId pulumi.StringPtrInput
}

func (FlowLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowLogState)(nil)).Elem()
}

type flowLogArgs struct {
	// ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
	DeliverCrossAccountRole *string `pulumi:"deliverCrossAccountRole"`
	// Describes the destination options for a flow log. More details below.
	DestinationOptions *FlowLogDestinationOptions `pulumi:"destinationOptions"`
	// Elastic Network Interface ID to attach to
	EniId *string `pulumi:"eniId"`
	// The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
	IamRoleArn *string `pulumi:"iamRoleArn"`
	// The ARN of the logging destination. Either `logDestination` or `logGroupName` must be set.
	LogDestination *string `pulumi:"logDestination"`
	// The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`, `kinesis-data-firehose`. Default: `cloud-watch-logs`.
	LogDestinationType *string `pulumi:"logDestinationType"`
	// The fields to include in the flow log record, in the order in which they should appear.
	LogFormat *string `pulumi:"logFormat"`
	// *Deprecated:* Use `logDestination` instead. The name of the CloudWatch log group. Either `logGroupName` or `logDestination` must be set.
	//
	// Deprecated: use 'log_destination' argument instead
	LogGroupName *string `pulumi:"logGroupName"`
	// The maximum interval of time
	// during which a flow of packets is captured and aggregated into a flow
	// log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
	// minutes). Default: `600`. When `transitGatewayId` or `transitGatewayAttachmentId` is specified, `maxAggregationInterval` *must* be 60 seconds (1 minute).
	MaxAggregationInterval *int `pulumi:"maxAggregationInterval"`
	// Subnet ID to attach to
	SubnetId *string `pulumi:"subnetId"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
	TrafficType *string `pulumi:"trafficType"`
	// Transit Gateway Attachment ID to attach to
	TransitGatewayAttachmentId *string `pulumi:"transitGatewayAttachmentId"`
	// Transit Gateway ID to attach to
	TransitGatewayId *string `pulumi:"transitGatewayId"`
	// VPC ID to attach to
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a FlowLog resource.
type FlowLogArgs struct {
	// ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
	DeliverCrossAccountRole pulumi.StringPtrInput
	// Describes the destination options for a flow log. More details below.
	DestinationOptions FlowLogDestinationOptionsPtrInput
	// Elastic Network Interface ID to attach to
	EniId pulumi.StringPtrInput
	// The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
	IamRoleArn pulumi.StringPtrInput
	// The ARN of the logging destination. Either `logDestination` or `logGroupName` must be set.
	LogDestination pulumi.StringPtrInput
	// The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`, `kinesis-data-firehose`. Default: `cloud-watch-logs`.
	LogDestinationType pulumi.StringPtrInput
	// The fields to include in the flow log record, in the order in which they should appear.
	LogFormat pulumi.StringPtrInput
	// *Deprecated:* Use `logDestination` instead. The name of the CloudWatch log group. Either `logGroupName` or `logDestination` must be set.
	//
	// Deprecated: use 'log_destination' argument instead
	LogGroupName pulumi.StringPtrInput
	// The maximum interval of time
	// during which a flow of packets is captured and aggregated into a flow
	// log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
	// minutes). Default: `600`. When `transitGatewayId` or `transitGatewayAttachmentId` is specified, `maxAggregationInterval` *must* be 60 seconds (1 minute).
	MaxAggregationInterval pulumi.IntPtrInput
	// Subnet ID to attach to
	SubnetId pulumi.StringPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
	TrafficType pulumi.StringPtrInput
	// Transit Gateway Attachment ID to attach to
	TransitGatewayAttachmentId pulumi.StringPtrInput
	// Transit Gateway ID to attach to
	TransitGatewayId pulumi.StringPtrInput
	// VPC ID to attach to
	VpcId pulumi.StringPtrInput
}

func (FlowLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowLogArgs)(nil)).Elem()
}

type FlowLogInput interface {
	pulumi.Input

	ToFlowLogOutput() FlowLogOutput
	ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput
}

func (*FlowLog) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowLog)(nil)).Elem()
}

func (i *FlowLog) ToFlowLogOutput() FlowLogOutput {
	return i.ToFlowLogOutputWithContext(context.Background())
}

func (i *FlowLog) ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowLogOutput)
}

// FlowLogArrayInput is an input type that accepts FlowLogArray and FlowLogArrayOutput values.
// You can construct a concrete instance of `FlowLogArrayInput` via:
//
//	FlowLogArray{ FlowLogArgs{...} }
type FlowLogArrayInput interface {
	pulumi.Input

	ToFlowLogArrayOutput() FlowLogArrayOutput
	ToFlowLogArrayOutputWithContext(context.Context) FlowLogArrayOutput
}

type FlowLogArray []FlowLogInput

func (FlowLogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlowLog)(nil)).Elem()
}

func (i FlowLogArray) ToFlowLogArrayOutput() FlowLogArrayOutput {
	return i.ToFlowLogArrayOutputWithContext(context.Background())
}

func (i FlowLogArray) ToFlowLogArrayOutputWithContext(ctx context.Context) FlowLogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowLogArrayOutput)
}

// FlowLogMapInput is an input type that accepts FlowLogMap and FlowLogMapOutput values.
// You can construct a concrete instance of `FlowLogMapInput` via:
//
//	FlowLogMap{ "key": FlowLogArgs{...} }
type FlowLogMapInput interface {
	pulumi.Input

	ToFlowLogMapOutput() FlowLogMapOutput
	ToFlowLogMapOutputWithContext(context.Context) FlowLogMapOutput
}

type FlowLogMap map[string]FlowLogInput

func (FlowLogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlowLog)(nil)).Elem()
}

func (i FlowLogMap) ToFlowLogMapOutput() FlowLogMapOutput {
	return i.ToFlowLogMapOutputWithContext(context.Background())
}

func (i FlowLogMap) ToFlowLogMapOutputWithContext(ctx context.Context) FlowLogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowLogMapOutput)
}

type FlowLogOutput struct{ *pulumi.OutputState }

func (FlowLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowLog)(nil)).Elem()
}

func (o FlowLogOutput) ToFlowLogOutput() FlowLogOutput {
	return o
}

func (o FlowLogOutput) ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput {
	return o
}

// The ARN of the Flow Log.
func (o FlowLogOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
func (o FlowLogOutput) DeliverCrossAccountRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringPtrOutput { return v.DeliverCrossAccountRole }).(pulumi.StringPtrOutput)
}

// Describes the destination options for a flow log. More details below.
func (o FlowLogOutput) DestinationOptions() FlowLogDestinationOptionsPtrOutput {
	return o.ApplyT(func(v *FlowLog) FlowLogDestinationOptionsPtrOutput { return v.DestinationOptions }).(FlowLogDestinationOptionsPtrOutput)
}

// Elastic Network Interface ID to attach to
func (o FlowLogOutput) EniId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringPtrOutput { return v.EniId }).(pulumi.StringPtrOutput)
}

// The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
func (o FlowLogOutput) IamRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringPtrOutput { return v.IamRoleArn }).(pulumi.StringPtrOutput)
}

// The ARN of the logging destination. Either `logDestination` or `logGroupName` must be set.
func (o FlowLogOutput) LogDestination() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringOutput { return v.LogDestination }).(pulumi.StringOutput)
}

// The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`, `kinesis-data-firehose`. Default: `cloud-watch-logs`.
func (o FlowLogOutput) LogDestinationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringPtrOutput { return v.LogDestinationType }).(pulumi.StringPtrOutput)
}

// The fields to include in the flow log record, in the order in which they should appear.
func (o FlowLogOutput) LogFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringOutput { return v.LogFormat }).(pulumi.StringOutput)
}

// *Deprecated:* Use `logDestination` instead. The name of the CloudWatch log group. Either `logGroupName` or `logDestination` must be set.
//
// Deprecated: use 'log_destination' argument instead
func (o FlowLogOutput) LogGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringOutput { return v.LogGroupName }).(pulumi.StringOutput)
}

// The maximum interval of time
// during which a flow of packets is captured and aggregated into a flow
// log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
// minutes). Default: `600`. When `transitGatewayId` or `transitGatewayAttachmentId` is specified, `maxAggregationInterval` *must* be 60 seconds (1 minute).
func (o FlowLogOutput) MaxAggregationInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.IntPtrOutput { return v.MaxAggregationInterval }).(pulumi.IntPtrOutput)
}

// Subnet ID to attach to
func (o FlowLogOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o FlowLogOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o FlowLogOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
func (o FlowLogOutput) TrafficType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringPtrOutput { return v.TrafficType }).(pulumi.StringPtrOutput)
}

// Transit Gateway Attachment ID to attach to
func (o FlowLogOutput) TransitGatewayAttachmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringPtrOutput { return v.TransitGatewayAttachmentId }).(pulumi.StringPtrOutput)
}

// Transit Gateway ID to attach to
func (o FlowLogOutput) TransitGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringPtrOutput { return v.TransitGatewayId }).(pulumi.StringPtrOutput)
}

// VPC ID to attach to
func (o FlowLogOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringPtrOutput { return v.VpcId }).(pulumi.StringPtrOutput)
}

type FlowLogArrayOutput struct{ *pulumi.OutputState }

func (FlowLogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlowLog)(nil)).Elem()
}

func (o FlowLogArrayOutput) ToFlowLogArrayOutput() FlowLogArrayOutput {
	return o
}

func (o FlowLogArrayOutput) ToFlowLogArrayOutputWithContext(ctx context.Context) FlowLogArrayOutput {
	return o
}

func (o FlowLogArrayOutput) Index(i pulumi.IntInput) FlowLogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FlowLog {
		return vs[0].([]*FlowLog)[vs[1].(int)]
	}).(FlowLogOutput)
}

type FlowLogMapOutput struct{ *pulumi.OutputState }

func (FlowLogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlowLog)(nil)).Elem()
}

func (o FlowLogMapOutput) ToFlowLogMapOutput() FlowLogMapOutput {
	return o
}

func (o FlowLogMapOutput) ToFlowLogMapOutputWithContext(ctx context.Context) FlowLogMapOutput {
	return o
}

func (o FlowLogMapOutput) MapIndex(k pulumi.StringInput) FlowLogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FlowLog {
		return vs[0].(map[string]*FlowLog)[vs[1].(string)]
	}).(FlowLogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlowLogInput)(nil)).Elem(), &FlowLog{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowLogArrayInput)(nil)).Elem(), FlowLogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlowLogMapInput)(nil)).Elem(), FlowLogMap{})
	pulumi.RegisterOutputType(FlowLogOutput{})
	pulumi.RegisterOutputType(FlowLogArrayOutput{})
	pulumi.RegisterOutputType(FlowLogMapOutput{})
}
