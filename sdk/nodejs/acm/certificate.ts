// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The ACM certificate resource allows requesting and management of certificates
 * from the Amazon Certificate Manager.
 *
 * ACM certificates can be created in three ways:
 * Amazon-issued, where AWS provides the certificate authority and automatically manages renewal;
 * imported certificates, issued by another certificate authority;
 * and private certificates, issued using an ACM Private Certificate Authority.
 *
 * ## Amazon-Issued Certificates
 *
 * For Amazon-issued certificates, this resource deals with requesting certificates and managing their attributes and life-cycle.
 * This resource does not deal with validation of a certificate but can provide inputs
 * for other resources implementing the validation.
 * It does not wait for a certificate to be issued.
 * Use a `aws.acm.CertificateValidation` resource for this.
 *
 * Most commonly, this resource is used together with `aws.route53.Record` and
 * `aws.acm.CertificateValidation` to request a DNS validated certificate,
 * deploy the required validation records and wait for validation to complete.
 *
 * Domain validation through email is also supported but should be avoided as it requires a manual step outside of this provider.
 *
 * ## Certificates Imported from Other Certificate Authority
 *
 * Imported certificates can be used to make certificates created with an external certificate authority available for AWS services.
 *
 * As they are not managed by AWS, imported certificates are not eligible for automatic renewal.
 * New certificate materials can be supplied to an existing imported certificate to update it in place.
 *
 * ## Private Certificates
 *
 * Private certificates are issued by an ACM Private Cerificate Authority, which can be created using the resource type `aws.acmpca.CertificateAuthority`.
 *
 * Private certificates created using this resource are eligible for managed renewal if they have been exported or associated with another AWS service.
 * See [managed renewal documentation](https://docs.aws.amazon.com/acm/latest/userguide/managed-renewal.html) for more information.
 * By default, a certificate is valid for 395 days and the managed renewal process will start 60 days before expiration.
 * To renew the certificate earlier than 60 days before expiration, configure `earlyRenewalDuration`.
 *
 * ## Example Usage
 * ### Create Certificate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const cert = new aws.acm.Certificate("cert", {
 *     domainName: "example.com",
 *     tags: {
 *         Environment: "test",
 *     },
 *     validationMethod: "DNS",
 * });
 * ```
 * ### Custom Domain Validation Options
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const cert = new aws.acm.Certificate("cert", {
 *     domainName: "testing.example.com",
 *     validationMethod: "EMAIL",
 *     validationOptions: [{
 *         domainName: "testing.example.com",
 *         validationDomain: "example.com",
 *     }],
 * });
 * ```
 * ### Existing Certificate Body Import
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as tls from "@pulumi/tls";
 *
 * const examplePrivateKey = new tls.PrivateKey("examplePrivateKey", {algorithm: "RSA"});
 * const exampleSelfSignedCert = new tls.SelfSignedCert("exampleSelfSignedCert", {
 *     keyAlgorithm: "RSA",
 *     privateKeyPem: examplePrivateKey.privateKeyPem,
 *     subjects: [{
 *         commonName: "example.com",
 *         organization: "ACME Examples, Inc",
 *     }],
 *     validityPeriodHours: 12,
 *     allowedUses: [
 *         "key_encipherment",
 *         "digital_signature",
 *         "server_auth",
 *     ],
 * });
 * const cert = new aws.acm.Certificate("cert", {
 *     privateKey: examplePrivateKey.privateKeyPem,
 *     certificateBody: exampleSelfSignedCert.certPem,
 * });
 * ```
 * ### Referencing domainValidationOptions With forEach Based Resources
 *
 * See the `aws.acm.CertificateValidation` resource for a full example of performing DNS validation.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example: aws.route53.Record[] = [];
 * for (const range of Object.entries(.reduce((__obj, dvo) => { ...__obj, [dvo.domainName]: {
 *     name: dvo.resourceRecordName,
 *     record: dvo.resourceRecordValue,
 *     type: dvo.resourceRecordType,
 * } })).map(([k, v]) => ({key: k, value: v}))) {
 *     example.push(new aws.route53.Record(`example-${range.key}`, {
 *         allowOverwrite: true,
 *         name: range.value.name,
 *         records: [range.value.record],
 *         ttl: 60,
 *         type: aws.route53.recordtype.RecordType[range.value.type],
 *         zoneId: aws_route53_zone.example.zone_id,
 *     }));
 * }
 * ```
 *
 * ## Import
 *
 * Certificates can be imported using their ARN, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:acm/certificate:Certificate cert arn:aws:acm:eu-central-1:123456789012:certificate/7e7a28d2-163f-4b8f-b9cd-822f96c08d6a
 * ```
 */
export class Certificate extends pulumi.CustomResource {
    /**
     * Get an existing Certificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CertificateState, opts?: pulumi.CustomResourceOptions): Certificate {
        return new Certificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:acm/certificate:Certificate';

    /**
     * Returns true if the given object is an instance of Certificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Certificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Certificate.__pulumiType;
    }

    /**
     * ARN of the certificate
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * ARN of an ACM PCA
     */
    public readonly certificateAuthorityArn!: pulumi.Output<string | undefined>;
    /**
     * Certificate's PEM-formatted public key
     */
    public readonly certificateBody!: pulumi.Output<string | undefined>;
    /**
     * Certificate's PEM-formatted chain
     * * Creating a private CA issued certificate
     */
    public readonly certificateChain!: pulumi.Output<string | undefined>;
    /**
     * Fully qualified domain name (FQDN) in the certificate.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * Set of domain validation objects which can be used to complete certificate validation.
     * Can have more than one element, e.g., if SANs are defined.
     * Only set if `DNS`-validation was used.
     */
    public /*out*/ readonly domainValidationOptions!: pulumi.Output<outputs.acm.CertificateDomainValidationOption[]>;
    /**
     * Amount of time to start automatic renewal process before expiration.
     * Has no effect if less than 60 days.
     * Represented by either
     * a subset of [RFC 3339 duration](https://www.rfc-editor.org/rfc/rfc3339) supporting years, months, and days (e.g., `P90D`),
     * or a string such as `2160h`.
     */
    public readonly earlyRenewalDuration!: pulumi.Output<string | undefined>;
    /**
     * Specifies the algorithm of the public and private key pair that your Amazon issued certificate uses to encrypt data. See [ACM Certificate characteristics](https://docs.aws.amazon.com/acm/latest/userguide/acm-certificate.html#algorithms) for more details.
     */
    public readonly keyAlgorithm!: pulumi.Output<string>;
    /**
     * Expiration date and time of the certificate.
     */
    public /*out*/ readonly notAfter!: pulumi.Output<string>;
    /**
     * Start of the validity period of the certificate.
     */
    public /*out*/ readonly notBefore!: pulumi.Output<string>;
    /**
     * Configuration block used to set certificate options. Detailed below.
     */
    public readonly options!: pulumi.Output<outputs.acm.CertificateOptions>;
    /**
     * `true` if a Private certificate eligible for managed renewal is within the `earlyRenewalDuration` period.
     */
    public /*out*/ readonly pendingRenewal!: pulumi.Output<boolean>;
    /**
     * Certificate's PEM-formatted private key
     */
    public readonly privateKey!: pulumi.Output<string | undefined>;
    /**
     * Whether the certificate is eligible for managed renewal.
     */
    public /*out*/ readonly renewalEligibility!: pulumi.Output<string>;
    /**
     * Contains information about the status of ACM's [managed renewal](https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html) for the certificate.
     */
    public /*out*/ readonly renewalSummaries!: pulumi.Output<outputs.acm.CertificateRenewalSummary[]>;
    /**
     * Status of the certificate.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Set of domains that should be SANs in the issued certificate.
     * To remove all elements of a previously configured list, set this value equal to an empty list (`[]`)
     */
    public readonly subjectAlternativeNames!: pulumi.Output<string[]>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Source of the certificate.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * List of addresses that received a validation email. Only set if `EMAIL` validation was used.
     */
    public /*out*/ readonly validationEmails!: pulumi.Output<string[]>;
    /**
     * Which method to use for validation. `DNS` or `EMAIL` are valid, `NONE` can be used for certificates that were imported into ACM and then into the provider.
     */
    public readonly validationMethod!: pulumi.Output<string>;
    /**
     * Configuration block used to specify information about the initial validation of each domain name. Detailed below.
     * * Importing an existing certificate
     */
    public readonly validationOptions!: pulumi.Output<outputs.acm.CertificateValidationOption[] | undefined>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateArgs | CertificateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CertificateState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["certificateAuthorityArn"] = state ? state.certificateAuthorityArn : undefined;
            resourceInputs["certificateBody"] = state ? state.certificateBody : undefined;
            resourceInputs["certificateChain"] = state ? state.certificateChain : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["domainValidationOptions"] = state ? state.domainValidationOptions : undefined;
            resourceInputs["earlyRenewalDuration"] = state ? state.earlyRenewalDuration : undefined;
            resourceInputs["keyAlgorithm"] = state ? state.keyAlgorithm : undefined;
            resourceInputs["notAfter"] = state ? state.notAfter : undefined;
            resourceInputs["notBefore"] = state ? state.notBefore : undefined;
            resourceInputs["options"] = state ? state.options : undefined;
            resourceInputs["pendingRenewal"] = state ? state.pendingRenewal : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["renewalEligibility"] = state ? state.renewalEligibility : undefined;
            resourceInputs["renewalSummaries"] = state ? state.renewalSummaries : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subjectAlternativeNames"] = state ? state.subjectAlternativeNames : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["validationEmails"] = state ? state.validationEmails : undefined;
            resourceInputs["validationMethod"] = state ? state.validationMethod : undefined;
            resourceInputs["validationOptions"] = state ? state.validationOptions : undefined;
        } else {
            const args = argsOrState as CertificateArgs | undefined;
            resourceInputs["certificateAuthorityArn"] = args ? args.certificateAuthorityArn : undefined;
            resourceInputs["certificateBody"] = args ? args.certificateBody : undefined;
            resourceInputs["certificateChain"] = args ? args.certificateChain : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["earlyRenewalDuration"] = args ? args.earlyRenewalDuration : undefined;
            resourceInputs["keyAlgorithm"] = args ? args.keyAlgorithm : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["privateKey"] = args?.privateKey ? pulumi.secret(args.privateKey) : undefined;
            resourceInputs["subjectAlternativeNames"] = args ? args.subjectAlternativeNames : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["validationMethod"] = args ? args.validationMethod : undefined;
            resourceInputs["validationOptions"] = args ? args.validationOptions : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["domainValidationOptions"] = undefined /*out*/;
            resourceInputs["notAfter"] = undefined /*out*/;
            resourceInputs["notBefore"] = undefined /*out*/;
            resourceInputs["pendingRenewal"] = undefined /*out*/;
            resourceInputs["renewalEligibility"] = undefined /*out*/;
            resourceInputs["renewalSummaries"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["validationEmails"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["privateKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Certificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Certificate resources.
 */
export interface CertificateState {
    /**
     * ARN of the certificate
     */
    arn?: pulumi.Input<string>;
    /**
     * ARN of an ACM PCA
     */
    certificateAuthorityArn?: pulumi.Input<string>;
    /**
     * Certificate's PEM-formatted public key
     */
    certificateBody?: pulumi.Input<string>;
    /**
     * Certificate's PEM-formatted chain
     * * Creating a private CA issued certificate
     */
    certificateChain?: pulumi.Input<string>;
    /**
     * Fully qualified domain name (FQDN) in the certificate.
     */
    domainName?: pulumi.Input<string>;
    /**
     * Set of domain validation objects which can be used to complete certificate validation.
     * Can have more than one element, e.g., if SANs are defined.
     * Only set if `DNS`-validation was used.
     */
    domainValidationOptions?: pulumi.Input<pulumi.Input<inputs.acm.CertificateDomainValidationOption>[]>;
    /**
     * Amount of time to start automatic renewal process before expiration.
     * Has no effect if less than 60 days.
     * Represented by either
     * a subset of [RFC 3339 duration](https://www.rfc-editor.org/rfc/rfc3339) supporting years, months, and days (e.g., `P90D`),
     * or a string such as `2160h`.
     */
    earlyRenewalDuration?: pulumi.Input<string>;
    /**
     * Specifies the algorithm of the public and private key pair that your Amazon issued certificate uses to encrypt data. See [ACM Certificate characteristics](https://docs.aws.amazon.com/acm/latest/userguide/acm-certificate.html#algorithms) for more details.
     */
    keyAlgorithm?: pulumi.Input<string>;
    /**
     * Expiration date and time of the certificate.
     */
    notAfter?: pulumi.Input<string>;
    /**
     * Start of the validity period of the certificate.
     */
    notBefore?: pulumi.Input<string>;
    /**
     * Configuration block used to set certificate options. Detailed below.
     */
    options?: pulumi.Input<inputs.acm.CertificateOptions>;
    /**
     * `true` if a Private certificate eligible for managed renewal is within the `earlyRenewalDuration` period.
     */
    pendingRenewal?: pulumi.Input<boolean>;
    /**
     * Certificate's PEM-formatted private key
     */
    privateKey?: pulumi.Input<string>;
    /**
     * Whether the certificate is eligible for managed renewal.
     */
    renewalEligibility?: pulumi.Input<string>;
    /**
     * Contains information about the status of ACM's [managed renewal](https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html) for the certificate.
     */
    renewalSummaries?: pulumi.Input<pulumi.Input<inputs.acm.CertificateRenewalSummary>[]>;
    /**
     * Status of the certificate.
     */
    status?: pulumi.Input<string>;
    /**
     * Set of domains that should be SANs in the issued certificate.
     * To remove all elements of a previously configured list, set this value equal to an empty list (`[]`)
     */
    subjectAlternativeNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Source of the certificate.
     */
    type?: pulumi.Input<string>;
    /**
     * List of addresses that received a validation email. Only set if `EMAIL` validation was used.
     */
    validationEmails?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Which method to use for validation. `DNS` or `EMAIL` are valid, `NONE` can be used for certificates that were imported into ACM and then into the provider.
     */
    validationMethod?: pulumi.Input<string>;
    /**
     * Configuration block used to specify information about the initial validation of each domain name. Detailed below.
     * * Importing an existing certificate
     */
    validationOptions?: pulumi.Input<pulumi.Input<inputs.acm.CertificateValidationOption>[]>;
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * ARN of an ACM PCA
     */
    certificateAuthorityArn?: pulumi.Input<string>;
    /**
     * Certificate's PEM-formatted public key
     */
    certificateBody?: pulumi.Input<string>;
    /**
     * Certificate's PEM-formatted chain
     * * Creating a private CA issued certificate
     */
    certificateChain?: pulumi.Input<string>;
    /**
     * Fully qualified domain name (FQDN) in the certificate.
     */
    domainName?: pulumi.Input<string>;
    /**
     * Amount of time to start automatic renewal process before expiration.
     * Has no effect if less than 60 days.
     * Represented by either
     * a subset of [RFC 3339 duration](https://www.rfc-editor.org/rfc/rfc3339) supporting years, months, and days (e.g., `P90D`),
     * or a string such as `2160h`.
     */
    earlyRenewalDuration?: pulumi.Input<string>;
    /**
     * Specifies the algorithm of the public and private key pair that your Amazon issued certificate uses to encrypt data. See [ACM Certificate characteristics](https://docs.aws.amazon.com/acm/latest/userguide/acm-certificate.html#algorithms) for more details.
     */
    keyAlgorithm?: pulumi.Input<string>;
    /**
     * Configuration block used to set certificate options. Detailed below.
     */
    options?: pulumi.Input<inputs.acm.CertificateOptions>;
    /**
     * Certificate's PEM-formatted private key
     */
    privateKey?: pulumi.Input<string>;
    /**
     * Set of domains that should be SANs in the issued certificate.
     * To remove all elements of a previously configured list, set this value equal to an empty list (`[]`)
     */
    subjectAlternativeNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Which method to use for validation. `DNS` or `EMAIL` are valid, `NONE` can be used for certificates that were imported into ACM and then into the provider.
     */
    validationMethod?: pulumi.Input<string>;
    /**
     * Configuration block used to specify information about the initial validation of each domain name. Detailed below.
     * * Importing an existing certificate
     */
    validationOptions?: pulumi.Input<pulumi.Input<inputs.acm.CertificateValidationOption>[]>;
}
