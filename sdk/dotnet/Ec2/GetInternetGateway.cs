// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static class GetInternetGateway
    {
        /// <summary>
        /// `aws.ec2.InternetGateway` provides details about a specific Internet Gateway.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///         var vpcId = config.RequireObject&lt;dynamic&gt;("vpcId");
        ///         var @default = Output.Create(Aws.Ec2.GetInternetGateway.InvokeAsync(new Aws.Ec2.GetInternetGatewayArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Aws.Ec2.Inputs.GetInternetGatewayFilterArgs
        ///                 {
        ///                     Name = "attachment.vpc-id",
        ///                     Values = 
        ///                     {
        ///                         vpcId,
        ///                     },
        ///                 },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInternetGatewayResult> InvokeAsync(GetInternetGatewayArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInternetGatewayResult>("aws:ec2/getInternetGateway:getInternetGateway", args ?? new GetInternetGatewayArgs(), options.WithVersion());
    }


    public sealed class GetInternetGatewayArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetInternetGatewayFilterArgs>? _filters;

        /// <summary>
        /// Custom filter block as described below.
        /// </summary>
        public List<Inputs.GetInternetGatewayFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetInternetGatewayFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The id of the specific Internet Gateway to retrieve.
        /// </summary>
        [Input("internetGatewayId")]
        public string? InternetGatewayId { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A map of tags, each pair of which must exactly match
        /// a pair on the desired Internet Gateway.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetInternetGatewayArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInternetGatewayResult
    {
        public readonly ImmutableArray<Outputs.GetInternetGatewayAttachmentResult> Attachments;
        public readonly ImmutableArray<Outputs.GetInternetGatewayFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InternetGatewayId;
        /// <summary>
        /// The ID of the AWS account that owns the internet gateway.
        /// </summary>
        public readonly string OwnerId;
        public readonly ImmutableDictionary<string, object> Tags;

        [OutputConstructor]
        private GetInternetGatewayResult(
            ImmutableArray<Outputs.GetInternetGatewayAttachmentResult> attachments,

            ImmutableArray<Outputs.GetInternetGatewayFilterResult> filters,

            string id,

            string internetGatewayId,

            string ownerId,

            ImmutableDictionary<string, object> tags)
        {
            Attachments = attachments;
            Filters = filters;
            Id = id;
            InternetGatewayId = internetGatewayId;
            OwnerId = ownerId;
            Tags = tags;
        }
    }
}
