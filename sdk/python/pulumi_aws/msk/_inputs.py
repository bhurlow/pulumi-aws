# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ClusterBrokerNodeGroupInfoArgs',
    'ClusterBrokerNodeGroupInfoConnectivityInfoArgs',
    'ClusterBrokerNodeGroupInfoConnectivityInfoPublicAccessArgs',
    'ClusterBrokerNodeGroupInfoStorageInfoArgs',
    'ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs',
    'ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughputArgs',
    'ClusterClientAuthenticationArgs',
    'ClusterClientAuthenticationSaslArgs',
    'ClusterClientAuthenticationTlsArgs',
    'ClusterConfigurationInfoArgs',
    'ClusterEncryptionInfoArgs',
    'ClusterEncryptionInfoEncryptionInTransitArgs',
    'ClusterLoggingInfoArgs',
    'ClusterLoggingInfoBrokerLogsArgs',
    'ClusterLoggingInfoBrokerLogsCloudwatchLogsArgs',
    'ClusterLoggingInfoBrokerLogsFirehoseArgs',
    'ClusterLoggingInfoBrokerLogsS3Args',
    'ClusterOpenMonitoringArgs',
    'ClusterOpenMonitoringPrometheusArgs',
    'ClusterOpenMonitoringPrometheusJmxExporterArgs',
    'ClusterOpenMonitoringPrometheusNodeExporterArgs',
]

