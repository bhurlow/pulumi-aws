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
    'GetInvocationResult',
    'AwaitableGetInvocationResult',
    'get_invocation',
    'get_invocation_output',
]

@pulumi.output_type
class GetInvocationResult:
    """
    A collection of values returned by getInvocation.
    """
    def __init__(__self__, function_name=None, id=None, input=None, qualifier=None, result=None):
        if function_name and not isinstance(function_name, str):
            raise TypeError("Expected argument 'function_name' to be a str")
        pulumi.set(__self__, "function_name", function_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if input and not isinstance(input, str):
            raise TypeError("Expected argument 'input' to be a str")
        pulumi.set(__self__, "input", input)
        if qualifier and not isinstance(qualifier, str):
            raise TypeError("Expected argument 'qualifier' to be a str")
        pulumi.set(__self__, "qualifier", qualifier)
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> str:
        return pulumi.get(self, "function_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def input(self) -> str:
        return pulumi.get(self, "input")

    @property
    @pulumi.getter
    def qualifier(self) -> Optional[str]:
        return pulumi.get(self, "qualifier")

    @property
    @pulumi.getter
    def result(self) -> str:
        """
        String result of the lambda function invocation.
        """
        return pulumi.get(self, "result")


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
            result=self.result)


def get_invocation(function_name: Optional[str] = None,
                   input: Optional[str] = None,
                   qualifier: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInvocationResult:
    """
    Use this data source to invoke custom lambda functions as data source.
    The lambda function is invoked with [RequestResponse](https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html#API_Invoke_RequestSyntax)
    invocation type.

    > **NOTE:** If you get a `KMSAccessDeniedException: Lambda was unable to decrypt the environment variables because KMS access was denied` error when invoking an `lambda.Function` with environment variables, the IAM role associated with the function may have been deleted and recreated _after_ the function was created. You can fix the problem two ways: 1) updating the function's role to another role and then updating it back again to the recreated role, or 2) by using Pulumi to `taint` the function and `apply` your configuration again to recreate the function. (When you create a function, Lambda grants permissions on the KMS key to the function's IAM role. If the IAM role is recreated, the grant is no longer valid. Changing the function's role or recreating the function causes Lambda to update the grant.)


    :param str function_name: Name of the lambda function.
    :param str input: String in JSON format that is passed as payload to the lambda function.
    :param str qualifier: Qualifier (a.k.a version) of the lambda function. Defaults
           to `$LATEST`.
    """
    __args__ = dict()
    __args__['functionName'] = function_name
    __args__['input'] = input
    __args__['qualifier'] = qualifier
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws:lambda/getInvocation:getInvocation', __args__, opts=opts, typ=GetInvocationResult).value

    return AwaitableGetInvocationResult(
        function_name=__ret__.function_name,
        id=__ret__.id,
        input=__ret__.input,
        qualifier=__ret__.qualifier,
        result=__ret__.result)


@_utilities.lift_output_func(get_invocation)
def get_invocation_output(function_name: Optional[pulumi.Input[str]] = None,
                          input: Optional[pulumi.Input[str]] = None,
                          qualifier: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInvocationResult]:
    """
    Use this data source to invoke custom lambda functions as data source.
    The lambda function is invoked with [RequestResponse](https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html#API_Invoke_RequestSyntax)
    invocation type.

    > **NOTE:** If you get a `KMSAccessDeniedException: Lambda was unable to decrypt the environment variables because KMS access was denied` error when invoking an `lambda.Function` with environment variables, the IAM role associated with the function may have been deleted and recreated _after_ the function was created. You can fix the problem two ways: 1) updating the function's role to another role and then updating it back again to the recreated role, or 2) by using Pulumi to `taint` the function and `apply` your configuration again to recreate the function. (When you create a function, Lambda grants permissions on the KMS key to the function's IAM role. If the IAM role is recreated, the grant is no longer valid. Changing the function's role or recreating the function causes Lambda to update the grant.)


    :param str function_name: Name of the lambda function.
    :param str input: String in JSON format that is passed as payload to the lambda function.
    :param str qualifier: Qualifier (a.k.a version) of the lambda function. Defaults
           to `$LATEST`.
    """
    ...
