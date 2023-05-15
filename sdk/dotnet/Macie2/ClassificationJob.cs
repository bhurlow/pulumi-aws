// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Macie2
{
    /// <summary>
    /// Provides a resource to manage an [AWS Macie Classification Job](https://docs.aws.amazon.com/macie/latest/APIReference/jobs.html).
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
    ///     var testAccount = new Aws.Macie2.Account("testAccount");
    /// 
    ///     var testClassificationJob = new Aws.Macie2.ClassificationJob("testClassificationJob", new()
    ///     {
    ///         JobType = "ONE_TIME",
    ///         S3JobDefinition = new Aws.Macie2.Inputs.ClassificationJobS3JobDefinitionArgs
    ///         {
    ///             BucketDefinitions = new[]
    ///             {
    ///                 new Aws.Macie2.Inputs.ClassificationJobS3JobDefinitionBucketDefinitionArgs
    ///                 {
    ///                     AccountId = "ACCOUNT ID",
    ///                     Buckets = new[]
    ///                     {
    ///                         "S3 BUCKET NAME",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             testAccount,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// `aws_macie2_classification_job` can be imported using the id, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:macie2/classificationJob:ClassificationJob example abcd1
    /// ```
    /// </summary>
    [AwsResourceType("aws:macie2/classificationJob:ClassificationJob")]
    public partial class ClassificationJob : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The date and time, in UTC and extended RFC 3339 format, when the job was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The custom data identifiers to use for data analysis and classification.
        /// </summary>
        [Output("customDataIdentifierIds")]
        public Output<ImmutableArray<string>> CustomDataIdentifierIds { get; private set; } = null!;

        /// <summary>
        /// A custom description of the job. The description can contain as many as 200 characters.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to analyze all existing, eligible objects immediately after the job is created.
        /// </summary>
        [Output("initialRun")]
        public Output<bool?> InitialRun { get; private set; } = null!;

        [Output("jobArn")]
        public Output<string> JobArn { get; private set; } = null!;

        [Output("jobId")]
        public Output<string> JobId { get; private set; } = null!;

        /// <summary>
        /// The status for the job. Valid values are: `CANCELLED`, `RUNNING` and `USER_PAUSED`
        /// </summary>
        [Output("jobStatus")]
        public Output<string> JobStatus { get; private set; } = null!;

        /// <summary>
        /// The schedule for running the job. Valid values are: `ONE_TIME` - Run the job only once. If you specify this value, don't specify a value for the `schedule_frequency` property. `SCHEDULED` - Run the job on a daily, weekly, or monthly basis. If you specify this value, use the `schedule_frequency` property to define the recurrence pattern for the job.
        /// </summary>
        [Output("jobType")]
        public Output<string> JobType { get; private set; } = null!;

        /// <summary>
        /// A custom name for the job. The name can contain as many as 500 characters. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// The S3 buckets that contain the objects to analyze, and the scope of that analysis. (documented below)
        /// </summary>
        [Output("s3JobDefinition")]
        public Output<Outputs.ClassificationJobS3JobDefinition> S3JobDefinition { get; private set; } = null!;

        /// <summary>
        /// The sampling depth, as a percentage, to apply when processing objects. This value determines the percentage of eligible objects that the job analyzes. If this value is less than 100, Amazon Macie selects the objects to analyze at random, up to the specified percentage, and analyzes all the data in those objects.
        /// </summary>
        [Output("samplingPercentage")]
        public Output<int> SamplingPercentage { get; private set; } = null!;

        /// <summary>
        /// The recurrence pattern for running the job. To run the job only once, don't specify a value for this property and set the value for the `job_type` property to `ONE_TIME`. (documented below)
        /// </summary>
        [Output("scheduleFrequency")]
        public Output<Outputs.ClassificationJobScheduleFrequency> ScheduleFrequency { get; private set; } = null!;

        /// <summary>
        /// A map of key-value pairs that specifies the tags to associate with the job. A job can have a maximum of 50 tags. Each tag consists of a tag key and an associated tag value. The maximum length of a tag key is 128 characters. The maximum length of a tag value is 256 characters.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// If the current status of the job is `USER_PAUSED`, specifies when the job was paused and when the job or job run will expire and be cancelled if it isn't resumed. This value is present only if the value for `job-status` is `USER_PAUSED`.
        /// </summary>
        [Output("userPausedDetails")]
        public Output<ImmutableArray<Outputs.ClassificationJobUserPausedDetail>> UserPausedDetails { get; private set; } = null!;


        /// <summary>
        /// Create a ClassificationJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClassificationJob(string name, ClassificationJobArgs args, CustomResourceOptions? options = null)
            : base("aws:macie2/classificationJob:ClassificationJob", name, args ?? new ClassificationJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClassificationJob(string name, Input<string> id, ClassificationJobState? state = null, CustomResourceOptions? options = null)
            : base("aws:macie2/classificationJob:ClassificationJob", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClassificationJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClassificationJob Get(string name, Input<string> id, ClassificationJobState? state = null, CustomResourceOptions? options = null)
        {
            return new ClassificationJob(name, id, state, options);
        }
    }

    public sealed class ClassificationJobArgs : global::Pulumi.ResourceArgs
    {
        [Input("customDataIdentifierIds")]
        private InputList<string>? _customDataIdentifierIds;

        /// <summary>
        /// The custom data identifiers to use for data analysis and classification.
        /// </summary>
        public InputList<string> CustomDataIdentifierIds
        {
            get => _customDataIdentifierIds ?? (_customDataIdentifierIds = new InputList<string>());
            set => _customDataIdentifierIds = value;
        }

        /// <summary>
        /// A custom description of the job. The description can contain as many as 200 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies whether to analyze all existing, eligible objects immediately after the job is created.
        /// </summary>
        [Input("initialRun")]
        public Input<bool>? InitialRun { get; set; }

        /// <summary>
        /// The status for the job. Valid values are: `CANCELLED`, `RUNNING` and `USER_PAUSED`
        /// </summary>
        [Input("jobStatus")]
        public Input<string>? JobStatus { get; set; }

        /// <summary>
        /// The schedule for running the job. Valid values are: `ONE_TIME` - Run the job only once. If you specify this value, don't specify a value for the `schedule_frequency` property. `SCHEDULED` - Run the job on a daily, weekly, or monthly basis. If you specify this value, use the `schedule_frequency` property to define the recurrence pattern for the job.
        /// </summary>
        [Input("jobType", required: true)]
        public Input<string> JobType { get; set; } = null!;

        /// <summary>
        /// A custom name for the job. The name can contain as many as 500 characters. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The S3 buckets that contain the objects to analyze, and the scope of that analysis. (documented below)
        /// </summary>
        [Input("s3JobDefinition", required: true)]
        public Input<Inputs.ClassificationJobS3JobDefinitionArgs> S3JobDefinition { get; set; } = null!;

        /// <summary>
        /// The sampling depth, as a percentage, to apply when processing objects. This value determines the percentage of eligible objects that the job analyzes. If this value is less than 100, Amazon Macie selects the objects to analyze at random, up to the specified percentage, and analyzes all the data in those objects.
        /// </summary>
        [Input("samplingPercentage")]
        public Input<int>? SamplingPercentage { get; set; }

        /// <summary>
        /// The recurrence pattern for running the job. To run the job only once, don't specify a value for this property and set the value for the `job_type` property to `ONE_TIME`. (documented below)
        /// </summary>
        [Input("scheduleFrequency")]
        public Input<Inputs.ClassificationJobScheduleFrequencyArgs>? ScheduleFrequency { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of key-value pairs that specifies the tags to associate with the job. A job can have a maximum of 50 tags. Each tag consists of a tag key and an associated tag value. The maximum length of a tag key is 128 characters. The maximum length of a tag value is 256 characters.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ClassificationJobArgs()
        {
        }
        public static new ClassificationJobArgs Empty => new ClassificationJobArgs();
    }

    public sealed class ClassificationJobState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time, in UTC and extended RFC 3339 format, when the job was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("customDataIdentifierIds")]
        private InputList<string>? _customDataIdentifierIds;

        /// <summary>
        /// The custom data identifiers to use for data analysis and classification.
        /// </summary>
        public InputList<string> CustomDataIdentifierIds
        {
            get => _customDataIdentifierIds ?? (_customDataIdentifierIds = new InputList<string>());
            set => _customDataIdentifierIds = value;
        }

        /// <summary>
        /// A custom description of the job. The description can contain as many as 200 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies whether to analyze all existing, eligible objects immediately after the job is created.
        /// </summary>
        [Input("initialRun")]
        public Input<bool>? InitialRun { get; set; }

        [Input("jobArn")]
        public Input<string>? JobArn { get; set; }

        [Input("jobId")]
        public Input<string>? JobId { get; set; }

        /// <summary>
        /// The status for the job. Valid values are: `CANCELLED`, `RUNNING` and `USER_PAUSED`
        /// </summary>
        [Input("jobStatus")]
        public Input<string>? JobStatus { get; set; }

        /// <summary>
        /// The schedule for running the job. Valid values are: `ONE_TIME` - Run the job only once. If you specify this value, don't specify a value for the `schedule_frequency` property. `SCHEDULED` - Run the job on a daily, weekly, or monthly basis. If you specify this value, use the `schedule_frequency` property to define the recurrence pattern for the job.
        /// </summary>
        [Input("jobType")]
        public Input<string>? JobType { get; set; }

        /// <summary>
        /// A custom name for the job. The name can contain as many as 500 characters. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The S3 buckets that contain the objects to analyze, and the scope of that analysis. (documented below)
        /// </summary>
        [Input("s3JobDefinition")]
        public Input<Inputs.ClassificationJobS3JobDefinitionGetArgs>? S3JobDefinition { get; set; }

        /// <summary>
        /// The sampling depth, as a percentage, to apply when processing objects. This value determines the percentage of eligible objects that the job analyzes. If this value is less than 100, Amazon Macie selects the objects to analyze at random, up to the specified percentage, and analyzes all the data in those objects.
        /// </summary>
        [Input("samplingPercentage")]
        public Input<int>? SamplingPercentage { get; set; }

        /// <summary>
        /// The recurrence pattern for running the job. To run the job only once, don't specify a value for this property and set the value for the `job_type` property to `ONE_TIME`. (documented below)
        /// </summary>
        [Input("scheduleFrequency")]
        public Input<Inputs.ClassificationJobScheduleFrequencyGetArgs>? ScheduleFrequency { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of key-value pairs that specifies the tags to associate with the job. A job can have a maximum of 50 tags. Each tag consists of a tag key and an associated tag value. The maximum length of a tag key is 128 characters. The maximum length of a tag value is 256 characters.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        [Input("userPausedDetails")]
        private InputList<Inputs.ClassificationJobUserPausedDetailGetArgs>? _userPausedDetails;

        /// <summary>
        /// If the current status of the job is `USER_PAUSED`, specifies when the job was paused and when the job or job run will expire and be cancelled if it isn't resumed. This value is present only if the value for `job-status` is `USER_PAUSED`.
        /// </summary>
        public InputList<Inputs.ClassificationJobUserPausedDetailGetArgs> UserPausedDetails
        {
            get => _userPausedDetails ?? (_userPausedDetails = new InputList<Inputs.ClassificationJobUserPausedDetailGetArgs>());
            set => _userPausedDetails = value;
        }

        public ClassificationJobState()
        {
        }
        public static new ClassificationJobState Empty => new ClassificationJobState();
    }
}
