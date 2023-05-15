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

__all__ = ['TemplateArgs', 'Template']

@pulumi.input_type
class TemplateArgs:
    def __init__(__self__, *,
                 template_id: pulumi.Input[str],
                 version_description: pulumi.Input[str],
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input['TemplatePermissionArgs']]]] = None,
                 source_entity: Optional[pulumi.Input['TemplateSourceEntityArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Template resource.
        :param pulumi.Input[str] template_id: Identifier for the template.
        :param pulumi.Input[str] version_description: A description of the current template version being created/updated.
        :param pulumi.Input[str] aws_account_id: AWS account ID.
        :param pulumi.Input[str] name: Display name for the template.
        :param pulumi.Input[Sequence[pulumi.Input['TemplatePermissionArgs']]] permissions: A set of resource permissions on the template. Maximum of 64 items. See permissions.
        :param pulumi.Input['TemplateSourceEntityArgs'] source_entity: The entity that you are using as a source when you create the template (analysis or template). Only one of `definition` or `source_entity` should be configured. See source_entity.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "template_id", template_id)
        pulumi.set(__self__, "version_description", version_description)
        if aws_account_id is not None:
            pulumi.set(__self__, "aws_account_id", aws_account_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if source_entity is not None:
            pulumi.set(__self__, "source_entity", source_entity)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Input[str]:
        """
        Identifier for the template.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "template_id", value)

    @property
    @pulumi.getter(name="versionDescription")
    def version_description(self) -> pulumi.Input[str]:
        """
        A description of the current template version being created/updated.
        """
        return pulumi.get(self, "version_description")

    @version_description.setter
    def version_description(self, value: pulumi.Input[str]):
        pulumi.set(self, "version_description", value)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        AWS account ID.
        """
        return pulumi.get(self, "aws_account_id")

    @aws_account_id.setter
    def aws_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_account_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name for the template.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TemplatePermissionArgs']]]]:
        """
        A set of resource permissions on the template. Maximum of 64 items. See permissions.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TemplatePermissionArgs']]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="sourceEntity")
    def source_entity(self) -> Optional[pulumi.Input['TemplateSourceEntityArgs']]:
        """
        The entity that you are using as a source when you create the template (analysis or template). Only one of `definition` or `source_entity` should be configured. See source_entity.
        """
        return pulumi.get(self, "source_entity")

    @source_entity.setter
    def source_entity(self, value: Optional[pulumi.Input['TemplateSourceEntityArgs']]):
        pulumi.set(self, "source_entity", value)

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
class _TemplateState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 created_time: Optional[pulumi.Input[str]] = None,
                 last_updated_time: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input['TemplatePermissionArgs']]]] = None,
                 source_entity: Optional[pulumi.Input['TemplateSourceEntityArgs']] = None,
                 source_entity_arn: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 version_description: Optional[pulumi.Input[str]] = None,
                 version_number: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Template resources.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the resource.
        :param pulumi.Input[str] aws_account_id: AWS account ID.
        :param pulumi.Input[str] created_time: The time that the template was created.
        :param pulumi.Input[str] last_updated_time: The time that the template was last updated.
        :param pulumi.Input[str] name: Display name for the template.
        :param pulumi.Input[Sequence[pulumi.Input['TemplatePermissionArgs']]] permissions: A set of resource permissions on the template. Maximum of 64 items. See permissions.
        :param pulumi.Input['TemplateSourceEntityArgs'] source_entity: The entity that you are using as a source when you create the template (analysis or template). Only one of `definition` or `source_entity` should be configured. See source_entity.
        :param pulumi.Input[str] source_entity_arn: Amazon Resource Name (ARN) of an analysis or template that was used to create this template.
        :param pulumi.Input[str] status: The template creation status.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[str] template_id: Identifier for the template.
        :param pulumi.Input[str] version_description: A description of the current template version being created/updated.
        :param pulumi.Input[int] version_number: The version number of the template version.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if aws_account_id is not None:
            pulumi.set(__self__, "aws_account_id", aws_account_id)
        if created_time is not None:
            pulumi.set(__self__, "created_time", created_time)
        if last_updated_time is not None:
            pulumi.set(__self__, "last_updated_time", last_updated_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if source_entity is not None:
            pulumi.set(__self__, "source_entity", source_entity)
        if source_entity_arn is not None:
            pulumi.set(__self__, "source_entity_arn", source_entity_arn)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if template_id is not None:
            pulumi.set(__self__, "template_id", template_id)
        if version_description is not None:
            pulumi.set(__self__, "version_description", version_description)
        if version_number is not None:
            pulumi.set(__self__, "version_number", version_number)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the resource.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        AWS account ID.
        """
        return pulumi.get(self, "aws_account_id")

    @aws_account_id.setter
    def aws_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_account_id", value)

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time that the template was created.
        """
        return pulumi.get(self, "created_time")

    @created_time.setter
    def created_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_time", value)

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time that the template was last updated.
        """
        return pulumi.get(self, "last_updated_time")

    @last_updated_time.setter
    def last_updated_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_updated_time", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name for the template.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TemplatePermissionArgs']]]]:
        """
        A set of resource permissions on the template. Maximum of 64 items. See permissions.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TemplatePermissionArgs']]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="sourceEntity")
    def source_entity(self) -> Optional[pulumi.Input['TemplateSourceEntityArgs']]:
        """
        The entity that you are using as a source when you create the template (analysis or template). Only one of `definition` or `source_entity` should be configured. See source_entity.
        """
        return pulumi.get(self, "source_entity")

    @source_entity.setter
    def source_entity(self, value: Optional[pulumi.Input['TemplateSourceEntityArgs']]):
        pulumi.set(self, "source_entity", value)

    @property
    @pulumi.getter(name="sourceEntityArn")
    def source_entity_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of an analysis or template that was used to create this template.
        """
        return pulumi.get(self, "source_entity_arn")

    @source_entity_arn.setter
    def source_entity_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_entity_arn", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The template creation status.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

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
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier for the template.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_id", value)

    @property
    @pulumi.getter(name="versionDescription")
    def version_description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the current template version being created/updated.
        """
        return pulumi.get(self, "version_description")

    @version_description.setter
    def version_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version_description", value)

    @property
    @pulumi.getter(name="versionNumber")
    def version_number(self) -> Optional[pulumi.Input[int]]:
        """
        The version number of the template version.
        """
        return pulumi.get(self, "version_number")

    @version_number.setter
    def version_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version_number", value)


class Template(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TemplatePermissionArgs']]]]] = None,
                 source_entity: Optional[pulumi.Input[pulumi.InputType['TemplateSourceEntityArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 version_description: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource for managing a QuickSight Template.

        ## Example Usage
        ### From Source Template

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.quicksight.Template("example",
            template_id="example-id",
            version_description="version",
            source_entity=aws.quicksight.TemplateSourceEntityArgs(
                source_template=aws.quicksight.TemplateSourceEntitySourceTemplateArgs(
                    arn=aws_quicksight_template["source"]["arn"],
                ),
            ))
        ```

        ## Import

        A QuickSight Template can be imported using the AWS account ID and template ID separated by a comma (`,`) e.g.,

        ```sh
         $ pulumi import aws:quicksight/template:Template example 123456789012,example-id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_account_id: AWS account ID.
        :param pulumi.Input[str] name: Display name for the template.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TemplatePermissionArgs']]]] permissions: A set of resource permissions on the template. Maximum of 64 items. See permissions.
        :param pulumi.Input[pulumi.InputType['TemplateSourceEntityArgs']] source_entity: The entity that you are using as a source when you create the template (analysis or template). Only one of `definition` or `source_entity` should be configured. See source_entity.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] template_id: Identifier for the template.
        :param pulumi.Input[str] version_description: A description of the current template version being created/updated.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for managing a QuickSight Template.

        ## Example Usage
        ### From Source Template

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.quicksight.Template("example",
            template_id="example-id",
            version_description="version",
            source_entity=aws.quicksight.TemplateSourceEntityArgs(
                source_template=aws.quicksight.TemplateSourceEntitySourceTemplateArgs(
                    arn=aws_quicksight_template["source"]["arn"],
                ),
            ))
        ```

        ## Import

        A QuickSight Template can be imported using the AWS account ID and template ID separated by a comma (`,`) e.g.,

        ```sh
         $ pulumi import aws:quicksight/template:Template example 123456789012,example-id
        ```

        :param str resource_name: The name of the resource.
        :param TemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TemplatePermissionArgs']]]]] = None,
                 source_entity: Optional[pulumi.Input[pulumi.InputType['TemplateSourceEntityArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 version_description: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TemplateArgs.__new__(TemplateArgs)

            __props__.__dict__["aws_account_id"] = aws_account_id
            __props__.__dict__["name"] = name
            __props__.__dict__["permissions"] = permissions
            __props__.__dict__["source_entity"] = source_entity
            __props__.__dict__["tags"] = tags
            if template_id is None and not opts.urn:
                raise TypeError("Missing required property 'template_id'")
            __props__.__dict__["template_id"] = template_id
            if version_description is None and not opts.urn:
                raise TypeError("Missing required property 'version_description'")
            __props__.__dict__["version_description"] = version_description
            __props__.__dict__["arn"] = None
            __props__.__dict__["created_time"] = None
            __props__.__dict__["last_updated_time"] = None
            __props__.__dict__["source_entity_arn"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["tags_all"] = None
            __props__.__dict__["version_number"] = None
        super(Template, __self__).__init__(
            'aws:quicksight/template:Template',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            aws_account_id: Optional[pulumi.Input[str]] = None,
            created_time: Optional[pulumi.Input[str]] = None,
            last_updated_time: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TemplatePermissionArgs']]]]] = None,
            source_entity: Optional[pulumi.Input[pulumi.InputType['TemplateSourceEntityArgs']]] = None,
            source_entity_arn: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            template_id: Optional[pulumi.Input[str]] = None,
            version_description: Optional[pulumi.Input[str]] = None,
            version_number: Optional[pulumi.Input[int]] = None) -> 'Template':
        """
        Get an existing Template resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the resource.
        :param pulumi.Input[str] aws_account_id: AWS account ID.
        :param pulumi.Input[str] created_time: The time that the template was created.
        :param pulumi.Input[str] last_updated_time: The time that the template was last updated.
        :param pulumi.Input[str] name: Display name for the template.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TemplatePermissionArgs']]]] permissions: A set of resource permissions on the template. Maximum of 64 items. See permissions.
        :param pulumi.Input[pulumi.InputType['TemplateSourceEntityArgs']] source_entity: The entity that you are using as a source when you create the template (analysis or template). Only one of `definition` or `source_entity` should be configured. See source_entity.
        :param pulumi.Input[str] source_entity_arn: Amazon Resource Name (ARN) of an analysis or template that was used to create this template.
        :param pulumi.Input[str] status: The template creation status.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[str] template_id: Identifier for the template.
        :param pulumi.Input[str] version_description: A description of the current template version being created/updated.
        :param pulumi.Input[int] version_number: The version number of the template version.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TemplateState.__new__(_TemplateState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["aws_account_id"] = aws_account_id
        __props__.__dict__["created_time"] = created_time
        __props__.__dict__["last_updated_time"] = last_updated_time
        __props__.__dict__["name"] = name
        __props__.__dict__["permissions"] = permissions
        __props__.__dict__["source_entity"] = source_entity
        __props__.__dict__["source_entity_arn"] = source_entity_arn
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["template_id"] = template_id
        __props__.__dict__["version_description"] = version_description
        __props__.__dict__["version_number"] = version_number
        return Template(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the resource.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> pulumi.Output[str]:
        """
        AWS account ID.
        """
        return pulumi.get(self, "aws_account_id")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> pulumi.Output[str]:
        """
        The time that the template was created.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> pulumi.Output[str]:
        """
        The time that the template was last updated.
        """
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Display name for the template.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Optional[Sequence['outputs.TemplatePermission']]]:
        """
        A set of resource permissions on the template. Maximum of 64 items. See permissions.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="sourceEntity")
    def source_entity(self) -> pulumi.Output[Optional['outputs.TemplateSourceEntity']]:
        """
        The entity that you are using as a source when you create the template (analysis or template). Only one of `definition` or `source_entity` should be configured. See source_entity.
        """
        return pulumi.get(self, "source_entity")

    @property
    @pulumi.getter(name="sourceEntityArn")
    def source_entity_arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of an analysis or template that was used to create this template.
        """
        return pulumi.get(self, "source_entity_arn")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The template creation status.
        """
        return pulumi.get(self, "status")

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
        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Output[str]:
        """
        Identifier for the template.
        """
        return pulumi.get(self, "template_id")

    @property
    @pulumi.getter(name="versionDescription")
    def version_description(self) -> pulumi.Output[str]:
        """
        A description of the current template version being created/updated.
        """
        return pulumi.get(self, "version_description")

    @property
    @pulumi.getter(name="versionNumber")
    def version_number(self) -> pulumi.Output[int]:
        """
        The version number of the template version.
        """
        return pulumi.get(self, "version_number")

