// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspectionBodyContainsArgs : global::Pulumi.ResourceArgs
    {
        [Input("failureStrings", required: true)]
        private InputList<string>? _failureStrings;

        /// <summary>
        /// Strings in the body of the response that indicate a failed login attempt.
        /// </summary>
        public InputList<string> FailureStrings
        {
            get => _failureStrings ?? (_failureStrings = new InputList<string>());
            set => _failureStrings = value;
        }

        [Input("successStrings", required: true)]
        private InputList<string>? _successStrings;

        /// <summary>
        /// Strings in the body of the response that indicate a successful login attempt.
        /// </summary>
        public InputList<string> SuccessStrings
        {
            get => _successStrings ?? (_successStrings = new InputList<string>());
            set => _successStrings = value;
        }

        public WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspectionBodyContainsArgs()
        {
        }
        public static new WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspectionBodyContainsArgs Empty => new WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspectionBodyContainsArgs();
    }
}
