// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

import {PolicyDocument} from "../iam";
import {RoutingRule} from "./index";

/**
 * Provides a S3 bucket resource.
 *
 * > This functionality is for managing S3 in an AWS Partition. To manage [S3 on Outposts](https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html), see the `aws.s3control.Bucket` resource.
 *
 * > **NOTE:** This resource might not work well if using an alternative s3-compatible provider. Please use `aws.s3.BucketV2` instead.
 *
 * ## Example Usage
 * ### Private Bucket w/ Tags
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const bucket = new aws.s3.Bucket("bucket", {
 *     acl: "private",
 *     tags: {
 *         Environment: "Dev",
 *         Name: "My bucket",
 *     },
 * });
 * ```
 * ### Static Website Hosting
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as fs from "fs";
 *
 * const bucket = new aws.s3.Bucket("bucket", {
 *     acl: "public-read",
 *     policy: fs.readFileSync("policy.json", "utf8"),
 *     website: {
 *         indexDocument: "index.html",
 *         errorDocument: "error.html",
 *         routingRules: `[{
 *     "Condition": {
 *         "KeyPrefixEquals": "docs/"
 *     },
 *     "Redirect": {
 *         "ReplaceKeyPrefixWith": "documents/"
 *     }
 * }]
 * `,
 *     },
 * });
 * ```
 * ### Using CORS
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const bucket = new aws.s3.Bucket("bucket", {
 *     acl: "public-read",
 *     corsRules: [{
 *         allowedHeaders: ["*"],
 *         allowedMethods: [
 *             "PUT",
 *             "POST",
 *         ],
 *         allowedOrigins: ["https://s3-website-test.mydomain.com"],
 *         exposeHeaders: ["ETag"],
 *         maxAgeSeconds: 3000,
 *     }],
 * });
 * ```
 * ### Using versioning
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const bucket = new aws.s3.Bucket("bucket", {
 *     acl: "private",
 *     versioning: {
 *         enabled: true,
 *     },
 * });
 * ```
 * ### Enable Logging
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const logBucket = new aws.s3.Bucket("logBucket", {acl: "log-delivery-write"});
 * const bucket = new aws.s3.Bucket("bucket", {
 *     acl: "private",
 *     loggings: [{
 *         targetBucket: logBucket.id,
 *         targetPrefix: "log/",
 *     }],
 * });
 * ```
 * ### Using object lifecycle
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const bucket = new aws.s3.Bucket("bucket", {
 *     acl: "private",
 *     lifecycleRules: [
 *         {
 *             enabled: true,
 *             expiration: {
 *                 days: 90,
 *             },
 *             id: "log",
 *             prefix: "log/",
 *             tags: {
 *                 autoclean: "true",
 *                 rule: "log",
 *             },
 *             transitions: [
 *                 {
 *                     days: 30,
 *                     storageClass: "STANDARD_IA",
 *                 },
 *                 {
 *                     days: 60,
 *                     storageClass: "GLACIER",
 *                 },
 *             ],
 *         },
 *         {
 *             enabled: true,
 *             expiration: {
 *                 date: "2016-01-12",
 *             },
 *             id: "tmp",
 *             prefix: "tmp/",
 *         },
 *     ],
 * });
 * const versioningBucket = new aws.s3.Bucket("versioningBucket", {
 *     acl: "private",
 *     lifecycleRules: [{
 *         enabled: true,
 *         noncurrentVersionExpiration: {
 *             days: 90,
 *         },
 *         noncurrentVersionTransitions: [
 *             {
 *                 days: 30,
 *                 storageClass: "STANDARD_IA",
 *             },
 *             {
 *                 days: 60,
 *                 storageClass: "GLACIER",
 *             },
 *         ],
 *         prefix: "config/",
 *     }],
 *     versioning: {
 *         enabled: true,
 *     },
 * });
 * ```
 * ### Using replication configuration
 *
 * > **NOTE:** See the `aws.s3.BucketReplicationConfig` resource to support bi-directional replication configuration and additional features.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const central = new aws.Provider("central", {region: "eu-central-1"});
 * const replicationRole = new aws.iam.Role("replicationRole", {assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "sts:AssumeRole",
 *       "Principal": {
 *         "Service": "s3.amazonaws.com"
 *       },
 *       "Effect": "Allow",
 *       "Sid": ""
 *     }
 *   ]
 * }
 * `});
 * const destination = new aws.s3.Bucket("destination", {versioning: {
 *     enabled: true,
 * }});
 * const source = new aws.s3.Bucket("source", {
 *     acl: "private",
 *     versioning: {
 *         enabled: true,
 *     },
 *     replicationConfiguration: {
 *         role: replicationRole.arn,
 *         rules: [{
 *             id: "foobar",
 *             status: "Enabled",
 *             filter: {
 *                 tags: {},
 *             },
 *             destination: {
 *                 bucket: destination.arn,
 *                 storageClass: "STANDARD",
 *                 replicationTime: {
 *                     status: "Enabled",
 *                     minutes: 15,
 *                 },
 *                 metrics: {
 *                     status: "Enabled",
 *                     minutes: 15,
 *                 },
 *             },
 *         }],
 *     },
 * }, {
 *     provider: aws.central,
 * });
 * const replicationPolicy = new aws.iam.Policy("replicationPolicy", {policy: pulumi.interpolate`{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": [
 *         "s3:GetReplicationConfiguration",
 *         "s3:ListBucket"
 *       ],
 *       "Effect": "Allow",
 *       "Resource": [
 *         "${source.arn}"
 *       ]
 *     },
 *     {
 *       "Action": [
 *         "s3:GetObjectVersionForReplication",
 *         "s3:GetObjectVersionAcl",
 *          "s3:GetObjectVersionTagging"
 *       ],
 *       "Effect": "Allow",
 *       "Resource": [
 *         "${source.arn}/*"
 *       ]
 *     },
 *     {
 *       "Action": [
 *         "s3:ReplicateObject",
 *         "s3:ReplicateDelete",
 *         "s3:ReplicateTags"
 *       ],
 *       "Effect": "Allow",
 *       "Resource": "${destination.arn}/*"
 *     }
 *   ]
 * }
 * `});
 * const replicationRolePolicyAttachment = new aws.iam.RolePolicyAttachment("replicationRolePolicyAttachment", {
 *     role: replicationRole.name,
 *     policyArn: replicationPolicy.arn,
 * });
 * ```
 * ### Enable Default Server Side Encryption
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const mykey = new aws.kms.Key("mykey", {
 *     description: "This key is used to encrypt bucket objects",
 *     deletionWindowInDays: 10,
 * });
 * const mybucket = new aws.s3.Bucket("mybucket", {serverSideEncryptionConfiguration: {
 *     rule: {
 *         applyServerSideEncryptionByDefault: {
 *             kmsMasterKeyId: mykey.arn,
 *             sseAlgorithm: "aws:kms",
 *         },
 *     },
 * }});
 * ```
 * ### Using ACL policy grants
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const currentUser = aws.s3.getCanonicalUserId({});
 * const bucket = new aws.s3.Bucket("bucket", {grants: [
 *     {
 *         id: currentUser.then(currentUser => currentUser.id),
 *         type: "CanonicalUser",
 *         permissions: ["FULL_CONTROL"],
 *     },
 *     {
 *         type: "Group",
 *         permissions: [
 *             "READ_ACP",
 *             "WRITE",
 *         ],
 *         uri: "http://acs.amazonaws.com/groups/s3/LogDelivery",
 *     },
 * ]});
 * ```
 *
 * ## Import
 *
 * S3 bucket can be imported using the `bucket`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:s3/bucket:Bucket bucket bucket-name
 * ```
 *  The `policy` argument is not imported and will be deprecated in a future version of the provider. Use the `aws_s3_bucket_policy` resource to manage the S3 Bucket Policy instead.
 */
