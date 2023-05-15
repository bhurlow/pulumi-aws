// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Registers a custom domain name for use with AWS API Gateway. Additional information about this functionality
 * can be found in the [API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html).
 *
 * This resource just establishes ownership of and the TLS settings for
 * a particular domain name. An API can be attached to a particular path
 * under the registered domain name using
 * the `aws.apigateway.BasePathMapping` resource.
 *
 * API Gateway domains can be defined as either 'edge-optimized' or 'regional'.  In an edge-optimized configuration,
 * API Gateway internally creates and manages a CloudFront distribution to route requests on the given hostname. In
 * addition to this resource it's necessary to create a DNS record corresponding to the given domain name which is an alias
 * (either Route53 alias or traditional CNAME) to the Cloudfront domain name exported in the `cloudfrontDomainName`
 * attribute.
 *
 * In a regional configuration, API Gateway does not create a CloudFront distribution to route requests to the API, though
 * a distribution can be created if needed. In either case, it is necessary to create a DNS record corresponding to the
 * given domain name which is an alias (either Route53 alias or traditional CNAME) to the regional domain name exported in
 * the `regionalDomainName` attribute.
 *
 * > **Note:** API Gateway requires the use of AWS Certificate Manager (ACM) certificates instead of Identity and Access Management (IAM) certificates in regions that support ACM. Regions that support ACM can be found in the [Regions and Endpoints Documentation](https://docs.aws.amazon.com/general/latest/gr/rande.html#acm_region). To import an existing private key and certificate into ACM or request an ACM certificate, see the `aws.acm.Certificate` resource.
 *
 * > **Note:** The `aws.apigateway.DomainName` resource expects dependency on the `aws.acm.CertificateValidation` as
 * only verified certificates can be used. This can be made either explicitly by adding the
 * `dependsOn = [aws_acm_certificate_validation.cert]` attribute. Or implicitly by referring certificate ARN
 * from the validation resource where it will be available after the resource creation:
 * `regionalCertificateArn = aws_acm_certificate_validation.cert.certificate_arn`.
 *
 * ## Example Usage
 * ### Edge Optimized (ACM Certificate)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleDomainName = new aws.apigateway.DomainName("exampleDomainName", {
 *     certificateArn: aws_acm_certificate_validation.example.certificate_arn,
 *     domainName: "api.example.com",
 * });
 * // Example DNS record using Route53.
 * // Route53 is not specifically required; any DNS host can be used.
 * const exampleRecord = new aws.route53.Record("exampleRecord", {
 *     name: exampleDomainName.domainName,
 *     type: "A",
 *     zoneId: aws_route53_zone.example.id,
 *     aliases: [{
 *         evaluateTargetHealth: true,
 *         name: exampleDomainName.cloudfrontDomainName,
 *         zoneId: exampleDomainName.cloudfrontZoneId,
 *     }],
 * });
 * ```
 * ### Edge Optimized (IAM Certificate)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as fs from "fs";
 *
 * const exampleDomainName = new aws.apigateway.DomainName("exampleDomainName", {
 *     domainName: "api.example.com",
 *     certificateName: "example-api",
 *     certificateBody: fs.readFileSync(`${path.module}/example.com/example.crt`),
 *     certificateChain: fs.readFileSync(`${path.module}/example.com/ca.crt`),
 *     certificatePrivateKey: fs.readFileSync(`${path.module}/example.com/example.key`),
 * });
 * // Example DNS record using Route53.
 * // Route53 is not specifically required; any DNS host can be used.
 * const exampleRecord = new aws.route53.Record("exampleRecord", {
 *     zoneId: aws_route53_zone.example.id,
 *     name: exampleDomainName.domainName,
 *     type: "A",
 *     aliases: [{
 *         name: exampleDomainName.cloudfrontDomainName,
 *         zoneId: exampleDomainName.cloudfrontZoneId,
 *         evaluateTargetHealth: true,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * API Gateway domain names can be imported using their `name`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:apigateway/domainName:DomainName example dev.example.com
 * ```
 */
