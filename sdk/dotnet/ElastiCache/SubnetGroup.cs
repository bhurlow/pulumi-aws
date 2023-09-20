// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElastiCache
{
    /// <summary>
    /// Provides an ElastiCache Subnet Group resource.
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
    ///     var fooVpc = new Aws.Ec2.Vpc("fooVpc", new()
    ///     {
    ///         CidrBlock = "10.0.0.0/16",
    ///         Tags = 
    ///         {
    ///             { "Name", "tf-test" },
    ///         },
    ///     });
    /// 
    ///     var fooSubnet = new Aws.Ec2.Subnet("fooSubnet", new()
    ///     {
    ///         VpcId = fooVpc.Id,
    ///         CidrBlock = "10.0.0.0/24",
    ///         AvailabilityZone = "us-west-2a",
    ///         Tags = 
    ///         {
    ///             { "Name", "tf-test" },
    ///         },
    ///     });
    /// 
    ///     var bar = new Aws.ElastiCache.SubnetGroup("bar", new()
    ///     {
    ///         SubnetIds = new[]
    ///         {
    ///             fooSubnet.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import ElastiCache Subnet Groups using the `name`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:elasticache/subnetGroup:SubnetGroup bar tf-test-cache-subnet
    /// ```
    /// </summary>
    [AwsResourceType("aws:elasticache/subnetGroup:SubnetGroup")]
    public partial class SubnetGroup : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Description for the cache subnet group. Defaults to "Managed by Pulumi".
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Name for the cache subnet group. ElastiCache converts this name to lowercase.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of VPC Subnet IDs for the cache subnet group
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a SubnetGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SubnetGroup(string name, SubnetGroupArgs args, CustomResourceOptions? options = null)
            : base("aws:elasticache/subnetGroup:SubnetGroup", name, args ?? new SubnetGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SubnetGroup(string name, Input<string> id, SubnetGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:elasticache/subnetGroup:SubnetGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SubnetGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SubnetGroup Get(string name, Input<string> id, SubnetGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new SubnetGroup(name, id, state, options);
        }
    }

    public sealed class SubnetGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description for the cache subnet group. Defaults to "Managed by Pulumi".
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name for the cache subnet group. ElastiCache converts this name to lowercase.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// List of VPC Subnet IDs for the cache subnet group
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public SubnetGroupArgs()
        {
            Description = "Managed by Pulumi";
        }
        public static new SubnetGroupArgs Empty => new SubnetGroupArgs();
    }

    public sealed class SubnetGroupState : global::Pulumi.ResourceArgs
    {
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Description for the cache subnet group. Defaults to "Managed by Pulumi".
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name for the cache subnet group. ElastiCache converts this name to lowercase.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// List of VPC Subnet IDs for the cache subnet group
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public SubnetGroupState()
        {
            Description = "Managed by Pulumi";
        }
        public static new SubnetGroupState Empty => new SubnetGroupState();
    }
}
