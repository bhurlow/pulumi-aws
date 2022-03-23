// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudSearch
{
    /// <summary>
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
    ///         var example = new Aws.CloudSearch.Domain("example", new Aws.CloudSearch.DomainArgs
    ///         {
    ///             IndexFields = 
    ///             {
    ///                 new Aws.CloudSearch.Inputs.DomainIndexFieldArgs
    ///                 {
    ///                     AnalysisScheme = "_en_default_",
    ///                     Highlight = false,
    ///                     Name = "headline",
    ///                     Return = true,
    ///                     Search = true,
    ///                     Sort = true,
    ///                     Type = "text",
    ///                 },
    ///                 new Aws.CloudSearch.Inputs.DomainIndexFieldArgs
    ///                 {
    ///                     Facet = true,
    ///                     Name = "price",
    ///                     Return = true,
    ///                     Search = true,
    ///                     Sort = true,
    ///                     Type = "double",
    ///                 },
    ///             },
    ///             ScalingParameters = new Aws.CloudSearch.Inputs.DomainScalingParametersArgs
    ///             {
    ///                 DesiredInstanceType = "search.medium",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CloudSearch Domains can be imported using the `name`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:cloudsearch/domain:Domain example example-domain
    /// ```
    /// </summary>
    [AwsResourceType("aws:cloudsearch/domain:Domain")]
    public partial class Domain : Pulumi.CustomResource
    {
        /// <summary>
        /// The domain's ARN.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The service endpoint for updating documents in a search domain.
        /// </summary>
        [Output("documentServiceEndpoint")]
        public Output<string> DocumentServiceEndpoint { get; private set; } = null!;

        /// <summary>
        /// An internally generated unique identifier for the domain.
        /// </summary>
        [Output("domainId")]
        public Output<string> DomainId { get; private set; } = null!;

        /// <summary>
        /// Domain endpoint options. Documented below.
        /// </summary>
        [Output("endpointOptions")]
        public Output<Outputs.DomainEndpointOptions> EndpointOptions { get; private set; } = null!;

        /// <summary>
        /// The index fields for documents added to the domain. Documented below.
        /// </summary>
        [Output("indexFields")]
        public Output<ImmutableArray<Outputs.DomainIndexField>> IndexFields { get; private set; } = null!;

        /// <summary>
        /// Whether or not to maintain extra instances for the domain in a second Availability Zone to ensure high availability.
        /// </summary>
        [Output("multiAz")]
        public Output<bool> MultiAz { get; private set; } = null!;

        /// <summary>
        /// A unique name for the field. Field names must begin with a letter and be at least 3 and no more than 64 characters long. The allowed characters are: `a`-`z` (lower-case letters), `0`-`9`, and `_` (underscore). The name `score` is reserved and cannot be used as a field name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Domain scaling parameters. Documented below.
        /// </summary>
        [Output("scalingParameters")]
        public Output<Outputs.DomainScalingParameters> ScalingParameters { get; private set; } = null!;

        /// <summary>
        /// The service endpoint for requesting search results from a search domain.
        /// </summary>
        [Output("searchServiceEndpoint")]
        public Output<string> SearchServiceEndpoint { get; private set; } = null!;


        /// <summary>
        /// Create a Domain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Domain(string name, DomainArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:cloudsearch/domain:Domain", name, args ?? new DomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Domain(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudsearch/domain:Domain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Domain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Domain Get(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
        {
            return new Domain(name, id, state, options);
        }
    }

    public sealed class DomainArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Domain endpoint options. Documented below.
        /// </summary>
        [Input("endpointOptions")]
        public Input<Inputs.DomainEndpointOptionsArgs>? EndpointOptions { get; set; }

        [Input("indexFields")]
        private InputList<Inputs.DomainIndexFieldArgs>? _indexFields;

        /// <summary>
        /// The index fields for documents added to the domain. Documented below.
        /// </summary>
        public InputList<Inputs.DomainIndexFieldArgs> IndexFields
        {
            get => _indexFields ?? (_indexFields = new InputList<Inputs.DomainIndexFieldArgs>());
            set => _indexFields = value;
        }

        /// <summary>
        /// Whether or not to maintain extra instances for the domain in a second Availability Zone to ensure high availability.
        /// </summary>
        [Input("multiAz")]
        public Input<bool>? MultiAz { get; set; }

        /// <summary>
        /// A unique name for the field. Field names must begin with a letter and be at least 3 and no more than 64 characters long. The allowed characters are: `a`-`z` (lower-case letters), `0`-`9`, and `_` (underscore). The name `score` is reserved and cannot be used as a field name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Domain scaling parameters. Documented below.
        /// </summary>
        [Input("scalingParameters")]
        public Input<Inputs.DomainScalingParametersArgs>? ScalingParameters { get; set; }

        public DomainArgs()
        {
        }
    }

    public sealed class DomainState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain's ARN.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The service endpoint for updating documents in a search domain.
        /// </summary>
        [Input("documentServiceEndpoint")]
        public Input<string>? DocumentServiceEndpoint { get; set; }

        /// <summary>
        /// An internally generated unique identifier for the domain.
        /// </summary>
        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        /// <summary>
        /// Domain endpoint options. Documented below.
        /// </summary>
        [Input("endpointOptions")]
        public Input<Inputs.DomainEndpointOptionsGetArgs>? EndpointOptions { get; set; }

        [Input("indexFields")]
        private InputList<Inputs.DomainIndexFieldGetArgs>? _indexFields;

        /// <summary>
        /// The index fields for documents added to the domain. Documented below.
        /// </summary>
        public InputList<Inputs.DomainIndexFieldGetArgs> IndexFields
        {
            get => _indexFields ?? (_indexFields = new InputList<Inputs.DomainIndexFieldGetArgs>());
            set => _indexFields = value;
        }

        /// <summary>
        /// Whether or not to maintain extra instances for the domain in a second Availability Zone to ensure high availability.
        /// </summary>
        [Input("multiAz")]
        public Input<bool>? MultiAz { get; set; }

        /// <summary>
        /// A unique name for the field. Field names must begin with a letter and be at least 3 and no more than 64 characters long. The allowed characters are: `a`-`z` (lower-case letters), `0`-`9`, and `_` (underscore). The name `score` is reserved and cannot be used as a field name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Domain scaling parameters. Documented below.
        /// </summary>
        [Input("scalingParameters")]
        public Input<Inputs.DomainScalingParametersGetArgs>? ScalingParameters { get; set; }

        /// <summary>
        /// The service endpoint for requesting search results from a search domain.
        /// </summary>
        [Input("searchServiceEndpoint")]
        public Input<string>? SearchServiceEndpoint { get; set; }

        public DomainState()
        {
        }
    }
}