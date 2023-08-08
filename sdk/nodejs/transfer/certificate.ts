// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a AWS Transfer AS2 Certificate resource.
 *
 * ## Example Usage
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as fs from "fs";
 *
 * const example = new aws.transfer.Certificate("example", {
 *     certificate: fs.readFileSync(`${path.module}/example.com/example.crt`),
 *     certificateChain: fs.readFileSync(`${path.module}/example.com/ca.crt`),
 *     privateKey: fs.readFileSync(`${path.module}/example.com/example.key`),
 *     description: "example",
 *     usage: "SIGNING",
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_transfer_certificate.example
 *
 *  id = "c-4221a88afd5f4362a" } Using `pulumi import`, import Transfer AS2 Certificate using the `certificate_id`. For exampleconsole % pulumi import aws_transfer_certificate.example c-4221a88afd5f4362a
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
    public static readonly __pulumiType = 'aws:transfer/certificate:Certificate';

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
     * An date when the certificate becomes active
     */
    public /*out*/ readonly activeDate!: pulumi.Output<string>;
    /**
     * The valid certificate file required for the transfer.
     */
    public readonly certificate!: pulumi.Output<string>;
    /**
     * The optional list of certificate that make up the chain for the certificate that is being imported.
     */
    public readonly certificateChain!: pulumi.Output<string | undefined>;
    /**
     * The unique identifier for the AS2 certificate
     */
    public /*out*/ readonly certificateId!: pulumi.Output<string>;
    /**
     * A short description that helps identify the certificate.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * An date when the certificate becomes inactive
     */
    public /*out*/ readonly inactiveDate!: pulumi.Output<string>;
    /**
     * The private key associated with the certificate being imported.
     */
    public readonly privateKey!: pulumi.Output<string | undefined>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Specifies if a certificate is being used for signing or encryption. The valid values are SIGNING and ENCRYPTION.
     */
    public readonly usage!: pulumi.Output<string>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateArgs | CertificateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CertificateState | undefined;
            resourceInputs["activeDate"] = state ? state.activeDate : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["certificateChain"] = state ? state.certificateChain : undefined;
            resourceInputs["certificateId"] = state ? state.certificateId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["inactiveDate"] = state ? state.inactiveDate : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["usage"] = state ? state.usage : undefined;
        } else {
            const args = argsOrState as CertificateArgs | undefined;
            if ((!args || args.certificate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificate'");
            }
            if ((!args || args.usage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'usage'");
            }
            resourceInputs["certificate"] = args?.certificate ? pulumi.secret(args.certificate) : undefined;
            resourceInputs["certificateChain"] = args?.certificateChain ? pulumi.secret(args.certificateChain) : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["privateKey"] = args?.privateKey ? pulumi.secret(args.privateKey) : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["usage"] = args ? args.usage : undefined;
            resourceInputs["activeDate"] = undefined /*out*/;
            resourceInputs["certificateId"] = undefined /*out*/;
            resourceInputs["inactiveDate"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["certificate", "certificateChain", "privateKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Certificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Certificate resources.
 */
export interface CertificateState {
    /**
     * An date when the certificate becomes active
     */
    activeDate?: pulumi.Input<string>;
    /**
     * The valid certificate file required for the transfer.
     */
    certificate?: pulumi.Input<string>;
    /**
     * The optional list of certificate that make up the chain for the certificate that is being imported.
     */
    certificateChain?: pulumi.Input<string>;
    /**
     * The unique identifier for the AS2 certificate
     */
    certificateId?: pulumi.Input<string>;
    /**
     * A short description that helps identify the certificate.
     */
    description?: pulumi.Input<string>;
    /**
     * An date when the certificate becomes inactive
     */
    inactiveDate?: pulumi.Input<string>;
    /**
     * The private key associated with the certificate being imported.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies if a certificate is being used for signing or encryption. The valid values are SIGNING and ENCRYPTION.
     */
    usage?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * The valid certificate file required for the transfer.
     */
    certificate: pulumi.Input<string>;
    /**
     * The optional list of certificate that make up the chain for the certificate that is being imported.
     */
    certificateChain?: pulumi.Input<string>;
    /**
     * A short description that helps identify the certificate.
     */
    description?: pulumi.Input<string>;
    /**
     * The private key associated with the certificate being imported.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies if a certificate is being used for signing or encryption. The valid values are SIGNING and ENCRYPTION.
     */
    usage: pulumi.Input<string>;
}
