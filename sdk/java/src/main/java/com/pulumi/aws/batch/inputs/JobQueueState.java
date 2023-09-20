// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.batch.inputs;

import com.pulumi.aws.batch.inputs.JobQueueTimeoutsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class JobQueueState extends com.pulumi.resources.ResourceArgs {

    public static final JobQueueState Empty = new JobQueueState();

    /**
     * The Amazon Resource Name of the job queue.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The Amazon Resource Name of the job queue.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Specifies the set of compute environments
     * mapped to a job queue and their order.  The position of the compute environments
     * in the list will dictate the order.
     * 
     */
    @Import(name="computeEnvironments")
    private @Nullable Output<List<String>> computeEnvironments;

    /**
     * @return Specifies the set of compute environments
     * mapped to a job queue and their order.  The position of the compute environments
     * in the list will dictate the order.
     * 
     */
    public Optional<Output<List<String>>> computeEnvironments() {
        return Optional.ofNullable(this.computeEnvironments);
    }

    /**
     * Specifies the name of the job queue.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Specifies the name of the job queue.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The priority of the job queue. Job queues with a higher priority
     * are evaluated first when associated with the same compute environment.
     * 
     */
    @Import(name="priority")
    private @Nullable Output<Integer> priority;

    /**
     * @return The priority of the job queue. Job queues with a higher priority
     * are evaluated first when associated with the same compute environment.
     * 
     */
    public Optional<Output<Integer>> priority() {
        return Optional.ofNullable(this.priority);
    }

    /**
     * The ARN of the fair share scheduling policy. If this parameter is specified, the job queue uses a fair share scheduling policy. If this parameter isn&#39;t specified, the job queue uses a first in, first out (FIFO) scheduling policy. After a job queue is created, you can replace but can&#39;t remove the fair share scheduling policy.
     * 
     */
    @Import(name="schedulingPolicyArn")
    private @Nullable Output<String> schedulingPolicyArn;

    /**
     * @return The ARN of the fair share scheduling policy. If this parameter is specified, the job queue uses a fair share scheduling policy. If this parameter isn&#39;t specified, the job queue uses a first in, first out (FIFO) scheduling policy. After a job queue is created, you can replace but can&#39;t remove the fair share scheduling policy.
     * 
     */
    public Optional<Output<String>> schedulingPolicyArn() {
        return Optional.ofNullable(this.schedulingPolicyArn);
    }

    /**
     * The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    /**
     * Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    @Import(name="timeouts")
    private @Nullable Output<JobQueueTimeoutsArgs> timeouts;

    public Optional<Output<JobQueueTimeoutsArgs>> timeouts() {
        return Optional.ofNullable(this.timeouts);
    }

    private JobQueueState() {}

    private JobQueueState(JobQueueState $) {
        this.arn = $.arn;
        this.computeEnvironments = $.computeEnvironments;
        this.name = $.name;
        this.priority = $.priority;
        this.schedulingPolicyArn = $.schedulingPolicyArn;
        this.state = $.state;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.timeouts = $.timeouts;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(JobQueueState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private JobQueueState $;

        public Builder() {
            $ = new JobQueueState();
        }

        public Builder(JobQueueState defaults) {
            $ = new JobQueueState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The Amazon Resource Name of the job queue.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The Amazon Resource Name of the job queue.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param computeEnvironments Specifies the set of compute environments
         * mapped to a job queue and their order.  The position of the compute environments
         * in the list will dictate the order.
         * 
         * @return builder
         * 
         */
        public Builder computeEnvironments(@Nullable Output<List<String>> computeEnvironments) {
            $.computeEnvironments = computeEnvironments;
            return this;
        }

        /**
         * @param computeEnvironments Specifies the set of compute environments
         * mapped to a job queue and their order.  The position of the compute environments
         * in the list will dictate the order.
         * 
         * @return builder
         * 
         */
        public Builder computeEnvironments(List<String> computeEnvironments) {
            return computeEnvironments(Output.of(computeEnvironments));
        }

        /**
         * @param computeEnvironments Specifies the set of compute environments
         * mapped to a job queue and their order.  The position of the compute environments
         * in the list will dictate the order.
         * 
         * @return builder
         * 
         */
        public Builder computeEnvironments(String... computeEnvironments) {
            return computeEnvironments(List.of(computeEnvironments));
        }

        /**
         * @param name Specifies the name of the job queue.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Specifies the name of the job queue.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param priority The priority of the job queue. Job queues with a higher priority
         * are evaluated first when associated with the same compute environment.
         * 
         * @return builder
         * 
         */
        public Builder priority(@Nullable Output<Integer> priority) {
            $.priority = priority;
            return this;
        }

        /**
         * @param priority The priority of the job queue. Job queues with a higher priority
         * are evaluated first when associated with the same compute environment.
         * 
         * @return builder
         * 
         */
        public Builder priority(Integer priority) {
            return priority(Output.of(priority));
        }

        /**
         * @param schedulingPolicyArn The ARN of the fair share scheduling policy. If this parameter is specified, the job queue uses a fair share scheduling policy. If this parameter isn&#39;t specified, the job queue uses a first in, first out (FIFO) scheduling policy. After a job queue is created, you can replace but can&#39;t remove the fair share scheduling policy.
         * 
         * @return builder
         * 
         */
        public Builder schedulingPolicyArn(@Nullable Output<String> schedulingPolicyArn) {
            $.schedulingPolicyArn = schedulingPolicyArn;
            return this;
        }

        /**
         * @param schedulingPolicyArn The ARN of the fair share scheduling policy. If this parameter is specified, the job queue uses a fair share scheduling policy. If this parameter isn&#39;t specified, the job queue uses a first in, first out (FIFO) scheduling policy. After a job queue is created, you can replace but can&#39;t remove the fair share scheduling policy.
         * 
         * @return builder
         * 
         */
        public Builder schedulingPolicyArn(String schedulingPolicyArn) {
            return schedulingPolicyArn(Output.of(schedulingPolicyArn));
        }

        /**
         * @param state The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        /**
         * @param tags Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        public Builder timeouts(@Nullable Output<JobQueueTimeoutsArgs> timeouts) {
            $.timeouts = timeouts;
            return this;
        }

        public Builder timeouts(JobQueueTimeoutsArgs timeouts) {
            return timeouts(Output.of(timeouts));
        }

        public JobQueueState build() {
            return $;
        }
    }

}
