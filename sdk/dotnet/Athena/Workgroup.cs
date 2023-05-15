// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Athena
{
    /// <summary>
    /// Provides an Athena Workgroup.
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
    ///     var example = new Aws.Athena.Workgroup("example", new()
    ///     {
    ///         Configuration = new Aws.Athena.Inputs.WorkgroupConfigurationArgs
    ///         {
    ///             EnforceWorkgroupConfiguration = true,
    ///             PublishCloudwatchMetricsEnabled = true,
    ///             ResultConfiguration = new Aws.Athena.Inputs.WorkgroupConfigurationResultConfigurationArgs
    ///             {
    ///                 OutputLocation = $"s3://{aws_s3_bucket.Example.Bucket}/output/",
    ///                 EncryptionConfiguration = new Aws.Athena.Inputs.WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs
    ///                 {
    ///                     EncryptionOption = "SSE_KMS",
    ///                     KmsKeyArn = aws_kms_key.Example.Arn,
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Athena Workgroups can be imported using their name, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:athena/workgroup:Workgroup example example
    /// ```
    /// </summary>
    [AwsResourceType("aws:athena/workgroup:Workgroup")]
    public partial class Workgroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN of the workgroup
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Configuration block with various settings for the workgroup. Documented below.
        /// </summary>
        [Output("configuration")]
        public Output<Outputs.WorkgroupConfiguration?> Configuration { get; private set; } = null!;

        /// <summary>
        /// Description of the workgroup.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Option to delete the workgroup and its contents even if the workgroup contains any named queries.
        /// </summary>
        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;

        /// <summary>
        /// Name of the workgroup.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
        /// </summary>
        [Output("state")]
        public Output<string?> State { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags for the workgroup. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a Workgroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Workgroup(string name, WorkgroupArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:athena/workgroup:Workgroup", name, args ?? new WorkgroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Workgroup(string name, Input<string> id, WorkgroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:athena/workgroup:Workgroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Workgroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Workgroup Get(string name, Input<string> id, WorkgroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Workgroup(name, id, state, options);
        }
    }

    public sealed class WorkgroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration block with various settings for the workgroup. Documented below.
        /// </summary>
        [Input("configuration")]
        public Input<Inputs.WorkgroupConfigurationArgs>? Configuration { get; set; }

        /// <summary>
        /// Description of the workgroup.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Option to delete the workgroup and its contents even if the workgroup contains any named queries.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// Name of the workgroup.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags for the workgroup. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public WorkgroupArgs()
        {
        }
        public static new WorkgroupArgs Empty => new WorkgroupArgs();
    }

    public sealed class WorkgroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the workgroup
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Configuration block with various settings for the workgroup. Documented below.
        /// </summary>
        [Input("configuration")]
        public Input<Inputs.WorkgroupConfigurationGetArgs>? Configuration { get; set; }

        /// <summary>
        /// Description of the workgroup.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Option to delete the workgroup and its contents even if the workgroup contains any named queries.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// Name of the workgroup.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags for the workgroup. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public WorkgroupState()
        {
        }
        public static new WorkgroupState Empty => new WorkgroupState();
    }
}
