// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.chime;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.chime.VoiceConnectorStreamingArgs;
import com.pulumi.aws.chime.inputs.VoiceConnectorStreamingState;
import com.pulumi.aws.chime.outputs.VoiceConnectorStreamingMediaInsightsConfiguration;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Adds a streaming configuration for the specified Amazon Chime Voice Connector. The streaming configuration specifies whether media streaming is enabled for sending to Amazon Kinesis.
 * It also sets the retention period, in hours, for the Amazon Kinesis data.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.chime.VoiceConnector;
 * import com.pulumi.aws.chime.VoiceConnectorArgs;
 * import com.pulumi.aws.chime.VoiceConnectorStreaming;
 * import com.pulumi.aws.chime.VoiceConnectorStreamingArgs;
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
 *         var defaultVoiceConnector = new VoiceConnector(&#34;defaultVoiceConnector&#34;, VoiceConnectorArgs.builder()        
 *             .requireEncryption(true)
 *             .build());
 * 
 *         var defaultVoiceConnectorStreaming = new VoiceConnectorStreaming(&#34;defaultVoiceConnectorStreaming&#34;, VoiceConnectorStreamingArgs.builder()        
 *             .disabled(false)
 *             .voiceConnectorId(defaultVoiceConnector.id())
 *             .dataRetention(7)
 *             .streamingNotificationTargets(&#34;SQS&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Example Usage With Media Insights
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.chime.VoiceConnector;
 * import com.pulumi.aws.chime.VoiceConnectorArgs;
 * import com.pulumi.aws.iam.IamFunctions;
 * import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
 * import com.pulumi.aws.iam.Role;
 * import com.pulumi.aws.iam.RoleArgs;
 * import com.pulumi.aws.kinesis.Stream;
 * import com.pulumi.aws.kinesis.StreamArgs;
 * import com.pulumi.aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration;
 * import com.pulumi.aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationArgs;
 * import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementArgs;
 * import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationArgs;
 * import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs;
 * import com.pulumi.aws.chime.VoiceConnectorStreaming;
 * import com.pulumi.aws.chime.VoiceConnectorStreamingArgs;
 * import com.pulumi.aws.chime.inputs.VoiceConnectorStreamingMediaInsightsConfigurationArgs;
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
 *         var defaultVoiceConnector = new VoiceConnector(&#34;defaultVoiceConnector&#34;, VoiceConnectorArgs.builder()        
 *             .requireEncryption(true)
 *             .build());
 * 
 *         final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .statements(GetPolicyDocumentStatementArgs.builder()
 *                 .effect(&#34;Allow&#34;)
 *                 .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
 *                     .type(&#34;Service&#34;)
 *                     .identifiers(&#34;mediapipelines.chime.amazonaws.com&#34;)
 *                     .build())
 *                 .actions(&#34;sts:AssumeRole&#34;)
 *                 .build())
 *             .build());
 * 
 *         var exampleRole = new Role(&#34;exampleRole&#34;, RoleArgs.builder()        
 *             .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json()))
 *             .build());
 * 
 *         var exampleStream = new Stream(&#34;exampleStream&#34;, StreamArgs.builder()        
 *             .shardCount(2)
 *             .build());
 * 
 *         var exampleMediaInsightsPipelineConfiguration = new MediaInsightsPipelineConfiguration(&#34;exampleMediaInsightsPipelineConfiguration&#34;, MediaInsightsPipelineConfigurationArgs.builder()        
 *             .resourceAccessRoleArn(exampleRole.arn())
 *             .elements(            
 *                 MediaInsightsPipelineConfigurationElementArgs.builder()
 *                     .type(&#34;AmazonTranscribeCallAnalyticsProcessor&#34;)
 *                     .amazonTranscribeCallAnalyticsProcessorConfiguration(MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationArgs.builder()
 *                         .languageCode(&#34;en-US&#34;)
 *                         .build())
 *                     .build(),
 *                 MediaInsightsPipelineConfigurationElementArgs.builder()
 *                     .type(&#34;KinesisDataStreamSink&#34;)
 *                     .kinesisDataStreamSinkConfiguration(MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs.builder()
 *                         .insightsTarget(exampleStream.arn())
 *                         .build())
 *                     .build())
 *             .build());
 * 
 *         var defaultVoiceConnectorStreaming = new VoiceConnectorStreaming(&#34;defaultVoiceConnectorStreaming&#34;, VoiceConnectorStreamingArgs.builder()        
 *             .disabled(false)
 *             .voiceConnectorId(defaultVoiceConnector.id())
 *             .dataRetention(7)
 *             .streamingNotificationTargets(&#34;SQS&#34;)
 *             .mediaInsightsConfiguration(VoiceConnectorStreamingMediaInsightsConfigurationArgs.builder()
 *                 .disabled(false)
 *                 .configurationArn(exampleMediaInsightsPipelineConfiguration.arn())
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
 *  to = aws_chime_voice_connector_streaming.default
 * 
 *  id = &#34;abcdef1ghij2klmno3pqr4&#34; } Using `pulumi import`, import Chime Voice Connector Streaming using the `voice_connector_id`. For exampleconsole % pulumi import aws_chime_voice_connector_streaming.default abcdef1ghij2klmno3pqr4
 * 
 */
@ResourceType(type="aws:chime/voiceConnectorStreaming:VoiceConnectorStreaming")
public class VoiceConnectorStreaming extends com.pulumi.resources.CustomResource {
    /**
     * The retention period, in hours, for the Amazon Kinesis data.
     * 
     */
    @Export(name="dataRetention", refs={Integer.class}, tree="[0]")
    private Output<Integer> dataRetention;

    /**
     * @return The retention period, in hours, for the Amazon Kinesis data.
     * 
     */
    public Output<Integer> dataRetention() {
        return this.dataRetention;
    }
    /**
     * When true, media streaming to Amazon Kinesis is turned off. Default: `false`
     * 
     */
    @Export(name="disabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disabled;

    /**
     * @return When true, media streaming to Amazon Kinesis is turned off. Default: `false`
     * 
     */
    public Output<Optional<Boolean>> disabled() {
        return Codegen.optional(this.disabled);
    }
    /**
     * The media insights configuration. See `media_insights_configuration`.
     * 
     */
    @Export(name="mediaInsightsConfiguration", refs={VoiceConnectorStreamingMediaInsightsConfiguration.class}, tree="[0]")
    private Output</* @Nullable */ VoiceConnectorStreamingMediaInsightsConfiguration> mediaInsightsConfiguration;

    /**
     * @return The media insights configuration. See `media_insights_configuration`.
     * 
     */
    public Output<Optional<VoiceConnectorStreamingMediaInsightsConfiguration>> mediaInsightsConfiguration() {
        return Codegen.optional(this.mediaInsightsConfiguration);
    }
    /**
     * The streaming notification targets. Valid Values: `EventBridge | SNS | SQS`
     * 
     */
    @Export(name="streamingNotificationTargets", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> streamingNotificationTargets;

    /**
     * @return The streaming notification targets. Valid Values: `EventBridge | SNS | SQS`
     * 
     */
    public Output<Optional<List<String>>> streamingNotificationTargets() {
        return Codegen.optional(this.streamingNotificationTargets);
    }
    /**
     * The Amazon Chime Voice Connector ID.
     * 
     */
    @Export(name="voiceConnectorId", refs={String.class}, tree="[0]")
    private Output<String> voiceConnectorId;

    /**
     * @return The Amazon Chime Voice Connector ID.
     * 
     */
    public Output<String> voiceConnectorId() {
        return this.voiceConnectorId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VoiceConnectorStreaming(String name) {
        this(name, VoiceConnectorStreamingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VoiceConnectorStreaming(String name, VoiceConnectorStreamingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VoiceConnectorStreaming(String name, VoiceConnectorStreamingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:chime/voiceConnectorStreaming:VoiceConnectorStreaming", name, args == null ? VoiceConnectorStreamingArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VoiceConnectorStreaming(String name, Output<String> id, @Nullable VoiceConnectorStreamingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:chime/voiceConnectorStreaming:VoiceConnectorStreaming", name, state, makeResourceOptions(options, id));
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
    public static VoiceConnectorStreaming get(String name, Output<String> id, @Nullable VoiceConnectorStreamingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VoiceConnectorStreaming(name, id, state, options);
    }
}
