// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a VPC/Subnet/ENI/Transit Gateway/Transit Gateway Attachment Flow Log to capture IP traffic for a specific network
 * interface, subnet, or VPC. Logs are sent to a CloudWatch Log Group, a S3 Bucket, or Amazon Kinesis Data Firehose
 *
 * ## Example Usage
 * ### CloudWatch Logging
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {});
 * const assumeRole = aws.iam.getPolicyDocument({
 *     statements: [{
 *         effect: "Allow",
 *         principals: [{
 *             type: "Service",
 *             identifiers: ["vpc-flow-logs.amazonaws.com"],
 *         }],
 *         actions: ["sts:AssumeRole"],
 *     }],
 * });
 * const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
 * const exampleFlowLog = new aws.ec2.FlowLog("exampleFlowLog", {
 *     iamRoleArn: exampleRole.arn,
 *     logDestination: exampleLogGroup.arn,
 *     trafficType: "ALL",
 *     vpcId: aws_vpc.example.id,
 * });
 * const examplePolicyDocument = aws.iam.getPolicyDocument({
 *     statements: [{
 *         effect: "Allow",
 *         actions: [
 *             "logs:CreateLogGroup",
 *             "logs:CreateLogStream",
 *             "logs:PutLogEvents",
 *             "logs:DescribeLogGroups",
 *             "logs:DescribeLogStreams",
 *         ],
 *         resources: ["*"],
 *     }],
 * });
 * const exampleRolePolicy = new aws.iam.RolePolicy("exampleRolePolicy", {
 *     role: exampleRole.id,
 *     policy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json),
 * });
 * ```
 * ### S3 Logging
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
 * const exampleFlowLog = new aws.ec2.FlowLog("exampleFlowLog", {
 *     logDestination: exampleBucketV2.arn,
 *     logDestinationType: "s3",
 *     trafficType: "ALL",
 *     vpcId: aws_vpc.example.id,
 * });
 * ```
 * ### S3 Logging in Apache Parquet format with per-hour partitions
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
 * const exampleFlowLog = new aws.ec2.FlowLog("exampleFlowLog", {
 *     logDestination: exampleBucketV2.arn,
 *     logDestinationType: "s3",
 *     trafficType: "ALL",
 *     vpcId: aws_vpc.example.id,
 *     destinationOptions: {
 *         fileFormat: "parquet",
 *         perHourPartition: true,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_flow_log.test_flow_log
 *
 *  id = "fl-1a2b3c4d" } Using `pulumi import`, import Flow Logs using the `id`. For exampleconsole % pulumi import aws_flow_log.test_flow_log fl-1a2b3c4d
 */
export class FlowLog extends pulumi.CustomResource {
    /**
     * Get an existing FlowLog resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FlowLogState, opts?: pulumi.CustomResourceOptions): FlowLog {
        return new FlowLog(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/flowLog:FlowLog';

    /**
     * Returns true if the given object is an instance of FlowLog.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FlowLog {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FlowLog.__pulumiType;
    }

    /**
     * The ARN of the Flow Log.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
     */
    public readonly deliverCrossAccountRole!: pulumi.Output<string | undefined>;
    /**
     * Describes the destination options for a flow log. More details below.
     */
    public readonly destinationOptions!: pulumi.Output<outputs.ec2.FlowLogDestinationOptions | undefined>;
    /**
     * Elastic Network Interface ID to attach to
     */
    public readonly eniId!: pulumi.Output<string | undefined>;
    /**
     * The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
     */
    public readonly iamRoleArn!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the logging destination. Either `logDestination` or `logGroupName` must be set.
     */
    public readonly logDestination!: pulumi.Output<string>;
    /**
     * The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`, `kinesis-data-firehose`. Default: `cloud-watch-logs`.
     */
    public readonly logDestinationType!: pulumi.Output<string | undefined>;
    /**
     * The fields to include in the flow log record, in the order in which they should appear.
     */
    public readonly logFormat!: pulumi.Output<string>;
    /**
     * **Deprecated:** Use `logDestination` instead. The name of the CloudWatch log group. Either `logGroupName` or `logDestination` must be set.
     *
     * @deprecated use 'log_destination' argument instead
     */
    public readonly logGroupName!: pulumi.Output<string>;
    /**
     * The maximum interval of time
     * during which a flow of packets is captured and aggregated into a flow
     * log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
     * minutes). Default: `600`. When `transitGatewayId` or `transitGatewayAttachmentId` is specified, `maxAggregationInterval` *must* be 60 seconds (1 minute).
     */
    public readonly maxAggregationInterval!: pulumi.Output<number | undefined>;
    /**
     * Subnet ID to attach to
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
     */
    public readonly trafficType!: pulumi.Output<string | undefined>;
    /**
     * Transit Gateway Attachment ID to attach to
     */
    public readonly transitGatewayAttachmentId!: pulumi.Output<string | undefined>;
    /**
     * Transit Gateway ID to attach to
     */
    public readonly transitGatewayId!: pulumi.Output<string | undefined>;
    /**
     * VPC ID to attach to
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;

    /**
     * Create a FlowLog resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FlowLogArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FlowLogArgs | FlowLogState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FlowLogState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["deliverCrossAccountRole"] = state ? state.deliverCrossAccountRole : undefined;
            resourceInputs["destinationOptions"] = state ? state.destinationOptions : undefined;
            resourceInputs["eniId"] = state ? state.eniId : undefined;
            resourceInputs["iamRoleArn"] = state ? state.iamRoleArn : undefined;
            resourceInputs["logDestination"] = state ? state.logDestination : undefined;
            resourceInputs["logDestinationType"] = state ? state.logDestinationType : undefined;
            resourceInputs["logFormat"] = state ? state.logFormat : undefined;
            resourceInputs["logGroupName"] = state ? state.logGroupName : undefined;
            resourceInputs["maxAggregationInterval"] = state ? state.maxAggregationInterval : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["trafficType"] = state ? state.trafficType : undefined;
            resourceInputs["transitGatewayAttachmentId"] = state ? state.transitGatewayAttachmentId : undefined;
            resourceInputs["transitGatewayId"] = state ? state.transitGatewayId : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as FlowLogArgs | undefined;
            resourceInputs["deliverCrossAccountRole"] = args ? args.deliverCrossAccountRole : undefined;
            resourceInputs["destinationOptions"] = args ? args.destinationOptions : undefined;
            resourceInputs["eniId"] = args ? args.eniId : undefined;
            resourceInputs["iamRoleArn"] = args ? args.iamRoleArn : undefined;
            resourceInputs["logDestination"] = args ? args.logDestination : undefined;
            resourceInputs["logDestinationType"] = args ? args.logDestinationType : undefined;
            resourceInputs["logFormat"] = args ? args.logFormat : undefined;
            resourceInputs["logGroupName"] = args ? args.logGroupName : undefined;
            resourceInputs["maxAggregationInterval"] = args ? args.maxAggregationInterval : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["trafficType"] = args ? args.trafficType : undefined;
            resourceInputs["transitGatewayAttachmentId"] = args ? args.transitGatewayAttachmentId : undefined;
            resourceInputs["transitGatewayId"] = args ? args.transitGatewayId : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FlowLog.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FlowLog resources.
 */
export interface FlowLogState {
    /**
     * The ARN of the Flow Log.
     */
    arn?: pulumi.Input<string>;
    /**
     * ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
     */
    deliverCrossAccountRole?: pulumi.Input<string>;
    /**
     * Describes the destination options for a flow log. More details below.
     */
    destinationOptions?: pulumi.Input<inputs.ec2.FlowLogDestinationOptions>;
    /**
     * Elastic Network Interface ID to attach to
     */
    eniId?: pulumi.Input<string>;
    /**
     * The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
     */
    iamRoleArn?: pulumi.Input<string>;
    /**
     * The ARN of the logging destination. Either `logDestination` or `logGroupName` must be set.
     */
    logDestination?: pulumi.Input<string>;
    /**
     * The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`, `kinesis-data-firehose`. Default: `cloud-watch-logs`.
     */
    logDestinationType?: pulumi.Input<string>;
    /**
     * The fields to include in the flow log record, in the order in which they should appear.
     */
    logFormat?: pulumi.Input<string>;
    /**
     * **Deprecated:** Use `logDestination` instead. The name of the CloudWatch log group. Either `logGroupName` or `logDestination` must be set.
     *
     * @deprecated use 'log_destination' argument instead
     */
    logGroupName?: pulumi.Input<string>;
    /**
     * The maximum interval of time
     * during which a flow of packets is captured and aggregated into a flow
     * log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
     * minutes). Default: `600`. When `transitGatewayId` or `transitGatewayAttachmentId` is specified, `maxAggregationInterval` *must* be 60 seconds (1 minute).
     */
    maxAggregationInterval?: pulumi.Input<number>;
    /**
     * Subnet ID to attach to
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
     */
    trafficType?: pulumi.Input<string>;
    /**
     * Transit Gateway Attachment ID to attach to
     */
    transitGatewayAttachmentId?: pulumi.Input<string>;
    /**
     * Transit Gateway ID to attach to
     */
    transitGatewayId?: pulumi.Input<string>;
    /**
     * VPC ID to attach to
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FlowLog resource.
 */
export interface FlowLogArgs {
    /**
     * ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
     */
    deliverCrossAccountRole?: pulumi.Input<string>;
    /**
     * Describes the destination options for a flow log. More details below.
     */
    destinationOptions?: pulumi.Input<inputs.ec2.FlowLogDestinationOptions>;
    /**
     * Elastic Network Interface ID to attach to
     */
    eniId?: pulumi.Input<string>;
    /**
     * The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
     */
    iamRoleArn?: pulumi.Input<string>;
    /**
     * The ARN of the logging destination. Either `logDestination` or `logGroupName` must be set.
     */
    logDestination?: pulumi.Input<string>;
    /**
     * The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`, `kinesis-data-firehose`. Default: `cloud-watch-logs`.
     */
    logDestinationType?: pulumi.Input<string>;
    /**
     * The fields to include in the flow log record, in the order in which they should appear.
     */
    logFormat?: pulumi.Input<string>;
    /**
     * **Deprecated:** Use `logDestination` instead. The name of the CloudWatch log group. Either `logGroupName` or `logDestination` must be set.
     *
     * @deprecated use 'log_destination' argument instead
     */
    logGroupName?: pulumi.Input<string>;
    /**
     * The maximum interval of time
     * during which a flow of packets is captured and aggregated into a flow
     * log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
     * minutes). Default: `600`. When `transitGatewayId` or `transitGatewayAttachmentId` is specified, `maxAggregationInterval` *must* be 60 seconds (1 minute).
     */
    maxAggregationInterval?: pulumi.Input<number>;
    /**
     * Subnet ID to attach to
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
     */
    trafficType?: pulumi.Input<string>;
    /**
     * Transit Gateway Attachment ID to attach to
     */
    transitGatewayAttachmentId?: pulumi.Input<string>;
    /**
     * Transit Gateway ID to attach to
     */
    transitGatewayId?: pulumi.Input<string>;
    /**
     * VPC ID to attach to
     */
    vpcId?: pulumi.Input<string>;
}
