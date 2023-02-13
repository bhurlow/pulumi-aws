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
    /// Provides a resource for subscribing to SNS topics. Requires that an SNS topic exist for the subscription to attach to. This resource allows you to automatically place messages sent to SNS topics in SQS queues, send them as HTTP(S) POST requests to a given endpoint, send SMS messages, or notify devices / applications. The most likely use case for provider users will probably be SQS queues.
    /// 
    /// &gt; **NOTE:** If the SNS topic and SQS queue are in different AWS regions, the `aws.sns.TopicSubscription` must use an AWS provider that is in the same region as the SNS topic. If the `aws.sns.TopicSubscription` uses a provider with a different region than the SNS topic, this provider will fail to create the subscription.
    /// 
    /// &gt; **NOTE:** Setup of cross-account subscriptions from SNS topics to SQS queues requires the provider to have access to BOTH accounts.
    /// 
    /// &gt; **NOTE:** If an SNS topic and SQS queue are in different AWS accounts but the same region, the `aws.sns.TopicSubscription` must use the AWS provider for the account with the SQS queue. If `aws.sns.TopicSubscription` uses a Provider with a different account than the SQS queue, this provider creates the subscription but does not keep state and tries to re-create the subscription at every `apply`.
    /// 
    /// &gt; **NOTE:** If an SNS topic and SQS queue are in different AWS accounts and different AWS regions, the subscription needs to be initiated from the account with the SQS queue but in the region of the SNS topic.
    /// 
    /// &gt; **NOTE:** You cannot unsubscribe to a subscription that is pending confirmation. If you use `email`, `email-json`, or `http`/`https` (without auto-confirmation enabled), until the subscription is confirmed (e.g., outside of this provider), AWS does not allow this provider to delete / unsubscribe the subscription. If you `destroy` an unconfirmed subscription, this provider will remove the subscription from its state but the subscription will still exist in AWS. However, if you delete an SNS topic, SNS [deletes all the subscriptions](https://docs.aws.amazon.com/sns/latest/dg/sns-delete-subscription-topic.html) associated with the topic. Also, you can import a subscription after confirmation and then have the capability to delete it.
    /// 
    /// ## Example Usage
    /// 
    /// You can directly supply a topic and ARN by hand in the `topic_arn` property along with the queue ARN:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var userUpdatesSqsTarget = new Aws.Sns.TopicSubscription("userUpdatesSqsTarget", new()
    ///     {
    ///         Endpoint = "arn:aws:sqs:us-west-2:432981146916:queue-too",
    ///         Protocol = "sqs",
    ///         Topic = "arn:aws:sns:us-west-2:432981146916:user-updates-topic",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Alternatively you can use the ARN properties of a managed SNS topic and SQS queue:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var userUpdates = new Aws.Sns.Topic("userUpdates");
    /// 
    ///     var userUpdatesQueue = new Aws.Sqs.Queue("userUpdatesQueue");
    /// 
    ///     var userUpdatesSqsTarget = new Aws.Sns.TopicSubscription("userUpdatesSqsTarget", new()
    ///     {
    ///         Topic = userUpdates.Arn,
    ///         Protocol = "sqs",
    ///         Endpoint = userUpdatesQueue.Arn,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// You can subscribe SNS topics to SQS queues in different Amazon accounts and regions:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var sns = config.GetObject&lt;dynamic&gt;("sns") ?? 
    ///     {
    ///         { "account-id", "111111111111" },
    ///         { "role-name", "service/service" },
    ///         { "name", "example-sns-topic" },
    ///         { "display_name", "example" },
    ///         { "region", "us-west-1" },
    ///     };
    ///     var sqs = config.GetObject&lt;dynamic&gt;("sqs") ?? 
    ///     {
    ///         { "account-id", "222222222222" },
    ///         { "role-name", "service/service" },
    ///         { "name", "example-sqs-queue" },
    ///         { "region", "us-east-1" },
    ///     };
    ///     var sns_topic_policy = Aws.Iam.GetPolicyDocument.Invoke(new()
    ///     {
    ///         PolicyId = "__default_policy_ID",
    ///         Statements = new[]
    ///         {
    ///             new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
    ///             {
    ///                 Actions = new[]
    ///                 {
    ///                     "SNS:Subscribe",
    ///                     "SNS:SetTopicAttributes",
    ///                     "SNS:RemovePermission",
    ///                     "SNS:Publish",
    ///                     "SNS:ListSubscriptionsByTopic",
    ///                     "SNS:GetTopicAttributes",
    ///                     "SNS:DeleteTopic",
    ///                     "SNS:AddPermission",
    ///                 },
    ///                 Conditions = new[]
    ///                 {
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
    ///                     {
    ///                         Test = "StringEquals",
    ///                         Variable = "AWS:SourceOwner",
    ///                         Values = new[]
    ///                         {
    ///                             sns.Account_id,
    ///                         },
    ///                     },
    ///                 },
    ///                 Effect = "Allow",
    ///                 Principals = new[]
    ///                 {
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
    ///                     {
    ///                         Type = "AWS",
    ///                         Identifiers = new[]
    ///                         {
    ///                             "*",
    ///                         },
    ///                     },
    ///                 },
    ///                 Resources = new[]
    ///                 {
    ///                     $"arn:aws:sns:{sns.Region}:{sns.Account_id}:{sns.Name}",
    ///                 },
    ///                 Sid = "__default_statement_ID",
    ///             },
    ///             new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
    ///             {
    ///                 Actions = new[]
    ///                 {
    ///                     "SNS:Subscribe",
    ///                     "SNS:Receive",
    ///                 },
    ///                 Conditions = new[]
    ///                 {
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
    ///                     {
    ///                         Test = "StringLike",
    ///                         Variable = "SNS:Endpoint",
    ///                         Values = new[]
    ///                         {
    ///                             $"arn:aws:sqs:{sqs.Region}:{sqs.Account_id}:{sqs.Name}",
    ///                         },
    ///                     },
    ///                 },
    ///                 Effect = "Allow",
    ///                 Principals = new[]
    ///                 {
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
    ///                     {
    ///                         Type = "AWS",
    ///                         Identifiers = new[]
    ///                         {
    ///                             "*",
    ///                         },
    ///                     },
    ///                 },
    ///                 Resources = new[]
    ///                 {
    ///                     $"arn:aws:sns:{sns.Region}:{sns.Account_id}:{sns.Name}",
    ///                 },
    ///                 Sid = "__console_sub_0",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var sqs_queue_policy = Aws.Iam.GetPolicyDocument.Invoke(new()
    ///     {
    ///         PolicyId = $"arn:aws:sqs:{sqs.Region}:{sqs.Account_id}:{sqs.Name}/SQSDefaultPolicy",
    ///         Statements = new[]
    ///         {
    ///             new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
    ///             {
    ///                 Sid = "example-sns-topic",
    ///                 Effect = "Allow",
    ///                 Principals = new[]
    ///                 {
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
    ///                     {
    ///                         Type = "AWS",
    ///                         Identifiers = new[]
    ///                         {
    ///                             "*",
    ///                         },
    ///                     },
    ///                 },
    ///                 Actions = new[]
    ///                 {
    ///                     "SQS:SendMessage",
    ///                 },
    ///                 Resources = new[]
    ///                 {
    ///                     $"arn:aws:sqs:{sqs.Region}:{sqs.Account_id}:{sqs.Name}",
    ///                 },
    ///                 Conditions = new[]
    ///                 {
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
    ///                     {
    ///                         Test = "ArnEquals",
    ///                         Variable = "aws:SourceArn",
    ///                         Values = new[]
    ///                         {
    ///                             $"arn:aws:sns:{sns.Region}:{sns.Account_id}:{sns.Name}",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     // provider to manage SNS topics
    ///     var awsSns = new Aws.Provider("awsSns", new()
    ///     {
    ///         Region = sns.Region,
    ///         AssumeRole = new Aws.Inputs.ProviderAssumeRoleArgs
    ///         {
    ///             RoleArn = $"arn:aws:iam::{sns.Account_id}:role/{sns.Role_name}",
    ///             SessionName = $"sns-{sns.Region}",
    ///         },
    ///     });
    /// 
    ///     // provider to manage SQS queues
    ///     var awsSqs = new Aws.Provider("awsSqs", new()
    ///     {
    ///         Region = sqs.Region,
    ///         AssumeRole = new Aws.Inputs.ProviderAssumeRoleArgs
    ///         {
    ///             RoleArn = $"arn:aws:iam::{sqs.Account_id}:role/{sqs.Role_name}",
    ///             SessionName = $"sqs-{sqs.Region}",
    ///         },
    ///     });
    /// 
    ///     // provider to subscribe SQS to SNS (using the SQS account but the SNS region)
    ///     var sns2sqs = new Aws.Provider("sns2sqs", new()
    ///     {
    ///         Region = sns.Region,
    ///         AssumeRole = new Aws.Inputs.ProviderAssumeRoleArgs
    ///         {
    ///             RoleArn = $"arn:aws:iam::{sqs.Account_id}:role/{sqs.Role_name}",
    ///             SessionName = $"sns2sqs-{sns.Region}",
    ///         },
    ///     });
    /// 
    ///     var sns_topicTopic = new Aws.Sns.Topic("sns-topicTopic", new()
    ///     {
    ///         DisplayName = sns.Display_name,
    ///         Policy = sns_topic_policy.Apply(sns_topic_policy =&gt; sns_topic_policy.Apply(getPolicyDocumentResult =&gt; getPolicyDocumentResult.Json)),
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = "aws.sns",
    ///     });
    /// 
    ///     var sqs_queue = new Aws.Sqs.Queue("sqs-queue", new()
    ///     {
    ///         Policy = sqs_queue_policy.Apply(sqs_queue_policy =&gt; sqs_queue_policy.Apply(getPolicyDocumentResult =&gt; getPolicyDocumentResult.Json)),
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = "aws.sqs",
    ///     });
    /// 
    ///     var sns_topicTopicSubscription = new Aws.Sns.TopicSubscription("sns-topicTopicSubscription", new()
    ///     {
    ///         Topic = sns_topicTopic.Arn,
    ///         Protocol = "sqs",
    ///         Endpoint = sqs_queue.Arn,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = "aws.sns2sqs",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SNS Topic Subscriptions can be imported using the `subscription arn`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:sns/topicSubscription:TopicSubscription user_updates_sqs_target arn:aws:sns:us-west-2:0123456789012:my-topic:8a21d249-4329-4871-acc6-7be709c6ea7f
    /// ```
    /// </summary>
    [AwsResourceType("aws:sns/topicSubscription:TopicSubscription")]
    public partial class TopicSubscription : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN of the subscription.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Integer indicating number of minutes to wait in retrying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols. Default is `1`.
        /// </summary>
        [Output("confirmationTimeoutInMinutes")]
        public Output<int?> ConfirmationTimeoutInMinutes { get; private set; } = null!;

        /// <summary>
        /// Whether the subscription confirmation request was authenticated.
        /// </summary>
        [Output("confirmationWasAuthenticated")]
        public Output<bool> ConfirmationWasAuthenticated { get; private set; } = null!;

        /// <summary>
        /// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
        /// </summary>
        [Output("deliveryPolicy")]
        public Output<string?> DeliveryPolicy { get; private set; } = null!;

        /// <summary>
        /// Endpoint to send data to. The contents vary with the protocol. See details below.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// Whether the endpoint is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) (e.g., PagerDuty). Default is `false`.
        /// </summary>
        [Output("endpointAutoConfirms")]
        public Output<bool?> EndpointAutoConfirms { get; private set; } = null!;

        /// <summary>
        /// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
        /// </summary>
        [Output("filterPolicy")]
        public Output<string?> FilterPolicy { get; private set; } = null!;

        /// <summary>
        /// Whether the `filter_policy` applies to `MessageAttributes` (default) or `MessageBody`.
        /// </summary>
        [Output("filterPolicyScope")]
        public Output<string> FilterPolicyScope { get; private set; } = null!;

        /// <summary>
        /// AWS account ID of the subscription's owner.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// Whether the subscription has not been confirmed.
        /// </summary>
        [Output("pendingConfirmation")]
        public Output<bool> PendingConfirmation { get; private set; } = null!;

        /// <summary>
        /// Protocol to use. Valid values are: `sqs`, `sms`, `lambda`, `firehose`, and `application`. Protocols `email`, `email-json`, `http` and `https` are also valid but partially supported. See details below.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Whether to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property). Default is `false`.
        /// </summary>
        [Output("rawMessageDelivery")]
        public Output<bool?> RawMessageDelivery { get; private set; } = null!;

        /// <summary>
        /// JSON String with the redrive policy that will be used in the subscription. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-dead-letter-queues.html#how-messages-moved-into-dead-letter-queue) for more details.
        /// </summary>
        [Output("redrivePolicy")]
        public Output<string?> RedrivePolicy { get; private set; } = null!;

        /// <summary>
        /// ARN of the IAM role to publish to Kinesis Data Firehose delivery stream. Refer to [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-firehose-as-subscriber.html).
        /// </summary>
        [Output("subscriptionRoleArn")]
        public Output<string?> SubscriptionRoleArn { get; private set; } = null!;

        /// <summary>
        /// ARN of the SNS topic to subscribe to.
        /// </summary>
        [Output("topic")]
        public Output<string> Topic { get; private set; } = null!;


        /// <summary>
        /// Create a TopicSubscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TopicSubscription(string name, TopicSubscriptionArgs args, CustomResourceOptions? options = null)
            : base("aws:sns/topicSubscription:TopicSubscription", name, args ?? new TopicSubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TopicSubscription(string name, Input<string> id, TopicSubscriptionState? state = null, CustomResourceOptions? options = null)
            : base("aws:sns/topicSubscription:TopicSubscription", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TopicSubscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TopicSubscription Get(string name, Input<string> id, TopicSubscriptionState? state = null, CustomResourceOptions? options = null)
        {
            return new TopicSubscription(name, id, state, options);
        }
    }

    public sealed class TopicSubscriptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Integer indicating number of minutes to wait in retrying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols. Default is `1`.
        /// </summary>
        [Input("confirmationTimeoutInMinutes")]
        public Input<int>? ConfirmationTimeoutInMinutes { get; set; }

        /// <summary>
        /// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
        /// </summary>
        [Input("deliveryPolicy")]
        public Input<string>? DeliveryPolicy { get; set; }

        /// <summary>
        /// Endpoint to send data to. The contents vary with the protocol. See details below.
        /// </summary>
        [Input("endpoint", required: true)]
        public Input<string> Endpoint { get; set; } = null!;

        /// <summary>
        /// Whether the endpoint is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) (e.g., PagerDuty). Default is `false`.
        /// </summary>
        [Input("endpointAutoConfirms")]
        public Input<bool>? EndpointAutoConfirms { get; set; }

        /// <summary>
        /// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
        /// </summary>
        [Input("filterPolicy")]
        public Input<string>? FilterPolicy { get; set; }

        /// <summary>
        /// Whether the `filter_policy` applies to `MessageAttributes` (default) or `MessageBody`.
        /// </summary>
        [Input("filterPolicyScope")]
        public Input<string>? FilterPolicyScope { get; set; }

        /// <summary>
        /// Protocol to use. Valid values are: `sqs`, `sms`, `lambda`, `firehose`, and `application`. Protocols `email`, `email-json`, `http` and `https` are also valid but partially supported. See details below.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// Whether to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property). Default is `false`.
        /// </summary>
        [Input("rawMessageDelivery")]
        public Input<bool>? RawMessageDelivery { get; set; }

        /// <summary>
        /// JSON String with the redrive policy that will be used in the subscription. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-dead-letter-queues.html#how-messages-moved-into-dead-letter-queue) for more details.
        /// </summary>
        [Input("redrivePolicy")]
        public Input<string>? RedrivePolicy { get; set; }

        /// <summary>
        /// ARN of the IAM role to publish to Kinesis Data Firehose delivery stream. Refer to [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-firehose-as-subscriber.html).
        /// </summary>
        [Input("subscriptionRoleArn")]
        public Input<string>? SubscriptionRoleArn { get; set; }

        /// <summary>
        /// ARN of the SNS topic to subscribe to.
        /// </summary>
        [Input("topic", required: true)]
        public Input<string> Topic { get; set; } = null!;

        public TopicSubscriptionArgs()
        {
        }
        public static new TopicSubscriptionArgs Empty => new TopicSubscriptionArgs();
    }

    public sealed class TopicSubscriptionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the subscription.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Integer indicating number of minutes to wait in retrying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols. Default is `1`.
        /// </summary>
        [Input("confirmationTimeoutInMinutes")]
        public Input<int>? ConfirmationTimeoutInMinutes { get; set; }

        /// <summary>
        /// Whether the subscription confirmation request was authenticated.
        /// </summary>
        [Input("confirmationWasAuthenticated")]
        public Input<bool>? ConfirmationWasAuthenticated { get; set; }

        /// <summary>
        /// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
        /// </summary>
        [Input("deliveryPolicy")]
        public Input<string>? DeliveryPolicy { get; set; }

        /// <summary>
        /// Endpoint to send data to. The contents vary with the protocol. See details below.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// Whether the endpoint is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) (e.g., PagerDuty). Default is `false`.
        /// </summary>
        [Input("endpointAutoConfirms")]
        public Input<bool>? EndpointAutoConfirms { get; set; }

        /// <summary>
        /// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
        /// </summary>
        [Input("filterPolicy")]
        public Input<string>? FilterPolicy { get; set; }

        /// <summary>
        /// Whether the `filter_policy` applies to `MessageAttributes` (default) or `MessageBody`.
        /// </summary>
        [Input("filterPolicyScope")]
        public Input<string>? FilterPolicyScope { get; set; }

        /// <summary>
        /// AWS account ID of the subscription's owner.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        /// <summary>
        /// Whether the subscription has not been confirmed.
        /// </summary>
        [Input("pendingConfirmation")]
        public Input<bool>? PendingConfirmation { get; set; }

        /// <summary>
        /// Protocol to use. Valid values are: `sqs`, `sms`, `lambda`, `firehose`, and `application`. Protocols `email`, `email-json`, `http` and `https` are also valid but partially supported. See details below.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Whether to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property). Default is `false`.
        /// </summary>
        [Input("rawMessageDelivery")]
        public Input<bool>? RawMessageDelivery { get; set; }

        /// <summary>
        /// JSON String with the redrive policy that will be used in the subscription. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-dead-letter-queues.html#how-messages-moved-into-dead-letter-queue) for more details.
        /// </summary>
        [Input("redrivePolicy")]
        public Input<string>? RedrivePolicy { get; set; }

        /// <summary>
        /// ARN of the IAM role to publish to Kinesis Data Firehose delivery stream. Refer to [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-firehose-as-subscriber.html).
        /// </summary>
        [Input("subscriptionRoleArn")]
        public Input<string>? SubscriptionRoleArn { get; set; }

        /// <summary>
        /// ARN of the SNS topic to subscribe to.
        /// </summary>
        [Input("topic")]
        public Input<string>? Topic { get; set; }

        public TopicSubscriptionState()
        {
        }
        public static new TopicSubscriptionState Empty => new TopicSubscriptionState();
    }
}
