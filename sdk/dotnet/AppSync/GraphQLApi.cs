// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppSync
{
    /// <summary>
    /// Provides an AppSync GraphQL API.
    /// 
    /// ## Example Usage
    /// ### API Key Authentication
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.AppSync.GraphQLApi("example", new()
    ///     {
    ///         AuthenticationType = "API_KEY",
    ///     });
    /// 
    /// });
    /// ```
    /// ### AWS IAM Authentication
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.AppSync.GraphQLApi("example", new()
    ///     {
    ///         AuthenticationType = "AWS_IAM",
    ///     });
    /// 
    /// });
    /// ```
    /// ### AWS Cognito User Pool Authentication
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.AppSync.GraphQLApi("example", new()
    ///     {
    ///         AuthenticationType = "AMAZON_COGNITO_USER_POOLS",
    ///         UserPoolConfig = new Aws.AppSync.Inputs.GraphQLApiUserPoolConfigArgs
    ///         {
    ///             AwsRegion = data.Aws_region.Current.Name,
    ///             DefaultAction = "DENY",
    ///             UserPoolId = aws_cognito_user_pool.Example.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### OpenID Connect Authentication
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.AppSync.GraphQLApi("example", new()
    ///     {
    ///         AuthenticationType = "OPENID_CONNECT",
    ///         OpenidConnectConfig = new Aws.AppSync.Inputs.GraphQLApiOpenidConnectConfigArgs
    ///         {
    ///             Issuer = "https://example.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### AWS Lambda Authorizer Authentication
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.AppSync.GraphQLApi("example", new()
    ///     {
    ///         AuthenticationType = "AWS_LAMBDA",
    ///         LambdaAuthorizerConfig = new Aws.AppSync.Inputs.GraphQLApiLambdaAuthorizerConfigArgs
    ///         {
    ///             AuthorizerUri = "arn:aws:lambda:us-east-1:123456789012:function:custom_lambda_authorizer",
    ///         },
    ///     });
    /// 
    ///     var appsyncLambdaAuthorizer = new Aws.Lambda.Permission("appsyncLambdaAuthorizer", new()
    ///     {
    ///         Action = "lambda:InvokeFunction",
    ///         Function = "custom_lambda_authorizer",
    ///         Principal = "appsync.amazonaws.com",
    ///         SourceArn = example.Arn,
    ///     });
    /// 
    /// });
    /// ```
    /// ### With Multiple Authentication Providers
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.AppSync.GraphQLApi("example", new()
    ///     {
    ///         AdditionalAuthenticationProviders = new[]
    ///         {
    ///             new Aws.AppSync.Inputs.GraphQLApiAdditionalAuthenticationProviderArgs
    ///             {
    ///                 AuthenticationType = "AWS_IAM",
    ///             },
    ///         },
    ///         AuthenticationType = "API_KEY",
    ///     });
    /// 
    /// });
    /// ```
    /// ### With Schema
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.AppSync.GraphQLApi("example", new()
    ///     {
    ///         AuthenticationType = "AWS_IAM",
    ///         Schema = @"schema {
    /// 	query: Query
    /// }
    /// type Query {
    ///   test: Int
    /// }
    /// 
    /// ",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Enabling Logging
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var assumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
    ///     {
    ///         Statements = new[]
    ///         {
    ///             new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
    ///             {
    ///                 Effect = "Allow",
    ///                 Principals = new[]
    ///                 {
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
    ///                     {
    ///                         Type = "Service",
    ///                         Identifiers = new[]
    ///                         {
    ///                             "appsync.amazonaws.com",
    ///                         },
    ///                     },
    ///                 },
    ///                 Actions = new[]
    ///                 {
    ///                     "sts:AssumeRole",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleRole = new Aws.Iam.Role("exampleRole", new()
    ///     {
    ///         AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult =&gt; getPolicyDocumentResult.Json),
    ///     });
    /// 
    ///     var exampleRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("exampleRolePolicyAttachment", new()
    ///     {
    ///         PolicyArn = "arn:aws:iam::aws:policy/service-role/AWSAppSyncPushToCloudWatchLogs",
    ///         Role = exampleRole.Name,
    ///     });
    /// 
    ///     // ... other configuration ...
    ///     var exampleGraphQLApi = new Aws.AppSync.GraphQLApi("exampleGraphQLApi", new()
    ///     {
    ///         LogConfig = new Aws.AppSync.Inputs.GraphQLApiLogConfigArgs
    ///         {
    ///             CloudwatchLogsRoleArn = exampleRole.Arn,
    ///             FieldLogLevel = "ERROR",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// AppSync GraphQL API can be imported using the GraphQL API ID, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:appsync/graphQLApi:GraphQLApi example 0123456789
    /// ```
    /// </summary>
    [AwsResourceType("aws:appsync/graphQLApi:GraphQLApi")]
    public partial class GraphQLApi : global::Pulumi.CustomResource
    {
        /// <summary>
        /// One or more additional authentication providers for the GraphqlApi. Defined below.
        /// </summary>
        [Output("additionalAuthenticationProviders")]
        public Output<ImmutableArray<Outputs.GraphQLApiAdditionalAuthenticationProvider>> AdditionalAuthenticationProviders { get; private set; } = null!;

        /// <summary>
        /// ARN
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`, `AWS_LAMBDA`
        /// </summary>
        [Output("authenticationType")]
        public Output<string> AuthenticationType { get; private set; } = null!;

        /// <summary>
        /// Nested argument containing Lambda authorizer configuration. Defined below.
        /// </summary>
        [Output("lambdaAuthorizerConfig")]
        public Output<Outputs.GraphQLApiLambdaAuthorizerConfig?> LambdaAuthorizerConfig { get; private set; } = null!;

        /// <summary>
        /// Nested argument containing logging configuration. Defined below.
        /// </summary>
        [Output("logConfig")]
        public Output<Outputs.GraphQLApiLogConfig?> LogConfig { get; private set; } = null!;

        /// <summary>
        /// User-supplied name for the GraphqlApi.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Nested argument containing OpenID Connect configuration. Defined below.
        /// </summary>
        [Output("openidConnectConfig")]
        public Output<Outputs.GraphQLApiOpenidConnectConfig?> OpenidConnectConfig { get; private set; } = null!;

        /// <summary>
        /// Schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
        /// </summary>
        [Output("schema")]
        public Output<string?> Schema { get; private set; } = null!;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Map of URIs associated with the APIE.g., `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
        /// </summary>
        [Output("uris")]
        public Output<ImmutableDictionary<string, string>> Uris { get; private set; } = null!;

        /// <summary>
        /// Amazon Cognito User Pool configuration. Defined below.
        /// </summary>
        [Output("userPoolConfig")]
        public Output<Outputs.GraphQLApiUserPoolConfig?> UserPoolConfig { get; private set; } = null!;

        /// <summary>
        /// Whether tracing with X-ray is enabled. Defaults to false.
        /// </summary>
        [Output("xrayEnabled")]
        public Output<bool?> XrayEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a GraphQLApi resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GraphQLApi(string name, GraphQLApiArgs args, CustomResourceOptions? options = null)
            : base("aws:appsync/graphQLApi:GraphQLApi", name, args ?? new GraphQLApiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GraphQLApi(string name, Input<string> id, GraphQLApiState? state = null, CustomResourceOptions? options = null)
            : base("aws:appsync/graphQLApi:GraphQLApi", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GraphQLApi resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GraphQLApi Get(string name, Input<string> id, GraphQLApiState? state = null, CustomResourceOptions? options = null)
        {
            return new GraphQLApi(name, id, state, options);
        }
    }

    public sealed class GraphQLApiArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalAuthenticationProviders")]
        private InputList<Inputs.GraphQLApiAdditionalAuthenticationProviderArgs>? _additionalAuthenticationProviders;

        /// <summary>
        /// One or more additional authentication providers for the GraphqlApi. Defined below.
        /// </summary>
        public InputList<Inputs.GraphQLApiAdditionalAuthenticationProviderArgs> AdditionalAuthenticationProviders
        {
            get => _additionalAuthenticationProviders ?? (_additionalAuthenticationProviders = new InputList<Inputs.GraphQLApiAdditionalAuthenticationProviderArgs>());
            set => _additionalAuthenticationProviders = value;
        }

        /// <summary>
        /// Authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`, `AWS_LAMBDA`
        /// </summary>
        [Input("authenticationType", required: true)]
        public Input<string> AuthenticationType { get; set; } = null!;

        /// <summary>
        /// Nested argument containing Lambda authorizer configuration. Defined below.
        /// </summary>
        [Input("lambdaAuthorizerConfig")]
        public Input<Inputs.GraphQLApiLambdaAuthorizerConfigArgs>? LambdaAuthorizerConfig { get; set; }

        /// <summary>
        /// Nested argument containing logging configuration. Defined below.
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.GraphQLApiLogConfigArgs>? LogConfig { get; set; }

        /// <summary>
        /// User-supplied name for the GraphqlApi.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Nested argument containing OpenID Connect configuration. Defined below.
        /// </summary>
        [Input("openidConnectConfig")]
        public Input<Inputs.GraphQLApiOpenidConnectConfigArgs>? OpenidConnectConfig { get; set; }

        /// <summary>
        /// Schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Amazon Cognito User Pool configuration. Defined below.
        /// </summary>
        [Input("userPoolConfig")]
        public Input<Inputs.GraphQLApiUserPoolConfigArgs>? UserPoolConfig { get; set; }

        /// <summary>
        /// Whether tracing with X-ray is enabled. Defaults to false.
        /// </summary>
        [Input("xrayEnabled")]
        public Input<bool>? XrayEnabled { get; set; }

        public GraphQLApiArgs()
        {
        }
        public static new GraphQLApiArgs Empty => new GraphQLApiArgs();
    }

    public sealed class GraphQLApiState : global::Pulumi.ResourceArgs
    {
        [Input("additionalAuthenticationProviders")]
        private InputList<Inputs.GraphQLApiAdditionalAuthenticationProviderGetArgs>? _additionalAuthenticationProviders;

        /// <summary>
        /// One or more additional authentication providers for the GraphqlApi. Defined below.
        /// </summary>
        public InputList<Inputs.GraphQLApiAdditionalAuthenticationProviderGetArgs> AdditionalAuthenticationProviders
        {
            get => _additionalAuthenticationProviders ?? (_additionalAuthenticationProviders = new InputList<Inputs.GraphQLApiAdditionalAuthenticationProviderGetArgs>());
            set => _additionalAuthenticationProviders = value;
        }

        /// <summary>
        /// ARN
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`, `AWS_LAMBDA`
        /// </summary>
        [Input("authenticationType")]
        public Input<string>? AuthenticationType { get; set; }

        /// <summary>
        /// Nested argument containing Lambda authorizer configuration. Defined below.
        /// </summary>
        [Input("lambdaAuthorizerConfig")]
        public Input<Inputs.GraphQLApiLambdaAuthorizerConfigGetArgs>? LambdaAuthorizerConfig { get; set; }

        /// <summary>
        /// Nested argument containing logging configuration. Defined below.
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.GraphQLApiLogConfigGetArgs>? LogConfig { get; set; }

        /// <summary>
        /// User-supplied name for the GraphqlApi.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Nested argument containing OpenID Connect configuration. Defined below.
        /// </summary>
        [Input("openidConnectConfig")]
        public Input<Inputs.GraphQLApiOpenidConnectConfigGetArgs>? OpenidConnectConfig { get; set; }

        /// <summary>
        /// Schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        [Input("uris")]
        private InputMap<string>? _uris;

        /// <summary>
        /// Map of URIs associated with the APIE.g., `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
        /// </summary>
        public InputMap<string> Uris
        {
            get => _uris ?? (_uris = new InputMap<string>());
            set => _uris = value;
        }

        /// <summary>
        /// Amazon Cognito User Pool configuration. Defined below.
        /// </summary>
        [Input("userPoolConfig")]
        public Input<Inputs.GraphQLApiUserPoolConfigGetArgs>? UserPoolConfig { get; set; }

        /// <summary>
        /// Whether tracing with X-ray is enabled. Defaults to false.
        /// </summary>
        [Input("xrayEnabled")]
        public Input<bool>? XrayEnabled { get; set; }

        public GraphQLApiState()
        {
        }
        public static new GraphQLApiState Empty => new GraphQLApiState();
    }
}
