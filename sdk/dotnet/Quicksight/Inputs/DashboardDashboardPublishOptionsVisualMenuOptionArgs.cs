// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Quicksight.Inputs
{

    public sealed class DashboardDashboardPublishOptionsVisualMenuOptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Availability status. Possibles values: ENABLED, DISABLED.
        /// </summary>
        [Input("availabilityStatus")]
        public Input<string>? AvailabilityStatus { get; set; }

        public DashboardDashboardPublishOptionsVisualMenuOptionArgs()
        {
        }
        public static new DashboardDashboardPublishOptionsVisualMenuOptionArgs Empty => new DashboardDashboardPublishOptionsVisualMenuOptionArgs();
    }
}
