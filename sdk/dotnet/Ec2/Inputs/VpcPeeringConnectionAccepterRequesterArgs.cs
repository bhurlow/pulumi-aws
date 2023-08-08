// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Inputs
{

    public sealed class VpcPeeringConnectionAccepterRequesterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether a local VPC can resolve public DNS hostnames to
        /// private IP addresses when queried from instances in a peer VPC.
        /// </summary>
        [Input("allowRemoteVpcDnsResolution")]
        public Input<bool>? AllowRemoteVpcDnsResolution { get; set; }

        public VpcPeeringConnectionAccepterRequesterArgs()
        {
        }
        public static new VpcPeeringConnectionAccepterRequesterArgs Empty => new VpcPeeringConnectionAccepterRequesterArgs();
    }
}
