// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Outputs
{

    [OutputType]
    public sealed class WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseBlock
    {
        /// <summary>
        /// Defines a custom response for the web request. See `custom_response` below for details.
        /// </summary>
        public readonly Outputs.WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseBlockCustomResponse? CustomResponse;

        [OutputConstructor]
        private WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseBlock(Outputs.WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseBlockCustomResponse? customResponse)
        {
            CustomResponse = customResponse;
        }
    }
}
