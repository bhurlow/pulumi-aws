// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppFlow.Outputs
{

    [OutputType]
    public sealed class FlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnector
    {
        /// <summary>
        /// The custom properties that are specific to the connector when it's used as a source in the flow. Maximum of 50 items.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? CustomProperties;
        /// <summary>
        /// The entity specified in the custom connector as a source in the flow.
        /// </summary>
        public readonly string EntityName;
        /// <summary>
        /// The settings that determine how Amazon AppFlow handles an error when placing data in the destination. See Error Handling Config for more details.
        /// </summary>
        public readonly Outputs.FlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnectorErrorHandlingConfig? ErrorHandlingConfig;
        /// <summary>
        /// The name of the field that Amazon AppFlow uses as an ID when performing a write operation such as update or delete.
        /// </summary>
        public readonly ImmutableArray<string> IdFieldNames;
        /// <summary>
        /// This specifies the type of write operation to be performed in Zendesk. When the value is `UPSERT`, then `id_field_names` is required. Valid values are `INSERT`, `UPSERT`, `UPDATE`, and `DELETE`.
        /// </summary>
        public readonly string? WriteOperationType;

        [OutputConstructor]
        private FlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnector(
            ImmutableDictionary<string, string>? customProperties,

            string entityName,

            Outputs.FlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnectorErrorHandlingConfig? errorHandlingConfig,

            ImmutableArray<string> idFieldNames,

            string? writeOperationType)
        {
            CustomProperties = customProperties;
            EntityName = entityName;
            ErrorHandlingConfig = errorHandlingConfig;
            IdFieldNames = idFieldNames;
            WriteOperationType = writeOperationType;
        }
    }
}
