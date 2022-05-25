// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker
{
    /// <summary>
    /// Provides a SageMaker Workforce resource.
    /// 
    /// ## Example Usage
    /// ### Cognito Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleUserPool = new Aws.Cognito.UserPool("exampleUserPool", new Aws.Cognito.UserPoolArgs
    ///         {
    ///         });
    ///         var exampleUserPoolClient = new Aws.Cognito.UserPoolClient("exampleUserPoolClient", new Aws.Cognito.UserPoolClientArgs
    ///         {
    ///             GenerateSecret = true,
    ///             UserPoolId = exampleUserPool.Id,
    ///         });
    ///         var exampleUserPoolDomain = new Aws.Cognito.UserPoolDomain("exampleUserPoolDomain", new Aws.Cognito.UserPoolDomainArgs
    ///         {
    ///             Domain = "example",
    ///             UserPoolId = exampleUserPool.Id,
    ///         });
    ///         var exampleWorkforce = new Aws.Sagemaker.Workforce("exampleWorkforce", new Aws.Sagemaker.WorkforceArgs
    ///         {
    ///             WorkforceName = "example",
    ///             CognitoConfig = new Aws.Sagemaker.Inputs.WorkforceCognitoConfigArgs
    ///             {
    ///                 ClientId = exampleUserPoolClient.Id,
    ///                 UserPool = exampleUserPoolDomain.UserPoolId,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Oidc Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Sagemaker.Workforce("example", new Aws.Sagemaker.WorkforceArgs
    ///         {
    ///             OidcConfig = new Aws.Sagemaker.Inputs.WorkforceOidcConfigArgs
    ///             {
    ///                 AuthorizationEndpoint = "https://example.com",
    ///                 ClientId = "example",
    ///                 ClientSecret = "example",
    ///                 Issuer = "https://example.com",
    ///                 JwksUri = "https://example.com",
    ///                 LogoutEndpoint = "https://example.com",
    ///                 TokenEndpoint = "https://example.com",
    ///                 UserInfoEndpoint = "https://example.com",
    ///             },
    ///             WorkforceName = "example",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// SageMaker Workforces can be imported using the `workforce_name`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:sagemaker/workforce:Workforce example example
    /// ```
    /// </summary>
    [AwsResourceType("aws:sagemaker/workforce:Workforce")]
    public partial class Workforce : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) assigned by AWS to this Workforce.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Use this parameter to configure an Amazon Cognito private workforce. A single Cognito workforce is created using and corresponds to a single Amazon Cognito user pool. Conflicts with `oidc_config`. see Cognito Config details below.
        /// </summary>
        [Output("cognitoConfig")]
        public Output<Outputs.WorkforceCognitoConfig?> CognitoConfig { get; private set; } = null!;

        /// <summary>
        /// Use this parameter to configure a private workforce using your own OIDC Identity Provider. Conflicts with `cognito_config`. see OIDC Config details below.
        /// </summary>
        [Output("oidcConfig")]
        public Output<Outputs.WorkforceOidcConfig?> OidcConfig { get; private set; } = null!;

        /// <summary>
        /// A list of IP address ranges Used to create an allow list of IP addresses for a private workforce. By default, a workforce isn't restricted to specific IP addresses. see Source Ip Config details below.
        /// </summary>
        [Output("sourceIpConfig")]
        public Output<Outputs.WorkforceSourceIpConfig> SourceIpConfig { get; private set; } = null!;

        /// <summary>
        /// The subdomain for your OIDC Identity Provider.
        /// </summary>
        [Output("subdomain")]
        public Output<string> Subdomain { get; private set; } = null!;

        /// <summary>
        /// The name of the Workforce (must be unique).
        /// </summary>
        [Output("workforceName")]
        public Output<string> WorkforceName { get; private set; } = null!;


        /// <summary>
        /// Create a Workforce resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Workforce(string name, WorkforceArgs args, CustomResourceOptions? options = null)
            : base("aws:sagemaker/workforce:Workforce", name, args ?? new WorkforceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Workforce(string name, Input<string> id, WorkforceState? state = null, CustomResourceOptions? options = null)
            : base("aws:sagemaker/workforce:Workforce", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Workforce resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Workforce Get(string name, Input<string> id, WorkforceState? state = null, CustomResourceOptions? options = null)
        {
            return new Workforce(name, id, state, options);
        }
    }

    public sealed class WorkforceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Use this parameter to configure an Amazon Cognito private workforce. A single Cognito workforce is created using and corresponds to a single Amazon Cognito user pool. Conflicts with `oidc_config`. see Cognito Config details below.
        /// </summary>
        [Input("cognitoConfig")]
        public Input<Inputs.WorkforceCognitoConfigArgs>? CognitoConfig { get; set; }

        /// <summary>
        /// Use this parameter to configure a private workforce using your own OIDC Identity Provider. Conflicts with `cognito_config`. see OIDC Config details below.
        /// </summary>
        [Input("oidcConfig")]
        public Input<Inputs.WorkforceOidcConfigArgs>? OidcConfig { get; set; }

        /// <summary>
        /// A list of IP address ranges Used to create an allow list of IP addresses for a private workforce. By default, a workforce isn't restricted to specific IP addresses. see Source Ip Config details below.
        /// </summary>
        [Input("sourceIpConfig")]
        public Input<Inputs.WorkforceSourceIpConfigArgs>? SourceIpConfig { get; set; }

        /// <summary>
        /// The name of the Workforce (must be unique).
        /// </summary>
        [Input("workforceName", required: true)]
        public Input<string> WorkforceName { get; set; } = null!;

        public WorkforceArgs()
        {
        }
    }

    public sealed class WorkforceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) assigned by AWS to this Workforce.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Use this parameter to configure an Amazon Cognito private workforce. A single Cognito workforce is created using and corresponds to a single Amazon Cognito user pool. Conflicts with `oidc_config`. see Cognito Config details below.
        /// </summary>
        [Input("cognitoConfig")]
        public Input<Inputs.WorkforceCognitoConfigGetArgs>? CognitoConfig { get; set; }

        /// <summary>
        /// Use this parameter to configure a private workforce using your own OIDC Identity Provider. Conflicts with `cognito_config`. see OIDC Config details below.
        /// </summary>
        [Input("oidcConfig")]
        public Input<Inputs.WorkforceOidcConfigGetArgs>? OidcConfig { get; set; }

        /// <summary>
        /// A list of IP address ranges Used to create an allow list of IP addresses for a private workforce. By default, a workforce isn't restricted to specific IP addresses. see Source Ip Config details below.
        /// </summary>
        [Input("sourceIpConfig")]
        public Input<Inputs.WorkforceSourceIpConfigGetArgs>? SourceIpConfig { get; set; }

        /// <summary>
        /// The subdomain for your OIDC Identity Provider.
        /// </summary>
        [Input("subdomain")]
        public Input<string>? Subdomain { get; set; }

        /// <summary>
        /// The name of the Workforce (must be unique).
        /// </summary>
        [Input("workforceName")]
        public Input<string>? WorkforceName { get; set; }

        public WorkforceState()
        {
        }
    }
}
