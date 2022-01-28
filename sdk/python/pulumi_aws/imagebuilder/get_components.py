# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetComponentsResult',
    'AwaitableGetComponentsResult',
    'get_components',
    'get_components_output',
]

@pulumi.output_type
class GetComponentsResult:
    """
    A collection of values returned by getComponents.
    """
    def __init__(__self__, arns=None, filters=None, id=None, names=None, owner=None):
        if arns and not isinstance(arns, list):
            raise TypeError("Expected argument 'arns' to be a list")
        pulumi.set(__self__, "arns", arns)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if owner and not isinstance(owner, str):
            raise TypeError("Expected argument 'owner' to be a str")
        pulumi.set(__self__, "owner", owner)

    @property
    @pulumi.getter
    def arns(self) -> Sequence[str]:
        """
        Set of ARNs of the matched Image Builder Components.
        """
        return pulumi.get(self, "arns")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetComponentsFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        Set of names of the matched Image Builder Components.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter
    def owner(self) -> Optional[str]:
        return pulumi.get(self, "owner")


class AwaitableGetComponentsResult(GetComponentsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetComponentsResult(
            arns=self.arns,
            filters=self.filters,
            id=self.id,
            names=self.names,
            owner=self.owner)


def get_components(filters: Optional[Sequence[pulumi.InputType['GetComponentsFilterArgs']]] = None,
                   owner: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetComponentsResult:
    """
    Use this data source to get the ARNs and names of Image Builder Components matching the specified criteria.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.imagebuilder.get_components(filters=[aws.imagebuilder.GetComponentsFilterArgs(
            name="platform",
            values=["Linux"],
        )],
        owner="Self")
    ```


    :param Sequence[pulumi.InputType['GetComponentsFilterArgs']] filters: Configuration block(s) for filtering. Detailed below.
    :param str owner: The owner of the image recipes. Valid values are `Self`, `Shared` and `Amazon`. Defaults to `Self`.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['owner'] = owner
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:imagebuilder/getComponents:getComponents', __args__, opts=opts, typ=GetComponentsResult).value

    return AwaitableGetComponentsResult(
        arns=__ret__.arns,
        filters=__ret__.filters,
        id=__ret__.id,
        names=__ret__.names,
        owner=__ret__.owner)


@_utilities.lift_output_func(get_components)
def get_components_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetComponentsFilterArgs']]]]] = None,
                          owner: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetComponentsResult]:
    """
    Use this data source to get the ARNs and names of Image Builder Components matching the specified criteria.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.imagebuilder.get_components(filters=[aws.imagebuilder.GetComponentsFilterArgs(
            name="platform",
            values=["Linux"],
        )],
        owner="Self")
    ```


    :param Sequence[pulumi.InputType['GetComponentsFilterArgs']] filters: Configuration block(s) for filtering. Detailed below.
    :param str owner: The owner of the image recipes. Valid values are `Self`, `Shared` and `Amazon`. Defaults to `Self`.
    """
    ...
