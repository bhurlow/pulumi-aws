// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SecretsManager
{
    public static class GetSecret
    {
        /// <summary>
        /// Retrieve metadata information about a Secrets Manager secret. To retrieve a secret value, see the `aws.secretsmanager.SecretVersion` data source.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### ARN
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var by_arn = Aws.SecretsManager.GetSecret.Invoke(new()
        ///     {
        ///         Arn = "arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% example %}}
        /// ### Name
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var by_name = Aws.SecretsManager.GetSecret.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSecretResult> InvokeAsync(GetSecretArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecretResult>("aws:secretsmanager/getSecret:getSecret", args ?? new GetSecretArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieve metadata information about a Secrets Manager secret. To retrieve a secret value, see the `aws.secretsmanager.SecretVersion` data source.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### ARN
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var by_arn = Aws.SecretsManager.GetSecret.Invoke(new()
        ///     {
        ///         Arn = "arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% example %}}
        /// ### Name
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var by_name = Aws.SecretsManager.GetSecret.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSecretResult> Invoke(GetSecretInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSecretResult>("aws:secretsmanager/getSecret:getSecret", args ?? new GetSecretInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecretArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ARN of the secret to retrieve.
        /// </summary>
        [Input("arn")]
        public string? Arn { get; set; }

        /// <summary>
        /// Name of the secret to retrieve.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetSecretArgs()
        {
        }
        public static new GetSecretArgs Empty => new GetSecretArgs();
    }

    public sealed class GetSecretInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ARN of the secret to retrieve.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Name of the secret to retrieve.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetSecretInvokeArgs()
        {
        }
        public static new GetSecretInvokeArgs Empty => new GetSecretInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecretResult
    {
        /// <summary>
        /// ARN of the secret.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Description of the secret.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Key Management Service (KMS) Customer Master Key (CMK) associated with the secret.
        /// </summary>
        public readonly string KmsKeyId;
        public readonly string Name;
        /// <summary>
        /// Resource-based policy document that's attached to the secret.
        /// </summary>
        public readonly string Policy;
        /// <summary>
        /// Whether rotation is enabled or not.
        /// </summary>
        public readonly bool RotationEnabled;
        /// <summary>
        /// Rotation Lambda function ARN if rotation is enabled.
        /// </summary>
        public readonly string RotationLambdaArn;
        /// <summary>
        /// Rotation rules if rotation is enabled.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSecretRotationRuleResult> RotationRules;
        /// <summary>
        /// Tags of the secret.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetSecretResult(
            string arn,

            string description,

            string id,

            string kmsKeyId,

            string name,

            string policy,

            bool rotationEnabled,

            string rotationLambdaArn,

            ImmutableArray<Outputs.GetSecretRotationRuleResult> rotationRules,

            ImmutableDictionary<string, string> tags)
        {
            Arn = arn;
            Description = description;
            Id = id;
            KmsKeyId = kmsKeyId;
            Name = name;
            Policy = policy;
            RotationEnabled = rotationEnabled;
            RotationLambdaArn = rotationLambdaArn;
            RotationRules = rotationRules;
            Tags = tags;
        }
    }
}
