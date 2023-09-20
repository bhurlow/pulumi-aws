// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.memorydb.inputs;

import com.pulumi.aws.memorydb.inputs.ClusterClusterEndpointArgs;
import com.pulumi.aws.memorydb.inputs.ClusterShardArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClusterState extends com.pulumi.resources.ResourceArgs {

    public static final ClusterState Empty = new ClusterState();

    /**
     * The name of the Access Control List to associate with the cluster.
     * 
     */
    @Import(name="aclName")
    private @Nullable Output<String> aclName;

    /**
     * @return The name of the Access Control List to associate with the cluster.
     * 
     */
    public Optional<Output<String>> aclName() {
        return Optional.ofNullable(this.aclName);
    }

    /**
     * The ARN of the cluster.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The ARN of the cluster.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * When set to `true`, the cluster will automatically receive minor engine version upgrades after launch. Defaults to `true`.
     * 
     */
    @Import(name="autoMinorVersionUpgrade")
    private @Nullable Output<Boolean> autoMinorVersionUpgrade;

    /**
     * @return When set to `true`, the cluster will automatically receive minor engine version upgrades after launch. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> autoMinorVersionUpgrade() {
        return Optional.ofNullable(this.autoMinorVersionUpgrade);
    }

    @Import(name="clusterEndpoints")
    private @Nullable Output<List<ClusterClusterEndpointArgs>> clusterEndpoints;

    public Optional<Output<List<ClusterClusterEndpointArgs>>> clusterEndpoints() {
        return Optional.ofNullable(this.clusterEndpoints);
    }

    /**
     * Enables data tiering. This option is not supported by all instance types. For more information, see [Data tiering](https://docs.aws.amazon.com/memorydb/latest/devguide/data-tiering.html).
     * 
     */
    @Import(name="dataTiering")
    private @Nullable Output<Boolean> dataTiering;

    /**
     * @return Enables data tiering. This option is not supported by all instance types. For more information, see [Data tiering](https://docs.aws.amazon.com/memorydb/latest/devguide/data-tiering.html).
     * 
     */
    public Optional<Output<Boolean>> dataTiering() {
        return Optional.ofNullable(this.dataTiering);
    }

    /**
     * Description for the cluster.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description for the cluster.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Patch version number of the Redis engine used by the cluster.
     * 
     */
    @Import(name="enginePatchVersion")
    private @Nullable Output<String> enginePatchVersion;

    /**
     * @return Patch version number of the Redis engine used by the cluster.
     * 
     */
    public Optional<Output<String>> enginePatchVersion() {
        return Optional.ofNullable(this.enginePatchVersion);
    }

    /**
     * Version number of the Redis engine to be used for the cluster. Downgrades are not supported.
     * 
     */
    @Import(name="engineVersion")
    private @Nullable Output<String> engineVersion;

    /**
     * @return Version number of the Redis engine to be used for the cluster. Downgrades are not supported.
     * 
     */
    public Optional<Output<String>> engineVersion() {
        return Optional.ofNullable(this.engineVersion);
    }

    /**
     * Name of the final cluster snapshot to be created when this resource is deleted. If omitted, no final snapshot will be made.
     * 
     */
    @Import(name="finalSnapshotName")
    private @Nullable Output<String> finalSnapshotName;

    /**
     * @return Name of the final cluster snapshot to be created when this resource is deleted. If omitted, no final snapshot will be made.
     * 
     */
    public Optional<Output<String>> finalSnapshotName() {
        return Optional.ofNullable(this.finalSnapshotName);
    }

    /**
     * ARN of the KMS key used to encrypt the cluster at rest.
     * 
     */
    @Import(name="kmsKeyArn")
    private @Nullable Output<String> kmsKeyArn;

    /**
     * @return ARN of the KMS key used to encrypt the cluster at rest.
     * 
     */
    public Optional<Output<String>> kmsKeyArn() {
        return Optional.ofNullable(this.kmsKeyArn);
    }

    /**
     * Specifies the weekly time range during which maintenance on the cluster is performed. Specify as a range in the format `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example: `sun:23:00-mon:01:30`.
     * 
     */
    @Import(name="maintenanceWindow")
    private @Nullable Output<String> maintenanceWindow;

    /**
     * @return Specifies the weekly time range during which maintenance on the cluster is performed. Specify as a range in the format `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example: `sun:23:00-mon:01:30`.
     * 
     */
    public Optional<Output<String>> maintenanceWindow() {
        return Optional.ofNullable(this.maintenanceWindow);
    }

    /**
     * Name of the cluster. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the cluster. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     * 
     */
    @Import(name="namePrefix")
    private @Nullable Output<String> namePrefix;

    /**
     * @return Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     * 
     */
    public Optional<Output<String>> namePrefix() {
        return Optional.ofNullable(this.namePrefix);
    }

    /**
     * The compute and memory capacity of the nodes in the cluster. See AWS documentation on [supported node types](https://docs.aws.amazon.com/memorydb/latest/devguide/nodes.supportedtypes.html) as well as [vertical scaling](https://docs.aws.amazon.com/memorydb/latest/devguide/cluster-vertical-scaling.html).
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="nodeType")
    private @Nullable Output<String> nodeType;

    /**
     * @return The compute and memory capacity of the nodes in the cluster. See AWS documentation on [supported node types](https://docs.aws.amazon.com/memorydb/latest/devguide/nodes.supportedtypes.html) as well as [vertical scaling](https://docs.aws.amazon.com/memorydb/latest/devguide/cluster-vertical-scaling.html).
     * 
     * The following arguments are optional:
     * 
     */
    public Optional<Output<String>> nodeType() {
        return Optional.ofNullable(this.nodeType);
    }

    /**
     * The number of replicas to apply to each shard, up to a maximum of 5. Defaults to `1` (i.e. 2 nodes per shard).
     * 
     */
    @Import(name="numReplicasPerShard")
    private @Nullable Output<Integer> numReplicasPerShard;

    /**
     * @return The number of replicas to apply to each shard, up to a maximum of 5. Defaults to `1` (i.e. 2 nodes per shard).
     * 
     */
    public Optional<Output<Integer>> numReplicasPerShard() {
        return Optional.ofNullable(this.numReplicasPerShard);
    }

    /**
     * The number of shards in the cluster. Defaults to `1`.
     * 
     */
    @Import(name="numShards")
    private @Nullable Output<Integer> numShards;

    /**
     * @return The number of shards in the cluster. Defaults to `1`.
     * 
     */
    public Optional<Output<Integer>> numShards() {
        return Optional.ofNullable(this.numShards);
    }

    /**
     * The name of the parameter group associated with the cluster.
     * 
     */
    @Import(name="parameterGroupName")
    private @Nullable Output<String> parameterGroupName;

    /**
     * @return The name of the parameter group associated with the cluster.
     * 
     */
    public Optional<Output<String>> parameterGroupName() {
        return Optional.ofNullable(this.parameterGroupName);
    }

    /**
     * The port number on which each of the nodes accepts connections. Defaults to `6379`.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return The port number on which each of the nodes accepts connections. Defaults to `6379`.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * Set of VPC Security Group ID-s to associate with this cluster.
     * 
     */
    @Import(name="securityGroupIds")
    private @Nullable Output<List<String>> securityGroupIds;

    /**
     * @return Set of VPC Security Group ID-s to associate with this cluster.
     * 
     */
    public Optional<Output<List<String>>> securityGroupIds() {
        return Optional.ofNullable(this.securityGroupIds);
    }

    /**
     * Set of shards in this cluster.
     * 
     */
    @Import(name="shards")
    private @Nullable Output<List<ClusterShardArgs>> shards;

    /**
     * @return Set of shards in this cluster.
     * 
     */
    public Optional<Output<List<ClusterShardArgs>>> shards() {
        return Optional.ofNullable(this.shards);
    }

    /**
     * List of ARN-s that uniquely identify RDB snapshot files stored in S3. The snapshot files will be used to populate the new cluster. Object names in the ARN-s cannot contain any commas.
     * 
     */
    @Import(name="snapshotArns")
    private @Nullable Output<List<String>> snapshotArns;

    /**
     * @return List of ARN-s that uniquely identify RDB snapshot files stored in S3. The snapshot files will be used to populate the new cluster. Object names in the ARN-s cannot contain any commas.
     * 
     */
    public Optional<Output<List<String>>> snapshotArns() {
        return Optional.ofNullable(this.snapshotArns);
    }

    /**
     * The name of a snapshot from which to restore data into the new cluster.
     * 
     */
    @Import(name="snapshotName")
    private @Nullable Output<String> snapshotName;

    /**
     * @return The name of a snapshot from which to restore data into the new cluster.
     * 
     */
    public Optional<Output<String>> snapshotName() {
        return Optional.ofNullable(this.snapshotName);
    }

    /**
     * The number of days for which MemoryDB retains automatic snapshots before deleting them. When set to `0`, automatic backups are disabled. Defaults to `0`.
     * 
     */
    @Import(name="snapshotRetentionLimit")
    private @Nullable Output<Integer> snapshotRetentionLimit;

    /**
     * @return The number of days for which MemoryDB retains automatic snapshots before deleting them. When set to `0`, automatic backups are disabled. Defaults to `0`.
     * 
     */
    public Optional<Output<Integer>> snapshotRetentionLimit() {
        return Optional.ofNullable(this.snapshotRetentionLimit);
    }

    /**
     * The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your shard. Example: `05:00-09:00`.
     * 
     */
    @Import(name="snapshotWindow")
    private @Nullable Output<String> snapshotWindow;

    /**
     * @return The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your shard. Example: `05:00-09:00`.
     * 
     */
    public Optional<Output<String>> snapshotWindow() {
        return Optional.ofNullable(this.snapshotWindow);
    }

    /**
     * ARN of the SNS topic to which cluster notifications are sent.
     * 
     */
    @Import(name="snsTopicArn")
    private @Nullable Output<String> snsTopicArn;

    /**
     * @return ARN of the SNS topic to which cluster notifications are sent.
     * 
     */
    public Optional<Output<String>> snsTopicArn() {
        return Optional.ofNullable(this.snsTopicArn);
    }

    /**
     * The name of the subnet group to be used for the cluster. Defaults to a subnet group consisting of default VPC subnets.
     * 
     */
    @Import(name="subnetGroupName")
    private @Nullable Output<String> subnetGroupName;

    /**
     * @return The name of the subnet group to be used for the cluster. Defaults to a subnet group consisting of default VPC subnets.
     * 
     */
    public Optional<Output<String>> subnetGroupName() {
        return Optional.ofNullable(this.subnetGroupName);
    }

    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * A flag to enable in-transit encryption on the cluster. When set to `false`, the `acl_name` must be `open-access`. Defaults to `true`.
     * 
     */
    @Import(name="tlsEnabled")
    private @Nullable Output<Boolean> tlsEnabled;

    /**
     * @return A flag to enable in-transit encryption on the cluster. When set to `false`, the `acl_name` must be `open-access`. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> tlsEnabled() {
        return Optional.ofNullable(this.tlsEnabled);
    }

    private ClusterState() {}

    private ClusterState(ClusterState $) {
        this.aclName = $.aclName;
        this.arn = $.arn;
        this.autoMinorVersionUpgrade = $.autoMinorVersionUpgrade;
        this.clusterEndpoints = $.clusterEndpoints;
        this.dataTiering = $.dataTiering;
        this.description = $.description;
        this.enginePatchVersion = $.enginePatchVersion;
        this.engineVersion = $.engineVersion;
        this.finalSnapshotName = $.finalSnapshotName;
        this.kmsKeyArn = $.kmsKeyArn;
        this.maintenanceWindow = $.maintenanceWindow;
        this.name = $.name;
        this.namePrefix = $.namePrefix;
        this.nodeType = $.nodeType;
        this.numReplicasPerShard = $.numReplicasPerShard;
        this.numShards = $.numShards;
        this.parameterGroupName = $.parameterGroupName;
        this.port = $.port;
        this.securityGroupIds = $.securityGroupIds;
        this.shards = $.shards;
        this.snapshotArns = $.snapshotArns;
        this.snapshotName = $.snapshotName;
        this.snapshotRetentionLimit = $.snapshotRetentionLimit;
        this.snapshotWindow = $.snapshotWindow;
        this.snsTopicArn = $.snsTopicArn;
        this.subnetGroupName = $.subnetGroupName;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.tlsEnabled = $.tlsEnabled;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterState $;

        public Builder() {
            $ = new ClusterState();
        }

        public Builder(ClusterState defaults) {
            $ = new ClusterState(Objects.requireNonNull(defaults));
        }

        /**
         * @param aclName The name of the Access Control List to associate with the cluster.
         * 
         * @return builder
         * 
         */
        public Builder aclName(@Nullable Output<String> aclName) {
            $.aclName = aclName;
            return this;
        }

        /**
         * @param aclName The name of the Access Control List to associate with the cluster.
         * 
         * @return builder
         * 
         */
        public Builder aclName(String aclName) {
            return aclName(Output.of(aclName));
        }

        /**
         * @param arn The ARN of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The ARN of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param autoMinorVersionUpgrade When set to `true`, the cluster will automatically receive minor engine version upgrades after launch. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder autoMinorVersionUpgrade(@Nullable Output<Boolean> autoMinorVersionUpgrade) {
            $.autoMinorVersionUpgrade = autoMinorVersionUpgrade;
            return this;
        }

        /**
         * @param autoMinorVersionUpgrade When set to `true`, the cluster will automatically receive minor engine version upgrades after launch. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder autoMinorVersionUpgrade(Boolean autoMinorVersionUpgrade) {
            return autoMinorVersionUpgrade(Output.of(autoMinorVersionUpgrade));
        }

        public Builder clusterEndpoints(@Nullable Output<List<ClusterClusterEndpointArgs>> clusterEndpoints) {
            $.clusterEndpoints = clusterEndpoints;
            return this;
        }

        public Builder clusterEndpoints(List<ClusterClusterEndpointArgs> clusterEndpoints) {
            return clusterEndpoints(Output.of(clusterEndpoints));
        }

        public Builder clusterEndpoints(ClusterClusterEndpointArgs... clusterEndpoints) {
            return clusterEndpoints(List.of(clusterEndpoints));
        }

        /**
         * @param dataTiering Enables data tiering. This option is not supported by all instance types. For more information, see [Data tiering](https://docs.aws.amazon.com/memorydb/latest/devguide/data-tiering.html).
         * 
         * @return builder
         * 
         */
        public Builder dataTiering(@Nullable Output<Boolean> dataTiering) {
            $.dataTiering = dataTiering;
            return this;
        }

        /**
         * @param dataTiering Enables data tiering. This option is not supported by all instance types. For more information, see [Data tiering](https://docs.aws.amazon.com/memorydb/latest/devguide/data-tiering.html).
         * 
         * @return builder
         * 
         */
        public Builder dataTiering(Boolean dataTiering) {
            return dataTiering(Output.of(dataTiering));
        }

        /**
         * @param description Description for the cluster.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description for the cluster.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param enginePatchVersion Patch version number of the Redis engine used by the cluster.
         * 
         * @return builder
         * 
         */
        public Builder enginePatchVersion(@Nullable Output<String> enginePatchVersion) {
            $.enginePatchVersion = enginePatchVersion;
            return this;
        }

        /**
         * @param enginePatchVersion Patch version number of the Redis engine used by the cluster.
         * 
         * @return builder
         * 
         */
        public Builder enginePatchVersion(String enginePatchVersion) {
            return enginePatchVersion(Output.of(enginePatchVersion));
        }

        /**
         * @param engineVersion Version number of the Redis engine to be used for the cluster. Downgrades are not supported.
         * 
         * @return builder
         * 
         */
        public Builder engineVersion(@Nullable Output<String> engineVersion) {
            $.engineVersion = engineVersion;
            return this;
        }

        /**
         * @param engineVersion Version number of the Redis engine to be used for the cluster. Downgrades are not supported.
         * 
         * @return builder
         * 
         */
        public Builder engineVersion(String engineVersion) {
            return engineVersion(Output.of(engineVersion));
        }

        /**
         * @param finalSnapshotName Name of the final cluster snapshot to be created when this resource is deleted. If omitted, no final snapshot will be made.
         * 
         * @return builder
         * 
         */
        public Builder finalSnapshotName(@Nullable Output<String> finalSnapshotName) {
            $.finalSnapshotName = finalSnapshotName;
            return this;
        }

        /**
         * @param finalSnapshotName Name of the final cluster snapshot to be created when this resource is deleted. If omitted, no final snapshot will be made.
         * 
         * @return builder
         * 
         */
        public Builder finalSnapshotName(String finalSnapshotName) {
            return finalSnapshotName(Output.of(finalSnapshotName));
        }

        /**
         * @param kmsKeyArn ARN of the KMS key used to encrypt the cluster at rest.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyArn(@Nullable Output<String> kmsKeyArn) {
            $.kmsKeyArn = kmsKeyArn;
            return this;
        }

        /**
         * @param kmsKeyArn ARN of the KMS key used to encrypt the cluster at rest.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyArn(String kmsKeyArn) {
            return kmsKeyArn(Output.of(kmsKeyArn));
        }

        /**
         * @param maintenanceWindow Specifies the weekly time range during which maintenance on the cluster is performed. Specify as a range in the format `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example: `sun:23:00-mon:01:30`.
         * 
         * @return builder
         * 
         */
        public Builder maintenanceWindow(@Nullable Output<String> maintenanceWindow) {
            $.maintenanceWindow = maintenanceWindow;
            return this;
        }

        /**
         * @param maintenanceWindow Specifies the weekly time range during which maintenance on the cluster is performed. Specify as a range in the format `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example: `sun:23:00-mon:01:30`.
         * 
         * @return builder
         * 
         */
        public Builder maintenanceWindow(String maintenanceWindow) {
            return maintenanceWindow(Output.of(maintenanceWindow));
        }

        /**
         * @param name Name of the cluster. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the cluster. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namePrefix Creates a unique name beginning with the specified prefix. Conflicts with `name`.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(@Nullable Output<String> namePrefix) {
            $.namePrefix = namePrefix;
            return this;
        }

        /**
         * @param namePrefix Creates a unique name beginning with the specified prefix. Conflicts with `name`.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(String namePrefix) {
            return namePrefix(Output.of(namePrefix));
        }

        /**
         * @param nodeType The compute and memory capacity of the nodes in the cluster. See AWS documentation on [supported node types](https://docs.aws.amazon.com/memorydb/latest/devguide/nodes.supportedtypes.html) as well as [vertical scaling](https://docs.aws.amazon.com/memorydb/latest/devguide/cluster-vertical-scaling.html).
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder nodeType(@Nullable Output<String> nodeType) {
            $.nodeType = nodeType;
            return this;
        }

        /**
         * @param nodeType The compute and memory capacity of the nodes in the cluster. See AWS documentation on [supported node types](https://docs.aws.amazon.com/memorydb/latest/devguide/nodes.supportedtypes.html) as well as [vertical scaling](https://docs.aws.amazon.com/memorydb/latest/devguide/cluster-vertical-scaling.html).
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder nodeType(String nodeType) {
            return nodeType(Output.of(nodeType));
        }

        /**
         * @param numReplicasPerShard The number of replicas to apply to each shard, up to a maximum of 5. Defaults to `1` (i.e. 2 nodes per shard).
         * 
         * @return builder
         * 
         */
        public Builder numReplicasPerShard(@Nullable Output<Integer> numReplicasPerShard) {
            $.numReplicasPerShard = numReplicasPerShard;
            return this;
        }

        /**
         * @param numReplicasPerShard The number of replicas to apply to each shard, up to a maximum of 5. Defaults to `1` (i.e. 2 nodes per shard).
         * 
         * @return builder
         * 
         */
        public Builder numReplicasPerShard(Integer numReplicasPerShard) {
            return numReplicasPerShard(Output.of(numReplicasPerShard));
        }

        /**
         * @param numShards The number of shards in the cluster. Defaults to `1`.
         * 
         * @return builder
         * 
         */
        public Builder numShards(@Nullable Output<Integer> numShards) {
            $.numShards = numShards;
            return this;
        }

        /**
         * @param numShards The number of shards in the cluster. Defaults to `1`.
         * 
         * @return builder
         * 
         */
        public Builder numShards(Integer numShards) {
            return numShards(Output.of(numShards));
        }

        /**
         * @param parameterGroupName The name of the parameter group associated with the cluster.
         * 
         * @return builder
         * 
         */
        public Builder parameterGroupName(@Nullable Output<String> parameterGroupName) {
            $.parameterGroupName = parameterGroupName;
            return this;
        }

        /**
         * @param parameterGroupName The name of the parameter group associated with the cluster.
         * 
         * @return builder
         * 
         */
        public Builder parameterGroupName(String parameterGroupName) {
            return parameterGroupName(Output.of(parameterGroupName));
        }

        /**
         * @param port The port number on which each of the nodes accepts connections. Defaults to `6379`.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The port number on which each of the nodes accepts connections. Defaults to `6379`.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param securityGroupIds Set of VPC Security Group ID-s to associate with this cluster.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(@Nullable Output<List<String>> securityGroupIds) {
            $.securityGroupIds = securityGroupIds;
            return this;
        }

        /**
         * @param securityGroupIds Set of VPC Security Group ID-s to associate with this cluster.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(List<String> securityGroupIds) {
            return securityGroupIds(Output.of(securityGroupIds));
        }

        /**
         * @param securityGroupIds Set of VPC Security Group ID-s to associate with this cluster.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(String... securityGroupIds) {
            return securityGroupIds(List.of(securityGroupIds));
        }

        /**
         * @param shards Set of shards in this cluster.
         * 
         * @return builder
         * 
         */
        public Builder shards(@Nullable Output<List<ClusterShardArgs>> shards) {
            $.shards = shards;
            return this;
        }

        /**
         * @param shards Set of shards in this cluster.
         * 
         * @return builder
         * 
         */
        public Builder shards(List<ClusterShardArgs> shards) {
            return shards(Output.of(shards));
        }

        /**
         * @param shards Set of shards in this cluster.
         * 
         * @return builder
         * 
         */
        public Builder shards(ClusterShardArgs... shards) {
            return shards(List.of(shards));
        }

        /**
         * @param snapshotArns List of ARN-s that uniquely identify RDB snapshot files stored in S3. The snapshot files will be used to populate the new cluster. Object names in the ARN-s cannot contain any commas.
         * 
         * @return builder
         * 
         */
        public Builder snapshotArns(@Nullable Output<List<String>> snapshotArns) {
            $.snapshotArns = snapshotArns;
            return this;
        }

        /**
         * @param snapshotArns List of ARN-s that uniquely identify RDB snapshot files stored in S3. The snapshot files will be used to populate the new cluster. Object names in the ARN-s cannot contain any commas.
         * 
         * @return builder
         * 
         */
        public Builder snapshotArns(List<String> snapshotArns) {
            return snapshotArns(Output.of(snapshotArns));
        }

        /**
         * @param snapshotArns List of ARN-s that uniquely identify RDB snapshot files stored in S3. The snapshot files will be used to populate the new cluster. Object names in the ARN-s cannot contain any commas.
         * 
         * @return builder
         * 
         */
        public Builder snapshotArns(String... snapshotArns) {
            return snapshotArns(List.of(snapshotArns));
        }

        /**
         * @param snapshotName The name of a snapshot from which to restore data into the new cluster.
         * 
         * @return builder
         * 
         */
        public Builder snapshotName(@Nullable Output<String> snapshotName) {
            $.snapshotName = snapshotName;
            return this;
        }

        /**
         * @param snapshotName The name of a snapshot from which to restore data into the new cluster.
         * 
         * @return builder
         * 
         */
        public Builder snapshotName(String snapshotName) {
            return snapshotName(Output.of(snapshotName));
        }

        /**
         * @param snapshotRetentionLimit The number of days for which MemoryDB retains automatic snapshots before deleting them. When set to `0`, automatic backups are disabled. Defaults to `0`.
         * 
         * @return builder
         * 
         */
        public Builder snapshotRetentionLimit(@Nullable Output<Integer> snapshotRetentionLimit) {
            $.snapshotRetentionLimit = snapshotRetentionLimit;
            return this;
        }

        /**
         * @param snapshotRetentionLimit The number of days for which MemoryDB retains automatic snapshots before deleting them. When set to `0`, automatic backups are disabled. Defaults to `0`.
         * 
         * @return builder
         * 
         */
        public Builder snapshotRetentionLimit(Integer snapshotRetentionLimit) {
            return snapshotRetentionLimit(Output.of(snapshotRetentionLimit));
        }

        /**
         * @param snapshotWindow The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your shard. Example: `05:00-09:00`.
         * 
         * @return builder
         * 
         */
        public Builder snapshotWindow(@Nullable Output<String> snapshotWindow) {
            $.snapshotWindow = snapshotWindow;
            return this;
        }

        /**
         * @param snapshotWindow The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your shard. Example: `05:00-09:00`.
         * 
         * @return builder
         * 
         */
        public Builder snapshotWindow(String snapshotWindow) {
            return snapshotWindow(Output.of(snapshotWindow));
        }

        /**
         * @param snsTopicArn ARN of the SNS topic to which cluster notifications are sent.
         * 
         * @return builder
         * 
         */
        public Builder snsTopicArn(@Nullable Output<String> snsTopicArn) {
            $.snsTopicArn = snsTopicArn;
            return this;
        }

        /**
         * @param snsTopicArn ARN of the SNS topic to which cluster notifications are sent.
         * 
         * @return builder
         * 
         */
        public Builder snsTopicArn(String snsTopicArn) {
            return snsTopicArn(Output.of(snsTopicArn));
        }

        /**
         * @param subnetGroupName The name of the subnet group to be used for the cluster. Defaults to a subnet group consisting of default VPC subnets.
         * 
         * @return builder
         * 
         */
        public Builder subnetGroupName(@Nullable Output<String> subnetGroupName) {
            $.subnetGroupName = subnetGroupName;
            return this;
        }

        /**
         * @param subnetGroupName The name of the subnet group to be used for the cluster. Defaults to a subnet group consisting of default VPC subnets.
         * 
         * @return builder
         * 
         */
        public Builder subnetGroupName(String subnetGroupName) {
            return subnetGroupName(Output.of(subnetGroupName));
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param tlsEnabled A flag to enable in-transit encryption on the cluster. When set to `false`, the `acl_name` must be `open-access`. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder tlsEnabled(@Nullable Output<Boolean> tlsEnabled) {
            $.tlsEnabled = tlsEnabled;
            return this;
        }

        /**
         * @param tlsEnabled A flag to enable in-transit encryption on the cluster. When set to `false`, the `acl_name` must be `open-access`. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder tlsEnabled(Boolean tlsEnabled) {
            return tlsEnabled(Output.of(tlsEnabled));
        }

        public ClusterState build() {
            return $;
        }
    }

}
