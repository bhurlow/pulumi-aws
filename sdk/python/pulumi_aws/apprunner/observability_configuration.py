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

__all__ = ['ObservabilityConfigurationArgs', 'ObservabilityConfiguration']

@pulumi.input_type
class ObservabilityConfigurationArgs:
    def __init__(__self__, *,
                 observability_configuration_name: pulumi.Input[str],
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 trace_configuration: Optional[pulumi.Input['ObservabilityConfigurationTraceConfigurationArgs']] = None):
        """
        The set of arguments for constructing a ObservabilityConfiguration resource.
        :param pulumi.Input[str] observability_configuration_name: Name of the observability configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input['ObservabilityConfigurationTraceConfigurationArgs'] trace_configuration: Configuration of the tracing feature within this observability configuration. If you don't specify it, App Runner doesn't enable tracing. See Trace Configuration below for more details.
        """
        pulumi.set(__self__, "observability_configuration_name", observability_configuration_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if trace_configuration is not None:
            pulumi.set(__self__, "trace_configuration", trace_configuration)

    @property
    @pulumi.getter(name="observabilityConfigurationName")
    def observability_configuration_name(self) -> pulumi.Input[str]:
        """
        Name of the observability configuration.
        """
        return pulumi.get(self, "observability_configuration_name")

    @observability_configuration_name.setter
    def observability_configuration_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "observability_configuration_name", value)

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
    @pulumi.getter(name="traceConfiguration")
    def trace_configuration(self) -> Optional[pulumi.Input['ObservabilityConfigurationTraceConfigurationArgs']]:
        """
        Configuration of the tracing feature within this observability configuration. If you don't specify it, App Runner doesn't enable tracing. See Trace Configuration below for more details.
        """
        return pulumi.get(self, "trace_configuration")

    @trace_configuration.setter
    def trace_configuration(self, value: Optional[pulumi.Input['ObservabilityConfigurationTraceConfigurationArgs']]):
        pulumi.set(self, "trace_configuration", value)


