// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGatewayV2
{
    public static class GetApi
    {
        /// <summary>
        /// Provides details about a specific Amazon API Gateway Version 2 API.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Aws.ApiGatewayV2.GetApi.Invoke(new()
        ///     {
        ///         ApiId = "aabbccddee",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetApiResult> InvokeAsync(GetApiArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApiResult>("aws:apigatewayv2/getApi:getApi", args ?? new GetApiArgs(), options.WithDefaults());

        /// <summary>
        /// Provides details about a specific Amazon API Gateway Version 2 API.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Aws.ApiGatewayV2.GetApi.Invoke(new()
        ///     {
        ///         ApiId = "aabbccddee",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetApiResult> Invoke(GetApiInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetApiResult>("aws:apigatewayv2/getApi:getApi", args ?? new GetApiInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApiArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// API identifier.
        /// </summary>
        [Input("apiId", required: true)]
        public string ApiId { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// Map of resource tags.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetApiArgs()
        {
        }
        public static new GetApiArgs Empty => new GetApiArgs();
    }

    public sealed class GetApiInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// API identifier.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetApiInvokeArgs()
        {
        }
        public static new GetApiInvokeArgs Empty => new GetApiInvokeArgs();
    }


    [OutputType]
    public sealed class GetApiResult
    {
        /// <summary>
        /// URI of the API, of the form `https://{api-id}.execute-api.{region}.amazonaws.com` for HTTP APIs and `wss://{api-id}.execute-api.{region}.amazonaws.com` for WebSocket APIs.
        /// </summary>
        public readonly string ApiEndpoint;
        public readonly string ApiId;
        /// <summary>
        /// An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
        /// Applicable for WebSocket APIs.
        /// </summary>
        public readonly string ApiKeySelectionExpression;
        /// <summary>
        /// ARN of the API.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html).
        /// Applicable for HTTP APIs.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetApiCorsConfigurationResult> CorsConfigurations;
        /// <summary>
        /// Description of the API.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Whether clients can invoke the API by using the default `execute-api` endpoint.
        /// </summary>
        public readonly bool DisableExecuteApiEndpoint;
        /// <summary>
        /// ARN prefix to be used in an `aws.lambda.Permission`'s `source_arn` attribute
        /// or in an `aws.iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
        /// See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
        /// </summary>
        public readonly string ExecutionArn;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the API.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// API protocol.
        /// </summary>
        public readonly string ProtocolType;
        /// <summary>
        /// The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
        /// </summary>
        public readonly string RouteSelectionExpression;
        /// <summary>
        /// Map of resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// Version identifier for the API.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetApiResult(
            string apiEndpoint,

            string apiId,

            string apiKeySelectionExpression,

            string arn,

            ImmutableArray<Outputs.GetApiCorsConfigurationResult> corsConfigurations,

            string description,

            bool disableExecuteApiEndpoint,

            string executionArn,

            string id,

            string name,

            string protocolType,

            string routeSelectionExpression,

            ImmutableDictionary<string, string> tags,

            string version)
        {
            ApiEndpoint = apiEndpoint;
            ApiId = apiId;
            ApiKeySelectionExpression = apiKeySelectionExpression;
            Arn = arn;
            CorsConfigurations = corsConfigurations;
            Description = description;
            DisableExecuteApiEndpoint = disableExecuteApiEndpoint;
            ExecutionArn = executionArn;
            Id = id;
            Name = name;
            ProtocolType = protocolType;
            RouteSelectionExpression = routeSelectionExpression;
            Tags = tags;
            Version = version;
        }
    }
}
