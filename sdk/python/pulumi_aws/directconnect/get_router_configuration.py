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

__all__ = [
    'GetRouterConfigurationResult',
    'AwaitableGetRouterConfigurationResult',
    'get_router_configuration',
    'get_router_configuration_output',
]

@pulumi.output_type
class GetRouterConfigurationResult:
    """
    A collection of values returned by getRouterConfiguration.
    """
    def __init__(__self__, customer_router_config=None, id=None, router_type_identifier=None, routers=None, virtual_interface_id=None, virtual_interface_name=None):
        if customer_router_config and not isinstance(customer_router_config, str):
            raise TypeError("Expected argument 'customer_router_config' to be a str")
        pulumi.set(__self__, "customer_router_config", customer_router_config)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if router_type_identifier and not isinstance(router_type_identifier, str):
            raise TypeError("Expected argument 'router_type_identifier' to be a str")
        pulumi.set(__self__, "router_type_identifier", router_type_identifier)
        if routers and not isinstance(routers, list):
            raise TypeError("Expected argument 'routers' to be a list")
        pulumi.set(__self__, "routers", routers)
        if virtual_interface_id and not isinstance(virtual_interface_id, str):
            raise TypeError("Expected argument 'virtual_interface_id' to be a str")
        pulumi.set(__self__, "virtual_interface_id", virtual_interface_id)
        if virtual_interface_name and not isinstance(virtual_interface_name, str):
            raise TypeError("Expected argument 'virtual_interface_name' to be a str")
        pulumi.set(__self__, "virtual_interface_name", virtual_interface_name)

    @property
    @pulumi.getter(name="customerRouterConfig")
    def customer_router_config(self) -> str:
        """
        Instructions for configuring your router
        """
        return pulumi.get(self, "customer_router_config")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="routerTypeIdentifier")
    def router_type_identifier(self) -> str:
        """
        Router type identifier
        """
        return pulumi.get(self, "router_type_identifier")

    @property
    @pulumi.getter
    def routers(self) -> Sequence['outputs.GetRouterConfigurationRouterResult']:
        """
        Block of the router type details
        """
        return pulumi.get(self, "routers")

    @property
    @pulumi.getter(name="virtualInterfaceId")
    def virtual_interface_id(self) -> str:
        return pulumi.get(self, "virtual_interface_id")

    @property
    @pulumi.getter(name="virtualInterfaceName")
    def virtual_interface_name(self) -> str:
        return pulumi.get(self, "virtual_interface_name")


class AwaitableGetRouterConfigurationResult(GetRouterConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterConfigurationResult(
            customer_router_config=self.customer_router_config,
            id=self.id,
            router_type_identifier=self.router_type_identifier,
            routers=self.routers,
            virtual_interface_id=self.virtual_interface_id,
            virtual_interface_name=self.virtual_interface_name)


def get_router_configuration(router_type_identifier: Optional[str] = None,
                             virtual_interface_id: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterConfigurationResult:
    """
    Data source for retrieving Router Configuration instructions for a given AWS Direct Connect Virtual Interface and Router Type.

    ## Example Usage
    ### Basic Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.directconnect.get_router_configuration(router_type_identifier="CiscoSystemsInc-2900SeriesRouters-IOS124",
        virtual_interface_id="dxvif-abcde123")
    ```


    :param str router_type_identifier: ID of the Router Type. For example: `CiscoSystemsInc-2900SeriesRouters-IOS124`
    :param str virtual_interface_id: ID of the Direct Connect Virtual Interface
    """
    __args__ = dict()
    __args__['routerTypeIdentifier'] = router_type_identifier
    __args__['virtualInterfaceId'] = virtual_interface_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws:directconnect/getRouterConfiguration:getRouterConfiguration', __args__, opts=opts, typ=GetRouterConfigurationResult).value

    return AwaitableGetRouterConfigurationResult(
        customer_router_config=__ret__.customer_router_config,
        id=__ret__.id,
        router_type_identifier=__ret__.router_type_identifier,
        routers=__ret__.routers,
        virtual_interface_id=__ret__.virtual_interface_id,
        virtual_interface_name=__ret__.virtual_interface_name)


@_utilities.lift_output_func(get_router_configuration)
def get_router_configuration_output(router_type_identifier: Optional[pulumi.Input[str]] = None,
                                    virtual_interface_id: Optional[pulumi.Input[str]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRouterConfigurationResult]:
    """
    Data source for retrieving Router Configuration instructions for a given AWS Direct Connect Virtual Interface and Router Type.

    ## Example Usage
    ### Basic Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.directconnect.get_router_configuration(router_type_identifier="CiscoSystemsInc-2900SeriesRouters-IOS124",
        virtual_interface_id="dxvif-abcde123")
    ```


    :param str router_type_identifier: ID of the Router Type. For example: `CiscoSystemsInc-2900SeriesRouters-IOS124`
    :param str virtual_interface_id: ID of the Direct Connect Virtual Interface
    """
    ...
