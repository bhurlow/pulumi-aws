// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kendra.Inputs
{

    public sealed class DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration of the condition used for the target document attribute or metadata field when ingesting documents into Amazon Kendra. See condition.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationConditionArgs>? Condition { get; set; }

        /// <summary>
        /// `TRUE` to delete content if the condition used for the target attribute is met.
        /// </summary>
        [Input("documentContentDeletion")]
        public Input<bool>? DocumentContentDeletion { get; set; }

        /// <summary>
        /// Configuration of the target document attribute or metadata field when ingesting documents into Amazon Kendra. You can also include a value. Detailed below.
        /// </summary>
        [Input("target")]
        public Input<Inputs.DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationTargetArgs>? Target { get; set; }

        public DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationArgs()
        {
        }
        public static new DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationArgs Empty => new DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationArgs();
    }
}
