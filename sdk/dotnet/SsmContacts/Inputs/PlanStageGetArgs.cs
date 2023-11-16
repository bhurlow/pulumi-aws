// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SsmContacts.Inputs
{

    public sealed class PlanStageGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time to wait until beginning the next stage. The duration can only be set to 0 if a target is specified.
        /// </summary>
        [Input("durationInMinutes", required: true)]
        public Input<int> DurationInMinutes { get; set; } = null!;

        [Input("targets")]
        private InputList<Inputs.PlanStageTargetGetArgs>? _targets;

        /// <summary>
        /// One or more configuration blocks for specifying the contacts or contact methods that the escalation plan or engagement plan is engaging. See Target below for more details.
        /// </summary>
        public InputList<Inputs.PlanStageTargetGetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.PlanStageTargetGetArgs>());
            set => _targets = value;
        }

        public PlanStageGetArgs()
        {
        }
        public static new PlanStageGetArgs Empty => new PlanStageGetArgs();
    }
}
