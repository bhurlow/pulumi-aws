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

__all__ = ['StreamArgs', 'Stream']

@pulumi.input_type
class StreamArgs:
    def __init__(__self__, *,
                 inclusive_start_time: pulumi.Input[str],
                 kinesis_configuration: pulumi.Input['StreamKinesisConfigurationArgs'],
                 ledger_name: pulumi.Input[str],
                 role_arn: pulumi.Input[str],
                 stream_name: pulumi.Input[str],
                 exclusive_end_time: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Stream resource.
        :param pulumi.Input[str] inclusive_start_time: The inclusive start date and time from which to start streaming journal data. This parameter must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: `"2019-06-13T21:36:34Z"`.  This cannot be in the future and must be before `exclusive_end_time`.  If you provide a value that is before the ledger's `CreationDateTime`, QLDB effectively defaults it to the ledger's `CreationDateTime`.
        :param pulumi.Input['StreamKinesisConfigurationArgs'] kinesis_configuration: The configuration settings of the Kinesis Data Streams destination for your stream request. Documented below.
        :param pulumi.Input[str] ledger_name: The name of the QLDB ledger.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for a journal stream to write data records to a Kinesis Data Streams resource.
        :param pulumi.Input[str] stream_name: The name that you want to assign to the QLDB journal stream. User-defined names can help identify and indicate the purpose of a stream.  Your stream name must be unique among other active streams for a given ledger. Stream names have the same naming constraints as ledger names, as defined in the [Amazon QLDB Developer Guide](https://docs.aws.amazon.com/qldb/latest/developerguide/limits.html#limits.naming).
        :param pulumi.Input[str] exclusive_end_time: The exclusive date and time that specifies when the stream ends. If you don't define this parameter, the stream runs indefinitely until you cancel it. It must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: `"2019-06-13T21:36:34Z"`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "inclusive_start_time", inclusive_start_time)
        pulumi.set(__self__, "kinesis_configuration", kinesis_configuration)
        pulumi.set(__self__, "ledger_name", ledger_name)
        pulumi.set(__self__, "role_arn", role_arn)
        pulumi.set(__self__, "stream_name", stream_name)
        if exclusive_end_time is not None:
            pulumi.set(__self__, "exclusive_end_time", exclusive_end_time)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="inclusiveStartTime")
    def inclusive_start_time(self) -> pulumi.Input[str]:
        """
        The inclusive start date and time from which to start streaming journal data. This parameter must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: `"2019-06-13T21:36:34Z"`.  This cannot be in the future and must be before `exclusive_end_time`.  If you provide a value that is before the ledger's `CreationDateTime`, QLDB effectively defaults it to the ledger's `CreationDateTime`.
        """
        return pulumi.get(self, "inclusive_start_time")

    @inclusive_start_time.setter
    def inclusive_start_time(self, value: pulumi.Input[str]):
        pulumi.set(self, "inclusive_start_time", value)

    @property
    @pulumi.getter(name="kinesisConfiguration")
    def kinesis_configuration(self) -> pulumi.Input['StreamKinesisConfigurationArgs']:
        """
        The configuration settings of the Kinesis Data Streams destination for your stream request. Documented below.
        """
        return pulumi.get(self, "kinesis_configuration")

    @kinesis_configuration.setter
    def kinesis_configuration(self, value: pulumi.Input['StreamKinesisConfigurationArgs']):
        pulumi.set(self, "kinesis_configuration", value)

    @property
    @pulumi.getter(name="ledgerName")
    def ledger_name(self) -> pulumi.Input[str]:
        """
        The name of the QLDB ledger.
        """
        return pulumi.get(self, "ledger_name")

    @ledger_name.setter
    def ledger_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "ledger_name", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for a journal stream to write data records to a Kinesis Data Streams resource.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="streamName")
    def stream_name(self) -> pulumi.Input[str]:
        """
        The name that you want to assign to the QLDB journal stream. User-defined names can help identify and indicate the purpose of a stream.  Your stream name must be unique among other active streams for a given ledger. Stream names have the same naming constraints as ledger names, as defined in the [Amazon QLDB Developer Guide](https://docs.aws.amazon.com/qldb/latest/developerguide/limits.html#limits.naming).
        """
        return pulumi.get(self, "stream_name")

    @stream_name.setter
    def stream_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "stream_name", value)

    @property
    @pulumi.getter(name="exclusiveEndTime")
    def exclusive_end_time(self) -> Optional[pulumi.Input[str]]:
        """
        The exclusive date and time that specifies when the stream ends. If you don't define this parameter, the stream runs indefinitely until you cancel it. It must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: `"2019-06-13T21:36:34Z"`.
        """
        return pulumi.get(self, "exclusive_end_time")

    @exclusive_end_time.setter
    def exclusive_end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exclusive_end_time", value)

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
class _StreamState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 exclusive_end_time: Optional[pulumi.Input[str]] = None,
                 inclusive_start_time: Optional[pulumi.Input[str]] = None,
                 kinesis_configuration: Optional[pulumi.Input['StreamKinesisConfigurationArgs']] = None,
                 ledger_name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 stream_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Stream resources.
        :param pulumi.Input[str] arn: The ARN of the QLDB Stream.
        :param pulumi.Input[str] exclusive_end_time: The exclusive date and time that specifies when the stream ends. If you don't define this parameter, the stream runs indefinitely until you cancel it. It must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: `"2019-06-13T21:36:34Z"`.
        :param pulumi.Input[str] inclusive_start_time: The inclusive start date and time from which to start streaming journal data. This parameter must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: `"2019-06-13T21:36:34Z"`.  This cannot be in the future and must be before `exclusive_end_time`.  If you provide a value that is before the ledger's `CreationDateTime`, QLDB effectively defaults it to the ledger's `CreationDateTime`.
        :param pulumi.Input['StreamKinesisConfigurationArgs'] kinesis_configuration: The configuration settings of the Kinesis Data Streams destination for your stream request. Documented below.
        :param pulumi.Input[str] ledger_name: The name of the QLDB ledger.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for a journal stream to write data records to a Kinesis Data Streams resource.
        :param pulumi.Input[str] stream_name: The name that you want to assign to the QLDB journal stream. User-defined names can help identify and indicate the purpose of a stream.  Your stream name must be unique among other active streams for a given ledger. Stream names have the same naming constraints as ledger names, as defined in the [Amazon QLDB Developer Guide](https://docs.aws.amazon.com/qldb/latest/developerguide/limits.html#limits.naming).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if exclusive_end_time is not None:
            pulumi.set(__self__, "exclusive_end_time", exclusive_end_time)
        if inclusive_start_time is not None:
            pulumi.set(__self__, "inclusive_start_time", inclusive_start_time)
        if kinesis_configuration is not None:
            pulumi.set(__self__, "kinesis_configuration", kinesis_configuration)
        if ledger_name is not None:
            pulumi.set(__self__, "ledger_name", ledger_name)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if stream_name is not None:
            pulumi.set(__self__, "stream_name", stream_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
            pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the QLDB Stream.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="exclusiveEndTime")
    def exclusive_end_time(self) -> Optional[pulumi.Input[str]]:
        """
        The exclusive date and time that specifies when the stream ends. If you don't define this parameter, the stream runs indefinitely until you cancel it. It must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: `"2019-06-13T21:36:34Z"`.
        """
        return pulumi.get(self, "exclusive_end_time")

    @exclusive_end_time.setter
    def exclusive_end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exclusive_end_time", value)

    @property
    @pulumi.getter(name="inclusiveStartTime")
    def inclusive_start_time(self) -> Optional[pulumi.Input[str]]:
        """
        The inclusive start date and time from which to start streaming journal data. This parameter must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: `"2019-06-13T21:36:34Z"`.  This cannot be in the future and must be before `exclusive_end_time`.  If you provide a value that is before the ledger's `CreationDateTime`, QLDB effectively defaults it to the ledger's `CreationDateTime`.
        """
        return pulumi.get(self, "inclusive_start_time")

    @inclusive_start_time.setter
    def inclusive_start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "inclusive_start_time", value)

    @property
    @pulumi.getter(name="kinesisConfiguration")
    def kinesis_configuration(self) -> Optional[pulumi.Input['StreamKinesisConfigurationArgs']]:
        """
        The configuration settings of the Kinesis Data Streams destination for your stream request. Documented below.
        """
        return pulumi.get(self, "kinesis_configuration")

    @kinesis_configuration.setter
    def kinesis_configuration(self, value: Optional[pulumi.Input['StreamKinesisConfigurationArgs']]):
        pulumi.set(self, "kinesis_configuration", value)

    @property
    @pulumi.getter(name="ledgerName")
    def ledger_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the QLDB ledger.
        """
        return pulumi.get(self, "ledger_name")

    @ledger_name.setter
    def ledger_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ledger_name", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for a journal stream to write data records to a Kinesis Data Streams resource.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="streamName")
    def stream_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name that you want to assign to the QLDB journal stream. User-defined names can help identify and indicate the purpose of a stream.  Your stream name must be unique among other active streams for a given ledger. Stream names have the same naming constraints as ledger names, as defined in the [Amazon QLDB Developer Guide](https://docs.aws.amazon.com/qldb/latest/developerguide/limits.html#limits.naming).
        """
        return pulumi.get(self, "stream_name")

    @stream_name.setter
    def stream_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stream_name", value)

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
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class Stream(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 exclusive_end_time: Optional[pulumi.Input[str]] = None,
                 inclusive_start_time: Optional[pulumi.Input[str]] = None,
                 kinesis_configuration: Optional[pulumi.Input[pulumi.InputType['StreamKinesisConfigurationArgs']]] = None,
                 ledger_name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 stream_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides an AWS Quantum Ledger Database (QLDB) Stream resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.qldb.Stream("example",
            inclusive_start_time="2021-01-01T00:00:00Z",
            kinesis_configuration=aws.qldb.StreamKinesisConfigurationArgs(
                aggregation_enabled=False,
                stream_arn="arn:aws:kinesis:us-east-1:xxxxxxxxxxxx:stream/example-kinesis-stream",
            ),
            ledger_name="existing-ledger-name",
            role_arn="sample-role-arn",
            stream_name="sample-ledger-stream",
            tags={
                "example": "tag",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] exclusive_end_time: The exclusive date and time that specifies when the stream ends. If you don't define this parameter, the stream runs indefinitely until you cancel it. It must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: `"2019-06-13T21:36:34Z"`.
        :param pulumi.Input[str] inclusive_start_time: The inclusive start date and time from which to start streaming journal data. This parameter must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: `"2019-06-13T21:36:34Z"`.  This cannot be in the future and must be before `exclusive_end_time`.  If you provide a value that is before the ledger's `CreationDateTime`, QLDB effectively defaults it to the ledger's `CreationDateTime`.
        :param pulumi.Input[pulumi.InputType['StreamKinesisConfigurationArgs']] kinesis_configuration: The configuration settings of the Kinesis Data Streams destination for your stream request. Documented below.
        :param pulumi.Input[str] ledger_name: The name of the QLDB ledger.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for a journal stream to write data records to a Kinesis Data Streams resource.
        :param pulumi.Input[str] stream_name: The name that you want to assign to the QLDB journal stream. User-defined names can help identify and indicate the purpose of a stream.  Your stream name must be unique among other active streams for a given ledger. Stream names have the same naming constraints as ledger names, as defined in the [Amazon QLDB Developer Guide](https://docs.aws.amazon.com/qldb/latest/developerguide/limits.html#limits.naming).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StreamArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an AWS Quantum Ledger Database (QLDB) Stream resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.qldb.Stream("example",
            inclusive_start_time="2021-01-01T00:00:00Z",
            kinesis_configuration=aws.qldb.StreamKinesisConfigurationArgs(
                aggregation_enabled=False,
                stream_arn="arn:aws:kinesis:us-east-1:xxxxxxxxxxxx:stream/example-kinesis-stream",
            ),
            ledger_name="existing-ledger-name",
            role_arn="sample-role-arn",
            stream_name="sample-ledger-stream",
            tags={
                "example": "tag",
            })
        ```

        :param str resource_name: The name of the resource.
        :param StreamArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StreamArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 exclusive_end_time: Optional[pulumi.Input[str]] = None,
                 inclusive_start_time: Optional[pulumi.Input[str]] = None,
                 kinesis_configuration: Optional[pulumi.Input[pulumi.InputType['StreamKinesisConfigurationArgs']]] = None,
                 ledger_name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 stream_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StreamArgs.__new__(StreamArgs)

            __props__.__dict__["exclusive_end_time"] = exclusive_end_time
            if inclusive_start_time is None and not opts.urn:
                raise TypeError("Missing required property 'inclusive_start_time'")
            __props__.__dict__["inclusive_start_time"] = inclusive_start_time
            if kinesis_configuration is None and not opts.urn:
                raise TypeError("Missing required property 'kinesis_configuration'")
            __props__.__dict__["kinesis_configuration"] = kinesis_configuration
            if ledger_name is None and not opts.urn:
                raise TypeError("Missing required property 'ledger_name'")
            __props__.__dict__["ledger_name"] = ledger_name
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__.__dict__["role_arn"] = role_arn
            if stream_name is None and not opts.urn:
                raise TypeError("Missing required property 'stream_name'")
            __props__.__dict__["stream_name"] = stream_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Stream, __self__).__init__(
            'aws:qldb/stream:Stream',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            exclusive_end_time: Optional[pulumi.Input[str]] = None,
            inclusive_start_time: Optional[pulumi.Input[str]] = None,
            kinesis_configuration: Optional[pulumi.Input[pulumi.InputType['StreamKinesisConfigurationArgs']]] = None,
            ledger_name: Optional[pulumi.Input[str]] = None,
            role_arn: Optional[pulumi.Input[str]] = None,
            stream_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Stream':
        """
        Get an existing Stream resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the QLDB Stream.
        :param pulumi.Input[str] exclusive_end_time: The exclusive date and time that specifies when the stream ends. If you don't define this parameter, the stream runs indefinitely until you cancel it. It must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: `"2019-06-13T21:36:34Z"`.
        :param pulumi.Input[str] inclusive_start_time: The inclusive start date and time from which to start streaming journal data. This parameter must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: `"2019-06-13T21:36:34Z"`.  This cannot be in the future and must be before `exclusive_end_time`.  If you provide a value that is before the ledger's `CreationDateTime`, QLDB effectively defaults it to the ledger's `CreationDateTime`.
        :param pulumi.Input[pulumi.InputType['StreamKinesisConfigurationArgs']] kinesis_configuration: The configuration settings of the Kinesis Data Streams destination for your stream request. Documented below.
        :param pulumi.Input[str] ledger_name: The name of the QLDB ledger.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for a journal stream to write data records to a Kinesis Data Streams resource.
        :param pulumi.Input[str] stream_name: The name that you want to assign to the QLDB journal stream. User-defined names can help identify and indicate the purpose of a stream.  Your stream name must be unique among other active streams for a given ledger. Stream names have the same naming constraints as ledger names, as defined in the [Amazon QLDB Developer Guide](https://docs.aws.amazon.com/qldb/latest/developerguide/limits.html#limits.naming).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StreamState.__new__(_StreamState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["exclusive_end_time"] = exclusive_end_time
        __props__.__dict__["inclusive_start_time"] = inclusive_start_time
        __props__.__dict__["kinesis_configuration"] = kinesis_configuration
        __props__.__dict__["ledger_name"] = ledger_name
        __props__.__dict__["role_arn"] = role_arn
        __props__.__dict__["stream_name"] = stream_name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return Stream(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the QLDB Stream.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="exclusiveEndTime")
    def exclusive_end_time(self) -> pulumi.Output[Optional[str]]:
        """
        The exclusive date and time that specifies when the stream ends. If you don't define this parameter, the stream runs indefinitely until you cancel it. It must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: `"2019-06-13T21:36:34Z"`.
        """
        return pulumi.get(self, "exclusive_end_time")

    @property
    @pulumi.getter(name="inclusiveStartTime")
    def inclusive_start_time(self) -> pulumi.Output[str]:
        """
        The inclusive start date and time from which to start streaming journal data. This parameter must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: `"2019-06-13T21:36:34Z"`.  This cannot be in the future and must be before `exclusive_end_time`.  If you provide a value that is before the ledger's `CreationDateTime`, QLDB effectively defaults it to the ledger's `CreationDateTime`.
        """
        return pulumi.get(self, "inclusive_start_time")

    @property
    @pulumi.getter(name="kinesisConfiguration")
    def kinesis_configuration(self) -> pulumi.Output['outputs.StreamKinesisConfiguration']:
        """
        The configuration settings of the Kinesis Data Streams destination for your stream request. Documented below.
        """
        return pulumi.get(self, "kinesis_configuration")

    @property
    @pulumi.getter(name="ledgerName")
    def ledger_name(self) -> pulumi.Output[str]:
        """
        The name of the QLDB ledger.
        """
        return pulumi.get(self, "ledger_name")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for a journal stream to write data records to a Kinesis Data Streams resource.
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="streamName")
    def stream_name(self) -> pulumi.Output[str]:
        """
        The name that you want to assign to the QLDB journal stream. User-defined names can help identify and indicate the purpose of a stream.  Your stream name must be unique among other active streams for a given ledger. Stream names have the same naming constraints as ledger names, as defined in the [Amazon QLDB Developer Guide](https://docs.aws.amazon.com/qldb/latest/developerguide/limits.html#limits.naming).
        """
        return pulumi.get(self, "stream_name")

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
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

