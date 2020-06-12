// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway
{
    public static class GetKey
    {
        /// <summary>
        /// Use this data source to get the name and value of a pre-existing API Key, for
        /// example to supply credentials for a dependency microservice.
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
        ///         var myApiKey = Output.Create(Aws.ApiGateway.GetKey.InvokeAsync(new Aws.ApiGateway.GetKeyArgs
        ///         {
        ///             Id = "ru3mpjgse6",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetKeyResult> InvokeAsync(GetKeyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKeyResult>("aws:apigateway/getKey:getKey", args ?? new GetKeyArgs(), options.WithVersion());
    }


    public sealed class GetKeyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the API Key to look up.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A map of tags for the resource.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetKeyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetKeyResult
    {
        /// <summary>
        /// The date and time when the API Key was created.
        /// </summary>
        public readonly string CreatedDate;
        /// <summary>
        /// The description of the API Key.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Specifies whether the API Key is enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// Set to the ID of the API Key.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The date and time when the API Key was last updated.
        /// </summary>
        public readonly string LastUpdatedDate;
        /// <summary>
        /// Set to the name of the API Key.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A map of tags for the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// Set to the value of the API Key.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetKeyResult(
            string createdDate,

            string description,

            bool enabled,

            string id,

            string lastUpdatedDate,

            string name,

            ImmutableDictionary<string, object> tags,

            string value)
        {
            CreatedDate = createdDate;
            Description = description;
            Enabled = enabled;
            Id = id;
            LastUpdatedDate = lastUpdatedDate;
            Name = name;
            Tags = tags;
            Value = value;
        }
    }
}
