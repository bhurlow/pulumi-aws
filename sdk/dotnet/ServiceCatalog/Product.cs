// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ServiceCatalog
{
    /// <summary>
    /// Manages a Service Catalog Product.
    /// 
    /// &gt; **NOTE:** The user or role that uses this resources must have the `cloudformation:GetTemplate` IAM policy permission. This policy permission is required when using the `template_physical_id` argument.
    /// 
    /// &gt; A "provisioning artifact" is also referred to as a "version." A "distributor" is also referred to as a "vendor."
    /// 
    /// ## Example Usage
    /// ### Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.ServiceCatalog.Product("example", new()
    ///     {
    ///         Owner = "example-owner",
    ///         ProvisioningArtifactParameters = new Aws.ServiceCatalog.Inputs.ProductProvisioningArtifactParametersArgs
    ///         {
    ///             TemplateUrl = "https://s3.amazonaws.com/cf-templates-ozkq9d3hgiq2-us-east-1/temp1.json",
    ///         },
    ///         Tags = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Type = "CLOUD_FORMATION_TEMPLATE",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import `aws_servicecatalog_product` using the product ID. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:servicecatalog/product:Product example prod-dnigbtea24ste
    /// ```
    /// </summary>
    [AwsResourceType("aws:servicecatalog/product:Product")]
    public partial class Product : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
        /// </summary>
        [Output("acceptLanguage")]
        public Output<string?> AcceptLanguage { get; private set; } = null!;

        /// <summary>
        /// ARN of the product.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Time when the product was created.
        /// </summary>
        [Output("createdTime")]
        public Output<string> CreatedTime { get; private set; } = null!;

        /// <summary>
        /// Description of the product.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Distributor (i.e., vendor) of the product.
        /// </summary>
        [Output("distributor")]
        public Output<string> Distributor { get; private set; } = null!;

        /// <summary>
        /// Whether the product has a default path. If the product does not have a default path, call `ListLaunchPaths` to disambiguate between paths.  Otherwise, `ListLaunchPaths` is not required, and the output of ProductViewSummary can be used directly with `DescribeProvisioningParameters`.
        /// </summary>
        [Output("hasDefaultPath")]
        public Output<bool> HasDefaultPath { get; private set; } = null!;

        /// <summary>
        /// Name of the product.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Owner of the product.
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// Configuration block for provisioning artifact (i.e., version) parameters. Detailed below.
        /// </summary>
        [Output("provisioningArtifactParameters")]
        public Output<Outputs.ProductProvisioningArtifactParameters> ProvisioningArtifactParameters { get; private set; } = null!;

        /// <summary>
        /// Status of the product.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Support information about the product.
        /// </summary>
        [Output("supportDescription")]
        public Output<string> SupportDescription { get; private set; } = null!;

        /// <summary>
        /// Contact email for product support.
        /// </summary>
        [Output("supportEmail")]
        public Output<string> SupportEmail { get; private set; } = null!;

        /// <summary>
        /// Contact URL for product support.
        /// </summary>
        [Output("supportUrl")]
        public Output<string> SupportUrl { get; private set; } = null!;

        /// <summary>
        /// Tags to apply to the product. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Type of product. See [AWS Docs](https://docs.aws.amazon.com/servicecatalog/latest/dg/API_CreateProduct.html#API_CreateProduct_RequestSyntax) for valid list of values.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Product resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Product(string name, ProductArgs args, CustomResourceOptions? options = null)
            : base("aws:servicecatalog/product:Product", name, args ?? new ProductArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Product(string name, Input<string> id, ProductState? state = null, CustomResourceOptions? options = null)
            : base("aws:servicecatalog/product:Product", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Product resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Product Get(string name, Input<string> id, ProductState? state = null, CustomResourceOptions? options = null)
        {
            return new Product(name, id, state, options);
        }
    }

    public sealed class ProductArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
        /// </summary>
        [Input("acceptLanguage")]
        public Input<string>? AcceptLanguage { get; set; }

        /// <summary>
        /// Description of the product.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Distributor (i.e., vendor) of the product.
        /// </summary>
        [Input("distributor")]
        public Input<string>? Distributor { get; set; }

        /// <summary>
        /// Name of the product.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Owner of the product.
        /// </summary>
        [Input("owner", required: true)]
        public Input<string> Owner { get; set; } = null!;

        /// <summary>
        /// Configuration block for provisioning artifact (i.e., version) parameters. Detailed below.
        /// </summary>
        [Input("provisioningArtifactParameters", required: true)]
        public Input<Inputs.ProductProvisioningArtifactParametersArgs> ProvisioningArtifactParameters { get; set; } = null!;

        /// <summary>
        /// Support information about the product.
        /// </summary>
        [Input("supportDescription")]
        public Input<string>? SupportDescription { get; set; }

        /// <summary>
        /// Contact email for product support.
        /// </summary>
        [Input("supportEmail")]
        public Input<string>? SupportEmail { get; set; }

        /// <summary>
        /// Contact URL for product support.
        /// </summary>
        [Input("supportUrl")]
        public Input<string>? SupportUrl { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Tags to apply to the product. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Type of product. See [AWS Docs](https://docs.aws.amazon.com/servicecatalog/latest/dg/API_CreateProduct.html#API_CreateProduct_RequestSyntax) for valid list of values.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ProductArgs()
        {
        }
        public static new ProductArgs Empty => new ProductArgs();
    }

    public sealed class ProductState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
        /// </summary>
        [Input("acceptLanguage")]
        public Input<string>? AcceptLanguage { get; set; }

        /// <summary>
        /// ARN of the product.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Time when the product was created.
        /// </summary>
        [Input("createdTime")]
        public Input<string>? CreatedTime { get; set; }

        /// <summary>
        /// Description of the product.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Distributor (i.e., vendor) of the product.
        /// </summary>
        [Input("distributor")]
        public Input<string>? Distributor { get; set; }

        /// <summary>
        /// Whether the product has a default path. If the product does not have a default path, call `ListLaunchPaths` to disambiguate between paths.  Otherwise, `ListLaunchPaths` is not required, and the output of ProductViewSummary can be used directly with `DescribeProvisioningParameters`.
        /// </summary>
        [Input("hasDefaultPath")]
        public Input<bool>? HasDefaultPath { get; set; }

        /// <summary>
        /// Name of the product.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Owner of the product.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        /// <summary>
        /// Configuration block for provisioning artifact (i.e., version) parameters. Detailed below.
        /// </summary>
        [Input("provisioningArtifactParameters")]
        public Input<Inputs.ProductProvisioningArtifactParametersGetArgs>? ProvisioningArtifactParameters { get; set; }

        /// <summary>
        /// Status of the product.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Support information about the product.
        /// </summary>
        [Input("supportDescription")]
        public Input<string>? SupportDescription { get; set; }

        /// <summary>
        /// Contact email for product support.
        /// </summary>
        [Input("supportEmail")]
        public Input<string>? SupportEmail { get; set; }

        /// <summary>
        /// Contact URL for product support.
        /// </summary>
        [Input("supportUrl")]
        public Input<string>? SupportUrl { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Tags to apply to the product. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        /// <summary>
        /// Type of product. See [AWS Docs](https://docs.aws.amazon.com/servicecatalog/latest/dg/API_CreateProduct.html#API_CreateProduct_RequestSyntax) for valid list of values.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ProductState()
        {
        }
        public static new ProductState Empty => new ProductState();
    }
}
