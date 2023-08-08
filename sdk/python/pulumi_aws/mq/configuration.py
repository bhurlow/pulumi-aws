# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ConfigurationArgs', 'Configuration']

@pulumi.input_type
class ConfigurationArgs:
    def __init__(__self__, *,
                 data: pulumi.Input[str],
                 engine_type: pulumi.Input[str],
                 engine_version: pulumi.Input[str],
                 authentication_strategy: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Configuration resource.
        :param pulumi.Input[str] data: Broker configuration in XML format. See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html) for supported parameters and format of the XML.
        :param pulumi.Input[str] engine_type: Type of broker engine. Valid values are `ActiveMQ` and `RabbitMQ`.
        :param pulumi.Input[str] engine_version: Version of the broker engine.
        :param pulumi.Input[str] authentication_strategy: Authentication strategy associated with the configuration. Valid values are `simple` and `ldap`. `ldap` is not supported for `engine_type` `RabbitMQ`.
        :param pulumi.Input[str] description: Description of the configuration.
        :param pulumi.Input[str] name: Name of the configuration.
               
               The following arguments are optional:
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "data", data)
        pulumi.set(__self__, "engine_type", engine_type)
        pulumi.set(__self__, "engine_version", engine_version)
        if authentication_strategy is not None:
            pulumi.set(__self__, "authentication_strategy", authentication_strategy)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def data(self) -> pulumi.Input[str]:
        """
        Broker configuration in XML format. See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html) for supported parameters and format of the XML.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: pulumi.Input[str]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter(name="engineType")
    def engine_type(self) -> pulumi.Input[str]:
        """
        Type of broker engine. Valid values are `ActiveMQ` and `RabbitMQ`.
        """
        return pulumi.get(self, "engine_type")

    @engine_type.setter
    def engine_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "engine_type", value)

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> pulumi.Input[str]:
        """
        Version of the broker engine.
        """
        return pulumi.get(self, "engine_version")

    @engine_version.setter
    def engine_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "engine_version", value)

    @property
    @pulumi.getter(name="authenticationStrategy")
    def authentication_strategy(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication strategy associated with the configuration. Valid values are `simple` and `ldap`. `ldap` is not supported for `engine_type` `RabbitMQ`.
        """
        return pulumi.get(self, "authentication_strategy")

    @authentication_strategy.setter
    def authentication_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authentication_strategy", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the configuration.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the configuration.

        The following arguments are optional:
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ConfigurationState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 authentication_strategy: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine_type: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 latest_revision: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Configuration resources.
        :param pulumi.Input[str] arn: ARN of the configuration.
        :param pulumi.Input[str] authentication_strategy: Authentication strategy associated with the configuration. Valid values are `simple` and `ldap`. `ldap` is not supported for `engine_type` `RabbitMQ`.
        :param pulumi.Input[str] data: Broker configuration in XML format. See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html) for supported parameters and format of the XML.
        :param pulumi.Input[str] description: Description of the configuration.
        :param pulumi.Input[str] engine_type: Type of broker engine. Valid values are `ActiveMQ` and `RabbitMQ`.
        :param pulumi.Input[str] engine_version: Version of the broker engine.
        :param pulumi.Input[int] latest_revision: Latest revision of the configuration.
        :param pulumi.Input[str] name: Name of the configuration.
               
               The following arguments are optional:
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if authentication_strategy is not None:
            pulumi.set(__self__, "authentication_strategy", authentication_strategy)
        if data is not None:
            pulumi.set(__self__, "data", data)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if engine_type is not None:
            pulumi.set(__self__, "engine_type", engine_type)
        if engine_version is not None:
            pulumi.set(__self__, "engine_version", engine_version)
        if latest_revision is not None:
            pulumi.set(__self__, "latest_revision", latest_revision)
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
        ARN of the configuration.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="authenticationStrategy")
    def authentication_strategy(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication strategy associated with the configuration. Valid values are `simple` and `ldap`. `ldap` is not supported for `engine_type` `RabbitMQ`.
        """
        return pulumi.get(self, "authentication_strategy")

    @authentication_strategy.setter
    def authentication_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authentication_strategy", value)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        """
        Broker configuration in XML format. See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html) for supported parameters and format of the XML.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the configuration.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="engineType")
    def engine_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of broker engine. Valid values are `ActiveMQ` and `RabbitMQ`.
        """
        return pulumi.get(self, "engine_type")

    @engine_type.setter
    def engine_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_type", value)

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> Optional[pulumi.Input[str]]:
        """
        Version of the broker engine.
        """
        return pulumi.get(self, "engine_version")

    @engine_version.setter
    def engine_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_version", value)

    @property
    @pulumi.getter(name="latestRevision")
    def latest_revision(self) -> Optional[pulumi.Input[int]]:
        """
        Latest revision of the configuration.
        """
        return pulumi.get(self, "latest_revision")

    @latest_revision.setter
    def latest_revision(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "latest_revision", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the configuration.

        The following arguments are optional:
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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


class Configuration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication_strategy: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine_type: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides an MQ Configuration Resource.

        For more information on Amazon MQ, see [Amazon MQ documentation](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/welcome.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.mq.Configuration("example",
            data=\"\"\"<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
        <broker xmlns="http://activemq.apache.org/schema/core">
          <plugins>
            <forcePersistencyModeBrokerPlugin persistenceFlag="true"/>
            <statisticsBrokerPlugin/>
            <timeStampingBrokerPlugin ttlCeiling="86400000" zeroExpirationOverride="86400000"/>
          </plugins>
        </broker>

        \"\"\",
            description="Example Configuration",
            engine_type="ActiveMQ",
            engine_version="5.15.0")
        ```

        ## Import

        terraform import {

         to = aws_mq_configuration.example

         id = "c-0187d1eb-88c8-475a-9b79-16ef5a10c94f" } Using `pulumi import`, import MQ Configurations using the configuration ID. For exampleconsole % pulumi import aws_mq_configuration.example c-0187d1eb-88c8-475a-9b79-16ef5a10c94f

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authentication_strategy: Authentication strategy associated with the configuration. Valid values are `simple` and `ldap`. `ldap` is not supported for `engine_type` `RabbitMQ`.
        :param pulumi.Input[str] data: Broker configuration in XML format. See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html) for supported parameters and format of the XML.
        :param pulumi.Input[str] description: Description of the configuration.
        :param pulumi.Input[str] engine_type: Type of broker engine. Valid values are `ActiveMQ` and `RabbitMQ`.
        :param pulumi.Input[str] engine_version: Version of the broker engine.
        :param pulumi.Input[str] name: Name of the configuration.
               
               The following arguments are optional:
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an MQ Configuration Resource.

        For more information on Amazon MQ, see [Amazon MQ documentation](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/welcome.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.mq.Configuration("example",
            data=\"\"\"<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
        <broker xmlns="http://activemq.apache.org/schema/core">
          <plugins>
            <forcePersistencyModeBrokerPlugin persistenceFlag="true"/>
            <statisticsBrokerPlugin/>
            <timeStampingBrokerPlugin ttlCeiling="86400000" zeroExpirationOverride="86400000"/>
          </plugins>
        </broker>

        \"\"\",
            description="Example Configuration",
            engine_type="ActiveMQ",
            engine_version="5.15.0")
        ```

        ## Import

        terraform import {

         to = aws_mq_configuration.example

         id = "c-0187d1eb-88c8-475a-9b79-16ef5a10c94f" } Using `pulumi import`, import MQ Configurations using the configuration ID. For exampleconsole % pulumi import aws_mq_configuration.example c-0187d1eb-88c8-475a-9b79-16ef5a10c94f

        :param str resource_name: The name of the resource.
        :param ConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication_strategy: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine_type: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConfigurationArgs.__new__(ConfigurationArgs)

            __props__.__dict__["authentication_strategy"] = authentication_strategy
            if data is None and not opts.urn:
                raise TypeError("Missing required property 'data'")
            __props__.__dict__["data"] = data
            __props__.__dict__["description"] = description
            if engine_type is None and not opts.urn:
                raise TypeError("Missing required property 'engine_type'")
            __props__.__dict__["engine_type"] = engine_type
            if engine_version is None and not opts.urn:
                raise TypeError("Missing required property 'engine_version'")
            __props__.__dict__["engine_version"] = engine_version
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["latest_revision"] = None
            __props__.__dict__["tags_all"] = None
        super(Configuration, __self__).__init__(
            'aws:mq/configuration:Configuration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            authentication_strategy: Optional[pulumi.Input[str]] = None,
            data: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            engine_type: Optional[pulumi.Input[str]] = None,
            engine_version: Optional[pulumi.Input[str]] = None,
            latest_revision: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Configuration':
        """
        Get an existing Configuration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: ARN of the configuration.
        :param pulumi.Input[str] authentication_strategy: Authentication strategy associated with the configuration. Valid values are `simple` and `ldap`. `ldap` is not supported for `engine_type` `RabbitMQ`.
        :param pulumi.Input[str] data: Broker configuration in XML format. See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html) for supported parameters and format of the XML.
        :param pulumi.Input[str] description: Description of the configuration.
        :param pulumi.Input[str] engine_type: Type of broker engine. Valid values are `ActiveMQ` and `RabbitMQ`.
        :param pulumi.Input[str] engine_version: Version of the broker engine.
        :param pulumi.Input[int] latest_revision: Latest revision of the configuration.
        :param pulumi.Input[str] name: Name of the configuration.
               
               The following arguments are optional:
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConfigurationState.__new__(_ConfigurationState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["authentication_strategy"] = authentication_strategy
        __props__.__dict__["data"] = data
        __props__.__dict__["description"] = description
        __props__.__dict__["engine_type"] = engine_type
        __props__.__dict__["engine_version"] = engine_version
        __props__.__dict__["latest_revision"] = latest_revision
        __props__.__dict__["name"] = name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return Configuration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN of the configuration.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="authenticationStrategy")
    def authentication_strategy(self) -> pulumi.Output[str]:
        """
        Authentication strategy associated with the configuration. Valid values are `simple` and `ldap`. `ldap` is not supported for `engine_type` `RabbitMQ`.
        """
        return pulumi.get(self, "authentication_strategy")

    @property
    @pulumi.getter
    def data(self) -> pulumi.Output[str]:
        """
        Broker configuration in XML format. See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html) for supported parameters and format of the XML.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the configuration.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="engineType")
    def engine_type(self) -> pulumi.Output[str]:
        """
        Type of broker engine. Valid values are `ActiveMQ` and `RabbitMQ`.
        """
        return pulumi.get(self, "engine_type")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> pulumi.Output[str]:
        """
        Version of the broker engine.
        """
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="latestRevision")
    def latest_revision(self) -> pulumi.Output[int]:
        """
        Latest revision of the configuration.
        """
        return pulumi.get(self, "latest_revision")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the configuration.

        The following arguments are optional:
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

