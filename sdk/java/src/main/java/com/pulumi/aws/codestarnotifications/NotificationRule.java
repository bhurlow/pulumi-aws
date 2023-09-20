// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.codestarnotifications;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.codestarnotifications.NotificationRuleArgs;
import com.pulumi.aws.codestarnotifications.inputs.NotificationRuleState;
import com.pulumi.aws.codestarnotifications.outputs.NotificationRuleTarget;
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
 * Provides a CodeStar Notifications Rule.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.codecommit.Repository;
 * import com.pulumi.aws.codecommit.RepositoryArgs;
 * import com.pulumi.aws.sns.Topic;
 * import com.pulumi.aws.iam.IamFunctions;
 * import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
 * import com.pulumi.aws.sns.TopicPolicy;
 * import com.pulumi.aws.sns.TopicPolicyArgs;
 * import com.pulumi.aws.codestarnotifications.NotificationRule;
 * import com.pulumi.aws.codestarnotifications.NotificationRuleArgs;
 * import com.pulumi.aws.codestarnotifications.inputs.NotificationRuleTargetArgs;
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
 *         var code = new Repository(&#34;code&#34;, RepositoryArgs.builder()        
 *             .repositoryName(&#34;example-code-repo&#34;)
 *             .build());
 * 
 *         var notif = new Topic(&#34;notif&#34;);
 * 
 *         final var notifAccess = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .statements(GetPolicyDocumentStatementArgs.builder()
 *                 .actions(&#34;sns:Publish&#34;)
 *                 .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
 *                     .type(&#34;Service&#34;)
 *                     .identifiers(&#34;codestar-notifications.amazonaws.com&#34;)
 *                     .build())
 *                 .resources(notif.arn())
 *                 .build())
 *             .build());
 * 
 *         var default_ = new TopicPolicy(&#34;default&#34;, TopicPolicyArgs.builder()        
 *             .arn(notif.arn())
 *             .policy(notifAccess.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult).applyValue(notifAccess -&gt; notifAccess.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json())))
 *             .build());
 * 
 *         var commits = new NotificationRule(&#34;commits&#34;, NotificationRuleArgs.builder()        
 *             .detailType(&#34;BASIC&#34;)
 *             .eventTypeIds(&#34;codecommit-repository-comments-on-commits&#34;)
 *             .resource(code.arn())
 *             .targets(NotificationRuleTargetArgs.builder()
 *                 .address(notif.arn())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * In TODO v1.5.0 and later, use an `import` block to import CodeStar notification rule using the ARN. For exampleterraform import {
 * 
 *  to = aws_codestarnotifications_notification_rule.foo
 * 
 *  id = &#34;arn:aws:codestar-notifications:us-west-1:0123456789:notificationrule/2cdc68a3-8f7c-4893-b6a5-45b362bd4f2b&#34; } Using `TODO import`, import CodeStar notification rule using the ARN. For exampleconsole % TODO import aws_codestarnotifications_notification_rule.foo arn:aws:codestar-notifications:us-west-1:0123456789:notificationrule/2cdc68a3-8f7c-4893-b6a5-45b362bd4f2b
 * 
 */
@ResourceType(type="aws:codestarnotifications/notificationRule:NotificationRule")
public class NotificationRule extends com.pulumi.resources.CustomResource {
    /**
     * The codestar notification rule ARN.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The codestar notification rule ARN.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The level of detail to include in the notifications for this resource. Possible values are `BASIC` and `FULL`.
     * 
     */
    @Export(name="detailType", refs={String.class}, tree="[0]")
    private Output<String> detailType;

    /**
     * @return The level of detail to include in the notifications for this resource. Possible values are `BASIC` and `FULL`.
     * 
     */
    public Output<String> detailType() {
        return this.detailType;
    }
    /**
     * A list of event types associated with this notification rule.
     * For list of allowed events see [here](https://docs.aws.amazon.com/codestar-notifications/latest/userguide/concepts.html#concepts-api).
     * 
     */
    @Export(name="eventTypeIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> eventTypeIds;

    /**
     * @return A list of event types associated with this notification rule.
     * For list of allowed events see [here](https://docs.aws.amazon.com/codestar-notifications/latest/userguide/concepts.html#concepts-api).
     * 
     */
    public Output<List<String>> eventTypeIds() {
        return this.eventTypeIds;
    }
    /**
     * The name of notification rule.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of notification rule.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ARN of the resource to associate with the notification rule.
     * 
     */
    @Export(name="resource", refs={String.class}, tree="[0]")
    private Output<String> resource;

    /**
     * @return The ARN of the resource to associate with the notification rule.
     * 
     */
    public Output<String> resource() {
        return this.resource;
    }
    /**
     * The status of the notification rule. Possible values are `ENABLED` and `DISABLED`, default is `ENABLED`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> status;

    /**
     * @return The status of the notification rule. Possible values are `ENABLED` and `DISABLED`, default is `ENABLED`.
     * 
     */
    public Output<Optional<String>> status() {
        return Codegen.optional(this.status);
    }
    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * Configuration blocks containing notification target information. Can be specified multiple times. At least one target must be specified on creation.
     * 
     */
    @Export(name="targets", refs={List.class,NotificationRuleTarget.class}, tree="[0,1]")
    private Output</* @Nullable */ List<NotificationRuleTarget>> targets;

    /**
     * @return Configuration blocks containing notification target information. Can be specified multiple times. At least one target must be specified on creation.
     * 
     */
    public Output<Optional<List<NotificationRuleTarget>>> targets() {
        return Codegen.optional(this.targets);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NotificationRule(String name) {
        this(name, NotificationRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NotificationRule(String name, NotificationRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NotificationRule(String name, NotificationRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:codestarnotifications/notificationRule:NotificationRule", name, args == null ? NotificationRuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NotificationRule(String name, Output<String> id, @Nullable NotificationRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:codestarnotifications/notificationRule:NotificationRule", name, state, makeResourceOptions(options, id));
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
    public static NotificationRule get(String name, Output<String> id, @Nullable NotificationRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NotificationRule(name, id, state, options);
    }
}
