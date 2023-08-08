// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Use this data source to get the id and rootResourceId of a REST API in
 * API Gateway. To fetch the REST API you must provide a name to match against.
 * As there is no unique name constraint on REST APIs this data source will
 * error if there is more than one match.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const myRestApi = aws.apigateway.getRestApi({
 *     name: "my-rest-api",
 * });
 * ```
 */
export function getRestApi(args: GetRestApiArgs, opts?: pulumi.InvokeOptions): Promise<GetRestApiResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws:apigateway/getRestApi:getRestApi", {
        "name": args.name,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getRestApi.
 */
export interface GetRestApiArgs {
    /**
     * Name of the REST API to look up. If no REST API is found with this name, an error will be returned. If multiple REST APIs are found with this name, an error will be returned.
     */
    name: string;
    /**
     * Key-value map of resource tags.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getRestApi.
 */
export interface GetRestApiResult {
    /**
     * Source of the API key for requests.
     */
    readonly apiKeySource: string;
    /**
     * ARN of the REST API.
     */
    readonly arn: string;
    /**
     * List of binary media types supported by the REST API.
     */
    readonly binaryMediaTypes: string[];
    /**
     * Description of the REST API.
     */
    readonly description: string;
    /**
     * The endpoint configuration of this RestApi showing the endpoint types of the API.
     */
    readonly endpointConfigurations: outputs.apigateway.GetRestApiEndpointConfiguration[];
    /**
     * Execution ARN part to be used in `lambdaPermission`'s `sourceArn` when allowing API Gateway to invoke a Lambda function, e.g., `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j`, which can be concatenated with allowed stage, method and resource path.
     */
    readonly executionArn: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Minimum response size to compress for the REST API.
     */
    readonly minimumCompressionSize: string;
    readonly name: string;
    /**
     * JSON formatted policy document that controls access to the API Gateway.
     */
    readonly policy: string;
    /**
     * Set to the ID of the API Gateway Resource on the found REST API where the route matches '/'.
     */
    readonly rootResourceId: string;
    /**
     * Key-value map of resource tags.
     */
    readonly tags: {[key: string]: string};
}
/**
 * Use this data source to get the id and rootResourceId of a REST API in
 * API Gateway. To fetch the REST API you must provide a name to match against.
 * As there is no unique name constraint on REST APIs this data source will
 * error if there is more than one match.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const myRestApi = aws.apigateway.getRestApi({
 *     name: "my-rest-api",
 * });
 * ```
 */
export function getRestApiOutput(args: GetRestApiOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRestApiResult> {
    return pulumi.output(args).apply((a: any) => getRestApi(a, opts))
}

/**
 * A collection of arguments for invoking getRestApi.
 */
export interface GetRestApiOutputArgs {
    /**
     * Name of the REST API to look up. If no REST API is found with this name, an error will be returned. If multiple REST APIs are found with this name, an error will be returned.
     */
    name: pulumi.Input<string>;
    /**
     * Key-value map of resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
