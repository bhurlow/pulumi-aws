// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {RestApi} from "./index";

/**
 * Provides an HTTP Method Response for an API Gateway Resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const myDemoAPI = new aws.apigateway.RestApi("myDemoAPI", {description: "This is my API for demonstration purposes"});
 * const myDemoResource = new aws.apigateway.Resource("myDemoResource", {
 *     restApi: myDemoAPI.id,
 *     parentId: myDemoAPI.rootResourceId,
 *     pathPart: "mydemoresource",
 * });
 * const myDemoMethod = new aws.apigateway.Method("myDemoMethod", {
 *     restApi: myDemoAPI.id,
 *     resourceId: myDemoResource.id,
 *     httpMethod: "GET",
 *     authorization: "NONE",
 * });
 * const myDemoIntegration = new aws.apigateway.Integration("myDemoIntegration", {
 *     restApi: myDemoAPI.id,
 *     resourceId: myDemoResource.id,
 *     httpMethod: myDemoMethod.httpMethod,
 *     type: "MOCK",
 * });
 * const response200 = new aws.apigateway.MethodResponse("response200", {
 *     restApi: myDemoAPI.id,
 *     resourceId: myDemoResource.id,
 *     httpMethod: myDemoMethod.httpMethod,
 *     statusCode: "200",
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_api_gateway_method_response.example
 *
 *  id = "12345abcde/67890fghij/GET/200" } Using `pulumi import`, import `aws_api_gateway_method_response` using `REST-API-ID/RESOURCE-ID/HTTP-METHOD/STATUS-CODE`. For exampleconsole % pulumi import aws_api_gateway_method_response.example 12345abcde/67890fghij/GET/200
 */
export class MethodResponse extends pulumi.CustomResource {
    /**
     * Get an existing MethodResponse resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MethodResponseState, opts?: pulumi.CustomResourceOptions): MethodResponse {
        return new MethodResponse(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigateway/methodResponse:MethodResponse';

    /**
     * Returns true if the given object is an instance of MethodResponse.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MethodResponse {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MethodResponse.__pulumiType;
    }

    /**
     * HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
     */
    public readonly httpMethod!: pulumi.Output<string>;
    /**
     * API resource ID
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * Map of the API models used for the response's content type
     */
    public readonly responseModels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of response parameters that can be sent to the caller.
     * For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
     * would define that the header `X-Some-Header` can be provided on the response.
     */
    public readonly responseParameters!: pulumi.Output<{[key: string]: boolean} | undefined>;
    /**
     * ID of the associated REST API
     */
    public readonly restApi!: pulumi.Output<string>;
    /**
     * HTTP status code
     */
    public readonly statusCode!: pulumi.Output<string>;

    /**
     * Create a MethodResponse resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MethodResponseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MethodResponseArgs | MethodResponseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MethodResponseState | undefined;
            resourceInputs["httpMethod"] = state ? state.httpMethod : undefined;
            resourceInputs["resourceId"] = state ? state.resourceId : undefined;
            resourceInputs["responseModels"] = state ? state.responseModels : undefined;
            resourceInputs["responseParameters"] = state ? state.responseParameters : undefined;
            resourceInputs["restApi"] = state ? state.restApi : undefined;
            resourceInputs["statusCode"] = state ? state.statusCode : undefined;
        } else {
            const args = argsOrState as MethodResponseArgs | undefined;
            if ((!args || args.httpMethod === undefined) && !opts.urn) {
                throw new Error("Missing required property 'httpMethod'");
            }
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            if ((!args || args.restApi === undefined) && !opts.urn) {
                throw new Error("Missing required property 'restApi'");
            }
            if ((!args || args.statusCode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'statusCode'");
            }
            resourceInputs["httpMethod"] = args ? args.httpMethod : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["responseModels"] = args ? args.responseModels : undefined;
            resourceInputs["responseParameters"] = args ? args.responseParameters : undefined;
            resourceInputs["restApi"] = args ? args.restApi : undefined;
            resourceInputs["statusCode"] = args ? args.statusCode : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MethodResponse.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MethodResponse resources.
 */
export interface MethodResponseState {
    /**
     * HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
     */
    httpMethod?: pulumi.Input<string>;
    /**
     * API resource ID
     */
    resourceId?: pulumi.Input<string>;
    /**
     * Map of the API models used for the response's content type
     */
    responseModels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of response parameters that can be sent to the caller.
     * For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
     * would define that the header `X-Some-Header` can be provided on the response.
     */
    responseParameters?: pulumi.Input<{[key: string]: pulumi.Input<boolean>}>;
    /**
     * ID of the associated REST API
     */
    restApi?: pulumi.Input<string | RestApi>;
    /**
     * HTTP status code
     */
    statusCode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MethodResponse resource.
 */
export interface MethodResponseArgs {
    /**
     * HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
     */
    httpMethod: pulumi.Input<string>;
    /**
     * API resource ID
     */
    resourceId: pulumi.Input<string>;
    /**
     * Map of the API models used for the response's content type
     */
    responseModels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of response parameters that can be sent to the caller.
     * For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
     * would define that the header `X-Some-Header` can be provided on the response.
     */
    responseParameters?: pulumi.Input<{[key: string]: pulumi.Input<boolean>}>;
    /**
     * ID of the associated REST API
     */
    restApi: pulumi.Input<string | RestApi>;
    /**
     * HTTP status code
     */
    statusCode: pulumi.Input<string>;
}
