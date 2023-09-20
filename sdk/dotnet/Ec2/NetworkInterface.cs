// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// Provides an Elastic network interface (ENI) resource.
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
    ///     var test = new Aws.Ec2.NetworkInterface("test", new()
    ///     {
    ///         SubnetId = aws_subnet.Public_a.Id,
    ///         PrivateIps = new[]
    ///         {
    ///             "10.0.0.50",
    ///         },
    ///         SecurityGroups = new[]
    ///         {
    ///             aws_security_group.Web.Id,
    ///         },
    ///         Attachments = new[]
    ///         {
    ///             new Aws.Ec2.Inputs.NetworkInterfaceAttachmentArgs
    ///             {
    ///                 Instance = aws_instance.Test.Id,
    ///                 DeviceIndex = 1,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Example of Managing Multiple IPs on a Network Interface
    /// 
    /// By default, private IPs are managed through the `private_ips` and `private_ips_count` arguments which manage IPs as a set of IPs that are configured without regard to order. For a new network interface, the same primary IP address is consistently selected from a given set of addresses, regardless of the order provided. However, modifications of the set of addresses of an existing interface will not alter the current primary IP address unless it has been removed from the set.
    /// 
    /// In order to manage the private IPs as a sequentially ordered list, configure `private_ip_list_enabled` to `true` and use `private_ip_list` to manage the IPs. This will disable the `private_ips` and `private_ips_count` settings, which must be removed from the config file but are still exported. Note that changing the first address of `private_ip_list`, which is the primary, always requires a new interface.
    /// 
    /// If you are managing a specific set or list of IPs, instead of just using `private_ips_count`, this is a potential workflow for also leveraging `private_ips_count` to have AWS automatically assign additional IP addresses:
    /// 
    /// 1. Comment out `private_ips`, `private_ip_list`, `private_ip_list_enabled` in your configuration
    /// 2. Set the desired `private_ips_count` (count of the number of secondaries, the primary is not included)
    /// 3. Apply to assign the extra IPs
    /// 4. Remove `private_ips_count` and restore your settings from the first step
    /// 5. Add the new IPs to your current settings
    /// 6. Apply again to update the stored state
    /// 
    /// This process can also be used to remove IP addresses in addition to the option of manually removing them. Adding IP addresses in a manually is more difficult because it requires knowledge of which addresses are available.
    /// 
    /// ## Import
    /// 
    /// In TODO v1.5.0 and later, use an `import` block to import Network Interfaces using the `id`. For exampleterraform import {
    /// 
    ///  to = aws_network_interface.test
    /// 
    ///  id = "eni-e5aa89a3" } Using `TODO import`, import Network Interfaces using the `id`. For exampleconsole % TODO import aws_network_interface.test eni-e5aa89a3
    /// </summary>
    [AwsResourceType("aws:ec2/networkInterface:NetworkInterface")]
    public partial class NetworkInterface : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN of the network interface.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Configuration block to define the attachment of the ENI. See Attachment below for more details!
        /// </summary>
        [Output("attachments")]
        public Output<ImmutableArray<Outputs.NetworkInterfaceAttachment>> Attachments { get; private set; } = null!;

        /// <summary>
        /// Description for the network interface.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Type of network interface to create. Set to `efa` for Elastic Fabric Adapter. Changing `interface_type` will cause the resource to be destroyed and re-created.
        /// </summary>
        [Output("interfaceType")]
        public Output<string> InterfaceType { get; private set; } = null!;

        /// <summary>
        /// Number of IPv4 prefixes that AWS automatically assigns to the network interface.
        /// </summary>
        [Output("ipv4PrefixCount")]
        public Output<int> Ipv4PrefixCount { get; private set; } = null!;

        /// <summary>
        /// One or more IPv4 prefixes assigned to the network interface.
        /// </summary>
        [Output("ipv4Prefixes")]
        public Output<ImmutableArray<string>> Ipv4Prefixes { get; private set; } = null!;

        /// <summary>
        /// Number of IPv6 addresses to assign to a network interface. You can't use this option if specifying specific `ipv6_addresses`. If your subnet has the AssignIpv6AddressOnCreation attribute set to `true`, you can specify `0` to override this setting.
        /// </summary>
        [Output("ipv6AddressCount")]
        public Output<int> Ipv6AddressCount { get; private set; } = null!;

        /// <summary>
        /// Whether `ipv6_address_list` is allowed and controls the IPs to assign to the ENI and `ipv6_addresses` and `ipv6_address_count` become read-only. Default false.
        /// </summary>
        [Output("ipv6AddressListEnabled")]
        public Output<bool?> Ipv6AddressListEnabled { get; private set; } = null!;

        /// <summary>
        /// List of private IPs to assign to the ENI in sequential order.
        /// </summary>
        [Output("ipv6AddressLists")]
        public Output<ImmutableArray<string>> Ipv6AddressLists { get; private set; } = null!;

        /// <summary>
        /// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. Addresses are assigned without regard to order. You can't use this option if you're specifying `ipv6_address_count`.
        /// </summary>
        [Output("ipv6Addresses")]
        public Output<ImmutableArray<string>> Ipv6Addresses { get; private set; } = null!;

        /// <summary>
        /// Number of IPv6 prefixes that AWS automatically assigns to the network interface.
        /// </summary>
        [Output("ipv6PrefixCount")]
        public Output<int> Ipv6PrefixCount { get; private set; } = null!;

        /// <summary>
        /// One or more IPv6 prefixes assigned to the network interface.
        /// </summary>
        [Output("ipv6Prefixes")]
        public Output<ImmutableArray<string>> Ipv6Prefixes { get; private set; } = null!;

        /// <summary>
        /// MAC address of the network interface.
        /// </summary>
        [Output("macAddress")]
        public Output<string> MacAddress { get; private set; } = null!;

        [Output("outpostArn")]
        public Output<string> OutpostArn { get; private set; } = null!;

        /// <summary>
        /// AWS account ID of the owner of the network interface.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// Private DNS name of the network interface (IPv4).
        /// </summary>
        [Output("privateDnsName")]
        public Output<string> PrivateDnsName { get; private set; } = null!;

        [Output("privateIp")]
        public Output<string> PrivateIp { get; private set; } = null!;

        /// <summary>
        /// Whether `private_ip_list` is allowed and controls the IPs to assign to the ENI and `private_ips` and `private_ips_count` become read-only. Default false.
        /// </summary>
        [Output("privateIpListEnabled")]
        public Output<bool?> PrivateIpListEnabled { get; private set; } = null!;

        /// <summary>
        /// List of private IPs to assign to the ENI in sequential order. Requires setting `private_ip_list_enabled` to `true`.
        /// </summary>
        [Output("privateIpLists")]
        public Output<ImmutableArray<string>> PrivateIpLists { get; private set; } = null!;

        /// <summary>
        /// List of private IPs to assign to the ENI without regard to order.
        /// </summary>
        [Output("privateIps")]
        public Output<ImmutableArray<string>> PrivateIps { get; private set; } = null!;

        /// <summary>
        /// Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + `private_ips_count`, as a primary private IP will be assiged to an ENI by default.
        /// </summary>
        [Output("privateIpsCount")]
        public Output<int> PrivateIpsCount { get; private set; } = null!;

        /// <summary>
        /// List of security group IDs to assign to the ENI.
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// Whether to enable source destination checking for the ENI. Default true.
        /// </summary>
        [Output("sourceDestCheck")]
        public Output<bool?> SourceDestCheck { get; private set; } = null!;

        /// <summary>
        /// Subnet ID to create the ENI in.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkInterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkInterface(string name, NetworkInterfaceArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/networkInterface:NetworkInterface", name, args ?? new NetworkInterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkInterface(string name, Input<string> id, NetworkInterfaceState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/networkInterface:NetworkInterface", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkInterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkInterface Get(string name, Input<string> id, NetworkInterfaceState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkInterface(name, id, state, options);
        }
    }

    public sealed class NetworkInterfaceArgs : global::Pulumi.ResourceArgs
    {
        [Input("attachments")]
        private InputList<Inputs.NetworkInterfaceAttachmentArgs>? _attachments;

        /// <summary>
        /// Configuration block to define the attachment of the ENI. See Attachment below for more details!
        /// </summary>
        public InputList<Inputs.NetworkInterfaceAttachmentArgs> Attachments
        {
            get => _attachments ?? (_attachments = new InputList<Inputs.NetworkInterfaceAttachmentArgs>());
            set => _attachments = value;
        }

        /// <summary>
        /// Description for the network interface.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Type of network interface to create. Set to `efa` for Elastic Fabric Adapter. Changing `interface_type` will cause the resource to be destroyed and re-created.
        /// </summary>
        [Input("interfaceType")]
        public Input<string>? InterfaceType { get; set; }

        /// <summary>
        /// Number of IPv4 prefixes that AWS automatically assigns to the network interface.
        /// </summary>
        [Input("ipv4PrefixCount")]
        public Input<int>? Ipv4PrefixCount { get; set; }

        [Input("ipv4Prefixes")]
        private InputList<string>? _ipv4Prefixes;

        /// <summary>
        /// One or more IPv4 prefixes assigned to the network interface.
        /// </summary>
        public InputList<string> Ipv4Prefixes
        {
            get => _ipv4Prefixes ?? (_ipv4Prefixes = new InputList<string>());
            set => _ipv4Prefixes = value;
        }

        /// <summary>
        /// Number of IPv6 addresses to assign to a network interface. You can't use this option if specifying specific `ipv6_addresses`. If your subnet has the AssignIpv6AddressOnCreation attribute set to `true`, you can specify `0` to override this setting.
        /// </summary>
        [Input("ipv6AddressCount")]
        public Input<int>? Ipv6AddressCount { get; set; }

        /// <summary>
        /// Whether `ipv6_address_list` is allowed and controls the IPs to assign to the ENI and `ipv6_addresses` and `ipv6_address_count` become read-only. Default false.
        /// </summary>
        [Input("ipv6AddressListEnabled")]
        public Input<bool>? Ipv6AddressListEnabled { get; set; }

        [Input("ipv6AddressLists")]
        private InputList<string>? _ipv6AddressLists;

        /// <summary>
        /// List of private IPs to assign to the ENI in sequential order.
        /// </summary>
        public InputList<string> Ipv6AddressLists
        {
            get => _ipv6AddressLists ?? (_ipv6AddressLists = new InputList<string>());
            set => _ipv6AddressLists = value;
        }

        [Input("ipv6Addresses")]
        private InputList<string>? _ipv6Addresses;

        /// <summary>
        /// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. Addresses are assigned without regard to order. You can't use this option if you're specifying `ipv6_address_count`.
        /// </summary>
        public InputList<string> Ipv6Addresses
        {
            get => _ipv6Addresses ?? (_ipv6Addresses = new InputList<string>());
            set => _ipv6Addresses = value;
        }

        /// <summary>
        /// Number of IPv6 prefixes that AWS automatically assigns to the network interface.
        /// </summary>
        [Input("ipv6PrefixCount")]
        public Input<int>? Ipv6PrefixCount { get; set; }

        [Input("ipv6Prefixes")]
        private InputList<string>? _ipv6Prefixes;

        /// <summary>
        /// One or more IPv6 prefixes assigned to the network interface.
        /// </summary>
        public InputList<string> Ipv6Prefixes
        {
            get => _ipv6Prefixes ?? (_ipv6Prefixes = new InputList<string>());
            set => _ipv6Prefixes = value;
        }

        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        /// <summary>
        /// Whether `private_ip_list` is allowed and controls the IPs to assign to the ENI and `private_ips` and `private_ips_count` become read-only. Default false.
        /// </summary>
        [Input("privateIpListEnabled")]
        public Input<bool>? PrivateIpListEnabled { get; set; }

        [Input("privateIpLists")]
        private InputList<string>? _privateIpLists;

        /// <summary>
        /// List of private IPs to assign to the ENI in sequential order. Requires setting `private_ip_list_enabled` to `true`.
        /// </summary>
        public InputList<string> PrivateIpLists
        {
            get => _privateIpLists ?? (_privateIpLists = new InputList<string>());
            set => _privateIpLists = value;
        }

        [Input("privateIps")]
        private InputList<string>? _privateIps;

        /// <summary>
        /// List of private IPs to assign to the ENI without regard to order.
        /// </summary>
        public InputList<string> PrivateIps
        {
            get => _privateIps ?? (_privateIps = new InputList<string>());
            set => _privateIps = value;
        }

        /// <summary>
        /// Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + `private_ips_count`, as a primary private IP will be assiged to an ENI by default.
        /// </summary>
        [Input("privateIpsCount")]
        public Input<int>? PrivateIpsCount { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// List of security group IDs to assign to the ENI.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// Whether to enable source destination checking for the ENI. Default true.
        /// </summary>
        [Input("sourceDestCheck")]
        public Input<bool>? SourceDestCheck { get; set; }

        /// <summary>
        /// Subnet ID to create the ENI in.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public NetworkInterfaceArgs()
        {
        }
        public static new NetworkInterfaceArgs Empty => new NetworkInterfaceArgs();
    }

    public sealed class NetworkInterfaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the network interface.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("attachments")]
        private InputList<Inputs.NetworkInterfaceAttachmentGetArgs>? _attachments;

        /// <summary>
        /// Configuration block to define the attachment of the ENI. See Attachment below for more details!
        /// </summary>
        public InputList<Inputs.NetworkInterfaceAttachmentGetArgs> Attachments
        {
            get => _attachments ?? (_attachments = new InputList<Inputs.NetworkInterfaceAttachmentGetArgs>());
            set => _attachments = value;
        }

        /// <summary>
        /// Description for the network interface.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Type of network interface to create. Set to `efa` for Elastic Fabric Adapter. Changing `interface_type` will cause the resource to be destroyed and re-created.
        /// </summary>
        [Input("interfaceType")]
        public Input<string>? InterfaceType { get; set; }

        /// <summary>
        /// Number of IPv4 prefixes that AWS automatically assigns to the network interface.
        /// </summary>
        [Input("ipv4PrefixCount")]
        public Input<int>? Ipv4PrefixCount { get; set; }

        [Input("ipv4Prefixes")]
        private InputList<string>? _ipv4Prefixes;

        /// <summary>
        /// One or more IPv4 prefixes assigned to the network interface.
        /// </summary>
        public InputList<string> Ipv4Prefixes
        {
            get => _ipv4Prefixes ?? (_ipv4Prefixes = new InputList<string>());
            set => _ipv4Prefixes = value;
        }

        /// <summary>
        /// Number of IPv6 addresses to assign to a network interface. You can't use this option if specifying specific `ipv6_addresses`. If your subnet has the AssignIpv6AddressOnCreation attribute set to `true`, you can specify `0` to override this setting.
        /// </summary>
        [Input("ipv6AddressCount")]
        public Input<int>? Ipv6AddressCount { get; set; }

        /// <summary>
        /// Whether `ipv6_address_list` is allowed and controls the IPs to assign to the ENI and `ipv6_addresses` and `ipv6_address_count` become read-only. Default false.
        /// </summary>
        [Input("ipv6AddressListEnabled")]
        public Input<bool>? Ipv6AddressListEnabled { get; set; }

        [Input("ipv6AddressLists")]
        private InputList<string>? _ipv6AddressLists;

        /// <summary>
        /// List of private IPs to assign to the ENI in sequential order.
        /// </summary>
        public InputList<string> Ipv6AddressLists
        {
            get => _ipv6AddressLists ?? (_ipv6AddressLists = new InputList<string>());
            set => _ipv6AddressLists = value;
        }

        [Input("ipv6Addresses")]
        private InputList<string>? _ipv6Addresses;

        /// <summary>
        /// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. Addresses are assigned without regard to order. You can't use this option if you're specifying `ipv6_address_count`.
        /// </summary>
        public InputList<string> Ipv6Addresses
        {
            get => _ipv6Addresses ?? (_ipv6Addresses = new InputList<string>());
            set => _ipv6Addresses = value;
        }

        /// <summary>
        /// Number of IPv6 prefixes that AWS automatically assigns to the network interface.
        /// </summary>
        [Input("ipv6PrefixCount")]
        public Input<int>? Ipv6PrefixCount { get; set; }

        [Input("ipv6Prefixes")]
        private InputList<string>? _ipv6Prefixes;

        /// <summary>
        /// One or more IPv6 prefixes assigned to the network interface.
        /// </summary>
        public InputList<string> Ipv6Prefixes
        {
            get => _ipv6Prefixes ?? (_ipv6Prefixes = new InputList<string>());
            set => _ipv6Prefixes = value;
        }

        /// <summary>
        /// MAC address of the network interface.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        [Input("outpostArn")]
        public Input<string>? OutpostArn { get; set; }

        /// <summary>
        /// AWS account ID of the owner of the network interface.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        /// <summary>
        /// Private DNS name of the network interface (IPv4).
        /// </summary>
        [Input("privateDnsName")]
        public Input<string>? PrivateDnsName { get; set; }

        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        /// <summary>
        /// Whether `private_ip_list` is allowed and controls the IPs to assign to the ENI and `private_ips` and `private_ips_count` become read-only. Default false.
        /// </summary>
        [Input("privateIpListEnabled")]
        public Input<bool>? PrivateIpListEnabled { get; set; }

        [Input("privateIpLists")]
        private InputList<string>? _privateIpLists;

        /// <summary>
        /// List of private IPs to assign to the ENI in sequential order. Requires setting `private_ip_list_enabled` to `true`.
        /// </summary>
        public InputList<string> PrivateIpLists
        {
            get => _privateIpLists ?? (_privateIpLists = new InputList<string>());
            set => _privateIpLists = value;
        }

        [Input("privateIps")]
        private InputList<string>? _privateIps;

        /// <summary>
        /// List of private IPs to assign to the ENI without regard to order.
        /// </summary>
        public InputList<string> PrivateIps
        {
            get => _privateIps ?? (_privateIps = new InputList<string>());
            set => _privateIps = value;
        }

        /// <summary>
        /// Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + `private_ips_count`, as a primary private IP will be assiged to an ENI by default.
        /// </summary>
        [Input("privateIpsCount")]
        public Input<int>? PrivateIpsCount { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// List of security group IDs to assign to the ENI.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// Whether to enable source destination checking for the ENI. Default true.
        /// </summary>
        [Input("sourceDestCheck")]
        public Input<bool>? SourceDestCheck { get; set; }

        /// <summary>
        /// Subnet ID to create the ENI in.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public NetworkInterfaceState()
        {
        }
        public static new NetworkInterfaceState Empty => new NetworkInterfaceState();
    }
}
