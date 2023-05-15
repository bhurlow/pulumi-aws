// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource for associating accounts to existing Inspector instances.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.inspector2.MemberAssociation("example", {accountId: "123456789012"});
 * ```
 *
 * ## Import
 *
 * Amazon Inspector Member Association can be imported using the `account_id`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:inspector2/memberAssociation:MemberAssociation example 123456789012
 * ```
 */
export class MemberAssociation extends pulumi.CustomResource {
    /**
     * Get an existing MemberAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MemberAssociationState, opts?: pulumi.CustomResourceOptions): MemberAssociation {
        return new MemberAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:inspector2/memberAssociation:MemberAssociation';

    /**
     * Returns true if the given object is an instance of MemberAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MemberAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MemberAssociation.__pulumiType;
    }

    /**
     * ID of the account to associate
     */
    public readonly accountId!: pulumi.Output<string>;
    /**
     * Account ID of the delegated administrator account
     */
    public /*out*/ readonly delegatedAdminAccountId!: pulumi.Output<string>;
    /**
     * Status of the member relationship
     */
    public /*out*/ readonly relationshipStatus!: pulumi.Output<string>;
    /**
     * Date and time of the last update of the relationship
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a MemberAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MemberAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MemberAssociationArgs | MemberAssociationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MemberAssociationState | undefined;
            resourceInputs["accountId"] = state ? state.accountId : undefined;
            resourceInputs["delegatedAdminAccountId"] = state ? state.delegatedAdminAccountId : undefined;
            resourceInputs["relationshipStatus"] = state ? state.relationshipStatus : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as MemberAssociationArgs | undefined;
            if ((!args || args.accountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountId'");
            }
            resourceInputs["accountId"] = args ? args.accountId : undefined;
            resourceInputs["delegatedAdminAccountId"] = undefined /*out*/;
            resourceInputs["relationshipStatus"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MemberAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MemberAssociation resources.
 */
export interface MemberAssociationState {
    /**
     * ID of the account to associate
     */
    accountId?: pulumi.Input<string>;
    /**
     * Account ID of the delegated administrator account
     */
    delegatedAdminAccountId?: pulumi.Input<string>;
    /**
     * Status of the member relationship
     */
    relationshipStatus?: pulumi.Input<string>;
    /**
     * Date and time of the last update of the relationship
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MemberAssociation resource.
 */
export interface MemberAssociationArgs {
    /**
     * ID of the account to associate
     */
    accountId: pulumi.Input<string>;
}
