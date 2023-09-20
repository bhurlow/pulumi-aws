// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.fsx;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.fsx.OpenZfsSnapshotArgs;
import com.pulumi.aws.fsx.inputs.OpenZfsSnapshotState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an Amazon FSx for OpenZFS volume.
 * See the [FSx OpenZFS User Guide](https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/what-is-fsx.html) for more information.
 * 
 * ## Example Usage
 * ### Root volume Example
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.fsx.OpenZfsFileSystem;
 * import com.pulumi.aws.fsx.OpenZfsFileSystemArgs;
 * import com.pulumi.aws.fsx.OpenZfsSnapshot;
 * import com.pulumi.aws.fsx.OpenZfsSnapshotArgs;
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
 *         var exampleOpenZfsFileSystem = new OpenZfsFileSystem(&#34;exampleOpenZfsFileSystem&#34;, OpenZfsFileSystemArgs.builder()        
 *             .storageCapacity(64)
 *             .subnetIds(aws_subnet.example().id())
 *             .deploymentType(&#34;SINGLE_AZ_1&#34;)
 *             .throughputCapacity(64)
 *             .build());
 * 
 *         var exampleOpenZfsSnapshot = new OpenZfsSnapshot(&#34;exampleOpenZfsSnapshot&#34;, OpenZfsSnapshotArgs.builder()        
 *             .volumeId(exampleOpenZfsFileSystem.rootVolumeId())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Child volume Example
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.fsx.OpenZfsFileSystem;
 * import com.pulumi.aws.fsx.OpenZfsFileSystemArgs;
 * import com.pulumi.aws.fsx.OpenZfsVolume;
 * import com.pulumi.aws.fsx.OpenZfsVolumeArgs;
 * import com.pulumi.aws.fsx.OpenZfsSnapshot;
 * import com.pulumi.aws.fsx.OpenZfsSnapshotArgs;
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
 *         var exampleOpenZfsFileSystem = new OpenZfsFileSystem(&#34;exampleOpenZfsFileSystem&#34;, OpenZfsFileSystemArgs.builder()        
 *             .storageCapacity(64)
 *             .subnetIds(aws_subnet.example().id())
 *             .deploymentType(&#34;SINGLE_AZ_1&#34;)
 *             .throughputCapacity(64)
 *             .build());
 * 
 *         var exampleOpenZfsVolume = new OpenZfsVolume(&#34;exampleOpenZfsVolume&#34;, OpenZfsVolumeArgs.builder()        
 *             .parentVolumeId(exampleOpenZfsFileSystem.rootVolumeId())
 *             .build());
 * 
 *         var exampleOpenZfsSnapshot = new OpenZfsSnapshot(&#34;exampleOpenZfsSnapshot&#34;, OpenZfsSnapshotArgs.builder()        
 *             .volumeId(exampleOpenZfsVolume.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import FSx OpenZFS snapshot using the `id`. For example:
 * 
 * ```sh
 *  $ pulumi import aws:fsx/openZfsSnapshot:OpenZfsSnapshot example fs-543ab12b1ca672f33
 * ```
 * 
 */
@ResourceType(type="aws:fsx/openZfsSnapshot:OpenZfsSnapshot")
public class OpenZfsSnapshot extends com.pulumi.resources.CustomResource {
    /**
     * Amazon Resource Name of the snapshot.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Amazon Resource Name of the snapshot.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    @Export(name="creationTime", refs={String.class}, tree="[0]")
    private Output<String> creationTime;

    public Output<String> creationTime() {
        return this.creationTime;
    }
    /**
     * The name of the Snapshot. You can use a maximum of 203 alphanumeric characters plus either _ or -  or : or . for the name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the Snapshot. You can use a maximum of 203 alphanumeric characters plus either _ or -  or : or . for the name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A map of tags to assign to the file system. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level. If you have set `copy_tags_to_backups` to true, and you specify one or more tags, no existing file system tags are copied from the file system to the backup.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the file system. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level. If you have set `copy_tags_to_backups` to true, and you specify one or more tags, no existing file system tags are copied from the file system to the backup.
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
     * The ID of the volume to snapshot. This can be the root volume or a child volume.
     * 
     */
    @Export(name="volumeId", refs={String.class}, tree="[0]")
    private Output<String> volumeId;

    /**
     * @return The ID of the volume to snapshot. This can be the root volume or a child volume.
     * 
     */
    public Output<String> volumeId() {
        return this.volumeId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public OpenZfsSnapshot(String name) {
        this(name, OpenZfsSnapshotArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OpenZfsSnapshot(String name, OpenZfsSnapshotArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OpenZfsSnapshot(String name, OpenZfsSnapshotArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:fsx/openZfsSnapshot:OpenZfsSnapshot", name, args == null ? OpenZfsSnapshotArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private OpenZfsSnapshot(String name, Output<String> id, @Nullable OpenZfsSnapshotState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:fsx/openZfsSnapshot:OpenZfsSnapshot", name, state, makeResourceOptions(options, id));
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
    public static OpenZfsSnapshot get(String name, Output<String> id, @Nullable OpenZfsSnapshotState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OpenZfsSnapshot(name, id, state, options);
    }
}
