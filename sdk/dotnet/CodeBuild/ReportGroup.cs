// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeBuild
{
    /// <summary>
    /// Provides a CodeBuild Report Groups Resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var current = Output.Create(Aws.GetCallerIdentity.InvokeAsync());
    ///         var exampleKey = new Aws.Kms.Key("exampleKey", new Aws.Kms.KeyArgs
    ///         {
    ///             Description = "my test kms key",
    ///             DeletionWindowInDays = 7,
    ///             Policy = current.Apply(current =&gt; @$"{{
    ///   ""Version"": ""2012-10-17"",
    ///   ""Id"": ""kms-tf-1"",
    ///   ""Statement"": [
    ///     {{
    ///       ""Sid"": ""Enable IAM User Permissions"",
    ///       ""Effect"": ""Allow"",
    ///       ""Principal"": {{
    ///         ""AWS"": ""arn:aws:iam::{current.AccountId}:root""
    ///       }},
    ///       ""Action"": ""kms:*"",
    ///       ""Resource"": ""*""
    ///     }}
    ///   ]
    /// }}
    /// "),
    ///         });
    ///         var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2", new Aws.S3.BucketV2Args
    ///         {
    ///         });
    ///         var exampleReportGroup = new Aws.CodeBuild.ReportGroup("exampleReportGroup", new Aws.CodeBuild.ReportGroupArgs
    ///         {
    ///             Type = "TEST",
    ///             ExportConfig = new Aws.CodeBuild.Inputs.ReportGroupExportConfigArgs
    ///             {
    ///                 Type = "S3",
    ///                 S3Destination = new Aws.CodeBuild.Inputs.ReportGroupExportConfigS3DestinationArgs
    ///                 {
    ///                     Bucket = exampleBucketV2.Id,
    ///                     EncryptionDisabled = false,
    ///                     EncryptionKey = exampleKey.Arn,
    ///                     Packaging = "NONE",
    ///                     Path = "/some",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CodeBuild Report Group can be imported using the CodeBuild Report Group arn, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:codebuild/reportGroup:ReportGroup example arn:aws:codebuild:us-west-2:123456789:report-group/report-group-name
    /// ```
    /// </summary>
    [AwsResourceType("aws:codebuild/reportGroup:ReportGroup")]
    public partial class ReportGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of Report Group.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The date and time this Report Group was created.
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// If `true`, deletes any reports that belong to a report group before deleting the report group. If `false`, you must delete any reports in the report group before deleting it. Default value is `false`.
        /// </summary>
        [Output("deleteReports")]
        public Output<bool?> DeleteReports { get; private set; } = null!;

        /// <summary>
        /// Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
        /// </summary>
        [Output("exportConfig")]
        public Output<Outputs.ReportGroupExportConfig> ExportConfig { get; private set; } = null!;

        /// <summary>
        /// The name of a Report Group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Key-value mapping of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The export configuration type. Valid values are `S3` and `NO_EXPORT`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ReportGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReportGroup(string name, ReportGroupArgs args, CustomResourceOptions? options = null)
            : base("aws:codebuild/reportGroup:ReportGroup", name, args ?? new ReportGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReportGroup(string name, Input<string> id, ReportGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:codebuild/reportGroup:ReportGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReportGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReportGroup Get(string name, Input<string> id, ReportGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ReportGroup(name, id, state, options);
        }
    }

    public sealed class ReportGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If `true`, deletes any reports that belong to a report group before deleting the report group. If `false`, you must delete any reports in the report group before deleting it. Default value is `false`.
        /// </summary>
        [Input("deleteReports")]
        public Input<bool>? DeleteReports { get; set; }

        /// <summary>
        /// Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
        /// </summary>
        [Input("exportConfig", required: true)]
        public Input<Inputs.ReportGroupExportConfigArgs> ExportConfig { get; set; } = null!;

        /// <summary>
        /// The name of a Report Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The export configuration type. Valid values are `S3` and `NO_EXPORT`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ReportGroupArgs()
        {
        }
    }

    public sealed class ReportGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of Report Group.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The date and time this Report Group was created.
        /// </summary>
        [Input("created")]
        public Input<string>? Created { get; set; }

        /// <summary>
        /// If `true`, deletes any reports that belong to a report group before deleting the report group. If `false`, you must delete any reports in the report group before deleting it. Default value is `false`.
        /// </summary>
        [Input("deleteReports")]
        public Input<bool>? DeleteReports { get; set; }

        /// <summary>
        /// Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
        /// </summary>
        [Input("exportConfig")]
        public Input<Inputs.ReportGroupExportConfigGetArgs>? ExportConfig { get; set; }

        /// <summary>
        /// The name of a Report Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// The export configuration type. Valid values are `S3` and `NO_EXPORT`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ReportGroupState()
        {
        }
    }
}
