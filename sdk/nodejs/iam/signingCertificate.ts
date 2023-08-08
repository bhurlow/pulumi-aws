// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an IAM Signing Certificate resource to upload Signing Certificates.
 *
 * > **Note:** All arguments including the certificate body will be stored in the raw state as plain-text.
 * ## Example Usage
 *
 * **Using certs on file:**
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as fs from "fs";
 *
 * const testCert = new aws.iam.SigningCertificate("testCert", {
 *     username: "some_test_cert",
 *     certificateBody: fs.readFileSync("self-ca-cert.pem"),
 * });
 * ```
 *
 * **Example with cert in-line:**
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testCertAlt = new aws.iam.SigningCertificate("testCertAlt", {
 *     certificateBody: `-----BEGIN CERTIFICATE-----
 * [......] # cert contents
 * -----END CERTIFICATE-----
 *
 * `,
 *     username: "some_test_cert",
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_iam_signing_certificate.certificate
 *
 *  id = "IDIDIDIDID:user-name" } Using `pulumi import`, import IAM Signing Certificates using the `id`. For exampleconsole % pulumi import aws_iam_signing_certificate.certificate IDIDIDIDID:user-name
 */
export class SigningCertificate extends pulumi.CustomResource {
    /**
     * Get an existing SigningCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SigningCertificateState, opts?: pulumi.CustomResourceOptions): SigningCertificate {
        return new SigningCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:iam/signingCertificate:SigningCertificate';

    /**
     * Returns true if the given object is an instance of SigningCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SigningCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SigningCertificate.__pulumiType;
    }

    /**
     * The contents of the signing certificate in PEM-encoded format.
     */
    public readonly certificateBody!: pulumi.Output<string>;
    /**
     * The ID for the signing certificate.
     */
    public /*out*/ readonly certificateId!: pulumi.Output<string>;
    /**
     * The status you want to assign to the certificate. `Active` means that the certificate can be used for programmatic calls to Amazon Web Services `Inactive` means that the certificate cannot be used.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * The name of the user the signing certificate is for.
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a SigningCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SigningCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SigningCertificateArgs | SigningCertificateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SigningCertificateState | undefined;
            resourceInputs["certificateBody"] = state ? state.certificateBody : undefined;
            resourceInputs["certificateId"] = state ? state.certificateId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as SigningCertificateArgs | undefined;
            if ((!args || args.certificateBody === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateBody'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["certificateBody"] = args ? args.certificateBody : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["certificateId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SigningCertificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SigningCertificate resources.
 */
export interface SigningCertificateState {
    /**
     * The contents of the signing certificate in PEM-encoded format.
     */
    certificateBody?: pulumi.Input<string>;
    /**
     * The ID for the signing certificate.
     */
    certificateId?: pulumi.Input<string>;
    /**
     * The status you want to assign to the certificate. `Active` means that the certificate can be used for programmatic calls to Amazon Web Services `Inactive` means that the certificate cannot be used.
     */
    status?: pulumi.Input<string>;
    /**
     * The name of the user the signing certificate is for.
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SigningCertificate resource.
 */
export interface SigningCertificateArgs {
    /**
     * The contents of the signing certificate in PEM-encoded format.
     */
    certificateBody: pulumi.Input<string>;
    /**
     * The status you want to assign to the certificate. `Active` means that the certificate can be used for programmatic calls to Amazon Web Services `Inactive` means that the certificate cannot be used.
     */
    status?: pulumi.Input<string>;
    /**
     * The name of the user the signing certificate is for.
     */
    userName: pulumi.Input<string>;
}
