// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SsmContacts.Inputs
{

    public sealed class PlanStageTargetContactTargetInfoGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the contact.
        /// </summary>
        [Input("contactId")]
        public Input<string>? ContactId { get; set; }

        /// <summary>
        /// A Boolean value determining if the contact's acknowledgement stops the progress of stages in the plan.
        /// </summary>
        [Input("isEssential", required: true)]
        public Input<bool> IsEssential { get; set; } = null!;

        public PlanStageTargetContactTargetInfoGetArgs()
        {
        }
        public static new PlanStageTargetContactTargetInfoGetArgs Empty => new PlanStageTargetContactTargetInfoGetArgs();
    }
}
