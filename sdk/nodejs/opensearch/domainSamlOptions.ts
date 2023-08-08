// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Manages SAML authentication options for an AWS OpenSearch Domain.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as fs from "fs";
 *
 * const exampleDomain = new aws.opensearch.Domain("exampleDomain", {
 *     engineVersion: "OpenSearch_1.1",
 *     clusterConfig: {
 *         instanceType: "r4.large.search",
 *     },
 *     snapshotOptions: {
 *         automatedSnapshotStartHour: 23,
 *     },
 *     tags: {
 *         Domain: "TestDomain",
 *     },
 * });
 * const exampleDomainSamlOptions = new aws.opensearch.DomainSamlOptions("exampleDomainSamlOptions", {
 *     domainName: exampleDomain.domainName,
 *     samlOptions: {
 *         enabled: true,
 *         idp: {
 *             entityId: "https://example.com",
 *             metadataContent: fs.readFileSync("./saml-metadata.xml"),
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_opensearch_domain_saml_options.example
 *
 *  id = "domain_name" } Using `pulumi import`, import OpenSearch domains using the `domain_name`. For exampleconsole % pulumi import aws_opensearch_domain_saml_options.example domain_name
 */
export class DomainSamlOptions extends pulumi.CustomResource {
    /**
     * Get an existing DomainSamlOptions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainSamlOptionsState, opts?: pulumi.CustomResourceOptions): DomainSamlOptions {
        return new DomainSamlOptions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:opensearch/domainSamlOptions:DomainSamlOptions';

    /**
     * Returns true if the given object is an instance of DomainSamlOptions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainSamlOptions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainSamlOptions.__pulumiType;
    }

    /**
     * Name of the domain.
     *
     * The following arguments are optional:
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * SAML authentication options for an AWS OpenSearch Domain.
     */
    public readonly samlOptions!: pulumi.Output<outputs.opensearch.DomainSamlOptionsSamlOptions | undefined>;

    /**
     * Create a DomainSamlOptions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainSamlOptionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainSamlOptionsArgs | DomainSamlOptionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainSamlOptionsState | undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["samlOptions"] = state ? state.samlOptions : undefined;
        } else {
            const args = argsOrState as DomainSamlOptionsArgs | undefined;
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["samlOptions"] = args ? args.samlOptions : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DomainSamlOptions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainSamlOptions resources.
 */
export interface DomainSamlOptionsState {
    /**
     * Name of the domain.
     *
     * The following arguments are optional:
     */
    domainName?: pulumi.Input<string>;
    /**
     * SAML authentication options for an AWS OpenSearch Domain.
     */
    samlOptions?: pulumi.Input<inputs.opensearch.DomainSamlOptionsSamlOptions>;
}

/**
 * The set of arguments for constructing a DomainSamlOptions resource.
 */
export interface DomainSamlOptionsArgs {
    /**
     * Name of the domain.
     *
     * The following arguments are optional:
     */
    domainName: pulumi.Input<string>;
    /**
     * SAML authentication options for an AWS OpenSearch Domain.
     */
    samlOptions?: pulumi.Input<inputs.opensearch.DomainSamlOptionsSamlOptions>;
}
