// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GuardDuty.Inputs
{

    public sealed class OrganizationConfigurationDatasourcesS3LogsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// *Deprecated:* Use `auto_enable_organization_members` instead. When this setting is enabled, all new accounts that are created in, or added to, the organization are added as a member accounts of the organization’s GuardDuty delegated administrator and GuardDuty is enabled in that AWS Region.
        /// </summary>
        [Input("autoEnable", required: true)]
        public Input<bool> AutoEnable { get; set; } = null!;

        public OrganizationConfigurationDatasourcesS3LogsGetArgs()
        {
        }
        public static new OrganizationConfigurationDatasourcesS3LogsGetArgs Empty => new OrganizationConfigurationDatasourcesS3LogsGetArgs();
    }
}
