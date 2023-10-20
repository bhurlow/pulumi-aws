# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LicenseConfigurationArgs', 'LicenseConfiguration']

@pulumi.input_type
class LicenseConfigurationArgs:
    def __init__(__self__, *,
                 license_counting_type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 license_count: Optional[pulumi.Input[int]] = None,
                 license_count_hard_limit: Optional[pulumi.Input[bool]] = None,
                 license_rules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a LicenseConfiguration resource.
        :param pulumi.Input[str] license_counting_type: Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
        :param pulumi.Input[str] description: Description of the license configuration.
        :param pulumi.Input[int] license_count: Number of licenses managed by the license configuration.
        :param pulumi.Input[bool] license_count_hard_limit: Sets the number of available licenses as a hard limit.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] license_rules: Array of configured License Manager rules.
        :param pulumi.Input[str] name: Name of the license configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "license_counting_type", license_counting_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if license_count is not None:
            pulumi.set(__self__, "license_count", license_count)
        if license_count_hard_limit is not None:
            pulumi.set(__self__, "license_count_hard_limit", license_count_hard_limit)
        if license_rules is not None:
            pulumi.set(__self__, "license_rules", license_rules)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="licenseCountingType")
    def license_counting_type(self) -> pulumi.Input[str]:
        """
        Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
        """
        return pulumi.get(self, "license_counting_type")

    @license_counting_type.setter
    def license_counting_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "license_counting_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the license configuration.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="licenseCount")
    def license_count(self) -> Optional[pulumi.Input[int]]:
        """
        Number of licenses managed by the license configuration.
        """
        return pulumi.get(self, "license_count")

    @license_count.setter
    def license_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "license_count", value)

    @property
    @pulumi.getter(name="licenseCountHardLimit")
    def license_count_hard_limit(self) -> Optional[pulumi.Input[bool]]:
        """
        Sets the number of available licenses as a hard limit.
        """
        return pulumi.get(self, "license_count_hard_limit")

    @license_count_hard_limit.setter
    def license_count_hard_limit(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "license_count_hard_limit", value)

    @property
    @pulumi.getter(name="licenseRules")
    def license_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Array of configured License Manager rules.
        """
        return pulumi.get(self, "license_rules")

    @license_rules.setter
    def license_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "license_rules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the license configuration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _LicenseConfigurationState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 license_count: Optional[pulumi.Input[int]] = None,
                 license_count_hard_limit: Optional[pulumi.Input[bool]] = None,
                 license_counting_type: Optional[pulumi.Input[str]] = None,
                 license_rules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_account_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering LicenseConfiguration resources.
        :param pulumi.Input[str] arn: The license configuration ARN.
        :param pulumi.Input[str] description: Description of the license configuration.
        :param pulumi.Input[int] license_count: Number of licenses managed by the license configuration.
        :param pulumi.Input[bool] license_count_hard_limit: Sets the number of available licenses as a hard limit.
        :param pulumi.Input[str] license_counting_type: Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] license_rules: Array of configured License Manager rules.
        :param pulumi.Input[str] name: Name of the license configuration.
        :param pulumi.Input[str] owner_account_id: Account ID of the owner of the license configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if license_count is not None:
            pulumi.set(__self__, "license_count", license_count)
        if license_count_hard_limit is not None:
            pulumi.set(__self__, "license_count_hard_limit", license_count_hard_limit)
        if license_counting_type is not None:
            pulumi.set(__self__, "license_counting_type", license_counting_type)
        if license_rules is not None:
            pulumi.set(__self__, "license_rules", license_rules)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if owner_account_id is not None:
            pulumi.set(__self__, "owner_account_id", owner_account_id)
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
        The license configuration ARN.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the license configuration.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="licenseCount")
    def license_count(self) -> Optional[pulumi.Input[int]]:
        """
        Number of licenses managed by the license configuration.
        """
        return pulumi.get(self, "license_count")

    @license_count.setter
    def license_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "license_count", value)

    @property
    @pulumi.getter(name="licenseCountHardLimit")
    def license_count_hard_limit(self) -> Optional[pulumi.Input[bool]]:
        """
        Sets the number of available licenses as a hard limit.
        """
        return pulumi.get(self, "license_count_hard_limit")

    @license_count_hard_limit.setter
    def license_count_hard_limit(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "license_count_hard_limit", value)

    @property
    @pulumi.getter(name="licenseCountingType")
    def license_counting_type(self) -> Optional[pulumi.Input[str]]:
        """
        Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
        """
        return pulumi.get(self, "license_counting_type")

    @license_counting_type.setter
    def license_counting_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "license_counting_type", value)

    @property
    @pulumi.getter(name="licenseRules")
    def license_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Array of configured License Manager rules.
        """
        return pulumi.get(self, "license_rules")

    @license_rules.setter
    def license_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "license_rules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the license configuration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ownerAccountId")
    def owner_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        Account ID of the owner of the license configuration.
        """
        return pulumi.get(self, "owner_account_id")

    @owner_account_id.setter
    def owner_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_account_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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


class LicenseConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 license_count: Optional[pulumi.Input[int]] = None,
                 license_count_hard_limit: Optional[pulumi.Input[bool]] = None,
                 license_counting_type: Optional[pulumi.Input[str]] = None,
                 license_rules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a License Manager license configuration resource.

        > **Note:** Removing the `license_count` attribute is not supported by the License Manager API.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.licensemanager.LicenseConfiguration("example",
            description="Example",
            license_count=10,
            license_count_hard_limit=True,
            license_counting_type="Socket",
            license_rules=["#minimumSockets=2"],
            tags={
                "foo": "barr",
            })
        ```
        ## Rules

        License rules should be in the format of `#RuleType=RuleValue`. Supported rule types:

        * `minimumVcpus` - Resource must have minimum vCPU count in order to use the license. Default: 1
        * `maximumVcpus` - Resource must have maximum vCPU count in order to use the license. Default: unbounded, limit: 10000
        * `minimumCores` - Resource must have minimum core count in order to use the license. Default: 1
        * `maximumCores` - Resource must have maximum core count in order to use the license. Default: unbounded, limit: 10000
        * `minimumSockets` - Resource must have minimum socket count in order to use the license. Default: 1
        * `maximumSockets` - Resource must have maximum socket count in order to use the license. Default: unbounded, limit: 10000
        * `allowedTenancy` - Defines where the license can be used. If set, restricts license usage to selected tenancies. Specify a comma delimited list of `EC2-Default`, `EC2-DedicatedHost`, `EC2-DedicatedInstance`

        ## Import

        Using `pulumi import`, import license configurations using the `id`. For example:

        ```sh
         $ pulumi import aws:licensemanager/licenseConfiguration:LicenseConfiguration example arn:aws:license-manager:eu-west-1:123456789012:license-configuration:lic-0123456789abcdef0123456789abcdef
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the license configuration.
        :param pulumi.Input[int] license_count: Number of licenses managed by the license configuration.
        :param pulumi.Input[bool] license_count_hard_limit: Sets the number of available licenses as a hard limit.
        :param pulumi.Input[str] license_counting_type: Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] license_rules: Array of configured License Manager rules.
        :param pulumi.Input[str] name: Name of the license configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LicenseConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a License Manager license configuration resource.

        > **Note:** Removing the `license_count` attribute is not supported by the License Manager API.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.licensemanager.LicenseConfiguration("example",
            description="Example",
            license_count=10,
            license_count_hard_limit=True,
            license_counting_type="Socket",
            license_rules=["#minimumSockets=2"],
            tags={
                "foo": "barr",
            })
        ```
        ## Rules

        License rules should be in the format of `#RuleType=RuleValue`. Supported rule types:

        * `minimumVcpus` - Resource must have minimum vCPU count in order to use the license. Default: 1
        * `maximumVcpus` - Resource must have maximum vCPU count in order to use the license. Default: unbounded, limit: 10000
        * `minimumCores` - Resource must have minimum core count in order to use the license. Default: 1
        * `maximumCores` - Resource must have maximum core count in order to use the license. Default: unbounded, limit: 10000
        * `minimumSockets` - Resource must have minimum socket count in order to use the license. Default: 1
        * `maximumSockets` - Resource must have maximum socket count in order to use the license. Default: unbounded, limit: 10000
        * `allowedTenancy` - Defines where the license can be used. If set, restricts license usage to selected tenancies. Specify a comma delimited list of `EC2-Default`, `EC2-DedicatedHost`, `EC2-DedicatedInstance`

        ## Import

        Using `pulumi import`, import license configurations using the `id`. For example:

        ```sh
         $ pulumi import aws:licensemanager/licenseConfiguration:LicenseConfiguration example arn:aws:license-manager:eu-west-1:123456789012:license-configuration:lic-0123456789abcdef0123456789abcdef
        ```

        :param str resource_name: The name of the resource.
        :param LicenseConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LicenseConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 license_count: Optional[pulumi.Input[int]] = None,
                 license_count_hard_limit: Optional[pulumi.Input[bool]] = None,
                 license_counting_type: Optional[pulumi.Input[str]] = None,
                 license_rules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LicenseConfigurationArgs.__new__(LicenseConfigurationArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["license_count"] = license_count
            __props__.__dict__["license_count_hard_limit"] = license_count_hard_limit
            if license_counting_type is None and not opts.urn:
                raise TypeError("Missing required property 'license_counting_type'")
            __props__.__dict__["license_counting_type"] = license_counting_type
            __props__.__dict__["license_rules"] = license_rules
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["owner_account_id"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(LicenseConfiguration, __self__).__init__(
            'aws:licensemanager/licenseConfiguration:LicenseConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            license_count: Optional[pulumi.Input[int]] = None,
            license_count_hard_limit: Optional[pulumi.Input[bool]] = None,
            license_counting_type: Optional[pulumi.Input[str]] = None,
            license_rules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            owner_account_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'LicenseConfiguration':
        """
        Get an existing LicenseConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The license configuration ARN.
        :param pulumi.Input[str] description: Description of the license configuration.
        :param pulumi.Input[int] license_count: Number of licenses managed by the license configuration.
        :param pulumi.Input[bool] license_count_hard_limit: Sets the number of available licenses as a hard limit.
        :param pulumi.Input[str] license_counting_type: Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] license_rules: Array of configured License Manager rules.
        :param pulumi.Input[str] name: Name of the license configuration.
        :param pulumi.Input[str] owner_account_id: Account ID of the owner of the license configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LicenseConfigurationState.__new__(_LicenseConfigurationState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["description"] = description
        __props__.__dict__["license_count"] = license_count
        __props__.__dict__["license_count_hard_limit"] = license_count_hard_limit
        __props__.__dict__["license_counting_type"] = license_counting_type
        __props__.__dict__["license_rules"] = license_rules
        __props__.__dict__["name"] = name
        __props__.__dict__["owner_account_id"] = owner_account_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return LicenseConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The license configuration ARN.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the license configuration.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="licenseCount")
    def license_count(self) -> pulumi.Output[Optional[int]]:
        """
        Number of licenses managed by the license configuration.
        """
        return pulumi.get(self, "license_count")

    @property
    @pulumi.getter(name="licenseCountHardLimit")
    def license_count_hard_limit(self) -> pulumi.Output[Optional[bool]]:
        """
        Sets the number of available licenses as a hard limit.
        """
        return pulumi.get(self, "license_count_hard_limit")

    @property
    @pulumi.getter(name="licenseCountingType")
    def license_counting_type(self) -> pulumi.Output[str]:
        """
        Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
        """
        return pulumi.get(self, "license_counting_type")

    @property
    @pulumi.getter(name="licenseRules")
    def license_rules(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Array of configured License Manager rules.
        """
        return pulumi.get(self, "license_rules")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the license configuration.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerAccountId")
    def owner_account_id(self) -> pulumi.Output[str]:
        """
        Account ID of the owner of the license configuration.
        """
        return pulumi.get(self, "owner_account_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

