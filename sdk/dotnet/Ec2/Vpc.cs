// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// Provides a VPC resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpc.html.markdown.
    /// </summary>
    public partial class Vpc : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of VPC
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Requests an Amazon-provided IPv6 CIDR
        /// block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or
        /// the size of the CIDR block. Default is `false`.
        /// </summary>
        [Output("assignGeneratedIpv6CidrBlock")]
        public Output<bool?> AssignGeneratedIpv6CidrBlock { get; private set; } = null!;

        /// <summary>
        /// The CIDR block for the VPC.
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
        /// A boolean flag to enable/disable ClassicLink
        /// for the VPC. Only valid in regions and accounts that support EC2 Classic.
        /// See the [ClassicLink documentation][1] for more information. Defaults false.
        /// </summary>
        [Output("enableClassiclink")]
        public Output<bool> EnableClassiclink { get; private set; } = null!;

        /// <summary>
        /// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
        /// Only valid in regions and accounts that support EC2 Classic.
        /// </summary>
        [Output("enableClassiclinkDnsSupport")]
        public Output<bool> EnableClassiclinkDnsSupport { get; private set; } = null!;

        /// <summary>
        /// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
        /// </summary>
        [Output("enableDnsHostnames")]
        public Output<bool> EnableDnsHostnames { get; private set; } = null!;

        /// <summary>
        /// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
        /// </summary>
        [Output("enableDnsSupport")]
        public Output<bool?> EnableDnsSupport { get; private set; } = null!;

        /// <summary>
        /// A tenancy option for instances launched into the VPC
        /// </summary>
        [Output("instanceTenancy")]
        public Output<string?> InstanceTenancy { get; private set; } = null!;

        /// <summary>
        /// The association ID for the IPv6 CIDR block.
        /// </summary>
        [Output("ipv6AssociationId")]
        public Output<string> Ipv6AssociationId { get; private set; } = null!;

        /// <summary>
        /// The IPv6 CIDR block.
        /// </summary>
        [Output("ipv6CidrBlock")]
        public Output<string> Ipv6CidrBlock { get; private set; } = null!;

        /// <summary>
        /// The ID of the main route table associated with
        /// this VPC. Note that you can change a VPC's main route table by using an
        /// [`aws.ec2.MainRouteTableAssociation`](https://www.terraform.io/docs/providers/aws/r/main_route_table_assoc.html).
        /// </summary>
        [Output("mainRouteTableId")]
        public Output<string> MainRouteTableId { get; private set; } = null!;

        /// <summary>
        /// The ID of the AWS account that owns the VPC.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Vpc resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vpc(string name, VpcArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/vpc:Vpc", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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

    public sealed class VpcArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Requests an Amazon-provided IPv6 CIDR
        /// block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or
        /// the size of the CIDR block. Default is `false`.
        /// </summary>
        [Input("assignGeneratedIpv6CidrBlock")]
        public Input<bool>? AssignGeneratedIpv6CidrBlock { get; set; }

        /// <summary>
        /// The CIDR block for the VPC.
        /// </summary>
        [Input("cidrBlock", required: true)]
        public Input<string> CidrBlock { get; set; } = null!;

        /// <summary>
        /// A boolean flag to enable/disable ClassicLink
        /// for the VPC. Only valid in regions and accounts that support EC2 Classic.
        /// See the [ClassicLink documentation][1] for more information. Defaults false.
        /// </summary>
        [Input("enableClassiclink")]
        public Input<bool>? EnableClassiclink { get; set; }

        /// <summary>
        /// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
        /// Only valid in regions and accounts that support EC2 Classic.
        /// </summary>
        [Input("enableClassiclinkDnsSupport")]
        public Input<bool>? EnableClassiclinkDnsSupport { get; set; }

        /// <summary>
        /// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
        /// </summary>
        [Input("enableDnsHostnames")]
        public Input<bool>? EnableDnsHostnames { get; set; }

        /// <summary>
        /// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
        /// </summary>
        [Input("enableDnsSupport")]
        public Input<bool>? EnableDnsSupport { get; set; }

        /// <summary>
        /// A tenancy option for instances launched into the VPC
        /// </summary>
        [Input("instanceTenancy")]
        public Input<string>? InstanceTenancy { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public VpcArgs()
        {
        }
    }

    public sealed class VpcState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of VPC
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Requests an Amazon-provided IPv6 CIDR
        /// block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or
        /// the size of the CIDR block. Default is `false`.
        /// </summary>
        [Input("assignGeneratedIpv6CidrBlock")]
        public Input<bool>? AssignGeneratedIpv6CidrBlock { get; set; }

        /// <summary>
        /// The CIDR block for the VPC.
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
        /// A boolean flag to enable/disable ClassicLink
        /// for the VPC. Only valid in regions and accounts that support EC2 Classic.
        /// See the [ClassicLink documentation][1] for more information. Defaults false.
        /// </summary>
        [Input("enableClassiclink")]
        public Input<bool>? EnableClassiclink { get; set; }

        /// <summary>
        /// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
        /// Only valid in regions and accounts that support EC2 Classic.
        /// </summary>
        [Input("enableClassiclinkDnsSupport")]
        public Input<bool>? EnableClassiclinkDnsSupport { get; set; }

        /// <summary>
        /// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
        /// </summary>
        [Input("enableDnsHostnames")]
        public Input<bool>? EnableDnsHostnames { get; set; }

        /// <summary>
        /// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
        /// </summary>
        [Input("enableDnsSupport")]
        public Input<bool>? EnableDnsSupport { get; set; }

        /// <summary>
        /// A tenancy option for instances launched into the VPC
        /// </summary>
        [Input("instanceTenancy")]
        public Input<string>? InstanceTenancy { get; set; }

        /// <summary>
        /// The association ID for the IPv6 CIDR block.
        /// </summary>
        [Input("ipv6AssociationId")]
        public Input<string>? Ipv6AssociationId { get; set; }

        /// <summary>
        /// The IPv6 CIDR block.
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        /// <summary>
        /// The ID of the main route table associated with
        /// this VPC. Note that you can change a VPC's main route table by using an
        /// [`aws.ec2.MainRouteTableAssociation`](https://www.terraform.io/docs/providers/aws/r/main_route_table_assoc.html).
        /// </summary>
        [Input("mainRouteTableId")]
        public Input<string>? MainRouteTableId { get; set; }

        /// <summary>
        /// The ID of the AWS account that owns the VPC.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public VpcState()
        {
        }
    }
}
