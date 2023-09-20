# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DedicatedIpPoolArgs', 'DedicatedIpPool']

@pulumi.input_type
class DedicatedIpPoolArgs:
    def __init__(__self__, *,
                 pool_name: pulumi.Input[str],
                 scaling_mode: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a DedicatedIpPool resource.
        :param pulumi.Input[str] pool_name: Name of the dedicated IP pool.
               
               The following arguments are optional:
        :param pulumi.Input[str] scaling_mode: IP pool scaling mode. Valid values: `STANDARD`, `MANAGED`. If omitted, the AWS API will default to a standard pool.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the pool. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "pool_name", pool_name)
        if scaling_mode is not None:
            pulumi.set(__self__, "scaling_mode", scaling_mode)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="poolName")
    def pool_name(self) -> pulumi.Input[str]:
        """
        Name of the dedicated IP pool.

        The following arguments are optional:
        """
        return pulumi.get(self, "pool_name")

    @pool_name.setter
    def pool_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "pool_name", value)

    @property
    @pulumi.getter(name="scalingMode")
    def scaling_mode(self) -> Optional[pulumi.Input[str]]:
        """
        IP pool scaling mode. Valid values: `STANDARD`, `MANAGED`. If omitted, the AWS API will default to a standard pool.
        """
        return pulumi.get(self, "scaling_mode")

    @scaling_mode.setter
    def scaling_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scaling_mode", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the pool. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _DedicatedIpPoolState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 pool_name: Optional[pulumi.Input[str]] = None,
                 scaling_mode: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering DedicatedIpPool resources.
        :param pulumi.Input[str] arn: ARN of the Dedicated IP Pool.
        :param pulumi.Input[str] pool_name: Name of the dedicated IP pool.
               
               The following arguments are optional:
        :param pulumi.Input[str] scaling_mode: IP pool scaling mode. Valid values: `STANDARD`, `MANAGED`. If omitted, the AWS API will default to a standard pool.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the pool. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if pool_name is not None:
            pulumi.set(__self__, "pool_name", pool_name)
        if scaling_mode is not None:
            pulumi.set(__self__, "scaling_mode", scaling_mode)
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
        ARN of the Dedicated IP Pool.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="poolName")
    def pool_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the dedicated IP pool.

        The following arguments are optional:
        """
        return pulumi.get(self, "pool_name")

    @pool_name.setter
    def pool_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pool_name", value)

    @property
    @pulumi.getter(name="scalingMode")
    def scaling_mode(self) -> Optional[pulumi.Input[str]]:
        """
        IP pool scaling mode. Valid values: `STANDARD`, `MANAGED`. If omitted, the AWS API will default to a standard pool.
        """
        return pulumi.get(self, "scaling_mode")

    @scaling_mode.setter
    def scaling_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scaling_mode", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the pool. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class DedicatedIpPool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 pool_name: Optional[pulumi.Input[str]] = None,
                 scaling_mode: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Resource for managing an AWS SESv2 (Simple Email V2) Dedicated IP Pool.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sesv2.DedicatedIpPool("example", pool_name="my-pool")
        ```
        ### Managed Pool

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sesv2.DedicatedIpPool("example",
            pool_name="my-managed-pool",
            scaling_mode="MANAGED")
        ```

        ## Import

        Using `pulumi import`, import SESv2 (Simple Email V2) Dedicated IP Pool using the `pool_name`. For example:

        ```sh
         $ pulumi import aws:sesv2/dedicatedIpPool:DedicatedIpPool example my-pool
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] pool_name: Name of the dedicated IP pool.
               
               The following arguments are optional:
        :param pulumi.Input[str] scaling_mode: IP pool scaling mode. Valid values: `STANDARD`, `MANAGED`. If omitted, the AWS API will default to a standard pool.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the pool. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DedicatedIpPoolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for managing an AWS SESv2 (Simple Email V2) Dedicated IP Pool.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sesv2.DedicatedIpPool("example", pool_name="my-pool")
        ```
        ### Managed Pool

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sesv2.DedicatedIpPool("example",
            pool_name="my-managed-pool",
            scaling_mode="MANAGED")
        ```

        ## Import

        Using `pulumi import`, import SESv2 (Simple Email V2) Dedicated IP Pool using the `pool_name`. For example:

        ```sh
         $ pulumi import aws:sesv2/dedicatedIpPool:DedicatedIpPool example my-pool
        ```

        :param str resource_name: The name of the resource.
        :param DedicatedIpPoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DedicatedIpPoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 pool_name: Optional[pulumi.Input[str]] = None,
                 scaling_mode: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DedicatedIpPoolArgs.__new__(DedicatedIpPoolArgs)

            if pool_name is None and not opts.urn:
                raise TypeError("Missing required property 'pool_name'")
            __props__.__dict__["pool_name"] = pool_name
            __props__.__dict__["scaling_mode"] = scaling_mode
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(DedicatedIpPool, __self__).__init__(
            'aws:sesv2/dedicatedIpPool:DedicatedIpPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            pool_name: Optional[pulumi.Input[str]] = None,
            scaling_mode: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'DedicatedIpPool':
        """
        Get an existing DedicatedIpPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: ARN of the Dedicated IP Pool.
        :param pulumi.Input[str] pool_name: Name of the dedicated IP pool.
               
               The following arguments are optional:
        :param pulumi.Input[str] scaling_mode: IP pool scaling mode. Valid values: `STANDARD`, `MANAGED`. If omitted, the AWS API will default to a standard pool.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the pool. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DedicatedIpPoolState.__new__(_DedicatedIpPoolState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["pool_name"] = pool_name
        __props__.__dict__["scaling_mode"] = scaling_mode
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return DedicatedIpPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN of the Dedicated IP Pool.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="poolName")
    def pool_name(self) -> pulumi.Output[str]:
        """
        Name of the dedicated IP pool.

        The following arguments are optional:
        """
        return pulumi.get(self, "pool_name")

    @property
    @pulumi.getter(name="scalingMode")
    def scaling_mode(self) -> pulumi.Output[str]:
        """
        IP pool scaling mode. Valid values: `STANDARD`, `MANAGED`. If omitted, the AWS API will default to a standard pool.
        """
        return pulumi.get(self, "scaling_mode")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the pool. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

