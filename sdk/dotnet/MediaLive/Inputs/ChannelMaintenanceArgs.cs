// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class ChannelMaintenanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The day of the week to use for maintenance.
        /// </summary>
        [Input("maintenanceDay", required: true)]
        public Input<string> MaintenanceDay { get; set; } = null!;

        /// <summary>
        /// The hour maintenance will start.
        /// </summary>
        [Input("maintenanceStartTime", required: true)]
        public Input<string> MaintenanceStartTime { get; set; } = null!;

        public ChannelMaintenanceArgs()
        {
        }
        public static new ChannelMaintenanceArgs Empty => new ChannelMaintenanceArgs();
    }
}
