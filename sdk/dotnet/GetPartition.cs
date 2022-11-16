// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws
{
    public static class GetPartition
    {
        /// <summary>
        /// Use this data source to lookup information about the current AWS partition in
        /// which the provider is working.
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
        ///     var current = Aws.GetPartition.Invoke();
        /// 
        ///     var s3Policy = Aws.Iam.GetPolicyDocument.Invoke(new()
        ///     {
        ///         Statements = new[]
        ///         {
        ///             new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
        ///             {
        ///                 Actions = new[]
        ///                 {
        ///                     "s3:ListBucket",
        ///                 },
        ///                 Resources = new[]
        ///                 {
        ///                     $"arn:{current.Apply(getPartitionResult =&gt; getPartitionResult.Partition)}:s3:::my-bucket",
        ///                 },
        ///                 Sid = "1",
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPartitionResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPartitionResult>("aws:index/getPartition:getPartition", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetPartitionResult
    {
        /// <summary>
        /// Base DNS domain name for the current partition (e.g., `amazonaws.com` in AWS Commercial, `amazonaws.com.cn` in AWS China).
        /// </summary>
        public readonly string DnsSuffix;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Identifier of the current partition (e.g., `aws` in AWS Commercial, `aws-cn` in AWS China).
        /// </summary>
        public readonly string Partition;
        /// <summary>
        /// Prefix of service names (e.g., `com.amazonaws` in AWS Commercial, `cn.com.amazonaws` in AWS China).
        /// </summary>
        public readonly string ReverseDnsPrefix;

        [OutputConstructor]
        private GetPartitionResult(
            string dnsSuffix,

            string id,

            string partition,

            string reverseDnsPrefix)
        {
            DnsSuffix = dnsSuffix;
            Id = id;
            Partition = partition;
            ReverseDnsPrefix = reverseDnsPrefix;
        }
    }
}
