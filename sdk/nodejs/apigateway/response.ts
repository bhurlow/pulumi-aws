// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an API Gateway Gateway Response for a REST API Gateway.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const main = new aws.apigateway.RestApi("main", {});
 * const test = new aws.apigateway.Response("test", {
 *     restApiId: main.id,
 *     statusCode: "401",
 *     responseType: "UNAUTHORIZED",
 *     responseTemplates: {
 *         "application/json": "{\"message\":$context.error.messageString}",
 *     },
 *     responseParameters: {
 *         "gatewayresponse.header.Authorization": "'Basic'",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_api_gateway_gateway_response.example
 *
 *  id = "12345abcde/UNAUTHORIZED" } Using `pulumi import`, import `aws_api_gateway_gateway_response` using `REST-API-ID/RESPONSE-TYPE`. For exampleconsole % pulumi import aws_api_gateway_gateway_response.example 12345abcde/UNAUTHORIZED
 */
export class Response extends pulumi.CustomResource {
    /**
     * Get an existing Response resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResponseState, opts?: pulumi.CustomResourceOptions): Response {
        return new Response(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigateway/response:Response';

    /**
     * Returns true if the given object is an instance of Response.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Response {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Response.__pulumiType;
    }

    /**
     * Map of parameters (paths, query strings and headers) of the Gateway Response.
     */
    public readonly responseParameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of templates used to transform the response body.
     */
    public readonly responseTemplates!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Response type of the associated GatewayResponse.
     */
    public readonly responseType!: pulumi.Output<string>;
    /**
     * String identifier of the associated REST API.
     */
    public readonly restApiId!: pulumi.Output<string>;
    /**
     * HTTP status code of the Gateway Response.
     */
    public readonly statusCode!: pulumi.Output<string | undefined>;

    /**
     * Create a Response resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResponseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResponseArgs | ResponseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResponseState | undefined;
            resourceInputs["responseParameters"] = state ? state.responseParameters : undefined;
            resourceInputs["responseTemplates"] = state ? state.responseTemplates : undefined;
            resourceInputs["responseType"] = state ? state.responseType : undefined;
            resourceInputs["restApiId"] = state ? state.restApiId : undefined;
            resourceInputs["statusCode"] = state ? state.statusCode : undefined;
        } else {
            const args = argsOrState as ResponseArgs | undefined;
            if ((!args || args.responseType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'responseType'");
            }
            if ((!args || args.restApiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'restApiId'");
            }
            resourceInputs["responseParameters"] = args ? args.responseParameters : undefined;
            resourceInputs["responseTemplates"] = args ? args.responseTemplates : undefined;
            resourceInputs["responseType"] = args ? args.responseType : undefined;
            resourceInputs["restApiId"] = args ? args.restApiId : undefined;
            resourceInputs["statusCode"] = args ? args.statusCode : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Response.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Response resources.
 */
export interface ResponseState {
    /**
     * Map of parameters (paths, query strings and headers) of the Gateway Response.
     */
    responseParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of templates used to transform the response body.
     */
    responseTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Response type of the associated GatewayResponse.
     */
    responseType?: pulumi.Input<string>;
    /**
     * String identifier of the associated REST API.
     */
    restApiId?: pulumi.Input<string>;
    /**
     * HTTP status code of the Gateway Response.
     */
    statusCode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Response resource.
 */
export interface ResponseArgs {
    /**
     * Map of parameters (paths, query strings and headers) of the Gateway Response.
     */
    responseParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of templates used to transform the response body.
     */
    responseTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Response type of the associated GatewayResponse.
     */
    responseType: pulumi.Input<string>;
    /**
     * String identifier of the associated REST API.
     */
    restApiId: pulumi.Input<string>;
    /**
     * HTTP status code of the Gateway Response.
     */
    statusCode?: pulumi.Input<string>;
}
