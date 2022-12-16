// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Inputs
{

    public sealed class UserProfileUserSettingsJupyterServerAppSettingsCodeRepositoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL of the Git repository.
        /// </summary>
        [Input("repositoryUrl", required: true)]
        public Input<string> RepositoryUrl { get; set; } = null!;

        public UserProfileUserSettingsJupyterServerAppSettingsCodeRepositoryArgs()
        {
        }
        public static new UserProfileUserSettingsJupyterServerAppSettingsCodeRepositoryArgs Empty => new UserProfileUserSettingsJupyterServerAppSettingsCodeRepositoryArgs();
    }
}
