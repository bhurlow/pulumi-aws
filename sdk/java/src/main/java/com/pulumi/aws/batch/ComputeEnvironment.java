// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.batch;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.batch.ComputeEnvironmentArgs;
import com.pulumi.aws.batch.inputs.ComputeEnvironmentState;
import com.pulumi.aws.batch.outputs.ComputeEnvironmentComputeResources;
import com.pulumi.aws.batch.outputs.ComputeEnvironmentEksConfiguration;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a AWS Batch compute environment. Compute environments contain the Amazon ECS container instances that are used to run containerized batch jobs.
 * 
 * For information about AWS Batch, see [What is AWS Batch?](http://docs.aws.amazon.com/batch/latest/userguide/what-is-batch.html) .
 * For information about compute environment, see [Compute Environments](http://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html) .
 * 
 * &gt; **Note:** To prevent a race condition during environment deletion, make sure to set `depends_on` to the related `aws.iam.RolePolicyAttachment`;
 * otherwise, the policy may be destroyed too soon and the compute environment will then get stuck in the `DELETING` state, see [Troubleshooting AWS Batch](http://docs.aws.amazon.com/batch/latest/userguide/troubleshooting.html) .
 * 
 * ## Example Usage
 * ### EC2 Type
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.iam.IamFunctions;
 * import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
 * import com.pulumi.aws.iam.Role;
 * import com.pulumi.aws.iam.RoleArgs;
 * import com.pulumi.aws.iam.RolePolicyAttachment;
 * import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
 * import com.pulumi.aws.iam.InstanceProfile;
 * import com.pulumi.aws.iam.InstanceProfileArgs;
 * import com.pulumi.aws.ec2.SecurityGroup;
 * import com.pulumi.aws.ec2.SecurityGroupArgs;
 * import com.pulumi.aws.ec2.inputs.SecurityGroupEgressArgs;
 * import com.pulumi.aws.ec2.Vpc;
 * import com.pulumi.aws.ec2.VpcArgs;
 * import com.pulumi.aws.ec2.Subnet;
 * import com.pulumi.aws.ec2.SubnetArgs;
 * import com.pulumi.aws.ec2.PlacementGroup;
 * import com.pulumi.aws.ec2.PlacementGroupArgs;
 * import com.pulumi.aws.batch.ComputeEnvironment;
 * import com.pulumi.aws.batch.ComputeEnvironmentArgs;
 * import com.pulumi.aws.batch.inputs.ComputeEnvironmentComputeResourcesArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         final var ec2AssumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .statements(GetPolicyDocumentStatementArgs.builder()
 *                 .effect(&#34;Allow&#34;)
 *                 .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
 *                     .type(&#34;Service&#34;)
 *                     .identifiers(&#34;ec2.amazonaws.com&#34;)
 *                     .build())
 *                 .actions(&#34;sts:AssumeRole&#34;)
 *                 .build())
 *             .build());
 * 
 *         var ecsInstanceRoleRole = new Role(&#34;ecsInstanceRoleRole&#34;, RoleArgs.builder()        
 *             .assumeRolePolicy(ec2AssumeRole.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json()))
 *             .build());
 * 
 *         var ecsInstanceRoleRolePolicyAttachment = new RolePolicyAttachment(&#34;ecsInstanceRoleRolePolicyAttachment&#34;, RolePolicyAttachmentArgs.builder()        
 *             .role(ecsInstanceRoleRole.name())
 *             .policyArn(&#34;arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role&#34;)
 *             .build());
 * 
 *         var ecsInstanceRoleInstanceProfile = new InstanceProfile(&#34;ecsInstanceRoleInstanceProfile&#34;, InstanceProfileArgs.builder()        
 *             .role(ecsInstanceRoleRole.name())
 *             .build());
 * 
 *         final var batchAssumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .statements(GetPolicyDocumentStatementArgs.builder()
 *                 .effect(&#34;Allow&#34;)
 *                 .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
 *                     .type(&#34;Service&#34;)
 *                     .identifiers(&#34;batch.amazonaws.com&#34;)
 *                     .build())
 *                 .actions(&#34;sts:AssumeRole&#34;)
 *                 .build())
 *             .build());
 * 
 *         var awsBatchServiceRoleRole = new Role(&#34;awsBatchServiceRoleRole&#34;, RoleArgs.builder()        
 *             .assumeRolePolicy(batchAssumeRole.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json()))
 *             .build());
 * 
 *         var awsBatchServiceRoleRolePolicyAttachment = new RolePolicyAttachment(&#34;awsBatchServiceRoleRolePolicyAttachment&#34;, RolePolicyAttachmentArgs.builder()        
 *             .role(awsBatchServiceRoleRole.name())
 *             .policyArn(&#34;arn:aws:iam::aws:policy/service-role/AWSBatchServiceRole&#34;)
 *             .build());
 * 
 *         var sampleSecurityGroup = new SecurityGroup(&#34;sampleSecurityGroup&#34;, SecurityGroupArgs.builder()        
 *             .egress(SecurityGroupEgressArgs.builder()
 *                 .fromPort(0)
 *                 .toPort(0)
 *                 .protocol(&#34;-1&#34;)
 *                 .cidrBlocks(&#34;0.0.0.0/0&#34;)
 *                 .build())
 *             .build());
 * 
 *         var sampleVpc = new Vpc(&#34;sampleVpc&#34;, VpcArgs.builder()        
 *             .cidrBlock(&#34;10.1.0.0/16&#34;)
 *             .build());
 * 
 *         var sampleSubnet = new Subnet(&#34;sampleSubnet&#34;, SubnetArgs.builder()        
 *             .vpcId(sampleVpc.id())
 *             .cidrBlock(&#34;10.1.1.0/24&#34;)
 *             .build());
 * 
 *         var samplePlacementGroup = new PlacementGroup(&#34;samplePlacementGroup&#34;, PlacementGroupArgs.builder()        
 *             .strategy(&#34;cluster&#34;)
 *             .build());
 * 
 *         var sampleComputeEnvironment = new ComputeEnvironment(&#34;sampleComputeEnvironment&#34;, ComputeEnvironmentArgs.builder()        
 *             .computeEnvironmentName(&#34;sample&#34;)
 *             .computeResources(ComputeEnvironmentComputeResourcesArgs.builder()
 *                 .instanceRole(ecsInstanceRoleInstanceProfile.arn())
 *                 .instanceTypes(&#34;c4.large&#34;)
 *                 .maxVcpus(16)
 *                 .minVcpus(0)
 *                 .placementGroup(samplePlacementGroup.name())
 *                 .securityGroupIds(sampleSecurityGroup.id())
 *                 .subnets(sampleSubnet.id())
 *                 .type(&#34;EC2&#34;)
 *                 .build())
 *             .serviceRole(awsBatchServiceRoleRole.arn())
 *             .type(&#34;MANAGED&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(awsBatchServiceRoleRolePolicyAttachment)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Fargate Type
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.batch.ComputeEnvironment;
 * import com.pulumi.aws.batch.ComputeEnvironmentArgs;
 * import com.pulumi.aws.batch.inputs.ComputeEnvironmentComputeResourcesArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var sample = new ComputeEnvironment(&#34;sample&#34;, ComputeEnvironmentArgs.builder()        
 *             .computeEnvironmentName(&#34;sample&#34;)
 *             .computeResources(ComputeEnvironmentComputeResourcesArgs.builder()
 *                 .maxVcpus(16)
 *                 .securityGroupIds(aws_security_group.sample().id())
 *                 .subnets(aws_subnet.sample().id())
 *                 .type(&#34;FARGATE&#34;)
 *                 .build())
 *             .serviceRole(aws_iam_role.aws_batch_service_role().arn())
 *             .type(&#34;MANAGED&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(aws_iam_role_policy_attachment.aws_batch_service_role())
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import AWS Batch compute using the `compute_environment_name`. For example:
 * 
 * ```sh
 *  $ pulumi import aws:batch/computeEnvironment:ComputeEnvironment sample sample
 * ```
 * 
 */
@ResourceType(type="aws:batch/computeEnvironment:ComputeEnvironment")
public class ComputeEnvironment extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name (ARN) of the compute environment.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) of the compute environment.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The name for your compute environment. Up to 128 letters (uppercase and lowercase), numbers, and underscores are allowed. If omitted, the provider will assign a random, unique name.
     * 
     */
    @Export(name="computeEnvironmentName", refs={String.class}, tree="[0]")
    private Output<String> computeEnvironmentName;

    /**
     * @return The name for your compute environment. Up to 128 letters (uppercase and lowercase), numbers, and underscores are allowed. If omitted, the provider will assign a random, unique name.
     * 
     */
    public Output<String> computeEnvironmentName() {
        return this.computeEnvironmentName;
    }
    /**
     * Creates a unique compute environment name beginning with the specified prefix. Conflicts with `compute_environment_name`.
     * 
     */
    @Export(name="computeEnvironmentNamePrefix", refs={String.class}, tree="[0]")
    private Output<String> computeEnvironmentNamePrefix;

    /**
     * @return Creates a unique compute environment name beginning with the specified prefix. Conflicts with `compute_environment_name`.
     * 
     */
    public Output<String> computeEnvironmentNamePrefix() {
        return this.computeEnvironmentNamePrefix;
    }
    /**
     * Details of the compute resources managed by the compute environment. This parameter is required for managed compute environments. See details below.
     * 
     */
    @Export(name="computeResources", refs={ComputeEnvironmentComputeResources.class}, tree="[0]")
    private Output</* @Nullable */ ComputeEnvironmentComputeResources> computeResources;

    /**
     * @return Details of the compute resources managed by the compute environment. This parameter is required for managed compute environments. See details below.
     * 
     */
    public Output<Optional<ComputeEnvironmentComputeResources>> computeResources() {
        return Codegen.optional(this.computeResources);
    }
    /**
     * The Amazon Resource Name (ARN) of the underlying Amazon ECS cluster used by the compute environment.
     * 
     */
    @Export(name="ecsClusterArn", refs={String.class}, tree="[0]")
    private Output<String> ecsClusterArn;

    /**
     * @return The Amazon Resource Name (ARN) of the underlying Amazon ECS cluster used by the compute environment.
     * 
     */
    public Output<String> ecsClusterArn() {
        return this.ecsClusterArn;
    }
    /**
     * Details for the Amazon EKS cluster that supports the compute environment. See details below.
     * 
     */
    @Export(name="eksConfiguration", refs={ComputeEnvironmentEksConfiguration.class}, tree="[0]")
    private Output</* @Nullable */ ComputeEnvironmentEksConfiguration> eksConfiguration;

    /**
     * @return Details for the Amazon EKS cluster that supports the compute environment. See details below.
     * 
     */
    public Output<Optional<ComputeEnvironmentEksConfiguration>> eksConfiguration() {
        return Codegen.optional(this.eksConfiguration);
    }
    /**
     * The full Amazon Resource Name (ARN) of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.
     * 
     */
    @Export(name="serviceRole", refs={String.class}, tree="[0]")
    private Output<String> serviceRole;

    /**
     * @return The full Amazon Resource Name (ARN) of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.
     * 
     */
    public Output<String> serviceRole() {
        return this.serviceRole;
    }
    /**
     * The state of the compute environment. If the state is `ENABLED`, then the compute environment accepts jobs from a queue and can scale out automatically based on queues. Valid items are `ENABLED` or `DISABLED`. Defaults to `ENABLED`.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> state;

    /**
     * @return The state of the compute environment. If the state is `ENABLED`, then the compute environment accepts jobs from a queue and can scale out automatically based on queues. Valid items are `ENABLED` or `DISABLED`. Defaults to `ENABLED`.
     * 
     */
    public Output<Optional<String>> state() {
        return Codegen.optional(this.state);
    }
    /**
     * The current status of the compute environment (for example, CREATING or VALID).
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The current status of the compute environment (for example, CREATING or VALID).
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * A short, human-readable string to provide additional details about the current status of the compute environment.
     * 
     */
    @Export(name="statusReason", refs={String.class}, tree="[0]")
    private Output<String> statusReason;

    /**
     * @return A short, human-readable string to provide additional details about the current status of the compute environment.
     * 
     */
    public Output<String> statusReason() {
        return this.statusReason;
    }
    /**
     * Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
    /**
     * The type of the compute environment. Valid items are `MANAGED` or `UNMANAGED`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of the compute environment. Valid items are `MANAGED` or `UNMANAGED`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ComputeEnvironment(String name) {
        this(name, ComputeEnvironmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ComputeEnvironment(String name, ComputeEnvironmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ComputeEnvironment(String name, ComputeEnvironmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:batch/computeEnvironment:ComputeEnvironment", name, args == null ? ComputeEnvironmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ComputeEnvironment(String name, Output<String> id, @Nullable ComputeEnvironmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:batch/computeEnvironment:ComputeEnvironment", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "tagsAll"
            ))
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
    public static ComputeEnvironment get(String name, Output<String> id, @Nullable ComputeEnvironmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ComputeEnvironment(name, id, state, options);
    }
}
