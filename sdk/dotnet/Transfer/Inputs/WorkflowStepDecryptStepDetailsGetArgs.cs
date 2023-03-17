// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Transfer.Inputs
{

    public sealed class WorkflowStepDecryptStepDetailsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the location for the file being copied. Use ${Transfer:username} in this field to parametrize the destination prefix by username.
        /// </summary>
        [Input("destinationFileLocation")]
        public Input<Inputs.WorkflowStepDecryptStepDetailsDestinationFileLocationGetArgs>? DestinationFileLocation { get; set; }

        /// <summary>
        /// The name of the step, used as an identifier.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A flag that indicates whether or not to overwrite an existing file of the same name. The default is `FALSE`. Valid values are `TRUE` and `FALSE`.
        /// </summary>
        [Input("overwriteExisting")]
        public Input<string>? OverwriteExisting { get; set; }

        /// <summary>
        /// Specifies which file to use as input to the workflow step: either the output from the previous step, or the originally uploaded file for the workflow. Enter ${previous.file} to use the previous file as the input. In this case, this workflow step uses the output file from the previous workflow step as input. This is the default value. Enter ${original.file} to use the originally-uploaded file location as input for this step.
        /// </summary>
        [Input("sourceFileLocation")]
        public Input<string>? SourceFileLocation { get; set; }

        /// <summary>
        /// The type of encryption used. Currently, this value must be `"PGP"`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public WorkflowStepDecryptStepDetailsGetArgs()
        {
        }
        public static new WorkflowStepDecryptStepDetailsGetArgs Empty => new WorkflowStepDecryptStepDetailsGetArgs();
    }
}
