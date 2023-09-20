// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.gamelift.inputs;

import com.pulumi.aws.gamelift.inputs.GameSessionQueuePlayerLatencyPolicyArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GameSessionQueueState extends com.pulumi.resources.ResourceArgs {

    public static final GameSessionQueueState Empty = new GameSessionQueueState();

    /**
     * Game Session Queue ARN.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return Game Session Queue ARN.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Information to be added to all events that are related to this game session queue.
     * 
     */
    @Import(name="customEventData")
    private @Nullable Output<String> customEventData;

    /**
     * @return Information to be added to all events that are related to this game session queue.
     * 
     */
    public Optional<Output<String>> customEventData() {
        return Optional.ofNullable(this.customEventData);
    }

    /**
     * List of fleet/alias ARNs used by session queue for placing game sessions.
     * 
     */
    @Import(name="destinations")
    private @Nullable Output<List<String>> destinations;

    /**
     * @return List of fleet/alias ARNs used by session queue for placing game sessions.
     * 
     */
    public Optional<Output<List<String>>> destinations() {
        return Optional.ofNullable(this.destinations);
    }

    /**
     * Name of the session queue.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the session queue.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * An SNS topic ARN that is set up to receive game session placement notifications.
     * 
     */
    @Import(name="notificationTarget")
    private @Nullable Output<String> notificationTarget;

    /**
     * @return An SNS topic ARN that is set up to receive game session placement notifications.
     * 
     */
    public Optional<Output<String>> notificationTarget() {
        return Optional.ofNullable(this.notificationTarget);
    }

    /**
     * One or more policies used to choose fleet based on player latency. See below.
     * 
     */
    @Import(name="playerLatencyPolicies")
    private @Nullable Output<List<GameSessionQueuePlayerLatencyPolicyArgs>> playerLatencyPolicies;

    /**
     * @return One or more policies used to choose fleet based on player latency. See below.
     * 
     */
    public Optional<Output<List<GameSessionQueuePlayerLatencyPolicyArgs>>> playerLatencyPolicies() {
        return Optional.ofNullable(this.playerLatencyPolicies);
    }

    /**
     * Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

    /**
     * Maximum time a game session request can remain in the queue.
     * 
     */
    @Import(name="timeoutInSeconds")
    private @Nullable Output<Integer> timeoutInSeconds;

    /**
     * @return Maximum time a game session request can remain in the queue.
     * 
     */
    public Optional<Output<Integer>> timeoutInSeconds() {
        return Optional.ofNullable(this.timeoutInSeconds);
    }

    private GameSessionQueueState() {}

    private GameSessionQueueState(GameSessionQueueState $) {
        this.arn = $.arn;
        this.customEventData = $.customEventData;
        this.destinations = $.destinations;
        this.name = $.name;
        this.notificationTarget = $.notificationTarget;
        this.playerLatencyPolicies = $.playerLatencyPolicies;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.timeoutInSeconds = $.timeoutInSeconds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GameSessionQueueState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GameSessionQueueState $;

        public Builder() {
            $ = new GameSessionQueueState();
        }

        public Builder(GameSessionQueueState defaults) {
            $ = new GameSessionQueueState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn Game Session Queue ARN.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn Game Session Queue ARN.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param customEventData Information to be added to all events that are related to this game session queue.
         * 
         * @return builder
         * 
         */
        public Builder customEventData(@Nullable Output<String> customEventData) {
            $.customEventData = customEventData;
            return this;
        }

        /**
         * @param customEventData Information to be added to all events that are related to this game session queue.
         * 
         * @return builder
         * 
         */
        public Builder customEventData(String customEventData) {
            return customEventData(Output.of(customEventData));
        }

        /**
         * @param destinations List of fleet/alias ARNs used by session queue for placing game sessions.
         * 
         * @return builder
         * 
         */
        public Builder destinations(@Nullable Output<List<String>> destinations) {
            $.destinations = destinations;
            return this;
        }

        /**
         * @param destinations List of fleet/alias ARNs used by session queue for placing game sessions.
         * 
         * @return builder
         * 
         */
        public Builder destinations(List<String> destinations) {
            return destinations(Output.of(destinations));
        }

        /**
         * @param destinations List of fleet/alias ARNs used by session queue for placing game sessions.
         * 
         * @return builder
         * 
         */
        public Builder destinations(String... destinations) {
            return destinations(List.of(destinations));
        }

        /**
         * @param name Name of the session queue.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the session queue.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param notificationTarget An SNS topic ARN that is set up to receive game session placement notifications.
         * 
         * @return builder
         * 
         */
        public Builder notificationTarget(@Nullable Output<String> notificationTarget) {
            $.notificationTarget = notificationTarget;
            return this;
        }

        /**
         * @param notificationTarget An SNS topic ARN that is set up to receive game session placement notifications.
         * 
         * @return builder
         * 
         */
        public Builder notificationTarget(String notificationTarget) {
            return notificationTarget(Output.of(notificationTarget));
        }

        /**
         * @param playerLatencyPolicies One or more policies used to choose fleet based on player latency. See below.
         * 
         * @return builder
         * 
         */
        public Builder playerLatencyPolicies(@Nullable Output<List<GameSessionQueuePlayerLatencyPolicyArgs>> playerLatencyPolicies) {
            $.playerLatencyPolicies = playerLatencyPolicies;
            return this;
        }

        /**
         * @param playerLatencyPolicies One or more policies used to choose fleet based on player latency. See below.
         * 
         * @return builder
         * 
         */
        public Builder playerLatencyPolicies(List<GameSessionQueuePlayerLatencyPolicyArgs> playerLatencyPolicies) {
            return playerLatencyPolicies(Output.of(playerLatencyPolicies));
        }

        /**
         * @param playerLatencyPolicies One or more policies used to choose fleet based on player latency. See below.
         * 
         * @return builder
         * 
         */
        public Builder playerLatencyPolicies(GameSessionQueuePlayerLatencyPolicyArgs... playerLatencyPolicies) {
            return playerLatencyPolicies(List.of(playerLatencyPolicies));
        }

        /**
         * @param tags Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        /**
         * @param timeoutInSeconds Maximum time a game session request can remain in the queue.
         * 
         * @return builder
         * 
         */
        public Builder timeoutInSeconds(@Nullable Output<Integer> timeoutInSeconds) {
            $.timeoutInSeconds = timeoutInSeconds;
            return this;
        }

        /**
         * @param timeoutInSeconds Maximum time a game session request can remain in the queue.
         * 
         * @return builder
         * 
         */
        public Builder timeoutInSeconds(Integer timeoutInSeconds) {
            return timeoutInSeconds(Output.of(timeoutInSeconds));
        }

        public GameSessionQueueState build() {
            return $;
        }
    }

}
