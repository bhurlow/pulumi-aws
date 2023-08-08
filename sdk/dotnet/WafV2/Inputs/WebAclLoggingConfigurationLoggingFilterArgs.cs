// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclLoggingConfigurationLoggingFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default handling for logs that don't match any of the specified filtering conditions. Valid values for `default_behavior` are `KEEP` or `DROP`.
        /// </summary>
        [Input("defaultBehavior", required: true)]
        public Input<string> DefaultBehavior { get; set; } = null!;

        [Input("filters", required: true)]
        private InputList<Inputs.WebAclLoggingConfigurationLoggingFilterFilterArgs>? _filters;

        /// <summary>
        /// Filter(s) that you want to apply to the logs. See Filter below for more details.
        /// </summary>
        public InputList<Inputs.WebAclLoggingConfigurationLoggingFilterFilterArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.WebAclLoggingConfigurationLoggingFilterFilterArgs>());
            set => _filters = value;
        }

        public WebAclLoggingConfigurationLoggingFilterArgs()
        {
        }
        public static new WebAclLoggingConfigurationLoggingFilterArgs Empty => new WebAclLoggingConfigurationLoggingFilterArgs();
    }
}
