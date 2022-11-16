// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.elasticache;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.elasticache.UserArgs;
import com.pulumi.aws.elasticache.inputs.UserState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an ElastiCache user resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.elasticache.User;
 * import com.pulumi.aws.elasticache.UserArgs;
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
 *         var test = new User(&#34;test&#34;, UserArgs.builder()        
 *             .accessString(&#34;on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember&#34;)
 *             .engine(&#34;REDIS&#34;)
 *             .passwords(&#34;password123456789&#34;)
 *             .userId(&#34;testUserId&#34;)
 *             .userName(&#34;testUserName&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ElastiCache users can be imported using the `user_id`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:elasticache/user:User my_user userId1
 * ```
 * 
 */
@ResourceType(type="aws:elasticache/user:User")
public class User extends com.pulumi.resources.CustomResource {
    /**
     * Access permissions string used for this user. See [Specifying Permissions Using an Access String](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.RBAC.html#Access-string) for more details.
     * 
     */
    @Export(name="accessString", type=String.class, parameters={})
    private Output<String> accessString;

    /**
     * @return Access permissions string used for this user. See [Specifying Permissions Using an Access String](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.RBAC.html#Access-string) for more details.
     * 
     */
    public Output<String> accessString() {
        return this.accessString;
    }
    /**
     * The ARN of the created ElastiCache User.
     * 
     */
    @Export(name="arn", type=String.class, parameters={})
    private Output<String> arn;

    /**
     * @return The ARN of the created ElastiCache User.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The current supported value is `REDIS`.
     * 
     */
    @Export(name="engine", type=String.class, parameters={})
    private Output<String> engine;

    /**
     * @return The current supported value is `REDIS`.
     * 
     */
    public Output<String> engine() {
        return this.engine;
    }
    /**
     * Indicates a password is not required for this user.
     * 
     */
    @Export(name="noPasswordRequired", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> noPasswordRequired;

    /**
     * @return Indicates a password is not required for this user.
     * 
     */
    public Output<Optional<Boolean>> noPasswordRequired() {
        return Codegen.optional(this.noPasswordRequired);
    }
    /**
     * Passwords used for this user. You can create up to two passwords for each user.
     * 
     */
    @Export(name="passwords", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> passwords;

    /**
     * @return Passwords used for this user. You can create up to two passwords for each user.
     * 
     */
    public Output<Optional<List<String>>> passwords() {
        return Codegen.optional(this.passwords);
    }
    /**
     * A list of tags to be added to this resource. A tag is a key-value pair.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A list of tags to be added to this resource. A tag is a key-value pair.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    @Export(name="tagsAll", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> tagsAll;

    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    /**
     * The ID of the user.
     * 
     */
    @Export(name="userId", type=String.class, parameters={})
    private Output<String> userId;

    /**
     * @return The ID of the user.
     * 
     */
    public Output<String> userId() {
        return this.userId;
    }
    /**
     * The username of the user.
     * 
     */
    @Export(name="userName", type=String.class, parameters={})
    private Output<String> userName;

    /**
     * @return The username of the user.
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
        super("aws:elasticache/user:User", name, args == null ? UserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private User(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:elasticache/user:User", name, state, makeResourceOptions(options, id));
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
