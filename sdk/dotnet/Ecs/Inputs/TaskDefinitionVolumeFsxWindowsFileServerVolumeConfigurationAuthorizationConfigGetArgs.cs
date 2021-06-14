// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecs.Inputs
{

    public sealed class TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authorization credential option to use. The authorization credential options can be provided using either the Amazon Resource Name (ARN) of an AWS Secrets Manager secret or AWS Systems Manager Parameter Store parameter. The ARNs refer to the stored credentials.
        /// </summary>
        [Input("credentialsParameter", required: true)]
        public Input<string> CredentialsParameter { get; set; } = null!;

        /// <summary>
        /// A fully qualified domain name hosted by an AWS Directory Service Managed Microsoft AD (Active Directory) or self-hosted AD on Amazon EC2.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        public TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfigGetArgs()
        {
        }
    }
}
