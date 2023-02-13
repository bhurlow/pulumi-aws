// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudTrail
{
    /// <summary>
    /// Provides a CloudTrail resource.
    /// 
    /// &gt; **Tip:** For a multi-region trail, this resource must be in the home region of the trail.
    /// 
    /// &gt; **Tip:** For an organization trail, this resource must be in the master account of the organization.
    /// 
    /// ## Example Usage
    /// ### Basic
    /// 
    /// Enable CloudTrail to capture all compatible management events in region.
    /// For capturing events from services like IAM, `include_global_service_events` must be enabled.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var current = Aws.GetCallerIdentity.Invoke();
    /// 
    ///     var fooBucketV2 = new Aws.S3.BucketV2("fooBucketV2", new()
    ///     {
    ///         ForceDestroy = true,
    ///     });
    /// 
    ///     var foobar = new Aws.CloudTrail.Trail("foobar", new()
    ///     {
    ///         S3BucketName = fooBucketV2.Id,
    ///         S3KeyPrefix = "prefix",
    ///         IncludeGlobalServiceEvents = false,
    ///     });
    /// 
    ///     var fooBucketPolicy = new Aws.S3.BucketPolicy("fooBucketPolicy", new()
    ///     {
    ///         Bucket = fooBucketV2.Id,
    ///         Policy = Output.Tuple(fooBucketV2.Arn, fooBucketV2.Arn, current).Apply(values =&gt;
    ///         {
    ///             var fooBucketV2Arn = values.Item1;
    ///             var fooBucketV2Arn1 = values.Item2;
    ///             var current = values.Item3;
    ///             return @$"{{
    ///     ""Version"": ""2012-10-17"",
    ///     ""Statement"": [
    ///         {{
    ///             ""Sid"": ""AWSCloudTrailAclCheck"",
    ///             ""Effect"": ""Allow"",
    ///             ""Principal"": {{
    ///               ""Service"": ""cloudtrail.amazonaws.com""
    ///             }},
    ///             ""Action"": ""s3:GetBucketAcl"",
    ///             ""Resource"": ""{fooBucketV2Arn}""
    ///         }},
    ///         {{
    ///             ""Sid"": ""AWSCloudTrailWrite"",
    ///             ""Effect"": ""Allow"",
    ///             ""Principal"": {{
    ///               ""Service"": ""cloudtrail.amazonaws.com""
    ///             }},
    ///             ""Action"": ""s3:PutObject"",
    ///             ""Resource"": ""{fooBucketV2Arn1}/prefix/AWSLogs/{current.Apply(getCallerIdentityResult =&gt; getCallerIdentityResult.AccountId)}/*"",
    ///             ""Condition"": {{
    ///                 ""StringEquals"": {{
    ///                     ""s3:x-amz-acl"": ""bucket-owner-full-control""
    ///                 }}
    ///             }}
    ///         }}
    ///     ]
    /// }}
    /// ";
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// ### Data Event Logging
    /// 
    /// CloudTrail can log [Data Events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html) for certain services such as S3 objects and Lambda function invocations. Additional information about data event configuration can be found in the following links:
    /// 
    /// * [CloudTrail API DataResource documentation](https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_DataResource.html) (for basic event selector).
    /// * [CloudTrail API AdvancedFieldSelector documentation](https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_AdvancedFieldSelector.html) (for advanced event selector).
    /// ### Logging All Lambda Function Invocations By Using Basic Event Selectors
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.CloudTrail.Trail("example", new()
    ///     {
    ///         EventSelectors = new[]
    ///         {
    ///             new Aws.CloudTrail.Inputs.TrailEventSelectorArgs
    ///             {
    ///                 DataResources = new[]
    ///                 {
    ///                     new Aws.CloudTrail.Inputs.TrailEventSelectorDataResourceArgs
    ///                     {
    ///                         Type = "AWS::Lambda::Function",
    ///                         Values = new[]
    ///                         {
    ///                             "arn:aws:lambda",
    ///                         },
    ///                     },
    ///                 },
    ///                 IncludeManagementEvents = true,
    ///                 ReadWriteType = "All",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Logging All S3 Object Events By Using Basic Event Selectors
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.CloudTrail.Trail("example", new()
    ///     {
    ///         EventSelectors = new[]
    ///         {
    ///             new Aws.CloudTrail.Inputs.TrailEventSelectorArgs
    ///             {
    ///                 DataResources = new[]
    ///                 {
    ///                     new Aws.CloudTrail.Inputs.TrailEventSelectorDataResourceArgs
    ///                     {
    ///                         Type = "AWS::S3::Object",
    ///                         Values = new[]
    ///                         {
    ///                             "arn:aws:s3",
    ///                         },
    ///                     },
    ///                 },
    ///                 IncludeManagementEvents = true,
    ///                 ReadWriteType = "All",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Logging Individual S3 Bucket Events By Using Basic Event Selectors
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var important_bucket = Aws.S3.GetBucket.Invoke(new()
    ///     {
    ///         Bucket = "important-bucket",
    ///     });
    /// 
    ///     var example = new Aws.CloudTrail.Trail("example", new()
    ///     {
    ///         EventSelectors = new[]
    ///         {
    ///             new Aws.CloudTrail.Inputs.TrailEventSelectorArgs
    ///             {
    ///                 DataResources = new[]
    ///                 {
    ///                     new Aws.CloudTrail.Inputs.TrailEventSelectorDataResourceArgs
    ///                     {
    ///                         Type = "AWS::S3::Object",
    ///                         Values = new[]
    ///                         {
    ///                             important_bucket.Apply(important_bucket =&gt; $"{important_bucket.Apply(getBucketResult =&gt; getBucketResult.Arn)}/"),
    ///                         },
    ///                     },
    ///                 },
    ///                 IncludeManagementEvents = true,
    ///                 ReadWriteType = "All",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Sending Events to CloudWatch Logs
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup");
    /// 
    ///     var exampleTrail = new Aws.CloudTrail.Trail("exampleTrail", new()
    ///     {
    ///         CloudWatchLogsGroupArn = exampleLogGroup.Arn.Apply(arn =&gt; $"{arn}:*"),
    ///     });
    /// 
    ///     // CloudTrail requires the Log Stream wildcard
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cloudtrails can be imported using the `name`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:cloudtrail/trail:Trail sample my-sample-trail
    /// ```
    /// </summary>
    [AwsResourceType("aws:cloudtrail/trail:Trail")]
    public partial class Trail : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies an advanced event selector for enabling data event logging. Fields documented below. Conflicts with `event_selector`.
        /// </summary>
        [Output("advancedEventSelectors")]
        public Output<ImmutableArray<Outputs.TrailAdvancedEventSelector>> AdvancedEventSelectors { get; private set; } = null!;

        /// <summary>
        /// ARN of the trail.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Log group name using an ARN that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
        /// </summary>
        [Output("cloudWatchLogsGroupArn")]
        public Output<string?> CloudWatchLogsGroupArn { get; private set; } = null!;

        /// <summary>
        /// Role for the CloudWatch Logs endpoint to assume to write to a user’s log group.
        /// </summary>
        [Output("cloudWatchLogsRoleArn")]
        public Output<string?> CloudWatchLogsRoleArn { get; private set; } = null!;

        /// <summary>
        /// Whether log file integrity validation is enabled. Defaults to `false`.
        /// </summary>
        [Output("enableLogFileValidation")]
        public Output<bool?> EnableLogFileValidation { get; private set; } = null!;

        /// <summary>
        /// Enables logging for the trail. Defaults to `true`. Setting this to `false` will pause logging.
        /// </summary>
        [Output("enableLogging")]
        public Output<bool?> EnableLogging { get; private set; } = null!;

        /// <summary>
        /// Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these. Conflicts with `advanced_event_selector`.
        /// </summary>
        [Output("eventSelectors")]
        public Output<ImmutableArray<Outputs.TrailEventSelector>> EventSelectors { get; private set; } = null!;

        /// <summary>
        /// Region in which the trail was created.
        /// </summary>
        [Output("homeRegion")]
        public Output<string> HomeRegion { get; private set; } = null!;

        /// <summary>
        /// Whether the trail is publishing events from global services such as IAM to the log files. Defaults to `true`.
        /// </summary>
        [Output("includeGlobalServiceEvents")]
        public Output<bool?> IncludeGlobalServiceEvents { get; private set; } = null!;

        /// <summary>
        /// Configuration block for identifying unusual operational activity. See details below.
        /// </summary>
        [Output("insightSelectors")]
        public Output<ImmutableArray<Outputs.TrailInsightSelector>> InsightSelectors { get; private set; } = null!;

        /// <summary>
        /// Whether the trail is created in the current region or in all regions. Defaults to `false`.
        /// </summary>
        [Output("isMultiRegionTrail")]
        public Output<bool?> IsMultiRegionTrail { get; private set; } = null!;

        /// <summary>
        /// Whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
        /// </summary>
        [Output("isOrganizationTrail")]
        public Output<bool?> IsOrganizationTrail { get; private set; } = null!;

        /// <summary>
        /// KMS key ARN to use to encrypt the logs delivered by CloudTrail.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string?> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// Name of the trail.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Name of the S3 bucket designated for publishing log files.
        /// </summary>
        [Output("s3BucketName")]
        public Output<string> S3BucketName { get; private set; } = null!;

        /// <summary>
        /// S3 key prefix that follows the name of the bucket you have designated for log file delivery.
        /// </summary>
        [Output("s3KeyPrefix")]
        public Output<string?> S3KeyPrefix { get; private set; } = null!;

        /// <summary>
        /// Name of the Amazon SNS topic defined for notification of log file delivery.
        /// </summary>
        [Output("snsTopicName")]
        public Output<string?> SnsTopicName { get; private set; } = null!;

        /// <summary>
        /// Map of tags to assign to the trail. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a Trail resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Trail(string name, TrailArgs args, CustomResourceOptions? options = null)
            : base("aws:cloudtrail/trail:Trail", name, args ?? new TrailArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Trail(string name, Input<string> id, TrailState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudtrail/trail:Trail", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Trail resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Trail Get(string name, Input<string> id, TrailState? state = null, CustomResourceOptions? options = null)
        {
            return new Trail(name, id, state, options);
        }
    }

    public sealed class TrailArgs : global::Pulumi.ResourceArgs
    {
        [Input("advancedEventSelectors")]
        private InputList<Inputs.TrailAdvancedEventSelectorArgs>? _advancedEventSelectors;

        /// <summary>
        /// Specifies an advanced event selector for enabling data event logging. Fields documented below. Conflicts with `event_selector`.
        /// </summary>
        public InputList<Inputs.TrailAdvancedEventSelectorArgs> AdvancedEventSelectors
        {
            get => _advancedEventSelectors ?? (_advancedEventSelectors = new InputList<Inputs.TrailAdvancedEventSelectorArgs>());
            set => _advancedEventSelectors = value;
        }

        /// <summary>
        /// Log group name using an ARN that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
        /// </summary>
        [Input("cloudWatchLogsGroupArn")]
        public Input<string>? CloudWatchLogsGroupArn { get; set; }

        /// <summary>
        /// Role for the CloudWatch Logs endpoint to assume to write to a user’s log group.
        /// </summary>
        [Input("cloudWatchLogsRoleArn")]
        public Input<string>? CloudWatchLogsRoleArn { get; set; }

        /// <summary>
        /// Whether log file integrity validation is enabled. Defaults to `false`.
        /// </summary>
        [Input("enableLogFileValidation")]
        public Input<bool>? EnableLogFileValidation { get; set; }

        /// <summary>
        /// Enables logging for the trail. Defaults to `true`. Setting this to `false` will pause logging.
        /// </summary>
        [Input("enableLogging")]
        public Input<bool>? EnableLogging { get; set; }

        [Input("eventSelectors")]
        private InputList<Inputs.TrailEventSelectorArgs>? _eventSelectors;

        /// <summary>
        /// Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these. Conflicts with `advanced_event_selector`.
        /// </summary>
        public InputList<Inputs.TrailEventSelectorArgs> EventSelectors
        {
            get => _eventSelectors ?? (_eventSelectors = new InputList<Inputs.TrailEventSelectorArgs>());
            set => _eventSelectors = value;
        }

        /// <summary>
        /// Whether the trail is publishing events from global services such as IAM to the log files. Defaults to `true`.
        /// </summary>
        [Input("includeGlobalServiceEvents")]
        public Input<bool>? IncludeGlobalServiceEvents { get; set; }

        [Input("insightSelectors")]
        private InputList<Inputs.TrailInsightSelectorArgs>? _insightSelectors;

        /// <summary>
        /// Configuration block for identifying unusual operational activity. See details below.
        /// </summary>
        public InputList<Inputs.TrailInsightSelectorArgs> InsightSelectors
        {
            get => _insightSelectors ?? (_insightSelectors = new InputList<Inputs.TrailInsightSelectorArgs>());
            set => _insightSelectors = value;
        }

        /// <summary>
        /// Whether the trail is created in the current region or in all regions. Defaults to `false`.
        /// </summary>
        [Input("isMultiRegionTrail")]
        public Input<bool>? IsMultiRegionTrail { get; set; }

        /// <summary>
        /// Whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
        /// </summary>
        [Input("isOrganizationTrail")]
        public Input<bool>? IsOrganizationTrail { get; set; }

        /// <summary>
        /// KMS key ARN to use to encrypt the logs delivered by CloudTrail.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Name of the trail.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of the S3 bucket designated for publishing log files.
        /// </summary>
        [Input("s3BucketName", required: true)]
        public Input<string> S3BucketName { get; set; } = null!;

        /// <summary>
        /// S3 key prefix that follows the name of the bucket you have designated for log file delivery.
        /// </summary>
        [Input("s3KeyPrefix")]
        public Input<string>? S3KeyPrefix { get; set; }

        /// <summary>
        /// Name of the Amazon SNS topic defined for notification of log file delivery.
        /// </summary>
        [Input("snsTopicName")]
        public Input<string>? SnsTopicName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the trail. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public TrailArgs()
        {
        }
        public static new TrailArgs Empty => new TrailArgs();
    }

    public sealed class TrailState : global::Pulumi.ResourceArgs
    {
        [Input("advancedEventSelectors")]
        private InputList<Inputs.TrailAdvancedEventSelectorGetArgs>? _advancedEventSelectors;

        /// <summary>
        /// Specifies an advanced event selector for enabling data event logging. Fields documented below. Conflicts with `event_selector`.
        /// </summary>
        public InputList<Inputs.TrailAdvancedEventSelectorGetArgs> AdvancedEventSelectors
        {
            get => _advancedEventSelectors ?? (_advancedEventSelectors = new InputList<Inputs.TrailAdvancedEventSelectorGetArgs>());
            set => _advancedEventSelectors = value;
        }

        /// <summary>
        /// ARN of the trail.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Log group name using an ARN that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
        /// </summary>
        [Input("cloudWatchLogsGroupArn")]
        public Input<string>? CloudWatchLogsGroupArn { get; set; }

        /// <summary>
        /// Role for the CloudWatch Logs endpoint to assume to write to a user’s log group.
        /// </summary>
        [Input("cloudWatchLogsRoleArn")]
        public Input<string>? CloudWatchLogsRoleArn { get; set; }

        /// <summary>
        /// Whether log file integrity validation is enabled. Defaults to `false`.
        /// </summary>
        [Input("enableLogFileValidation")]
        public Input<bool>? EnableLogFileValidation { get; set; }

        /// <summary>
        /// Enables logging for the trail. Defaults to `true`. Setting this to `false` will pause logging.
        /// </summary>
        [Input("enableLogging")]
        public Input<bool>? EnableLogging { get; set; }

        [Input("eventSelectors")]
        private InputList<Inputs.TrailEventSelectorGetArgs>? _eventSelectors;

        /// <summary>
        /// Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these. Conflicts with `advanced_event_selector`.
        /// </summary>
        public InputList<Inputs.TrailEventSelectorGetArgs> EventSelectors
        {
            get => _eventSelectors ?? (_eventSelectors = new InputList<Inputs.TrailEventSelectorGetArgs>());
            set => _eventSelectors = value;
        }

        /// <summary>
        /// Region in which the trail was created.
        /// </summary>
        [Input("homeRegion")]
        public Input<string>? HomeRegion { get; set; }

        /// <summary>
        /// Whether the trail is publishing events from global services such as IAM to the log files. Defaults to `true`.
        /// </summary>
        [Input("includeGlobalServiceEvents")]
        public Input<bool>? IncludeGlobalServiceEvents { get; set; }

        [Input("insightSelectors")]
        private InputList<Inputs.TrailInsightSelectorGetArgs>? _insightSelectors;

        /// <summary>
        /// Configuration block for identifying unusual operational activity. See details below.
        /// </summary>
        public InputList<Inputs.TrailInsightSelectorGetArgs> InsightSelectors
        {
            get => _insightSelectors ?? (_insightSelectors = new InputList<Inputs.TrailInsightSelectorGetArgs>());
            set => _insightSelectors = value;
        }

        /// <summary>
        /// Whether the trail is created in the current region or in all regions. Defaults to `false`.
        /// </summary>
        [Input("isMultiRegionTrail")]
        public Input<bool>? IsMultiRegionTrail { get; set; }

        /// <summary>
        /// Whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
        /// </summary>
        [Input("isOrganizationTrail")]
        public Input<bool>? IsOrganizationTrail { get; set; }

        /// <summary>
        /// KMS key ARN to use to encrypt the logs delivered by CloudTrail.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Name of the trail.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of the S3 bucket designated for publishing log files.
        /// </summary>
        [Input("s3BucketName")]
        public Input<string>? S3BucketName { get; set; }

        /// <summary>
        /// S3 key prefix that follows the name of the bucket you have designated for log file delivery.
        /// </summary>
        [Input("s3KeyPrefix")]
        public Input<string>? S3KeyPrefix { get; set; }

        /// <summary>
        /// Name of the Amazon SNS topic defined for notification of log file delivery.
        /// </summary>
        [Input("snsTopicName")]
        public Input<string>? SnsTopicName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the trail. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public TrailState()
        {
        }
        public static new TrailState Empty => new TrailState();
    }
}
