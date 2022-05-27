// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a CloudFront response headers policy resource.
 * A response headers policy contains information about a set of HTTP response headers and their values.
 * After you create a response headers policy, you can use its ID to attach it to one or more cache behaviors in a CloudFront distribution.
 * When it’s attached to a cache behavior, CloudFront adds the headers in the policy to every response that it sends for requests that match the cache behavior.
 *
 * ## Example Usage
 *
 * The example below creates a CloudFront response headers policy.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.cloudfront.ResponseHeadersPolicy("example", {
 *     comment: "test comment",
 *     corsConfig: {
 *         accessControlAllowCredentials: true,
 *         accessControlAllowHeaders: {
 *             items: ["test"],
 *         },
 *         accessControlAllowMethods: {
 *             items: ["GET"],
 *         },
 *         accessControlAllowOrigins: {
 *             items: ["test.example.comtest"],
 *         },
 *         originOverride: true,
 *     },
 * });
 * ```
 *
 * The example below creates a CloudFront response headers policy with a custom headers config.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.cloudfront.ResponseHeadersPolicy("example", {
 *     customHeadersConfig: {
 *         items: [
 *             {
 *                 header: "X-Permitted-Cross-Domain-Policies",
 *                 override: true,
 *                 value: "none",
 *             },
 *             {
 *                 header: "X-Test",
 *                 override: true,
 *                 value: "none",
 *             },
 *         ],
 *     },
 * });
 * ```
 *
 * The example below creates a CloudFront response headers policy with a custom headers config and server timing headers config.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.cloudfront.ResponseHeadersPolicy("example", {
 *     customHeadersConfig: {
 *         items: [{
 *             header: "X-Permitted-Cross-Domain-Policies",
 *             override: true,
 *             value: "none",
 *         }],
 *     },
 *     serverTimingHeadersConfig: {
 *         enabled: true,
 *         samplingRate: 50,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Cloudfront Response Headers Policies can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:cloudfront/responseHeadersPolicy:ResponseHeadersPolicy policy 658327ea-f89d-4fab-a63d-7e88639e58f9
 * ```
 */
export class ResponseHeadersPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ResponseHeadersPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResponseHeadersPolicyState, opts?: pulumi.CustomResourceOptions): ResponseHeadersPolicy {
        return new ResponseHeadersPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudfront/responseHeadersPolicy:ResponseHeadersPolicy';

    /**
     * Returns true if the given object is an instance of ResponseHeadersPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResponseHeadersPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResponseHeadersPolicy.__pulumiType;
    }

    /**
     * A comment to describe the response headers policy. The comment cannot be longer than 128 characters.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * A configuration for a set of HTTP response headers that are used for Cross-Origin Resource Sharing (CORS). See Cors Config for more information.
     */
    public readonly corsConfig!: pulumi.Output<outputs.cloudfront.ResponseHeadersPolicyCorsConfig | undefined>;
    /**
     * Object that contains an attribute `items` that contains a list of custom headers. See Custom Header for more information.
     */
    public readonly customHeadersConfig!: pulumi.Output<outputs.cloudfront.ResponseHeadersPolicyCustomHeadersConfig | undefined>;
    /**
     * The current version of the response headers policy.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * A unique name to identify the response headers policy.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A configuration for a set of security-related HTTP response headers. See Security Headers Config for more information.
     */
    public readonly securityHeadersConfig!: pulumi.Output<outputs.cloudfront.ResponseHeadersPolicySecurityHeadersConfig | undefined>;
    /**
     * A configuration for enabling the Server-Timing header in HTTP responses sent from CloudFront. See Server Timing Headers Config for more information.
     */
    public readonly serverTimingHeadersConfig!: pulumi.Output<outputs.cloudfront.ResponseHeadersPolicyServerTimingHeadersConfig | undefined>;

    /**
     * Create a ResponseHeadersPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ResponseHeadersPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResponseHeadersPolicyArgs | ResponseHeadersPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResponseHeadersPolicyState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["corsConfig"] = state ? state.corsConfig : undefined;
            resourceInputs["customHeadersConfig"] = state ? state.customHeadersConfig : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["securityHeadersConfig"] = state ? state.securityHeadersConfig : undefined;
            resourceInputs["serverTimingHeadersConfig"] = state ? state.serverTimingHeadersConfig : undefined;
        } else {
            const args = argsOrState as ResponseHeadersPolicyArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["corsConfig"] = args ? args.corsConfig : undefined;
            resourceInputs["customHeadersConfig"] = args ? args.customHeadersConfig : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["securityHeadersConfig"] = args ? args.securityHeadersConfig : undefined;
            resourceInputs["serverTimingHeadersConfig"] = args ? args.serverTimingHeadersConfig : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ResponseHeadersPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResponseHeadersPolicy resources.
 */
export interface ResponseHeadersPolicyState {
    /**
     * A comment to describe the response headers policy. The comment cannot be longer than 128 characters.
     */
    comment?: pulumi.Input<string>;
    /**
     * A configuration for a set of HTTP response headers that are used for Cross-Origin Resource Sharing (CORS). See Cors Config for more information.
     */
    corsConfig?: pulumi.Input<inputs.cloudfront.ResponseHeadersPolicyCorsConfig>;
    /**
     * Object that contains an attribute `items` that contains a list of custom headers. See Custom Header for more information.
     */
    customHeadersConfig?: pulumi.Input<inputs.cloudfront.ResponseHeadersPolicyCustomHeadersConfig>;
    /**
     * The current version of the response headers policy.
     */
    etag?: pulumi.Input<string>;
    /**
     * A unique name to identify the response headers policy.
     */
    name?: pulumi.Input<string>;
    /**
     * A configuration for a set of security-related HTTP response headers. See Security Headers Config for more information.
     */
    securityHeadersConfig?: pulumi.Input<inputs.cloudfront.ResponseHeadersPolicySecurityHeadersConfig>;
    /**
     * A configuration for enabling the Server-Timing header in HTTP responses sent from CloudFront. See Server Timing Headers Config for more information.
     */
    serverTimingHeadersConfig?: pulumi.Input<inputs.cloudfront.ResponseHeadersPolicyServerTimingHeadersConfig>;
}

/**
 * The set of arguments for constructing a ResponseHeadersPolicy resource.
 */
export interface ResponseHeadersPolicyArgs {
    /**
     * A comment to describe the response headers policy. The comment cannot be longer than 128 characters.
     */
    comment?: pulumi.Input<string>;
    /**
     * A configuration for a set of HTTP response headers that are used for Cross-Origin Resource Sharing (CORS). See Cors Config for more information.
     */
    corsConfig?: pulumi.Input<inputs.cloudfront.ResponseHeadersPolicyCorsConfig>;
    /**
     * Object that contains an attribute `items` that contains a list of custom headers. See Custom Header for more information.
     */
    customHeadersConfig?: pulumi.Input<inputs.cloudfront.ResponseHeadersPolicyCustomHeadersConfig>;
    /**
     * The current version of the response headers policy.
     */
    etag?: pulumi.Input<string>;
    /**
     * A unique name to identify the response headers policy.
     */
    name?: pulumi.Input<string>;
    /**
     * A configuration for a set of security-related HTTP response headers. See Security Headers Config for more information.
     */
    securityHeadersConfig?: pulumi.Input<inputs.cloudfront.ResponseHeadersPolicySecurityHeadersConfig>;
    /**
     * A configuration for enabling the Server-Timing header in HTTP responses sent from CloudFront. See Server Timing Headers Config for more information.
     */
    serverTimingHeadersConfig?: pulumi.Input<inputs.cloudfront.ResponseHeadersPolicyServerTimingHeadersConfig>;
}
