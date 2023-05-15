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
    public sealed class GetGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementVcpuCountResult
    {
        /// <summary>
        /// Maximum.
        /// </summary>
        public readonly int Max;
        /// <summary>
        /// Minimum.
        /// </summary>
        public readonly int Min;

        [OutputConstructor]
        private GetGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementVcpuCountResult(
            int max,

            int min)
        {
            Max = max;
            Min = min;
        }
    }
}
