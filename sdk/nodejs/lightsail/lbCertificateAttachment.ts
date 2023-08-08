// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Attaches a Lightsail Load Balancer Certificate to a Lightsail Load Balancer.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testLb = new aws.lightsail.Lb("testLb", {
 *     healthCheckPath: "/",
 *     instancePort: 80,
 *     tags: {
 *         foo: "bar",
 *     },
 * });
 * const testLbCertificate = new aws.lightsail.LbCertificate("testLbCertificate", {
 *     lbName: testLb.id,
 *     domainName: "test.com",
 * });
 * const testLbCertificateAttachment = new aws.lightsail.LbCertificateAttachment("testLbCertificateAttachment", {
 *     lbName: testLb.name,
 *     certificateName: testLbCertificate.name,
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_lightsail_lb_certificate_attachment.test
 *
 *  id = "example-load-balancer,example-certificate" } Using `pulumi import`, import `aws_lightsail_lb_certificate_attachment` using the name attribute. For exampleconsole % pulumi import aws_lightsail_lb_certificate_attachment.test example-load-balancer,example-certificate
 */
export class LbCertificateAttachment extends pulumi.CustomResource {
    /**
     * Get an existing LbCertificateAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LbCertificateAttachmentState, opts?: pulumi.CustomResourceOptions): LbCertificateAttachment {
        return new LbCertificateAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:lightsail/lbCertificateAttachment:LbCertificateAttachment';

    /**
     * Returns true if the given object is an instance of LbCertificateAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LbCertificateAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LbCertificateAttachment.__pulumiType;
    }

    /**
     * The name of your SSL/TLS certificate.
     */
    public readonly certificateName!: pulumi.Output<string>;
    /**
     * The name of the load balancer to which you want to associate the SSL/TLS certificate.
     */
    public readonly lbName!: pulumi.Output<string>;

    /**
     * Create a LbCertificateAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LbCertificateAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LbCertificateAttachmentArgs | LbCertificateAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LbCertificateAttachmentState | undefined;
            resourceInputs["certificateName"] = state ? state.certificateName : undefined;
            resourceInputs["lbName"] = state ? state.lbName : undefined;
        } else {
            const args = argsOrState as LbCertificateAttachmentArgs | undefined;
            if ((!args || args.certificateName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateName'");
            }
            if ((!args || args.lbName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lbName'");
            }
            resourceInputs["certificateName"] = args ? args.certificateName : undefined;
            resourceInputs["lbName"] = args ? args.lbName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LbCertificateAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LbCertificateAttachment resources.
 */
export interface LbCertificateAttachmentState {
    /**
     * The name of your SSL/TLS certificate.
     */
    certificateName?: pulumi.Input<string>;
    /**
     * The name of the load balancer to which you want to associate the SSL/TLS certificate.
     */
    lbName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LbCertificateAttachment resource.
 */
export interface LbCertificateAttachmentArgs {
    /**
     * The name of your SSL/TLS certificate.
     */
    certificateName: pulumi.Input<string>;
    /**
     * The name of the load balancer to which you want to associate the SSL/TLS certificate.
     */
    lbName: pulumi.Input<string>;
}
