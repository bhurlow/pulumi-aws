// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static class GetVpcEndpoint
    {
        /// <summary>
        /// The VPC Endpoint data source provides details about
        /// a specific VPC endpoint.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var s3 = Aws.Ec2.GetVpcEndpoint.Invoke(new()
        ///     {
        ///         VpcId = aws_vpc.Foo.Id,
        ///         ServiceName = "com.amazonaws.us-west-2.s3",
        ///     });
        /// 
        ///     var privateS3 = new Aws.Ec2.VpcEndpointRouteTableAssociation("privateS3", new()
        ///     {
        ///         VpcEndpointId = s3.Apply(getVpcEndpointResult =&gt; getVpcEndpointResult.Id),
        ///         RouteTableId = aws_route_table.Private.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVpcEndpointResult> InvokeAsync(GetVpcEndpointArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcEndpointResult>("aws:ec2/getVpcEndpoint:getVpcEndpoint", args ?? new GetVpcEndpointArgs(), options.WithDefaults());

        /// <summary>
        /// The VPC Endpoint data source provides details about
        /// a specific VPC endpoint.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var s3 = Aws.Ec2.GetVpcEndpoint.Invoke(new()
        ///     {
        ///         VpcId = aws_vpc.Foo.Id,
        ///         ServiceName = "com.amazonaws.us-west-2.s3",
        ///     });
        /// 
        ///     var privateS3 = new Aws.Ec2.VpcEndpointRouteTableAssociation("privateS3", new()
        ///     {
        ///         VpcEndpointId = s3.Apply(getVpcEndpointResult =&gt; getVpcEndpointResult.Id),
        ///         RouteTableId = aws_route_table.Private.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetVpcEndpointResult> Invoke(GetVpcEndpointInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcEndpointResult>("aws:ec2/getVpcEndpoint:getVpcEndpoint", args ?? new GetVpcEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcEndpointArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetVpcEndpointFilterArgs>? _filters;

        /// <summary>
        /// Custom filter block as described below.
        /// </summary>
        public List<Inputs.GetVpcEndpointFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetVpcEndpointFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// ID of the specific VPC Endpoint to retrieve.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// Service name of the specific VPC Endpoint to retrieve. For AWS services the service name is usually in the form `com.amazonaws.&lt;region&gt;.&lt;service&gt;` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.&lt;region&gt;.notebook`).
        /// </summary>
        [Input("serviceName")]
        public string? ServiceName { get; set; }

        /// <summary>
        /// State of the specific VPC Endpoint to retrieve.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// Map of tags, each pair of which must exactly match
        /// a pair on the specific VPC Endpoint to retrieve.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        /// <summary>
        /// ID of the VPC in which the specific VPC Endpoint is used.
        /// 
        /// More complex filters can be expressed using one or more `filter` sub-blocks,
        /// which take the following arguments:
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetVpcEndpointArgs()
        {
        }
        public static new GetVpcEndpointArgs Empty => new GetVpcEndpointArgs();
    }

    public sealed class GetVpcEndpointInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetVpcEndpointFilterInputArgs>? _filters;

        /// <summary>
        /// Custom filter block as described below.
        /// </summary>
        public InputList<Inputs.GetVpcEndpointFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetVpcEndpointFilterInputArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// ID of the specific VPC Endpoint to retrieve.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Service name of the specific VPC Endpoint to retrieve. For AWS services the service name is usually in the form `com.amazonaws.&lt;region&gt;.&lt;service&gt;` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.&lt;region&gt;.notebook`).
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// State of the specific VPC Endpoint to retrieve.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags, each pair of which must exactly match
        /// a pair on the specific VPC Endpoint to retrieve.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// ID of the VPC in which the specific VPC Endpoint is used.
        /// 
        /// More complex filters can be expressed using one or more `filter` sub-blocks,
        /// which take the following arguments:
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GetVpcEndpointInvokeArgs()
        {
        }
        public static new GetVpcEndpointInvokeArgs Empty => new GetVpcEndpointInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcEndpointResult
    {
        /// <summary>
        /// ARN of the VPC endpoint.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// List of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
        /// </summary>
        public readonly ImmutableArray<string> CidrBlocks;
        /// <summary>
        /// DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS entry blocks are documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVpcEndpointDnsEntryResult> DnsEntries;
        /// <summary>
        /// DNS options for the VPC Endpoint. DNS options blocks are documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVpcEndpointDnsOptionResult> DnsOptions;
        public readonly ImmutableArray<Outputs.GetVpcEndpointFilterResult> Filters;
        public readonly string Id;
        public readonly string IpAddressType;
        /// <summary>
        /// One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
        /// </summary>
        public readonly ImmutableArray<string> NetworkInterfaceIds;
        /// <summary>
        /// ID of the AWS account that owns the VPC endpoint.
        /// </summary>
        public readonly string OwnerId;
        /// <summary>
        /// Policy document associated with the VPC Endpoint. Applicable for endpoints of type `Gateway`.
        /// </summary>
        public readonly string Policy;
        /// <summary>
        /// Prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
        /// </summary>
        public readonly string PrefixListId;
        /// <summary>
        /// Whether or not the VPC is associated with a private hosted zone - `true` or `false`. Applicable for endpoints of type `Interface`.
        /// </summary>
        public readonly bool PrivateDnsEnabled;
        /// <summary>
        /// Whether or not the VPC Endpoint is being managed by its service - `true` or `false`.
        /// </summary>
        public readonly bool RequesterManaged;
        /// <summary>
        /// One or more route tables associated with the VPC Endpoint. Applicable for endpoints of type `Gateway`.
        /// </summary>
        public readonly ImmutableArray<string> RouteTableIds;
        /// <summary>
        /// One or more security groups associated with the network interfaces. Applicable for endpoints of type `Interface`.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        public readonly string ServiceName;
        public readonly string State;
        /// <summary>
        /// One or more subnets in which the VPC Endpoint is located. Applicable for endpoints of type `Interface`.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// VPC Endpoint type, `Gateway` or `Interface`.
        /// </summary>
        public readonly string VpcEndpointType;
        public readonly string VpcId;

        [OutputConstructor]
        private GetVpcEndpointResult(
            string arn,

            ImmutableArray<string> cidrBlocks,

            ImmutableArray<Outputs.GetVpcEndpointDnsEntryResult> dnsEntries,

            ImmutableArray<Outputs.GetVpcEndpointDnsOptionResult> dnsOptions,

            ImmutableArray<Outputs.GetVpcEndpointFilterResult> filters,

            string id,

            string ipAddressType,

            ImmutableArray<string> networkInterfaceIds,

            string ownerId,

            string policy,

            string prefixListId,

            bool privateDnsEnabled,

            bool requesterManaged,

            ImmutableArray<string> routeTableIds,

            ImmutableArray<string> securityGroupIds,

            string serviceName,

            string state,

            ImmutableArray<string> subnetIds,

            ImmutableDictionary<string, string> tags,

            string vpcEndpointType,

            string vpcId)
        {
            Arn = arn;
            CidrBlocks = cidrBlocks;
            DnsEntries = dnsEntries;
            DnsOptions = dnsOptions;
            Filters = filters;
            Id = id;
            IpAddressType = ipAddressType;
            NetworkInterfaceIds = networkInterfaceIds;
            OwnerId = ownerId;
            Policy = policy;
            PrefixListId = prefixListId;
            PrivateDnsEnabled = privateDnsEnabled;
            RequesterManaged = requesterManaged;
            RouteTableIds = routeTableIds;
            SecurityGroupIds = securityGroupIds;
            ServiceName = serviceName;
            State = state;
            SubnetIds = subnetIds;
            Tags = tags;
            VpcEndpointType = vpcEndpointType;
            VpcId = vpcId;
        }
    }
}
