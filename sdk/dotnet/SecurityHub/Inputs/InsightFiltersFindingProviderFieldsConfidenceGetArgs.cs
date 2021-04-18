// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SecurityHub.Inputs
{

    public sealed class InsightFiltersFindingProviderFieldsConfidenceGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The equal-to condition to be applied to a single field when querying for findings, provided as a String.
        /// </summary>
        [Input("eq")]
        public Input<string>? Eq { get; set; }

        /// <summary>
        /// The greater-than-equal condition to be applied to a single field when querying for findings, provided as a String.
        /// </summary>
        [Input("gte")]
        public Input<string>? Gte { get; set; }

        /// <summary>
        /// The less-than-equal condition to be applied to a single field when querying for findings, provided as a String.
        /// </summary>
        [Input("lte")]
        public Input<string>? Lte { get; set; }

        public InsightFiltersFindingProviderFieldsConfidenceGetArgs()
        {
        }
    }
}