export class Bucket extends pulumi.CustomResource {
    /**
     * Get an existing Bucket resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketState, opts?: pulumi.CustomResourceOptions): Bucket {
        return new Bucket(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:s3/bucket:Bucket';

    /**
     * Returns true if the given object is an instance of Bucket.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Bucket {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Bucket.__pulumiType;
    }

    /**
     * Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
     */
    public readonly accelerationStatus!: pulumi.Output<string>;
    /**
     * The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Valid values are `private`, `public-read`, `public-read-write`, `aws-exec-read`, `authenticated-read`, and `log-delivery-write`. Defaults to `private`.  Conflicts with `grant`.
     */
    public readonly acl!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
     */
    public readonly arn!: pulumi.Output<string>;
    /**
     * The name of the bucket. If omitted, this provider will assign a random, unique name. Must be lowercase and less than or equal to 63 characters in length. A full list of bucket naming rules [may be found here](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html).
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
     */
    public /*out*/ readonly bucketDomainName!: pulumi.Output<string>;
    /**
     * Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`. Must be lowercase and less than or equal to 37 characters in length. A full list of bucket naming rules [may be found here](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html).
     */
    public readonly bucketPrefix!: pulumi.Output<string | undefined>;
    /**
     * The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.
     */
    public /*out*/ readonly bucketRegionalDomainName!: pulumi.Output<string>;
    /**
     * A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
     */
    public readonly corsRules!: pulumi.Output<outputs.s3.BucketCorsRule[] | undefined>;
    /**
     * A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
     */
    public readonly forceDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * An [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl) (documented below). Conflicts with `acl`.
     */
    public readonly grants!: pulumi.Output<outputs.s3.BucketGrant[] | undefined>;
    /**
     * The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
     */
    public readonly hostedZoneId!: pulumi.Output<string>;
    /**
     * A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
     */
    public readonly lifecycleRules!: pulumi.Output<outputs.s3.BucketLifecycleRule[] | undefined>;
    /**
     * A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
     */
    public readonly loggings!: pulumi.Output<outputs.s3.BucketLogging[] | undefined>;
    /**
     * A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
     *
     * > **NOTE:** You cannot use `accelerationStatus` in `cn-north-1` or `us-gov-west-1`
     */
    public readonly objectLockConfiguration!: pulumi.Output<outputs.s3.BucketObjectLockConfiguration | undefined>;
    /**
     * A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), this provider may view the policy as constantly changing in a `pulumi preview`. In this case, please make sure you use the verbose/specific version of the policy.
     */
    public readonly policy!: pulumi.Output<string | undefined>;
    /**
     * The AWS region this bucket resides in.
     */
    public /*out*/ readonly region!: pulumi.Output<string>;
    /**
     * A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
     */
    public readonly replicationConfiguration!: pulumi.Output<outputs.s3.BucketReplicationConfiguration | undefined>;
    /**
     * Specifies who should bear the cost of Amazon S3 data transfer.
     * Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
     * the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
     * developer guide for more information.
     */
    public readonly requestPayer!: pulumi.Output<string>;
    /**
     * A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
     */
    public readonly serverSideEncryptionConfiguration!: pulumi.Output<outputs.s3.BucketServerSideEncryptionConfiguration>;
    /**
     * A map of tags to assign to the bucket. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
     */
    public readonly versioning!: pulumi.Output<outputs.s3.BucketVersioning>;
    /**
     * A website object (documented below).
     */
    public readonly website!: pulumi.Output<outputs.s3.BucketWebsite | undefined>;
    /**
     * The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
     */
    public readonly websiteDomain!: pulumi.Output<string>;
    /**
     * The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
     */
    public readonly websiteEndpoint!: pulumi.Output<string>;

