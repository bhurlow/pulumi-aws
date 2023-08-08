// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("statements", required: true)]
        private InputList<Inputs.WebAclRuleStatementGetArgs>? _statements;

        /// <summary>
        /// The statements to combine.
        /// </summary>
        public InputList<Inputs.WebAclRuleStatementGetArgs> Statements
        {
            get => _statements ?? (_statements = new InputList<Inputs.WebAclRuleStatementGetArgs>());
            set => _statements = value;
        }

        public WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementGetArgs()
        {
        }
        public static new WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementGetArgs Empty => new WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementGetArgs();
    }
}
