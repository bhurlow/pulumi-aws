// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Transfer.Outputs
{

    [OutputType]
    public sealed class WorkflowStep
    {
        /// <summary>
        /// Details for a step that performs a file copy. See Copy Step Details below.
        /// </summary>
        public readonly Outputs.WorkflowStepCopyStepDetails? CopyStepDetails;
        /// <summary>
        /// Details for a step that invokes a lambda function.
        /// </summary>
        public readonly Outputs.WorkflowStepCustomStepDetails? CustomStepDetails;
        /// <summary>
        /// Details for a step that decrypts the file.
        /// </summary>
        public readonly Outputs.WorkflowStepDecryptStepDetails? DecryptStepDetails;
        /// <summary>
        /// Details for a step that deletes the file.
        /// </summary>
        public readonly Outputs.WorkflowStepDeleteStepDetails? DeleteStepDetails;
        /// <summary>
        /// Details for a step that creates one or more tags.
        /// </summary>
        public readonly Outputs.WorkflowStepTagStepDetails? TagStepDetails;
        /// <summary>
        /// One of the following step types are supported. `COPY`, `CUSTOM`, `DECRYPT`, `DELETE`, and `TAG`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private WorkflowStep(
            Outputs.WorkflowStepCopyStepDetails? copyStepDetails,

            Outputs.WorkflowStepCustomStepDetails? customStepDetails,

            Outputs.WorkflowStepDecryptStepDetails? decryptStepDetails,

            Outputs.WorkflowStepDeleteStepDetails? deleteStepDetails,

            Outputs.WorkflowStepTagStepDetails? tagStepDetails,

            string type)
        {
            CopyStepDetails = copyStepDetails;
            CustomStepDetails = customStepDetails;
            DecryptStepDetails = decryptStepDetails;
            DeleteStepDetails = deleteStepDetails;
            TagStepDetails = tagStepDetails;
            Type = type;
        }
    }
}
