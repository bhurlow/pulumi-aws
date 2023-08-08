// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Msk
{
    /// <summary>
    /// Manages an Amazon MSK cluster.
    /// 
    /// &gt; **Note:** This resource manages _provisioned_ clusters. To manage a _serverless_ Amazon MSK cluster, use the `aws.msk.ServerlessCluster` resource.
    /// 
    /// ## Example Usage
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var vpc = new Aws.Ec2.Vpc("vpc", new()
    ///     {
    ///         CidrBlock = "192.168.0.0/22",
    ///     });
    /// 
    ///     var azs = Aws.GetAvailabilityZones.Invoke(new()
    ///     {
    ///         State = "available",
    ///     });
    /// 
    ///     var subnetAz1 = new Aws.Ec2.Subnet("subnetAz1", new()
    ///     {
    ///         AvailabilityZone = azs.Apply(getAvailabilityZonesResult =&gt; getAvailabilityZonesResult.Names[0]),
    ///         CidrBlock = "192.168.0.0/24",
    ///         VpcId = vpc.Id,
    ///     });
    /// 
    ///     var subnetAz2 = new Aws.Ec2.Subnet("subnetAz2", new()
    ///     {
    ///         AvailabilityZone = azs.Apply(getAvailabilityZonesResult =&gt; getAvailabilityZonesResult.Names[1]),
    ///         CidrBlock = "192.168.1.0/24",
    ///         VpcId = vpc.Id,
    ///     });
    /// 
    ///     var subnetAz3 = new Aws.Ec2.Subnet("subnetAz3", new()
    ///     {
    ///         AvailabilityZone = azs.Apply(getAvailabilityZonesResult =&gt; getAvailabilityZonesResult.Names[2]),
    ///         CidrBlock = "192.168.2.0/24",
    ///         VpcId = vpc.Id,
    ///     });
    /// 
    ///     var sg = new Aws.Ec2.SecurityGroup("sg", new()
    ///     {
    ///         VpcId = vpc.Id,
    ///     });
    /// 
    ///     var kms = new Aws.Kms.Key("kms", new()
    ///     {
    ///         Description = "example",
    ///     });
    /// 
    ///     var test = new Aws.CloudWatch.LogGroup("test");
    /// 
    ///     var bucket = new Aws.S3.BucketV2("bucket");
    /// 
    ///     var bucketAcl = new Aws.S3.BucketAclV2("bucketAcl", new()
    ///     {
    ///         Bucket = bucket.Id,
    ///         Acl = "private",
    ///     });
    /// 
    ///     var assumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
    ///     {
    ///         Statements = new[]
    ///         {
    ///             new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
    ///             {
    ///                 Effect = "Allow",
    ///                 Principals = new[]
    ///                 {
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
    ///                     {
    ///                         Type = "Service",
    ///                         Identifiers = new[]
    ///                         {
    ///                             "firehose.amazonaws.com",
    ///                         },
    ///                     },
    ///                 },
    ///                 Actions = new[]
    ///                 {
    ///                     "sts:AssumeRole",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var firehoseRole = new Aws.Iam.Role("firehoseRole", new()
    ///     {
    ///         AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult =&gt; getPolicyDocumentResult.Json),
    ///     });
    /// 
    ///     var testStream = new Aws.Kinesis.FirehoseDeliveryStream("testStream", new()
    ///     {
    ///         Destination = "extended_s3",
    ///         ExtendedS3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationArgs
    ///         {
    ///             RoleArn = firehoseRole.Arn,
    ///             BucketArn = bucket.Arn,
    ///         },
    ///         Tags = 
    ///         {
    ///             { "LogDeliveryEnabled", "placeholder" },
    ///         },
    ///     });
    /// 
    ///     var example = new Aws.Msk.Cluster("example", new()
    ///     {
    ///         KafkaVersion = "3.2.0",
    ///         NumberOfBrokerNodes = 3,
    ///         BrokerNodeGroupInfo = new Aws.Msk.Inputs.ClusterBrokerNodeGroupInfoArgs
    ///         {
    ///             InstanceType = "kafka.m5.large",
    ///             ClientSubnets = new[]
    ///             {
    ///                 subnetAz1.Id,
    ///                 subnetAz2.Id,
    ///                 subnetAz3.Id,
    ///             },
    ///             StorageInfo = new Aws.Msk.Inputs.ClusterBrokerNodeGroupInfoStorageInfoArgs
    ///             {
    ///                 EbsStorageInfo = new Aws.Msk.Inputs.ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs
    ///                 {
    ///                     VolumeSize = 1000,
    ///                 },
    ///             },
    ///             SecurityGroups = new[]
    ///             {
    ///                 sg.Id,
    ///             },
    ///         },
    ///         EncryptionInfo = new Aws.Msk.Inputs.ClusterEncryptionInfoArgs
    ///         {
    ///             EncryptionAtRestKmsKeyArn = kms.Arn,
    ///         },
    ///         OpenMonitoring = new Aws.Msk.Inputs.ClusterOpenMonitoringArgs
    ///         {
    ///             Prometheus = new Aws.Msk.Inputs.ClusterOpenMonitoringPrometheusArgs
    ///             {
    ///                 JmxExporter = new Aws.Msk.Inputs.ClusterOpenMonitoringPrometheusJmxExporterArgs
    ///                 {
    ///                     EnabledInBroker = true,
    ///                 },
    ///                 NodeExporter = new Aws.Msk.Inputs.ClusterOpenMonitoringPrometheusNodeExporterArgs
    ///                 {
    ///                     EnabledInBroker = true,
    ///                 },
    ///             },
    ///         },
    ///         LoggingInfo = new Aws.Msk.Inputs.ClusterLoggingInfoArgs
    ///         {
    ///             BrokerLogs = new Aws.Msk.Inputs.ClusterLoggingInfoBrokerLogsArgs
    ///             {
    ///                 CloudwatchLogs = new Aws.Msk.Inputs.ClusterLoggingInfoBrokerLogsCloudwatchLogsArgs
    ///                 {
    ///                     Enabled = true,
    ///                     LogGroup = test.Name,
    ///                 },
    ///                 Firehose = new Aws.Msk.Inputs.ClusterLoggingInfoBrokerLogsFirehoseArgs
    ///                 {
    ///                     Enabled = true,
    ///                     DeliveryStream = testStream.Name,
    ///                 },
    ///                 S3 = new Aws.Msk.Inputs.ClusterLoggingInfoBrokerLogsS3Args
    ///                 {
    ///                     Enabled = true,
    ///                     Bucket = bucket.Id,
    ///                     Prefix = "logs/msk-",
    ///                 },
    ///             },
    ///         },
    ///         Tags = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["zookeeperConnectString"] = example.ZookeeperConnectString,
    ///         ["bootstrapBrokersTls"] = example.BootstrapBrokersTls,
    ///     };
    /// });
    /// ```
    /// ### With volume_throughput argument
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.Msk.Cluster("example", new()
    ///     {
    ///         KafkaVersion = "2.7.1",
    ///         NumberOfBrokerNodes = 3,
    ///         BrokerNodeGroupInfo = new Aws.Msk.Inputs.ClusterBrokerNodeGroupInfoArgs
    ///         {
    ///             InstanceType = "kafka.m5.4xlarge",
    ///             ClientSubnets = new[]
    ///             {
    ///                 aws_subnet.Subnet_az1.Id,
    ///                 aws_subnet.Subnet_az2.Id,
    ///                 aws_subnet.Subnet_az3.Id,
    ///             },
    ///             StorageInfo = new Aws.Msk.Inputs.ClusterBrokerNodeGroupInfoStorageInfoArgs
    ///             {
    ///                 EbsStorageInfo = new Aws.Msk.Inputs.ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs
    ///                 {
    ///                     ProvisionedThroughput = new Aws.Msk.Inputs.ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughputArgs
    ///                     {
    ///                         Enabled = true,
    ///                         VolumeThroughput = 250,
    ///                     },
    ///                     VolumeSize = 1000,
    ///                 },
    ///             },
    ///             SecurityGroups = new[]
    ///             {
    ///                 aws_security_group.Sg.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_msk_cluster.example
    /// 
    ///  id = "arn:aws:kafka:us-west-2:123456789012:cluster/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3" } Using `pulumi import`, import MSK clusters using the cluster `arn`. For exampleconsole % pulumi import aws_msk_cluster.example arn:aws:kafka:us-west-2:123456789012:cluster/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3
    /// </summary>
    [AwsResourceType("aws:msk/cluster:Cluster")]
    public partial class Cluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Comma separated list of one or more hostname:port pairs of kafka brokers suitable to bootstrap connectivity to the kafka cluster. Contains a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `PLAINTEXT` or `TLS_PLAINTEXT`. The resource sorts values alphabetically. AWS may not always return all endpoints so this value is not guaranteed to be stable across applies.
        /// </summary>
        [Output("bootstrapBrokers")]
        public Output<string> BootstrapBrokers { get; private set; } = null!;

        /// <summary>
        /// One or more DNS names (or IP addresses) and SASL IAM port pairs. For example, `b-1-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9198,b-2-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9198,b-3-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9198`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.iam` is set to `true` and `broker_node_group_info.0.connectivity_info.0.public_access.0.type` is set to `SERVICE_PROVIDED_EIPS` and the cluster fulfill all other requirements for public access. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
        /// </summary>
        [Output("bootstrapBrokersPublicSaslIam")]
        public Output<string> BootstrapBrokersPublicSaslIam { get; private set; } = null!;

        /// <summary>
        /// One or more DNS names (or IP addresses) and SASL SCRAM port pairs. For example, `b-1-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9196,b-2-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9196,b-3-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9196`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.scram` is set to `true` and `broker_node_group_info.0.connectivity_info.0.public_access.0.type` is set to `SERVICE_PROVIDED_EIPS` and the cluster fulfill all other requirements for public access. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
        /// </summary>
        [Output("bootstrapBrokersPublicSaslScram")]
        public Output<string> BootstrapBrokersPublicSaslScram { get; private set; } = null!;

        /// <summary>
        /// One or more DNS names (or IP addresses) and TLS port pairs. For example, `b-1-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9194,b-2-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9194,b-3-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9194`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `broker_node_group_info.0.connectivity_info.0.public_access.0.type` is set to `SERVICE_PROVIDED_EIPS` and the cluster fulfill all other requirements for public access. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
        /// </summary>
        [Output("bootstrapBrokersPublicTls")]
        public Output<string> BootstrapBrokersPublicTls { get; private set; } = null!;

        /// <summary>
        /// One or more DNS names (or IP addresses) and SASL IAM port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.iam` is set to `true`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
        /// </summary>
        [Output("bootstrapBrokersSaslIam")]
        public Output<string> BootstrapBrokersSaslIam { get; private set; } = null!;

        /// <summary>
        /// One or more DNS names (or IP addresses) and SASL SCRAM port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.scram` is set to `true`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
        /// </summary>
        [Output("bootstrapBrokersSaslScram")]
        public Output<string> BootstrapBrokersSaslScram { get; private set; } = null!;

        /// <summary>
        /// One or more DNS names (or IP addresses) and TLS port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
        /// </summary>
        [Output("bootstrapBrokersTls")]
        public Output<string> BootstrapBrokersTls { get; private set; } = null!;

        /// <summary>
        /// Configuration block for the broker nodes of the Kafka cluster.
        /// </summary>
        [Output("brokerNodeGroupInfo")]
        public Output<Outputs.ClusterBrokerNodeGroupInfo> BrokerNodeGroupInfo { get; private set; } = null!;

        /// <summary>
        /// Configuration block for specifying a client authentication. See below.
        /// </summary>
        [Output("clientAuthentication")]
        public Output<Outputs.ClusterClientAuthentication?> ClientAuthentication { get; private set; } = null!;

        /// <summary>
        /// Name of the MSK cluster.
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
        /// </summary>
        [Output("configurationInfo")]
        public Output<Outputs.ClusterConfigurationInfo?> ConfigurationInfo { get; private set; } = null!;

        /// <summary>
        /// Current version of the MSK Cluster used for updates, e.g., `K13V1IB3VIYZZH`
        /// * `encryption_info.0.encryption_at_rest_kms_key_arn` - The ARN of the KMS key used for encryption at rest of the broker data volumes.
        /// </summary>
        [Output("currentVersion")]
        public Output<string> CurrentVersion { get; private set; } = null!;

        /// <summary>
        /// Configuration block for specifying encryption. See below.
        /// </summary>
        [Output("encryptionInfo")]
        public Output<Outputs.ClusterEncryptionInfo?> EncryptionInfo { get; private set; } = null!;

        /// <summary>
        /// Specify the desired enhanced MSK CloudWatch monitoring level. See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
        /// </summary>
        [Output("enhancedMonitoring")]
        public Output<string?> EnhancedMonitoring { get; private set; } = null!;

        /// <summary>
        /// Specify the desired Kafka software version.
        /// </summary>
        [Output("kafkaVersion")]
        public Output<string> KafkaVersion { get; private set; } = null!;

        /// <summary>
        /// Configuration block for streaming broker logs to Cloudwatch/S3/Kinesis Firehose. See below.
        /// </summary>
        [Output("loggingInfo")]
        public Output<Outputs.ClusterLoggingInfo?> LoggingInfo { get; private set; } = null!;

        /// <summary>
        /// The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
        /// </summary>
        [Output("numberOfBrokerNodes")]
        public Output<int> NumberOfBrokerNodes { get; private set; } = null!;

        /// <summary>
        /// Configuration block for JMX and Node monitoring for the MSK cluster. See below.
        /// </summary>
        [Output("openMonitoring")]
        public Output<Outputs.ClusterOpenMonitoring?> OpenMonitoring { get; private set; } = null!;

        /// <summary>
        /// Controls storage mode for supported storage tiers. Valid values are: `LOCAL` or `TIERED`.
        /// </summary>
        [Output("storageMode")]
        public Output<string> StorageMode { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster. The returned values are sorted alphabetically. The AWS API may not return all endpoints, so this value is not guaranteed to be stable across applies.
        /// </summary>
        [Output("zookeeperConnectString")]
        public Output<string> ZookeeperConnectString { get; private set; } = null!;

        /// <summary>
        /// A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster via TLS. The returned values are sorted alphabetically. The AWS API may not return all endpoints, so this value is not guaranteed to be stable across applies.
        /// </summary>
        [Output("zookeeperConnectStringTls")]
        public Output<string> ZookeeperConnectStringTls { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("aws:msk/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("aws:msk/cluster:Cluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, state, options);
        }
    }

    public sealed class ClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration block for the broker nodes of the Kafka cluster.
        /// </summary>
        [Input("brokerNodeGroupInfo", required: true)]
        public Input<Inputs.ClusterBrokerNodeGroupInfoArgs> BrokerNodeGroupInfo { get; set; } = null!;

        /// <summary>
        /// Configuration block for specifying a client authentication. See below.
        /// </summary>
        [Input("clientAuthentication")]
        public Input<Inputs.ClusterClientAuthenticationArgs>? ClientAuthentication { get; set; }

        /// <summary>
        /// Name of the MSK cluster.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
        /// </summary>
        [Input("configurationInfo")]
        public Input<Inputs.ClusterConfigurationInfoArgs>? ConfigurationInfo { get; set; }

        /// <summary>
        /// Configuration block for specifying encryption. See below.
        /// </summary>
        [Input("encryptionInfo")]
        public Input<Inputs.ClusterEncryptionInfoArgs>? EncryptionInfo { get; set; }

        /// <summary>
        /// Specify the desired enhanced MSK CloudWatch monitoring level. See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
        /// </summary>
        [Input("enhancedMonitoring")]
        public Input<string>? EnhancedMonitoring { get; set; }

        /// <summary>
        /// Specify the desired Kafka software version.
        /// </summary>
        [Input("kafkaVersion", required: true)]
        public Input<string> KafkaVersion { get; set; } = null!;

        /// <summary>
        /// Configuration block for streaming broker logs to Cloudwatch/S3/Kinesis Firehose. See below.
        /// </summary>
        [Input("loggingInfo")]
        public Input<Inputs.ClusterLoggingInfoArgs>? LoggingInfo { get; set; }

        /// <summary>
        /// The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
        /// </summary>
        [Input("numberOfBrokerNodes", required: true)]
        public Input<int> NumberOfBrokerNodes { get; set; } = null!;

        /// <summary>
        /// Configuration block for JMX and Node monitoring for the MSK cluster. See below.
        /// </summary>
        [Input("openMonitoring")]
        public Input<Inputs.ClusterOpenMonitoringArgs>? OpenMonitoring { get; set; }

        /// <summary>
        /// Controls storage mode for supported storage tiers. Valid values are: `LOCAL` or `TIERED`.
        /// </summary>
        [Input("storageMode")]
        public Input<string>? StorageMode { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ClusterArgs()
        {
        }
        public static new ClusterArgs Empty => new ClusterArgs();
    }

    public sealed class ClusterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Comma separated list of one or more hostname:port pairs of kafka brokers suitable to bootstrap connectivity to the kafka cluster. Contains a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `PLAINTEXT` or `TLS_PLAINTEXT`. The resource sorts values alphabetically. AWS may not always return all endpoints so this value is not guaranteed to be stable across applies.
        /// </summary>
        [Input("bootstrapBrokers")]
        public Input<string>? BootstrapBrokers { get; set; }

        /// <summary>
        /// One or more DNS names (or IP addresses) and SASL IAM port pairs. For example, `b-1-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9198,b-2-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9198,b-3-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9198`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.iam` is set to `true` and `broker_node_group_info.0.connectivity_info.0.public_access.0.type` is set to `SERVICE_PROVIDED_EIPS` and the cluster fulfill all other requirements for public access. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
        /// </summary>
        [Input("bootstrapBrokersPublicSaslIam")]
        public Input<string>? BootstrapBrokersPublicSaslIam { get; set; }

        /// <summary>
        /// One or more DNS names (or IP addresses) and SASL SCRAM port pairs. For example, `b-1-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9196,b-2-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9196,b-3-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9196`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.scram` is set to `true` and `broker_node_group_info.0.connectivity_info.0.public_access.0.type` is set to `SERVICE_PROVIDED_EIPS` and the cluster fulfill all other requirements for public access. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
        /// </summary>
        [Input("bootstrapBrokersPublicSaslScram")]
        public Input<string>? BootstrapBrokersPublicSaslScram { get; set; }

        /// <summary>
        /// One or more DNS names (or IP addresses) and TLS port pairs. For example, `b-1-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9194,b-2-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9194,b-3-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9194`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `broker_node_group_info.0.connectivity_info.0.public_access.0.type` is set to `SERVICE_PROVIDED_EIPS` and the cluster fulfill all other requirements for public access. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
        /// </summary>
        [Input("bootstrapBrokersPublicTls")]
        public Input<string>? BootstrapBrokersPublicTls { get; set; }

        /// <summary>
        /// One or more DNS names (or IP addresses) and SASL IAM port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.iam` is set to `true`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
        /// </summary>
        [Input("bootstrapBrokersSaslIam")]
        public Input<string>? BootstrapBrokersSaslIam { get; set; }

        /// <summary>
        /// One or more DNS names (or IP addresses) and SASL SCRAM port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.scram` is set to `true`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
        /// </summary>
        [Input("bootstrapBrokersSaslScram")]
        public Input<string>? BootstrapBrokersSaslScram { get; set; }

        /// <summary>
        /// One or more DNS names (or IP addresses) and TLS port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
        /// </summary>
        [Input("bootstrapBrokersTls")]
        public Input<string>? BootstrapBrokersTls { get; set; }

        /// <summary>
        /// Configuration block for the broker nodes of the Kafka cluster.
        /// </summary>
        [Input("brokerNodeGroupInfo")]
        public Input<Inputs.ClusterBrokerNodeGroupInfoGetArgs>? BrokerNodeGroupInfo { get; set; }

        /// <summary>
        /// Configuration block for specifying a client authentication. See below.
        /// </summary>
        [Input("clientAuthentication")]
        public Input<Inputs.ClusterClientAuthenticationGetArgs>? ClientAuthentication { get; set; }

        /// <summary>
        /// Name of the MSK cluster.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
        /// </summary>
        [Input("configurationInfo")]
        public Input<Inputs.ClusterConfigurationInfoGetArgs>? ConfigurationInfo { get; set; }

        /// <summary>
        /// Current version of the MSK Cluster used for updates, e.g., `K13V1IB3VIYZZH`
        /// * `encryption_info.0.encryption_at_rest_kms_key_arn` - The ARN of the KMS key used for encryption at rest of the broker data volumes.
        /// </summary>
        [Input("currentVersion")]
        public Input<string>? CurrentVersion { get; set; }

        /// <summary>
        /// Configuration block for specifying encryption. See below.
        /// </summary>
        [Input("encryptionInfo")]
        public Input<Inputs.ClusterEncryptionInfoGetArgs>? EncryptionInfo { get; set; }

        /// <summary>
        /// Specify the desired enhanced MSK CloudWatch monitoring level. See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
        /// </summary>
        [Input("enhancedMonitoring")]
        public Input<string>? EnhancedMonitoring { get; set; }

        /// <summary>
        /// Specify the desired Kafka software version.
        /// </summary>
        [Input("kafkaVersion")]
        public Input<string>? KafkaVersion { get; set; }

        /// <summary>
        /// Configuration block for streaming broker logs to Cloudwatch/S3/Kinesis Firehose. See below.
        /// </summary>
        [Input("loggingInfo")]
        public Input<Inputs.ClusterLoggingInfoGetArgs>? LoggingInfo { get; set; }

        /// <summary>
        /// The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
        /// </summary>
        [Input("numberOfBrokerNodes")]
        public Input<int>? NumberOfBrokerNodes { get; set; }

        /// <summary>
        /// Configuration block for JMX and Node monitoring for the MSK cluster. See below.
        /// </summary>
        [Input("openMonitoring")]
        public Input<Inputs.ClusterOpenMonitoringGetArgs>? OpenMonitoring { get; set; }

        /// <summary>
        /// Controls storage mode for supported storage tiers. Valid values are: `LOCAL` or `TIERED`.
        /// </summary>
        [Input("storageMode")]
        public Input<string>? StorageMode { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster. The returned values are sorted alphabetically. The AWS API may not return all endpoints, so this value is not guaranteed to be stable across applies.
        /// </summary>
        [Input("zookeeperConnectString")]
        public Input<string>? ZookeeperConnectString { get; set; }

        /// <summary>
        /// A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster via TLS. The returned values are sorted alphabetically. The AWS API may not return all endpoints, so this value is not guaranteed to be stable across applies.
        /// </summary>
        [Input("zookeeperConnectStringTls")]
        public Input<string>? ZookeeperConnectStringTls { get; set; }

        public ClusterState()
        {
        }
        public static new ClusterState Empty => new ClusterState();
    }
}
