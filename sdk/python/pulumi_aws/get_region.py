# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetRegionResult',
    'AwaitableGetRegionResult',
    'get_region',
    'get_region_output',
]

@pulumi.output_type
class GetRegionResult:
    """
    A collection of values returned by getRegion.
    """
    def __init__(__self__, description=None, endpoint=None, id=None, name=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        pulumi.set(__self__, "endpoint", endpoint)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Region's description in this format: "Location (Region name)".
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        """
        EC2 endpoint for the selected region.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the selected region.
        """
        return pulumi.get(self, "name")


class AwaitableGetRegionResult(GetRegionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionResult(
            description=self.description,
            endpoint=self.endpoint,
            id=self.id,
            name=self.name)


def get_region(endpoint: Optional[str] = None,
               id: Optional[str] = None,
               name: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegionResult:
    """
    `get_region` provides details about a specific AWS region.

    As well as validating a given region name this resource can be used to
    discover the name of the region configured within the provider. The latter
    can be useful in a child module which is inheriting an AWS provider
    configuration from its parent module.

    ## Example Usage

    The following example shows how the resource might be used to obtain
    the name of the AWS region configured on the provider.

    ```python
    import pulumi
    import pulumi_aws as aws

    current = aws.get_region()
    ```


    :param str endpoint: EC2 endpoint of the region to select.
    :param str name: Full name of the region to select.
    """
    __args__ = dict()
    __args__['endpoint'] = endpoint
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws:index/getRegion:getRegion', __args__, opts=opts, typ=GetRegionResult).value

    return AwaitableGetRegionResult(
        description=pulumi.get(__ret__, 'description'),
        endpoint=pulumi.get(__ret__, 'endpoint'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'))


@_utilities.lift_output_func(get_region)
def get_region_output(endpoint: Optional[pulumi.Input[Optional[str]]] = None,
                      id: Optional[pulumi.Input[Optional[str]]] = None,
                      name: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRegionResult]:
    """
    `get_region` provides details about a specific AWS region.

    As well as validating a given region name this resource can be used to
    discover the name of the region configured within the provider. The latter
    can be useful in a child module which is inheriting an AWS provider
    configuration from its parent module.

    ## Example Usage

    The following example shows how the resource might be used to obtain
    the name of the AWS region configured on the provider.

    ```python
    import pulumi
    import pulumi_aws as aws

    current = aws.get_region()
    ```


    :param str endpoint: EC2 endpoint of the region to select.
    :param str name: Full name of the region to select.
    """
    ...
