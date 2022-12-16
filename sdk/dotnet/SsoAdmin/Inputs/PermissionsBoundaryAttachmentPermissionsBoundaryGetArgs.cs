// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SsoAdmin.Inputs
{

    public sealed class PermissionsBoundaryAttachmentPermissionsBoundaryGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name and path of a customer managed policy. See below.
        /// </summary>
        [Input("customerManagedPolicyReference")]
        public Input<Inputs.PermissionsBoundaryAttachmentPermissionsBoundaryCustomerManagedPolicyReferenceGetArgs>? CustomerManagedPolicyReference { get; set; }

        /// <summary>
        /// AWS-managed IAM policy ARN to use as the permissions boundary.
        /// </summary>
        [Input("managedPolicyArn")]
        public Input<string>? ManagedPolicyArn { get; set; }

        public PermissionsBoundaryAttachmentPermissionsBoundaryGetArgs()
        {
        }
        public static new PermissionsBoundaryAttachmentPermissionsBoundaryGetArgs Empty => new PermissionsBoundaryAttachmentPermissionsBoundaryGetArgs();
    }
}
