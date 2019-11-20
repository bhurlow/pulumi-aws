// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway
{
    /// <summary>
    /// Provides an API Gateway Authorizer.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_authorizer.html.markdown.
    /// </summary>
    public partial class Authorizer : Pulumi.CustomResource
    {
        /// <summary>
        /// The credentials required for the authorizer.
        /// To specify an IAM Role for API Gateway to assume, use the IAM Role ARN.
        /// </summary>
        [Output("authorizerCredentials")]
        public Output<string?> AuthorizerCredentials { get; private set; } = null!;

        /// <summary>
        /// The TTL of cached authorizer results in seconds.
        /// Defaults to `300`.
        /// </summary>
        [Output("authorizerResultTtlInSeconds")]
        public Output<int?> AuthorizerResultTtlInSeconds { get; private set; } = null!;

        /// <summary>
        /// The authorizer's Uniform Resource Identifier (URI).
        /// This must be a well-formed Lambda function URI in the form of `arn:aws:apigateway:{region}:lambda:path/{service_api}`,
        /// e.g. `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:012345678912:function:my-function/invocations`
        /// </summary>
        [Output("authorizerUri")]
        public Output<string?> AuthorizerUri { get; private set; } = null!;

        /// <summary>
        /// The source of the identity in an incoming request.
        /// Defaults to `method.request.header.Authorization`. For `REQUEST` type, this may be a comma-separated list of values, including headers, query string parameters and stage variables - e.g. `"method.request.header.SomeHeaderName,method.request.querystring.SomeQueryStringName,stageVariables.SomeStageVariableName"`
        /// </summary>
        [Output("identitySource")]
        public Output<string?> IdentitySource { get; private set; } = null!;

        /// <summary>
        /// A validation expression for the incoming identity.
        /// For `TOKEN` type, this value should be a regular expression. The incoming token from the client is matched
        /// against this expression, and will proceed if the token matches. If the token doesn't match,
        /// the client receives a 401 Unauthorized response.
        /// </summary>
        [Output("identityValidationExpression")]
        public Output<string?> IdentityValidationExpression { get; private set; } = null!;

        /// <summary>
        /// The name of the authorizer
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of the Amazon Cognito user pool ARNs.
        /// Each element is of this format: `arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}`.
        /// </summary>
        [Output("providerArns")]
        public Output<ImmutableArray<string>> ProviderArns { get; private set; } = null!;

        /// <summary>
        /// The ID of the associated REST API
        /// </summary>
        [Output("restApi")]
        public Output<string> RestApi { get; private set; } = null!;

        /// <summary>
        /// The type of the authorizer. Possible values are `TOKEN` for a Lambda function using a single authorization token submitted in a custom header, `REQUEST` for a Lambda function using incoming request parameters, or `COGNITO_USER_POOLS` for using an Amazon Cognito user pool.
        /// Defaults to `TOKEN`.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Authorizer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Authorizer(string name, AuthorizerArgs args, CustomResourceOptions? options = null)
            : base("aws:apigateway/authorizer:Authorizer", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Authorizer(string name, Input<string> id, AuthorizerState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigateway/authorizer:Authorizer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Authorizer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Authorizer Get(string name, Input<string> id, AuthorizerState? state = null, CustomResourceOptions? options = null)
        {
            return new Authorizer(name, id, state, options);
        }
    }

    public sealed class AuthorizerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The credentials required for the authorizer.
        /// To specify an IAM Role for API Gateway to assume, use the IAM Role ARN.
        /// </summary>
        [Input("authorizerCredentials")]
        public Input<string>? AuthorizerCredentials { get; set; }

        /// <summary>
        /// The TTL of cached authorizer results in seconds.
        /// Defaults to `300`.
        /// </summary>
        [Input("authorizerResultTtlInSeconds")]
        public Input<int>? AuthorizerResultTtlInSeconds { get; set; }

        /// <summary>
        /// The authorizer's Uniform Resource Identifier (URI).
        /// This must be a well-formed Lambda function URI in the form of `arn:aws:apigateway:{region}:lambda:path/{service_api}`,
        /// e.g. `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:012345678912:function:my-function/invocations`
        /// </summary>
        [Input("authorizerUri")]
        public Input<string>? AuthorizerUri { get; set; }

        /// <summary>
        /// The source of the identity in an incoming request.
        /// Defaults to `method.request.header.Authorization`. For `REQUEST` type, this may be a comma-separated list of values, including headers, query string parameters and stage variables - e.g. `"method.request.header.SomeHeaderName,method.request.querystring.SomeQueryStringName,stageVariables.SomeStageVariableName"`
        /// </summary>
        [Input("identitySource")]
        public Input<string>? IdentitySource { get; set; }

        /// <summary>
        /// A validation expression for the incoming identity.
        /// For `TOKEN` type, this value should be a regular expression. The incoming token from the client is matched
        /// against this expression, and will proceed if the token matches. If the token doesn't match,
        /// the client receives a 401 Unauthorized response.
        /// </summary>
        [Input("identityValidationExpression")]
        public Input<string>? IdentityValidationExpression { get; set; }

        /// <summary>
        /// The name of the authorizer
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("providerArns")]
        private InputList<string>? _providerArns;

        /// <summary>
        /// A list of the Amazon Cognito user pool ARNs.
        /// Each element is of this format: `arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}`.
        /// </summary>
        public InputList<string> ProviderArns
        {
            get => _providerArns ?? (_providerArns = new InputList<string>());
            set => _providerArns = value;
        }

        /// <summary>
        /// The ID of the associated REST API
        /// </summary>
        [Input("restApi", required: true)]
        public Input<string> RestApi { get; set; } = null!;

        /// <summary>
        /// The type of the authorizer. Possible values are `TOKEN` for a Lambda function using a single authorization token submitted in a custom header, `REQUEST` for a Lambda function using incoming request parameters, or `COGNITO_USER_POOLS` for using an Amazon Cognito user pool.
        /// Defaults to `TOKEN`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AuthorizerArgs()
        {
        }
    }

    public sealed class AuthorizerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The credentials required for the authorizer.
        /// To specify an IAM Role for API Gateway to assume, use the IAM Role ARN.
        /// </summary>
        [Input("authorizerCredentials")]
        public Input<string>? AuthorizerCredentials { get; set; }

        /// <summary>
        /// The TTL of cached authorizer results in seconds.
        /// Defaults to `300`.
        /// </summary>
        [Input("authorizerResultTtlInSeconds")]
        public Input<int>? AuthorizerResultTtlInSeconds { get; set; }

        /// <summary>
        /// The authorizer's Uniform Resource Identifier (URI).
        /// This must be a well-formed Lambda function URI in the form of `arn:aws:apigateway:{region}:lambda:path/{service_api}`,
        /// e.g. `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:012345678912:function:my-function/invocations`
        /// </summary>
        [Input("authorizerUri")]
        public Input<string>? AuthorizerUri { get; set; }

        /// <summary>
        /// The source of the identity in an incoming request.
        /// Defaults to `method.request.header.Authorization`. For `REQUEST` type, this may be a comma-separated list of values, including headers, query string parameters and stage variables - e.g. `"method.request.header.SomeHeaderName,method.request.querystring.SomeQueryStringName,stageVariables.SomeStageVariableName"`
        /// </summary>
        [Input("identitySource")]
        public Input<string>? IdentitySource { get; set; }

        /// <summary>
        /// A validation expression for the incoming identity.
        /// For `TOKEN` type, this value should be a regular expression. The incoming token from the client is matched
        /// against this expression, and will proceed if the token matches. If the token doesn't match,
        /// the client receives a 401 Unauthorized response.
        /// </summary>
        [Input("identityValidationExpression")]
        public Input<string>? IdentityValidationExpression { get; set; }

        /// <summary>
        /// The name of the authorizer
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("providerArns")]
        private InputList<string>? _providerArns;

        /// <summary>
        /// A list of the Amazon Cognito user pool ARNs.
        /// Each element is of this format: `arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}`.
        /// </summary>
        public InputList<string> ProviderArns
        {
            get => _providerArns ?? (_providerArns = new InputList<string>());
            set => _providerArns = value;
        }

        /// <summary>
        /// The ID of the associated REST API
        /// </summary>
        [Input("restApi")]
        public Input<string>? RestApi { get; set; }

        /// <summary>
        /// The type of the authorizer. Possible values are `TOKEN` for a Lambda function using a single authorization token submitted in a custom header, `REQUEST` for a Lambda function using incoming request parameters, or `COGNITO_USER_POOLS` for using an Amazon Cognito user pool.
        /// Defaults to `TOKEN`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AuthorizerState()
        {
        }
    }
}
