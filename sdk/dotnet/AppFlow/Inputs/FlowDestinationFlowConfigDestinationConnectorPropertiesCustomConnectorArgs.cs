// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppFlow.Inputs
{

    public sealed class FlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnectorArgs : Pulumi.ResourceArgs
    {
        [Input("customProperties")]
        private InputMap<string>? _customProperties;

        /// <summary>
        /// The custom properties that are specific to the connector when it's used as a source in the flow. Maximum of 50 items.
        /// </summary>
        public InputMap<string> CustomProperties
        {
            get => _customProperties ?? (_customProperties = new InputMap<string>());
            set => _customProperties = value;
        }

        /// <summary>
        /// The entity specified in the custom connector as a source in the flow.
        /// </summary>
        [Input("entityName", required: true)]
        public Input<string> EntityName { get; set; } = null!;

        /// <summary>
        /// The settings that determine how Amazon AppFlow handles an error when placing data in the destination. See Error Handling Config for more details.
        /// </summary>
        [Input("errorHandlingConfig")]
        public Input<Inputs.FlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnectorErrorHandlingConfigArgs>? ErrorHandlingConfig { get; set; }

        [Input("idFieldNames")]
        private InputList<string>? _idFieldNames;

        /// <summary>
        /// The name of the field that Amazon AppFlow uses as an ID when performing a write operation such as update or delete.
        /// </summary>
        public InputList<string> IdFieldNames
        {
            get => _idFieldNames ?? (_idFieldNames = new InputList<string>());
            set => _idFieldNames = value;
        }

        /// <summary>
        /// This specifies the type of write operation to be performed in Zendesk. When the value is `UPSERT`, then `id_field_names` is required. Valid values are `INSERT`, `UPSERT`, `UPDATE`, and `DELETE`.
        /// </summary>
        [Input("writeOperationType")]
        public Input<string>? WriteOperationType { get; set; }

        public FlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnectorArgs()
        {
        }
    }
}
