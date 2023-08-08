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
    'GetConnectionResult',
    'AwaitableGetConnectionResult',
    'get_connection',
    'get_connection_output',
]

@pulumi.output_type
class GetConnectionResult:
    """
    A collection of values returned by getConnection.
    """
    def __init__(__self__, arn=None, aws_device=None, bandwidth=None, id=None, location=None, name=None, owner_account_id=None, partner_name=None, provider_name=None, tags=None, vlan_id=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if aws_device and not isinstance(aws_device, str):
            raise TypeError("Expected argument 'aws_device' to be a str")
        pulumi.set(__self__, "aws_device", aws_device)
        if bandwidth and not isinstance(bandwidth, str):
            raise TypeError("Expected argument 'bandwidth' to be a str")
        pulumi.set(__self__, "bandwidth", bandwidth)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner_account_id and not isinstance(owner_account_id, str):
            raise TypeError("Expected argument 'owner_account_id' to be a str")
        pulumi.set(__self__, "owner_account_id", owner_account_id)
        if partner_name and not isinstance(partner_name, str):
            raise TypeError("Expected argument 'partner_name' to be a str")
        pulumi.set(__self__, "partner_name", partner_name)
        if provider_name and not isinstance(provider_name, str):
            raise TypeError("Expected argument 'provider_name' to be a str")
        pulumi.set(__self__, "provider_name", provider_name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if vlan_id and not isinstance(vlan_id, int):
            raise TypeError("Expected argument 'vlan_id' to be a int")
        pulumi.set(__self__, "vlan_id", vlan_id)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        ARN of the connection.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsDevice")
    def aws_device(self) -> str:
        """
        Direct Connect endpoint on which the physical connection terminates.
        """
        return pulumi.get(self, "aws_device")

    @property
    @pulumi.getter
    def bandwidth(self) -> str:
        """
        Bandwidth of the connection.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        AWS Direct Connect location where the connection is located.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerAccountId")
    def owner_account_id(self) -> str:
        """
        ID of the AWS account that owns the connection.
        """
        return pulumi.get(self, "owner_account_id")

    @property
    @pulumi.getter(name="partnerName")
    def partner_name(self) -> str:
        """
        The name of the AWS Direct Connect service provider associated with the connection.
        """
        return pulumi.get(self, "partner_name")

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> str:
        """
        Name of the service provider associated with the connection.
        """
        return pulumi.get(self, "provider_name")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Map of tags for the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vlanId")
    def vlan_id(self) -> int:
        """
        The VLAN ID.
        """
        return pulumi.get(self, "vlan_id")


class AwaitableGetConnectionResult(GetConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConnectionResult(
            arn=self.arn,
            aws_device=self.aws_device,
            bandwidth=self.bandwidth,
            id=self.id,
            location=self.location,
            name=self.name,
            owner_account_id=self.owner_account_id,
            partner_name=self.partner_name,
            provider_name=self.provider_name,
            tags=self.tags,
            vlan_id=self.vlan_id)


def get_connection(name: Optional[str] = None,
                   tags: Optional[Mapping[str, str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConnectionResult:
    """
    Retrieve information about a Direct Connect Connection.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.directconnect.get_connection(name="tf-dx-connection")
    ```


    :param str name: Name of the connection to retrieve.
    :param Mapping[str, str] tags: Map of tags for the resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws:directconnect/getConnection:getConnection', __args__, opts=opts, typ=GetConnectionResult).value

    return AwaitableGetConnectionResult(
        arn=pulumi.get(__ret__, 'arn'),
        aws_device=pulumi.get(__ret__, 'aws_device'),
        bandwidth=pulumi.get(__ret__, 'bandwidth'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        owner_account_id=pulumi.get(__ret__, 'owner_account_id'),
        partner_name=pulumi.get(__ret__, 'partner_name'),
        provider_name=pulumi.get(__ret__, 'provider_name'),
        tags=pulumi.get(__ret__, 'tags'),
        vlan_id=pulumi.get(__ret__, 'vlan_id'))


@_utilities.lift_output_func(get_connection)
def get_connection_output(name: Optional[pulumi.Input[str]] = None,
                          tags: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConnectionResult]:
    """
    Retrieve information about a Direct Connect Connection.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.directconnect.get_connection(name="tf-dx-connection")
    ```


    :param str name: Name of the connection to retrieve.
    :param Mapping[str, str] tags: Map of tags for the resource.
    """
    ...
