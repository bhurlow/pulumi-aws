// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ebs;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ebs.SnapshotCopyArgs;
import com.pulumi.aws.ebs.inputs.SnapshotCopyState;
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
 * Creates a Snapshot of a snapshot.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ebs.Volume;
 * import com.pulumi.aws.ebs.VolumeArgs;
 * import com.pulumi.aws.ebs.Snapshot;
 * import com.pulumi.aws.ebs.SnapshotArgs;
 * import com.pulumi.aws.ebs.SnapshotCopy;
 * import com.pulumi.aws.ebs.SnapshotCopyArgs;
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
 *         var example = new Volume(&#34;example&#34;, VolumeArgs.builder()        
 *             .availabilityZone(&#34;us-west-2a&#34;)
 *             .size(40)
 *             .tags(Map.of(&#34;Name&#34;, &#34;HelloWorld&#34;))
 *             .build());
 * 
 *         var exampleSnapshot = new Snapshot(&#34;exampleSnapshot&#34;, SnapshotArgs.builder()        
 *             .volumeId(example.id())
 *             .tags(Map.of(&#34;Name&#34;, &#34;HelloWorld_snap&#34;))
 *             .build());
 * 
 *         var exampleCopy = new SnapshotCopy(&#34;exampleCopy&#34;, SnapshotCopyArgs.builder()        
 *             .sourceSnapshotId(exampleSnapshot.id())
 *             .sourceRegion(&#34;us-west-2&#34;)
 *             .tags(Map.of(&#34;Name&#34;, &#34;HelloWorld_copy_snap&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="aws:ebs/snapshotCopy:SnapshotCopy")
public class SnapshotCopy extends com.pulumi.resources.CustomResource {
    /**
     * Amazon Resource Name (ARN) of the EBS Snapshot.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of the EBS Snapshot.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The data encryption key identifier for the snapshot.
     * 
     */
    @Export(name="dataEncryptionKeyId", refs={String.class}, tree="[0]")
    private Output<String> dataEncryptionKeyId;

    /**
     * @return The data encryption key identifier for the snapshot.
     * 
     */
    public Output<String> dataEncryptionKeyId() {
        return this.dataEncryptionKeyId;
    }
    /**
     * A description of what the snapshot is.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of what the snapshot is.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Whether the snapshot is encrypted.
     * 
     */
    @Export(name="encrypted", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> encrypted;

    /**
     * @return Whether the snapshot is encrypted.
     * 
     */
    public Output<Optional<Boolean>> encrypted() {
        return Codegen.optional(this.encrypted);
    }
    /**
     * The ARN for the KMS encryption key.
     * 
     */
    @Export(name="kmsKeyId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsKeyId;

    /**
     * @return The ARN for the KMS encryption key.
     * 
     */
    public Output<Optional<String>> kmsKeyId() {
        return Codegen.optional(this.kmsKeyId);
    }
    @Export(name="outpostArn", refs={String.class}, tree="[0]")
    private Output<String> outpostArn;

    public Output<String> outpostArn() {
        return this.outpostArn;
    }
    /**
     * Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
     * 
     */
    @Export(name="ownerAlias", refs={String.class}, tree="[0]")
    private Output<String> ownerAlias;

    /**
     * @return Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
     * 
     */
    public Output<String> ownerAlias() {
        return this.ownerAlias;
    }
    /**
     * The AWS account ID of the snapshot owner.
     * 
     */
    @Export(name="ownerId", refs={String.class}, tree="[0]")
    private Output<String> ownerId;

    /**
     * @return The AWS account ID of the snapshot owner.
     * 
     */
    public Output<String> ownerId() {
        return this.ownerId;
    }
    /**
     * Indicates whether to permanently restore an archived snapshot.
     * 
     */
    @Export(name="permanentRestore", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> permanentRestore;

    /**
     * @return Indicates whether to permanently restore an archived snapshot.
     * 
     */
    public Output<Optional<Boolean>> permanentRestore() {
        return Codegen.optional(this.permanentRestore);
    }
    /**
     * The region of the source snapshot.
     * 
     */
    @Export(name="sourceRegion", refs={String.class}, tree="[0]")
    private Output<String> sourceRegion;

    /**
     * @return The region of the source snapshot.
     * 
     */
    public Output<String> sourceRegion() {
        return this.sourceRegion;
    }
    /**
     * The ARN for the snapshot to be copied.
     * 
     */
    @Export(name="sourceSnapshotId", refs={String.class}, tree="[0]")
    private Output<String> sourceSnapshotId;

    /**
     * @return The ARN for the snapshot to be copied.
     * 
     */
    public Output<String> sourceSnapshotId() {
        return this.sourceSnapshotId;
    }
    /**
     * The name of the storage tier. Valid values are `archive` and `standard`. Default value is `standard`.
     * 
     */
    @Export(name="storageTier", refs={String.class}, tree="[0]")
    private Output<String> storageTier;

    /**
     * @return The name of the storage tier. Valid values are `archive` and `standard`. Default value is `standard`.
     * 
     */
    public Output<String> storageTier() {
        return this.storageTier;
    }
    /**
     * A map of tags for the snapshot.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags for the snapshot.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
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
     * Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
     * 
     */
    @Export(name="temporaryRestoreDays", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> temporaryRestoreDays;

    /**
     * @return Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
     * 
     */
    public Output<Optional<Integer>> temporaryRestoreDays() {
        return Codegen.optional(this.temporaryRestoreDays);
    }
    @Export(name="volumeId", refs={String.class}, tree="[0]")
    private Output<String> volumeId;

    public Output<String> volumeId() {
        return this.volumeId;
    }
    /**
     * The size of the drive in GiBs.
     * 
     */
    @Export(name="volumeSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> volumeSize;

    /**
     * @return The size of the drive in GiBs.
     * 
     */
    public Output<Integer> volumeSize() {
        return this.volumeSize;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SnapshotCopy(String name) {
        this(name, SnapshotCopyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SnapshotCopy(String name, SnapshotCopyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SnapshotCopy(String name, SnapshotCopyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ebs/snapshotCopy:SnapshotCopy", name, args == null ? SnapshotCopyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SnapshotCopy(String name, Output<String> id, @Nullable SnapshotCopyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ebs/snapshotCopy:SnapshotCopy", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "tagsAll"
            ))
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
    public static SnapshotCopy get(String name, Output<String> id, @Nullable SnapshotCopyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SnapshotCopy(name, id, state, options);
    }
}
