// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.LicenseManager
{
    /// <summary>
    /// Provides a License Manager license configuration resource.
    /// 
    /// &gt; **Note:** Removing the `license_count` attribute is not supported by the License Manager API - recreate the resource instead.
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
    ///     var example = new Aws.LicenseManager.LicenseConfiguration("example", new()
    ///     {
    ///         Description = "Example",
    ///         LicenseCount = 10,
    ///         LicenseCountHardLimit = true,
    ///         LicenseCountingType = "Socket",
    ///         LicenseRules = new[]
    ///         {
    ///             "#minimumSockets=2",
    ///         },
    ///         Tags = 
    ///         {
    ///             { "foo", "barr" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ## Rules
    /// 
    /// License rules should be in the format of `#RuleType=RuleValue`. Supported rule types:
    /// 
    /// * `minimumVcpus` - Resource must have minimum vCPU count in order to use the license. Default: 1
    /// * `maximumVcpus` - Resource must have maximum vCPU count in order to use the license. Default: unbounded, limit: 10000
    /// * `minimumCores` - Resource must have minimum core count in order to use the license. Default: 1
    /// * `maximumCores` - Resource must have maximum core count in order to use the license. Default: unbounded, limit: 10000
    /// * `minimumSockets` - Resource must have minimum socket count in order to use the license. Default: 1
    /// * `maximumSockets` - Resource must have maximum socket count in order to use the license. Default: unbounded, limit: 10000
    /// * `allowedTenancy` - Defines where the license can be used. If set, restricts license usage to selected tenancies. Specify a comma delimited list of `EC2-Default`, `EC2-DedicatedHost`, `EC2-DedicatedInstance`
    /// 
    /// ## Import
    /// 
    /// License configurations can be imported using the `id`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:licensemanager/licenseConfiguration:LicenseConfiguration example arn:aws:license-manager:eu-west-1:123456789012:license-configuration:lic-0123456789abcdef0123456789abcdef
    /// ```
    /// </summary>
    [AwsResourceType("aws:licensemanager/licenseConfiguration:LicenseConfiguration")]
    public partial class LicenseConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The license configuration ARN.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Description of the license configuration.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Number of licenses managed by the license configuration.
        /// </summary>
        [Output("licenseCount")]
        public Output<int?> LicenseCount { get; private set; } = null!;

        /// <summary>
        /// Sets the number of available licenses as a hard limit.
        /// </summary>
        [Output("licenseCountHardLimit")]
        public Output<bool?> LicenseCountHardLimit { get; private set; } = null!;

        /// <summary>
        /// Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
        /// </summary>
        [Output("licenseCountingType")]
        public Output<string> LicenseCountingType { get; private set; } = null!;

        /// <summary>
        /// Array of configured License Manager rules.
        /// </summary>
        [Output("licenseRules")]
        public Output<ImmutableArray<string>> LicenseRules { get; private set; } = null!;

        /// <summary>
        /// Name of the license configuration.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Account ID of the owner of the license configuration.
        /// </summary>
        [Output("ownerAccountId")]
        public Output<string> OwnerAccountId { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a LicenseConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LicenseConfiguration(string name, LicenseConfigurationArgs args, CustomResourceOptions? options = null)
            : base("aws:licensemanager/licenseConfiguration:LicenseConfiguration", name, args ?? new LicenseConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LicenseConfiguration(string name, Input<string> id, LicenseConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("aws:licensemanager/licenseConfiguration:LicenseConfiguration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LicenseConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LicenseConfiguration Get(string name, Input<string> id, LicenseConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new LicenseConfiguration(name, id, state, options);
        }
    }

    public sealed class LicenseConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the license configuration.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Number of licenses managed by the license configuration.
        /// </summary>
        [Input("licenseCount")]
        public Input<int>? LicenseCount { get; set; }

        /// <summary>
        /// Sets the number of available licenses as a hard limit.
        /// </summary>
        [Input("licenseCountHardLimit")]
        public Input<bool>? LicenseCountHardLimit { get; set; }

        /// <summary>
        /// Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
        /// </summary>
        [Input("licenseCountingType", required: true)]
        public Input<string> LicenseCountingType { get; set; } = null!;

        [Input("licenseRules")]
        private InputList<string>? _licenseRules;

        /// <summary>
        /// Array of configured License Manager rules.
        /// </summary>
        public InputList<string> LicenseRules
        {
            get => _licenseRules ?? (_licenseRules = new InputList<string>());
            set => _licenseRules = value;
        }

        /// <summary>
        /// Name of the license configuration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public LicenseConfigurationArgs()
        {
        }
        public static new LicenseConfigurationArgs Empty => new LicenseConfigurationArgs();
    }

    public sealed class LicenseConfigurationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The license configuration ARN.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Description of the license configuration.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Number of licenses managed by the license configuration.
        /// </summary>
        [Input("licenseCount")]
        public Input<int>? LicenseCount { get; set; }

        /// <summary>
        /// Sets the number of available licenses as a hard limit.
        /// </summary>
        [Input("licenseCountHardLimit")]
        public Input<bool>? LicenseCountHardLimit { get; set; }

        /// <summary>
        /// Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
        /// </summary>
        [Input("licenseCountingType")]
        public Input<string>? LicenseCountingType { get; set; }

        [Input("licenseRules")]
        private InputList<string>? _licenseRules;

        /// <summary>
        /// Array of configured License Manager rules.
        /// </summary>
        public InputList<string> LicenseRules
        {
            get => _licenseRules ?? (_licenseRules = new InputList<string>());
            set => _licenseRules = value;
        }

        /// <summary>
        /// Name of the license configuration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Account ID of the owner of the license configuration.
        /// </summary>
        [Input("ownerAccountId")]
        public Input<string>? OwnerAccountId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public LicenseConfigurationState()
        {
        }
        public static new LicenseConfigurationState Empty => new LicenseConfigurationState();
    }
}
