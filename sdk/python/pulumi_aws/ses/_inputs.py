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
    'ConfigurationSetDeliveryOptionsArgs',
    'ConfigurationSetTrackingOptionsArgs',
    'EventDestinationCloudwatchDestinationArgs',
    'EventDestinationKinesisDestinationArgs',
    'EventDestinationSnsDestinationArgs',
    'ReceiptRuleAddHeaderActionArgs',
    'ReceiptRuleBounceActionArgs',
    'ReceiptRuleLambdaActionArgs',
    'ReceiptRuleS3ActionArgs',
    'ReceiptRuleSnsActionArgs',
    'ReceiptRuleStopActionArgs',
    'ReceiptRuleWorkmailActionArgs',
]

@pulumi.input_type
class ConfigurationSetDeliveryOptionsArgs:
    def __init__(__self__, *,
                 tls_policy: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] tls_policy: Whether messages that use the configuration set are required to use Transport Layer Security (TLS). If the value is `Require`, messages are only delivered if a TLS connection can be established. If the value is `Optional`, messages can be delivered in plain text if a TLS connection can't be established. Valid values: `Require` or `Optional`. Defaults to `Optional`.
        """
        if tls_policy is not None:
            pulumi.set(__self__, "tls_policy", tls_policy)

    @property
    @pulumi.getter(name="tlsPolicy")
    def tls_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Whether messages that use the configuration set are required to use Transport Layer Security (TLS). If the value is `Require`, messages are only delivered if a TLS connection can be established. If the value is `Optional`, messages can be delivered in plain text if a TLS connection can't be established. Valid values: `Require` or `Optional`. Defaults to `Optional`.
        """
        return pulumi.get(self, "tls_policy")

    @tls_policy.setter
    def tls_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_policy", value)


@pulumi.input_type
class ConfigurationSetTrackingOptionsArgs:
    def __init__(__self__, *,
                 custom_redirect_domain: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] custom_redirect_domain: Custom subdomain that is used to redirect email recipients to the Amazon SES event tracking domain.
        """
        if custom_redirect_domain is not None:
            pulumi.set(__self__, "custom_redirect_domain", custom_redirect_domain)

    @property
    @pulumi.getter(name="customRedirectDomain")
    def custom_redirect_domain(self) -> Optional[pulumi.Input[str]]:
        """
        Custom subdomain that is used to redirect email recipients to the Amazon SES event tracking domain.
        """
        return pulumi.get(self, "custom_redirect_domain")

    @custom_redirect_domain.setter
    def custom_redirect_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_redirect_domain", value)


@pulumi.input_type
class EventDestinationCloudwatchDestinationArgs:
    def __init__(__self__, *,
                 default_value: pulumi.Input[str],
                 dimension_name: pulumi.Input[str],
                 value_source: pulumi.Input[str]):
        """
        :param pulumi.Input[str] default_value: The default value for the event
        :param pulumi.Input[str] dimension_name: The name for the dimension
        :param pulumi.Input[str] value_source: The source for the value. May be any of `"messageTag"`, `"emailHeader"` or `"linkTag"`.
        """
        pulumi.set(__self__, "default_value", default_value)
        pulumi.set(__self__, "dimension_name", dimension_name)
        pulumi.set(__self__, "value_source", value_source)

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> pulumi.Input[str]:
        """
        The default value for the event
        """
        return pulumi.get(self, "default_value")

    @default_value.setter
    def default_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "default_value", value)

    @property
    @pulumi.getter(name="dimensionName")
    def dimension_name(self) -> pulumi.Input[str]:
        """
        The name for the dimension
        """
        return pulumi.get(self, "dimension_name")

    @dimension_name.setter
    def dimension_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "dimension_name", value)

    @property
    @pulumi.getter(name="valueSource")
    def value_source(self) -> pulumi.Input[str]:
        """
        The source for the value. May be any of `"messageTag"`, `"emailHeader"` or `"linkTag"`.
        """
        return pulumi.get(self, "value_source")

    @value_source.setter
    def value_source(self, value: pulumi.Input[str]):
        pulumi.set(self, "value_source", value)


@pulumi.input_type
class EventDestinationKinesisDestinationArgs:
    def __init__(__self__, *,
                 role_arn: pulumi.Input[str],
                 stream_arn: pulumi.Input[str]):
        """
        :param pulumi.Input[str] role_arn: The ARN of the role that has permissions to access the Kinesis Stream
        :param pulumi.Input[str] stream_arn: The ARN of the Kinesis Stream
        """
        pulumi.set(__self__, "role_arn", role_arn)
        pulumi.set(__self__, "stream_arn", stream_arn)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        """
        The ARN of the role that has permissions to access the Kinesis Stream
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="streamArn")
    def stream_arn(self) -> pulumi.Input[str]:
        """
        The ARN of the Kinesis Stream
        """
        return pulumi.get(self, "stream_arn")

    @stream_arn.setter
    def stream_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "stream_arn", value)


