// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.EcrPublic
{
    /// <summary>
    /// Provides a Public Elastic Container Registry Repository.
    /// 
    /// &gt; **NOTE:** This resource can only be used in the `us-east-1` region.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System;
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// 	private static string ReadFileBase64(string path) {
    /// 		return Convert.ToBase64String(System.Text.Encoding.UTF8.GetBytes(File.ReadAllText(path)));
    /// 	}
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var usEast1 = new Aws.Provider("usEast1", new()
    ///     {
    ///         Region = "us-east-1",
    ///     });
    /// 
    ///     var foo = new Aws.EcrPublic.Repository("foo", new()
    ///     {
    ///         RepositoryName = "bar",
    ///         CatalogData = new Aws.EcrPublic.Inputs.RepositoryCatalogDataArgs
    ///         {
    ///             AboutText = "About Text",
    ///             Architectures = new[]
    ///             {
    ///                 "ARM",
    ///             },
    ///             Description = "Description",
    ///             LogoImageBlob = ReadFileBase64(image.Png),
    ///             OperatingSystems = new[]
    ///             {
    ///                 "Linux",
    ///             },
    ///             UsageText = "Usage Text",
    ///         },
    ///         Tags = 
    ///         {
    ///             { "env", "production" },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = aws.Us_east_1,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECR Public Repositories can be imported using the `repository_name`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:ecrpublic/repository:Repository example example
    /// ```
    /// </summary>
    [AwsResourceType("aws:ecrpublic/repository:Repository")]
    public partial class Repository : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Full ARN of the repository.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Catalog data configuration for the repository. See below for schema.
        /// </summary>
        [Output("catalogData")]
        public Output<Outputs.RepositoryCatalogData?> CatalogData { get; private set; } = null!;

        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;

        /// <summary>
        /// The registry ID where the repository was created.
        /// </summary>
        [Output("registryId")]
        public Output<string> RegistryId { get; private set; } = null!;

        /// <summary>
        /// Name of the repository.
        /// </summary>
        [Output("repositoryName")]
        public Output<string> RepositoryName { get; private set; } = null!;

        /// <summary>
        /// The URI of the repository.
        /// </summary>
        [Output("repositoryUri")]
        public Output<string> RepositoryUri { get; private set; } = null!;

        /// <summary>
        /// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a Repository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Repository(string name, RepositoryArgs args, CustomResourceOptions? options = null)
            : base("aws:ecrpublic/repository:Repository", name, args ?? new RepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Repository(string name, Input<string> id, RepositoryState? state = null, CustomResourceOptions? options = null)
            : base("aws:ecrpublic/repository:Repository", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Repository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Repository Get(string name, Input<string> id, RepositoryState? state = null, CustomResourceOptions? options = null)
        {
            return new Repository(name, id, state, options);
        }
    }

    public sealed class RepositoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Catalog data configuration for the repository. See below for schema.
        /// </summary>
        [Input("catalogData")]
        public Input<Inputs.RepositoryCatalogDataArgs>? CatalogData { get; set; }

        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// Name of the repository.
        /// </summary>
        [Input("repositoryName", required: true)]
        public Input<string> RepositoryName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public RepositoryArgs()
        {
        }
        public static new RepositoryArgs Empty => new RepositoryArgs();
    }

    public sealed class RepositoryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Full ARN of the repository.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Catalog data configuration for the repository. See below for schema.
        /// </summary>
        [Input("catalogData")]
        public Input<Inputs.RepositoryCatalogDataGetArgs>? CatalogData { get; set; }

        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// The registry ID where the repository was created.
        /// </summary>
        [Input("registryId")]
        public Input<string>? RegistryId { get; set; }

        /// <summary>
        /// Name of the repository.
        /// </summary>
        [Input("repositoryName")]
        public Input<string>? RepositoryName { get; set; }

        /// <summary>
        /// The URI of the repository.
        /// </summary>
        [Input("repositoryUri")]
        public Input<string>? RepositoryUri { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public RepositoryState()
        {
        }
        public static new RepositoryState Empty => new RepositoryState();
    }
}
