// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Creates a Signer Signing Job.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testSp = new aws.signer.SigningProfile("testSp", {platformId: "AWSLambda-SHA384-ECDSA"});
 * const buildSigningJob = new aws.signer.SigningJob("buildSigningJob", {
 *     profileName: testSp.name,
 *     source: {
 *         s3: {
 *             bucket: "s3-bucket-name",
 *             key: "object-to-be-signed.zip",
 *             version: "jADjFYYYEXAMPLETszPjOmCMFDzd9dN1",
 *         },
 *     },
 *     destination: {
 *         s3: {
 *             bucket: "s3-bucket-name",
 *             prefix: "signed/",
 *         },
 *     },
 *     ignoreSigningJobFailure: true,
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_signer_signing_job.test_signer_signing_job
 *
 *  id = "9ed7e5c3-b8d4-4da0-8459-44e0b068f7ee" } Using `pulumi import`, import Signer signing jobs using the `job_id`. For exampleconsole % pulumi import aws_signer_signing_job.test_signer_signing_job 9ed7e5c3-b8d4-4da0-8459-44e0b068f7ee
 */
export class SigningJob extends pulumi.CustomResource {
    /**
     * Get an existing SigningJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SigningJobState, opts?: pulumi.CustomResourceOptions): SigningJob {
        return new SigningJob(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:signer/signingJob:SigningJob';

    /**
     * Returns true if the given object is an instance of SigningJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SigningJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SigningJob.__pulumiType;
    }

    /**
     * Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the signing job was completed.
     */
    public /*out*/ readonly completedAt!: pulumi.Output<string>;
    /**
     * Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the signing job was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The S3 bucket in which to save your signed object. See Destination below for details.
     */
    public readonly destination!: pulumi.Output<outputs.signer.SigningJobDestination>;
    /**
     * Set this argument to `true` to ignore signing job failures and retrieve failed status and reason. Default `false`.
     */
    public readonly ignoreSigningJobFailure!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the signing job on output.
     */
    public /*out*/ readonly jobId!: pulumi.Output<string>;
    /**
     * The IAM entity that initiated the signing job.
     */
    public /*out*/ readonly jobInvoker!: pulumi.Output<string>;
    /**
     * The AWS account ID of the job owner.
     */
    public /*out*/ readonly jobOwner!: pulumi.Output<string>;
    /**
     * A human-readable name for the signing platform associated with the signing job.
     */
    public /*out*/ readonly platformDisplayName!: pulumi.Output<string>;
    /**
     * The platform to which your signed code image will be distributed.
     */
    public /*out*/ readonly platformId!: pulumi.Output<string>;
    /**
     * The name of the profile to initiate the signing operation.
     */
    public readonly profileName!: pulumi.Output<string>;
    /**
     * The version of the signing profile used to initiate the signing job.
     */
    public /*out*/ readonly profileVersion!: pulumi.Output<string>;
    /**
     * The IAM principal that requested the signing job.
     */
    public /*out*/ readonly requestedBy!: pulumi.Output<string>;
    /**
     * A revocation record if the signature generated by the signing job has been revoked. Contains a timestamp and the ID of the IAM entity that revoked the signature.
     */
    public /*out*/ readonly revocationRecords!: pulumi.Output<outputs.signer.SigningJobRevocationRecord[]>;
    /**
     * The time when the signature of a signing job expires.
     */
    public /*out*/ readonly signatureExpiresAt!: pulumi.Output<string>;
    /**
     * Name of the S3 bucket where the signed code image is saved by code signing.
     */
    public /*out*/ readonly signedObjects!: pulumi.Output<outputs.signer.SigningJobSignedObject[]>;
    /**
     * The S3 bucket that contains the object to sign. See Source below for details.
     */
    public readonly source!: pulumi.Output<outputs.signer.SigningJobSource>;
    /**
     * Status of the signing job.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * String value that contains the status reason.
     */
    public /*out*/ readonly statusReason!: pulumi.Output<string>;

    /**
     * Create a SigningJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SigningJobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SigningJobArgs | SigningJobState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SigningJobState | undefined;
            resourceInputs["completedAt"] = state ? state.completedAt : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["destination"] = state ? state.destination : undefined;
            resourceInputs["ignoreSigningJobFailure"] = state ? state.ignoreSigningJobFailure : undefined;
            resourceInputs["jobId"] = state ? state.jobId : undefined;
            resourceInputs["jobInvoker"] = state ? state.jobInvoker : undefined;
            resourceInputs["jobOwner"] = state ? state.jobOwner : undefined;
            resourceInputs["platformDisplayName"] = state ? state.platformDisplayName : undefined;
            resourceInputs["platformId"] = state ? state.platformId : undefined;
            resourceInputs["profileName"] = state ? state.profileName : undefined;
            resourceInputs["profileVersion"] = state ? state.profileVersion : undefined;
            resourceInputs["requestedBy"] = state ? state.requestedBy : undefined;
            resourceInputs["revocationRecords"] = state ? state.revocationRecords : undefined;
            resourceInputs["signatureExpiresAt"] = state ? state.signatureExpiresAt : undefined;
            resourceInputs["signedObjects"] = state ? state.signedObjects : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["statusReason"] = state ? state.statusReason : undefined;
        } else {
            const args = argsOrState as SigningJobArgs | undefined;
            if ((!args || args.destination === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destination'");
            }
            if ((!args || args.profileName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'profileName'");
            }
            if ((!args || args.source === undefined) && !opts.urn) {
                throw new Error("Missing required property 'source'");
            }
            resourceInputs["destination"] = args ? args.destination : undefined;
            resourceInputs["ignoreSigningJobFailure"] = args ? args.ignoreSigningJobFailure : undefined;
            resourceInputs["profileName"] = args ? args.profileName : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["completedAt"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["jobId"] = undefined /*out*/;
            resourceInputs["jobInvoker"] = undefined /*out*/;
            resourceInputs["jobOwner"] = undefined /*out*/;
            resourceInputs["platformDisplayName"] = undefined /*out*/;
            resourceInputs["platformId"] = undefined /*out*/;
            resourceInputs["profileVersion"] = undefined /*out*/;
            resourceInputs["requestedBy"] = undefined /*out*/;
            resourceInputs["revocationRecords"] = undefined /*out*/;
            resourceInputs["signatureExpiresAt"] = undefined /*out*/;
            resourceInputs["signedObjects"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusReason"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SigningJob.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SigningJob resources.
 */
export interface SigningJobState {
    /**
     * Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the signing job was completed.
     */
    completedAt?: pulumi.Input<string>;
    /**
     * Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the signing job was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The S3 bucket in which to save your signed object. See Destination below for details.
     */
    destination?: pulumi.Input<inputs.signer.SigningJobDestination>;
    /**
     * Set this argument to `true` to ignore signing job failures and retrieve failed status and reason. Default `false`.
     */
    ignoreSigningJobFailure?: pulumi.Input<boolean>;
    /**
     * The ID of the signing job on output.
     */
    jobId?: pulumi.Input<string>;
    /**
     * The IAM entity that initiated the signing job.
     */
    jobInvoker?: pulumi.Input<string>;
    /**
     * The AWS account ID of the job owner.
     */
    jobOwner?: pulumi.Input<string>;
    /**
     * A human-readable name for the signing platform associated with the signing job.
     */
    platformDisplayName?: pulumi.Input<string>;
    /**
     * The platform to which your signed code image will be distributed.
     */
    platformId?: pulumi.Input<string>;
    /**
     * The name of the profile to initiate the signing operation.
     */
    profileName?: pulumi.Input<string>;
    /**
     * The version of the signing profile used to initiate the signing job.
     */
    profileVersion?: pulumi.Input<string>;
    /**
     * The IAM principal that requested the signing job.
     */
    requestedBy?: pulumi.Input<string>;
    /**
     * A revocation record if the signature generated by the signing job has been revoked. Contains a timestamp and the ID of the IAM entity that revoked the signature.
     */
    revocationRecords?: pulumi.Input<pulumi.Input<inputs.signer.SigningJobRevocationRecord>[]>;
    /**
     * The time when the signature of a signing job expires.
     */
    signatureExpiresAt?: pulumi.Input<string>;
    /**
     * Name of the S3 bucket where the signed code image is saved by code signing.
     */
    signedObjects?: pulumi.Input<pulumi.Input<inputs.signer.SigningJobSignedObject>[]>;
    /**
     * The S3 bucket that contains the object to sign. See Source below for details.
     */
    source?: pulumi.Input<inputs.signer.SigningJobSource>;
    /**
     * Status of the signing job.
     */
    status?: pulumi.Input<string>;
    /**
     * String value that contains the status reason.
     */
    statusReason?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SigningJob resource.
 */
export interface SigningJobArgs {
    /**
     * The S3 bucket in which to save your signed object. See Destination below for details.
     */
    destination: pulumi.Input<inputs.signer.SigningJobDestination>;
    /**
     * Set this argument to `true` to ignore signing job failures and retrieve failed status and reason. Default `false`.
     */
    ignoreSigningJobFailure?: pulumi.Input<boolean>;
    /**
     * The name of the profile to initiate the signing operation.
     */
    profileName: pulumi.Input<string>;
    /**
     * The S3 bucket that contains the object to sign. See Source below for details.
     */
    source: pulumi.Input<inputs.signer.SigningJobSource>;
}
