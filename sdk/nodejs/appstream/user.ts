// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an AppStream user.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.appstream.User("example", {
 *     authenticationType: "USERPOOL",
 *     firstName: "FIRST NAME",
 *     lastName: "LAST NAME",
 *     userName: "EMAIL ADDRESS",
 * });
 * ```
 *
 * ## Import
 *
 * `aws_appstream_user` can be imported using the `user_name` and `authentication_type` separated by a slash (`/`), e.g.,
 *
 * ```sh
 *  $ pulumi import aws:appstream/user:User example UserName/AuthenticationType
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
    public static readonly __pulumiType = 'aws:appstream/user:User';

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
     * ARN of the appstream user.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Authentication type for the user. You must specify USERPOOL. Valid values: `API`, `SAML`, `USERPOOL`
     */
    public readonly authenticationType!: pulumi.Output<string>;
    /**
     * Date and time, in UTC and extended RFC 3339 format, when the user was created.
     */
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    /**
     * Specifies whether the user in the user pool is enabled.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * First name, or given name, of the user.
     */
    public readonly firstName!: pulumi.Output<string | undefined>;
    /**
     * Last name, or surname, of the user.
     */
    public readonly lastName!: pulumi.Output<string | undefined>;
    /**
     * Send an email notification.
     */
    public readonly sendEmailNotification!: pulumi.Output<boolean | undefined>;
    /**
     * Email address of the user.
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["authenticationType"] = state ? state.authenticationType : undefined;
            inputs["createdTime"] = state ? state.createdTime : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["firstName"] = state ? state.firstName : undefined;
            inputs["lastName"] = state ? state.lastName : undefined;
            inputs["sendEmailNotification"] = state ? state.sendEmailNotification : undefined;
            inputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.authenticationType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authenticationType'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            inputs["authenticationType"] = args ? args.authenticationType : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["firstName"] = args ? args.firstName : undefined;
            inputs["lastName"] = args ? args.lastName : undefined;
            inputs["sendEmailNotification"] = args ? args.sendEmailNotification : undefined;
            inputs["userName"] = args ? args.userName : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["createdTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(User.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * ARN of the appstream user.
     */
    arn?: pulumi.Input<string>;
    /**
     * Authentication type for the user. You must specify USERPOOL. Valid values: `API`, `SAML`, `USERPOOL`
     */
    authenticationType?: pulumi.Input<string>;
    /**
     * Date and time, in UTC and extended RFC 3339 format, when the user was created.
     */
    createdTime?: pulumi.Input<string>;
    /**
     * Specifies whether the user in the user pool is enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * First name, or given name, of the user.
     */
    firstName?: pulumi.Input<string>;
    /**
     * Last name, or surname, of the user.
     */
    lastName?: pulumi.Input<string>;
    /**
     * Send an email notification.
     */
    sendEmailNotification?: pulumi.Input<boolean>;
    /**
     * Email address of the user.
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * Authentication type for the user. You must specify USERPOOL. Valid values: `API`, `SAML`, `USERPOOL`
     */
    authenticationType: pulumi.Input<string>;
    /**
     * Specifies whether the user in the user pool is enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * First name, or given name, of the user.
     */
    firstName?: pulumi.Input<string>;
    /**
     * Last name, or surname, of the user.
     */
    lastName?: pulumi.Input<string>;
    /**
     * Send an email notification.
     */
    sendEmailNotification?: pulumi.Input<boolean>;
    /**
     * Email address of the user.
     */
    userName: pulumi.Input<string>;
}