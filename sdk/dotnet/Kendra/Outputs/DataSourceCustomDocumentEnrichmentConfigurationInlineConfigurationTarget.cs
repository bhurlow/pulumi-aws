// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kendra.Outputs
{

    [OutputType]
    public sealed class DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationTarget
    {
        /// <summary>
        /// The identifier of the target document attribute or metadata field. For example, 'Department' could be an identifier for the target attribute or metadata field that includes the department names associated with the documents.
        /// </summary>
        public readonly string? TargetDocumentAttributeKey;
        /// <summary>
        /// The target value you want to create for the target attribute. For example, 'Finance' could be the target value for the target attribute key 'Department'. See target_document_attribute_value.
        /// </summary>
        public readonly Outputs.DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationTargetTargetDocumentAttributeValue? TargetDocumentAttributeValue;
        /// <summary>
        /// `TRUE` to delete the existing target value for your specified target attribute key. You cannot create a target value and set this to `TRUE`. To create a target value (`TargetDocumentAttributeValue`), set this to `FALSE`.
        /// </summary>
        public readonly bool? TargetDocumentAttributeValueDeletion;

        [OutputConstructor]
        private DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationTarget(
            string? targetDocumentAttributeKey,

            Outputs.DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationTargetTargetDocumentAttributeValue? targetDocumentAttributeValue,

            bool? targetDocumentAttributeValueDeletion)
        {
            TargetDocumentAttributeKey = targetDocumentAttributeKey;
            TargetDocumentAttributeValue = targetDocumentAttributeValue;
            TargetDocumentAttributeValueDeletion = targetDocumentAttributeValueDeletion;
        }
    }
}
