// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cloudwatch.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EventRuleState extends com.pulumi.resources.ResourceArgs {

    public static final EventRuleState Empty = new EventRuleState();

    /**
     * The Amazon Resource Name (ARN) of the rule.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) of the rule.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The description of the rule.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the rule.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The name or ARN of the event bus to associate with this rule.
     * If you omit this, the `default` event bus is used.
     * 
     */
    @Import(name="eventBusName")
    private @Nullable Output<String> eventBusName;

    /**
     * @return The name or ARN of the event bus to associate with this rule.
     * If you omit this, the `default` event bus is used.
     * 
     */
    public Optional<Output<String>> eventBusName() {
        return Optional.ofNullable(this.eventBusName);
    }

    /**
     * The event pattern described a JSON object. At least one of `schedule_expression` or `event_pattern` is required. See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details.
     * 
     */
    @Import(name="eventPattern")
    private @Nullable Output<String> eventPattern;

    /**
     * @return The event pattern described a JSON object. At least one of `schedule_expression` or `event_pattern` is required. See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details.
     * 
     */
    public Optional<Output<String>> eventPattern() {
        return Optional.ofNullable(this.eventPattern);
    }

    /**
     * Whether the rule should be enabled (defaults to `true`).
     * 
     */
    @Import(name="isEnabled")
    private @Nullable Output<Boolean> isEnabled;

    /**
     * @return Whether the rule should be enabled (defaults to `true`).
     * 
     */
    public Optional<Output<Boolean>> isEnabled() {
        return Optional.ofNullable(this.isEnabled);
    }

    /**
     * The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`. **Note**: Due to the length of the generated suffix, must be 38 characters or less.
     * 
     */
    @Import(name="namePrefix")
    private @Nullable Output<String> namePrefix;

    /**
     * @return Creates a unique name beginning with the specified prefix. Conflicts with `name`. **Note**: Due to the length of the generated suffix, must be 38 characters or less.
     * 
     */
    public Optional<Output<String>> namePrefix() {
        return Optional.ofNullable(this.namePrefix);
    }

    /**
     * The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
     * 
     */
    @Import(name="roleArn")
    private @Nullable Output<String> roleArn;

    /**
     * @return The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
     * 
     */
    public Optional<Output<String>> roleArn() {
        return Optional.ofNullable(this.roleArn);
    }

    /**
     * The scheduling expression. For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`. At least one of `schedule_expression` or `event_pattern` is required. Can only be used on the default event bus. For more information, refer to the AWS documentation [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
     * 
     */
    @Import(name="scheduleExpression")
    private @Nullable Output<String> scheduleExpression;

    /**
     * @return The scheduling expression. For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`. At least one of `schedule_expression` or `event_pattern` is required. Can only be used on the default event bus. For more information, refer to the AWS documentation [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
     * 
     */
    public Optional<Output<String>> scheduleExpression() {
        return Optional.ofNullable(this.scheduleExpression);
    }

    /**
     * A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

    private EventRuleState() {}

    private EventRuleState(EventRuleState $) {
        this.arn = $.arn;
        this.description = $.description;
        this.eventBusName = $.eventBusName;
        this.eventPattern = $.eventPattern;
        this.isEnabled = $.isEnabled;
        this.name = $.name;
        this.namePrefix = $.namePrefix;
        this.roleArn = $.roleArn;
        this.scheduleExpression = $.scheduleExpression;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EventRuleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EventRuleState $;

        public Builder() {
            $ = new EventRuleState();
        }

        public Builder(EventRuleState defaults) {
            $ = new EventRuleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The Amazon Resource Name (ARN) of the rule.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The Amazon Resource Name (ARN) of the rule.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param description The description of the rule.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the rule.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param eventBusName The name or ARN of the event bus to associate with this rule.
         * If you omit this, the `default` event bus is used.
         * 
         * @return builder
         * 
         */
        public Builder eventBusName(@Nullable Output<String> eventBusName) {
            $.eventBusName = eventBusName;
            return this;
        }

        /**
         * @param eventBusName The name or ARN of the event bus to associate with this rule.
         * If you omit this, the `default` event bus is used.
         * 
         * @return builder
         * 
         */
        public Builder eventBusName(String eventBusName) {
            return eventBusName(Output.of(eventBusName));
        }

        /**
         * @param eventPattern The event pattern described a JSON object. At least one of `schedule_expression` or `event_pattern` is required. See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details.
         * 
         * @return builder
         * 
         */
        public Builder eventPattern(@Nullable Output<String> eventPattern) {
            $.eventPattern = eventPattern;
            return this;
        }

        /**
         * @param eventPattern The event pattern described a JSON object. At least one of `schedule_expression` or `event_pattern` is required. See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details.
         * 
         * @return builder
         * 
         */
        public Builder eventPattern(String eventPattern) {
            return eventPattern(Output.of(eventPattern));
        }

        /**
         * @param isEnabled Whether the rule should be enabled (defaults to `true`).
         * 
         * @return builder
         * 
         */
        public Builder isEnabled(@Nullable Output<Boolean> isEnabled) {
            $.isEnabled = isEnabled;
            return this;
        }

        /**
         * @param isEnabled Whether the rule should be enabled (defaults to `true`).
         * 
         * @return builder
         * 
         */
        public Builder isEnabled(Boolean isEnabled) {
            return isEnabled(Output.of(isEnabled));
        }

        /**
         * @param name The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namePrefix Creates a unique name beginning with the specified prefix. Conflicts with `name`. **Note**: Due to the length of the generated suffix, must be 38 characters or less.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(@Nullable Output<String> namePrefix) {
            $.namePrefix = namePrefix;
            return this;
        }

        /**
         * @param namePrefix Creates a unique name beginning with the specified prefix. Conflicts with `name`. **Note**: Due to the length of the generated suffix, must be 38 characters or less.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(String namePrefix) {
            return namePrefix(Output.of(namePrefix));
        }

        /**
         * @param roleArn The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(@Nullable Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        /**
         * @param scheduleExpression The scheduling expression. For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`. At least one of `schedule_expression` or `event_pattern` is required. Can only be used on the default event bus. For more information, refer to the AWS documentation [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
         * 
         * @return builder
         * 
         */
        public Builder scheduleExpression(@Nullable Output<String> scheduleExpression) {
            $.scheduleExpression = scheduleExpression;
            return this;
        }

        /**
         * @param scheduleExpression The scheduling expression. For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`. At least one of `schedule_expression` or `event_pattern` is required. Can only be used on the default event bus. For more information, refer to the AWS documentation [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
         * 
         * @return builder
         * 
         */
        public Builder scheduleExpression(String scheduleExpression) {
            return scheduleExpression(Output.of(scheduleExpression));
        }

        /**
         * @param tags A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public EventRuleState build() {
            return $;
        }
    }

}
