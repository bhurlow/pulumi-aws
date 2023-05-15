// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DeviceFarm
{
    /// <summary>
    /// Provides a resource to manage AWS Device Farm Instance Profiles.
    /// ∂
    /// &gt; **NOTE:** AWS currently has limited regional support for Device Farm (e.g., `us-west-2`). See [AWS Device Farm endpoints and quotas](https://docs.aws.amazon.com/general/latest/gr/devicefarm.html) for information on supported regions.
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
    ///     var example = new Aws.DeviceFarm.InstanceProfile("example");
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// DeviceFarm Instance Profiles can be imported by their arn
    /// 
    /// ```sh
    ///  $ pulumi import aws:devicefarm/instanceProfile:InstanceProfile example arn:aws:devicefarm:us-west-2:123456789012:instanceprofile:4fa784c7-ccb4-4dbf-ba4f-02198320daa1
    /// ```
    /// </summary>
    [AwsResourceType("aws:devicefarm/instanceProfile:InstanceProfile")]
    public partial class InstanceProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name of this instance profile.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The description of the instance profile.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// An array of strings that specifies the list of app packages that should not be cleaned up from the device after a test run.
        /// </summary>
        [Output("excludeAppPackagesFromCleanups")]
        public Output<ImmutableArray<string>> ExcludeAppPackagesFromCleanups { get; private set; } = null!;

        /// <summary>
        /// The name for the instance profile.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// When set to `true`, Device Farm removes app packages after a test run. The default value is `false` for private devices.
        /// </summary>
        [Output("packageCleanup")]
        public Output<bool?> PackageCleanup { get; private set; } = null!;

        /// <summary>
        /// When set to `true`, Device Farm reboots the instance after a test run. The default value is `true`.
        /// </summary>
        [Output("rebootAfterUse")]
        public Output<bool?> RebootAfterUse { get; private set; } = null!;

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
        /// Create a InstanceProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceProfile(string name, InstanceProfileArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:devicefarm/instanceProfile:InstanceProfile", name, args ?? new InstanceProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceProfile(string name, Input<string> id, InstanceProfileState? state = null, CustomResourceOptions? options = null)
            : base("aws:devicefarm/instanceProfile:InstanceProfile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceProfile Get(string name, Input<string> id, InstanceProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceProfile(name, id, state, options);
        }
    }

    public sealed class InstanceProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the instance profile.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("excludeAppPackagesFromCleanups")]
        private InputList<string>? _excludeAppPackagesFromCleanups;

        /// <summary>
        /// An array of strings that specifies the list of app packages that should not be cleaned up from the device after a test run.
        /// </summary>
        public InputList<string> ExcludeAppPackagesFromCleanups
        {
            get => _excludeAppPackagesFromCleanups ?? (_excludeAppPackagesFromCleanups = new InputList<string>());
            set => _excludeAppPackagesFromCleanups = value;
        }

        /// <summary>
        /// The name for the instance profile.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// When set to `true`, Device Farm removes app packages after a test run. The default value is `false` for private devices.
        /// </summary>
        [Input("packageCleanup")]
        public Input<bool>? PackageCleanup { get; set; }

        /// <summary>
        /// When set to `true`, Device Farm reboots the instance after a test run. The default value is `true`.
        /// </summary>
        [Input("rebootAfterUse")]
        public Input<bool>? RebootAfterUse { get; set; }

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

        public InstanceProfileArgs()
        {
        }
        public static new InstanceProfileArgs Empty => new InstanceProfileArgs();
    }

    public sealed class InstanceProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name of this instance profile.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The description of the instance profile.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("excludeAppPackagesFromCleanups")]
        private InputList<string>? _excludeAppPackagesFromCleanups;

        /// <summary>
        /// An array of strings that specifies the list of app packages that should not be cleaned up from the device after a test run.
        /// </summary>
        public InputList<string> ExcludeAppPackagesFromCleanups
        {
            get => _excludeAppPackagesFromCleanups ?? (_excludeAppPackagesFromCleanups = new InputList<string>());
            set => _excludeAppPackagesFromCleanups = value;
        }

        /// <summary>
        /// The name for the instance profile.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// When set to `true`, Device Farm removes app packages after a test run. The default value is `false` for private devices.
        /// </summary>
        [Input("packageCleanup")]
        public Input<bool>? PackageCleanup { get; set; }

        /// <summary>
        /// When set to `true`, Device Farm reboots the instance after a test run. The default value is `true`.
        /// </summary>
        [Input("rebootAfterUse")]
        public Input<bool>? RebootAfterUse { get; set; }

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

        public InstanceProfileState()
        {
        }
        public static new InstanceProfileState Empty => new InstanceProfileState();
    }
}
