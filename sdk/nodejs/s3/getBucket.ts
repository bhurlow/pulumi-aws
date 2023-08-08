// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides details about a specific S3 bucket.
 *
 * This resource may prove useful when setting up a Route53 record, or an origin for a CloudFront
 * Distribution.
 *
 * ## Example Usage
 * ### Route53 Record
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const selected = aws.s3.getBucket({
 *     bucket: "bucket.test.com",
 * });
 * const testZone = aws.route53.getZone({
 *     name: "test.com.",
 * });
 * const example = new aws.route53.Record("example", {
 *     zoneId: testZone.then(testZone => testZone.id),
 *     name: "bucket",
 *     type: "A",
 *     aliases: [{
 *         name: selected.then(selected => selected.websiteDomain),
 *         zoneId: selected.then(selected => selected.hostedZoneId),
 *     }],
 * });
 * ```
 * ### CloudFront Origin
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const selected = aws.s3.getBucket({
 *     bucket: "a-test-bucket",
 * });
 * const test = new aws.cloudfront.Distribution("test", {origins: [{
 *     domainName: selected.then(selected => selected.bucketDomainName),
 *     originId: "s3-selected-bucket",
 * }]});
 * ```
 */
export function getBucket(args: GetBucketArgs, opts?: pulumi.InvokeOptions): Promise<GetBucketResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws:s3/getBucket:getBucket", {
        "bucket": args.bucket,
    }, opts);
}

/**
 * A collection of arguments for invoking getBucket.
 */
export interface GetBucketArgs {
    /**
     * Name of the bucket
     */
    bucket: string;
}

/**
 * A collection of values returned by getBucket.
 */
export interface GetBucketResult {
    /**
     * ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
     */
    readonly arn: string;
    readonly bucket: string;
    /**
     * Bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
     */
    readonly bucketDomainName: string;
    /**
     * The bucket region-specific domain name. The bucket domain name including the region name. Please refer to the [S3 endpoints reference](https://docs.aws.amazon.com/general/latest/gr/s3.html#s3_region) for format. Note: AWS CloudFront allows specifying an S3 region-specific endpoint when creating an S3 origin. This will prevent redirect issues from CloudFront to the S3 Origin URL. For more information, see the [Virtual Hosted-Style Requests for Other Regions](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html#deprecated-global-endpoint) section in the AWS S3 User Guide.
     */
    readonly bucketRegionalDomainName: string;
    /**
     * The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
     */
    readonly hostedZoneId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * AWS region this bucket resides in.
     */
    readonly region: string;
    /**
     * Domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
     */
    readonly websiteDomain: string;
    /**
     * Website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
     */
    readonly websiteEndpoint: string;
}
/**
 * Provides details about a specific S3 bucket.
 *
 * This resource may prove useful when setting up a Route53 record, or an origin for a CloudFront
 * Distribution.
 *
 * ## Example Usage
 * ### Route53 Record
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const selected = aws.s3.getBucket({
 *     bucket: "bucket.test.com",
 * });
 * const testZone = aws.route53.getZone({
 *     name: "test.com.",
 * });
 * const example = new aws.route53.Record("example", {
 *     zoneId: testZone.then(testZone => testZone.id),
 *     name: "bucket",
 *     type: "A",
 *     aliases: [{
 *         name: selected.then(selected => selected.websiteDomain),
 *         zoneId: selected.then(selected => selected.hostedZoneId),
 *     }],
 * });
 * ```
 * ### CloudFront Origin
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const selected = aws.s3.getBucket({
 *     bucket: "a-test-bucket",
 * });
 * const test = new aws.cloudfront.Distribution("test", {origins: [{
 *     domainName: selected.then(selected => selected.bucketDomainName),
 *     originId: "s3-selected-bucket",
 * }]});
 * ```
 */
export function getBucketOutput(args: GetBucketOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBucketResult> {
    return pulumi.output(args).apply((a: any) => getBucket(a, opts))
}

/**
 * A collection of arguments for invoking getBucket.
 */
export interface GetBucketOutputArgs {
    /**
     * Name of the bucket
     */
    bucket: pulumi.Input<string>;
}
