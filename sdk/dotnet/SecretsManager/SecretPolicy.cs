// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SecretsManager
{
    /// <summary>
    /// Provides a resource to manage AWS Secrets Manager secret policy.
    /// 
    /// ## Example Usage
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleSecret = new Aws.SecretsManager.Secret("exampleSecret");
    /// 
    ///     var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    ///     {
    ///         Statements = new[]
    ///         {
    ///             new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
    ///             {
    ///                 Sid = "EnableAnotherAWSAccountToReadTheSecret",
    ///                 Effect = "Allow",
    ///                 Principals = new[]
    ///                 {
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
    ///                     {
    ///                         Type = "AWS",
    ///                         Identifiers = new[]
    ///                         {
    ///                             "arn:aws:iam::123456789012:root",
    ///                         },
    ///                     },
    ///                 },
    ///                 Actions = new[]
    ///                 {
    ///                     "secretsmanager:GetSecretValue",
    ///                 },
    ///                 Resources = new[]
    ///                 {
    ///                     "*",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleSecretPolicy = new Aws.SecretsManager.SecretPolicy("exampleSecretPolicy", new()
    ///     {
    ///         SecretArn = exampleSecret.Arn,
    ///         Policy = examplePolicyDocument.Apply(getPolicyDocumentResult =&gt; getPolicyDocumentResult.Json),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_secretsmanager_secret_policy.example
    /// 
    ///  id = "arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456" } Using `pulumi import`, import `aws_secretsmanager_secret_policy` using the secret Amazon Resource Name (ARN). For exampleconsole % pulumi import aws_secretsmanager_secret_policy.example arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456
    /// </summary>
    [AwsResourceType("aws:secretsmanager/secretPolicy:SecretPolicy")]
    public partial class SecretPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Makes an optional API call to Zelkova to validate the Resource Policy to prevent broad access to your secret.
        /// </summary>
        [Output("blockPublicPolicy")]
        public Output<bool?> BlockPublicPolicy { get; private set; } = null!;

        /// <summary>
        /// Valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html). Unlike `aws.secretsmanager.Secret`, where `policy` can be set to `"{}"` to delete the policy, `"{}"` is not a valid policy since `policy` is required.
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;

        /// <summary>
        /// Secret ARN.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("secretArn")]
        public Output<string> SecretArn { get; private set; } = null!;


        /// <summary>
        /// Create a SecretPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretPolicy(string name, SecretPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:secretsmanager/secretPolicy:SecretPolicy", name, args ?? new SecretPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretPolicy(string name, Input<string> id, SecretPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:secretsmanager/secretPolicy:SecretPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretPolicy Get(string name, Input<string> id, SecretPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretPolicy(name, id, state, options);
        }
    }

    public sealed class SecretPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Makes an optional API call to Zelkova to validate the Resource Policy to prevent broad access to your secret.
        /// </summary>
        [Input("blockPublicPolicy")]
        public Input<bool>? BlockPublicPolicy { get; set; }

        /// <summary>
        /// Valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html). Unlike `aws.secretsmanager.Secret`, where `policy` can be set to `"{}"` to delete the policy, `"{}"` is not a valid policy since `policy` is required.
        /// </summary>
        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        /// <summary>
        /// Secret ARN.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("secretArn", required: true)]
        public Input<string> SecretArn { get; set; } = null!;

        public SecretPolicyArgs()
        {
        }
        public static new SecretPolicyArgs Empty => new SecretPolicyArgs();
    }

    public sealed class SecretPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Makes an optional API call to Zelkova to validate the Resource Policy to prevent broad access to your secret.
        /// </summary>
        [Input("blockPublicPolicy")]
        public Input<bool>? BlockPublicPolicy { get; set; }

        /// <summary>
        /// Valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html). Unlike `aws.secretsmanager.Secret`, where `policy` can be set to `"{}"` to delete the policy, `"{}"` is not a valid policy since `policy` is required.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// Secret ARN.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("secretArn")]
        public Input<string>? SecretArn { get; set; }

        public SecretPolicyState()
        {
        }
        public static new SecretPolicyState Empty => new SecretPolicyState();
    }
}
