// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.codestarconnections;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.codestarconnections.ConnectionArgs;
import com.pulumi.aws.codestarconnections.inputs.ConnectionState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a CodeStar Connection.
 * 
 * &gt; **NOTE:** The `aws.codestarconnections.Connection` resource is created in the state `PENDING`. Authentication with the connection provider must be completed in the AWS Console.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.codestarconnections.Connection;
 * import com.pulumi.aws.codestarconnections.ConnectionArgs;
 * import com.pulumi.aws.codepipeline.Pipeline;
 * import com.pulumi.aws.codepipeline.PipelineArgs;
 * import com.pulumi.aws.codepipeline.inputs.PipelineArtifactStoreArgs;
 * import com.pulumi.aws.codepipeline.inputs.PipelineStageArgs;
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
 *         var exampleConnection = new Connection(&#34;exampleConnection&#34;, ConnectionArgs.builder()        
 *             .providerType(&#34;Bitbucket&#34;)
 *             .build());
 * 
 *         var examplePipeline = new Pipeline(&#34;examplePipeline&#34;, PipelineArgs.builder()        
 *             .roleArn(aws_iam_role.codepipeline_role().arn())
 *             .artifactStores()
 *             .stages(            
 *                 PipelineStageArgs.builder()
 *                     .name(&#34;Source&#34;)
 *                     .actions(PipelineStageActionArgs.builder()
 *                         .name(&#34;Source&#34;)
 *                         .category(&#34;Source&#34;)
 *                         .owner(&#34;AWS&#34;)
 *                         .provider(&#34;CodeStarSourceConnection&#34;)
 *                         .version(&#34;1&#34;)
 *                         .outputArtifacts(&#34;source_output&#34;)
 *                         .configuration(Map.ofEntries(
 *                             Map.entry(&#34;ConnectionArn&#34;, exampleConnection.arn()),
 *                             Map.entry(&#34;FullRepositoryId&#34;, &#34;my-organization/test&#34;),
 *                             Map.entry(&#34;BranchName&#34;, &#34;main&#34;)
 *                         ))
 *                         .build())
 *                     .build(),
 *                 PipelineStageArgs.builder()
 *                     .name(&#34;Build&#34;)
 *                     .actions()
 *                     .build(),
 *                 PipelineStageArgs.builder()
 *                     .name(&#34;Deploy&#34;)
 *                     .actions()
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * CodeStar connections can be imported using the ARN, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:codestarconnections/connection:Connection test-connection arn:aws:codestar-connections:us-west-1:0123456789:connection/79d4d357-a2ee-41e4-b350-2fe39ae59448
 * ```
 * 
 */
@ResourceType(type="aws:codestarconnections/connection:Connection")
public class Connection extends com.pulumi.resources.CustomResource {
    /**
     * The codestar connection ARN.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The codestar connection ARN.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The codestar connection status. Possible values are `PENDING`, `AVAILABLE` and `ERROR`.
     * 
     */
    @Export(name="connectionStatus", refs={String.class}, tree="[0]")
    private Output<String> connectionStatus;

    /**
     * @return The codestar connection status. Possible values are `PENDING`, `AVAILABLE` and `ERROR`.
     * 
     */
    public Output<String> connectionStatus() {
        return this.connectionStatus;
    }
    /**
     * The Amazon Resource Name (ARN) of the host associated with the connection. Conflicts with `provider_type`
     * 
     */
    @Export(name="hostArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> hostArn;

    /**
     * @return The Amazon Resource Name (ARN) of the host associated with the connection. Conflicts with `provider_type`
     * 
     */
    public Output<Optional<String>> hostArn() {
        return Codegen.optional(this.hostArn);
    }
    /**
     * The name of the connection to be created. The name must be unique in the calling AWS account. Changing `name` will create a new resource.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the connection to be created. The name must be unique in the calling AWS account. Changing `name` will create a new resource.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The name of the external provider where your third-party code repository is configured. Valid values are `Bitbucket`, `GitHub` or `GitHubEnterpriseServer`. Changing `provider_type` will create a new resource. Conflicts with `host_arn`
     * 
     */
    @Export(name="providerType", refs={String.class}, tree="[0]")
    private Output<String> providerType;

    /**
     * @return The name of the external provider where your third-party code repository is configured. Valid values are `Bitbucket`, `GitHub` or `GitHubEnterpriseServer`. Changing `provider_type` will create a new resource. Conflicts with `host_arn`
     * 
     */
    public Output<String> providerType() {
        return this.providerType;
    }
    /**
     * Map of key-value resource tags to associate with the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Map of key-value resource tags to associate with the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
    public Connection(String name) {
        this(name, ConnectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Connection(String name, @Nullable ConnectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Connection(String name, @Nullable ConnectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:codestarconnections/connection:Connection", name, args == null ? ConnectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Connection(String name, Output<String> id, @Nullable ConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:codestarconnections/connection:Connection", name, state, makeResourceOptions(options, id));
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
