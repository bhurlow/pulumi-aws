// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.s3control;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.s3control.BucketPolicyArgs;
import com.pulumi.aws.s3control.inputs.BucketPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a resource to manage an S3 Control Bucket Policy.
 * 
 * &gt; This functionality is for managing [S3 on Outposts](https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html). To manage S3 Bucket Policies in an AWS Partition, see the `aws.s3.BucketPolicy` resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.s3control.BucketPolicy;
 * import com.pulumi.aws.s3control.BucketPolicyArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
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
 *         var example = new BucketPolicy(&#34;example&#34;, BucketPolicyArgs.builder()        
 *             .bucket(aws_s3control_bucket.example().arn())
 *             .policy(serializeJson(
 *                 jsonObject(
 *                     jsonProperty(&#34;Id&#34;, &#34;testBucketPolicy&#34;),
 *                     jsonProperty(&#34;Statement&#34;, jsonArray(jsonObject(
 *                         jsonProperty(&#34;Action&#34;, &#34;s3-outposts:PutBucketLifecycleConfiguration&#34;),
 *                         jsonProperty(&#34;Effect&#34;, &#34;Deny&#34;),
 *                         jsonProperty(&#34;Principal&#34;, jsonObject(
 *                             jsonProperty(&#34;AWS&#34;, &#34;*&#34;)
 *                         )),
 *                         jsonProperty(&#34;Resource&#34;, aws_s3control_bucket.example().arn()),
 *                         jsonProperty(&#34;Sid&#34;, &#34;statement1&#34;)
 *                     ))),
 *                     jsonProperty(&#34;Version&#34;, &#34;2012-10-17&#34;)
 *                 )))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * S3 Control Bucket Policies can be imported using the Amazon Resource Name (ARN), e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:s3control/bucketPolicy:BucketPolicy example arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-12345678/bucket/example
 * ```
 * 
 */
@ResourceType(type="aws:s3control/bucketPolicy:BucketPolicy")
public class BucketPolicy extends com.pulumi.resources.CustomResource {
    /**
     * Amazon Resource Name (ARN) of the bucket.
     * 
     */
    @Export(name="bucket", refs={String.class}, tree="[0]")
    private Output<String> bucket;

    /**
     * @return Amazon Resource Name (ARN) of the bucket.
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
    }
    /**
     * JSON string of the resource policy.
     * 
     */
    @Export(name="policy", refs={String.class}, tree="[0]")
    private Output<String> policy;

    /**
     * @return JSON string of the resource policy.
     * 
     */
    public Output<String> policy() {
        return this.policy;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BucketPolicy(String name) {
        this(name, BucketPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BucketPolicy(String name, BucketPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BucketPolicy(String name, BucketPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:s3control/bucketPolicy:BucketPolicy", name, args == null ? BucketPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BucketPolicy(String name, Output<String> id, @Nullable BucketPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:s3control/bucketPolicy:BucketPolicy", name, state, makeResourceOptions(options, id));
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
    public static BucketPolicy get(String name, Output<String> id, @Nullable BucketPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BucketPolicy(name, id, state, options);
    }
}
