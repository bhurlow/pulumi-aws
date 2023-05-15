// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker
{
    /// <summary>
    /// Provides a SageMaker Project resource.
    /// 
    ///  &gt; Note: If you are trying to use SageMaker projects with SageMaker studio you will need to add a tag with the key `sagemaker:studio-visibility` with value `true`. For more on requirements to use projects and permission needed see [AWS Docs](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-projects-templates-custom.html).
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
    ///     var example = new Aws.Sagemaker.Project("example", new()
    ///     {
    ///         ProjectName = "example",
    ///         ServiceCatalogProvisioningDetails = new Aws.Sagemaker.Inputs.ProjectServiceCatalogProvisioningDetailsArgs
    ///         {
    ///             ProductId = aws_servicecatalog_product.Example.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SageMaker Projects can be imported using the `project_name`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:sagemaker/project:Project example example
    /// ```
    /// </summary>
    [AwsResourceType("aws:sagemaker/project:Project")]
    public partial class Project : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) assigned by AWS to this Project.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A description for the project.
        /// </summary>
        [Output("projectDescription")]
        public Output<string?> ProjectDescription { get; private set; } = null!;

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The name of the Project.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// The product ID and provisioning artifact ID to provision a service catalog. See Service Catalog Provisioning Details below.
        /// </summary>
        [Output("serviceCatalogProvisioningDetails")]
        public Output<Outputs.ProjectServiceCatalogProvisioningDetails> ServiceCatalogProvisioningDetails { get; private set; } = null!;

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
        /// Create a Project resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Project(string name, ProjectArgs args, CustomResourceOptions? options = null)
            : base("aws:sagemaker/project:Project", name, args ?? new ProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Project(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
            : base("aws:sagemaker/project:Project", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Project resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Project Get(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
        {
            return new Project(name, id, state, options);
        }
    }

    public sealed class ProjectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description for the project.
        /// </summary>
        [Input("projectDescription")]
        public Input<string>? ProjectDescription { get; set; }

        /// <summary>
        /// The name of the Project.
        /// </summary>
        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        /// <summary>
        /// The product ID and provisioning artifact ID to provision a service catalog. See Service Catalog Provisioning Details below.
        /// </summary>
        [Input("serviceCatalogProvisioningDetails", required: true)]
        public Input<Inputs.ProjectServiceCatalogProvisioningDetailsArgs> ServiceCatalogProvisioningDetails { get; set; } = null!;

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

        public ProjectArgs()
        {
        }
        public static new ProjectArgs Empty => new ProjectArgs();
    }

    public sealed class ProjectState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) assigned by AWS to this Project.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// A description for the project.
        /// </summary>
        [Input("projectDescription")]
        public Input<string>? ProjectDescription { get; set; }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The name of the Project.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The product ID and provisioning artifact ID to provision a service catalog. See Service Catalog Provisioning Details below.
        /// </summary>
        [Input("serviceCatalogProvisioningDetails")]
        public Input<Inputs.ProjectServiceCatalogProvisioningDetailsGetArgs>? ServiceCatalogProvisioningDetails { get; set; }

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

        public ProjectState()
        {
        }
        public static new ProjectState Empty => new ProjectState();
    }
}
