// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.rolesanywhere;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.rolesanywhere.ProfileArgs;
import com.pulumi.aws.rolesanywhere.inputs.ProfileState;
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
 * Resource for managing a Roles Anywhere Profile.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.iam.Role;
 * import com.pulumi.aws.iam.RoleArgs;
 * import com.pulumi.aws.rolesanywhere.Profile;
 * import com.pulumi.aws.rolesanywhere.ProfileArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
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
 *         var testRole = new Role(&#34;testRole&#34;, RoleArgs.builder()        
 *             .path(&#34;/&#34;)
 *             .assumeRolePolicy(serializeJson(
 *                 jsonObject(
 *                     jsonProperty(&#34;Version&#34;, &#34;2012-10-17&#34;),
 *                     jsonProperty(&#34;Statement&#34;, jsonArray(jsonObject(
 *                         jsonProperty(&#34;Action&#34;, jsonArray(
 *                             &#34;sts:AssumeRole&#34;, 
 *                             &#34;sts:TagSession&#34;, 
 *                             &#34;sts:SetSourceIdentity&#34;
 *                         )),
 *                         jsonProperty(&#34;Principal&#34;, jsonObject(
 *                             jsonProperty(&#34;Service&#34;, &#34;rolesanywhere.amazonaws.com&#34;)
 *                         )),
 *                         jsonProperty(&#34;Effect&#34;, &#34;Allow&#34;),
 *                         jsonProperty(&#34;Sid&#34;, &#34;&#34;)
 *                     )))
 *                 )))
 *             .build());
 * 
 *         var testProfile = new Profile(&#34;testProfile&#34;, ProfileArgs.builder()        
 *             .roleArns(testRole.arn())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * `aws_rolesanywhere_profile` can be imported using its `id`, e.g.
 * 
 * ```sh
 *  $ pulumi import aws:rolesanywhere/profile:Profile example db138a85-8925-4f9f-a409-08231233cacf
 * ```
 * 
 */
@ResourceType(type="aws:rolesanywhere/profile:Profile")
public class Profile extends com.pulumi.resources.CustomResource {
    /**
     * Amazon Resource Name (ARN) of the Profile
     * 
     */
    @Export(name="arn", type=String.class, parameters={})
    private Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of the Profile
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The number of seconds the vended session credentials are valid for. Defaults to 3600.
     * 
     */
    @Export(name="durationSeconds", type=Integer.class, parameters={})
    private Output<Integer> durationSeconds;

    /**
     * @return The number of seconds the vended session credentials are valid for. Defaults to 3600.
     * 
     */
    public Output<Integer> durationSeconds() {
        return this.durationSeconds;
    }
    /**
     * Whether or not the Profile is enabled.
     * 
     */
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Whether or not the Profile is enabled.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * A list of managed policy ARNs that apply to the vended session credentials.
     * 
     */
    @Export(name="managedPolicyArns", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> managedPolicyArns;

    /**
     * @return A list of managed policy ARNs that apply to the vended session credentials.
     * 
     */
    public Output<Optional<List<String>>> managedPolicyArns() {
        return Codegen.optional(this.managedPolicyArns);
    }
    /**
     * The name of the Profile.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the Profile.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Specifies whether instance properties are required in [CreateSession](https://docs.aws.amazon.com/rolesanywhere/latest/APIReference/API_CreateSession.html) requests with this profile.
     * 
     */
    @Export(name="requireInstanceProperties", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> requireInstanceProperties;

    /**
     * @return Specifies whether instance properties are required in [CreateSession](https://docs.aws.amazon.com/rolesanywhere/latest/APIReference/API_CreateSession.html) requests with this profile.
     * 
     */
    public Output<Optional<Boolean>> requireInstanceProperties() {
        return Codegen.optional(this.requireInstanceProperties);
    }
    /**
     * A list of IAM roles that this profile can assume
     * 
     */
    @Export(name="roleArns", type=List.class, parameters={String.class})
    private Output<List<String>> roleArns;

    /**
     * @return A list of IAM roles that this profile can assume
     * 
     */
    public Output<List<String>> roleArns() {
        return this.roleArns;
    }
    /**
     * A session policy that applies to the trust boundary of the vended session credentials.
     * 
     */
    @Export(name="sessionPolicy", type=String.class, parameters={})
    private Output</* @Nullable */ String> sessionPolicy;

    /**
     * @return A session policy that applies to the trust boundary of the vended session credentials.
     * 
     */
    public Output<Optional<String>> sessionPolicy() {
        return Codegen.optional(this.sessionPolicy);
    }
    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, String.class})
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
    @Export(name="tagsAll", type=Map.class, parameters={String.class, String.class})
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
    public Profile(String name) {
        this(name, ProfileArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Profile(String name, ProfileArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Profile(String name, ProfileArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:rolesanywhere/profile:Profile", name, args == null ? ProfileArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Profile(String name, Output<String> id, @Nullable ProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:rolesanywhere/profile:Profile", name, state, makeResourceOptions(options, id));
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
    public static Profile get(String name, Output<String> id, @Nullable ProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Profile(name, id, state, options);
    }
}
