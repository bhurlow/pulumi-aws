// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Outputs
{

    [OutputType]
    public sealed class SpaceSpaceSettings
    {
        /// <summary>
        /// The type of app created within the space.
        /// </summary>
        public readonly string? AppType;
        /// <summary>
        /// The Code Editor application settings. See Code Editor App Settings below.
        /// </summary>
        public readonly Outputs.SpaceSpaceSettingsCodeEditorAppSettings? CodeEditorAppSettings;
        /// <summary>
        /// A file system, created by you, that you assign to a space for an Amazon SageMaker Domain. See Custom File System below.
        /// </summary>
        public readonly ImmutableArray<Outputs.SpaceSpaceSettingsCustomFileSystem> CustomFileSystems;
        /// <summary>
        /// The settings for the JupyterLab application. See Jupyter Lab App Settings below.
        /// </summary>
        public readonly Outputs.SpaceSpaceSettingsJupyterLabAppSettings? JupyterLabAppSettings;
        /// <summary>
        /// The Jupyter server's app settings. See Jupyter Server App Settings below.
        /// </summary>
        public readonly Outputs.SpaceSpaceSettingsJupyterServerAppSettings? JupyterServerAppSettings;
        /// <summary>
        /// The kernel gateway app settings. See Kernel Gateway App Settings below.
        /// </summary>
        public readonly Outputs.SpaceSpaceSettingsKernelGatewayAppSettings? KernelGatewayAppSettings;
        public readonly Outputs.SpaceSpaceSettingsSpaceStorageSettings? SpaceStorageSettings;

        [OutputConstructor]
        private SpaceSpaceSettings(
            string? appType,

            Outputs.SpaceSpaceSettingsCodeEditorAppSettings? codeEditorAppSettings,

            ImmutableArray<Outputs.SpaceSpaceSettingsCustomFileSystem> customFileSystems,

            Outputs.SpaceSpaceSettingsJupyterLabAppSettings? jupyterLabAppSettings,

            Outputs.SpaceSpaceSettingsJupyterServerAppSettings? jupyterServerAppSettings,

            Outputs.SpaceSpaceSettingsKernelGatewayAppSettings? kernelGatewayAppSettings,

            Outputs.SpaceSpaceSettingsSpaceStorageSettings? spaceStorageSettings)
        {
            AppType = appType;
            CodeEditorAppSettings = codeEditorAppSettings;
            CustomFileSystems = customFileSystems;
            JupyterLabAppSettings = jupyterLabAppSettings;
            JupyterServerAppSettings = jupyterServerAppSettings;
            KernelGatewayAppSettings = kernelGatewayAppSettings;
            SpaceStorageSettings = spaceStorageSettings;
        }
    }
}
