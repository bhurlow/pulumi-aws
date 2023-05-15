// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScaling.Outputs
{

    [OutputType]
    public sealed class GetGroupMixedInstancesPolicyResult
    {
        /// <summary>
        /// List of instances distribution objects.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupMixedInstancesPolicyInstancesDistributionResult> InstancesDistributions;
        /// <summary>
        /// List of launch templates along with the overrides.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupMixedInstancesPolicyLaunchTemplateResult> LaunchTemplates;

        [OutputConstructor]
        private GetGroupMixedInstancesPolicyResult(
            ImmutableArray<Outputs.GetGroupMixedInstancesPolicyInstancesDistributionResult> instancesDistributions,

            ImmutableArray<Outputs.GetGroupMixedInstancesPolicyLaunchTemplateResult> launchTemplates)
        {
            InstancesDistributions = instancesDistributions;
            LaunchTemplates = launchTemplates;
        }
    }
}
