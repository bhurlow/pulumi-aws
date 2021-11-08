// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Eks
{
    public static class GetClusters
    {
        /// <summary>
        /// Retrieve EKS Clusters list
        /// </summary>
        public static Task<GetClustersResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClustersResult>("aws:eks/getClusters:getClusters", InvokeArgs.Empty, options.WithVersion());
    }


    [OutputType]
    public sealed class GetClustersResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Set of EKS clusters names
        /// </summary>
        public readonly ImmutableArray<string> Names;

        [OutputConstructor]
        private GetClustersResult(
            string id,

            ImmutableArray<string> names)
        {
            Id = id;
            Names = names;
        }
    }
}