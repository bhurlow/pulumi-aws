// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway
{
    /// <summary>
    /// Provides an HTTP Method Integration Response for an API Gateway Resource.
    /// 
    /// &gt; **Note:** Depends on having `aws.apigateway.Integration` inside your rest api. To ensure this
    /// you might need to add an explicit `depends_on` for clean runs.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_integration_response.html.markdown.
    /// </summary>
    public partial class IntegrationResponse : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
        /// </summary>
        [Output("contentHandling")]
        public Output<string?> ContentHandling { get; private set; } = null!;

        /// <summary>
        /// The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
        /// </summary>
        [Output("httpMethod")]
        public Output<string> HttpMethod { get; private set; } = null!;

        /// <summary>
        /// The API resource ID
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// A map of response parameters that can be read from the backend response.
        /// For example: `response_parameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }`
        /// </summary>
        [Output("responseParameters")]
        public Output<ImmutableDictionary<string, string>?> ResponseParameters { get; private set; } = null!;

        /// <summary>
        /// A map specifying the templates used to transform the integration response body
        /// </summary>
        [Output("responseTemplates")]
        public Output<ImmutableDictionary<string, string>?> ResponseTemplates { get; private set; } = null!;

        /// <summary>
        /// The ID of the associated REST API
        /// </summary>
        [Output("restApi")]
        public Output<string> RestApi { get; private set; } = null!;

        /// <summary>
        /// Specifies the regular expression pattern used to choose
        /// an integration response based on the response from the backend. Setting this to `-` makes the integration the default one.
        /// If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched.
        /// For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
        /// </summary>
        [Output("selectionPattern")]
        public Output<string?> SelectionPattern { get; private set; } = null!;

        /// <summary>
        /// The HTTP status code
        /// </summary>
        [Output("statusCode")]
        public Output<string> StatusCode { get; private set; } = null!;


        /// <summary>
        /// Create a IntegrationResponse resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IntegrationResponse(string name, IntegrationResponseArgs args, CustomResourceOptions? options = null)
            : base("aws:apigateway/integrationResponse:IntegrationResponse", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private IntegrationResponse(string name, Input<string> id, IntegrationResponseState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigateway/integrationResponse:IntegrationResponse", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IntegrationResponse resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IntegrationResponse Get(string name, Input<string> id, IntegrationResponseState? state = null, CustomResourceOptions? options = null)
        {
            return new IntegrationResponse(name, id, state, options);
        }
    }

    public sealed class IntegrationResponseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
        /// </summary>
        [Input("contentHandling")]
        public Input<string>? ContentHandling { get; set; }

        /// <summary>
        /// The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
        /// </summary>
        [Input("httpMethod", required: true)]
        public Input<string> HttpMethod { get; set; } = null!;

        /// <summary>
        /// The API resource ID
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        [Input("responseParameters")]
        private InputMap<string>? _responseParameters;

        /// <summary>
        /// A map of response parameters that can be read from the backend response.
        /// For example: `response_parameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }`
        /// </summary>
        public InputMap<string> ResponseParameters
        {
            get => _responseParameters ?? (_responseParameters = new InputMap<string>());
            set => _responseParameters = value;
        }

        [Input("responseTemplates")]
        private InputMap<string>? _responseTemplates;

        /// <summary>
        /// A map specifying the templates used to transform the integration response body
        /// </summary>
        public InputMap<string> ResponseTemplates
        {
            get => _responseTemplates ?? (_responseTemplates = new InputMap<string>());
            set => _responseTemplates = value;
        }

        /// <summary>
        /// The ID of the associated REST API
        /// </summary>
        [Input("restApi", required: true)]
        public Input<string> RestApi { get; set; } = null!;

        /// <summary>
        /// Specifies the regular expression pattern used to choose
        /// an integration response based on the response from the backend. Setting this to `-` makes the integration the default one.
        /// If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched.
        /// For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
        /// </summary>
        [Input("selectionPattern")]
        public Input<string>? SelectionPattern { get; set; }

        /// <summary>
        /// The HTTP status code
        /// </summary>
        [Input("statusCode", required: true)]
        public Input<string> StatusCode { get; set; } = null!;

        public IntegrationResponseArgs()
        {
        }
    }

    public sealed class IntegrationResponseState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
        /// </summary>
        [Input("contentHandling")]
        public Input<string>? ContentHandling { get; set; }

        /// <summary>
        /// The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
        /// </summary>
        [Input("httpMethod")]
        public Input<string>? HttpMethod { get; set; }

        /// <summary>
        /// The API resource ID
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        [Input("responseParameters")]
        private InputMap<string>? _responseParameters;

        /// <summary>
        /// A map of response parameters that can be read from the backend response.
        /// For example: `response_parameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }`
        /// </summary>
        public InputMap<string> ResponseParameters
        {
            get => _responseParameters ?? (_responseParameters = new InputMap<string>());
            set => _responseParameters = value;
        }

        [Input("responseTemplates")]
        private InputMap<string>? _responseTemplates;

        /// <summary>
        /// A map specifying the templates used to transform the integration response body
        /// </summary>
        public InputMap<string> ResponseTemplates
        {
            get => _responseTemplates ?? (_responseTemplates = new InputMap<string>());
            set => _responseTemplates = value;
        }

        /// <summary>
        /// The ID of the associated REST API
        /// </summary>
        [Input("restApi")]
        public Input<string>? RestApi { get; set; }

        /// <summary>
        /// Specifies the regular expression pattern used to choose
        /// an integration response based on the response from the backend. Setting this to `-` makes the integration the default one.
        /// If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched.
        /// For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
        /// </summary>
        [Input("selectionPattern")]
        public Input<string>? SelectionPattern { get; set; }

        /// <summary>
        /// The HTTP status code
        /// </summary>
        [Input("statusCode")]
        public Input<string>? StatusCode { get; set; }

        public IntegrationResponseState()
        {
        }
    }
}
