// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a CloudFront real-time log configuration resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const assumeRole = aws.iam.getPolicyDocument({
 *     statements: [{
 *         effect: "Allow",
 *         principals: [{
 *             type: "Service",
 *             identifiers: ["cloudfront.amazonaws.com"],
 *         }],
 *         actions: ["sts:AssumeRole"],
 *     }],
 * });
 * const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
 * const examplePolicyDocument = aws.iam.getPolicyDocument({
 *     statements: [{
 *         effect: "Allow",
 *         actions: [
 *             "kinesis:DescribeStreamSummary",
 *             "kinesis:DescribeStream",
 *             "kinesis:PutRecord",
 *             "kinesis:PutRecords",
 *         ],
 *         resources: [aws_kinesis_stream.example.arn],
 *     }],
 * });
 * const exampleRolePolicy = new aws.iam.RolePolicy("exampleRolePolicy", {
 *     role: exampleRole.id,
 *     policy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json),
 * });
 * const exampleRealtimeLogConfig = new aws.cloudfront.RealtimeLogConfig("exampleRealtimeLogConfig", {
 *     samplingRate: 75,
 *     fields: [
 *         "timestamp",
 *         "c-ip",
 *     ],
 *     endpoint: {
 *         streamType: "Kinesis",
 *         kinesisStreamConfig: {
 *             roleArn: exampleRole.arn,
 *             streamArn: aws_kinesis_stream.example.arn,
 *         },
 *     },
 * }, {
 *     dependsOn: [exampleRolePolicy],
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_cloudfront_realtime_log_config.example
 *
 *  id = "arn:aws:cloudfront::111122223333:realtime-log-config/ExampleNameForRealtimeLogConfig" } Using `pulumi import`, import CloudFront real-time log configurations using the ARN. For exampleconsole % pulumi import aws_cloudfront_realtime_log_config.example arn:aws:cloudfront::111122223333:realtime-log-config/ExampleNameForRealtimeLogConfig
 */
export class RealtimeLogConfig extends pulumi.CustomResource {
    /**
     * Get an existing RealtimeLogConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RealtimeLogConfigState, opts?: pulumi.CustomResourceOptions): RealtimeLogConfig {
        return new RealtimeLogConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudfront/realtimeLogConfig:RealtimeLogConfig';

    /**
     * Returns true if the given object is an instance of RealtimeLogConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RealtimeLogConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RealtimeLogConfig.__pulumiType;
    }

    /**
     * The ARN (Amazon Resource Name) of the CloudFront real-time log configuration.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The Amazon Kinesis data streams where real-time log data is sent.
     */
    public readonly endpoint!: pulumi.Output<outputs.cloudfront.RealtimeLogConfigEndpoint>;
    /**
     * The fields that are included in each real-time log record. See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) for supported values.
     */
    public readonly fields!: pulumi.Output<string[]>;
    /**
     * The unique name to identify this real-time log configuration.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between `1` and `100`, inclusive.
     */
    public readonly samplingRate!: pulumi.Output<number>;

    /**
     * Create a RealtimeLogConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RealtimeLogConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RealtimeLogConfigArgs | RealtimeLogConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RealtimeLogConfigState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["fields"] = state ? state.fields : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["samplingRate"] = state ? state.samplingRate : undefined;
        } else {
            const args = argsOrState as RealtimeLogConfigArgs | undefined;
            if ((!args || args.endpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpoint'");
            }
            if ((!args || args.fields === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fields'");
            }
            if ((!args || args.samplingRate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'samplingRate'");
            }
            resourceInputs["endpoint"] = args ? args.endpoint : undefined;
            resourceInputs["fields"] = args ? args.fields : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["samplingRate"] = args ? args.samplingRate : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RealtimeLogConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RealtimeLogConfig resources.
 */
export interface RealtimeLogConfigState {
    /**
     * The ARN (Amazon Resource Name) of the CloudFront real-time log configuration.
     */
    arn?: pulumi.Input<string>;
    /**
     * The Amazon Kinesis data streams where real-time log data is sent.
     */
    endpoint?: pulumi.Input<inputs.cloudfront.RealtimeLogConfigEndpoint>;
    /**
     * The fields that are included in each real-time log record. See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) for supported values.
     */
    fields?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The unique name to identify this real-time log configuration.
     */
    name?: pulumi.Input<string>;
    /**
     * The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between `1` and `100`, inclusive.
     */
    samplingRate?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RealtimeLogConfig resource.
 */
export interface RealtimeLogConfigArgs {
    /**
     * The Amazon Kinesis data streams where real-time log data is sent.
     */
    endpoint: pulumi.Input<inputs.cloudfront.RealtimeLogConfigEndpoint>;
    /**
     * The fields that are included in each real-time log record. See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) for supported values.
     */
    fields: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The unique name to identify this real-time log configuration.
     */
    name?: pulumi.Input<string>;
    /**
     * The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between `1` and `100`, inclusive.
     */
    samplingRate: pulumi.Input<number>;
}
