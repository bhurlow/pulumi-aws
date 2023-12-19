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
    /// Provides a CloudFront Function resource. With CloudFront Functions in Amazon CloudFront, you can write lightweight functions in JavaScript for high-scale, latency-sensitive CDN customizations.
    /// 
    /// See [CloudFront Functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-functions.html)
    /// 
    /// &gt; **NOTE:** You cannot delete a function if it’s associated with a cache behavior. First, update your distributions to remove the function association from all cache behaviors, then delete the function.
    /// 
    /// ## Example Usage
    /// ### Basic Example
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = new Aws.CloudFront.Function("test", new()
    ///     {
    ///         Runtime = "cloudfront-js-2.0",
    ///         Comment = "my function",
    ///         Publish = true,
    ///         Code = File.ReadAllText($"{path.Module}/function.js"),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import CloudFront Functions using the `name`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:cloudfront/function:Function test my_test_function
    /// ```
    /// </summary>
    [AwsResourceType("aws:cloudfront/function:Function")]
    public partial class Function : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) identifying your CloudFront Function.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Source code of the function
        /// </summary>
        [Output("code")]
        public Output<string> Code { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// ETag hash of the function. This is the value for the `DEVELOPMENT` stage of the function.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// ETag hash of any `LIVE` stage of the function.
        /// </summary>
        [Output("liveStageEtag")]
        public Output<string> LiveStageEtag { get; private set; } = null!;

        /// <summary>
        /// Unique name for your CloudFront Function.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Whether to publish creation/change as Live CloudFront Function Version. Defaults to `true`.
        /// </summary>
        [Output("publish")]
        public Output<bool?> Publish { get; private set; } = null!;

        /// <summary>
        /// Identifier of the function's runtime. Valid values are `cloudfront-js-1.0` and `cloudfront-js-2.0`.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("runtime")]
        public Output<string> Runtime { get; private set; } = null!;

        /// <summary>
        /// Status of the function. Can be `UNPUBLISHED`, `UNASSOCIATED` or `ASSOCIATED`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Function resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Function(string name, FunctionArgs args, CustomResourceOptions? options = null)
            : base("aws:cloudfront/function:Function", name, args ?? new FunctionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Function(string name, Input<string> id, FunctionState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudfront/function:Function", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Function resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Function Get(string name, Input<string> id, FunctionState? state = null, CustomResourceOptions? options = null)
        {
            return new Function(name, id, state, options);
        }
    }

    public sealed class FunctionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Source code of the function
        /// </summary>
        [Input("code", required: true)]
        public Input<string> Code { get; set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Unique name for your CloudFront Function.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether to publish creation/change as Live CloudFront Function Version. Defaults to `true`.
        /// </summary>
        [Input("publish")]
        public Input<bool>? Publish { get; set; }

        /// <summary>
        /// Identifier of the function's runtime. Valid values are `cloudfront-js-1.0` and `cloudfront-js-2.0`.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("runtime", required: true)]
        public Input<string> Runtime { get; set; } = null!;

        public FunctionArgs()
        {
        }
        public static new FunctionArgs Empty => new FunctionArgs();
    }

    public sealed class FunctionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) identifying your CloudFront Function.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Source code of the function
        /// </summary>
        [Input("code")]
        public Input<string>? Code { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// ETag hash of the function. This is the value for the `DEVELOPMENT` stage of the function.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// ETag hash of any `LIVE` stage of the function.
        /// </summary>
        [Input("liveStageEtag")]
        public Input<string>? LiveStageEtag { get; set; }

        /// <summary>
        /// Unique name for your CloudFront Function.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether to publish creation/change as Live CloudFront Function Version. Defaults to `true`.
        /// </summary>
        [Input("publish")]
        public Input<bool>? Publish { get; set; }

        /// <summary>
        /// Identifier of the function's runtime. Valid values are `cloudfront-js-1.0` and `cloudfront-js-2.0`.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("runtime")]
        public Input<string>? Runtime { get; set; }

        /// <summary>
        /// Status of the function. Can be `UNPUBLISHED`, `UNASSOCIATED` or `ASSOCIATED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public FunctionState()
        {
        }
        public static new FunctionState Empty => new FunctionState();
    }
}
