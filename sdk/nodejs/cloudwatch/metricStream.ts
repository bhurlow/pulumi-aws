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
 * const streamsAssumeRole = aws.iam.getPolicyDocument({
 *     statements: [{
 *         effect: "Allow",
 *         principals: [{
 *             type: "Service",
 *             identifiers: ["streams.metrics.cloudwatch.amazonaws.com"],
 *         }],
 *         actions: ["sts:AssumeRole"],
 *     }],
 * });
 * const metricStreamToFirehoseRole = new aws.iam.Role("metricStreamToFirehoseRole", {assumeRolePolicy: streamsAssumeRole.then(streamsAssumeRole => streamsAssumeRole.json)});
 * const bucket = new aws.s3.BucketV2("bucket", {});
 * const firehoseAssumeRole = aws.iam.getPolicyDocument({
 *     statements: [{
 *         effect: "Allow",
 *         principals: [{
 *             type: "Service",
 *             identifiers: ["firehose.amazonaws.com"],
 *         }],
 *         actions: ["sts:AssumeRole"],
 *     }],
 * });
 * const firehoseToS3Role = new aws.iam.Role("firehoseToS3Role", {assumeRolePolicy: firehoseAssumeRole.then(firehoseAssumeRole => firehoseAssumeRole.json)});
 * const s3Stream = new aws.kinesis.FirehoseDeliveryStream("s3Stream", {
 *     destination: "extended_s3",
 *     extendedS3Configuration: {
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
 *             metricNames: [
 *                 "CPUUtilization",
 *                 "NetworkOut",
 *             ],
 *         },
 *         {
 *             namespace: "AWS/EBS",
 *             metricNames: [],
 *         },
 *     ],
 * });
 * const metricStreamToFirehosePolicyDocument = aws.iam.getPolicyDocumentOutput({
 *     statements: [{
 *         effect: "Allow",
 *         actions: [
 *             "firehose:PutRecord",
 *             "firehose:PutRecordBatch",
 *         ],
 *         resources: [s3Stream.arn],
 *     }],
 * });
 * const metricStreamToFirehoseRolePolicy = new aws.iam.RolePolicy("metricStreamToFirehoseRolePolicy", {
 *     role: metricStreamToFirehoseRole.id,
 *     policy: metricStreamToFirehosePolicyDocument.apply(metricStreamToFirehosePolicyDocument => metricStreamToFirehosePolicyDocument.json),
 * });
 * const bucketAcl = new aws.s3.BucketAclV2("bucketAcl", {
 *     bucket: bucket.id,
 *     acl: "private",
 * });
 * const firehoseToS3PolicyDocument = aws.iam.getPolicyDocumentOutput({
 *     statements: [{
 *         effect: "Allow",
 *         actions: [
 *             "s3:AbortMultipartUpload",
 *             "s3:GetBucketLocation",
 *             "s3:GetObject",
 *             "s3:ListBucket",
 *             "s3:ListBucketMultipartUploads",
 *             "s3:PutObject",
 *         ],
 *         resources: [
 *             bucket.arn,
 *             pulumi.interpolate`${bucket.arn}/*`,
 *         ],
 *     }],
 * });
 * const firehoseToS3RolePolicy = new aws.iam.RolePolicy("firehoseToS3RolePolicy", {
 *     role: firehoseToS3Role.id,
 *     policy: firehoseToS3PolicyDocument.apply(firehoseToS3PolicyDocument => firehoseToS3PolicyDocument.json),
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
 * Using `pulumi import`, import CloudWatch metric streams using the `name`. For example:
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
     * List of exclusive metric filters. If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces and the conditional metric names that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is excluded. Conflicts with `includeFilter`.
     */
    public readonly excludeFilters!: pulumi.Output<outputs.cloudwatch.MetricStreamExcludeFilter[] | undefined>;
    /**
     * ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
     */
    public readonly firehoseArn!: pulumi.Output<string>;
    /**
     * List of inclusive metric filters. If you specify this parameter, the stream sends only the conditional metric names from the metric namespaces that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is included. Conflicts with `excludeFilter`.
     */
    public readonly includeFilters!: pulumi.Output<outputs.cloudwatch.MetricStreamIncludeFilter[] | undefined>;
    /**
     * If you are creating a metric stream in a monitoring account, specify true to include metrics from source accounts that are linked to this monitoring account, in the metric stream. The default is false. For more information about linking accounts, see [CloudWatch cross-account observability](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html).
     */
    public readonly includeLinkedAccountsMetrics!: pulumi.Output<boolean | undefined>;
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
     * Output format for the stream. Possible values are `json`, `opentelemetry0.7`, and `opentelemetry1.0`. For more information about output formats, see [Metric streams output formats](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html).
     *
     * The following arguments are optional:
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
     * For each entry in this array, you specify one or more metrics and the list of additional statistics to stream for those metrics. The additional statistics that you can stream depend on the stream's `outputFormat`. If the OutputFormat is `json`, you can stream any additional statistic that is supported by CloudWatch, listed in [CloudWatch statistics definitions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.html). If the OutputFormat is `opentelemetry0.7` or `opentelemetry1.0`, you can stream percentile statistics (p99 etc.). See details below.
     */
    public readonly statisticsConfigurations!: pulumi.Output<outputs.cloudwatch.MetricStreamStatisticsConfiguration[] | undefined>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
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
            resourceInputs["includeLinkedAccountsMetrics"] = state ? state.includeLinkedAccountsMetrics : undefined;
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
            resourceInputs["includeLinkedAccountsMetrics"] = args ? args.includeLinkedAccountsMetrics : undefined;
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
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
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
     * List of exclusive metric filters. If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces and the conditional metric names that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is excluded. Conflicts with `includeFilter`.
     */
    excludeFilters?: pulumi.Input<pulumi.Input<inputs.cloudwatch.MetricStreamExcludeFilter>[]>;
    /**
     * ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
     */
    firehoseArn?: pulumi.Input<string>;
    /**
     * List of inclusive metric filters. If you specify this parameter, the stream sends only the conditional metric names from the metric namespaces that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is included. Conflicts with `excludeFilter`.
     */
    includeFilters?: pulumi.Input<pulumi.Input<inputs.cloudwatch.MetricStreamIncludeFilter>[]>;
    /**
     * If you are creating a metric stream in a monitoring account, specify true to include metrics from source accounts that are linked to this monitoring account, in the metric stream. The default is false. For more information about linking accounts, see [CloudWatch cross-account observability](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html).
     */
    includeLinkedAccountsMetrics?: pulumi.Input<boolean>;
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
     * Output format for the stream. Possible values are `json`, `opentelemetry0.7`, and `opentelemetry1.0`. For more information about output formats, see [Metric streams output formats](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html).
     *
     * The following arguments are optional:
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
     * For each entry in this array, you specify one or more metrics and the list of additional statistics to stream for those metrics. The additional statistics that you can stream depend on the stream's `outputFormat`. If the OutputFormat is `json`, you can stream any additional statistic that is supported by CloudWatch, listed in [CloudWatch statistics definitions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.html). If the OutputFormat is `opentelemetry0.7` or `opentelemetry1.0`, you can stream percentile statistics (p99 etc.). See details below.
     */
    statisticsConfigurations?: pulumi.Input<pulumi.Input<inputs.cloudwatch.MetricStreamStatisticsConfiguration>[]>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a MetricStream resource.
 */
export interface MetricStreamArgs {
    /**
     * List of exclusive metric filters. If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces and the conditional metric names that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is excluded. Conflicts with `includeFilter`.
     */
    excludeFilters?: pulumi.Input<pulumi.Input<inputs.cloudwatch.MetricStreamExcludeFilter>[]>;
    /**
     * ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
     */
    firehoseArn: pulumi.Input<string>;
    /**
     * List of inclusive metric filters. If you specify this parameter, the stream sends only the conditional metric names from the metric namespaces that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is included. Conflicts with `excludeFilter`.
     */
    includeFilters?: pulumi.Input<pulumi.Input<inputs.cloudwatch.MetricStreamIncludeFilter>[]>;
    /**
     * If you are creating a metric stream in a monitoring account, specify true to include metrics from source accounts that are linked to this monitoring account, in the metric stream. The default is false. For more information about linking accounts, see [CloudWatch cross-account observability](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html).
     */
    includeLinkedAccountsMetrics?: pulumi.Input<boolean>;
    /**
     * Friendly name of the metric stream. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * Output format for the stream. Possible values are `json`, `opentelemetry0.7`, and `opentelemetry1.0`. For more information about output formats, see [Metric streams output formats](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html).
     *
     * The following arguments are optional:
     */
    outputFormat: pulumi.Input<string>;
    /**
     * ARN of the IAM role that this metric stream will use to access Amazon Kinesis Firehose resources. For more information about role permissions, see [Trust between CloudWatch and Kinesis Data Firehose](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-trustpolicy.html).
     */
    roleArn: pulumi.Input<string>;
    /**
     * For each entry in this array, you specify one or more metrics and the list of additional statistics to stream for those metrics. The additional statistics that you can stream depend on the stream's `outputFormat`. If the OutputFormat is `json`, you can stream any additional statistic that is supported by CloudWatch, listed in [CloudWatch statistics definitions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.html). If the OutputFormat is `opentelemetry0.7` or `opentelemetry1.0`, you can stream percentile statistics (p99 etc.). See details below.
     */
    statisticsConfigurations?: pulumi.Input<pulumi.Input<inputs.cloudwatch.MetricStreamStatisticsConfiguration>[]>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
