# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LfTagArgs', 'LfTag']

@pulumi.input_type
class LfTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 values: pulumi.Input[Sequence[pulumi.Input[str]]],
                 catalog_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LfTag resource.
        :param pulumi.Input[str] key: Key-name for the tag.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] values: List of possible values an attribute can take.
        :param pulumi.Input[str] catalog_id: ID of the Data Catalog to create the tag in. If omitted, this defaults to the AWS Account ID.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "values", values)
        if catalog_id is not None:
            pulumi.set(__self__, "catalog_id", catalog_id)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        Key-name for the tag.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def values(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of possible values an attribute can take.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "values", value)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the Data Catalog to create the tag in. If omitted, this defaults to the AWS Account ID.
        """
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "catalog_id", value)


@pulumi.input_type
class _LfTagState:
    def __init__(__self__, *,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering LfTag resources.
        :param pulumi.Input[str] catalog_id: ID of the Data Catalog to create the tag in. If omitted, this defaults to the AWS Account ID.
        :param pulumi.Input[str] key: Key-name for the tag.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] values: List of possible values an attribute can take.
        """
        if catalog_id is not None:
            pulumi.set(__self__, "catalog_id", catalog_id)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the Data Catalog to create the tag in. If omitted, this defaults to the AWS Account ID.
        """
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "catalog_id", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        Key-name for the tag.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of possible values an attribute can take.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "values", value)


class LfTag(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Creates an LF-Tag with the specified name and values. Each key must have at least one value. The maximum number of values permitted is 15.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.lakeformation.LfTag("example",
            key="module",
            values=[
                "Orders",
                "Sales",
                "Customers",
            ])
        ```

        ## Import

        terraform import {

         to = aws_lakeformation_lf_tag.example

         id = "123456789012:some_key" } Using `pulumi import`, import Lake Formation LF-Tags using the `catalog_id:key`. If you have not set a Catalog ID specify the AWS Account ID that the database is in. For exampleconsole % pulumi import aws_lakeformation_lf_tag.example 123456789012:some_key

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: ID of the Data Catalog to create the tag in. If omitted, this defaults to the AWS Account ID.
        :param pulumi.Input[str] key: Key-name for the tag.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] values: List of possible values an attribute can take.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LfTagArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an LF-Tag with the specified name and values. Each key must have at least one value. The maximum number of values permitted is 15.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.lakeformation.LfTag("example",
            key="module",
            values=[
                "Orders",
                "Sales",
                "Customers",
            ])
        ```

        ## Import

        terraform import {

         to = aws_lakeformation_lf_tag.example

         id = "123456789012:some_key" } Using `pulumi import`, import Lake Formation LF-Tags using the `catalog_id:key`. If you have not set a Catalog ID specify the AWS Account ID that the database is in. For exampleconsole % pulumi import aws_lakeformation_lf_tag.example 123456789012:some_key

        :param str resource_name: The name of the resource.
        :param LfTagArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LfTagArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LfTagArgs.__new__(LfTagArgs)

            __props__.__dict__["catalog_id"] = catalog_id
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            if values is None and not opts.urn:
                raise TypeError("Missing required property 'values'")
            __props__.__dict__["values"] = values
        super(LfTag, __self__).__init__(
            'aws:lakeformation/lfTag:LfTag',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            catalog_id: Optional[pulumi.Input[str]] = None,
            key: Optional[pulumi.Input[str]] = None,
            values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'LfTag':
        """
        Get an existing LfTag resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: ID of the Data Catalog to create the tag in. If omitted, this defaults to the AWS Account ID.
        :param pulumi.Input[str] key: Key-name for the tag.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] values: List of possible values an attribute can take.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LfTagState.__new__(_LfTagState)

        __props__.__dict__["catalog_id"] = catalog_id
        __props__.__dict__["key"] = key
        __props__.__dict__["values"] = values
        return LfTag(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> pulumi.Output[str]:
        """
        ID of the Data Catalog to create the tag in. If omitted, this defaults to the AWS Account ID.
        """
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        Key-name for the tag.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def values(self) -> pulumi.Output[Sequence[str]]:
        """
        List of possible values an attribute can take.
        """
        return pulumi.get(self, "values")

