// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CostExplorer.Inputs
{

    public sealed class CostCategorySplitChargeRuleParameterGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Parameter type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// Parameter values.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public CostCategorySplitChargeRuleParameterGetArgs()
        {
        }
    }
}
