// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.s3;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.s3.BucketServerSideEncryptionConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketServerSideEncryptionConfigurationV2State;
import com.pulumi.aws.s3.outputs.BucketServerSideEncryptionConfigurationV2Rule;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a S3 bucket server-side encryption configuration resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.kms.Key;
 * import com.pulumi.aws.kms.KeyArgs;
 * import com.pulumi.aws.s3.BucketV2;
 * import com.pulumi.aws.s3.BucketServerSideEncryptionConfigurationV2;
 * import com.pulumi.aws.s3.BucketServerSideEncryptionConfigurationV2Args;
 * import com.pulumi.aws.s3.inputs.BucketServerSideEncryptionConfigurationV2RuleArgs;
 * import com.pulumi.aws.s3.inputs.BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs;
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
 *         var mykey = new Key(&#34;mykey&#34;, KeyArgs.builder()        
 *             .description(&#34;This key is used to encrypt bucket objects&#34;)
 *             .deletionWindowInDays(10)
 *             .build());
 * 
 *         var mybucket = new BucketV2(&#34;mybucket&#34;);
 * 
 *         var example = new BucketServerSideEncryptionConfigurationV2(&#34;example&#34;, BucketServerSideEncryptionConfigurationV2Args.builder()        
 *             .bucket(mybucket.bucket())
 *             .rules(BucketServerSideEncryptionConfigurationV2RuleArgs.builder()
 *                 .applyServerSideEncryptionByDefault(BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs.builder()
 *                     .kmsMasterKeyId(mykey.arn())
 *                     .sseAlgorithm(&#34;aws:kms&#34;)
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
 * S3 bucket server-side encryption configuration can be imported in one of two ways. If the owner (account ID) of the source bucket is the same account used to configure the AWS Provider, the S3 server-side encryption configuration resource should be imported using the `bucket` e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:s3/bucketServerSideEncryptionConfigurationV2:BucketServerSideEncryptionConfigurationV2 example bucket-name
 * ```
 * 
 *  If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, the S3 bucket server-side encryption configuration resource should be imported using the `bucket` and `expected_bucket_owner` separated by a comma (`,`) e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:s3/bucketServerSideEncryptionConfigurationV2:BucketServerSideEncryptionConfigurationV2 example bucket-name,123456789012
 * ```
 * 
 */
@ResourceType(type="aws:s3/bucketServerSideEncryptionConfigurationV2:BucketServerSideEncryptionConfigurationV2")
public class BucketServerSideEncryptionConfigurationV2 extends com.pulumi.resources.CustomResource {
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
     * Set of server-side encryption configuration rules. documented below. Currently, only a single rule is supported.
     * 
     */
    @Export(name="rules", refs={List.class,BucketServerSideEncryptionConfigurationV2Rule.class}, tree="[0,1]")
    private Output<List<BucketServerSideEncryptionConfigurationV2Rule>> rules;

    /**
     * @return Set of server-side encryption configuration rules. documented below. Currently, only a single rule is supported.
     * 
     */
    public Output<List<BucketServerSideEncryptionConfigurationV2Rule>> rules() {
        return this.rules;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BucketServerSideEncryptionConfigurationV2(String name) {
        this(name, BucketServerSideEncryptionConfigurationV2Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BucketServerSideEncryptionConfigurationV2(String name, BucketServerSideEncryptionConfigurationV2Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BucketServerSideEncryptionConfigurationV2(String name, BucketServerSideEncryptionConfigurationV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:s3/bucketServerSideEncryptionConfigurationV2:BucketServerSideEncryptionConfigurationV2", name, args == null ? BucketServerSideEncryptionConfigurationV2Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BucketServerSideEncryptionConfigurationV2(String name, Output<String> id, @Nullable BucketServerSideEncryptionConfigurationV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:s3/bucketServerSideEncryptionConfigurationV2:BucketServerSideEncryptionConfigurationV2", name, state, makeResourceOptions(options, id));
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
    public static BucketServerSideEncryptionConfigurationV2 get(String name, Output<String> id, @Nullable BucketServerSideEncryptionConfigurationV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BucketServerSideEncryptionConfigurationV2(name, id, state, options);
    }
}
