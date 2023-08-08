// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies that AWS WAF should allow requests by default. See `allow` below for details.
        /// </summary>
        [Input("allow")]
        public Input<Inputs.WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseAllowGetArgs>? Allow { get; set; }

        /// <summary>
        /// Specifies that AWS WAF should block requests by default. See `block` below for details.
        /// </summary>
        [Input("block")]
        public Input<Inputs.WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseBlockGetArgs>? Block { get; set; }

        /// <summary>
        /// Instructs AWS WAF to run a Captcha check against the web request. See `captcha` below for details.
        /// </summary>
        [Input("captcha")]
        public Input<Inputs.WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCaptchaGetArgs>? Captcha { get; set; }

        /// <summary>
        /// Instructs AWS WAF to count the web request and allow it. See `count` below for details.
        /// </summary>
        [Input("count")]
        public Input<Inputs.WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCountGetArgs>? Count { get; set; }

        public WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseGetArgs()
        {
        }
        public static new WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseGetArgs Empty => new WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseGetArgs();
    }
}
