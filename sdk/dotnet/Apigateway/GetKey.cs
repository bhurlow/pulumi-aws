// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get the name and value of a pre-existing API Key, for
        /// example to supply credentials for a dependency microservice.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/api_gateway_api_key.html.markdown.
        /// </summary>
        public static Task<GetKeyResult> GetKey(GetKeyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKeyResult>("aws:apigateway/getKey:getKey", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the API Key to look up.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetKeyArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetKeyResult
    {
        /// <summary>
        /// Set to the ID of the API Key.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Set to the name of the API Key.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Set to the value of the API Key.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetKeyResult(
            string id,
            string name,
            string value)
        {
            Id = id;
            Name = name;
            Value = value;
        }
    }
}
