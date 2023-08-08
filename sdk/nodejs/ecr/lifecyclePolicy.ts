// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {LifecyclePolicyDocument} from "./index";

/**
 * Manages an ECR repository lifecycle policy.
 *
 * > **NOTE:** Only one `aws.ecr.LifecyclePolicy` resource can be used with the same ECR repository. To apply multiple rules, they must be combined in the `policy` JSON.
 *
 * > **NOTE:** The AWS ECR API seems to reorder rules based on `rulePriority`. If you define multiple rules that are not sorted in ascending `rulePriority` order in the this provider code, the resource will be flagged for recreation every deployment.
 *
 * ## Example Usage
 * ### Policy on untagged image
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const foo = new aws.ecr.Repository("foo", {});
 * const foopolicy = new aws.ecr.LifecyclePolicy("foopolicy", {
 *     repository: foo.name,
 *     policy: `{
 *     "rules": [
 *         {
 *             "rulePriority": 1,
 *             "description": "Expire images older than 14 days",
 *             "selection": {
 *                 "tagStatus": "untagged",
 *                 "countType": "sinceImagePushed",
 *                 "countUnit": "days",
 *                 "countNumber": 14
 *             },
 *             "action": {
 *                 "type": "expire"
 *             }
 *         }
 *     ]
 * }
 * `,
 * });
 * ```
 * ### Policy on tagged image
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const foo = new aws.ecr.Repository("foo", {});
 * const foopolicy = new aws.ecr.LifecyclePolicy("foopolicy", {
 *     repository: foo.name,
 *     policy: `{
 *     "rules": [
 *         {
 *             "rulePriority": 1,
 *             "description": "Keep last 30 images",
 *             "selection": {
 *                 "tagStatus": "tagged",
 *                 "tagPrefixList": ["v"],
 *                 "countType": "imageCountMoreThan",
 *                 "countNumber": 30
 *             },
 *             "action": {
 *                 "type": "expire"
 *             }
 *         }
 *     ]
 * }
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_ecr_lifecycle_policy.example
 *
 *  id = "tf-example" } Using `pulumi import`, import ECR Lifecycle Policy using the name of the repository. For exampleconsole % pulumi import aws_ecr_lifecycle_policy.example tf-example
 */
export class LifecyclePolicy extends pulumi.CustomResource {
    /**
     * Get an existing LifecyclePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LifecyclePolicyState, opts?: pulumi.CustomResourceOptions): LifecyclePolicy {
        return new LifecyclePolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ecr/lifecyclePolicy:LifecyclePolicy';

    /**
     * Returns true if the given object is an instance of LifecyclePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LifecyclePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LifecyclePolicy.__pulumiType;
    }

    /**
     * The policy document. This is a JSON formatted string. See more details about [Policy Parameters](http://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html#lifecycle_policy_parameters) in the official AWS docs.
     */
    public readonly policy!: pulumi.Output<string>;
    /**
     * The registry ID where the repository was created.
     */
    public /*out*/ readonly registryId!: pulumi.Output<string>;
    /**
     * Name of the repository to apply the policy.
     */
    public readonly repository!: pulumi.Output<string>;

    /**
     * Create a LifecyclePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LifecyclePolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LifecyclePolicyArgs | LifecyclePolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LifecyclePolicyState | undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
            resourceInputs["registryId"] = state ? state.registryId : undefined;
            resourceInputs["repository"] = state ? state.repository : undefined;
        } else {
            const args = argsOrState as LifecyclePolicyArgs | undefined;
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            if ((!args || args.repository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repository'");
            }
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["repository"] = args ? args.repository : undefined;
            resourceInputs["registryId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LifecyclePolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LifecyclePolicy resources.
 */
export interface LifecyclePolicyState {
    /**
     * The policy document. This is a JSON formatted string. See more details about [Policy Parameters](http://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html#lifecycle_policy_parameters) in the official AWS docs.
     */
    policy?: pulumi.Input<string | LifecyclePolicyDocument>;
    /**
     * The registry ID where the repository was created.
     */
    registryId?: pulumi.Input<string>;
    /**
     * Name of the repository to apply the policy.
     */
    repository?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LifecyclePolicy resource.
 */
export interface LifecyclePolicyArgs {
    /**
     * The policy document. This is a JSON formatted string. See more details about [Policy Parameters](http://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html#lifecycle_policy_parameters) in the official AWS docs.
     */
    policy: pulumi.Input<string | LifecyclePolicyDocument>;
    /**
     * Name of the repository to apply the policy.
     */
    repository: pulumi.Input<string>;
}
