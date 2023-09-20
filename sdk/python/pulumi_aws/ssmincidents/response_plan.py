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

__all__ = ['ResponsePlanArgs', 'ResponsePlan']

@pulumi.input_type
class ResponsePlanArgs:
    def __init__(__self__, *,
                 incident_template: pulumi.Input['ResponsePlanIncidentTemplateArgs'],
                 action: Optional[pulumi.Input['ResponsePlanActionArgs']] = None,
                 chat_channels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 engagements: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 integration: Optional[pulumi.Input['ResponsePlanIntegrationArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ResponsePlan resource.
        :param pulumi.Input['ResponsePlanActionArgs'] action: The actions that the response plan starts at the beginning of an incident.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] chat_channels: The Chatbot chat channel used for collaboration during an incident.
        :param pulumi.Input[str] display_name: The long format of the response plan name. This field can contain spaces.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] engagements: The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.
        :param pulumi.Input['ResponsePlanIntegrationArgs'] integration: Information about third-party services integrated into the response plan. The following values are supported:
        :param pulumi.Input[str] name: The name of the response plan.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags applied to the response plan.
        """
        pulumi.set(__self__, "incident_template", incident_template)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if chat_channels is not None:
            pulumi.set(__self__, "chat_channels", chat_channels)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if engagements is not None:
            pulumi.set(__self__, "engagements", engagements)
        if integration is not None:
            pulumi.set(__self__, "integration", integration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="incidentTemplate")
    def incident_template(self) -> pulumi.Input['ResponsePlanIncidentTemplateArgs']:
        return pulumi.get(self, "incident_template")

    @incident_template.setter
    def incident_template(self, value: pulumi.Input['ResponsePlanIncidentTemplateArgs']):
        pulumi.set(self, "incident_template", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input['ResponsePlanActionArgs']]:
        """
        The actions that the response plan starts at the beginning of an incident.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input['ResponsePlanActionArgs']]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="chatChannels")
    def chat_channels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The Chatbot chat channel used for collaboration during an incident.
        """
        return pulumi.get(self, "chat_channels")

    @chat_channels.setter
    def chat_channels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "chat_channels", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The long format of the response plan name. This field can contain spaces.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def engagements(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.
        """
        return pulumi.get(self, "engagements")

    @engagements.setter
    def engagements(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "engagements", value)

    @property
    @pulumi.getter
    def integration(self) -> Optional[pulumi.Input['ResponsePlanIntegrationArgs']]:
        """
        Information about third-party services integrated into the response plan. The following values are supported:
        """
        return pulumi.get(self, "integration")

    @integration.setter
    def integration(self, value: Optional[pulumi.Input['ResponsePlanIntegrationArgs']]):
        pulumi.set(self, "integration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the response plan.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The tags applied to the response plan.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ResponsePlanState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input['ResponsePlanActionArgs']] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 chat_channels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 engagements: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 incident_template: Optional[pulumi.Input['ResponsePlanIncidentTemplateArgs']] = None,
                 integration: Optional[pulumi.Input['ResponsePlanIntegrationArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ResponsePlan resources.
        :param pulumi.Input['ResponsePlanActionArgs'] action: The actions that the response plan starts at the beginning of an incident.
        :param pulumi.Input[str] arn: The ARN of the response plan.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] chat_channels: The Chatbot chat channel used for collaboration during an incident.
        :param pulumi.Input[str] display_name: The long format of the response plan name. This field can contain spaces.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] engagements: The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.
        :param pulumi.Input['ResponsePlanIntegrationArgs'] integration: Information about third-party services integrated into the response plan. The following values are supported:
        :param pulumi.Input[str] name: The name of the response plan.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags applied to the response plan.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if chat_channels is not None:
            pulumi.set(__self__, "chat_channels", chat_channels)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if engagements is not None:
            pulumi.set(__self__, "engagements", engagements)
        if incident_template is not None:
            pulumi.set(__self__, "incident_template", incident_template)
        if integration is not None:
            pulumi.set(__self__, "integration", integration)
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
    def action(self) -> Optional[pulumi.Input['ResponsePlanActionArgs']]:
        """
        The actions that the response plan starts at the beginning of an incident.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input['ResponsePlanActionArgs']]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the response plan.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="chatChannels")
    def chat_channels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The Chatbot chat channel used for collaboration during an incident.
        """
        return pulumi.get(self, "chat_channels")

    @chat_channels.setter
    def chat_channels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "chat_channels", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The long format of the response plan name. This field can contain spaces.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def engagements(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.
        """
        return pulumi.get(self, "engagements")

    @engagements.setter
    def engagements(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "engagements", value)

    @property
    @pulumi.getter(name="incidentTemplate")
    def incident_template(self) -> Optional[pulumi.Input['ResponsePlanIncidentTemplateArgs']]:
        return pulumi.get(self, "incident_template")

    @incident_template.setter
    def incident_template(self, value: Optional[pulumi.Input['ResponsePlanIncidentTemplateArgs']]):
        pulumi.set(self, "incident_template", value)

    @property
    @pulumi.getter
    def integration(self) -> Optional[pulumi.Input['ResponsePlanIntegrationArgs']]:
        """
        Information about third-party services integrated into the response plan. The following values are supported:
        """
        return pulumi.get(self, "integration")

    @integration.setter
    def integration(self, value: Optional[pulumi.Input['ResponsePlanIntegrationArgs']]):
        pulumi.set(self, "integration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the response plan.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The tags applied to the response plan.
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


class ResponsePlan(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[pulumi.InputType['ResponsePlanActionArgs']]] = None,
                 chat_channels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 engagements: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 incident_template: Optional[pulumi.Input[pulumi.InputType['ResponsePlanIncidentTemplateArgs']]] = None,
                 integration: Optional[pulumi.Input[pulumi.InputType['ResponsePlanIntegrationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to manage response plans in AWS Systems Manager Incident Manager.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ssmincidents.ResponsePlan("example",
            incident_template=aws.ssmincidents.ResponsePlanIncidentTemplateArgs(
                title="title",
                impact=3,
            ),
            tags={
                "key": "value",
            },
            opts=pulumi.ResourceOptions(depends_on=[aws_ssmincidents_replication_set["example"]]))
        ```
        ### Usage With All Fields

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ssmincidents.ResponsePlan("example",
            incident_template=aws.ssmincidents.ResponsePlanIncidentTemplateArgs(
                title="title",
                impact=3,
                dedupe_string="dedupe",
                incident_tags={
                    "key": "value",
                },
                notification_targets=[
                    aws.ssmincidents.ResponsePlanIncidentTemplateNotificationTargetArgs(
                        sns_topic_arn=aws_sns_topic["example1"]["arn"],
                    ),
                    aws.ssmincidents.ResponsePlanIncidentTemplateNotificationTargetArgs(
                        sns_topic_arn=aws_sns_topic["example2"]["arn"],
                    ),
                ],
                summary="summary",
            ),
            display_name="display name",
            chat_channels=[aws_sns_topic["topic"]["arn"]],
            engagements=["arn:aws:ssm-contacts:us-east-2:111122223333:contact/test1"],
            action=aws.ssmincidents.ResponsePlanActionArgs(
                ssm_automations=[aws.ssmincidents.ResponsePlanActionSsmAutomationArgs(
                    document_name=aws_ssm_document["document1"]["name"],
                    role_arn=aws_iam_role["role1"]["arn"],
                    document_version="version1",
                    target_account="RESPONSE_PLAN_OWNER_ACCOUNT",
                    parameters=[
                        aws.ssmincidents.ResponsePlanActionSsmAutomationParameterArgs(
                            name="key",
                            values=[
                                "value1",
                                "value2",
                            ],
                        ),
                        aws.ssmincidents.ResponsePlanActionSsmAutomationParameterArgs(
                            name="foo",
                            values=["bar"],
                        ),
                    ],
                    dynamic_parameters={
                        "someKey": "INVOLVED_RESOURCES",
                        "anotherKey": "INCIDENT_RECORD_ARN",
                    },
                )],
            ),
            integration=aws.ssmincidents.ResponsePlanIntegrationArgs(
                pagerduties=[aws.ssmincidents.ResponsePlanIntegrationPagerdutyArgs(
                    name="pagerdutyIntergration",
                    service_id="example",
                    secret_id="example",
                )],
            ),
            tags={
                "key": "value",
            },
            opts=pulumi.ResourceOptions(depends_on=[aws_ssmincidents_replication_set["example"]]))
        ```

        ## Import

        Using `pulumi import`, import an Incident Manager response plan using the response plan ARN. You can find the response plan ARN in the AWS Management Console. For example:

        ```sh
         $ pulumi import aws:ssmincidents/responsePlan:ResponsePlan responsePlanName ARNValue
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ResponsePlanActionArgs']] action: The actions that the response plan starts at the beginning of an incident.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] chat_channels: The Chatbot chat channel used for collaboration during an incident.
        :param pulumi.Input[str] display_name: The long format of the response plan name. This field can contain spaces.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] engagements: The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.
        :param pulumi.Input[pulumi.InputType['ResponsePlanIntegrationArgs']] integration: Information about third-party services integrated into the response plan. The following values are supported:
        :param pulumi.Input[str] name: The name of the response plan.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags applied to the response plan.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResponsePlanArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage response plans in AWS Systems Manager Incident Manager.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ssmincidents.ResponsePlan("example",
            incident_template=aws.ssmincidents.ResponsePlanIncidentTemplateArgs(
                title="title",
                impact=3,
            ),
            tags={
                "key": "value",
            },
            opts=pulumi.ResourceOptions(depends_on=[aws_ssmincidents_replication_set["example"]]))
        ```
        ### Usage With All Fields

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ssmincidents.ResponsePlan("example",
            incident_template=aws.ssmincidents.ResponsePlanIncidentTemplateArgs(
                title="title",
                impact=3,
                dedupe_string="dedupe",
                incident_tags={
                    "key": "value",
                },
                notification_targets=[
                    aws.ssmincidents.ResponsePlanIncidentTemplateNotificationTargetArgs(
                        sns_topic_arn=aws_sns_topic["example1"]["arn"],
                    ),
                    aws.ssmincidents.ResponsePlanIncidentTemplateNotificationTargetArgs(
                        sns_topic_arn=aws_sns_topic["example2"]["arn"],
                    ),
                ],
                summary="summary",
            ),
            display_name="display name",
            chat_channels=[aws_sns_topic["topic"]["arn"]],
            engagements=["arn:aws:ssm-contacts:us-east-2:111122223333:contact/test1"],
            action=aws.ssmincidents.ResponsePlanActionArgs(
                ssm_automations=[aws.ssmincidents.ResponsePlanActionSsmAutomationArgs(
                    document_name=aws_ssm_document["document1"]["name"],
                    role_arn=aws_iam_role["role1"]["arn"],
                    document_version="version1",
                    target_account="RESPONSE_PLAN_OWNER_ACCOUNT",
                    parameters=[
                        aws.ssmincidents.ResponsePlanActionSsmAutomationParameterArgs(
                            name="key",
                            values=[
                                "value1",
                                "value2",
                            ],
                        ),
                        aws.ssmincidents.ResponsePlanActionSsmAutomationParameterArgs(
                            name="foo",
                            values=["bar"],
                        ),
                    ],
                    dynamic_parameters={
                        "someKey": "INVOLVED_RESOURCES",
                        "anotherKey": "INCIDENT_RECORD_ARN",
                    },
                )],
            ),
            integration=aws.ssmincidents.ResponsePlanIntegrationArgs(
                pagerduties=[aws.ssmincidents.ResponsePlanIntegrationPagerdutyArgs(
                    name="pagerdutyIntergration",
                    service_id="example",
                    secret_id="example",
                )],
            ),
            tags={
                "key": "value",
            },
            opts=pulumi.ResourceOptions(depends_on=[aws_ssmincidents_replication_set["example"]]))
        ```

        ## Import

        Using `pulumi import`, import an Incident Manager response plan using the response plan ARN. You can find the response plan ARN in the AWS Management Console. For example:

        ```sh
         $ pulumi import aws:ssmincidents/responsePlan:ResponsePlan responsePlanName ARNValue
        ```

        :param str resource_name: The name of the resource.
        :param ResponsePlanArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResponsePlanArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[pulumi.InputType['ResponsePlanActionArgs']]] = None,
                 chat_channels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 engagements: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 incident_template: Optional[pulumi.Input[pulumi.InputType['ResponsePlanIncidentTemplateArgs']]] = None,
                 integration: Optional[pulumi.Input[pulumi.InputType['ResponsePlanIntegrationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResponsePlanArgs.__new__(ResponsePlanArgs)

            __props__.__dict__["action"] = action
            __props__.__dict__["chat_channels"] = chat_channels
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["engagements"] = engagements
            if incident_template is None and not opts.urn:
                raise TypeError("Missing required property 'incident_template'")
            __props__.__dict__["incident_template"] = incident_template
            __props__.__dict__["integration"] = integration
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ResponsePlan, __self__).__init__(
            'aws:ssmincidents/responsePlan:ResponsePlan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[pulumi.InputType['ResponsePlanActionArgs']]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            chat_channels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            engagements: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            incident_template: Optional[pulumi.Input[pulumi.InputType['ResponsePlanIncidentTemplateArgs']]] = None,
            integration: Optional[pulumi.Input[pulumi.InputType['ResponsePlanIntegrationArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'ResponsePlan':
        """
        Get an existing ResponsePlan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ResponsePlanActionArgs']] action: The actions that the response plan starts at the beginning of an incident.
        :param pulumi.Input[str] arn: The ARN of the response plan.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] chat_channels: The Chatbot chat channel used for collaboration during an incident.
        :param pulumi.Input[str] display_name: The long format of the response plan name. This field can contain spaces.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] engagements: The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.
        :param pulumi.Input[pulumi.InputType['ResponsePlanIntegrationArgs']] integration: Information about third-party services integrated into the response plan. The following values are supported:
        :param pulumi.Input[str] name: The name of the response plan.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags applied to the response plan.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ResponsePlanState.__new__(_ResponsePlanState)

        __props__.__dict__["action"] = action
        __props__.__dict__["arn"] = arn
        __props__.__dict__["chat_channels"] = chat_channels
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["engagements"] = engagements
        __props__.__dict__["incident_template"] = incident_template
        __props__.__dict__["integration"] = integration
        __props__.__dict__["name"] = name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return ResponsePlan(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[Optional['outputs.ResponsePlanAction']]:
        """
        The actions that the response plan starts at the beginning of an incident.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the response plan.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="chatChannels")
    def chat_channels(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The Chatbot chat channel used for collaboration during an incident.
        """
        return pulumi.get(self, "chat_channels")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        The long format of the response plan name. This field can contain spaces.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def engagements(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.
        """
        return pulumi.get(self, "engagements")

    @property
    @pulumi.getter(name="incidentTemplate")
    def incident_template(self) -> pulumi.Output['outputs.ResponsePlanIncidentTemplate']:
        return pulumi.get(self, "incident_template")

    @property
    @pulumi.getter
    def integration(self) -> pulumi.Output[Optional['outputs.ResponsePlanIntegration']]:
        """
        Information about third-party services integrated into the response plan. The following values are supported:
        """
        return pulumi.get(self, "integration")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the response plan.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The tags applied to the response plan.
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

