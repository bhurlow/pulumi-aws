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

__all__ = ['PartitionIndexArgs', 'PartitionIndex']

@pulumi.input_type
class PartitionIndexArgs:
    def __init__(__self__, *,
                 database_name: pulumi.Input[str],
                 partition_index: pulumi.Input['PartitionIndexPartitionIndexArgs'],
                 table_name: pulumi.Input[str],
                 catalog_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PartitionIndex resource.
        :param pulumi.Input[str] database_name: Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
        :param pulumi.Input['PartitionIndexPartitionIndexArgs'] partition_index: Configuration block for a partition index. See `partition_index` below.
        :param pulumi.Input[str] table_name: Name of the table. For Hive compatibility, this must be entirely lowercase.
        :param pulumi.Input[str] catalog_id: The catalog ID where the table resides.
        """
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "partition_index", partition_index)
        pulumi.set(__self__, "table_name", table_name)
        if catalog_id is not None:
            pulumi.set(__self__, "catalog_id", catalog_id)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[str]:
        """
        Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="partitionIndex")
    def partition_index(self) -> pulumi.Input['PartitionIndexPartitionIndexArgs']:
        """
        Configuration block for a partition index. See `partition_index` below.
        """
        return pulumi.get(self, "partition_index")

    @partition_index.setter
    def partition_index(self, value: pulumi.Input['PartitionIndexPartitionIndexArgs']):
        pulumi.set(self, "partition_index", value)

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> pulumi.Input[str]:
        """
        Name of the table. For Hive compatibility, this must be entirely lowercase.
        """
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "table_name", value)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[pulumi.Input[str]]:
        """
        The catalog ID where the table resides.
        """
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "catalog_id", value)


@pulumi.input_type
class _PartitionIndexState:
    def __init__(__self__, *,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 partition_index: Optional[pulumi.Input['PartitionIndexPartitionIndexArgs']] = None,
                 table_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PartitionIndex resources.
        :param pulumi.Input[str] catalog_id: The catalog ID where the table resides.
        :param pulumi.Input[str] database_name: Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
        :param pulumi.Input['PartitionIndexPartitionIndexArgs'] partition_index: Configuration block for a partition index. See `partition_index` below.
        :param pulumi.Input[str] table_name: Name of the table. For Hive compatibility, this must be entirely lowercase.
        """
        if catalog_id is not None:
            pulumi.set(__self__, "catalog_id", catalog_id)
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if partition_index is not None:
            pulumi.set(__self__, "partition_index", partition_index)
        if table_name is not None:
            pulumi.set(__self__, "table_name", table_name)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[pulumi.Input[str]]:
        """
        The catalog ID where the table resides.
        """
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "catalog_id", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="partitionIndex")
    def partition_index(self) -> Optional[pulumi.Input['PartitionIndexPartitionIndexArgs']]:
        """
        Configuration block for a partition index. See `partition_index` below.
        """
        return pulumi.get(self, "partition_index")

    @partition_index.setter
    def partition_index(self, value: Optional[pulumi.Input['PartitionIndexPartitionIndexArgs']]):
        pulumi.set(self, "partition_index", value)

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the table. For Hive compatibility, this must be entirely lowercase.
        """
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "table_name", value)