@pulumi.input_type
class ClusterBrokerNodeGroupInfoArgs:
    def __init__(__self__, *,
                 client_subnets: pulumi.Input[Sequence[pulumi.Input[str]]],
                 instance_type: pulumi.Input[str],
                 security_groups: pulumi.Input[Sequence[pulumi.Input[str]]],
                 az_distribution: Optional[pulumi.Input[str]] = None,
                 connectivity_info: Optional[pulumi.Input['ClusterBrokerNodeGroupInfoConnectivityInfoArgs']] = None,
                 ebs_volume_size: Optional[pulumi.Input[int]] = None,
                 storage_info: Optional[pulumi.Input['ClusterBrokerNodeGroupInfoStorageInfoArgs']] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] client_subnets: A list of subnets to connect to in client VPC ([documentation](https://docs.aws.amazon.com/msk/1.0/apireference/clusters.html#clusters-prop-brokernodegroupinfo-clientsubnets)).
        :param pulumi.Input[str] instance_type: Specify the instance type to use for the kafka brokersE.g., kafka.m5.large. ([Pricing info](https://aws.amazon.com/msk/pricing/))
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups: A list of the security groups to associate with the elastic network interfaces to control who can communicate with the cluster.
        :param pulumi.Input[str] az_distribution: The distribution of broker nodes across availability zones ([documentation](https://docs.aws.amazon.com/msk/1.0/apireference/clusters.html#clusters-model-brokerazdistribution)). Currently the only valid value is `DEFAULT`.
        :param pulumi.Input['ClusterBrokerNodeGroupInfoConnectivityInfoArgs'] connectivity_info: Information about the cluster access configuration. See below. For security reasons, you can't turn on public access while creating an MSK cluster. However, you can update an existing cluster to make it publicly accessible. You can also create a new cluster and then update it to make it publicly accessible ([documentation](https://docs.aws.amazon.com/msk/latest/developerguide/public-access.html)).
        :param pulumi.Input[int] ebs_volume_size: The size in GiB of the EBS volume for the data drive on each broker node.
        :param pulumi.Input['ClusterBrokerNodeGroupInfoStorageInfoArgs'] storage_info: A block that contains information about storage volumes attached to MSK broker nodes. See below.
        """
        pulumi.set(__self__, "client_subnets", client_subnets)
        pulumi.set(__self__, "instance_type", instance_type)
        pulumi.set(__self__, "security_groups", security_groups)
        if az_distribution is not None:
            pulumi.set(__self__, "az_distribution", az_distribution)
        if connectivity_info is not None:
            pulumi.set(__self__, "connectivity_info", connectivity_info)
        if ebs_volume_size is not None:
            warnings.warn("""use 'storage_info' argument instead""", DeprecationWarning)
            pulumi.log.warn("""ebs_volume_size is deprecated: use 'storage_info' argument instead""")
        if ebs_volume_size is not None:
            pulumi.set(__self__, "ebs_volume_size", ebs_volume_size)
        if storage_info is not None:
            pulumi.set(__self__, "storage_info", storage_info)

    @property
    @pulumi.getter(name="clientSubnets")
    def client_subnets(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of subnets to connect to in client VPC ([documentation](https://docs.aws.amazon.com/msk/1.0/apireference/clusters.html#clusters-prop-brokernodegroupinfo-clientsubnets)).
        """
        return pulumi.get(self, "client_subnets")

    @client_subnets.setter
    def client_subnets(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "client_subnets", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Input[str]:
        """
        Specify the instance type to use for the kafka brokersE.g., kafka.m5.large. ([Pricing info](https://aws.amazon.com/msk/pricing/))
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of the security groups to associate with the elastic network interfaces to control who can communicate with the cluster.
        """
        return pulumi.get(self, "security_groups")

    @security_groups.setter
    def security_groups(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "security_groups", value)

    @property
    @pulumi.getter(name="azDistribution")
    def az_distribution(self) -> Optional[pulumi.Input[str]]:
        """
        The distribution of broker nodes across availability zones ([documentation](https://docs.aws.amazon.com/msk/1.0/apireference/clusters.html#clusters-model-brokerazdistribution)). Currently the only valid value is `DEFAULT`.
        """
        return pulumi.get(self, "az_distribution")

    @az_distribution.setter
    def az_distribution(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "az_distribution", value)

    @property
    @pulumi.getter(name="connectivityInfo")
    def connectivity_info(self) -> Optional[pulumi.Input['ClusterBrokerNodeGroupInfoConnectivityInfoArgs']]:
        """
        Information about the cluster access configuration. See below. For security reasons, you can't turn on public access while creating an MSK cluster. However, you can update an existing cluster to make it publicly accessible. You can also create a new cluster and then update it to make it publicly accessible ([documentation](https://docs.aws.amazon.com/msk/latest/developerguide/public-access.html)).
        """
        return pulumi.get(self, "connectivity_info")

    @connectivity_info.setter
    def connectivity_info(self, value: Optional[pulumi.Input['ClusterBrokerNodeGroupInfoConnectivityInfoArgs']]):
        pulumi.set(self, "connectivity_info", value)

    @property
    @pulumi.getter(name="ebsVolumeSize")
    def ebs_volume_size(self) -> Optional[pulumi.Input[int]]:
        """
        The size in GiB of the EBS volume for the data drive on each broker node.
        """
        return pulumi.get(self, "ebs_volume_size")

    @ebs_volume_size.setter
    def ebs_volume_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ebs_volume_size", value)

    @property
    @pulumi.getter(name="storageInfo")
    def storage_info(self) -> Optional[pulumi.Input['ClusterBrokerNodeGroupInfoStorageInfoArgs']]:
        """
        A block that contains information about storage volumes attached to MSK broker nodes. See below.
        """
        return pulumi.get(self, "storage_info")

    @storage_info.setter
    def storage_info(self, value: Optional[pulumi.Input['ClusterBrokerNodeGroupInfoStorageInfoArgs']]):
        pulumi.set(self, "storage_info", value)


@pulumi.input_type
class ClusterBrokerNodeGroupInfoConnectivityInfoArgs:
    def __init__(__self__, *,
                 public_access: Optional[pulumi.Input['ClusterBrokerNodeGroupInfoConnectivityInfoPublicAccessArgs']] = None):
        """
        :param pulumi.Input['ClusterBrokerNodeGroupInfoConnectivityInfoPublicAccessArgs'] public_access: Access control settings for brokers. See below.
        """
        if public_access is not None:
            pulumi.set(__self__, "public_access", public_access)

    @property
    @pulumi.getter(name="publicAccess")
    def public_access(self) -> Optional[pulumi.Input['ClusterBrokerNodeGroupInfoConnectivityInfoPublicAccessArgs']]:
        """
        Access control settings for brokers. See below.
        """
        return pulumi.get(self, "public_access")

    @public_access.setter
    def public_access(self, value: Optional[pulumi.Input['ClusterBrokerNodeGroupInfoConnectivityInfoPublicAccessArgs']]):
        pulumi.set(self, "public_access", value)


@pulumi.input_type
class ClusterBrokerNodeGroupInfoConnectivityInfoPublicAccessArgs:
    def __init__(__self__, *,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: Public access type. Valida values: `DISABLED`, `SERVICE_PROVIDED_EIPS`.
        """
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Public access type. Valida values: `DISABLED`, `SERVICE_PROVIDED_EIPS`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ClusterBrokerNodeGroupInfoStorageInfoArgs:
    def __init__(__self__, *,
                 ebs_storage_info: Optional[pulumi.Input['ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs']] = None):
        """
        :param pulumi.Input['ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs'] ebs_storage_info: A block that contains EBS volume information. See below.
        """
        if ebs_storage_info is not None:
            pulumi.set(__self__, "ebs_storage_info", ebs_storage_info)

    @property
    @pulumi.getter(name="ebsStorageInfo")
    def ebs_storage_info(self) -> Optional[pulumi.Input['ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs']]:
        """
        A block that contains EBS volume information. See below.
        """
        return pulumi.get(self, "ebs_storage_info")

    @ebs_storage_info.setter
    def ebs_storage_info(self, value: Optional[pulumi.Input['ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs']]):
        pulumi.set(self, "ebs_storage_info", value)


@pulumi.input_type
class ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoArgs:
    def __init__(__self__, *,
                 provisioned_throughput: Optional[pulumi.Input['ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughputArgs']] = None,
                 volume_size: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input['ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughputArgs'] provisioned_throughput: A block that contains EBS volume provisioned throughput information. To provision storage throughput, you must choose broker type kafka.m5.4xlarge or larger. See below.
        :param pulumi.Input[int] volume_size: The size in GiB of the EBS volume for the data drive on each broker node. Minimum value of `1` and maximum value of `16384`.
        """
        if provisioned_throughput is not None:
            pulumi.set(__self__, "provisioned_throughput", provisioned_throughput)
        if volume_size is not None:
            pulumi.set(__self__, "volume_size", volume_size)

    @property
    @pulumi.getter(name="provisionedThroughput")
    def provisioned_throughput(self) -> Optional[pulumi.Input['ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughputArgs']]:
        """
        A block that contains EBS volume provisioned throughput information. To provision storage throughput, you must choose broker type kafka.m5.4xlarge or larger. See below.
        """
        return pulumi.get(self, "provisioned_throughput")

    @provisioned_throughput.setter
    def provisioned_throughput(self, value: Optional[pulumi.Input['ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughputArgs']]):
        pulumi.set(self, "provisioned_throughput", value)

    @property
    @pulumi.getter(name="volumeSize")
    def volume_size(self) -> Optional[pulumi.Input[int]]:
        """
        The size in GiB of the EBS volume for the data drive on each broker node. Minimum value of `1` and maximum value of `16384`.
        """
        return pulumi.get(self, "volume_size")

    @volume_size.setter
    def volume_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "volume_size", value)


@pulumi.input_type
class ClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughputArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 volume_throughput: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[bool] enabled: Controls whether provisioned throughput is enabled or not. Default value: `false`.
        :param pulumi.Input[int] volume_throughput: Throughput value of the EBS volumes for the data drive on each kafka broker node in MiB per second. The minimum value is `250`. The maximum value varies between broker type. You can refer to the valid values for the maximum volume throughput at the following [documentation on throughput bottlenecks](https://docs.aws.amazon.com/msk/latest/developerguide/msk-provision-throughput.html#throughput-bottlenecks)
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if volume_throughput is not None:
            pulumi.set(__self__, "volume_throughput", volume_throughput)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Controls whether provisioned throughput is enabled or not. Default value: `false`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="volumeThroughput")
    def volume_throughput(self) -> Optional[pulumi.Input[int]]:
        """
        Throughput value of the EBS volumes for the data drive on each kafka broker node in MiB per second. The minimum value is `250`. The maximum value varies between broker type. You can refer to the valid values for the maximum volume throughput at the following [documentation on throughput bottlenecks](https://docs.aws.amazon.com/msk/latest/developerguide/msk-provision-throughput.html#throughput-bottlenecks)
        """
        return pulumi.get(self, "volume_throughput")

    @volume_throughput.setter
    def volume_throughput(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "volume_throughput", value)


@pulumi.input_type
class ClusterClientAuthenticationArgs:
    def __init__(__self__, *,
                 sasl: Optional[pulumi.Input['ClusterClientAuthenticationSaslArgs']] = None,
                 tls: Optional[pulumi.Input['ClusterClientAuthenticationTlsArgs']] = None,
                 unauthenticated: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input['ClusterClientAuthenticationSaslArgs'] sasl: Configuration block for specifying SASL client authentication. See below.
        :param pulumi.Input['ClusterClientAuthenticationTlsArgs'] tls: Configuration block for specifying TLS client authentication. See below.
        :param pulumi.Input[bool] unauthenticated: Enables unauthenticated access.
        """
        if sasl is not None:
            pulumi.set(__self__, "sasl", sasl)
        if tls is not None:
            pulumi.set(__self__, "tls", tls)
        if unauthenticated is not None:
            pulumi.set(__self__, "unauthenticated", unauthenticated)

    @property
    @pulumi.getter
    def sasl(self) -> Optional[pulumi.Input['ClusterClientAuthenticationSaslArgs']]:
        """
        Configuration block for specifying SASL client authentication. See below.
        """
        return pulumi.get(self, "sasl")

    @sasl.setter
    def sasl(self, value: Optional[pulumi.Input['ClusterClientAuthenticationSaslArgs']]):
        pulumi.set(self, "sasl", value)

    @property
    @pulumi.getter
    def tls(self) -> Optional[pulumi.Input['ClusterClientAuthenticationTlsArgs']]:
        """
        Configuration block for specifying TLS client authentication. See below.
        """
        return pulumi.get(self, "tls")

    @tls.setter
    def tls(self, value: Optional[pulumi.Input['ClusterClientAuthenticationTlsArgs']]):
        pulumi.set(self, "tls", value)

    @property
    @pulumi.getter
    def unauthenticated(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables unauthenticated access.
        """
        return pulumi.get(self, "unauthenticated")

    @unauthenticated.setter
    def unauthenticated(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "unauthenticated", value)


@pulumi.input_type
class ClusterClientAuthenticationSaslArgs:
    def __init__(__self__, *,
                 iam: Optional[pulumi.Input[bool]] = None,
                 scram: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] iam: Enables IAM client authentication. Defaults to `false`.
        :param pulumi.Input[bool] scram: Enables SCRAM client authentication via AWS Secrets Manager. Defaults to `false`.
        """
        if iam is not None:
            pulumi.set(__self__, "iam", iam)
        if scram is not None:
            pulumi.set(__self__, "scram", scram)

    @property
    @pulumi.getter
    def iam(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables IAM client authentication. Defaults to `false`.
        """
        return pulumi.get(self, "iam")

    @iam.setter
    def iam(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "iam", value)

    @property
    @pulumi.getter
    def scram(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables SCRAM client authentication via AWS Secrets Manager. Defaults to `false`.
        """
        return pulumi.get(self, "scram")

    @scram.setter
    def scram(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "scram", value)


@pulumi.input_type
class ClusterClientAuthenticationTlsArgs:
    def __init__(__self__, *,
                 certificate_authority_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] certificate_authority_arns: List of ACM Certificate Authority Amazon Resource Names (ARNs).
        """
        if certificate_authority_arns is not None:
            pulumi.set(__self__, "certificate_authority_arns", certificate_authority_arns)

    @property
    @pulumi.getter(name="certificateAuthorityArns")
    def certificate_authority_arns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of ACM Certificate Authority Amazon Resource Names (ARNs).
        """
        return pulumi.get(self, "certificate_authority_arns")

    @certificate_authority_arns.setter
    def certificate_authority_arns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "certificate_authority_arns", value)


@pulumi.input_type
class ClusterConfigurationInfoArgs:
    def __init__(__self__, *,
                 arn: pulumi.Input[str],
                 revision: pulumi.Input[int]):
        """
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
        :param pulumi.Input[int] revision: Revision of the MSK Configuration to use in the cluster.
        """
        pulumi.set(__self__, "arn", arn)
        pulumi.set(__self__, "revision", revision)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Input[str]:
        """
        Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def revision(self) -> pulumi.Input[int]:
        """
        Revision of the MSK Configuration to use in the cluster.
        """
        return pulumi.get(self, "revision")

    @revision.setter
    def revision(self, value: pulumi.Input[int]):
        pulumi.set(self, "revision", value)


@pulumi.input_type
class ClusterEncryptionInfoArgs:
    def __init__(__self__, *,
                 encryption_at_rest_kms_key_arn: Optional[pulumi.Input[str]] = None,
                 encryption_in_transit: Optional[pulumi.Input['ClusterEncryptionInfoEncryptionInTransitArgs']] = None):
        """
        :param pulumi.Input[str] encryption_at_rest_kms_key_arn: You may specify a KMS key short ID or ARN (it will always output an ARN) to use for encrypting your data at rest.  If no key is specified, an AWS managed KMS ('aws/msk' managed service) key will be used for encrypting the data at rest.
        :param pulumi.Input['ClusterEncryptionInfoEncryptionInTransitArgs'] encryption_in_transit: Configuration block to specify encryption in transit. See below.
        """
        if encryption_at_rest_kms_key_arn is not None:
            pulumi.set(__self__, "encryption_at_rest_kms_key_arn", encryption_at_rest_kms_key_arn)
        if encryption_in_transit is not None:
            pulumi.set(__self__, "encryption_in_transit", encryption_in_transit)

    @property
    @pulumi.getter(name="encryptionAtRestKmsKeyArn")
    def encryption_at_rest_kms_key_arn(self) -> Optional[pulumi.Input[str]]:
        """
        You may specify a KMS key short ID or ARN (it will always output an ARN) to use for encrypting your data at rest.  If no key is specified, an AWS managed KMS ('aws/msk' managed service) key will be used for encrypting the data at rest.
        """
        return pulumi.get(self, "encryption_at_rest_kms_key_arn")

    @encryption_at_rest_kms_key_arn.setter
    def encryption_at_rest_kms_key_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encryption_at_rest_kms_key_arn", value)

    @property
    @pulumi.getter(name="encryptionInTransit")
    def encryption_in_transit(self) -> Optional[pulumi.Input['ClusterEncryptionInfoEncryptionInTransitArgs']]:
        """
        Configuration block to specify encryption in transit. See below.
        """
        return pulumi.get(self, "encryption_in_transit")

    @encryption_in_transit.setter
    def encryption_in_transit(self, value: Optional[pulumi.Input['ClusterEncryptionInfoEncryptionInTransitArgs']]):
        pulumi.set(self, "encryption_in_transit", value)


@pulumi.input_type
class ClusterEncryptionInfoEncryptionInTransitArgs:
    def __init__(__self__, *,
                 client_broker: Optional[pulumi.Input[str]] = None,
                 in_cluster: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[str] client_broker: Encryption setting for data in transit between clients and brokers. Valid values: `TLS`, `TLS_PLAINTEXT`, and `PLAINTEXT`. Default value is `TLS`.
        :param pulumi.Input[bool] in_cluster: Whether data communication among broker nodes is encrypted. Default value: `true`.
        """
        if client_broker is not None:
            pulumi.set(__self__, "client_broker", client_broker)
        if in_cluster is not None:
            pulumi.set(__self__, "in_cluster", in_cluster)

    @property
    @pulumi.getter(name="clientBroker")
    def client_broker(self) -> Optional[pulumi.Input[str]]:
        """
        Encryption setting for data in transit between clients and brokers. Valid values: `TLS`, `TLS_PLAINTEXT`, and `PLAINTEXT`. Default value is `TLS`.
        """
        return pulumi.get(self, "client_broker")

    @client_broker.setter
    def client_broker(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_broker", value)

    @property
    @pulumi.getter(name="inCluster")
    def in_cluster(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether data communication among broker nodes is encrypted. Default value: `true`.
        """
        return pulumi.get(self, "in_cluster")

    @in_cluster.setter
    def in_cluster(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "in_cluster", value)


@pulumi.input_type
class ClusterLoggingInfoArgs:
    def __init__(__self__, *,
                 broker_logs: pulumi.Input['ClusterLoggingInfoBrokerLogsArgs']):
        """
        :param pulumi.Input['ClusterLoggingInfoBrokerLogsArgs'] broker_logs: Configuration block for Broker Logs settings for logging info. See below.
        """
        pulumi.set(__self__, "broker_logs", broker_logs)

    @property
    @pulumi.getter(name="brokerLogs")
    def broker_logs(self) -> pulumi.Input['ClusterLoggingInfoBrokerLogsArgs']:
        """
        Configuration block for Broker Logs settings for logging info. See below.
        """
        return pulumi.get(self, "broker_logs")

    @broker_logs.setter
    def broker_logs(self, value: pulumi.Input['ClusterLoggingInfoBrokerLogsArgs']):
        pulumi.set(self, "broker_logs", value)


@pulumi.input_type
class ClusterLoggingInfoBrokerLogsArgs:
    def __init__(__self__, *,
                 cloudwatch_logs: Optional[pulumi.Input['ClusterLoggingInfoBrokerLogsCloudwatchLogsArgs']] = None,
                 firehose: Optional[pulumi.Input['ClusterLoggingInfoBrokerLogsFirehoseArgs']] = None,
                 s3: Optional[pulumi.Input['ClusterLoggingInfoBrokerLogsS3Args']] = None):
        if cloudwatch_logs is not None:
            pulumi.set(__self__, "cloudwatch_logs", cloudwatch_logs)
        if firehose is not None:
            pulumi.set(__self__, "firehose", firehose)
        if s3 is not None:
            pulumi.set(__self__, "s3", s3)

    @property
    @pulumi.getter(name="cloudwatchLogs")
    def cloudwatch_logs(self) -> Optional[pulumi.Input['ClusterLoggingInfoBrokerLogsCloudwatchLogsArgs']]:
        return pulumi.get(self, "cloudwatch_logs")

    @cloudwatch_logs.setter
    def cloudwatch_logs(self, value: Optional[pulumi.Input['ClusterLoggingInfoBrokerLogsCloudwatchLogsArgs']]):
        pulumi.set(self, "cloudwatch_logs", value)

    @property
    @pulumi.getter
    def firehose(self) -> Optional[pulumi.Input['ClusterLoggingInfoBrokerLogsFirehoseArgs']]:
        return pulumi.get(self, "firehose")

    @firehose.setter
    def firehose(self, value: Optional[pulumi.Input['ClusterLoggingInfoBrokerLogsFirehoseArgs']]):
        pulumi.set(self, "firehose", value)

    @property
    @pulumi.getter
    def s3(self) -> Optional[pulumi.Input['ClusterLoggingInfoBrokerLogsS3Args']]:
        return pulumi.get(self, "s3")

    @s3.setter
    def s3(self, value: Optional[pulumi.Input['ClusterLoggingInfoBrokerLogsS3Args']]):
        pulumi.set(self, "s3", value)


@pulumi.input_type
class ClusterLoggingInfoBrokerLogsCloudwatchLogsArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 log_group: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] enabled: Controls whether provisioned throughput is enabled or not. Default value: `false`.
        :param pulumi.Input[str] log_group: Name of the Cloudwatch Log Group to deliver logs to.
        """
        pulumi.set(__self__, "enabled", enabled)
        if log_group is not None:
            pulumi.set(__self__, "log_group", log_group)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        Controls whether provisioned throughput is enabled or not. Default value: `false`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="logGroup")
    def log_group(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Cloudwatch Log Group to deliver logs to.
        """
        return pulumi.get(self, "log_group")

    @log_group.setter
    def log_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_group", value)


@pulumi.input_type
class ClusterLoggingInfoBrokerLogsFirehoseArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 delivery_stream: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] enabled: Controls whether provisioned throughput is enabled or not. Default value: `false`.
        :param pulumi.Input[str] delivery_stream: Name of the Kinesis Data Firehose delivery stream to deliver logs to.
        """
        pulumi.set(__self__, "enabled", enabled)
        if delivery_stream is not None:
            pulumi.set(__self__, "delivery_stream", delivery_stream)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        Controls whether provisioned throughput is enabled or not. Default value: `false`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="deliveryStream")
    def delivery_stream(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Kinesis Data Firehose delivery stream to deliver logs to.
        """
        return pulumi.get(self, "delivery_stream")

    @delivery_stream.setter
    def delivery_stream(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delivery_stream", value)


@pulumi.input_type
class ClusterLoggingInfoBrokerLogsS3Args:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 bucket: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] enabled: Controls whether provisioned throughput is enabled or not. Default value: `false`.
        :param pulumi.Input[str] bucket: Name of the S3 bucket to deliver logs to.
        :param pulumi.Input[str] prefix: Prefix to append to the folder name.
        """
        pulumi.set(__self__, "enabled", enabled)
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        Controls whether provisioned throughput is enabled or not. Default value: `false`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the S3 bucket to deliver logs to.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Prefix to append to the folder name.
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)


@pulumi.input_type
class ClusterOpenMonitoringArgs:
    def __init__(__self__, *,
                 prometheus: pulumi.Input['ClusterOpenMonitoringPrometheusArgs']):
        """
        :param pulumi.Input['ClusterOpenMonitoringPrometheusArgs'] prometheus: Configuration block for Prometheus settings for open monitoring. See below.
        """
        pulumi.set(__self__, "prometheus", prometheus)

    @property
    @pulumi.getter
    def prometheus(self) -> pulumi.Input['ClusterOpenMonitoringPrometheusArgs']:
        """
        Configuration block for Prometheus settings for open monitoring. See below.
        """
        return pulumi.get(self, "prometheus")

    @prometheus.setter
    def prometheus(self, value: pulumi.Input['ClusterOpenMonitoringPrometheusArgs']):
        pulumi.set(self, "prometheus", value)


@pulumi.input_type
class ClusterOpenMonitoringPrometheusArgs:
    def __init__(__self__, *,
                 jmx_exporter: Optional[pulumi.Input['ClusterOpenMonitoringPrometheusJmxExporterArgs']] = None,
                 node_exporter: Optional[pulumi.Input['ClusterOpenMonitoringPrometheusNodeExporterArgs']] = None):
        """
        :param pulumi.Input['ClusterOpenMonitoringPrometheusJmxExporterArgs'] jmx_exporter: Configuration block for JMX Exporter. See below.
        :param pulumi.Input['ClusterOpenMonitoringPrometheusNodeExporterArgs'] node_exporter: Configuration block for Node Exporter. See below.
        """
        if jmx_exporter is not None:
            pulumi.set(__self__, "jmx_exporter", jmx_exporter)
        if node_exporter is not None:
            pulumi.set(__self__, "node_exporter", node_exporter)

    @property
    @pulumi.getter(name="jmxExporter")
    def jmx_exporter(self) -> Optional[pulumi.Input['ClusterOpenMonitoringPrometheusJmxExporterArgs']]:
        """
        Configuration block for JMX Exporter. See below.
        """
        return pulumi.get(self, "jmx_exporter")

    @jmx_exporter.setter
    def jmx_exporter(self, value: Optional[pulumi.Input['ClusterOpenMonitoringPrometheusJmxExporterArgs']]):
        pulumi.set(self, "jmx_exporter", value)

    @property
    @pulumi.getter(name="nodeExporter")
    def node_exporter(self) -> Optional[pulumi.Input['ClusterOpenMonitoringPrometheusNodeExporterArgs']]:
        """
        Configuration block for Node Exporter. See below.
        """
        return pulumi.get(self, "node_exporter")

    @node_exporter.setter
    def node_exporter(self, value: Optional[pulumi.Input['ClusterOpenMonitoringPrometheusNodeExporterArgs']]):
        pulumi.set(self, "node_exporter", value)


@pulumi.input_type
class ClusterOpenMonitoringPrometheusJmxExporterArgs:
    def __init__(__self__, *,
                 enabled_in_broker: pulumi.Input[bool]):
        """
        :param pulumi.Input[bool] enabled_in_broker: Indicates whether you want to enable or disable the JMX Exporter.
        """
        pulumi.set(__self__, "enabled_in_broker", enabled_in_broker)

    @property
    @pulumi.getter(name="enabledInBroker")
    def enabled_in_broker(self) -> pulumi.Input[bool]:
        """
        Indicates whether you want to enable or disable the JMX Exporter.
        """
        return pulumi.get(self, "enabled_in_broker")

    @enabled_in_broker.setter
    def enabled_in_broker(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled_in_broker", value)


@pulumi.input_type
class ClusterOpenMonitoringPrometheusNodeExporterArgs:
    def __init__(__self__, *,
                 enabled_in_broker: pulumi.Input[bool]):
        """
        :param pulumi.Input[bool] enabled_in_broker: Indicates whether you want to enable or disable the JMX Exporter.
        """
        pulumi.set(__self__, "enabled_in_broker", enabled_in_broker)

    @property
    @pulumi.getter(name="enabledInBroker")
    def enabled_in_broker(self) -> pulumi.Input[bool]:
        """
        Indicates whether you want to enable or disable the JMX Exporter.
        """
        return pulumi.get(self, "enabled_in_broker")

    @enabled_in_broker.setter
    def enabled_in_broker(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled_in_broker", value)


