// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage an S3 Multi-Region Access Point access control policy.
 *
 * ## Example Usage
 * ### Basic Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const currentCallerIdentity = aws.getCallerIdentity({});
 * const currentPartition = aws.getPartition({});
 * const fooBucket = new aws.s3.BucketV2("fooBucket", {});
 * const exampleMultiRegionAccessPoint = new aws.s3control.MultiRegionAccessPoint("exampleMultiRegionAccessPoint", {details: {
 *     name: "example",
 *     regions: [{
 *         bucket: fooBucket.id,
 *     }],
 * }});
 * const exampleMultiRegionAccessPointPolicy = new aws.s3control.MultiRegionAccessPointPolicy("exampleMultiRegionAccessPointPolicy", {details: {
 *     name: exampleMultiRegionAccessPoint.id.apply(id => id.split(":"))[1],
 *     policy: pulumi.all([currentCallerIdentity, currentPartition, currentCallerIdentity, exampleMultiRegionAccessPoint.alias]).apply(([currentCallerIdentity, currentPartition, currentCallerIdentity1, alias]) => JSON.stringify({
 *         Version: "2012-10-17",
 *         Statement: [{
 *             Sid: "Example",
 *             Effect: "Allow",
 *             Principal: {
 *                 AWS: currentCallerIdentity.accountId,
 *             },
 *             Action: [
 *                 "s3:GetObject",
 *                 "s3:PutObject",
 *             ],
 *             Resource: `arn:${currentPartition.partition}:s3::${currentCallerIdentity1.accountId}:accesspoint/${alias}/object/*`,
 *         }],
 *     })),
 * }});
 * ```
 *
 * ## Import
 *
 * Multi-Region Access Point Policies can be imported using the `account_id` and `name` of the Multi-Region Access Point separated by a colon (`:`), e.g.
 *
 * ```sh
 *  $ pulumi import aws:s3control/multiRegionAccessPointPolicy:MultiRegionAccessPointPolicy example 123456789012:example
 * ```
 */
export class MultiRegionAccessPointPolicy extends pulumi.CustomResource {
    /**
     * Get an existing MultiRegionAccessPointPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MultiRegionAccessPointPolicyState, opts?: pulumi.CustomResourceOptions): MultiRegionAccessPointPolicy {
        return new MultiRegionAccessPointPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:s3control/multiRegionAccessPointPolicy:MultiRegionAccessPointPolicy';

    /**
     * Returns true if the given object is an instance of MultiRegionAccessPointPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MultiRegionAccessPointPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MultiRegionAccessPointPolicy.__pulumiType;
    }

    /**
     * The AWS account ID for the owner of the Multi-Region Access Point. Defaults to automatically determined account ID of the AWS provider.
     */
    public readonly accountId!: pulumi.Output<string>;
    /**
     * A configuration block containing details about the policy for the Multi-Region Access Point. See Details Configuration Block below for more details
     */
    public readonly details!: pulumi.Output<outputs.s3control.MultiRegionAccessPointPolicyDetails>;
    /**
     * The last established policy for the Multi-Region Access Point.
     */
    public /*out*/ readonly established!: pulumi.Output<string>;
    /**
     * The proposed policy for the Multi-Region Access Point.
     */
    public /*out*/ readonly proposed!: pulumi.Output<string>;

    /**
     * Create a MultiRegionAccessPointPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MultiRegionAccessPointPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MultiRegionAccessPointPolicyArgs | MultiRegionAccessPointPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MultiRegionAccessPointPolicyState | undefined;
            resourceInputs["accountId"] = state ? state.accountId : undefined;
            resourceInputs["details"] = state ? state.details : undefined;
            resourceInputs["established"] = state ? state.established : undefined;
            resourceInputs["proposed"] = state ? state.proposed : undefined;
        } else {
            const args = argsOrState as MultiRegionAccessPointPolicyArgs | undefined;
            if ((!args || args.details === undefined) && !opts.urn) {
                throw new Error("Missing required property 'details'");
            }
            resourceInputs["accountId"] = args ? args.accountId : undefined;
            resourceInputs["details"] = args ? args.details : undefined;
            resourceInputs["established"] = undefined /*out*/;
            resourceInputs["proposed"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MultiRegionAccessPointPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MultiRegionAccessPointPolicy resources.
 */
export interface MultiRegionAccessPointPolicyState {
    /**
     * The AWS account ID for the owner of the Multi-Region Access Point. Defaults to automatically determined account ID of the AWS provider.
     */
    accountId?: pulumi.Input<string>;
    /**
     * A configuration block containing details about the policy for the Multi-Region Access Point. See Details Configuration Block below for more details
     */
    details?: pulumi.Input<inputs.s3control.MultiRegionAccessPointPolicyDetails>;
    /**
     * The last established policy for the Multi-Region Access Point.
     */
    established?: pulumi.Input<string>;
    /**
     * The proposed policy for the Multi-Region Access Point.
     */
    proposed?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MultiRegionAccessPointPolicy resource.
 */
export interface MultiRegionAccessPointPolicyArgs {
    /**
     * The AWS account ID for the owner of the Multi-Region Access Point. Defaults to automatically determined account ID of the AWS provider.
     */
    accountId?: pulumi.Input<string>;
    /**
     * A configuration block containing details about the policy for the Multi-Region Access Point. See Details Configuration Block below for more details
     */
    details: pulumi.Input<inputs.s3control.MultiRegionAccessPointPolicyDetails>;
}
