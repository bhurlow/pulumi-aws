// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sns
{
    /// <summary>
    /// Provides an SNS topic resource
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var userUpdates = new Aws.Sns.Topic("userUpdates");
    /// 
    /// });
    /// ```
    /// ## Example with Delivery Policy
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var userUpdates = new Aws.Sns.Topic("userUpdates", new()
    ///     {
    ///         DeliveryPolicy = @"{
    ///   ""http"": {
    ///     ""defaultHealthyRetryPolicy"": {
    ///       ""minDelayTarget"": 20,
    ///       ""maxDelayTarget"": 20,
    ///       ""numRetries"": 3,
    ///       ""numMaxDelayRetries"": 0,
    ///       ""numNoDelayRetries"": 0,
    ///       ""numMinDelayRetries"": 0,
    ///       ""backoffFunction"": ""linear""
    ///     },
    ///     ""disableSubscriptionOverrides"": false,
    ///     ""defaultThrottlePolicy"": {
    ///       ""maxReceivesPerSecond"": 1
    ///     }
    ///   }
    /// }
    /// 
    /// ",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Example with Server-side encryption (SSE)
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var userUpdates = new Aws.Sns.Topic("userUpdates", new()
    ///     {
    ///         KmsMasterKeyId = "alias/aws/sns",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Example with First-In-First-Out (FIFO)
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var userUpdates = new Aws.Sns.Topic("userUpdates", new()
    ///     {
    ///         ContentBasedDeduplication = true,
    ///         FifoTopic = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Message Delivery Status Arguments
    /// 
    /// The `&lt;endpoint&gt;_success_feedback_role_arn` and `&lt;endpoint&gt;_failure_feedback_role_arn` arguments are used to give Amazon SNS write access to use CloudWatch Logs on your behalf. The `&lt;endpoint&gt;_success_feedback_sample_rate` argument is for specifying the sample rate percentage (0-100) of successfully delivered messages. After you configure the  `&lt;endpoint&gt;_failure_feedback_role_arn` argument, then all failed message deliveries generate CloudWatch Logs.
    /// 
    /// ## Import
    /// 
    /// SNS Topics can be imported using the `topic arn`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:sns/topic:Topic user_updates arn:aws:sns:us-west-2:0123456789012:my-topic
    /// ```
    /// </summary>
    [AwsResourceType("aws:sns/topic:Topic")]
    public partial class Topic : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IAM role for failure feedback
        /// </summary>
        [Output("applicationFailureFeedbackRoleArn")]
        public Output<string?> ApplicationFailureFeedbackRoleArn { get; private set; } = null!;

        /// <summary>
        /// The IAM role permitted to receive success feedback for this topic
        /// </summary>
        [Output("applicationSuccessFeedbackRoleArn")]
        public Output<string?> ApplicationSuccessFeedbackRoleArn { get; private set; } = null!;

        /// <summary>
        /// Percentage of success to sample
        /// </summary>
        [Output("applicationSuccessFeedbackSampleRate")]
        public Output<int?> ApplicationSuccessFeedbackSampleRate { get; private set; } = null!;

        /// <summary>
        /// The ARN of the SNS topic, as a more obvious property (clone of id)
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Enables content-based deduplication for FIFO topics. For more information, see the [related documentation](https://docs.aws.amazon.com/sns/latest/dg/fifo-message-dedup.html)
        /// </summary>
        [Output("contentBasedDeduplication")]
        public Output<bool?> ContentBasedDeduplication { get; private set; } = null!;

        /// <summary>
        /// The SNS delivery policy. More on [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html)
        /// </summary>
        [Output("deliveryPolicy")]
        public Output<string?> DeliveryPolicy { get; private set; } = null!;

        /// <summary>
        /// The display name for the topic
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Boolean indicating whether or not to create a FIFO (first-in-first-out) topic (default is `false`).
        /// </summary>
        [Output("fifoTopic")]
        public Output<bool?> FifoTopic { get; private set; } = null!;

        /// <summary>
        /// IAM role for failure feedback
        /// </summary>
        [Output("firehoseFailureFeedbackRoleArn")]
        public Output<string?> FirehoseFailureFeedbackRoleArn { get; private set; } = null!;

        /// <summary>
        /// The IAM role permitted to receive success feedback for this topic
        /// </summary>
        [Output("firehoseSuccessFeedbackRoleArn")]
        public Output<string?> FirehoseSuccessFeedbackRoleArn { get; private set; } = null!;

        /// <summary>
        /// Percentage of success to sample
        /// </summary>
        [Output("firehoseSuccessFeedbackSampleRate")]
        public Output<int?> FirehoseSuccessFeedbackSampleRate { get; private set; } = null!;

        /// <summary>
        /// IAM role for failure feedback
        /// </summary>
        [Output("httpFailureFeedbackRoleArn")]
        public Output<string?> HttpFailureFeedbackRoleArn { get; private set; } = null!;

        /// <summary>
        /// The IAM role permitted to receive success feedback for this topic
        /// </summary>
        [Output("httpSuccessFeedbackRoleArn")]
        public Output<string?> HttpSuccessFeedbackRoleArn { get; private set; } = null!;

        /// <summary>
        /// Percentage of success to sample
        /// </summary>
        [Output("httpSuccessFeedbackSampleRate")]
        public Output<int?> HttpSuccessFeedbackSampleRate { get; private set; } = null!;

        /// <summary>
        /// The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms)
        /// </summary>
        [Output("kmsMasterKeyId")]
        public Output<string?> KmsMasterKeyId { get; private set; } = null!;

        /// <summary>
        /// IAM role for failure feedback
        /// </summary>
        [Output("lambdaFailureFeedbackRoleArn")]
        public Output<string?> LambdaFailureFeedbackRoleArn { get; private set; } = null!;

        /// <summary>
        /// The IAM role permitted to receive success feedback for this topic
        /// </summary>
        [Output("lambdaSuccessFeedbackRoleArn")]
        public Output<string?> LambdaSuccessFeedbackRoleArn { get; private set; } = null!;

        /// <summary>
        /// Percentage of success to sample
        /// </summary>
        [Output("lambdaSuccessFeedbackSampleRate")]
        public Output<int?> LambdaSuccessFeedbackSampleRate { get; private set; } = null!;

        /// <summary>
        /// The name of the topic. Topic names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. For a FIFO (first-in-first-out) topic, the name must end with the `.fifo` suffix. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`
        /// </summary>
        [Output("namePrefix")]
        public Output<string> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// The AWS Account ID of the SNS topic owner
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// The fully-formed AWS policy as JSON.
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;

        /// <summary>
        /// If `SignatureVersion` should be [1 (SHA1) or 2 (SHA256)](https://docs.aws.amazon.com/sns/latest/dg/sns-verify-signature-of-message.html). The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS.
        /// </summary>
        [Output("signatureVersion")]
        public Output<int> SignatureVersion { get; private set; } = null!;

        /// <summary>
        /// IAM role for failure feedback
        /// </summary>
        [Output("sqsFailureFeedbackRoleArn")]
        public Output<string?> SqsFailureFeedbackRoleArn { get; private set; } = null!;

        /// <summary>
        /// The IAM role permitted to receive success feedback for this topic
        /// </summary>
        [Output("sqsSuccessFeedbackRoleArn")]
        public Output<string?> SqsSuccessFeedbackRoleArn { get; private set; } = null!;

        /// <summary>
        /// Percentage of success to sample
        /// </summary>
        [Output("sqsSuccessFeedbackSampleRate")]
        public Output<int?> SqsSuccessFeedbackSampleRate { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Tracing mode of an Amazon SNS topic. Valid values: `"PassThrough"`, `"Active"`.
        /// </summary>
        [Output("tracingConfig")]
        public Output<string> TracingConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Topic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Topic(string name, TopicArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:sns/topic:Topic", name, args ?? new TopicArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Topic(string name, Input<string> id, TopicState? state = null, CustomResourceOptions? options = null)
            : base("aws:sns/topic:Topic", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Topic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Topic Get(string name, Input<string> id, TopicState? state = null, CustomResourceOptions? options = null)
        {
            return new Topic(name, id, state, options);
        }
    }

    public sealed class TopicArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IAM role for failure feedback
        /// </summary>
        [Input("applicationFailureFeedbackRoleArn")]
        public Input<string>? ApplicationFailureFeedbackRoleArn { get; set; }

        /// <summary>
        /// The IAM role permitted to receive success feedback for this topic
        /// </summary>
        [Input("applicationSuccessFeedbackRoleArn")]
        public Input<string>? ApplicationSuccessFeedbackRoleArn { get; set; }

        /// <summary>
        /// Percentage of success to sample
        /// </summary>
        [Input("applicationSuccessFeedbackSampleRate")]
        public Input<int>? ApplicationSuccessFeedbackSampleRate { get; set; }

        /// <summary>
        /// Enables content-based deduplication for FIFO topics. For more information, see the [related documentation](https://docs.aws.amazon.com/sns/latest/dg/fifo-message-dedup.html)
        /// </summary>
        [Input("contentBasedDeduplication")]
        public Input<bool>? ContentBasedDeduplication { get; set; }

        /// <summary>
        /// The SNS delivery policy. More on [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html)
        /// </summary>
        [Input("deliveryPolicy")]
        public Input<string>? DeliveryPolicy { get; set; }

        /// <summary>
        /// The display name for the topic
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Boolean indicating whether or not to create a FIFO (first-in-first-out) topic (default is `false`).
        /// </summary>
        [Input("fifoTopic")]
        public Input<bool>? FifoTopic { get; set; }

        /// <summary>
        /// IAM role for failure feedback
        /// </summary>
        [Input("firehoseFailureFeedbackRoleArn")]
        public Input<string>? FirehoseFailureFeedbackRoleArn { get; set; }

        /// <summary>
        /// The IAM role permitted to receive success feedback for this topic
        /// </summary>
        [Input("firehoseSuccessFeedbackRoleArn")]
        public Input<string>? FirehoseSuccessFeedbackRoleArn { get; set; }

        /// <summary>
        /// Percentage of success to sample
        /// </summary>
        [Input("firehoseSuccessFeedbackSampleRate")]
        public Input<int>? FirehoseSuccessFeedbackSampleRate { get; set; }

        /// <summary>
        /// IAM role for failure feedback
        /// </summary>
        [Input("httpFailureFeedbackRoleArn")]
        public Input<string>? HttpFailureFeedbackRoleArn { get; set; }

        /// <summary>
        /// The IAM role permitted to receive success feedback for this topic
        /// </summary>
        [Input("httpSuccessFeedbackRoleArn")]
        public Input<string>? HttpSuccessFeedbackRoleArn { get; set; }

        /// <summary>
        /// Percentage of success to sample
        /// </summary>
        [Input("httpSuccessFeedbackSampleRate")]
        public Input<int>? HttpSuccessFeedbackSampleRate { get; set; }

        /// <summary>
        /// The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms)
        /// </summary>
        [Input("kmsMasterKeyId")]
        public Input<string>? KmsMasterKeyId { get; set; }

        /// <summary>
        /// IAM role for failure feedback
        /// </summary>
        [Input("lambdaFailureFeedbackRoleArn")]
        public Input<string>? LambdaFailureFeedbackRoleArn { get; set; }

        /// <summary>
        /// The IAM role permitted to receive success feedback for this topic
        /// </summary>
        [Input("lambdaSuccessFeedbackRoleArn")]
        public Input<string>? LambdaSuccessFeedbackRoleArn { get; set; }

        /// <summary>
        /// Percentage of success to sample
        /// </summary>
        [Input("lambdaSuccessFeedbackSampleRate")]
        public Input<int>? LambdaSuccessFeedbackSampleRate { get; set; }

        /// <summary>
        /// The name of the topic. Topic names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. For a FIFO (first-in-first-out) topic, the name must end with the `.fifo` suffix. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The fully-formed AWS policy as JSON.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// If `SignatureVersion` should be [1 (SHA1) or 2 (SHA256)](https://docs.aws.amazon.com/sns/latest/dg/sns-verify-signature-of-message.html). The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS.
        /// </summary>
        [Input("signatureVersion")]
        public Input<int>? SignatureVersion { get; set; }

        /// <summary>
        /// IAM role for failure feedback
        /// </summary>
        [Input("sqsFailureFeedbackRoleArn")]
        public Input<string>? SqsFailureFeedbackRoleArn { get; set; }

        /// <summary>
        /// The IAM role permitted to receive success feedback for this topic
        /// </summary>
        [Input("sqsSuccessFeedbackRoleArn")]
        public Input<string>? SqsSuccessFeedbackRoleArn { get; set; }

        /// <summary>
        /// Percentage of success to sample
        /// </summary>
        [Input("sqsSuccessFeedbackSampleRate")]
        public Input<int>? SqsSuccessFeedbackSampleRate { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Tracing mode of an Amazon SNS topic. Valid values: `"PassThrough"`, `"Active"`.
        /// </summary>
        [Input("tracingConfig")]
        public Input<string>? TracingConfig { get; set; }

        public TopicArgs()
        {
        }
        public static new TopicArgs Empty => new TopicArgs();
    }

    public sealed class TopicState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IAM role for failure feedback
        /// </summary>
        [Input("applicationFailureFeedbackRoleArn")]
        public Input<string>? ApplicationFailureFeedbackRoleArn { get; set; }

        /// <summary>
        /// The IAM role permitted to receive success feedback for this topic
        /// </summary>
        [Input("applicationSuccessFeedbackRoleArn")]
        public Input<string>? ApplicationSuccessFeedbackRoleArn { get; set; }

        /// <summary>
        /// Percentage of success to sample
        /// </summary>
        [Input("applicationSuccessFeedbackSampleRate")]
        public Input<int>? ApplicationSuccessFeedbackSampleRate { get; set; }

        /// <summary>
        /// The ARN of the SNS topic, as a more obvious property (clone of id)
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Enables content-based deduplication for FIFO topics. For more information, see the [related documentation](https://docs.aws.amazon.com/sns/latest/dg/fifo-message-dedup.html)
        /// </summary>
        [Input("contentBasedDeduplication")]
        public Input<bool>? ContentBasedDeduplication { get; set; }

        /// <summary>
        /// The SNS delivery policy. More on [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html)
        /// </summary>
        [Input("deliveryPolicy")]
        public Input<string>? DeliveryPolicy { get; set; }

        /// <summary>
        /// The display name for the topic
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Boolean indicating whether or not to create a FIFO (first-in-first-out) topic (default is `false`).
        /// </summary>
        [Input("fifoTopic")]
        public Input<bool>? FifoTopic { get; set; }

        /// <summary>
        /// IAM role for failure feedback
        /// </summary>
        [Input("firehoseFailureFeedbackRoleArn")]
        public Input<string>? FirehoseFailureFeedbackRoleArn { get; set; }

        /// <summary>
        /// The IAM role permitted to receive success feedback for this topic
        /// </summary>
        [Input("firehoseSuccessFeedbackRoleArn")]
        public Input<string>? FirehoseSuccessFeedbackRoleArn { get; set; }

        /// <summary>
        /// Percentage of success to sample
        /// </summary>
        [Input("firehoseSuccessFeedbackSampleRate")]
        public Input<int>? FirehoseSuccessFeedbackSampleRate { get; set; }

        /// <summary>
        /// IAM role for failure feedback
        /// </summary>
        [Input("httpFailureFeedbackRoleArn")]
        public Input<string>? HttpFailureFeedbackRoleArn { get; set; }

        /// <summary>
        /// The IAM role permitted to receive success feedback for this topic
        /// </summary>
        [Input("httpSuccessFeedbackRoleArn")]
        public Input<string>? HttpSuccessFeedbackRoleArn { get; set; }

        /// <summary>
        /// Percentage of success to sample
        /// </summary>
        [Input("httpSuccessFeedbackSampleRate")]
        public Input<int>? HttpSuccessFeedbackSampleRate { get; set; }

        /// <summary>
        /// The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms)
        /// </summary>
        [Input("kmsMasterKeyId")]
        public Input<string>? KmsMasterKeyId { get; set; }

        /// <summary>
        /// IAM role for failure feedback
        /// </summary>
        [Input("lambdaFailureFeedbackRoleArn")]
        public Input<string>? LambdaFailureFeedbackRoleArn { get; set; }

        /// <summary>
        /// The IAM role permitted to receive success feedback for this topic
        /// </summary>
        [Input("lambdaSuccessFeedbackRoleArn")]
        public Input<string>? LambdaSuccessFeedbackRoleArn { get; set; }

        /// <summary>
        /// Percentage of success to sample
        /// </summary>
        [Input("lambdaSuccessFeedbackSampleRate")]
        public Input<int>? LambdaSuccessFeedbackSampleRate { get; set; }

        /// <summary>
        /// The name of the topic. Topic names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. For a FIFO (first-in-first-out) topic, the name must end with the `.fifo` suffix. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The AWS Account ID of the SNS topic owner
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        /// <summary>
        /// The fully-formed AWS policy as JSON.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// If `SignatureVersion` should be [1 (SHA1) or 2 (SHA256)](https://docs.aws.amazon.com/sns/latest/dg/sns-verify-signature-of-message.html). The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS.
        /// </summary>
        [Input("signatureVersion")]
        public Input<int>? SignatureVersion { get; set; }

        /// <summary>
        /// IAM role for failure feedback
        /// </summary>
        [Input("sqsFailureFeedbackRoleArn")]
        public Input<string>? SqsFailureFeedbackRoleArn { get; set; }

        /// <summary>
        /// The IAM role permitted to receive success feedback for this topic
        /// </summary>
        [Input("sqsSuccessFeedbackRoleArn")]
        public Input<string>? SqsSuccessFeedbackRoleArn { get; set; }

        /// <summary>
        /// Percentage of success to sample
        /// </summary>
        [Input("sqsSuccessFeedbackSampleRate")]
        public Input<int>? SqsSuccessFeedbackSampleRate { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// Tracing mode of an Amazon SNS topic. Valid values: `"PassThrough"`, `"Active"`.
        /// </summary>
        [Input("tracingConfig")]
        public Input<string>? TracingConfig { get; set; }

        public TopicState()
        {
        }
        public static new TopicState Empty => new TopicState();
    }
}
