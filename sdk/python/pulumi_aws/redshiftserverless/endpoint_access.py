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

__all__ = ['EndpointAccessArgs', 'EndpointAccess']

@pulumi.input_type
class EndpointAccessArgs:
    def __init__(__self__, *,
                 endpoint_name: pulumi.Input[str],
                 subnet_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 workgroup_name: pulumi.Input[str],
                 vpc_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a EndpointAccess resource.
        :param pulumi.Input[str] endpoint_name: The name of the endpoint.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: An array of VPC subnet IDs to associate with the endpoint.
        :param pulumi.Input[str] workgroup_name: The name of the workgroup.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_security_group_ids: An array of security group IDs to associate with the workgroup.
        """
        pulumi.set(__self__, "endpoint_name", endpoint_name)
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        pulumi.set(__self__, "workgroup_name", workgroup_name)
        if vpc_security_group_ids is not None:
            pulumi.set(__self__, "vpc_security_group_ids", vpc_security_group_ids)

    @property
    @pulumi.getter(name="endpointName")
    def endpoint_name(self) -> pulumi.Input[str]:
        """
        The name of the endpoint.
        """
        return pulumi.get(self, "endpoint_name")

    @endpoint_name.setter
    def endpoint_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint_name", value)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        An array of VPC subnet IDs to associate with the endpoint.
        """
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "subnet_ids", value)

    @property
    @pulumi.getter(name="workgroupName")
    def workgroup_name(self) -> pulumi.Input[str]:
        """
        The name of the workgroup.
        """
        return pulumi.get(self, "workgroup_name")

    @workgroup_name.setter
    def workgroup_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "workgroup_name", value)

    @property
    @pulumi.getter(name="vpcSecurityGroupIds")
    def vpc_security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of security group IDs to associate with the workgroup.
        """
        return pulumi.get(self, "vpc_security_group_ids")

    @vpc_security_group_ids.setter
    def vpc_security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vpc_security_group_ids", value)


@pulumi.input_type
class _EndpointAccessState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 endpoint_name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpc_endpoints: Optional[pulumi.Input[Sequence[pulumi.Input['EndpointAccessVpcEndpointArgs']]]] = None,
                 vpc_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 workgroup_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EndpointAccess resources.
        :param pulumi.Input[str] address: The DNS address of the VPC endpoint.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the Redshift Serverless Endpoint Access.
        :param pulumi.Input[str] endpoint_name: The name of the endpoint.
        :param pulumi.Input[int] port: The port that Amazon Redshift Serverless listens on.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: An array of VPC subnet IDs to associate with the endpoint.
        :param pulumi.Input[Sequence[pulumi.Input['EndpointAccessVpcEndpointArgs']]] vpc_endpoints: The VPC endpoint or the Redshift Serverless workgroup. See `VPC Endpoint` below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_security_group_ids: An array of security group IDs to associate with the workgroup.
        :param pulumi.Input[str] workgroup_name: The name of the workgroup.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if endpoint_name is not None:
            pulumi.set(__self__, "endpoint_name", endpoint_name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if subnet_ids is not None:
            pulumi.set(__self__, "subnet_ids", subnet_ids)
        if vpc_endpoints is not None:
            pulumi.set(__self__, "vpc_endpoints", vpc_endpoints)
        if vpc_security_group_ids is not None:
            pulumi.set(__self__, "vpc_security_group_ids", vpc_security_group_ids)
        if workgroup_name is not None:
            pulumi.set(__self__, "workgroup_name", workgroup_name)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        The DNS address of the VPC endpoint.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the Redshift Serverless Endpoint Access.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="endpointName")
    def endpoint_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the endpoint.
        """
        return pulumi.get(self, "endpoint_name")

    @endpoint_name.setter
    def endpoint_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The port that Amazon Redshift Serverless listens on.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of VPC subnet IDs to associate with the endpoint.
        """
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "subnet_ids", value)

    @property
    @pulumi.getter(name="vpcEndpoints")
    def vpc_endpoints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EndpointAccessVpcEndpointArgs']]]]:
        """
        The VPC endpoint or the Redshift Serverless workgroup. See `VPC Endpoint` below.
        """
        return pulumi.get(self, "vpc_endpoints")

    @vpc_endpoints.setter
    def vpc_endpoints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EndpointAccessVpcEndpointArgs']]]]):
        pulumi.set(self, "vpc_endpoints", value)

    @property
    @pulumi.getter(name="vpcSecurityGroupIds")
    def vpc_security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of security group IDs to associate with the workgroup.
        """
        return pulumi.get(self, "vpc_security_group_ids")

    @vpc_security_group_ids.setter
    def vpc_security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vpc_security_group_ids", value)

    @property
    @pulumi.getter(name="workgroupName")
    def workgroup_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the workgroup.
        """
        return pulumi.get(self, "workgroup_name")

    @workgroup_name.setter
    def workgroup_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workgroup_name", value)


class EndpointAccess(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint_name: Optional[pulumi.Input[str]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpc_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 workgroup_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new Amazon Redshift Serverless Endpoint Access.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.redshiftserverless.EndpointAccess("example",
            endpoint_name="example",
            workgroup_name="example")
        ```

        ## Import

        terraform import {

         to = aws_redshiftserverless_endpoint_access.example

         id = "example" } Using `pulumi import`, import Redshift Serverless Endpoint Access using the `endpoint_name`. For exampleconsole % pulumi import aws_redshiftserverless_endpoint_access.example example

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint_name: The name of the endpoint.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: An array of VPC subnet IDs to associate with the endpoint.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_security_group_ids: An array of security group IDs to associate with the workgroup.
        :param pulumi.Input[str] workgroup_name: The name of the workgroup.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EndpointAccessArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new Amazon Redshift Serverless Endpoint Access.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.redshiftserverless.EndpointAccess("example",
            endpoint_name="example",
            workgroup_name="example")
        ```

        ## Import

        terraform import {

         to = aws_redshiftserverless_endpoint_access.example

         id = "example" } Using `pulumi import`, import Redshift Serverless Endpoint Access using the `endpoint_name`. For exampleconsole % pulumi import aws_redshiftserverless_endpoint_access.example example

        :param str resource_name: The name of the resource.
        :param EndpointAccessArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndpointAccessArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint_name: Optional[pulumi.Input[str]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpc_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 workgroup_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EndpointAccessArgs.__new__(EndpointAccessArgs)

            if endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint_name'")
            __props__.__dict__["endpoint_name"] = endpoint_name
            if subnet_ids is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_ids'")
            __props__.__dict__["subnet_ids"] = subnet_ids
            __props__.__dict__["vpc_security_group_ids"] = vpc_security_group_ids
            if workgroup_name is None and not opts.urn:
                raise TypeError("Missing required property 'workgroup_name'")
            __props__.__dict__["workgroup_name"] = workgroup_name
            __props__.__dict__["address"] = None
            __props__.__dict__["arn"] = None
            __props__.__dict__["port"] = None
            __props__.__dict__["vpc_endpoints"] = None
        super(EndpointAccess, __self__).__init__(
            'aws:redshiftserverless/endpointAccess:EndpointAccess',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            endpoint_name: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            vpc_endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointAccessVpcEndpointArgs']]]]] = None,
            vpc_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            workgroup_name: Optional[pulumi.Input[str]] = None) -> 'EndpointAccess':
        """
        Get an existing EndpointAccess resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The DNS address of the VPC endpoint.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the Redshift Serverless Endpoint Access.
        :param pulumi.Input[str] endpoint_name: The name of the endpoint.
        :param pulumi.Input[int] port: The port that Amazon Redshift Serverless listens on.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: An array of VPC subnet IDs to associate with the endpoint.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointAccessVpcEndpointArgs']]]] vpc_endpoints: The VPC endpoint or the Redshift Serverless workgroup. See `VPC Endpoint` below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_security_group_ids: An array of security group IDs to associate with the workgroup.
        :param pulumi.Input[str] workgroup_name: The name of the workgroup.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EndpointAccessState.__new__(_EndpointAccessState)

        __props__.__dict__["address"] = address
        __props__.__dict__["arn"] = arn
        __props__.__dict__["endpoint_name"] = endpoint_name
        __props__.__dict__["port"] = port
        __props__.__dict__["subnet_ids"] = subnet_ids
        __props__.__dict__["vpc_endpoints"] = vpc_endpoints
        __props__.__dict__["vpc_security_group_ids"] = vpc_security_group_ids
        __props__.__dict__["workgroup_name"] = workgroup_name
        return EndpointAccess(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        The DNS address of the VPC endpoint.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the Redshift Serverless Endpoint Access.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="endpointName")
    def endpoint_name(self) -> pulumi.Output[str]:
        """
        The name of the endpoint.
        """
        return pulumi.get(self, "endpoint_name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        The port that Amazon Redshift Serverless listens on.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        An array of VPC subnet IDs to associate with the endpoint.
        """
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter(name="vpcEndpoints")
    def vpc_endpoints(self) -> pulumi.Output[Sequence['outputs.EndpointAccessVpcEndpoint']]:
        """
        The VPC endpoint or the Redshift Serverless workgroup. See `VPC Endpoint` below.
        """
        return pulumi.get(self, "vpc_endpoints")

    @property
    @pulumi.getter(name="vpcSecurityGroupIds")
    def vpc_security_group_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        An array of security group IDs to associate with the workgroup.
        """
        return pulumi.get(self, "vpc_security_group_ids")

    @property
    @pulumi.getter(name="workgroupName")
    def workgroup_name(self) -> pulumi.Output[str]:
        """
        The name of the workgroup.
        """
        return pulumi.get(self, "workgroup_name")

