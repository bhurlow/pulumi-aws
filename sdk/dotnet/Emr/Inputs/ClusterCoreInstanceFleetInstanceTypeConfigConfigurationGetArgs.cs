// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Emr.Inputs
{

    public sealed class ClusterCoreInstanceFleetInstanceTypeConfigConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Classification within a configuration.
        /// </summary>
        [Input("classification")]
        public Input<string>? Classification { get; set; }

        [Input("properties")]
        private InputMap<object>? _properties;

        /// <summary>
        /// Key-Value map of Java properties that are set when the step runs. You can use these properties to pass key value pairs to your main function.
        /// </summary>
        public InputMap<object> Properties
        {
            get => _properties ?? (_properties = new InputMap<object>());
            set => _properties = value;
        }

        public ClusterCoreInstanceFleetInstanceTypeConfigConfigurationGetArgs()
        {
        }
    }
}
