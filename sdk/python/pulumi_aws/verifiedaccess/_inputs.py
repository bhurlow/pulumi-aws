# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'InstanceLoggingConfigurationAccessLogsArgs',
    'InstanceLoggingConfigurationAccessLogsCloudwatchLogsArgs',
    'InstanceLoggingConfigurationAccessLogsKinesisDataFirehoseArgs',
    'InstanceLoggingConfigurationAccessLogsS3Args',
    'InstanceVerifiedAccessTrustProviderArgs',
    'TrustProviderDeviceOptionsArgs',
    'TrustProviderOidcOptionsArgs',
]

@pulumi.input_type
class InstanceLoggingConfigurationAccessLogsArgs:
    def __init__(__self__, *,
                 cloudwatch_logs: Optional[pulumi.Input['InstanceLoggingConfigurationAccessLogsCloudwatchLogsArgs']] = None,
                 include_trust_context: Optional[pulumi.Input[bool]] = None,
                 kinesis_data_firehose: Optional[pulumi.Input['InstanceLoggingConfigurationAccessLogsKinesisDataFirehoseArgs']] = None,
                 log_version: Optional[pulumi.Input[str]] = None,
                 s3: Optional[pulumi.Input['InstanceLoggingConfigurationAccessLogsS3Args']] = None):
        """
        :param pulumi.Input['InstanceLoggingConfigurationAccessLogsCloudwatchLogsArgs'] cloudwatch_logs: A block that specifies configures sending Verified Access logs to CloudWatch Logs. Detailed below.
        :param pulumi.Input[bool] include_trust_context: Include trust data sent by trust providers into the logs.
        :param pulumi.Input['InstanceLoggingConfigurationAccessLogsKinesisDataFirehoseArgs'] kinesis_data_firehose: A block that specifies configures sending Verified Access logs to Kinesis. Detailed below.
        :param pulumi.Input[str] log_version: The logging version to use. Refer to [VerifiedAccessLogOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VerifiedAccessLogOptions.html) for the allowed values.
        :param pulumi.Input['InstanceLoggingConfigurationAccessLogsS3Args'] s3: A block that specifies configures sending Verified Access logs to S3. Detailed below.
        """
        if cloudwatch_logs is not None:
            pulumi.set(__self__, "cloudwatch_logs", cloudwatch_logs)
        if include_trust_context is not None:
            pulumi.set(__self__, "include_trust_context", include_trust_context)
        if kinesis_data_firehose is not None:
            pulumi.set(__self__, "kinesis_data_firehose", kinesis_data_firehose)
        if log_version is not None:
            pulumi.set(__self__, "log_version", log_version)
        if s3 is not None:
            pulumi.set(__self__, "s3", s3)

    @property
    @pulumi.getter(name="cloudwatchLogs")
    def cloudwatch_logs(self) -> Optional[pulumi.Input['InstanceLoggingConfigurationAccessLogsCloudwatchLogsArgs']]:
        """
        A block that specifies configures sending Verified Access logs to CloudWatch Logs. Detailed below.
        """
        return pulumi.get(self, "cloudwatch_logs")

    @cloudwatch_logs.setter
    def cloudwatch_logs(self, value: Optional[pulumi.Input['InstanceLoggingConfigurationAccessLogsCloudwatchLogsArgs']]):
        pulumi.set(self, "cloudwatch_logs", value)

    @property
    @pulumi.getter(name="includeTrustContext")
    def include_trust_context(self) -> Optional[pulumi.Input[bool]]:
        """
        Include trust data sent by trust providers into the logs.
        """
        return pulumi.get(self, "include_trust_context")

    @include_trust_context.setter
    def include_trust_context(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_trust_context", value)

    @property
    @pulumi.getter(name="kinesisDataFirehose")
    def kinesis_data_firehose(self) -> Optional[pulumi.Input['InstanceLoggingConfigurationAccessLogsKinesisDataFirehoseArgs']]:
        """
        A block that specifies configures sending Verified Access logs to Kinesis. Detailed below.
        """
        return pulumi.get(self, "kinesis_data_firehose")

    @kinesis_data_firehose.setter
    def kinesis_data_firehose(self, value: Optional[pulumi.Input['InstanceLoggingConfigurationAccessLogsKinesisDataFirehoseArgs']]):
        pulumi.set(self, "kinesis_data_firehose", value)

    @property
    @pulumi.getter(name="logVersion")
    def log_version(self) -> Optional[pulumi.Input[str]]:
        """
        The logging version to use. Refer to [VerifiedAccessLogOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VerifiedAccessLogOptions.html) for the allowed values.
        """
        return pulumi.get(self, "log_version")

    @log_version.setter
    def log_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_version", value)

    @property
    @pulumi.getter
    def s3(self) -> Optional[pulumi.Input['InstanceLoggingConfigurationAccessLogsS3Args']]:
        """
        A block that specifies configures sending Verified Access logs to S3. Detailed below.
        """
        return pulumi.get(self, "s3")

    @s3.setter
    def s3(self, value: Optional[pulumi.Input['InstanceLoggingConfigurationAccessLogsS3Args']]):
        pulumi.set(self, "s3", value)


@pulumi.input_type
class InstanceLoggingConfigurationAccessLogsCloudwatchLogsArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 log_group: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] enabled: Indicates whether logging is enabled.
        :param pulumi.Input[str] log_group: The name of the CloudWatch Logs Log Group.
        """
        pulumi.set(__self__, "enabled", enabled)
        if log_group is not None:
            pulumi.set(__self__, "log_group", log_group)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        Indicates whether logging is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="logGroup")
    def log_group(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the CloudWatch Logs Log Group.
        """
        return pulumi.get(self, "log_group")

    @log_group.setter
    def log_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_group", value)