@pulumi.input_type
class EventDestinationSnsDestinationArgs:
    def __init__(__self__, *,
                 topic_arn: pulumi.Input[str]):
        """
        :param pulumi.Input[str] topic_arn: The ARN of the SNS topic
        """
        pulumi.set(__self__, "topic_arn", topic_arn)

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> pulumi.Input[str]:
        """
        The ARN of the SNS topic
        """
        return pulumi.get(self, "topic_arn")

    @topic_arn.setter
    def topic_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic_arn", value)


@pulumi.input_type
class ReceiptRuleAddHeaderActionArgs:
    def __init__(__self__, *,
                 header_name: pulumi.Input[str],
                 header_value: pulumi.Input[str],
                 position: pulumi.Input[int]):
        """
        :param pulumi.Input[str] header_name: The name of the header to add
        :param pulumi.Input[str] header_value: The value of the header to add
        :param pulumi.Input[int] position: The position of the action in the receipt rule
        """
        pulumi.set(__self__, "header_name", header_name)
        pulumi.set(__self__, "header_value", header_value)
        pulumi.set(__self__, "position", position)

    @property
    @pulumi.getter(name="headerName")
    def header_name(self) -> pulumi.Input[str]:
        """
        The name of the header to add
        """
        return pulumi.get(self, "header_name")

    @header_name.setter
    def header_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "header_name", value)

    @property
    @pulumi.getter(name="headerValue")
    def header_value(self) -> pulumi.Input[str]:
        """
        The value of the header to add
        """
        return pulumi.get(self, "header_value")

    @header_value.setter
    def header_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "header_value", value)

    @property
    @pulumi.getter
    def position(self) -> pulumi.Input[int]:
        """
        The position of the action in the receipt rule
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: pulumi.Input[int]):
        pulumi.set(self, "position", value)


@pulumi.input_type
class ReceiptRuleBounceActionArgs:
    def __init__(__self__, *,
                 message: pulumi.Input[str],
                 position: pulumi.Input[int],
                 sender: pulumi.Input[str],
                 smtp_reply_code: pulumi.Input[str],
                 status_code: Optional[pulumi.Input[str]] = None,
                 topic_arn: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] message: The message to send
        :param pulumi.Input[int] position: The position of the action in the receipt rule
        :param pulumi.Input[str] sender: The email address of the sender
        :param pulumi.Input[str] smtp_reply_code: The RFC 5321 SMTP reply code
        :param pulumi.Input[str] status_code: The RFC 3463 SMTP enhanced status code
        :param pulumi.Input[str] topic_arn: The ARN of an SNS topic to notify
        """
        pulumi.set(__self__, "message", message)
        pulumi.set(__self__, "position", position)
        pulumi.set(__self__, "sender", sender)
        pulumi.set(__self__, "smtp_reply_code", smtp_reply_code)
        if status_code is not None:
            pulumi.set(__self__, "status_code", status_code)
        if topic_arn is not None:
            pulumi.set(__self__, "topic_arn", topic_arn)

    @property
    @pulumi.getter
    def message(self) -> pulumi.Input[str]:
        """
        The message to send
        """
        return pulumi.get(self, "message")

    @message.setter
    def message(self, value: pulumi.Input[str]):
        pulumi.set(self, "message", value)

    @property
    @pulumi.getter
    def position(self) -> pulumi.Input[int]:
        """
        The position of the action in the receipt rule
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: pulumi.Input[int]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter
    def sender(self) -> pulumi.Input[str]:
        """
        The email address of the sender
        """
        return pulumi.get(self, "sender")

    @sender.setter
    def sender(self, value: pulumi.Input[str]):
        pulumi.set(self, "sender", value)

    @property
    @pulumi.getter(name="smtpReplyCode")
    def smtp_reply_code(self) -> pulumi.Input[str]:
        """
        The RFC 5321 SMTP reply code
        """
        return pulumi.get(self, "smtp_reply_code")

    @smtp_reply_code.setter
    def smtp_reply_code(self, value: pulumi.Input[str]):
        pulumi.set(self, "smtp_reply_code", value)

    @property
    @pulumi.getter(name="statusCode")
    def status_code(self) -> Optional[pulumi.Input[str]]:
        """
        The RFC 3463 SMTP enhanced status code
        """
        return pulumi.get(self, "status_code")

    @status_code.setter
    def status_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status_code", value)

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of an SNS topic to notify
        """
        return pulumi.get(self, "topic_arn")

    @topic_arn.setter
    def topic_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_arn", value)


