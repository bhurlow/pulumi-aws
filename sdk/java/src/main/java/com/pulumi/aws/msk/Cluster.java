// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.msk;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.msk.ClusterArgs;
import com.pulumi.aws.msk.inputs.ClusterState;
import com.pulumi.aws.msk.outputs.ClusterBrokerNodeGroupInfo;
import com.pulumi.aws.msk.outputs.ClusterClientAuthentication;
import com.pulumi.aws.msk.outputs.ClusterConfigurationInfo;
import com.pulumi.aws.msk.outputs.ClusterEncryptionInfo;
import com.pulumi.aws.msk.outputs.ClusterLoggingInfo;
import com.pulumi.aws.msk.outputs.ClusterOpenMonitoring;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an Amazon MSK cluster.
 * 
 * &gt; **Note:** This resource manages _provisioned_ clusters. To manage a _serverless_ Amazon MSK cluster, use the `aws.msk.ServerlessCluster` resource.
 * 
 * ## Example Usage
 * ### Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.Vpc;
 * import com.pulumi.aws.ec2.VpcArgs;
 * import com.pulumi.aws.AwsFunctions;
 * import com.pulumi.aws.inputs.GetAvailabilityZonesArgs;
 * import com.pulumi.aws.ec2.Subnet;
 * import com.pulumi.aws.ec2.SubnetArgs;
 * import com.pulumi.aws.ec2.SecurityGroup;
 * import com.pulumi.aws.ec2.SecurityGroupArgs;
 * import com.pulumi.aws.kms.Key;
 * import com.pulumi.aws.kms.KeyArgs;
 * import com.pulumi.aws.cloudwatch.LogGroup;
 * import com.pulumi.aws.s3.BucketV2;
 * import com.pulumi.aws.s3.BucketAclV2;
 * import com.pulumi.aws.s3.BucketAclV2Args;
 * import com.pulumi.aws.iam.Role;
 * import com.pulumi.aws.iam.RoleArgs;
 * import com.pulumi.aws.kinesis.FirehoseDeliveryStream;
 * import com.pulumi.aws.kinesis.FirehoseDeliveryStreamArgs;
 * import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamS3ConfigurationArgs;
 * import com.pulumi.aws.msk.Cluster;
 * import com.pulumi.aws.msk.ClusterArgs;
 * import com.pulumi.aws.msk.inputs.ClusterBrokerNodeGroupInfoArgs;
 * import com.pulumi.aws.msk.inputs.ClusterBrokerNodeGroupInfoStorageInfoArgs;
 * import com.pulumi.aws.msk.inputs.ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs;
 * import com.pulumi.aws.msk.inputs.ClusterEncryptionInfoArgs;
 * import com.pulumi.aws.msk.inputs.ClusterOpenMonitoringArgs;
 * import com.pulumi.aws.msk.inputs.ClusterOpenMonitoringPrometheusArgs;
 * import com.pulumi.aws.msk.inputs.ClusterOpenMonitoringPrometheusJmxExporterArgs;
 * import com.pulumi.aws.msk.inputs.ClusterOpenMonitoringPrometheusNodeExporterArgs;
 * import com.pulumi.aws.msk.inputs.ClusterLoggingInfoArgs;
 * import com.pulumi.aws.msk.inputs.ClusterLoggingInfoBrokerLogsArgs;
 * import com.pulumi.aws.msk.inputs.ClusterLoggingInfoBrokerLogsCloudwatchLogsArgs;
 * import com.pulumi.aws.msk.inputs.ClusterLoggingInfoBrokerLogsFirehoseArgs;
 * import com.pulumi.aws.msk.inputs.ClusterLoggingInfoBrokerLogsS3Args;
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
 *         var vpc = new Vpc(&#34;vpc&#34;, VpcArgs.builder()        
 *             .cidrBlock(&#34;192.168.0.0/22&#34;)
 *             .build());
 * 
 *         final var azs = AwsFunctions.getAvailabilityZones(GetAvailabilityZonesArgs.builder()
 *             .state(&#34;available&#34;)
 *             .build());
 * 
 *         var subnetAz1 = new Subnet(&#34;subnetAz1&#34;, SubnetArgs.builder()        
 *             .availabilityZone(azs.applyValue(getAvailabilityZonesResult -&gt; getAvailabilityZonesResult.names()[0]))
 *             .cidrBlock(&#34;192.168.0.0/24&#34;)
 *             .vpcId(vpc.id())
 *             .build());
 * 
 *         var subnetAz2 = new Subnet(&#34;subnetAz2&#34;, SubnetArgs.builder()        
 *             .availabilityZone(azs.applyValue(getAvailabilityZonesResult -&gt; getAvailabilityZonesResult.names()[1]))
 *             .cidrBlock(&#34;192.168.1.0/24&#34;)
 *             .vpcId(vpc.id())
 *             .build());
 * 
 *         var subnetAz3 = new Subnet(&#34;subnetAz3&#34;, SubnetArgs.builder()        
 *             .availabilityZone(azs.applyValue(getAvailabilityZonesResult -&gt; getAvailabilityZonesResult.names()[2]))
 *             .cidrBlock(&#34;192.168.2.0/24&#34;)
 *             .vpcId(vpc.id())
 *             .build());
 * 
 *         var sg = new SecurityGroup(&#34;sg&#34;, SecurityGroupArgs.builder()        
 *             .vpcId(vpc.id())
 *             .build());
 * 
 *         var kms = new Key(&#34;kms&#34;, KeyArgs.builder()        
 *             .description(&#34;example&#34;)
 *             .build());
 * 
 *         var test = new LogGroup(&#34;test&#34;);
 * 
 *         var bucket = new BucketV2(&#34;bucket&#34;);
 * 
 *         var bucketAcl = new BucketAclV2(&#34;bucketAcl&#34;, BucketAclV2Args.builder()        
 *             .bucket(bucket.id())
 *             .acl(&#34;private&#34;)
 *             .build());
 * 
 *         var firehoseRole = new Role(&#34;firehoseRole&#34;, RoleArgs.builder()        
 *             .assumeRolePolicy(&#34;&#34;&#34;
 * {
 * &#34;Version&#34;: &#34;2012-10-17&#34;,
 * &#34;Statement&#34;: [
 *   {
 *     &#34;Action&#34;: &#34;sts:AssumeRole&#34;,
 *     &#34;Principal&#34;: {
 *       &#34;Service&#34;: &#34;firehose.amazonaws.com&#34;
 *     },
 *     &#34;Effect&#34;: &#34;Allow&#34;,
 *     &#34;Sid&#34;: &#34;&#34;
 *   }
 *   ]
 * }
 *             &#34;&#34;&#34;)
 *             .build());
 * 
 *         var testStream = new FirehoseDeliveryStream(&#34;testStream&#34;, FirehoseDeliveryStreamArgs.builder()        
 *             .destination(&#34;s3&#34;)
 *             .s3Configuration(FirehoseDeliveryStreamS3ConfigurationArgs.builder()
 *                 .roleArn(firehoseRole.arn())
 *                 .bucketArn(bucket.arn())
 *                 .build())
 *             .tags(Map.of(&#34;LogDeliveryEnabled&#34;, &#34;placeholder&#34;))
 *             .build());
 * 
 *         var example = new Cluster(&#34;example&#34;, ClusterArgs.builder()        
 *             .kafkaVersion(&#34;3.2.0&#34;)
 *             .numberOfBrokerNodes(3)
 *             .brokerNodeGroupInfo(ClusterBrokerNodeGroupInfoArgs.builder()
 *                 .instanceType(&#34;kafka.m5.large&#34;)
 *                 .clientSubnets(                
 *                     subnetAz1.id(),
 *                     subnetAz2.id(),
 *                     subnetAz3.id())
 *                 .storageInfo(ClusterBrokerNodeGroupInfoStorageInfoArgs.builder()
 *                     .ebsStorageInfo(ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs.builder()
 *                         .volumeSize(1000)
 *                         .build())
 *                     .build())
 *                 .securityGroups(sg.id())
 *                 .build())
 *             .encryptionInfo(ClusterEncryptionInfoArgs.builder()
 *                 .encryptionAtRestKmsKeyArn(kms.arn())
 *                 .build())
 *             .openMonitoring(ClusterOpenMonitoringArgs.builder()
 *                 .prometheus(ClusterOpenMonitoringPrometheusArgs.builder()
 *                     .jmxExporter(ClusterOpenMonitoringPrometheusJmxExporterArgs.builder()
 *                         .enabledInBroker(true)
 *                         .build())
 *                     .nodeExporter(ClusterOpenMonitoringPrometheusNodeExporterArgs.builder()
 *                         .enabledInBroker(true)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .loggingInfo(ClusterLoggingInfoArgs.builder()
 *                 .brokerLogs(ClusterLoggingInfoBrokerLogsArgs.builder()
 *                     .cloudwatchLogs(ClusterLoggingInfoBrokerLogsCloudwatchLogsArgs.builder()
 *                         .enabled(true)
 *                         .logGroup(test.name())
 *                         .build())
 *                     .firehose(ClusterLoggingInfoBrokerLogsFirehoseArgs.builder()
 *                         .enabled(true)
 *                         .deliveryStream(testStream.name())
 *                         .build())
 *                     .s3(ClusterLoggingInfoBrokerLogsS3Args.builder()
 *                         .enabled(true)
 *                         .bucket(bucket.id())
 *                         .prefix(&#34;logs/msk-&#34;)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .tags(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .build());
 * 
 *         ctx.export(&#34;zookeeperConnectString&#34;, example.zookeeperConnectString());
 *         ctx.export(&#34;bootstrapBrokersTls&#34;, example.bootstrapBrokersTls());
 *     }
 * }
 * ```
 * ### With volume_throughput argument
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.msk.Cluster;
 * import com.pulumi.aws.msk.ClusterArgs;
 * import com.pulumi.aws.msk.inputs.ClusterBrokerNodeGroupInfoArgs;
 * import com.pulumi.aws.msk.inputs.ClusterBrokerNodeGroupInfoStorageInfoArgs;
 * import com.pulumi.aws.msk.inputs.ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs;
 * import com.pulumi.aws.msk.inputs.ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughputArgs;
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
 *             .kafkaVersion(&#34;2.7.1&#34;)
 *             .numberOfBrokerNodes(3)
 *             .brokerNodeGroupInfo(ClusterBrokerNodeGroupInfoArgs.builder()
 *                 .instanceType(&#34;kafka.m5.4xlarge&#34;)
 *                 .clientSubnets(                
 *                     aws_subnet.subnet_az1().id(),
 *                     aws_subnet.subnet_az2().id(),
 *                     aws_subnet.subnet_az3().id())
 *                 .storageInfo(ClusterBrokerNodeGroupInfoStorageInfoArgs.builder()
 *                     .ebsStorageInfo(ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs.builder()
 *                         .provisionedThroughput(ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughputArgs.builder()
 *                             .enabled(true)
 *                             .volumeThroughput(250)
 *                             .build())
 *                         .volumeSize(1000)
 *                         .build())
 *                     .build())
 *                 .securityGroups(aws_security_group.sg().id())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * MSK clusters can be imported using the cluster `arn`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:msk/cluster:Cluster example arn:aws:kafka:us-west-2:123456789012:cluster/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3
 * ```
 * 
 */
@ResourceType(type="aws:msk/cluster:Cluster")
public class Cluster extends com.pulumi.resources.CustomResource {
    /**
     * Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
     * 
     */
    @Export(name="arn", type=String.class, parameters={})
    private Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Comma separated list of one or more hostname:port pairs of kafka brokers suitable to bootstrap connectivity to the kafka cluster. Contains a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `PLAINTEXT` or `TLS_PLAINTEXT`. The resource sorts values alphabetically. AWS may not always return all endpoints so this value is not guaranteed to be stable across applies.
     * 
     */
    @Export(name="bootstrapBrokers", type=String.class, parameters={})
    private Output<String> bootstrapBrokers;

    /**
     * @return Comma separated list of one or more hostname:port pairs of kafka brokers suitable to bootstrap connectivity to the kafka cluster. Contains a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `PLAINTEXT` or `TLS_PLAINTEXT`. The resource sorts values alphabetically. AWS may not always return all endpoints so this value is not guaranteed to be stable across applies.
     * 
     */
    public Output<String> bootstrapBrokers() {
        return this.bootstrapBrokers;
    }
    /**
     * One or more DNS names (or IP addresses) and SASL IAM port pairs. For example, `b-1-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9198,b-2-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9198,b-3-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9198`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.iam` is set to `true` and `broker_node_group_info.0.connectivity_info.0.public_access.0.type` is set to `SERVICE_PROVIDED_EIPS` and the cluster fulfill all other requirements for public access. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
     * 
     */
    @Export(name="bootstrapBrokersPublicSaslIam", type=String.class, parameters={})
    private Output<String> bootstrapBrokersPublicSaslIam;

    /**
     * @return One or more DNS names (or IP addresses) and SASL IAM port pairs. For example, `b-1-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9198,b-2-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9198,b-3-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9198`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.iam` is set to `true` and `broker_node_group_info.0.connectivity_info.0.public_access.0.type` is set to `SERVICE_PROVIDED_EIPS` and the cluster fulfill all other requirements for public access. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
     * 
     */
    public Output<String> bootstrapBrokersPublicSaslIam() {
        return this.bootstrapBrokersPublicSaslIam;
    }
    /**
     * One or more DNS names (or IP addresses) and SASL SCRAM port pairs. For example, `b-1-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9196,b-2-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9196,b-3-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9196`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.scram` is set to `true` and `broker_node_group_info.0.connectivity_info.0.public_access.0.type` is set to `SERVICE_PROVIDED_EIPS` and the cluster fulfill all other requirements for public access. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
     * 
     */
    @Export(name="bootstrapBrokersPublicSaslScram", type=String.class, parameters={})
    private Output<String> bootstrapBrokersPublicSaslScram;

    /**
     * @return One or more DNS names (or IP addresses) and SASL SCRAM port pairs. For example, `b-1-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9196,b-2-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9196,b-3-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9196`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.scram` is set to `true` and `broker_node_group_info.0.connectivity_info.0.public_access.0.type` is set to `SERVICE_PROVIDED_EIPS` and the cluster fulfill all other requirements for public access. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
     * 
     */
    public Output<String> bootstrapBrokersPublicSaslScram() {
        return this.bootstrapBrokersPublicSaslScram;
    }
    /**
     * One or more DNS names (or IP addresses) and TLS port pairs. For example, `b-1-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9194,b-2-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9194,b-3-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9194`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `broker_node_group_info.0.connectivity_info.0.public_access.0.type` is set to `SERVICE_PROVIDED_EIPS` and the cluster fulfill all other requirements for public access. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
     * 
     */
    @Export(name="bootstrapBrokersPublicTls", type=String.class, parameters={})
    private Output<String> bootstrapBrokersPublicTls;

    /**
     * @return One or more DNS names (or IP addresses) and TLS port pairs. For example, `b-1-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9194,b-2-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9194,b-3-public.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9194`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `broker_node_group_info.0.connectivity_info.0.public_access.0.type` is set to `SERVICE_PROVIDED_EIPS` and the cluster fulfill all other requirements for public access. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
     * 
     */
    public Output<String> bootstrapBrokersPublicTls() {
        return this.bootstrapBrokersPublicTls;
    }
    /**
     * One or more DNS names (or IP addresses) and SASL IAM port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.iam` is set to `true`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
     * 
     */
    @Export(name="bootstrapBrokersSaslIam", type=String.class, parameters={})
    private Output<String> bootstrapBrokersSaslIam;

    /**
     * @return One or more DNS names (or IP addresses) and SASL IAM port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.iam` is set to `true`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
     * 
     */
    public Output<String> bootstrapBrokersSaslIam() {
        return this.bootstrapBrokersSaslIam;
    }
    /**
     * One or more DNS names (or IP addresses) and SASL SCRAM port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.scram` is set to `true`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
     * 
     */
    @Export(name="bootstrapBrokersSaslScram", type=String.class, parameters={})
    private Output<String> bootstrapBrokersSaslScram;

    /**
     * @return One or more DNS names (or IP addresses) and SASL SCRAM port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.scram` is set to `true`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
     * 
     */
    public Output<String> bootstrapBrokersSaslScram() {
        return this.bootstrapBrokersSaslScram;
    }
    /**
     * One or more DNS names (or IP addresses) and TLS port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
     * 
     */
    @Export(name="bootstrapBrokersTls", type=String.class, parameters={})
    private Output<String> bootstrapBrokersTls;

    /**
     * @return One or more DNS names (or IP addresses) and TLS port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
     * 
     */
    public Output<String> bootstrapBrokersTls() {
        return this.bootstrapBrokersTls;
    }
    /**
     * Configuration block for the broker nodes of the Kafka cluster.
     * 
     */
    @Export(name="brokerNodeGroupInfo", type=ClusterBrokerNodeGroupInfo.class, parameters={})
    private Output<ClusterBrokerNodeGroupInfo> brokerNodeGroupInfo;

    /**
     * @return Configuration block for the broker nodes of the Kafka cluster.
     * 
     */
    public Output<ClusterBrokerNodeGroupInfo> brokerNodeGroupInfo() {
        return this.brokerNodeGroupInfo;
    }
    /**
     * Configuration block for specifying a client authentication. See below.
     * 
     */
    @Export(name="clientAuthentication", type=ClusterClientAuthentication.class, parameters={})
    private Output</* @Nullable */ ClusterClientAuthentication> clientAuthentication;

    /**
     * @return Configuration block for specifying a client authentication. See below.
     * 
     */
    public Output<Optional<ClusterClientAuthentication>> clientAuthentication() {
        return Codegen.optional(this.clientAuthentication);
    }
    /**
     * Name of the MSK cluster.
     * 
     */
    @Export(name="clusterName", type=String.class, parameters={})
    private Output<String> clusterName;

    /**
     * @return Name of the MSK cluster.
     * 
     */
    public Output<String> clusterName() {
        return this.clusterName;
    }
    /**
     * Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
     * 
     */
    @Export(name="configurationInfo", type=ClusterConfigurationInfo.class, parameters={})
    private Output</* @Nullable */ ClusterConfigurationInfo> configurationInfo;

    /**
     * @return Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
     * 
     */
    public Output<Optional<ClusterConfigurationInfo>> configurationInfo() {
        return Codegen.optional(this.configurationInfo);
    }
    /**
     * Current version of the MSK Cluster used for updates, e.g., `K13V1IB3VIYZZH`
     * * `encryption_info.0.encryption_at_rest_kms_key_arn` - The ARN of the KMS key used for encryption at rest of the broker data volumes.
     * 
     */
    @Export(name="currentVersion", type=String.class, parameters={})
    private Output<String> currentVersion;

    /**
     * @return Current version of the MSK Cluster used for updates, e.g., `K13V1IB3VIYZZH`
     * * `encryption_info.0.encryption_at_rest_kms_key_arn` - The ARN of the KMS key used for encryption at rest of the broker data volumes.
     * 
     */
    public Output<String> currentVersion() {
        return this.currentVersion;
    }
    /**
     * Configuration block for specifying encryption. See below.
     * 
     */
    @Export(name="encryptionInfo", type=ClusterEncryptionInfo.class, parameters={})
    private Output</* @Nullable */ ClusterEncryptionInfo> encryptionInfo;

    /**
     * @return Configuration block for specifying encryption. See below.
     * 
     */
    public Output<Optional<ClusterEncryptionInfo>> encryptionInfo() {
        return Codegen.optional(this.encryptionInfo);
    }
    /**
     * Specify the desired enhanced MSK CloudWatch monitoring level. See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
     * 
     */
    @Export(name="enhancedMonitoring", type=String.class, parameters={})
    private Output</* @Nullable */ String> enhancedMonitoring;

    /**
     * @return Specify the desired enhanced MSK CloudWatch monitoring level. See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
     * 
     */
    public Output<Optional<String>> enhancedMonitoring() {
        return Codegen.optional(this.enhancedMonitoring);
    }
    /**
     * Specify the desired Kafka software version.
     * 
     */
    @Export(name="kafkaVersion", type=String.class, parameters={})
    private Output<String> kafkaVersion;

    /**
     * @return Specify the desired Kafka software version.
     * 
     */
    public Output<String> kafkaVersion() {
        return this.kafkaVersion;
    }
    /**
     * Configuration block for streaming broker logs to Cloudwatch/S3/Kinesis Firehose. See below.
     * 
     */
    @Export(name="loggingInfo", type=ClusterLoggingInfo.class, parameters={})
    private Output</* @Nullable */ ClusterLoggingInfo> loggingInfo;

    /**
     * @return Configuration block for streaming broker logs to Cloudwatch/S3/Kinesis Firehose. See below.
     * 
     */
    public Output<Optional<ClusterLoggingInfo>> loggingInfo() {
        return Codegen.optional(this.loggingInfo);
    }
    /**
     * The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
     * 
     */
    @Export(name="numberOfBrokerNodes", type=Integer.class, parameters={})
    private Output<Integer> numberOfBrokerNodes;

    /**
     * @return The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
     * 
     */
    public Output<Integer> numberOfBrokerNodes() {
        return this.numberOfBrokerNodes;
    }
    /**
     * Configuration block for JMX and Node monitoring for the MSK cluster. See below.
     * 
     */
    @Export(name="openMonitoring", type=ClusterOpenMonitoring.class, parameters={})
    private Output</* @Nullable */ ClusterOpenMonitoring> openMonitoring;

    /**
     * @return Configuration block for JMX and Node monitoring for the MSK cluster. See below.
     * 
     */
    public Output<Optional<ClusterOpenMonitoring>> openMonitoring() {
        return Codegen.optional(this.openMonitoring);
    }
    /**
     * A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Export(name="tagsAll", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    /**
     * A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster. The returned values are sorted alphabetically. The AWS API may not return all endpoints, so this value is not guaranteed to be stable across applies.
     * 
     */
    @Export(name="zookeeperConnectString", type=String.class, parameters={})
    private Output<String> zookeeperConnectString;

    /**
     * @return A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster. The returned values are sorted alphabetically. The AWS API may not return all endpoints, so this value is not guaranteed to be stable across applies.
     * 
     */
    public Output<String> zookeeperConnectString() {
        return this.zookeeperConnectString;
    }
    /**
     * A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster via TLS. The returned values are sorted alphabetically. The AWS API may not return all endpoints, so this value is not guaranteed to be stable across applies.
     * 
     */
    @Export(name="zookeeperConnectStringTls", type=String.class, parameters={})
    private Output<String> zookeeperConnectStringTls;

    /**
     * @return A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster via TLS. The returned values are sorted alphabetically. The AWS API may not return all endpoints, so this value is not guaranteed to be stable across applies.
     * 
     */
    public Output<String> zookeeperConnectStringTls() {
        return this.zookeeperConnectStringTls;
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
    public Cluster(String name, ClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Cluster(String name, ClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:msk/cluster:Cluster", name, args == null ? ClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Cluster(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:msk/cluster:Cluster", name, state, makeResourceOptions(options, id));
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
