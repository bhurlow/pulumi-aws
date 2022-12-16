// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.iam;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.iam.OpenIdConnectProviderArgs;
import com.pulumi.aws.iam.inputs.OpenIdConnectProviderState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an IAM OpenID Connect provider.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.iam.OpenIdConnectProvider;
 * import com.pulumi.aws.iam.OpenIdConnectProviderArgs;
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
 *         var default_ = new OpenIdConnectProvider(&#34;default&#34;, OpenIdConnectProviderArgs.builder()        
 *             .clientIdLists(&#34;266362248691-342342xasdasdasda-apps.googleusercontent.com&#34;)
 *             .thumbprintLists()
 *             .url(&#34;https://accounts.google.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * IAM OpenID Connect Providers can be imported using the `arn`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:iam/openIdConnectProvider:OpenIdConnectProvider default arn:aws:iam::123456789012:oidc-provider/accounts.google.com
 * ```
 * 
 */
@ResourceType(type="aws:iam/openIdConnectProvider:OpenIdConnectProvider")
public class OpenIdConnectProvider extends com.pulumi.resources.CustomResource {
    /**
     * The ARN assigned by AWS for this provider.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN assigned by AWS for this provider.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that&#39;s sent as the client_id parameter on OAuth requests.)
     * 
     */
    @Export(name="clientIdLists", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> clientIdLists;

    /**
     * @return A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that&#39;s sent as the client_id parameter on OAuth requests.)
     * 
     */
    public Output<List<String>> clientIdLists() {
        return this.clientIdLists;
    }
    /**
     * Map of resource tags for the IAM OIDC provider. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Map of resource tags for the IAM OIDC provider. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider&#39;s server certificate(s).
     * 
     */
    @Export(name="thumbprintLists", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> thumbprintLists;

    /**
     * @return A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider&#39;s server certificate(s).
     * 
     */
    public Output<List<String>> thumbprintLists() {
        return this.thumbprintLists;
    }
    /**
     * The URL of the identity provider. Corresponds to the _iss_ claim.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The URL of the identity provider. Corresponds to the _iss_ claim.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public OpenIdConnectProvider(String name) {
        this(name, OpenIdConnectProviderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OpenIdConnectProvider(String name, OpenIdConnectProviderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OpenIdConnectProvider(String name, OpenIdConnectProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:iam/openIdConnectProvider:OpenIdConnectProvider", name, args == null ? OpenIdConnectProviderArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private OpenIdConnectProvider(String name, Output<String> id, @Nullable OpenIdConnectProviderState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:iam/openIdConnectProvider:OpenIdConnectProvider", name, state, makeResourceOptions(options, id));
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
    public static OpenIdConnectProvider get(String name, Output<String> id, @Nullable OpenIdConnectProviderState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OpenIdConnectProvider(name, id, state, options);
    }
}
