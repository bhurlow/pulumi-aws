// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.memorydb;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.memorydb.UserArgs;
import com.pulumi.aws.memorydb.inputs.UserState;
import com.pulumi.aws.memorydb.outputs.UserAuthenticationMode;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a MemoryDB User.
 * 
 * More information about users and ACL-s can be found in the [MemoryDB User Guide](https://docs.aws.amazon.com/memorydb/latest/devguide/clusters.acls.html).
 * 
 * &gt; **Note:** All arguments including the username and passwords will be stored in the raw state as plain-text.
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.random.RandomPassword;
 * import com.pulumi.random.RandomPasswordArgs;
 * import com.pulumi.aws.memorydb.User;
 * import com.pulumi.aws.memorydb.UserArgs;
 * import com.pulumi.aws.memorydb.inputs.UserAuthenticationModeArgs;
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
 *         var exampleRandomPassword = new RandomPassword(&#34;exampleRandomPassword&#34;, RandomPasswordArgs.builder()        
 *             .length(16)
 *             .build());
 * 
 *         var exampleUser = new User(&#34;exampleUser&#34;, UserArgs.builder()        
 *             .userName(&#34;my-user&#34;)
 *             .accessString(&#34;on ~* &amp;* +@all&#34;)
 *             .authenticationMode(UserAuthenticationModeArgs.builder()
 *                 .type(&#34;password&#34;)
 *                 .passwords(exampleRandomPassword.result())
 *                 .build())
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
 *  to = aws_memorydb_user.example
 * 
 *  id = &#34;my-user&#34; } Using `pulumi import`, import a user using the `user_name`. For exampleconsole % pulumi import aws_memorydb_user.example my-user The `passwords` are not available for imported resources, as this information cannot be read back from the MemoryDB API.
 * 
 */
@ResourceType(type="aws:memorydb/user:User")
public class User extends com.pulumi.resources.CustomResource {
    /**
     * The access permissions string used for this user.
     * 
     */
    @Export(name="accessString", refs={String.class}, tree="[0]")
    private Output<String> accessString;

    /**
     * @return The access permissions string used for this user.
     * 
     */
    public Output<String> accessString() {
        return this.accessString;
    }
    /**
     * The ARN of the user.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the user.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Denotes the user&#39;s authentication properties. Detailed below.
     * 
     */
    @Export(name="authenticationMode", refs={UserAuthenticationMode.class}, tree="[0]")
    private Output<UserAuthenticationMode> authenticationMode;

    /**
     * @return Denotes the user&#39;s authentication properties. Detailed below.
     * 
     */
    public Output<UserAuthenticationMode> authenticationMode() {
        return this.authenticationMode;
    }
    /**
     * The minimum engine version supported for the user.
     * 
     */
    @Export(name="minimumEngineVersion", refs={String.class}, tree="[0]")
    private Output<String> minimumEngineVersion;

    /**
     * @return The minimum engine version supported for the user.
     * 
     */
    public Output<String> minimumEngineVersion() {
        return this.minimumEngineVersion;
    }
    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * Name of the MemoryDB user. Up to 40 characters.
     * 
     * The following arguments are optional:
     * 
     */
    @Export(name="userName", refs={String.class}, tree="[0]")
    private Output<String> userName;

    /**
     * @return Name of the MemoryDB user. Up to 40 characters.
     * 
     * The following arguments are optional:
     * 
     */
    public Output<String> userName() {
        return this.userName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public User(String name) {
        this(name, UserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public User(String name, UserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public User(String name, UserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:memorydb/user:User", name, args == null ? UserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private User(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:memorydb/user:User", name, state, makeResourceOptions(options, id));
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
    public static User get(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new User(name, id, state, options);
    }
}
