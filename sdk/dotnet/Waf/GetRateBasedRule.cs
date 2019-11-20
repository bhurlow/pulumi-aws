// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Waf
{
    public static partial class Invokes
    {
        /// <summary>
        /// `aws.waf.RateBasedRule` Retrieves a WAF Rate Based Rule Resource Id.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/waf_rate_based_rule.html.markdown.
        /// </summary>
        public static Task<GetRateBasedRuleResult> GetRateBasedRule(GetRateBasedRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRateBasedRuleResult>("aws:waf/getRateBasedRule:getRateBasedRule", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetRateBasedRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the WAF rate based rule.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetRateBasedRuleArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetRateBasedRuleResult
    {
        public readonly string Name;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetRateBasedRuleResult(
            string name,
            string id)
        {
            Name = name;
            Id = id;
        }
    }
}
