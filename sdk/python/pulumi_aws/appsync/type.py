# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TypeArgs', 'Type']

@pulumi.input_type
class TypeArgs:
    def __init__(__self__, *,
                 api_id: pulumi.Input[str],
                 definition: pulumi.Input[str],
                 format: pulumi.Input[str]):
        """
        The set of arguments for constructing a Type resource.
        :param pulumi.Input[str] api_id: GraphQL API ID.
        :param pulumi.Input[str] definition: The type definition.
        :param pulumi.Input[str] format: The type format: `SDL` or `JSON`.
        """
        pulumi.set(__self__, "api_id", api_id)
        pulumi.set(__self__, "definition", definition)
        pulumi.set(__self__, "format", format)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Input[str]:
        """
        GraphQL API ID.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter
    def definition(self) -> pulumi.Input[str]:
        """
        The type definition.
        """
        return pulumi.get(self, "definition")

    @definition.setter
    def definition(self, value: pulumi.Input[str]):
        pulumi.set(self, "definition", value)

    @property
    @pulumi.getter
    def format(self) -> pulumi.Input[str]:
        """
        The type format: `SDL` or `JSON`.
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: pulumi.Input[str]):
        pulumi.set(self, "format", value)


@pulumi.input_type
class _TypeState:
    def __init__(__self__, *,
                 api_id: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 definition: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Type resources.
        :param pulumi.Input[str] api_id: GraphQL API ID.
        :param pulumi.Input[str] arn: The ARN of the type.
        :param pulumi.Input[str] definition: The type definition.
        :param pulumi.Input[str] description: The type description.
        :param pulumi.Input[str] format: The type format: `SDL` or `JSON`.
        :param pulumi.Input[str] name: The type name.
        """
        if api_id is not None:
            pulumi.set(__self__, "api_id", api_id)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if definition is not None:
            pulumi.set(__self__, "definition", definition)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if format is not None:
            pulumi.set(__self__, "format", format)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> Optional[pulumi.Input[str]]:
        """
        GraphQL API ID.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the type.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def definition(self) -> Optional[pulumi.Input[str]]:
        """
        The type definition.
        """
        return pulumi.get(self, "definition")

    @definition.setter
    def definition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "definition", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The type description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        """
        The type format: `SDL` or `JSON`.
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The type name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class Type(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 definition: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an AppSync Type.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_graph_ql_api = aws.appsync.GraphQLApi("exampleGraphQLApi", authentication_type="API_KEY")
        example_type = aws.appsync.Type("exampleType",
            api_id=example_graph_ql_api.id,
            format="SDL",
            definition=\"\"\"type Mutation

        {
        putPost(id: ID!,title: String! ): Post

        }
        \"\"\")
        ```

        ## Import

        terraform import {

         to = aws_appsync_type.example

         id = "api-id:format:name" } Using `pulumi import`, import Appsync Types using the `id`. For exampleconsole % pulumi import aws_appsync_type.example api-id:format:name

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: GraphQL API ID.
        :param pulumi.Input[str] definition: The type definition.
        :param pulumi.Input[str] format: The type format: `SDL` or `JSON`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TypeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an AppSync Type.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_graph_ql_api = aws.appsync.GraphQLApi("exampleGraphQLApi", authentication_type="API_KEY")
        example_type = aws.appsync.Type("exampleType",
            api_id=example_graph_ql_api.id,
            format="SDL",
            definition=\"\"\"type Mutation

        {
        putPost(id: ID!,title: String! ): Post

        }
        \"\"\")
        ```

        ## Import

        terraform import {

         to = aws_appsync_type.example

         id = "api-id:format:name" } Using `pulumi import`, import Appsync Types using the `id`. For exampleconsole % pulumi import aws_appsync_type.example api-id:format:name

        :param str resource_name: The name of the resource.
        :param TypeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TypeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 definition: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TypeArgs.__new__(TypeArgs)

            if api_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_id'")
            __props__.__dict__["api_id"] = api_id
            if definition is None and not opts.urn:
                raise TypeError("Missing required property 'definition'")
            __props__.__dict__["definition"] = definition
            if format is None and not opts.urn:
                raise TypeError("Missing required property 'format'")
            __props__.__dict__["format"] = format
            __props__.__dict__["arn"] = None
            __props__.__dict__["description"] = None
            __props__.__dict__["name"] = None
        super(Type, __self__).__init__(
            'aws:appsync/type:Type',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_id: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            definition: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            format: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'Type':
        """
        Get an existing Type resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: GraphQL API ID.
        :param pulumi.Input[str] arn: The ARN of the type.
        :param pulumi.Input[str] definition: The type definition.
        :param pulumi.Input[str] description: The type description.
        :param pulumi.Input[str] format: The type format: `SDL` or `JSON`.
        :param pulumi.Input[str] name: The type name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TypeState.__new__(_TypeState)

        __props__.__dict__["api_id"] = api_id
        __props__.__dict__["arn"] = arn
        __props__.__dict__["definition"] = definition
        __props__.__dict__["description"] = description
        __props__.__dict__["format"] = format
        __props__.__dict__["name"] = name
        return Type(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[str]:
        """
        GraphQL API ID.
        """
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the type.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def definition(self) -> pulumi.Output[str]:
        """
        The type definition.
        """
        return pulumi.get(self, "definition")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The type description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def format(self) -> pulumi.Output[str]:
        """
        The type format: `SDL` or `JSON`.
        """
        return pulumi.get(self, "format")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The type name.
        """
        return pulumi.get(self, "name")

