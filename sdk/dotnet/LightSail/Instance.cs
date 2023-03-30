// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.LightSail
{
    /// <summary>
    /// Provides a Lightsail Instance. Amazon Lightsail is a service to provide easy virtual private servers
    /// with custom software already setup. See [What is Amazon Lightsail?](https://lightsail.aws.amazon.com/ls/docs/getting-started/article/what-is-amazon-lightsail)
    /// for more information.
    /// 
    /// &gt; **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details
    /// 
    /// ## Example Usage
    /// ### Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create a new GitLab Lightsail Instance
    ///     var gitlabTest = new Aws.LightSail.Instance("gitlabTest", new()
    ///     {
    ///         AvailabilityZone = "us-east-1b",
    ///         BlueprintId = "amazon_linux_2",
    ///         BundleId = "nano_1_0",
    ///         KeyPairName = "some_key_name",
    ///         Tags = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Example With User Data
    /// 
    /// Lightsail user data is handled differently than ec2 user data. Lightsail user data only accepts a single lined string. The below example shows installing apache and creating the index page.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var custom = new Aws.LightSail.Instance("custom", new()
    ///     {
    ///         AvailabilityZone = "us-east-1b",
    ///         BlueprintId = "amazon_linux_2",
    ///         BundleId = "nano_1_0",
    ///         UserData = "sudo yum install -y httpd &amp;&amp; sudo systemctl start httpd &amp;&amp; sudo systemctl enable httpd &amp;&amp; echo '&lt;h1&gt;Deployed via Pulumi&lt;/h1&gt;' | sudo tee /var/www/html/index.html",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Enable Auto Snapshots
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = new Aws.LightSail.Instance("test", new()
    ///     {
    ///         AddOn = new Aws.LightSail.Inputs.InstanceAddOnArgs
    ///         {
    ///             SnapshotTime = "06:00",
    ///             Status = "Enabled",
    ///             Type = "AutoSnapshot",
    ///         },
    ///         AvailabilityZone = "us-east-1b",
    ///         BlueprintId = "amazon_linux_2",
    ///         BundleId = "nano_1_0",
    ///         Tags = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ## Availability Zones
    /// 
    /// Lightsail currently supports the following Availability Zones (e.g., `us-east-1a`):
    /// 
    /// - `ap-northeast-1{a,c,d}`
    /// - `ap-northeast-2{a,c}`
    /// - `ap-south-1{a,b}`
    /// - `ap-southeast-1{a,b,c}`
    /// - `ap-southeast-2{a,b,c}`
    /// - `ca-central-1{a,b}`
    /// - `eu-central-1{a,b,c}`
    /// - `eu-west-1{a,b,c}`
    /// - `eu-west-2{a,b,c}`
    /// - `eu-west-3{a,b,c}`
    /// - `us-east-1{a,b,c,d,e,f}`
    /// - `us-east-2{a,b,c}`
    /// - `us-west-2{a,b,c}`
    /// 
    /// ## Bundles
    /// 
    /// Lightsail currently supports the following Bundle IDs (e.g., an instance in `ap-northeast-1` would use `small_2_0`):
    /// 
    /// ### Prefix
    /// 
    /// A Bundle ID starts with one of the below size prefixes:
    /// 
    /// - `nano_`
    /// - `micro_`
    /// - `small_`
    /// - `medium_`
    /// - `large_`
    /// - `xlarge_`
    /// - `2xlarge_`
    /// 
    /// ### Suffix
    /// 
    /// A Bundle ID ends with one of the following suffixes depending on Availability Zone:
    /// 
    /// - ap-northeast-1: `2_0`
    /// - ap-northeast-2: `2_0`
    /// - ap-south-1: `2_1`
    /// - ap-southeast-1: `2_0`
    /// - ap-southeast-2: `2_2`
    /// - ca-central-1: `2_0`
    /// - eu-central-1: `2_0`
    /// - eu-west-1: `2_0`
    /// - eu-west-2: `2_0`
    /// - eu-west-3: `2_0`
    /// - us-east-1: `2_0`
    /// - us-east-2: `2_0`
    /// - us-west-2: `2_0`
    /// 
    /// ## Import
    /// 
    /// Lightsail Instances can be imported using their name, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:lightsail/instance:Instance gitlab_test 'custom_gitlab'
    /// ```
    /// </summary>
    [AwsResourceType("aws:lightsail/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The add on configuration for the instance. Detailed below.
        /// </summary>
        [Output("addOn")]
        public Output<Outputs.InstanceAddOn?> AddOn { get; private set; } = null!;

        /// <summary>
        /// The ARN of the Lightsail instance (matches `id`).
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The Availability Zone in which to create your
        /// instance (see list below)
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// The ID for a virtual private server image. A list of available blueprint IDs can be obtained using the AWS CLI command: `aws lightsail get-blueprints`
        /// </summary>
        [Output("blueprintId")]
        public Output<string> BlueprintId { get; private set; } = null!;

        /// <summary>
        /// The bundle of specification information (see list below)
        /// </summary>
        [Output("bundleId")]
        public Output<string> BundleId { get; private set; } = null!;

        /// <summary>
        /// The number of vCPUs the instance has.
        /// </summary>
        [Output("cpuCount")]
        public Output<int> CpuCount { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the instance was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The IP address type of the Lightsail Instance. Valid Values: `dualstack` | `ipv4`.
        /// </summary>
        [Output("ipAddressType")]
        public Output<string?> IpAddressType { get; private set; } = null!;

        /// <summary>
        /// (**Deprecated**) The first IPv6 address of the Lightsail instance. Use `ipv6_addresses` attribute instead.
        /// </summary>
        [Output("ipv6Address")]
        public Output<string> Ipv6Address { get; private set; } = null!;

        /// <summary>
        /// List of IPv6 addresses for the Lightsail instance.
        /// </summary>
        [Output("ipv6Addresses")]
        public Output<ImmutableArray<string>> Ipv6Addresses { get; private set; } = null!;

        /// <summary>
        /// A Boolean value indicating whether this instance has a static IP assigned to it.
        /// </summary>
        [Output("isStaticIp")]
        public Output<bool> IsStaticIp { get; private set; } = null!;

        /// <summary>
        /// The name of your key pair. Created in the
        /// Lightsail console (cannot use `aws.ec2.KeyPair` at this time)
        /// </summary>
        [Output("keyPairName")]
        public Output<string?> KeyPairName { get; private set; } = null!;

        /// <summary>
        /// The name of the Lightsail Instance. Names be unique within each AWS Region in your Lightsail account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The private IP address of the instance.
        /// </summary>
        [Output("privateIpAddress")]
        public Output<string> PrivateIpAddress { get; private set; } = null!;

        /// <summary>
        /// The public IP address of the instance.
        /// </summary>
        [Output("publicIpAddress")]
        public Output<string> PublicIpAddress { get; private set; } = null!;

        /// <summary>
        /// The amount of RAM in GB on the instance (e.g., 1.0).
        /// </summary>
        [Output("ramSize")]
        public Output<double> RamSize { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Single lined launch script as a string to configure server with additional user data
        /// </summary>
        [Output("userData")]
        public Output<string?> UserData { get; private set; } = null!;

        /// <summary>
        /// The user name for connecting to the instance (e.g., ec2-user).
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("aws:lightsail/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("aws:lightsail/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The add on configuration for the instance. Detailed below.
        /// </summary>
        [Input("addOn")]
        public Input<Inputs.InstanceAddOnArgs>? AddOn { get; set; }

        /// <summary>
        /// The Availability Zone in which to create your
        /// instance (see list below)
        /// </summary>
        [Input("availabilityZone", required: true)]
        public Input<string> AvailabilityZone { get; set; } = null!;

        /// <summary>
        /// The ID for a virtual private server image. A list of available blueprint IDs can be obtained using the AWS CLI command: `aws lightsail get-blueprints`
        /// </summary>
        [Input("blueprintId", required: true)]
        public Input<string> BlueprintId { get; set; } = null!;

        /// <summary>
        /// The bundle of specification information (see list below)
        /// </summary>
        [Input("bundleId", required: true)]
        public Input<string> BundleId { get; set; } = null!;

        /// <summary>
        /// The IP address type of the Lightsail Instance. Valid Values: `dualstack` | `ipv4`.
        /// </summary>
        [Input("ipAddressType")]
        public Input<string>? IpAddressType { get; set; }

        /// <summary>
        /// The name of your key pair. Created in the
        /// Lightsail console (cannot use `aws.ec2.KeyPair` at this time)
        /// </summary>
        [Input("keyPairName")]
        public Input<string>? KeyPairName { get; set; }

        /// <summary>
        /// The name of the Lightsail Instance. Names be unique within each AWS Region in your Lightsail account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Single lined launch script as a string to configure server with additional user data
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The add on configuration for the instance. Detailed below.
        /// </summary>
        [Input("addOn")]
        public Input<Inputs.InstanceAddOnGetArgs>? AddOn { get; set; }

        /// <summary>
        /// The ARN of the Lightsail instance (matches `id`).
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The Availability Zone in which to create your
        /// instance (see list below)
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The ID for a virtual private server image. A list of available blueprint IDs can be obtained using the AWS CLI command: `aws lightsail get-blueprints`
        /// </summary>
        [Input("blueprintId")]
        public Input<string>? BlueprintId { get; set; }

        /// <summary>
        /// The bundle of specification information (see list below)
        /// </summary>
        [Input("bundleId")]
        public Input<string>? BundleId { get; set; }

        /// <summary>
        /// The number of vCPUs the instance has.
        /// </summary>
        [Input("cpuCount")]
        public Input<int>? CpuCount { get; set; }

        /// <summary>
        /// The timestamp when the instance was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The IP address type of the Lightsail Instance. Valid Values: `dualstack` | `ipv4`.
        /// </summary>
        [Input("ipAddressType")]
        public Input<string>? IpAddressType { get; set; }

        /// <summary>
        /// (**Deprecated**) The first IPv6 address of the Lightsail instance. Use `ipv6_addresses` attribute instead.
        /// </summary>
        [Input("ipv6Address")]
        public Input<string>? Ipv6Address { get; set; }

        [Input("ipv6Addresses")]
        private InputList<string>? _ipv6Addresses;

        /// <summary>
        /// List of IPv6 addresses for the Lightsail instance.
        /// </summary>
        public InputList<string> Ipv6Addresses
        {
            get => _ipv6Addresses ?? (_ipv6Addresses = new InputList<string>());
            set => _ipv6Addresses = value;
        }

        /// <summary>
        /// A Boolean value indicating whether this instance has a static IP assigned to it.
        /// </summary>
        [Input("isStaticIp")]
        public Input<bool>? IsStaticIp { get; set; }

        /// <summary>
        /// The name of your key pair. Created in the
        /// Lightsail console (cannot use `aws.ec2.KeyPair` at this time)
        /// </summary>
        [Input("keyPairName")]
        public Input<string>? KeyPairName { get; set; }

        /// <summary>
        /// The name of the Lightsail Instance. Names be unique within each AWS Region in your Lightsail account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The private IP address of the instance.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        /// <summary>
        /// The public IP address of the instance.
        /// </summary>
        [Input("publicIpAddress")]
        public Input<string>? PublicIpAddress { get; set; }

        /// <summary>
        /// The amount of RAM in GB on the instance (e.g., 1.0).
        /// </summary>
        [Input("ramSize")]
        public Input<double>? RamSize { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        /// <summary>
        /// Single lined launch script as a string to configure server with additional user data
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        /// <summary>
        /// The user name for connecting to the instance (e.g., ec2-user).
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}
