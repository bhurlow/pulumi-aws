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

__all__ = ['WorkteamArgs', 'Workteam']

@pulumi.input_type
class WorkteamArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 member_definitions: pulumi.Input[Sequence[pulumi.Input['WorkteamMemberDefinitionArgs']]],
                 workforce_name: pulumi.Input[str],
                 workteam_name: pulumi.Input[str],
                 notification_configuration: Optional[pulumi.Input['WorkteamNotificationConfigurationArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Workteam resource.
        :param pulumi.Input[str] description: A description of the work team.
        :param pulumi.Input[Sequence[pulumi.Input['WorkteamMemberDefinitionArgs']]] member_definitions: A list of Member Definitions that contains objects that identify the workers that make up the work team. Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use `cognito_member_definition`. For workforces created using your own OIDC identity provider (IdP) use `oidc_member_definition`. Do not provide input for both of these parameters in a single request. see Member Definition details below.
        :param pulumi.Input[str] workforce_name: The name of the Workteam (must be unique).
        :param pulumi.Input[str] workteam_name: The name of the workforce.
        :param pulumi.Input['WorkteamNotificationConfigurationArgs'] notification_configuration: Configures notification of workers regarding available or expiring work items. see Notification Configuration details below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "member_definitions", member_definitions)
        pulumi.set(__self__, "workforce_name", workforce_name)
        pulumi.set(__self__, "workteam_name", workteam_name)
        if notification_configuration is not None:
            pulumi.set(__self__, "notification_configuration", notification_configuration)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        A description of the work team.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="memberDefinitions")
    def member_definitions(self) -> pulumi.Input[Sequence[pulumi.Input['WorkteamMemberDefinitionArgs']]]:
        """
        A list of Member Definitions that contains objects that identify the workers that make up the work team. Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use `cognito_member_definition`. For workforces created using your own OIDC identity provider (IdP) use `oidc_member_definition`. Do not provide input for both of these parameters in a single request. see Member Definition details below.
        """
        return pulumi.get(self, "member_definitions")

    @member_definitions.setter
    def member_definitions(self, value: pulumi.Input[Sequence[pulumi.Input['WorkteamMemberDefinitionArgs']]]):
        pulumi.set(self, "member_definitions", value)

    @property
    @pulumi.getter(name="workforceName")
    def workforce_name(self) -> pulumi.Input[str]:
        """
        The name of the Workteam (must be unique).
        """
        return pulumi.get(self, "workforce_name")

    @workforce_name.setter
    def workforce_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "workforce_name", value)

    @property
    @pulumi.getter(name="workteamName")
    def workteam_name(self) -> pulumi.Input[str]:
        """
        The name of the workforce.
        """
        return pulumi.get(self, "workteam_name")

    @workteam_name.setter
    def workteam_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "workteam_name", value)

    @property
    @pulumi.getter(name="notificationConfiguration")
    def notification_configuration(self) -> Optional[pulumi.Input['WorkteamNotificationConfigurationArgs']]:
        """
        Configures notification of workers regarding available or expiring work items. see Notification Configuration details below.
        """
        return pulumi.get(self, "notification_configuration")

    @notification_configuration.setter
    def notification_configuration(self, value: Optional[pulumi.Input['WorkteamNotificationConfigurationArgs']]):
        pulumi.set(self, "notification_configuration", value)

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
class _WorkteamState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 member_definitions: Optional[pulumi.Input[Sequence[pulumi.Input['WorkteamMemberDefinitionArgs']]]] = None,
                 notification_configuration: Optional[pulumi.Input['WorkteamNotificationConfigurationArgs']] = None,
                 subdomain: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 workforce_name: Optional[pulumi.Input[str]] = None,
                 workteam_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Workteam resources.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) assigned by AWS to this Workteam.
        :param pulumi.Input[str] description: A description of the work team.
        :param pulumi.Input[Sequence[pulumi.Input['WorkteamMemberDefinitionArgs']]] member_definitions: A list of Member Definitions that contains objects that identify the workers that make up the work team. Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use `cognito_member_definition`. For workforces created using your own OIDC identity provider (IdP) use `oidc_member_definition`. Do not provide input for both of these parameters in a single request. see Member Definition details below.
        :param pulumi.Input['WorkteamNotificationConfigurationArgs'] notification_configuration: Configures notification of workers regarding available or expiring work items. see Notification Configuration details below.
        :param pulumi.Input[str] subdomain: The subdomain for your OIDC Identity Provider.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[str] workforce_name: The name of the Workteam (must be unique).
        :param pulumi.Input[str] workteam_name: The name of the workforce.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if member_definitions is not None:
            pulumi.set(__self__, "member_definitions", member_definitions)
        if notification_configuration is not None:
            pulumi.set(__self__, "notification_configuration", notification_configuration)
        if subdomain is not None:
            pulumi.set(__self__, "subdomain", subdomain)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if workforce_name is not None:
            pulumi.set(__self__, "workforce_name", workforce_name)
        if workteam_name is not None:
            pulumi.set(__self__, "workteam_name", workteam_name)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) assigned by AWS to this Workteam.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the work team.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="memberDefinitions")
    def member_definitions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WorkteamMemberDefinitionArgs']]]]:
        """
        A list of Member Definitions that contains objects that identify the workers that make up the work team. Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use `cognito_member_definition`. For workforces created using your own OIDC identity provider (IdP) use `oidc_member_definition`. Do not provide input for both of these parameters in a single request. see Member Definition details below.
        """
        return pulumi.get(self, "member_definitions")

    @member_definitions.setter
    def member_definitions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WorkteamMemberDefinitionArgs']]]]):
        pulumi.set(self, "member_definitions", value)

    @property
    @pulumi.getter(name="notificationConfiguration")
    def notification_configuration(self) -> Optional[pulumi.Input['WorkteamNotificationConfigurationArgs']]:
        """
        Configures notification of workers regarding available or expiring work items. see Notification Configuration details below.
        """
        return pulumi.get(self, "notification_configuration")

    @notification_configuration.setter
    def notification_configuration(self, value: Optional[pulumi.Input['WorkteamNotificationConfigurationArgs']]):
        pulumi.set(self, "notification_configuration", value)

    @property
    @pulumi.getter
    def subdomain(self) -> Optional[pulumi.Input[str]]:
        """
        The subdomain for your OIDC Identity Provider.
        """
        return pulumi.get(self, "subdomain")

    @subdomain.setter
    def subdomain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subdomain", value)

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
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter(name="workforceName")
    def workforce_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Workteam (must be unique).
        """
        return pulumi.get(self, "workforce_name")

    @workforce_name.setter
    def workforce_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workforce_name", value)

    @property
    @pulumi.getter(name="workteamName")
    def workteam_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the workforce.
        """
        return pulumi.get(self, "workteam_name")

    @workteam_name.setter
    def workteam_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workteam_name", value)


