# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AccountArgs', 'Account']

@pulumi.input_type
class AccountArgs:
    def __init__(__self__, *,
                 enable_default_standards: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Account resource.
        :param pulumi.Input[bool] enable_default_standards: Whether to enable the security standards that Security Hub has designated as automatically enabled including: ` AWS Foundational Security Best Practices v1.0.0` and `CIS AWS Foundations Benchmark v1.2.0`. Defaults to `true`.
        """
        if enable_default_standards is not None:
            pulumi.set(__self__, "enable_default_standards", enable_default_standards)

    @property
    @pulumi.getter(name="enableDefaultStandards")
    def enable_default_standards(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable the security standards that Security Hub has designated as automatically enabled including: ` AWS Foundational Security Best Practices v1.0.0` and `CIS AWS Foundations Benchmark v1.2.0`. Defaults to `true`.
        """
        return pulumi.get(self, "enable_default_standards")

    @enable_default_standards.setter
    def enable_default_standards(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_default_standards", value)


@pulumi.input_type
class _AccountState:
    def __init__(__self__, *,
                 enable_default_standards: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering Account resources.
        :param pulumi.Input[bool] enable_default_standards: Whether to enable the security standards that Security Hub has designated as automatically enabled including: ` AWS Foundational Security Best Practices v1.0.0` and `CIS AWS Foundations Benchmark v1.2.0`. Defaults to `true`.
        """
        if enable_default_standards is not None:
            pulumi.set(__self__, "enable_default_standards", enable_default_standards)

    @property
    @pulumi.getter(name="enableDefaultStandards")
    def enable_default_standards(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable the security standards that Security Hub has designated as automatically enabled including: ` AWS Foundational Security Best Practices v1.0.0` and `CIS AWS Foundations Benchmark v1.2.0`. Defaults to `true`.
        """
        return pulumi.get(self, "enable_default_standards")

    @enable_default_standards.setter
    def enable_default_standards(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_default_standards", value)


class Account(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_default_standards: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Enables Security Hub for this AWS account.

        > **NOTE:** Destroying this resource will disable Security Hub for this AWS account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.securityhub.Account("example")
        ```

        ## Import

        An existing Security Hub enabled account can be imported using the AWS account ID, e.g.,

        ```sh
         $ pulumi import aws:securityhub/account:Account example 123456789012
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enable_default_standards: Whether to enable the security standards that Security Hub has designated as automatically enabled including: ` AWS Foundational Security Best Practices v1.0.0` and `CIS AWS Foundations Benchmark v1.2.0`. Defaults to `true`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AccountArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Enables Security Hub for this AWS account.

        > **NOTE:** Destroying this resource will disable Security Hub for this AWS account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.securityhub.Account("example")
        ```

        ## Import

        An existing Security Hub enabled account can be imported using the AWS account ID, e.g.,

        ```sh
         $ pulumi import aws:securityhub/account:Account example 123456789012
        ```

        :param str resource_name: The name of the resource.
        :param AccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_default_standards: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccountArgs.__new__(AccountArgs)

            __props__.__dict__["enable_default_standards"] = enable_default_standards
        super(Account, __self__).__init__(
            'aws:securityhub/account:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enable_default_standards: Optional[pulumi.Input[bool]] = None) -> 'Account':
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enable_default_standards: Whether to enable the security standards that Security Hub has designated as automatically enabled including: ` AWS Foundational Security Best Practices v1.0.0` and `CIS AWS Foundations Benchmark v1.2.0`. Defaults to `true`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccountState.__new__(_AccountState)

        __props__.__dict__["enable_default_standards"] = enable_default_standards
        return Account(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="enableDefaultStandards")
    def enable_default_standards(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable the security standards that Security Hub has designated as automatically enabled including: ` AWS Foundational Security Best Practices v1.0.0` and `CIS AWS Foundations Benchmark v1.2.0`. Defaults to `true`.
        """
        return pulumi.get(self, "enable_default_standards")

