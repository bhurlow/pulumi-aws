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

__all__ = ['ViewArgs', 'View']

@pulumi.input_type
class ViewArgs:
    def __init__(__self__, *,
                 default_view: Optional[pulumi.Input[bool]] = None,
                 filters: Optional[pulumi.Input['ViewFiltersArgs']] = None,
                 included_properties: Optional[pulumi.Input[Sequence[pulumi.Input['ViewIncludedPropertyArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a View resource.
        :param pulumi.Input[bool] default_view: Specifies whether the view is the [_default view_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-views-about.html#manage-views-about-default) for the AWS Region. Default: `false`.
        :param pulumi.Input['ViewFiltersArgs'] filters: Specifies which resources are included in the results of queries made using this view. See Filters below for more details.
        :param pulumi.Input[Sequence[pulumi.Input['ViewIncludedPropertyArgs']]] included_properties: Optional fields to be included in search results from this view. See Included Properties below for more details.
        :param pulumi.Input[str] name: The name of the view. The name must be no more than 64 characters long, and can include letters, digits, and the dash (-) character. The name must be unique within its AWS Region.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        if default_view is not None:
            pulumi.set(__self__, "default_view", default_view)
        if filters is not None:
            pulumi.set(__self__, "filters", filters)
        if included_properties is not None:
            pulumi.set(__self__, "included_properties", included_properties)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="defaultView")
    def default_view(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether the view is the [_default view_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-views-about.html#manage-views-about-default) for the AWS Region. Default: `false`.
        """
        return pulumi.get(self, "default_view")

    @default_view.setter
    def default_view(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default_view", value)

    @property
    @pulumi.getter
    def filters(self) -> Optional[pulumi.Input['ViewFiltersArgs']]:
        """
        Specifies which resources are included in the results of queries made using this view. See Filters below for more details.
        """
        return pulumi.get(self, "filters")

    @filters.setter
    def filters(self, value: Optional[pulumi.Input['ViewFiltersArgs']]):
        pulumi.set(self, "filters", value)

    @property
    @pulumi.getter(name="includedProperties")
    def included_properties(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ViewIncludedPropertyArgs']]]]:
        """
        Optional fields to be included in search results from this view. See Included Properties below for more details.
        """
        return pulumi.get(self, "included_properties")

    @included_properties.setter
    def included_properties(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ViewIncludedPropertyArgs']]]]):
        pulumi.set(self, "included_properties", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the view. The name must be no more than 64 characters long, and can include letters, digits, and the dash (-) character. The name must be unique within its AWS Region.
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
class _ViewState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 default_view: Optional[pulumi.Input[bool]] = None,
                 filters: Optional[pulumi.Input['ViewFiltersArgs']] = None,
                 included_properties: Optional[pulumi.Input[Sequence[pulumi.Input['ViewIncludedPropertyArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering View resources.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the Resource Explorer view.
        :param pulumi.Input[bool] default_view: Specifies whether the view is the [_default view_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-views-about.html#manage-views-about-default) for the AWS Region. Default: `false`.
        :param pulumi.Input['ViewFiltersArgs'] filters: Specifies which resources are included in the results of queries made using this view. See Filters below for more details.
        :param pulumi.Input[Sequence[pulumi.Input['ViewIncludedPropertyArgs']]] included_properties: Optional fields to be included in search results from this view. See Included Properties below for more details.
        :param pulumi.Input[str] name: The name of the view. The name must be no more than 64 characters long, and can include letters, digits, and the dash (-) character. The name must be unique within its AWS Region.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if default_view is not None:
            pulumi.set(__self__, "default_view", default_view)
        if filters is not None:
            pulumi.set(__self__, "filters", filters)
        if included_properties is not None:
            pulumi.set(__self__, "included_properties", included_properties)
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
        Amazon Resource Name (ARN) of the Resource Explorer view.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="defaultView")
    def default_view(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether the view is the [_default view_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-views-about.html#manage-views-about-default) for the AWS Region. Default: `false`.
        """
        return pulumi.get(self, "default_view")

    @default_view.setter
    def default_view(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default_view", value)

    @property
    @pulumi.getter
    def filters(self) -> Optional[pulumi.Input['ViewFiltersArgs']]:
        """
        Specifies which resources are included in the results of queries made using this view. See Filters below for more details.
        """
        return pulumi.get(self, "filters")

    @filters.setter
    def filters(self, value: Optional[pulumi.Input['ViewFiltersArgs']]):
        pulumi.set(self, "filters", value)

    @property
    @pulumi.getter(name="includedProperties")
    def included_properties(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ViewIncludedPropertyArgs']]]]:
        """
        Optional fields to be included in search results from this view. See Included Properties below for more details.
        """
        return pulumi.get(self, "included_properties")

    @included_properties.setter
    def included_properties(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ViewIncludedPropertyArgs']]]]):
        pulumi.set(self, "included_properties", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the view. The name must be no more than 64 characters long, and can include letters, digits, and the dash (-) character. The name must be unique within its AWS Region.
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


class View(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_view: Optional[pulumi.Input[bool]] = None,
                 filters: Optional[pulumi.Input[pulumi.InputType['ViewFiltersArgs']]] = None,
                 included_properties: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ViewIncludedPropertyArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to manage a Resource Explorer view.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_index = aws.resourceexplorer.Index("exampleIndex", type="LOCAL")
        example_view = aws.resourceexplorer.View("exampleView",
            filters=aws.resourceexplorer.ViewFiltersArgs(
                filter_string="resourcetype:ec2:instance",
            ),
            included_properties=[aws.resourceexplorer.ViewIncludedPropertyArgs(
                name="tags",
            )],
            opts=pulumi.ResourceOptions(depends_on=[example_index]))
        ```

        ## Import

        Using `pulumi import`, import Resource Explorer views using the `arn`. For example:

        ```sh
         $ pulumi import aws:resourceexplorer/view:View example arn:aws:resource-explorer-2:us-west-2:123456789012:view/exampleview/e0914f6c-6c27-4b47-b5d4-6b28381a2421
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] default_view: Specifies whether the view is the [_default view_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-views-about.html#manage-views-about-default) for the AWS Region. Default: `false`.
        :param pulumi.Input[pulumi.InputType['ViewFiltersArgs']] filters: Specifies which resources are included in the results of queries made using this view. See Filters below for more details.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ViewIncludedPropertyArgs']]]] included_properties: Optional fields to be included in search results from this view. See Included Properties below for more details.
        :param pulumi.Input[str] name: The name of the view. The name must be no more than 64 characters long, and can include letters, digits, and the dash (-) character. The name must be unique within its AWS Region.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ViewArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage a Resource Explorer view.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_index = aws.resourceexplorer.Index("exampleIndex", type="LOCAL")
        example_view = aws.resourceexplorer.View("exampleView",
            filters=aws.resourceexplorer.ViewFiltersArgs(
                filter_string="resourcetype:ec2:instance",
            ),
            included_properties=[aws.resourceexplorer.ViewIncludedPropertyArgs(
                name="tags",
            )],
            opts=pulumi.ResourceOptions(depends_on=[example_index]))
        ```

        ## Import

        Using `pulumi import`, import Resource Explorer views using the `arn`. For example:

        ```sh
         $ pulumi import aws:resourceexplorer/view:View example arn:aws:resource-explorer-2:us-west-2:123456789012:view/exampleview/e0914f6c-6c27-4b47-b5d4-6b28381a2421
        ```

        :param str resource_name: The name of the resource.
        :param ViewArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ViewArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_view: Optional[pulumi.Input[bool]] = None,
                 filters: Optional[pulumi.Input[pulumi.InputType['ViewFiltersArgs']]] = None,
                 included_properties: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ViewIncludedPropertyArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ViewArgs.__new__(ViewArgs)

            __props__.__dict__["default_view"] = default_view
            __props__.__dict__["filters"] = filters
            __props__.__dict__["included_properties"] = included_properties
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(View, __self__).__init__(
            'aws:resourceexplorer/view:View',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            default_view: Optional[pulumi.Input[bool]] = None,
            filters: Optional[pulumi.Input[pulumi.InputType['ViewFiltersArgs']]] = None,
            included_properties: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ViewIncludedPropertyArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'View':
        """
        Get an existing View resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the Resource Explorer view.
        :param pulumi.Input[bool] default_view: Specifies whether the view is the [_default view_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-views-about.html#manage-views-about-default) for the AWS Region. Default: `false`.
        :param pulumi.Input[pulumi.InputType['ViewFiltersArgs']] filters: Specifies which resources are included in the results of queries made using this view. See Filters below for more details.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ViewIncludedPropertyArgs']]]] included_properties: Optional fields to be included in search results from this view. See Included Properties below for more details.
        :param pulumi.Input[str] name: The name of the view. The name must be no more than 64 characters long, and can include letters, digits, and the dash (-) character. The name must be unique within its AWS Region.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ViewState.__new__(_ViewState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["default_view"] = default_view
        __props__.__dict__["filters"] = filters
        __props__.__dict__["included_properties"] = included_properties
        __props__.__dict__["name"] = name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return View(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the Resource Explorer view.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="defaultView")
    def default_view(self) -> pulumi.Output[bool]:
        """
        Specifies whether the view is the [_default view_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-views-about.html#manage-views-about-default) for the AWS Region. Default: `false`.
        """
        return pulumi.get(self, "default_view")

    @property
    @pulumi.getter
    def filters(self) -> pulumi.Output[Optional['outputs.ViewFilters']]:
        """
        Specifies which resources are included in the results of queries made using this view. See Filters below for more details.
        """
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter(name="includedProperties")
    def included_properties(self) -> pulumi.Output[Optional[Sequence['outputs.ViewIncludedProperty']]]:
        """
        Optional fields to be included in search results from this view. See Included Properties below for more details.
        """
        return pulumi.get(self, "included_properties")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the view. The name must be no more than 64 characters long, and can include letters, digits, and the dash (-) character. The name must be unique within its AWS Region.
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

