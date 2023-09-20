// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.neptune;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.neptune.ClusterInstanceArgs;
import com.pulumi.aws.neptune.inputs.ClusterInstanceState;
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
 * A Cluster Instance Resource defines attributes that are specific to a single instance in a Neptune Cluster.
 * 
 * You can simply add neptune instances and Neptune manages the replication. You can use the count
 * meta-parameter to make multiple instances and join them all to the same Neptune Cluster, or you may specify different Cluster Instance resources with various `instance_class` sizes.
 * 
 * ## Example Usage
 * 
 * The following example will create a neptune cluster with two neptune instances(one writer and one reader).
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.neptune.Cluster;
 * import com.pulumi.aws.neptune.ClusterArgs;
 * import com.pulumi.aws.neptune.ClusterInstance;
 * import com.pulumi.aws.neptune.ClusterInstanceArgs;
 * import com.pulumi.codegen.internal.KeyedValue;
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
 *             .clusterIdentifier(&#34;neptune-cluster-demo&#34;)
 *             .engine(&#34;neptune&#34;)
 *             .backupRetentionPeriod(5)
 *             .preferredBackupWindow(&#34;07:00-09:00&#34;)
 *             .skipFinalSnapshot(true)
 *             .iamDatabaseAuthenticationEnabled(true)
 *             .applyImmediately(true)
 *             .build());
 * 
 *         for (var i = 0; i &lt; 2; i++) {
 *             new ClusterInstance(&#34;example-&#34; + i, ClusterInstanceArgs.builder()            
 *                 .clusterIdentifier(default_.id())
 *                 .engine(&#34;neptune&#34;)
 *                 .instanceClass(&#34;db.r4.large&#34;)
 *                 .applyImmediately(true)
 *                 .build());
 * 
 *         
 * }
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import `aws_neptune_cluster_instance` using the instance identifier. For example:
 * 
 * ```sh
 *  $ pulumi import aws:neptune/clusterInstance:ClusterInstance example my-instance
 * ```
 * 
 */
@ResourceType(type="aws:neptune/clusterInstance:ClusterInstance")
public class ClusterInstance extends com.pulumi.resources.CustomResource {
    /**
     * The hostname of the instance. See also `endpoint` and `port`.
     * 
     */
    @Export(name="address", refs={String.class}, tree="[0]")
    private Output<String> address;

    /**
     * @return The hostname of the instance. See also `endpoint` and `port`.
     * 
     */
    public Output<String> address() {
        return this.address;
    }
    /**
     * Specifies whether any instance modifications
     * are applied immediately, or during the next maintenance window. Default is`false`.
     * 
     */
    @Export(name="applyImmediately", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> applyImmediately;

    /**
     * @return Specifies whether any instance modifications
     * are applied immediately, or during the next maintenance window. Default is`false`.
     * 
     */
    public Output<Boolean> applyImmediately() {
        return this.applyImmediately;
    }
    /**
     * Amazon Resource Name (ARN) of neptune instance
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of neptune instance
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Indicates that minor engine upgrades will be applied automatically to the instance during the maintenance window. Default is `true`.
     * 
     */
    @Export(name="autoMinorVersionUpgrade", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoMinorVersionUpgrade;

    /**
     * @return Indicates that minor engine upgrades will be applied automatically to the instance during the maintenance window. Default is `true`.
     * 
     */
    public Output<Optional<Boolean>> autoMinorVersionUpgrade() {
        return Codegen.optional(this.autoMinorVersionUpgrade);
    }
    /**
     * The EC2 Availability Zone that the neptune instance is created in.
     * 
     */
    @Export(name="availabilityZone", refs={String.class}, tree="[0]")
    private Output<String> availabilityZone;

    /**
     * @return The EC2 Availability Zone that the neptune instance is created in.
     * 
     */
    public Output<String> availabilityZone() {
        return this.availabilityZone;
    }
    /**
     * The identifier of the `aws.neptune.Cluster` in which to launch this instance.
     * 
     */
    @Export(name="clusterIdentifier", refs={String.class}, tree="[0]")
    private Output<String> clusterIdentifier;

    /**
     * @return The identifier of the `aws.neptune.Cluster` in which to launch this instance.
     * 
     */
    public Output<String> clusterIdentifier() {
        return this.clusterIdentifier;
    }
    /**
     * The region-unique, immutable identifier for the neptune instance.
     * 
     */
    @Export(name="dbiResourceId", refs={String.class}, tree="[0]")
    private Output<String> dbiResourceId;

    /**
     * @return The region-unique, immutable identifier for the neptune instance.
     * 
     */
    public Output<String> dbiResourceId() {
        return this.dbiResourceId;
    }
    /**
     * The connection endpoint in `address:port` format.
     * 
     */
    @Export(name="endpoint", refs={String.class}, tree="[0]")
    private Output<String> endpoint;

    /**
     * @return The connection endpoint in `address:port` format.
     * 
     */
    public Output<String> endpoint() {
        return this.endpoint;
    }
    /**
     * The name of the database engine to be used for the neptune instance. Defaults to `neptune`. Valid Values: `neptune`.
     * 
     */
    @Export(name="engine", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> engine;

    /**
     * @return The name of the database engine to be used for the neptune instance. Defaults to `neptune`. Valid Values: `neptune`.
     * 
     */
    public Output<Optional<String>> engine() {
        return Codegen.optional(this.engine);
    }
    /**
     * The neptune engine version.
     * 
     */
    @Export(name="engineVersion", refs={String.class}, tree="[0]")
    private Output<String> engineVersion;

    /**
     * @return The neptune engine version.
     * 
     */
    public Output<String> engineVersion() {
        return this.engineVersion;
    }
    /**
     * The identifier for the neptune instance, if omitted, this provider will assign a random, unique identifier.
     * 
     */
    @Export(name="identifier", refs={String.class}, tree="[0]")
    private Output<String> identifier;

    /**
     * @return The identifier for the neptune instance, if omitted, this provider will assign a random, unique identifier.
     * 
     */
    public Output<String> identifier() {
        return this.identifier;
    }
    /**
     * Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
     * 
     */
    @Export(name="identifierPrefix", refs={String.class}, tree="[0]")
    private Output<String> identifierPrefix;

    /**
     * @return Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
     * 
     */
    public Output<String> identifierPrefix() {
        return this.identifierPrefix;
    }
    /**
     * The instance class to use.
     * 
     */
    @Export(name="instanceClass", refs={String.class}, tree="[0]")
    private Output<String> instanceClass;

    /**
     * @return The instance class to use.
     * 
     */
    public Output<String> instanceClass() {
        return this.instanceClass;
    }
    /**
     * The ARN for the KMS encryption key if one is set to the neptune cluster.
     * 
     */
    @Export(name="kmsKeyArn", refs={String.class}, tree="[0]")
    private Output<String> kmsKeyArn;

    /**
     * @return The ARN for the KMS encryption key if one is set to the neptune cluster.
     * 
     */
    public Output<String> kmsKeyArn() {
        return this.kmsKeyArn;
    }
    /**
     * The name of the neptune parameter group to associate with this instance.
     * 
     */
    @Export(name="neptuneParameterGroupName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> neptuneParameterGroupName;

    /**
     * @return The name of the neptune parameter group to associate with this instance.
     * 
     */
    public Output<Optional<String>> neptuneParameterGroupName() {
        return Codegen.optional(this.neptuneParameterGroupName);
    }
    /**
     * A subnet group to associate with this neptune instance. **NOTE:** This must match the `neptune_subnet_group_name` of the attached `aws.neptune.Cluster`.
     * 
     */
    @Export(name="neptuneSubnetGroupName", refs={String.class}, tree="[0]")
    private Output<String> neptuneSubnetGroupName;

    /**
     * @return A subnet group to associate with this neptune instance. **NOTE:** This must match the `neptune_subnet_group_name` of the attached `aws.neptune.Cluster`.
     * 
     */
    public Output<String> neptuneSubnetGroupName() {
        return this.neptuneSubnetGroupName;
    }
    /**
     * The port on which the DB accepts connections. Defaults to `8182`.
     * 
     */
    @Export(name="port", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> port;

    /**
     * @return The port on which the DB accepts connections. Defaults to `8182`.
     * 
     */
    public Output<Optional<Integer>> port() {
        return Codegen.optional(this.port);
    }
    /**
     * The daily time range during which automated backups are created if automated backups are enabled. Eg: &#34;04:00-09:00&#34;
     * 
     */
    @Export(name="preferredBackupWindow", refs={String.class}, tree="[0]")
    private Output<String> preferredBackupWindow;

    /**
     * @return The daily time range during which automated backups are created if automated backups are enabled. Eg: &#34;04:00-09:00&#34;
     * 
     */
    public Output<String> preferredBackupWindow() {
        return this.preferredBackupWindow;
    }
    /**
     * The window to perform maintenance in.
     * Syntax: &#34;ddd:hh24:mi-ddd:hh24:mi&#34;. Eg: &#34;Mon:00:00-Mon:03:00&#34;.
     * 
     */
    @Export(name="preferredMaintenanceWindow", refs={String.class}, tree="[0]")
    private Output<String> preferredMaintenanceWindow;

    /**
     * @return The window to perform maintenance in.
     * Syntax: &#34;ddd:hh24:mi-ddd:hh24:mi&#34;. Eg: &#34;Mon:00:00-Mon:03:00&#34;.
     * 
     */
    public Output<String> preferredMaintenanceWindow() {
        return this.preferredMaintenanceWindow;
    }
    /**
     * Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
     * 
     */
    @Export(name="promotionTier", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> promotionTier;

    /**
     * @return Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
     * 
     */
    public Output<Optional<Integer>> promotionTier() {
        return Codegen.optional(this.promotionTier);
    }
    /**
     * Bool to control if instance is publicly accessible. Default is `false`.
     * 
     */
    @Export(name="publiclyAccessible", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> publiclyAccessible;

    /**
     * @return Bool to control if instance is publicly accessible. Default is `false`.
     * 
     */
    public Output<Optional<Boolean>> publiclyAccessible() {
        return Codegen.optional(this.publiclyAccessible);
    }
    /**
     * Specifies whether the neptune cluster is encrypted.
     * 
     */
    @Export(name="storageEncrypted", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> storageEncrypted;

    /**
     * @return Specifies whether the neptune cluster is encrypted.
     * 
     */
    public Output<Boolean> storageEncrypted() {
        return this.storageEncrypted;
    }
    /**
     * A map of tags to assign to the instance. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the instance. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
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
     * Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
     * 
     */
    @Export(name="writer", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> writer;

    /**
     * @return Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
     * 
     */
    public Output<Boolean> writer() {
        return this.writer;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ClusterInstance(String name) {
        this(name, ClusterInstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ClusterInstance(String name, ClusterInstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ClusterInstance(String name, ClusterInstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:neptune/clusterInstance:ClusterInstance", name, args == null ? ClusterInstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ClusterInstance(String name, Output<String> id, @Nullable ClusterInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:neptune/clusterInstance:ClusterInstance", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "tagsAll"
            ))
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
    public static ClusterInstance get(String name, Output<String> id, @Nullable ClusterInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ClusterInstance(name, id, state, options);
    }
}
