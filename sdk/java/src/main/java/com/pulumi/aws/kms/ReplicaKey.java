// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kms;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.kms.ReplicaKeyArgs;
import com.pulumi.aws.kms.inputs.ReplicaKeyState;
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
 * Manages a KMS multi-Region replica key.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.Provider;
 * import com.pulumi.aws.ProviderArgs;
 * import com.pulumi.aws.kms.Key;
 * import com.pulumi.aws.kms.KeyArgs;
 * import com.pulumi.aws.kms.ReplicaKey;
 * import com.pulumi.aws.kms.ReplicaKeyArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var primary = new Provider(&#34;primary&#34;, ProviderArgs.builder()        
 *             .region(&#34;us-east-1&#34;)
 *             .build());
 * 
 *         var primaryKey = new Key(&#34;primaryKey&#34;, KeyArgs.builder()        
 *             .description(&#34;Multi-Region primary key&#34;)
 *             .deletionWindowInDays(30)
 *             .multiRegion(true)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(aws.primary())
 *                 .build());
 * 
 *         var replica = new ReplicaKey(&#34;replica&#34;, ReplicaKeyArgs.builder()        
 *             .description(&#34;Multi-Region replica key&#34;)
 *             .deletionWindowInDays(7)
 *             .primaryKeyArn(primaryKey.arn())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import KMS multi-Region replica keys using the `id`. For example:
 * 
 * ```sh
 *  $ pulumi import aws:kms/replicaKey:ReplicaKey example 1234abcd-12ab-34cd-56ef-1234567890ab
 * ```
 * 
 */
@ResourceType(type="aws:kms/replicaKey:ReplicaKey")
public class ReplicaKey extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name (ARN) of the replica key. The key ARNs of related multi-Region keys differ only in the Region value.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) of the replica key. The key ARNs of related multi-Region keys differ only in the Region value.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * A flag to indicate whether to bypass the key policy lockout safety check.
     * Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
     * For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
     * The default value is `false`.
     * 
     */
    @Export(name="bypassPolicyLockoutSafetyCheck", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> bypassPolicyLockoutSafetyCheck;

    /**
     * @return A flag to indicate whether to bypass the key policy lockout safety check.
     * Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
     * For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
     * The default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> bypassPolicyLockoutSafetyCheck() {
        return Codegen.optional(this.bypassPolicyLockoutSafetyCheck);
    }
    /**
     * The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
     * If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
     * 
     */
    @Export(name="deletionWindowInDays", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> deletionWindowInDays;

    /**
     * @return The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
     * If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
     * 
     */
    public Output<Optional<Integer>> deletionWindowInDays() {
        return Codegen.optional(this.deletionWindowInDays);
    }
    /**
     * A description of the KMS key.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the KMS key.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. The default value is `true`.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. The default value is `true`.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * The key ID of the replica key. Related multi-Region keys have the same key ID.
     * 
     */
    @Export(name="keyId", refs={String.class}, tree="[0]")
    private Output<String> keyId;

    /**
     * @return The key ID of the replica key. Related multi-Region keys have the same key ID.
     * 
     */
    public Output<String> keyId() {
        return this.keyId;
    }
    /**
     * A Boolean value that specifies whether key rotation is enabled. This is a shared property of multi-Region keys.
     * 
     */
    @Export(name="keyRotationEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> keyRotationEnabled;

    /**
     * @return A Boolean value that specifies whether key rotation is enabled. This is a shared property of multi-Region keys.
     * 
     */
    public Output<Boolean> keyRotationEnabled() {
        return this.keyRotationEnabled;
    }
    /**
     * The type of key material in the KMS key. This is a shared property of multi-Region keys.
     * 
     */
    @Export(name="keySpec", refs={String.class}, tree="[0]")
    private Output<String> keySpec;

    /**
     * @return The type of key material in the KMS key. This is a shared property of multi-Region keys.
     * 
     */
    public Output<String> keySpec() {
        return this.keySpec;
    }
    /**
     * The [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the KMS key. This is a shared property of multi-Region keys.
     * 
     */
    @Export(name="keyUsage", refs={String.class}, tree="[0]")
    private Output<String> keyUsage;

    /**
     * @return The [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the KMS key. This is a shared property of multi-Region keys.
     * 
     */
    public Output<String> keyUsage() {
        return this.keyUsage;
    }
    /**
     * The key policy to attach to the KMS key. If you do not specify a key policy, AWS KMS attaches the [default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default) to the KMS key.
     * 
     */
    @Export(name="policy", refs={String.class}, tree="[0]")
    private Output<String> policy;

    /**
     * @return The key policy to attach to the KMS key. If you do not specify a key policy, AWS KMS attaches the [default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default) to the KMS key.
     * 
     */
    public Output<String> policy() {
        return this.policy;
    }
    /**
     * The ARN of the multi-Region primary key to replicate. The primary key must be in a different AWS Region of the same AWS Partition. You can create only one replica of a given primary key in each AWS Region.
     * 
     */
    @Export(name="primaryKeyArn", refs={String.class}, tree="[0]")
    private Output<String> primaryKeyArn;

    /**
     * @return The ARN of the multi-Region primary key to replicate. The primary key must be in a different AWS Region of the same AWS Partition. You can create only one replica of a given primary key in each AWS Region.
     * 
     */
    public Output<String> primaryKeyArn() {
        return this.primaryKeyArn;
    }
    /**
     * A map of tags to assign to the replica key. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the replica key. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ReplicaKey(String name) {
        this(name, ReplicaKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ReplicaKey(String name, ReplicaKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ReplicaKey(String name, ReplicaKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:kms/replicaKey:ReplicaKey", name, args == null ? ReplicaKeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ReplicaKey(String name, Output<String> id, @Nullable ReplicaKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:kms/replicaKey:ReplicaKey", name, state, makeResourceOptions(options, id));
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
    public static ReplicaKey get(String name, Output<String> id, @Nullable ReplicaKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ReplicaKey(name, id, state, options);
    }
}
