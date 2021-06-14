// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecs.Outputs
{

    [OutputType]
    public sealed class TaskDefinitionVolumeFsxWindowsFileServerVolumeConfiguration
    {
        /// <summary>
        /// Configuration block for authorization for the Amazon FSx for Windows File Server file system detailed below.
        /// </summary>
        public readonly Outputs.TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfig AuthorizationConfig;
        /// <summary>
        /// The Amazon FSx for Windows File Server file system ID to use.
        /// </summary>
        public readonly string FileSystemId;
        /// <summary>
        /// The directory within the Amazon FSx for Windows File Server file system to mount as the root directory inside the host.
        /// </summary>
        public readonly string RootDirectory;

        [OutputConstructor]
        private TaskDefinitionVolumeFsxWindowsFileServerVolumeConfiguration(
            Outputs.TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfig authorizationConfig,

            string fileSystemId,

            string rootDirectory)
        {
            AuthorizationConfig = authorizationConfig;
            FileSystemId = fileSystemId;
            RootDirectory = rootDirectory;
        }
    }
}
