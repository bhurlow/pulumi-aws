// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.apprunner;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.apprunner.ConnectionArgs;
import com.pulumi.aws.apprunner.inputs.ConnectionState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an App Runner Connection.
 * 
 * &gt; **NOTE:** After creation, you must complete the authentication handshake using the App Runner console.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.apprunner.Connection;
 * import com.pulumi.aws.apprunner.ConnectionArgs;
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
 *         var example = new Connection(&#34;example&#34;, ConnectionArgs.builder()        
 *             .connectionName(&#34;example&#34;)
 *             .providerType(&#34;GITHUB&#34;)
 *             .tags(Map.of(&#34;Name&#34;, &#34;example-apprunner-connection&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * App Runner Connections can be imported by using the `connection_name`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:apprunner/connection:Connection example example
 * ```
 * 
 */
@ResourceType(type="aws:apprunner/connection:Connection")
public class Connection extends com.pulumi.resources.CustomResource {
    /**
     * ARN of the connection.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return ARN of the connection.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Name of the connection.
     * 
     */
    @Export(name="connectionName", refs={String.class}, tree="[0]")
    private Output<String> connectionName;

    /**
     * @return Name of the connection.
     * 
     */
    public Output<String> connectionName() {
        return this.connectionName;
    }
    /**
     * Source repository provider. Valid values: `GITHUB`.
     * 
     */
    @Export(name="providerType", refs={String.class}, tree="[0]")
    private Output<String> providerType;

    /**
     * @return Source repository provider. Valid values: `GITHUB`.
     * 
     */
    public Output<String> providerType() {
        return this.providerType;
    }
    /**
     * Current state of the App Runner connection. When the state is `AVAILABLE`, you can use the connection to create an `aws.apprunner.Service` resource.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Current state of the App Runner connection. When the state is `AVAILABLE`, you can use the connection to create an `aws.apprunner.Service` resource.
     * 
     */
    public Output<String> status() {
        return this.status;
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
     * Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    /**
     * @return Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Connection(String name) {
        this(name, ConnectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Connection(String name, ConnectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Connection(String name, ConnectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:apprunner/connection:Connection", name, args == null ? ConnectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Connection(String name, Output<String> id, @Nullable ConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:apprunner/connection:Connection", name, state, makeResourceOptions(options, id));
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
    public static Connection get(String name, Output<String> id, @Nullable ConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Connection(name, id, state, options);
    }
}
