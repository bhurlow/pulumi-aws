// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Route53
{
    public static class GetResolverFirewallConfig
    {
        /// <summary>
        /// `aws.route53.ResolverFirewallConfig` provides details about a specific a Route 53 Resolver DNS Firewall config.
        /// 
        /// This data source allows to find a details about a specific a Route 53 Resolver DNS Firewall config.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// The following example shows how to get a firewall config using the VPC ID.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Aws.Route53.GetResolverFirewallConfig.Invoke(new()
        ///     {
        ///         ResourceId = "vpc-exampleid",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetResolverFirewallConfigResult> InvokeAsync(GetResolverFirewallConfigArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResolverFirewallConfigResult>("aws:route53/getResolverFirewallConfig:getResolverFirewallConfig", args ?? new GetResolverFirewallConfigArgs(), options.WithDefaults());

        /// <summary>
        /// `aws.route53.ResolverFirewallConfig` provides details about a specific a Route 53 Resolver DNS Firewall config.
        /// 
        /// This data source allows to find a details about a specific a Route 53 Resolver DNS Firewall config.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// The following example shows how to get a firewall config using the VPC ID.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Aws.Route53.GetResolverFirewallConfig.Invoke(new()
        ///     {
        ///         ResourceId = "vpc-exampleid",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetResolverFirewallConfigResult> Invoke(GetResolverFirewallConfigInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetResolverFirewallConfigResult>("aws:route53/getResolverFirewallConfig:getResolverFirewallConfig", args ?? new GetResolverFirewallConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResolverFirewallConfigArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the VPC from Amazon VPC that the configuration is for.
        /// </summary>
        [Input("resourceId", required: true)]
        public string ResourceId { get; set; } = null!;

        public GetResolverFirewallConfigArgs()
        {
        }
        public static new GetResolverFirewallConfigArgs Empty => new GetResolverFirewallConfigArgs();
    }

    public sealed class GetResolverFirewallConfigInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the VPC from Amazon VPC that the configuration is for.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        public GetResolverFirewallConfigInvokeArgs()
        {
        }
        public static new GetResolverFirewallConfigInvokeArgs Empty => new GetResolverFirewallConfigInvokeArgs();
    }


    [OutputType]
    public sealed class GetResolverFirewallConfigResult
    {
        public readonly string FirewallFailOpen;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string OwnerId;
        public readonly string ResourceId;

        [OutputConstructor]
        private GetResolverFirewallConfigResult(
            string firewallFailOpen,

            string id,

            string ownerId,

            string resourceId)
        {
            FirewallFailOpen = firewallFailOpen;
            Id = id;
            OwnerId = ownerId;
            ResourceId = resourceId;
        }
    }
}
