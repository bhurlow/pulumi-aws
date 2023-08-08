# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ProxyEndpointArgs', 'ProxyEndpoint']

@pulumi.input_type
class ProxyEndpointArgs:
    def __init__(__self__, *,
                 db_proxy_endpoint_name: pulumi.Input[str],
                 db_proxy_name: pulumi.Input[str],
                 vpc_subnet_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_role: Optional[pulumi.Input[str]] = None,
                 vpc_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ProxyEndpoint resource.
        :param pulumi.Input[str] db_proxy_endpoint_name: The identifier for the proxy endpoint. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
        :param pulumi.Input[str] db_proxy_name: The name of the DB proxy associated with the DB proxy endpoint that you create.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_subnet_ids: One or more VPC subnet IDs to associate with the new proxy.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] target_role: Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is `READ_WRITE`. Valid values are `READ_WRITE` and `READ_ONLY`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_security_group_ids: One or more VPC security group IDs to associate with the new proxy.
        """
        pulumi.set(__self__, "db_proxy_endpoint_name", db_proxy_endpoint_name)
        pulumi.set(__self__, "db_proxy_name", db_proxy_name)
        pulumi.set(__self__, "vpc_subnet_ids", vpc_subnet_ids)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if target_role is not None:
            pulumi.set(__self__, "target_role", target_role)
        if vpc_security_group_ids is not None:
            pulumi.set(__self__, "vpc_security_group_ids", vpc_security_group_ids)

    @property
    @pulumi.getter(name="dbProxyEndpointName")
    def db_proxy_endpoint_name(self) -> pulumi.Input[str]:
        """
        The identifier for the proxy endpoint. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
        """
        return pulumi.get(self, "db_proxy_endpoint_name")

    @db_proxy_endpoint_name.setter
    def db_proxy_endpoint_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_proxy_endpoint_name", value)

    @property
    @pulumi.getter(name="dbProxyName")
    def db_proxy_name(self) -> pulumi.Input[str]:
        """
        The name of the DB proxy associated with the DB proxy endpoint that you create.
        """
        return pulumi.get(self, "db_proxy_name")

    @db_proxy_name.setter
    def db_proxy_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_proxy_name", value)

    @property
    @pulumi.getter(name="vpcSubnetIds")
    def vpc_subnet_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        One or more VPC subnet IDs to associate with the new proxy.
        """
        return pulumi.get(self, "vpc_subnet_ids")

    @vpc_subnet_ids.setter
    def vpc_subnet_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "vpc_subnet_ids", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="targetRole")
    def target_role(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is `READ_WRITE`. Valid values are `READ_WRITE` and `READ_ONLY`.
        """
        return pulumi.get(self, "target_role")

    @target_role.setter
    def target_role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_role", value)

    @property
    @pulumi.getter(name="vpcSecurityGroupIds")
    def vpc_security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        One or more VPC security group IDs to associate with the new proxy.
        """
        return pulumi.get(self, "vpc_security_group_ids")

    @vpc_security_group_ids.setter
    def vpc_security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vpc_security_group_ids", value)


@pulumi.input_type
class _ProxyEndpointState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 db_proxy_endpoint_name: Optional[pulumi.Input[str]] = None,
                 db_proxy_name: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_role: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vpc_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpc_subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ProxyEndpoint resources.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) for the proxy endpoint.
        :param pulumi.Input[str] db_proxy_endpoint_name: The identifier for the proxy endpoint. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
        :param pulumi.Input[str] db_proxy_name: The name of the DB proxy associated with the DB proxy endpoint that you create.
        :param pulumi.Input[str] endpoint: The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
        :param pulumi.Input[bool] is_default: Indicates whether this endpoint is the default endpoint for the associated DB proxy.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] target_role: Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is `READ_WRITE`. Valid values are `READ_WRITE` and `READ_ONLY`.
        :param pulumi.Input[str] vpc_id: The VPC ID of the DB proxy endpoint.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_security_group_ids: One or more VPC security group IDs to associate with the new proxy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_subnet_ids: One or more VPC subnet IDs to associate with the new proxy.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if db_proxy_endpoint_name is not None:
            pulumi.set(__self__, "db_proxy_endpoint_name", db_proxy_endpoint_name)
        if db_proxy_name is not None:
            pulumi.set(__self__, "db_proxy_name", db_proxy_name)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if is_default is not None:
            pulumi.set(__self__, "is_default", is_default)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if target_role is not None:
            pulumi.set(__self__, "target_role", target_role)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if vpc_security_group_ids is not None:
            pulumi.set(__self__, "vpc_security_group_ids", vpc_security_group_ids)
        if vpc_subnet_ids is not None:
            pulumi.set(__self__, "vpc_subnet_ids", vpc_subnet_ids)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) for the proxy endpoint.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="dbProxyEndpointName")
    def db_proxy_endpoint_name(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier for the proxy endpoint. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
        """
        return pulumi.get(self, "db_proxy_endpoint_name")

    @db_proxy_endpoint_name.setter
    def db_proxy_endpoint_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_proxy_endpoint_name", value)

    @property
    @pulumi.getter(name="dbProxyName")
    def db_proxy_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the DB proxy associated with the DB proxy endpoint that you create.
        """
        return pulumi.get(self, "db_proxy_name")

    @db_proxy_name.setter
    def db_proxy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_proxy_name", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether this endpoint is the default endpoint for the associated DB proxy.
        """
        return pulumi.get(self, "is_default")

    @is_default.setter
    def is_default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_default", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter(name="targetRole")
    def target_role(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is `READ_WRITE`. Valid values are `READ_WRITE` and `READ_ONLY`.
        """
        return pulumi.get(self, "target_role")

    @target_role.setter
    def target_role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_role", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The VPC ID of the DB proxy endpoint.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vpcSecurityGroupIds")
    def vpc_security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        One or more VPC security group IDs to associate with the new proxy.
        """
        return pulumi.get(self, "vpc_security_group_ids")

    @vpc_security_group_ids.setter
    def vpc_security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vpc_security_group_ids", value)

    @property
    @pulumi.getter(name="vpcSubnetIds")
    def vpc_subnet_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        One or more VPC subnet IDs to associate with the new proxy.
        """
        return pulumi.get(self, "vpc_subnet_ids")

    @vpc_subnet_ids.setter
    def vpc_subnet_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vpc_subnet_ids", value)


class ProxyEndpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_proxy_endpoint_name: Optional[pulumi.Input[str]] = None,
                 db_proxy_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_role: Optional[pulumi.Input[str]] = None,
                 vpc_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpc_subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides an RDS DB proxy endpoint resource. For additional information, see the [RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy-endpoints.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.rds.ProxyEndpoint("example",
            db_proxy_name=aws_db_proxy["test"]["name"],
            db_proxy_endpoint_name="example",
            vpc_subnet_ids=[__item["id"] for __item in aws_subnet["test"]],
            target_role="READ_ONLY")
        ```

        ## Import

        terraform import {

         to = aws_db_proxy_endpoint.example

         id = "example/example" } Using `pulumi import`, import DB proxy endpoints using the `DB-PROXY-NAME/DB-PROXY-ENDPOINT-NAME`. For exampleconsole % pulumi import aws_db_proxy_endpoint.example example/example

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_proxy_endpoint_name: The identifier for the proxy endpoint. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
        :param pulumi.Input[str] db_proxy_name: The name of the DB proxy associated with the DB proxy endpoint that you create.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] target_role: Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is `READ_WRITE`. Valid values are `READ_WRITE` and `READ_ONLY`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_security_group_ids: One or more VPC security group IDs to associate with the new proxy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_subnet_ids: One or more VPC subnet IDs to associate with the new proxy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProxyEndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an RDS DB proxy endpoint resource. For additional information, see the [RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy-endpoints.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.rds.ProxyEndpoint("example",
            db_proxy_name=aws_db_proxy["test"]["name"],
            db_proxy_endpoint_name="example",
            vpc_subnet_ids=[__item["id"] for __item in aws_subnet["test"]],
            target_role="READ_ONLY")
        ```

        ## Import

        terraform import {

         to = aws_db_proxy_endpoint.example

         id = "example/example" } Using `pulumi import`, import DB proxy endpoints using the `DB-PROXY-NAME/DB-PROXY-ENDPOINT-NAME`. For exampleconsole % pulumi import aws_db_proxy_endpoint.example example/example

        :param str resource_name: The name of the resource.
        :param ProxyEndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProxyEndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_proxy_endpoint_name: Optional[pulumi.Input[str]] = None,
                 db_proxy_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_role: Optional[pulumi.Input[str]] = None,
                 vpc_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpc_subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProxyEndpointArgs.__new__(ProxyEndpointArgs)

            if db_proxy_endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'db_proxy_endpoint_name'")
            __props__.__dict__["db_proxy_endpoint_name"] = db_proxy_endpoint_name
            if db_proxy_name is None and not opts.urn:
                raise TypeError("Missing required property 'db_proxy_name'")
            __props__.__dict__["db_proxy_name"] = db_proxy_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["target_role"] = target_role
            __props__.__dict__["vpc_security_group_ids"] = vpc_security_group_ids
            if vpc_subnet_ids is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_subnet_ids'")
            __props__.__dict__["vpc_subnet_ids"] = vpc_subnet_ids
            __props__.__dict__["arn"] = None
            __props__.__dict__["endpoint"] = None
            __props__.__dict__["is_default"] = None
            __props__.__dict__["tags_all"] = None
            __props__.__dict__["vpc_id"] = None
        super(ProxyEndpoint, __self__).__init__(
            'aws:rds/proxyEndpoint:ProxyEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            db_proxy_endpoint_name: Optional[pulumi.Input[str]] = None,
            db_proxy_name: Optional[pulumi.Input[str]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            is_default: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            target_role: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            vpc_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            vpc_subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'ProxyEndpoint':
        """
        Get an existing ProxyEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) for the proxy endpoint.
        :param pulumi.Input[str] db_proxy_endpoint_name: The identifier for the proxy endpoint. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
        :param pulumi.Input[str] db_proxy_name: The name of the DB proxy associated with the DB proxy endpoint that you create.
        :param pulumi.Input[str] endpoint: The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
        :param pulumi.Input[bool] is_default: Indicates whether this endpoint is the default endpoint for the associated DB proxy.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] target_role: Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is `READ_WRITE`. Valid values are `READ_WRITE` and `READ_ONLY`.
        :param pulumi.Input[str] vpc_id: The VPC ID of the DB proxy endpoint.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_security_group_ids: One or more VPC security group IDs to associate with the new proxy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_subnet_ids: One or more VPC subnet IDs to associate with the new proxy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProxyEndpointState.__new__(_ProxyEndpointState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["db_proxy_endpoint_name"] = db_proxy_endpoint_name
        __props__.__dict__["db_proxy_name"] = db_proxy_name
        __props__.__dict__["endpoint"] = endpoint
        __props__.__dict__["is_default"] = is_default
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["target_role"] = target_role
        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["vpc_security_group_ids"] = vpc_security_group_ids
        __props__.__dict__["vpc_subnet_ids"] = vpc_subnet_ids
        return ProxyEndpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) for the proxy endpoint.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="dbProxyEndpointName")
    def db_proxy_endpoint_name(self) -> pulumi.Output[str]:
        """
        The identifier for the proxy endpoint. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
        """
        return pulumi.get(self, "db_proxy_endpoint_name")

    @property
    @pulumi.getter(name="dbProxyName")
    def db_proxy_name(self) -> pulumi.Output[str]:
        """
        The name of the DB proxy associated with the DB proxy endpoint that you create.
        """
        return pulumi.get(self, "db_proxy_name")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> pulumi.Output[bool]:
        """
        Indicates whether this endpoint is the default endpoint for the associated DB proxy.
        """
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter(name="targetRole")
    def target_role(self) -> pulumi.Output[Optional[str]]:
        """
        Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is `READ_WRITE`. Valid values are `READ_WRITE` and `READ_ONLY`.
        """
        return pulumi.get(self, "target_role")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The VPC ID of the DB proxy endpoint.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vpcSecurityGroupIds")
    def vpc_security_group_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        One or more VPC security group IDs to associate with the new proxy.
        """
        return pulumi.get(self, "vpc_security_group_ids")

    @property
    @pulumi.getter(name="vpcSubnetIds")
    def vpc_subnet_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        One or more VPC subnet IDs to associate with the new proxy.
        """
        return pulumi.get(self, "vpc_subnet_ids")

