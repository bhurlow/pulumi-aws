// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SesV2.Outputs
{

    [OutputType]
    public sealed class ConfigurationSetVdmOptions
    {
        /// <summary>
        /// Specifies additional settings for your VDM configuration as applicable to the Dashboard.
        /// </summary>
        public readonly Outputs.ConfigurationSetVdmOptionsDashboardOptions? DashboardOptions;
        /// <summary>
        /// Specifies additional settings for your VDM configuration as applicable to the Guardian.
        /// </summary>
        public readonly Outputs.ConfigurationSetVdmOptionsGuardianOptions? GuardianOptions;

        [OutputConstructor]
        private ConfigurationSetVdmOptions(
            Outputs.ConfigurationSetVdmOptionsDashboardOptions? dashboardOptions,

            Outputs.ConfigurationSetVdmOptionsGuardianOptions? guardianOptions)
        {
            DashboardOptions = dashboardOptions;
            GuardianOptions = guardianOptions;
        }
    }
}
