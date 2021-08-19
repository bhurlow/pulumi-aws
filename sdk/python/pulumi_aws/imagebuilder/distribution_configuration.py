# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DistributionConfigurationArgs', 'DistributionConfiguration']

@pulumi.input_type
class DistributionConfigurationArgs:
    def __init__(__self__, *,
                 distributions: pulumi.Input[Sequence[pulumi.Input['DistributionConfigurationDistributionArgs']]],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a DistributionConfiguration resource.
        :param pulumi.Input[Sequence[pulumi.Input['DistributionConfigurationDistributionArgs']]] distributions: One or more configuration blocks with distribution settings. Detailed below.
        :param pulumi.Input[str] description: Description to apply to the distributed AMI.
        :param pulumi.Input[str] name: Name to apply to the distributed AMI.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags for the distribution configuration. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "distributions", distributions)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def distributions(self) -> pulumi.Input[Sequence[pulumi.Input['DistributionConfigurationDistributionArgs']]]:
        """
        One or more configuration blocks with distribution settings. Detailed below.
        """
        return pulumi.get(self, "distributions")

    @distributions.setter
    def distributions(self, value: pulumi.Input[Sequence[pulumi.Input['DistributionConfigurationDistributionArgs']]]):
        pulumi.set(self, "distributions", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description to apply to the distributed AMI.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name to apply to the distributed AMI.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags for the distribution configuration. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _DistributionConfigurationState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 date_created: Optional[pulumi.Input[str]] = None,
                 date_updated: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 distributions: Optional[pulumi.Input[Sequence[pulumi.Input['DistributionConfigurationDistributionArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering DistributionConfiguration resources.
        :param pulumi.Input[str] arn: (Required) Amazon Resource Name (ARN) of the distribution configuration.
        :param pulumi.Input[str] date_created: Date the distribution configuration was created.
        :param pulumi.Input[str] date_updated: Date the distribution configuration was updated.
        :param pulumi.Input[str] description: Description to apply to the distributed AMI.
        :param pulumi.Input[Sequence[pulumi.Input['DistributionConfigurationDistributionArgs']]] distributions: One or more configuration blocks with distribution settings. Detailed below.
        :param pulumi.Input[str] name: Name to apply to the distributed AMI.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags for the distribution configuration. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if date_created is not None:
            pulumi.set(__self__, "date_created", date_created)
        if date_updated is not None:
            pulumi.set(__self__, "date_updated", date_updated)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if distributions is not None:
            pulumi.set(__self__, "distributions", distributions)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        (Required) Amazon Resource Name (ARN) of the distribution configuration.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> Optional[pulumi.Input[str]]:
        """
        Date the distribution configuration was created.
        """
        return pulumi.get(self, "date_created")

    @date_created.setter
    def date_created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "date_created", value)

    @property
    @pulumi.getter(name="dateUpdated")
    def date_updated(self) -> Optional[pulumi.Input[str]]:
        """
        Date the distribution configuration was updated.
        """
        return pulumi.get(self, "date_updated")

    @date_updated.setter
    def date_updated(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "date_updated", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description to apply to the distributed AMI.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def distributions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DistributionConfigurationDistributionArgs']]]]:
        """
        One or more configuration blocks with distribution settings. Detailed below.
        """
        return pulumi.get(self, "distributions")

    @distributions.setter
    def distributions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DistributionConfigurationDistributionArgs']]]]):
        pulumi.set(self, "distributions", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name to apply to the distributed AMI.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags for the distribution configuration. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class DistributionConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 distributions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DistributionConfigurationDistributionArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manages an Image Builder Distribution Configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.imagebuilder.DistributionConfiguration("example", distributions=[aws.imagebuilder.DistributionConfigurationDistributionArgs(
            ami_distribution_configuration=aws.imagebuilder.DistributionConfigurationDistributionAmiDistributionConfigurationArgs(
                ami_tags={
                    "CostCenter": "IT",
                },
                launch_permission=aws.imagebuilder.DistributionConfigurationDistributionAmiDistributionConfigurationLaunchPermissionArgs(
                    user_ids=["123456789012"],
                ),
                name="example-{{ imagebuilder:buildDate }}",
            ),
            region="us-east-1",
        )])
        ```

        ## Import

        `aws_imagebuilder_distribution_configurations` resources can be imported by using the Amazon Resource Name (ARN), e.g.

        ```sh
         $ pulumi import aws:imagebuilder/distributionConfiguration:DistributionConfiguration example arn:aws:imagebuilder:us-east-1:123456789012:distribution-configuration/example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description to apply to the distributed AMI.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DistributionConfigurationDistributionArgs']]]] distributions: One or more configuration blocks with distribution settings. Detailed below.
        :param pulumi.Input[str] name: Name to apply to the distributed AMI.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags for the distribution configuration. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DistributionConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Image Builder Distribution Configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.imagebuilder.DistributionConfiguration("example", distributions=[aws.imagebuilder.DistributionConfigurationDistributionArgs(
            ami_distribution_configuration=aws.imagebuilder.DistributionConfigurationDistributionAmiDistributionConfigurationArgs(
                ami_tags={
                    "CostCenter": "IT",
                },
                launch_permission=aws.imagebuilder.DistributionConfigurationDistributionAmiDistributionConfigurationLaunchPermissionArgs(
                    user_ids=["123456789012"],
                ),
                name="example-{{ imagebuilder:buildDate }}",
            ),
            region="us-east-1",
        )])
        ```

        ## Import

        `aws_imagebuilder_distribution_configurations` resources can be imported by using the Amazon Resource Name (ARN), e.g.

        ```sh
         $ pulumi import aws:imagebuilder/distributionConfiguration:DistributionConfiguration example arn:aws:imagebuilder:us-east-1:123456789012:distribution-configuration/example
        ```

        :param str resource_name: The name of the resource.
        :param DistributionConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DistributionConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 distributions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DistributionConfigurationDistributionArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DistributionConfigurationArgs.__new__(DistributionConfigurationArgs)

            __props__.__dict__["description"] = description
            if distributions is None and not opts.urn:
                raise TypeError("Missing required property 'distributions'")
            __props__.__dict__["distributions"] = distributions
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["date_created"] = None
            __props__.__dict__["date_updated"] = None
            __props__.__dict__["tags_all"] = None
        super(DistributionConfiguration, __self__).__init__(
            'aws:imagebuilder/distributionConfiguration:DistributionConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            date_created: Optional[pulumi.Input[str]] = None,
            date_updated: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            distributions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DistributionConfigurationDistributionArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'DistributionConfiguration':
        """
        Get an existing DistributionConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: (Required) Amazon Resource Name (ARN) of the distribution configuration.
        :param pulumi.Input[str] date_created: Date the distribution configuration was created.
        :param pulumi.Input[str] date_updated: Date the distribution configuration was updated.
        :param pulumi.Input[str] description: Description to apply to the distributed AMI.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DistributionConfigurationDistributionArgs']]]] distributions: One or more configuration blocks with distribution settings. Detailed below.
        :param pulumi.Input[str] name: Name to apply to the distributed AMI.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags for the distribution configuration. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DistributionConfigurationState.__new__(_DistributionConfigurationState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["date_created"] = date_created
        __props__.__dict__["date_updated"] = date_updated
        __props__.__dict__["description"] = description
        __props__.__dict__["distributions"] = distributions
        __props__.__dict__["name"] = name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return DistributionConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        (Required) Amazon Resource Name (ARN) of the distribution configuration.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> pulumi.Output[str]:
        """
        Date the distribution configuration was created.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter(name="dateUpdated")
    def date_updated(self) -> pulumi.Output[str]:
        """
        Date the distribution configuration was updated.
        """
        return pulumi.get(self, "date_updated")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description to apply to the distributed AMI.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def distributions(self) -> pulumi.Output[Sequence['outputs.DistributionConfigurationDistribution']]:
        """
        One or more configuration blocks with distribution settings. Detailed below.
        """
        return pulumi.get(self, "distributions")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name to apply to the distributed AMI.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags for the distribution configuration. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

