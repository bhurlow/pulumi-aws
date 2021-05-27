// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get the Account ID of the [AWS Elastic Load Balancing Service Account](http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-access-logs.html#attach-bucket-policy)
 * in a given region for the purpose of permitting in S3 bucket policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const main = aws.elb.getServiceAccount({});
 * const elbLogs = new aws.s3.Bucket("elbLogs", {
 *     acl: "private",
 *     policy: main.then(main => `{
 *   "Id": "Policy",
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": [
 *         "s3:PutObject"
 *       ],
 *       "Effect": "Allow",
 *       "Resource": "arn:aws:s3:::my-elb-tf-test-bucket/AWSLogs/*",
 *       "Principal": {
 *         "AWS": [
 *           "${main.arn}"
 *         ]
 *       }
 *     }
 *   ]
 * }
 * `),
 * });
 * const bar = new aws.elb.LoadBalancer("bar", {
 *     availabilityZones: ["us-west-2a"],
 *     accessLogs: {
 *         bucket: elbLogs.bucket,
 *         interval: 5,
 *     },
 *     listeners: [{
 *         instancePort: 8000,
 *         instanceProtocol: "http",
 *         lbPort: 80,
 *         lbProtocol: "http",
 *     }],
 * });
 * ```
 */
export function getServiceAccount(args?: GetServiceAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceAccountResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:elb/getServiceAccount:getServiceAccount", {
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceAccount.
 */
export interface GetServiceAccountArgs {
    /**
     * Name of the region whose AWS ELB account ID is desired.
     * Defaults to the region from the AWS provider configuration.
     */
    region?: string;
}

/**
 * A collection of values returned by getServiceAccount.
 */
export interface GetServiceAccountResult {
    /**
     * The ARN of the AWS ELB service account in the selected region.
     */
    readonly arn: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly region?: string;
}
