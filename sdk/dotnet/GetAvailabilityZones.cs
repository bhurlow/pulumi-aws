// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws
{
    public static partial class Invokes
    {
        /// <summary>
        /// The Availability Zones data source allows access to the list of AWS
        /// Availability Zones which can be accessed by an AWS account within the region
        /// configured in the provider.
        /// 
        /// This is different from the `aws..getAvailabilityZone` (singular) data source,
        /// which provides some details about a specific availability zone.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/availability_zones.html.markdown.
        /// </summary>
        public static Task<GetAvailabilityZonesResult> GetAvailabilityZones(GetAvailabilityZonesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAvailabilityZonesResult>("aws:index/getAvailabilityZones:getAvailabilityZones", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetAvailabilityZonesArgs : Pulumi.ResourceArgs
    {
        [Input("blacklistedNames")]
        private InputList<string>? _blacklistedNames;

        /// <summary>
        /// List of blacklisted Availability Zone names.
        /// </summary>
        public InputList<string> BlacklistedNames
        {
            get => _blacklistedNames ?? (_blacklistedNames = new InputList<string>());
            set => _blacklistedNames = value;
        }

        [Input("blacklistedZoneIds")]
        private InputList<string>? _blacklistedZoneIds;

        /// <summary>
        /// List of blacklisted Availability Zone IDs.
        /// </summary>
        public InputList<string> BlacklistedZoneIds
        {
            get => _blacklistedZoneIds ?? (_blacklistedZoneIds = new InputList<string>());
            set => _blacklistedZoneIds = value;
        }

        /// <summary>
        /// Allows to filter list of Availability Zones based on their
        /// current state. Can be either `"available"`, `"information"`, `"impaired"` or
        /// `"unavailable"`. By default the list includes a complete set of Availability Zones
        /// to which the underlying AWS account has access, regardless of their state.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public GetAvailabilityZonesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetAvailabilityZonesResult
    {
        public readonly ImmutableArray<string> BlacklistedNames;
        public readonly ImmutableArray<string> BlacklistedZoneIds;
        /// <summary>
        /// A list of the Availability Zone names available to the account.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? State;
        /// <summary>
        /// A list of the Availability Zone IDs available to the account.
        /// </summary>
        public readonly ImmutableArray<string> ZoneIds;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAvailabilityZonesResult(
            ImmutableArray<string> blacklistedNames,
            ImmutableArray<string> blacklistedZoneIds,
            ImmutableArray<string> names,
            string? state,
            ImmutableArray<string> zoneIds,
            string id)
        {
            BlacklistedNames = blacklistedNames;
            BlacklistedZoneIds = blacklistedZoneIds;
            Names = names;
            State = state;
            ZoneIds = zoneIds;
            Id = id;
        }
    }
}
