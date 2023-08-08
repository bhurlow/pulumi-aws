// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Inputs
{

    public sealed class UserProfileUserSettingsCanvasAppSettingsWorkspaceSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon S3 bucket used to store artifacts generated by Canvas. Updating the Amazon S3 location impacts existing configuration settings, and Canvas users no longer have access to their artifacts. Canvas users must log out and log back in to apply the new location.
        /// </summary>
        [Input("s3ArtifactPath")]
        public Input<string>? S3ArtifactPath { get; set; }

        /// <summary>
        /// The Amazon Web Services Key Management Service (KMS) encryption key ID that is used to encrypt artifacts generated by Canvas in the Amazon S3 bucket.
        /// </summary>
        [Input("s3KmsKeyId")]
        public Input<string>? S3KmsKeyId { get; set; }

        public UserProfileUserSettingsCanvasAppSettingsWorkspaceSettingsArgs()
        {
        }
        public static new UserProfileUserSettingsCanvasAppSettingsWorkspaceSettingsArgs Empty => new UserProfileUserSettingsCanvasAppSettingsWorkspaceSettingsArgs();
    }
}
