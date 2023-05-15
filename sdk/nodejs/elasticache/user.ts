// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides an ElastiCache user resource.
 *
 * > **Note:** All arguments including the username and passwords will be stored in the raw state as plain-text.
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.elasticache.User("test", {
 *     accessString: "on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember",
 *     engine: "REDIS",
 *     passwords: ["password123456789"],
 *     userId: "testUserId",
 *     userName: "testUserName",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.elasticache.User("test", {
 *     accessString: "on ~* +@all",
 *     authenticationMode: {
 *         type: "iam",
 *     },
 *     engine: "REDIS",
 *     userId: "testUserId",
 *     userName: "testUserName",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.elasticache.User("test", {
 *     accessString: "on ~* +@all",
 *     authenticationMode: {
 *         passwords: [
 *             "password1",
 *             "password2",
 *         ],
 *         type: "password",
 *     },
 *     engine: "REDIS",
 *     userId: "testUserId",
 *     userName: "testUserName",
 * });
 * ```
 *
 * ## Import
 *
 * ElastiCache users can be imported using the `user_id`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:elasticache/user:User my_user userId1
 * ```
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:elasticache/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * Access permissions string used for this user. See [Specifying Permissions Using an Access String](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.RBAC.html#Access-string) for more details.
     */
    public readonly accessString!: pulumi.Output<string>;
    /**
     * The ARN of the created ElastiCache User.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Denotes the user's authentication properties. Detailed below.
     */
    public readonly authenticationMode!: pulumi.Output<outputs.elasticache.UserAuthenticationMode>;
    /**
     * The current supported value is `REDIS`.
     */
    public readonly engine!: pulumi.Output<string>;
    /**
     * Indicates a password is not required for this user.
     */
    public readonly noPasswordRequired!: pulumi.Output<boolean | undefined>;
    /**
     * Passwords used for this user. You can create up to two passwords for each user.
     */
    public readonly passwords!: pulumi.Output<string[] | undefined>;
    /**
     * A list of tags to be added to this resource. A tag is a key-value pair.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The ID of the user.
     */
    public readonly userId!: pulumi.Output<string>;
    /**
     * The username of the user.
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            resourceInputs["accessString"] = state ? state.accessString : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["authenticationMode"] = state ? state.authenticationMode : undefined;
            resourceInputs["engine"] = state ? state.engine : undefined;
            resourceInputs["noPasswordRequired"] = state ? state.noPasswordRequired : undefined;
            resourceInputs["passwords"] = state ? state.passwords : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.accessString === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessString'");
            }
            if ((!args || args.engine === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engine'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["accessString"] = args ? args.accessString : undefined;
            resourceInputs["authenticationMode"] = args ? args.authenticationMode : undefined;
            resourceInputs["engine"] = args ? args.engine : undefined;
            resourceInputs["noPasswordRequired"] = args ? args.noPasswordRequired : undefined;
            resourceInputs["passwords"] = args?.passwords ? pulumi.secret(args.passwords) : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["passwords"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * Access permissions string used for this user. See [Specifying Permissions Using an Access String](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.RBAC.html#Access-string) for more details.
     */
    accessString?: pulumi.Input<string>;
    /**
     * The ARN of the created ElastiCache User.
     */
    arn?: pulumi.Input<string>;
    /**
     * Denotes the user's authentication properties. Detailed below.
     */
    authenticationMode?: pulumi.Input<inputs.elasticache.UserAuthenticationMode>;
    /**
     * The current supported value is `REDIS`.
     */
    engine?: pulumi.Input<string>;
    /**
     * Indicates a password is not required for this user.
     */
    noPasswordRequired?: pulumi.Input<boolean>;
    /**
     * Passwords used for this user. You can create up to two passwords for each user.
     */
    passwords?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of tags to be added to this resource. A tag is a key-value pair.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the user.
     */
    userId?: pulumi.Input<string>;
    /**
     * The username of the user.
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * Access permissions string used for this user. See [Specifying Permissions Using an Access String](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.RBAC.html#Access-string) for more details.
     */
    accessString: pulumi.Input<string>;
    /**
     * Denotes the user's authentication properties. Detailed below.
     */
    authenticationMode?: pulumi.Input<inputs.elasticache.UserAuthenticationMode>;
    /**
     * The current supported value is `REDIS`.
     */
    engine: pulumi.Input<string>;
    /**
     * Indicates a password is not required for this user.
     */
    noPasswordRequired?: pulumi.Input<boolean>;
    /**
     * Passwords used for this user. You can create up to two passwords for each user.
     */
    passwords?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of tags to be added to this resource. A tag is a key-value pair.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the user.
     */
    userId: pulumi.Input<string>;
    /**
     * The username of the user.
     */
    userName: pulumi.Input<string>;
}
