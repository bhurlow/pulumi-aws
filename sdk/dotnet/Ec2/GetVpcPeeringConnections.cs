// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static class GetVpcPeeringConnections
    {
        /// <summary>
        /// Use this data source to get IDs of Amazon VPC peering connections
        /// To get more details on each connection, use the data resource `aws.ec2.VpcPeeringConnection`
        /// </summary>
        public static Task<GetVpcPeeringConnectionsResult> InvokeAsync(GetVpcPeeringConnectionsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcPeeringConnectionsResult>("aws:ec2/getVpcPeeringConnections:getVpcPeeringConnections", args ?? new GetVpcPeeringConnectionsArgs(), options.WithVersion());
    }


    public sealed class GetVpcPeeringConnectionsArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetVpcPeeringConnectionsFilterArgs>? _filters;

        /// <summary>
        /// Custom filter block as described below.
        /// </summary>
        public List<Inputs.GetVpcPeeringConnectionsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetVpcPeeringConnectionsFilterArgs>());
            set => _filters = value;
        }

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// A mapping of tags, each pair of which must exactly match
        /// a pair on the desired VPC Peering Connection.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetVpcPeeringConnectionsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVpcPeeringConnectionsResult
    {
        public readonly ImmutableArray<Outputs.GetVpcPeeringConnectionsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The IDs of the VPC Peering Connections.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetVpcPeeringConnectionsResult(
            ImmutableArray<Outputs.GetVpcPeeringConnectionsFilterResult> filters,

            string id,

            ImmutableArray<string> ids,

            ImmutableDictionary<string, string> tags)
        {
            Filters = filters;
            Id = id;
            Ids = ids;
            Tags = tags;
        }
    }
}
