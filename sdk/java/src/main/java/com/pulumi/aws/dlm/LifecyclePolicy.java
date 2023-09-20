// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.dlm;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.dlm.LifecyclePolicyArgs;
import com.pulumi.aws.dlm.inputs.LifecyclePolicyState;
import com.pulumi.aws.dlm.outputs.LifecyclePolicyPolicyDetails;
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
 * Provides a [Data Lifecycle Manager (DLM) lifecycle policy](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshot-lifecycle.html) for managing snapshots.
 * 
 * ## Example Usage
 * ### Basic
 * 
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
 * import com.pulumi.aws.iam.RolePolicy;
 * import com.pulumi.aws.iam.RolePolicyArgs;
 * import com.pulumi.aws.dlm.LifecyclePolicy;
 * import com.pulumi.aws.dlm.LifecyclePolicyArgs;
 * import com.pulumi.aws.dlm.inputs.LifecyclePolicyPolicyDetailsArgs;
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
 *         final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .statements(GetPolicyDocumentStatementArgs.builder()
 *                 .effect(&#34;Allow&#34;)
 *                 .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
 *                     .type(&#34;Service&#34;)
 *                     .identifiers(&#34;dlm.amazonaws.com&#34;)
 *                     .build())
 *                 .actions(&#34;sts:AssumeRole&#34;)
 *                 .build())
 *             .build());
 * 
 *         var dlmLifecycleRole = new Role(&#34;dlmLifecycleRole&#34;, RoleArgs.builder()        
 *             .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json()))
 *             .build());
 * 
 *         final var dlmLifecyclePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .statements(            
 *                 GetPolicyDocumentStatementArgs.builder()
 *                     .effect(&#34;Allow&#34;)
 *                     .actions(                    
 *                         &#34;ec2:CreateSnapshot&#34;,
 *                         &#34;ec2:CreateSnapshots&#34;,
 *                         &#34;ec2:DeleteSnapshot&#34;,
 *                         &#34;ec2:DescribeInstances&#34;,
 *                         &#34;ec2:DescribeVolumes&#34;,
 *                         &#34;ec2:DescribeSnapshots&#34;)
 *                     .resources(&#34;*&#34;)
 *                     .build(),
 *                 GetPolicyDocumentStatementArgs.builder()
 *                     .effect(&#34;Allow&#34;)
 *                     .actions(&#34;ec2:CreateTags&#34;)
 *                     .resources(&#34;arn:aws:ec2:*::snapshot/*&#34;)
 *                     .build())
 *             .build());
 * 
 *         var dlmLifecycleRolePolicy = new RolePolicy(&#34;dlmLifecycleRolePolicy&#34;, RolePolicyArgs.builder()        
 *             .role(dlmLifecycleRole.id())
 *             .policy(dlmLifecyclePolicyDocument.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json()))
 *             .build());
 * 
 *         var example = new LifecyclePolicy(&#34;example&#34;, LifecyclePolicyArgs.builder()        
 *             .description(&#34;example DLM lifecycle policy&#34;)
 *             .executionRoleArn(dlmLifecycleRole.arn())
 *             .state(&#34;ENABLED&#34;)
 *             .policyDetails(LifecyclePolicyPolicyDetailsArgs.builder()
 *                 .resourceTypes(&#34;VOLUME&#34;)
 *                 .schedules(LifecyclePolicyPolicyDetailsScheduleArgs.builder()
 *                     .name(&#34;2 weeks of daily snapshots&#34;)
 *                     .createRule(LifecyclePolicyPolicyDetailsScheduleCreateRuleArgs.builder()
 *                         .interval(24)
 *                         .intervalUnit(&#34;HOURS&#34;)
 *                         .times(&#34;23:45&#34;)
 *                         .build())
 *                     .retainRule(LifecyclePolicyPolicyDetailsScheduleRetainRuleArgs.builder()
 *                         .count(14)
 *                         .build())
 *                     .tagsToAdd(Map.of(&#34;SnapshotCreator&#34;, &#34;DLM&#34;))
 *                     .copyTags(false)
 *                     .build())
 *                 .targetTags(Map.of(&#34;Snapshot&#34;, &#34;true&#34;))
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Example Cross-Region Snapshot Copy Usage
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.AwsFunctions;
 * import com.pulumi.aws.inputs.GetCallerIdentityArgs;
 * import com.pulumi.aws.iam.IamFunctions;
 * import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
 * import com.pulumi.aws.kms.Key;
 * import com.pulumi.aws.kms.KeyArgs;
 * import com.pulumi.aws.dlm.LifecyclePolicy;
 * import com.pulumi.aws.dlm.LifecyclePolicyArgs;
 * import com.pulumi.aws.dlm.inputs.LifecyclePolicyPolicyDetailsArgs;
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
 *         final var current = AwsFunctions.getCallerIdentity();
 * 
 *         final var key = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .statements(GetPolicyDocumentStatementArgs.builder()
 *                 .sid(&#34;Enable IAM User Permissions&#34;)
 *                 .effect(&#34;Allow&#34;)
 *                 .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
 *                     .type(&#34;AWS&#34;)
 *                     .identifiers(String.format(&#34;arn:aws:iam::%s:root&#34;, current.applyValue(getCallerIdentityResult -&gt; getCallerIdentityResult.accountId())))
 *                     .build())
 *                 .actions(&#34;kms:*&#34;)
 *                 .resources(&#34;*&#34;)
 *                 .build())
 *             .build());
 * 
 *         var dlmCrossRegionCopyCmk = new Key(&#34;dlmCrossRegionCopyCmk&#34;, KeyArgs.builder()        
 *             .description(&#34;Example Alternate Region KMS Key&#34;)
 *             .policy(key.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json()))
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(aws.alternate())
 *                 .build());
 * 
 *         var example = new LifecyclePolicy(&#34;example&#34;, LifecyclePolicyArgs.builder()        
 *             .description(&#34;example DLM lifecycle policy&#34;)
 *             .executionRoleArn(aws_iam_role.dlm_lifecycle_role().arn())
 *             .state(&#34;ENABLED&#34;)
 *             .policyDetails(LifecyclePolicyPolicyDetailsArgs.builder()
 *                 .resourceTypes(&#34;VOLUME&#34;)
 *                 .schedules(LifecyclePolicyPolicyDetailsScheduleArgs.builder()
 *                     .name(&#34;2 weeks of daily snapshots&#34;)
 *                     .createRule(LifecyclePolicyPolicyDetailsScheduleCreateRuleArgs.builder()
 *                         .interval(24)
 *                         .intervalUnit(&#34;HOURS&#34;)
 *                         .times(&#34;23:45&#34;)
 *                         .build())
 *                     .retainRule(LifecyclePolicyPolicyDetailsScheduleRetainRuleArgs.builder()
 *                         .count(14)
 *                         .build())
 *                     .tagsToAdd(Map.of(&#34;SnapshotCreator&#34;, &#34;DLM&#34;))
 *                     .copyTags(false)
 *                     .crossRegionCopyRules(LifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleArgs.builder()
 *                         .target(&#34;us-west-2&#34;)
 *                         .encrypted(true)
 *                         .cmkArn(dlmCrossRegionCopyCmk.arn())
 *                         .copyTags(true)
 *                         .retainRule(LifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleArgs.builder()
 *                             .interval(30)
 *                             .intervalUnit(&#34;DAYS&#34;)
 *                             .build())
 *                         .build())
 *                     .build())
 *                 .targetTags(Map.of(&#34;Snapshot&#34;, &#34;true&#34;))
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Example Event Based Policy Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.AwsFunctions;
 * import com.pulumi.aws.inputs.GetCallerIdentityArgs;
 * import com.pulumi.aws.dlm.LifecyclePolicy;
 * import com.pulumi.aws.dlm.LifecyclePolicyArgs;
 * import com.pulumi.aws.dlm.inputs.LifecyclePolicyPolicyDetailsArgs;
 * import com.pulumi.aws.dlm.inputs.LifecyclePolicyPolicyDetailsActionArgs;
 * import com.pulumi.aws.dlm.inputs.LifecyclePolicyPolicyDetailsEventSourceArgs;
 * import com.pulumi.aws.dlm.inputs.LifecyclePolicyPolicyDetailsEventSourceParametersArgs;
 * import com.pulumi.aws.iam.IamFunctions;
 * import com.pulumi.aws.iam.inputs.GetPolicyArgs;
 * import com.pulumi.aws.iam.RolePolicyAttachment;
 * import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
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
 *         final var current = AwsFunctions.getCallerIdentity();
 * 
 *         var exampleLifecyclePolicy = new LifecyclePolicy(&#34;exampleLifecyclePolicy&#34;, LifecyclePolicyArgs.builder()        
 *             .description(&#34;tf-acc-basic&#34;)
 *             .executionRoleArn(aws_iam_role.example().arn())
 *             .policyDetails(LifecyclePolicyPolicyDetailsArgs.builder()
 *                 .policyType(&#34;EVENT_BASED_POLICY&#34;)
 *                 .action(LifecyclePolicyPolicyDetailsActionArgs.builder()
 *                     .name(&#34;tf-acc-basic&#34;)
 *                     .crossRegionCopies(LifecyclePolicyPolicyDetailsActionCrossRegionCopyArgs.builder()
 *                         .encryptionConfiguration()
 *                         .retainRule(LifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleArgs.builder()
 *                             .interval(15)
 *                             .intervalUnit(&#34;MONTHS&#34;)
 *                             .build())
 *                         .target(&#34;us-east-1&#34;)
 *                         .build())
 *                     .build())
 *                 .eventSource(LifecyclePolicyPolicyDetailsEventSourceArgs.builder()
 *                     .type(&#34;MANAGED_CWE&#34;)
 *                     .parameters(LifecyclePolicyPolicyDetailsEventSourceParametersArgs.builder()
 *                         .descriptionRegex(&#34;^.*Created for policy: policy-1234567890abcdef0.*$&#34;)
 *                         .eventType(&#34;shareSnapshot&#34;)
 *                         .snapshotOwners(current.applyValue(getCallerIdentityResult -&gt; getCallerIdentityResult.accountId()))
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         final var examplePolicy = IamFunctions.getPolicy(GetPolicyArgs.builder()
 *             .name(&#34;AWSDataLifecycleManagerServiceRole&#34;)
 *             .build());
 * 
 *         var exampleRolePolicyAttachment = new RolePolicyAttachment(&#34;exampleRolePolicyAttachment&#34;, RolePolicyAttachmentArgs.builder()        
 *             .role(aws_iam_role.example().id())
 *             .policyArn(examplePolicy.applyValue(getPolicyResult -&gt; getPolicyResult.arn()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * In TODO v1.5.0 and later, use an `import` block to import DLM lifecycle policies using their policy ID. For exampleterraform import {
 * 
 *  to = aws_dlm_lifecycle_policy.example
 * 
 *  id = &#34;policy-abcdef12345678901&#34; } Using `TODO import`, import DLM lifecycle policies using their policy ID. For exampleconsole % TODO import aws_dlm_lifecycle_policy.example policy-abcdef12345678901
 * 
 */
@ResourceType(type="aws:dlm/lifecyclePolicy:LifecyclePolicy")
public class LifecyclePolicy extends com.pulumi.resources.CustomResource {
    /**
     * Amazon Resource Name (ARN) of the DLM Lifecycle Policy.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of the DLM Lifecycle Policy.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * A description for the DLM lifecycle policy.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return A description for the DLM lifecycle policy.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The ARN of an IAM role that is able to be assumed by the DLM service.
     * 
     */
    @Export(name="executionRoleArn", refs={String.class}, tree="[0]")
    private Output<String> executionRoleArn;

    /**
     * @return The ARN of an IAM role that is able to be assumed by the DLM service.
     * 
     */
    public Output<String> executionRoleArn() {
        return this.executionRoleArn;
    }
    /**
     * See the `policy_details` configuration block. Max of 1.
     * 
     */
    @Export(name="policyDetails", refs={LifecyclePolicyPolicyDetails.class}, tree="[0]")
    private Output<LifecyclePolicyPolicyDetails> policyDetails;

    /**
     * @return See the `policy_details` configuration block. Max of 1.
     * 
     */
    public Output<LifecyclePolicyPolicyDetails> policyDetails() {
        return this.policyDetails;
    }
    /**
     * Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> state;

    /**
     * @return Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
     * 
     */
    public Output<Optional<String>> state() {
        return Codegen.optional(this.state);
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LifecyclePolicy(String name) {
        this(name, LifecyclePolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LifecyclePolicy(String name, LifecyclePolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LifecyclePolicy(String name, LifecyclePolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:dlm/lifecyclePolicy:LifecyclePolicy", name, args == null ? LifecyclePolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LifecyclePolicy(String name, Output<String> id, @Nullable LifecyclePolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:dlm/lifecyclePolicy:LifecyclePolicy", name, state, makeResourceOptions(options, id));
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
    public static LifecyclePolicy get(String name, Output<String> id, @Nullable LifecyclePolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LifecyclePolicy(name, id, state, options);
    }
}
