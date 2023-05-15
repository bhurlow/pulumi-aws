// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Data source for managing an AWS SSM Contact.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.ssmcontacts.getContact({
 *     arn: "arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias",
 * });
 * ```
 */
export function getContact(args: GetContactArgs, opts?: pulumi.InvokeOptions): Promise<GetContactResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws:ssmcontacts/getContact:getContact", {
        "arn": args.arn,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getContact.
 */
export interface GetContactArgs {
    /**
     * The Amazon Resource Name (ARN) of the contact or escalation plan.
     */
    arn: string;
    /**
     * Map of tags to assign to the resource.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getContact.
 */
export interface GetContactResult {
    /**
     * A unique and identifiable alias of the contact or escalation plan.
     */
    readonly alias: string;
    readonly arn: string;
    /**
     * Full friendly name of the contact or escalation plan.
     */
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Map of tags to assign to the resource.
     */
    readonly tags: {[key: string]: string};
    /**
     * The type of contact engaged. A single contact is type `PERSONAL` and an escalation plan is type `ESCALATION`.
     */
    readonly type: string;
}
/**
 * Data source for managing an AWS SSM Contact.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.ssmcontacts.getContact({
 *     arn: "arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias",
 * });
 * ```
 */
export function getContactOutput(args: GetContactOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetContactResult> {
    return pulumi.output(args).apply((a: any) => getContact(a, opts))
}

/**
 * A collection of arguments for invoking getContact.
 */
export interface GetContactOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the contact or escalation plan.
     */
    arn: pulumi.Input<string>;
    /**
     * Map of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
