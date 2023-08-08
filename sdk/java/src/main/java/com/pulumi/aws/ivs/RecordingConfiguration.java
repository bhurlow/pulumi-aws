// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ivs;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ivs.RecordingConfigurationArgs;
import com.pulumi.aws.ivs.inputs.RecordingConfigurationState;
import com.pulumi.aws.ivs.outputs.RecordingConfigurationDestinationConfiguration;
import com.pulumi.aws.ivs.outputs.RecordingConfigurationThumbnailConfiguration;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Resource for managing an AWS IVS (Interactive Video) Recording Configuration.
 * 
 * ## Example Usage
 * ### Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ivs.RecordingConfiguration;
 * import com.pulumi.aws.ivs.RecordingConfigurationArgs;
 * import com.pulumi.aws.ivs.inputs.RecordingConfigurationDestinationConfigurationArgs;
 * import com.pulumi.aws.ivs.inputs.RecordingConfigurationDestinationConfigurationS3Args;
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
 *         var example = new RecordingConfiguration(&#34;example&#34;, RecordingConfigurationArgs.builder()        
 *             .destinationConfiguration(RecordingConfigurationDestinationConfigurationArgs.builder()
 *                 .s3(RecordingConfigurationDestinationConfigurationS3Args.builder()
 *                     .bucketName(&#34;ivs-stream-archive&#34;)
 *                     .build())
 *                 .build())
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
 *  to = aws_ivs_recording_configuration.example
 * 
 *  id = &#34;arn:aws:ivs:us-west-2:326937407773:recording-configuration/KAk1sHBl2L47&#34; } Using `pulumi import`, import IVS (Interactive Video) Recording Configuration using the ARN. For exampleconsole % pulumi import aws_ivs_recording_configuration.example arn:aws:ivs:us-west-2:326937407773:recording-configuration/KAk1sHBl2L47
 * 
 */
@ResourceType(type="aws:ivs/recordingConfiguration:RecordingConfiguration")
public class RecordingConfiguration extends com.pulumi.resources.CustomResource {
    /**
     * ARN of the Recording Configuration.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return ARN of the Recording Configuration.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Object containing destination configuration for where recorded video will be stored.
     * 
     */
    @Export(name="destinationConfiguration", refs={RecordingConfigurationDestinationConfiguration.class}, tree="[0]")
    private Output<RecordingConfigurationDestinationConfiguration> destinationConfiguration;

    /**
     * @return Object containing destination configuration for where recorded video will be stored.
     * 
     */
    public Output<RecordingConfigurationDestinationConfiguration> destinationConfiguration() {
        return this.destinationConfiguration;
    }
    /**
     * Recording Configuration name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Recording Configuration name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.
     * 
     */
    @Export(name="recordingReconnectWindowSeconds", refs={Integer.class}, tree="[0]")
    private Output<Integer> recordingReconnectWindowSeconds;

    /**
     * @return If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.
     * 
     */
    public Output<Integer> recordingReconnectWindowSeconds() {
        return this.recordingReconnectWindowSeconds;
    }
    /**
     * The current state of the Recording Configuration.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return The current state of the Recording Configuration.
     * 
     */
    public Output<String> state() {
        return this.state;
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
     * Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    /**
     * @return Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    /**
     * Object containing information to enable/disable the recording of thumbnails for a live session and modify the interval at which thumbnails are generated for the live session.
     * 
     */
    @Export(name="thumbnailConfiguration", refs={RecordingConfigurationThumbnailConfiguration.class}, tree="[0]")
    private Output<RecordingConfigurationThumbnailConfiguration> thumbnailConfiguration;

    /**
     * @return Object containing information to enable/disable the recording of thumbnails for a live session and modify the interval at which thumbnails are generated for the live session.
     * 
     */
    public Output<RecordingConfigurationThumbnailConfiguration> thumbnailConfiguration() {
        return this.thumbnailConfiguration;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RecordingConfiguration(String name) {
        this(name, RecordingConfigurationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RecordingConfiguration(String name, RecordingConfigurationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RecordingConfiguration(String name, RecordingConfigurationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ivs/recordingConfiguration:RecordingConfiguration", name, args == null ? RecordingConfigurationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RecordingConfiguration(String name, Output<String> id, @Nullable RecordingConfigurationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ivs/recordingConfiguration:RecordingConfiguration", name, state, makeResourceOptions(options, id));
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
    public static RecordingConfiguration get(String name, Output<String> id, @Nullable RecordingConfigurationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RecordingConfiguration(name, id, state, options);
    }
}
