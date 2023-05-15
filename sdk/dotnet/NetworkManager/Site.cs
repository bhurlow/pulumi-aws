// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.NetworkManager
{
    /// <summary>
    /// Creates a site in a global network.
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
    ///     var exampleGlobalNetwork = new Aws.NetworkManager.GlobalNetwork("exampleGlobalNetwork");
    /// 
    ///     var exampleSite = new Aws.NetworkManager.Site("exampleSite", new()
    ///     {
    ///         GlobalNetworkId = exampleGlobalNetwork.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// `aws_networkmanager_site` can be imported using the site ARN, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:networkmanager/site:Site example arn:aws:networkmanager::123456789012:site/global-network-0d47f6t230mz46dy4/site-444555aaabbb11223
    /// ```
    /// </summary>
    [AwsResourceType("aws:networkmanager/site:Site")]
    public partial class Site : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Site Amazon Resource Name (ARN)
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Description of the Site.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the Global Network to create the site in.
        /// </summary>
        [Output("globalNetworkId")]
        public Output<string> GlobalNetworkId { get; private set; } = null!;

        /// <summary>
        /// The site location as documented below.
        /// </summary>
        [Output("location")]
        public Output<Outputs.SiteLocation?> Location { get; private set; } = null!;

        /// <summary>
        /// Key-value tags for the Site. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a Site resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Site(string name, SiteArgs args, CustomResourceOptions? options = null)
            : base("aws:networkmanager/site:Site", name, args ?? new SiteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Site(string name, Input<string> id, SiteState? state = null, CustomResourceOptions? options = null)
            : base("aws:networkmanager/site:Site", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Site resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Site Get(string name, Input<string> id, SiteState? state = null, CustomResourceOptions? options = null)
        {
            return new Site(name, id, state, options);
        }
    }

    public sealed class SiteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the Site.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the Global Network to create the site in.
        /// </summary>
        [Input("globalNetworkId", required: true)]
        public Input<string> GlobalNetworkId { get; set; } = null!;

        /// <summary>
        /// The site location as documented below.
        /// </summary>
        [Input("location")]
        public Input<Inputs.SiteLocationArgs>? Location { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value tags for the Site. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public SiteArgs()
        {
        }
        public static new SiteArgs Empty => new SiteArgs();
    }

    public sealed class SiteState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Site Amazon Resource Name (ARN)
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Description of the Site.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the Global Network to create the site in.
        /// </summary>
        [Input("globalNetworkId")]
        public Input<string>? GlobalNetworkId { get; set; }

        /// <summary>
        /// The site location as documented below.
        /// </summary>
        [Input("location")]
        public Input<Inputs.SiteLocationGetArgs>? Location { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value tags for the Site. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public SiteState()
        {
        }
        public static new SiteState Empty => new SiteState();
    }
}
