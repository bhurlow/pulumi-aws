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
    'GetServiceResult',
    'AwaitableGetServiceResult',
    'get_service',
    'get_service_output',
]

@pulumi.output_type
class GetServiceResult:
    """
    A collection of values returned by getService.
    """
    def __init__(__self__, dns_name=None, id=None, partition=None, region=None, reverse_dns_name=None, reverse_dns_prefix=None, service_id=None, supported=None):
        if dns_name and not isinstance(dns_name, str):
            raise TypeError("Expected argument 'dns_name' to be a str")
        pulumi.set(__self__, "dns_name", dns_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if partition and not isinstance(partition, str):
            raise TypeError("Expected argument 'partition' to be a str")
        pulumi.set(__self__, "partition", partition)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if reverse_dns_name and not isinstance(reverse_dns_name, str):
            raise TypeError("Expected argument 'reverse_dns_name' to be a str")
        pulumi.set(__self__, "reverse_dns_name", reverse_dns_name)
        if reverse_dns_prefix and not isinstance(reverse_dns_prefix, str):
            raise TypeError("Expected argument 'reverse_dns_prefix' to be a str")
        pulumi.set(__self__, "reverse_dns_prefix", reverse_dns_prefix)
        if service_id and not isinstance(service_id, str):
            raise TypeError("Expected argument 'service_id' to be a str")
        pulumi.set(__self__, "service_id", service_id)
        if supported and not isinstance(supported, bool):
            raise TypeError("Expected argument 'supported' to be a bool")
        pulumi.set(__self__, "supported", supported)

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> str:
        return pulumi.get(self, "dns_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def partition(self) -> str:
        return pulumi.get(self, "partition")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="reverseDnsName")
    def reverse_dns_name(self) -> str:
        return pulumi.get(self, "reverse_dns_name")

    @property
    @pulumi.getter(name="reverseDnsPrefix")
    def reverse_dns_prefix(self) -> str:
        return pulumi.get(self, "reverse_dns_prefix")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> str:
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter
    def supported(self) -> bool:
        """
        Whether the service is supported in the region's partition. New services may not be listed immediately as supported.
        """
        return pulumi.get(self, "supported")


class AwaitableGetServiceResult(GetServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceResult(
            dns_name=self.dns_name,
            id=self.id,
            partition=self.partition,
            region=self.region,
            reverse_dns_name=self.reverse_dns_name,
            reverse_dns_prefix=self.reverse_dns_prefix,
            service_id=self.service_id,
            supported=self.supported)


def get_service(dns_name: Optional[str] = None,
                id: Optional[str] = None,
                region: Optional[str] = None,
                reverse_dns_name: Optional[str] = None,
                reverse_dns_prefix: Optional[str] = None,
                service_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceResult:
    """
    Use this data source to compose and decompose AWS service DNS names.

    ## Example Usage
    ### Get Service DNS Name

    ```python
    import pulumi
    import pulumi_aws as aws

    current = aws.get_region()
    test = aws.get_service(region=current.name,
        service_id="ec2")
    ```
    ### Use Service Reverse DNS Name to Get Components

    ```python
    import pulumi
    import pulumi_aws as aws

    s3 = aws.get_service(reverse_dns_name="cn.com.amazonaws.cn-north-1.s3")
    ```
    ### Determine Regional Support for a Service

    ```python
    import pulumi
    import pulumi_aws as aws

    s3 = aws.get_service(reverse_dns_name="com.amazonaws.us-gov-west-1.waf")
    ```


    :param str dns_name: DNS name of the service (_e.g.,_ `rds.us-east-1.amazonaws.com`). One of `dns_name`, `reverse_dns_name`, or `service_id` is required.
    :param str region: Region of the service (_e.g.,_ `us-west-2`, `ap-northeast-1`).
    :param str reverse_dns_name: Reverse DNS name of the service (_e.g.,_ `com.amazonaws.us-west-2.s3`). One of `dns_name`, `reverse_dns_name`, or `service_id` is required.
    :param str reverse_dns_prefix: Prefix of the service (_e.g.,_ `com.amazonaws` in AWS Commercial, `cn.com.amazonaws` in AWS China).
    :param str service_id: Service (_e.g.,_ `s3`, `rds`, `ec2`). One of `dns_name`, `reverse_dns_name`, or `service_id` is required.
    """
    __args__ = dict()
    __args__['dnsName'] = dns_name
    __args__['id'] = id
    __args__['region'] = region
    __args__['reverseDnsName'] = reverse_dns_name
    __args__['reverseDnsPrefix'] = reverse_dns_prefix
    __args__['serviceId'] = service_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws:index/getService:getService', __args__, opts=opts, typ=GetServiceResult).value

    return AwaitableGetServiceResult(
        dns_name=pulumi.get(__ret__, 'dns_name'),
        id=pulumi.get(__ret__, 'id'),
        partition=pulumi.get(__ret__, 'partition'),
        region=pulumi.get(__ret__, 'region'),
        reverse_dns_name=pulumi.get(__ret__, 'reverse_dns_name'),
        reverse_dns_prefix=pulumi.get(__ret__, 'reverse_dns_prefix'),
        service_id=pulumi.get(__ret__, 'service_id'),
        supported=pulumi.get(__ret__, 'supported'))


@_utilities.lift_output_func(get_service)
def get_service_output(dns_name: Optional[pulumi.Input[Optional[str]]] = None,
                       id: Optional[pulumi.Input[Optional[str]]] = None,
                       region: Optional[pulumi.Input[Optional[str]]] = None,
                       reverse_dns_name: Optional[pulumi.Input[Optional[str]]] = None,
                       reverse_dns_prefix: Optional[pulumi.Input[Optional[str]]] = None,
                       service_id: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServiceResult]:
    """
    Use this data source to compose and decompose AWS service DNS names.

    ## Example Usage
    ### Get Service DNS Name

    ```python
    import pulumi
    import pulumi_aws as aws

    current = aws.get_region()
    test = aws.get_service(region=current.name,
        service_id="ec2")
    ```
    ### Use Service Reverse DNS Name to Get Components

    ```python
    import pulumi
    import pulumi_aws as aws

    s3 = aws.get_service(reverse_dns_name="cn.com.amazonaws.cn-north-1.s3")
    ```
    ### Determine Regional Support for a Service

    ```python
    import pulumi
    import pulumi_aws as aws

    s3 = aws.get_service(reverse_dns_name="com.amazonaws.us-gov-west-1.waf")
    ```


    :param str dns_name: DNS name of the service (_e.g.,_ `rds.us-east-1.amazonaws.com`). One of `dns_name`, `reverse_dns_name`, or `service_id` is required.
    :param str region: Region of the service (_e.g.,_ `us-west-2`, `ap-northeast-1`).
    :param str reverse_dns_name: Reverse DNS name of the service (_e.g.,_ `com.amazonaws.us-west-2.s3`). One of `dns_name`, `reverse_dns_name`, or `service_id` is required.
    :param str reverse_dns_prefix: Prefix of the service (_e.g.,_ `com.amazonaws` in AWS Commercial, `cn.com.amazonaws` in AWS China).
    :param str service_id: Service (_e.g.,_ `s3`, `rds`, `ec2`). One of `dns_name`, `reverse_dns_name`, or `service_id` is required.
    """
    ...
