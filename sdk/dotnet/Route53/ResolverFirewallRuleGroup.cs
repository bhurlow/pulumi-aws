// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Route53
{
    /// <summary>
    /// Provides a Route 53 Resolver DNS Firewall rule group resource.
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
    ///     var example = new Aws.Route53.ResolverFirewallRuleGroup("example");
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_route53_resolver_firewall_rule_group.example
    /// 
    ///  id = "rslvr-frg-0123456789abcdef" } Using `pulumi import`, import
    /// 
    /// Route 53 Resolver DNS Firewall rule groups using the Route 53 Resolver DNS Firewall rule group ID. For exampleconsole % pulumi import aws_route53_resolver_firewall_rule_group.example rslvr-frg-0123456789abcdef
    /// </summary>
    [AwsResourceType("aws:route53/resolverFirewallRuleGroup:ResolverFirewallRuleGroup")]
    public partial class ResolverFirewallRuleGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN (Amazon Resource Name) of the rule group.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A name that lets you identify the rule group, to manage and use it.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The AWS account ID for the account that created the rule group. When a rule group is shared with your account, this is the account that has shared the rule group with you.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// Whether the rule group is shared with other AWS accounts, or was shared with the current account by another AWS account. Sharing is configured through AWS Resource Access Manager (AWS RAM). Valid values: `NOT_SHARED`, `SHARED_BY_ME`, `SHARED_WITH_ME`
        /// </summary>
        [Output("shareStatus")]
        public Output<string> ShareStatus { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a ResolverFirewallRuleGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResolverFirewallRuleGroup(string name, ResolverFirewallRuleGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:route53/resolverFirewallRuleGroup:ResolverFirewallRuleGroup", name, args ?? new ResolverFirewallRuleGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResolverFirewallRuleGroup(string name, Input<string> id, ResolverFirewallRuleGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:route53/resolverFirewallRuleGroup:ResolverFirewallRuleGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResolverFirewallRuleGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResolverFirewallRuleGroup Get(string name, Input<string> id, ResolverFirewallRuleGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ResolverFirewallRuleGroup(name, id, state, options);
        }
    }

    public sealed class ResolverFirewallRuleGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A name that lets you identify the rule group, to manage and use it.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ResolverFirewallRuleGroupArgs()
        {
        }
        public static new ResolverFirewallRuleGroupArgs Empty => new ResolverFirewallRuleGroupArgs();
    }

    public sealed class ResolverFirewallRuleGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN (Amazon Resource Name) of the rule group.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// A name that lets you identify the rule group, to manage and use it.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The AWS account ID for the account that created the rule group. When a rule group is shared with your account, this is the account that has shared the rule group with you.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        /// <summary>
        /// Whether the rule group is shared with other AWS accounts, or was shared with the current account by another AWS account. Sharing is configured through AWS Resource Access Manager (AWS RAM). Valid values: `NOT_SHARED`, `SHARED_BY_ME`, `SHARED_WITH_ME`
        /// </summary>
        [Input("shareStatus")]
        public Input<string>? ShareStatus { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public ResolverFirewallRuleGroupState()
        {
        }
        public static new ResolverFirewallRuleGroupState Empty => new ResolverFirewallRuleGroupState();
    }
}
