// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront
{
    /// <summary>
    /// Resource for managing an AWS CloudFront Key Value Store.
    /// 
    /// ## Example Usage
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
    ///     var example = new Aws.CloudFront.KeyValueStore("example", new()
    ///     {
    ///         Comment = "This is an example key value store",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import CloudFront Key Value Store using the `name`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:cloudfront/keyValueStore:KeyValueStore example example_store
    /// ```
    /// </summary>
    [AwsResourceType("aws:cloudfront/keyValueStore:KeyValueStore")]
    public partial class KeyValueStore : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) identifying your CloudFront KeyValueStore.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// ETag hash of the KeyValueStore.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("lastModifiedTime")]
        public Output<string> LastModifiedTime { get; private set; } = null!;

        /// <summary>
        /// Unique name for your CloudFront KeyValueStore.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("timeouts")]
        public Output<Outputs.KeyValueStoreTimeouts?> Timeouts { get; private set; } = null!;


        /// <summary>
        /// Create a KeyValueStore resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KeyValueStore(string name, KeyValueStoreArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:cloudfront/keyValueStore:KeyValueStore", name, args ?? new KeyValueStoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KeyValueStore(string name, Input<string> id, KeyValueStoreState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudfront/keyValueStore:KeyValueStore", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KeyValueStore resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KeyValueStore Get(string name, Input<string> id, KeyValueStoreState? state = null, CustomResourceOptions? options = null)
        {
            return new KeyValueStore(name, id, state, options);
        }
    }

    public sealed class KeyValueStoreArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Unique name for your CloudFront KeyValueStore.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("timeouts")]
        public Input<Inputs.KeyValueStoreTimeoutsArgs>? Timeouts { get; set; }

        public KeyValueStoreArgs()
        {
        }
        public static new KeyValueStoreArgs Empty => new KeyValueStoreArgs();
    }

    public sealed class KeyValueStoreState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) identifying your CloudFront KeyValueStore.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// ETag hash of the KeyValueStore.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("lastModifiedTime")]
        public Input<string>? LastModifiedTime { get; set; }

        /// <summary>
        /// Unique name for your CloudFront KeyValueStore.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("timeouts")]
        public Input<Inputs.KeyValueStoreTimeoutsGetArgs>? Timeouts { get; set; }

        public KeyValueStoreState()
        {
        }
        public static new KeyValueStoreState Empty => new KeyValueStoreState();
    }
}
