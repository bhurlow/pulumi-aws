// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appconfig.inputs;

import com.pulumi.aws.appconfig.inputs.EventIntegrationEventFilterArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EventIntegrationState extends com.pulumi.resources.ResourceArgs {

    public static final EventIntegrationState Empty = new EventIntegrationState();

    /**
     * ARN of the Event Integration.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return ARN of the Event Integration.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Description of the Event Integration.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the Event Integration.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Block that defines the configuration information for the event filter. The Event Filter block is documented below.
     * 
     */
    @Import(name="eventFilter")
    private @Nullable Output<EventIntegrationEventFilterArgs> eventFilter;

    /**
     * @return Block that defines the configuration information for the event filter. The Event Filter block is documented below.
     * 
     */
    public Optional<Output<EventIntegrationEventFilterArgs>> eventFilter() {
        return Optional.ofNullable(this.eventFilter);
    }

    /**
     * EventBridge bus.
     * 
     */
    @Import(name="eventbridgeBus")
    private @Nullable Output<String> eventbridgeBus;

    /**
     * @return EventBridge bus.
     * 
     */
    public Optional<Output<String>> eventbridgeBus() {
        return Optional.ofNullable(this.eventbridgeBus);
    }

    /**
     * Name of the Event Integration.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the Event Integration.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Tags to apply to the Event Integration. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Tags to apply to the Event Integration. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

    private EventIntegrationState() {}

    private EventIntegrationState(EventIntegrationState $) {
        this.arn = $.arn;
        this.description = $.description;
        this.eventFilter = $.eventFilter;
        this.eventbridgeBus = $.eventbridgeBus;
        this.name = $.name;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EventIntegrationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EventIntegrationState $;

        public Builder() {
            $ = new EventIntegrationState();
        }

        public Builder(EventIntegrationState defaults) {
            $ = new EventIntegrationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn ARN of the Event Integration.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn ARN of the Event Integration.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param description Description of the Event Integration.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the Event Integration.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param eventFilter Block that defines the configuration information for the event filter. The Event Filter block is documented below.
         * 
         * @return builder
         * 
         */
        public Builder eventFilter(@Nullable Output<EventIntegrationEventFilterArgs> eventFilter) {
            $.eventFilter = eventFilter;
            return this;
        }

        /**
         * @param eventFilter Block that defines the configuration information for the event filter. The Event Filter block is documented below.
         * 
         * @return builder
         * 
         */
        public Builder eventFilter(EventIntegrationEventFilterArgs eventFilter) {
            return eventFilter(Output.of(eventFilter));
        }

        /**
         * @param eventbridgeBus EventBridge bus.
         * 
         * @return builder
         * 
         */
        public Builder eventbridgeBus(@Nullable Output<String> eventbridgeBus) {
            $.eventbridgeBus = eventbridgeBus;
            return this;
        }

        /**
         * @param eventbridgeBus EventBridge bus.
         * 
         * @return builder
         * 
         */
        public Builder eventbridgeBus(String eventbridgeBus) {
            return eventbridgeBus(Output.of(eventbridgeBus));
        }

        /**
         * @param name Name of the Event Integration.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the Event Integration.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param tags Tags to apply to the Event Integration. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Tags to apply to the Event Integration. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public EventIntegrationState build() {
            return $;
        }
    }

}
