// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ses
{
    /// <summary>
    /// Provides a resource to designate the active SES receipt rule set
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_active_receipt_rule_set.html.markdown.
    /// </summary>
    public partial class ActiveReceiptRuleSet : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the rule set
        /// </summary>
        [Output("ruleSetName")]
        public Output<string> RuleSetName { get; private set; } = null!;


        /// <summary>
        /// Create a ActiveReceiptRuleSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ActiveReceiptRuleSet(string name, ActiveReceiptRuleSetArgs args, CustomResourceOptions? options = null)
            : base("aws:ses/activeReceiptRuleSet:ActiveReceiptRuleSet", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ActiveReceiptRuleSet(string name, Input<string> id, ActiveReceiptRuleSetState? state = null, CustomResourceOptions? options = null)
            : base("aws:ses/activeReceiptRuleSet:ActiveReceiptRuleSet", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ActiveReceiptRuleSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ActiveReceiptRuleSet Get(string name, Input<string> id, ActiveReceiptRuleSetState? state = null, CustomResourceOptions? options = null)
        {
            return new ActiveReceiptRuleSet(name, id, state, options);
        }
    }

    public sealed class ActiveReceiptRuleSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the rule set
        /// </summary>
        [Input("ruleSetName", required: true)]
        public Input<string> RuleSetName { get; set; } = null!;

        public ActiveReceiptRuleSetArgs()
        {
        }
    }

    public sealed class ActiveReceiptRuleSetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the rule set
        /// </summary>
        [Input("ruleSetName")]
        public Input<string>? RuleSetName { get; set; }

        public ActiveReceiptRuleSetState()
        {
        }
    }
}
