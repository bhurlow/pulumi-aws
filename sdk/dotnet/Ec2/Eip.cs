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
    /// Provides an Elastic IP resource.
    /// 
    /// &gt; **Note:** EIP may require IGW to exist prior to association. Use `depends_on` to set an explicit dependency on the IGW.
    /// 
    /// &gt; **Note:** Do not use `network_interface` to associate the EIP to `aws.lb.LoadBalancer` or `aws.ec2.NatGateway` resources. Instead use the `allocation_id` available in those resources to allow AWS to manage the association, otherwise you will see `AuthFailure` errors.
    /// 
    /// ## Example Usage
    /// ### Single EIP associated with an instance
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var lb = new Aws.Ec2.Eip("lb", new()
    ///     {
    ///         Instance = aws_instance.Web.Id,
    ///         Vpc = true,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Multiple EIPs associated with a single network interface
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var multi_ip = new Aws.Ec2.NetworkInterface("multi-ip", new()
    ///     {
    ///         SubnetId = aws_subnet.Main.Id,
    ///         PrivateIps = new[]
    ///         {
    ///             "10.0.0.10",
    ///             "10.0.0.11",
    ///         },
    ///     });
    /// 
    ///     var one = new Aws.Ec2.Eip("one", new()
    ///     {
    ///         Vpc = true,
    ///         NetworkInterface = multi_ip.Id,
    ///         AssociateWithPrivateIp = "10.0.0.10",
    ///     });
    /// 
    ///     var two = new Aws.Ec2.Eip("two", new()
    ///     {
    ///         Vpc = true,
    ///         NetworkInterface = multi_ip.Id,
    ///         AssociateWithPrivateIp = "10.0.0.11",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Attaching an EIP to an Instance with a pre-assigned private ip (VPC Only)
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Aws.Ec2.Vpc("default", new()
    ///     {
    ///         CidrBlock = "10.0.0.0/16",
    ///         EnableDnsHostnames = true,
    ///     });
    /// 
    ///     var gw = new Aws.Ec2.InternetGateway("gw", new()
    ///     {
    ///         VpcId = @default.Id,
    ///     });
    /// 
    ///     var myTestSubnet = new Aws.Ec2.Subnet("myTestSubnet", new()
    ///     {
    ///         VpcId = @default.Id,
    ///         CidrBlock = "10.0.0.0/24",
    ///         MapPublicIpOnLaunch = true,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             gw,
    ///         },
    ///     });
    /// 
    ///     var foo = new Aws.Ec2.Instance("foo", new()
    ///     {
    ///         Ami = "ami-5189a661",
    ///         InstanceType = "t2.micro",
    ///         PrivateIp = "10.0.0.12",
    ///         SubnetId = myTestSubnet.Id,
    ///     });
    /// 
    ///     var bar = new Aws.Ec2.Eip("bar", new()
    ///     {
    ///         Vpc = true,
    ///         Instance = foo.Id,
    ///         AssociateWithPrivateIp = "10.0.0.12",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             gw,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Allocating EIP from the BYOIP pool
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var byoip_ip = new Aws.Ec2.Eip("byoip-ip", new()
    ///     {
    ///         PublicIpv4Pool = "ipv4pool-ec2-012345",
    ///         Vpc = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// EIPs in a VPC can be imported using their Allocation ID, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2/eip:Eip bar eipalloc-00a10e96
    /// ```
    /// 
    ///  EIPs in EC2-Classic can be imported using their Public IP, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2/eip:Eip bar 52.0.0.0
    /// ```
    /// </summary>
    [AwsResourceType("aws:ec2/eip:Eip")]
    public partial class Eip : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IP address from an EC2 BYOIP pool. This option is only available for VPC EIPs.
        /// </summary>
        [Output("address")]
        public Output<string?> Address { get; private set; } = null!;

        /// <summary>
        /// ID that AWS assigns to represent the allocation of the Elastic IP address for use with instances in a VPC.
        /// </summary>
        [Output("allocationId")]
        public Output<string> AllocationId { get; private set; } = null!;

        /// <summary>
        /// User-specified primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.
        /// </summary>
        [Output("associateWithPrivateIp")]
        public Output<string?> AssociateWithPrivateIp { get; private set; } = null!;

        /// <summary>
        /// ID representing the association of the address with an instance in a VPC.
        /// </summary>
        [Output("associationId")]
        public Output<string> AssociationId { get; private set; } = null!;

        /// <summary>
        /// Carrier IP address.
        /// </summary>
        [Output("carrierIp")]
        public Output<string> CarrierIp { get; private set; } = null!;

        /// <summary>
        /// Customer owned IP.
        /// </summary>
        [Output("customerOwnedIp")]
        public Output<string> CustomerOwnedIp { get; private set; } = null!;

        /// <summary>
        /// ID  of a customer-owned address pool. For more on customer owned IP addressed check out [Customer-owned IP addresses guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing).
        /// </summary>
        [Output("customerOwnedIpv4Pool")]
        public Output<string?> CustomerOwnedIpv4Pool { get; private set; } = null!;

        /// <summary>
        /// Indicates if this EIP is for use in VPC (`vpc`) or EC2-Classic (`standard`).
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// EC2 instance ID.
        /// </summary>
        [Output("instance")]
        public Output<string> Instance { get; private set; } = null!;

        /// <summary>
        /// Location from which the IP address is advertised. Use this parameter to limit the address to this location.
        /// </summary>
        [Output("networkBorderGroup")]
        public Output<string> NetworkBorderGroup { get; private set; } = null!;

        /// <summary>
        /// Network interface ID to associate with.
        /// </summary>
        [Output("networkInterface")]
        public Output<string> NetworkInterface { get; private set; } = null!;

        /// <summary>
        /// The Private DNS associated with the Elastic IP address (if in VPC).
        /// </summary>
        [Output("privateDns")]
        public Output<string> PrivateDns { get; private set; } = null!;

        /// <summary>
        /// Contains the private IP address (if in VPC).
        /// </summary>
        [Output("privateIp")]
        public Output<string> PrivateIp { get; private set; } = null!;

        /// <summary>
        /// Public DNS associated with the Elastic IP address.
        /// </summary>
        [Output("publicDns")]
        public Output<string> PublicDns { get; private set; } = null!;

        /// <summary>
        /// Contains the public IP address.
        /// </summary>
        [Output("publicIp")]
        public Output<string> PublicIp { get; private set; } = null!;

        /// <summary>
        /// EC2 IPv4 address pool identifier or `amazon`.
        /// This option is only available for VPC EIPs.
        /// </summary>
        [Output("publicIpv4Pool")]
        public Output<string> PublicIpv4Pool { get; private set; } = null!;

        /// <summary>
        /// Map of tags to assign to the resource. Tags can only be applied to EIPs in a VPC. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Boolean if the EIP is in a VPC or not.
        /// Defaults to `true` unless the region supports EC2-Classic.
        /// </summary>
        [Output("vpc")]
        public Output<bool> Vpc { get; private set; } = null!;


        /// <summary>
        /// Create a Eip resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Eip(string name, EipArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:ec2/eip:Eip", name, args ?? new EipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Eip(string name, Input<string> id, EipState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/eip:Eip", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Eip resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Eip Get(string name, Input<string> id, EipState? state = null, CustomResourceOptions? options = null)
        {
            return new Eip(name, id, state, options);
        }
    }

    public sealed class EipArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address from an EC2 BYOIP pool. This option is only available for VPC EIPs.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// User-specified primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.
        /// </summary>
        [Input("associateWithPrivateIp")]
        public Input<string>? AssociateWithPrivateIp { get; set; }

        /// <summary>
        /// ID  of a customer-owned address pool. For more on customer owned IP addressed check out [Customer-owned IP addresses guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing).
        /// </summary>
        [Input("customerOwnedIpv4Pool")]
        public Input<string>? CustomerOwnedIpv4Pool { get; set; }

        /// <summary>
        /// EC2 instance ID.
        /// </summary>
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        /// <summary>
        /// Location from which the IP address is advertised. Use this parameter to limit the address to this location.
        /// </summary>
        [Input("networkBorderGroup")]
        public Input<string>? NetworkBorderGroup { get; set; }

        /// <summary>
        /// Network interface ID to associate with.
        /// </summary>
        [Input("networkInterface")]
        public Input<string>? NetworkInterface { get; set; }

        /// <summary>
        /// EC2 IPv4 address pool identifier or `amazon`.
        /// This option is only available for VPC EIPs.
        /// </summary>
        [Input("publicIpv4Pool")]
        public Input<string>? PublicIpv4Pool { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the resource. Tags can only be applied to EIPs in a VPC. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Boolean if the EIP is in a VPC or not.
        /// Defaults to `true` unless the region supports EC2-Classic.
        /// </summary>
        [Input("vpc")]
        public Input<bool>? Vpc { get; set; }

        public EipArgs()
        {
        }
        public static new EipArgs Empty => new EipArgs();
    }

    public sealed class EipState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address from an EC2 BYOIP pool. This option is only available for VPC EIPs.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// ID that AWS assigns to represent the allocation of the Elastic IP address for use with instances in a VPC.
        /// </summary>
        [Input("allocationId")]
        public Input<string>? AllocationId { get; set; }

        /// <summary>
        /// User-specified primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.
        /// </summary>
        [Input("associateWithPrivateIp")]
        public Input<string>? AssociateWithPrivateIp { get; set; }

        /// <summary>
        /// ID representing the association of the address with an instance in a VPC.
        /// </summary>
        [Input("associationId")]
        public Input<string>? AssociationId { get; set; }

        /// <summary>
        /// Carrier IP address.
        /// </summary>
        [Input("carrierIp")]
        public Input<string>? CarrierIp { get; set; }

        /// <summary>
        /// Customer owned IP.
        /// </summary>
        [Input("customerOwnedIp")]
        public Input<string>? CustomerOwnedIp { get; set; }

        /// <summary>
        /// ID  of a customer-owned address pool. For more on customer owned IP addressed check out [Customer-owned IP addresses guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing).
        /// </summary>
        [Input("customerOwnedIpv4Pool")]
        public Input<string>? CustomerOwnedIpv4Pool { get; set; }

        /// <summary>
        /// Indicates if this EIP is for use in VPC (`vpc`) or EC2-Classic (`standard`).
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// EC2 instance ID.
        /// </summary>
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        /// <summary>
        /// Location from which the IP address is advertised. Use this parameter to limit the address to this location.
        /// </summary>
        [Input("networkBorderGroup")]
        public Input<string>? NetworkBorderGroup { get; set; }

        /// <summary>
        /// Network interface ID to associate with.
        /// </summary>
        [Input("networkInterface")]
        public Input<string>? NetworkInterface { get; set; }

        /// <summary>
        /// The Private DNS associated with the Elastic IP address (if in VPC).
        /// </summary>
        [Input("privateDns")]
        public Input<string>? PrivateDns { get; set; }

        /// <summary>
        /// Contains the private IP address (if in VPC).
        /// </summary>
        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        /// <summary>
        /// Public DNS associated with the Elastic IP address.
        /// </summary>
        [Input("publicDns")]
        public Input<string>? PublicDns { get; set; }

        /// <summary>
        /// Contains the public IP address.
        /// </summary>
        [Input("publicIp")]
        public Input<string>? PublicIp { get; set; }

        /// <summary>
        /// EC2 IPv4 address pool identifier or `amazon`.
        /// This option is only available for VPC EIPs.
        /// </summary>
        [Input("publicIpv4Pool")]
        public Input<string>? PublicIpv4Pool { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the resource. Tags can only be applied to EIPs in a VPC. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        /// Boolean if the EIP is in a VPC or not.
        /// Defaults to `true` unless the region supports EC2-Classic.
        /// </summary>
        [Input("vpc")]
        public Input<bool>? Vpc { get; set; }

        public EipState()
        {
        }
        public static new EipState Empty => new EipState();
    }
}
