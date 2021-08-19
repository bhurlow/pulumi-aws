// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an ElastiCache user group resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testUser = new aws.elasticache.User("testUser", {
 *     userId: "testUserId",
 *     userName: "default",
 *     accessString: "on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember",
 *     engine: "REDIS",
 *     passwords: ["password123456789"],
 * });
 * const testUserGroup = new aws.elasticache.UserGroup("testUserGroup", {
 *     engine: "REDIS",
 *     userGroupId: "userGroupId",
 *     userIds: [testUser.userId],
 * });
 * ```
 *
 * ## Import
 *
 * ElastiCache user groups can be imported using the `user_group_id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:elasticache/userGroup:UserGroup my_user_group userGoupId1
 * ```
 */
export class UserGroup extends pulumi.CustomResource {
    /**
     * Get an existing UserGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserGroupState, opts?: pulumi.CustomResourceOptions): UserGroup {
        return new UserGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:elasticache/userGroup:UserGroup';

    /**
     * Returns true if the given object is an instance of UserGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserGroup.__pulumiType;
    }

    public readonly arn!: pulumi.Output<string>;
    /**
     * The current supported value is `REDIS`.
     */
    public readonly engine!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The ID of the user group.
     */
    public readonly userGroupId!: pulumi.Output<string>;
    /**
     * The list of user IDs that belong to the user group.
     */
    public readonly userIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a UserGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserGroupArgs | UserGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserGroupState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["engine"] = state ? state.engine : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["userGroupId"] = state ? state.userGroupId : undefined;
            inputs["userIds"] = state ? state.userIds : undefined;
        } else {
            const args = argsOrState as UserGroupArgs | undefined;
            if ((!args || args.engine === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engine'");
            }
            if ((!args || args.userGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userGroupId'");
            }
            inputs["arn"] = args ? args.arn : undefined;
            inputs["engine"] = args ? args.engine : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["userGroupId"] = args ? args.userGroupId : undefined;
            inputs["userIds"] = args ? args.userIds : undefined;
            inputs["tagsAll"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(UserGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserGroup resources.
 */
export interface UserGroupState {
    arn?: pulumi.Input<string>;
    /**
     * The current supported value is `REDIS`.
     */
    engine?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the user group.
     */
    userGroupId?: pulumi.Input<string>;
    /**
     * The list of user IDs that belong to the user group.
     */
    userIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a UserGroup resource.
 */
export interface UserGroupArgs {
    arn?: pulumi.Input<string>;
    /**
     * The current supported value is `REDIS`.
     */
    engine: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the user group.
     */
    userGroupId: pulumi.Input<string>;
    /**
     * The list of user IDs that belong to the user group.
     */
    userIds?: pulumi.Input<pulumi.Input<string>[]>;
}
