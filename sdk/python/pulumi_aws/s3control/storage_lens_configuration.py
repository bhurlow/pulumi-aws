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

__all__ = ['StorageLensConfigurationArgs', 'StorageLensConfiguration']

@pulumi.input_type
class StorageLensConfigurationArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[str],
                 storage_lens_configuration: pulumi.Input['StorageLensConfigurationStorageLensConfigurationArgs'],
                 account_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a StorageLensConfiguration resource.
        :param pulumi.Input[str] config_id: The ID of the S3 Storage Lens configuration.
        :param pulumi.Input['StorageLensConfigurationStorageLensConfigurationArgs'] storage_lens_configuration: The S3 Storage Lens configuration. See Storage Lens Configuration below for more details.
        :param pulumi.Input[str] account_id: The AWS account ID for the S3 Storage Lens configuration. Defaults to automatically determined account ID of the AWS provider.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "storage_lens_configuration", storage_lens_configuration)
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[str]:
        """
        The ID of the S3 Storage Lens configuration.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="storageLensConfiguration")
    def storage_lens_configuration(self) -> pulumi.Input['StorageLensConfigurationStorageLensConfigurationArgs']:
        """
        The S3 Storage Lens configuration. See Storage Lens Configuration below for more details.
        """
        return pulumi.get(self, "storage_lens_configuration")

    @storage_lens_configuration.setter
    def storage_lens_configuration(self, value: pulumi.Input['StorageLensConfigurationStorageLensConfigurationArgs']):
        pulumi.set(self, "storage_lens_configuration", value)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS account ID for the S3 Storage Lens configuration. Defaults to automatically determined account ID of the AWS provider.
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _StorageLensConfigurationState:
    def __init__(__self__, *,
                 account_id: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 config_id: Optional[pulumi.Input[str]] = None,
                 storage_lens_configuration: Optional[pulumi.Input['StorageLensConfigurationStorageLensConfigurationArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering StorageLensConfiguration resources.
        :param pulumi.Input[str] account_id: The AWS account ID for the S3 Storage Lens configuration. Defaults to automatically determined account ID of the AWS provider.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the Amazon Web Services organization.
        :param pulumi.Input[str] config_id: The ID of the S3 Storage Lens configuration.
        :param pulumi.Input['StorageLensConfigurationStorageLensConfigurationArgs'] storage_lens_configuration: The S3 Storage Lens configuration. See Storage Lens Configuration below for more details.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if storage_lens_configuration is not None:
            pulumi.set(__self__, "storage_lens_configuration", storage_lens_configuration)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS account ID for the S3 Storage Lens configuration. Defaults to automatically determined account ID of the AWS provider.
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the Amazon Web Services organization.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the S3 Storage Lens configuration.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="storageLensConfiguration")
    def storage_lens_configuration(self) -> Optional[pulumi.Input['StorageLensConfigurationStorageLensConfigurationArgs']]:
        """
        The S3 Storage Lens configuration. See Storage Lens Configuration below for more details.
        """
        return pulumi.get(self, "storage_lens_configuration")

    @storage_lens_configuration.setter
    def storage_lens_configuration(self, value: Optional[pulumi.Input['StorageLensConfigurationStorageLensConfigurationArgs']]):
        pulumi.set(self, "storage_lens_configuration", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class StorageLensConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 config_id: Optional[pulumi.Input[str]] = None,
                 storage_lens_configuration: Optional[pulumi.Input[pulumi.InputType['StorageLensConfigurationStorageLensConfigurationArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to manage an S3 Storage Lens configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        current = aws.get_caller_identity()
        example = aws.s3control.StorageLensConfiguration("example",
            config_id="example-1",
            storage_lens_configuration=aws.s3control.StorageLensConfigurationStorageLensConfigurationArgs(
                enabled=True,
                account_level=aws.s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelArgs(
                    activity_metrics=aws.s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelActivityMetricsArgs(
                        enabled=True,
                    ),
                    bucket_level=aws.s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelArgs(
                        activity_metrics=aws.s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelActivityMetricsArgs(
                            enabled=True,
                        ),
                    ),
                ),
                data_export=aws.s3control.StorageLensConfigurationStorageLensConfigurationDataExportArgs(
                    cloud_watch_metrics=aws.s3control.StorageLensConfigurationStorageLensConfigurationDataExportCloudWatchMetricsArgs(
                        enabled=True,
                    ),
                    s3_bucket_destination=aws.s3control.StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationArgs(
                        account_id=current.account_id,
                        arn=aws_s3_bucket["target"]["arn"],
                        format="CSV",
                        output_schema_version="V_1",
                        encryption=aws.s3control.StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationEncryptionArgs(
                            sse_s3s=[aws.s3control.StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationEncryptionSseS3Args()],
                        ),
                    ),
                ),
                exclude=aws.s3control.StorageLensConfigurationStorageLensConfigurationExcludeArgs(
                    buckets=[
                        aws_s3_bucket["b1"]["arn"],
                        aws_s3_bucket["b2"]["arn"],
                    ],
                    regions=["us-east-2"],
                ),
            ))
        ```

        ## Import

        S3 Storage Lens configurations can be imported using the `account_id` and `config_id`, separated by a colon (`:`), e.g.

        ```sh
         $ pulumi import aws:s3control/storageLensConfiguration:StorageLensConfiguration example 123456789012:example-1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The AWS account ID for the S3 Storage Lens configuration. Defaults to automatically determined account ID of the AWS provider.
        :param pulumi.Input[str] config_id: The ID of the S3 Storage Lens configuration.
        :param pulumi.Input[pulumi.InputType['StorageLensConfigurationStorageLensConfigurationArgs']] storage_lens_configuration: The S3 Storage Lens configuration. See Storage Lens Configuration below for more details.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StorageLensConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage an S3 Storage Lens configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        current = aws.get_caller_identity()
        example = aws.s3control.StorageLensConfiguration("example",
            config_id="example-1",
            storage_lens_configuration=aws.s3control.StorageLensConfigurationStorageLensConfigurationArgs(
                enabled=True,
                account_level=aws.s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelArgs(
                    activity_metrics=aws.s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelActivityMetricsArgs(
                        enabled=True,
                    ),
                    bucket_level=aws.s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelArgs(
                        activity_metrics=aws.s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelActivityMetricsArgs(
                            enabled=True,
                        ),
                    ),
                ),
                data_export=aws.s3control.StorageLensConfigurationStorageLensConfigurationDataExportArgs(
                    cloud_watch_metrics=aws.s3control.StorageLensConfigurationStorageLensConfigurationDataExportCloudWatchMetricsArgs(
                        enabled=True,
                    ),
                    s3_bucket_destination=aws.s3control.StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationArgs(
                        account_id=current.account_id,
                        arn=aws_s3_bucket["target"]["arn"],
                        format="CSV",
                        output_schema_version="V_1",
                        encryption=aws.s3control.StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationEncryptionArgs(
                            sse_s3s=[aws.s3control.StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationEncryptionSseS3Args()],
                        ),
                    ),
                ),
                exclude=aws.s3control.StorageLensConfigurationStorageLensConfigurationExcludeArgs(
                    buckets=[
                        aws_s3_bucket["b1"]["arn"],
                        aws_s3_bucket["b2"]["arn"],
                    ],
                    regions=["us-east-2"],
                ),
            ))
        ```

        ## Import

        S3 Storage Lens configurations can be imported using the `account_id` and `config_id`, separated by a colon (`:`), e.g.

        ```sh
         $ pulumi import aws:s3control/storageLensConfiguration:StorageLensConfiguration example 123456789012:example-1
        ```

        :param str resource_name: The name of the resource.
        :param StorageLensConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StorageLensConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 config_id: Optional[pulumi.Input[str]] = None,
                 storage_lens_configuration: Optional[pulumi.Input[pulumi.InputType['StorageLensConfigurationStorageLensConfigurationArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StorageLensConfigurationArgs.__new__(StorageLensConfigurationArgs)

            __props__.__dict__["account_id"] = account_id
            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if storage_lens_configuration is None and not opts.urn:
                raise TypeError("Missing required property 'storage_lens_configuration'")
            __props__.__dict__["storage_lens_configuration"] = storage_lens_configuration
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        super(StorageLensConfiguration, __self__).__init__(
            'aws:s3control/storageLensConfiguration:StorageLensConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_id: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            config_id: Optional[pulumi.Input[str]] = None,
            storage_lens_configuration: Optional[pulumi.Input[pulumi.InputType['StorageLensConfigurationStorageLensConfigurationArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'StorageLensConfiguration':
        """
        Get an existing StorageLensConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The AWS account ID for the S3 Storage Lens configuration. Defaults to automatically determined account ID of the AWS provider.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the Amazon Web Services organization.
        :param pulumi.Input[str] config_id: The ID of the S3 Storage Lens configuration.
        :param pulumi.Input[pulumi.InputType['StorageLensConfigurationStorageLensConfigurationArgs']] storage_lens_configuration: The S3 Storage Lens configuration. See Storage Lens Configuration below for more details.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StorageLensConfigurationState.__new__(_StorageLensConfigurationState)

        __props__.__dict__["account_id"] = account_id
        __props__.__dict__["arn"] = arn
        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["storage_lens_configuration"] = storage_lens_configuration
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return StorageLensConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[str]:
        """
        The AWS account ID for the S3 Storage Lens configuration. Defaults to automatically determined account ID of the AWS provider.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the Amazon Web Services organization.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[str]:
        """
        The ID of the S3 Storage Lens configuration.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="storageLensConfiguration")
    def storage_lens_configuration(self) -> pulumi.Output['outputs.StorageLensConfigurationStorageLensConfiguration']:
        """
        The S3 Storage Lens configuration. See Storage Lens Configuration below for more details.
        """
        return pulumi.get(self, "storage_lens_configuration")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

