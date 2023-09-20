# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ModelPackageGroupArgs', 'ModelPackageGroup']

@pulumi.input_type
class ModelPackageGroupArgs:
    def __init__(__self__, *,
                 model_package_group_name: pulumi.Input[str],
                 model_package_group_description: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ModelPackageGroup resource.
        :param pulumi.Input[str] model_package_group_name: The name of the model group.
        :param pulumi.Input[str] model_package_group_description: A description for the model group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "model_package_group_name", model_package_group_name)
        if model_package_group_description is not None:
            pulumi.set(__self__, "model_package_group_description", model_package_group_description)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="modelPackageGroupName")
    def model_package_group_name(self) -> pulumi.Input[str]:
        """
        The name of the model group.
        """
        return pulumi.get(self, "model_package_group_name")

    @model_package_group_name.setter
    def model_package_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "model_package_group_name", value)

    @property
    @pulumi.getter(name="modelPackageGroupDescription")
    def model_package_group_description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the model group.
        """
        return pulumi.get(self, "model_package_group_description")

    @model_package_group_description.setter
    def model_package_group_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "model_package_group_description", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ModelPackageGroupState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 model_package_group_description: Optional[pulumi.Input[str]] = None,
                 model_package_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ModelPackageGroup resources.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) assigned by AWS to this Model Package Group.
        :param pulumi.Input[str] model_package_group_description: A description for the model group.
        :param pulumi.Input[str] model_package_group_name: The name of the model group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if model_package_group_description is not None:
            pulumi.set(__self__, "model_package_group_description", model_package_group_description)
        if model_package_group_name is not None:
            pulumi.set(__self__, "model_package_group_name", model_package_group_name)
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
        The Amazon Resource Name (ARN) assigned by AWS to this Model Package Group.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="modelPackageGroupDescription")
    def model_package_group_description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the model group.
        """
        return pulumi.get(self, "model_package_group_description")

    @model_package_group_description.setter
    def model_package_group_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "model_package_group_description", value)

    @property
    @pulumi.getter(name="modelPackageGroupName")
    def model_package_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the model group.
        """
        return pulumi.get(self, "model_package_group_name")

    @model_package_group_name.setter
    def model_package_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "model_package_group_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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


class ModelPackageGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 model_package_group_description: Optional[pulumi.Input[str]] = None,
                 model_package_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a SageMaker Model Package Group resource.

        ## Example Usage
        ### Basic usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sagemaker.ModelPackageGroup("example", model_package_group_name="example")
        ```

        ## Import

        Using `pulumi import`, import SageMaker Model Package Groups using the `name`. For example:

        ```sh
         $ pulumi import aws:sagemaker/modelPackageGroup:ModelPackageGroup test_model_package_group my-code-repo
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] model_package_group_description: A description for the model group.
        :param pulumi.Input[str] model_package_group_name: The name of the model group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ModelPackageGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a SageMaker Model Package Group resource.

        ## Example Usage
        ### Basic usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sagemaker.ModelPackageGroup("example", model_package_group_name="example")
        ```

        ## Import

        Using `pulumi import`, import SageMaker Model Package Groups using the `name`. For example:

        ```sh
         $ pulumi import aws:sagemaker/modelPackageGroup:ModelPackageGroup test_model_package_group my-code-repo
        ```

        :param str resource_name: The name of the resource.
        :param ModelPackageGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ModelPackageGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 model_package_group_description: Optional[pulumi.Input[str]] = None,
                 model_package_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ModelPackageGroupArgs.__new__(ModelPackageGroupArgs)

            __props__.__dict__["model_package_group_description"] = model_package_group_description
            if model_package_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'model_package_group_name'")
            __props__.__dict__["model_package_group_name"] = model_package_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ModelPackageGroup, __self__).__init__(
            'aws:sagemaker/modelPackageGroup:ModelPackageGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            model_package_group_description: Optional[pulumi.Input[str]] = None,
            model_package_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'ModelPackageGroup':
        """
        Get an existing ModelPackageGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) assigned by AWS to this Model Package Group.
        :param pulumi.Input[str] model_package_group_description: A description for the model group.
        :param pulumi.Input[str] model_package_group_name: The name of the model group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ModelPackageGroupState.__new__(_ModelPackageGroupState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["model_package_group_description"] = model_package_group_description
        __props__.__dict__["model_package_group_name"] = model_package_group_name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return ModelPackageGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) assigned by AWS to this Model Package Group.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="modelPackageGroupDescription")
    def model_package_group_description(self) -> pulumi.Output[Optional[str]]:
        """
        A description for the model group.
        """
        return pulumi.get(self, "model_package_group_description")

    @property
    @pulumi.getter(name="modelPackageGroupName")
    def model_package_group_name(self) -> pulumi.Output[str]:
        """
        The name of the model group.
        """
        return pulumi.get(self, "model_package_group_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

