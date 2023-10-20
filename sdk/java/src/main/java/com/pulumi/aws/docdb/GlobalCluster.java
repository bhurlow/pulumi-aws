// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.docdb;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.docdb.GlobalClusterArgs;
import com.pulumi.aws.docdb.inputs.GlobalClusterState;
import com.pulumi.aws.docdb.outputs.GlobalClusterGlobalClusterMember;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an DocumentDB Global Cluster. A global cluster consists of one primary region and up to five read-only secondary regions. You issue write operations directly to the primary cluster in the primary region and Amazon DocumentDB automatically replicates the data to the secondary regions using dedicated infrastructure.
 * 
 * More information about DocumentDB Global Clusters can be found in the [DocumentDB Developer Guide](https://docs.aws.amazon.com/documentdb/latest/developerguide/global-clusters.html).
 * 
 * ## Example Usage
 * ### New DocumentDB Global Cluster
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.Provider;
 * import com.pulumi.aws.ProviderArgs;
 * import com.pulumi.aws.docdb.GlobalCluster;
 * import com.pulumi.aws.docdb.GlobalClusterArgs;
 * import com.pulumi.aws.docdb.Cluster;
 * import com.pulumi.aws.docdb.ClusterArgs;
 * import com.pulumi.aws.docdb.ClusterInstance;
 * import com.pulumi.aws.docdb.ClusterInstanceArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var primary = new Provider(&#34;primary&#34;, ProviderArgs.builder()        
 *             .region(&#34;us-east-2&#34;)
 *             .build());
 * 
 *         var secondary = new Provider(&#34;secondary&#34;, ProviderArgs.builder()        
 *             .region(&#34;us-east-1&#34;)
 *             .build());
 * 
 *         var example = new GlobalCluster(&#34;example&#34;, GlobalClusterArgs.builder()        
 *             .globalClusterIdentifier(&#34;global-test&#34;)
 *             .engine(&#34;docdb&#34;)
 *             .engineVersion(&#34;4.0.0&#34;)
 *             .build());
 * 
 *         var primaryCluster = new Cluster(&#34;primaryCluster&#34;, ClusterArgs.builder()        
 *             .engine(example.engine())
 *             .engineVersion(example.engineVersion())
 *             .clusterIdentifier(&#34;test-primary-cluster&#34;)
 *             .masterUsername(&#34;username&#34;)
 *             .masterPassword(&#34;somepass123&#34;)
 *             .globalClusterIdentifier(example.id())
 *             .dbSubnetGroupName(&#34;default&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(aws.primary())
 *                 .build());
 * 
 *         var primaryClusterInstance = new ClusterInstance(&#34;primaryClusterInstance&#34;, ClusterInstanceArgs.builder()        
 *             .engine(example.engine())
 *             .identifier(&#34;test-primary-cluster-instance&#34;)
 *             .clusterIdentifier(primaryCluster.id())
 *             .instanceClass(&#34;db.r5.large&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(aws.primary())
 *                 .build());
 * 
 *         var secondaryCluster = new Cluster(&#34;secondaryCluster&#34;, ClusterArgs.builder()        
 *             .engine(example.engine())
 *             .engineVersion(example.engineVersion())
 *             .clusterIdentifier(&#34;test-secondary-cluster&#34;)
 *             .globalClusterIdentifier(example.id())
 *             .dbSubnetGroupName(&#34;default&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(aws.secondary())
 *                 .dependsOn(primaryCluster)
 *                 .build());
 * 
 *         var secondaryClusterInstance = new ClusterInstance(&#34;secondaryClusterInstance&#34;, ClusterInstanceArgs.builder()        
 *             .engine(example.engine())
 *             .identifier(&#34;test-secondary-cluster-instance&#34;)
 *             .clusterIdentifier(secondaryCluster.id())
 *             .instanceClass(&#34;db.r5.large&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(aws.secondary())
 *                 .dependsOn(primaryClusterInstance)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### New Global Cluster From Existing DB Cluster
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.docdb.Cluster;
 * import com.pulumi.aws.docdb.GlobalCluster;
 * import com.pulumi.aws.docdb.GlobalClusterArgs;
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
 *         var exampleCluster = new Cluster(&#34;exampleCluster&#34;);
 * 
 *         var exampleGlobalCluster = new GlobalCluster(&#34;exampleGlobalCluster&#34;, GlobalClusterArgs.builder()        
 *             .globalClusterIdentifier(&#34;example&#34;)
 *             .sourceDbClusterIdentifier(exampleCluster.arn())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import `aws_docdb_global_cluster` using the Global Cluster identifier. For example:
 * 
 * ```sh
 *  $ pulumi import aws:docdb/globalCluster:GlobalCluster example example
 * ```
 *  Certain resource arguments, like `source_db_cluster_identifier`, do not have an API method for reading the information after creation. If the argument is set in the Pulumi program on an imported resource, Pulumi will always show a difference. To workaround this behavior, either omit the argument from the Pulumi program or use `ignore_changes` to hide the difference. For example:
 * 
 */
@ResourceType(type="aws:docdb/globalCluster:GlobalCluster")
public class GlobalCluster extends com.pulumi.resources.CustomResource {
    /**
     * Global Cluster Amazon Resource Name (ARN)
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Global Cluster Amazon Resource Name (ARN)
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Name for an automatically created database on cluster creation.
     * 
     */
    @Export(name="databaseName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> databaseName;

    /**
     * @return Name for an automatically created database on cluster creation.
     * 
     */
    public Output<Optional<String>> databaseName() {
        return Codegen.optional(this.databaseName);
    }
    /**
     * If the Global Cluster should have deletion protection enabled. The database can&#39;t be deleted when this value is set to `true`. The default is `false`.
     * 
     */
    @Export(name="deletionProtection", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deletionProtection;

    /**
     * @return If the Global Cluster should have deletion protection enabled. The database can&#39;t be deleted when this value is set to `true`. The default is `false`.
     * 
     */
    public Output<Optional<Boolean>> deletionProtection() {
        return Codegen.optional(this.deletionProtection);
    }
    /**
     * Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Current Valid values: `docdb`. Defaults to `docdb`. Conflicts with `source_db_cluster_identifier`.
     * 
     */
    @Export(name="engine", refs={String.class}, tree="[0]")
    private Output<String> engine;

    /**
     * @return Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Current Valid values: `docdb`. Defaults to `docdb`. Conflicts with `source_db_cluster_identifier`.
     * 
     */
    public Output<String> engine() {
        return this.engine;
    }
    /**
     * Engine version of the global database. Upgrading the engine version will result in all cluster members being immediately updated and will.
     * * **NOTE:** Upgrading major versions is not supported.
     * 
     */
    @Export(name="engineVersion", refs={String.class}, tree="[0]")
    private Output<String> engineVersion;

    /**
     * @return Engine version of the global database. Upgrading the engine version will result in all cluster members being immediately updated and will.
     * * **NOTE:** Upgrading major versions is not supported.
     * 
     */
    public Output<String> engineVersion() {
        return this.engineVersion;
    }
    /**
     * The global cluster identifier.
     * 
     */
    @Export(name="globalClusterIdentifier", refs={String.class}, tree="[0]")
    private Output<String> globalClusterIdentifier;

    /**
     * @return The global cluster identifier.
     * 
     */
    public Output<String> globalClusterIdentifier() {
        return this.globalClusterIdentifier;
    }
    /**
     * Set of objects containing Global Cluster members.
     * 
     */
    @Export(name="globalClusterMembers", refs={List.class,GlobalClusterGlobalClusterMember.class}, tree="[0,1]")
    private Output<List<GlobalClusterGlobalClusterMember>> globalClusterMembers;

    /**
     * @return Set of objects containing Global Cluster members.
     * 
     */
    public Output<List<GlobalClusterGlobalClusterMember>> globalClusterMembers() {
        return this.globalClusterMembers;
    }
    /**
     * AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed.
     * 
     */
    @Export(name="globalClusterResourceId", refs={String.class}, tree="[0]")
    private Output<String> globalClusterResourceId;

    /**
     * @return AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed.
     * 
     */
    public Output<String> globalClusterResourceId() {
        return this.globalClusterResourceId;
    }
    /**
     * Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
     * 
     */
    @Export(name="sourceDbClusterIdentifier", refs={String.class}, tree="[0]")
    private Output<String> sourceDbClusterIdentifier;

    /**
     * @return Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
     * 
     */
    public Output<String> sourceDbClusterIdentifier() {
        return this.sourceDbClusterIdentifier;
    }
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    public Output<String> status() {
        return this.status;
    }
    /**
     * Specifies whether the DB cluster is encrypted. The default is `false` unless `source_db_cluster_identifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
     * 
     */
    @Export(name="storageEncrypted", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> storageEncrypted;

    /**
     * @return Specifies whether the DB cluster is encrypted. The default is `false` unless `source_db_cluster_identifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
     * 
     */
    public Output<Boolean> storageEncrypted() {
        return this.storageEncrypted;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GlobalCluster(String name) {
        this(name, GlobalClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GlobalCluster(String name, GlobalClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GlobalCluster(String name, GlobalClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:docdb/globalCluster:GlobalCluster", name, args == null ? GlobalClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GlobalCluster(String name, Output<String> id, @Nullable GlobalClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:docdb/globalCluster:GlobalCluster", name, state, makeResourceOptions(options, id));
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
    public static GlobalCluster get(String name, Output<String> id, @Nullable GlobalClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GlobalCluster(name, id, state, options);
    }
}
