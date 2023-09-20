// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ssm.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MaintenanceWindowState extends com.pulumi.resources.ResourceArgs {

    public static final MaintenanceWindowState Empty = new MaintenanceWindowState();

    /**
     * Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
     * 
     */
    @Import(name="allowUnassociatedTargets")
    private @Nullable Output<Boolean> allowUnassociatedTargets;

    /**
     * @return Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
     * 
     */
    public Optional<Output<Boolean>> allowUnassociatedTargets() {
        return Optional.ofNullable(this.allowUnassociatedTargets);
    }

    /**
     * The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
     * 
     */
    @Import(name="cutoff")
    private @Nullable Output<Integer> cutoff;

    /**
     * @return The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
     * 
     */
    public Optional<Output<Integer>> cutoff() {
        return Optional.ofNullable(this.cutoff);
    }

    /**
     * A description for the maintenance window.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description for the maintenance window.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The duration of the Maintenance Window in hours.
     * 
     */
    @Import(name="duration")
    private @Nullable Output<Integer> duration;

    /**
     * @return The duration of the Maintenance Window in hours.
     * 
     */
    public Optional<Output<Integer>> duration() {
        return Optional.ofNullable(this.duration);
    }

    /**
     * Whether the maintenance window is enabled. Default: `true`.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Whether the maintenance window is enabled. Default: `true`.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
     * 
     */
    @Import(name="endDate")
    private @Nullable Output<String> endDate;

    /**
     * @return Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
     * 
     */
    public Optional<Output<String>> endDate() {
        return Optional.ofNullable(this.endDate);
    }

    /**
     * The name of the maintenance window.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the maintenance window.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The schedule of the Maintenance Window in the form of a [cron or rate expression](https://docs.aws.amazon.com/systems-manager/latest/userguide/reference-cron-and-rate-expressions.html).
     * 
     */
    @Import(name="schedule")
    private @Nullable Output<String> schedule;

    /**
     * @return The schedule of the Maintenance Window in the form of a [cron or rate expression](https://docs.aws.amazon.com/systems-manager/latest/userguide/reference-cron-and-rate-expressions.html).
     * 
     */
    public Optional<Output<String>> schedule() {
        return Optional.ofNullable(this.schedule);
    }

    /**
     * The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
     * 
     */
    @Import(name="scheduleOffset")
    private @Nullable Output<Integer> scheduleOffset;

    /**
     * @return The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
     * 
     */
    public Optional<Output<Integer>> scheduleOffset() {
        return Optional.ofNullable(this.scheduleOffset);
    }

    /**
     * Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
     * 
     */
    @Import(name="scheduleTimezone")
    private @Nullable Output<String> scheduleTimezone;

    /**
     * @return Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
     * 
     */
    public Optional<Output<String>> scheduleTimezone() {
        return Optional.ofNullable(this.scheduleTimezone);
    }

    /**
     * Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
     * 
     */
    @Import(name="startDate")
    private @Nullable Output<String> startDate;

    /**
     * @return Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
     * 
     */
    public Optional<Output<String>> startDate() {
        return Optional.ofNullable(this.startDate);
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

    private MaintenanceWindowState() {}

    private MaintenanceWindowState(MaintenanceWindowState $) {
        this.allowUnassociatedTargets = $.allowUnassociatedTargets;
        this.cutoff = $.cutoff;
        this.description = $.description;
        this.duration = $.duration;
        this.enabled = $.enabled;
        this.endDate = $.endDate;
        this.name = $.name;
        this.schedule = $.schedule;
        this.scheduleOffset = $.scheduleOffset;
        this.scheduleTimezone = $.scheduleTimezone;
        this.startDate = $.startDate;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MaintenanceWindowState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MaintenanceWindowState $;

        public Builder() {
            $ = new MaintenanceWindowState();
        }

        public Builder(MaintenanceWindowState defaults) {
            $ = new MaintenanceWindowState(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowUnassociatedTargets Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
         * 
         * @return builder
         * 
         */
        public Builder allowUnassociatedTargets(@Nullable Output<Boolean> allowUnassociatedTargets) {
            $.allowUnassociatedTargets = allowUnassociatedTargets;
            return this;
        }

        /**
         * @param allowUnassociatedTargets Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
         * 
         * @return builder
         * 
         */
        public Builder allowUnassociatedTargets(Boolean allowUnassociatedTargets) {
            return allowUnassociatedTargets(Output.of(allowUnassociatedTargets));
        }

        /**
         * @param cutoff The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
         * 
         * @return builder
         * 
         */
        public Builder cutoff(@Nullable Output<Integer> cutoff) {
            $.cutoff = cutoff;
            return this;
        }

        /**
         * @param cutoff The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
         * 
         * @return builder
         * 
         */
        public Builder cutoff(Integer cutoff) {
            return cutoff(Output.of(cutoff));
        }

        /**
         * @param description A description for the maintenance window.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description for the maintenance window.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param duration The duration of the Maintenance Window in hours.
         * 
         * @return builder
         * 
         */
        public Builder duration(@Nullable Output<Integer> duration) {
            $.duration = duration;
            return this;
        }

        /**
         * @param duration The duration of the Maintenance Window in hours.
         * 
         * @return builder
         * 
         */
        public Builder duration(Integer duration) {
            return duration(Output.of(duration));
        }

        /**
         * @param enabled Whether the maintenance window is enabled. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Whether the maintenance window is enabled. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param endDate Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
         * 
         * @return builder
         * 
         */
        public Builder endDate(@Nullable Output<String> endDate) {
            $.endDate = endDate;
            return this;
        }

        /**
         * @param endDate Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
         * 
         * @return builder
         * 
         */
        public Builder endDate(String endDate) {
            return endDate(Output.of(endDate));
        }

        /**
         * @param name The name of the maintenance window.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the maintenance window.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param schedule The schedule of the Maintenance Window in the form of a [cron or rate expression](https://docs.aws.amazon.com/systems-manager/latest/userguide/reference-cron-and-rate-expressions.html).
         * 
         * @return builder
         * 
         */
        public Builder schedule(@Nullable Output<String> schedule) {
            $.schedule = schedule;
            return this;
        }

        /**
         * @param schedule The schedule of the Maintenance Window in the form of a [cron or rate expression](https://docs.aws.amazon.com/systems-manager/latest/userguide/reference-cron-and-rate-expressions.html).
         * 
         * @return builder
         * 
         */
        public Builder schedule(String schedule) {
            return schedule(Output.of(schedule));
        }

        /**
         * @param scheduleOffset The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
         * 
         * @return builder
         * 
         */
        public Builder scheduleOffset(@Nullable Output<Integer> scheduleOffset) {
            $.scheduleOffset = scheduleOffset;
            return this;
        }

        /**
         * @param scheduleOffset The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
         * 
         * @return builder
         * 
         */
        public Builder scheduleOffset(Integer scheduleOffset) {
            return scheduleOffset(Output.of(scheduleOffset));
        }

        /**
         * @param scheduleTimezone Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
         * 
         * @return builder
         * 
         */
        public Builder scheduleTimezone(@Nullable Output<String> scheduleTimezone) {
            $.scheduleTimezone = scheduleTimezone;
            return this;
        }

        /**
         * @param scheduleTimezone Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
         * 
         * @return builder
         * 
         */
        public Builder scheduleTimezone(String scheduleTimezone) {
            return scheduleTimezone(Output.of(scheduleTimezone));
        }

        /**
         * @param startDate Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
         * 
         * @return builder
         * 
         */
        public Builder startDate(@Nullable Output<String> startDate) {
            $.startDate = startDate;
            return this;
        }

        /**
         * @param startDate Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
         * 
         * @return builder
         * 
         */
        public Builder startDate(String startDate) {
            return startDate(Output.of(startDate));
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

        public MaintenanceWindowState build() {
            return $;
        }
    }

}
