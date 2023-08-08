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
    public sealed class WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverride
    {
        /// <summary>
        /// Override action to use, in place of the configured action of the rule in the rule group. See `action` for details.
        /// </summary>
        public readonly Outputs.WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUse ActionToUse;
        /// <summary>
        /// Name of the rule to override. See the [documentation](https://docs.aws.amazon.com/waf/latest/developerguide/aws-managed-rule-groups-list.html) for a list of names in the appropriate rule group in use.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverride(
            Outputs.WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUse actionToUse,

            string name)
        {
            ActionToUse = actionToUse;
            Name = name;
        }
    }
}
