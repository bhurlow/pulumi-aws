// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspectionHeaderArgs : global::Pulumi.ResourceArgs
    {
        [Input("failureValues", required: true)]
        private InputList<string>? _failureValues;

        /// <summary>
        /// Values in the response header with the specified name that indicate a failed login attempt.
        /// </summary>
        public InputList<string> FailureValues
        {
            get => _failureValues ?? (_failureValues = new InputList<string>());
            set => _failureValues = value;
        }

        /// <summary>
        /// The name of the header to use.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("successValues", required: true)]
        private InputList<string>? _successValues;

        /// <summary>
        /// Values in the response header with the specified name that indicate a successful login attempt.
        /// </summary>
        public InputList<string> SuccessValues
        {
            get => _successValues ?? (_successValues = new InputList<string>());
            set => _successValues = value;
        }

        public WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspectionHeaderArgs()
        {
        }
        public static new WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspectionHeaderArgs Empty => new WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspectionHeaderArgs();
    }
}
