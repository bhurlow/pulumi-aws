// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cloudwatch;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.cloudwatch.EventBusArgs;
import com.pulumi.aws.cloudwatch.inputs.EventBusState;
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
 * Provides an EventBridge event bus resource.
 * 
 * &gt; **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.cloudwatch.EventBus;
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
 *         var messenger = new EventBus(&#34;messenger&#34;);
 * 
 *     }
 * }
 * ```
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.cloudwatch.CloudwatchFunctions;
 * import com.pulumi.aws.cloudwatch.inputs.GetEventSourceArgs;
 * import com.pulumi.aws.cloudwatch.EventBus;
 * import com.pulumi.aws.cloudwatch.EventBusArgs;
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
 *         final var examplepartnerEventSource = CloudwatchFunctions.getEventSource(GetEventSourceArgs.builder()
 *             .namePrefix(&#34;aws.partner/examplepartner.com&#34;)
 *             .build());
 * 
 *         var examplepartnerEventBus = new EventBus(&#34;examplepartnerEventBus&#34;, EventBusArgs.builder()        
 *             .eventSourceName(examplepartnerEventSource.applyValue(getEventSourceResult -&gt; getEventSourceResult.name()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import EventBridge event buses using the `name` (which can also be a partner event source name). For example:
 * 
 * ```sh
 *  $ pulumi import aws:cloudwatch/eventBus:EventBus messenger chat-messages
 * ```
 * 
 */
@ResourceType(type="aws:cloudwatch/eventBus:EventBus")
public class EventBus extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name (ARN) of the event bus.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) of the event bus.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The partner event source that the new event bus will be matched with. Must match `name`.
     * 
     */
    @Export(name="eventSourceName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> eventSourceName;

    /**
     * @return The partner event source that the new event bus will be matched with. Must match `name`.
     * 
     */
    public Output<Optional<String>> eventSourceName() {
        return Codegen.optional(this.eventSourceName);
    }
    /**
     * The name of the new event bus. The names of custom event buses can&#39;t contain the / character. To create a partner event bus, ensure the `name` matches the `event_source_name`.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the new event bus. The names of custom event buses can&#39;t contain the / character. To create a partner event bus, ensure the `name` matches the `event_source_name`.
     * 
     */
    public Output<String> name() {
        return this.name;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EventBus(String name) {
        this(name, EventBusArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EventBus(String name, @Nullable EventBusArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EventBus(String name, @Nullable EventBusArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cloudwatch/eventBus:EventBus", name, args == null ? EventBusArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EventBus(String name, Output<String> id, @Nullable EventBusState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cloudwatch/eventBus:EventBus", name, state, makeResourceOptions(options, id));
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
    public static EventBus get(String name, Output<String> id, @Nullable EventBusState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EventBus(name, id, state, options);
    }
}
