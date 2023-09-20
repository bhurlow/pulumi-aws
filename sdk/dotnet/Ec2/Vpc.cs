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
    /// Provides a VPC resource.
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
    ///     var main = new Aws.Ec2.Vpc("main", new()
    ///     {
    ///         CidrBlock = "10.0.0.0/16",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Basic usage with tags:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Aws.Ec2.Vpc("main", new()
    ///     {
    ///         CidrBlock = "10.0.0.0/16",
    ///         InstanceTenancy = "default",
    ///         Tags = 
    ///         {
    ///             { "Name", "main" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// VPC with CIDR from AWS IPAM:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var current = Aws.GetRegion.Invoke();
    /// 
    ///     var testVpcIpam = new Aws.Ec2.VpcIpam("testVpcIpam", new()
    ///     {
    ///         OperatingRegions = new[]
    ///         {
    ///             new Aws.Ec2.Inputs.VpcIpamOperatingRegionArgs
    ///             {
    ///                 RegionName = current.Apply(getRegionResult =&gt; getRegionResult.Name),
    ///             },
    ///         },
    ///     });
    /// 
    ///     var testVpcIpamPool = new Aws.Ec2.VpcIpamPool("testVpcIpamPool", new()
    ///     {
    ///         AddressFamily = "ipv4",
    ///         IpamScopeId = testVpcIpam.PrivateDefaultScopeId,
    ///         Locale = current.Apply(getRegionResult =&gt; getRegionResult.Name),
    ///     });
    /// 
    ///     var testVpcIpamPoolCidr = new Aws.Ec2.VpcIpamPoolCidr("testVpcIpamPoolCidr", new()
    ///     {
    ///         IpamPoolId = testVpcIpamPool.Id,
    ///         Cidr = "172.20.0.0/16",
    ///     });
    /// 
    ///     var testVpc = new Aws.Ec2.Vpc("testVpc", new()
    ///     {
    ///         Ipv4IpamPoolId = testVpcIpamPool.Id,
    ///         Ipv4NetmaskLength = 28,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             testVpcIpamPoolCidr,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import VPCs using the VPC `id`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2/vpc:Vpc test_vpc vpc-a01106c2
    /// ```
    /// </summary>
    [AwsResourceType("aws:ec2/vpc:Vpc")]
    public partial class Vpc : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of VPC
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block. Default is `false`. Conflicts with `ipv6_ipam_pool_id`
        /// </summary>
        [Output("assignGeneratedIpv6CidrBlock")]
        public Output<bool?> AssignGeneratedIpv6CidrBlock { get; private set; } = null!;

        /// <summary>
        /// The IPv4 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv4_netmask_length`.
        /// </summary>
        [Output("cidrBlock")]
        public Output<string> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// The ID of the network ACL created by default on VPC creation
        /// </summary>
        [Output("defaultNetworkAclId")]
        public Output<string> DefaultNetworkAclId { get; private set; } = null!;

        /// <summary>
        /// The ID of the route table created by default on VPC creation
        /// </summary>
        [Output("defaultRouteTableId")]
        public Output<string> DefaultRouteTableId { get; private set; } = null!;

        /// <summary>
        /// The ID of the security group created by default on VPC creation
        /// </summary>
        [Output("defaultSecurityGroupId")]
        public Output<string> DefaultSecurityGroupId { get; private set; } = null!;

        [Output("dhcpOptionsId")]
        public Output<string> DhcpOptionsId { get; private set; } = null!;

        /// <summary>
        /// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
        /// </summary>
        [Output("enableDnsHostnames")]
        public Output<bool> EnableDnsHostnames { get; private set; } = null!;

        /// <summary>
        /// A boolean flag to enable/disable DNS support in the VPC. Defaults to true.
        /// </summary>
        [Output("enableDnsSupport")]
        public Output<bool?> EnableDnsSupport { get; private set; } = null!;

        /// <summary>
        /// Indicates whether Network Address Usage metrics are enabled for your VPC. Defaults to false.
        /// </summary>
        [Output("enableNetworkAddressUsageMetrics")]
        public Output<bool> EnableNetworkAddressUsageMetrics { get; private set; } = null!;

        /// <summary>
        /// A tenancy option for instances launched into the VPC. Default is `default`, which ensures that EC2 instances launched in this VPC use the EC2 instance tenancy attribute specified when the EC2 instance is launched. The only other option is `dedicated`, which ensures that EC2 instances launched in this VPC are run on dedicated tenancy instances regardless of the tenancy attribute specified at launch. This has a dedicated per region fee of $2 per hour, plus an hourly per instance usage fee.
        /// </summary>
        [Output("instanceTenancy")]
        public Output<string?> InstanceTenancy { get; private set; } = null!;

        /// <summary>
        /// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts. Using IPAM you can monitor IP address usage throughout your AWS Organization.
        /// </summary>
        [Output("ipv4IpamPoolId")]
        public Output<string?> Ipv4IpamPoolId { get; private set; } = null!;

        /// <summary>
        /// The netmask length of the IPv4 CIDR you want to allocate to this VPC. Requires specifying a `ipv4_ipam_pool_id`.
        /// </summary>
        [Output("ipv4NetmaskLength")]
        public Output<int?> Ipv4NetmaskLength { get; private set; } = null!;

        /// <summary>
        /// The association ID for the IPv6 CIDR block.
        /// </summary>
        [Output("ipv6AssociationId")]
        public Output<string> Ipv6AssociationId { get; private set; } = null!;

        /// <summary>
        /// IPv6 CIDR block to request from an IPAM Pool. Can be set explicitly or derived from IPAM using `ipv6_netmask_length`.
        /// </summary>
        [Output("ipv6CidrBlock")]
        public Output<string> Ipv6CidrBlock { get; private set; } = null!;

        /// <summary>
        /// By default when an IPv6 CIDR is assigned to a VPC a default ipv6_cidr_block_network_border_group will be set to the region of the VPC. This can be changed to restrict advertisement of public addresses to specific Network Border Groups such as LocalZones.
        /// </summary>
        [Output("ipv6CidrBlockNetworkBorderGroup")]
        public Output<string> Ipv6CidrBlockNetworkBorderGroup { get; private set; } = null!;

        /// <summary>
        /// IPAM Pool ID for a IPv6 pool. Conflicts with `assign_generated_ipv6_cidr_block`.
        /// </summary>
        [Output("ipv6IpamPoolId")]
        public Output<string?> Ipv6IpamPoolId { get; private set; } = null!;

        /// <summary>
        /// Netmask length to request from IPAM Pool. Conflicts with `ipv6_cidr_block`. This can be omitted if IPAM pool as a `allocation_default_netmask_length` set. Valid values: `56`.
        /// </summary>
        [Output("ipv6NetmaskLength")]
        public Output<int?> Ipv6NetmaskLength { get; private set; } = null!;

        /// <summary>
        /// The ID of the main route table associated with
        /// this VPC. Note that you can change a VPC's main route table by using an
        /// `aws.ec2.MainRouteTableAssociation`.
        /// </summary>
        [Output("mainRouteTableId")]
        public Output<string> MainRouteTableId { get; private set; } = null!;

        /// <summary>
        /// The ID of the AWS account that owns the VPC.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

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
        /// Create a Vpc resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vpc(string name, VpcArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:ec2/vpc:Vpc", name, args ?? new VpcArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Vpc(string name, Input<string> id, VpcState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/vpc:Vpc", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Vpc resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vpc Get(string name, Input<string> id, VpcState? state = null, CustomResourceOptions? options = null)
        {
            return new Vpc(name, id, state, options);
        }
    }

    public sealed class VpcArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block. Default is `false`. Conflicts with `ipv6_ipam_pool_id`
        /// </summary>
        [Input("assignGeneratedIpv6CidrBlock")]
        public Input<bool>? AssignGeneratedIpv6CidrBlock { get; set; }

        /// <summary>
        /// The IPv4 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv4_netmask_length`.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
        /// </summary>
        [Input("enableDnsHostnames")]
        public Input<bool>? EnableDnsHostnames { get; set; }

        /// <summary>
        /// A boolean flag to enable/disable DNS support in the VPC. Defaults to true.
        /// </summary>
        [Input("enableDnsSupport")]
        public Input<bool>? EnableDnsSupport { get; set; }

        /// <summary>
        /// Indicates whether Network Address Usage metrics are enabled for your VPC. Defaults to false.
        /// </summary>
        [Input("enableNetworkAddressUsageMetrics")]
        public Input<bool>? EnableNetworkAddressUsageMetrics { get; set; }

        /// <summary>
        /// A tenancy option for instances launched into the VPC. Default is `default`, which ensures that EC2 instances launched in this VPC use the EC2 instance tenancy attribute specified when the EC2 instance is launched. The only other option is `dedicated`, which ensures that EC2 instances launched in this VPC are run on dedicated tenancy instances regardless of the tenancy attribute specified at launch. This has a dedicated per region fee of $2 per hour, plus an hourly per instance usage fee.
        /// </summary>
        [Input("instanceTenancy")]
        public Input<string>? InstanceTenancy { get; set; }

        /// <summary>
        /// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts. Using IPAM you can monitor IP address usage throughout your AWS Organization.
        /// </summary>
        [Input("ipv4IpamPoolId")]
        public Input<string>? Ipv4IpamPoolId { get; set; }

        /// <summary>
        /// The netmask length of the IPv4 CIDR you want to allocate to this VPC. Requires specifying a `ipv4_ipam_pool_id`.
        /// </summary>
        [Input("ipv4NetmaskLength")]
        public Input<int>? Ipv4NetmaskLength { get; set; }

        /// <summary>
        /// IPv6 CIDR block to request from an IPAM Pool. Can be set explicitly or derived from IPAM using `ipv6_netmask_length`.
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        /// <summary>
        /// By default when an IPv6 CIDR is assigned to a VPC a default ipv6_cidr_block_network_border_group will be set to the region of the VPC. This can be changed to restrict advertisement of public addresses to specific Network Border Groups such as LocalZones.
        /// </summary>
        [Input("ipv6CidrBlockNetworkBorderGroup")]
        public Input<string>? Ipv6CidrBlockNetworkBorderGroup { get; set; }

        /// <summary>
        /// IPAM Pool ID for a IPv6 pool. Conflicts with `assign_generated_ipv6_cidr_block`.
        /// </summary>
        [Input("ipv6IpamPoolId")]
        public Input<string>? Ipv6IpamPoolId { get; set; }

        /// <summary>
        /// Netmask length to request from IPAM Pool. Conflicts with `ipv6_cidr_block`. This can be omitted if IPAM pool as a `allocation_default_netmask_length` set. Valid values: `56`.
        /// </summary>
        [Input("ipv6NetmaskLength")]
        public Input<int>? Ipv6NetmaskLength { get; set; }

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

        public VpcArgs()
        {
        }
        public static new VpcArgs Empty => new VpcArgs();
    }

    public sealed class VpcState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of VPC
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block. Default is `false`. Conflicts with `ipv6_ipam_pool_id`
        /// </summary>
        [Input("assignGeneratedIpv6CidrBlock")]
        public Input<bool>? AssignGeneratedIpv6CidrBlock { get; set; }

        /// <summary>
        /// The IPv4 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv4_netmask_length`.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// The ID of the network ACL created by default on VPC creation
        /// </summary>
        [Input("defaultNetworkAclId")]
        public Input<string>? DefaultNetworkAclId { get; set; }

        /// <summary>
        /// The ID of the route table created by default on VPC creation
        /// </summary>
        [Input("defaultRouteTableId")]
        public Input<string>? DefaultRouteTableId { get; set; }

        /// <summary>
        /// The ID of the security group created by default on VPC creation
        /// </summary>
        [Input("defaultSecurityGroupId")]
        public Input<string>? DefaultSecurityGroupId { get; set; }

        [Input("dhcpOptionsId")]
        public Input<string>? DhcpOptionsId { get; set; }

        /// <summary>
        /// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
        /// </summary>
        [Input("enableDnsHostnames")]
        public Input<bool>? EnableDnsHostnames { get; set; }

        /// <summary>
        /// A boolean flag to enable/disable DNS support in the VPC. Defaults to true.
        /// </summary>
        [Input("enableDnsSupport")]
        public Input<bool>? EnableDnsSupport { get; set; }

        /// <summary>
        /// Indicates whether Network Address Usage metrics are enabled for your VPC. Defaults to false.
        /// </summary>
        [Input("enableNetworkAddressUsageMetrics")]
        public Input<bool>? EnableNetworkAddressUsageMetrics { get; set; }

        /// <summary>
        /// A tenancy option for instances launched into the VPC. Default is `default`, which ensures that EC2 instances launched in this VPC use the EC2 instance tenancy attribute specified when the EC2 instance is launched. The only other option is `dedicated`, which ensures that EC2 instances launched in this VPC are run on dedicated tenancy instances regardless of the tenancy attribute specified at launch. This has a dedicated per region fee of $2 per hour, plus an hourly per instance usage fee.
        /// </summary>
        [Input("instanceTenancy")]
        public Input<string>? InstanceTenancy { get; set; }

        /// <summary>
        /// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts. Using IPAM you can monitor IP address usage throughout your AWS Organization.
        /// </summary>
        [Input("ipv4IpamPoolId")]
        public Input<string>? Ipv4IpamPoolId { get; set; }

        /// <summary>
        /// The netmask length of the IPv4 CIDR you want to allocate to this VPC. Requires specifying a `ipv4_ipam_pool_id`.
        /// </summary>
        [Input("ipv4NetmaskLength")]
        public Input<int>? Ipv4NetmaskLength { get; set; }

        /// <summary>
        /// The association ID for the IPv6 CIDR block.
        /// </summary>
        [Input("ipv6AssociationId")]
        public Input<string>? Ipv6AssociationId { get; set; }

        /// <summary>
        /// IPv6 CIDR block to request from an IPAM Pool. Can be set explicitly or derived from IPAM using `ipv6_netmask_length`.
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        /// <summary>
        /// By default when an IPv6 CIDR is assigned to a VPC a default ipv6_cidr_block_network_border_group will be set to the region of the VPC. This can be changed to restrict advertisement of public addresses to specific Network Border Groups such as LocalZones.
        /// </summary>
        [Input("ipv6CidrBlockNetworkBorderGroup")]
        public Input<string>? Ipv6CidrBlockNetworkBorderGroup { get; set; }

        /// <summary>
        /// IPAM Pool ID for a IPv6 pool. Conflicts with `assign_generated_ipv6_cidr_block`.
        /// </summary>
        [Input("ipv6IpamPoolId")]
        public Input<string>? Ipv6IpamPoolId { get; set; }

        /// <summary>
        /// Netmask length to request from IPAM Pool. Conflicts with `ipv6_cidr_block`. This can be omitted if IPAM pool as a `allocation_default_netmask_length` set. Valid values: `56`.
        /// </summary>
        [Input("ipv6NetmaskLength")]
        public Input<int>? Ipv6NetmaskLength { get; set; }

        /// <summary>
        /// The ID of the main route table associated with
        /// this VPC. Note that you can change a VPC's main route table by using an
        /// `aws.ec2.MainRouteTableAssociation`.
        /// </summary>
        [Input("mainRouteTableId")]
        public Input<string>? MainRouteTableId { get; set; }

        /// <summary>
        /// The ID of the AWS account that owns the VPC.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

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

        public VpcState()
        {
        }
        public static new VpcState Empty => new VpcState();
    }
}
