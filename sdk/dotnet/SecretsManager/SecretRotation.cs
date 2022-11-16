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
    /// Provides a resource to manage AWS Secrets Manager secret rotation. To manage a secret, see the `aws.secretsmanager.Secret` resource. To manage a secret value, see the `aws.secretsmanager.SecretVersion` resource.
    /// 
    /// ## Example Usage
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.SecretsManager.SecretRotation("example", new()
    ///     {
    ///         SecretId = aws_secretsmanager_secret.Example.Id,
    ///         RotationLambdaArn = aws_lambda_function.Example.Arn,
    ///         RotationRules = new Aws.SecretsManager.Inputs.SecretRotationRotationRulesArgs
    ///         {
    ///             AutomaticallyAfterDays = 30,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Rotation Configuration
    /// 
    /// To enable automatic secret rotation, the Secrets Manager service requires usage of a Lambda function. The [Rotate Secrets section in the Secrets Manager User Guide](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets.html) provides additional information about deploying a prebuilt Lambda functions for supported credential rotation (e.g., RDS) or deploying a custom Lambda function.
    /// 
    /// &gt; **NOTE:** Configuring rotation causes the secret to rotate once as soon as you enable rotation. Before you do this, you must ensure that all of your applications that use the credentials stored in the secret are updated to retrieve the secret from AWS Secrets Manager. The old credentials might no longer be usable after the initial rotation and any applications that you fail to update will break as soon as the old credentials are no longer valid.
    /// 
    /// &gt; **NOTE:** If you cancel a rotation that is in progress (by removing the `rotation` configuration), it can leave the VersionStage labels in an unexpected state. Depending on what step of the rotation was in progress, you might need to remove the staging label AWSPENDING from the partially created version, specified by the SecretVersionId response value. You should also evaluate the partially rotated new version to see if it should be deleted, which you can do by removing all staging labels from the new version's VersionStage field.
    /// 
    /// ## Import
    /// 
    /// `aws_secretsmanager_secret_rotation` can be imported by using the secret Amazon Resource Name (ARN), e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:secretsmanager/secretRotation:SecretRotation example arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456
    /// ```
    /// </summary>
    [AwsResourceType("aws:secretsmanager/secretRotation:SecretRotation")]
    public partial class SecretRotation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether automatic rotation is enabled for this secret.
        /// </summary>
        [Output("rotationEnabled")]
        public Output<bool> RotationEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the ARN of the Lambda function that can rotate the secret.
        /// </summary>
        [Output("rotationLambdaArn")]
        public Output<string> RotationLambdaArn { get; private set; } = null!;

        /// <summary>
        /// A structure that defines the rotation configuration for this secret. Defined below.
        /// </summary>
        [Output("rotationRules")]
        public Output<Outputs.SecretRotationRotationRules> RotationRules { get; private set; } = null!;

        /// <summary>
        /// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
        /// </summary>
        [Output("secretId")]
        public Output<string> SecretId { get; private set; } = null!;


        /// <summary>
        /// Create a SecretRotation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretRotation(string name, SecretRotationArgs args, CustomResourceOptions? options = null)
            : base("aws:secretsmanager/secretRotation:SecretRotation", name, args ?? new SecretRotationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretRotation(string name, Input<string> id, SecretRotationState? state = null, CustomResourceOptions? options = null)
            : base("aws:secretsmanager/secretRotation:SecretRotation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretRotation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretRotation Get(string name, Input<string> id, SecretRotationState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretRotation(name, id, state, options);
        }
    }

    public sealed class SecretRotationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the ARN of the Lambda function that can rotate the secret.
        /// </summary>
        [Input("rotationLambdaArn", required: true)]
        public Input<string> RotationLambdaArn { get; set; } = null!;

        /// <summary>
        /// A structure that defines the rotation configuration for this secret. Defined below.
        /// </summary>
        [Input("rotationRules", required: true)]
        public Input<Inputs.SecretRotationRotationRulesArgs> RotationRules { get; set; } = null!;

        /// <summary>
        /// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
        /// </summary>
        [Input("secretId", required: true)]
        public Input<string> SecretId { get; set; } = null!;

        public SecretRotationArgs()
        {
        }
        public static new SecretRotationArgs Empty => new SecretRotationArgs();
    }

    public sealed class SecretRotationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether automatic rotation is enabled for this secret.
        /// </summary>
        [Input("rotationEnabled")]
        public Input<bool>? RotationEnabled { get; set; }

        /// <summary>
        /// Specifies the ARN of the Lambda function that can rotate the secret.
        /// </summary>
        [Input("rotationLambdaArn")]
        public Input<string>? RotationLambdaArn { get; set; }

        /// <summary>
        /// A structure that defines the rotation configuration for this secret. Defined below.
        /// </summary>
        [Input("rotationRules")]
        public Input<Inputs.SecretRotationRotationRulesGetArgs>? RotationRules { get; set; }

        /// <summary>
        /// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
        /// </summary>
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        public SecretRotationState()
        {
        }
        public static new SecretRotationState Empty => new SecretRotationState();
    }
}
