// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage an [AWS Macie Classification Job](https://docs.aws.amazon.com/macie/latest/APIReference/jobs.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testAccount = new aws.macie2.Account("testAccount", {});
 * const testClassificationJob = new aws.macie2.ClassificationJob("testClassificationJob", {
 *     jobType: "ONE_TIME",
 *     s3JobDefinition: {
 *         bucketDefinitions: [{
 *             accountId: "ACCOUNT ID",
 *             buckets: ["S3 BUCKET NAME"],
 *         }],
 *     },
 * }, {
 *     dependsOn: [testAccount],
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_macie2_classification_job.example
 *
 *  id = "abcd1" } Using `pulumi import`, import `aws_macie2_classification_job` using the id. For exampleconsole % pulumi import aws_macie2_classification_job.example abcd1
 */
export class ClassificationJob extends pulumi.CustomResource {
    /**
     * Get an existing ClassificationJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClassificationJobState, opts?: pulumi.CustomResourceOptions): ClassificationJob {
        return new ClassificationJob(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:macie2/classificationJob:ClassificationJob';

    /**
     * Returns true if the given object is an instance of ClassificationJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClassificationJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClassificationJob.__pulumiType;
    }

    /**
     * The date and time, in UTC and extended RFC 3339 format, when the job was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The custom data identifiers to use for data analysis and classification.
     */
    public readonly customDataIdentifierIds!: pulumi.Output<string[]>;
    /**
     * A custom description of the job. The description can contain as many as 200 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Specifies whether to analyze all existing, eligible objects immediately after the job is created.
     */
    public readonly initialRun!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly jobArn!: pulumi.Output<string>;
    public /*out*/ readonly jobId!: pulumi.Output<string>;
    /**
     * The status for the job. Valid values are: `CANCELLED`, `RUNNING` and `USER_PAUSED`
     */
    public readonly jobStatus!: pulumi.Output<string>;
    /**
     * The schedule for running the job. Valid values are: `ONE_TIME` - Run the job only once. If you specify this value, don't specify a value for the `scheduleFrequency` property. `SCHEDULED` - Run the job on a daily, weekly, or monthly basis. If you specify this value, use the `scheduleFrequency` property to define the recurrence pattern for the job.
     */
    public readonly jobType!: pulumi.Output<string>;
    /**
     * A custom name for the job. The name can contain as many as 500 characters. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string>;
    /**
     * The S3 buckets that contain the objects to analyze, and the scope of that analysis. (documented below)
     */
    public readonly s3JobDefinition!: pulumi.Output<outputs.macie2.ClassificationJobS3JobDefinition>;
    /**
     * The sampling depth, as a percentage, to apply when processing objects. This value determines the percentage of eligible objects that the job analyzes. If this value is less than 100, Amazon Macie selects the objects to analyze at random, up to the specified percentage, and analyzes all the data in those objects.
     */
    public readonly samplingPercentage!: pulumi.Output<number>;
    /**
     * The recurrence pattern for running the job. To run the job only once, don't specify a value for this property and set the value for the `jobType` property to `ONE_TIME`. (documented below)
     */
    public readonly scheduleFrequency!: pulumi.Output<outputs.macie2.ClassificationJobScheduleFrequency>;
    /**
     * A map of key-value pairs that specifies the tags to associate with the job. A job can have a maximum of 50 tags. Each tag consists of a tag key and an associated tag value. The maximum length of a tag key is 128 characters. The maximum length of a tag value is 256 characters.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * If the current status of the job is `USER_PAUSED`, specifies when the job was paused and when the job or job run will expire and be cancelled if it isn't resumed. This value is present only if the value for `job-status` is `USER_PAUSED`.
     */
    public /*out*/ readonly userPausedDetails!: pulumi.Output<outputs.macie2.ClassificationJobUserPausedDetail[]>;

    /**
     * Create a ClassificationJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClassificationJobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClassificationJobArgs | ClassificationJobState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClassificationJobState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["customDataIdentifierIds"] = state ? state.customDataIdentifierIds : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["initialRun"] = state ? state.initialRun : undefined;
            resourceInputs["jobArn"] = state ? state.jobArn : undefined;
            resourceInputs["jobId"] = state ? state.jobId : undefined;
            resourceInputs["jobStatus"] = state ? state.jobStatus : undefined;
            resourceInputs["jobType"] = state ? state.jobType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namePrefix"] = state ? state.namePrefix : undefined;
            resourceInputs["s3JobDefinition"] = state ? state.s3JobDefinition : undefined;
            resourceInputs["samplingPercentage"] = state ? state.samplingPercentage : undefined;
            resourceInputs["scheduleFrequency"] = state ? state.scheduleFrequency : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["userPausedDetails"] = state ? state.userPausedDetails : undefined;
        } else {
            const args = argsOrState as ClassificationJobArgs | undefined;
            if ((!args || args.jobType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobType'");
            }
            if ((!args || args.s3JobDefinition === undefined) && !opts.urn) {
                throw new Error("Missing required property 's3JobDefinition'");
            }
            resourceInputs["customDataIdentifierIds"] = args ? args.customDataIdentifierIds : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["initialRun"] = args ? args.initialRun : undefined;
            resourceInputs["jobStatus"] = args ? args.jobStatus : undefined;
            resourceInputs["jobType"] = args ? args.jobType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["s3JobDefinition"] = args ? args.s3JobDefinition : undefined;
            resourceInputs["samplingPercentage"] = args ? args.samplingPercentage : undefined;
            resourceInputs["scheduleFrequency"] = args ? args.scheduleFrequency : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["jobArn"] = undefined /*out*/;
            resourceInputs["jobId"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["userPausedDetails"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClassificationJob.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClassificationJob resources.
 */
export interface ClassificationJobState {
    /**
     * The date and time, in UTC and extended RFC 3339 format, when the job was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The custom data identifiers to use for data analysis and classification.
     */
    customDataIdentifierIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A custom description of the job. The description can contain as many as 200 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies whether to analyze all existing, eligible objects immediately after the job is created.
     */
    initialRun?: pulumi.Input<boolean>;
    jobArn?: pulumi.Input<string>;
    jobId?: pulumi.Input<string>;
    /**
     * The status for the job. Valid values are: `CANCELLED`, `RUNNING` and `USER_PAUSED`
     */
    jobStatus?: pulumi.Input<string>;
    /**
     * The schedule for running the job. Valid values are: `ONE_TIME` - Run the job only once. If you specify this value, don't specify a value for the `scheduleFrequency` property. `SCHEDULED` - Run the job on a daily, weekly, or monthly basis. If you specify this value, use the `scheduleFrequency` property to define the recurrence pattern for the job.
     */
    jobType?: pulumi.Input<string>;
    /**
     * A custom name for the job. The name can contain as many as 500 characters. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The S3 buckets that contain the objects to analyze, and the scope of that analysis. (documented below)
     */
    s3JobDefinition?: pulumi.Input<inputs.macie2.ClassificationJobS3JobDefinition>;
    /**
     * The sampling depth, as a percentage, to apply when processing objects. This value determines the percentage of eligible objects that the job analyzes. If this value is less than 100, Amazon Macie selects the objects to analyze at random, up to the specified percentage, and analyzes all the data in those objects.
     */
    samplingPercentage?: pulumi.Input<number>;
    /**
     * The recurrence pattern for running the job. To run the job only once, don't specify a value for this property and set the value for the `jobType` property to `ONE_TIME`. (documented below)
     */
    scheduleFrequency?: pulumi.Input<inputs.macie2.ClassificationJobScheduleFrequency>;
    /**
     * A map of key-value pairs that specifies the tags to associate with the job. A job can have a maximum of 50 tags. Each tag consists of a tag key and an associated tag value. The maximum length of a tag key is 128 characters. The maximum length of a tag value is 256 characters.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * If the current status of the job is `USER_PAUSED`, specifies when the job was paused and when the job or job run will expire and be cancelled if it isn't resumed. This value is present only if the value for `job-status` is `USER_PAUSED`.
     */
    userPausedDetails?: pulumi.Input<pulumi.Input<inputs.macie2.ClassificationJobUserPausedDetail>[]>;
}

/**
 * The set of arguments for constructing a ClassificationJob resource.
 */
export interface ClassificationJobArgs {
    /**
     * The custom data identifiers to use for data analysis and classification.
     */
    customDataIdentifierIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A custom description of the job. The description can contain as many as 200 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies whether to analyze all existing, eligible objects immediately after the job is created.
     */
    initialRun?: pulumi.Input<boolean>;
    /**
     * The status for the job. Valid values are: `CANCELLED`, `RUNNING` and `USER_PAUSED`
     */
    jobStatus?: pulumi.Input<string>;
    /**
     * The schedule for running the job. Valid values are: `ONE_TIME` - Run the job only once. If you specify this value, don't specify a value for the `scheduleFrequency` property. `SCHEDULED` - Run the job on a daily, weekly, or monthly basis. If you specify this value, use the `scheduleFrequency` property to define the recurrence pattern for the job.
     */
    jobType: pulumi.Input<string>;
    /**
     * A custom name for the job. The name can contain as many as 500 characters. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The S3 buckets that contain the objects to analyze, and the scope of that analysis. (documented below)
     */
    s3JobDefinition: pulumi.Input<inputs.macie2.ClassificationJobS3JobDefinition>;
    /**
     * The sampling depth, as a percentage, to apply when processing objects. This value determines the percentage of eligible objects that the job analyzes. If this value is less than 100, Amazon Macie selects the objects to analyze at random, up to the specified percentage, and analyzes all the data in those objects.
     */
    samplingPercentage?: pulumi.Input<number>;
    /**
     * The recurrence pattern for running the job. To run the job only once, don't specify a value for this property and set the value for the `jobType` property to `ONE_TIME`. (documented below)
     */
    scheduleFrequency?: pulumi.Input<inputs.macie2.ClassificationJobScheduleFrequency>;
    /**
     * A map of key-value pairs that specifies the tags to associate with the job. A job can have a maximum of 50 tags. Each tag consists of a tag key and an associated tag value. The maximum length of a tag key is 128 characters. The maximum length of a tag value is 256 characters.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
