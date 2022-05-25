// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Msk.Inputs
{

    public sealed class ClusterBrokerNodeGroupInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The distribution of broker nodes across availability zones ([documentation](https://docs.aws.amazon.com/msk/1.0/apireference/clusters.html#clusters-model-brokerazdistribution)). Currently the only valid value is `DEFAULT`.
        /// </summary>
        [Input("azDistribution")]
        public Input<string>? AzDistribution { get; set; }

        [Input("clientSubnets", required: true)]
        private InputList<string>? _clientSubnets;

        /// <summary>
        /// A list of subnets to connect to in client VPC ([documentation](https://docs.aws.amazon.com/msk/1.0/apireference/clusters.html#clusters-prop-brokernodegroupinfo-clientsubnets)).
        /// </summary>
        public InputList<string> ClientSubnets
        {
            get => _clientSubnets ?? (_clientSubnets = new InputList<string>());
            set => _clientSubnets = value;
        }

        /// <summary>
        /// Information about the cluster access configuration. See below. For security reasons, you can't turn on public access while creating an MSK cluster. However, you can update an existing cluster to make it publicly accessible. You can also create a new cluster and then update it to make it publicly accessible ([documentation](https://docs.aws.amazon.com/msk/latest/developerguide/public-access.html)).
        /// </summary>
        [Input("connectivityInfo")]
        public Input<Inputs.ClusterBrokerNodeGroupInfoConnectivityInfoArgs>? ConnectivityInfo { get; set; }

        /// <summary>
        /// The size in GiB of the EBS volume for the data drive on each broker node.
        /// </summary>
        [Input("ebsVolumeSize", required: true)]
        public Input<int> EbsVolumeSize { get; set; } = null!;

        /// <summary>
        /// Specify the instance type to use for the kafka brokersE.g., kafka.m5.large. ([Pricing info](https://aws.amazon.com/msk/pricing/))
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        [Input("securityGroups", required: true)]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// A list of the security groups to associate with the elastic network interfaces to control who can communicate with the cluster.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        public ClusterBrokerNodeGroupInfoArgs()
        {
        }
    }
}
