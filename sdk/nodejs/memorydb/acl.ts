// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a MemoryDB ACL.
 *
 * More information about users and ACL-s can be found in the [MemoryDB User Guide](https://docs.aws.amazon.com/memorydb/latest/devguide/clusters.acls.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.memorydb.Acl("example", {userNames: [
 *     "my-user-1",
 *     "my-user-2",
 * ]});
 * ```
 *
 * ## Import
 *
 * Use the `name` to import an ACL. For example
 *
 * ```sh
 *  $ pulumi import aws:memorydb/acl:Acl example my-acl
 * ```
 */
export class Acl extends pulumi.CustomResource {
    /**
     * Get an existing Acl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AclState, opts?: pulumi.CustomResourceOptions): Acl {
        return new Acl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:memorydb/acl:Acl';

    /**
     * Returns true if the given object is an instance of Acl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Acl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Acl.__pulumiType;
    }

    /**
     * The ARN of the ACL.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The minimum engine version supported by the ACL.
     */
    public /*out*/ readonly minimumEngineVersion!: pulumi.Output<string>;
    /**
     * Name of the ACL. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Set of MemoryDB user names to be included in this ACL.
     */
    public readonly userNames!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Acl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AclArgs | AclState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AclState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["minimumEngineVersion"] = state ? state.minimumEngineVersion : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namePrefix"] = state ? state.namePrefix : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["userNames"] = state ? state.userNames : undefined;
        } else {
            const args = argsOrState as AclArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userNames"] = args ? args.userNames : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["minimumEngineVersion"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Acl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Acl resources.
 */
export interface AclState {
    /**
     * The ARN of the ACL.
     */
    arn?: pulumi.Input<string>;
    /**
     * The minimum engine version supported by the ACL.
     */
    minimumEngineVersion?: pulumi.Input<string>;
    /**
     * Name of the ACL. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Set of MemoryDB user names to be included in this ACL.
     */
    userNames?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Acl resource.
 */
export interface AclArgs {
    /**
     * Name of the ACL. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Set of MemoryDB user names to be included in this ACL.
     */
    userNames?: pulumi.Input<pulumi.Input<string>[]>;
}
