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
 * const example = new aws.cloudfront.MonitoringSubscription("example", {
 *     distributionId: aws_cloudfront_distribution.example.id,
 *     monitoringSubscription: {
 *         realtimeMetricsSubscriptionConfig: {
 *             realtimeMetricsSubscriptionStatus: "Enabled",
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_cloudfront_monitoring_subscription.example
 *
 *  id = "E3QYSUHO4VYRGB" } Using `pulumi import`, import CloudFront monitoring subscription using the id. For exampleconsole % pulumi import aws_cloudfront_monitoring_subscription.example E3QYSUHO4VYRGB
 */
export class MonitoringSubscription extends pulumi.CustomResource {
    /**
     * Get an existing MonitoringSubscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MonitoringSubscriptionState, opts?: pulumi.CustomResourceOptions): MonitoringSubscription {
        return new MonitoringSubscription(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudfront/monitoringSubscription:MonitoringSubscription';

    /**
     * Returns true if the given object is an instance of MonitoringSubscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MonitoringSubscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MonitoringSubscription.__pulumiType;
    }

    /**
     * The ID of the distribution that you are enabling metrics for.
     */
    public readonly distributionId!: pulumi.Output<string>;
    /**
     * A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
     */
    public readonly monitoringSubscription!: pulumi.Output<outputs.cloudfront.MonitoringSubscriptionMonitoringSubscription>;

    /**
     * Create a MonitoringSubscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MonitoringSubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MonitoringSubscriptionArgs | MonitoringSubscriptionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MonitoringSubscriptionState | undefined;
            resourceInputs["distributionId"] = state ? state.distributionId : undefined;
            resourceInputs["monitoringSubscription"] = state ? state.monitoringSubscription : undefined;
        } else {
            const args = argsOrState as MonitoringSubscriptionArgs | undefined;
            if ((!args || args.distributionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'distributionId'");
            }
            if ((!args || args.monitoringSubscription === undefined) && !opts.urn) {
                throw new Error("Missing required property 'monitoringSubscription'");
            }
            resourceInputs["distributionId"] = args ? args.distributionId : undefined;
            resourceInputs["monitoringSubscription"] = args ? args.monitoringSubscription : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MonitoringSubscription.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MonitoringSubscription resources.
 */
export interface MonitoringSubscriptionState {
    /**
     * The ID of the distribution that you are enabling metrics for.
     */
    distributionId?: pulumi.Input<string>;
    /**
     * A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
     */
    monitoringSubscription?: pulumi.Input<inputs.cloudfront.MonitoringSubscriptionMonitoringSubscription>;
}

/**
 * The set of arguments for constructing a MonitoringSubscription resource.
 */
export interface MonitoringSubscriptionArgs {
    /**
     * The ID of the distribution that you are enabling metrics for.
     */
    distributionId: pulumi.Input<string>;
    /**
     * A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
     */
    monitoringSubscription: pulumi.Input<inputs.cloudfront.MonitoringSubscriptionMonitoringSubscription>;
}
