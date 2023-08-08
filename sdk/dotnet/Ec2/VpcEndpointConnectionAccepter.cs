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
    /// Provides a resource to accept a pending VPC Endpoint Connection accept request to VPC Endpoint Service.
    /// 
    /// ## Example Usage
    /// ### Accept cross-account request
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleVpcEndpointService = new Aws.Ec2.VpcEndpointService("exampleVpcEndpointService", new()
    ///     {
    ///         AcceptanceRequired = false,
    ///         NetworkLoadBalancerArns = new[]
    ///         {
    ///             aws_lb.Example.Arn,
    ///         },
    ///     });
    /// 
    ///     var exampleVpcEndpoint = new Aws.Ec2.VpcEndpoint("exampleVpcEndpoint", new()
    ///     {
    ///         VpcId = aws_vpc.Test_alternate.Id,
    ///         ServiceName = aws_vpc_endpoint_service.Test.Service_name,
    ///         VpcEndpointType = "Interface",
    ///         PrivateDnsEnabled = false,
    ///         SecurityGroupIds = new[]
    ///         {
    ///             aws_security_group.Test.Id,
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = aws.Alternate,
    ///     });
    /// 
    ///     var exampleVpcEndpointConnectionAccepter = new Aws.Ec2.VpcEndpointConnectionAccepter("exampleVpcEndpointConnectionAccepter", new()
    ///     {
    ///         VpcEndpointServiceId = exampleVpcEndpointService.Id,
    ///         VpcEndpointId = exampleVpcEndpoint.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_vpc_endpoint_connection_accepter.foo
    /// 
    ///  id = "vpce-svc-0f97a19d3fa8220bc_vpce-010601a6db371e263" } Using `pulumi import`, import VPC Endpoint Services using ID of the connection, which is the `VPC Endpoint Service ID` and `VPC Endpoint ID` separated by underscore (`_`).. For exampleconsole % pulumi import aws_vpc_endpoint_connection_accepter.foo vpce-svc-0f97a19d3fa8220bc_vpce-010601a6db371e263
    /// </summary>
    [AwsResourceType("aws:ec2/vpcEndpointConnectionAccepter:VpcEndpointConnectionAccepter")]
    public partial class VpcEndpointConnectionAccepter : global::Pulumi.CustomResource
    {
        /// <summary>
        /// AWS VPC Endpoint ID.
        /// </summary>
        [Output("vpcEndpointId")]
        public Output<string> VpcEndpointId { get; private set; } = null!;

        /// <summary>
        /// AWS VPC Endpoint Service ID.
        /// </summary>
        [Output("vpcEndpointServiceId")]
        public Output<string> VpcEndpointServiceId { get; private set; } = null!;

        /// <summary>
        /// State of the VPC Endpoint.
        /// </summary>
        [Output("vpcEndpointState")]
        public Output<string> VpcEndpointState { get; private set; } = null!;


        /// <summary>
        /// Create a VpcEndpointConnectionAccepter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcEndpointConnectionAccepter(string name, VpcEndpointConnectionAccepterArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/vpcEndpointConnectionAccepter:VpcEndpointConnectionAccepter", name, args ?? new VpcEndpointConnectionAccepterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcEndpointConnectionAccepter(string name, Input<string> id, VpcEndpointConnectionAccepterState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/vpcEndpointConnectionAccepter:VpcEndpointConnectionAccepter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcEndpointConnectionAccepter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcEndpointConnectionAccepter Get(string name, Input<string> id, VpcEndpointConnectionAccepterState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcEndpointConnectionAccepter(name, id, state, options);
        }
    }

    public sealed class VpcEndpointConnectionAccepterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AWS VPC Endpoint ID.
        /// </summary>
        [Input("vpcEndpointId", required: true)]
        public Input<string> VpcEndpointId { get; set; } = null!;

        /// <summary>
        /// AWS VPC Endpoint Service ID.
        /// </summary>
        [Input("vpcEndpointServiceId", required: true)]
        public Input<string> VpcEndpointServiceId { get; set; } = null!;

        public VpcEndpointConnectionAccepterArgs()
        {
        }
        public static new VpcEndpointConnectionAccepterArgs Empty => new VpcEndpointConnectionAccepterArgs();
    }

    public sealed class VpcEndpointConnectionAccepterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AWS VPC Endpoint ID.
        /// </summary>
        [Input("vpcEndpointId")]
        public Input<string>? VpcEndpointId { get; set; }

        /// <summary>
        /// AWS VPC Endpoint Service ID.
        /// </summary>
        [Input("vpcEndpointServiceId")]
        public Input<string>? VpcEndpointServiceId { get; set; }

        /// <summary>
        /// State of the VPC Endpoint.
        /// </summary>
        [Input("vpcEndpointState")]
        public Input<string>? VpcEndpointState { get; set; }

        public VpcEndpointConnectionAccepterState()
        {
        }
        public static new VpcEndpointConnectionAccepterState Empty => new VpcEndpointConnectionAccepterState();
    }
}
