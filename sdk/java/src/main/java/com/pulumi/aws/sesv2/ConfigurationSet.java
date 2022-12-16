// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sesv2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.sesv2.ConfigurationSetArgs;
import com.pulumi.aws.sesv2.inputs.ConfigurationSetState;
import com.pulumi.aws.sesv2.outputs.ConfigurationSetDeliveryOptions;
import com.pulumi.aws.sesv2.outputs.ConfigurationSetReputationOptions;
import com.pulumi.aws.sesv2.outputs.ConfigurationSetSendingOptions;
import com.pulumi.aws.sesv2.outputs.ConfigurationSetSuppressionOptions;
import com.pulumi.aws.sesv2.outputs.ConfigurationSetTrackingOptions;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Resource for managing an AWS SESv2 (Simple Email V2) Configuration Set.
 * 
 * ## Example Usage
 * ### Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.sesv2.ConfigurationSet;
 * import com.pulumi.aws.sesv2.ConfigurationSetArgs;
 * import com.pulumi.aws.sesv2.inputs.ConfigurationSetDeliveryOptionsArgs;
 * import com.pulumi.aws.sesv2.inputs.ConfigurationSetReputationOptionsArgs;
 * import com.pulumi.aws.sesv2.inputs.ConfigurationSetSendingOptionsArgs;
 * import com.pulumi.aws.sesv2.inputs.ConfigurationSetSuppressionOptionsArgs;
 * import com.pulumi.aws.sesv2.inputs.ConfigurationSetTrackingOptionsArgs;
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
 *         var example = new ConfigurationSet(&#34;example&#34;, ConfigurationSetArgs.builder()        
 *             .configurationSetName(&#34;example&#34;)
 *             .deliveryOptions(ConfigurationSetDeliveryOptionsArgs.builder()
 *                 .tlsPolicy(&#34;REQUIRE&#34;)
 *                 .build())
 *             .reputationOptions(ConfigurationSetReputationOptionsArgs.builder()
 *                 .reputationMetricsEnabled(false)
 *                 .build())
 *             .sendingOptions(ConfigurationSetSendingOptionsArgs.builder()
 *                 .sendingEnabled(true)
 *                 .build())
 *             .suppressionOptions(ConfigurationSetSuppressionOptionsArgs.builder()
 *                 .suppressedReasons(                
 *                     &#34;BOUNCE&#34;,
 *                     &#34;COMPLAINT&#34;)
 *                 .build())
 *             .trackingOptions(ConfigurationSetTrackingOptionsArgs.builder()
 *                 .customRedirectDomain(&#34;example.com&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * SESv2 (Simple Email V2) Configuration Set can be imported using the `configuration_set_name`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:sesv2/configurationSet:ConfigurationSet example example
 * ```
 * 
 */
@ResourceType(type="aws:sesv2/configurationSet:ConfigurationSet")
public class ConfigurationSet extends com.pulumi.resources.CustomResource {
    /**
     * ARN of the Configuration Set.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return ARN of the Configuration Set.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The name of the configuration set.
     * 
     */
    @Export(name="configurationSetName", refs={String.class}, tree="[0]")
    private Output<String> configurationSetName;

    /**
     * @return The name of the configuration set.
     * 
     */
    public Output<String> configurationSetName() {
        return this.configurationSetName;
    }
    /**
     * An object that defines the dedicated IP pool that is used to send emails that you send using the configuration set.
     * 
     */
    @Export(name="deliveryOptions", refs={ConfigurationSetDeliveryOptions.class}, tree="[0]")
    private Output</* @Nullable */ ConfigurationSetDeliveryOptions> deliveryOptions;

    /**
     * @return An object that defines the dedicated IP pool that is used to send emails that you send using the configuration set.
     * 
     */
    public Output<Optional<ConfigurationSetDeliveryOptions>> deliveryOptions() {
        return Codegen.optional(this.deliveryOptions);
    }
    /**
     * An object that defines whether or not Amazon SES collects reputation metrics for the emails that you send that use the configuration set.
     * 
     */
    @Export(name="reputationOptions", refs={ConfigurationSetReputationOptions.class}, tree="[0]")
    private Output<ConfigurationSetReputationOptions> reputationOptions;

    /**
     * @return An object that defines whether or not Amazon SES collects reputation metrics for the emails that you send that use the configuration set.
     * 
     */
    public Output<ConfigurationSetReputationOptions> reputationOptions() {
        return this.reputationOptions;
    }
    /**
     * An object that defines whether or not Amazon SES can send email that you send using the configuration set.
     * 
     */
    @Export(name="sendingOptions", refs={ConfigurationSetSendingOptions.class}, tree="[0]")
    private Output<ConfigurationSetSendingOptions> sendingOptions;

    /**
     * @return An object that defines whether or not Amazon SES can send email that you send using the configuration set.
     * 
     */
    public Output<ConfigurationSetSendingOptions> sendingOptions() {
        return this.sendingOptions;
    }
    /**
     * An object that contains information about the suppression list preferences for your account.
     * 
     */
    @Export(name="suppressionOptions", refs={ConfigurationSetSuppressionOptions.class}, tree="[0]")
    private Output</* @Nullable */ ConfigurationSetSuppressionOptions> suppressionOptions;

    /**
     * @return An object that contains information about the suppression list preferences for your account.
     * 
     */
    public Output<Optional<ConfigurationSetSuppressionOptions>> suppressionOptions() {
        return Codegen.optional(this.suppressionOptions);
    }
    /**
     * A map of tags to assign to the service. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the service. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    /**
     * An object that defines the open and click tracking options for emails that you send using the configuration set.
     * 
     */
    @Export(name="trackingOptions", refs={ConfigurationSetTrackingOptions.class}, tree="[0]")
    private Output</* @Nullable */ ConfigurationSetTrackingOptions> trackingOptions;

    /**
     * @return An object that defines the open and click tracking options for emails that you send using the configuration set.
     * 
     */
    public Output<Optional<ConfigurationSetTrackingOptions>> trackingOptions() {
        return Codegen.optional(this.trackingOptions);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ConfigurationSet(String name) {
        this(name, ConfigurationSetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ConfigurationSet(String name, ConfigurationSetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ConfigurationSet(String name, ConfigurationSetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:sesv2/configurationSet:ConfigurationSet", name, args == null ? ConfigurationSetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ConfigurationSet(String name, Output<String> id, @Nullable ConfigurationSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:sesv2/configurationSet:ConfigurationSet", name, state, makeResourceOptions(options, id));
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
    public static ConfigurationSet get(String name, Output<String> id, @Nullable ConfigurationSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ConfigurationSet(name, id, state, options);
    }
}
