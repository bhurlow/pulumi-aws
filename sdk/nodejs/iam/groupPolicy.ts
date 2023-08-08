// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {PolicyDocument} from "./index";

/**
 * Provides an IAM policy attached to a group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const myDevelopers = new aws.iam.Group("myDevelopers", {path: "/users/"});
 * const myDeveloperPolicy = new aws.iam.GroupPolicy("myDeveloperPolicy", {
 *     group: myDevelopers.name,
 *     policy: JSON.stringify({
 *         Version: "2012-10-17",
 *         Statement: [{
 *             Action: ["ec2:Describe*"],
 *             Effect: "Allow",
 *             Resource: "*",
 *         }],
 *     }),
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_iam_group_policy.mypolicy
 *
 *  id = "group_of_mypolicy_name:mypolicy_name" } Using `pulumi import`, import IAM Group Policies using the `group_name:group_policy_name`. For exampleconsole % pulumi import aws_iam_group_policy.mypolicy group_of_mypolicy_name:mypolicy_name
 */
export class GroupPolicy extends pulumi.CustomResource {
    /**
     * Get an existing GroupPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupPolicyState, opts?: pulumi.CustomResourceOptions): GroupPolicy {
        return new GroupPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:iam/groupPolicy:GroupPolicy';

    /**
     * Returns true if the given object is an instance of GroupPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupPolicy.__pulumiType;
    }

    /**
     * The IAM group to attach to the policy.
     */
    public readonly group!: pulumi.Output<string>;
    /**
     * The name of the policy. If omitted, the provider will
     * assign a random, unique name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * The policy document. This is a JSON formatted string.
     */
    public readonly policy!: pulumi.Output<string>;

    /**
     * Create a GroupPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupPolicyArgs | GroupPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupPolicyState | undefined;
            resourceInputs["group"] = state ? state.group : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namePrefix"] = state ? state.namePrefix : undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
        } else {
            const args = argsOrState as GroupPolicyArgs | undefined;
            if ((!args || args.group === undefined) && !opts.urn) {
                throw new Error("Missing required property 'group'");
            }
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GroupPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupPolicy resources.
 */
export interface GroupPolicyState {
    /**
     * The IAM group to attach to the policy.
     */
    group?: pulumi.Input<string>;
    /**
     * The name of the policy. If omitted, the provider will
     * assign a random, unique name.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The policy document. This is a JSON formatted string.
     */
    policy?: pulumi.Input<string | PolicyDocument>;
}

/**
 * The set of arguments for constructing a GroupPolicy resource.
 */
export interface GroupPolicyArgs {
    /**
     * The IAM group to attach to the policy.
     */
    group: pulumi.Input<string>;
    /**
     * The name of the policy. If omitted, the provider will
     * assign a random, unique name.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The policy document. This is a JSON formatted string.
     */
    policy: pulumi.Input<string | PolicyDocument>;
}
