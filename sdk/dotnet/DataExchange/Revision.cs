// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DataExchange
{
    /// <summary>
    /// Provides a resource to manage AWS Data Exchange Revisions.
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
    ///     var example = new Aws.DataExchange.Revision("example", new()
    ///     {
    ///         DataSetId = aws_dataexchange_data_set.Example.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// DataExchange Revisions can be imported by their `data-set-id:revision-id`
    /// 
    /// ```sh
    ///  $ pulumi import aws:dataexchange/revision:Revision example 4fa784c7-ccb4-4dbf-ba4f-02198320daa1:4fa784c7-ccb4-4dbf-ba4f-02198320daa1
    /// ```
    /// </summary>
    [AwsResourceType("aws:dataexchange/revision:Revision")]
    public partial class Revision : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name of this data set.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// An optional comment about the revision.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// The dataset id.
        /// </summary>
        [Output("dataSetId")]
        public Output<string> DataSetId { get; private set; } = null!;

        /// <summary>
        /// The Id of the revision.
        /// </summary>
        [Output("revisionId")]
        public Output<string> RevisionId { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a Revision resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Revision(string name, RevisionArgs args, CustomResourceOptions? options = null)
            : base("aws:dataexchange/revision:Revision", name, args ?? new RevisionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Revision(string name, Input<string> id, RevisionState? state = null, CustomResourceOptions? options = null)
            : base("aws:dataexchange/revision:Revision", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Revision resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Revision Get(string name, Input<string> id, RevisionState? state = null, CustomResourceOptions? options = null)
        {
            return new Revision(name, id, state, options);
        }
    }

    public sealed class RevisionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional comment about the revision.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// The dataset id.
        /// </summary>
        [Input("dataSetId", required: true)]
        public Input<string> DataSetId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public RevisionArgs()
        {
        }
        public static new RevisionArgs Empty => new RevisionArgs();
    }

    public sealed class RevisionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name of this data set.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// An optional comment about the revision.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// The dataset id.
        /// </summary>
        [Input("dataSetId")]
        public Input<string>? DataSetId { get; set; }

        /// <summary>
        /// The Id of the revision.
        /// </summary>
        [Input("revisionId")]
        public Input<string>? RevisionId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public RevisionState()
        {
        }
        public static new RevisionState Empty => new RevisionState();
    }
}
