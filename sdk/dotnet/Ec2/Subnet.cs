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
    /// Provides an VPC subnet resource.
    /// 
    /// &gt; **NOTE:** Due to [AWS Lambda improved VPC networking changes that began deploying in September 2019](https://aws.amazon.com/blogs/compute/announcing-improved-vpc-networking-for-aws-lambda-functions/), subnets associated with Lambda Functions can take up to 45 minutes to successfully delete.
    /// 
    /// ## Example Usage
    /// ### Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Aws.Ec2.Subnet("main", new()
    ///     {
    ///         VpcId = aws_vpc.Main.Id,
    ///         CidrBlock = "10.0.1.0/24",
    ///         Tags = 
    ///         {
    ///             { "Name", "Main" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Subnets In Secondary VPC CIDR Blocks
    /// 
    /// When managing subnets in one of a VPC's secondary CIDR blocks created using a `aws.ec2.VpcIpv4CidrBlockAssociation`
    /// resource, it is recommended to reference that resource's `vpc_id` attribute to ensure correct dependency ordering.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var secondaryCidr = new Aws.Ec2.VpcIpv4CidrBlockAssociation("secondaryCidr", new()
    ///     {
    ///         VpcId = aws_vpc.Main.Id,
    ///         CidrBlock = "172.2.0.0/16",
    ///     });
    /// 
    ///     var inSecondaryCidr = new Aws.Ec2.Subnet("inSecondaryCidr", new()
    ///     {
    ///         VpcId = secondaryCidr.VpcId,
    ///         CidrBlock = "172.2.0.0/24",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Subnets can be imported using the `subnet id`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2/subnet:Subnet public_subnet subnet-9d4a7b6c
    /// ```
    /// </summary>
    [AwsResourceType("aws:ec2/subnet:Subnet")]
    public partial class Subnet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the subnet.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Specify true to indicate
        /// that network interfaces created in the specified subnet should be
        /// assigned an IPv6 address. Default is `false`
        /// </summary>
        [Output("assignIpv6AddressOnCreation")]
        public Output<bool?> AssignIpv6AddressOnCreation { get; private set; } = null!;

        /// <summary>
        /// AZ for the subnet.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// AZ ID of the subnet. This argument is not supported in all regions or partitions. If necessary, use `availability_zone` instead.
        /// </summary>
        [Output("availabilityZoneId")]
        public Output<string> AvailabilityZoneId { get; private set; } = null!;

        /// <summary>
        /// The IPv4 CIDR block for the subnet.
        /// </summary>
        [Output("cidrBlock")]
        public Output<string?> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// The customer owned IPv4 address pool. Typically used with the `map_customer_owned_ip_on_launch` argument. The `outpost_arn` argument must be specified when configured.
        /// </summary>
        [Output("customerOwnedIpv4Pool")]
        public Output<string?> CustomerOwnedIpv4Pool { get; private set; } = null!;

        /// <summary>
        /// Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations. Default: `false`.
        /// </summary>
        [Output("enableDns64")]
        public Output<bool?> EnableDns64 { get; private set; } = null!;

        /// <summary>
        /// Indicates the device position for local network interfaces in this subnet. For example, 1 indicates local network interfaces in this subnet are the secondary network interface (eth1). A local network interface cannot be the primary network interface (eth0).
        /// </summary>
        [Output("enableLniAtDeviceIndex")]
        public Output<int?> EnableLniAtDeviceIndex { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to respond to DNS queries for instance hostnames with DNS A records. Default: `false`.
        /// </summary>
        [Output("enableResourceNameDnsARecordOnLaunch")]
        public Output<bool?> EnableResourceNameDnsARecordOnLaunch { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records. Default: `false`.
        /// </summary>
        [Output("enableResourceNameDnsAaaaRecordOnLaunch")]
        public Output<bool?> EnableResourceNameDnsAaaaRecordOnLaunch { get; private set; } = null!;

        /// <summary>
        /// The IPv6 network range for the subnet,
        /// in CIDR notation. The subnet size must use a /64 prefix length.
        /// </summary>
        [Output("ipv6CidrBlock")]
        public Output<string?> Ipv6CidrBlock { get; private set; } = null!;

        /// <summary>
        /// The association ID for the IPv6 CIDR block.
        /// </summary>
        [Output("ipv6CidrBlockAssociationId")]
        public Output<string> Ipv6CidrBlockAssociationId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to create an IPv6-only subnet. Default: `false`.
        /// </summary>
        [Output("ipv6Native")]
        public Output<bool?> Ipv6Native { get; private set; } = null!;

        /// <summary>
        /// Specify `true` to indicate that network interfaces created in the subnet should be assigned a customer owned IP address. The `customer_owned_ipv4_pool` and `outpost_arn` arguments must be specified when set to `true`. Default is `false`.
        /// </summary>
        [Output("mapCustomerOwnedIpOnLaunch")]
        public Output<bool?> MapCustomerOwnedIpOnLaunch { get; private set; } = null!;

        /// <summary>
        /// Specify true to indicate
        /// that instances launched into the subnet should be assigned
        /// a public IP address. Default is `false`.
        /// </summary>
        [Output("mapPublicIpOnLaunch")]
        public Output<bool?> MapPublicIpOnLaunch { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Outpost.
        /// </summary>
        [Output("outpostArn")]
        public Output<string?> OutpostArn { get; private set; } = null!;

        /// <summary>
        /// The ID of the AWS account that owns the subnet.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// The type of hostnames to assign to instances in the subnet at launch. For IPv6-only subnets, an instance DNS name must be based on the instance ID. For dual-stack and IPv4-only subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. Valid values: `ip-name`, `resource-name`.
        /// </summary>
        [Output("privateDnsHostnameTypeOnLaunch")]
        public Output<string> PrivateDnsHostnameTypeOnLaunch { get; private set; } = null!;

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
        /// The VPC ID.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Subnet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subnet(string name, SubnetArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/subnet:Subnet", name, args ?? new SubnetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Subnet(string name, Input<string> id, SubnetState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/subnet:Subnet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Subnet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Subnet Get(string name, Input<string> id, SubnetState? state = null, CustomResourceOptions? options = null)
        {
            return new Subnet(name, id, state, options);
        }
    }

    public sealed class SubnetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify true to indicate
        /// that network interfaces created in the specified subnet should be
        /// assigned an IPv6 address. Default is `false`
        /// </summary>
        [Input("assignIpv6AddressOnCreation")]
        public Input<bool>? AssignIpv6AddressOnCreation { get; set; }

        /// <summary>
        /// AZ for the subnet.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// AZ ID of the subnet. This argument is not supported in all regions or partitions. If necessary, use `availability_zone` instead.
        /// </summary>
        [Input("availabilityZoneId")]
        public Input<string>? AvailabilityZoneId { get; set; }

        /// <summary>
        /// The IPv4 CIDR block for the subnet.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// The customer owned IPv4 address pool. Typically used with the `map_customer_owned_ip_on_launch` argument. The `outpost_arn` argument must be specified when configured.
        /// </summary>
        [Input("customerOwnedIpv4Pool")]
        public Input<string>? CustomerOwnedIpv4Pool { get; set; }

        /// <summary>
        /// Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations. Default: `false`.
        /// </summary>
        [Input("enableDns64")]
        public Input<bool>? EnableDns64 { get; set; }

        /// <summary>
        /// Indicates the device position for local network interfaces in this subnet. For example, 1 indicates local network interfaces in this subnet are the secondary network interface (eth1). A local network interface cannot be the primary network interface (eth0).
        /// </summary>
        [Input("enableLniAtDeviceIndex")]
        public Input<int>? EnableLniAtDeviceIndex { get; set; }

        /// <summary>
        /// Indicates whether to respond to DNS queries for instance hostnames with DNS A records. Default: `false`.
        /// </summary>
        [Input("enableResourceNameDnsARecordOnLaunch")]
        public Input<bool>? EnableResourceNameDnsARecordOnLaunch { get; set; }

        /// <summary>
        /// Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records. Default: `false`.
        /// </summary>
        [Input("enableResourceNameDnsAaaaRecordOnLaunch")]
        public Input<bool>? EnableResourceNameDnsAaaaRecordOnLaunch { get; set; }

        /// <summary>
        /// The IPv6 network range for the subnet,
        /// in CIDR notation. The subnet size must use a /64 prefix length.
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        /// <summary>
        /// Indicates whether to create an IPv6-only subnet. Default: `false`.
        /// </summary>
        [Input("ipv6Native")]
        public Input<bool>? Ipv6Native { get; set; }

        /// <summary>
        /// Specify `true` to indicate that network interfaces created in the subnet should be assigned a customer owned IP address. The `customer_owned_ipv4_pool` and `outpost_arn` arguments must be specified when set to `true`. Default is `false`.
        /// </summary>
        [Input("mapCustomerOwnedIpOnLaunch")]
        public Input<bool>? MapCustomerOwnedIpOnLaunch { get; set; }

        /// <summary>
        /// Specify true to indicate
        /// that instances launched into the subnet should be assigned
        /// a public IP address. Default is `false`.
        /// </summary>
        [Input("mapPublicIpOnLaunch")]
        public Input<bool>? MapPublicIpOnLaunch { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Outpost.
        /// </summary>
        [Input("outpostArn")]
        public Input<string>? OutpostArn { get; set; }

        /// <summary>
        /// The type of hostnames to assign to instances in the subnet at launch. For IPv6-only subnets, an instance DNS name must be based on the instance ID. For dual-stack and IPv4-only subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. Valid values: `ip-name`, `resource-name`.
        /// </summary>
        [Input("privateDnsHostnameTypeOnLaunch")]
        public Input<string>? PrivateDnsHostnameTypeOnLaunch { get; set; }

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

        /// <summary>
        /// The VPC ID.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public SubnetArgs()
        {
        }
        public static new SubnetArgs Empty => new SubnetArgs();
    }

    public sealed class SubnetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the subnet.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Specify true to indicate
        /// that network interfaces created in the specified subnet should be
        /// assigned an IPv6 address. Default is `false`
        /// </summary>
        [Input("assignIpv6AddressOnCreation")]
        public Input<bool>? AssignIpv6AddressOnCreation { get; set; }

        /// <summary>
        /// AZ for the subnet.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// AZ ID of the subnet. This argument is not supported in all regions or partitions. If necessary, use `availability_zone` instead.
        /// </summary>
        [Input("availabilityZoneId")]
        public Input<string>? AvailabilityZoneId { get; set; }

        /// <summary>
        /// The IPv4 CIDR block for the subnet.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// The customer owned IPv4 address pool. Typically used with the `map_customer_owned_ip_on_launch` argument. The `outpost_arn` argument must be specified when configured.
        /// </summary>
        [Input("customerOwnedIpv4Pool")]
        public Input<string>? CustomerOwnedIpv4Pool { get; set; }

        /// <summary>
        /// Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations. Default: `false`.
        /// </summary>
        [Input("enableDns64")]
        public Input<bool>? EnableDns64 { get; set; }

        /// <summary>
        /// Indicates the device position for local network interfaces in this subnet. For example, 1 indicates local network interfaces in this subnet are the secondary network interface (eth1). A local network interface cannot be the primary network interface (eth0).
        /// </summary>
        [Input("enableLniAtDeviceIndex")]
        public Input<int>? EnableLniAtDeviceIndex { get; set; }

        /// <summary>
        /// Indicates whether to respond to DNS queries for instance hostnames with DNS A records. Default: `false`.
        /// </summary>
        [Input("enableResourceNameDnsARecordOnLaunch")]
        public Input<bool>? EnableResourceNameDnsARecordOnLaunch { get; set; }

        /// <summary>
        /// Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records. Default: `false`.
        /// </summary>
        [Input("enableResourceNameDnsAaaaRecordOnLaunch")]
        public Input<bool>? EnableResourceNameDnsAaaaRecordOnLaunch { get; set; }

        /// <summary>
        /// The IPv6 network range for the subnet,
        /// in CIDR notation. The subnet size must use a /64 prefix length.
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        /// <summary>
        /// The association ID for the IPv6 CIDR block.
        /// </summary>
        [Input("ipv6CidrBlockAssociationId")]
        public Input<string>? Ipv6CidrBlockAssociationId { get; set; }

        /// <summary>
        /// Indicates whether to create an IPv6-only subnet. Default: `false`.
        /// </summary>
        [Input("ipv6Native")]
        public Input<bool>? Ipv6Native { get; set; }

        /// <summary>
        /// Specify `true` to indicate that network interfaces created in the subnet should be assigned a customer owned IP address. The `customer_owned_ipv4_pool` and `outpost_arn` arguments must be specified when set to `true`. Default is `false`.
        /// </summary>
        [Input("mapCustomerOwnedIpOnLaunch")]
        public Input<bool>? MapCustomerOwnedIpOnLaunch { get; set; }

        /// <summary>
        /// Specify true to indicate
        /// that instances launched into the subnet should be assigned
        /// a public IP address. Default is `false`.
        /// </summary>
        [Input("mapPublicIpOnLaunch")]
        public Input<bool>? MapPublicIpOnLaunch { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Outpost.
        /// </summary>
        [Input("outpostArn")]
        public Input<string>? OutpostArn { get; set; }

        /// <summary>
        /// The ID of the AWS account that owns the subnet.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        /// <summary>
        /// The type of hostnames to assign to instances in the subnet at launch. For IPv6-only subnets, an instance DNS name must be based on the instance ID. For dual-stack and IPv4-only subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. Valid values: `ip-name`, `resource-name`.
        /// </summary>
        [Input("privateDnsHostnameTypeOnLaunch")]
        public Input<string>? PrivateDnsHostnameTypeOnLaunch { get; set; }

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

        /// <summary>
        /// The VPC ID.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public SubnetState()
        {
        }
        public static new SubnetState Empty => new SubnetState();
    }
}