export class DomainName extends pulumi.CustomResource {
    /**
     * Get an existing DomainName resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainNameState, opts?: pulumi.CustomResourceOptions): DomainName {
        return new DomainName(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigateway/domainName:DomainName';

    /**
     * Returns true if the given object is an instance of DomainName.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainName {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainName.__pulumiType;
    }

    /**
     * ARN of domain name.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificateName`, `certificateBody`, `certificateChain`, `certificatePrivateKey`, `regionalCertificateArn`, and `regionalCertificateName`.
     */
    public readonly certificateArn!: pulumi.Output<string | undefined>;
    /**
     * Certificate issued for the domain name being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`.
     */
    public readonly certificateBody!: pulumi.Output<string | undefined>;
    /**
     * Certificate for the CA that issued the certificate, along with any intermediate CA certificates required to create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`.
     */
    public readonly certificateChain!: pulumi.Output<string | undefined>;
    /**
     * Unique name to use when registering this certificate as an IAM server certificate. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`. Required if `certificateArn` is not set.
     */
    public readonly certificateName!: pulumi.Output<string | undefined>;
    /**
     * Private key associated with the domain certificate given in `certificateBody`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`.
     */
    public readonly certificatePrivateKey!: pulumi.Output<string | undefined>;
    /**
     * Upload date associated with the domain certificate.
     */
    public /*out*/ readonly certificateUploadDate!: pulumi.Output<string>;
    /**
     * Hostname created by Cloudfront to represent the distribution that implements this domain name mapping.
     */
    public /*out*/ readonly cloudfrontDomainName!: pulumi.Output<string>;
    /**
     * For convenience, the hosted zone ID (`Z2FDTNDATAQYW2`) that can be used to create a Route53 alias record for the distribution.
     */
    public /*out*/ readonly cloudfrontZoneId!: pulumi.Output<string>;
    /**
     * Fully-qualified domain name to register.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * Configuration block defining API endpoint information including type. See below.
     */
    public readonly endpointConfiguration!: pulumi.Output<outputs.apigateway.DomainNameEndpointConfiguration>;
    /**
     * Mutual TLS authentication configuration for the domain name. See below.
     */
    public readonly mutualTlsAuthentication!: pulumi.Output<outputs.apigateway.DomainNameMutualTlsAuthentication | undefined>;
    /**
     * ARN of the AWS-issued certificate used to validate custom domain ownership (when `certificateArn` is issued via an ACM Private CA or `mutualTlsAuthentication` is configured with an ACM-imported certificate.)
     */
    public readonly ownershipVerificationCertificateArn!: pulumi.Output<string>;
    /**
     * ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and `certificatePrivateKey`.
     */
    public readonly regionalCertificateArn!: pulumi.Output<string | undefined>;
    /**
     * User-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and `certificatePrivateKey`.
     */
    public readonly regionalCertificateName!: pulumi.Output<string | undefined>;
    /**
     * Hostname for the custom domain's regional endpoint.
     */
    public /*out*/ readonly regionalDomainName!: pulumi.Output<string>;
    /**
     * Hosted zone ID that can be used to create a Route53 alias record for the regional endpoint.
     */
    public /*out*/ readonly regionalZoneId!: pulumi.Output<string>;
    /**
     * Transport Layer Security (TLS) version + cipher suite for this DomainName. Valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
     */
    public readonly securityPolicy!: pulumi.Output<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a DomainName resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainNameArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainNameArgs | DomainNameState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainNameState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["certificateArn"] = state ? state.certificateArn : undefined;
            resourceInputs["certificateBody"] = state ? state.certificateBody : undefined;
            resourceInputs["certificateChain"] = state ? state.certificateChain : undefined;
            resourceInputs["certificateName"] = state ? state.certificateName : undefined;
            resourceInputs["certificatePrivateKey"] = state ? state.certificatePrivateKey : undefined;
            resourceInputs["certificateUploadDate"] = state ? state.certificateUploadDate : undefined;
            resourceInputs["cloudfrontDomainName"] = state ? state.cloudfrontDomainName : undefined;
            resourceInputs["cloudfrontZoneId"] = state ? state.cloudfrontZoneId : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["endpointConfiguration"] = state ? state.endpointConfiguration : undefined;
            resourceInputs["mutualTlsAuthentication"] = state ? state.mutualTlsAuthentication : undefined;
            resourceInputs["ownershipVerificationCertificateArn"] = state ? state.ownershipVerificationCertificateArn : undefined;
            resourceInputs["regionalCertificateArn"] = state ? state.regionalCertificateArn : undefined;
            resourceInputs["regionalCertificateName"] = state ? state.regionalCertificateName : undefined;
            resourceInputs["regionalDomainName"] = state ? state.regionalDomainName : undefined;
            resourceInputs["regionalZoneId"] = state ? state.regionalZoneId : undefined;
            resourceInputs["securityPolicy"] = state ? state.securityPolicy : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as DomainNameArgs | undefined;
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            resourceInputs["certificateArn"] = args ? args.certificateArn : undefined;
            resourceInputs["certificateBody"] = args ? args.certificateBody : undefined;
            resourceInputs["certificateChain"] = args ? args.certificateChain : undefined;
            resourceInputs["certificateName"] = args ? args.certificateName : undefined;
            resourceInputs["certificatePrivateKey"] = args?.certificatePrivateKey ? pulumi.secret(args.certificatePrivateKey) : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["endpointConfiguration"] = args ? args.endpointConfiguration : undefined;
            resourceInputs["mutualTlsAuthentication"] = args ? args.mutualTlsAuthentication : undefined;
            resourceInputs["ownershipVerificationCertificateArn"] = args ? args.ownershipVerificationCertificateArn : undefined;
            resourceInputs["regionalCertificateArn"] = args ? args.regionalCertificateArn : undefined;
            resourceInputs["regionalCertificateName"] = args ? args.regionalCertificateName : undefined;
            resourceInputs["securityPolicy"] = args ? args.securityPolicy : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["certificateUploadDate"] = undefined /*out*/;
            resourceInputs["cloudfrontDomainName"] = undefined /*out*/;
            resourceInputs["cloudfrontZoneId"] = undefined /*out*/;
            resourceInputs["regionalDomainName"] = undefined /*out*/;
            resourceInputs["regionalZoneId"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["certificatePrivateKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(DomainName.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainName resources.
 */
export interface DomainNameState {
    /**
     * ARN of domain name.
     */
    arn?: pulumi.Input<string>;
    /**
     * ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificateName`, `certificateBody`, `certificateChain`, `certificatePrivateKey`, `regionalCertificateArn`, and `regionalCertificateName`.
     */
    certificateArn?: pulumi.Input<string>;
    /**
     * Certificate issued for the domain name being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`.
     */
    certificateBody?: pulumi.Input<string>;
    /**
     * Certificate for the CA that issued the certificate, along with any intermediate CA certificates required to create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`.
     */
    certificateChain?: pulumi.Input<string>;
    /**
     * Unique name to use when registering this certificate as an IAM server certificate. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`. Required if `certificateArn` is not set.
     */
    certificateName?: pulumi.Input<string>;
    /**
     * Private key associated with the domain certificate given in `certificateBody`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`.
     */
    certificatePrivateKey?: pulumi.Input<string>;
    /**
     * Upload date associated with the domain certificate.
     */
    certificateUploadDate?: pulumi.Input<string>;
    /**
     * Hostname created by Cloudfront to represent the distribution that implements this domain name mapping.
     */
    cloudfrontDomainName?: pulumi.Input<string>;
    /**
     * For convenience, the hosted zone ID (`Z2FDTNDATAQYW2`) that can be used to create a Route53 alias record for the distribution.
     */
    cloudfrontZoneId?: pulumi.Input<string>;
    /**
     * Fully-qualified domain name to register.
     */
    domainName?: pulumi.Input<string>;
    /**
     * Configuration block defining API endpoint information including type. See below.
     */
    endpointConfiguration?: pulumi.Input<inputs.apigateway.DomainNameEndpointConfiguration>;
    /**
     * Mutual TLS authentication configuration for the domain name. See below.
     */
    mutualTlsAuthentication?: pulumi.Input<inputs.apigateway.DomainNameMutualTlsAuthentication>;
    /**
     * ARN of the AWS-issued certificate used to validate custom domain ownership (when `certificateArn` is issued via an ACM Private CA or `mutualTlsAuthentication` is configured with an ACM-imported certificate.)
     */
    ownershipVerificationCertificateArn?: pulumi.Input<string>;
    /**
     * ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and `certificatePrivateKey`.
     */
    regionalCertificateArn?: pulumi.Input<string>;
    /**
     * User-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and `certificatePrivateKey`.
     */
    regionalCertificateName?: pulumi.Input<string>;
    /**
     * Hostname for the custom domain's regional endpoint.
     */
    regionalDomainName?: pulumi.Input<string>;
    /**
     * Hosted zone ID that can be used to create a Route53 alias record for the regional endpoint.
     */
    regionalZoneId?: pulumi.Input<string>;
    /**
     * Transport Layer Security (TLS) version + cipher suite for this DomainName. Valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
     */
    securityPolicy?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a DomainName resource.
 */
export interface DomainNameArgs {
    /**
     * ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificateName`, `certificateBody`, `certificateChain`, `certificatePrivateKey`, `regionalCertificateArn`, and `regionalCertificateName`.
     */
    certificateArn?: pulumi.Input<string>;
    /**
     * Certificate issued for the domain name being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`.
     */
    certificateBody?: pulumi.Input<string>;
    /**
     * Certificate for the CA that issued the certificate, along with any intermediate CA certificates required to create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`.
     */
    certificateChain?: pulumi.Input<string>;
    /**
     * Unique name to use when registering this certificate as an IAM server certificate. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`. Required if `certificateArn` is not set.
     */
    certificateName?: pulumi.Input<string>;
    /**
     * Private key associated with the domain certificate given in `certificateBody`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`.
     */
    certificatePrivateKey?: pulumi.Input<string>;
    /**
     * Fully-qualified domain name to register.
     */
    domainName: pulumi.Input<string>;
    /**
     * Configuration block defining API endpoint information including type. See below.
     */
    endpointConfiguration?: pulumi.Input<inputs.apigateway.DomainNameEndpointConfiguration>;
    /**
     * Mutual TLS authentication configuration for the domain name. See below.
     */
    mutualTlsAuthentication?: pulumi.Input<inputs.apigateway.DomainNameMutualTlsAuthentication>;
    /**
     * ARN of the AWS-issued certificate used to validate custom domain ownership (when `certificateArn` is issued via an ACM Private CA or `mutualTlsAuthentication` is configured with an ACM-imported certificate.)
     */
    ownershipVerificationCertificateArn?: pulumi.Input<string>;
    /**
     * ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and `certificatePrivateKey`.
     */
    regionalCertificateArn?: pulumi.Input<string>;
    /**
     * User-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and `certificatePrivateKey`.
     */
    regionalCertificateName?: pulumi.Input<string>;
    /**
     * Transport Layer Security (TLS) version + cipher suite for this DomainName. Valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
     */
    securityPolicy?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
