// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch.Inputs
{

    public sealed class EventTargetInputTransformerGetArgs : Pulumi.ResourceArgs
    {
        [Input("inputPaths")]
        private InputMap<string>? _inputPaths;

        /// <summary>
        /// Key value pairs specified in the form of JSONPath (for example, time = $.time)
        /// * You can have as many as 10 key-value pairs.
        /// * You must use JSON dot notation, not bracket notation.
        /// * The keys can't start with "AWS".
        /// </summary>
        public InputMap<string> InputPaths
        {
            get => _inputPaths ?? (_inputPaths = new InputMap<string>());
            set => _inputPaths = value;
        }

        /// <summary>
        /// Template to customize data sent to the target. Must be valid JSON. To send a string value, the string value must include double quotes. Values must be escaped for both JSON and the provider, e.g. `"\"Your string goes here.\\nA new line.\""`
        /// </summary>
        [Input("inputTemplate", required: true)]
        public Input<string> InputTemplate { get; set; } = null!;

        public EventTargetInputTransformerGetArgs()
        {
        }
    }
}
