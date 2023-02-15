// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Lambda
{
    public static class GetFunction
    {
        /// <summary>
        /// Provides information about a Lambda Function.
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
        ///     var config = new Config();
        ///     var functionName = config.Require("functionName");
        ///     var existing = Aws.Lambda.GetFunction.Invoke(new()
        ///     {
        ///         FunctionName = functionName,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFunctionResult> InvokeAsync(GetFunctionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFunctionResult>("aws:lambda/getFunction:getFunction", args ?? new GetFunctionArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information about a Lambda Function.
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
        ///     var config = new Config();
        ///     var functionName = config.Require("functionName");
        ///     var existing = Aws.Lambda.GetFunction.Invoke(new()
        ///     {
        ///         FunctionName = functionName,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFunctionResult> Invoke(GetFunctionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFunctionResult>("aws:lambda/getFunction:getFunction", args ?? new GetFunctionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFunctionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the lambda function.
        /// </summary>
        [Input("functionName", required: true)]
        public string FunctionName { get; set; } = null!;

        /// <summary>
        /// Alias name or version number of the lambda functionE.g., `$LATEST`, `my-alias`, or `1`. When not included: the data source resolves to the most recent published version; if no published version exists: it resolves to the most recent unpublished version.
        /// </summary>
        [Input("qualifier")]
        public string? Qualifier { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetFunctionArgs()
        {
        }
        public static new GetFunctionArgs Empty => new GetFunctionArgs();
    }

    public sealed class GetFunctionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the lambda function.
        /// </summary>
        [Input("functionName", required: true)]
        public Input<string> FunctionName { get; set; } = null!;

        /// <summary>
        /// Alias name or version number of the lambda functionE.g., `$LATEST`, `my-alias`, or `1`. When not included: the data source resolves to the most recent published version; if no published version exists: it resolves to the most recent unpublished version.
        /// </summary>
        [Input("qualifier")]
        public Input<string>? Qualifier { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetFunctionInvokeArgs()
        {
        }
        public static new GetFunctionInvokeArgs Empty => new GetFunctionInvokeArgs();
    }


    [OutputType]
    public sealed class GetFunctionResult
    {
        /// <summary>
        /// Instruction set architecture for the Lambda function.
        /// </summary>
        public readonly ImmutableArray<string> Architectures;
        /// <summary>
        /// Unqualified (no `:QUALIFIER` or `:VERSION` suffix) ARN identifying your Lambda Function. See also `qualified_arn`.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// ARN for a Code Signing Configuration.
        /// </summary>
        public readonly string CodeSigningConfigArn;
        /// <summary>
        /// Configure the function's *dead letter queue*.
        /// </summary>
        public readonly Outputs.GetFunctionDeadLetterConfigResult DeadLetterConfig;
        /// <summary>
        /// Description of what your Lambda Function does.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Lambda environment's configuration settings.
        /// </summary>
        public readonly Outputs.GetFunctionEnvironmentResult Environment;
        /// <summary>
        /// Amount of Ephemeral storage(`/tmp`) allocated for the Lambda Function.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionEphemeralStorageResult> EphemeralStorages;
        /// <summary>
        /// Connection settings for an Amazon EFS file system.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionFileSystemConfigResult> FileSystemConfigs;
        public readonly string FunctionName;
        /// <summary>
        /// Function entrypoint in your code.
        /// </summary>
        public readonly string Handler;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// URI of the container image.
        /// </summary>
        public readonly string ImageUri;
        /// <summary>
        /// ARN to be used for invoking Lambda Function from API Gateway. **NOTE:** Starting with `v4.51.0` of the provider, this will *not* include the qualifier.
        /// </summary>
        public readonly string InvokeArn;
        /// <summary>
        /// ARN for the KMS encryption key.
        /// </summary>
        public readonly string KmsKeyArn;
        /// <summary>
        /// Date this resource was last modified.
        /// </summary>
        public readonly string LastModified;
        /// <summary>
        /// List of Lambda Layer ARNs attached to your Lambda Function.
        /// </summary>
        public readonly ImmutableArray<string> Layers;
        /// <summary>
        /// Amount of memory in MB your Lambda Function can use at runtime.
        /// </summary>
        public readonly int MemorySize;
        /// <summary>
        /// Qualified (`:QUALIFIER` or `:VERSION` suffix) ARN identifying your Lambda Function. See also `arn`.
        /// </summary>
        public readonly string QualifiedArn;
        /// <summary>
        /// Qualified (`:QUALIFIER` or `:VERSION` suffix) ARN to be used for invoking Lambda Function from API Gateway. See also `invoke_arn`.
        /// </summary>
        public readonly string QualifiedInvokeArn;
        public readonly string? Qualifier;
        /// <summary>
        /// The amount of reserved concurrent executions for this lambda function or `-1` if unreserved.
        /// </summary>
        public readonly int ReservedConcurrentExecutions;
        /// <summary>
        /// IAM role attached to the Lambda Function.
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// Runtime environment for the Lambda function.
        /// </summary>
        public readonly string Runtime;
        /// <summary>
        /// ARN of a signing job.
        /// </summary>
        public readonly string SigningJobArn;
        /// <summary>
        /// The ARN for a signing profile version.
        /// </summary>
        public readonly string SigningProfileVersionArn;
        /// <summary>
        /// Base64-encoded representation of raw SHA-256 sum of the zip file.
        /// </summary>
        public readonly string SourceCodeHash;
        /// <summary>
        /// Size in bytes of the function .zip file.
        /// </summary>
        public readonly int SourceCodeSize;
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// Function execution time at which Lambda should terminate the function.
        /// </summary>
        public readonly int Timeout;
        /// <summary>
        /// Tracing settings of the function.
        /// </summary>
        public readonly Outputs.GetFunctionTracingConfigResult TracingConfig;
        /// <summary>
        /// The version of the Lambda function returned. If `qualifier` is not set, this will resolve to the most recent published version. If no published version of the function exists, `version` will resolve to `$LATEST`.
        /// </summary>
        public readonly string Version;
        /// <summary>
        /// VPC configuration associated with your Lambda function.
        /// </summary>
        public readonly Outputs.GetFunctionVpcConfigResult VpcConfig;

        [OutputConstructor]
        private GetFunctionResult(
            ImmutableArray<string> architectures,

            string arn,

            string codeSigningConfigArn,

            Outputs.GetFunctionDeadLetterConfigResult deadLetterConfig,

            string description,

            Outputs.GetFunctionEnvironmentResult environment,

            ImmutableArray<Outputs.GetFunctionEphemeralStorageResult> ephemeralStorages,

            ImmutableArray<Outputs.GetFunctionFileSystemConfigResult> fileSystemConfigs,

            string functionName,

            string handler,

            string id,

            string imageUri,

            string invokeArn,

            string kmsKeyArn,

            string lastModified,

            ImmutableArray<string> layers,

            int memorySize,

            string qualifiedArn,

            string qualifiedInvokeArn,

            string? qualifier,

            int reservedConcurrentExecutions,

            string role,

            string runtime,

            string signingJobArn,

            string signingProfileVersionArn,

            string sourceCodeHash,

            int sourceCodeSize,

            ImmutableDictionary<string, string> tags,

            int timeout,

            Outputs.GetFunctionTracingConfigResult tracingConfig,

            string version,

            Outputs.GetFunctionVpcConfigResult vpcConfig)
        {
            Architectures = architectures;
            Arn = arn;
            CodeSigningConfigArn = codeSigningConfigArn;
            DeadLetterConfig = deadLetterConfig;
            Description = description;
            Environment = environment;
            EphemeralStorages = ephemeralStorages;
            FileSystemConfigs = fileSystemConfigs;
            FunctionName = functionName;
            Handler = handler;
            Id = id;
            ImageUri = imageUri;
            InvokeArn = invokeArn;
            KmsKeyArn = kmsKeyArn;
            LastModified = lastModified;
            Layers = layers;
            MemorySize = memorySize;
            QualifiedArn = qualifiedArn;
            QualifiedInvokeArn = qualifiedInvokeArn;
            Qualifier = qualifier;
            ReservedConcurrentExecutions = reservedConcurrentExecutions;
            Role = role;
            Runtime = runtime;
            SigningJobArn = signingJobArn;
            SigningProfileVersionArn = signingProfileVersionArn;
            SourceCodeHash = sourceCodeHash;
            SourceCodeSize = sourceCodeSize;
            Tags = tags;
            Timeout = timeout;
            TracingConfig = tracingConfig;
            Version = version;
            VpcConfig = vpcConfig;
        }
    }
}
