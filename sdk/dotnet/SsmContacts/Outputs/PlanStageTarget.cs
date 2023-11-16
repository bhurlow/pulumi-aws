// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SsmContacts.Outputs
{

    [OutputType]
    public sealed class PlanStageTarget
    {
        /// <summary>
        /// A configuration block for specifying information about the contact channel that Incident Manager engages. See Channel Target Info for more details.
        /// </summary>
        public readonly Outputs.PlanStageTargetChannelTargetInfo? ChannelTargetInfo;
        /// <summary>
        /// A configuration block for specifying information about the contact that Incident Manager engages. See Contact Target Info for more details.
        /// </summary>
        public readonly Outputs.PlanStageTargetContactTargetInfo? ContactTargetInfo;

        [OutputConstructor]
        private PlanStageTarget(
            Outputs.PlanStageTargetChannelTargetInfo? channelTargetInfo,

            Outputs.PlanStageTargetContactTargetInfo? contactTargetInfo)
        {
            ChannelTargetInfo = channelTargetInfo;
            ContactTargetInfo = contactTargetInfo;
        }
    }
}
