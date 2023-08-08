// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Transfer
{
    public static class GetServer
    {
        /// <summary>
        /// Use this data source to get the ARN of an AWS Transfer Server for use in other
        /// resources.
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
        ///     var example = Aws.Transfer.GetServer.Invoke(new()
        ///     {
        ///         ServerId = "s-1234567",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServerResult> InvokeAsync(GetServerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServerResult>("aws:transfer/getServer:getServer", args ?? new GetServerArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the ARN of an AWS Transfer Server for use in other
        /// resources.
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
        ///     var example = Aws.Transfer.GetServer.Invoke(new()
        ///     {
        ///         ServerId = "s-1234567",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetServerResult> Invoke(GetServerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServerResult>("aws:transfer/getServer:getServer", args ?? new GetServerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID for an SFTP server.
        /// </summary>
        [Input("serverId", required: true)]
        public string ServerId { get; set; } = null!;

        public GetServerArgs()
        {
        }
        public static new GetServerArgs Empty => new GetServerArgs();
    }

    public sealed class GetServerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID for an SFTP server.
        /// </summary>
        [Input("serverId", required: true)]
        public Input<string> ServerId { get; set; } = null!;

        public GetServerInvokeArgs()
        {
        }
        public static new GetServerInvokeArgs Empty => new GetServerInvokeArgs();
    }


    [OutputType]
    public sealed class GetServerResult
    {
        /// <summary>
        /// ARN of Transfer Server.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// ARN of any certificate.
        /// </summary>
        public readonly string Certificate;
        /// <summary>
        /// The domain of the storage system that is used for file transfers.
        /// </summary>
        public readonly string Domain;
        /// <summary>
        /// Endpoint of the Transfer Server (e.g., `s-12345678.server.transfer.REGION.amazonaws.com`).
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// Type of endpoint that the server is connected to.
        /// </summary>
        public readonly string EndpointType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The mode of authentication enabled for this service. The default value is `SERVICE_MANAGED`, which allows you to store and access SFTP user credentials within the service. `API_GATEWAY` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice.
        /// </summary>
        public readonly string IdentityProviderType;
        /// <summary>
        /// ARN of the IAM role used to authenticate the user account with an `identity_provider_type` of `API_GATEWAY`.
        /// </summary>
        public readonly string InvocationRole;
        /// <summary>
        /// ARN of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
        /// </summary>
        public readonly string LoggingRole;
        /// <summary>
        /// File transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint.
        /// </summary>
        public readonly ImmutableArray<string> Protocols;
        /// <summary>
        /// The name of the security policy that is attached to the server.
        /// </summary>
        public readonly string SecurityPolicyName;
        public readonly string ServerId;
        public readonly ImmutableArray<string> StructuredLogDestinations;
        /// <summary>
        /// URL of the service endpoint used to authenticate users with an `identity_provider_type` of `API_GATEWAY`.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetServerResult(
            string arn,

            string certificate,

            string domain,

            string endpoint,

            string endpointType,

            string id,

            string identityProviderType,

            string invocationRole,

            string loggingRole,

            ImmutableArray<string> protocols,

            string securityPolicyName,

            string serverId,

            ImmutableArray<string> structuredLogDestinations,

            string url)
        {
            Arn = arn;
            Certificate = certificate;
            Domain = domain;
            Endpoint = endpoint;
            EndpointType = endpointType;
            Id = id;
            IdentityProviderType = identityProviderType;
            InvocationRole = invocationRole;
            LoggingRole = loggingRole;
            Protocols = protocols;
            SecurityPolicyName = securityPolicyName;
            ServerId = serverId;
            StructuredLogDestinations = structuredLogDestinations;
            Url = url;
        }
    }
}
