// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a Kinesis Stream resource. Amazon Kinesis is a managed service that
 * scales elastically for real-time processing of streaming big data.
 *
 * For more details, see the [Amazon Kinesis Documentation](https://aws.amazon.com/documentation/kinesis/).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testStream = new aws.kinesis.Stream("testStream", {
 *     retentionPeriod: 48,
 *     shardCount: 1,
 *     shardLevelMetrics: [
 *         "IncomingBytes",
 *         "OutgoingBytes",
 *     ],
 *     streamModeDetails: {
 *         streamMode: "PROVISIONED",
 *     },
 *     tags: {
 *         Environment: "test",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_kinesis_stream.test_stream
 *
 *  id = "TODO-kinesis-test" } Using `pulumi import`, import Kinesis Streams using the `name`. For exampleconsole % pulumi import aws_kinesis_stream.test_stream TODO-kinesis-test [1]https://aws.amazon.com/documentation/kinesis/ [2]https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html [3]https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html
 */
export class Stream extends pulumi.CustomResource {
    /**
     * Get an existing Stream resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StreamState, opts?: pulumi.CustomResourceOptions): Stream {
        return new Stream(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:kinesis/stream:Stream';

    /**
     * Returns true if the given object is an instance of Stream.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Stream {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Stream.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
     */
    public readonly arn!: pulumi.Output<string>;
    /**
     * The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
     */
    public readonly encryptionType!: pulumi.Output<string | undefined>;
    /**
     * A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
     */
    public readonly enforceConsumerDeletion!: pulumi.Output<boolean | undefined>;
    /**
     * The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 8760 hours. Minimum value is 24. Default is 24.
     */
    public readonly retentionPeriod!: pulumi.Output<number | undefined>;
    /**
     * The number of shards that the stream will use. If the `streamMode` is `PROVISIONED`, this field is required.
     * Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
     */
    public readonly shardCount!: pulumi.Output<number | undefined>;
    /**
     * A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
     */
    public readonly shardLevelMetrics!: pulumi.Output<string[] | undefined>;
    /**
     * Indicates the [capacity mode](https://docs.aws.amazon.com/streams/latest/dev/how-do-i-size-a-stream.html) of the data stream. Detailed below.
     */
    public readonly streamModeDetails!: pulumi.Output<outputs.kinesis.StreamStreamModeDetails>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Stream resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: StreamArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StreamArgs | StreamState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StreamState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["encryptionType"] = state ? state.encryptionType : undefined;
            resourceInputs["enforceConsumerDeletion"] = state ? state.enforceConsumerDeletion : undefined;
            resourceInputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["retentionPeriod"] = state ? state.retentionPeriod : undefined;
            resourceInputs["shardCount"] = state ? state.shardCount : undefined;
            resourceInputs["shardLevelMetrics"] = state ? state.shardLevelMetrics : undefined;
            resourceInputs["streamModeDetails"] = state ? state.streamModeDetails : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as StreamArgs | undefined;
            resourceInputs["arn"] = args ? args.arn : undefined;
            resourceInputs["encryptionType"] = args ? args.encryptionType : undefined;
            resourceInputs["enforceConsumerDeletion"] = args ? args.enforceConsumerDeletion : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["retentionPeriod"] = args ? args.retentionPeriod : undefined;
            resourceInputs["shardCount"] = args ? args.shardCount : undefined;
            resourceInputs["shardLevelMetrics"] = args ? args.shardLevelMetrics : undefined;
            resourceInputs["streamModeDetails"] = args ? args.streamModeDetails : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Stream.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Stream resources.
 */
export interface StreamState {
    /**
     * The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
     */
    arn?: pulumi.Input<string>;
    /**
     * The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
     */
    encryptionType?: pulumi.Input<string>;
    /**
     * A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
     */
    enforceConsumerDeletion?: pulumi.Input<boolean>;
    /**
     * The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
     */
    name?: pulumi.Input<string>;
    /**
     * Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 8760 hours. Minimum value is 24. Default is 24.
     */
    retentionPeriod?: pulumi.Input<number>;
    /**
     * The number of shards that the stream will use. If the `streamMode` is `PROVISIONED`, this field is required.
     * Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
     */
    shardCount?: pulumi.Input<number>;
    /**
     * A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
     */
    shardLevelMetrics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates the [capacity mode](https://docs.aws.amazon.com/streams/latest/dev/how-do-i-size-a-stream.html) of the data stream. Detailed below.
     */
    streamModeDetails?: pulumi.Input<inputs.kinesis.StreamStreamModeDetails>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Stream resource.
 */
export interface StreamArgs {
    /**
     * The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
     */
    arn?: pulumi.Input<string>;
    /**
     * The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
     */
    encryptionType?: pulumi.Input<string>;
    /**
     * A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
     */
    enforceConsumerDeletion?: pulumi.Input<boolean>;
    /**
     * The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
     */
    name?: pulumi.Input<string>;
    /**
     * Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 8760 hours. Minimum value is 24. Default is 24.
     */
    retentionPeriod?: pulumi.Input<number>;
    /**
     * The number of shards that the stream will use. If the `streamMode` is `PROVISIONED`, this field is required.
     * Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
     */
    shardCount?: pulumi.Input<number>;
    /**
     * A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
     */
    shardLevelMetrics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates the [capacity mode](https://docs.aws.amazon.com/streams/latest/dev/how-do-i-size-a-stream.html) of the data stream. Detailed below.
     */
    streamModeDetails?: pulumi.Input<inputs.kinesis.StreamStreamModeDetails>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