@pulumi.input_type
class InstanceLoggingConfigurationAccessLogsKinesisDataFirehoseArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 delivery_stream: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] enabled: Indicates whether logging is enabled.
        :param pulumi.Input[str] delivery_stream: The name of the delivery stream.
        """
        pulumi.set(__self__, "enabled", enabled)
        if delivery_stream is not None:
            pulumi.set(__self__, "delivery_stream", delivery_stream)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        Indicates whether logging is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="deliveryStream")
    def delivery_stream(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the delivery stream.
        """
        return pulumi.get(self, "delivery_stream")

    @delivery_stream.setter
    def delivery_stream(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delivery_stream", value)


@pulumi.input_type
class InstanceLoggingConfigurationAccessLogsS3Args:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 bucket_owner: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] enabled: Indicates whether logging is enabled.
        :param pulumi.Input[str] bucket_name: The name of S3 bucket.
        :param pulumi.Input[str] bucket_owner: The ID of the AWS account that owns the Amazon S3 bucket.
        :param pulumi.Input[str] prefix: The bucket prefix.
        """
        pulumi.set(__self__, "enabled", enabled)
        if bucket_name is not None:
            pulumi.set(__self__, "bucket_name", bucket_name)
        if bucket_owner is not None:
            pulumi.set(__self__, "bucket_owner", bucket_owner)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        Indicates whether logging is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of S3 bucket.
        """
        return pulumi.get(self, "bucket_name")

    @bucket_name.setter
    def bucket_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket_name", value)

    @property
    @pulumi.getter(name="bucketOwner")
    def bucket_owner(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the AWS account that owns the Amazon S3 bucket.
        """
        return pulumi.get(self, "bucket_owner")

    @bucket_owner.setter
    def bucket_owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket_owner", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The bucket prefix.
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)


@pulumi.input_type
class InstanceVerifiedAccessTrustProviderArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 device_trust_provider_type: Optional[pulumi.Input[str]] = None,
                 trust_provider_type: Optional[pulumi.Input[str]] = None,
                 user_trust_provider_type: Optional[pulumi.Input[str]] = None,
                 verified_access_trust_provider_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] description: A description for the AWS Verified Access Instance.
        :param pulumi.Input[str] device_trust_provider_type: The type of device-based trust provider.
        :param pulumi.Input[str] trust_provider_type: The type of trust provider (user- or device-based).
        :param pulumi.Input[str] user_trust_provider_type: The type of user-based trust provider.
        :param pulumi.Input[str] verified_access_trust_provider_id: The ID of the trust provider.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if device_trust_provider_type is not None:
            pulumi.set(__self__, "device_trust_provider_type", device_trust_provider_type)
        if trust_provider_type is not None:
            pulumi.set(__self__, "trust_provider_type", trust_provider_type)
        if user_trust_provider_type is not None:
            pulumi.set(__self__, "user_trust_provider_type", user_trust_provider_type)
        if verified_access_trust_provider_id is not None:
            pulumi.set(__self__, "verified_access_trust_provider_id", verified_access_trust_provider_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the AWS Verified Access Instance.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="deviceTrustProviderType")
    def device_trust_provider_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of device-based trust provider.
        """
        return pulumi.get(self, "device_trust_provider_type")

    @device_trust_provider_type.setter
    def device_trust_provider_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_trust_provider_type", value)

    @property
    @pulumi.getter(name="trustProviderType")
    def trust_provider_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of trust provider (user- or device-based).
        """
        return pulumi.get(self, "trust_provider_type")

    @trust_provider_type.setter
    def trust_provider_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trust_provider_type", value)

    @property
    @pulumi.getter(name="userTrustProviderType")
    def user_trust_provider_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of user-based trust provider.
        """
        return pulumi.get(self, "user_trust_provider_type")

    @user_trust_provider_type.setter
    def user_trust_provider_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_trust_provider_type", value)

    @property
    @pulumi.getter(name="verifiedAccessTrustProviderId")
    def verified_access_trust_provider_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the trust provider.
        """
        return pulumi.get(self, "verified_access_trust_provider_id")

    @verified_access_trust_provider_id.setter
    def verified_access_trust_provider_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "verified_access_trust_provider_id", value)


@pulumi.input_type
class TrustProviderDeviceOptionsArgs:
    def __init__(__self__, *,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


@pulumi.input_type
class TrustProviderOidcOptionsArgs:
    def __init__(__self__, *,
                 client_secret: pulumi.Input[str],
                 authorization_endpoint: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 token_endpoint: Optional[pulumi.Input[str]] = None,
                 user_info_endpoint: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "client_secret", client_secret)
        if authorization_endpoint is not None:
            pulumi.set(__self__, "authorization_endpoint", authorization_endpoint)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if issuer is not None:
            pulumi.set(__self__, "issuer", issuer)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if token_endpoint is not None:
            pulumi.set(__self__, "token_endpoint", token_endpoint)
        if user_info_endpoint is not None:
            pulumi.set(__self__, "user_info_endpoint", user_info_endpoint)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Input[str]:
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter(name="authorizationEndpoint")
    def authorization_endpoint(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "authorization_endpoint")

    @authorization_endpoint.setter
    def authorization_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization_endpoint", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter
    def issuer(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "issuer")

    @issuer.setter
    def issuer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuer", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter(name="tokenEndpoint")
    def token_endpoint(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "token_endpoint")

    @token_endpoint.setter
    def token_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token_endpoint", value)

    @property
    @pulumi.getter(name="userInfoEndpoint")
    def user_info_endpoint(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "user_info_endpoint")

    @user_info_endpoint.setter
    def user_info_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_info_endpoint", value)


