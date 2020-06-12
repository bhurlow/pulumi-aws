# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetInvocationResult:
    """
    A collection of values returned by getInvocation.
    """
    def __init__(__self__, function_name=None, id=None, input=None, qualifier=None, result=None, result_map=None):
        if function_name and not isinstance(function_name, str):
            raise TypeError("Expected argument 'function_name' to be a str")
        __self__.function_name = function_name
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if input and not isinstance(input, str):
            raise TypeError("Expected argument 'input' to be a str")
        __self__.input = input
        if qualifier and not isinstance(qualifier, str):
            raise TypeError("Expected argument 'qualifier' to be a str")
        __self__.qualifier = qualifier
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        __self__.result = result
        """
        String result of the lambda function invocation.
        """
        if result_map and not isinstance(result_map, dict):
            raise TypeError("Expected argument 'result_map' to be a dict")
        if result_map is not None:
            warnings.warn("use `result` attribute with jsondecode() function", DeprecationWarning)
            pulumi.log.warn("result_map is deprecated: use `result` attribute with jsondecode() function")
        __self__.result_map = result_map
        """
        (**DEPRECATED**) This field is set only if result is a map of primitive types, where the map is string keys and string values.
        """
class AwaitableGetInvocationResult(GetInvocationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInvocationResult(
            function_name=self.function_name,
            id=self.id,
            input=self.input,
            qualifier=self.qualifier,
            result=self.result,
            result_map=self.result_map)

def get_invocation(function_name=None,input=None,qualifier=None,opts=None):
    """
    Use this data source to invoke custom lambda functions as data source.
    The lambda function is invoked with [RequestResponse](https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html#API_Invoke_RequestSyntax)
    invocation type.


    :param str function_name: The name of the lambda function.
    :param str input: A string in JSON format that is passed as payload to the lambda function.
    :param str qualifier: The qualifier (a.k.a version) of the lambda function. Defaults
           to `$LATEST`.
    """
    __args__ = dict()


    __args__['functionName'] = function_name
    __args__['input'] = input
    __args__['qualifier'] = qualifier
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:lambda/getInvocation:getInvocation', __args__, opts=opts).value

    return AwaitableGetInvocationResult(
        function_name=__ret__.get('functionName'),
        id=__ret__.get('id'),
        input=__ret__.get('input'),
        qualifier=__ret__.get('qualifier'),
        result=__ret__.get('result'),
        result_map=__ret__.get('resultMap'))
