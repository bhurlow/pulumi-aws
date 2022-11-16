// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a CloudWatch Metric Stream resource.
 *
 * ## Example Usage
 * ### Filters
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-trustpolicy.html
 * const metricStreamToFirehoseRole = new aws.iam.Role("metricStreamToFirehoseRole", {assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "sts:AssumeRole",
 *       "Principal": {
 *         "Service": "streams.metrics.cloudwatch.amazonaws.com"
 *       },
 *       "Effect": "Allow",
 *       "Sid": ""
 *     }
 *   ]
 * }
 * `});
 * const bucket = new aws.s3.BucketV2("bucket", {});
 * const firehoseToS3Role = new aws.iam.Role("firehoseToS3Role", {assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "sts:AssumeRole",
 *       "Principal": {
 *         "Service": "firehose.amazonaws.com"
 *       },
 *       "Effect": "Allow",
 *       "Sid": ""
 *     }
 *   ]
 * }
 * `});
 * const s3Stream = new aws.kinesis.FirehoseDeliveryStream("s3Stream", {
 *     destination: "s3",
 *     s3Configuration: {
 *         roleArn: firehoseToS3Role.arn,
 *         bucketArn: bucket.arn,
 *     },
 * });
 * const main = new aws.cloudwatch.MetricStream("main", {
 *     roleArn: metricStreamToFirehoseRole.arn,
 *     firehoseArn: s3Stream.arn,
 *     outputFormat: "json",
 *     includeFilters: [
 *         {
 *             namespace: "AWS/EC2",
 *         },
 *         {
 *             namespace: "AWS/EBS",
 *         },
 *     ],
 * });
 * // https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-trustpolicy.html
 * const metricStreamToFirehoseRolePolicy = new aws.iam.RolePolicy("metricStreamToFirehoseRolePolicy", {
 *     role: metricStreamToFirehoseRole.id,
 *     policy: pulumi.interpolate`{
 *     "Version": "2012-10-17",
 *     "Statement": [
 *         {
 *             "Effect": "Allow",
 *             "Action": [
 *                 "firehose:PutRecord",
 *                 "firehose:PutRecordBatch"
 *             ],
 *             "Resource": "${s3Stream.arn}"
 *         }
 *     ]
 * }
 * `,
 * });
 * const bucketAcl = new aws.s3.BucketAclV2("bucketAcl", {
 *     bucket: bucket.id,
 *     acl: "private",
 * });
 * const firehoseToS3RolePolicy = new aws.iam.RolePolicy("firehoseToS3RolePolicy", {
 *     role: firehoseToS3Role.id,
 *     policy: pulumi.interpolate`{
 *     "Version": "2012-10-17",
 *     "Statement": [
 *         {
 *             "Effect": "Allow",
 *             "Action": [
 *                 "s3:AbortMultipartUpload",
 *                 "s3:GetBucketLocation",
 *                 "s3:GetObject",
 *                 "s3:ListBucket",
 *                 "s3:ListBucketMultipartUploads",
 *                 "s3:PutObject"
 *             ],
 *             "Resource": [
 *                 "${bucket.arn}",
 *                 "${bucket.arn}/*"
 *             ]
 *         }
 *     ]
 * }
 * `,
 * });
 * ```
 * ### Additional Statistics
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const main = new aws.cloudwatch.MetricStream("main", {
 *     roleArn: aws_iam_role.metric_stream_to_firehose.arn,
 *     firehoseArn: aws_kinesis_firehose_delivery_stream.s3_stream.arn,
 *     outputFormat: "json",
 *     statisticsConfigurations: [
 *         {
 *             additionalStatistics: [
 *                 "p1",
 *                 "tm99",
 *             ],
 *             includeMetrics: [{
 *                 metricName: "CPUUtilization",
 *                 namespace: "AWS/EC2",
 *             }],
 *         },
 *         {
 *             additionalStatistics: ["TS(50.5:)"],
 *             includeMetrics: [{
 *                 metricName: "CPUUtilization",
 *                 namespace: "AWS/EC2",
 *             }],
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * CloudWatch metric streams can be imported using the `name`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:cloudwatch/metricStream:MetricStream sample sample-stream-name
 * ```
 */
