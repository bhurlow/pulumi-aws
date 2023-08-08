// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.gamelift;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.gamelift.GameSessionQueueArgs;
import com.pulumi.aws.gamelift.inputs.GameSessionQueueState;
import com.pulumi.aws.gamelift.outputs.GameSessionQueuePlayerLatencyPolicy;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an GameLift Game Session Queue resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.gamelift.GameSessionQueue;
 * import com.pulumi.aws.gamelift.GameSessionQueueArgs;
 * import com.pulumi.aws.gamelift.inputs.GameSessionQueuePlayerLatencyPolicyArgs;
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
 *         var test = new GameSessionQueue(&#34;test&#34;, GameSessionQueueArgs.builder()        
 *             .destinations(            
 *                 aws_gamelift_fleet.us_west_2_fleet().arn(),
 *                 aws_gamelift_fleet.eu_central_1_fleet().arn())
 *             .notificationTarget(aws_sns_topic.game_session_queue_notifications().arn())
 *             .playerLatencyPolicies(            
 *                 GameSessionQueuePlayerLatencyPolicyArgs.builder()
 *                     .maximumIndividualPlayerLatencyMilliseconds(100)
 *                     .policyDurationSeconds(5)
 *                     .build(),
 *                 GameSessionQueuePlayerLatencyPolicyArgs.builder()
 *                     .maximumIndividualPlayerLatencyMilliseconds(200)
 *                     .build())
 *             .timeoutInSeconds(60)
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
 *  to = aws_gamelift_game_session_queue.example
 * 
 *  id = &#34;example&#34; } Using `pulumi import`, import GameLift Game Session Queues using their `name`. For exampleconsole % pulumi import aws_gamelift_game_session_queue.example example
 * 
 */
@ResourceType(type="aws:gamelift/gameSessionQueue:GameSessionQueue")
public class GameSessionQueue extends com.pulumi.resources.CustomResource {
    /**
     * Game Session Queue ARN.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Game Session Queue ARN.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Information to be added to all events that are related to this game session queue.
     * 
     */
    @Export(name="customEventData", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> customEventData;

    /**
     * @return Information to be added to all events that are related to this game session queue.
     * 
     */
    public Output<Optional<String>> customEventData() {
        return Codegen.optional(this.customEventData);
    }
    /**
     * List of fleet/alias ARNs used by session queue for placing game sessions.
     * 
     */
    @Export(name="destinations", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> destinations;

    /**
     * @return List of fleet/alias ARNs used by session queue for placing game sessions.
     * 
     */
    public Output<Optional<List<String>>> destinations() {
        return Codegen.optional(this.destinations);
    }
    /**
     * Name of the session queue.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the session queue.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * An SNS topic ARN that is set up to receive game session placement notifications.
     * 
     */
    @Export(name="notificationTarget", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> notificationTarget;

    /**
     * @return An SNS topic ARN that is set up to receive game session placement notifications.
     * 
     */
    public Output<Optional<String>> notificationTarget() {
        return Codegen.optional(this.notificationTarget);
    }
    /**
     * One or more policies used to choose fleet based on player latency. See below.
     * 
     */
    @Export(name="playerLatencyPolicies", refs={List.class,GameSessionQueuePlayerLatencyPolicy.class}, tree="[0,1]")
    private Output</* @Nullable */ List<GameSessionQueuePlayerLatencyPolicy>> playerLatencyPolicies;

    /**
     * @return One or more policies used to choose fleet based on player latency. See below.
     * 
     */
    public Output<Optional<List<GameSessionQueuePlayerLatencyPolicy>>> playerLatencyPolicies() {
        return Codegen.optional(this.playerLatencyPolicies);
    }
    /**
     * Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
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
     * Maximum time a game session request can remain in the queue.
     * 
     */
    @Export(name="timeoutInSeconds", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> timeoutInSeconds;

    /**
     * @return Maximum time a game session request can remain in the queue.
     * 
     */
    public Output<Optional<Integer>> timeoutInSeconds() {
        return Codegen.optional(this.timeoutInSeconds);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GameSessionQueue(String name) {
        this(name, GameSessionQueueArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GameSessionQueue(String name, @Nullable GameSessionQueueArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GameSessionQueue(String name, @Nullable GameSessionQueueArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:gamelift/gameSessionQueue:GameSessionQueue", name, args == null ? GameSessionQueueArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GameSessionQueue(String name, Output<String> id, @Nullable GameSessionQueueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:gamelift/gameSessionQueue:GameSessionQueue", name, state, makeResourceOptions(options, id));
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
    public static GameSessionQueue get(String name, Output<String> id, @Nullable GameSessionQueueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GameSessionQueue(name, id, state, options);
    }
}
