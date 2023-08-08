// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Data source for managing an AWS SESv2 (Simple Email V2) Email Identity.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.sesv2.getEmailIdentity({
 *     emailIdentity: "example.com",
 * });
 * ```
 */
export function getEmailIdentity(args: GetEmailIdentityArgs, opts?: pulumi.InvokeOptions): Promise<GetEmailIdentityResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws:sesv2/getEmailIdentity:getEmailIdentity", {
        "emailIdentity": args.emailIdentity,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getEmailIdentity.
 */
export interface GetEmailIdentityArgs {
    /**
     * The name of the email identity.
     */
    emailIdentity: string;
    /**
     * Key-value mapping of resource tags.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getEmailIdentity.
 */
export interface GetEmailIdentityResult {
    /**
     * ARN of the Email Identity.
     */
    readonly arn: string;
    readonly configurationSetName: string;
    /**
     * A list of objects that contains at most one element with information about the private key and selector that you want to use to configure DKIM for the identity for Bring Your Own DKIM (BYODKIM) for the identity, or, configures the key length to be used for Easy DKIM.
     */
    readonly dkimSigningAttributes: outputs.sesv2.GetEmailIdentityDkimSigningAttribute[];
    readonly emailIdentity: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The email identity type. Valid values: `EMAIL_ADDRESS`, `DOMAIN`.
     */
    readonly identityType: string;
    /**
     * Key-value mapping of resource tags.
     */
    readonly tags: {[key: string]: string};
    /**
     * Specifies whether or not the identity is verified.
     */
    readonly verifiedForSendingStatus: boolean;
}
/**
 * Data source for managing an AWS SESv2 (Simple Email V2) Email Identity.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.sesv2.getEmailIdentity({
 *     emailIdentity: "example.com",
 * });
 * ```
 */
export function getEmailIdentityOutput(args: GetEmailIdentityOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEmailIdentityResult> {
    return pulumi.output(args).apply((a: any) => getEmailIdentity(a, opts))
}

/**
 * A collection of arguments for invoking getEmailIdentity.
 */
export interface GetEmailIdentityOutputArgs {
    /**
     * The name of the email identity.
     */
    emailIdentity: pulumi.Input<string>;
    /**
     * Key-value mapping of resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
