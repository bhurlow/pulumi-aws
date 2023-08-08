// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElastiCache
{
    public static class GetCluster
    {
        /// <summary>
        /// Use this data source to get information about an ElastiCache Cluster
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myCluster = Aws.ElastiCache.GetCluster.Invoke(new()
        ///     {
        ///         ClusterId = "my-cluster-id",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("aws:elasticache/getCluster:getCluster", args ?? new GetClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about an ElastiCache Cluster
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myCluster = Aws.ElastiCache.GetCluster.Invoke(new()
        ///     {
        ///         ClusterId = "my-cluster-id",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterResult>("aws:elasticache/getCluster:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Group identifier.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// Tags assigned to the resource
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetClusterArgs()
        {
        }
        public static new GetClusterArgs Empty => new GetClusterArgs();
    }

    public sealed class GetClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Group identifier.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Tags assigned to the resource
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetClusterInvokeArgs()
        {
        }
        public static new GetClusterInvokeArgs Empty => new GetClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        public readonly string Arn;
        /// <summary>
        /// Availability Zone for the cache cluster.
        /// </summary>
        public readonly string AvailabilityZone;
        /// <summary>
        /// List of node objects including `id`, `address`, `port`, `availability_zone` and `outpost_arn`.
        /// Referenceable e.g., as `${data.aws_elasticache_cluster.bar.cache_nodes.0.address}`
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClusterCacheNodeResult> CacheNodes;
        /// <summary>
        /// (Memcached only) DNS name of the cache cluster without the port appended.
        /// </summary>
        public readonly string ClusterAddress;
        public readonly string ClusterId;
        /// <summary>
        /// (Memcached only) Configuration endpoint to allow host discovery.
        /// </summary>
        public readonly string ConfigurationEndpoint;
        /// <summary>
        /// Name of the cache engine.
        /// </summary>
        public readonly string Engine;
        /// <summary>
        /// Version number of the cache engine.
        /// </summary>
        public readonly string EngineVersion;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The IP version advertised in the discovery protocol.
        /// </summary>
        public readonly string IpDiscovery;
        /// <summary>
        /// Redis [SLOWLOG](https://redis.io/commands/slowlog) or Redis [Engine Log](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Log_Delivery.html#Log_contents-engine-log) delivery settings.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClusterLogDeliveryConfigurationResult> LogDeliveryConfigurations;
        /// <summary>
        /// Specifies the weekly time range for when maintenance
        /// on the cache cluster is performed.
        /// </summary>
        public readonly string MaintenanceWindow;
        /// <summary>
        /// The IP versions for cache cluster connections.
        /// </summary>
        public readonly string NetworkType;
        /// <summary>
        /// The cluster node type.
        /// </summary>
        public readonly string NodeType;
        /// <summary>
        /// An ARN of an
        /// SNS topic that ElastiCache notifications get sent to.
        /// </summary>
        public readonly string NotificationTopicArn;
        /// <summary>
        /// The number of cache nodes that the cache cluster has.
        /// </summary>
        public readonly int NumCacheNodes;
        /// <summary>
        /// Name of the parameter group associated with this cache cluster.
        /// </summary>
        public readonly string ParameterGroupName;
        /// <summary>
        /// The port number on which each of the cache nodes will
        /// accept connections.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The outpost ARN in which the cache cluster was created if created in outpost.
        /// </summary>
        public readonly string PreferredOutpostArn;
        /// <summary>
        /// The replication group to which this cache cluster belongs.
        /// </summary>
        public readonly string ReplicationGroupId;
        /// <summary>
        /// List VPC security groups associated with the cache cluster.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// The number of days for which ElastiCache will
        /// retain automatic cache cluster snapshots before deleting them.
        /// </summary>
        public readonly int SnapshotRetentionLimit;
        /// <summary>
        /// Daily time range (in UTC) during which ElastiCache will
        /// begin taking a daily snapshot of the cache cluster.
        /// </summary>
        public readonly string SnapshotWindow;
        /// <summary>
        /// Name of the subnet group associated to the cache cluster.
        /// </summary>
        public readonly string SubnetGroupName;
        /// <summary>
        /// Tags assigned to the resource
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetClusterResult(
            string arn,

            string availabilityZone,

            ImmutableArray<Outputs.GetClusterCacheNodeResult> cacheNodes,

            string clusterAddress,

            string clusterId,

            string configurationEndpoint,

            string engine,

            string engineVersion,

            string id,

            string ipDiscovery,

            ImmutableArray<Outputs.GetClusterLogDeliveryConfigurationResult> logDeliveryConfigurations,

            string maintenanceWindow,

            string networkType,

            string nodeType,

            string notificationTopicArn,

            int numCacheNodes,

            string parameterGroupName,

            int port,

            string preferredOutpostArn,

            string replicationGroupId,

            ImmutableArray<string> securityGroupIds,

            int snapshotRetentionLimit,

            string snapshotWindow,

            string subnetGroupName,

            ImmutableDictionary<string, string> tags)
        {
            Arn = arn;
            AvailabilityZone = availabilityZone;
            CacheNodes = cacheNodes;
            ClusterAddress = clusterAddress;
            ClusterId = clusterId;
            ConfigurationEndpoint = configurationEndpoint;
            Engine = engine;
            EngineVersion = engineVersion;
            Id = id;
            IpDiscovery = ipDiscovery;
            LogDeliveryConfigurations = logDeliveryConfigurations;
            MaintenanceWindow = maintenanceWindow;
            NetworkType = networkType;
            NodeType = nodeType;
            NotificationTopicArn = notificationTopicArn;
            NumCacheNodes = numCacheNodes;
            ParameterGroupName = parameterGroupName;
            Port = port;
            PreferredOutpostArn = preferredOutpostArn;
            ReplicationGroupId = replicationGroupId;
            SecurityGroupIds = securityGroupIds;
            SnapshotRetentionLimit = snapshotRetentionLimit;
            SnapshotWindow = snapshotWindow;
            SubnetGroupName = subnetGroupName;
            Tags = tags;
        }
    }
}
