// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCountCustomRequestHandlingInsertHeaderGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the custom header. For custom request header insertion, when AWS WAF inserts the header into the request, it prefixes this name `x-amzn-waf-`, to avoid confusion with the headers that are already in the request. For example, for the header name `sample`, AWS WAF inserts the header `x-amzn-waf-sample`.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Value of the custom header.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCountCustomRequestHandlingInsertHeaderGetArgs()
        {
        }
        public static new WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCountCustomRequestHandlingInsertHeaderGetArgs Empty => new WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCountCustomRequestHandlingInsertHeaderGetArgs();
    }
}
