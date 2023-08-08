// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.backup;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.backup.SelectionArgs;
import com.pulumi.aws.backup.inputs.SelectionState;
import com.pulumi.aws.backup.outputs.SelectionCondition;
import com.pulumi.aws.backup.outputs.SelectionSelectionTag;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages selection conditions for AWS Backup plan resources.
 * 
 * ## Example Usage
 * ### IAM Role
 * 
 * &gt; For more information about creating and managing IAM Roles for backups and restores, see the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/iam-service-roles.html).
 * 
 * The below example creates an IAM role with the default managed IAM Policy for allowing AWS Backup to create backups.
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
 * import com.pulumi.aws.backup.Selection;
 * import com.pulumi.aws.backup.SelectionArgs;
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
 *                     .identifiers(&#34;backup.amazonaws.com&#34;)
 *                     .build())
 *                 .actions(&#34;sts:AssumeRole&#34;)
 *                 .build())
 *             .build());
 * 
 *         var exampleRole = new Role(&#34;exampleRole&#34;, RoleArgs.builder()        
 *             .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json()))
 *             .build());
 * 
 *         var exampleRolePolicyAttachment = new RolePolicyAttachment(&#34;exampleRolePolicyAttachment&#34;, RolePolicyAttachmentArgs.builder()        
 *             .policyArn(&#34;arn:aws:iam::aws:policy/service-role/AWSBackupServiceRolePolicyForBackup&#34;)
 *             .role(exampleRole.name())
 *             .build());
 * 
 *         var exampleSelection = new Selection(&#34;exampleSelection&#34;, SelectionArgs.builder()        
 *             .iamRoleArn(exampleRole.arn())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Selecting Backups By Tag
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.backup.Selection;
 * import com.pulumi.aws.backup.SelectionArgs;
 * import com.pulumi.aws.backup.inputs.SelectionSelectionTagArgs;
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
 *         var example = new Selection(&#34;example&#34;, SelectionArgs.builder()        
 *             .iamRoleArn(aws_iam_role.example().arn())
 *             .planId(aws_backup_plan.example().id())
 *             .selectionTags(SelectionSelectionTagArgs.builder()
 *                 .type(&#34;STRINGEQUALS&#34;)
 *                 .key(&#34;foo&#34;)
 *                 .value(&#34;bar&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Selecting Backups By Conditions
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.backup.Selection;
 * import com.pulumi.aws.backup.SelectionArgs;
 * import com.pulumi.aws.backup.inputs.SelectionConditionArgs;
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
 *         var example = new Selection(&#34;example&#34;, SelectionArgs.builder()        
 *             .iamRoleArn(aws_iam_role.example().arn())
 *             .planId(aws_backup_plan.example().id())
 *             .resources(&#34;*&#34;)
 *             .conditions(SelectionConditionArgs.builder()
 *                 .stringEquals(SelectionConditionStringEqualArgs.builder()
 *                     .key(&#34;aws:ResourceTag/Component&#34;)
 *                     .value(&#34;rds&#34;)
 *                     .build())
 *                 .stringLikes(SelectionConditionStringLikeArgs.builder()
 *                     .key(&#34;aws:ResourceTag/Application&#34;)
 *                     .value(&#34;app*&#34;)
 *                     .build())
 *                 .stringNotEquals(SelectionConditionStringNotEqualArgs.builder()
 *                     .key(&#34;aws:ResourceTag/Backup&#34;)
 *                     .value(&#34;false&#34;)
 *                     .build())
 *                 .stringNotLikes(SelectionConditionStringNotLikeArgs.builder()
 *                     .key(&#34;aws:ResourceTag/Environment&#34;)
 *                     .value(&#34;test*&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Selecting Backups By Resource
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.backup.Selection;
 * import com.pulumi.aws.backup.SelectionArgs;
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
 *         var example = new Selection(&#34;example&#34;, SelectionArgs.builder()        
 *             .iamRoleArn(aws_iam_role.example().arn())
 *             .planId(aws_backup_plan.example().id())
 *             .resources(            
 *                 aws_db_instance.example().arn(),
 *                 aws_ebs_volume.example().arn(),
 *                 aws_efs_file_system.example().arn())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Selecting Backups By Not Resource
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.backup.Selection;
 * import com.pulumi.aws.backup.SelectionArgs;
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
 *         var example = new Selection(&#34;example&#34;, SelectionArgs.builder()        
 *             .iamRoleArn(aws_iam_role.example().arn())
 *             .planId(aws_backup_plan.example().id())
 *             .notResources(            
 *                 aws_db_instance.example().arn(),
 *                 aws_ebs_volume.example().arn(),
 *                 aws_efs_file_system.example().arn())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * terraform import {
 * 
 *  to = aws_backup_selection.example
 * 
 *  id = &#34;plan-id|selection-id&#34; } Using `pulumi import`, import Backup selection using the role plan_id and id separated by `|`. For exampleconsole % pulumi import aws_backup_selection.example plan-id|selection-id
 * 
 */
@ResourceType(type="aws:backup/selection:Selection")
public class Selection extends com.pulumi.resources.CustomResource {
    /**
     * A list of conditions that you define to assign resources to your backup plans using tags.
     * 
     */
    @Export(name="conditions", refs={List.class,SelectionCondition.class}, tree="[0,1]")
    private Output<List<SelectionCondition>> conditions;

    /**
     * @return A list of conditions that you define to assign resources to your backup plans using tags.
     * 
     */
    public Output<List<SelectionCondition>> conditions() {
        return this.conditions;
    }
    /**
     * The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
     * 
     */
    @Export(name="iamRoleArn", refs={String.class}, tree="[0]")
    private Output<String> iamRoleArn;

    /**
     * @return The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
     * 
     */
    public Output<String> iamRoleArn() {
        return this.iamRoleArn;
    }
    /**
     * The display name of a resource selection document.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The display name of a resource selection document.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to exclude from a backup plan.
     * 
     */
    @Export(name="notResources", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> notResources;

    /**
     * @return An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to exclude from a backup plan.
     * 
     */
    public Output<List<String>> notResources() {
        return this.notResources;
    }
    /**
     * The backup plan ID to be associated with the selection of resources.
     * 
     */
    @Export(name="planId", refs={String.class}, tree="[0]")
    private Output<String> planId;

    /**
     * @return The backup plan ID to be associated with the selection of resources.
     * 
     */
    public Output<String> planId() {
        return this.planId;
    }
    /**
     * An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan.
     * 
     */
    @Export(name="resources", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> resources;

    /**
     * @return An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan.
     * 
     */
    public Output<Optional<List<String>>> resources() {
        return Codegen.optional(this.resources);
    }
    /**
     * Tag-based conditions used to specify a set of resources to assign to a backup plan.
     * 
     */
    @Export(name="selectionTags", refs={List.class,SelectionSelectionTag.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SelectionSelectionTag>> selectionTags;

    /**
     * @return Tag-based conditions used to specify a set of resources to assign to a backup plan.
     * 
     */
    public Output<Optional<List<SelectionSelectionTag>>> selectionTags() {
        return Codegen.optional(this.selectionTags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Selection(String name) {
        this(name, SelectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Selection(String name, SelectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Selection(String name, SelectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:backup/selection:Selection", name, args == null ? SelectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Selection(String name, Output<String> id, @Nullable SelectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:backup/selection:Selection", name, state, makeResourceOptions(options, id));
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
    public static Selection get(String name, Output<String> id, @Nullable SelectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Selection(name, id, state, options);
    }
}
