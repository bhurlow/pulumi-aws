// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static class GetRouteTable
    {
        /// <summary>
        /// `aws.ec2.RouteTable` provides details about a specific Route Table.
        /// 
        /// This resource can prove useful when a module accepts a Subnet ID as an input variable and needs to, for example, add a route in the Route Table.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// The following example shows how one might accept a Route Table ID as a variable and use this data source to obtain the data necessary to create a route.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var config = new Config();
        ///         var subnetId = config.RequireObject&lt;dynamic&gt;("subnetId");
        ///         var selected = Output.Create(Aws.Ec2.GetRouteTable.InvokeAsync(new Aws.Ec2.GetRouteTableArgs
        ///         {
        ///             SubnetId = subnetId,
        ///         }));
        ///         var route = new Aws.Ec2.Route("route", new Aws.Ec2.RouteArgs
        ///         {
        ///             RouteTableId = selected.Apply(selected =&gt; selected.Id),
        ///             DestinationCidrBlock = "10.0.1.0/22",
        ///             VpcPeeringConnectionId = "pcx-45ff3dc1",
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRouteTableResult> InvokeAsync(GetRouteTableArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRouteTableResult>("aws:ec2/getRouteTable:getRouteTable", args ?? new GetRouteTableArgs(), options.WithVersion());
    }


    public sealed class GetRouteTableArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetRouteTableFilterArgs>? _filters;

        /// <summary>
        /// Configuration block. Detailed below.
        /// </summary>
        public List<Inputs.GetRouteTableFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetRouteTableFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// ID of an Internet Gateway or Virtual Private Gateway which is connected to the Route Table (not exported if not passed as a parameter).
        /// </summary>
        [Input("gatewayId")]
        public string? GatewayId { get; set; }

        /// <summary>
        /// ID of the specific Route Table to retrieve.
        /// </summary>
        [Input("routeTableId")]
        public string? RouteTableId { get; set; }

        /// <summary>
        /// ID of a Subnet which is connected to the Route Table (not exported if not passed as a parameter).
        /// </summary>
        [Input("subnetId")]
        public string? SubnetId { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// Map of tags, each pair of which must exactly match a pair on the desired Route Table.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        /// <summary>
        /// ID of the VPC that the desired Route Table belongs to.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetRouteTableArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRouteTableResult
    {
        /// <summary>
        /// ARN of the route table.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// List of associations with attributes detailed below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouteTableAssociationResult> Associations;
        public readonly ImmutableArray<Outputs.GetRouteTableFilterResult> Filters;
        /// <summary>
        /// Gateway ID. Only set when associated with an Internet Gateway or Virtual Private Gateway.
        /// </summary>
        public readonly string GatewayId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// ID of the AWS account that owns the route table.
        /// </summary>
        public readonly string OwnerId;
        /// <summary>
        /// Route Table ID.
        /// </summary>
        public readonly string RouteTableId;
        /// <summary>
        /// List of routes with attributes detailed below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouteTableRouteResult> Routes;
        /// <summary>
        /// Subnet ID. Only set when associated with a subnet.
        /// </summary>
        public readonly string SubnetId;
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly string VpcId;

        [OutputConstructor]
        private GetRouteTableResult(
            string arn,

            ImmutableArray<Outputs.GetRouteTableAssociationResult> associations,

            ImmutableArray<Outputs.GetRouteTableFilterResult> filters,

            string gatewayId,

            string id,

            string ownerId,

            string routeTableId,

            ImmutableArray<Outputs.GetRouteTableRouteResult> routes,

            string subnetId,

            ImmutableDictionary<string, string> tags,

            string vpcId)
        {
            Arn = arn;
            Associations = associations;
            Filters = filters;
            GatewayId = gatewayId;
            Id = id;
            OwnerId = ownerId;
            RouteTableId = routeTableId;
            Routes = routes;
            SubnetId = subnetId;
            Tags = tags;
            VpcId = vpcId;
        }
    }
}
