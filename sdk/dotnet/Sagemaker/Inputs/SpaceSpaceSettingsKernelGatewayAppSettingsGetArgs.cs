// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Inputs
{

    public sealed class SpaceSpaceSettingsKernelGatewayAppSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("customImages")]
        private InputList<Inputs.SpaceSpaceSettingsKernelGatewayAppSettingsCustomImageGetArgs>? _customImages;

        /// <summary>
        /// A list of custom SageMaker images that are configured to run as a KernelGateway app. see Custom Image below.
        /// </summary>
        public InputList<Inputs.SpaceSpaceSettingsKernelGatewayAppSettingsCustomImageGetArgs> CustomImages
        {
            get => _customImages ?? (_customImages = new InputList<Inputs.SpaceSpaceSettingsKernelGatewayAppSettingsCustomImageGetArgs>());
            set => _customImages = value;
        }

        /// <summary>
        /// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
        /// </summary>
        [Input("defaultResourceSpec", required: true)]
        public Input<Inputs.SpaceSpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecGetArgs> DefaultResourceSpec { get; set; } = null!;

        [Input("lifecycleConfigArns")]
        private InputList<string>? _lifecycleConfigArns;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Lifecycle Configurations.
        /// </summary>
        public InputList<string> LifecycleConfigArns
        {
            get => _lifecycleConfigArns ?? (_lifecycleConfigArns = new InputList<string>());
            set => _lifecycleConfigArns = value;
        }

        public SpaceSpaceSettingsKernelGatewayAppSettingsGetArgs()
        {
        }
        public static new SpaceSpaceSettingsKernelGatewayAppSettingsGetArgs Empty => new SpaceSpaceSettingsKernelGatewayAppSettingsGetArgs();
    }
}