class Workteam(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 member_definitions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WorkteamMemberDefinitionArgs']]]]] = None,
                 notification_configuration: Optional[pulumi.Input[pulumi.InputType['WorkteamNotificationConfigurationArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 workforce_name: Optional[pulumi.Input[str]] = None,
                 workteam_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a SageMaker Workteam resource.

        ## Example Usage
        ### Cognito Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sagemaker.Workteam("example",
            workteam_name="example",
            workforce_name=aws_sagemaker_workforce["example"]["id"],
            description="example",
            member_definitions=[aws.sagemaker.WorkteamMemberDefinitionArgs(
                cognito_member_definition=aws.sagemaker.WorkteamMemberDefinitionCognitoMemberDefinitionArgs(
                    client_id=aws_cognito_user_pool_client["example"]["id"],
                    user_pool=aws_cognito_user_pool_domain["example"]["user_pool_id"],
                    user_group=aws_cognito_user_group["example"]["id"],
                ),
            )])
        ```
        ### Oidc Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sagemaker.Workteam("example",
            workteam_name="example",
            workforce_name=aws_sagemaker_workforce["example"]["id"],
            description="example",
            member_definitions=[aws.sagemaker.WorkteamMemberDefinitionArgs(
                oidc_member_definition=aws.sagemaker.WorkteamMemberDefinitionOidcMemberDefinitionArgs(
                    groups=["example"],
                ),
            )])
        ```

        ## Import

        SageMaker Workteams can be imported using the `workteam_name`, e.g.,

        ```sh
         $ pulumi import aws:sagemaker/workteam:Workteam example example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the work team.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WorkteamMemberDefinitionArgs']]]] member_definitions: A list of Member Definitions that contains objects that identify the workers that make up the work team. Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use `cognito_member_definition`. For workforces created using your own OIDC identity provider (IdP) use `oidc_member_definition`. Do not provide input for both of these parameters in a single request. see Member Definition details below.
        :param pulumi.Input[pulumi.InputType['WorkteamNotificationConfigurationArgs']] notification_configuration: Configures notification of workers regarding available or expiring work items. see Notification Configuration details below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] workforce_name: The name of the Workteam (must be unique).
        :param pulumi.Input[str] workteam_name: The name of the workforce.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WorkteamArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a SageMaker Workteam resource.

        ## Example Usage
        ### Cognito Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sagemaker.Workteam("example",
            workteam_name="example",
            workforce_name=aws_sagemaker_workforce["example"]["id"],
            description="example",
            member_definitions=[aws.sagemaker.WorkteamMemberDefinitionArgs(
                cognito_member_definition=aws.sagemaker.WorkteamMemberDefinitionCognitoMemberDefinitionArgs(
                    client_id=aws_cognito_user_pool_client["example"]["id"],
                    user_pool=aws_cognito_user_pool_domain["example"]["user_pool_id"],
                    user_group=aws_cognito_user_group["example"]["id"],
                ),
            )])
        ```
        ### Oidc Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sagemaker.Workteam("example",
            workteam_name="example",
            workforce_name=aws_sagemaker_workforce["example"]["id"],
            description="example",
            member_definitions=[aws.sagemaker.WorkteamMemberDefinitionArgs(
                oidc_member_definition=aws.sagemaker.WorkteamMemberDefinitionOidcMemberDefinitionArgs(
                    groups=["example"],
                ),
            )])
        ```

        ## Import

        SageMaker Workteams can be imported using the `workteam_name`, e.g.,

        ```sh
         $ pulumi import aws:sagemaker/workteam:Workteam example example
        ```

        :param str resource_name: The name of the resource.
        :param WorkteamArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WorkteamArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 member_definitions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WorkteamMemberDefinitionArgs']]]]] = None,
                 notification_configuration: Optional[pulumi.Input[pulumi.InputType['WorkteamNotificationConfigurationArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 workforce_name: Optional[pulumi.Input[str]] = None,
                 workteam_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WorkteamArgs.__new__(WorkteamArgs)

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            if member_definitions is None and not opts.urn:
                raise TypeError("Missing required property 'member_definitions'")
            __props__.__dict__["member_definitions"] = member_definitions
            __props__.__dict__["notification_configuration"] = notification_configuration
            __props__.__dict__["tags"] = tags
            if workforce_name is None and not opts.urn:
                raise TypeError("Missing required property 'workforce_name'")
            __props__.__dict__["workforce_name"] = workforce_name
            if workteam_name is None and not opts.urn:
                raise TypeError("Missing required property 'workteam_name'")
            __props__.__dict__["workteam_name"] = workteam_name
            __props__.__dict__["arn"] = None
            __props__.__dict__["subdomain"] = None
            __props__.__dict__["tags_all"] = None
        super(Workteam, __self__).__init__(
            'aws:sagemaker/workteam:Workteam',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            member_definitions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WorkteamMemberDefinitionArgs']]]]] = None,
            notification_configuration: Optional[pulumi.Input[pulumi.InputType['WorkteamNotificationConfigurationArgs']]] = None,
            subdomain: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            workforce_name: Optional[pulumi.Input[str]] = None,
            workteam_name: Optional[pulumi.Input[str]] = None) -> 'Workteam':
        """
        Get an existing Workteam resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) assigned by AWS to this Workteam.
        :param pulumi.Input[str] description: A description of the work team.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WorkteamMemberDefinitionArgs']]]] member_definitions: A list of Member Definitions that contains objects that identify the workers that make up the work team. Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use `cognito_member_definition`. For workforces created using your own OIDC identity provider (IdP) use `oidc_member_definition`. Do not provide input for both of these parameters in a single request. see Member Definition details below.
        :param pulumi.Input[pulumi.InputType['WorkteamNotificationConfigurationArgs']] notification_configuration: Configures notification of workers regarding available or expiring work items. see Notification Configuration details below.
        :param pulumi.Input[str] subdomain: The subdomain for your OIDC Identity Provider.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[str] workforce_name: The name of the Workteam (must be unique).
        :param pulumi.Input[str] workteam_name: The name of the workforce.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WorkteamState.__new__(_WorkteamState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["description"] = description
        __props__.__dict__["member_definitions"] = member_definitions
        __props__.__dict__["notification_configuration"] = notification_configuration
        __props__.__dict__["subdomain"] = subdomain
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["workforce_name"] = workforce_name
        __props__.__dict__["workteam_name"] = workteam_name
        return Workteam(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) assigned by AWS to this Workteam.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        A description of the work team.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="memberDefinitions")
    def member_definitions(self) -> pulumi.Output[Sequence['outputs.WorkteamMemberDefinition']]:
        """
        A list of Member Definitions that contains objects that identify the workers that make up the work team. Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use `cognito_member_definition`. For workforces created using your own OIDC identity provider (IdP) use `oidc_member_definition`. Do not provide input for both of these parameters in a single request. see Member Definition details below.
        """
        return pulumi.get(self, "member_definitions")

    @property
    @pulumi.getter(name="notificationConfiguration")
    def notification_configuration(self) -> pulumi.Output[Optional['outputs.WorkteamNotificationConfiguration']]:
        """
        Configures notification of workers regarding available or expiring work items. see Notification Configuration details below.
        """
        return pulumi.get(self, "notification_configuration")

    @property
    @pulumi.getter
    def subdomain(self) -> pulumi.Output[str]:
        """
        The subdomain for your OIDC Identity Provider.
        """
        return pulumi.get(self, "subdomain")

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
        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter(name="workforceName")
    def workforce_name(self) -> pulumi.Output[str]:
        """
        The name of the Workteam (must be unique).
        """
        return pulumi.get(self, "workforce_name")

    @property
    @pulumi.getter(name="workteamName")
    def workteam_name(self) -> pulumi.Output[str]:
        """
        The name of the workforce.
        """
        return pulumi.get(self, "workteam_name")

