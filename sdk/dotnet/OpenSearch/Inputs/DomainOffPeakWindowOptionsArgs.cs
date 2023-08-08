// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.OpenSearch.Inputs
{

    public sealed class DomainOffPeakWindowOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enabled disabled toggle for off-peak update window.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("offPeakWindow")]
        public Input<Inputs.DomainOffPeakWindowOptionsOffPeakWindowArgs>? OffPeakWindow { get; set; }

        public DomainOffPeakWindowOptionsArgs()
        {
        }
        public static new DomainOffPeakWindowOptionsArgs Empty => new DomainOffPeakWindowOptionsArgs();
    }
}
