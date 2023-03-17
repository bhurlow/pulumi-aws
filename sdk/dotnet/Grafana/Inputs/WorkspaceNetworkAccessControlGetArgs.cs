// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Grafana.Inputs
{

    public sealed class WorkspaceNetworkAccessControlGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("prefixListIds", required: true)]
        private InputList<string>? _prefixListIds;

        /// <summary>
        /// An array of prefix list IDs.
        /// </summary>
        public InputList<string> PrefixListIds
        {
            get => _prefixListIds ?? (_prefixListIds = new InputList<string>());
            set => _prefixListIds = value;
        }

        [Input("vpceIds", required: true)]
        private InputList<string>? _vpceIds;

        /// <summary>
        /// An array of Amazon VPC endpoint IDs for the workspace. The only VPC endpoints that can be specified here are interface VPC endpoints for Grafana workspaces (using the com.amazonaws.[region].grafana-workspace service endpoint). Other VPC endpoints will be ignored.
        /// </summary>
        public InputList<string> VpceIds
        {
            get => _vpceIds ?? (_vpceIds = new InputList<string>());
            set => _vpceIds = value;
        }

        public WorkspaceNetworkAccessControlGetArgs()
        {
        }
        public static new WorkspaceNetworkAccessControlGetArgs Empty => new WorkspaceNetworkAccessControlGetArgs();
    }
}
