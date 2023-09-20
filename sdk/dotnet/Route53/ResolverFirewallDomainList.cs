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
    /// Provides a Route 53 Resolver DNS Firewall domain list resource.
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
    ///     var example = new Aws.Route53.ResolverFirewallDomainList("example");
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// In TODO v1.5.0 and later, use an `import` block to import
    /// 
    /// Route 53 Resolver DNS Firewall domain lists using the Route 53 Resolver DNS Firewall domain list ID. For exampleterraform import {
    /// 
    ///  to = aws_route53_resolver_firewall_domain_list.example
    /// 
    ///  id = "rslvr-fdl-0123456789abcdef" } Using `TODO import`, import
    /// 
    /// Route 53 Resolver DNS Firewall domain lists using the Route 53 Resolver DNS Firewall domain list ID. For exampleconsole % TODO import aws_route53_resolver_firewall_domain_list.example rslvr-fdl-0123456789abcdef
    /// </summary>
    [AwsResourceType("aws:route53/resolverFirewallDomainList:ResolverFirewallDomainList")]
    public partial class ResolverFirewallDomainList : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN (Amazon Resource Name) of the domain list.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A array of domains for the firewall domain list.
        /// </summary>
        [Output("domains")]
        public Output<ImmutableArray<string>> Domains { get; private set; } = null!;

        /// <summary>
        /// A name that lets you identify the domain list, to manage and use it.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. f configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a ResolverFirewallDomainList resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResolverFirewallDomainList(string name, ResolverFirewallDomainListArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:route53/resolverFirewallDomainList:ResolverFirewallDomainList", name, args ?? new ResolverFirewallDomainListArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResolverFirewallDomainList(string name, Input<string> id, ResolverFirewallDomainListState? state = null, CustomResourceOptions? options = null)
            : base("aws:route53/resolverFirewallDomainList:ResolverFirewallDomainList", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "tagsAll",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ResolverFirewallDomainList resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResolverFirewallDomainList Get(string name, Input<string> id, ResolverFirewallDomainListState? state = null, CustomResourceOptions? options = null)
        {
            return new ResolverFirewallDomainList(name, id, state, options);
        }
    }

    public sealed class ResolverFirewallDomainListArgs : global::Pulumi.ResourceArgs
    {
        [Input("domains")]
        private InputList<string>? _domains;

        /// <summary>
        /// A array of domains for the firewall domain list.
        /// </summary>
        public InputList<string> Domains
        {
            get => _domains ?? (_domains = new InputList<string>());
            set => _domains = value;
        }

        /// <summary>
        /// A name that lets you identify the domain list, to manage and use it.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. f configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ResolverFirewallDomainListArgs()
        {
        }
        public static new ResolverFirewallDomainListArgs Empty => new ResolverFirewallDomainListArgs();
    }

    public sealed class ResolverFirewallDomainListState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN (Amazon Resource Name) of the domain list.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("domains")]
        private InputList<string>? _domains;

        /// <summary>
        /// A array of domains for the firewall domain list.
        /// </summary>
        public InputList<string> Domains
        {
            get => _domains ?? (_domains = new InputList<string>());
            set => _domains = value;
        }

        /// <summary>
        /// A name that lets you identify the domain list, to manage and use it.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. f configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        [Obsolete(@"Please use `tags` instead.")]
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _tagsAll = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        public ResolverFirewallDomainListState()
        {
        }
        public static new ResolverFirewallDomainListState Empty => new ResolverFirewallDomainListState();
    }
}
