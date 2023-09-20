// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.RedshiftServerless
{
    /// <summary>
    /// Creates a new Amazon Redshift Serverless Workgroup.
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
    ///     var example = new Aws.RedshiftServerless.Workgroup("example", new()
    ///     {
    ///         NamespaceName = "concurrency-scaling",
    ///         WorkgroupName = "concurrency-scaling",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import Redshift Serverless Workgroups using the `workgroup_name`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:redshiftserverless/workgroup:Workgroup example example
    /// ```
    /// </summary>
    [AwsResourceType("aws:redshiftserverless/workgroup:Workgroup")]
    public partial class Workgroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the Redshift Serverless Workgroup.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The base data warehouse capacity of the workgroup in Redshift Processing Units (RPUs).
        /// </summary>
        [Output("baseCapacity")]
        public Output<int> BaseCapacity { get; private set; } = null!;

        /// <summary>
        /// An array of parameters to set for more control over a serverless database. See `Config Parameter` below.
        /// </summary>
        [Output("configParameters")]
        public Output<ImmutableArray<Outputs.WorkgroupConfigParameter>> ConfigParameters { get; private set; } = null!;

        /// <summary>
        /// The endpoint that is created from the workgroup. See `Endpoint` below.
        /// </summary>
        [Output("endpoints")]
        public Output<ImmutableArray<Outputs.WorkgroupEndpoint>> Endpoints { get; private set; } = null!;

        /// <summary>
        /// The value that specifies whether to turn on enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC instead of over the internet.
        /// </summary>
        [Output("enhancedVpcRouting")]
        public Output<bool?> EnhancedVpcRouting { get; private set; } = null!;

        /// <summary>
        /// The name of the namespace.
        /// </summary>
        [Output("namespaceName")]
        public Output<string> NamespaceName { get; private set; } = null!;

        /// <summary>
        /// A value that specifies whether the workgroup can be accessed from a public network.
        /// </summary>
        [Output("publiclyAccessible")]
        public Output<bool?> PubliclyAccessible { get; private set; } = null!;

        /// <summary>
        /// An array of security group IDs to associate with the workgroup.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// An array of VPC subnet IDs to associate with the workgroup. When set, must contain at least three subnets spanning three Availability Zones. A minimum number of IP addresses is required and scales with the Base Capacity. For more information, see the following [AWS document](https://docs.aws.amazon.com/redshift/latest/mgmt/serverless-known-issues.html).
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

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
        /// The Redshift Workgroup ID.
        /// </summary>
        [Output("workgroupId")]
        public Output<string> WorkgroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the workgroup.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("workgroupName")]
        public Output<string> WorkgroupName { get; private set; } = null!;


        /// <summary>
        /// Create a Workgroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Workgroup(string name, WorkgroupArgs args, CustomResourceOptions? options = null)
            : base("aws:redshiftserverless/workgroup:Workgroup", name, args ?? new WorkgroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Workgroup(string name, Input<string> id, WorkgroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:redshiftserverless/workgroup:Workgroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Workgroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Workgroup Get(string name, Input<string> id, WorkgroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Workgroup(name, id, state, options);
        }
    }

    public sealed class WorkgroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The base data warehouse capacity of the workgroup in Redshift Processing Units (RPUs).
        /// </summary>
        [Input("baseCapacity")]
        public Input<int>? BaseCapacity { get; set; }

        [Input("configParameters")]
        private InputList<Inputs.WorkgroupConfigParameterArgs>? _configParameters;

        /// <summary>
        /// An array of parameters to set for more control over a serverless database. See `Config Parameter` below.
        /// </summary>
        public InputList<Inputs.WorkgroupConfigParameterArgs> ConfigParameters
        {
            get => _configParameters ?? (_configParameters = new InputList<Inputs.WorkgroupConfigParameterArgs>());
            set => _configParameters = value;
        }

        /// <summary>
        /// The value that specifies whether to turn on enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC instead of over the internet.
        /// </summary>
        [Input("enhancedVpcRouting")]
        public Input<bool>? EnhancedVpcRouting { get; set; }

        /// <summary>
        /// The name of the namespace.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// A value that specifies whether the workgroup can be accessed from a public network.
        /// </summary>
        [Input("publiclyAccessible")]
        public Input<bool>? PubliclyAccessible { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// An array of security group IDs to associate with the workgroup.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// An array of VPC subnet IDs to associate with the workgroup. When set, must contain at least three subnets spanning three Availability Zones. A minimum number of IP addresses is required and scales with the Base Capacity. For more information, see the following [AWS document](https://docs.aws.amazon.com/redshift/latest/mgmt/serverless-known-issues.html).
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

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

        /// <summary>
        /// The name of the workgroup.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("workgroupName", required: true)]
        public Input<string> WorkgroupName { get; set; } = null!;

        public WorkgroupArgs()
        {
        }
        public static new WorkgroupArgs Empty => new WorkgroupArgs();
    }

    public sealed class WorkgroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the Redshift Serverless Workgroup.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The base data warehouse capacity of the workgroup in Redshift Processing Units (RPUs).
        /// </summary>
        [Input("baseCapacity")]
        public Input<int>? BaseCapacity { get; set; }

        [Input("configParameters")]
        private InputList<Inputs.WorkgroupConfigParameterGetArgs>? _configParameters;

        /// <summary>
        /// An array of parameters to set for more control over a serverless database. See `Config Parameter` below.
        /// </summary>
        public InputList<Inputs.WorkgroupConfigParameterGetArgs> ConfigParameters
        {
            get => _configParameters ?? (_configParameters = new InputList<Inputs.WorkgroupConfigParameterGetArgs>());
            set => _configParameters = value;
        }

        [Input("endpoints")]
        private InputList<Inputs.WorkgroupEndpointGetArgs>? _endpoints;

        /// <summary>
        /// The endpoint that is created from the workgroup. See `Endpoint` below.
        /// </summary>
        public InputList<Inputs.WorkgroupEndpointGetArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.WorkgroupEndpointGetArgs>());
            set => _endpoints = value;
        }

        /// <summary>
        /// The value that specifies whether to turn on enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC instead of over the internet.
        /// </summary>
        [Input("enhancedVpcRouting")]
        public Input<bool>? EnhancedVpcRouting { get; set; }

        /// <summary>
        /// The name of the namespace.
        /// </summary>
        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        /// <summary>
        /// A value that specifies whether the workgroup can be accessed from a public network.
        /// </summary>
        [Input("publiclyAccessible")]
        public Input<bool>? PubliclyAccessible { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// An array of security group IDs to associate with the workgroup.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// An array of VPC subnet IDs to associate with the workgroup. When set, must contain at least three subnets spanning three Availability Zones. A minimum number of IP addresses is required and scales with the Base Capacity. For more information, see the following [AWS document](https://docs.aws.amazon.com/redshift/latest/mgmt/serverless-known-issues.html).
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

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

        /// <summary>
        /// The Redshift Workgroup ID.
        /// </summary>
        [Input("workgroupId")]
        public Input<string>? WorkgroupId { get; set; }

        /// <summary>
        /// The name of the workgroup.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("workgroupName")]
        public Input<string>? WorkgroupName { get; set; }

        public WorkgroupState()
        {
        }
        public static new WorkgroupState Empty => new WorkgroupState();
    }
}