export class MetricStream extends pulumi.CustomResource {
    /**
     * Get an existing MetricStream resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MetricStreamState, opts?: pulumi.CustomResourceOptions): MetricStream {
        return new MetricStream(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudwatch/metricStream:MetricStream';

    /**
     * Returns true if the given object is an instance of MetricStream.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MetricStream {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MetricStream.__pulumiType;
    }

    /**
     * ARN of the metric stream.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the metric stream was created.
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * List of exclusive metric filters. If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces that you specify here. Conflicts with `includeFilter`.
     */
    public readonly excludeFilters!: pulumi.Output<outputs.cloudwatch.MetricStreamExcludeFilter[] | undefined>;
    /**
     * ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
     */
    public readonly firehoseArn!: pulumi.Output<string>;
    /**
     * List of inclusive metric filters. If you specify this parameter, the stream sends only the metrics from the metric namespaces that you specify here. Conflicts with `excludeFilter`.
     */
    public readonly includeFilters!: pulumi.Output<outputs.cloudwatch.MetricStreamIncludeFilter[] | undefined>;
    /**
     * Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the metric stream was last updated.
     */
    public /*out*/ readonly lastUpdateDate!: pulumi.Output<string>;
    /**
     * Friendly name of the metric stream. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string>;
    /**
     * Output format for the stream. Possible values are `json` and `opentelemetry0.7`. For more information about output formats, see [Metric streams output formats](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html).
     */
    public readonly outputFormat!: pulumi.Output<string>;
    /**
     * ARN of the IAM role that this metric stream will use to access Amazon Kinesis Firehose resources. For more information about role permissions, see [Trust between CloudWatch and Kinesis Data Firehose](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-trustpolicy.html).
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * State of the metric stream. Possible values are `running` and `stopped`.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * For each entry in this array, you specify one or more metrics and the list of additional statistics to stream for those metrics. The additional statistics that you can stream depend on the stream's `outputFormat`. If the OutputFormat is `json`, you can stream any additional statistic that is supported by CloudWatch, listed in [CloudWatch statistics definitions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.html). If the OutputFormat is `opentelemetry0.7`, you can stream percentile statistics (p99 etc.). See details below.
     */
    public readonly statisticsConfigurations!: pulumi.Output<outputs.cloudwatch.MetricStreamStatisticsConfiguration[] | undefined>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a MetricStream resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MetricStreamArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MetricStreamArgs | MetricStreamState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MetricStreamState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["creationDate"] = state ? state.creationDate : undefined;
            resourceInputs["excludeFilters"] = state ? state.excludeFilters : undefined;
            resourceInputs["firehoseArn"] = state ? state.firehoseArn : undefined;
            resourceInputs["includeFilters"] = state ? state.includeFilters : undefined;
            resourceInputs["lastUpdateDate"] = state ? state.lastUpdateDate : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namePrefix"] = state ? state.namePrefix : undefined;
            resourceInputs["outputFormat"] = state ? state.outputFormat : undefined;
            resourceInputs["roleArn"] = state ? state.roleArn : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["statisticsConfigurations"] = state ? state.statisticsConfigurations : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as MetricStreamArgs | undefined;
            if ((!args || args.firehoseArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'firehoseArn'");
            }
            if ((!args || args.outputFormat === undefined) && !opts.urn) {
                throw new Error("Missing required property 'outputFormat'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            resourceInputs["excludeFilters"] = args ? args.excludeFilters : undefined;
            resourceInputs["firehoseArn"] = args ? args.firehoseArn : undefined;
            resourceInputs["includeFilters"] = args ? args.includeFilters : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["outputFormat"] = args ? args.outputFormat : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["statisticsConfigurations"] = args ? args.statisticsConfigurations : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["lastUpdateDate"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MetricStream.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MetricStream resources.
 */
export interface MetricStreamState {
    /**
     * ARN of the metric stream.
     */
    arn?: pulumi.Input<string>;
    /**
     * Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the metric stream was created.
     */
    creationDate?: pulumi.Input<string>;
    /**
     * List of exclusive metric filters. If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces that you specify here. Conflicts with `includeFilter`.
     */
    excludeFilters?: pulumi.Input<pulumi.Input<inputs.cloudwatch.MetricStreamExcludeFilter>[]>;
    /**
     * ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
     */
    firehoseArn?: pulumi.Input<string>;
    /**
     * List of inclusive metric filters. If you specify this parameter, the stream sends only the metrics from the metric namespaces that you specify here. Conflicts with `excludeFilter`.
     */
    includeFilters?: pulumi.Input<pulumi.Input<inputs.cloudwatch.MetricStreamIncludeFilter>[]>;
    /**
     * Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the metric stream was last updated.
     */
    lastUpdateDate?: pulumi.Input<string>;
    /**
     * Friendly name of the metric stream. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * Output format for the stream. Possible values are `json` and `opentelemetry0.7`. For more information about output formats, see [Metric streams output formats](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html).
     */
    outputFormat?: pulumi.Input<string>;
    /**
     * ARN of the IAM role that this metric stream will use to access Amazon Kinesis Firehose resources. For more information about role permissions, see [Trust between CloudWatch and Kinesis Data Firehose](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-trustpolicy.html).
     */
    roleArn?: pulumi.Input<string>;
    /**
     * State of the metric stream. Possible values are `running` and `stopped`.
     */
    state?: pulumi.Input<string>;
    /**
     * For each entry in this array, you specify one or more metrics and the list of additional statistics to stream for those metrics. The additional statistics that you can stream depend on the stream's `outputFormat`. If the OutputFormat is `json`, you can stream any additional statistic that is supported by CloudWatch, listed in [CloudWatch statistics definitions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.html). If the OutputFormat is `opentelemetry0.7`, you can stream percentile statistics (p99 etc.). See details below.
     */
    statisticsConfigurations?: pulumi.Input<pulumi.Input<inputs.cloudwatch.MetricStreamStatisticsConfiguration>[]>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a MetricStream resource.
 */
export interface MetricStreamArgs {
    /**
     * List of exclusive metric filters. If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces that you specify here. Conflicts with `includeFilter`.
     */
    excludeFilters?: pulumi.Input<pulumi.Input<inputs.cloudwatch.MetricStreamExcludeFilter>[]>;
    /**
     * ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
     */
    firehoseArn: pulumi.Input<string>;
    /**
     * List of inclusive metric filters. If you specify this parameter, the stream sends only the metrics from the metric namespaces that you specify here. Conflicts with `excludeFilter`.
     */
    includeFilters?: pulumi.Input<pulumi.Input<inputs.cloudwatch.MetricStreamIncludeFilter>[]>;
    /**
     * Friendly name of the metric stream. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * Output format for the stream. Possible values are `json` and `opentelemetry0.7`. For more information about output formats, see [Metric streams output formats](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html).
     */
    outputFormat: pulumi.Input<string>;
    /**
     * ARN of the IAM role that this metric stream will use to access Amazon Kinesis Firehose resources. For more information about role permissions, see [Trust between CloudWatch and Kinesis Data Firehose](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-trustpolicy.html).
     */
    roleArn: pulumi.Input<string>;
    /**
     * For each entry in this array, you specify one or more metrics and the list of additional statistics to stream for those metrics. The additional statistics that you can stream depend on the stream's `outputFormat`. If the OutputFormat is `json`, you can stream any additional statistic that is supported by CloudWatch, listed in [CloudWatch statistics definitions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.html). If the OutputFormat is `opentelemetry0.7`, you can stream percentile statistics (p99 etc.). See details below.
     */
    statisticsConfigurations?: pulumi.Input<pulumi.Input<inputs.cloudwatch.MetricStreamStatisticsConfiguration>[]>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
