// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.codeartifact;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.codeartifact.RepositoryArgs;
import com.pulumi.aws.codeartifact.inputs.RepositoryState;
import com.pulumi.aws.codeartifact.outputs.RepositoryExternalConnections;
import com.pulumi.aws.codeartifact.outputs.RepositoryUpstream;
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
 * Provides a CodeArtifact Repository Resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.kms.Key;
 * import com.pulumi.aws.kms.KeyArgs;
 * import com.pulumi.aws.codeartifact.Domain;
 * import com.pulumi.aws.codeartifact.DomainArgs;
 * import com.pulumi.aws.codeartifact.Repository;
 * import com.pulumi.aws.codeartifact.RepositoryArgs;
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
 *         var exampleKey = new Key(&#34;exampleKey&#34;, KeyArgs.builder()        
 *             .description(&#34;domain key&#34;)
 *             .build());
 * 
 *         var exampleDomain = new Domain(&#34;exampleDomain&#34;, DomainArgs.builder()        
 *             .domain(&#34;example&#34;)
 *             .encryptionKey(exampleKey.arn())
 *             .build());
 * 
 *         var test = new Repository(&#34;test&#34;, RepositoryArgs.builder()        
 *             .repository(&#34;example&#34;)
 *             .domain(exampleDomain.domain())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### With Upstream Repository
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.codeartifact.Repository;
 * import com.pulumi.aws.codeartifact.RepositoryArgs;
 * import com.pulumi.aws.codeartifact.inputs.RepositoryUpstreamArgs;
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
 *         var upstream = new Repository(&#34;upstream&#34;, RepositoryArgs.builder()        
 *             .repository(&#34;upstream&#34;)
 *             .domain(aws_codeartifact_domain.test().domain())
 *             .build());
 * 
 *         var test = new Repository(&#34;test&#34;, RepositoryArgs.builder()        
 *             .repository(&#34;example&#34;)
 *             .domain(aws_codeartifact_domain.example().domain())
 *             .upstreams(RepositoryUpstreamArgs.builder()
 *                 .repositoryName(upstream.repository())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### With External Connection
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.codeartifact.Repository;
 * import com.pulumi.aws.codeartifact.RepositoryArgs;
 * import com.pulumi.aws.codeartifact.inputs.RepositoryExternalConnectionsArgs;
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
 *         var upstream = new Repository(&#34;upstream&#34;, RepositoryArgs.builder()        
 *             .repository(&#34;upstream&#34;)
 *             .domain(aws_codeartifact_domain.test().domain())
 *             .build());
 * 
 *         var test = new Repository(&#34;test&#34;, RepositoryArgs.builder()        
 *             .repository(&#34;example&#34;)
 *             .domain(aws_codeartifact_domain.example().domain())
 *             .externalConnections(RepositoryExternalConnectionsArgs.builder()
 *                 .externalConnectionName(&#34;public:npmjs&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * CodeArtifact Repository can be imported using the CodeArtifact Repository ARN, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:codeartifact/repository:Repository example arn:aws:codeartifact:us-west-2:012345678912:repository/tf-acc-test-6968272603913957763/tf-acc-test-6968272603913957763
 * ```
 * 
 */
@ResourceType(type="aws:codeartifact/repository:Repository")
public class Repository extends com.pulumi.resources.CustomResource {
    /**
     * The account number of the AWS account that manages the repository.
     * 
     */
    @Export(name="administratorAccount", refs={String.class}, tree="[0]")
    private Output<String> administratorAccount;

    /**
     * @return The account number of the AWS account that manages the repository.
     * 
     */
    public Output<String> administratorAccount() {
        return this.administratorAccount;
    }
    /**
     * The ARN of the repository.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the repository.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The description of the repository.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the repository.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The domain that contains the created repository.
     * 
     */
    @Export(name="domain", refs={String.class}, tree="[0]")
    private Output<String> domain;

    /**
     * @return The domain that contains the created repository.
     * 
     */
    public Output<String> domain() {
        return this.domain;
    }
    /**
     * The account number of the AWS account that owns the domain.
     * 
     */
    @Export(name="domainOwner", refs={String.class}, tree="[0]")
    private Output<String> domainOwner;

    /**
     * @return The account number of the AWS account that owns the domain.
     * 
     */
    public Output<String> domainOwner() {
        return this.domainOwner;
    }
    /**
     * An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
     * 
     */
    @Export(name="externalConnections", refs={RepositoryExternalConnections.class}, tree="[0]")
    private Output</* @Nullable */ RepositoryExternalConnections> externalConnections;

    /**
     * @return An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
     * 
     */
    public Output<Optional<RepositoryExternalConnections>> externalConnections() {
        return Codegen.optional(this.externalConnections);
    }
    /**
     * The name of the repository to create.
     * 
     */
    @Export(name="repository", refs={String.class}, tree="[0]")
    private Output<String> repository;

    /**
     * @return The name of the repository to create.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }
    /**
     * Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
     * 
     */
    @Export(name="upstreams", refs={List.class,RepositoryUpstream.class}, tree="[0,1]")
    private Output</* @Nullable */ List<RepositoryUpstream>> upstreams;

    /**
     * @return A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
     * 
     */
    public Output<Optional<List<RepositoryUpstream>>> upstreams() {
        return Codegen.optional(this.upstreams);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Repository(String name) {
        this(name, RepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Repository(String name, RepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Repository(String name, RepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:codeartifact/repository:Repository", name, args == null ? RepositoryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Repository(String name, Output<String> id, @Nullable RepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:codeartifact/repository:Repository", name, state, makeResourceOptions(options, id));
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
    public static Repository get(String name, Output<String> id, @Nullable RepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Repository(name, id, state, options);
    }
}
