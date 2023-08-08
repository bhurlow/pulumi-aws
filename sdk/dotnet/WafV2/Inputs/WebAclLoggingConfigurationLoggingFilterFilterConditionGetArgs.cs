// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclLoggingConfigurationLoggingFilterFilterConditionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration for a single action condition. See Action Condition below for more details.
        /// </summary>
        [Input("actionCondition")]
        public Input<Inputs.WebAclLoggingConfigurationLoggingFilterFilterConditionActionConditionGetArgs>? ActionCondition { get; set; }

        /// <summary>
        /// Condition for a single label name. See Label Name Condition below for more details.
        /// </summary>
        [Input("labelNameCondition")]
        public Input<Inputs.WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameConditionGetArgs>? LabelNameCondition { get; set; }

        public WebAclLoggingConfigurationLoggingFilterFilterConditionGetArgs()
        {
        }
        public static new WebAclLoggingConfigurationLoggingFilterFilterConditionGetArgs Empty => new WebAclLoggingConfigurationLoggingFilterFilterConditionGetArgs();
    }
}
