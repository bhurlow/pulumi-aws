// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Grafana.Outputs
{

    [OutputType]
    public sealed class WorkspaceNetworkAccessControl
    {
        /// <summary>
        /// An array of prefix list IDs.
        /// </summary>
        public readonly ImmutableArray<string> PrefixListIds;
        /// <summary>
        /// An array of Amazon VPC endpoint IDs for the workspace. The only VPC endpoints that can be specified here are interface VPC endpoints for Grafana workspaces (using the com.amazonaws.[region].grafana-workspace service endpoint). Other VPC endpoints will be ignored.
        /// </summary>
        public readonly ImmutableArray<string> VpceIds;

        [OutputConstructor]
        private WorkspaceNetworkAccessControl(
            ImmutableArray<string> prefixListIds,

            ImmutableArray<string> vpceIds)
        {
            PrefixListIds = prefixListIds;
            VpceIds = vpceIds;
        }
    }
}
