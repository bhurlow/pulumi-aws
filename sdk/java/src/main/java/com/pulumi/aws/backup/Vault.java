// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.backup;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.backup.VaultArgs;
import com.pulumi.aws.backup.inputs.VaultState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an AWS Backup vault resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.backup.Vault;
 * import com.pulumi.aws.backup.VaultArgs;
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
 *         var example = new Vault(&#34;example&#34;, VaultArgs.builder()        
 *             .kmsKeyArn(aws_kms_key.example().arn())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Backup vault can be imported using the `name`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:backup/vault:Vault test-vault TestVault
 * ```
 * 
 */
@ResourceType(type="aws:backup/vault:Vault")
public class Vault extends com.pulumi.resources.CustomResource {
    /**
     * The ARN of the vault.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the vault.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * A boolean that indicates that all recovery points stored in the vault are deleted so that the vault can be destroyed without error.
     * 
     */
    @Export(name="forceDestroy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceDestroy;

    /**
     * @return A boolean that indicates that all recovery points stored in the vault are deleted so that the vault can be destroyed without error.
     * 
     */
    public Output<Optional<Boolean>> forceDestroy() {
        return Codegen.optional(this.forceDestroy);
    }
    /**
     * The server-side encryption key that is used to protect your backups.
     * 
     */
    @Export(name="kmsKeyArn", refs={String.class}, tree="[0]")
    private Output<String> kmsKeyArn;

    /**
     * @return The server-side encryption key that is used to protect your backups.
     * 
     */
    public Output<String> kmsKeyArn() {
        return this.kmsKeyArn;
    }
    /**
     * Name of the backup vault to create.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the backup vault to create.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The number of recovery points that are stored in a backup vault.
     * 
     */
    @Export(name="recoveryPoints", refs={Integer.class}, tree="[0]")
    private Output<Integer> recoveryPoints;

    /**
     * @return The number of recovery points that are stored in a backup vault.
     * 
     */
    public Output<Integer> recoveryPoints() {
        return this.recoveryPoints;
    }
    /**
     * Metadata that you can assign to help organize the resources that you create. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Metadata that you can assign to help organize the resources that you create. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
    public Vault(String name) {
        this(name, VaultArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Vault(String name, @Nullable VaultArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Vault(String name, @Nullable VaultArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:backup/vault:Vault", name, args == null ? VaultArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Vault(String name, Output<String> id, @Nullable VaultState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:backup/vault:Vault", name, state, makeResourceOptions(options, id));
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
    public static Vault get(String name, Output<String> id, @Nullable VaultState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Vault(name, id, state, options);
    }
}
