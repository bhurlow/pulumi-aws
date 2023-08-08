// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The part of a web request that you want AWS WAF to inspect. See `field_to_match` below for details.
        /// </summary>
        [Input("fieldToMatch")]
        public Input<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchArgs>? FieldToMatch { get; set; }

        /// <summary>
        /// String representing the regular expression. Minimum of `1` and maximum of `512` characters.
        /// </summary>
        [Input("regexString", required: true)]
        public Input<string> RegexString { get; set; } = null!;

        [Input("textTransformations", required: true)]
        private InputList<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementTextTransformationArgs>? _textTransformations;

        /// <summary>
        /// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection. At least one transformation is required. See `text_transformation` below for details.
        /// </summary>
        public InputList<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementTextTransformationArgs> TextTransformations
        {
            get => _textTransformations ?? (_textTransformations = new InputList<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementTextTransformationArgs>());
            set => _textTransformations = value;
        }

        public WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementArgs()
        {
        }
        public static new WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementArgs Empty => new WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementArgs();
    }
}
