// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Msk
{
    /// <summary>
    /// Manages an Amazon MSK Serverless cluster.
    /// 
    /// &gt; **Note:** To manage a _provisioned_ Amazon MSK cluster, use the `aws.msk.Cluster` resource.
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import MSK serverless clusters using the cluster `arn`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:msk/serverlessCluster:ServerlessCluster example arn:aws:kafka:us-west-2:123456789012:cluster/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3
    /// ```
    /// </summary>
    [AwsResourceType("aws:msk/serverlessCluster:ServerlessCluster")]
    public partial class ServerlessCluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the serverless cluster.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Specifies client authentication information for the serverless cluster. See below.
        /// </summary>
        [Output("clientAuthentication")]
        public Output<Outputs.ServerlessClusterClientAuthentication> ClientAuthentication { get; private set; } = null!;

        /// <summary>
        /// The name of the serverless cluster.
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

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
        /// VPC configuration information. See below.
        /// </summary>
        [Output("vpcConfigs")]
        public Output<ImmutableArray<Outputs.ServerlessClusterVpcConfig>> VpcConfigs { get; private set; } = null!;


        /// <summary>
        /// Create a ServerlessCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerlessCluster(string name, ServerlessClusterArgs args, CustomResourceOptions? options = null)
            : base("aws:msk/serverlessCluster:ServerlessCluster", name, args ?? new ServerlessClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerlessCluster(string name, Input<string> id, ServerlessClusterState? state = null, CustomResourceOptions? options = null)
            : base("aws:msk/serverlessCluster:ServerlessCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServerlessCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerlessCluster Get(string name, Input<string> id, ServerlessClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new ServerlessCluster(name, id, state, options);
        }
    }

    public sealed class ServerlessClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies client authentication information for the serverless cluster. See below.
        /// </summary>
        [Input("clientAuthentication", required: true)]
        public Input<Inputs.ServerlessClusterClientAuthenticationArgs> ClientAuthentication { get; set; } = null!;

        /// <summary>
        /// The name of the serverless cluster.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

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

        [Input("vpcConfigs", required: true)]
        private InputList<Inputs.ServerlessClusterVpcConfigArgs>? _vpcConfigs;

        /// <summary>
        /// VPC configuration information. See below.
        /// </summary>
        public InputList<Inputs.ServerlessClusterVpcConfigArgs> VpcConfigs
        {
            get => _vpcConfigs ?? (_vpcConfigs = new InputList<Inputs.ServerlessClusterVpcConfigArgs>());
            set => _vpcConfigs = value;
        }

        public ServerlessClusterArgs()
        {
        }
        public static new ServerlessClusterArgs Empty => new ServerlessClusterArgs();
    }

    public sealed class ServerlessClusterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the serverless cluster.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Specifies client authentication information for the serverless cluster. See below.
        /// </summary>
        [Input("clientAuthentication")]
        public Input<Inputs.ServerlessClusterClientAuthenticationGetArgs>? ClientAuthentication { get; set; }

        /// <summary>
        /// The name of the serverless cluster.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

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

        [Input("vpcConfigs")]
        private InputList<Inputs.ServerlessClusterVpcConfigGetArgs>? _vpcConfigs;

        /// <summary>
        /// VPC configuration information. See below.
        /// </summary>
        public InputList<Inputs.ServerlessClusterVpcConfigGetArgs> VpcConfigs
        {
            get => _vpcConfigs ?? (_vpcConfigs = new InputList<Inputs.ServerlessClusterVpcConfigGetArgs>());
            set => _vpcConfigs = value;
        }

        public ServerlessClusterState()
        {
        }
        public static new ServerlessClusterState Empty => new ServerlessClusterState();
    }
}
