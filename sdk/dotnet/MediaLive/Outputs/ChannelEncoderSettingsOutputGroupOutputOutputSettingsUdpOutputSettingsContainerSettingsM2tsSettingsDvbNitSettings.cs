// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive.Outputs
{

    [OutputType]
    public sealed class ChannelEncoderSettingsOutputGroupOutputOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsDvbNitSettings
    {
        public readonly int NetworkId;
        public readonly string NetworkName;
        public readonly int? RepInterval;

        [OutputConstructor]
        private ChannelEncoderSettingsOutputGroupOutputOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsDvbNitSettings(
            int networkId,

            string networkName,

            int? repInterval)
        {
            NetworkId = networkId;
            NetworkName = networkName;
            RepInterval = repInterval;
        }
    }
}
