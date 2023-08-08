// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.efs;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.efs.AccessPointArgs;
import com.pulumi.aws.efs.inputs.AccessPointState;
import com.pulumi.aws.efs.outputs.AccessPointPosixUser;
import com.pulumi.aws.efs.outputs.AccessPointRootDirectory;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Elastic File System (EFS) access point.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.efs.AccessPoint;
 * import com.pulumi.aws.efs.AccessPointArgs;
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
 *         var test = new AccessPoint(&#34;test&#34;, AccessPointArgs.builder()        
 *             .fileSystemId(aws_efs_file_system.foo().id())
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
 *  to = aws_efs_access_point.test
 * 
 *  id = &#34;fsap-52a643fb&#34; } Using `pulumi import`, import the EFS access points using the `id`. For exampleconsole % pulumi import aws_efs_access_point.test fsap-52a643fb
 * 
 */
@ResourceType(type="aws:efs/accessPoint:AccessPoint")
public class AccessPoint extends com.pulumi.resources.CustomResource {
    /**
     * ARN of the access point.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return ARN of the access point.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * ARN of the file system.
     * 
     */
    @Export(name="fileSystemArn", refs={String.class}, tree="[0]")
    private Output<String> fileSystemArn;

    /**
     * @return ARN of the file system.
     * 
     */
    public Output<String> fileSystemArn() {
        return this.fileSystemArn;
    }
    /**
     * ID of the file system for which the access point is intended.
     * 
     */
    @Export(name="fileSystemId", refs={String.class}, tree="[0]")
    private Output<String> fileSystemId;

    /**
     * @return ID of the file system for which the access point is intended.
     * 
     */
    public Output<String> fileSystemId() {
        return this.fileSystemId;
    }
    @Export(name="ownerId", refs={String.class}, tree="[0]")
    private Output<String> ownerId;

    public Output<String> ownerId() {
        return this.ownerId;
    }
    /**
     * Operating system user and group applied to all file system requests made using the access point. Detailed below.
     * 
     */
    @Export(name="posixUser", refs={AccessPointPosixUser.class}, tree="[0]")
    private Output</* @Nullable */ AccessPointPosixUser> posixUser;

    /**
     * @return Operating system user and group applied to all file system requests made using the access point. Detailed below.
     * 
     */
    public Output<Optional<AccessPointPosixUser>> posixUser() {
        return Codegen.optional(this.posixUser);
    }
    /**
     * Directory on the Amazon EFS file system that the access point provides access to. Detailed below.
     * 
     */
    @Export(name="rootDirectory", refs={AccessPointRootDirectory.class}, tree="[0]")
    private Output<AccessPointRootDirectory> rootDirectory;

    /**
     * @return Directory on the Amazon EFS file system that the access point provides access to. Detailed below.
     * 
     */
    public Output<AccessPointRootDirectory> rootDirectory() {
        return this.rootDirectory;
    }
    /**
     * Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AccessPoint(String name) {
        this(name, AccessPointArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AccessPoint(String name, AccessPointArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AccessPoint(String name, AccessPointArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:efs/accessPoint:AccessPoint", name, args == null ? AccessPointArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AccessPoint(String name, Output<String> id, @Nullable AccessPointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:efs/accessPoint:AccessPoint", name, state, makeResourceOptions(options, id));
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
    public static AccessPoint get(String name, Output<String> id, @Nullable AccessPointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AccessPoint(name, id, state, options);
    }
}
