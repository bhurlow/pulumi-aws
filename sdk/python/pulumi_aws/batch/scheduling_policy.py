# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SchedulingPolicyArgs', 'SchedulingPolicy']

@pulumi.input_type
class SchedulingPolicyArgs:
    def __init__(__self__, *,
                 fair_share_policy: Optional[pulumi.Input['SchedulingPolicyFairSharePolicyArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SchedulingPolicy resource.
        :param pulumi.Input[str] name: Specifies the name of the scheduling policy.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        if fair_share_policy is not None:
            pulumi.set(__self__, "fair_share_policy", fair_share_policy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="fairSharePolicy")
    def fair_share_policy(self) -> Optional[pulumi.Input['SchedulingPolicyFairSharePolicyArgs']]:
        return pulumi.get(self, "fair_share_policy")

    @fair_share_policy.setter
    def fair_share_policy(self, value: Optional[pulumi.Input['SchedulingPolicyFairSharePolicyArgs']]):
        pulumi.set(self, "fair_share_policy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the scheduling policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _SchedulingPolicyState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 fair_share_policy: Optional[pulumi.Input['SchedulingPolicyFairSharePolicyArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering SchedulingPolicy resources.
        :param pulumi.Input[str] arn: The Amazon Resource Name of the scheduling policy.
        :param pulumi.Input[str] name: Specifies the name of the scheduling policy.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if fair_share_policy is not None:
            pulumi.set(__self__, "fair_share_policy", fair_share_policy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
            pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name of the scheduling policy.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="fairSharePolicy")
    def fair_share_policy(self) -> Optional[pulumi.Input['SchedulingPolicyFairSharePolicyArgs']]:
        return pulumi.get(self, "fair_share_policy")

    @fair_share_policy.setter
    def fair_share_policy(self, value: Optional[pulumi.Input['SchedulingPolicyFairSharePolicyArgs']]):
        pulumi.set(self, "fair_share_policy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the scheduling policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class SchedulingPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fair_share_policy: Optional[pulumi.Input[pulumi.InputType['SchedulingPolicyFairSharePolicyArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a Batch Scheduling Policy resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.batch.SchedulingPolicy("example",
            fair_share_policy=aws.batch.SchedulingPolicyFairSharePolicyArgs(
                compute_reservation=1,
                share_decay_seconds=3600,
                share_distributions=[
                    aws.batch.SchedulingPolicyFairSharePolicyShareDistributionArgs(
                        share_identifier="A1*",
                        weight_factor=0.1,
                    ),
                    aws.batch.SchedulingPolicyFairSharePolicyShareDistributionArgs(
                        share_identifier="A2",
                        weight_factor=0.2,
                    ),
                ],
            ),
            tags={
                "Name": "Example Batch Scheduling Policy",
            })
        ```

        ## Import

        Using `pulumi import`, import Batch Scheduling Policy using the `arn`. For example:

        ```sh
         $ pulumi import aws:batch/schedulingPolicy:SchedulingPolicy test_policy arn:aws:batch:us-east-1:123456789012:scheduling-policy/sample
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Specifies the name of the scheduling policy.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SchedulingPolicyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Batch Scheduling Policy resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.batch.SchedulingPolicy("example",
            fair_share_policy=aws.batch.SchedulingPolicyFairSharePolicyArgs(
                compute_reservation=1,
                share_decay_seconds=3600,
                share_distributions=[
                    aws.batch.SchedulingPolicyFairSharePolicyShareDistributionArgs(
                        share_identifier="A1*",
                        weight_factor=0.1,
                    ),
                    aws.batch.SchedulingPolicyFairSharePolicyShareDistributionArgs(
                        share_identifier="A2",
                        weight_factor=0.2,
                    ),
                ],
            ),
            tags={
                "Name": "Example Batch Scheduling Policy",
            })
        ```

        ## Import

        Using `pulumi import`, import Batch Scheduling Policy using the `arn`. For example:

        ```sh
         $ pulumi import aws:batch/schedulingPolicy:SchedulingPolicy test_policy arn:aws:batch:us-east-1:123456789012:scheduling-policy/sample
        ```

        :param str resource_name: The name of the resource.
        :param SchedulingPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SchedulingPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fair_share_policy: Optional[pulumi.Input[pulumi.InputType['SchedulingPolicyFairSharePolicyArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SchedulingPolicyArgs.__new__(SchedulingPolicyArgs)

            __props__.__dict__["fair_share_policy"] = fair_share_policy
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SchedulingPolicy, __self__).__init__(
            'aws:batch/schedulingPolicy:SchedulingPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            fair_share_policy: Optional[pulumi.Input[pulumi.InputType['SchedulingPolicyFairSharePolicyArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'SchedulingPolicy':
        """
        Get an existing SchedulingPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name of the scheduling policy.
        :param pulumi.Input[str] name: Specifies the name of the scheduling policy.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SchedulingPolicyState.__new__(_SchedulingPolicyState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["fair_share_policy"] = fair_share_policy
        __props__.__dict__["name"] = name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return SchedulingPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name of the scheduling policy.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="fairSharePolicy")
    def fair_share_policy(self) -> pulumi.Output[Optional['outputs.SchedulingPolicyFairSharePolicy']]:
        return pulumi.get(self, "fair_share_policy")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the scheduling policy.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

