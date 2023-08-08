// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kinesis;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.kinesis.StreamArgs;
import com.pulumi.aws.kinesis.inputs.StreamState;
import com.pulumi.aws.kinesis.outputs.StreamStreamModeDetails;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Kinesis Stream resource. Amazon Kinesis is a managed service that
 * scales elastically for real-time processing of streaming big data.
 * 
 * For more details, see the [Amazon Kinesis Documentation](https://aws.amazon.com/documentation/kinesis/).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.kinesis.Stream;
 * import com.pulumi.aws.kinesis.StreamArgs;
 * import com.pulumi.aws.kinesis.inputs.StreamStreamModeDetailsArgs;
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
 *         var testStream = new Stream(&#34;testStream&#34;, StreamArgs.builder()        
 *             .retentionPeriod(48)
 *             .shardCount(1)
 *             .shardLevelMetrics(            
 *                 &#34;IncomingBytes&#34;,
 *                 &#34;OutgoingBytes&#34;)
 *             .streamModeDetails(StreamStreamModeDetailsArgs.builder()
 *                 .streamMode(&#34;PROVISIONED&#34;)
 *                 .build())
 *             .tags(Map.of(&#34;Environment&#34;, &#34;test&#34;))
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
 *  to = aws_kinesis_stream.test_stream
 * 
 *  id = &#34;TODO-kinesis-test&#34; } Using `pulumi import`, import Kinesis Streams using the `name`. For exampleconsole % pulumi import aws_kinesis_stream.test_stream TODO-kinesis-test [1]https://aws.amazon.com/documentation/kinesis/ [2]https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html [3]https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html
 * 
 */
@ResourceType(type="aws:kinesis/stream:Stream")
public class Stream extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
     * 
     */
    @Export(name="encryptionType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> encryptionType;

    /**
     * @return The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
     * 
     */
    public Output<Optional<String>> encryptionType() {
        return Codegen.optional(this.encryptionType);
    }
    /**
     * A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
     * 
     */
    @Export(name="enforceConsumerDeletion", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enforceConsumerDeletion;

    /**
     * @return A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> enforceConsumerDeletion() {
        return Codegen.optional(this.enforceConsumerDeletion);
    }
    /**
     * The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
     * 
     */
    @Export(name="kmsKeyId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsKeyId;

    /**
     * @return The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
     * 
     */
    public Output<Optional<String>> kmsKeyId() {
        return Codegen.optional(this.kmsKeyId);
    }
    /**
     * A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Length of time data records are accessible after they are added to the stream. The maximum value of a stream&#39;s retention period is 8760 hours. Minimum value is 24. Default is 24.
     * 
     */
    @Export(name="retentionPeriod", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> retentionPeriod;

    /**
     * @return Length of time data records are accessible after they are added to the stream. The maximum value of a stream&#39;s retention period is 8760 hours. Minimum value is 24. Default is 24.
     * 
     */
    public Output<Optional<Integer>> retentionPeriod() {
        return Codegen.optional(this.retentionPeriod);
    }
    /**
     * The number of shards that the stream will use. If the `stream_mode` is `PROVISIONED`, this field is required.
     * Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
     * 
     */
    @Export(name="shardCount", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> shardCount;

    /**
     * @return The number of shards that the stream will use. If the `stream_mode` is `PROVISIONED`, this field is required.
     * Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
     * 
     */
    public Output<Optional<Integer>> shardCount() {
        return Codegen.optional(this.shardCount);
    }
    /**
     * A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
     * 
     */
    @Export(name="shardLevelMetrics", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> shardLevelMetrics;

    /**
     * @return A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
     * 
     */
    public Output<Optional<List<String>>> shardLevelMetrics() {
        return Codegen.optional(this.shardLevelMetrics);
    }
    /**
     * Indicates the [capacity mode](https://docs.aws.amazon.com/streams/latest/dev/how-do-i-size-a-stream.html) of the data stream. Detailed below.
     * 
     */
    @Export(name="streamModeDetails", refs={StreamStreamModeDetails.class}, tree="[0]")
    private Output<StreamStreamModeDetails> streamModeDetails;

    /**
     * @return Indicates the [capacity mode](https://docs.aws.amazon.com/streams/latest/dev/how-do-i-size-a-stream.html) of the data stream. Detailed below.
     * 
     */
    public Output<StreamStreamModeDetails> streamModeDetails() {
        return this.streamModeDetails;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Stream(String name) {
        this(name, StreamArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Stream(String name, @Nullable StreamArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Stream(String name, @Nullable StreamArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:kinesis/stream:Stream", name, args == null ? StreamArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Stream(String name, Output<String> id, @Nullable StreamState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:kinesis/stream:Stream", name, state, makeResourceOptions(options, id));
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
    public static Stream get(String name, Output<String> id, @Nullable StreamState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Stream(name, id, state, options);
    }
}
