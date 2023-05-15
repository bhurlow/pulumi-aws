// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.docdb;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EventSubscriptionArgs extends com.pulumi.resources.ResourceArgs {

    public static final EventSubscriptionArgs Empty = new EventSubscriptionArgs();

    /**
     * A boolean flag to enable/disable the subscription. Defaults to true.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return A boolean flag to enable/disable the subscription. Defaults to true.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Event.html or run `aws docdb describe-event-categories`.
     * 
     */
    @Import(name="eventCategories")
    private @Nullable Output<List<String>> eventCategories;

    /**
     * @return A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Event.html or run `aws docdb describe-event-categories`.
     * 
     */
    public Optional<Output<List<String>>> eventCategories() {
        return Optional.ofNullable(this.eventCategories);
    }

    /**
     * The name of the DocumentDB event subscription. By default generated by this provider.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the DocumentDB event subscription. By default generated by this provider.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The name of the DocumentDB event subscription. Conflicts with `name`.
     * 
     */
    @Import(name="namePrefix")
    private @Nullable Output<String> namePrefix;

    /**
     * @return The name of the DocumentDB event subscription. Conflicts with `name`.
     * 
     */
    public Optional<Output<String>> namePrefix() {
        return Optional.ofNullable(this.namePrefix);
    }

    @Import(name="snsTopicArn", required=true)
    private Output<String> snsTopicArn;

    public Output<String> snsTopicArn() {
        return this.snsTopicArn;
    }

    /**
     * A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a source_type must also be specified.
     * 
     */
    @Import(name="sourceIds")
    private @Nullable Output<List<String>> sourceIds;

    /**
     * @return A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a source_type must also be specified.
     * 
     */
    public Optional<Output<List<String>>> sourceIds() {
        return Optional.ofNullable(this.sourceIds);
    }

    /**
     * The type of source that will be generating the events. Valid options are `db-instance`, `db-cluster`, `db-parameter-group`, `db-security-group`,`  db-cluster-snapshot `. If not set, all sources will be subscribed to.
     * 
     */
    @Import(name="sourceType")
    private @Nullable Output<String> sourceType;

    /**
     * @return The type of source that will be generating the events. Valid options are `db-instance`, `db-cluster`, `db-parameter-group`, `db-security-group`,`  db-cluster-snapshot `. If not set, all sources will be subscribed to.
     * 
     */
    public Optional<Output<String>> sourceType() {
        return Optional.ofNullable(this.sourceType);
    }

    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private EventSubscriptionArgs() {}

    private EventSubscriptionArgs(EventSubscriptionArgs $) {
        this.enabled = $.enabled;
        this.eventCategories = $.eventCategories;
        this.name = $.name;
        this.namePrefix = $.namePrefix;
        this.snsTopicArn = $.snsTopicArn;
        this.sourceIds = $.sourceIds;
        this.sourceType = $.sourceType;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EventSubscriptionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EventSubscriptionArgs $;

        public Builder() {
            $ = new EventSubscriptionArgs();
        }

        public Builder(EventSubscriptionArgs defaults) {
            $ = new EventSubscriptionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enabled A boolean flag to enable/disable the subscription. Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled A boolean flag to enable/disable the subscription. Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param eventCategories A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Event.html or run `aws docdb describe-event-categories`.
         * 
         * @return builder
         * 
         */
        public Builder eventCategories(@Nullable Output<List<String>> eventCategories) {
            $.eventCategories = eventCategories;
            return this;
        }

        /**
         * @param eventCategories A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Event.html or run `aws docdb describe-event-categories`.
         * 
         * @return builder
         * 
         */
        public Builder eventCategories(List<String> eventCategories) {
            return eventCategories(Output.of(eventCategories));
        }

        /**
         * @param eventCategories A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Event.html or run `aws docdb describe-event-categories`.
         * 
         * @return builder
         * 
         */
        public Builder eventCategories(String... eventCategories) {
            return eventCategories(List.of(eventCategories));
        }

        /**
         * @param name The name of the DocumentDB event subscription. By default generated by this provider.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the DocumentDB event subscription. By default generated by this provider.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namePrefix The name of the DocumentDB event subscription. Conflicts with `name`.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(@Nullable Output<String> namePrefix) {
            $.namePrefix = namePrefix;
            return this;
        }

        /**
         * @param namePrefix The name of the DocumentDB event subscription. Conflicts with `name`.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(String namePrefix) {
            return namePrefix(Output.of(namePrefix));
        }

        public Builder snsTopicArn(Output<String> snsTopicArn) {
            $.snsTopicArn = snsTopicArn;
            return this;
        }

        public Builder snsTopicArn(String snsTopicArn) {
            return snsTopicArn(Output.of(snsTopicArn));
        }

        /**
         * @param sourceIds A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a source_type must also be specified.
         * 
         * @return builder
         * 
         */
        public Builder sourceIds(@Nullable Output<List<String>> sourceIds) {
            $.sourceIds = sourceIds;
            return this;
        }

        /**
         * @param sourceIds A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a source_type must also be specified.
         * 
         * @return builder
         * 
         */
        public Builder sourceIds(List<String> sourceIds) {
            return sourceIds(Output.of(sourceIds));
        }

        /**
         * @param sourceIds A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a source_type must also be specified.
         * 
         * @return builder
         * 
         */
        public Builder sourceIds(String... sourceIds) {
            return sourceIds(List.of(sourceIds));
        }

        /**
         * @param sourceType The type of source that will be generating the events. Valid options are `db-instance`, `db-cluster`, `db-parameter-group`, `db-security-group`,`  db-cluster-snapshot `. If not set, all sources will be subscribed to.
         * 
         * @return builder
         * 
         */
        public Builder sourceType(@Nullable Output<String> sourceType) {
            $.sourceType = sourceType;
            return this;
        }

        /**
         * @param sourceType The type of source that will be generating the events. Valid options are `db-instance`, `db-cluster`, `db-parameter-group`, `db-security-group`,`  db-cluster-snapshot `. If not set, all sources will be subscribed to.
         * 
         * @return builder
         * 
         */
        public Builder sourceType(String sourceType) {
            return sourceType(Output.of(sourceType));
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public EventSubscriptionArgs build() {
            $.snsTopicArn = Objects.requireNonNull($.snsTopicArn, "expected parameter 'snsTopicArn' to be non-null");
            return $;
        }
    }

}