class PartitionIndex(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 partition_index: Optional[pulumi.Input[pulumi.InputType['PartitionIndexPartitionIndexArgs']]] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_catalog_database = aws.glue.CatalogDatabase("exampleCatalogDatabase", name="example")
        example_catalog_table = aws.glue.CatalogTable("exampleCatalogTable",
            name="example",
            database_name=example_catalog_database.name,
            owner="my_owner",
            retention=1,
            table_type="VIRTUAL_VIEW",
            view_expanded_text="view_expanded_text_1",
            view_original_text="view_original_text_1",
            storage_descriptor=aws.glue.CatalogTableStorageDescriptorArgs(
                bucket_columns=["bucket_column_1"],
                compressed=False,
                input_format="SequenceFileInputFormat",
                location="my_location",
                number_of_buckets=1,
                output_format="SequenceFileInputFormat",
                stored_as_sub_directories=False,
                parameters={
                    "param1": "param1_val",
                },
                columns=[
                    aws.glue.CatalogTableStorageDescriptorColumnArgs(
                        name="my_column_1",
                        type="int",
                        comment="my_column1_comment",
                    ),
                    aws.glue.CatalogTableStorageDescriptorColumnArgs(
                        name="my_column_2",
                        type="string",
                        comment="my_column2_comment",
                    ),
                ],
                ser_de_info=aws.glue.CatalogTableStorageDescriptorSerDeInfoArgs(
                    name="ser_de_name",
                    parameters={
                        "param1": "param_val_1",
                    },
                    serialization_library="org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe",
                ),
                sort_columns=[aws.glue.CatalogTableStorageDescriptorSortColumnArgs(
                    column="my_column_1",
                    sort_order=1,
                )],
                skewed_info=aws.glue.CatalogTableStorageDescriptorSkewedInfoArgs(
                    skewed_column_names=["my_column_1"],
                    skewed_column_value_location_maps={
                        "my_column_1": "my_column_1_val_loc_map",
                    },
                    skewed_column_values=["skewed_val_1"],
                ),
            ),
            partition_keys=[
                aws.glue.CatalogTablePartitionKeyArgs(
                    name="my_column_1",
                    type="int",
                    comment="my_column_1_comment",
                ),
                aws.glue.CatalogTablePartitionKeyArgs(
                    name="my_column_2",
                    type="string",
                    comment="my_column_2_comment",
                ),
            ],
            parameters={
                "param1": "param1_val",
            })
        example_partition_index = aws.glue.PartitionIndex("examplePartitionIndex",
            database_name=example_catalog_database.name,
            table_name=example_catalog_table.name,
            partition_index=aws.glue.PartitionIndexPartitionIndexArgs(
                index_name="example",
                keys=[
                    "my_column_1",
                    "my_column_2",
                ],
            ))
        ```

        ## Import

        terraform import {

         to = aws_glue_partition_index.example

         id = "123456789012:MyDatabase:MyTable:index-name" } Using `pulumi import`, import Glue Partition Indexes using the catalog ID (usually AWS account ID), database name, table name, and index name. For exampleconsole % pulumi import aws_glue_partition_index.example 123456789012:MyDatabase:MyTable:index-name

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: The catalog ID where the table resides.
        :param pulumi.Input[str] database_name: Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
        :param pulumi.Input[pulumi.InputType['PartitionIndexPartitionIndexArgs']] partition_index: Configuration block for a partition index. See `partition_index` below.
        :param pulumi.Input[str] table_name: Name of the table. For Hive compatibility, this must be entirely lowercase.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PartitionIndexArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_catalog_database = aws.glue.CatalogDatabase("exampleCatalogDatabase", name="example")
        example_catalog_table = aws.glue.CatalogTable("exampleCatalogTable",
            name="example",
            database_name=example_catalog_database.name,
            owner="my_owner",
            retention=1,
            table_type="VIRTUAL_VIEW",
            view_expanded_text="view_expanded_text_1",
            view_original_text="view_original_text_1",
            storage_descriptor=aws.glue.CatalogTableStorageDescriptorArgs(
                bucket_columns=["bucket_column_1"],
                compressed=False,
                input_format="SequenceFileInputFormat",
                location="my_location",
                number_of_buckets=1,
                output_format="SequenceFileInputFormat",
                stored_as_sub_directories=False,
                parameters={
                    "param1": "param1_val",
                },
                columns=[
                    aws.glue.CatalogTableStorageDescriptorColumnArgs(
                        name="my_column_1",
                        type="int",
                        comment="my_column1_comment",
                    ),
                    aws.glue.CatalogTableStorageDescriptorColumnArgs(
                        name="my_column_2",
                        type="string",
                        comment="my_column2_comment",
                    ),
                ],
                ser_de_info=aws.glue.CatalogTableStorageDescriptorSerDeInfoArgs(
                    name="ser_de_name",
                    parameters={
                        "param1": "param_val_1",
                    },
                    serialization_library="org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe",
                ),
                sort_columns=[aws.glue.CatalogTableStorageDescriptorSortColumnArgs(
                    column="my_column_1",
                    sort_order=1,
                )],
                skewed_info=aws.glue.CatalogTableStorageDescriptorSkewedInfoArgs(
                    skewed_column_names=["my_column_1"],
                    skewed_column_value_location_maps={
                        "my_column_1": "my_column_1_val_loc_map",
                    },
                    skewed_column_values=["skewed_val_1"],
                ),
            ),
            partition_keys=[
                aws.glue.CatalogTablePartitionKeyArgs(
                    name="my_column_1",
                    type="int",
                    comment="my_column_1_comment",
                ),
                aws.glue.CatalogTablePartitionKeyArgs(
                    name="my_column_2",
                    type="string",
                    comment="my_column_2_comment",
                ),
            ],
            parameters={
                "param1": "param1_val",
            })
        example_partition_index = aws.glue.PartitionIndex("examplePartitionIndex",
            database_name=example_catalog_database.name,
            table_name=example_catalog_table.name,
            partition_index=aws.glue.PartitionIndexPartitionIndexArgs(
                index_name="example",
                keys=[
                    "my_column_1",
                    "my_column_2",
                ],
            ))
        ```

        ## Import

        terraform import {

         to = aws_glue_partition_index.example

         id = "123456789012:MyDatabase:MyTable:index-name" } Using `pulumi import`, import Glue Partition Indexes using the catalog ID (usually AWS account ID), database name, table name, and index name. For exampleconsole % pulumi import aws_glue_partition_index.example 123456789012:MyDatabase:MyTable:index-name

        :param str resource_name: The name of the resource.
        :param PartitionIndexArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PartitionIndexArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 partition_index: Optional[pulumi.Input[pulumi.InputType['PartitionIndexPartitionIndexArgs']]] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PartitionIndexArgs.__new__(PartitionIndexArgs)

            __props__.__dict__["catalog_id"] = catalog_id
            if database_name is None and not opts.urn:
                raise TypeError("Missing required property 'database_name'")
            __props__.__dict__["database_name"] = database_name
            if partition_index is None and not opts.urn:
                raise TypeError("Missing required property 'partition_index'")
            __props__.__dict__["partition_index"] = partition_index
            if table_name is None and not opts.urn:
                raise TypeError("Missing required property 'table_name'")
            __props__.__dict__["table_name"] = table_name
        super(PartitionIndex, __self__).__init__(
            'aws:glue/partitionIndex:PartitionIndex',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            catalog_id: Optional[pulumi.Input[str]] = None,
            database_name: Optional[pulumi.Input[str]] = None,
            partition_index: Optional[pulumi.Input[pulumi.InputType['PartitionIndexPartitionIndexArgs']]] = None,
            table_name: Optional[pulumi.Input[str]] = None) -> 'PartitionIndex':
        """
        Get an existing PartitionIndex resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: The catalog ID where the table resides.
        :param pulumi.Input[str] database_name: Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
        :param pulumi.Input[pulumi.InputType['PartitionIndexPartitionIndexArgs']] partition_index: Configuration block for a partition index. See `partition_index` below.
        :param pulumi.Input[str] table_name: Name of the table. For Hive compatibility, this must be entirely lowercase.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PartitionIndexState.__new__(_PartitionIndexState)

        __props__.__dict__["catalog_id"] = catalog_id
        __props__.__dict__["database_name"] = database_name
        __props__.__dict__["partition_index"] = partition_index
        __props__.__dict__["table_name"] = table_name
        return PartitionIndex(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> pulumi.Output[str]:
        """
        The catalog ID where the table resides.
        """
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[str]:
        """
        Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="partitionIndex")
    def partition_index(self) -> pulumi.Output['outputs.PartitionIndexPartitionIndex']:
        """
        Configuration block for a partition index. See `partition_index` below.
        """
        return pulumi.get(self, "partition_index")

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> pulumi.Output[str]:
        """
        Name of the table. For Hive compatibility, this must be entirely lowercase.
        """
        return pulumi.get(self, "table_name")

