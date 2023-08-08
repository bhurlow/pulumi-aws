// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.workspaces;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.workspaces.WorkspaceArgs;
import com.pulumi.aws.workspaces.inputs.WorkspaceState;
import com.pulumi.aws.workspaces.outputs.WorkspaceWorkspaceProperties;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a workspace in [AWS Workspaces](https://docs.aws.amazon.com/workspaces/latest/adminguide/amazon-workspaces.html) Service
 * 
 * &gt; **NOTE:** AWS WorkSpaces service requires [`workspaces_DefaultRole`](https://docs.aws.amazon.com/workspaces/latest/adminguide/workspaces-access-control.html#create-default-role) IAM role to operate normally.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.workspaces.WorkspacesFunctions;
 * import com.pulumi.aws.workspaces.inputs.GetBundleArgs;
 * import com.pulumi.aws.workspaces.Workspace;
 * import com.pulumi.aws.workspaces.WorkspaceArgs;
 * import com.pulumi.aws.workspaces.inputs.WorkspaceWorkspacePropertiesArgs;
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
 *         final var valueWindows10 = WorkspacesFunctions.getBundle(GetBundleArgs.builder()
 *             .bundleId(&#34;wsb-bh8rsxt14&#34;)
 *             .build());
 * 
 *         var example = new Workspace(&#34;example&#34;, WorkspaceArgs.builder()        
 *             .directoryId(aws_workspaces_directory.example().id())
 *             .bundleId(valueWindows10.applyValue(getBundleResult -&gt; getBundleResult.id()))
 *             .userName(&#34;john.doe&#34;)
 *             .rootVolumeEncryptionEnabled(true)
 *             .userVolumeEncryptionEnabled(true)
 *             .volumeEncryptionKey(&#34;alias/aws/workspaces&#34;)
 *             .workspaceProperties(WorkspaceWorkspacePropertiesArgs.builder()
 *                 .computeTypeName(&#34;VALUE&#34;)
 *                 .userVolumeSizeGib(10)
 *                 .rootVolumeSizeGib(80)
 *                 .runningMode(&#34;AUTO_STOP&#34;)
 *                 .runningModeAutoStopTimeoutInMinutes(60)
 *                 .build())
 *             .tags(Map.of(&#34;Department&#34;, &#34;IT&#34;))
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
 *  to = aws_workspaces_workspace.example
 * 
 *  id = &#34;ws-9z9zmbkhv&#34; } Using `pulumi import`, import Workspaces using their ID. For exampleconsole % pulumi import aws_workspaces_workspace.example ws-9z9zmbkhv
 * 
 */
@ResourceType(type="aws:workspaces/workspace:Workspace")
public class Workspace extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the bundle for the WorkSpace.
     * 
     */
    @Export(name="bundleId", refs={String.class}, tree="[0]")
    private Output<String> bundleId;

    /**
     * @return The ID of the bundle for the WorkSpace.
     * 
     */
    public Output<String> bundleId() {
        return this.bundleId;
    }
    /**
     * The name of the WorkSpace, as seen by the operating system.
     * 
     */
    @Export(name="computerName", refs={String.class}, tree="[0]")
    private Output<String> computerName;

    /**
     * @return The name of the WorkSpace, as seen by the operating system.
     * 
     */
    public Output<String> computerName() {
        return this.computerName;
    }
    /**
     * The ID of the directory for the WorkSpace.
     * 
     */
    @Export(name="directoryId", refs={String.class}, tree="[0]")
    private Output<String> directoryId;

    /**
     * @return The ID of the directory for the WorkSpace.
     * 
     */
    public Output<String> directoryId() {
        return this.directoryId;
    }
    /**
     * The IP address of the WorkSpace.
     * 
     */
    @Export(name="ipAddress", refs={String.class}, tree="[0]")
    private Output<String> ipAddress;

    /**
     * @return The IP address of the WorkSpace.
     * 
     */
    public Output<String> ipAddress() {
        return this.ipAddress;
    }
    /**
     * Indicates whether the data stored on the root volume is encrypted.
     * 
     */
    @Export(name="rootVolumeEncryptionEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> rootVolumeEncryptionEnabled;

    /**
     * @return Indicates whether the data stored on the root volume is encrypted.
     * 
     */
    public Output<Optional<Boolean>> rootVolumeEncryptionEnabled() {
        return Codegen.optional(this.rootVolumeEncryptionEnabled);
    }
    /**
     * The operational state of the WorkSpace.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return The operational state of the WorkSpace.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The tags for the WorkSpace. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return The tags for the WorkSpace. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * The user name of the user for the WorkSpace. This user name must exist in the directory for the WorkSpace.
     * 
     */
    @Export(name="userName", refs={String.class}, tree="[0]")
    private Output<String> userName;

    /**
     * @return The user name of the user for the WorkSpace. This user name must exist in the directory for the WorkSpace.
     * 
     */
    public Output<String> userName() {
        return this.userName;
    }
    /**
     * Indicates whether the data stored on the user volume is encrypted.
     * 
     */
    @Export(name="userVolumeEncryptionEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> userVolumeEncryptionEnabled;

    /**
     * @return Indicates whether the data stored on the user volume is encrypted.
     * 
     */
    public Output<Optional<Boolean>> userVolumeEncryptionEnabled() {
        return Codegen.optional(this.userVolumeEncryptionEnabled);
    }
    /**
     * The symmetric AWS KMS customer master key (CMK) used to encrypt data stored on your WorkSpace. Amazon WorkSpaces does not support asymmetric CMKs.
     * 
     */
    @Export(name="volumeEncryptionKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> volumeEncryptionKey;

    /**
     * @return The symmetric AWS KMS customer master key (CMK) used to encrypt data stored on your WorkSpace. Amazon WorkSpaces does not support asymmetric CMKs.
     * 
     */
    public Output<Optional<String>> volumeEncryptionKey() {
        return Codegen.optional(this.volumeEncryptionKey);
    }
    /**
     * The WorkSpace properties.
     * 
     */
    @Export(name="workspaceProperties", refs={WorkspaceWorkspaceProperties.class}, tree="[0]")
    private Output<WorkspaceWorkspaceProperties> workspaceProperties;

    /**
     * @return The WorkSpace properties.
     * 
     */
    public Output<WorkspaceWorkspaceProperties> workspaceProperties() {
        return this.workspaceProperties;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Workspace(String name) {
        this(name, WorkspaceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Workspace(String name, WorkspaceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Workspace(String name, WorkspaceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:workspaces/workspace:Workspace", name, args == null ? WorkspaceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Workspace(String name, Output<String> id, @Nullable WorkspaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:workspaces/workspace:Workspace", name, state, makeResourceOptions(options, id));
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
    public static Workspace get(String name, Output<String> id, @Nullable WorkspaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Workspace(name, id, state, options);
    }
}
