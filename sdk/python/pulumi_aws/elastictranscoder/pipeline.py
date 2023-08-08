# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['PipelineArgs', 'Pipeline']

@pulumi.input_type
class PipelineArgs:
    def __init__(__self__, *,
                 input_bucket: pulumi.Input[str],
                 role: pulumi.Input[str],
                 aws_kms_key_arn: Optional[pulumi.Input[str]] = None,
                 content_config: Optional[pulumi.Input['PipelineContentConfigArgs']] = None,
                 content_config_permissions: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineContentConfigPermissionArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notifications: Optional[pulumi.Input['PipelineNotificationsArgs']] = None,
                 output_bucket: Optional[pulumi.Input[str]] = None,
                 thumbnail_config: Optional[pulumi.Input['PipelineThumbnailConfigArgs']] = None,
                 thumbnail_config_permissions: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineThumbnailConfigPermissionArgs']]]] = None):
        """
        The set of arguments for constructing a Pipeline resource.
        :param pulumi.Input[str] input_bucket: The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
        :param pulumi.Input[str] role: The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
        :param pulumi.Input[str] aws_kms_key_arn: The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
        :param pulumi.Input['PipelineContentConfigArgs'] content_config: The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
        :param pulumi.Input[Sequence[pulumi.Input['PipelineContentConfigPermissionArgs']]] content_config_permissions: The permissions for the `content_config` object. (documented below)
        :param pulumi.Input[str] name: The name of the pipeline. Maximum 40 characters
        :param pulumi.Input['PipelineNotificationsArgs'] notifications: The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
        :param pulumi.Input[str] output_bucket: The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
        :param pulumi.Input['PipelineThumbnailConfigArgs'] thumbnail_config: The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
        :param pulumi.Input[Sequence[pulumi.Input['PipelineThumbnailConfigPermissionArgs']]] thumbnail_config_permissions: The permissions for the `thumbnail_config` object. (documented below)
               
               The `content_config` object specifies information about the Amazon S3 bucket in
               which you want Elastic Transcoder to save transcoded files and playlists: which
               bucket to use, and the storage class that you want to assign to the files. If
               you specify values for `content_config`, you must also specify values for
               `thumbnail_config`. If you specify values for `content_config` and
               `thumbnail_config`, omit the `output_bucket` object.
        """
        pulumi.set(__self__, "input_bucket", input_bucket)
        pulumi.set(__self__, "role", role)
        if aws_kms_key_arn is not None:
            pulumi.set(__self__, "aws_kms_key_arn", aws_kms_key_arn)
        if content_config is not None:
            pulumi.set(__self__, "content_config", content_config)
        if content_config_permissions is not None:
            pulumi.set(__self__, "content_config_permissions", content_config_permissions)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notifications is not None:
            pulumi.set(__self__, "notifications", notifications)
        if output_bucket is not None:
            pulumi.set(__self__, "output_bucket", output_bucket)
        if thumbnail_config is not None:
            pulumi.set(__self__, "thumbnail_config", thumbnail_config)
        if thumbnail_config_permissions is not None:
            pulumi.set(__self__, "thumbnail_config_permissions", thumbnail_config_permissions)

    @property
    @pulumi.getter(name="inputBucket")
    def input_bucket(self) -> pulumi.Input[str]:
        """
        The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
        """
        return pulumi.get(self, "input_bucket")

    @input_bucket.setter
    def input_bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "input_bucket", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="awsKmsKeyArn")
    def aws_kms_key_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
        """
        return pulumi.get(self, "aws_kms_key_arn")

    @aws_kms_key_arn.setter
    def aws_kms_key_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_kms_key_arn", value)

    @property
    @pulumi.getter(name="contentConfig")
    def content_config(self) -> Optional[pulumi.Input['PipelineContentConfigArgs']]:
        """
        The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
        """
        return pulumi.get(self, "content_config")

    @content_config.setter
    def content_config(self, value: Optional[pulumi.Input['PipelineContentConfigArgs']]):
        pulumi.set(self, "content_config", value)

    @property
    @pulumi.getter(name="contentConfigPermissions")
    def content_config_permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PipelineContentConfigPermissionArgs']]]]:
        """
        The permissions for the `content_config` object. (documented below)
        """
        return pulumi.get(self, "content_config_permissions")

    @content_config_permissions.setter
    def content_config_permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineContentConfigPermissionArgs']]]]):
        pulumi.set(self, "content_config_permissions", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the pipeline. Maximum 40 characters
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def notifications(self) -> Optional[pulumi.Input['PipelineNotificationsArgs']]:
        """
        The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
        """
        return pulumi.get(self, "notifications")

    @notifications.setter
    def notifications(self, value: Optional[pulumi.Input['PipelineNotificationsArgs']]):
        pulumi.set(self, "notifications", value)

    @property
    @pulumi.getter(name="outputBucket")
    def output_bucket(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
        """
        return pulumi.get(self, "output_bucket")

    @output_bucket.setter
    def output_bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_bucket", value)

    @property
    @pulumi.getter(name="thumbnailConfig")
    def thumbnail_config(self) -> Optional[pulumi.Input['PipelineThumbnailConfigArgs']]:
        """
        The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
        """
        return pulumi.get(self, "thumbnail_config")

    @thumbnail_config.setter
    def thumbnail_config(self, value: Optional[pulumi.Input['PipelineThumbnailConfigArgs']]):
        pulumi.set(self, "thumbnail_config", value)

    @property
    @pulumi.getter(name="thumbnailConfigPermissions")
    def thumbnail_config_permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PipelineThumbnailConfigPermissionArgs']]]]:
        """
        The permissions for the `thumbnail_config` object. (documented below)

        The `content_config` object specifies information about the Amazon S3 bucket in
        which you want Elastic Transcoder to save transcoded files and playlists: which
        bucket to use, and the storage class that you want to assign to the files. If
        you specify values for `content_config`, you must also specify values for
        `thumbnail_config`. If you specify values for `content_config` and
        `thumbnail_config`, omit the `output_bucket` object.
        """
        return pulumi.get(self, "thumbnail_config_permissions")

    @thumbnail_config_permissions.setter
    def thumbnail_config_permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineThumbnailConfigPermissionArgs']]]]):
        pulumi.set(self, "thumbnail_config_permissions", value)


