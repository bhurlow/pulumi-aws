// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Macie
{
    /// <summary>
    /// Provides a resource to manage an [Amazon Macie Findings Filter](https://docs.aws.amazon.com/macie/latest/APIReference/findingsfilters-id.html).
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
    ///     var example = new Aws.Macie2.Account("example");
    /// 
    ///     var test = new Aws.Macie.FindingsFilter("test", new()
    ///     {
    ///         Description = "DESCRIPTION",
    ///         Position = 1,
    ///         Action = "ARCHIVE",
    ///         FindingCriteria = new Aws.Macie.Inputs.FindingsFilterFindingCriteriaArgs
    ///         {
    ///             Criterions = new[]
    ///             {
    ///                 new Aws.Macie.Inputs.FindingsFilterFindingCriteriaCriterionArgs
    ///                 {
    ///                     Field = "region",
    ///                     Eqs = new[]
    ///                     {
    ///                         data.Aws_region.Current.Name,
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             aws_macie2_account.Test,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// `aws_macie2_findings_filter` can be imported using the id, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:macie/findingsFilter:FindingsFilter example abcd1
    /// ```
    /// </summary>
    [AwsResourceType("aws:macie/findingsFilter:FindingsFilter")]
    public partial class FindingsFilter : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The action to perform on findings that meet the filter criteria (`finding_criteria`). Valid values are: `ARCHIVE`, suppress (automatically archive) the findings; and, `NOOP`, don't perform any action on the findings.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Findings Filter.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A custom description of the filter. The description can contain as many as 512 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The criteria to use to filter findings.
        /// </summary>
        [Output("findingCriteria")]
        public Output<Outputs.FindingsFilterFindingCriteria> FindingCriteria { get; private set; } = null!;

        /// <summary>
        /// A custom name for the filter. The name must contain at least 3 characters and can contain as many as 64 characters. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// The position of the filter in the list of saved filters on the Amazon Macie console. This value also determines the order in which the filter is applied to findings, relative to other filters that are also applied to the findings.
        /// </summary>
        [Output("position")]
        public Output<int> Position { get; private set; } = null!;

        /// <summary>
        /// A map of key-value pairs that specifies the tags to associate with the filter.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a FindingsFilter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FindingsFilter(string name, FindingsFilterArgs args, CustomResourceOptions? options = null)
            : base("aws:macie/findingsFilter:FindingsFilter", name, args ?? new FindingsFilterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FindingsFilter(string name, Input<string> id, FindingsFilterState? state = null, CustomResourceOptions? options = null)
            : base("aws:macie/findingsFilter:FindingsFilter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FindingsFilter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FindingsFilter Get(string name, Input<string> id, FindingsFilterState? state = null, CustomResourceOptions? options = null)
        {
            return new FindingsFilter(name, id, state, options);
        }
    }

    public sealed class FindingsFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action to perform on findings that meet the filter criteria (`finding_criteria`). Valid values are: `ARCHIVE`, suppress (automatically archive) the findings; and, `NOOP`, don't perform any action on the findings.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// A custom description of the filter. The description can contain as many as 512 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The criteria to use to filter findings.
        /// </summary>
        [Input("findingCriteria", required: true)]
        public Input<Inputs.FindingsFilterFindingCriteriaArgs> FindingCriteria { get; set; } = null!;

        /// <summary>
        /// A custom name for the filter. The name must contain at least 3 characters and can contain as many as 64 characters. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The position of the filter in the list of saved filters on the Amazon Macie console. This value also determines the order in which the filter is applied to findings, relative to other filters that are also applied to the findings.
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of key-value pairs that specifies the tags to associate with the filter.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public FindingsFilterArgs()
        {
        }
        public static new FindingsFilterArgs Empty => new FindingsFilterArgs();
    }

    public sealed class FindingsFilterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action to perform on findings that meet the filter criteria (`finding_criteria`). Valid values are: `ARCHIVE`, suppress (automatically archive) the findings; and, `NOOP`, don't perform any action on the findings.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Findings Filter.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// A custom description of the filter. The description can contain as many as 512 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The criteria to use to filter findings.
        /// </summary>
        [Input("findingCriteria")]
        public Input<Inputs.FindingsFilterFindingCriteriaGetArgs>? FindingCriteria { get; set; }

        /// <summary>
        /// A custom name for the filter. The name must contain at least 3 characters and can contain as many as 64 characters. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The position of the filter in the list of saved filters on the Amazon Macie console. This value also determines the order in which the filter is applied to findings, relative to other filters that are also applied to the findings.
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of key-value pairs that specifies the tags to associate with the filter.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public FindingsFilterState()
        {
        }
        public static new FindingsFilterState Empty => new FindingsFilterState();
    }
}
