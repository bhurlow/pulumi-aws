// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.lightsail;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.lightsail.KeyPairArgs;
import com.pulumi.aws.lightsail.inputs.KeyPairState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Lightsail Key Pair, for use with Lightsail Instances. These key pairs
 * are separate from EC2 Key Pairs, and must be created or imported for use with
 * Lightsail.
 * 
 * &gt; **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see [&#34;Regions and Availability Zones in Amazon Lightsail&#34;](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details
 * 
 * ## Example Usage
 * ### Create New Key Pair
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.lightsail.KeyPair;
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
 *         var lgKeyPair = new KeyPair(&#34;lgKeyPair&#34;);
 * 
 *     }
 * }
 * ```
 * ### Create New Key Pair with PGP Encrypted Private Key
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.lightsail.KeyPair;
 * import com.pulumi.aws.lightsail.KeyPairArgs;
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
 *         var lgKeyPair = new KeyPair(&#34;lgKeyPair&#34;, KeyPairArgs.builder()        
 *             .pgpKey(&#34;keybase:keybaseusername&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Existing Public Key Import
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.lightsail.KeyPair;
 * import com.pulumi.aws.lightsail.KeyPairArgs;
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
 *         var lgKeyPair = new KeyPair(&#34;lgKeyPair&#34;, KeyPairArgs.builder()        
 *             .publicKey(Files.readString(Paths.get(&#34;~/.ssh/id_rsa.pub&#34;)))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * You cannot import Lightsail Key Pairs because the private and public key are only available on initial creation.
 * 
 */
@ResourceType(type="aws:lightsail/keyPair:KeyPair")
public class KeyPair extends com.pulumi.resources.CustomResource {
    /**
     * The ARN of the Lightsail key pair.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the Lightsail key pair.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The MD5 public key fingerprint for the encrypted private key.
     * 
     */
    @Export(name="encryptedFingerprint", refs={String.class}, tree="[0]")
    private Output<String> encryptedFingerprint;

    /**
     * @return The MD5 public key fingerprint for the encrypted private key.
     * 
     */
    public Output<String> encryptedFingerprint() {
        return this.encryptedFingerprint;
    }
    /**
     * the private key material, base 64 encoded and encrypted with the given `pgp_key`. This is only populated when creating a new key and `pgp_key` is supplied.
     * 
     */
    @Export(name="encryptedPrivateKey", refs={String.class}, tree="[0]")
    private Output<String> encryptedPrivateKey;

    /**
     * @return the private key material, base 64 encoded and encrypted with the given `pgp_key`. This is only populated when creating a new key and `pgp_key` is supplied.
     * 
     */
    public Output<String> encryptedPrivateKey() {
        return this.encryptedPrivateKey;
    }
    /**
     * The MD5 public key fingerprint as specified in section 4 of RFC 4716.
     * 
     */
    @Export(name="fingerprint", refs={String.class}, tree="[0]")
    private Output<String> fingerprint;

    /**
     * @return The MD5 public key fingerprint as specified in section 4 of RFC 4716.
     * 
     */
    public Output<String> fingerprint() {
        return this.fingerprint;
    }
    /**
     * The name of the Lightsail Key Pair. If omitted, a unique name will be generated by this provider
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the Lightsail Key Pair. If omitted, a unique name will be generated by this provider
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="namePrefix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namePrefix;

    public Output<Optional<String>> namePrefix() {
        return Codegen.optional(this.namePrefix);
    }
    /**
     * An optional PGP key to encrypt the resulting private key material. Only used when creating a new key pair
     * 
     */
    @Export(name="pgpKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pgpKey;

    /**
     * @return An optional PGP key to encrypt the resulting private key material. Only used when creating a new key pair
     * 
     */
    public Output<Optional<String>> pgpKey() {
        return Codegen.optional(this.pgpKey);
    }
    /**
     * the private key, base64 encoded. This is only populated when creating a new key, and when no `pgp_key` is provided.
     * 
     */
    @Export(name="privateKey", refs={String.class}, tree="[0]")
    private Output<String> privateKey;

    /**
     * @return the private key, base64 encoded. This is only populated when creating a new key, and when no `pgp_key` is provided.
     * 
     */
    public Output<String> privateKey() {
        return this.privateKey;
    }
    /**
     * The public key material. This public key will be imported into Lightsail
     * 
     */
    @Export(name="publicKey", refs={String.class}, tree="[0]")
    private Output<String> publicKey;

    /**
     * @return The public key material. This public key will be imported into Lightsail
     * 
     */
    public Output<String> publicKey() {
        return this.publicKey;
    }
    /**
     * A map of tags to assign to the collection. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     * &gt; **NOTE:** a PGP key is not required, however it is strongly encouraged. Without a PGP key, the private key material will be stored in state unencrypted.`pgp_key` is ignored if `public_key` is supplied.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the collection. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     * &gt; **NOTE:** a PGP key is not required, however it is strongly encouraged. Without a PGP key, the private key material will be stored in state unencrypted.`pgp_key` is ignored if `public_key` is supplied.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public KeyPair(String name) {
        this(name, KeyPairArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public KeyPair(String name, @Nullable KeyPairArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public KeyPair(String name, @Nullable KeyPairArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:lightsail/keyPair:KeyPair", name, args == null ? KeyPairArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private KeyPair(String name, Output<String> id, @Nullable KeyPairState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:lightsail/keyPair:KeyPair", name, state, makeResourceOptions(options, id));
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
    public static KeyPair get(String name, Output<String> id, @Nullable KeyPairState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new KeyPair(name, id, state, options);
    }
}
