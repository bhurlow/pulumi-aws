// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Inputs
{

    public sealed class SpaceSpaceSettingsKernelGatewayAppSettingsCustomImageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the App Image Config.
        /// </summary>
        [Input("appImageConfigName", required: true)]
        public Input<string> AppImageConfigName { get; set; } = null!;

        /// <summary>
        /// The name of the Custom Image.
        /// </summary>
        [Input("imageName", required: true)]
        public Input<string> ImageName { get; set; } = null!;

        /// <summary>
        /// The version number of the Custom Image.
        /// </summary>
        [Input("imageVersionNumber")]
        public Input<int>? ImageVersionNumber { get; set; }

        public SpaceSpaceSettingsKernelGatewayAppSettingsCustomImageArgs()
        {
        }
        public static new SpaceSpaceSettingsKernelGatewayAppSettingsCustomImageArgs Empty => new SpaceSpaceSettingsKernelGatewayAppSettingsCustomImageArgs();
    }
}
