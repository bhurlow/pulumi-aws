// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Quicksight
{
    /// <summary>
    /// Resource for managing a QuickSight Template.
    /// 
    /// ## Example Usage
    /// ### From Source Template
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.Quicksight.Template("example", new()
    ///     {
    ///         TemplateId = "example-id",
    ///         VersionDescription = "version",
    ///         SourceEntity = new Aws.Quicksight.Inputs.TemplateSourceEntityArgs
    ///         {
    ///             SourceTemplate = new Aws.Quicksight.Inputs.TemplateSourceEntitySourceTemplateArgs
    ///             {
    ///                 Arn = aws_quicksight_template.Source.Arn,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// A QuickSight Template can be imported using the AWS account ID and template ID separated by a comma (`,`) e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:quicksight/template:Template example 123456789012,example-id
    /// ```
    /// </summary>
    [AwsResourceType("aws:quicksight/template:Template")]
    public partial class Template : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the resource.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// AWS account ID.
        /// </summary>
        [Output("awsAccountId")]
        public Output<string> AwsAccountId { get; private set; } = null!;

        /// <summary>
        /// The time that the template was created.
        /// </summary>
        [Output("createdTime")]
        public Output<string> CreatedTime { get; private set; } = null!;

        /// <summary>
        /// The time that the template was last updated.
        /// </summary>
        [Output("lastUpdatedTime")]
        public Output<string> LastUpdatedTime { get; private set; } = null!;

        /// <summary>
        /// Display name for the template.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A set of resource permissions on the template. Maximum of 64 items. See permissions.
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableArray<Outputs.TemplatePermission>> Permissions { get; private set; } = null!;

        /// <summary>
        /// The entity that you are using as a source when you create the template (analysis or template). Only one of `definition` or `source_entity` should be configured. See source_entity.
        /// </summary>
        [Output("sourceEntity")]
        public Output<Outputs.TemplateSourceEntity?> SourceEntity { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of an analysis or template that was used to create this template.
        /// </summary>
        [Output("sourceEntityArn")]
        public Output<string> SourceEntityArn { get; private set; } = null!;

        /// <summary>
        /// The template creation status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Identifier for the template.
        /// </summary>
        [Output("templateId")]
        public Output<string> TemplateId { get; private set; } = null!;

        /// <summary>
        /// A description of the current template version being created/updated.
        /// </summary>
        [Output("versionDescription")]
        public Output<string> VersionDescription { get; private set; } = null!;

        /// <summary>
        /// The version number of the template version.
        /// </summary>
        [Output("versionNumber")]
        public Output<int> VersionNumber { get; private set; } = null!;


        /// <summary>
        /// Create a Template resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Template(string name, TemplateArgs args, CustomResourceOptions? options = null)
            : base("aws:quicksight/template:Template", name, args ?? new TemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Template(string name, Input<string> id, TemplateState? state = null, CustomResourceOptions? options = null)
            : base("aws:quicksight/template:Template", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Template resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Template Get(string name, Input<string> id, TemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new Template(name, id, state, options);
        }
    }

    public sealed class TemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AWS account ID.
        /// </summary>
        [Input("awsAccountId")]
        public Input<string>? AwsAccountId { get; set; }

        /// <summary>
        /// Display name for the template.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("permissions")]
        private InputList<Inputs.TemplatePermissionArgs>? _permissions;

        /// <summary>
        /// A set of resource permissions on the template. Maximum of 64 items. See permissions.
        /// </summary>
        public InputList<Inputs.TemplatePermissionArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.TemplatePermissionArgs>());
            set => _permissions = value;
        }

        /// <summary>
        /// The entity that you are using as a source when you create the template (analysis or template). Only one of `definition` or `source_entity` should be configured. See source_entity.
        /// </summary>
        [Input("sourceEntity")]
        public Input<Inputs.TemplateSourceEntityArgs>? SourceEntity { get; set; }

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
        /// Identifier for the template.
        /// </summary>
        [Input("templateId", required: true)]
        public Input<string> TemplateId { get; set; } = null!;

        /// <summary>
        /// A description of the current template version being created/updated.
        /// </summary>
        [Input("versionDescription", required: true)]
        public Input<string> VersionDescription { get; set; } = null!;

        public TemplateArgs()
        {
        }
        public static new TemplateArgs Empty => new TemplateArgs();
    }

    public sealed class TemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the resource.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// AWS account ID.
        /// </summary>
        [Input("awsAccountId")]
        public Input<string>? AwsAccountId { get; set; }

        /// <summary>
        /// The time that the template was created.
        /// </summary>
        [Input("createdTime")]
        public Input<string>? CreatedTime { get; set; }

        /// <summary>
        /// The time that the template was last updated.
        /// </summary>
        [Input("lastUpdatedTime")]
        public Input<string>? LastUpdatedTime { get; set; }

        /// <summary>
        /// Display name for the template.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("permissions")]
        private InputList<Inputs.TemplatePermissionGetArgs>? _permissions;

        /// <summary>
        /// A set of resource permissions on the template. Maximum of 64 items. See permissions.
        /// </summary>
        public InputList<Inputs.TemplatePermissionGetArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.TemplatePermissionGetArgs>());
            set => _permissions = value;
        }

        /// <summary>
        /// The entity that you are using as a source when you create the template (analysis or template). Only one of `definition` or `source_entity` should be configured. See source_entity.
        /// </summary>
        [Input("sourceEntity")]
        public Input<Inputs.TemplateSourceEntityGetArgs>? SourceEntity { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of an analysis or template that was used to create this template.
        /// </summary>
        [Input("sourceEntityArn")]
        public Input<string>? SourceEntityArn { get; set; }

        /// <summary>
        /// The template creation status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

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
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// Identifier for the template.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        /// <summary>
        /// A description of the current template version being created/updated.
        /// </summary>
        [Input("versionDescription")]
        public Input<string>? VersionDescription { get; set; }

        /// <summary>
        /// The version number of the template version.
        /// </summary>
        [Input("versionNumber")]
        public Input<int>? VersionNumber { get; set; }

        public TemplateState()
        {
        }
        public static new TemplateState Empty => new TemplateState();
    }
}
