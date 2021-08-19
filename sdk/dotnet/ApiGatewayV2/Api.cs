// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGatewayV2
{
    /// <summary>
    /// Manages an Amazon API Gateway Version 2 API.
    /// 
    /// &gt; **Note:** Amazon API Gateway Version 2 resources are used for creating and deploying WebSocket and HTTP APIs. To create and deploy REST APIs, use Amazon API Gateway Version 1.
    /// 
    /// ## Example Usage
    /// ### Basic WebSocket API
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.ApiGatewayV2.Api("example", new Aws.ApiGatewayV2.ApiArgs
    ///         {
    ///             ProtocolType = "WEBSOCKET",
    ///             RouteSelectionExpression = "$request.body.action",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Basic HTTP API
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.ApiGatewayV2.Api("example", new Aws.ApiGatewayV2.ApiArgs
    ///         {
    ///             ProtocolType = "HTTP",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// `aws_apigatewayv2_api` can be imported by using the API identifier, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:apigatewayv2/api:Api example aabbccddee
    /// ```
    /// </summary>
    [AwsResourceType("aws:apigatewayv2/api:Api")]
    public partial class Api : Pulumi.CustomResource
    {
        /// <summary>
        /// The URI of the API, of the form `https://{api-id}.execute-api.{region}.amazonaws.com` for HTTP APIs and `wss://{api-id}.execute-api.{region}.amazonaws.com` for WebSocket APIs.
        /// </summary>
        [Output("apiEndpoint")]
        public Output<string> ApiEndpoint { get; private set; } = null!;

        /// <summary>
        /// An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
        /// Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
        /// Applicable for WebSocket APIs.
        /// </summary>
        [Output("apiKeySelectionExpression")]
        public Output<string?> ApiKeySelectionExpression { get; private set; } = null!;

        /// <summary>
        /// The ARN of the API.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
        /// </summary>
        [Output("body")]
        public Output<string?> Body { get; private set; } = null!;

        /// <summary>
        /// The cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
        /// </summary>
        [Output("corsConfiguration")]
        public Output<Outputs.ApiCorsConfiguration?> CorsConfiguration { get; private set; } = null!;

        /// <summary>
        /// Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
        /// </summary>
        [Output("credentialsArn")]
        public Output<string?> CredentialsArn { get; private set; } = null!;

        /// <summary>
        /// The description of the API. Must be less than or equal to 1024 characters in length.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether clients can invoke the API by using the default `execute-api` endpoint.
        /// By default, clients can invoke the API with the default `{api_id}.execute-api.{region}.amazonaws.com endpoint`.
        /// To require that clients use a custom domain name to invoke the API, disable the default endpoint.
        /// </summary>
        [Output("disableExecuteApiEndpoint")]
        public Output<bool?> DisableExecuteApiEndpoint { get; private set; } = null!;

        /// <summary>
        /// The ARN prefix to be used in an `aws.lambda.Permission`'s `source_arn` attribute
        /// or in an `aws.iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
        /// See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
        /// </summary>
        [Output("executionArn")]
        public Output<string> ExecutionArn { get; private set; } = null!;

        /// <summary>
        /// Whether warnings should return an error while API Gateway is creating or updating the resource using an OpenAPI specification. Defaults to `false`. Applicable for HTTP APIs.
        /// </summary>
        [Output("failOnWarnings")]
        public Output<bool?> FailOnWarnings { get; private set; } = null!;

        /// <summary>
        /// The name of the API. Must be less than or equal to 128 characters in length.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The API protocol. Valid values: `HTTP`, `WEBSOCKET`.
        /// </summary>
        [Output("protocolType")]
        public Output<string> ProtocolType { get; private set; } = null!;

        /// <summary>
        /// Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
        /// </summary>
        [Output("routeKey")]
        public Output<string?> RouteKey { get; private set; } = null!;

        /// <summary>
        /// The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
        /// Defaults to `$request.method $request.path`.
        /// </summary>
        [Output("routeSelectionExpression")]
        public Output<string?> RouteSelectionExpression { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the API. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
        /// For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
        /// The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
        /// </summary>
        [Output("target")]
        public Output<string?> Target { get; private set; } = null!;

        /// <summary>
        /// A version identifier for the API. Must be between 1 and 64 characters in length.
        /// </summary>
        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Api resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Api(string name, ApiArgs args, CustomResourceOptions? options = null)
            : base("aws:apigatewayv2/api:Api", name, args ?? new ApiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Api(string name, Input<string> id, ApiState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigatewayv2/api:Api", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Api resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Api Get(string name, Input<string> id, ApiState? state = null, CustomResourceOptions? options = null)
        {
            return new Api(name, id, state, options);
        }
    }

    public sealed class ApiArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
        /// Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
        /// Applicable for WebSocket APIs.
        /// </summary>
        [Input("apiKeySelectionExpression")]
        public Input<string>? ApiKeySelectionExpression { get; set; }

        /// <summary>
        /// An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        /// <summary>
        /// The cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
        /// </summary>
        [Input("corsConfiguration")]
        public Input<Inputs.ApiCorsConfigurationArgs>? CorsConfiguration { get; set; }

        /// <summary>
        /// Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
        /// </summary>
        [Input("credentialsArn")]
        public Input<string>? CredentialsArn { get; set; }

        /// <summary>
        /// The description of the API. Must be less than or equal to 1024 characters in length.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether clients can invoke the API by using the default `execute-api` endpoint.
        /// By default, clients can invoke the API with the default `{api_id}.execute-api.{region}.amazonaws.com endpoint`.
        /// To require that clients use a custom domain name to invoke the API, disable the default endpoint.
        /// </summary>
        [Input("disableExecuteApiEndpoint")]
        public Input<bool>? DisableExecuteApiEndpoint { get; set; }

        /// <summary>
        /// Whether warnings should return an error while API Gateway is creating or updating the resource using an OpenAPI specification. Defaults to `false`. Applicable for HTTP APIs.
        /// </summary>
        [Input("failOnWarnings")]
        public Input<bool>? FailOnWarnings { get; set; }

        /// <summary>
        /// The name of the API. Must be less than or equal to 128 characters in length.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The API protocol. Valid values: `HTTP`, `WEBSOCKET`.
        /// </summary>
        [Input("protocolType", required: true)]
        public Input<string> ProtocolType { get; set; } = null!;

        /// <summary>
        /// Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
        /// </summary>
        [Input("routeKey")]
        public Input<string>? RouteKey { get; set; }

        /// <summary>
        /// The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
        /// Defaults to `$request.method $request.path`.
        /// </summary>
        [Input("routeSelectionExpression")]
        public Input<string>? RouteSelectionExpression { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the API. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
        /// For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
        /// The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        /// <summary>
        /// A version identifier for the API. Must be between 1 and 64 characters in length.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ApiArgs()
        {
        }
    }

    public sealed class ApiState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URI of the API, of the form `https://{api-id}.execute-api.{region}.amazonaws.com` for HTTP APIs and `wss://{api-id}.execute-api.{region}.amazonaws.com` for WebSocket APIs.
        /// </summary>
        [Input("apiEndpoint")]
        public Input<string>? ApiEndpoint { get; set; }

        /// <summary>
        /// An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
        /// Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
        /// Applicable for WebSocket APIs.
        /// </summary>
        [Input("apiKeySelectionExpression")]
        public Input<string>? ApiKeySelectionExpression { get; set; }

        /// <summary>
        /// The ARN of the API.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        /// <summary>
        /// The cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
        /// </summary>
        [Input("corsConfiguration")]
        public Input<Inputs.ApiCorsConfigurationGetArgs>? CorsConfiguration { get; set; }

        /// <summary>
        /// Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
        /// </summary>
        [Input("credentialsArn")]
        public Input<string>? CredentialsArn { get; set; }

        /// <summary>
        /// The description of the API. Must be less than or equal to 1024 characters in length.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether clients can invoke the API by using the default `execute-api` endpoint.
        /// By default, clients can invoke the API with the default `{api_id}.execute-api.{region}.amazonaws.com endpoint`.
        /// To require that clients use a custom domain name to invoke the API, disable the default endpoint.
        /// </summary>
        [Input("disableExecuteApiEndpoint")]
        public Input<bool>? DisableExecuteApiEndpoint { get; set; }

        /// <summary>
        /// The ARN prefix to be used in an `aws.lambda.Permission`'s `source_arn` attribute
        /// or in an `aws.iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
        /// See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
        /// </summary>
        [Input("executionArn")]
        public Input<string>? ExecutionArn { get; set; }

        /// <summary>
        /// Whether warnings should return an error while API Gateway is creating or updating the resource using an OpenAPI specification. Defaults to `false`. Applicable for HTTP APIs.
        /// </summary>
        [Input("failOnWarnings")]
        public Input<bool>? FailOnWarnings { get; set; }

        /// <summary>
        /// The name of the API. Must be less than or equal to 128 characters in length.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The API protocol. Valid values: `HTTP`, `WEBSOCKET`.
        /// </summary>
        [Input("protocolType")]
        public Input<string>? ProtocolType { get; set; }

        /// <summary>
        /// Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
        /// </summary>
        [Input("routeKey")]
        public Input<string>? RouteKey { get; set; }

        /// <summary>
        /// The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
        /// Defaults to `$request.method $request.path`.
        /// </summary>
        [Input("routeSelectionExpression")]
        public Input<string>? RouteSelectionExpression { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the API. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
        /// For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
        /// The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        /// <summary>
        /// A version identifier for the API. Must be between 1 and 64 characters in length.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ApiState()
        {
        }
    }
}
