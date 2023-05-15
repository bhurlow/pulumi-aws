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
    /// Provides a Route 53 Resolver DNS Firewall rule group association resource.
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
    ///     var exampleResolverFirewallRuleGroup = new Aws.Route53.ResolverFirewallRuleGroup("exampleResolverFirewallRuleGroup");
    /// 
    ///     var exampleResolverFirewallRuleGroupAssociation = new Aws.Route53.ResolverFirewallRuleGroupAssociation("exampleResolverFirewallRuleGroupAssociation", new()
    ///     {
    ///         FirewallRuleGroupId = exampleResolverFirewallRuleGroup.Id,
    ///         Priority = 100,
    ///         VpcId = aws_vpc.Example.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Route 53 Resolver DNS Firewall rule group associations can be imported using the Route 53 Resolver DNS Firewall rule group association ID, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:route53/resolverFirewallRuleGroupAssociation:ResolverFirewallRuleGroupAssociation example rslvr-frgassoc-0123456789abcdef
    /// ```
    /// </summary>
    [AwsResourceType("aws:route53/resolverFirewallRuleGroupAssociation:ResolverFirewallRuleGroupAssociation")]
    public partial class ResolverFirewallRuleGroupAssociation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN (Amazon Resource Name) of the firewall rule group association.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the firewall rule group.
        /// </summary>
        [Output("firewallRuleGroupId")]
        public Output<string> FirewallRuleGroupId { get; private set; } = null!;

        /// <summary>
        /// If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections. Valid values: `ENABLED`, `DISABLED`.
        /// </summary>
        [Output("mutationProtection")]
        public Output<string> MutationProtection { get; private set; } = null!;

        /// <summary>
        /// A name that lets you identify the rule group association, to manage and use it.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The setting that determines the processing order of the rule group among the rule groups that you associate with the specified VPC. DNS Firewall filters VPC traffic starting from the rule group with the lowest numeric priority setting.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the VPC that you want to associate with the rule group.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a ResolverFirewallRuleGroupAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResolverFirewallRuleGroupAssociation(string name, ResolverFirewallRuleGroupAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws:route53/resolverFirewallRuleGroupAssociation:ResolverFirewallRuleGroupAssociation", name, args ?? new ResolverFirewallRuleGroupAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResolverFirewallRuleGroupAssociation(string name, Input<string> id, ResolverFirewallRuleGroupAssociationState? state = null, CustomResourceOptions? options = null)
            : base("aws:route53/resolverFirewallRuleGroupAssociation:ResolverFirewallRuleGroupAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResolverFirewallRuleGroupAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResolverFirewallRuleGroupAssociation Get(string name, Input<string> id, ResolverFirewallRuleGroupAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new ResolverFirewallRuleGroupAssociation(name, id, state, options);
        }
    }

    public sealed class ResolverFirewallRuleGroupAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique identifier of the firewall rule group.
        /// </summary>
        [Input("firewallRuleGroupId", required: true)]
        public Input<string> FirewallRuleGroupId { get; set; } = null!;

        /// <summary>
        /// If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections. Valid values: `ENABLED`, `DISABLED`.
        /// </summary>
        [Input("mutationProtection")]
        public Input<string>? MutationProtection { get; set; }

        /// <summary>
        /// A name that lets you identify the rule group association, to manage and use it.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The setting that determines the processing order of the rule group among the rule groups that you associate with the specified VPC. DNS Firewall filters VPC traffic starting from the rule group with the lowest numeric priority setting.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The unique identifier of the VPC that you want to associate with the rule group.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public ResolverFirewallRuleGroupAssociationArgs()
        {
        }
        public static new ResolverFirewallRuleGroupAssociationArgs Empty => new ResolverFirewallRuleGroupAssociationArgs();
    }

    public sealed class ResolverFirewallRuleGroupAssociationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN (Amazon Resource Name) of the firewall rule group association.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The unique identifier of the firewall rule group.
        /// </summary>
        [Input("firewallRuleGroupId")]
        public Input<string>? FirewallRuleGroupId { get; set; }

        /// <summary>
        /// If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections. Valid values: `ENABLED`, `DISABLED`.
        /// </summary>
        [Input("mutationProtection")]
        public Input<string>? MutationProtection { get; set; }

        /// <summary>
        /// A name that lets you identify the rule group association, to manage and use it.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The setting that determines the processing order of the rule group among the rule groups that you associate with the specified VPC. DNS Firewall filters VPC traffic starting from the rule group with the lowest numeric priority setting.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        /// <summary>
        /// The unique identifier of the VPC that you want to associate with the rule group.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public ResolverFirewallRuleGroupAssociationState()
        {
        }
        public static new ResolverFirewallRuleGroupAssociationState Empty => new ResolverFirewallRuleGroupAssociationState();
    }
}
