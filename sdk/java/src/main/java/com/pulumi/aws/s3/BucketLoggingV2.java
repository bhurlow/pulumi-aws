// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.s3;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.s3.BucketLoggingV2Args;
import com.pulumi.aws.s3.inputs.BucketLoggingV2State;
import com.pulumi.aws.s3.outputs.BucketLoggingV2TargetGrant;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an S3 bucket (server access) logging resource. For more information, see [Logging requests using server access logging](https://docs.aws.amazon.com/AmazonS3/latest/userguide/ServerLogs.html)
 * in the AWS S3 User Guide.
 * 
 * &gt; **Note:** Amazon S3 supports server access logging, AWS CloudTrail, or a combination of both. Refer to the [Logging options for Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/userguide/logging-with-S3.html)
 * to decide which method meets your requirements.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.s3.BucketV2;
 * import com.pulumi.aws.s3.BucketAclV2;
 * import com.pulumi.aws.s3.BucketAclV2Args;
 * import com.pulumi.aws.s3.BucketLoggingV2;
 * import com.pulumi.aws.s3.BucketLoggingV2Args;
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
 *         var exampleBucketV2 = new BucketV2(&#34;exampleBucketV2&#34;);
 * 
 *         var exampleBucketAclV2 = new BucketAclV2(&#34;exampleBucketAclV2&#34;, BucketAclV2Args.builder()        
 *             .bucket(exampleBucketV2.id())
 *             .acl(&#34;private&#34;)
 *             .build());
 * 
 *         var logBucket = new BucketV2(&#34;logBucket&#34;);
 * 
 *         var logBucketAcl = new BucketAclV2(&#34;logBucketAcl&#34;, BucketAclV2Args.builder()        
 *             .bucket(logBucket.id())
 *             .acl(&#34;log-delivery-write&#34;)
 *             .build());
 * 
 *         var exampleBucketLoggingV2 = new BucketLoggingV2(&#34;exampleBucketLoggingV2&#34;, BucketLoggingV2Args.builder()        
 *             .bucket(exampleBucketV2.id())
 *             .targetBucket(logBucket.id())
 *             .targetPrefix(&#34;log/&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * S3 bucket logging can be imported in one of two ways. If the owner (account ID) of the source bucket is the same account used to configure the AWS Provider, the S3 bucket logging resource should be imported using the `bucket` e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:s3/bucketLoggingV2:BucketLoggingV2 example bucket-name
 * ```
 * 
 *  If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, the S3 bucket logging resource should be imported using the `bucket` and `expected_bucket_owner` separated by a comma (`,`) e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:s3/bucketLoggingV2:BucketLoggingV2 example bucket-name,123456789012
 * ```
 * 
 */
@ResourceType(type="aws:s3/bucketLoggingV2:BucketLoggingV2")
public class BucketLoggingV2 extends com.pulumi.resources.CustomResource {
    /**
     * The name of the bucket.
     * 
     */
    @Export(name="bucket", refs={String.class}, tree="[0]")
    private Output<String> bucket;

    /**
     * @return The name of the bucket.
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
    }
    /**
     * The account ID of the expected bucket owner.
     * 
     */
    @Export(name="expectedBucketOwner", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> expectedBucketOwner;

    /**
     * @return The account ID of the expected bucket owner.
     * 
     */
    public Output<Optional<String>> expectedBucketOwner() {
        return Codegen.optional(this.expectedBucketOwner);
    }
    /**
     * The name of the bucket where you want Amazon S3 to store server access logs.
     * 
     */
    @Export(name="targetBucket", refs={String.class}, tree="[0]")
    private Output<String> targetBucket;

    /**
     * @return The name of the bucket where you want Amazon S3 to store server access logs.
     * 
     */
    public Output<String> targetBucket() {
        return this.targetBucket;
    }
    /**
     * Set of configuration blocks with information for granting permissions documented below.
     * 
     */
    @Export(name="targetGrants", refs={List.class,BucketLoggingV2TargetGrant.class}, tree="[0,1]")
    private Output</* @Nullable */ List<BucketLoggingV2TargetGrant>> targetGrants;

    /**
     * @return Set of configuration blocks with information for granting permissions documented below.
     * 
     */
    public Output<Optional<List<BucketLoggingV2TargetGrant>>> targetGrants() {
        return Codegen.optional(this.targetGrants);
    }
    /**
     * A prefix for all log object keys.
     * 
     */
    @Export(name="targetPrefix", refs={String.class}, tree="[0]")
    private Output<String> targetPrefix;

    /**
     * @return A prefix for all log object keys.
     * 
     */
    public Output<String> targetPrefix() {
        return this.targetPrefix;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BucketLoggingV2(String name) {
        this(name, BucketLoggingV2Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BucketLoggingV2(String name, BucketLoggingV2Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BucketLoggingV2(String name, BucketLoggingV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:s3/bucketLoggingV2:BucketLoggingV2", name, args == null ? BucketLoggingV2Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BucketLoggingV2(String name, Output<String> id, @Nullable BucketLoggingV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:s3/bucketLoggingV2:BucketLoggingV2", name, state, makeResourceOptions(options, id));
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
    public static BucketLoggingV2 get(String name, Output<String> id, @Nullable BucketLoggingV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BucketLoggingV2(name, id, state, options);
    }
}
