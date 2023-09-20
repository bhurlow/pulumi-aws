// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AccessAnalyzer
{
    /// <summary>
    /// Manages an Access Analyzer Analyzer. More information can be found in the [Access Analyzer User Guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/what-is-access-analyzer.html).
    /// 
    /// ## Example Usage
    /// ### Account Analyzer
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.AccessAnalyzer.Analyzer("example", new()
    ///     {
    ///         AnalyzerName = "example",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Organization Analyzer
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleOrganization = new Aws.Organizations.Organization("exampleOrganization", new()
    ///     {
    ///         AwsServiceAccessPrincipals = new[]
    ///         {
    ///             "access-analyzer.amazonaws.com",
    ///         },
    ///     });
    /// 
    ///     var exampleAnalyzer = new Aws.AccessAnalyzer.Analyzer("exampleAnalyzer", new()
    ///     {
    ///         AnalyzerName = "example",
    ///         Type = "ORGANIZATION",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             exampleOrganization,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import Access Analyzer Analyzers using the `analyzer_name`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:accessanalyzer/analyzer:Analyzer example example
    /// ```
    /// </summary>
    [AwsResourceType("aws:accessanalyzer/analyzer:Analyzer")]
    public partial class Analyzer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the Analyzer.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("analyzerName")]
        public Output<string> AnalyzerName { get; private set; } = null!;

        /// <summary>
        /// ARN of the Analyzer.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Type of Analyzer. Valid values are `ACCOUNT` or `ORGANIZATION`. Defaults to `ACCOUNT`.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Analyzer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Analyzer(string name, AnalyzerArgs args, CustomResourceOptions? options = null)
            : base("aws:accessanalyzer/analyzer:Analyzer", name, args ?? new AnalyzerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Analyzer(string name, Input<string> id, AnalyzerState? state = null, CustomResourceOptions? options = null)
            : base("aws:accessanalyzer/analyzer:Analyzer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Analyzer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Analyzer Get(string name, Input<string> id, AnalyzerState? state = null, CustomResourceOptions? options = null)
        {
            return new Analyzer(name, id, state, options);
        }
    }

    public sealed class AnalyzerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Analyzer.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("analyzerName", required: true)]
        public Input<string> AnalyzerName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Type of Analyzer. Valid values are `ACCOUNT` or `ORGANIZATION`. Defaults to `ACCOUNT`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AnalyzerArgs()
        {
        }
        public static new AnalyzerArgs Empty => new AnalyzerArgs();
    }

    public sealed class AnalyzerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Analyzer.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("analyzerName")]
        public Input<string>? AnalyzerName { get; set; }

        /// <summary>
        /// ARN of the Analyzer.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
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

        /// <summary>
        /// Type of Analyzer. Valid values are `ACCOUNT` or `ORGANIZATION`. Defaults to `ACCOUNT`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AnalyzerState()
        {
        }
        public static new AnalyzerState Empty => new AnalyzerState();
    }
}