@pulumi.input_type
class ReceiptRuleLambdaActionArgs:
    def __init__(__self__, *,
                 function_arn: pulumi.Input[str],
                 position: pulumi.Input[int],
                 invocation_type: Optional[pulumi.Input[str]] = None,
                 topic_arn: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] function_arn: The ARN of the Lambda function to invoke
        :param pulumi.Input[int] position: The position of the action in the receipt rule
        :param pulumi.Input[str] invocation_type: `Event` or `RequestResponse`
        :param pulumi.Input[str] topic_arn: The ARN of an SNS topic to notify
        """
        pulumi.set(__self__, "function_arn", function_arn)
        pulumi.set(__self__, "position", position)
        if invocation_type is not None:
            pulumi.set(__self__, "invocation_type", invocation_type)
        if topic_arn is not None:
            pulumi.set(__self__, "topic_arn", topic_arn)

    @property
    @pulumi.getter(name="functionArn")
    def function_arn(self) -> pulumi.Input[str]:
        """
        The ARN of the Lambda function to invoke
        """
        return pulumi.get(self, "function_arn")

    @function_arn.setter
    def function_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "function_arn", value)

    @property
    @pulumi.getter
    def position(self) -> pulumi.Input[int]:
        """
        The position of the action in the receipt rule
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: pulumi.Input[int]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter(name="invocationType")
    def invocation_type(self) -> Optional[pulumi.Input[str]]:
        """
        `Event` or `RequestResponse`
        """
        return pulumi.get(self, "invocation_type")

    @invocation_type.setter
    def invocation_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "invocation_type", value)

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of an SNS topic to notify
        """
        return pulumi.get(self, "topic_arn")

    @topic_arn.setter
    def topic_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_arn", value)


@pulumi.input_type
class ReceiptRuleS3ActionArgs:
    def __init__(__self__, *,
                 bucket_name: pulumi.Input[str],
                 position: pulumi.Input[int],
                 kms_key_arn: Optional[pulumi.Input[str]] = None,
                 object_key_prefix: Optional[pulumi.Input[str]] = None,
                 topic_arn: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] bucket_name: The name of the S3 bucket
        :param pulumi.Input[int] position: The position of the action in the receipt rule
        :param pulumi.Input[str] kms_key_arn: The ARN of the KMS key
        :param pulumi.Input[str] object_key_prefix: The key prefix of the S3 bucket
        :param pulumi.Input[str] topic_arn: The ARN of an SNS topic to notify
        """
        pulumi.set(__self__, "bucket_name", bucket_name)
        pulumi.set(__self__, "position", position)
        if kms_key_arn is not None:
            pulumi.set(__self__, "kms_key_arn", kms_key_arn)
        if object_key_prefix is not None:
            pulumi.set(__self__, "object_key_prefix", object_key_prefix)
        if topic_arn is not None:
            pulumi.set(__self__, "topic_arn", topic_arn)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> pulumi.Input[str]:
        """
        The name of the S3 bucket
        """
        return pulumi.get(self, "bucket_name")

    @bucket_name.setter
    def bucket_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket_name", value)

    @property
    @pulumi.getter
    def position(self) -> pulumi.Input[int]:
        """
        The position of the action in the receipt rule
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: pulumi.Input[int]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the KMS key
        """
        return pulumi.get(self, "kms_key_arn")

    @kms_key_arn.setter
    def kms_key_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_arn", value)

    @property
    @pulumi.getter(name="objectKeyPrefix")
    def object_key_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The key prefix of the S3 bucket
        """
        return pulumi.get(self, "object_key_prefix")

    @object_key_prefix.setter
    def object_key_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object_key_prefix", value)

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of an SNS topic to notify
        """
        return pulumi.get(self, "topic_arn")

    @topic_arn.setter
    def topic_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_arn", value)


