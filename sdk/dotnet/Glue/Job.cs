// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue
{
    /// <summary>
    /// Provides a Glue Job resource.
    /// 
    /// &gt; Glue functionality, such as monitoring and logging of jobs, is typically managed with the `default_arguments` argument. See the [Special Parameters Used by AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html) topic in the Glue developer guide for additional information.
    /// 
    /// ## Example Usage
    /// ### Python Job
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.Glue.Job("example", new()
    ///     {
    ///         RoleArn = aws_iam_role.Example.Arn,
    ///         Command = new Aws.Glue.Inputs.JobCommandArgs
    ///         {
    ///             ScriptLocation = $"s3://{aws_s3_bucket.Example.Bucket}/example.py",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Scala Job
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.Glue.Job("example", new()
    ///     {
    ///         RoleArn = aws_iam_role.Example.Arn,
    ///         Command = new Aws.Glue.Inputs.JobCommandArgs
    ///         {
    ///             ScriptLocation = $"s3://{aws_s3_bucket.Example.Bucket}/example.scala",
    ///         },
    ///         DefaultArguments = 
    ///         {
    ///             { "--job-language", "scala" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Streaming Job
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.Glue.Job("example", new()
    ///     {
    ///         RoleArn = aws_iam_role.Example.Arn,
    ///         Command = new Aws.Glue.Inputs.JobCommandArgs
    ///         {
    ///             Name = "gluestreaming",
    ///             ScriptLocation = $"s3://{aws_s3_bucket.Example.Bucket}/example.script",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Enabling CloudWatch Logs and Metrics
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup", new()
    ///     {
    ///         RetentionInDays = 14,
    ///     });
    /// 
    ///     // ... other configuration ...
    ///     var exampleJob = new Aws.Glue.Job("exampleJob", new()
    ///     {
    ///         DefaultArguments = 
    ///         {
    ///             { "--continuous-log-logGroup", exampleLogGroup.Name },
    ///             { "--enable-continuous-cloudwatch-log", "true" },
    ///             { "--enable-continuous-log-filter", "true" },
    ///             { "--enable-metrics", "" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Glue Jobs can be imported using `name`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:glue/job:Job MyJob MyJob
    /// ```
    /// </summary>
    [AwsResourceType("aws:glue/job:Job")]
    public partial class Job : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of Glue Job
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The command of the job. Defined below.
        /// </summary>
        [Output("command")]
        public Output<Outputs.JobCommand> Command { get; private set; } = null!;

        /// <summary>
        /// The list of connections used for this job.
        /// </summary>
        [Output("connections")]
        public Output<ImmutableArray<string>> Connections { get; private set; } = null!;

        /// <summary>
        /// The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
        /// </summary>
        [Output("defaultArguments")]
        public Output<ImmutableDictionary<string, string>?> DefaultArguments { get; private set; } = null!;

        /// <summary>
        /// Description of the job.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the job is run with a standard or flexible execution class. The standard execution class is ideal for time-sensitive workloads that require fast job startup and dedicated resources. Valid value: `FLEX`, `STANDARD`.
        /// </summary>
        [Output("executionClass")]
        public Output<string?> ExecutionClass { get; private set; } = null!;

        /// <summary>
        /// Execution property of the job. Defined below.
        /// </summary>
        [Output("executionProperty")]
        public Output<Outputs.JobExecutionProperty> ExecutionProperty { get; private set; } = null!;

        /// <summary>
        /// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
        /// </summary>
        [Output("glueVersion")]
        public Output<string> GlueVersion { get; private set; } = null!;

        /// <summary>
        /// The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`. Use `number_of_workers` and `worker_type` arguments instead with `glue_version` `2.0` and above.
        /// </summary>
        [Output("maxCapacity")]
        public Output<double> MaxCapacity { get; private set; } = null!;

        /// <summary>
        /// The maximum number of times to retry this job if it fails.
        /// </summary>
        [Output("maxRetries")]
        public Output<int?> MaxRetries { get; private set; } = null!;

        /// <summary>
        /// The name you assign to this job. It must be unique in your account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Non-overridable arguments for this job, specified as name-value pairs.
        /// </summary>
        [Output("nonOverridableArguments")]
        public Output<ImmutableDictionary<string, string>?> NonOverridableArguments { get; private set; } = null!;

        /// <summary>
        /// Notification property of the job. Defined below.
        /// </summary>
        [Output("notificationProperty")]
        public Output<Outputs.JobNotificationProperty> NotificationProperty { get; private set; } = null!;

        /// <summary>
        /// The number of workers of a defined workerType that are allocated when a job runs.
        /// </summary>
        [Output("numberOfWorkers")]
        public Output<int?> NumberOfWorkers { get; private set; } = null!;

        /// <summary>
        /// The ARN of the IAM role associated with this job.
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// The name of the Security Configuration to be associated with the job.
        /// </summary>
        [Output("securityConfiguration")]
        public Output<string?> SecurityConfiguration { get; private set; } = null!;

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
        /// The job timeout in minutes. The default is 2880 minutes (48 hours) for `glueetl` and `pythonshell` jobs, and null (unlimited) for `gluestreaming` jobs.
        /// </summary>
        [Output("timeout")]
        public Output<int> Timeout { get; private set; } = null!;

        /// <summary>
        /// The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.
        /// </summary>
        [Output("workerType")]
        public Output<string?> WorkerType { get; private set; } = null!;


        /// <summary>
        /// Create a Job resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Job(string name, JobArgs args, CustomResourceOptions? options = null)
            : base("aws:glue/job:Job", name, args ?? new JobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Job(string name, Input<string> id, JobState? state = null, CustomResourceOptions? options = null)
            : base("aws:glue/job:Job", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Job resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Job Get(string name, Input<string> id, JobState? state = null, CustomResourceOptions? options = null)
        {
            return new Job(name, id, state, options);
        }
    }

    public sealed class JobArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The command of the job. Defined below.
        /// </summary>
        [Input("command", required: true)]
        public Input<Inputs.JobCommandArgs> Command { get; set; } = null!;

        [Input("connections")]
        private InputList<string>? _connections;

        /// <summary>
        /// The list of connections used for this job.
        /// </summary>
        public InputList<string> Connections
        {
            get => _connections ?? (_connections = new InputList<string>());
            set => _connections = value;
        }

        [Input("defaultArguments")]
        private InputMap<string>? _defaultArguments;

        /// <summary>
        /// The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
        /// </summary>
        public InputMap<string> DefaultArguments
        {
            get => _defaultArguments ?? (_defaultArguments = new InputMap<string>());
            set => _defaultArguments = value;
        }

        /// <summary>
        /// Description of the job.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates whether the job is run with a standard or flexible execution class. The standard execution class is ideal for time-sensitive workloads that require fast job startup and dedicated resources. Valid value: `FLEX`, `STANDARD`.
        /// </summary>
        [Input("executionClass")]
        public Input<string>? ExecutionClass { get; set; }

        /// <summary>
        /// Execution property of the job. Defined below.
        /// </summary>
        [Input("executionProperty")]
        public Input<Inputs.JobExecutionPropertyArgs>? ExecutionProperty { get; set; }

        /// <summary>
        /// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
        /// </summary>
        [Input("glueVersion")]
        public Input<string>? GlueVersion { get; set; }

        /// <summary>
        /// The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`. Use `number_of_workers` and `worker_type` arguments instead with `glue_version` `2.0` and above.
        /// </summary>
        [Input("maxCapacity")]
        public Input<double>? MaxCapacity { get; set; }

        /// <summary>
        /// The maximum number of times to retry this job if it fails.
        /// </summary>
        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        /// <summary>
        /// The name you assign to this job. It must be unique in your account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nonOverridableArguments")]
        private InputMap<string>? _nonOverridableArguments;

        /// <summary>
        /// Non-overridable arguments for this job, specified as name-value pairs.
        /// </summary>
        public InputMap<string> NonOverridableArguments
        {
            get => _nonOverridableArguments ?? (_nonOverridableArguments = new InputMap<string>());
            set => _nonOverridableArguments = value;
        }

        /// <summary>
        /// Notification property of the job. Defined below.
        /// </summary>
        [Input("notificationProperty")]
        public Input<Inputs.JobNotificationPropertyArgs>? NotificationProperty { get; set; }

        /// <summary>
        /// The number of workers of a defined workerType that are allocated when a job runs.
        /// </summary>
        [Input("numberOfWorkers")]
        public Input<int>? NumberOfWorkers { get; set; }

        /// <summary>
        /// The ARN of the IAM role associated with this job.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// The name of the Security Configuration to be associated with the job.
        /// </summary>
        [Input("securityConfiguration")]
        public Input<string>? SecurityConfiguration { get; set; }

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
        /// The job timeout in minutes. The default is 2880 minutes (48 hours) for `glueetl` and `pythonshell` jobs, and null (unlimited) for `gluestreaming` jobs.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.
        /// </summary>
        [Input("workerType")]
        public Input<string>? WorkerType { get; set; }

        public JobArgs()
        {
        }
        public static new JobArgs Empty => new JobArgs();
    }

    public sealed class JobState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of Glue Job
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The command of the job. Defined below.
        /// </summary>
        [Input("command")]
        public Input<Inputs.JobCommandGetArgs>? Command { get; set; }

        [Input("connections")]
        private InputList<string>? _connections;

        /// <summary>
        /// The list of connections used for this job.
        /// </summary>
        public InputList<string> Connections
        {
            get => _connections ?? (_connections = new InputList<string>());
            set => _connections = value;
        }

        [Input("defaultArguments")]
        private InputMap<string>? _defaultArguments;

        /// <summary>
        /// The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
        /// </summary>
        public InputMap<string> DefaultArguments
        {
            get => _defaultArguments ?? (_defaultArguments = new InputMap<string>());
            set => _defaultArguments = value;
        }

        /// <summary>
        /// Description of the job.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates whether the job is run with a standard or flexible execution class. The standard execution class is ideal for time-sensitive workloads that require fast job startup and dedicated resources. Valid value: `FLEX`, `STANDARD`.
        /// </summary>
        [Input("executionClass")]
        public Input<string>? ExecutionClass { get; set; }

        /// <summary>
        /// Execution property of the job. Defined below.
        /// </summary>
        [Input("executionProperty")]
        public Input<Inputs.JobExecutionPropertyGetArgs>? ExecutionProperty { get; set; }

        /// <summary>
        /// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
        /// </summary>
        [Input("glueVersion")]
        public Input<string>? GlueVersion { get; set; }

        /// <summary>
        /// The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`. Use `number_of_workers` and `worker_type` arguments instead with `glue_version` `2.0` and above.
        /// </summary>
        [Input("maxCapacity")]
        public Input<double>? MaxCapacity { get; set; }

        /// <summary>
        /// The maximum number of times to retry this job if it fails.
        /// </summary>
        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        /// <summary>
        /// The name you assign to this job. It must be unique in your account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nonOverridableArguments")]
        private InputMap<string>? _nonOverridableArguments;

        /// <summary>
        /// Non-overridable arguments for this job, specified as name-value pairs.
        /// </summary>
        public InputMap<string> NonOverridableArguments
        {
            get => _nonOverridableArguments ?? (_nonOverridableArguments = new InputMap<string>());
            set => _nonOverridableArguments = value;
        }

        /// <summary>
        /// Notification property of the job. Defined below.
        /// </summary>
        [Input("notificationProperty")]
        public Input<Inputs.JobNotificationPropertyGetArgs>? NotificationProperty { get; set; }

        /// <summary>
        /// The number of workers of a defined workerType that are allocated when a job runs.
        /// </summary>
        [Input("numberOfWorkers")]
        public Input<int>? NumberOfWorkers { get; set; }

        /// <summary>
        /// The ARN of the IAM role associated with this job.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// The name of the Security Configuration to be associated with the job.
        /// </summary>
        [Input("securityConfiguration")]
        public Input<string>? SecurityConfiguration { get; set; }

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
        /// The job timeout in minutes. The default is 2880 minutes (48 hours) for `glueetl` and `pythonshell` jobs, and null (unlimited) for `gluestreaming` jobs.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.
        /// </summary>
        [Input("workerType")]
        public Input<string>? WorkerType { get; set; }

        public JobState()
        {
        }
        public static new JobState Empty => new JobState();
    }
}
