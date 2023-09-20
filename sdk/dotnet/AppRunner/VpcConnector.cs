// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppRunner
{
    /// <summary>
    /// Manages an App Runner VPC Connector.
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
    ///     var connector = new Aws.AppRunner.VpcConnector("connector", new()
    ///     {
    ///         SecurityGroups = new[]
    ///         {
    ///             "sg1",
    ///             "sg2",
    ///         },
    ///         Subnets = new[]
    ///         {
    ///             "subnet1",
    ///             "subnet2",
    ///         },
    ///         VpcConnectorName = "name",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import App Runner vpc connector using the `arn`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:apprunner/vpcConnector:VpcConnector example arn:aws:apprunner:us-east-1:1234567890:vpcconnector/example/1/0a03292a89764e5882c41d8f991c82fe
    /// ```
    /// </summary>
    [AwsResourceType("aws:apprunner/vpcConnector:VpcConnector")]
    public partial class VpcConnector : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN of VPC connector.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// List of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// Current state of the VPC connector. If the status of a connector revision is INACTIVE, it was deleted and can't be used. Inactive connector revisions are permanently removed some time after they are deleted.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// List of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
        /// </summary>
        [Output("subnets")]
        public Output<ImmutableArray<string>> Subnets { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Name for the VPC connector.
        /// </summary>
        [Output("vpcConnectorName")]
        public Output<string> VpcConnectorName { get; private set; } = null!;

        /// <summary>
        /// The revision of VPC connector. It's unique among all the active connectors ("Status": "ACTIVE") that share the same Name.
        /// </summary>
        [Output("vpcConnectorRevision")]
        public Output<int> VpcConnectorRevision { get; private set; } = null!;


        /// <summary>
        /// Create a VpcConnector resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcConnector(string name, VpcConnectorArgs args, CustomResourceOptions? options = null)
            : base("aws:apprunner/vpcConnector:VpcConnector", name, args ?? new VpcConnectorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcConnector(string name, Input<string> id, VpcConnectorState? state = null, CustomResourceOptions? options = null)
            : base("aws:apprunner/vpcConnector:VpcConnector", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcConnector resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcConnector Get(string name, Input<string> id, VpcConnectorState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcConnector(name, id, state, options);
        }
    }

    public sealed class VpcConnectorArgs : global::Pulumi.ResourceArgs
    {
        [Input("securityGroups", required: true)]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// List of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("subnets", required: true)]
        private InputList<string>? _subnets;

        /// <summary>
        /// List of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
        /// </summary>
        public InputList<string> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<string>());
            set => _subnets = value;
        }

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
        /// Name for the VPC connector.
        /// </summary>
        [Input("vpcConnectorName", required: true)]
        public Input<string> VpcConnectorName { get; set; } = null!;

        public VpcConnectorArgs()
        {
        }
        public static new VpcConnectorArgs Empty => new VpcConnectorArgs();
    }

    public sealed class VpcConnectorState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of VPC connector.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// List of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// Current state of the VPC connector. If the status of a connector revision is INACTIVE, it was deleted and can't be used. Inactive connector revisions are permanently removed some time after they are deleted.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("subnets")]
        private InputList<string>? _subnets;

        /// <summary>
        /// List of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
        /// </summary>
        public InputList<string> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<string>());
            set => _subnets = value;
        }

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
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
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
        /// Name for the VPC connector.
        /// </summary>
        [Input("vpcConnectorName")]
        public Input<string>? VpcConnectorName { get; set; }

        /// <summary>
        /// The revision of VPC connector. It's unique among all the active connectors ("Status": "ACTIVE") that share the same Name.
        /// </summary>
        [Input("vpcConnectorRevision")]
        public Input<int>? VpcConnectorRevision { get; set; }

        public VpcConnectorState()
        {
        }
        public static new VpcConnectorState Empty => new VpcConnectorState();
    }
}