@pulumi.input_type
class _ObservabilityConfigurationState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 latest: Optional[pulumi.Input[bool]] = None,
                 observability_configuration_name: Optional[pulumi.Input[str]] = None,
                 observability_configuration_revision: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 trace_configuration: Optional[pulumi.Input['ObservabilityConfigurationTraceConfigurationArgs']] = None):
        """
        Input properties used for looking up and filtering ObservabilityConfiguration resources.
        :param pulumi.Input[str] arn: ARN of this observability configuration.
        :param pulumi.Input[bool] latest: Whether the observability configuration has the highest `observability_configuration_revision` among all configurations that share the same `observability_configuration_name`.
        :param pulumi.Input[str] observability_configuration_name: Name of the observability configuration.
        :param pulumi.Input[int] observability_configuration_revision: The revision of this observability configuration.
        :param pulumi.Input[str] status: Current state of the observability configuration. An INACTIVE configuration revision has been deleted and can't be used. It is permanently removed some time after deletion.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input['ObservabilityConfigurationTraceConfigurationArgs'] trace_configuration: Configuration of the tracing feature within this observability configuration. If you don't specify it, App Runner doesn't enable tracing. See Trace Configuration below for more details.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if latest is not None:
            pulumi.set(__self__, "latest", latest)
        if observability_configuration_name is not None:
            pulumi.set(__self__, "observability_configuration_name", observability_configuration_name)
        if observability_configuration_revision is not None:
            pulumi.set(__self__, "observability_configuration_revision", observability_configuration_revision)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if trace_configuration is not None:
            pulumi.set(__self__, "trace_configuration", trace_configuration)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of this observability configuration.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def latest(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the observability configuration has the highest `observability_configuration_revision` among all configurations that share the same `observability_configuration_name`.
        """
        return pulumi.get(self, "latest")

    @latest.setter
    def latest(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "latest", value)

    @property
    @pulumi.getter(name="observabilityConfigurationName")
    def observability_configuration_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the observability configuration.
        """
        return pulumi.get(self, "observability_configuration_name")

    @observability_configuration_name.setter
    def observability_configuration_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "observability_configuration_name", value)

    @property
    @pulumi.getter(name="observabilityConfigurationRevision")
    def observability_configuration_revision(self) -> Optional[pulumi.Input[int]]:
        """
        The revision of this observability configuration.
        """
        return pulumi.get(self, "observability_configuration_revision")

    @observability_configuration_revision.setter
    def observability_configuration_revision(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "observability_configuration_revision", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Current state of the observability configuration. An INACTIVE configuration revision has been deleted and can't be used. It is permanently removed some time after deletion.
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
        Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter(name="traceConfiguration")
    def trace_configuration(self) -> Optional[pulumi.Input['ObservabilityConfigurationTraceConfigurationArgs']]:
        """
        Configuration of the tracing feature within this observability configuration. If you don't specify it, App Runner doesn't enable tracing. See Trace Configuration below for more details.
        """
        return pulumi.get(self, "trace_configuration")

    @trace_configuration.setter
    def trace_configuration(self, value: Optional[pulumi.Input['ObservabilityConfigurationTraceConfigurationArgs']]):
        pulumi.set(self, "trace_configuration", value)


class ObservabilityConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 observability_configuration_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 trace_configuration: Optional[pulumi.Input[pulumi.InputType['ObservabilityConfigurationTraceConfigurationArgs']]] = None,
                 __props__=None):
        """
        Manages an App Runner Observability Configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.apprunner.ObservabilityConfiguration("example",
            observability_configuration_name="example",
            tags={
                "Name": "example-apprunner-observability-configuration",
            },
            trace_configuration=aws.apprunner.ObservabilityConfigurationTraceConfigurationArgs(
                vendor="AWSXRAY",
            ))
        ```

        ## Import

        terraform import {

         to = aws_apprunner_observability_configuration.example

         id = "arn:aws:apprunner:us-east-1:1234567890:observabilityconfiguration/example/1/d75bc7ea55b71e724fe5c23452fe22a1" } Using `pulumi import`, import App Runner Observability Configuration using the `arn`. For exampleconsole % pulumi import aws_apprunner_observability_configuration.example arn:aws:apprunner:us-east-1:1234567890:observabilityconfiguration/example/1/d75bc7ea55b71e724fe5c23452fe22a1

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] observability_configuration_name: Name of the observability configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[pulumi.InputType['ObservabilityConfigurationTraceConfigurationArgs']] trace_configuration: Configuration of the tracing feature within this observability configuration. If you don't specify it, App Runner doesn't enable tracing. See Trace Configuration below for more details.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ObservabilityConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an App Runner Observability Configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.apprunner.ObservabilityConfiguration("example",
            observability_configuration_name="example",
            tags={
                "Name": "example-apprunner-observability-configuration",
            },
            trace_configuration=aws.apprunner.ObservabilityConfigurationTraceConfigurationArgs(
                vendor="AWSXRAY",
            ))
        ```

        ## Import

        terraform import {

         to = aws_apprunner_observability_configuration.example

         id = "arn:aws:apprunner:us-east-1:1234567890:observabilityconfiguration/example/1/d75bc7ea55b71e724fe5c23452fe22a1" } Using `pulumi import`, import App Runner Observability Configuration using the `arn`. For exampleconsole % pulumi import aws_apprunner_observability_configuration.example arn:aws:apprunner:us-east-1:1234567890:observabilityconfiguration/example/1/d75bc7ea55b71e724fe5c23452fe22a1

        :param str resource_name: The name of the resource.
        :param ObservabilityConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ObservabilityConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 observability_configuration_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 trace_configuration: Optional[pulumi.Input[pulumi.InputType['ObservabilityConfigurationTraceConfigurationArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ObservabilityConfigurationArgs.__new__(ObservabilityConfigurationArgs)

            if observability_configuration_name is None and not opts.urn:
                raise TypeError("Missing required property 'observability_configuration_name'")
            __props__.__dict__["observability_configuration_name"] = observability_configuration_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["trace_configuration"] = trace_configuration
            __props__.__dict__["arn"] = None
            __props__.__dict__["latest"] = None
            __props__.__dict__["observability_configuration_revision"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["tags_all"] = None
        super(ObservabilityConfiguration, __self__).__init__(
            'aws:apprunner/observabilityConfiguration:ObservabilityConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            latest: Optional[pulumi.Input[bool]] = None,
            observability_configuration_name: Optional[pulumi.Input[str]] = None,
            observability_configuration_revision: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            trace_configuration: Optional[pulumi.Input[pulumi.InputType['ObservabilityConfigurationTraceConfigurationArgs']]] = None) -> 'ObservabilityConfiguration':
        """
        Get an existing ObservabilityConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: ARN of this observability configuration.
        :param pulumi.Input[bool] latest: Whether the observability configuration has the highest `observability_configuration_revision` among all configurations that share the same `observability_configuration_name`.
        :param pulumi.Input[str] observability_configuration_name: Name of the observability configuration.
        :param pulumi.Input[int] observability_configuration_revision: The revision of this observability configuration.
        :param pulumi.Input[str] status: Current state of the observability configuration. An INACTIVE configuration revision has been deleted and can't be used. It is permanently removed some time after deletion.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[pulumi.InputType['ObservabilityConfigurationTraceConfigurationArgs']] trace_configuration: Configuration of the tracing feature within this observability configuration. If you don't specify it, App Runner doesn't enable tracing. See Trace Configuration below for more details.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ObservabilityConfigurationState.__new__(_ObservabilityConfigurationState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["latest"] = latest
        __props__.__dict__["observability_configuration_name"] = observability_configuration_name
        __props__.__dict__["observability_configuration_revision"] = observability_configuration_revision
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["trace_configuration"] = trace_configuration
        return ObservabilityConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN of this observability configuration.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def latest(self) -> pulumi.Output[bool]:
        """
        Whether the observability configuration has the highest `observability_configuration_revision` among all configurations that share the same `observability_configuration_name`.
        """
        return pulumi.get(self, "latest")

    @property
    @pulumi.getter(name="observabilityConfigurationName")
    def observability_configuration_name(self) -> pulumi.Output[str]:
        """
        Name of the observability configuration.
        """
        return pulumi.get(self, "observability_configuration_name")

    @property
    @pulumi.getter(name="observabilityConfigurationRevision")
    def observability_configuration_revision(self) -> pulumi.Output[int]:
        """
        The revision of this observability configuration.
        """
        return pulumi.get(self, "observability_configuration_revision")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Current state of the observability configuration. An INACTIVE configuration revision has been deleted and can't be used. It is permanently removed some time after deletion.
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
        Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter(name="traceConfiguration")
    def trace_configuration(self) -> pulumi.Output[Optional['outputs.ObservabilityConfigurationTraceConfiguration']]:
        """
        Configuration of the tracing feature within this observability configuration. If you don't specify it, App Runner doesn't enable tracing. See Trace Configuration below for more details.
        """
        return pulumi.get(self, "trace_configuration")

