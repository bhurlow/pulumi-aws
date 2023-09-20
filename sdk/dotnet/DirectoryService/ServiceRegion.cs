// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DirectoryService
{
    /// <summary>
    /// Manages a replicated Region and directory for Multi-Region replication.
    /// Multi-Region replication is only supported for the Enterprise Edition of AWS Managed Microsoft AD.
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import Replicated Regions using directory ID,Region name. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:directoryservice/serviceRegion:ServiceRegion example d-9267651497,us-east-2
    /// ```
    /// </summary>
    [AwsResourceType("aws:directoryservice/serviceRegion:ServiceRegion")]
    public partial class ServiceRegion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The number of domain controllers desired in the replicated directory. Minimum value of `2`.
        /// </summary>
        [Output("desiredNumberOfDomainControllers")]
        public Output<int> DesiredNumberOfDomainControllers { get; private set; } = null!;

        /// <summary>
        /// The identifier of the directory to which you want to add Region replication.
        /// </summary>
        [Output("directoryId")]
        public Output<string> DirectoryId { get; private set; } = null!;

        /// <summary>
        /// The name of the Region where you want to add domain controllers for replication.
        /// </summary>
        [Output("regionName")]
        public Output<string> RegionName { get; private set; } = null!;

        /// <summary>
        /// Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// VPC information in the replicated Region. Detailed below.
        /// </summary>
        [Output("vpcSettings")]
        public Output<Outputs.ServiceRegionVpcSettings> VpcSettings { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceRegion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceRegion(string name, ServiceRegionArgs args, CustomResourceOptions? options = null)
            : base("aws:directoryservice/serviceRegion:ServiceRegion", name, args ?? new ServiceRegionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceRegion(string name, Input<string> id, ServiceRegionState? state = null, CustomResourceOptions? options = null)
            : base("aws:directoryservice/serviceRegion:ServiceRegion", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceRegion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceRegion Get(string name, Input<string> id, ServiceRegionState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceRegion(name, id, state, options);
        }
    }

    public sealed class ServiceRegionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of domain controllers desired in the replicated directory. Minimum value of `2`.
        /// </summary>
        [Input("desiredNumberOfDomainControllers")]
        public Input<int>? DesiredNumberOfDomainControllers { get; set; }

        /// <summary>
        /// The identifier of the directory to which you want to add Region replication.
        /// </summary>
        [Input("directoryId", required: true)]
        public Input<string> DirectoryId { get; set; } = null!;

        /// <summary>
        /// The name of the Region where you want to add domain controllers for replication.
        /// </summary>
        [Input("regionName", required: true)]
        public Input<string> RegionName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// VPC information in the replicated Region. Detailed below.
        /// </summary>
        [Input("vpcSettings", required: true)]
        public Input<Inputs.ServiceRegionVpcSettingsArgs> VpcSettings { get; set; } = null!;

        public ServiceRegionArgs()
        {
        }
        public static new ServiceRegionArgs Empty => new ServiceRegionArgs();
    }

    public sealed class ServiceRegionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of domain controllers desired in the replicated directory. Minimum value of `2`.
        /// </summary>
        [Input("desiredNumberOfDomainControllers")]
        public Input<int>? DesiredNumberOfDomainControllers { get; set; }

        /// <summary>
        /// The identifier of the directory to which you want to add Region replication.
        /// </summary>
        [Input("directoryId")]
        public Input<string>? DirectoryId { get; set; }

        /// <summary>
        /// The name of the Region where you want to add domain controllers for replication.
        /// </summary>
        [Input("regionName")]
        public Input<string>? RegionName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        /// VPC information in the replicated Region. Detailed below.
        /// </summary>
        [Input("vpcSettings")]
        public Input<Inputs.ServiceRegionVpcSettingsGetArgs>? VpcSettings { get; set; }

        public ServiceRegionState()
        {
        }
        public static new ServiceRegionState Empty => new ServiceRegionState();
    }
}
