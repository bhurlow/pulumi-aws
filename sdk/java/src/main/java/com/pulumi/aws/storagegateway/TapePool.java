// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.storagegateway;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.storagegateway.TapePoolArgs;
import com.pulumi.aws.storagegateway.inputs.TapePoolState;
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
 * Manages an AWS Storage Gateway Tape Pool.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.storagegateway.TapePool;
 * import com.pulumi.aws.storagegateway.TapePoolArgs;
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
 *         var example = new TapePool(&#34;example&#34;, TapePoolArgs.builder()        
 *             .poolName(&#34;example&#34;)
 *             .storageClass(&#34;GLACIER&#34;)
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
 *  to = aws_storagegateway_tape_pool.example
 * 
 *  id = &#34;arn:aws:storagegateway:us-east-1:123456789012:tapepool/pool-12345678&#34; } Using `pulumi import`, import `aws_storagegateway_tape_pool` using the volume Amazon Resource Name (ARN). For exampleconsole % pulumi import aws_storagegateway_tape_pool.example arn:aws:storagegateway:us-east-1:123456789012:tapepool/pool-12345678
 * 
 */
@ResourceType(type="aws:storagegateway/tapePool:TapePool")
public class TapePool extends com.pulumi.resources.CustomResource {
    /**
     * Volume Amazon Resource Name (ARN), e.g., `aws_storagegateway_tape_pool.example arn:aws:storagegateway:us-east-1:123456789012:tapepool/pool-12345678`.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Volume Amazon Resource Name (ARN), e.g., `aws_storagegateway_tape_pool.example arn:aws:storagegateway:us-east-1:123456789012:tapepool/pool-12345678`.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The name of the new custom tape pool.
     * 
     */
    @Export(name="poolName", refs={String.class}, tree="[0]")
    private Output<String> poolName;

    /**
     * @return The name of the new custom tape pool.
     * 
     */
    public Output<String> poolName() {
        return this.poolName;
    }
    /**
     * Tape retention lock time is set in days. Tape retention lock can be enabled for up to 100 years (36,500 days). Default value is 0.
     * 
     */
    @Export(name="retentionLockTimeInDays", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> retentionLockTimeInDays;

    /**
     * @return Tape retention lock time is set in days. Tape retention lock can be enabled for up to 100 years (36,500 days). Default value is 0.
     * 
     */
    public Output<Optional<Integer>> retentionLockTimeInDays() {
        return Codegen.optional(this.retentionLockTimeInDays);
    }
    /**
     * Tape retention lock can be configured in two modes. When configured in governance mode, AWS accounts with specific IAM permissions are authorized to remove the tape retention lock from archived virtual tapes. When configured in compliance mode, the tape retention lock cannot be removed by any user, including the root AWS account. Possible values are `COMPLIANCE`, `GOVERNANCE`, and `NONE`. Default value is `NONE`.
     * 
     */
    @Export(name="retentionLockType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> retentionLockType;

    /**
     * @return Tape retention lock can be configured in two modes. When configured in governance mode, AWS accounts with specific IAM permissions are authorized to remove the tape retention lock from archived virtual tapes. When configured in compliance mode, the tape retention lock cannot be removed by any user, including the root AWS account. Possible values are `COMPLIANCE`, `GOVERNANCE`, and `NONE`. Default value is `NONE`.
     * 
     */
    public Output<Optional<String>> retentionLockType() {
        return Codegen.optional(this.retentionLockType);
    }
    /**
     * The storage class that is associated with the new custom pool. When you use your backup application to eject the tape, the tape is archived directly into the storage class that corresponds to the pool. Possible values are `DEEP_ARCHIVE` or `GLACIER`.
     * 
     */
    @Export(name="storageClass", refs={String.class}, tree="[0]")
    private Output<String> storageClass;

    /**
     * @return The storage class that is associated with the new custom pool. When you use your backup application to eject the tape, the tape is archived directly into the storage class that corresponds to the pool. Possible values are `DEEP_ARCHIVE` or `GLACIER`.
     * 
     */
    public Output<String> storageClass() {
        return this.storageClass;
    }
    /**
     * Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
    public TapePool(String name) {
        this(name, TapePoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TapePool(String name, TapePoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TapePool(String name, TapePoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:storagegateway/tapePool:TapePool", name, args == null ? TapePoolArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TapePool(String name, Output<String> id, @Nullable TapePoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:storagegateway/tapePool:TapePool", name, state, makeResourceOptions(options, id));
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
    public static TapePool get(String name, Output<String> id, @Nullable TapePoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TapePool(name, id, state, options);
    }
}
