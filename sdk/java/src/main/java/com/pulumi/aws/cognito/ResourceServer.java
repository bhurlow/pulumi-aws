// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cognito;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.cognito.ResourceServerArgs;
import com.pulumi.aws.cognito.inputs.ResourceServerState;
import com.pulumi.aws.cognito.outputs.ResourceServerScope;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Cognito Resource Server.
 * 
 * ## Example Usage
 * ### Create a basic resource server
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.cognito.UserPool;
 * import com.pulumi.aws.cognito.ResourceServer;
 * import com.pulumi.aws.cognito.ResourceServerArgs;
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
 *         var pool = new UserPool(&#34;pool&#34;);
 * 
 *         var resource = new ResourceServer(&#34;resource&#34;, ResourceServerArgs.builder()        
 *             .identifier(&#34;https://example.com&#34;)
 *             .userPoolId(pool.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Create a resource server with sample-scope
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.cognito.UserPool;
 * import com.pulumi.aws.cognito.ResourceServer;
 * import com.pulumi.aws.cognito.ResourceServerArgs;
 * import com.pulumi.aws.cognito.inputs.ResourceServerScopeArgs;
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
 *         var pool = new UserPool(&#34;pool&#34;);
 * 
 *         var resource = new ResourceServer(&#34;resource&#34;, ResourceServerArgs.builder()        
 *             .identifier(&#34;https://example.com&#34;)
 *             .scopes(ResourceServerScopeArgs.builder()
 *                 .scopeName(&#34;sample-scope&#34;)
 *                 .scopeDescription(&#34;a Sample Scope Description&#34;)
 *                 .build())
 *             .userPoolId(pool.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * `aws_cognito_resource_server` can be imported using their User Pool ID and Identifier, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:cognito/resourceServer:ResourceServer example us-west-2_abc123:https://example.com
 * ```
 * 
 */
@ResourceType(type="aws:cognito/resourceServer:ResourceServer")
public class ResourceServer extends com.pulumi.resources.CustomResource {
    /**
     * An identifier for the resource server.
     * 
     */
    @Export(name="identifier", refs={String.class}, tree="[0]")
    private Output<String> identifier;

    /**
     * @return An identifier for the resource server.
     * 
     */
    public Output<String> identifier() {
        return this.identifier;
    }
    /**
     * A name for the resource server.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A name for the resource server.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A list of all scopes configured for this resource server in the format identifier/scope_name.
     * 
     */
    @Export(name="scopeIdentifiers", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> scopeIdentifiers;

    /**
     * @return A list of all scopes configured for this resource server in the format identifier/scope_name.
     * 
     */
    public Output<List<String>> scopeIdentifiers() {
        return this.scopeIdentifiers;
    }
    /**
     * A list of Authorization Scope.
     * 
     */
    @Export(name="scopes", refs={List.class,ResourceServerScope.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ResourceServerScope>> scopes;

    /**
     * @return A list of Authorization Scope.
     * 
     */
    public Output<Optional<List<ResourceServerScope>>> scopes() {
        return Codegen.optional(this.scopes);
    }
    @Export(name="userPoolId", refs={String.class}, tree="[0]")
    private Output<String> userPoolId;

    public Output<String> userPoolId() {
        return this.userPoolId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ResourceServer(String name) {
        this(name, ResourceServerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ResourceServer(String name, ResourceServerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ResourceServer(String name, ResourceServerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cognito/resourceServer:ResourceServer", name, args == null ? ResourceServerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ResourceServer(String name, Output<String> id, @Nullable ResourceServerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cognito/resourceServer:ResourceServer", name, state, makeResourceOptions(options, id));
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
    public static ResourceServer get(String name, Output<String> id, @Nullable ResourceServerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ResourceServer(name, id, state, options);
    }
}
