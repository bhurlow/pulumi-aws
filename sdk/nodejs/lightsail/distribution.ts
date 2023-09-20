// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource for managing an AWS Lightsail Distribution.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * Below is a basic example with a bucket as an origin.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testBucket = new aws.lightsail.Bucket("testBucket", {bundleId: "small_1_0"});
 * const testDistribution = new aws.lightsail.Distribution("testDistribution", {
 *     bundleId: "small_1_0",
 *     origin: {
 *         name: testBucket.name,
 *         regionName: testBucket.region,
 *     },
 *     defaultCacheBehavior: {
 *         behavior: "cache",
 *     },
 *     cacheBehaviorSettings: {
 *         allowedHttpMethods: "GET,HEAD,OPTIONS,PUT,PATCH,POST,DELETE",
 *         cachedHttpMethods: "GET,HEAD",
 *         defaultTtl: 86400,
 *         maximumTtl: 31536000,
 *         minimumTtl: 0,
 *         forwardedCookies: {
 *             option: "none",
 *         },
 *         forwardedHeaders: {
 *             option: "default",
 *         },
 *         forwardedQueryStrings: {
 *             option: false,
 *         },
 *     },
 * });
 * ```
 * ### instance origin example
 *
 * Below is an example of an instance as the origin.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const available = aws.getAvailabilityZones({
 *     state: "available",
 *     filters: [{
 *         name: "opt-in-status",
 *         values: ["opt-in-not-required"],
 *     }],
 * });
 * const testStaticIp = new aws.lightsail.StaticIp("testStaticIp", {});
 * const testInstance = new aws.lightsail.Instance("testInstance", {
 *     availabilityZone: available.then(available => available.names?.[0]),
 *     blueprintId: "amazon_linux_2",
 *     bundleId: "micro_1_0",
 * });
 * const testStaticIpAttachment = new aws.lightsail.StaticIpAttachment("testStaticIpAttachment", {
 *     staticIpName: testStaticIp.name,
 *     instanceName: testInstance.name,
 * });
 * const testDistribution = new aws.lightsail.Distribution("testDistribution", {
 *     bundleId: "small_1_0",
 *     origin: {
 *         name: testInstance.name,
 *         regionName: available.then(available => available.id),
 *     },
 *     defaultCacheBehavior: {
 *         behavior: "cache",
 *     },
 * }, {
 *     dependsOn: [testStaticIpAttachment],
 * });
 * ```
 * ### lb origin example
 *
 * Below is an example with a load balancer as an origin
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const available = aws.getAvailabilityZones({
 *     state: "available",
 *     filters: [{
 *         name: "opt-in-status",
 *         values: ["opt-in-not-required"],
 *     }],
 * });
 * const testLb = new aws.lightsail.Lb("testLb", {
 *     healthCheckPath: "/",
 *     instancePort: 80,
 *     tags: {
 *         foo: "bar",
 *     },
 * });
 * const testInstance = new aws.lightsail.Instance("testInstance", {
 *     availabilityZone: available.then(available => available.names?.[0]),
 *     blueprintId: "amazon_linux_2",
 *     bundleId: "nano_1_0",
 * });
 * const testLbAttachment = new aws.lightsail.LbAttachment("testLbAttachment", {
 *     lbName: testLb.name,
 *     instanceName: testInstance.name,
 * });
 * const testDistribution = new aws.lightsail.Distribution("testDistribution", {
 *     bundleId: "small_1_0",
 *     origin: {
 *         name: testLb.name,
 *         regionName: available.then(available => available.id),
 *     },
 *     defaultCacheBehavior: {
 *         behavior: "cache",
 *     },
 * }, {
 *     dependsOn: [testLbAttachment],
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import Lightsail Distribution using the `id`. For example:
 *
 * ```sh
 *  $ pulumi import aws:lightsail/distribution:Distribution example rft-8012925589
 * ```
 */
export class Distribution extends pulumi.CustomResource {
    /**
     * Get an existing Distribution resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DistributionState, opts?: pulumi.CustomResourceOptions): Distribution {
        return new Distribution(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:lightsail/distribution:Distribution';

    /**
     * Returns true if the given object is an instance of Distribution.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Distribution {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Distribution.__pulumiType;
    }

    /**
     * The alternate domain names of the distribution.
     */
    public /*out*/ readonly alternativeDomainNames!: pulumi.Output<string[]>;
    /**
     * The Amazon Resource Name (ARN) of the distribution.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Bundle ID to use for the distribution.
     */
    public readonly bundleId!: pulumi.Output<string>;
    /**
     * An object that describes the cache behavior settings of the distribution. Detailed below
     *
     * The following arguments are optional:
     */
    public readonly cacheBehaviorSettings!: pulumi.Output<outputs.lightsail.DistributionCacheBehaviorSettings | undefined>;
    /**
     * A set of configuration blocks that describe the per-path cache behavior of the distribution. Detailed below
     */
    public readonly cacheBehaviors!: pulumi.Output<outputs.lightsail.DistributionCacheBehavior[] | undefined>;
    /**
     * The name of the SSL/TLS certificate attached to the distribution, if any.
     */
    public readonly certificateName!: pulumi.Output<string | undefined>;
    /**
     * The timestamp when the distribution was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Object that describes the default cache behavior of the distribution. Detailed below
     */
    public readonly defaultCacheBehavior!: pulumi.Output<outputs.lightsail.DistributionDefaultCacheBehavior>;
    /**
     * The domain name of the distribution.
     */
    public /*out*/ readonly domainName!: pulumi.Output<string>;
    /**
     * The IP address type of the distribution. Default: `dualstack`.
     */
    public readonly ipAddressType!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether the distribution is enabled. Default: `true`.
     */
    public readonly isEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * An object that describes the location of the distribution, such as the AWS Region and Availability Zone. Detailed below
     */
    public /*out*/ readonly locations!: pulumi.Output<outputs.lightsail.DistributionLocation[]>;
    /**
     * Name of the distribution.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Object that describes the origin resource of the distribution, such as a Lightsail instance, bucket, or load balancer. Detailed below
     */
    public readonly origin!: pulumi.Output<outputs.lightsail.DistributionOrigin>;
    /**
     * The public DNS of the origin.
     */
    public /*out*/ readonly originPublicDns!: pulumi.Output<string>;
    /**
     * The resource type of the origin resource (e.g., Instance).
     */
    public /*out*/ readonly resourceType!: pulumi.Output<string>;
    /**
     * The status of the distribution.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The support code. Include this code in your email to support when you have questions about your Lightsail distribution. This code enables our support team to look up your Lightsail information more easily.
     */
    public /*out*/ readonly supportCode!: pulumi.Output<string>;
    /**
     * Map of tags for the Lightsail Distribution. If
     * configured with a provider
     * `defaultTags` configuration block
     * present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Distribution resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DistributionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DistributionArgs | DistributionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DistributionState | undefined;
            resourceInputs["alternativeDomainNames"] = state ? state.alternativeDomainNames : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["bundleId"] = state ? state.bundleId : undefined;
            resourceInputs["cacheBehaviorSettings"] = state ? state.cacheBehaviorSettings : undefined;
            resourceInputs["cacheBehaviors"] = state ? state.cacheBehaviors : undefined;
            resourceInputs["certificateName"] = state ? state.certificateName : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["defaultCacheBehavior"] = state ? state.defaultCacheBehavior : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["ipAddressType"] = state ? state.ipAddressType : undefined;
            resourceInputs["isEnabled"] = state ? state.isEnabled : undefined;
            resourceInputs["locations"] = state ? state.locations : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["origin"] = state ? state.origin : undefined;
            resourceInputs["originPublicDns"] = state ? state.originPublicDns : undefined;
            resourceInputs["resourceType"] = state ? state.resourceType : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["supportCode"] = state ? state.supportCode : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as DistributionArgs | undefined;
            if ((!args || args.bundleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bundleId'");
            }
            if ((!args || args.defaultCacheBehavior === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultCacheBehavior'");
            }
            if ((!args || args.origin === undefined) && !opts.urn) {
                throw new Error("Missing required property 'origin'");
            }
            resourceInputs["bundleId"] = args ? args.bundleId : undefined;
            resourceInputs["cacheBehaviorSettings"] = args ? args.cacheBehaviorSettings : undefined;
            resourceInputs["cacheBehaviors"] = args ? args.cacheBehaviors : undefined;
            resourceInputs["certificateName"] = args ? args.certificateName : undefined;
            resourceInputs["defaultCacheBehavior"] = args ? args.defaultCacheBehavior : undefined;
            resourceInputs["ipAddressType"] = args ? args.ipAddressType : undefined;
            resourceInputs["isEnabled"] = args ? args.isEnabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["origin"] = args ? args.origin : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["alternativeDomainNames"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["domainName"] = undefined /*out*/;
            resourceInputs["locations"] = undefined /*out*/;
            resourceInputs["originPublicDns"] = undefined /*out*/;
            resourceInputs["resourceType"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["supportCode"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Distribution.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Distribution resources.
 */
export interface DistributionState {
    /**
     * The alternate domain names of the distribution.
     */
    alternativeDomainNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Amazon Resource Name (ARN) of the distribution.
     */
    arn?: pulumi.Input<string>;
    /**
     * Bundle ID to use for the distribution.
     */
    bundleId?: pulumi.Input<string>;
    /**
     * An object that describes the cache behavior settings of the distribution. Detailed below
     *
     * The following arguments are optional:
     */
    cacheBehaviorSettings?: pulumi.Input<inputs.lightsail.DistributionCacheBehaviorSettings>;
    /**
     * A set of configuration blocks that describe the per-path cache behavior of the distribution. Detailed below
     */
    cacheBehaviors?: pulumi.Input<pulumi.Input<inputs.lightsail.DistributionCacheBehavior>[]>;
    /**
     * The name of the SSL/TLS certificate attached to the distribution, if any.
     */
    certificateName?: pulumi.Input<string>;
    /**
     * The timestamp when the distribution was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Object that describes the default cache behavior of the distribution. Detailed below
     */
    defaultCacheBehavior?: pulumi.Input<inputs.lightsail.DistributionDefaultCacheBehavior>;
    /**
     * The domain name of the distribution.
     */
    domainName?: pulumi.Input<string>;
    /**
     * The IP address type of the distribution. Default: `dualstack`.
     */
    ipAddressType?: pulumi.Input<string>;
    /**
     * Indicates whether the distribution is enabled. Default: `true`.
     */
    isEnabled?: pulumi.Input<boolean>;
    /**
     * An object that describes the location of the distribution, such as the AWS Region and Availability Zone. Detailed below
     */
    locations?: pulumi.Input<pulumi.Input<inputs.lightsail.DistributionLocation>[]>;
    /**
     * Name of the distribution.
     */
    name?: pulumi.Input<string>;
    /**
     * Object that describes the origin resource of the distribution, such as a Lightsail instance, bucket, or load balancer. Detailed below
     */
    origin?: pulumi.Input<inputs.lightsail.DistributionOrigin>;
    /**
     * The public DNS of the origin.
     */
    originPublicDns?: pulumi.Input<string>;
    /**
     * The resource type of the origin resource (e.g., Instance).
     */
    resourceType?: pulumi.Input<string>;
    /**
     * The status of the distribution.
     */
    status?: pulumi.Input<string>;
    /**
     * The support code. Include this code in your email to support when you have questions about your Lightsail distribution. This code enables our support team to look up your Lightsail information more easily.
     */
    supportCode?: pulumi.Input<string>;
    /**
     * Map of tags for the Lightsail Distribution. If
     * configured with a provider
     * `defaultTags` configuration block
     * present, tags with matching keys will overwrite those defined at the provider-level.
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
 * The set of arguments for constructing a Distribution resource.
 */
export interface DistributionArgs {
    /**
     * Bundle ID to use for the distribution.
     */
    bundleId: pulumi.Input<string>;
    /**
     * An object that describes the cache behavior settings of the distribution. Detailed below
     *
     * The following arguments are optional:
     */
    cacheBehaviorSettings?: pulumi.Input<inputs.lightsail.DistributionCacheBehaviorSettings>;
    /**
     * A set of configuration blocks that describe the per-path cache behavior of the distribution. Detailed below
     */
    cacheBehaviors?: pulumi.Input<pulumi.Input<inputs.lightsail.DistributionCacheBehavior>[]>;
    /**
     * The name of the SSL/TLS certificate attached to the distribution, if any.
     */
    certificateName?: pulumi.Input<string>;
    /**
     * Object that describes the default cache behavior of the distribution. Detailed below
     */
    defaultCacheBehavior: pulumi.Input<inputs.lightsail.DistributionDefaultCacheBehavior>;
    /**
     * The IP address type of the distribution. Default: `dualstack`.
     */
    ipAddressType?: pulumi.Input<string>;
    /**
     * Indicates whether the distribution is enabled. Default: `true`.
     */
    isEnabled?: pulumi.Input<boolean>;
    /**
     * Name of the distribution.
     */
    name?: pulumi.Input<string>;
    /**
     * Object that describes the origin resource of the distribution, such as a Lightsail instance, bucket, or load balancer. Detailed below
     */
    origin: pulumi.Input<inputs.lightsail.DistributionOrigin>;
    /**
     * Map of tags for the Lightsail Distribution. If
     * configured with a provider
     * `defaultTags` configuration block
     * present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
