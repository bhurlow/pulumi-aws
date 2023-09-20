// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cloudwatch.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EventBusState extends com.pulumi.resources.ResourceArgs {

    public static final EventBusState Empty = new EventBusState();

    /**
     * The Amazon Resource Name (ARN) of the event bus.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) of the event bus.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The partner event source that the new event bus will be matched with. Must match `name`.
     * 
     */
    @Import(name="eventSourceName")
    private @Nullable Output<String> eventSourceName;

    /**
     * @return The partner event source that the new event bus will be matched with. Must match `name`.
     * 
     */
    public Optional<Output<String>> eventSourceName() {
        return Optional.ofNullable(this.eventSourceName);
    }

    /**
     * The name of the new event bus. The names of custom event buses can&#39;t contain the / character. To create a partner event bus, ensure the `name` matches the `event_source_name`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the new event bus. The names of custom event buses can&#39;t contain the / character. To create a partner event bus, ensure the `name` matches the `event_source_name`.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
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

    private EventBusState() {}

    private EventBusState(EventBusState $) {
        this.arn = $.arn;
        this.eventSourceName = $.eventSourceName;
        this.name = $.name;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EventBusState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EventBusState $;

        public Builder() {
            $ = new EventBusState();
        }

        public Builder(EventBusState defaults) {
            $ = new EventBusState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The Amazon Resource Name (ARN) of the event bus.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The Amazon Resource Name (ARN) of the event bus.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param eventSourceName The partner event source that the new event bus will be matched with. Must match `name`.
         * 
         * @return builder
         * 
         */
        public Builder eventSourceName(@Nullable Output<String> eventSourceName) {
            $.eventSourceName = eventSourceName;
            return this;
        }

        /**
         * @param eventSourceName The partner event source that the new event bus will be matched with. Must match `name`.
         * 
         * @return builder
         * 
         */
        public Builder eventSourceName(String eventSourceName) {
            return eventSourceName(Output.of(eventSourceName));
        }

        /**
         * @param name The name of the new event bus. The names of custom event buses can&#39;t contain the / character. To create a partner event bus, ensure the `name` matches the `event_source_name`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the new event bus. The names of custom event buses can&#39;t contain the / character. To create a partner event bus, ensure the `name` matches the `event_source_name`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
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

        public EventBusState build() {
            return $;
        }
    }

}
