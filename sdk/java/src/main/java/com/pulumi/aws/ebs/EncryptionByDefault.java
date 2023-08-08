// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ebs;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ebs.EncryptionByDefaultArgs;
import com.pulumi.aws.ebs.inputs.EncryptionByDefaultState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource to manage whether default EBS encryption is enabled for your AWS account in the current AWS region. To manage the default KMS key for the region, see the `aws.ebs.DefaultKmsKey` resource.
 * 
 * &gt; **NOTE:** Removing this resource disables default EBS encryption.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ebs.EncryptionByDefault;
 * import com.pulumi.aws.ebs.EncryptionByDefaultArgs;
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
 *         var example = new EncryptionByDefault(&#34;example&#34;, EncryptionByDefaultArgs.builder()        
 *             .enabled(true)
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
 *  to = aws_ebs_encryption_by_default.example
 * 
 *  id = &#34;default&#34; } Using `pulumi import`, import the default EBS encryption state. For exampleconsole % pulumi import aws_ebs_encryption_by_default.example default
 * 
 */
@ResourceType(type="aws:ebs/encryptionByDefault:EncryptionByDefault")
public class EncryptionByDefault extends com.pulumi.resources.CustomResource {
    /**
     * Whether or not default EBS encryption is enabled. Valid values are `true` or `false`. Defaults to `true`.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Whether or not default EBS encryption is enabled. Valid values are `true` or `false`. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EncryptionByDefault(String name) {
        this(name, EncryptionByDefaultArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EncryptionByDefault(String name, @Nullable EncryptionByDefaultArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EncryptionByDefault(String name, @Nullable EncryptionByDefaultArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ebs/encryptionByDefault:EncryptionByDefault", name, args == null ? EncryptionByDefaultArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EncryptionByDefault(String name, Output<String> id, @Nullable EncryptionByDefaultState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ebs/encryptionByDefault:EncryptionByDefault", name, state, makeResourceOptions(options, id));
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
    public static EncryptionByDefault get(String name, Output<String> id, @Nullable EncryptionByDefaultState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EncryptionByDefault(name, id, state, options);
    }
}
