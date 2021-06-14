# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NatGatewayArgs', 'NatGateway']

@pulumi.input_type
class NatGatewayArgs:
    def __init__(__self__, *,
                 subnet_id: pulumi.Input[str],
                 allocation_id: Optional[pulumi.Input[str]] = None,
                 connectivity_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a NatGateway resource.
        :param pulumi.Input[str] subnet_id: The Subnet ID of the subnet in which to place the gateway.
        :param pulumi.Input[str] allocation_id: The Allocation ID of the Elastic IP address for the gateway. Required for `connectivity_type` of `public`.
        :param pulumi.Input[str] connectivity_type: Connectivity type for the gateway. Valid values are `private` and `public`. Defaults to `public`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        pulumi.set(__self__, "subnet_id", subnet_id)
        if allocation_id is not None:
            pulumi.set(__self__, "allocation_id", allocation_id)
        if connectivity_type is not None:
            pulumi.set(__self__, "connectivity_type", connectivity_type)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        The Subnet ID of the subnet in which to place the gateway.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="allocationId")
    def allocation_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Allocation ID of the Elastic IP address for the gateway. Required for `connectivity_type` of `public`.
        """
        return pulumi.get(self, "allocation_id")

    @allocation_id.setter
    def allocation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allocation_id", value)

    @property
    @pulumi.getter(name="connectivityType")
    def connectivity_type(self) -> Optional[pulumi.Input[str]]:
        """
        Connectivity type for the gateway. Valid values are `private` and `public`. Defaults to `public`.
        """
        return pulumi.get(self, "connectivity_type")

    @connectivity_type.setter
    def connectivity_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connectivity_type", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


@pulumi.input_type
class _NatGatewayState:
    def __init__(__self__, *,
                 allocation_id: Optional[pulumi.Input[str]] = None,
                 connectivity_type: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 private_ip: Optional[pulumi.Input[str]] = None,
                 public_ip: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering NatGateway resources.
        :param pulumi.Input[str] allocation_id: The Allocation ID of the Elastic IP address for the gateway. Required for `connectivity_type` of `public`.
        :param pulumi.Input[str] connectivity_type: Connectivity type for the gateway. Valid values are `private` and `public`. Defaults to `public`.
        :param pulumi.Input[str] network_interface_id: The ENI ID of the network interface created by the NAT gateway.
        :param pulumi.Input[str] private_ip: The private IP address of the NAT Gateway.
        :param pulumi.Input[str] public_ip: The public IP address of the NAT Gateway.
        :param pulumi.Input[str] subnet_id: The Subnet ID of the subnet in which to place the gateway.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        if allocation_id is not None:
            pulumi.set(__self__, "allocation_id", allocation_id)
        if connectivity_type is not None:
            pulumi.set(__self__, "connectivity_type", connectivity_type)
        if network_interface_id is not None:
            pulumi.set(__self__, "network_interface_id", network_interface_id)
        if private_ip is not None:
            pulumi.set(__self__, "private_ip", private_ip)
        if public_ip is not None:
            pulumi.set(__self__, "public_ip", public_ip)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter(name="allocationId")
    def allocation_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Allocation ID of the Elastic IP address for the gateway. Required for `connectivity_type` of `public`.
        """
        return pulumi.get(self, "allocation_id")

    @allocation_id.setter
    def allocation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allocation_id", value)

    @property
    @pulumi.getter(name="connectivityType")
    def connectivity_type(self) -> Optional[pulumi.Input[str]]:
        """
        Connectivity type for the gateway. Valid values are `private` and `public`. Defaults to `public`.
        """
        return pulumi.get(self, "connectivity_type")

    @connectivity_type.setter
    def connectivity_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connectivity_type", value)

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ENI ID of the network interface created by the NAT gateway.
        """
        return pulumi.get(self, "network_interface_id")

    @network_interface_id.setter
    def network_interface_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_interface_id", value)

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The private IP address of the NAT Gateway.
        """
        return pulumi.get(self, "private_ip")

    @private_ip.setter
    def private_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_ip", value)

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The public IP address of the NAT Gateway.
        """
        return pulumi.get(self, "public_ip")

    @public_ip.setter
    def public_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_ip", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Subnet ID of the subnet in which to place the gateway.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class NatGateway(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allocation_id: Optional[pulumi.Input[str]] = None,
                 connectivity_type: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a VPC NAT Gateway.

        ## Example Usage
        ### Public NAT

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2.NatGateway("example",
            allocation_id=aws_eip["example"]["id"],
            subnet_id=aws_subnet["example"]["id"],
            tags={
                "Name": "gw NAT",
            },
            opts=pulumi.ResourceOptions(depends_on=[aws_internet_gateway["example"]]))
        ```
        ### Private NAT

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2.NatGateway("example",
            connectivity_type="private",
            subnet_id=aws_subnet["example"]["id"])
        ```

        ## Import

        NAT Gateways can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:ec2/natGateway:NatGateway private_gw nat-05dba92075d71c408
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allocation_id: The Allocation ID of the Elastic IP address for the gateway. Required for `connectivity_type` of `public`.
        :param pulumi.Input[str] connectivity_type: Connectivity type for the gateway. Valid values are `private` and `public`. Defaults to `public`.
        :param pulumi.Input[str] subnet_id: The Subnet ID of the subnet in which to place the gateway.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NatGatewayArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a VPC NAT Gateway.

        ## Example Usage
        ### Public NAT

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2.NatGateway("example",
            allocation_id=aws_eip["example"]["id"],
            subnet_id=aws_subnet["example"]["id"],
            tags={
                "Name": "gw NAT",
            },
            opts=pulumi.ResourceOptions(depends_on=[aws_internet_gateway["example"]]))
        ```
        ### Private NAT

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2.NatGateway("example",
            connectivity_type="private",
            subnet_id=aws_subnet["example"]["id"])
        ```

        ## Import

        NAT Gateways can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:ec2/natGateway:NatGateway private_gw nat-05dba92075d71c408
        ```

        :param str resource_name: The name of the resource.
        :param NatGatewayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NatGatewayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allocation_id: Optional[pulumi.Input[str]] = None,
                 connectivity_type: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NatGatewayArgs.__new__(NatGatewayArgs)

            __props__.__dict__["allocation_id"] = allocation_id
            __props__.__dict__["connectivity_type"] = connectivity_type
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tags_all"] = tags_all
            __props__.__dict__["network_interface_id"] = None
            __props__.__dict__["private_ip"] = None
            __props__.__dict__["public_ip"] = None
        super(NatGateway, __self__).__init__(
            'aws:ec2/natGateway:NatGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allocation_id: Optional[pulumi.Input[str]] = None,
            connectivity_type: Optional[pulumi.Input[str]] = None,
            network_interface_id: Optional[pulumi.Input[str]] = None,
            private_ip: Optional[pulumi.Input[str]] = None,
            public_ip: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'NatGateway':
        """
        Get an existing NatGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allocation_id: The Allocation ID of the Elastic IP address for the gateway. Required for `connectivity_type` of `public`.
        :param pulumi.Input[str] connectivity_type: Connectivity type for the gateway. Valid values are `private` and `public`. Defaults to `public`.
        :param pulumi.Input[str] network_interface_id: The ENI ID of the network interface created by the NAT gateway.
        :param pulumi.Input[str] private_ip: The private IP address of the NAT Gateway.
        :param pulumi.Input[str] public_ip: The public IP address of the NAT Gateway.
        :param pulumi.Input[str] subnet_id: The Subnet ID of the subnet in which to place the gateway.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NatGatewayState.__new__(_NatGatewayState)

        __props__.__dict__["allocation_id"] = allocation_id
        __props__.__dict__["connectivity_type"] = connectivity_type
        __props__.__dict__["network_interface_id"] = network_interface_id
        __props__.__dict__["private_ip"] = private_ip
        __props__.__dict__["public_ip"] = public_ip
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return NatGateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allocationId")
    def allocation_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Allocation ID of the Elastic IP address for the gateway. Required for `connectivity_type` of `public`.
        """
        return pulumi.get(self, "allocation_id")

    @property
    @pulumi.getter(name="connectivityType")
    def connectivity_type(self) -> pulumi.Output[Optional[str]]:
        """
        Connectivity type for the gateway. Valid values are `private` and `public`. Defaults to `public`.
        """
        return pulumi.get(self, "connectivity_type")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Output[str]:
        """
        The ENI ID of the network interface created by the NAT gateway.
        """
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> pulumi.Output[str]:
        """
        The private IP address of the NAT Gateway.
        """
        return pulumi.get(self, "private_ip")

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> pulumi.Output[str]:
        """
        The public IP address of the NAT Gateway.
        """
        return pulumi.get(self, "public_ip")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        The Subnet ID of the subnet in which to place the gateway.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

