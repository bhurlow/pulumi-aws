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

__all__ = ['ListenerArgs', 'Listener']

@pulumi.input_type
class ListenerArgs:
    def __init__(__self__, *,
                 default_action: pulumi.Input['ListenerDefaultActionArgs'],
                 protocol: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 service_arn: Optional[pulumi.Input[str]] = None,
                 service_identifier: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Listener resource.
        :param pulumi.Input['ListenerDefaultActionArgs'] default_action: Default action block for the default listener rule. Default action blocks are defined below.
        :param pulumi.Input[str] protocol: Protocol for the listener. Supported values are `HTTP` or `HTTPS`
        :param pulumi.Input[str] name: Name of the listener. A listener name must be unique within a service. Valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
        :param pulumi.Input[int] port: Listener port. You can specify a value from 1 to 65535. If `port` is not specified and `protocol` is HTTP, the value will default to 80. If `port` is not specified and `protocol` is HTTPS, the value will default to 443.
        :param pulumi.Input[str] service_arn: Amazon Resource Name (ARN) of the VPC Lattice service. You must include either the `service_arn` or `service_identifier` arguments.
        :param pulumi.Input[str] service_identifier: ID of the VPC Lattice service. You must include either the `service_arn` or `service_identifier` arguments.
               > **NOTE:** You must specify one of the following arguments: `service_arn` or `service_identifier`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "default_action", default_action)
        pulumi.set(__self__, "protocol", protocol)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if service_arn is not None:
            pulumi.set(__self__, "service_arn", service_arn)
        if service_identifier is not None:
            pulumi.set(__self__, "service_identifier", service_identifier)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> pulumi.Input['ListenerDefaultActionArgs']:
        """
        Default action block for the default listener rule. Default action blocks are defined below.
        """
        return pulumi.get(self, "default_action")

    @default_action.setter
    def default_action(self, value: pulumi.Input['ListenerDefaultActionArgs']):
        pulumi.set(self, "default_action", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        """
        Protocol for the listener. Supported values are `HTTP` or `HTTPS`
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the listener. A listener name must be unique within a service. Valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Listener port. You can specify a value from 1 to 65535. If `port` is not specified and `protocol` is HTTP, the value will default to 80. If `port` is not specified and `protocol` is HTTPS, the value will default to 443.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serviceArn")
    def service_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the VPC Lattice service. You must include either the `service_arn` or `service_identifier` arguments.
        """
        return pulumi.get(self, "service_arn")

    @service_arn.setter
    def service_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_arn", value)

    @property
    @pulumi.getter(name="serviceIdentifier")
    def service_identifier(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the VPC Lattice service. You must include either the `service_arn` or `service_identifier` arguments.
        > **NOTE:** You must specify one of the following arguments: `service_arn` or `service_identifier`.
        """
        return pulumi.get(self, "service_identifier")

    @service_identifier.setter
    def service_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_identifier", value)

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
class _ListenerState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 default_action: Optional[pulumi.Input['ListenerDefaultActionArgs']] = None,
                 last_updated_at: Optional[pulumi.Input[str]] = None,
                 listener_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 service_arn: Optional[pulumi.Input[str]] = None,
                 service_identifier: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Listener resources.
        :param pulumi.Input[str] arn: ARN of the listener.
        :param pulumi.Input[str] created_at: Date and time that the listener was created, specified in ISO-8601 format.
        :param pulumi.Input['ListenerDefaultActionArgs'] default_action: Default action block for the default listener rule. Default action blocks are defined below.
        :param pulumi.Input[str] listener_id: Standalone ID of the listener, e.g. `listener-0a1b2c3d4e5f6g`.
        :param pulumi.Input[str] name: Name of the listener. A listener name must be unique within a service. Valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
        :param pulumi.Input[int] port: Listener port. You can specify a value from 1 to 65535. If `port` is not specified and `protocol` is HTTP, the value will default to 80. If `port` is not specified and `protocol` is HTTPS, the value will default to 443.
        :param pulumi.Input[str] protocol: Protocol for the listener. Supported values are `HTTP` or `HTTPS`
        :param pulumi.Input[str] service_arn: Amazon Resource Name (ARN) of the VPC Lattice service. You must include either the `service_arn` or `service_identifier` arguments.
        :param pulumi.Input[str] service_identifier: ID of the VPC Lattice service. You must include either the `service_arn` or `service_identifier` arguments.
               > **NOTE:** You must specify one of the following arguments: `service_arn` or `service_identifier`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if default_action is not None:
            pulumi.set(__self__, "default_action", default_action)
        if last_updated_at is not None:
            pulumi.set(__self__, "last_updated_at", last_updated_at)
        if listener_id is not None:
            pulumi.set(__self__, "listener_id", listener_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if service_arn is not None:
            pulumi.set(__self__, "service_arn", service_arn)
        if service_identifier is not None:
            pulumi.set(__self__, "service_identifier", service_identifier)
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
        ARN of the listener.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date and time that the listener was created, specified in ISO-8601 format.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> Optional[pulumi.Input['ListenerDefaultActionArgs']]:
        """
        Default action block for the default listener rule. Default action blocks are defined below.
        """
        return pulumi.get(self, "default_action")

    @default_action.setter
    def default_action(self, value: Optional[pulumi.Input['ListenerDefaultActionArgs']]):
        pulumi.set(self, "default_action", value)

    @property
    @pulumi.getter(name="lastUpdatedAt")
    def last_updated_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "last_updated_at")

    @last_updated_at.setter
    def last_updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_updated_at", value)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> Optional[pulumi.Input[str]]:
        """
        Standalone ID of the listener, e.g. `listener-0a1b2c3d4e5f6g`.
        """
        return pulumi.get(self, "listener_id")

    @listener_id.setter
    def listener_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "listener_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the listener. A listener name must be unique within a service. Valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Listener port. You can specify a value from 1 to 65535. If `port` is not specified and `protocol` is HTTP, the value will default to 80. If `port` is not specified and `protocol` is HTTPS, the value will default to 443.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Protocol for the listener. Supported values are `HTTP` or `HTTPS`
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="serviceArn")
    def service_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the VPC Lattice service. You must include either the `service_arn` or `service_identifier` arguments.
        """
        return pulumi.get(self, "service_arn")

    @service_arn.setter
    def service_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_arn", value)

    @property
    @pulumi.getter(name="serviceIdentifier")
    def service_identifier(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the VPC Lattice service. You must include either the `service_arn` or `service_identifier` arguments.
        > **NOTE:** You must specify one of the following arguments: `service_arn` or `service_identifier`.
        """
        return pulumi.get(self, "service_identifier")

    @service_identifier.setter
    def service_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_identifier", value)

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
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class Listener(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_action: Optional[pulumi.Input[pulumi.InputType['ListenerDefaultActionArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 service_arn: Optional[pulumi.Input[str]] = None,
                 service_identifier: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Resource for managing an AWS VPC Lattice Listener.

        ## Example Usage
        ### Forward action

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.vpclattice.Service("test")
        example_target_group = aws.vpclattice.TargetGroup("exampleTargetGroup",
            type="INSTANCE",
            config=aws.vpclattice.TargetGroupConfigArgs(
                port=80,
                protocol="HTTP",
                vpc_identifier=aws_vpc["test"]["id"],
            ))
        example_listener = aws.vpclattice.Listener("exampleListener",
            protocol="HTTP",
            service_identifier=aws_vpclattice_service["example"]["id"],
            default_action=aws.vpclattice.ListenerDefaultActionArgs(
                forwards=[aws.vpclattice.ListenerDefaultActionForwardArgs(
                    target_groups=[aws.vpclattice.ListenerDefaultActionForwardTargetGroupArgs(
                        target_group_identifier=example_target_group.id,
                    )],
                )],
            ))
        ```
        ### Forward action with weighted target groups

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.vpclattice.Service("test")
        example1 = aws.vpclattice.TargetGroup("example1",
            type="INSTANCE",
            config=aws.vpclattice.TargetGroupConfigArgs(
                port=80,
                protocol="HTTP",
                vpc_identifier=aws_vpc["test"]["id"],
            ))
        example2 = aws.vpclattice.TargetGroup("example2",
            type="INSTANCE",
            config=aws.vpclattice.TargetGroupConfigArgs(
                port=8080,
                protocol="HTTP",
                vpc_identifier=aws_vpc["test"]["id"],
            ))
        example = aws.vpclattice.Listener("example",
            protocol="HTTP",
            service_identifier=aws_vpclattice_service["example"]["id"],
            default_action=aws.vpclattice.ListenerDefaultActionArgs(
                forwards=[aws.vpclattice.ListenerDefaultActionForwardArgs(
                    target_groups=[
                        aws.vpclattice.ListenerDefaultActionForwardTargetGroupArgs(
                            target_group_identifier=example1.id,
                            weight=80,
                        ),
                        aws.vpclattice.ListenerDefaultActionForwardTargetGroupArgs(
                            target_group_identifier=example2.id,
                            weight=20,
                        ),
                    ],
                )],
            ))
        ```

        ## Import

        Using `pulumi import`, import VPC Lattice Listener using the `listener_id` of the listener and the `id` of the VPC Lattice service combined with a `/` character. For example:

        ```sh
         $ pulumi import aws:vpclattice/listener:Listener example svc-1a2b3c4d/listener-987654321
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ListenerDefaultActionArgs']] default_action: Default action block for the default listener rule. Default action blocks are defined below.
        :param pulumi.Input[str] name: Name of the listener. A listener name must be unique within a service. Valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
        :param pulumi.Input[int] port: Listener port. You can specify a value from 1 to 65535. If `port` is not specified and `protocol` is HTTP, the value will default to 80. If `port` is not specified and `protocol` is HTTPS, the value will default to 443.
        :param pulumi.Input[str] protocol: Protocol for the listener. Supported values are `HTTP` or `HTTPS`
        :param pulumi.Input[str] service_arn: Amazon Resource Name (ARN) of the VPC Lattice service. You must include either the `service_arn` or `service_identifier` arguments.
        :param pulumi.Input[str] service_identifier: ID of the VPC Lattice service. You must include either the `service_arn` or `service_identifier` arguments.
               > **NOTE:** You must specify one of the following arguments: `service_arn` or `service_identifier`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ListenerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for managing an AWS VPC Lattice Listener.

        ## Example Usage
        ### Forward action

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.vpclattice.Service("test")
        example_target_group = aws.vpclattice.TargetGroup("exampleTargetGroup",
            type="INSTANCE",
            config=aws.vpclattice.TargetGroupConfigArgs(
                port=80,
                protocol="HTTP",
                vpc_identifier=aws_vpc["test"]["id"],
            ))
        example_listener = aws.vpclattice.Listener("exampleListener",
            protocol="HTTP",
            service_identifier=aws_vpclattice_service["example"]["id"],
            default_action=aws.vpclattice.ListenerDefaultActionArgs(
                forwards=[aws.vpclattice.ListenerDefaultActionForwardArgs(
                    target_groups=[aws.vpclattice.ListenerDefaultActionForwardTargetGroupArgs(
                        target_group_identifier=example_target_group.id,
                    )],
                )],
            ))
        ```
        ### Forward action with weighted target groups

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.vpclattice.Service("test")
        example1 = aws.vpclattice.TargetGroup("example1",
            type="INSTANCE",
            config=aws.vpclattice.TargetGroupConfigArgs(
                port=80,
                protocol="HTTP",
                vpc_identifier=aws_vpc["test"]["id"],
            ))
        example2 = aws.vpclattice.TargetGroup("example2",
            type="INSTANCE",
            config=aws.vpclattice.TargetGroupConfigArgs(
                port=8080,
                protocol="HTTP",
                vpc_identifier=aws_vpc["test"]["id"],
            ))
        example = aws.vpclattice.Listener("example",
            protocol="HTTP",
            service_identifier=aws_vpclattice_service["example"]["id"],
            default_action=aws.vpclattice.ListenerDefaultActionArgs(
                forwards=[aws.vpclattice.ListenerDefaultActionForwardArgs(
                    target_groups=[
                        aws.vpclattice.ListenerDefaultActionForwardTargetGroupArgs(
                            target_group_identifier=example1.id,
                            weight=80,
                        ),
                        aws.vpclattice.ListenerDefaultActionForwardTargetGroupArgs(
                            target_group_identifier=example2.id,
                            weight=20,
                        ),
                    ],
                )],
            ))
        ```

        ## Import

        Using `pulumi import`, import VPC Lattice Listener using the `listener_id` of the listener and the `id` of the VPC Lattice service combined with a `/` character. For example:

        ```sh
         $ pulumi import aws:vpclattice/listener:Listener example svc-1a2b3c4d/listener-987654321
        ```

        :param str resource_name: The name of the resource.
        :param ListenerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ListenerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_action: Optional[pulumi.Input[pulumi.InputType['ListenerDefaultActionArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 service_arn: Optional[pulumi.Input[str]] = None,
                 service_identifier: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ListenerArgs.__new__(ListenerArgs)

            if default_action is None and not opts.urn:
                raise TypeError("Missing required property 'default_action'")
            __props__.__dict__["default_action"] = default_action
            __props__.__dict__["name"] = name
            __props__.__dict__["port"] = port
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["service_arn"] = service_arn
            __props__.__dict__["service_identifier"] = service_identifier
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["last_updated_at"] = None
            __props__.__dict__["listener_id"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Listener, __self__).__init__(
            'aws:vpclattice/listener:Listener',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            default_action: Optional[pulumi.Input[pulumi.InputType['ListenerDefaultActionArgs']]] = None,
            last_updated_at: Optional[pulumi.Input[str]] = None,
            listener_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            service_arn: Optional[pulumi.Input[str]] = None,
            service_identifier: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Listener':
        """
        Get an existing Listener resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: ARN of the listener.
        :param pulumi.Input[str] created_at: Date and time that the listener was created, specified in ISO-8601 format.
        :param pulumi.Input[pulumi.InputType['ListenerDefaultActionArgs']] default_action: Default action block for the default listener rule. Default action blocks are defined below.
        :param pulumi.Input[str] listener_id: Standalone ID of the listener, e.g. `listener-0a1b2c3d4e5f6g`.
        :param pulumi.Input[str] name: Name of the listener. A listener name must be unique within a service. Valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
        :param pulumi.Input[int] port: Listener port. You can specify a value from 1 to 65535. If `port` is not specified and `protocol` is HTTP, the value will default to 80. If `port` is not specified and `protocol` is HTTPS, the value will default to 443.
        :param pulumi.Input[str] protocol: Protocol for the listener. Supported values are `HTTP` or `HTTPS`
        :param pulumi.Input[str] service_arn: Amazon Resource Name (ARN) of the VPC Lattice service. You must include either the `service_arn` or `service_identifier` arguments.
        :param pulumi.Input[str] service_identifier: ID of the VPC Lattice service. You must include either the `service_arn` or `service_identifier` arguments.
               > **NOTE:** You must specify one of the following arguments: `service_arn` or `service_identifier`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ListenerState.__new__(_ListenerState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["default_action"] = default_action
        __props__.__dict__["last_updated_at"] = last_updated_at
        __props__.__dict__["listener_id"] = listener_id
        __props__.__dict__["name"] = name
        __props__.__dict__["port"] = port
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["service_arn"] = service_arn
        __props__.__dict__["service_identifier"] = service_identifier
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return Listener(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN of the listener.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Date and time that the listener was created, specified in ISO-8601 format.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> pulumi.Output['outputs.ListenerDefaultAction']:
        """
        Default action block for the default listener rule. Default action blocks are defined below.
        """
        return pulumi.get(self, "default_action")

    @property
    @pulumi.getter(name="lastUpdatedAt")
    def last_updated_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "last_updated_at")

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> pulumi.Output[str]:
        """
        Standalone ID of the listener, e.g. `listener-0a1b2c3d4e5f6g`.
        """
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the listener. A listener name must be unique within a service. Valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        Listener port. You can specify a value from 1 to 65535. If `port` is not specified and `protocol` is HTTP, the value will default to 80. If `port` is not specified and `protocol` is HTTPS, the value will default to 443.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        Protocol for the listener. Supported values are `HTTP` or `HTTPS`
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="serviceArn")
    def service_arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the VPC Lattice service. You must include either the `service_arn` or `service_identifier` arguments.
        """
        return pulumi.get(self, "service_arn")

    @property
    @pulumi.getter(name="serviceIdentifier")
    def service_identifier(self) -> pulumi.Output[str]:
        """
        ID of the VPC Lattice service. You must include either the `service_arn` or `service_identifier` arguments.
        > **NOTE:** You must specify one of the following arguments: `service_arn` or `service_identifier`.
        """
        return pulumi.get(self, "service_identifier")

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
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

