// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2.KeyPairArgs;
import com.pulumi.aws.ec2.inputs.KeyPairState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an [EC2 key pair](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) resource. A key pair is used to control login access to EC2 instances.
 * 
 * Currently this resource requires an existing user-supplied key pair. This key pair&#39;s public key will be registered with AWS to allow logging-in to EC2 instances.
 * 
 * When importing an existing key pair the public key material may be in any format supported by AWS. Supported formats (per the [AWS documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html#how-to-generate-your-own-key-and-import-it-to-aws)) are:
 * 
 * * OpenSSH public key format (the format in ~/.ssh/authorized_keys)
 * * Base64 encoded DER format
 * * SSH public key file format as specified in RFC4716
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.KeyPair;
 * import com.pulumi.aws.ec2.KeyPairArgs;
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
 *         var deployer = new KeyPair(&#34;deployer&#34;, KeyPairArgs.builder()        
 *             .publicKey(&#34;ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQD3F6tyPEFEzV0LX3X8BsXdMsQz1x2cEikKDEY0aIj41qgxMCP/iteneqXSIFZBp5vizPvaoIR3Um9xK7PGoW8giupGn+EPuxIA4cDM4vzOqOkiMPhz5XK0whEjkVzTo4+S0puvDZuwIsdiW9mxhJc7tgBNL0cYlWSYVkz4G/fslNfRPW5mYAM49f4fhtxPb5ok4Q2Lg9dPKVHO/Bgeu5woMc7RY0p1ej6D4CKFE6lymSDJpW0YHX/wqE9+cfEauh7xZcG0q9t2ta6F6fmX0agvpFyZo8aFbXeUBr7osSCJNgvavWbM/06niWrOvYX2xwWdhXmXSrbX8ZbabVohBK41 email@example.com&#34;)
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
 *  to = aws_key_pair.deployer
 * 
 *  id = &#34;deployer-key&#34; } Using `pulumi import`, import Key Pairs using the `key_name`. For exampleconsole % pulumi import aws_key_pair.deployer deployer-key
 * 
 */
@ResourceType(type="aws:ec2/keyPair:KeyPair")
public class KeyPair extends com.pulumi.resources.CustomResource {
    /**
     * The key pair ARN.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The key pair ARN.
     * 
     */
    public Output<String> arn() {
        return this.arn;
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
     * The name for the key pair. If neither `key_name` nor `key_name_prefix` is provided, the provider will create a unique key name.
     * 
     */
    @Export(name="keyName", refs={String.class}, tree="[0]")
    private Output<String> keyName;

    /**
     * @return The name for the key pair. If neither `key_name` nor `key_name_prefix` is provided, the provider will create a unique key name.
     * 
     */
    public Output<String> keyName() {
        return this.keyName;
    }
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `key_name`. If neither `key_name` nor `key_name_prefix` is provided, the provider will create a unique key name.
     * 
     */
    @Export(name="keyNamePrefix", refs={String.class}, tree="[0]")
    private Output<String> keyNamePrefix;

    /**
     * @return Creates a unique name beginning with the specified prefix. Conflicts with `key_name`. If neither `key_name` nor `key_name_prefix` is provided, the provider will create a unique key name.
     * 
     */
    public Output<String> keyNamePrefix() {
        return this.keyNamePrefix;
    }
    /**
     * The key pair ID.
     * 
     */
    @Export(name="keyPairId", refs={String.class}, tree="[0]")
    private Output<String> keyPairId;

    /**
     * @return The key pair ID.
     * 
     */
    public Output<String> keyPairId() {
        return this.keyPairId;
    }
    /**
     * The type of key pair.
     * 
     */
    @Export(name="keyType", refs={String.class}, tree="[0]")
    private Output<String> keyType;

    /**
     * @return The type of key pair.
     * 
     */
    public Output<String> keyType() {
        return this.keyType;
    }
    /**
     * The public key material.
     * 
     */
    @Export(name="publicKey", refs={String.class}, tree="[0]")
    private Output<String> publicKey;

    /**
     * @return The public key material.
     * 
     */
    public Output<String> publicKey() {
        return this.publicKey;
    }
    /**
     * Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
    public KeyPair(String name) {
        this(name, KeyPairArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public KeyPair(String name, KeyPairArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public KeyPair(String name, KeyPairArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/keyPair:KeyPair", name, args == null ? KeyPairArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private KeyPair(String name, Output<String> id, @Nullable KeyPairState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/keyPair:KeyPair", name, state, makeResourceOptions(options, id));
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
