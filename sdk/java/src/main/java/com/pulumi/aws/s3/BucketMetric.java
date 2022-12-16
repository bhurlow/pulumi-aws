// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.s3;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.s3.BucketMetricArgs;
import com.pulumi.aws.s3.inputs.BucketMetricState;
import com.pulumi.aws.s3.outputs.BucketMetricFilter;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a S3 bucket [metrics configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html) resource.
 * 
 * ## Example Usage
 * ### Add metrics configuration for entire S3 bucket
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.s3.BucketV2;
 * import com.pulumi.aws.s3.BucketMetric;
 * import com.pulumi.aws.s3.BucketMetricArgs;
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
 *         var example = new BucketV2(&#34;example&#34;);
 * 
 *         var example_entire_bucket = new BucketMetric(&#34;example-entire-bucket&#34;, BucketMetricArgs.builder()        
 *             .bucket(example.bucket())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Add metrics configuration with S3 object filter
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.s3.BucketV2;
 * import com.pulumi.aws.s3.BucketMetric;
 * import com.pulumi.aws.s3.BucketMetricArgs;
 * import com.pulumi.aws.s3.inputs.BucketMetricFilterArgs;
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
 *         var example = new BucketV2(&#34;example&#34;);
 * 
 *         var example_filtered = new BucketMetric(&#34;example-filtered&#34;, BucketMetricArgs.builder()        
 *             .bucket(example.bucket())
 *             .filter(BucketMetricFilterArgs.builder()
 *                 .prefix(&#34;documents/&#34;)
 *                 .tags(Map.ofEntries(
 *                     Map.entry(&#34;priority&#34;, &#34;high&#34;),
 *                     Map.entry(&#34;class&#34;, &#34;blue&#34;)
 *                 ))
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * S3 bucket metric configurations can be imported using `bucket:metric`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:s3/bucketMetric:BucketMetric my-bucket-entire-bucket my-bucket:EntireBucket
 * ```
 * 
 */
@ResourceType(type="aws:s3/bucketMetric:BucketMetric")
public class BucketMetric extends com.pulumi.resources.CustomResource {
    /**
     * The name of the bucket to put metric configuration.
     * 
     */
    @Export(name="bucket", refs={String.class}, tree="[0]")
    private Output<String> bucket;

    /**
     * @return The name of the bucket to put metric configuration.
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
    }
    /**
     * [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
     * 
     */
    @Export(name="filter", refs={BucketMetricFilter.class}, tree="[0]")
    private Output</* @Nullable */ BucketMetricFilter> filter;

    /**
     * @return [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
     * 
     */
    public Output<Optional<BucketMetricFilter>> filter() {
        return Codegen.optional(this.filter);
    }
    /**
     * Unique identifier of the metrics configuration for the bucket. Must be less than or equal to 64 characters in length.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Unique identifier of the metrics configuration for the bucket. Must be less than or equal to 64 characters in length.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BucketMetric(String name) {
        this(name, BucketMetricArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BucketMetric(String name, BucketMetricArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BucketMetric(String name, BucketMetricArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:s3/bucketMetric:BucketMetric", name, args == null ? BucketMetricArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BucketMetric(String name, Output<String> id, @Nullable BucketMetricState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:s3/bucketMetric:BucketMetric", name, state, makeResourceOptions(options, id));
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
    public static BucketMetric get(String name, Output<String> id, @Nullable BucketMetricState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BucketMetric(name, id, state, options);
    }
}
