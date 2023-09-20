// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Waf
{
    /// <summary>
    /// Provides a WAF Rate Based Rule Resource
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
    ///     var ipset = new Aws.Waf.IpSet("ipset", new()
    ///     {
    ///         IpSetDescriptors = new[]
    ///         {
    ///             new Aws.Waf.Inputs.IpSetIpSetDescriptorArgs
    ///             {
    ///                 Type = "IPV4",
    ///                 Value = "192.0.7.0/24",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var wafrule = new Aws.Waf.RateBasedRule("wafrule", new()
    ///     {
    ///         MetricName = "tfWAFRule",
    ///         RateKey = "IP",
    ///         RateLimit = 100,
    ///         Predicates = new[]
    ///         {
    ///             new Aws.Waf.Inputs.RateBasedRulePredicateArgs
    ///             {
    ///                 DataId = ipset.Id,
    ///                 Negated = false,
    ///                 Type = "IPMatch",
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             ipset,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import WAF Rated Based Rule using the id. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:waf/rateBasedRule:RateBasedRule wafrule a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
    /// ```
    /// </summary>
    [AwsResourceType("aws:waf/rateBasedRule:RateBasedRule")]
    public partial class RateBasedRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name or description for the Amazon CloudWatch metric of this rule.
        /// </summary>
        [Output("metricName")]
        public Output<string> MetricName { get; private set; } = null!;

        /// <summary>
        /// The name or description of the rule.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The objects to include in a rule (documented below).
        /// </summary>
        [Output("predicates")]
        public Output<ImmutableArray<Outputs.RateBasedRulePredicate>> Predicates { get; private set; } = null!;

        /// <summary>
        /// Valid value is IP.
        /// </summary>
        [Output("rateKey")]
        public Output<string> RateKey { get; private set; } = null!;

        /// <summary>
        /// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
        /// </summary>
        [Output("rateLimit")]
        public Output<int> RateLimit { get; private set; } = null!;

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
        /// Create a RateBasedRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RateBasedRule(string name, RateBasedRuleArgs args, CustomResourceOptions? options = null)
            : base("aws:waf/rateBasedRule:RateBasedRule", name, args ?? new RateBasedRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RateBasedRule(string name, Input<string> id, RateBasedRuleState? state = null, CustomResourceOptions? options = null)
            : base("aws:waf/rateBasedRule:RateBasedRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RateBasedRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RateBasedRule Get(string name, Input<string> id, RateBasedRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new RateBasedRule(name, id, state, options);
        }
    }

    public sealed class RateBasedRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name or description for the Amazon CloudWatch metric of this rule.
        /// </summary>
        [Input("metricName", required: true)]
        public Input<string> MetricName { get; set; } = null!;

        /// <summary>
        /// The name or description of the rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("predicates")]
        private InputList<Inputs.RateBasedRulePredicateArgs>? _predicates;

        /// <summary>
        /// The objects to include in a rule (documented below).
        /// </summary>
        public InputList<Inputs.RateBasedRulePredicateArgs> Predicates
        {
            get => _predicates ?? (_predicates = new InputList<Inputs.RateBasedRulePredicateArgs>());
            set => _predicates = value;
        }

        /// <summary>
        /// Valid value is IP.
        /// </summary>
        [Input("rateKey", required: true)]
        public Input<string> RateKey { get; set; } = null!;

        /// <summary>
        /// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
        /// </summary>
        [Input("rateLimit", required: true)]
        public Input<int> RateLimit { get; set; } = null!;

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

        public RateBasedRuleArgs()
        {
        }
        public static new RateBasedRuleArgs Empty => new RateBasedRuleArgs();
    }

    public sealed class RateBasedRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name or description for the Amazon CloudWatch metric of this rule.
        /// </summary>
        [Input("metricName")]
        public Input<string>? MetricName { get; set; }

        /// <summary>
        /// The name or description of the rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("predicates")]
        private InputList<Inputs.RateBasedRulePredicateGetArgs>? _predicates;

        /// <summary>
        /// The objects to include in a rule (documented below).
        /// </summary>
        public InputList<Inputs.RateBasedRulePredicateGetArgs> Predicates
        {
            get => _predicates ?? (_predicates = new InputList<Inputs.RateBasedRulePredicateGetArgs>());
            set => _predicates = value;
        }

        /// <summary>
        /// Valid value is IP.
        /// </summary>
        [Input("rateKey")]
        public Input<string>? RateKey { get; set; }

        /// <summary>
        /// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
        /// </summary>
        [Input("rateLimit")]
        public Input<int>? RateLimit { get; set; }

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

        public RateBasedRuleState()
        {
        }
        public static new RateBasedRuleState Empty => new RateBasedRuleState();
    }
}
