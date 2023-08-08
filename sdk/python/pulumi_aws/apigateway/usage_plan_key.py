# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['UsagePlanKeyArgs', 'UsagePlanKey']

@pulumi.input_type
class UsagePlanKeyArgs:
    def __init__(__self__, *,
                 key_id: pulumi.Input[str],
                 key_type: pulumi.Input[str],
                 usage_plan_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a UsagePlanKey resource.
        :param pulumi.Input[str] key_id: Identifier of the API key resource.
        :param pulumi.Input[str] key_type: Type of the API key resource. Currently, the valid key type is API_KEY.
        :param pulumi.Input[str] usage_plan_id: Id of the usage plan resource representing to associate the key to.
        """
        pulumi.set(__self__, "key_id", key_id)
        pulumi.set(__self__, "key_type", key_type)
        pulumi.set(__self__, "usage_plan_id", usage_plan_id)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Input[str]:
        """
        Identifier of the API key resource.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> pulumi.Input[str]:
        """
        Type of the API key resource. Currently, the valid key type is API_KEY.
        """
        return pulumi.get(self, "key_type")

    @key_type.setter
    def key_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_type", value)

    @property
    @pulumi.getter(name="usagePlanId")
    def usage_plan_id(self) -> pulumi.Input[str]:
        """
        Id of the usage plan resource representing to associate the key to.
        """
        return pulumi.get(self, "usage_plan_id")

    @usage_plan_id.setter
    def usage_plan_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "usage_plan_id", value)


@pulumi.input_type
class _UsagePlanKeyState:
    def __init__(__self__, *,
                 key_id: Optional[pulumi.Input[str]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 usage_plan_id: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UsagePlanKey resources.
        :param pulumi.Input[str] key_id: Identifier of the API key resource.
        :param pulumi.Input[str] key_type: Type of the API key resource. Currently, the valid key type is API_KEY.
        :param pulumi.Input[str] name: Name of a usage plan key.
        :param pulumi.Input[str] usage_plan_id: Id of the usage plan resource representing to associate the key to.
        :param pulumi.Input[str] value: Value of a usage plan key.
        """
        if key_id is not None:
            pulumi.set(__self__, "key_id", key_id)
        if key_type is not None:
            pulumi.set(__self__, "key_type", key_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if usage_plan_id is not None:
            pulumi.set(__self__, "usage_plan_id", usage_plan_id)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the API key resource.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the API key resource. Currently, the valid key type is API_KEY.
        """
        return pulumi.get(self, "key_type")

    @key_type.setter
    def key_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of a usage plan key.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="usagePlanId")
    def usage_plan_id(self) -> Optional[pulumi.Input[str]]:
        """
        Id of the usage plan resource representing to associate the key to.
        """
        return pulumi.get(self, "usage_plan_id")

    @usage_plan_id.setter
    def usage_plan_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "usage_plan_id", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        Value of a usage plan key.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


class UsagePlanKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 usage_plan_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an API Gateway Usage Plan Key.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.apigateway.RestApi("test")
        # ...
        myusageplan = aws.apigateway.UsagePlan("myusageplan", api_stages=[aws.apigateway.UsagePlanApiStageArgs(
            api_id=test.id,
            stage=aws_api_gateway_stage["foo"]["stage_name"],
        )])
        mykey = aws.apigateway.ApiKey("mykey")
        main = aws.apigateway.UsagePlanKey("main",
            key_id=mykey.id,
            key_type="API_KEY",
            usage_plan_id=myusageplan.id)
        ```

        ## Import

        terraform import {

         to = aws_api_gateway_usage_plan_key.key

         id = "12345abcde/zzz" } Using `pulumi import`, import AWS API Gateway Usage Plan Key using the `USAGE-PLAN-ID/USAGE-PLAN-KEY-ID`. For exampleconsole % pulumi import aws_api_gateway_usage_plan_key.key 12345abcde/zzz

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_id: Identifier of the API key resource.
        :param pulumi.Input[str] key_type: Type of the API key resource. Currently, the valid key type is API_KEY.
        :param pulumi.Input[str] usage_plan_id: Id of the usage plan resource representing to associate the key to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UsagePlanKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an API Gateway Usage Plan Key.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.apigateway.RestApi("test")
        # ...
        myusageplan = aws.apigateway.UsagePlan("myusageplan", api_stages=[aws.apigateway.UsagePlanApiStageArgs(
            api_id=test.id,
            stage=aws_api_gateway_stage["foo"]["stage_name"],
        )])
        mykey = aws.apigateway.ApiKey("mykey")
        main = aws.apigateway.UsagePlanKey("main",
            key_id=mykey.id,
            key_type="API_KEY",
            usage_plan_id=myusageplan.id)
        ```

        ## Import

        terraform import {

         to = aws_api_gateway_usage_plan_key.key

         id = "12345abcde/zzz" } Using `pulumi import`, import AWS API Gateway Usage Plan Key using the `USAGE-PLAN-ID/USAGE-PLAN-KEY-ID`. For exampleconsole % pulumi import aws_api_gateway_usage_plan_key.key 12345abcde/zzz

        :param str resource_name: The name of the resource.
        :param UsagePlanKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UsagePlanKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 usage_plan_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UsagePlanKeyArgs.__new__(UsagePlanKeyArgs)

            if key_id is None and not opts.urn:
                raise TypeError("Missing required property 'key_id'")
            __props__.__dict__["key_id"] = key_id
            if key_type is None and not opts.urn:
                raise TypeError("Missing required property 'key_type'")
            __props__.__dict__["key_type"] = key_type
            if usage_plan_id is None and not opts.urn:
                raise TypeError("Missing required property 'usage_plan_id'")
            __props__.__dict__["usage_plan_id"] = usage_plan_id
            __props__.__dict__["name"] = None
            __props__.__dict__["value"] = None
        super(UsagePlanKey, __self__).__init__(
            'aws:apigateway/usagePlanKey:UsagePlanKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            key_id: Optional[pulumi.Input[str]] = None,
            key_type: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            usage_plan_id: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None) -> 'UsagePlanKey':
        """
        Get an existing UsagePlanKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_id: Identifier of the API key resource.
        :param pulumi.Input[str] key_type: Type of the API key resource. Currently, the valid key type is API_KEY.
        :param pulumi.Input[str] name: Name of a usage plan key.
        :param pulumi.Input[str] usage_plan_id: Id of the usage plan resource representing to associate the key to.
        :param pulumi.Input[str] value: Value of a usage plan key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UsagePlanKeyState.__new__(_UsagePlanKeyState)

        __props__.__dict__["key_id"] = key_id
        __props__.__dict__["key_type"] = key_type
        __props__.__dict__["name"] = name
        __props__.__dict__["usage_plan_id"] = usage_plan_id
        __props__.__dict__["value"] = value
        return UsagePlanKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Output[str]:
        """
        Identifier of the API key resource.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> pulumi.Output[str]:
        """
        Type of the API key resource. Currently, the valid key type is API_KEY.
        """
        return pulumi.get(self, "key_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of a usage plan key.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="usagePlanId")
    def usage_plan_id(self) -> pulumi.Output[str]:
        """
        Id of the usage plan resource representing to associate the key to.
        """
        return pulumi.get(self, "usage_plan_id")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        Value of a usage plan key.
        """
        return pulumi.get(self, "value")

