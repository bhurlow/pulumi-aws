// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage an [Amazon Macie Member](https://docs.aws.amazon.com/macie/latest/APIReference/members-id.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleAccount = new aws.macie2.Account("exampleAccount", {});
 * const exampleMember = new aws.macie2.Member("exampleMember", {
 *     accountId: "AWS ACCOUNT ID",
 *     email: "EMAIL",
 *     invite: true,
 *     invitationMessage: "Message of the invitation",
 *     invitationDisableEmailNotification: true,
 * }, {
 *     dependsOn: [exampleAccount],
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import `aws_macie2_member` using the account ID of the member account. For example:
 *
 * ```sh
 *  $ pulumi import aws:macie2/member:Member example 123456789012
 * ```
 */
export class Member extends pulumi.CustomResource {
    /**
     * Get an existing Member resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MemberState, opts?: pulumi.CustomResourceOptions): Member {
        return new Member(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:macie2/member:Member';

    /**
     * Returns true if the given object is an instance of Member.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Member {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Member.__pulumiType;
    }

    /**
     * The AWS account ID for the account.
     */
    public readonly accountId!: pulumi.Output<string>;
    /**
     * The AWS account ID for the administrator account.
     */
    public /*out*/ readonly administratorAccountId!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the account.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The email address for the account.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * Specifies whether to send an email notification to the root user of each account that the invitation will be sent to. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. To send an email notification to the root user of each account, set this value to `true`.
     */
    public readonly invitationDisableEmailNotification!: pulumi.Output<boolean | undefined>;
    /**
     * A custom message to include in the invitation. Amazon Macie adds this message to the standard content that it sends for an invitation.
     */
    public readonly invitationMessage!: pulumi.Output<string | undefined>;
    /**
     * Send an invitation to a member
     */
    public readonly invite!: pulumi.Output<boolean>;
    /**
     * The date and time, in UTC and extended RFC 3339 format, when an Amazon Macie membership invitation was last sent to the account. This value is null if a Macie invitation hasn't been sent to the account.
     */
    public /*out*/ readonly invitedAt!: pulumi.Output<string>;
    public /*out*/ readonly masterAccountId!: pulumi.Output<string>;
    /**
     * The current status of the relationship between the account and the administrator account.
     */
    public /*out*/ readonly relationshipStatus!: pulumi.Output<string>;
    /**
     * Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to `ENABLED`. Valid values are `ENABLED` or `PAUSED`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * A map of key-value pairs that specifies the tags to associate with the account in Amazon Macie.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The date and time, in UTC and extended RFC 3339 format, of the most recent change to the status of the relationship between the account and the administrator account.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a Member resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MemberArgs | MemberState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MemberState | undefined;
            resourceInputs["accountId"] = state ? state.accountId : undefined;
            resourceInputs["administratorAccountId"] = state ? state.administratorAccountId : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["invitationDisableEmailNotification"] = state ? state.invitationDisableEmailNotification : undefined;
            resourceInputs["invitationMessage"] = state ? state.invitationMessage : undefined;
            resourceInputs["invite"] = state ? state.invite : undefined;
            resourceInputs["invitedAt"] = state ? state.invitedAt : undefined;
            resourceInputs["masterAccountId"] = state ? state.masterAccountId : undefined;
            resourceInputs["relationshipStatus"] = state ? state.relationshipStatus : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as MemberArgs | undefined;
            if ((!args || args.accountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountId'");
            }
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            resourceInputs["accountId"] = args ? args.accountId : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["invitationDisableEmailNotification"] = args ? args.invitationDisableEmailNotification : undefined;
            resourceInputs["invitationMessage"] = args ? args.invitationMessage : undefined;
            resourceInputs["invite"] = args ? args.invite : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["administratorAccountId"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["invitedAt"] = undefined /*out*/;
            resourceInputs["masterAccountId"] = undefined /*out*/;
            resourceInputs["relationshipStatus"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Member.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Member resources.
 */
export interface MemberState {
    /**
     * The AWS account ID for the account.
     */
    accountId?: pulumi.Input<string>;
    /**
     * The AWS account ID for the administrator account.
     */
    administratorAccountId?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the account.
     */
    arn?: pulumi.Input<string>;
    /**
     * The email address for the account.
     */
    email?: pulumi.Input<string>;
    /**
     * Specifies whether to send an email notification to the root user of each account that the invitation will be sent to. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. To send an email notification to the root user of each account, set this value to `true`.
     */
    invitationDisableEmailNotification?: pulumi.Input<boolean>;
    /**
     * A custom message to include in the invitation. Amazon Macie adds this message to the standard content that it sends for an invitation.
     */
    invitationMessage?: pulumi.Input<string>;
    /**
     * Send an invitation to a member
     */
    invite?: pulumi.Input<boolean>;
    /**
     * The date and time, in UTC and extended RFC 3339 format, when an Amazon Macie membership invitation was last sent to the account. This value is null if a Macie invitation hasn't been sent to the account.
     */
    invitedAt?: pulumi.Input<string>;
    masterAccountId?: pulumi.Input<string>;
    /**
     * The current status of the relationship between the account and the administrator account.
     */
    relationshipStatus?: pulumi.Input<string>;
    /**
     * Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to `ENABLED`. Valid values are `ENABLED` or `PAUSED`.
     */
    status?: pulumi.Input<string>;
    /**
     * A map of key-value pairs that specifies the tags to associate with the account in Amazon Macie.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The date and time, in UTC and extended RFC 3339 format, of the most recent change to the status of the relationship between the account and the administrator account.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Member resource.
 */
export interface MemberArgs {
    /**
     * The AWS account ID for the account.
     */
    accountId: pulumi.Input<string>;
    /**
     * The email address for the account.
     */
    email: pulumi.Input<string>;
    /**
     * Specifies whether to send an email notification to the root user of each account that the invitation will be sent to. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. To send an email notification to the root user of each account, set this value to `true`.
     */
    invitationDisableEmailNotification?: pulumi.Input<boolean>;
    /**
     * A custom message to include in the invitation. Amazon Macie adds this message to the standard content that it sends for an invitation.
     */
    invitationMessage?: pulumi.Input<string>;
    /**
     * Send an invitation to a member
     */
    invite?: pulumi.Input<boolean>;
    /**
     * Specifies the status for the account. To enable Amazon Macie and start all Macie activities for the account, set this value to `ENABLED`. Valid values are `ENABLED` or `PAUSED`.
     */
    status?: pulumi.Input<string>;
    /**
     * A map of key-value pairs that specifies the tags to associate with the account in Amazon Macie.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
