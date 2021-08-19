// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2ClientVpn
{
    /// <summary>
    /// Provides an AWS Client VPN endpoint for OpenVPN clients. For more information on usage, please see the
    /// [AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Ec2ClientVpn.Endpoint("example", new Aws.Ec2ClientVpn.EndpointArgs
    ///         {
    ///             Description = "clientvpn-example",
    ///             ServerCertificateArn = aws_acm_certificate.Cert.Arn,
    ///             ClientCidrBlock = "10.0.0.0/16",
    ///             AuthenticationOptions = 
    ///             {
    ///                 new Aws.Ec2ClientVpn.Inputs.EndpointAuthenticationOptionArgs
    ///                 {
    ///                     Type = "certificate-authentication",
    ///                     RootCertificateChainArn = aws_acm_certificate.Root_cert.Arn,
    ///                 },
    ///             },
    ///             ConnectionLogOptions = new Aws.Ec2ClientVpn.Inputs.EndpointConnectionLogOptionsArgs
    ///             {
    ///                 Enabled = true,
    ///                 CloudwatchLogGroup = aws_cloudwatch_log_group.Lg.Name,
    ///                 CloudwatchLogStream = aws_cloudwatch_log_stream.Ls.Name,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// AWS Client VPN endpoints can be imported using the `id` value found via `aws ec2 describe-client-vpn-endpoints`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2clientvpn/endpoint:Endpoint example cvpn-endpoint-0ac3a1abbccddd666
    /// ```
    /// </summary>
    [AwsResourceType("aws:ec2clientvpn/endpoint:Endpoint")]
    public partial class Endpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the Client VPN endpoint.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Information about the authentication method to be used to authenticate clients.
        /// </summary>
        [Output("authenticationOptions")]
        public Output<ImmutableArray<Outputs.EndpointAuthenticationOption>> AuthenticationOptions { get; private set; } = null!;

        /// <summary>
        /// The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
        /// </summary>
        [Output("clientCidrBlock")]
        public Output<string> ClientCidrBlock { get; private set; } = null!;

        /// <summary>
        /// Information about the client connection logging options.
        /// </summary>
        [Output("connectionLogOptions")]
        public Output<Outputs.EndpointConnectionLogOptions> ConnectionLogOptions { get; private set; } = null!;

        /// <summary>
        /// Name of the repository.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The DNS name to be used by clients when establishing their VPN session.
        /// </summary>
        [Output("dnsName")]
        public Output<string> DnsName { get; private set; } = null!;

        /// <summary>
        /// Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
        /// </summary>
        [Output("dnsServers")]
        public Output<ImmutableArray<string>> DnsServers { get; private set; } = null!;

        /// <summary>
        /// The ARN of the ACM server certificate.
        /// </summary>
        [Output("serverCertificateArn")]
        public Output<string> ServerCertificateArn { get; private set; } = null!;

        /// <summary>
        /// Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
        /// </summary>
        [Output("splitTunnel")]
        public Output<bool?> SplitTunnel { get; private set; } = null!;

        /// <summary>
        /// The current state of the Client VPN endpoint.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The transport protocol to be used by the VPN session. Default value is `udp`.
        /// </summary>
        [Output("transportProtocol")]
        public Output<string?> TransportProtocol { get; private set; } = null!;


        /// <summary>
        /// Create a Endpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Endpoint(string name, EndpointArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2clientvpn/endpoint:Endpoint", name, args ?? new EndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Endpoint(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2clientvpn/endpoint:Endpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Endpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Endpoint Get(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new Endpoint(name, id, state, options);
        }
    }

    public sealed class EndpointArgs : Pulumi.ResourceArgs
    {
        [Input("authenticationOptions", required: true)]
        private InputList<Inputs.EndpointAuthenticationOptionArgs>? _authenticationOptions;

        /// <summary>
        /// Information about the authentication method to be used to authenticate clients.
        /// </summary>
        public InputList<Inputs.EndpointAuthenticationOptionArgs> AuthenticationOptions
        {
            get => _authenticationOptions ?? (_authenticationOptions = new InputList<Inputs.EndpointAuthenticationOptionArgs>());
            set => _authenticationOptions = value;
        }

        /// <summary>
        /// The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
        /// </summary>
        [Input("clientCidrBlock", required: true)]
        public Input<string> ClientCidrBlock { get; set; } = null!;

        /// <summary>
        /// Information about the client connection logging options.
        /// </summary>
        [Input("connectionLogOptions", required: true)]
        public Input<Inputs.EndpointConnectionLogOptionsArgs> ConnectionLogOptions { get; set; } = null!;

        /// <summary>
        /// Name of the repository.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("dnsServers")]
        private InputList<string>? _dnsServers;

        /// <summary>
        /// Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
        /// </summary>
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        /// <summary>
        /// The ARN of the ACM server certificate.
        /// </summary>
        [Input("serverCertificateArn", required: true)]
        public Input<string> ServerCertificateArn { get; set; } = null!;

        /// <summary>
        /// Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
        /// </summary>
        [Input("splitTunnel")]
        public Input<bool>? SplitTunnel { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The transport protocol to be used by the VPN session. Default value is `udp`.
        /// </summary>
        [Input("transportProtocol")]
        public Input<string>? TransportProtocol { get; set; }

        public EndpointArgs()
        {
        }
    }

    public sealed class EndpointState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Client VPN endpoint.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("authenticationOptions")]
        private InputList<Inputs.EndpointAuthenticationOptionGetArgs>? _authenticationOptions;

        /// <summary>
        /// Information about the authentication method to be used to authenticate clients.
        /// </summary>
        public InputList<Inputs.EndpointAuthenticationOptionGetArgs> AuthenticationOptions
        {
            get => _authenticationOptions ?? (_authenticationOptions = new InputList<Inputs.EndpointAuthenticationOptionGetArgs>());
            set => _authenticationOptions = value;
        }

        /// <summary>
        /// The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
        /// </summary>
        [Input("clientCidrBlock")]
        public Input<string>? ClientCidrBlock { get; set; }

        /// <summary>
        /// Information about the client connection logging options.
        /// </summary>
        [Input("connectionLogOptions")]
        public Input<Inputs.EndpointConnectionLogOptionsGetArgs>? ConnectionLogOptions { get; set; }

        /// <summary>
        /// Name of the repository.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The DNS name to be used by clients when establishing their VPN session.
        /// </summary>
        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        [Input("dnsServers")]
        private InputList<string>? _dnsServers;

        /// <summary>
        /// Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
        /// </summary>
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        /// <summary>
        /// The ARN of the ACM server certificate.
        /// </summary>
        [Input("serverCertificateArn")]
        public Input<string>? ServerCertificateArn { get; set; }

        /// <summary>
        /// Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
        /// </summary>
        [Input("splitTunnel")]
        public Input<bool>? SplitTunnel { get; set; }

        /// <summary>
        /// The current state of the Client VPN endpoint.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// The transport protocol to be used by the VPN session. Default value is `udp`.
        /// </summary>
        [Input("transportProtocol")]
        public Input<string>? TransportProtocol { get; set; }

        public EndpointState()
        {
        }
    }
}
