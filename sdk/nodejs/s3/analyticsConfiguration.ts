// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a S3 bucket [analytics configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/analytics-storage-class.html) resource.
 *
 * ## Example Usage
 * ### Add analytics configuration for entire S3 bucket and export results to a second S3 bucket
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.s3.BucketV2("example", {});
 * const analytics = new aws.s3.BucketV2("analytics", {});
 * const example_entire_bucket = new aws.s3.AnalyticsConfiguration("example-entire-bucket", {
 *     bucket: example.id,
 *     storageClassAnalysis: {
 *         dataExport: {
 *             destination: {
 *                 s3BucketDestination: {
 *                     bucketArn: analytics.arn,
 *                 },
 *             },
 *         },
 *     },
 * });
 * ```
 * ### Add analytics configuration with S3 object filter
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.s3.BucketV2("example", {});
 * const example_filtered = new aws.s3.AnalyticsConfiguration("example-filtered", {
 *     bucket: example.id,
 *     filter: {
 *         prefix: "documents/",
 *         tags: {
 *             priority: "high",
 *             "class": "blue",
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_s3_bucket_analytics_configuration.my-bucket-entire-bucket
 *
 *  id = "my-bucket:EntireBucket" } Using `pulumi import`, import S3 bucket analytics configurations using `bucket:analytics`. For exampleconsole % pulumi import aws_s3_bucket_analytics_configuration.my-bucket-entire-bucket my-bucket:EntireBucket
 */
export class AnalyticsConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing AnalyticsConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AnalyticsConfigurationState, opts?: pulumi.CustomResourceOptions): AnalyticsConfiguration {
        return new AnalyticsConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:s3/analyticsConfiguration:AnalyticsConfiguration';

    /**
     * Returns true if the given object is an instance of AnalyticsConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AnalyticsConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AnalyticsConfiguration.__pulumiType;
    }

    /**
     * Name of the bucket this analytics configuration is associated with.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
     */
    public readonly filter!: pulumi.Output<outputs.s3.AnalyticsConfigurationFilter | undefined>;
    /**
     * Unique identifier of the analytics configuration for the bucket.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configuration for the analytics data export (documented below).
     */
    public readonly storageClassAnalysis!: pulumi.Output<outputs.s3.AnalyticsConfigurationStorageClassAnalysis | undefined>;

    /**
     * Create a AnalyticsConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AnalyticsConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AnalyticsConfigurationArgs | AnalyticsConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AnalyticsConfigurationState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["filter"] = state ? state.filter : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["storageClassAnalysis"] = state ? state.storageClassAnalysis : undefined;
        } else {
            const args = argsOrState as AnalyticsConfigurationArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["filter"] = args ? args.filter : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["storageClassAnalysis"] = args ? args.storageClassAnalysis : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AnalyticsConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AnalyticsConfiguration resources.
 */
export interface AnalyticsConfigurationState {
    /**
     * Name of the bucket this analytics configuration is associated with.
     */
    bucket?: pulumi.Input<string>;
    /**
     * Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
     */
    filter?: pulumi.Input<inputs.s3.AnalyticsConfigurationFilter>;
    /**
     * Unique identifier of the analytics configuration for the bucket.
     */
    name?: pulumi.Input<string>;
    /**
     * Configuration for the analytics data export (documented below).
     */
    storageClassAnalysis?: pulumi.Input<inputs.s3.AnalyticsConfigurationStorageClassAnalysis>;
}

/**
 * The set of arguments for constructing a AnalyticsConfiguration resource.
 */
export interface AnalyticsConfigurationArgs {
    /**
     * Name of the bucket this analytics configuration is associated with.
     */
    bucket: pulumi.Input<string>;
    /**
     * Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
     */
    filter?: pulumi.Input<inputs.s3.AnalyticsConfigurationFilter>;
    /**
     * Unique identifier of the analytics configuration for the bucket.
     */
    name?: pulumi.Input<string>;
    /**
     * Configuration for the analytics data export (documented below).
     */
    storageClassAnalysis?: pulumi.Input<inputs.s3.AnalyticsConfigurationStorageClassAnalysis>;
}
