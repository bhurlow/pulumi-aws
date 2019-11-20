// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElastiCache
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get information about an Elasticache Cluster
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elasticache_cluster.html.markdown.
        /// </summary>
        public static Task<GetClusterResult> GetCluster(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("aws:elasticache/getCluster:getCluster", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Group identifier.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GetClusterArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetClusterResult
    {
        public readonly string Arn;
        /// <summary>
        /// The Availability Zone for the cache cluster.
        /// </summary>
        public readonly string AvailabilityZone;
        /// <summary>
        /// List of node objects including `id`, `address`, `port` and `availability_zone`.
        /// Referenceable e.g. as `${data.aws_elasticache_cluster.bar.cache_nodes.0.address}`
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClusterCacheNodesResult> CacheNodes;
        /// <summary>
        /// (Memcached only) The DNS name of the cache cluster without the port appended.
        /// </summary>
        public readonly string ClusterAddress;
        public readonly string ClusterId;
        /// <summary>
        /// (Memcached only) The configuration endpoint to allow host discovery.
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
        /// Specifies the weekly time range for when maintenance
        /// on the cache cluster is performed.
        /// </summary>
        public readonly string MaintenanceWindow;
        /// <summary>
        /// The cluster node type.
        /// </summary>
        public readonly string NodeType;
        /// <summary>
        /// An Amazon Resource Name (ARN) of an
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
        /// The replication group to which this cache cluster belongs.
        /// </summary>
        public readonly string ReplicationGroupId;
        /// <summary>
        /// List VPC security groups associated with the cache cluster.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// List of security group names associated with this cache cluster.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupNames;
        /// <summary>
        /// The number of days for which ElastiCache will
        /// retain automatic cache cluster snapshots before deleting them.
        /// </summary>
        public readonly int SnapshotRetentionLimit;
        /// <summary>
        /// The daily time range (in UTC) during which ElastiCache will
        /// begin taking a daily snapshot of the cache cluster.
        /// </summary>
        public readonly string SnapshotWindow;
        /// <summary>
        /// Name of the subnet group associated to the cache cluster.
        /// </summary>
        public readonly string SubnetGroupName;
        /// <summary>
        /// The tags assigned to the resource
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetClusterResult(
            string arn,
            string availabilityZone,
            ImmutableArray<Outputs.GetClusterCacheNodesResult> cacheNodes,
            string clusterAddress,
            string clusterId,
            string configurationEndpoint,
            string engine,
            string engineVersion,
            string maintenanceWindow,
            string nodeType,
            string notificationTopicArn,
            int numCacheNodes,
            string parameterGroupName,
            int port,
            string replicationGroupId,
            ImmutableArray<string> securityGroupIds,
            ImmutableArray<string> securityGroupNames,
            int snapshotRetentionLimit,
            string snapshotWindow,
            string subnetGroupName,
            ImmutableDictionary<string, object> tags,
            string id)
        {
            Arn = arn;
            AvailabilityZone = availabilityZone;
            CacheNodes = cacheNodes;
            ClusterAddress = clusterAddress;
            ClusterId = clusterId;
            ConfigurationEndpoint = configurationEndpoint;
            Engine = engine;
            EngineVersion = engineVersion;
            MaintenanceWindow = maintenanceWindow;
            NodeType = nodeType;
            NotificationTopicArn = notificationTopicArn;
            NumCacheNodes = numCacheNodes;
            ParameterGroupName = parameterGroupName;
            Port = port;
            ReplicationGroupId = replicationGroupId;
            SecurityGroupIds = securityGroupIds;
            SecurityGroupNames = securityGroupNames;
            SnapshotRetentionLimit = snapshotRetentionLimit;
            SnapshotWindow = snapshotWindow;
            SubnetGroupName = subnetGroupName;
            Tags = tags;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetClusterCacheNodesResult
    {
        public readonly string Address;
        /// <summary>
        /// The Availability Zone for the cache cluster.
        /// </summary>
        public readonly string AvailabilityZone;
        public readonly string Id;
        /// <summary>
        /// The port number on which each of the cache nodes will
        /// accept connections.
        /// </summary>
        public readonly int Port;

        [OutputConstructor]
        private GetClusterCacheNodesResult(
            string address,
            string availabilityZone,
            string id,
            int port)
        {
            Address = address;
            AvailabilityZone = availabilityZone;
            Id = id;
            Port = port;
        }
    }
    }
}
