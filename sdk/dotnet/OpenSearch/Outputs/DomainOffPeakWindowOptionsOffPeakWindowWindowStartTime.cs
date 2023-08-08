// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.OpenSearch.Outputs
{

    [OutputType]
    public sealed class DomainOffPeakWindowOptionsOffPeakWindowWindowStartTime
    {
        /// <summary>
        /// Starting hour of the 10-hour window for updates
        /// </summary>
        public readonly int? Hours;
        /// <summary>
        /// Starting minute of the 10-hour window for updates
        /// </summary>
        public readonly int? Minutes;

        [OutputConstructor]
        private DomainOffPeakWindowOptionsOffPeakWindowWindowStartTime(
            int? hours,

            int? minutes)
        {
            Hours = hours;
            Minutes = minutes;
        }
    }
}
