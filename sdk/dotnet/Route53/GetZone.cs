// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Route53
{
    public static partial class Invokes
    {
        /// <summary>
        /// `aws.route53.Zone` provides details about a specific Route 53 Hosted Zone.
        /// 
        /// This data source allows to find a Hosted Zone ID given Hosted Zone name and certain search criteria.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/route53_zone.html.markdown.
        /// </summary>
        public static Task<GetZoneResult> GetZone(GetZoneArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetZoneResult>("aws:route53/getZone:getZone", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetZoneArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Hosted Zone name of the desired Hosted Zone.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Used with `name` field to get a private Hosted Zone.
        /// </summary>
        [Input("privateZone")]
        public Input<bool>? PrivateZone { get; set; }

        [Input("resourceRecordSetCount")]
        public Input<int>? ResourceRecordSetCount { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Used with `name` field. A mapping of tags, each pair of which must exactly match
        /// a pair on the desired Hosted Zone.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Used with `name` field to get a private Hosted Zone associated with the vpc_id (in this case, private_zone is not mandatory).
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The Hosted Zone id of the desired Hosted Zone.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public GetZoneArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetZoneResult
    {
        /// <summary>
        /// Caller Reference of the Hosted Zone.
        /// </summary>
        public readonly string CallerReference;
        /// <summary>
        /// The comment field of the Hosted Zone.
        /// </summary>
        public readonly string Comment;
        /// <summary>
        /// The description provided by the service that created the Hosted Zone (e.g. `arn:aws:servicediscovery:us-east-1:1234567890:namespace/ns-xxxxxxxxxxxxxxxx`).
        /// </summary>
        public readonly string LinkedServiceDescription;
        /// <summary>
        /// The service that created the Hosted Zone (e.g. `servicediscovery.amazonaws.com`).
        /// </summary>
        public readonly string LinkedServicePrincipal;
        public readonly string Name;
        /// <summary>
        /// The list of DNS name servers for the Hosted Zone.
        /// </summary>
        public readonly ImmutableArray<string> NameServers;
        public readonly bool? PrivateZone;
        /// <summary>
        /// The number of Record Set in the Hosted Zone.
        /// </summary>
        public readonly int ResourceRecordSetCount;
        public readonly ImmutableDictionary<string, object> Tags;
        public readonly string VpcId;
        public readonly string ZoneId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetZoneResult(
            string callerReference,
            string comment,
            string linkedServiceDescription,
            string linkedServicePrincipal,
            string name,
            ImmutableArray<string> nameServers,
            bool? privateZone,
            int resourceRecordSetCount,
            ImmutableDictionary<string, object> tags,
            string vpcId,
            string zoneId,
            string id)
        {
            CallerReference = callerReference;
            Comment = comment;
            LinkedServiceDescription = linkedServiceDescription;
            LinkedServicePrincipal = linkedServicePrincipal;
            Name = name;
            NameServers = nameServers;
            PrivateZone = privateZone;
            ResourceRecordSetCount = resourceRecordSetCount;
            Tags = tags;
            VpcId = vpcId;
            ZoneId = zoneId;
            Id = id;
        }
    }
}