    /**
     * Create a Bucket resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: BucketArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketArgs | BucketState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BucketState | undefined;
            resourceInputs["accelerationStatus"] = state ? state.accelerationStatus : undefined;
            resourceInputs["acl"] = state ? state.acl : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["bucketDomainName"] = state ? state.bucketDomainName : undefined;
            resourceInputs["bucketPrefix"] = state ? state.bucketPrefix : undefined;
            resourceInputs["bucketRegionalDomainName"] = state ? state.bucketRegionalDomainName : undefined;
            resourceInputs["corsRules"] = state ? state.corsRules : undefined;
            resourceInputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            resourceInputs["grants"] = state ? state.grants : undefined;
            resourceInputs["hostedZoneId"] = state ? state.hostedZoneId : undefined;
            resourceInputs["lifecycleRules"] = state ? state.lifecycleRules : undefined;
            resourceInputs["loggings"] = state ? state.loggings : undefined;
            resourceInputs["objectLockConfiguration"] = state ? state.objectLockConfiguration : undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["replicationConfiguration"] = state ? state.replicationConfiguration : undefined;
            resourceInputs["requestPayer"] = state ? state.requestPayer : undefined;
            resourceInputs["serverSideEncryptionConfiguration"] = state ? state.serverSideEncryptionConfiguration : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["versioning"] = state ? state.versioning : undefined;
            resourceInputs["website"] = state ? state.website : undefined;
            resourceInputs["websiteDomain"] = state ? state.websiteDomain : undefined;
            resourceInputs["websiteEndpoint"] = state ? state.websiteEndpoint : undefined;
        } else {
            const args = argsOrState as BucketArgs | undefined;
            resourceInputs["accelerationStatus"] = args ? args.accelerationStatus : undefined;
            resourceInputs["acl"] = args ? args.acl : undefined;
            resourceInputs["arn"] = args ? args.arn : undefined;
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["bucketPrefix"] = args ? args.bucketPrefix : undefined;
            resourceInputs["corsRules"] = args ? args.corsRules : undefined;
            resourceInputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            resourceInputs["grants"] = args ? args.grants : undefined;
            resourceInputs["hostedZoneId"] = args ? args.hostedZoneId : undefined;
            resourceInputs["lifecycleRules"] = args ? args.lifecycleRules : undefined;
            resourceInputs["loggings"] = args ? args.loggings : undefined;
            resourceInputs["objectLockConfiguration"] = args ? args.objectLockConfiguration : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["replicationConfiguration"] = args ? args.replicationConfiguration : undefined;
            resourceInputs["requestPayer"] = args ? args.requestPayer : undefined;
            resourceInputs["serverSideEncryptionConfiguration"] = args ? args.serverSideEncryptionConfiguration : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["versioning"] = args ? args.versioning : undefined;
            resourceInputs["website"] = args ? args.website : undefined;
            resourceInputs["websiteDomain"] = args ? args.websiteDomain : undefined;
            resourceInputs["websiteEndpoint"] = args ? args.websiteEndpoint : undefined;
            resourceInputs["bucketDomainName"] = undefined /*out*/;
            resourceInputs["bucketRegionalDomainName"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Bucket.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Bucket resources.
 */
export interface BucketState {
    /**
     * Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
     */
    accelerationStatus?: pulumi.Input<string>;
    /**
     * The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Valid values are `private`, `public-read`, `public-read-write`, `aws-exec-read`, `authenticated-read`, and `log-delivery-write`. Defaults to `private`.  Conflicts with `grant`.
     */
    acl?: pulumi.Input<string | enums.s3.CannedAcl>;
    /**
     * The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
     */
    arn?: pulumi.Input<string>;
    /**
     * The name of the bucket. If omitted, this provider will assign a random, unique name. Must be lowercase and less than or equal to 63 characters in length. A full list of bucket naming rules [may be found here](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html).
     */
    bucket?: pulumi.Input<string>;
    /**
     * The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
     */
    bucketDomainName?: pulumi.Input<string>;
    /**
     * Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`. Must be lowercase and less than or equal to 37 characters in length. A full list of bucket naming rules [may be found here](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html).
     */
    bucketPrefix?: pulumi.Input<string>;
    /**
     * The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.
     */
    bucketRegionalDomainName?: pulumi.Input<string>;
    /**
     * A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
     */
    corsRules?: pulumi.Input<pulumi.Input<inputs.s3.BucketCorsRule>[]>;
    /**
     * A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * An [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl) (documented below). Conflicts with `acl`.
     */
    grants?: pulumi.Input<pulumi.Input<inputs.s3.BucketGrant>[]>;
    /**
     * The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
     */
    hostedZoneId?: pulumi.Input<string>;
    /**
     * A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
     */
    lifecycleRules?: pulumi.Input<pulumi.Input<inputs.s3.BucketLifecycleRule>[]>;
    /**
     * A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
     */
    loggings?: pulumi.Input<pulumi.Input<inputs.s3.BucketLogging>[]>;
    /**
     * A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
     *
     * > **NOTE:** You cannot use `accelerationStatus` in `cn-north-1` or `us-gov-west-1`
     */
    objectLockConfiguration?: pulumi.Input<inputs.s3.BucketObjectLockConfiguration>;
    /**
     * A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), this provider may view the policy as constantly changing in a `pulumi preview`. In this case, please make sure you use the verbose/specific version of the policy.
     */
    policy?: pulumi.Input<string | PolicyDocument>;
    /**
     * The AWS region this bucket resides in.
     */
    region?: pulumi.Input<string>;
    /**
     * A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
     */
    replicationConfiguration?: pulumi.Input<inputs.s3.BucketReplicationConfiguration>;
    /**
     * Specifies who should bear the cost of Amazon S3 data transfer.
     * Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
     * the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
     * developer guide for more information.
     */
    requestPayer?: pulumi.Input<string>;
    /**
     * A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
     */
    serverSideEncryptionConfiguration?: pulumi.Input<inputs.s3.BucketServerSideEncryptionConfiguration>;
    /**
     * A map of tags to assign to the bucket. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
     */
    versioning?: pulumi.Input<inputs.s3.BucketVersioning>;
    /**
     * A website object (documented below).
     */
    website?: pulumi.Input<inputs.s3.BucketWebsite>;
    /**
     * The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
     */
    websiteDomain?: pulumi.Input<string>;
    /**
     * The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
     */
    websiteEndpoint?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Bucket resource.
 */
export interface BucketArgs {
    /**
     * Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
     */
    accelerationStatus?: pulumi.Input<string>;
    /**
     * The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Valid values are `private`, `public-read`, `public-read-write`, `aws-exec-read`, `authenticated-read`, and `log-delivery-write`. Defaults to `private`.  Conflicts with `grant`.
     */
    acl?: pulumi.Input<string | enums.s3.CannedAcl>;
    /**
     * The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
     */
    arn?: pulumi.Input<string>;
    /**
     * The name of the bucket. If omitted, this provider will assign a random, unique name. Must be lowercase and less than or equal to 63 characters in length. A full list of bucket naming rules [may be found here](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html).
     */
    bucket?: pulumi.Input<string>;
    /**
     * Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`. Must be lowercase and less than or equal to 37 characters in length. A full list of bucket naming rules [may be found here](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html).
     */
    bucketPrefix?: pulumi.Input<string>;
    /**
     * A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
     */
    corsRules?: pulumi.Input<pulumi.Input<inputs.s3.BucketCorsRule>[]>;
    /**
     * A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * An [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl) (documented below). Conflicts with `acl`.
     */
    grants?: pulumi.Input<pulumi.Input<inputs.s3.BucketGrant>[]>;
    /**
     * The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
     */
    hostedZoneId?: pulumi.Input<string>;
    /**
     * A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
     */
    lifecycleRules?: pulumi.Input<pulumi.Input<inputs.s3.BucketLifecycleRule>[]>;
    /**
     * A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
     */
    loggings?: pulumi.Input<pulumi.Input<inputs.s3.BucketLogging>[]>;
    /**
     * A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
     *
     * > **NOTE:** You cannot use `accelerationStatus` in `cn-north-1` or `us-gov-west-1`
     */
    objectLockConfiguration?: pulumi.Input<inputs.s3.BucketObjectLockConfiguration>;
    /**
     * A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), this provider may view the policy as constantly changing in a `pulumi preview`. In this case, please make sure you use the verbose/specific version of the policy.
     */
    policy?: pulumi.Input<string | PolicyDocument>;
    /**
     * A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
     */
    replicationConfiguration?: pulumi.Input<inputs.s3.BucketReplicationConfiguration>;
    /**
     * Specifies who should bear the cost of Amazon S3 data transfer.
     * Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
     * the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
     * developer guide for more information.
     */
    requestPayer?: pulumi.Input<string>;
    /**
     * A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
     */
    serverSideEncryptionConfiguration?: pulumi.Input<inputs.s3.BucketServerSideEncryptionConfiguration>;
    /**
     * A map of tags to assign to the bucket. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
     */
    versioning?: pulumi.Input<inputs.s3.BucketVersioning>;
    /**
     * A website object (documented below).
     */
    website?: pulumi.Input<inputs.s3.BucketWebsite>;
    /**
     * The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
     */
    websiteDomain?: pulumi.Input<string>;
    /**
     * The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
     */
    websiteEndpoint?: pulumi.Input<string>;
}
