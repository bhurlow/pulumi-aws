// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Inputs
{

    public sealed class ChannelDestinationMediaPackageSettingGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the channel in MediaPackage that is the destination for this output group.
        /// </summary>
        [Input("channelId", required: true)]
        public Input<string> ChannelId { get; set; } = null!;

        public ChannelDestinationMediaPackageSettingGetArgs()
        {
        }
        public static new ChannelDestinationMediaPackageSettingGetArgs Empty => new ChannelDestinationMediaPackageSettingGetArgs();
    }
}
