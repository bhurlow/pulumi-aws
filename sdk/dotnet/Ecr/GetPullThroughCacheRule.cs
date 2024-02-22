// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecr
{
    public static class GetPullThroughCacheRule
    {
        /// <summary>
        /// The ECR Pull Through Cache Rule data source allows the upstream registry URL and registry ID to be retrieved for a Pull Through Cache Rule.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ecrPublic = Aws.Ecr.GetPullThroughCacheRule.Invoke(new()
        ///     {
        ///         EcrRepositoryPrefix = "ecr-public",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPullThroughCacheRuleResult> InvokeAsync(GetPullThroughCacheRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPullThroughCacheRuleResult>("aws:ecr/getPullThroughCacheRule:getPullThroughCacheRule", args ?? new GetPullThroughCacheRuleArgs(), options.WithDefaults());

        /// <summary>
        /// The ECR Pull Through Cache Rule data source allows the upstream registry URL and registry ID to be retrieved for a Pull Through Cache Rule.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ecrPublic = Aws.Ecr.GetPullThroughCacheRule.Invoke(new()
        ///     {
        ///         EcrRepositoryPrefix = "ecr-public",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetPullThroughCacheRuleResult> Invoke(GetPullThroughCacheRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPullThroughCacheRuleResult>("aws:ecr/getPullThroughCacheRule:getPullThroughCacheRule", args ?? new GetPullThroughCacheRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPullThroughCacheRuleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The repository name prefix to use when caching images from the source registry.
        /// </summary>
        [Input("ecrRepositoryPrefix", required: true)]
        public string EcrRepositoryPrefix { get; set; } = null!;

        public GetPullThroughCacheRuleArgs()
        {
        }
        public static new GetPullThroughCacheRuleArgs Empty => new GetPullThroughCacheRuleArgs();
    }

    public sealed class GetPullThroughCacheRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The repository name prefix to use when caching images from the source registry.
        /// </summary>
        [Input("ecrRepositoryPrefix", required: true)]
        public Input<string> EcrRepositoryPrefix { get; set; } = null!;

        public GetPullThroughCacheRuleInvokeArgs()
        {
        }
        public static new GetPullThroughCacheRuleInvokeArgs Empty => new GetPullThroughCacheRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetPullThroughCacheRuleResult
    {
        /// <summary>
        /// ARN of the Secret which will be used to authenticate against the registry.
        /// </summary>
        public readonly string CredentialArn;
        public readonly string EcrRepositoryPrefix;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The registry ID where the repository was created.
        /// </summary>
        public readonly string RegistryId;
        /// <summary>
        /// The registry URL of the upstream public registry to use as the source.
        /// </summary>
        public readonly string UpstreamRegistryUrl;

        [OutputConstructor]
        private GetPullThroughCacheRuleResult(
            string credentialArn,

            string ecrRepositoryPrefix,

            string id,

            string registryId,

            string upstreamRegistryUrl)
        {
            CredentialArn = credentialArn;
            EcrRepositoryPrefix = ecrRepositoryPrefix;
            Id = id;
            RegistryId = registryId;
            UpstreamRegistryUrl = upstreamRegistryUrl;
        }
    }
}
