// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.vpclattice.inputs;

import com.pulumi.aws.vpclattice.inputs.ListenerRuleActionArgs;
import com.pulumi.aws.vpclattice.inputs.ListenerRuleMatchArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ListenerRuleState extends com.pulumi.resources.ResourceArgs {

    public static final ListenerRuleState Empty = new ListenerRuleState();

    /**
     * The action for the default rule.
     * 
     */
    @Import(name="action")
    private @Nullable Output<ListenerRuleActionArgs> action;

    /**
     * @return The action for the default rule.
     * 
     */
    public Optional<Output<ListenerRuleActionArgs>> action() {
        return Optional.ofNullable(this.action);
    }

    /**
     * ARN of the target group.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return ARN of the target group.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The ID or Amazon Resource Name (ARN) of the listener.
     * 
     */
    @Import(name="listenerIdentifier")
    private @Nullable Output<String> listenerIdentifier;

    /**
     * @return The ID or Amazon Resource Name (ARN) of the listener.
     * 
     */
    public Optional<Output<String>> listenerIdentifier() {
        return Optional.ofNullable(this.listenerIdentifier);
    }

    /**
     * The rule match.
     * 
     */
    @Import(name="match")
    private @Nullable Output<ListenerRuleMatchArgs> match;

    /**
     * @return The rule match.
     * 
     */
    public Optional<Output<ListenerRuleMatchArgs>> match() {
        return Optional.ofNullable(this.match);
    }

    /**
     * The name of the rule. The name must be unique within the listener. The valid characters are a-z, 0-9, and hyphens (-). You can&#39;t use a hyphen as the first or last character, or immediately after another hyphen.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the rule. The name must be unique within the listener. The valid characters are a-z, 0-9, and hyphens (-). You can&#39;t use a hyphen as the first or last character, or immediately after another hyphen.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The priority assigned to the rule. Each rule for a specific listener must have a unique priority. The lower the priority number the higher the priority.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="priority")
    private @Nullable Output<Integer> priority;

    /**
     * @return The priority assigned to the rule. Each rule for a specific listener must have a unique priority. The lower the priority number the higher the priority.
     * 
     * The following arguments are optional:
     * 
     */
    public Optional<Output<Integer>> priority() {
        return Optional.ofNullable(this.priority);
    }

    /**
     * Unique identifier for the target group.
     * 
     */
    @Import(name="ruleId")
    private @Nullable Output<String> ruleId;

    /**
     * @return Unique identifier for the target group.
     * 
     */
    public Optional<Output<String>> ruleId() {
        return Optional.ofNullable(this.ruleId);
    }

    /**
     * The ID or Amazon Resource Identifier (ARN) of the service.
     * 
     */
    @Import(name="serviceIdentifier")
    private @Nullable Output<String> serviceIdentifier;

    /**
     * @return The ID or Amazon Resource Identifier (ARN) of the service.
     * 
     */
    public Optional<Output<String>> serviceIdentifier() {
        return Optional.ofNullable(this.serviceIdentifier);
    }

    /**
     * Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    private ListenerRuleState() {}

    private ListenerRuleState(ListenerRuleState $) {
        this.action = $.action;
        this.arn = $.arn;
        this.listenerIdentifier = $.listenerIdentifier;
        this.match = $.match;
        this.name = $.name;
        this.priority = $.priority;
        this.ruleId = $.ruleId;
        this.serviceIdentifier = $.serviceIdentifier;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ListenerRuleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ListenerRuleState $;

        public Builder() {
            $ = new ListenerRuleState();
        }

        public Builder(ListenerRuleState defaults) {
            $ = new ListenerRuleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param action The action for the default rule.
         * 
         * @return builder
         * 
         */
        public Builder action(@Nullable Output<ListenerRuleActionArgs> action) {
            $.action = action;
            return this;
        }

        /**
         * @param action The action for the default rule.
         * 
         * @return builder
         * 
         */
        public Builder action(ListenerRuleActionArgs action) {
            return action(Output.of(action));
        }

        /**
         * @param arn ARN of the target group.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn ARN of the target group.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param listenerIdentifier The ID or Amazon Resource Name (ARN) of the listener.
         * 
         * @return builder
         * 
         */
        public Builder listenerIdentifier(@Nullable Output<String> listenerIdentifier) {
            $.listenerIdentifier = listenerIdentifier;
            return this;
        }

        /**
         * @param listenerIdentifier The ID or Amazon Resource Name (ARN) of the listener.
         * 
         * @return builder
         * 
         */
        public Builder listenerIdentifier(String listenerIdentifier) {
            return listenerIdentifier(Output.of(listenerIdentifier));
        }

        /**
         * @param match The rule match.
         * 
         * @return builder
         * 
         */
        public Builder match(@Nullable Output<ListenerRuleMatchArgs> match) {
            $.match = match;
            return this;
        }

        /**
         * @param match The rule match.
         * 
         * @return builder
         * 
         */
        public Builder match(ListenerRuleMatchArgs match) {
            return match(Output.of(match));
        }

        /**
         * @param name The name of the rule. The name must be unique within the listener. The valid characters are a-z, 0-9, and hyphens (-). You can&#39;t use a hyphen as the first or last character, or immediately after another hyphen.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the rule. The name must be unique within the listener. The valid characters are a-z, 0-9, and hyphens (-). You can&#39;t use a hyphen as the first or last character, or immediately after another hyphen.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param priority The priority assigned to the rule. Each rule for a specific listener must have a unique priority. The lower the priority number the higher the priority.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder priority(@Nullable Output<Integer> priority) {
            $.priority = priority;
            return this;
        }

        /**
         * @param priority The priority assigned to the rule. Each rule for a specific listener must have a unique priority. The lower the priority number the higher the priority.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder priority(Integer priority) {
            return priority(Output.of(priority));
        }

        /**
         * @param ruleId Unique identifier for the target group.
         * 
         * @return builder
         * 
         */
        public Builder ruleId(@Nullable Output<String> ruleId) {
            $.ruleId = ruleId;
            return this;
        }

        /**
         * @param ruleId Unique identifier for the target group.
         * 
         * @return builder
         * 
         */
        public Builder ruleId(String ruleId) {
            return ruleId(Output.of(ruleId));
        }

        /**
         * @param serviceIdentifier The ID or Amazon Resource Identifier (ARN) of the service.
         * 
         * @return builder
         * 
         */
        public Builder serviceIdentifier(@Nullable Output<String> serviceIdentifier) {
            $.serviceIdentifier = serviceIdentifier;
            return this;
        }

        /**
         * @param serviceIdentifier The ID or Amazon Resource Identifier (ARN) of the service.
         * 
         * @return builder
         * 
         */
        public Builder serviceIdentifier(String serviceIdentifier) {
            return serviceIdentifier(Output.of(serviceIdentifier));
        }

        /**
         * @param tags Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
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
         * @param tagsAll Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
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

        public ListenerRuleState build() {
            return $;
        }
    }

}