@pulumi.input_type
class _PipelineState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 aws_kms_key_arn: Optional[pulumi.Input[str]] = None,
                 content_config: Optional[pulumi.Input['PipelineContentConfigArgs']] = None,
                 content_config_permissions: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineContentConfigPermissionArgs']]]] = None,
                 input_bucket: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notifications: Optional[pulumi.Input['PipelineNotificationsArgs']] = None,
                 output_bucket: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 thumbnail_config: Optional[pulumi.Input['PipelineThumbnailConfigArgs']] = None,
                 thumbnail_config_permissions: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineThumbnailConfigPermissionArgs']]]] = None):
        """
        Input properties used for looking up and filtering Pipeline resources.
        :param pulumi.Input[str] arn: The ARN of the Elastictranscoder pipeline.
        :param pulumi.Input[str] aws_kms_key_arn: The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
        :param pulumi.Input['PipelineContentConfigArgs'] content_config: The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
        :param pulumi.Input[Sequence[pulumi.Input['PipelineContentConfigPermissionArgs']]] content_config_permissions: The permissions for the `content_config` object. (documented below)
        :param pulumi.Input[str] input_bucket: The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
        :param pulumi.Input[str] name: The name of the pipeline. Maximum 40 characters
        :param pulumi.Input['PipelineNotificationsArgs'] notifications: The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
        :param pulumi.Input[str] output_bucket: The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
        :param pulumi.Input[str] role: The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
        :param pulumi.Input['PipelineThumbnailConfigArgs'] thumbnail_config: The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
        :param pulumi.Input[Sequence[pulumi.Input['PipelineThumbnailConfigPermissionArgs']]] thumbnail_config_permissions: The permissions for the `thumbnail_config` object. (documented below)
               
               The `content_config` object specifies information about the Amazon S3 bucket in
               which you want Elastic Transcoder to save transcoded files and playlists: which
               bucket to use, and the storage class that you want to assign to the files. If
               you specify values for `content_config`, you must also specify values for
               `thumbnail_config`. If you specify values for `content_config` and
               `thumbnail_config`, omit the `output_bucket` object.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if aws_kms_key_arn is not None:
            pulumi.set(__self__, "aws_kms_key_arn", aws_kms_key_arn)
        if content_config is not None:
            pulumi.set(__self__, "content_config", content_config)
        if content_config_permissions is not None:
            pulumi.set(__self__, "content_config_permissions", content_config_permissions)
        if input_bucket is not None:
            pulumi.set(__self__, "input_bucket", input_bucket)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notifications is not None:
            pulumi.set(__self__, "notifications", notifications)
        if output_bucket is not None:
            pulumi.set(__self__, "output_bucket", output_bucket)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if thumbnail_config is not None:
            pulumi.set(__self__, "thumbnail_config", thumbnail_config)
        if thumbnail_config_permissions is not None:
            pulumi.set(__self__, "thumbnail_config_permissions", thumbnail_config_permissions)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the Elastictranscoder pipeline.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="awsKmsKeyArn")
    def aws_kms_key_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
        """
        return pulumi.get(self, "aws_kms_key_arn")

    @aws_kms_key_arn.setter
    def aws_kms_key_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_kms_key_arn", value)

    @property
    @pulumi.getter(name="contentConfig")
    def content_config(self) -> Optional[pulumi.Input['PipelineContentConfigArgs']]:
        """
        The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
        """
        return pulumi.get(self, "content_config")

    @content_config.setter
    def content_config(self, value: Optional[pulumi.Input['PipelineContentConfigArgs']]):
        pulumi.set(self, "content_config", value)

    @property
    @pulumi.getter(name="contentConfigPermissions")
    def content_config_permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PipelineContentConfigPermissionArgs']]]]:
        """
        The permissions for the `content_config` object. (documented below)
        """
        return pulumi.get(self, "content_config_permissions")

    @content_config_permissions.setter
    def content_config_permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineContentConfigPermissionArgs']]]]):
        pulumi.set(self, "content_config_permissions", value)

    @property
    @pulumi.getter(name="inputBucket")
    def input_bucket(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
        """
        return pulumi.get(self, "input_bucket")

    @input_bucket.setter
    def input_bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "input_bucket", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the pipeline. Maximum 40 characters
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def notifications(self) -> Optional[pulumi.Input['PipelineNotificationsArgs']]:
        """
        The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
        """
        return pulumi.get(self, "notifications")

    @notifications.setter
    def notifications(self, value: Optional[pulumi.Input['PipelineNotificationsArgs']]):
        pulumi.set(self, "notifications", value)

    @property
    @pulumi.getter(name="outputBucket")
    def output_bucket(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
        """
        return pulumi.get(self, "output_bucket")

    @output_bucket.setter
    def output_bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_bucket", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="thumbnailConfig")
    def thumbnail_config(self) -> Optional[pulumi.Input['PipelineThumbnailConfigArgs']]:
        """
        The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
        """
        return pulumi.get(self, "thumbnail_config")

    @thumbnail_config.setter
    def thumbnail_config(self, value: Optional[pulumi.Input['PipelineThumbnailConfigArgs']]):
        pulumi.set(self, "thumbnail_config", value)

    @property
    @pulumi.getter(name="thumbnailConfigPermissions")
    def thumbnail_config_permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PipelineThumbnailConfigPermissionArgs']]]]:
        """
        The permissions for the `thumbnail_config` object. (documented below)

        The `content_config` object specifies information about the Amazon S3 bucket in
        which you want Elastic Transcoder to save transcoded files and playlists: which
        bucket to use, and the storage class that you want to assign to the files. If
        you specify values for `content_config`, you must also specify values for
        `thumbnail_config`. If you specify values for `content_config` and
        `thumbnail_config`, omit the `output_bucket` object.
        """
        return pulumi.get(self, "thumbnail_config_permissions")

    @thumbnail_config_permissions.setter
    def thumbnail_config_permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineThumbnailConfigPermissionArgs']]]]):
        pulumi.set(self, "thumbnail_config_permissions", value)


class Pipeline(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_kms_key_arn: Optional[pulumi.Input[str]] = None,
                 content_config: Optional[pulumi.Input[pulumi.InputType['PipelineContentConfigArgs']]] = None,
                 content_config_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineContentConfigPermissionArgs']]]]] = None,
                 input_bucket: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notifications: Optional[pulumi.Input[pulumi.InputType['PipelineNotificationsArgs']]] = None,
                 output_bucket: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 thumbnail_config: Optional[pulumi.Input[pulumi.InputType['PipelineThumbnailConfigArgs']]] = None,
                 thumbnail_config_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineThumbnailConfigPermissionArgs']]]]] = None,
                 __props__=None):
        """
        Provides an Elastic Transcoder pipeline resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        bar = aws.elastictranscoder.Pipeline("bar",
            input_bucket=aws_s3_bucket["input_bucket"]["id"],
            role=aws_iam_role["test_role"]["arn"],
            content_config=aws.elastictranscoder.PipelineContentConfigArgs(
                bucket=aws_s3_bucket["content_bucket"]["id"],
                storage_class="Standard",
            ),
            thumbnail_config=aws.elastictranscoder.PipelineThumbnailConfigArgs(
                bucket=aws_s3_bucket["thumb_bucket"]["id"],
                storage_class="Standard",
            ))
        ```

        ## Import

        terraform import {

         to = aws_elastictranscoder_pipeline.basic_pipeline

         id = "1407981661351-cttk8b" } Using `pulumi import`, import Elastic Transcoder pipelines using the `id`. For exampleconsole % pulumi import aws_elastictranscoder_pipeline.basic_pipeline 1407981661351-cttk8b

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_kms_key_arn: The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
        :param pulumi.Input[pulumi.InputType['PipelineContentConfigArgs']] content_config: The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineContentConfigPermissionArgs']]]] content_config_permissions: The permissions for the `content_config` object. (documented below)
        :param pulumi.Input[str] input_bucket: The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
        :param pulumi.Input[str] name: The name of the pipeline. Maximum 40 characters
        :param pulumi.Input[pulumi.InputType['PipelineNotificationsArgs']] notifications: The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
        :param pulumi.Input[str] output_bucket: The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
        :param pulumi.Input[str] role: The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
        :param pulumi.Input[pulumi.InputType['PipelineThumbnailConfigArgs']] thumbnail_config: The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineThumbnailConfigPermissionArgs']]]] thumbnail_config_permissions: The permissions for the `thumbnail_config` object. (documented below)
               
               The `content_config` object specifies information about the Amazon S3 bucket in
               which you want Elastic Transcoder to save transcoded files and playlists: which
               bucket to use, and the storage class that you want to assign to the files. If
               you specify values for `content_config`, you must also specify values for
               `thumbnail_config`. If you specify values for `content_config` and
               `thumbnail_config`, omit the `output_bucket` object.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PipelineArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Elastic Transcoder pipeline resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        bar = aws.elastictranscoder.Pipeline("bar",
            input_bucket=aws_s3_bucket["input_bucket"]["id"],
            role=aws_iam_role["test_role"]["arn"],
            content_config=aws.elastictranscoder.PipelineContentConfigArgs(
                bucket=aws_s3_bucket["content_bucket"]["id"],
                storage_class="Standard",
            ),
            thumbnail_config=aws.elastictranscoder.PipelineThumbnailConfigArgs(
                bucket=aws_s3_bucket["thumb_bucket"]["id"],
                storage_class="Standard",
            ))
        ```

        ## Import

        terraform import {

         to = aws_elastictranscoder_pipeline.basic_pipeline

         id = "1407981661351-cttk8b" } Using `pulumi import`, import Elastic Transcoder pipelines using the `id`. For exampleconsole % pulumi import aws_elastictranscoder_pipeline.basic_pipeline 1407981661351-cttk8b

        :param str resource_name: The name of the resource.
        :param PipelineArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PipelineArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_kms_key_arn: Optional[pulumi.Input[str]] = None,
                 content_config: Optional[pulumi.Input[pulumi.InputType['PipelineContentConfigArgs']]] = None,
                 content_config_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineContentConfigPermissionArgs']]]]] = None,
                 input_bucket: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notifications: Optional[pulumi.Input[pulumi.InputType['PipelineNotificationsArgs']]] = None,
                 output_bucket: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 thumbnail_config: Optional[pulumi.Input[pulumi.InputType['PipelineThumbnailConfigArgs']]] = None,
                 thumbnail_config_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineThumbnailConfigPermissionArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PipelineArgs.__new__(PipelineArgs)

            __props__.__dict__["aws_kms_key_arn"] = aws_kms_key_arn
            __props__.__dict__["content_config"] = content_config
            __props__.__dict__["content_config_permissions"] = content_config_permissions
            if input_bucket is None and not opts.urn:
                raise TypeError("Missing required property 'input_bucket'")
            __props__.__dict__["input_bucket"] = input_bucket
            __props__.__dict__["name"] = name
            __props__.__dict__["notifications"] = notifications
            __props__.__dict__["output_bucket"] = output_bucket
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["thumbnail_config"] = thumbnail_config
            __props__.__dict__["thumbnail_config_permissions"] = thumbnail_config_permissions
            __props__.__dict__["arn"] = None
        super(Pipeline, __self__).__init__(
            'aws:elastictranscoder/pipeline:Pipeline',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            aws_kms_key_arn: Optional[pulumi.Input[str]] = None,
            content_config: Optional[pulumi.Input[pulumi.InputType['PipelineContentConfigArgs']]] = None,
            content_config_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineContentConfigPermissionArgs']]]]] = None,
            input_bucket: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notifications: Optional[pulumi.Input[pulumi.InputType['PipelineNotificationsArgs']]] = None,
            output_bucket: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            thumbnail_config: Optional[pulumi.Input[pulumi.InputType['PipelineThumbnailConfigArgs']]] = None,
            thumbnail_config_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineThumbnailConfigPermissionArgs']]]]] = None) -> 'Pipeline':
        """
        Get an existing Pipeline resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the Elastictranscoder pipeline.
        :param pulumi.Input[str] aws_kms_key_arn: The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
        :param pulumi.Input[pulumi.InputType['PipelineContentConfigArgs']] content_config: The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineContentConfigPermissionArgs']]]] content_config_permissions: The permissions for the `content_config` object. (documented below)
        :param pulumi.Input[str] input_bucket: The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
        :param pulumi.Input[str] name: The name of the pipeline. Maximum 40 characters
        :param pulumi.Input[pulumi.InputType['PipelineNotificationsArgs']] notifications: The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
        :param pulumi.Input[str] output_bucket: The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
        :param pulumi.Input[str] role: The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
        :param pulumi.Input[pulumi.InputType['PipelineThumbnailConfigArgs']] thumbnail_config: The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineThumbnailConfigPermissionArgs']]]] thumbnail_config_permissions: The permissions for the `thumbnail_config` object. (documented below)
               
               The `content_config` object specifies information about the Amazon S3 bucket in
               which you want Elastic Transcoder to save transcoded files and playlists: which
               bucket to use, and the storage class that you want to assign to the files. If
               you specify values for `content_config`, you must also specify values for
               `thumbnail_config`. If you specify values for `content_config` and
               `thumbnail_config`, omit the `output_bucket` object.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PipelineState.__new__(_PipelineState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["aws_kms_key_arn"] = aws_kms_key_arn
        __props__.__dict__["content_config"] = content_config
        __props__.__dict__["content_config_permissions"] = content_config_permissions
        __props__.__dict__["input_bucket"] = input_bucket
        __props__.__dict__["name"] = name
        __props__.__dict__["notifications"] = notifications
        __props__.__dict__["output_bucket"] = output_bucket
        __props__.__dict__["role"] = role
        __props__.__dict__["thumbnail_config"] = thumbnail_config
        __props__.__dict__["thumbnail_config_permissions"] = thumbnail_config_permissions
        return Pipeline(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the Elastictranscoder pipeline.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsKmsKeyArn")
    def aws_kms_key_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
        """
        return pulumi.get(self, "aws_kms_key_arn")

    @property
    @pulumi.getter(name="contentConfig")
    def content_config(self) -> pulumi.Output['outputs.PipelineContentConfig']:
        """
        The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
        """
        return pulumi.get(self, "content_config")

    @property
    @pulumi.getter(name="contentConfigPermissions")
    def content_config_permissions(self) -> pulumi.Output[Optional[Sequence['outputs.PipelineContentConfigPermission']]]:
        """
        The permissions for the `content_config` object. (documented below)
        """
        return pulumi.get(self, "content_config_permissions")

    @property
    @pulumi.getter(name="inputBucket")
    def input_bucket(self) -> pulumi.Output[str]:
        """
        The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
        """
        return pulumi.get(self, "input_bucket")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the pipeline. Maximum 40 characters
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notifications(self) -> pulumi.Output[Optional['outputs.PipelineNotifications']]:
        """
        The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
        """
        return pulumi.get(self, "notifications")

    @property
    @pulumi.getter(name="outputBucket")
    def output_bucket(self) -> pulumi.Output[str]:
        """
        The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
        """
        return pulumi.get(self, "output_bucket")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="thumbnailConfig")
    def thumbnail_config(self) -> pulumi.Output['outputs.PipelineThumbnailConfig']:
        """
        The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
        """
        return pulumi.get(self, "thumbnail_config")

    @property
    @pulumi.getter(name="thumbnailConfigPermissions")
    def thumbnail_config_permissions(self) -> pulumi.Output[Optional[Sequence['outputs.PipelineThumbnailConfigPermission']]]:
        """
        The permissions for the `thumbnail_config` object. (documented below)

        The `content_config` object specifies information about the Amazon S3 bucket in
        which you want Elastic Transcoder to save transcoded files and playlists: which
        bucket to use, and the storage class that you want to assign to the files. If
        you specify values for `content_config`, you must also specify values for
        `thumbnail_config`. If you specify values for `content_config` and
        `thumbnail_config`, omit the `output_bucket` object.
        """
        return pulumi.get(self, "thumbnail_config_permissions")

