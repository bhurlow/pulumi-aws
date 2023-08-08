// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Oam
{
    /// <summary>
    /// Resource for managing an AWS CloudWatch Observability Access Manager Sink Policy.
    /// 
    /// ## Example Usage
    /// ### Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleSink = new Aws.Oam.Sink("exampleSink");
    /// 
    ///     var exampleSinkPolicy = new Aws.Oam.SinkPolicy("exampleSinkPolicy", new()
    ///     {
    ///         SinkIdentifier = exampleSink.Id,
    ///         Policy = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["Version"] = "2012-10-17",
    ///             ["Statement"] = new[]
    ///             {
    ///                 new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["Action"] = new[]
    ///                     {
    ///                         "oam:CreateLink",
    ///                         "oam:UpdateLink",
    ///                     },
    ///                     ["Effect"] = "Allow",
    ///                     ["Resource"] = "*",
    ///                     ["Principal"] = new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["AWS"] = new[]
    ///                         {
    ///                             "1111111111111",
    ///                             "222222222222",
    ///                         },
    ///                     },
    ///                     ["Condition"] = new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["ForAllValues:StringEquals"] = new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             ["oam:ResourceTypes"] = new[]
    ///                             {
    ///                                 "AWS::CloudWatch::Metric",
    ///                                 "AWS::Logs::LogGroup",
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_oam_sink_policy.example
    /// 
    ///  id = "arn:aws:oam:us-west-2:123456789012:sink/sink-id" } Using `pulumi import`, import CloudWatch Observability Access Manager Sink Policy using the `sink_identifier`. For exampleconsole % pulumi import aws_oam_sink_policy.example arn:aws:oam:us-west-2:123456789012:sink/sink-id
    /// </summary>
    [AwsResourceType("aws:oam/sinkPolicy:SinkPolicy")]
    public partial class SinkPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN of the Sink.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// JSON policy to use. If you are updating an existing policy, the entire existing policy is replaced by what you specify here.
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;

        /// <summary>
        /// ID string that AWS generated as part of the sink ARN.
        /// </summary>
        [Output("sinkId")]
        public Output<string> SinkId { get; private set; } = null!;

        /// <summary>
        /// ARN of the sink to attach this policy to.
        /// </summary>
        [Output("sinkIdentifier")]
        public Output<string> SinkIdentifier { get; private set; } = null!;


        /// <summary>
        /// Create a SinkPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SinkPolicy(string name, SinkPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:oam/sinkPolicy:SinkPolicy", name, args ?? new SinkPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SinkPolicy(string name, Input<string> id, SinkPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:oam/sinkPolicy:SinkPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SinkPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SinkPolicy Get(string name, Input<string> id, SinkPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new SinkPolicy(name, id, state, options);
        }
    }

    public sealed class SinkPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// JSON policy to use. If you are updating an existing policy, the entire existing policy is replaced by what you specify here.
        /// </summary>
        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        /// <summary>
        /// ARN of the sink to attach this policy to.
        /// </summary>
        [Input("sinkIdentifier", required: true)]
        public Input<string> SinkIdentifier { get; set; } = null!;

        public SinkPolicyArgs()
        {
        }
        public static new SinkPolicyArgs Empty => new SinkPolicyArgs();
    }

    public sealed class SinkPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the Sink.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// JSON policy to use. If you are updating an existing policy, the entire existing policy is replaced by what you specify here.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// ID string that AWS generated as part of the sink ARN.
        /// </summary>
        [Input("sinkId")]
        public Input<string>? SinkId { get; set; }

        /// <summary>
        /// ARN of the sink to attach this policy to.
        /// </summary>
        [Input("sinkIdentifier")]
        public Input<string>? SinkIdentifier { get; set; }

        public SinkPolicyState()
        {
        }
        public static new SinkPolicyState Empty => new SinkPolicyState();
    }
}
