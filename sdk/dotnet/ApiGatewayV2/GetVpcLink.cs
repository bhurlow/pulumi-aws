// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGatewayV2
{
    public static class GetVpcLink
    {
        /// <summary>
        /// Data source for managing an AWS API Gateway V2 VPC Link.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Aws.ApiGatewayV2.GetVpcLink.Invoke(new()
        ///     {
        ///         VpcLinkId = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVpcLinkResult> InvokeAsync(GetVpcLinkArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcLinkResult>("aws:apigatewayv2/getVpcLink:getVpcLink", args ?? new GetVpcLinkArgs(), options.WithDefaults());

        /// <summary>
        /// Data source for managing an AWS API Gateway V2 VPC Link.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Aws.ApiGatewayV2.GetVpcLink.Invoke(new()
        ///     {
        ///         VpcLinkId = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetVpcLinkResult> Invoke(GetVpcLinkInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcLinkResult>("aws:apigatewayv2/getVpcLink:getVpcLink", args ?? new GetVpcLinkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcLinkArgs : global::Pulumi.InvokeArgs
    {
        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// VPC Link Tags.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        /// <summary>
        /// VPC Link ID
        /// </summary>
        [Input("vpcLinkId", required: true)]
        public string VpcLinkId { get; set; } = null!;

        public GetVpcLinkArgs()
        {
        }
        public static new GetVpcLinkArgs Empty => new GetVpcLinkArgs();
    }

    public sealed class GetVpcLinkInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// VPC Link Tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// VPC Link ID
        /// </summary>
        [Input("vpcLinkId", required: true)]
        public Input<string> VpcLinkId { get; set; } = null!;

        public GetVpcLinkInvokeArgs()
        {
        }
        public static new GetVpcLinkInvokeArgs Empty => new GetVpcLinkInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcLinkResult
    {
        /// <summary>
        /// ARN of the VPC Link.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// VPC Link Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of security groups associated with the VPC Link.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// List of subnets attached to the VPC Link.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        /// <summary>
        /// VPC Link Tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        public readonly string VpcLinkId;

        [OutputConstructor]
        private GetVpcLinkResult(
            string arn,

            string id,

            string name,

            ImmutableArray<string> securityGroupIds,

            ImmutableArray<string> subnetIds,

            ImmutableDictionary<string, string>? tags,

            string vpcLinkId)
        {
            Arn = arn;
            Id = id;
            Name = name;
            SecurityGroupIds = securityGroupIds;
            SubnetIds = subnetIds;
            Tags = tags;
            VpcLinkId = vpcLinkId;
        }
    }
}
