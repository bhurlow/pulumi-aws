// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.rds;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.rds.ClusterEndpointArgs;
import com.pulumi.aws.rds.inputs.ClusterEndpointState;
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
 * Manages an RDS Aurora Cluster Endpoint.
 * You can refer to the [User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Overview.Endpoints.html#Aurora.Endpoints.Cluster).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.rds.Cluster;
 * import com.pulumi.aws.rds.ClusterArgs;
 * import com.pulumi.aws.rds.ClusterInstance;
 * import com.pulumi.aws.rds.ClusterInstanceArgs;
 * import com.pulumi.aws.rds.ClusterEndpoint;
 * import com.pulumi.aws.rds.ClusterEndpointArgs;
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
 *         var default_ = new Cluster(&#34;default&#34;, ClusterArgs.builder()        
 *             .clusterIdentifier(&#34;aurora-cluster-demo&#34;)
 *             .availabilityZones(            
 *                 &#34;us-west-2a&#34;,
 *                 &#34;us-west-2b&#34;,
 *                 &#34;us-west-2c&#34;)
 *             .databaseName(&#34;mydb&#34;)
 *             .masterUsername(&#34;foo&#34;)
 *             .masterPassword(&#34;bar&#34;)
 *             .backupRetentionPeriod(5)
 *             .preferredBackupWindow(&#34;07:00-09:00&#34;)
 *             .build());
 * 
 *         var test1 = new ClusterInstance(&#34;test1&#34;, ClusterInstanceArgs.builder()        
 *             .applyImmediately(true)
 *             .clusterIdentifier(default_.id())
 *             .identifier(&#34;test1&#34;)
 *             .instanceClass(&#34;db.t2.small&#34;)
 *             .engine(default_.engine())
 *             .engineVersion(default_.engineVersion())
 *             .build());
 * 
 *         var test2 = new ClusterInstance(&#34;test2&#34;, ClusterInstanceArgs.builder()        
 *             .applyImmediately(true)
 *             .clusterIdentifier(default_.id())
 *             .identifier(&#34;test2&#34;)
 *             .instanceClass(&#34;db.t2.small&#34;)
 *             .engine(default_.engine())
 *             .engineVersion(default_.engineVersion())
 *             .build());
 * 
 *         var test3 = new ClusterInstance(&#34;test3&#34;, ClusterInstanceArgs.builder()        
 *             .applyImmediately(true)
 *             .clusterIdentifier(default_.id())
 *             .identifier(&#34;test3&#34;)
 *             .instanceClass(&#34;db.t2.small&#34;)
 *             .engine(default_.engine())
 *             .engineVersion(default_.engineVersion())
 *             .build());
 * 
 *         var eligible = new ClusterEndpoint(&#34;eligible&#34;, ClusterEndpointArgs.builder()        
 *             .clusterIdentifier(default_.id())
 *             .clusterEndpointIdentifier(&#34;reader&#34;)
 *             .customEndpointType(&#34;READER&#34;)
 *             .excludedMembers(            
 *                 test1.id(),
 *                 test2.id())
 *             .build());
 * 
 *         var static_ = new ClusterEndpoint(&#34;static&#34;, ClusterEndpointArgs.builder()        
 *             .clusterIdentifier(default_.id())
 *             .clusterEndpointIdentifier(&#34;static&#34;)
 *             .customEndpointType(&#34;READER&#34;)
 *             .staticMembers(            
 *                 test1.id(),
 *                 test3.id())
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
 *  to = aws_rds_cluster_endpoint.custom_reader
 * 
 *  id = &#34;aurora-prod-cluster-custom-reader&#34; } Using `pulumi import`, import RDS Clusters Endpoint using the `cluster_endpoint_identifier`. For exampleconsole % pulumi import aws_rds_cluster_endpoint.custom_reader aurora-prod-cluster-custom-reader [1]https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Overview.Endpoints.html#Aurora.Endpoints.Cluster
 * 
 */
@ResourceType(type="aws:rds/clusterEndpoint:ClusterEndpoint")
public class ClusterEndpoint extends com.pulumi.resources.CustomResource {
    /**
     * Amazon Resource Name (ARN) of cluster
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of cluster
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The identifier to use for the new endpoint. This parameter is stored as a lowercase string.
     * 
     */
    @Export(name="clusterEndpointIdentifier", refs={String.class}, tree="[0]")
    private Output<String> clusterEndpointIdentifier;

    /**
     * @return The identifier to use for the new endpoint. This parameter is stored as a lowercase string.
     * 
     */
    public Output<String> clusterEndpointIdentifier() {
        return this.clusterEndpointIdentifier;
    }
    /**
     * The cluster identifier.
     * 
     */
    @Export(name="clusterIdentifier", refs={String.class}, tree="[0]")
    private Output<String> clusterIdentifier;

    /**
     * @return The cluster identifier.
     * 
     */
    public Output<String> clusterIdentifier() {
        return this.clusterIdentifier;
    }
    /**
     * The type of the endpoint. One of: READER , ANY .
     * 
     */
    @Export(name="customEndpointType", refs={String.class}, tree="[0]")
    private Output<String> customEndpointType;

    /**
     * @return The type of the endpoint. One of: READER , ANY .
     * 
     */
    public Output<String> customEndpointType() {
        return this.customEndpointType;
    }
    /**
     * A custom endpoint for the Aurora cluster
     * 
     */
    @Export(name="endpoint", refs={String.class}, tree="[0]")
    private Output<String> endpoint;

    /**
     * @return A custom endpoint for the Aurora cluster
     * 
     */
    public Output<String> endpoint() {
        return this.endpoint;
    }
    /**
     * List of DB instance identifiers that aren&#39;t part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Conflicts with `static_members`.
     * 
     */
    @Export(name="excludedMembers", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> excludedMembers;

    /**
     * @return List of DB instance identifiers that aren&#39;t part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Conflicts with `static_members`.
     * 
     */
    public Output<Optional<List<String>>> excludedMembers() {
        return Codegen.optional(this.excludedMembers);
    }
    /**
     * List of DB instance identifiers that are part of the custom endpoint group. Conflicts with `excluded_members`.
     * 
     */
    @Export(name="staticMembers", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> staticMembers;

    /**
     * @return List of DB instance identifiers that are part of the custom endpoint group. Conflicts with `excluded_members`.
     * 
     */
    public Output<Optional<List<String>>> staticMembers() {
        return Codegen.optional(this.staticMembers);
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ClusterEndpoint(String name) {
        this(name, ClusterEndpointArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ClusterEndpoint(String name, ClusterEndpointArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ClusterEndpoint(String name, ClusterEndpointArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:rds/clusterEndpoint:ClusterEndpoint", name, args == null ? ClusterEndpointArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ClusterEndpoint(String name, Output<String> id, @Nullable ClusterEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:rds/clusterEndpoint:ClusterEndpoint", name, state, makeResourceOptions(options, id));
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
    public static ClusterEndpoint get(String name, Output<String> id, @Nullable ClusterEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ClusterEndpoint(name, id, state, options);
    }
}
