// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Inputs
{

    public sealed class SpaceSpaceSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of app created within the space.
        /// </summary>
        [Input("appType")]
        public Input<string>? AppType { get; set; }

        /// <summary>
        /// The Code Editor application settings. See Code Editor App Settings below.
        /// </summary>
        [Input("codeEditorAppSettings")]
        public Input<Inputs.SpaceSpaceSettingsCodeEditorAppSettingsGetArgs>? CodeEditorAppSettings { get; set; }

        [Input("customFileSystems")]
        private InputList<Inputs.SpaceSpaceSettingsCustomFileSystemGetArgs>? _customFileSystems;

        /// <summary>
        /// A file system, created by you, that you assign to a space for an Amazon SageMaker Domain. See Custom File System below.
        /// </summary>
        public InputList<Inputs.SpaceSpaceSettingsCustomFileSystemGetArgs> CustomFileSystems
        {
            get => _customFileSystems ?? (_customFileSystems = new InputList<Inputs.SpaceSpaceSettingsCustomFileSystemGetArgs>());
            set => _customFileSystems = value;
        }

        /// <summary>
        /// The settings for the JupyterLab application. See Jupyter Lab App Settings below.
        /// </summary>
        [Input("jupyterLabAppSettings")]
        public Input<Inputs.SpaceSpaceSettingsJupyterLabAppSettingsGetArgs>? JupyterLabAppSettings { get; set; }

        /// <summary>
        /// The Jupyter server's app settings. See Jupyter Server App Settings below.
        /// </summary>
        [Input("jupyterServerAppSettings")]
        public Input<Inputs.SpaceSpaceSettingsJupyterServerAppSettingsGetArgs>? JupyterServerAppSettings { get; set; }

        /// <summary>
        /// The kernel gateway app settings. See Kernel Gateway App Settings below.
        /// </summary>
        [Input("kernelGatewayAppSettings")]
        public Input<Inputs.SpaceSpaceSettingsKernelGatewayAppSettingsGetArgs>? KernelGatewayAppSettings { get; set; }

        [Input("spaceStorageSettings")]
        public Input<Inputs.SpaceSpaceSettingsSpaceStorageSettingsGetArgs>? SpaceStorageSettings { get; set; }

        public SpaceSpaceSettingsGetArgs()
        {
        }
        public static new SpaceSpaceSettingsGetArgs Empty => new SpaceSpaceSettingsGetArgs();
    }
}
