// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ses
{
    /// <summary>
    /// Provides an SES receipt rule set resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Aws.Ses.ReceiptRuleSet("main", new()
    ///     {
    ///         RuleSetName = "primary-rules",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_ses_receipt_rule_set.my_rule_set
    /// 
    ///  id = "my_rule_set_name" } Using `pulumi import`, import SES receipt rule sets using the rule set name. For exampleconsole % pulumi import aws_ses_receipt_rule_set.my_rule_set my_rule_set_name
    /// </summary>
    [AwsResourceType("aws:ses/receiptRuleSet:ReceiptRuleSet")]
    public partial class ReceiptRuleSet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// SES receipt rule set ARN.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Name of the rule set.
        /// </summary>
        [Output("ruleSetName")]
        public Output<string> RuleSetName { get; private set; } = null!;


        /// <summary>
        /// Create a ReceiptRuleSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReceiptRuleSet(string name, ReceiptRuleSetArgs args, CustomResourceOptions? options = null)
            : base("aws:ses/receiptRuleSet:ReceiptRuleSet", name, args ?? new ReceiptRuleSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReceiptRuleSet(string name, Input<string> id, ReceiptRuleSetState? state = null, CustomResourceOptions? options = null)
            : base("aws:ses/receiptRuleSet:ReceiptRuleSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReceiptRuleSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReceiptRuleSet Get(string name, Input<string> id, ReceiptRuleSetState? state = null, CustomResourceOptions? options = null)
        {
            return new ReceiptRuleSet(name, id, state, options);
        }
    }

    public sealed class ReceiptRuleSetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the rule set.
        /// </summary>
        [Input("ruleSetName", required: true)]
        public Input<string> RuleSetName { get; set; } = null!;

        public ReceiptRuleSetArgs()
        {
        }
        public static new ReceiptRuleSetArgs Empty => new ReceiptRuleSetArgs();
    }

    public sealed class ReceiptRuleSetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SES receipt rule set ARN.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Name of the rule set.
        /// </summary>
        [Input("ruleSetName")]
        public Input<string>? RuleSetName { get; set; }

        public ReceiptRuleSetState()
        {
        }
        public static new ReceiptRuleSetState Empty => new ReceiptRuleSetState();
    }
}
