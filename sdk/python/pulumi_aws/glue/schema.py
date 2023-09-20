# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SchemaArgs', 'Schema']

@pulumi.input_type
class SchemaArgs:
    def __init__(__self__, *,
                 compatibility: pulumi.Input[str],
                 data_format: pulumi.Input[str],
                 schema_definition: pulumi.Input[str],
                 schema_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 registry_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Schema resource.
        :param pulumi.Input[str] compatibility: The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
        :param pulumi.Input[str] data_format: The data format of the schema definition. Valid values are `AVRO`, `JSON` and `PROTOBUF`.
        :param pulumi.Input[str] schema_definition: The schema definition using the `data_format` setting for `schema_name`.
        :param pulumi.Input[str] schema_name: The Name of the schema.
        :param pulumi.Input[str] description: A description of the schema.
        :param pulumi.Input[str] registry_arn: The ARN of the Glue Registry to create the schema in.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "compatibility", compatibility)
        pulumi.set(__self__, "data_format", data_format)
        pulumi.set(__self__, "schema_definition", schema_definition)
        pulumi.set(__self__, "schema_name", schema_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if registry_arn is not None:
            pulumi.set(__self__, "registry_arn", registry_arn)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def compatibility(self) -> pulumi.Input[str]:
        """
        The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
        """
        return pulumi.get(self, "compatibility")

    @compatibility.setter
    def compatibility(self, value: pulumi.Input[str]):
        pulumi.set(self, "compatibility", value)

    @property
    @pulumi.getter(name="dataFormat")
    def data_format(self) -> pulumi.Input[str]:
        """
        The data format of the schema definition. Valid values are `AVRO`, `JSON` and `PROTOBUF`.
        """
        return pulumi.get(self, "data_format")

    @data_format.setter
    def data_format(self, value: pulumi.Input[str]):
        pulumi.set(self, "data_format", value)

    @property
    @pulumi.getter(name="schemaDefinition")
    def schema_definition(self) -> pulumi.Input[str]:
        """
        The schema definition using the `data_format` setting for `schema_name`.
        """
        return pulumi.get(self, "schema_definition")

    @schema_definition.setter
    def schema_definition(self, value: pulumi.Input[str]):
        pulumi.set(self, "schema_definition", value)

    @property
    @pulumi.getter(name="schemaName")
    def schema_name(self) -> pulumi.Input[str]:
        """
        The Name of the schema.
        """
        return pulumi.get(self, "schema_name")

    @schema_name.setter
    def schema_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "schema_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the schema.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="registryArn")
    def registry_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the Glue Registry to create the schema in.
        """
        return pulumi.get(self, "registry_arn")

    @registry_arn.setter
    def registry_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registry_arn", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _SchemaState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 compatibility: Optional[pulumi.Input[str]] = None,
                 data_format: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 latest_schema_version: Optional[pulumi.Input[int]] = None,
                 next_schema_version: Optional[pulumi.Input[int]] = None,
                 registry_arn: Optional[pulumi.Input[str]] = None,
                 registry_name: Optional[pulumi.Input[str]] = None,
                 schema_checkpoint: Optional[pulumi.Input[int]] = None,
                 schema_definition: Optional[pulumi.Input[str]] = None,
                 schema_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Schema resources.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the schema.
        :param pulumi.Input[str] compatibility: The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
        :param pulumi.Input[str] data_format: The data format of the schema definition. Valid values are `AVRO`, `JSON` and `PROTOBUF`.
        :param pulumi.Input[str] description: A description of the schema.
        :param pulumi.Input[int] latest_schema_version: The latest version of the schema associated with the returned schema definition.
        :param pulumi.Input[int] next_schema_version: The next version of the schema associated with the returned schema definition.
        :param pulumi.Input[str] registry_arn: The ARN of the Glue Registry to create the schema in.
        :param pulumi.Input[str] registry_name: The name of the Glue Registry.
        :param pulumi.Input[int] schema_checkpoint: The version number of the checkpoint (the last time the compatibility mode was changed).
        :param pulumi.Input[str] schema_definition: The schema definition using the `data_format` setting for `schema_name`.
        :param pulumi.Input[str] schema_name: The Name of the schema.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if compatibility is not None:
            pulumi.set(__self__, "compatibility", compatibility)
        if data_format is not None:
            pulumi.set(__self__, "data_format", data_format)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if latest_schema_version is not None:
            pulumi.set(__self__, "latest_schema_version", latest_schema_version)
        if next_schema_version is not None:
            pulumi.set(__self__, "next_schema_version", next_schema_version)
        if registry_arn is not None:
            pulumi.set(__self__, "registry_arn", registry_arn)
        if registry_name is not None:
            pulumi.set(__self__, "registry_name", registry_name)
        if schema_checkpoint is not None:
            pulumi.set(__self__, "schema_checkpoint", schema_checkpoint)
        if schema_definition is not None:
            pulumi.set(__self__, "schema_definition", schema_definition)
        if schema_name is not None:
            pulumi.set(__self__, "schema_name", schema_name)
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
        Amazon Resource Name (ARN) of the schema.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def compatibility(self) -> Optional[pulumi.Input[str]]:
        """
        The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
        """
        return pulumi.get(self, "compatibility")

    @compatibility.setter
    def compatibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compatibility", value)

    @property
    @pulumi.getter(name="dataFormat")
    def data_format(self) -> Optional[pulumi.Input[str]]:
        """
        The data format of the schema definition. Valid values are `AVRO`, `JSON` and `PROTOBUF`.
        """
        return pulumi.get(self, "data_format")

    @data_format.setter
    def data_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_format", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the schema.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="latestSchemaVersion")
    def latest_schema_version(self) -> Optional[pulumi.Input[int]]:
        """
        The latest version of the schema associated with the returned schema definition.
        """
        return pulumi.get(self, "latest_schema_version")

    @latest_schema_version.setter
    def latest_schema_version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "latest_schema_version", value)

    @property
    @pulumi.getter(name="nextSchemaVersion")
    def next_schema_version(self) -> Optional[pulumi.Input[int]]:
        """
        The next version of the schema associated with the returned schema definition.
        """
        return pulumi.get(self, "next_schema_version")

    @next_schema_version.setter
    def next_schema_version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "next_schema_version", value)

    @property
    @pulumi.getter(name="registryArn")
    def registry_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the Glue Registry to create the schema in.
        """
        return pulumi.get(self, "registry_arn")

    @registry_arn.setter
    def registry_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registry_arn", value)

    @property
    @pulumi.getter(name="registryName")
    def registry_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Glue Registry.
        """
        return pulumi.get(self, "registry_name")

    @registry_name.setter
    def registry_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registry_name", value)

    @property
    @pulumi.getter(name="schemaCheckpoint")
    def schema_checkpoint(self) -> Optional[pulumi.Input[int]]:
        """
        The version number of the checkpoint (the last time the compatibility mode was changed).
        """
        return pulumi.get(self, "schema_checkpoint")

    @schema_checkpoint.setter
    def schema_checkpoint(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "schema_checkpoint", value)

    @property
    @pulumi.getter(name="schemaDefinition")
    def schema_definition(self) -> Optional[pulumi.Input[str]]:
        """
        The schema definition using the `data_format` setting for `schema_name`.
        """
        return pulumi.get(self, "schema_definition")

    @schema_definition.setter
    def schema_definition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schema_definition", value)

    @property
    @pulumi.getter(name="schemaName")
    def schema_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Name of the schema.
        """
        return pulumi.get(self, "schema_name")

    @schema_name.setter
    def schema_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schema_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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


class Schema(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compatibility: Optional[pulumi.Input[str]] = None,
                 data_format: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 registry_arn: Optional[pulumi.Input[str]] = None,
                 schema_definition: Optional[pulumi.Input[str]] = None,
                 schema_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a Glue Schema resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.glue.Schema("example",
            schema_name="example",
            registry_arn=aws_glue_registry["test"]["arn"],
            data_format="AVRO",
            compatibility="NONE",
            schema_definition="{\\"type\\": \\"record\\", \\"name\\": \\"r1\\", \\"fields\\": [ {\\"name\\": \\"f1\\", \\"type\\": \\"int\\"}, {\\"name\\": \\"f2\\", \\"type\\": \\"string\\"} ]}")
        ```

        ## Import

        Using `pulumi import`, import Glue Registries using `arn`. For example:

        ```sh
         $ pulumi import aws:glue/schema:Schema example arn:aws:glue:us-west-2:123456789012:schema/example/example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compatibility: The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
        :param pulumi.Input[str] data_format: The data format of the schema definition. Valid values are `AVRO`, `JSON` and `PROTOBUF`.
        :param pulumi.Input[str] description: A description of the schema.
        :param pulumi.Input[str] registry_arn: The ARN of the Glue Registry to create the schema in.
        :param pulumi.Input[str] schema_definition: The schema definition using the `data_format` setting for `schema_name`.
        :param pulumi.Input[str] schema_name: The Name of the schema.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SchemaArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Glue Schema resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.glue.Schema("example",
            schema_name="example",
            registry_arn=aws_glue_registry["test"]["arn"],
            data_format="AVRO",
            compatibility="NONE",
            schema_definition="{\\"type\\": \\"record\\", \\"name\\": \\"r1\\", \\"fields\\": [ {\\"name\\": \\"f1\\", \\"type\\": \\"int\\"}, {\\"name\\": \\"f2\\", \\"type\\": \\"string\\"} ]}")
        ```

        ## Import

        Using `pulumi import`, import Glue Registries using `arn`. For example:

        ```sh
         $ pulumi import aws:glue/schema:Schema example arn:aws:glue:us-west-2:123456789012:schema/example/example
        ```

        :param str resource_name: The name of the resource.
        :param SchemaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SchemaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compatibility: Optional[pulumi.Input[str]] = None,
                 data_format: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 registry_arn: Optional[pulumi.Input[str]] = None,
                 schema_definition: Optional[pulumi.Input[str]] = None,
                 schema_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SchemaArgs.__new__(SchemaArgs)

            if compatibility is None and not opts.urn:
                raise TypeError("Missing required property 'compatibility'")
            __props__.__dict__["compatibility"] = compatibility
            if data_format is None and not opts.urn:
                raise TypeError("Missing required property 'data_format'")
            __props__.__dict__["data_format"] = data_format
            __props__.__dict__["description"] = description
            __props__.__dict__["registry_arn"] = registry_arn
            if schema_definition is None and not opts.urn:
                raise TypeError("Missing required property 'schema_definition'")
            __props__.__dict__["schema_definition"] = schema_definition
            if schema_name is None and not opts.urn:
                raise TypeError("Missing required property 'schema_name'")
            __props__.__dict__["schema_name"] = schema_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["latest_schema_version"] = None
            __props__.__dict__["next_schema_version"] = None
            __props__.__dict__["registry_name"] = None
            __props__.__dict__["schema_checkpoint"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Schema, __self__).__init__(
            'aws:glue/schema:Schema',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            compatibility: Optional[pulumi.Input[str]] = None,
            data_format: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            latest_schema_version: Optional[pulumi.Input[int]] = None,
            next_schema_version: Optional[pulumi.Input[int]] = None,
            registry_arn: Optional[pulumi.Input[str]] = None,
            registry_name: Optional[pulumi.Input[str]] = None,
            schema_checkpoint: Optional[pulumi.Input[int]] = None,
            schema_definition: Optional[pulumi.Input[str]] = None,
            schema_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Schema':
        """
        Get an existing Schema resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the schema.
        :param pulumi.Input[str] compatibility: The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
        :param pulumi.Input[str] data_format: The data format of the schema definition. Valid values are `AVRO`, `JSON` and `PROTOBUF`.
        :param pulumi.Input[str] description: A description of the schema.
        :param pulumi.Input[int] latest_schema_version: The latest version of the schema associated with the returned schema definition.
        :param pulumi.Input[int] next_schema_version: The next version of the schema associated with the returned schema definition.
        :param pulumi.Input[str] registry_arn: The ARN of the Glue Registry to create the schema in.
        :param pulumi.Input[str] registry_name: The name of the Glue Registry.
        :param pulumi.Input[int] schema_checkpoint: The version number of the checkpoint (the last time the compatibility mode was changed).
        :param pulumi.Input[str] schema_definition: The schema definition using the `data_format` setting for `schema_name`.
        :param pulumi.Input[str] schema_name: The Name of the schema.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SchemaState.__new__(_SchemaState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["compatibility"] = compatibility
        __props__.__dict__["data_format"] = data_format
        __props__.__dict__["description"] = description
        __props__.__dict__["latest_schema_version"] = latest_schema_version
        __props__.__dict__["next_schema_version"] = next_schema_version
        __props__.__dict__["registry_arn"] = registry_arn
        __props__.__dict__["registry_name"] = registry_name
        __props__.__dict__["schema_checkpoint"] = schema_checkpoint
        __props__.__dict__["schema_definition"] = schema_definition
        __props__.__dict__["schema_name"] = schema_name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return Schema(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the schema.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def compatibility(self) -> pulumi.Output[str]:
        """
        The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
        """
        return pulumi.get(self, "compatibility")

    @property
    @pulumi.getter(name="dataFormat")
    def data_format(self) -> pulumi.Output[str]:
        """
        The data format of the schema definition. Valid values are `AVRO`, `JSON` and `PROTOBUF`.
        """
        return pulumi.get(self, "data_format")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the schema.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="latestSchemaVersion")
    def latest_schema_version(self) -> pulumi.Output[int]:
        """
        The latest version of the schema associated with the returned schema definition.
        """
        return pulumi.get(self, "latest_schema_version")

    @property
    @pulumi.getter(name="nextSchemaVersion")
    def next_schema_version(self) -> pulumi.Output[int]:
        """
        The next version of the schema associated with the returned schema definition.
        """
        return pulumi.get(self, "next_schema_version")

    @property
    @pulumi.getter(name="registryArn")
    def registry_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the Glue Registry to create the schema in.
        """
        return pulumi.get(self, "registry_arn")

    @property
    @pulumi.getter(name="registryName")
    def registry_name(self) -> pulumi.Output[str]:
        """
        The name of the Glue Registry.
        """
        return pulumi.get(self, "registry_name")

    @property
    @pulumi.getter(name="schemaCheckpoint")
    def schema_checkpoint(self) -> pulumi.Output[int]:
        """
        The version number of the checkpoint (the last time the compatibility mode was changed).
        """
        return pulumi.get(self, "schema_checkpoint")

    @property
    @pulumi.getter(name="schemaDefinition")
    def schema_definition(self) -> pulumi.Output[str]:
        """
        The schema definition using the `data_format` setting for `schema_name`.
        """
        return pulumi.get(self, "schema_definition")

    @property
    @pulumi.getter(name="schemaName")
    def schema_name(self) -> pulumi.Output[str]:
        """
        The Name of the schema.
        """
        return pulumi.get(self, "schema_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

