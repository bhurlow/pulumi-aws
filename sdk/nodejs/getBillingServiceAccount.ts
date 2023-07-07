// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get the Account ID of the [AWS Billing and Cost Management Service Account](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-getting-started.html#step-2) for the purpose of permitting in S3 bucket policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const main = aws.getBillingServiceAccount({});
 * const billingLogs = new aws.s3.BucketV2("billingLogs", {});
 * const billingLogsAcl = new aws.s3.BucketAclV2("billingLogsAcl", {
 *     bucket: billingLogs.id,
 *     acl: "private",
 * });
 * const allowBillingLoggingPolicyDocument = pulumi.all([main, billingLogs.arn, main, billingLogs.arn]).apply(([main, billingLogsArn, main1, billingLogsArn1]) => aws.iam.getPolicyDocumentOutput({
 *     statements: [
 *         {
 *             effect: "Allow",
 *             principals: [{
 *                 type: "AWS",
 *                 identifiers: [main.arn],
 *             }],
 *             actions: [
 *                 "s3:GetBucketAcl",
 *                 "s3:GetBucketPolicy",
 *             ],
 *             resources: [billingLogsArn],
 *         },
 *         {
 *             effect: "Allow",
 *             principals: [{
 *                 type: "AWS",
 *                 identifiers: [main1.arn],
 *             }],
 *             actions: ["s3:PutObject"],
 *             resources: [`${billingLogsArn1}/*`],
 *         },
 *     ],
 * }));
 * const allowBillingLoggingBucketPolicy = new aws.s3.BucketPolicy("allowBillingLoggingBucketPolicy", {
 *     bucket: billingLogs.id,
 *     policy: allowBillingLoggingPolicyDocument.apply(allowBillingLoggingPolicyDocument => allowBillingLoggingPolicyDocument.json),
 * });
 * ```
 */
export function getBillingServiceAccount(args?: GetBillingServiceAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetBillingServiceAccountResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws:index/getBillingServiceAccount:getBillingServiceAccount", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getBillingServiceAccount.
 */
export interface GetBillingServiceAccountArgs {
    /**
     * ID of the AWS billing service account.
     */
    id?: string;
}

/**
 * A collection of values returned by getBillingServiceAccount.
 */
export interface GetBillingServiceAccountResult {
    /**
     * ARN of the AWS billing service account.
     */
    readonly arn: string;
    /**
     * ID of the AWS billing service account.
     */
    readonly id: string;
}
/**
 * Use this data source to get the Account ID of the [AWS Billing and Cost Management Service Account](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-getting-started.html#step-2) for the purpose of permitting in S3 bucket policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const main = aws.getBillingServiceAccount({});
 * const billingLogs = new aws.s3.BucketV2("billingLogs", {});
 * const billingLogsAcl = new aws.s3.BucketAclV2("billingLogsAcl", {
 *     bucket: billingLogs.id,
 *     acl: "private",
 * });
 * const allowBillingLoggingPolicyDocument = pulumi.all([main, billingLogs.arn, main, billingLogs.arn]).apply(([main, billingLogsArn, main1, billingLogsArn1]) => aws.iam.getPolicyDocumentOutput({
 *     statements: [
 *         {
 *             effect: "Allow",
 *             principals: [{
 *                 type: "AWS",
 *                 identifiers: [main.arn],
 *             }],
 *             actions: [
 *                 "s3:GetBucketAcl",
 *                 "s3:GetBucketPolicy",
 *             ],
 *             resources: [billingLogsArn],
 *         },
 *         {
 *             effect: "Allow",
 *             principals: [{
 *                 type: "AWS",
 *                 identifiers: [main1.arn],
 *             }],
 *             actions: ["s3:PutObject"],
 *             resources: [`${billingLogsArn1}/*`],
 *         },
 *     ],
 * }));
 * const allowBillingLoggingBucketPolicy = new aws.s3.BucketPolicy("allowBillingLoggingBucketPolicy", {
 *     bucket: billingLogs.id,
 *     policy: allowBillingLoggingPolicyDocument.apply(allowBillingLoggingPolicyDocument => allowBillingLoggingPolicyDocument.json),
 * });
 * ```
 */
export function getBillingServiceAccountOutput(args?: GetBillingServiceAccountOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBillingServiceAccountResult> {
    return pulumi.output(args).apply((a: any) => getBillingServiceAccount(a, opts))
}

/**
 * A collection of arguments for invoking getBillingServiceAccount.
 */
export interface GetBillingServiceAccountOutputArgs {
    /**
     * ID of the AWS billing service account.
     */
    id?: pulumi.Input<string>;
}
