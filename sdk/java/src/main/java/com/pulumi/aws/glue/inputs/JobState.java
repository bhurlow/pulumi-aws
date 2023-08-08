// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.glue.inputs;

import com.pulumi.aws.glue.inputs.JobCommandArgs;
import com.pulumi.aws.glue.inputs.JobExecutionPropertyArgs;
import com.pulumi.aws.glue.inputs.JobNotificationPropertyArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class JobState extends com.pulumi.resources.ResourceArgs {

    public static final JobState Empty = new JobState();

    /**
     * Amazon Resource Name (ARN) of Glue Job
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of Glue Job
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The command of the job. Defined below.
     * 
     */
    @Import(name="command")
    private @Nullable Output<JobCommandArgs> command;

    /**
     * @return The command of the job. Defined below.
     * 
     */
    public Optional<Output<JobCommandArgs>> command() {
        return Optional.ofNullable(this.command);
    }

    /**
     * The list of connections used for this job.
     * 
     */
    @Import(name="connections")
    private @Nullable Output<List<String>> connections;

    /**
     * @return The list of connections used for this job.
     * 
     */
    public Optional<Output<List<String>>> connections() {
        return Optional.ofNullable(this.connections);
    }

    /**
     * The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
     * 
     */
    @Import(name="defaultArguments")
    private @Nullable Output<Map<String,String>> defaultArguments;

    /**
     * @return The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
     * 
     */
    public Optional<Output<Map<String,String>>> defaultArguments() {
        return Optional.ofNullable(this.defaultArguments);
    }

    /**
     * Description of the job.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the job.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Indicates whether the job is run with a standard or flexible execution class. The standard execution class is ideal for time-sensitive workloads that require fast job startup and dedicated resources. Valid value: `FLEX`, `STANDARD`.
     * 
     */
    @Import(name="executionClass")
    private @Nullable Output<String> executionClass;

    /**
     * @return Indicates whether the job is run with a standard or flexible execution class. The standard execution class is ideal for time-sensitive workloads that require fast job startup and dedicated resources. Valid value: `FLEX`, `STANDARD`.
     * 
     */
    public Optional<Output<String>> executionClass() {
        return Optional.ofNullable(this.executionClass);
    }

    /**
     * Execution property of the job. Defined below.
     * 
     */
    @Import(name="executionProperty")
    private @Nullable Output<JobExecutionPropertyArgs> executionProperty;

    /**
     * @return Execution property of the job. Defined below.
     * 
     */
    public Optional<Output<JobExecutionPropertyArgs>> executionProperty() {
        return Optional.ofNullable(this.executionProperty);
    }

    /**
     * The version of glue to use, for example &#34;1.0&#34;. Ray jobs should set this to 4.0 or greater. For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
     * 
     */
    @Import(name="glueVersion")
    private @Nullable Output<String> glueVersion;

    /**
     * @return The version of glue to use, for example &#34;1.0&#34;. Ray jobs should set this to 4.0 or greater. For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
     * 
     */
    public Optional<Output<String>> glueVersion() {
        return Optional.ofNullable(this.glueVersion);
    }

    /**
     * The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`. Use `number_of_workers` and `worker_type` arguments instead with `glue_version` `2.0` and above.
     * 
     */
    @Import(name="maxCapacity")
    private @Nullable Output<Double> maxCapacity;

    /**
     * @return The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`. Use `number_of_workers` and `worker_type` arguments instead with `glue_version` `2.0` and above.
     * 
     */
    public Optional<Output<Double>> maxCapacity() {
        return Optional.ofNullable(this.maxCapacity);
    }

    /**
     * The maximum number of times to retry this job if it fails.
     * 
     */
    @Import(name="maxRetries")
    private @Nullable Output<Integer> maxRetries;

    /**
     * @return The maximum number of times to retry this job if it fails.
     * 
     */
    public Optional<Output<Integer>> maxRetries() {
        return Optional.ofNullable(this.maxRetries);
    }

    /**
     * The name you assign to this job. It must be unique in your account.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name you assign to this job. It must be unique in your account.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Non-overridable arguments for this job, specified as name-value pairs.
     * 
     */
    @Import(name="nonOverridableArguments")
    private @Nullable Output<Map<String,String>> nonOverridableArguments;

    /**
     * @return Non-overridable arguments for this job, specified as name-value pairs.
     * 
     */
    public Optional<Output<Map<String,String>>> nonOverridableArguments() {
        return Optional.ofNullable(this.nonOverridableArguments);
    }

    /**
     * Notification property of the job. Defined below.
     * 
     */
    @Import(name="notificationProperty")
    private @Nullable Output<JobNotificationPropertyArgs> notificationProperty;

    /**
     * @return Notification property of the job. Defined below.
     * 
     */
    public Optional<Output<JobNotificationPropertyArgs>> notificationProperty() {
        return Optional.ofNullable(this.notificationProperty);
    }

    /**
     * The number of workers of a defined workerType that are allocated when a job runs.
     * 
     */
    @Import(name="numberOfWorkers")
    private @Nullable Output<Integer> numberOfWorkers;

    /**
     * @return The number of workers of a defined workerType that are allocated when a job runs.
     * 
     */
    public Optional<Output<Integer>> numberOfWorkers() {
        return Optional.ofNullable(this.numberOfWorkers);
    }

    /**
     * The ARN of the IAM role associated with this job.
     * 
     */
    @Import(name="roleArn")
    private @Nullable Output<String> roleArn;

    /**
     * @return The ARN of the IAM role associated with this job.
     * 
     */
    public Optional<Output<String>> roleArn() {
        return Optional.ofNullable(this.roleArn);
    }

    /**
     * The name of the Security Configuration to be associated with the job.
     * 
     */
    @Import(name="securityConfiguration")
    private @Nullable Output<String> securityConfiguration;

    /**
     * @return The name of the Security Configuration to be associated with the job.
     * 
     */
    public Optional<Output<String>> securityConfiguration() {
        return Optional.ofNullable(this.securityConfiguration);
    }

    /**
     * Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * The job timeout in minutes. The default is 2880 minutes (48 hours) for `glueetl` and `pythonshell` jobs, and null (unlimited) for `gluestreaming` jobs.
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    /**
     * @return The job timeout in minutes. The default is 2880 minutes (48 hours) for `glueetl` and `pythonshell` jobs, and null (unlimited) for `gluestreaming` jobs.
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    /**
     * The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, G.2X, or G.025X for Spark jobs. Accepts the value Z.2X for Ray jobs.
     * * For the Standard worker type, each worker provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.
     * * For the G.1X worker type, each worker maps to 1 DPU (4 vCPU, 16 GB of memory, 64 GB disk), and provides 1 executor per worker. Recommended for memory-intensive jobs.
     * * For the G.2X worker type, each worker maps to 2 DPU (8 vCPU, 32 GB of memory, 128 GB disk), and provides 1 executor per worker. Recommended for memory-intensive jobs.
     * * For the G.025X worker type, each worker maps to 0.25 DPU (2 vCPU, 4GB of memory, 64 GB disk), and provides 1 executor per worker. Recommended for low volume streaming jobs. Only available for Glue version 3.0.
     * * For the Z.2X worker type, each worker maps to 2 M-DPU (8vCPU, 64 GB of m emory, 128 GB disk), and provides up to 8 Ray workers based on the autoscaler.
     * 
     */
    @Import(name="workerType")
    private @Nullable Output<String> workerType;

    /**
     * @return The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, G.2X, or G.025X for Spark jobs. Accepts the value Z.2X for Ray jobs.
     * * For the Standard worker type, each worker provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.
     * * For the G.1X worker type, each worker maps to 1 DPU (4 vCPU, 16 GB of memory, 64 GB disk), and provides 1 executor per worker. Recommended for memory-intensive jobs.
     * * For the G.2X worker type, each worker maps to 2 DPU (8 vCPU, 32 GB of memory, 128 GB disk), and provides 1 executor per worker. Recommended for memory-intensive jobs.
     * * For the G.025X worker type, each worker maps to 0.25 DPU (2 vCPU, 4GB of memory, 64 GB disk), and provides 1 executor per worker. Recommended for low volume streaming jobs. Only available for Glue version 3.0.
     * * For the Z.2X worker type, each worker maps to 2 M-DPU (8vCPU, 64 GB of m emory, 128 GB disk), and provides up to 8 Ray workers based on the autoscaler.
     * 
     */
    public Optional<Output<String>> workerType() {
        return Optional.ofNullable(this.workerType);
    }

    private JobState() {}

    private JobState(JobState $) {
        this.arn = $.arn;
        this.command = $.command;
        this.connections = $.connections;
        this.defaultArguments = $.defaultArguments;
        this.description = $.description;
        this.executionClass = $.executionClass;
        this.executionProperty = $.executionProperty;
        this.glueVersion = $.glueVersion;
        this.maxCapacity = $.maxCapacity;
        this.maxRetries = $.maxRetries;
        this.name = $.name;
        this.nonOverridableArguments = $.nonOverridableArguments;
        this.notificationProperty = $.notificationProperty;
        this.numberOfWorkers = $.numberOfWorkers;
        this.roleArn = $.roleArn;
        this.securityConfiguration = $.securityConfiguration;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.timeout = $.timeout;
        this.workerType = $.workerType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(JobState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private JobState $;

        public Builder() {
            $ = new JobState();
        }

        public Builder(JobState defaults) {
            $ = new JobState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn Amazon Resource Name (ARN) of Glue Job
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn Amazon Resource Name (ARN) of Glue Job
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param command The command of the job. Defined below.
         * 
         * @return builder
         * 
         */
        public Builder command(@Nullable Output<JobCommandArgs> command) {
            $.command = command;
            return this;
        }

        /**
         * @param command The command of the job. Defined below.
         * 
         * @return builder
         * 
         */
        public Builder command(JobCommandArgs command) {
            return command(Output.of(command));
        }

        /**
         * @param connections The list of connections used for this job.
         * 
         * @return builder
         * 
         */
        public Builder connections(@Nullable Output<List<String>> connections) {
            $.connections = connections;
            return this;
        }

        /**
         * @param connections The list of connections used for this job.
         * 
         * @return builder
         * 
         */
        public Builder connections(List<String> connections) {
            return connections(Output.of(connections));
        }

        /**
         * @param connections The list of connections used for this job.
         * 
         * @return builder
         * 
         */
        public Builder connections(String... connections) {
            return connections(List.of(connections));
        }

        /**
         * @param defaultArguments The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
         * 
         * @return builder
         * 
         */
        public Builder defaultArguments(@Nullable Output<Map<String,String>> defaultArguments) {
            $.defaultArguments = defaultArguments;
            return this;
        }

        /**
         * @param defaultArguments The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
         * 
         * @return builder
         * 
         */
        public Builder defaultArguments(Map<String,String> defaultArguments) {
            return defaultArguments(Output.of(defaultArguments));
        }

        /**
         * @param description Description of the job.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the job.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param executionClass Indicates whether the job is run with a standard or flexible execution class. The standard execution class is ideal for time-sensitive workloads that require fast job startup and dedicated resources. Valid value: `FLEX`, `STANDARD`.
         * 
         * @return builder
         * 
         */
        public Builder executionClass(@Nullable Output<String> executionClass) {
            $.executionClass = executionClass;
            return this;
        }

        /**
         * @param executionClass Indicates whether the job is run with a standard or flexible execution class. The standard execution class is ideal for time-sensitive workloads that require fast job startup and dedicated resources. Valid value: `FLEX`, `STANDARD`.
         * 
         * @return builder
         * 
         */
        public Builder executionClass(String executionClass) {
            return executionClass(Output.of(executionClass));
        }

        /**
         * @param executionProperty Execution property of the job. Defined below.
         * 
         * @return builder
         * 
         */
        public Builder executionProperty(@Nullable Output<JobExecutionPropertyArgs> executionProperty) {
            $.executionProperty = executionProperty;
            return this;
        }

        /**
         * @param executionProperty Execution property of the job. Defined below.
         * 
         * @return builder
         * 
         */
        public Builder executionProperty(JobExecutionPropertyArgs executionProperty) {
            return executionProperty(Output.of(executionProperty));
        }

        /**
         * @param glueVersion The version of glue to use, for example &#34;1.0&#34;. Ray jobs should set this to 4.0 or greater. For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
         * 
         * @return builder
         * 
         */
        public Builder glueVersion(@Nullable Output<String> glueVersion) {
            $.glueVersion = glueVersion;
            return this;
        }

        /**
         * @param glueVersion The version of glue to use, for example &#34;1.0&#34;. Ray jobs should set this to 4.0 or greater. For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
         * 
         * @return builder
         * 
         */
        public Builder glueVersion(String glueVersion) {
            return glueVersion(Output.of(glueVersion));
        }

        /**
         * @param maxCapacity The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`. Use `number_of_workers` and `worker_type` arguments instead with `glue_version` `2.0` and above.
         * 
         * @return builder
         * 
         */
        public Builder maxCapacity(@Nullable Output<Double> maxCapacity) {
            $.maxCapacity = maxCapacity;
            return this;
        }

        /**
         * @param maxCapacity The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`. Use `number_of_workers` and `worker_type` arguments instead with `glue_version` `2.0` and above.
         * 
         * @return builder
         * 
         */
        public Builder maxCapacity(Double maxCapacity) {
            return maxCapacity(Output.of(maxCapacity));
        }

        /**
         * @param maxRetries The maximum number of times to retry this job if it fails.
         * 
         * @return builder
         * 
         */
        public Builder maxRetries(@Nullable Output<Integer> maxRetries) {
            $.maxRetries = maxRetries;
            return this;
        }

        /**
         * @param maxRetries The maximum number of times to retry this job if it fails.
         * 
         * @return builder
         * 
         */
        public Builder maxRetries(Integer maxRetries) {
            return maxRetries(Output.of(maxRetries));
        }

        /**
         * @param name The name you assign to this job. It must be unique in your account.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name you assign to this job. It must be unique in your account.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param nonOverridableArguments Non-overridable arguments for this job, specified as name-value pairs.
         * 
         * @return builder
         * 
         */
        public Builder nonOverridableArguments(@Nullable Output<Map<String,String>> nonOverridableArguments) {
            $.nonOverridableArguments = nonOverridableArguments;
            return this;
        }

        /**
         * @param nonOverridableArguments Non-overridable arguments for this job, specified as name-value pairs.
         * 
         * @return builder
         * 
         */
        public Builder nonOverridableArguments(Map<String,String> nonOverridableArguments) {
            return nonOverridableArguments(Output.of(nonOverridableArguments));
        }

        /**
         * @param notificationProperty Notification property of the job. Defined below.
         * 
         * @return builder
         * 
         */
        public Builder notificationProperty(@Nullable Output<JobNotificationPropertyArgs> notificationProperty) {
            $.notificationProperty = notificationProperty;
            return this;
        }

        /**
         * @param notificationProperty Notification property of the job. Defined below.
         * 
         * @return builder
         * 
         */
        public Builder notificationProperty(JobNotificationPropertyArgs notificationProperty) {
            return notificationProperty(Output.of(notificationProperty));
        }

        /**
         * @param numberOfWorkers The number of workers of a defined workerType that are allocated when a job runs.
         * 
         * @return builder
         * 
         */
        public Builder numberOfWorkers(@Nullable Output<Integer> numberOfWorkers) {
            $.numberOfWorkers = numberOfWorkers;
            return this;
        }

        /**
         * @param numberOfWorkers The number of workers of a defined workerType that are allocated when a job runs.
         * 
         * @return builder
         * 
         */
        public Builder numberOfWorkers(Integer numberOfWorkers) {
            return numberOfWorkers(Output.of(numberOfWorkers));
        }

        /**
         * @param roleArn The ARN of the IAM role associated with this job.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(@Nullable Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn The ARN of the IAM role associated with this job.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        /**
         * @param securityConfiguration The name of the Security Configuration to be associated with the job.
         * 
         * @return builder
         * 
         */
        public Builder securityConfiguration(@Nullable Output<String> securityConfiguration) {
            $.securityConfiguration = securityConfiguration;
            return this;
        }

        /**
         * @param securityConfiguration The name of the Security Configuration to be associated with the job.
         * 
         * @return builder
         * 
         */
        public Builder securityConfiguration(String securityConfiguration) {
            return securityConfiguration(Output.of(securityConfiguration));
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param timeout The job timeout in minutes. The default is 2880 minutes (48 hours) for `glueetl` and `pythonshell` jobs, and null (unlimited) for `gluestreaming` jobs.
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout The job timeout in minutes. The default is 2880 minutes (48 hours) for `glueetl` and `pythonshell` jobs, and null (unlimited) for `gluestreaming` jobs.
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        /**
         * @param workerType The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, G.2X, or G.025X for Spark jobs. Accepts the value Z.2X for Ray jobs.
         * * For the Standard worker type, each worker provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.
         * * For the G.1X worker type, each worker maps to 1 DPU (4 vCPU, 16 GB of memory, 64 GB disk), and provides 1 executor per worker. Recommended for memory-intensive jobs.
         * * For the G.2X worker type, each worker maps to 2 DPU (8 vCPU, 32 GB of memory, 128 GB disk), and provides 1 executor per worker. Recommended for memory-intensive jobs.
         * * For the G.025X worker type, each worker maps to 0.25 DPU (2 vCPU, 4GB of memory, 64 GB disk), and provides 1 executor per worker. Recommended for low volume streaming jobs. Only available for Glue version 3.0.
         * * For the Z.2X worker type, each worker maps to 2 M-DPU (8vCPU, 64 GB of m emory, 128 GB disk), and provides up to 8 Ray workers based on the autoscaler.
         * 
         * @return builder
         * 
         */
        public Builder workerType(@Nullable Output<String> workerType) {
            $.workerType = workerType;
            return this;
        }

        /**
         * @param workerType The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, G.2X, or G.025X for Spark jobs. Accepts the value Z.2X for Ray jobs.
         * * For the Standard worker type, each worker provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.
         * * For the G.1X worker type, each worker maps to 1 DPU (4 vCPU, 16 GB of memory, 64 GB disk), and provides 1 executor per worker. Recommended for memory-intensive jobs.
         * * For the G.2X worker type, each worker maps to 2 DPU (8 vCPU, 32 GB of memory, 128 GB disk), and provides 1 executor per worker. Recommended for memory-intensive jobs.
         * * For the G.025X worker type, each worker maps to 0.25 DPU (2 vCPU, 4GB of memory, 64 GB disk), and provides 1 executor per worker. Recommended for low volume streaming jobs. Only available for Glue version 3.0.
         * * For the Z.2X worker type, each worker maps to 2 M-DPU (8vCPU, 64 GB of m emory, 128 GB disk), and provides up to 8 Ray workers based on the autoscaler.
         * 
         * @return builder
         * 
         */
        public Builder workerType(String workerType) {
            return workerType(Output.of(workerType));
        }

        public JobState build() {
            return $;
        }
    }

}
