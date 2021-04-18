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
    /// Provides a security group rule resource. Represents a single `ingress` or
    /// `egress` group rule, which can be added to external Security Groups.
    /// 
    /// &gt; **NOTE on Security Groups and Security Group Rules:** This provider currently
    /// provides both a standalone Security Group Rule resource (a single `ingress` or
    /// `egress` rule), and a Security Group resource with `ingress` and `egress` rules
    /// defined in-line. At this time you cannot use a Security Group with in-line rules
    /// in conjunction with any Security Group Rule resources. Doing so will cause
    /// a conflict of rule settings and will overwrite rules.
    /// 
    /// &gt; **NOTE:** Setting `protocol = "all"` or `protocol = -1` with `from_port` and `to_port` will result in the EC2 API creating a security group rule with all ports open. This API behavior cannot be controlled by this provider and may generate warnings in the future.
    /// 
    /// &gt; **NOTE:** Referencing Security Groups across VPC peering has certain restrictions. More information is available in the [VPC Peering User Guide](https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-security-groups.html).
    /// 
    /// ## Example Usage
    /// 
    /// Basic usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Ec2.SecurityGroupRule("example", new Aws.Ec2.SecurityGroupRuleArgs
    ///         {
    ///             Type = "ingress",
    ///             FromPort = 0,
    ///             ToPort = 65535,
    ///             Protocol = "tcp",
    ///             CidrBlocks = 
    ///             {
    ///                 aws_vpc.Example.Cidr_block,
    ///             },
    ///             Ipv6CidrBlocks = 
    ///             {
    ///                 aws_vpc.Example.Ipv6_cidr_block,
    ///             },
    ///             SecurityGroupId = "sg-123456",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Usage With Prefix List IDs
    /// 
    /// Prefix Lists are either managed by AWS internally, or created by the customer using a
    /// Managed Prefix List resource. Prefix Lists provided by
    /// AWS are associated with a prefix list name, or service name, that is linked to a specific region.
    /// 
    /// Prefix list IDs are exported on VPC Endpoints, so you can use this format:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // ...
    ///         var myEndpoint = new Aws.Ec2.VpcEndpoint("myEndpoint", new Aws.Ec2.VpcEndpointArgs
    ///         {
    ///         });
    ///         // ...
    ///         var allowAll = new Aws.Ec2.SecurityGroupRule("allowAll", new Aws.Ec2.SecurityGroupRuleArgs
    ///         {
    ///             Type = "egress",
    ///             ToPort = 0,
    ///             Protocol = "-1",
    ///             PrefixListIds = 
    ///             {
    ///                 myEndpoint.PrefixListId,
    ///             },
    ///             FromPort = 0,
    ///             SecurityGroupId = "sg-123456",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// You can also find a specific Prefix List using the `aws.ec2.getPrefixList` data source.
    /// 
    /// ## Import
    /// 
    /// Security Group Rules can be imported using the `security_group_id`, `type`, `protocol`, `from_port`, `to_port`, and source(s)/destination(s) (e.g. `cidr_block`) separated by underscores (`_`). All parts are required. Not all rule permissions (e.g., not all of a rule's CIDR blocks) need to be imported for this provider to manage rule permissions. However, importing some of a rule's permissions but not others, and then making changes to the rule will result in the creation of an additional rule to capture the updated permissions. Rule permissions that were not imported are left intact in the original rule. Import an ingress rule in security group `sg-6e616f6d69` for TCP port 8000 with an IPv4 destination CIDR of `10.0.3.0/24`console
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2/securityGroupRule:SecurityGroupRule ingress sg-6e616f6d69_ingress_tcp_8000_8000_10.0.3.0/24
    /// ```
    /// 
    ///  Import a rule with various IPv4 and IPv6 source CIDR blocksconsole
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2/securityGroupRule:SecurityGroupRule ingress sg-4973616163_ingress_tcp_100_121_10.1.0.0/16_2001:db8::/48_10.2.0.0/16_2002:db8::/48
    /// ```
    /// 
    ///  Import a rule, applicable to all ports, with a protocol other than TCP/UDP/ICMP/ALL, e.g., Multicast Transport Protocol (MTP), using the IANA protocol number, e.g., 92. console
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2/securityGroupRule:SecurityGroupRule ingress sg-6777656e646f6c796e_ingress_92_0_65536_10.0.3.0/24_10.0.4.0/24
    /// ```
    /// 
    ///  Import an egress rule with a prefix list ID destinationconsole
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2/securityGroupRule:SecurityGroupRule egress sg-62726f6479_egress_tcp_8000_8000_pl-6469726b
    /// ```
    /// 
    ///  Import a rule applicable to all protocols and ports with a security group sourceconsole
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2/securityGroupRule:SecurityGroupRule ingress_rule sg-7472697374616e_ingress_all_0_65536_sg-6176657279
    /// ```
    /// 
    ///  Import a rule that has itself and an IPv6 CIDR block as sourcesconsole
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2/securityGroupRule:SecurityGroupRule rule_name sg-656c65616e6f72_ingress_tcp_80_80_self_2001:db8::/48
    /// ```
    /// </summary>
    [AwsResourceType("aws:ec2/securityGroupRule:SecurityGroupRule")]
    public partial class SecurityGroupRule : Pulumi.CustomResource
    {
        /// <summary>
        /// List of CIDR blocks. Cannot be specified with `source_security_group_id`.
        /// </summary>
        [Output("cidrBlocks")]
        public Output<ImmutableArray<string>> CidrBlocks { get; private set; } = null!;

        /// <summary>
        /// Description of the rule.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Start port (or ICMP type number if protocol is "icmp" or "icmpv6").
        /// </summary>
        [Output("fromPort")]
        public Output<int> FromPort { get; private set; } = null!;

        /// <summary>
        /// List of IPv6 CIDR blocks.
        /// </summary>
        [Output("ipv6CidrBlocks")]
        public Output<ImmutableArray<string>> Ipv6CidrBlocks { get; private set; } = null!;

        /// <summary>
        /// List of Prefix List IDs.
        /// </summary>
        [Output("prefixListIds")]
        public Output<ImmutableArray<string>> PrefixListIds { get; private set; } = null!;

        /// <summary>
        /// Protocol. If not icmp, icmpv6, tcp, udp, or all use the [protocol number](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Security group to apply this rule to.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// Whether the security group itself will be added as a source to this ingress rule. Cannot be specified with `source_security_group_id`.
        /// </summary>
        [Output("self")]
        public Output<bool?> Self { get; private set; } = null!;

        /// <summary>
        /// Security group id to allow access to/from, depending on the `type`. Cannot be specified with `cidr_blocks` and `self`.
        /// </summary>
        [Output("sourceSecurityGroupId")]
        public Output<string> SourceSecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// End port (or ICMP code if protocol is "icmp").
        /// </summary>
        [Output("toPort")]
        public Output<int> ToPort { get; private set; } = null!;

        /// <summary>
        /// Type of rule being created. Valid options are `ingress` (inbound)
        /// or `egress` (outbound).
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityGroupRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityGroupRule(string name, SecurityGroupRuleArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/securityGroupRule:SecurityGroupRule", name, args ?? new SecurityGroupRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityGroupRule(string name, Input<string> id, SecurityGroupRuleState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/securityGroupRule:SecurityGroupRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityGroupRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityGroupRule Get(string name, Input<string> id, SecurityGroupRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityGroupRule(name, id, state, options);
        }
    }

    public sealed class SecurityGroupRuleArgs : Pulumi.ResourceArgs
    {
        [Input("cidrBlocks")]
        private InputList<string>? _cidrBlocks;

        /// <summary>
        /// List of CIDR blocks. Cannot be specified with `source_security_group_id`.
        /// </summary>
        public InputList<string> CidrBlocks
        {
            get => _cidrBlocks ?? (_cidrBlocks = new InputList<string>());
            set => _cidrBlocks = value;
        }

        /// <summary>
        /// Description of the rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Start port (or ICMP type number if protocol is "icmp" or "icmpv6").
        /// </summary>
        [Input("fromPort", required: true)]
        public Input<int> FromPort { get; set; } = null!;

        [Input("ipv6CidrBlocks")]
        private InputList<string>? _ipv6CidrBlocks;

        /// <summary>
        /// List of IPv6 CIDR blocks.
        /// </summary>
        public InputList<string> Ipv6CidrBlocks
        {
            get => _ipv6CidrBlocks ?? (_ipv6CidrBlocks = new InputList<string>());
            set => _ipv6CidrBlocks = value;
        }

        [Input("prefixListIds")]
        private InputList<string>? _prefixListIds;

        /// <summary>
        /// List of Prefix List IDs.
        /// </summary>
        public InputList<string> PrefixListIds
        {
            get => _prefixListIds ?? (_prefixListIds = new InputList<string>());
            set => _prefixListIds = value;
        }

        /// <summary>
        /// Protocol. If not icmp, icmpv6, tcp, udp, or all use the [protocol number](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
        /// </summary>
        [Input("protocol", required: true)]
        public InputUnion<string, Pulumi.Aws.Ec2.ProtocolType> Protocol { get; set; } = null!;

        /// <summary>
        /// Security group to apply this rule to.
        /// </summary>
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        /// <summary>
        /// Whether the security group itself will be added as a source to this ingress rule. Cannot be specified with `source_security_group_id`.
        /// </summary>
        [Input("self")]
        public Input<bool>? Self { get; set; }

        /// <summary>
        /// Security group id to allow access to/from, depending on the `type`. Cannot be specified with `cidr_blocks` and `self`.
        /// </summary>
        [Input("sourceSecurityGroupId")]
        public Input<string>? SourceSecurityGroupId { get; set; }

        /// <summary>
        /// End port (or ICMP code if protocol is "icmp").
        /// </summary>
        [Input("toPort", required: true)]
        public Input<int> ToPort { get; set; } = null!;

        /// <summary>
        /// Type of rule being created. Valid options are `ingress` (inbound)
        /// or `egress` (outbound).
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public SecurityGroupRuleArgs()
        {
        }
    }

    public sealed class SecurityGroupRuleState : Pulumi.ResourceArgs
    {
        [Input("cidrBlocks")]
        private InputList<string>? _cidrBlocks;

        /// <summary>
        /// List of CIDR blocks. Cannot be specified with `source_security_group_id`.
        /// </summary>
        public InputList<string> CidrBlocks
        {
            get => _cidrBlocks ?? (_cidrBlocks = new InputList<string>());
            set => _cidrBlocks = value;
        }

        /// <summary>
        /// Description of the rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Start port (or ICMP type number if protocol is "icmp" or "icmpv6").
        /// </summary>
        [Input("fromPort")]
        public Input<int>? FromPort { get; set; }

        [Input("ipv6CidrBlocks")]
        private InputList<string>? _ipv6CidrBlocks;

        /// <summary>
        /// List of IPv6 CIDR blocks.
        /// </summary>
        public InputList<string> Ipv6CidrBlocks
        {
            get => _ipv6CidrBlocks ?? (_ipv6CidrBlocks = new InputList<string>());
            set => _ipv6CidrBlocks = value;
        }

        [Input("prefixListIds")]
        private InputList<string>? _prefixListIds;

        /// <summary>
        /// List of Prefix List IDs.
        /// </summary>
        public InputList<string> PrefixListIds
        {
            get => _prefixListIds ?? (_prefixListIds = new InputList<string>());
            set => _prefixListIds = value;
        }

        /// <summary>
        /// Protocol. If not icmp, icmpv6, tcp, udp, or all use the [protocol number](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
        /// </summary>
        [Input("protocol")]
        public InputUnion<string, Pulumi.Aws.Ec2.ProtocolType>? Protocol { get; set; }

        /// <summary>
        /// Security group to apply this rule to.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// Whether the security group itself will be added as a source to this ingress rule. Cannot be specified with `source_security_group_id`.
        /// </summary>
        [Input("self")]
        public Input<bool>? Self { get; set; }

        /// <summary>
        /// Security group id to allow access to/from, depending on the `type`. Cannot be specified with `cidr_blocks` and `self`.
        /// </summary>
        [Input("sourceSecurityGroupId")]
        public Input<string>? SourceSecurityGroupId { get; set; }

        /// <summary>
        /// End port (or ICMP code if protocol is "icmp").
        /// </summary>
        [Input("toPort")]
        public Input<int>? ToPort { get; set; }

        /// <summary>
        /// Type of rule being created. Valid options are `ingress` (inbound)
        /// or `egress` (outbound).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public SecurityGroupRuleState()
        {
        }
    }
}
