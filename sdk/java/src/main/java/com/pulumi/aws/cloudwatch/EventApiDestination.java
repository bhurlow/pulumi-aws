// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cloudwatch;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.cloudwatch.EventApiDestinationArgs;
import com.pulumi.aws.cloudwatch.inputs.EventApiDestinationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an EventBridge event API Destination resource.
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
 * import com.pulumi.aws.cloudwatch.EventApiDestination;
 * import com.pulumi.aws.cloudwatch.EventApiDestinationArgs;
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
 *         var test = new EventApiDestination(&#34;test&#34;, EventApiDestinationArgs.builder()        
 *             .description(&#34;An API Destination&#34;)
 *             .invocationEndpoint(&#34;https://api.destination.com/endpoint&#34;)
 *             .httpMethod(&#34;POST&#34;)
 *             .invocationRateLimitPerSecond(20)
 *             .connectionArn(aws_cloudwatch_event_connection.test().arn())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * EventBridge API Destinations can be imported using the `name`, e.g., console
 * 
 * ```sh
 *  $ pulumi import aws:cloudwatch/eventApiDestination:EventApiDestination test api-destination
 * ```
 * 
 */
@ResourceType(type="aws:cloudwatch/eventApiDestination:EventApiDestination")
public class EventApiDestination extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name (ARN) of the event API Destination.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) of the event API Destination.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * ARN of the EventBridge Connection to use for the API Destination.
     * 
     */
    @Export(name="connectionArn", refs={String.class}, tree="[0]")
    private Output<String> connectionArn;

    /**
     * @return ARN of the EventBridge Connection to use for the API Destination.
     * 
     */
    public Output<String> connectionArn() {
        return this.connectionArn;
    }
    /**
     * The description of the new API Destination. Maximum of 512 characters.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the new API Destination. Maximum of 512 characters.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Select the HTTP method used for the invocation endpoint, such as GET, POST, PUT, etc.
     * 
     */
    @Export(name="httpMethod", refs={String.class}, tree="[0]")
    private Output<String> httpMethod;

    /**
     * @return Select the HTTP method used for the invocation endpoint, such as GET, POST, PUT, etc.
     * 
     */
    public Output<String> httpMethod() {
        return this.httpMethod;
    }
    /**
     * URL endpoint to invoke as a target. This could be a valid endpoint generated by a partner service. You can include &#34;*&#34; as path parameters wildcards to be set from the Target HttpParameters.
     * 
     */
    @Export(name="invocationEndpoint", refs={String.class}, tree="[0]")
    private Output<String> invocationEndpoint;

    /**
     * @return URL endpoint to invoke as a target. This could be a valid endpoint generated by a partner service. You can include &#34;*&#34; as path parameters wildcards to be set from the Target HttpParameters.
     * 
     */
    public Output<String> invocationEndpoint() {
        return this.invocationEndpoint;
    }
    /**
     * Enter the maximum number of invocations per second to allow for this destination. Enter a value greater than 0 (default 300).
     * 
     */
    @Export(name="invocationRateLimitPerSecond", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> invocationRateLimitPerSecond;

    /**
     * @return Enter the maximum number of invocations per second to allow for this destination. Enter a value greater than 0 (default 300).
     * 
     */
    public Output<Optional<Integer>> invocationRateLimitPerSecond() {
        return Codegen.optional(this.invocationRateLimitPerSecond);
    }
    /**
     * The name of the new API Destination. The name must be unique for your account. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the new API Destination. The name must be unique for your account. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EventApiDestination(String name) {
        this(name, EventApiDestinationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EventApiDestination(String name, EventApiDestinationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EventApiDestination(String name, EventApiDestinationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cloudwatch/eventApiDestination:EventApiDestination", name, args == null ? EventApiDestinationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EventApiDestination(String name, Output<String> id, @Nullable EventApiDestinationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cloudwatch/eventApiDestination:EventApiDestination", name, state, makeResourceOptions(options, id));
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
    public static EventApiDestination get(String name, Output<String> id, @Nullable EventApiDestinationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EventApiDestination(name, id, state, options);
    }
}
