// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an SES domain DKIM generation resource.
 *
 * Domain ownership needs to be confirmed first using sesDomainIdentity Resource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleDomainIdentity = new aws.ses.DomainIdentity("exampleDomainIdentity", {domain: "example.com"});
 * const exampleDomainDkim = new aws.ses.DomainDkim("exampleDomainDkim", {domain: exampleDomainIdentity.domain});
 * const exampleAmazonsesDkimRecord: aws.route53.Record[] = [];
 * for (const range = {value: 0}; range.value < 3; range.value++) {
 *     exampleAmazonsesDkimRecord.push(new aws.route53.Record(`exampleAmazonsesDkimRecord-${range.value}`, {
 *         zoneId: "ABCDEFGHIJ123",
 *         name: pulumi.interpolate`${exampleDomainDkim.dkimTokens}._domainkey`,
 *         type: "CNAME",
 *         ttl: 600,
 *         records: [pulumi.interpolate`${exampleDomainDkim.dkimTokens}.dkim.amazonses.com`],
 *     }));
 * }
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_ses_domain_dkim.example
 *
 *  id = "example.com" } Using `pulumi import`, import DKIM tokens using the `domain` attribute. For exampleconsole % pulumi import aws_ses_domain_dkim.example example.com
 */
export class DomainDkim extends pulumi.CustomResource {
    /**
     * Get an existing DomainDkim resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainDkimState, opts?: pulumi.CustomResourceOptions): DomainDkim {
        return new DomainDkim(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ses/domainDkim:DomainDkim';

    /**
     * Returns true if the given object is an instance of DomainDkim.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainDkim {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainDkim.__pulumiType;
    }

    /**
     * DKIM tokens generated by SES.
     * These tokens should be used to create CNAME records used to verify SES Easy DKIM.
     * See below for an example of how this might be achieved
     * when the domain is hosted in Route 53 and managed by this provider.
     * Find out more about verifying domains in Amazon SES
     * in the [AWS SES docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim-dns-records.html).
     */
    public /*out*/ readonly dkimTokens!: pulumi.Output<string[]>;
    /**
     * Verified domain name to generate DKIM tokens for.
     */
    public readonly domain!: pulumi.Output<string>;

    /**
     * Create a DomainDkim resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainDkimArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainDkimArgs | DomainDkimState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainDkimState | undefined;
            resourceInputs["dkimTokens"] = state ? state.dkimTokens : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
        } else {
            const args = argsOrState as DomainDkimArgs | undefined;
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["dkimTokens"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DomainDkim.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainDkim resources.
 */
export interface DomainDkimState {
    /**
     * DKIM tokens generated by SES.
     * These tokens should be used to create CNAME records used to verify SES Easy DKIM.
     * See below for an example of how this might be achieved
     * when the domain is hosted in Route 53 and managed by this provider.
     * Find out more about verifying domains in Amazon SES
     * in the [AWS SES docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim-dns-records.html).
     */
    dkimTokens?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Verified domain name to generate DKIM tokens for.
     */
    domain?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DomainDkim resource.
 */
export interface DomainDkimArgs {
    /**
     * Verified domain name to generate DKIM tokens for.
     */
    domain: pulumi.Input<string>;
}
