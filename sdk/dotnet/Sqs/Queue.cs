// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sqs
{
    /// <summary>
    /// ## Example Usage
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
    ///     var queue = new Aws.Sqs.Queue("queue", new()
    ///     {
    ///         DelaySeconds = 90,
    ///         MaxMessageSize = 2048,
    ///         MessageRetentionSeconds = 86400,
    ///         ReceiveWaitTimeSeconds = 10,
    ///         RedrivePolicy = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["deadLetterTargetArn"] = aws_sqs_queue.Queue_deadletter.Arn,
    ///             ["maxReceiveCount"] = 4,
    ///         }),
    ///         Tags = 
    ///         {
    ///             { "Environment", "production" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ## FIFO queue
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var queue = new Aws.Sqs.Queue("queue", new()
    ///     {
    ///         ContentBasedDeduplication = true,
    ///         FifoQueue = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## High-throughput FIFO queue
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var queue = new Aws.Sqs.Queue("queue", new()
    ///     {
    ///         DeduplicationScope = "messageGroup",
    ///         FifoQueue = true,
    ///         FifoThroughputLimit = "perMessageGroupId",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Dead-letter queue
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
    ///     var exampleQueueDeadletter = new Aws.Sqs.Queue("exampleQueueDeadletter", new()
    ///     {
    ///         RedriveAllowPolicy = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["redrivePermission"] = "byQueue",
    ///             ["sourceQueueArns"] = new[]
    ///             {
    ///                 aws_sqs_queue.Example_queue.Arn,
    ///             },
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Server-side encryption (SSE)
    /// 
    /// Using [SSE-SQS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sqs-sse-queue.html):
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var queue = new Aws.Sqs.Queue("queue", new()
    ///     {
    ///         SqsManagedSseEnabled = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Using [SSE-KMS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sse-existing-queue.html):
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var queue = new Aws.Sqs.Queue("queue", new()
    ///     {
    ///         KmsDataKeyReusePeriodSeconds = 300,
    ///         KmsMasterKeyId = "alias/aws/sqs",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import SQS Queues using the queue `url`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:sqs/queue:Queue public_queue https://queue.amazonaws.com/80398EXAMPLE/MyQueue
    /// ```
    /// </summary>
    [AwsResourceType("aws:sqs/queue:Queue")]
    public partial class Queue : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the SQS queue
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Enables content-based deduplication for FIFO queues. For more information, see the [related documentation](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
        /// </summary>
        [Output("contentBasedDeduplication")]
        public Output<bool?> ContentBasedDeduplication { get; private set; } = null!;

        /// <summary>
        /// Specifies whether message deduplication occurs at the message group or queue level. Valid values are `messageGroup` and `queue` (default).
        /// </summary>
        [Output("deduplicationScope")]
        public Output<string> DeduplicationScope { get; private set; } = null!;

        /// <summary>
        /// The time in seconds that the delivery of all messages in the queue will be delayed. An integer from 0 to 900 (15 minutes). The default for this attribute is 0 seconds.
        /// </summary>
        [Output("delaySeconds")]
        public Output<int?> DelaySeconds { get; private set; } = null!;

        /// <summary>
        /// Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
        /// </summary>
        [Output("fifoQueue")]
        public Output<bool?> FifoQueue { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are `perQueue` (default) and `perMessageGroupId`.
        /// </summary>
        [Output("fifoThroughputLimit")]
        public Output<string> FifoThroughputLimit { get; private set; } = null!;

        /// <summary>
        /// The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours). The default is 300 (5 minutes).
        /// </summary>
        [Output("kmsDataKeyReusePeriodSeconds")]
        public Output<int> KmsDataKeyReusePeriodSeconds { get; private set; } = null!;

        /// <summary>
        /// The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK. For more information, see [Key Terms](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
        /// </summary>
        [Output("kmsMasterKeyId")]
        public Output<string?> KmsMasterKeyId { get; private set; } = null!;

        /// <summary>
        /// The limit of how many bytes a message can contain before Amazon SQS rejects it. An integer from 1024 bytes (1 KiB) up to 262144 bytes (256 KiB). The default for this attribute is 262144 (256 KiB).
        /// </summary>
        [Output("maxMessageSize")]
        public Output<int?> MaxMessageSize { get; private set; } = null!;

        /// <summary>
        /// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days). The default for this attribute is 345600 (4 days).
        /// </summary>
        [Output("messageRetentionSeconds")]
        public Output<int?> MessageRetentionSeconds { get; private set; } = null!;

        /// <summary>
        /// The name of the queue. Queue names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 80 characters long. For a FIFO (first-in-first-out) queue, the name must end with the `.fifo` suffix. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`
        /// </summary>
        [Output("namePrefix")]
        public Output<string> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// The JSON policy for the SQS queue.
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;

        /// <summary>
        /// The time for which a ReceiveMessage call will wait for a message to arrive (long polling) before returning. An integer from 0 to 20 (seconds). The default for this attribute is 0, meaning that the call will return immediately.
        /// </summary>
        [Output("receiveWaitTimeSeconds")]
        public Output<int?> ReceiveWaitTimeSeconds { get; private set; } = null!;

        /// <summary>
        /// The JSON policy to set up the Dead Letter Queue redrive permission, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html).
        /// </summary>
        [Output("redriveAllowPolicy")]
        public Output<string> RedriveAllowPolicy { get; private set; } = null!;

        /// <summary>
        /// The JSON policy to set up the Dead Letter Queue, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html). **Note:** when specifying `maxReceiveCount`, you must specify it as an integer (`5`), and not a string (`"5"`).
        /// </summary>
        [Output("redrivePolicy")]
        public Output<string> RedrivePolicy { get; private set; } = null!;

        /// <summary>
        /// Boolean to enable server-side encryption (SSE) of message content with SQS-owned encryption keys. See [Encryption at rest](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html). The provider will only perform drift detection of its value when present in a configuration.
        /// </summary>
        [Output("sqsManagedSseEnabled")]
        public Output<bool> SqsManagedSseEnabled { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the queue. If configured with a provider `default_tags` configuration block) present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Same as `id`: The URL for the created Amazon SQS queue.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours). The default for this attribute is 30. For more information about visibility timeout, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/AboutVT.html).
        /// </summary>
        [Output("visibilityTimeoutSeconds")]
        public Output<int?> VisibilityTimeoutSeconds { get; private set; } = null!;


        /// <summary>
        /// Create a Queue resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Queue(string name, QueueArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:sqs/queue:Queue", name, args ?? new QueueArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Queue(string name, Input<string> id, QueueState? state = null, CustomResourceOptions? options = null)
            : base("aws:sqs/queue:Queue", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Queue resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Queue Get(string name, Input<string> id, QueueState? state = null, CustomResourceOptions? options = null)
        {
            return new Queue(name, id, state, options);
        }
    }

    public sealed class QueueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables content-based deduplication for FIFO queues. For more information, see the [related documentation](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
        /// </summary>
        [Input("contentBasedDeduplication")]
        public Input<bool>? ContentBasedDeduplication { get; set; }

        /// <summary>
        /// Specifies whether message deduplication occurs at the message group or queue level. Valid values are `messageGroup` and `queue` (default).
        /// </summary>
        [Input("deduplicationScope")]
        public Input<string>? DeduplicationScope { get; set; }

        /// <summary>
        /// The time in seconds that the delivery of all messages in the queue will be delayed. An integer from 0 to 900 (15 minutes). The default for this attribute is 0 seconds.
        /// </summary>
        [Input("delaySeconds")]
        public Input<int>? DelaySeconds { get; set; }

        /// <summary>
        /// Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
        /// </summary>
        [Input("fifoQueue")]
        public Input<bool>? FifoQueue { get; set; }

        /// <summary>
        /// Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are `perQueue` (default) and `perMessageGroupId`.
        /// </summary>
        [Input("fifoThroughputLimit")]
        public Input<string>? FifoThroughputLimit { get; set; }

        /// <summary>
        /// The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours). The default is 300 (5 minutes).
        /// </summary>
        [Input("kmsDataKeyReusePeriodSeconds")]
        public Input<int>? KmsDataKeyReusePeriodSeconds { get; set; }

        /// <summary>
        /// The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK. For more information, see [Key Terms](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
        /// </summary>
        [Input("kmsMasterKeyId")]
        public Input<string>? KmsMasterKeyId { get; set; }

        /// <summary>
        /// The limit of how many bytes a message can contain before Amazon SQS rejects it. An integer from 1024 bytes (1 KiB) up to 262144 bytes (256 KiB). The default for this attribute is 262144 (256 KiB).
        /// </summary>
        [Input("maxMessageSize")]
        public Input<int>? MaxMessageSize { get; set; }

        /// <summary>
        /// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days). The default for this attribute is 345600 (4 days).
        /// </summary>
        [Input("messageRetentionSeconds")]
        public Input<int>? MessageRetentionSeconds { get; set; }

        /// <summary>
        /// The name of the queue. Queue names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 80 characters long. For a FIFO (first-in-first-out) queue, the name must end with the `.fifo` suffix. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The JSON policy for the SQS queue.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// The time for which a ReceiveMessage call will wait for a message to arrive (long polling) before returning. An integer from 0 to 20 (seconds). The default for this attribute is 0, meaning that the call will return immediately.
        /// </summary>
        [Input("receiveWaitTimeSeconds")]
        public Input<int>? ReceiveWaitTimeSeconds { get; set; }

        /// <summary>
        /// The JSON policy to set up the Dead Letter Queue redrive permission, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html).
        /// </summary>
        [Input("redriveAllowPolicy")]
        public Input<string>? RedriveAllowPolicy { get; set; }

        /// <summary>
        /// The JSON policy to set up the Dead Letter Queue, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html). **Note:** when specifying `maxReceiveCount`, you must specify it as an integer (`5`), and not a string (`"5"`).
        /// </summary>
        [Input("redrivePolicy")]
        public Input<string>? RedrivePolicy { get; set; }

        /// <summary>
        /// Boolean to enable server-side encryption (SSE) of message content with SQS-owned encryption keys. See [Encryption at rest](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html). The provider will only perform drift detection of its value when present in a configuration.
        /// </summary>
        [Input("sqsManagedSseEnabled")]
        public Input<bool>? SqsManagedSseEnabled { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the queue. If configured with a provider `default_tags` configuration block) present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours). The default for this attribute is 30. For more information about visibility timeout, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/AboutVT.html).
        /// </summary>
        [Input("visibilityTimeoutSeconds")]
        public Input<int>? VisibilityTimeoutSeconds { get; set; }

        public QueueArgs()
        {
        }
        public static new QueueArgs Empty => new QueueArgs();
    }

    public sealed class QueueState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the SQS queue
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Enables content-based deduplication for FIFO queues. For more information, see the [related documentation](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
        /// </summary>
        [Input("contentBasedDeduplication")]
        public Input<bool>? ContentBasedDeduplication { get; set; }

        /// <summary>
        /// Specifies whether message deduplication occurs at the message group or queue level. Valid values are `messageGroup` and `queue` (default).
        /// </summary>
        [Input("deduplicationScope")]
        public Input<string>? DeduplicationScope { get; set; }

        /// <summary>
        /// The time in seconds that the delivery of all messages in the queue will be delayed. An integer from 0 to 900 (15 minutes). The default for this attribute is 0 seconds.
        /// </summary>
        [Input("delaySeconds")]
        public Input<int>? DelaySeconds { get; set; }

        /// <summary>
        /// Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
        /// </summary>
        [Input("fifoQueue")]
        public Input<bool>? FifoQueue { get; set; }

        /// <summary>
        /// Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are `perQueue` (default) and `perMessageGroupId`.
        /// </summary>
        [Input("fifoThroughputLimit")]
        public Input<string>? FifoThroughputLimit { get; set; }

        /// <summary>
        /// The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours). The default is 300 (5 minutes).
        /// </summary>
        [Input("kmsDataKeyReusePeriodSeconds")]
        public Input<int>? KmsDataKeyReusePeriodSeconds { get; set; }

        /// <summary>
        /// The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK. For more information, see [Key Terms](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
        /// </summary>
        [Input("kmsMasterKeyId")]
        public Input<string>? KmsMasterKeyId { get; set; }

        /// <summary>
        /// The limit of how many bytes a message can contain before Amazon SQS rejects it. An integer from 1024 bytes (1 KiB) up to 262144 bytes (256 KiB). The default for this attribute is 262144 (256 KiB).
        /// </summary>
        [Input("maxMessageSize")]
        public Input<int>? MaxMessageSize { get; set; }

        /// <summary>
        /// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days). The default for this attribute is 345600 (4 days).
        /// </summary>
        [Input("messageRetentionSeconds")]
        public Input<int>? MessageRetentionSeconds { get; set; }

        /// <summary>
        /// The name of the queue. Queue names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 80 characters long. For a FIFO (first-in-first-out) queue, the name must end with the `.fifo` suffix. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The JSON policy for the SQS queue.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// The time for which a ReceiveMessage call will wait for a message to arrive (long polling) before returning. An integer from 0 to 20 (seconds). The default for this attribute is 0, meaning that the call will return immediately.
        /// </summary>
        [Input("receiveWaitTimeSeconds")]
        public Input<int>? ReceiveWaitTimeSeconds { get; set; }

        /// <summary>
        /// The JSON policy to set up the Dead Letter Queue redrive permission, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html).
        /// </summary>
        [Input("redriveAllowPolicy")]
        public Input<string>? RedriveAllowPolicy { get; set; }

        /// <summary>
        /// The JSON policy to set up the Dead Letter Queue, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html). **Note:** when specifying `maxReceiveCount`, you must specify it as an integer (`5`), and not a string (`"5"`).
        /// </summary>
        [Input("redrivePolicy")]
        public Input<string>? RedrivePolicy { get; set; }

        /// <summary>
        /// Boolean to enable server-side encryption (SSE) of message content with SQS-owned encryption keys. See [Encryption at rest](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html). The provider will only perform drift detection of its value when present in a configuration.
        /// </summary>
        [Input("sqsManagedSseEnabled")]
        public Input<bool>? SqsManagedSseEnabled { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the queue. If configured with a provider `default_tags` configuration block) present, tags with matching keys will overwrite those defined at the provider-level.
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
        /// Same as `id`: The URL for the created Amazon SQS queue.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours). The default for this attribute is 30. For more information about visibility timeout, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/AboutVT.html).
        /// </summary>
        [Input("visibilityTimeoutSeconds")]
        public Input<int>? VisibilityTimeoutSeconds { get; set; }

        public QueueState()
        {
        }
        public static new QueueState Empty => new QueueState();
    }
}
