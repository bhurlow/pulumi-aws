// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DirectConnect
{
    public static class GetConnection
    {
        /// <summary>
        /// Retrieve information about a Direct Connect Connection.
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
        ///     var example = Aws.DirectConnect.GetConnection.Invoke(new()
        ///     {
        ///         Name = "tf-dx-connection",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetConnectionResult> InvokeAsync(GetConnectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConnectionResult>("aws:directconnect/getConnection:getConnection", args ?? new GetConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieve information about a Direct Connect Connection.
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
        ///     var example = Aws.DirectConnect.GetConnection.Invoke(new()
        ///     {
        ///         Name = "tf-dx-connection",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetConnectionResult> Invoke(GetConnectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConnectionResult>("aws:directconnect/getConnection:getConnection", args ?? new GetConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConnectionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the connection to retrieve.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// Map of tags for the resource.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetConnectionArgs()
        {
        }
        public static new GetConnectionArgs Empty => new GetConnectionArgs();
    }

    public sealed class GetConnectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the connection to retrieve.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags for the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetConnectionInvokeArgs()
        {
        }
        public static new GetConnectionInvokeArgs Empty => new GetConnectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetConnectionResult
    {
        /// <summary>
        /// ARN of the connection.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Direct Connect endpoint on which the physical connection terminates.
        /// </summary>
        public readonly string AwsDevice;
        /// <summary>
        /// Bandwidth of the connection.
        /// </summary>
        public readonly string Bandwidth;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// AWS Direct Connect location where the connection is located.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        /// <summary>
        /// ID of the AWS account that owns the connection.
        /// </summary>
        public readonly string OwnerAccountId;
        /// <summary>
        /// The name of the AWS Direct Connect service provider associated with the connection.
        /// </summary>
        public readonly string PartnerName;
        /// <summary>
        /// Name of the service provider associated with the connection.
        /// </summary>
        public readonly string ProviderName;
        /// <summary>
        /// Map of tags for the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The VLAN ID.
        /// </summary>
        public readonly int VlanId;

        [OutputConstructor]
        private GetConnectionResult(
            string arn,

            string awsDevice,

            string bandwidth,

            string id,

            string location,

            string name,

            string ownerAccountId,

            string partnerName,

            string providerName,

            ImmutableDictionary<string, string> tags,

            int vlanId)
        {
            Arn = arn;
            AwsDevice = awsDevice;
            Bandwidth = bandwidth;
            Id = id;
            Location = location;
            Name = name;
            OwnerAccountId = ownerAccountId;
            PartnerName = partnerName;
            ProviderName = providerName;
            Tags = tags;
            VlanId = vlanId;
        }
    }
}
