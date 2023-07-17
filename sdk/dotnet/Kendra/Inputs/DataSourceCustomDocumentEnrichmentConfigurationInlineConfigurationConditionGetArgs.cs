// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kendra.Inputs
{

    public sealed class DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationConditionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identifier of the document attribute used for the condition. For example, `_source_uri` could be an identifier for the attribute or metadata field that contains source URIs associated with the documents. Amazon Kendra currently does not support `_document_body` as an attribute key used for the condition.
        /// </summary>
        [Input("conditionDocumentAttributeKey", required: true)]
        public Input<string> ConditionDocumentAttributeKey { get; set; } = null!;

        /// <summary>
        /// The value used by the operator. For example, you can specify the value 'financial' for strings in the `_source_uri` field that partially match or contain this value. See condition_on_value.
        /// </summary>
        [Input("conditionOnValue")]
        public Input<Inputs.DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationConditionConditionOnValueGetArgs>? ConditionOnValue { get; set; }

        /// <summary>
        /// The condition operator. For example, you can use `Contains` to partially match a string. Valid Values: `GreaterThan` | `GreaterThanOrEquals` | `LessThan` | `LessThanOrEquals` | `Equals` | `NotEquals` | `Contains` | `NotContains` | `Exists` | `NotExists` | `BeginsWith`.
        /// </summary>
        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        public DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationConditionGetArgs()
        {
        }
        public static new DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationConditionGetArgs Empty => new DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationConditionGetArgs();
    }
}
