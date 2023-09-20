// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cloud9
{
    /// <summary>
    /// Provides a Cloud9 EC2 Development Environment.
    /// 
    /// ## Example Usage
    /// 
    /// Basic usage:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.Cloud9.EnvironmentEC2("example", new()
    ///     {
    ///         InstanceType = "t2.micro",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Get the URL of the Cloud9 environment after creation:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.Cloud9.EnvironmentEC2("example", new()
    ///     {
    ///         InstanceType = "t2.micro",
    ///     });
    /// 
    ///     var cloud9Instance = Aws.Ec2.GetInstance.Invoke(new()
    ///     {
    ///         Filters = new[]
    ///         {
    ///             new Aws.Ec2.Inputs.GetInstanceFilterInputArgs
    ///             {
    ///                 Name = "tag:aws:cloud9:environment",
    ///                 Values = new[]
    ///                 {
    ///                     example.Id,
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["cloud9Url"] = example.Id.Apply(id =&gt; $"https://{@var.Region}.console.aws.amazon.com/cloud9/ide/{id}"),
    ///     };
    /// });
    /// ```
    /// 
    /// Allocate a static IP to the Cloud9 environment:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.Cloud9.EnvironmentEC2("example", new()
    ///     {
    ///         InstanceType = "t2.micro",
    ///     });
    /// 
    ///     var cloud9Instance = Aws.Ec2.GetInstance.Invoke(new()
    ///     {
    ///         Filters = new[]
    ///         {
    ///             new Aws.Ec2.Inputs.GetInstanceFilterInputArgs
    ///             {
    ///                 Name = "tag:aws:cloud9:environment",
    ///                 Values = new[]
    ///                 {
    ///                     example.Id,
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var cloud9Eip = new Aws.Ec2.Eip("cloud9Eip", new()
    ///     {
    ///         Instance = cloud9Instance.Apply(getInstanceResult =&gt; getInstanceResult.Id),
    ///         Domain = "vpc",
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["cloud9PublicIp"] = cloud9Eip.PublicIp,
    ///     };
    /// });
    /// ```
    /// </summary>
    [AwsResourceType("aws:cloud9/environmentEC2:EnvironmentEC2")]
    public partial class EnvironmentEC2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the environment.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The number of minutes until the running instance is shut down after the environment has last been used.
        /// </summary>
        [Output("automaticStopTimeMinutes")]
        public Output<int?> AutomaticStopTimeMinutes { get; private set; } = null!;

        /// <summary>
        /// The connection type used for connecting to an Amazon EC2 environment. Valid values are `CONNECT_SSH` and `CONNECT_SSM`. For more information please refer [AWS documentation for Cloud9](https://docs.aws.amazon.com/cloud9/latest/user-guide/ec2-ssm.html).
        /// </summary>
        [Output("connectionType")]
        public Output<string?> ConnectionType { get; private set; } = null!;

        /// <summary>
        /// The description of the environment.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The identifier for the Amazon Machine Image (AMI) that's used to create the EC2 instance. Valid values are
        /// * `amazonlinux-1-x86_64`
        /// * `amazonlinux-2-x86_64`
        /// * `ubuntu-18.04-x86_64`
        /// * `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-1-x86_64`
        /// * `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-2-x86_64`
        /// * `resolve:ssm:/aws/service/cloud9/amis/ubuntu-18.04-x86_64`
        /// </summary>
        [Output("imageId")]
        public Output<string?> ImageId { get; private set; } = null!;

        /// <summary>
        /// The type of instance to connect to the environment, e.g., `t2.micro`.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The name of the environment.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
        /// </summary>
        [Output("ownerArn")]
        public Output<string> OwnerArn { get; private set; } = null!;

        /// <summary>
        /// The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
        /// </summary>
        [Output("subnetId")]
        public Output<string?> SubnetId { get; private set; } = null!;

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
        /// The type of the environment (e.g., `ssh` or `ec2`)
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a EnvironmentEC2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EnvironmentEC2(string name, EnvironmentEC2Args args, CustomResourceOptions? options = null)
            : base("aws:cloud9/environmentEC2:EnvironmentEC2", name, args ?? new EnvironmentEC2Args(), MakeResourceOptions(options, ""))
        {
        }

        private EnvironmentEC2(string name, Input<string> id, EnvironmentEC2State? state = null, CustomResourceOptions? options = null)
            : base("aws:cloud9/environmentEC2:EnvironmentEC2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EnvironmentEC2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EnvironmentEC2 Get(string name, Input<string> id, EnvironmentEC2State? state = null, CustomResourceOptions? options = null)
        {
            return new EnvironmentEC2(name, id, state, options);
        }
    }

    public sealed class EnvironmentEC2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of minutes until the running instance is shut down after the environment has last been used.
        /// </summary>
        [Input("automaticStopTimeMinutes")]
        public Input<int>? AutomaticStopTimeMinutes { get; set; }

        /// <summary>
        /// The connection type used for connecting to an Amazon EC2 environment. Valid values are `CONNECT_SSH` and `CONNECT_SSM`. For more information please refer [AWS documentation for Cloud9](https://docs.aws.amazon.com/cloud9/latest/user-guide/ec2-ssm.html).
        /// </summary>
        [Input("connectionType")]
        public Input<string>? ConnectionType { get; set; }

        /// <summary>
        /// The description of the environment.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The identifier for the Amazon Machine Image (AMI) that's used to create the EC2 instance. Valid values are
        /// * `amazonlinux-1-x86_64`
        /// * `amazonlinux-2-x86_64`
        /// * `ubuntu-18.04-x86_64`
        /// * `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-1-x86_64`
        /// * `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-2-x86_64`
        /// * `resolve:ssm:/aws/service/cloud9/amis/ubuntu-18.04-x86_64`
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// The type of instance to connect to the environment, e.g., `t2.micro`.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// The name of the environment.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
        /// </summary>
        [Input("ownerArn")]
        public Input<string>? OwnerArn { get; set; }

        /// <summary>
        /// The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

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

        public EnvironmentEC2Args()
        {
        }
        public static new EnvironmentEC2Args Empty => new EnvironmentEC2Args();
    }

    public sealed class EnvironmentEC2State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the environment.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The number of minutes until the running instance is shut down after the environment has last been used.
        /// </summary>
        [Input("automaticStopTimeMinutes")]
        public Input<int>? AutomaticStopTimeMinutes { get; set; }

        /// <summary>
        /// The connection type used for connecting to an Amazon EC2 environment. Valid values are `CONNECT_SSH` and `CONNECT_SSM`. For more information please refer [AWS documentation for Cloud9](https://docs.aws.amazon.com/cloud9/latest/user-guide/ec2-ssm.html).
        /// </summary>
        [Input("connectionType")]
        public Input<string>? ConnectionType { get; set; }

        /// <summary>
        /// The description of the environment.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The identifier for the Amazon Machine Image (AMI) that's used to create the EC2 instance. Valid values are
        /// * `amazonlinux-1-x86_64`
        /// * `amazonlinux-2-x86_64`
        /// * `ubuntu-18.04-x86_64`
        /// * `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-1-x86_64`
        /// * `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-2-x86_64`
        /// * `resolve:ssm:/aws/service/cloud9/amis/ubuntu-18.04-x86_64`
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// The type of instance to connect to the environment, e.g., `t2.micro`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The name of the environment.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
        /// </summary>
        [Input("ownerArn")]
        public Input<string>? OwnerArn { get; set; }

        /// <summary>
        /// The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

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

        /// <summary>
        /// The type of the environment (e.g., `ssh` or `ec2`)
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public EnvironmentEC2State()
        {
        }
        public static new EnvironmentEC2State Empty => new EnvironmentEC2State();
    }
}
