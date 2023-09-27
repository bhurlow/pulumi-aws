# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetOrganizationalUnitResult',
    'AwaitableGetOrganizationalUnitResult',
    'get_organizational_unit',
    'get_organizational_unit_output',
]

@pulumi.output_type
class GetOrganizationalUnitResult:
    """
    A collection of values returned by getOrganizationalUnit.
    """
    def __init__(__self__, arn=None, id=None, name=None, parent_id=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parent_id and not isinstance(parent_id, str):
            raise TypeError("Expected argument 'parent_id' to be a str")
        pulumi.set(__self__, "parent_id", parent_id)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        ARN of the organizational unit
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> str:
        return pulumi.get(self, "parent_id")


class AwaitableGetOrganizationalUnitResult(GetOrganizationalUnitResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrganizationalUnitResult(
            arn=self.arn,
            id=self.id,
            name=self.name,
            parent_id=self.parent_id)


def get_organizational_unit(name: Optional[str] = None,
                            parent_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrganizationalUnitResult:
    """
    Data source for getting an AWS Organizations Organizational Unit.

    ## Example Usage
    ### Basic Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    org = aws.organizations.get_organization()
    ou = aws.organizations.get_organizational_unit(parent_id=org.roots[0].id,
        name="dev")
    ```


    :param str name: Name of the organizational unit
    :param str parent_id: Parent ID of the organizational unit.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['parentId'] = parent_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws:organizations/getOrganizationalUnit:getOrganizationalUnit', __args__, opts=opts, typ=GetOrganizationalUnitResult).value

    return AwaitableGetOrganizationalUnitResult(
        arn=pulumi.get(__ret__, 'arn'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        parent_id=pulumi.get(__ret__, 'parent_id'))


@_utilities.lift_output_func(get_organizational_unit)
def get_organizational_unit_output(name: Optional[pulumi.Input[str]] = None,
                                   parent_id: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOrganizationalUnitResult]:
    """
    Data source for getting an AWS Organizations Organizational Unit.

    ## Example Usage
    ### Basic Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    org = aws.organizations.get_organization()
    ou = aws.organizations.get_organizational_unit(parent_id=org.roots[0].id,
        name="dev")
    ```


    :param str name: Name of the organizational unit
    :param str parent_id: Parent ID of the organizational unit.
    """
    ...
