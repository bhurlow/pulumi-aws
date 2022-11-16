// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Manages an Amazon API Gateway Version 2 API.
 *
 * > **Note:** Amazon API Gateway Version 2 resources are used for creating and deploying WebSocket and HTTP APIs. To create and deploy REST APIs, use Amazon API Gateway Version 1 resources.
 *
 * ## Example Usage
 * ### Basic WebSocket API
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.apigatewayv2.Api("example", {
 *     protocolType: "WEBSOCKET",
 *     routeSelectionExpression: "$request.body.action",
 * });
 * ```
 * ### Basic HTTP API
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.apigatewayv2.Api("example", {
 *     protocolType: "HTTP",
 * });
 * ```
 *
 * ## Import
 *
 * `aws_apigatewayv2_api` can be imported by using the API identifier, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:apigatewayv2/api:Api example aabbccddee
 * ```
 */
export class Api extends pulumi.CustomResource {
    /**
     * Get an existing Api resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiState, opts?: pulumi.CustomResourceOptions): Api {
        return new Api(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigatewayv2/api:Api';

    /**
     * Returns true if the given object is an instance of Api.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Api {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Api.__pulumiType;
    }

    /**
     * URI of the API, of the form `https://{api-id}.execute-api.{region}.amazonaws.com` for HTTP APIs and `wss://{api-id}.execute-api.{region}.amazonaws.com` for WebSocket APIs.
     */
    public /*out*/ readonly apiEndpoint!: pulumi.Output<string>;
    /**
     * An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
     * Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
     * Applicable for WebSocket APIs.
     */
    public readonly apiKeySelectionExpression!: pulumi.Output<string | undefined>;
    /**
     * ARN of the API.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
     */
    public readonly body!: pulumi.Output<string | undefined>;
    /**
     * Cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
     */
    public readonly corsConfiguration!: pulumi.Output<outputs.apigatewayv2.ApiCorsConfiguration | undefined>;
    /**
     * Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
     */
    public readonly credentialsArn!: pulumi.Output<string | undefined>;
    /**
     * Description of the API. Must be less than or equal to 1024 characters in length.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether clients can invoke the API by using the default `execute-api` endpoint.
     * By default, clients can invoke the API with the default `{api_id}.execute-api.{region}.amazonaws.com endpoint`.
     * To require that clients use a custom domain name to invoke the API, disable the default endpoint.
     */
    public readonly disableExecuteApiEndpoint!: pulumi.Output<boolean | undefined>;
    /**
     * ARN prefix to be used in an `aws.lambda.Permission`'s `sourceArn` attribute
     * or in an `aws.iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
     * See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
     */
    public /*out*/ readonly executionArn!: pulumi.Output<string>;
    /**
     * Whether warnings should return an error while API Gateway is creating or updating the resource using an OpenAPI specification. Defaults to `false`. Applicable for HTTP APIs.
     */
    public readonly failOnWarnings!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the API. Must be less than or equal to 128 characters in length.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * API protocol. Valid values: `HTTP`, `WEBSOCKET`.
     */
    public readonly protocolType!: pulumi.Output<string>;
    /**
     * Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
     */
    public readonly routeKey!: pulumi.Output<string | undefined>;
    /**
     * The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
     * Defaults to `$request.method $request.path`.
     */
    public readonly routeSelectionExpression!: pulumi.Output<string | undefined>;
    /**
     * Map of tags to assign to the API. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
     * For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
     * The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
     */
    public readonly target!: pulumi.Output<string | undefined>;
    /**
     * Version identifier for the API. Must be between 1 and 64 characters in length.
     */
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a Api resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiArgs | ApiState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApiState | undefined;
            resourceInputs["apiEndpoint"] = state ? state.apiEndpoint : undefined;
            resourceInputs["apiKeySelectionExpression"] = state ? state.apiKeySelectionExpression : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["body"] = state ? state.body : undefined;
            resourceInputs["corsConfiguration"] = state ? state.corsConfiguration : undefined;
            resourceInputs["credentialsArn"] = state ? state.credentialsArn : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disableExecuteApiEndpoint"] = state ? state.disableExecuteApiEndpoint : undefined;
            resourceInputs["executionArn"] = state ? state.executionArn : undefined;
            resourceInputs["failOnWarnings"] = state ? state.failOnWarnings : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protocolType"] = state ? state.protocolType : undefined;
            resourceInputs["routeKey"] = state ? state.routeKey : undefined;
            resourceInputs["routeSelectionExpression"] = state ? state.routeSelectionExpression : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["target"] = state ? state.target : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as ApiArgs | undefined;
            if ((!args || args.protocolType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocolType'");
            }
            resourceInputs["apiKeySelectionExpression"] = args ? args.apiKeySelectionExpression : undefined;
            resourceInputs["body"] = args ? args.body : undefined;
            resourceInputs["corsConfiguration"] = args ? args.corsConfiguration : undefined;
            resourceInputs["credentialsArn"] = args ? args.credentialsArn : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disableExecuteApiEndpoint"] = args ? args.disableExecuteApiEndpoint : undefined;
            resourceInputs["failOnWarnings"] = args ? args.failOnWarnings : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["protocolType"] = args ? args.protocolType : undefined;
            resourceInputs["routeKey"] = args ? args.routeKey : undefined;
            resourceInputs["routeSelectionExpression"] = args ? args.routeSelectionExpression : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["target"] = args ? args.target : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["apiEndpoint"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["executionArn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Api.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Api resources.
 */
export interface ApiState {
    /**
     * URI of the API, of the form `https://{api-id}.execute-api.{region}.amazonaws.com` for HTTP APIs and `wss://{api-id}.execute-api.{region}.amazonaws.com` for WebSocket APIs.
     */
    apiEndpoint?: pulumi.Input<string>;
    /**
     * An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
     * Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
     * Applicable for WebSocket APIs.
     */
    apiKeySelectionExpression?: pulumi.Input<string>;
    /**
     * ARN of the API.
     */
    arn?: pulumi.Input<string>;
    /**
     * An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
     */
    body?: pulumi.Input<string>;
    /**
     * Cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
     */
    corsConfiguration?: pulumi.Input<inputs.apigatewayv2.ApiCorsConfiguration>;
    /**
     * Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
     */
    credentialsArn?: pulumi.Input<string>;
    /**
     * Description of the API. Must be less than or equal to 1024 characters in length.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether clients can invoke the API by using the default `execute-api` endpoint.
     * By default, clients can invoke the API with the default `{api_id}.execute-api.{region}.amazonaws.com endpoint`.
     * To require that clients use a custom domain name to invoke the API, disable the default endpoint.
     */
    disableExecuteApiEndpoint?: pulumi.Input<boolean>;
    /**
     * ARN prefix to be used in an `aws.lambda.Permission`'s `sourceArn` attribute
     * or in an `aws.iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
     * See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
     */
    executionArn?: pulumi.Input<string>;
    /**
     * Whether warnings should return an error while API Gateway is creating or updating the resource using an OpenAPI specification. Defaults to `false`. Applicable for HTTP APIs.
     */
    failOnWarnings?: pulumi.Input<boolean>;
    /**
     * Name of the API. Must be less than or equal to 128 characters in length.
     */
    name?: pulumi.Input<string>;
    /**
     * API protocol. Valid values: `HTTP`, `WEBSOCKET`.
     */
    protocolType?: pulumi.Input<string>;
    /**
     * Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
     */
    routeKey?: pulumi.Input<string>;
    /**
     * The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
     * Defaults to `$request.method $request.path`.
     */
    routeSelectionExpression?: pulumi.Input<string>;
    /**
     * Map of tags to assign to the API. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
     * For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
     * The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
     */
    target?: pulumi.Input<string>;
    /**
     * Version identifier for the API. Must be between 1 and 64 characters in length.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Api resource.
 */
export interface ApiArgs {
    /**
     * An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
     * Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
     * Applicable for WebSocket APIs.
     */
    apiKeySelectionExpression?: pulumi.Input<string>;
    /**
     * An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
     */
    body?: pulumi.Input<string>;
    /**
     * Cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
     */
    corsConfiguration?: pulumi.Input<inputs.apigatewayv2.ApiCorsConfiguration>;
    /**
     * Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
     */
    credentialsArn?: pulumi.Input<string>;
    /**
     * Description of the API. Must be less than or equal to 1024 characters in length.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether clients can invoke the API by using the default `execute-api` endpoint.
     * By default, clients can invoke the API with the default `{api_id}.execute-api.{region}.amazonaws.com endpoint`.
     * To require that clients use a custom domain name to invoke the API, disable the default endpoint.
     */
    disableExecuteApiEndpoint?: pulumi.Input<boolean>;
    /**
     * Whether warnings should return an error while API Gateway is creating or updating the resource using an OpenAPI specification. Defaults to `false`. Applicable for HTTP APIs.
     */
    failOnWarnings?: pulumi.Input<boolean>;
    /**
     * Name of the API. Must be less than or equal to 128 characters in length.
     */
    name?: pulumi.Input<string>;
    /**
     * API protocol. Valid values: `HTTP`, `WEBSOCKET`.
     */
    protocolType: pulumi.Input<string>;
    /**
     * Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
     */
    routeKey?: pulumi.Input<string>;
    /**
     * The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
     * Defaults to `$request.method $request.path`.
     */
    routeSelectionExpression?: pulumi.Input<string>;
    /**
     * Map of tags to assign to the API. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
     * For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
     * The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
     */
    target?: pulumi.Input<string>;
    /**
     * Version identifier for the API. Must be between 1 and 64 characters in length.
     */
    version?: pulumi.Input<string>;
}
