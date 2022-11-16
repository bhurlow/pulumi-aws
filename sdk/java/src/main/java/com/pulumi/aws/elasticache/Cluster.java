// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.elasticache;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.elasticache.ClusterArgs;
import com.pulumi.aws.elasticache.inputs.ClusterState;
import com.pulumi.aws.elasticache.outputs.ClusterCacheNode;
import com.pulumi.aws.elasticache.outputs.ClusterLogDeliveryConfiguration;
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
 * Provides an ElastiCache Cluster resource, which manages either a
 * [Memcached cluster](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/WhatIs.html), a
 * [single-node Redis instance](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/WhatIs.html), or a
 * [read replica in a Redis (Cluster Mode Enabled) replication group].
 * 
 * For working with Redis (Cluster Mode Enabled) replication groups, see the
 * `aws.elasticache.ReplicationGroup` resource.
 * 
 * &gt; **Note:** When you change an attribute, such as `num_cache_nodes`, by default
 * it is applied in the next maintenance window. Because of this, this provider may report
 * a difference in its planning phase because the actual modification has not yet taken
 * place. You can use the `apply_immediately` flag to instruct the service to apply the
 * change immediately. Using `apply_immediately` can result in a brief downtime as the server reboots.
 * See the AWS Documentation on Modifying an ElastiCache Cache Cluster for
 * [ElastiCache for Memcached](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/Clusters.Modify.html) or
 * [ElastiCache for Redis](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.Modify.html)
 * for more information.
 * 
 * &gt; **Note:** Any attribute changes that re-create the resource will be applied immediately, regardless of the value of `apply_immediately`.
 * 
 * ## Example Usage
 * ### Memcached Cluster
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.elasticache.Cluster;
 * import com.pulumi.aws.elasticache.ClusterArgs;
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
 *         var example = new Cluster(&#34;example&#34;, ClusterArgs.builder()        
 *             .engine(&#34;memcached&#34;)
 *             .nodeType(&#34;cache.m4.large&#34;)
 *             .numCacheNodes(2)
 *             .parameterGroupName(&#34;default.memcached1.4&#34;)
 *             .port(11211)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Redis Instance
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.elasticache.Cluster;
 * import com.pulumi.aws.elasticache.ClusterArgs;
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
 *         var example = new Cluster(&#34;example&#34;, ClusterArgs.builder()        
 *             .engine(&#34;redis&#34;)
 *             .engineVersion(&#34;3.2.10&#34;)
 *             .nodeType(&#34;cache.m4.large&#34;)
 *             .numCacheNodes(1)
 *             .parameterGroupName(&#34;default.redis3.2&#34;)
 *             .port(6379)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Redis Cluster Mode Disabled Read Replica Instance
 * 
 * These inherit their settings from the replication group.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.elasticache.Cluster;
 * import com.pulumi.aws.elasticache.ClusterArgs;
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
 *         var replica = new Cluster(&#34;replica&#34;, ClusterArgs.builder()        
 *             .replicationGroupId(aws_elasticache_replication_group.example().id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Redis Log Delivery configuration
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.elasticache.Cluster;
 * import com.pulumi.aws.elasticache.ClusterArgs;
 * import com.pulumi.aws.elasticache.inputs.ClusterLogDeliveryConfigurationArgs;
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
 *         var test = new Cluster(&#34;test&#34;, ClusterArgs.builder()        
 *             .engine(&#34;redis&#34;)
 *             .nodeType(&#34;cache.t3.micro&#34;)
 *             .numCacheNodes(1)
 *             .port(6379)
 *             .applyImmediately(true)
 *             .logDeliveryConfigurations(            
 *                 ClusterLogDeliveryConfigurationArgs.builder()
 *                     .destination(aws_cloudwatch_log_group.example().name())
 *                     .destinationType(&#34;cloudwatch-logs&#34;)
 *                     .logFormat(&#34;text&#34;)
 *                     .logType(&#34;slow-log&#34;)
 *                     .build(),
 *                 ClusterLogDeliveryConfigurationArgs.builder()
 *                     .destination(aws_kinesis_firehose_delivery_stream.example().name())
 *                     .destinationType(&#34;kinesis-firehose&#34;)
 *                     .logFormat(&#34;json&#34;)
 *                     .logType(&#34;engine-log&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ElastiCache Clusters can be imported using the `cluster_id`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:elasticache/cluster:Cluster my_cluster my_cluster
 * ```
 * 
 */
@ResourceType(type="aws:elasticache/cluster:Cluster")
public class Cluster extends com.pulumi.resources.CustomResource {
    /**
     * Whether any database modifications are applied immediately, or during the next maintenance window. Default is `false`. See [Amazon ElastiCache Documentation for more information.](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheCluster.html).
     * 
     */
    @Export(name="applyImmediately", type=Boolean.class, parameters={})
    private Output<Boolean> applyImmediately;

    /**
     * @return Whether any database modifications are applied immediately, or during the next maintenance window. Default is `false`. See [Amazon ElastiCache Documentation for more information.](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheCluster.html).
     * 
     */
    public Output<Boolean> applyImmediately() {
        return this.applyImmediately;
    }
    /**
     * The ARN of the created ElastiCache Cluster.
     * 
     */
    @Export(name="arn", type=String.class, parameters={})
    private Output<String> arn;

    /**
     * @return The ARN of the created ElastiCache Cluster.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Specifies whether minor version engine upgrades will be applied automatically to the underlying Cache Cluster instances during the maintenance window.
     * Only supported for engine type `&#34;redis&#34;` and if the engine version is 6 or higher.
     * Defaults to `true`.
     * 
     */
    @Export(name="autoMinorVersionUpgrade", type=String.class, parameters={})
    private Output</* @Nullable */ String> autoMinorVersionUpgrade;

    /**
     * @return Specifies whether minor version engine upgrades will be applied automatically to the underlying Cache Cluster instances during the maintenance window.
     * Only supported for engine type `&#34;redis&#34;` and if the engine version is 6 or higher.
     * Defaults to `true`.
     * 
     */
    public Output<Optional<String>> autoMinorVersionUpgrade() {
        return Codegen.optional(this.autoMinorVersionUpgrade);
    }
    /**
     * Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferred_availability_zones` instead. Default: System chosen Availability Zone. Changing this value will re-create the resource.
     * 
     */
    @Export(name="availabilityZone", type=String.class, parameters={})
    private Output<String> availabilityZone;

    /**
     * @return Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferred_availability_zones` instead. Default: System chosen Availability Zone. Changing this value will re-create the resource.
     * 
     */
    public Output<String> availabilityZone() {
        return this.availabilityZone;
    }
    /**
     * Whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster&#39;s region. Valid values for this parameter are `single-az` or `cross-az`, default is `single-az`. If you want to choose `cross-az`, `num_cache_nodes` must be greater than `1`.
     * 
     */
    @Export(name="azMode", type=String.class, parameters={})
    private Output<String> azMode;

    /**
     * @return Whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster&#39;s region. Valid values for this parameter are `single-az` or `cross-az`, default is `single-az`. If you want to choose `cross-az`, `num_cache_nodes` must be greater than `1`.
     * 
     */
    public Output<String> azMode() {
        return this.azMode;
    }
    /**
     * List of node objects including `id`, `address`, `port` and `availability_zone`.
     * 
     */
    @Export(name="cacheNodes", type=List.class, parameters={ClusterCacheNode.class})
    private Output<List<ClusterCacheNode>> cacheNodes;

    /**
     * @return List of node objects including `id`, `address`, `port` and `availability_zone`.
     * 
     */
    public Output<List<ClusterCacheNode>> cacheNodes() {
        return this.cacheNodes;
    }
    /**
     * (Memcached only) DNS name of the cache cluster without the port appended.
     * 
     */
    @Export(name="clusterAddress", type=String.class, parameters={})
    private Output<String> clusterAddress;

    /**
     * @return (Memcached only) DNS name of the cache cluster without the port appended.
     * 
     */
    public Output<String> clusterAddress() {
        return this.clusterAddress;
    }
    /**
     * Group identifier. ElastiCache converts this name to lowercase. Changing this value will re-create the resource.
     * 
     */
    @Export(name="clusterId", type=String.class, parameters={})
    private Output<String> clusterId;

    /**
     * @return Group identifier. ElastiCache converts this name to lowercase. Changing this value will re-create the resource.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * (Memcached only) Configuration endpoint to allow host discovery.
     * 
     */
    @Export(name="configurationEndpoint", type=String.class, parameters={})
    private Output<String> configurationEndpoint;

    /**
     * @return (Memcached only) Configuration endpoint to allow host discovery.
     * 
     */
    public Output<String> configurationEndpoint() {
        return this.configurationEndpoint;
    }
    /**
     * Name of the cache engine to be used for this cache cluster. Valid values are `memcached` or `redis`.
     * 
     */
    @Export(name="engine", type=String.class, parameters={})
    private Output<String> engine;

    /**
     * @return Name of the cache engine to be used for this cache cluster. Valid values are `memcached` or `redis`.
     * 
     */
    public Output<String> engine() {
        return this.engine;
    }
    /**
     * Version number of the cache engine to be used.
     * If not set, defaults to the latest version.
     * See [Describe Cache Engine Versions](https://docs.aws.amazon.com/cli/latest/reference/elasticache/describe-cache-engine-versions.html) in the AWS Documentation for supported versions.
     * When `engine` is `redis` and the version is 6 or higher, the major and minor version can be set, e.g., `6.2`,
     * or the minor version can be unspecified which will use the latest version at creation time, e.g., `6.x`.
     * Otherwise, specify the full version desired, e.g., `5.0.6`.
     * The actual engine version used is returned in the attribute `engine_version_actual`, see Attributes Reference below.
     * 
     */
    @Export(name="engineVersion", type=String.class, parameters={})
    private Output<String> engineVersion;

    /**
     * @return Version number of the cache engine to be used.
     * If not set, defaults to the latest version.
     * See [Describe Cache Engine Versions](https://docs.aws.amazon.com/cli/latest/reference/elasticache/describe-cache-engine-versions.html) in the AWS Documentation for supported versions.
     * When `engine` is `redis` and the version is 6 or higher, the major and minor version can be set, e.g., `6.2`,
     * or the minor version can be unspecified which will use the latest version at creation time, e.g., `6.x`.
     * Otherwise, specify the full version desired, e.g., `5.0.6`.
     * The actual engine version used is returned in the attribute `engine_version_actual`, see Attributes Reference below.
     * 
     */
    public Output<String> engineVersion() {
        return this.engineVersion;
    }
    /**
     * Because ElastiCache pulls the latest minor or patch for a version, this attribute returns the running version of the cache engine.
     * 
     */
    @Export(name="engineVersionActual", type=String.class, parameters={})
    private Output<String> engineVersionActual;

    /**
     * @return Because ElastiCache pulls the latest minor or patch for a version, this attribute returns the running version of the cache engine.
     * 
     */
    public Output<String> engineVersionActual() {
        return this.engineVersionActual;
    }
    /**
     * Name of your final cluster snapshot. If omitted, no final snapshot will be made.
     * 
     */
    @Export(name="finalSnapshotIdentifier", type=String.class, parameters={})
    private Output</* @Nullable */ String> finalSnapshotIdentifier;

    /**
     * @return Name of your final cluster snapshot. If omitted, no final snapshot will be made.
     * 
     */
    public Output<Optional<String>> finalSnapshotIdentifier() {
        return Codegen.optional(this.finalSnapshotIdentifier);
    }
    /**
     * Specifies the destination and format of Redis [SLOWLOG](https://redis.io/commands/slowlog) or Redis [Engine Log](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Log_Delivery.html#Log_contents-engine-log). See the documentation on [Amazon ElastiCache](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Log_Delivery.html). See Log Delivery Configuration below for more details.
     * 
     */
    @Export(name="logDeliveryConfigurations", type=List.class, parameters={ClusterLogDeliveryConfiguration.class})
    private Output</* @Nullable */ List<ClusterLogDeliveryConfiguration>> logDeliveryConfigurations;

    /**
     * @return Specifies the destination and format of Redis [SLOWLOG](https://redis.io/commands/slowlog) or Redis [Engine Log](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Log_Delivery.html#Log_contents-engine-log). See the documentation on [Amazon ElastiCache](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Log_Delivery.html). See Log Delivery Configuration below for more details.
     * 
     */
    public Output<Optional<List<ClusterLogDeliveryConfiguration>>> logDeliveryConfigurations() {
        return Codegen.optional(this.logDeliveryConfigurations);
    }
    /**
     * Specifies the weekly time range for when maintenance
     * on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
     * The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`.
     * 
     */
    @Export(name="maintenanceWindow", type=String.class, parameters={})
    private Output<String> maintenanceWindow;

    /**
     * @return Specifies the weekly time range for when maintenance
     * on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
     * The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`.
     * 
     */
    public Output<String> maintenanceWindow() {
        return this.maintenanceWindow;
    }
    /**
     * The instance class used. See AWS documentation for information on [supported node types for Redis](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html) and [guidance on selecting node types for Redis](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/nodes-select-size.html). See AWS documentation for information on [supported node types for Memcached](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/CacheNodes.SupportedTypes.html) and [guidance on selecting node types for Memcached](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/nodes-select-size.html). For Memcached, changing this value will re-create the resource.
     * 
     */
    @Export(name="nodeType", type=String.class, parameters={})
    private Output<String> nodeType;

    /**
     * @return The instance class used. See AWS documentation for information on [supported node types for Redis](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html) and [guidance on selecting node types for Redis](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/nodes-select-size.html). See AWS documentation for information on [supported node types for Memcached](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/CacheNodes.SupportedTypes.html) and [guidance on selecting node types for Memcached](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/nodes-select-size.html). For Memcached, changing this value will re-create the resource.
     * 
     */
    public Output<String> nodeType() {
        return this.nodeType;
    }
    /**
     * ARN of an SNS topic to send ElastiCache notifications to. Example: `arn:aws:sns:us-east-1:012345678999:my_sns_topic`.
     * 
     */
    @Export(name="notificationTopicArn", type=String.class, parameters={})
    private Output</* @Nullable */ String> notificationTopicArn;

    /**
     * @return ARN of an SNS topic to send ElastiCache notifications to. Example: `arn:aws:sns:us-east-1:012345678999:my_sns_topic`.
     * 
     */
    public Output<Optional<String>> notificationTopicArn() {
        return Codegen.optional(this.notificationTopicArn);
    }
    /**
     * The initial number of cache nodes that the cache cluster will have. For Redis, this value must be 1. For Memcached, this value must be between 1 and 40. If this number is reduced on subsequent runs, the highest numbered nodes will be removed.
     * 
     */
    @Export(name="numCacheNodes", type=Integer.class, parameters={})
    private Output<Integer> numCacheNodes;

    /**
     * @return The initial number of cache nodes that the cache cluster will have. For Redis, this value must be 1. For Memcached, this value must be between 1 and 40. If this number is reduced on subsequent runs, the highest numbered nodes will be removed.
     * 
     */
    public Output<Integer> numCacheNodes() {
        return this.numCacheNodes;
    }
    /**
     * The name of the parameter group to associate with this cache cluster.
     * 
     */
    @Export(name="parameterGroupName", type=String.class, parameters={})
    private Output<String> parameterGroupName;

    /**
     * @return The name of the parameter group to associate with this cache cluster.
     * 
     */
    public Output<String> parameterGroupName() {
        return this.parameterGroupName;
    }
    /**
     * The port number on which each of the cache nodes will accept connections. For Memcached the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replication_group_id`. Changing this value will re-create the resource.
     * 
     */
    @Export(name="port", type=Integer.class, parameters={})
    private Output<Integer> port;

    /**
     * @return The port number on which each of the cache nodes will accept connections. For Memcached the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replication_group_id`. Changing this value will re-create the resource.
     * 
     */
    public Output<Integer> port() {
        return this.port;
    }
    /**
     * List of the Availability Zones in which cache nodes are created. If you are creating your cluster in an Amazon VPC you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group. The number of Availability Zones listed must equal the value of `num_cache_nodes`. If you want all the nodes in the same Availability Zone, use `availability_zone` instead, or repeat the Availability Zone multiple times in the list. Default: System chosen Availability Zones. Detecting drift of existing node availability zone is not currently supported. Updating this argument by itself to migrate existing node availability zones is not currently supported and will show a perpetual difference.
     * 
     */
    @Export(name="preferredAvailabilityZones", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> preferredAvailabilityZones;

    /**
     * @return List of the Availability Zones in which cache nodes are created. If you are creating your cluster in an Amazon VPC you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group. The number of Availability Zones listed must equal the value of `num_cache_nodes`. If you want all the nodes in the same Availability Zone, use `availability_zone` instead, or repeat the Availability Zone multiple times in the list. Default: System chosen Availability Zones. Detecting drift of existing node availability zone is not currently supported. Updating this argument by itself to migrate existing node availability zones is not currently supported and will show a perpetual difference.
     * 
     */
    public Output<Optional<List<String>>> preferredAvailabilityZones() {
        return Codegen.optional(this.preferredAvailabilityZones);
    }
    /**
     * ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
     * 
     */
    @Export(name="replicationGroupId", type=String.class, parameters={})
    private Output<String> replicationGroupId;

    /**
     * @return ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
     * 
     */
    public Output<String> replicationGroupId() {
        return this.replicationGroupId;
    }
    /**
     * One or more VPC security groups associated with the cache cluster
     * 
     */
    @Export(name="securityGroupIds", type=List.class, parameters={String.class})
    private Output<List<String>> securityGroupIds;

    /**
     * @return One or more VPC security groups associated with the cache cluster
     * 
     */
    public Output<List<String>> securityGroupIds() {
        return this.securityGroupIds;
    }
    /**
     * List of security group names to associate with this cache cluster. Changing this value will re-create the resource.
     * 
     * @deprecated
     * With the retirement of EC2-Classic the security_group_names attribute has been deprecated and will be removed in a future version.
     * 
     */
    @Deprecated /* With the retirement of EC2-Classic the security_group_names attribute has been deprecated and will be removed in a future version. */
    @Export(name="securityGroupNames", type=List.class, parameters={String.class})
    private Output<List<String>> securityGroupNames;

    /**
     * @return List of security group names to associate with this cache cluster. Changing this value will re-create the resource.
     * 
     */
    public Output<List<String>> securityGroupNames() {
        return this.securityGroupNames;
    }
    /**
     * Single-element string list containing an Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3. The object name cannot contain any commas. Changing `snapshot_arns` forces a new resource.
     * 
     */
    @Export(name="snapshotArns", type=String.class, parameters={})
    private Output</* @Nullable */ String> snapshotArns;

    /**
     * @return Single-element string list containing an Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3. The object name cannot contain any commas. Changing `snapshot_arns` forces a new resource.
     * 
     */
    public Output<Optional<String>> snapshotArns() {
        return Codegen.optional(this.snapshotArns);
    }
    /**
     * Name of a snapshot from which to restore data into the new node group. Changing `snapshot_name` forces a new resource.
     * 
     */
    @Export(name="snapshotName", type=String.class, parameters={})
    private Output</* @Nullable */ String> snapshotName;

    /**
     * @return Name of a snapshot from which to restore data into the new node group. Changing `snapshot_name` forces a new resource.
     * 
     */
    public Output<Optional<String>> snapshotName() {
        return Codegen.optional(this.snapshotName);
    }
    /**
     * Number of days for which ElastiCache will retain automatic cache cluster snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off. Please note that setting a `snapshot_retention_limit` is not supported on cache.t1.micro cache nodes
     * 
     */
    @Export(name="snapshotRetentionLimit", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> snapshotRetentionLimit;

    /**
     * @return Number of days for which ElastiCache will retain automatic cache cluster snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off. Please note that setting a `snapshot_retention_limit` is not supported on cache.t1.micro cache nodes
     * 
     */
    public Output<Optional<Integer>> snapshotRetentionLimit() {
        return Codegen.optional(this.snapshotRetentionLimit);
    }
    /**
     * Daily time range (in UTC) during which ElastiCache will begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
     * 
     */
    @Export(name="snapshotWindow", type=String.class, parameters={})
    private Output<String> snapshotWindow;

    /**
     * @return Daily time range (in UTC) during which ElastiCache will begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
     * 
     */
    public Output<String> snapshotWindow() {
        return this.snapshotWindow;
    }
    /**
     * Name of the subnet group to be used for the cache cluster. Changing this value will re-create the resource.
     * 
     */
    @Export(name="subnetGroupName", type=String.class, parameters={})
    private Output<String> subnetGroupName;

    /**
     * @return Name of the subnet group to be used for the cache cluster. Changing this value will re-create the resource.
     * 
     */
    public Output<String> subnetGroupName() {
        return this.subnetGroupName;
    }
    /**
     * Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Export(name="tagsAll", type=Map.class, parameters={String.class, String.class})
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
    public Cluster(String name) {
        this(name, ClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Cluster(String name, @Nullable ClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Cluster(String name, @Nullable ClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:elasticache/cluster:Cluster", name, args == null ? ClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Cluster(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:elasticache/cluster:Cluster", name, state, makeResourceOptions(options, id));
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
    public static Cluster get(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Cluster(name, id, state, options);
    }
}
