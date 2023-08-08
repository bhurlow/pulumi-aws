// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Vpc
{
    /// <summary>
    /// Manages an inbound (ingress) rule for a security group.
    /// 
    /// When specifying an inbound rule for your security group in a VPC, the configuration must include a source for the traffic.
    /// 
    /// &gt; **NOTE on Security Groups and Security Group Rules:** this provider currently provides a Security Group resource with `ingress` and `egress` rules defined in-line and a Security Group Rule resource which manages one or more `ingress` or
    /// `egress` rules. Both of these resource were added before AWS assigned a [security group rule unique ID](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules.html), and they do not work well in all scenarios using the`description` and `tags` attributes, which rely on the unique ID.
    /// The `aws.vpc.SecurityGroupIngressRule` resource has been added to address these limitations and should be used for all new security group rules.
    /// You should not use the `aws.vpc.SecurityGroupIngressRule` resource in conjunction with an `aws.ec2.SecurityGroup` resource with in-line rules or with `aws.ec2.SecurityGroupRule` resources defined for the same Security Group, as rule conflicts may occur and rules will be overwritten.
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
    ///     var example = new Aws.Vpc.SecurityGroupIngressRule("example", new()
    ///     {
    ///         SecurityGroupId = aws_security_group.Example.Id,
    ///         CidrIpv4 = "10.0.0.0/8",
    ///         FromPort = 80,
    ///         IpProtocol = "tcp",
    ///         ToPort = 80,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_vpc_security_group_ingress_rule.example
    /// 
    ///  id = "sgr-02108b27edd666983" } Using `pulumi import`, import security group ingress rules using the `security_group_rule_id`. For exampleconsole % pulumi import aws_vpc_security_group_ingress_rule.example sgr-02108b27edd666983
    /// </summary>
    [AwsResourceType("aws:vpc/securityGroupIngressRule:SecurityGroupIngressRule")]
    public partial class SecurityGroupIngressRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the security group rule.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The source IPv4 CIDR range.
        /// </summary>
        [Output("cidrIpv4")]
        public Output<string?> CidrIpv4 { get; private set; } = null!;

        /// <summary>
        /// The source IPv6 CIDR range.
        /// </summary>
        [Output("cidrIpv6")]
        public Output<string?> CidrIpv6 { get; private set; } = null!;

        /// <summary>
        /// The security group rule description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type.
        /// </summary>
        [Output("fromPort")]
        public Output<int?> FromPort { get; private set; } = null!;

        /// <summary>
        /// The IP protocol name or number. Use `-1` to specify all protocols. Note that if `ip_protocol` is set to `-1`, it translates to all protocols, all port ranges, and `from_port` and `to_port` values should not be defined.
        /// </summary>
        [Output("ipProtocol")]
        public Output<string> IpProtocol { get; private set; } = null!;

        /// <summary>
        /// The ID of the source prefix list.
        /// </summary>
        [Output("prefixListId")]
        public Output<string?> PrefixListId { get; private set; } = null!;

        /// <summary>
        /// The source security group that is referenced in the rule.
        /// </summary>
        [Output("referencedSecurityGroupId")]
        public Output<string?> ReferencedSecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// The ID of the security group.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// The ID of the security group rule.
        /// </summary>
        [Output("securityGroupRuleId")]
        public Output<string> SecurityGroupRuleId { get; private set; } = null!;

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
        /// The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
        /// </summary>
        [Output("toPort")]
        public Output<int?> ToPort { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityGroupIngressRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityGroupIngressRule(string name, SecurityGroupIngressRuleArgs args, CustomResourceOptions? options = null)
            : base("aws:vpc/securityGroupIngressRule:SecurityGroupIngressRule", name, args ?? new SecurityGroupIngressRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityGroupIngressRule(string name, Input<string> id, SecurityGroupIngressRuleState? state = null, CustomResourceOptions? options = null)
            : base("aws:vpc/securityGroupIngressRule:SecurityGroupIngressRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityGroupIngressRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityGroupIngressRule Get(string name, Input<string> id, SecurityGroupIngressRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityGroupIngressRule(name, id, state, options);
        }
    }

    public sealed class SecurityGroupIngressRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The source IPv4 CIDR range.
        /// </summary>
        [Input("cidrIpv4")]
        public Input<string>? CidrIpv4 { get; set; }

        /// <summary>
        /// The source IPv6 CIDR range.
        /// </summary>
        [Input("cidrIpv6")]
        public Input<string>? CidrIpv6 { get; set; }

        /// <summary>
        /// The security group rule description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type.
        /// </summary>
        [Input("fromPort")]
        public Input<int>? FromPort { get; set; }

        /// <summary>
        /// The IP protocol name or number. Use `-1` to specify all protocols. Note that if `ip_protocol` is set to `-1`, it translates to all protocols, all port ranges, and `from_port` and `to_port` values should not be defined.
        /// </summary>
        [Input("ipProtocol", required: true)]
        public Input<string> IpProtocol { get; set; } = null!;

        /// <summary>
        /// The ID of the source prefix list.
        /// </summary>
        [Input("prefixListId")]
        public Input<string>? PrefixListId { get; set; }

        /// <summary>
        /// The source security group that is referenced in the rule.
        /// </summary>
        [Input("referencedSecurityGroupId")]
        public Input<string>? ReferencedSecurityGroupId { get; set; }

        /// <summary>
        /// The ID of the security group.
        /// </summary>
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

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
        /// The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
        /// </summary>
        [Input("toPort")]
        public Input<int>? ToPort { get; set; }

        public SecurityGroupIngressRuleArgs()
        {
        }
        public static new SecurityGroupIngressRuleArgs Empty => new SecurityGroupIngressRuleArgs();
    }

    public sealed class SecurityGroupIngressRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the security group rule.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The source IPv4 CIDR range.
        /// </summary>
        [Input("cidrIpv4")]
        public Input<string>? CidrIpv4 { get; set; }

        /// <summary>
        /// The source IPv6 CIDR range.
        /// </summary>
        [Input("cidrIpv6")]
        public Input<string>? CidrIpv6 { get; set; }

        /// <summary>
        /// The security group rule description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type.
        /// </summary>
        [Input("fromPort")]
        public Input<int>? FromPort { get; set; }

        /// <summary>
        /// The IP protocol name or number. Use `-1` to specify all protocols. Note that if `ip_protocol` is set to `-1`, it translates to all protocols, all port ranges, and `from_port` and `to_port` values should not be defined.
        /// </summary>
        [Input("ipProtocol")]
        public Input<string>? IpProtocol { get; set; }

        /// <summary>
        /// The ID of the source prefix list.
        /// </summary>
        [Input("prefixListId")]
        public Input<string>? PrefixListId { get; set; }

        /// <summary>
        /// The source security group that is referenced in the rule.
        /// </summary>
        [Input("referencedSecurityGroupId")]
        public Input<string>? ReferencedSecurityGroupId { get; set; }

        /// <summary>
        /// The ID of the security group.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// The ID of the security group rule.
        /// </summary>
        [Input("securityGroupRuleId")]
        public Input<string>? SecurityGroupRuleId { get; set; }

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
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
        /// </summary>
        [Input("toPort")]
        public Input<int>? ToPort { get; set; }

        public SecurityGroupIngressRuleState()
        {
        }
        public static new SecurityGroupIngressRuleState Empty => new SecurityGroupIngressRuleState();
    }
}
