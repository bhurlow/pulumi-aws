// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.S3Outposts
{
    /// <summary>
    /// Provides a resource to manage an S3 Outposts Endpoint.
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
    ///     var example = new Aws.S3Outposts.Endpoint("example", new()
    ///     {
    ///         OutpostId = data.Aws_outposts_outpost.Example.Id,
    ///         SecurityGroupId = aws_security_group.Example.Id,
    ///         SubnetId = aws_subnet.Example.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_s3outposts_endpoint.example
    /// 
    ///  id = "arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-12345678/endpoint/0123456789abcdef,sg-12345678,subnet-12345678" } Using `pulumi import`, import S3 Outposts Endpoints using Amazon Resource Name (ARN), EC2 Security Group identifier, and EC2 Subnet identifier, separated by commas (`,`). For exampleconsole % pulumi import aws_s3outposts_endpoint.example arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-12345678/endpoint/0123456789abcdef,sg-12345678,subnet-12345678
    /// </summary>
    [AwsResourceType("aws:s3outposts/endpoint:Endpoint")]
    public partial class Endpoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Type of access for the network connectivity. Valid values are `Private` or `CustomerOwnedIp`.
        /// </summary>
        [Output("accessType")]
        public Output<string> AccessType { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the endpoint.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// VPC CIDR block of the endpoint.
        /// </summary>
        [Output("cidrBlock")]
        public Output<string> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// UTC creation time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// The ID of a Customer Owned IP Pool. For more on customer owned IP addresses see the [User Guide](https://docs.aws.amazon.com/outposts/latest/userguide/local-rack.html#local-gateway-subnet).
        /// </summary>
        [Output("customerOwnedIpv4Pool")]
        public Output<string?> CustomerOwnedIpv4Pool { get; private set; } = null!;

        /// <summary>
        /// Set of nested attributes for associated Elastic Network Interfaces (ENIs).
        /// </summary>
        [Output("networkInterfaces")]
        public Output<ImmutableArray<Outputs.EndpointNetworkInterface>> NetworkInterfaces { get; private set; } = null!;

        /// <summary>
        /// Identifier of the Outpost to contain this endpoint.
        /// </summary>
        [Output("outpostId")]
        public Output<string> OutpostId { get; private set; } = null!;

        /// <summary>
        /// Identifier of the EC2 Security Group.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// Identifier of the EC2 Subnet.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;


        /// <summary>
        /// Create a Endpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Endpoint(string name, EndpointArgs args, CustomResourceOptions? options = null)
            : base("aws:s3outposts/endpoint:Endpoint", name, args ?? new EndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Endpoint(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
            : base("aws:s3outposts/endpoint:Endpoint", name, state, MakeResourceOptions(options, id))
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

    public sealed class EndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of access for the network connectivity. Valid values are `Private` or `CustomerOwnedIp`.
        /// </summary>
        [Input("accessType")]
        public Input<string>? AccessType { get; set; }

        /// <summary>
        /// The ID of a Customer Owned IP Pool. For more on customer owned IP addresses see the [User Guide](https://docs.aws.amazon.com/outposts/latest/userguide/local-rack.html#local-gateway-subnet).
        /// </summary>
        [Input("customerOwnedIpv4Pool")]
        public Input<string>? CustomerOwnedIpv4Pool { get; set; }

        /// <summary>
        /// Identifier of the Outpost to contain this endpoint.
        /// </summary>
        [Input("outpostId", required: true)]
        public Input<string> OutpostId { get; set; } = null!;

        /// <summary>
        /// Identifier of the EC2 Security Group.
        /// </summary>
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        /// <summary>
        /// Identifier of the EC2 Subnet.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public EndpointArgs()
        {
        }
        public static new EndpointArgs Empty => new EndpointArgs();
    }

    public sealed class EndpointState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of access for the network connectivity. Valid values are `Private` or `CustomerOwnedIp`.
        /// </summary>
        [Input("accessType")]
        public Input<string>? AccessType { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the endpoint.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// VPC CIDR block of the endpoint.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// UTC creation time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
        /// </summary>
        [Input("creationTime")]
        public Input<string>? CreationTime { get; set; }

        /// <summary>
        /// The ID of a Customer Owned IP Pool. For more on customer owned IP addresses see the [User Guide](https://docs.aws.amazon.com/outposts/latest/userguide/local-rack.html#local-gateway-subnet).
        /// </summary>
        [Input("customerOwnedIpv4Pool")]
        public Input<string>? CustomerOwnedIpv4Pool { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.EndpointNetworkInterfaceGetArgs>? _networkInterfaces;

        /// <summary>
        /// Set of nested attributes for associated Elastic Network Interfaces (ENIs).
        /// </summary>
        public InputList<Inputs.EndpointNetworkInterfaceGetArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.EndpointNetworkInterfaceGetArgs>());
            set => _networkInterfaces = value;
        }

        /// <summary>
        /// Identifier of the Outpost to contain this endpoint.
        /// </summary>
        [Input("outpostId")]
        public Input<string>? OutpostId { get; set; }

        /// <summary>
        /// Identifier of the EC2 Security Group.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// Identifier of the EC2 Subnet.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public EndpointState()
        {
        }
        public static new EndpointState Empty => new EndpointState();
    }
}
