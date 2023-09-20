// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.NetworkManager
{
    /// <summary>
    /// Resource for managing an AWS NetworkManager ConnectAttachment.
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import `aws_networkmanager_connect_attachment` using the attachment ID. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:networkmanager/connectAttachment:ConnectAttachment example attachment-0f8fa60d2238d1bd8
    /// ```
    /// </summary>
    [AwsResourceType("aws:networkmanager/connectAttachment:ConnectAttachment")]
    public partial class ConnectAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the attachment.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("attachmentId")]
        public Output<string> AttachmentId { get; private set; } = null!;

        /// <summary>
        /// The policy rule number associated with the attachment.
        /// </summary>
        [Output("attachmentPolicyRuleNumber")]
        public Output<int> AttachmentPolicyRuleNumber { get; private set; } = null!;

        /// <summary>
        /// The type of attachment.
        /// </summary>
        [Output("attachmentType")]
        public Output<string> AttachmentType { get; private set; } = null!;

        /// <summary>
        /// The ARN of a core network.
        /// </summary>
        [Output("coreNetworkArn")]
        public Output<string> CoreNetworkArn { get; private set; } = null!;

        /// <summary>
        /// The ID of a core network where you want to create the attachment.
        /// </summary>
        [Output("coreNetworkId")]
        public Output<string> CoreNetworkId { get; private set; } = null!;

        /// <summary>
        /// The Region where the edge is located.
        /// </summary>
        [Output("edgeLocation")]
        public Output<string> EdgeLocation { get; private set; } = null!;

        /// <summary>
        /// Options for creating an attachment.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("options")]
        public Output<Outputs.ConnectAttachmentOptions> Options { get; private set; } = null!;

        /// <summary>
        /// The ID of the attachment account owner.
        /// </summary>
        [Output("ownerAccountId")]
        public Output<string> OwnerAccountId { get; private set; } = null!;

        /// <summary>
        /// The attachment resource ARN.
        /// </summary>
        [Output("resourceArn")]
        public Output<string> ResourceArn { get; private set; } = null!;

        /// <summary>
        /// The name of the segment attachment.
        /// </summary>
        [Output("segmentName")]
        public Output<string> SegmentName { get; private set; } = null!;

        /// <summary>
        /// The state of the attachment.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The ID of the attachment between the two connections.
        /// </summary>
        [Output("transportAttachmentId")]
        public Output<string> TransportAttachmentId { get; private set; } = null!;


        /// <summary>
        /// Create a ConnectAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConnectAttachment(string name, ConnectAttachmentArgs args, CustomResourceOptions? options = null)
            : base("aws:networkmanager/connectAttachment:ConnectAttachment", name, args ?? new ConnectAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConnectAttachment(string name, Input<string> id, ConnectAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("aws:networkmanager/connectAttachment:ConnectAttachment", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "tagsAll",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ConnectAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConnectAttachment Get(string name, Input<string> id, ConnectAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ConnectAttachment(name, id, state, options);
        }
    }

    public sealed class ConnectAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of a core network where you want to create the attachment.
        /// </summary>
        [Input("coreNetworkId", required: true)]
        public Input<string> CoreNetworkId { get; set; } = null!;

        /// <summary>
        /// The Region where the edge is located.
        /// </summary>
        [Input("edgeLocation", required: true)]
        public Input<string> EdgeLocation { get; set; } = null!;

        /// <summary>
        /// Options for creating an attachment.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("options", required: true)]
        public Input<Inputs.ConnectAttachmentOptionsArgs> Options { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the attachment between the two connections.
        /// </summary>
        [Input("transportAttachmentId", required: true)]
        public Input<string> TransportAttachmentId { get; set; } = null!;

        public ConnectAttachmentArgs()
        {
        }
        public static new ConnectAttachmentArgs Empty => new ConnectAttachmentArgs();
    }

    public sealed class ConnectAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the attachment.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("attachmentId")]
        public Input<string>? AttachmentId { get; set; }

        /// <summary>
        /// The policy rule number associated with the attachment.
        /// </summary>
        [Input("attachmentPolicyRuleNumber")]
        public Input<int>? AttachmentPolicyRuleNumber { get; set; }

        /// <summary>
        /// The type of attachment.
        /// </summary>
        [Input("attachmentType")]
        public Input<string>? AttachmentType { get; set; }

        /// <summary>
        /// The ARN of a core network.
        /// </summary>
        [Input("coreNetworkArn")]
        public Input<string>? CoreNetworkArn { get; set; }

        /// <summary>
        /// The ID of a core network where you want to create the attachment.
        /// </summary>
        [Input("coreNetworkId")]
        public Input<string>? CoreNetworkId { get; set; }

        /// <summary>
        /// The Region where the edge is located.
        /// </summary>
        [Input("edgeLocation")]
        public Input<string>? EdgeLocation { get; set; }

        /// <summary>
        /// Options for creating an attachment.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("options")]
        public Input<Inputs.ConnectAttachmentOptionsGetArgs>? Options { get; set; }

        /// <summary>
        /// The ID of the attachment account owner.
        /// </summary>
        [Input("ownerAccountId")]
        public Input<string>? OwnerAccountId { get; set; }

        /// <summary>
        /// The attachment resource ARN.
        /// </summary>
        [Input("resourceArn")]
        public Input<string>? ResourceArn { get; set; }

        /// <summary>
        /// The name of the segment attachment.
        /// </summary>
        [Input("segmentName")]
        public Input<string>? SegmentName { get; set; }

        /// <summary>
        /// The state of the attachment.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Obsolete(@"Please use `tags` instead.")]
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _tagsAll = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// The ID of the attachment between the two connections.
        /// </summary>
        [Input("transportAttachmentId")]
        public Input<string>? TransportAttachmentId { get; set; }

        public ConnectAttachmentState()
        {
        }
        public static new ConnectAttachmentState Empty => new ConnectAttachmentState();
    }
}
