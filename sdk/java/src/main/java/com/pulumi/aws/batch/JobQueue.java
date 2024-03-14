// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.batch;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.batch.JobQueueArgs;
import com.pulumi.aws.batch.inputs.JobQueueState;
import com.pulumi.aws.batch.outputs.JobQueueComputeEnvironmentOrder;
import com.pulumi.aws.batch.outputs.JobQueueTimeouts;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Batch Job Queue resource.
 * 
 * ## Example Usage
 * 
 * ### Basic Job Queue
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.batch.JobQueue;
 * import com.pulumi.aws.batch.JobQueueArgs;
 * import com.pulumi.aws.batch.inputs.JobQueueComputeEnvironmentOrderArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var testQueue = new JobQueue(&#34;testQueue&#34;, JobQueueArgs.builder()        
 *             .name(&#34;tf-test-batch-job-queue&#34;)
 *             .state(&#34;ENABLED&#34;)
 *             .priority(1)
 *             .computeEnvironmentOrders(            
 *                 JobQueueComputeEnvironmentOrderArgs.builder()
 *                     .order(1)
 *                     .computeEnvironment(testEnvironment1.arn())
 *                     .build(),
 *                 JobQueueComputeEnvironmentOrderArgs.builder()
 *                     .order(2)
 *                     .computeEnvironment(testEnvironment2.arn())
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Job Queue with a fair share scheduling policy
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.batch.SchedulingPolicy;
 * import com.pulumi.aws.batch.SchedulingPolicyArgs;
 * import com.pulumi.aws.batch.inputs.SchedulingPolicyFairSharePolicyArgs;
 * import com.pulumi.aws.batch.JobQueue;
 * import com.pulumi.aws.batch.JobQueueArgs;
 * import com.pulumi.aws.batch.inputs.JobQueueComputeEnvironmentOrderArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new SchedulingPolicy(&#34;example&#34;, SchedulingPolicyArgs.builder()        
 *             .name(&#34;example&#34;)
 *             .fairSharePolicy(SchedulingPolicyFairSharePolicyArgs.builder()
 *                 .computeReservation(1)
 *                 .shareDecaySeconds(3600)
 *                 .shareDistributions(SchedulingPolicyFairSharePolicyShareDistributionArgs.builder()
 *                     .shareIdentifier(&#34;A1*&#34;)
 *                     .weightFactor(0.1)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var exampleJobQueue = new JobQueue(&#34;exampleJobQueue&#34;, JobQueueArgs.builder()        
 *             .name(&#34;tf-test-batch-job-queue&#34;)
 *             .schedulingPolicyArn(example.arn())
 *             .state(&#34;ENABLED&#34;)
 *             .priority(1)
 *             .computeEnvironmentOrders(            
 *                 JobQueueComputeEnvironmentOrderArgs.builder()
 *                     .order(1)
 *                     .computeEnvironment(testEnvironment1.arn())
 *                     .build(),
 *                 JobQueueComputeEnvironmentOrderArgs.builder()
 *                     .order(2)
 *                     .computeEnvironment(testEnvironment2.arn())
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Using `pulumi import`, import Batch Job Queue using the `arn`. For example:
 * 
 * ```sh
 * $ pulumi import aws:batch/jobQueue:JobQueue test_queue arn:aws:batch:us-east-1:123456789012:job-queue/sample
 * ```
 * 
 */
@ResourceType(type="aws:batch/jobQueue:JobQueue")
public class JobQueue extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name of the job queue.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name of the job queue.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The set of compute environments mapped to a job queue and their order relative to each other. The job scheduler uses this parameter to determine which compute environment runs a specific job. Compute environments must be in the VALID state before you can associate them with a job queue. You can associate up to three compute environments with a job queue.
     * 
     */
    @Export(name="computeEnvironmentOrders", refs={List.class,JobQueueComputeEnvironmentOrder.class}, tree="[0,1]")
    private Output</* @Nullable */ List<JobQueueComputeEnvironmentOrder>> computeEnvironmentOrders;

    /**
     * @return The set of compute environments mapped to a job queue and their order relative to each other. The job scheduler uses this parameter to determine which compute environment runs a specific job. Compute environments must be in the VALID state before you can associate them with a job queue. You can associate up to three compute environments with a job queue.
     * 
     */
    public Output<Optional<List<JobQueueComputeEnvironmentOrder>>> computeEnvironmentOrders() {
        return Codegen.optional(this.computeEnvironmentOrders);
    }
    /**
     * (Optional) This parameter is deprecated, please use `compute_environment_order` instead. List of compute environment ARNs mapped to a job queue. The position of the compute environments in the list will dictate the order. When importing a AWS Batch Job Queue, the parameter `compute_environments` will always be used over `compute_environment_order`. Please adjust your HCL accordingly.
     * 
     * @deprecated
     * This parameter will be replaced by `compute_environments_order`.
     * 
     */
    @Deprecated /* This parameter will be replaced by `compute_environments_order`. */
    @Export(name="computeEnvironments", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> computeEnvironments;

    /**
     * @return (Optional) This parameter is deprecated, please use `compute_environment_order` instead. List of compute environment ARNs mapped to a job queue. The position of the compute environments in the list will dictate the order. When importing a AWS Batch Job Queue, the parameter `compute_environments` will always be used over `compute_environment_order`. Please adjust your HCL accordingly.
     * 
     */
    public Output<Optional<List<String>>> computeEnvironments() {
        return Codegen.optional(this.computeEnvironments);
    }
    /**
     * Specifies the name of the job queue.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Specifies the name of the job queue.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The priority of the job queue. Job queues with a higher priority
     * are evaluated first when associated with the same compute environment.
     * 
     */
    @Export(name="priority", refs={Integer.class}, tree="[0]")
    private Output<Integer> priority;

    /**
     * @return The priority of the job queue. Job queues with a higher priority
     * are evaluated first when associated with the same compute environment.
     * 
     */
    public Output<Integer> priority() {
        return this.priority;
    }
    /**
     * The ARN of the fair share scheduling policy. If this parameter is specified, the job queue uses a fair share scheduling policy. If this parameter isn&#39;t specified, the job queue uses a first in, first out (FIFO) scheduling policy. After a job queue is created, you can replace but can&#39;t remove the fair share scheduling policy.
     * 
     */
    @Export(name="schedulingPolicyArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> schedulingPolicyArn;

    /**
     * @return The ARN of the fair share scheduling policy. If this parameter is specified, the job queue uses a fair share scheduling policy. If this parameter isn&#39;t specified, the job queue uses a first in, first out (FIFO) scheduling policy. After a job queue is created, you can replace but can&#39;t remove the fair share scheduling policy.
     * 
     */
    public Output<Optional<String>> schedulingPolicyArn() {
        return Codegen.optional(this.schedulingPolicyArn);
    }
    /**
     * The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    @Export(name="timeouts", refs={JobQueueTimeouts.class}, tree="[0]")
    private Output</* @Nullable */ JobQueueTimeouts> timeouts;

    public Output<Optional<JobQueueTimeouts>> timeouts() {
        return Codegen.optional(this.timeouts);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public JobQueue(String name) {
        this(name, JobQueueArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public JobQueue(String name, JobQueueArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public JobQueue(String name, JobQueueArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:batch/jobQueue:JobQueue", name, args == null ? JobQueueArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private JobQueue(String name, Output<String> id, @Nullable JobQueueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:batch/jobQueue:JobQueue", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static JobQueue get(String name, Output<String> id, @Nullable JobQueueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new JobQueue(name, id, state, options);
    }
}
