// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppMesh
{
    public static class GetVirtualGateway
    {
        /// <summary>
        /// Data source for managing an AWS App Mesh Virtual Gateway.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVirtualGatewayResult> InvokeAsync(GetVirtualGatewayArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualGatewayResult>("aws:appmesh/getVirtualGateway:getVirtualGateway", args ?? new GetVirtualGatewayArgs(), options.WithDefaults());

        /// <summary>
        /// Data source for managing an AWS App Mesh Virtual Gateway.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetVirtualGatewayResult> Invoke(GetVirtualGatewayInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualGatewayResult>("aws:appmesh/getVirtualGateway:getVirtualGateway", args ?? new GetVirtualGatewayInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualGatewayArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the service mesh in which the virtual gateway exists.
        /// </summary>
        [Input("meshName", required: true)]
        public string MeshName { get; set; } = null!;

        /// <summary>
        /// Name of the virtual gateway.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// Map of tags.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetVirtualGatewayArgs()
        {
        }
        public static new GetVirtualGatewayArgs Empty => new GetVirtualGatewayArgs();
    }

    public sealed class GetVirtualGatewayInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the service mesh in which the virtual gateway exists.
        /// </summary>
        [Input("meshName", required: true)]
        public Input<string> MeshName { get; set; } = null!;

        /// <summary>
        /// Name of the virtual gateway.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetVirtualGatewayInvokeArgs()
        {
        }
        public static new GetVirtualGatewayInvokeArgs Empty => new GetVirtualGatewayInvokeArgs();
    }


    [OutputType]
    public sealed class GetVirtualGatewayResult
    {
        /// <summary>
        /// ARN of the virtual gateway.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Creation date of the virtual gateway.
        /// </summary>
        public readonly string CreatedDate;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Last update date of the virtual gateway.
        /// </summary>
        public readonly string LastUpdatedDate;
        public readonly string MeshName;
        public readonly string MeshOwner;
        public readonly string Name;
        /// <summary>
        /// Resource owner's AWS account ID.
        /// </summary>
        public readonly string ResourceOwner;
        /// <summary>
        /// Virtual gateway specification
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualGatewaySpecResult> Specs;
        /// <summary>
        /// Map of tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetVirtualGatewayResult(
            string arn,

            string createdDate,

            string id,

            string lastUpdatedDate,

            string meshName,

            string meshOwner,

            string name,

            string resourceOwner,

            ImmutableArray<Outputs.GetVirtualGatewaySpecResult> specs,

            ImmutableDictionary<string, string> tags)
        {
            Arn = arn;
            CreatedDate = createdDate;
            Id = id;
            LastUpdatedDate = lastUpdatedDate;
            MeshName = meshName;
            MeshOwner = meshOwner;
            Name = name;
            ResourceOwner = resourceOwner;
            Specs = specs;
            Tags = tags;
        }
    }
}
