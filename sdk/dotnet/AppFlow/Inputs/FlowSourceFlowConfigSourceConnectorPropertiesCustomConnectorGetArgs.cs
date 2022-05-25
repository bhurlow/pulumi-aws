// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppFlow.Inputs
{

    public sealed class FlowSourceFlowConfigSourceConnectorPropertiesCustomConnectorGetArgs : Pulumi.ResourceArgs
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

        public FlowSourceFlowConfigSourceConnectorPropertiesCustomConnectorGetArgs()
        {
        }
    }
}
