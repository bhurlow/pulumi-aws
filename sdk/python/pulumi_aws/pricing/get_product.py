# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetProductResult:
    """
    A collection of values returned by getProduct.
    """
    def __init__(__self__, filters=None, id=None, result=None, service_code=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        __self__.result = result
        """
        Set to the product returned from the API.
        """
        if service_code and not isinstance(service_code, str):
            raise TypeError("Expected argument 'service_code' to be a str")
        __self__.service_code = service_code
class AwaitableGetProductResult(GetProductResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProductResult(
            filters=self.filters,
            id=self.id,
            result=self.result,
            service_code=self.service_code)

def get_product(filters=None,service_code=None,opts=None):
    """
    Use this data source to get the pricing information of all products in AWS.
    This data source is only available in a us-east-1 or ap-south-1 provider.

    ## Example Usage



    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.pricing.get_product(filters=[
            {
                "field": "instanceType",
                "value": "c5.xlarge",
            },
            {
                "field": "operatingSystem",
                "value": "Linux",
            },
            {
                "field": "location",
                "value": "US East (N. Virginia)",
            },
            {
                "field": "preInstalledSw",
                "value": "NA",
            },
            {
                "field": "licenseModel",
                "value": "No License required",
            },
            {
                "field": "tenancy",
                "value": "Shared",
            },
            {
                "field": "capacitystatus",
                "value": "Used",
            },
        ],
        service_code="AmazonEC2")
    ```


    :param list filters: A list of filters. Passed directly to the API (see GetProducts API reference). These filters must describe a single product, this resource will fail if more than one product is returned by the API.
    :param str service_code: The code of the service. Available service codes can be fetched using the DescribeServices pricing API call.

    The **filters** object supports the following:

      * `field` (`str`) - The product attribute name that you want to filter on.
      * `value` (`str`) - The product attribute value that you want to filter on.
    """
    __args__ = dict()


    __args__['filters'] = filters
    __args__['serviceCode'] = service_code
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:pricing/getProduct:getProduct', __args__, opts=opts).value

    return AwaitableGetProductResult(
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        result=__ret__.get('result'),
        service_code=__ret__.get('serviceCode'))