@pulumi.input_type
class ReceiptRuleSnsActionArgs:
    def __init__(__self__, *,
                 position: pulumi.Input[int],
                 topic_arn: pulumi.Input[str],
                 encoding: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] position: The position of the action in the receipt rule
        :param pulumi.Input[str] topic_arn: The ARN of an SNS topic to notify
        :param pulumi.Input[str] encoding: The encoding to use for the email within the Amazon SNS notification. Default value is `UTF-8`.
        """
        pulumi.set(__self__, "position", position)
        pulumi.set(__self__, "topic_arn", topic_arn)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)

    @property
    @pulumi.getter
    def position(self) -> pulumi.Input[int]:
        """
        The position of the action in the receipt rule
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: pulumi.Input[int]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> pulumi.Input[str]:
        """
        The ARN of an SNS topic to notify
        """
        return pulumi.get(self, "topic_arn")

    @topic_arn.setter
    def topic_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic_arn", value)

    @property
    @pulumi.getter
    def encoding(self) -> Optional[pulumi.Input[str]]:
        """
        The encoding to use for the email within the Amazon SNS notification. Default value is `UTF-8`.
        """
        return pulumi.get(self, "encoding")

    @encoding.setter
    def encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encoding", value)


@pulumi.input_type
class ReceiptRuleStopActionArgs:
    def __init__(__self__, *,
                 position: pulumi.Input[int],
                 scope: pulumi.Input[str],
                 topic_arn: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] position: The position of the action in the receipt rule
        :param pulumi.Input[str] scope: The scope to apply. The only acceptable value is `RuleSet`.
        :param pulumi.Input[str] topic_arn: The ARN of an SNS topic to notify
        """
        pulumi.set(__self__, "position", position)
        pulumi.set(__self__, "scope", scope)
        if topic_arn is not None:
            pulumi.set(__self__, "topic_arn", topic_arn)

    @property
    @pulumi.getter
    def position(self) -> pulumi.Input[int]:
        """
        The position of the action in the receipt rule
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: pulumi.Input[int]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Input[str]:
        """
        The scope to apply. The only acceptable value is `RuleSet`.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: pulumi.Input[str]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of an SNS topic to notify
        """
        return pulumi.get(self, "topic_arn")

    @topic_arn.setter
    def topic_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_arn", value)


@pulumi.input_type
class ReceiptRuleWorkmailActionArgs:
    def __init__(__self__, *,
                 organization_arn: pulumi.Input[str],
                 position: pulumi.Input[int],
                 topic_arn: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] organization_arn: The ARN of the WorkMail organization
        :param pulumi.Input[int] position: The position of the action in the receipt rule
        :param pulumi.Input[str] topic_arn: The ARN of an SNS topic to notify
        """
        pulumi.set(__self__, "organization_arn", organization_arn)
        pulumi.set(__self__, "position", position)
        if topic_arn is not None:
            pulumi.set(__self__, "topic_arn", topic_arn)

    @property
    @pulumi.getter(name="organizationArn")
    def organization_arn(self) -> pulumi.Input[str]:
        """
        The ARN of the WorkMail organization
        """
        return pulumi.get(self, "organization_arn")

    @organization_arn.setter
    def organization_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization_arn", value)

    @property
    @pulumi.getter
    def position(self) -> pulumi.Input[int]:
        """
        The position of the action in the receipt rule
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: pulumi.Input[int]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of an SNS topic to notify
        """
        return pulumi.get(self, "topic_arn")

    @topic_arn.setter
    def topic_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_arn", value)


