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

__all__ = ['ConfigurationTemplateArgs', 'ConfigurationTemplate']

@pulumi.input_type
class ConfigurationTemplateArgs:
    def __init__(__self__, *,
                 application: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input[Sequence[pulumi.Input['ConfigurationTemplateSettingArgs']]]] = None,
                 solution_stack_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ConfigurationTemplate resource.
        :param pulumi.Input[str] application: name of the application to associate with this configuration template
        :param pulumi.Input[str] description: Short description of the Template
        :param pulumi.Input[str] environment_id: The ID of the environment used with this configuration template
        :param pulumi.Input[str] name: A unique name for this Template.
        :param pulumi.Input[Sequence[pulumi.Input['ConfigurationTemplateSettingArgs']]] settings: Option settings to configure the new Environment. These
               override specific values that are set as defaults. The format is detailed
               below in Option Settings
        :param pulumi.Input[str] solution_stack_name: A solution stack to base your Template
               off of. Example stacks can be found in the [Amazon API documentation][1]
        """
        pulumi.set(__self__, "application", application)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if environment_id is not None:
            pulumi.set(__self__, "environment_id", environment_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if settings is not None:
            pulumi.set(__self__, "settings", settings)
        if solution_stack_name is not None:
            pulumi.set(__self__, "solution_stack_name", solution_stack_name)

    @property
    @pulumi.getter
    def application(self) -> pulumi.Input[str]:
        """
        name of the application to associate with this configuration template
        """
        return pulumi.get(self, "application")

    @application.setter
    def application(self, value: pulumi.Input[str]):
        pulumi.set(self, "application", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Short description of the Template
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the environment used with this configuration template
        """
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for this Template.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def settings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ConfigurationTemplateSettingArgs']]]]:
        """
        Option settings to configure the new Environment. These
        override specific values that are set as defaults. The format is detailed
        below in Option Settings
        """
        return pulumi.get(self, "settings")

    @settings.setter
    def settings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ConfigurationTemplateSettingArgs']]]]):
        pulumi.set(self, "settings", value)

    @property
    @pulumi.getter(name="solutionStackName")
    def solution_stack_name(self) -> Optional[pulumi.Input[str]]:
        """
        A solution stack to base your Template
        off of. Example stacks can be found in the [Amazon API documentation][1]
        """
        return pulumi.get(self, "solution_stack_name")

    @solution_stack_name.setter
    def solution_stack_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "solution_stack_name", value)


@pulumi.input_type
class _ConfigurationTemplateState:
    def __init__(__self__, *,
                 application: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input[Sequence[pulumi.Input['ConfigurationTemplateSettingArgs']]]] = None,
                 solution_stack_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ConfigurationTemplate resources.
        :param pulumi.Input[str] application: name of the application to associate with this configuration template
        :param pulumi.Input[str] description: Short description of the Template
        :param pulumi.Input[str] environment_id: The ID of the environment used with this configuration template
        :param pulumi.Input[str] name: A unique name for this Template.
        :param pulumi.Input[Sequence[pulumi.Input['ConfigurationTemplateSettingArgs']]] settings: Option settings to configure the new Environment. These
               override specific values that are set as defaults. The format is detailed
               below in Option Settings
        :param pulumi.Input[str] solution_stack_name: A solution stack to base your Template
               off of. Example stacks can be found in the [Amazon API documentation][1]
        """
        if application is not None:
            pulumi.set(__self__, "application", application)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if environment_id is not None:
            pulumi.set(__self__, "environment_id", environment_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if settings is not None:
            pulumi.set(__self__, "settings", settings)
        if solution_stack_name is not None:
            pulumi.set(__self__, "solution_stack_name", solution_stack_name)

    @property
    @pulumi.getter
    def application(self) -> Optional[pulumi.Input[str]]:
        """
        name of the application to associate with this configuration template
        """
        return pulumi.get(self, "application")

    @application.setter
    def application(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Short description of the Template
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the environment used with this configuration template
        """
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for this Template.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def settings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ConfigurationTemplateSettingArgs']]]]:
        """
        Option settings to configure the new Environment. These
        override specific values that are set as defaults. The format is detailed
        below in Option Settings
        """
        return pulumi.get(self, "settings")

    @settings.setter
    def settings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ConfigurationTemplateSettingArgs']]]]):
        pulumi.set(self, "settings", value)

    @property
    @pulumi.getter(name="solutionStackName")
    def solution_stack_name(self) -> Optional[pulumi.Input[str]]:
        """
        A solution stack to base your Template
        off of. Example stacks can be found in the [Amazon API documentation][1]
        """
        return pulumi.get(self, "solution_stack_name")

    @solution_stack_name.setter
    def solution_stack_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "solution_stack_name", value)


class ConfigurationTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConfigurationTemplateSettingArgs']]]]] = None,
                 solution_stack_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an Elastic Beanstalk Configuration Template, which are associated with
        a specific application and are used to deploy different versions of the
        application with the same configuration settings.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        tftest = aws.elasticbeanstalk.Application("tftest", description="tf-test-desc")
        my_template = aws.elasticbeanstalk.ConfigurationTemplate("myTemplate",
            application=tftest.name,
            solution_stack_name="64bit Amazon Linux 2015.09 v2.0.8 running Go 1.4")
        ```
        ## Option Settings

        The `setting` field supports the following format:

        * `namespace` - unique namespace identifying the option's associated AWS resource
        * `name` - name of the configuration option
        * `value` - value for the configuration option
        * `resource` - (Optional) resource name for [scheduled action](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options-general.html#command-options-general-autoscalingscheduledaction)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application: name of the application to associate with this configuration template
        :param pulumi.Input[str] description: Short description of the Template
        :param pulumi.Input[str] environment_id: The ID of the environment used with this configuration template
        :param pulumi.Input[str] name: A unique name for this Template.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConfigurationTemplateSettingArgs']]]] settings: Option settings to configure the new Environment. These
               override specific values that are set as defaults. The format is detailed
               below in Option Settings
        :param pulumi.Input[str] solution_stack_name: A solution stack to base your Template
               off of. Example stacks can be found in the [Amazon API documentation][1]
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConfigurationTemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Elastic Beanstalk Configuration Template, which are associated with
        a specific application and are used to deploy different versions of the
        application with the same configuration settings.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        tftest = aws.elasticbeanstalk.Application("tftest", description="tf-test-desc")
        my_template = aws.elasticbeanstalk.ConfigurationTemplate("myTemplate",
            application=tftest.name,
            solution_stack_name="64bit Amazon Linux 2015.09 v2.0.8 running Go 1.4")
        ```
        ## Option Settings

        The `setting` field supports the following format:

        * `namespace` - unique namespace identifying the option's associated AWS resource
        * `name` - name of the configuration option
        * `value` - value for the configuration option
        * `resource` - (Optional) resource name for [scheduled action](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options-general.html#command-options-general-autoscalingscheduledaction)

        :param str resource_name: The name of the resource.
        :param ConfigurationTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConfigurationTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConfigurationTemplateSettingArgs']]]]] = None,
                 solution_stack_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConfigurationTemplateArgs.__new__(ConfigurationTemplateArgs)

            if application is None and not opts.urn:
                raise TypeError("Missing required property 'application'")
            __props__.__dict__["application"] = application
            __props__.__dict__["description"] = description
            __props__.__dict__["environment_id"] = environment_id
            __props__.__dict__["name"] = name
            __props__.__dict__["settings"] = settings
            __props__.__dict__["solution_stack_name"] = solution_stack_name
        super(ConfigurationTemplate, __self__).__init__(
            'aws:elasticbeanstalk/configurationTemplate:ConfigurationTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            environment_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            settings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConfigurationTemplateSettingArgs']]]]] = None,
            solution_stack_name: Optional[pulumi.Input[str]] = None) -> 'ConfigurationTemplate':
        """
        Get an existing ConfigurationTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application: name of the application to associate with this configuration template
        :param pulumi.Input[str] description: Short description of the Template
        :param pulumi.Input[str] environment_id: The ID of the environment used with this configuration template
        :param pulumi.Input[str] name: A unique name for this Template.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConfigurationTemplateSettingArgs']]]] settings: Option settings to configure the new Environment. These
               override specific values that are set as defaults. The format is detailed
               below in Option Settings
        :param pulumi.Input[str] solution_stack_name: A solution stack to base your Template
               off of. Example stacks can be found in the [Amazon API documentation][1]
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConfigurationTemplateState.__new__(_ConfigurationTemplateState)

        __props__.__dict__["application"] = application
        __props__.__dict__["description"] = description
        __props__.__dict__["environment_id"] = environment_id
        __props__.__dict__["name"] = name
        __props__.__dict__["settings"] = settings
        __props__.__dict__["solution_stack_name"] = solution_stack_name
        return ConfigurationTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def application(self) -> pulumi.Output[str]:
        """
        name of the application to associate with this configuration template
        """
        return pulumi.get(self, "application")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Short description of the Template
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the environment used with this configuration template
        """
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique name for this Template.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def settings(self) -> pulumi.Output[Sequence['outputs.ConfigurationTemplateSetting']]:
        """
        Option settings to configure the new Environment. These
        override specific values that are set as defaults. The format is detailed
        below in Option Settings
        """
        return pulumi.get(self, "settings")

    @property
    @pulumi.getter(name="solutionStackName")
    def solution_stack_name(self) -> pulumi.Output[Optional[str]]:
        """
        A solution stack to base your Template
        off of. Example stacks can be found in the [Amazon API documentation][1]
        """
        return pulumi.get(self, "solution_stack_name")

