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

__all__ = ['ClusterArgs', 'Cluster']

@pulumi.input_type
class ClusterArgs:
    def __init__(__self__, *,
                 capacity_providers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 configuration: Optional[pulumi.Input['ClusterConfigurationArgs']] = None,
                 default_capacity_provider_strategies: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterDefaultCapacityProviderStrategyArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 service_connect_defaults: Optional[pulumi.Input['ClusterServiceConnectDefaultsArgs']] = None,
                 settings: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterSettingArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Cluster resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] capacity_providers: List of short names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
        :param pulumi.Input['ClusterConfigurationArgs'] configuration: The execute command configuration for the cluster. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input['ClusterDefaultCapacityProviderStrategyArgs']]] default_capacity_provider_strategies: Configuration block for capacity provider strategy to use by default for the cluster. Can be one or more. Detailed below.
        :param pulumi.Input[str] name: Name of the setting to manage. Valid values: `containerInsights`.
        :param pulumi.Input['ClusterServiceConnectDefaultsArgs'] service_connect_defaults: Configures a default Service Connect namespace. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input['ClusterSettingArgs']]] settings: Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Detailed below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        if capacity_providers is not None:
            warnings.warn("""Use the aws_ecs_cluster_capacity_providers resource instead""", DeprecationWarning)
            pulumi.log.warn("""capacity_providers is deprecated: Use the aws_ecs_cluster_capacity_providers resource instead""")
        if capacity_providers is not None:
            pulumi.set(__self__, "capacity_providers", capacity_providers)
        if configuration is not None:
            pulumi.set(__self__, "configuration", configuration)
        if default_capacity_provider_strategies is not None:
            warnings.warn("""Use the aws_ecs_cluster_capacity_providers resource instead""", DeprecationWarning)
            pulumi.log.warn("""default_capacity_provider_strategies is deprecated: Use the aws_ecs_cluster_capacity_providers resource instead""")
        if default_capacity_provider_strategies is not None:
            pulumi.set(__self__, "default_capacity_provider_strategies", default_capacity_provider_strategies)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if service_connect_defaults is not None:
            pulumi.set(__self__, "service_connect_defaults", service_connect_defaults)
        if settings is not None:
            pulumi.set(__self__, "settings", settings)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="capacityProviders")
    def capacity_providers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of short names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
        """
        return pulumi.get(self, "capacity_providers")

    @capacity_providers.setter
    def capacity_providers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "capacity_providers", value)

    @property
    @pulumi.getter
    def configuration(self) -> Optional[pulumi.Input['ClusterConfigurationArgs']]:
        """
        The execute command configuration for the cluster. Detailed below.
        """
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: Optional[pulumi.Input['ClusterConfigurationArgs']]):
        pulumi.set(self, "configuration", value)

    @property
    @pulumi.getter(name="defaultCapacityProviderStrategies")
    def default_capacity_provider_strategies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterDefaultCapacityProviderStrategyArgs']]]]:
        """
        Configuration block for capacity provider strategy to use by default for the cluster. Can be one or more. Detailed below.
        """
        return pulumi.get(self, "default_capacity_provider_strategies")

    @default_capacity_provider_strategies.setter
    def default_capacity_provider_strategies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterDefaultCapacityProviderStrategyArgs']]]]):
        pulumi.set(self, "default_capacity_provider_strategies", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the setting to manage. Valid values: `containerInsights`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="serviceConnectDefaults")
    def service_connect_defaults(self) -> Optional[pulumi.Input['ClusterServiceConnectDefaultsArgs']]:
        """
        Configures a default Service Connect namespace. Detailed below.
        """
        return pulumi.get(self, "service_connect_defaults")

    @service_connect_defaults.setter
    def service_connect_defaults(self, value: Optional[pulumi.Input['ClusterServiceConnectDefaultsArgs']]):
        pulumi.set(self, "service_connect_defaults", value)

    @property
    @pulumi.getter
    def settings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterSettingArgs']]]]:
        """
        Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Detailed below.
        """
        return pulumi.get(self, "settings")

    @settings.setter
    def settings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterSettingArgs']]]]):
        pulumi.set(self, "settings", value)

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
class _ClusterState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 capacity_providers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 configuration: Optional[pulumi.Input['ClusterConfigurationArgs']] = None,
                 default_capacity_provider_strategies: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterDefaultCapacityProviderStrategyArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 service_connect_defaults: Optional[pulumi.Input['ClusterServiceConnectDefaultsArgs']] = None,
                 settings: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterSettingArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Cluster resources.
        :param pulumi.Input[str] arn: ARN that identifies the cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] capacity_providers: List of short names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
        :param pulumi.Input['ClusterConfigurationArgs'] configuration: The execute command configuration for the cluster. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input['ClusterDefaultCapacityProviderStrategyArgs']]] default_capacity_provider_strategies: Configuration block for capacity provider strategy to use by default for the cluster. Can be one or more. Detailed below.
        :param pulumi.Input[str] name: Name of the setting to manage. Valid values: `containerInsights`.
        :param pulumi.Input['ClusterServiceConnectDefaultsArgs'] service_connect_defaults: Configures a default Service Connect namespace. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input['ClusterSettingArgs']]] settings: Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Detailed below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if capacity_providers is not None:
            warnings.warn("""Use the aws_ecs_cluster_capacity_providers resource instead""", DeprecationWarning)
            pulumi.log.warn("""capacity_providers is deprecated: Use the aws_ecs_cluster_capacity_providers resource instead""")
        if capacity_providers is not None:
            pulumi.set(__self__, "capacity_providers", capacity_providers)
        if configuration is not None:
            pulumi.set(__self__, "configuration", configuration)
        if default_capacity_provider_strategies is not None:
            warnings.warn("""Use the aws_ecs_cluster_capacity_providers resource instead""", DeprecationWarning)
            pulumi.log.warn("""default_capacity_provider_strategies is deprecated: Use the aws_ecs_cluster_capacity_providers resource instead""")
        if default_capacity_provider_strategies is not None:
            pulumi.set(__self__, "default_capacity_provider_strategies", default_capacity_provider_strategies)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if service_connect_defaults is not None:
            pulumi.set(__self__, "service_connect_defaults", service_connect_defaults)
        if settings is not None:
            pulumi.set(__self__, "settings", settings)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN that identifies the cluster.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="capacityProviders")
    def capacity_providers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of short names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
        """
        return pulumi.get(self, "capacity_providers")

    @capacity_providers.setter
    def capacity_providers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "capacity_providers", value)

    @property
    @pulumi.getter
    def configuration(self) -> Optional[pulumi.Input['ClusterConfigurationArgs']]:
        """
        The execute command configuration for the cluster. Detailed below.
        """
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: Optional[pulumi.Input['ClusterConfigurationArgs']]):
        pulumi.set(self, "configuration", value)

    @property
    @pulumi.getter(name="defaultCapacityProviderStrategies")
    def default_capacity_provider_strategies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterDefaultCapacityProviderStrategyArgs']]]]:
        """
        Configuration block for capacity provider strategy to use by default for the cluster. Can be one or more. Detailed below.
        """
        return pulumi.get(self, "default_capacity_provider_strategies")

    @default_capacity_provider_strategies.setter
    def default_capacity_provider_strategies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterDefaultCapacityProviderStrategyArgs']]]]):
        pulumi.set(self, "default_capacity_provider_strategies", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the setting to manage. Valid values: `containerInsights`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="serviceConnectDefaults")
    def service_connect_defaults(self) -> Optional[pulumi.Input['ClusterServiceConnectDefaultsArgs']]:
        """
        Configures a default Service Connect namespace. Detailed below.
        """
        return pulumi.get(self, "service_connect_defaults")

    @service_connect_defaults.setter
    def service_connect_defaults(self, value: Optional[pulumi.Input['ClusterServiceConnectDefaultsArgs']]):
        pulumi.set(self, "service_connect_defaults", value)

    @property
    @pulumi.getter
    def settings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterSettingArgs']]]]:
        """
        Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Detailed below.
        """
        return pulumi.get(self, "settings")

    @settings.setter
    def settings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterSettingArgs']]]]):
        pulumi.set(self, "settings", value)

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


class Cluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capacity_providers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 configuration: Optional[pulumi.Input[pulumi.InputType['ClusterConfigurationArgs']]] = None,
                 default_capacity_provider_strategies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterDefaultCapacityProviderStrategyArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 service_connect_defaults: Optional[pulumi.Input[pulumi.InputType['ClusterServiceConnectDefaultsArgs']]] = None,
                 settings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterSettingArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides an ECS cluster.

        > **NOTE on Clusters and Cluster Capacity Providers:** this provider provides both a standalone `ecs.ClusterCapacityProviders` resource, as well as allowing the capacity providers and default strategies to be managed in-line by the `ecs.Cluster` resource. You cannot use a Cluster with in-line capacity providers in conjunction with the Capacity Providers resource, nor use more than one Capacity Providers resource with a single Cluster, as doing so will cause a conflict and will lead to mutual overwrites.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        foo = aws.ecs.Cluster("foo", settings=[aws.ecs.ClusterSettingArgs(
            name="containerInsights",
            value="enabled",
        )])
        ```
        ### Example with Log Configuration

        ```python
        import pulumi
        import pulumi_aws as aws

        example_key = aws.kms.Key("exampleKey",
            description="example",
            deletion_window_in_days=7)
        example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup")
        test = aws.ecs.Cluster("test", configuration=aws.ecs.ClusterConfigurationArgs(
            execute_command_configuration=aws.ecs.ClusterConfigurationExecuteCommandConfigurationArgs(
                kms_key_id=example_key.arn,
                logging="OVERRIDE",
                log_configuration=aws.ecs.ClusterConfigurationExecuteCommandConfigurationLogConfigurationArgs(
                    cloud_watch_encryption_enabled=True,
                    cloud_watch_log_group_name=example_log_group.name,
                ),
            ),
        ))
        ```
        ### Example with Capacity Providers

        ```python
        import pulumi
        import pulumi_aws as aws

        example_cluster = aws.ecs.Cluster("exampleCluster")
        example_capacity_provider = aws.ecs.CapacityProvider("exampleCapacityProvider", auto_scaling_group_provider=aws.ecs.CapacityProviderAutoScalingGroupProviderArgs(
            auto_scaling_group_arn=aws_autoscaling_group["example"]["arn"],
        ))
        example_cluster_capacity_providers = aws.ecs.ClusterCapacityProviders("exampleClusterCapacityProviders",
            cluster_name=example_cluster.name,
            capacity_providers=[example_capacity_provider.name],
            default_capacity_provider_strategies=[aws.ecs.ClusterCapacityProvidersDefaultCapacityProviderStrategyArgs(
                base=1,
                weight=100,
                capacity_provider=example_capacity_provider.name,
            )])
        ```

        ## Import

        ECS clusters can be imported using the `name`, e.g.,

        ```sh
         $ pulumi import aws:ecs/cluster:Cluster stateless stateless-app
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] capacity_providers: List of short names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
        :param pulumi.Input[pulumi.InputType['ClusterConfigurationArgs']] configuration: The execute command configuration for the cluster. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterDefaultCapacityProviderStrategyArgs']]]] default_capacity_provider_strategies: Configuration block for capacity provider strategy to use by default for the cluster. Can be one or more. Detailed below.
        :param pulumi.Input[str] name: Name of the setting to manage. Valid values: `containerInsights`.
        :param pulumi.Input[pulumi.InputType['ClusterServiceConnectDefaultsArgs']] service_connect_defaults: Configures a default Service Connect namespace. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterSettingArgs']]]] settings: Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Detailed below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ClusterArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an ECS cluster.

        > **NOTE on Clusters and Cluster Capacity Providers:** this provider provides both a standalone `ecs.ClusterCapacityProviders` resource, as well as allowing the capacity providers and default strategies to be managed in-line by the `ecs.Cluster` resource. You cannot use a Cluster with in-line capacity providers in conjunction with the Capacity Providers resource, nor use more than one Capacity Providers resource with a single Cluster, as doing so will cause a conflict and will lead to mutual overwrites.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        foo = aws.ecs.Cluster("foo", settings=[aws.ecs.ClusterSettingArgs(
            name="containerInsights",
            value="enabled",
        )])
        ```
        ### Example with Log Configuration

        ```python
        import pulumi
        import pulumi_aws as aws

        example_key = aws.kms.Key("exampleKey",
            description="example",
            deletion_window_in_days=7)
        example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup")
        test = aws.ecs.Cluster("test", configuration=aws.ecs.ClusterConfigurationArgs(
            execute_command_configuration=aws.ecs.ClusterConfigurationExecuteCommandConfigurationArgs(
                kms_key_id=example_key.arn,
                logging="OVERRIDE",
                log_configuration=aws.ecs.ClusterConfigurationExecuteCommandConfigurationLogConfigurationArgs(
                    cloud_watch_encryption_enabled=True,
                    cloud_watch_log_group_name=example_log_group.name,
                ),
            ),
        ))
        ```
        ### Example with Capacity Providers

        ```python
        import pulumi
        import pulumi_aws as aws

        example_cluster = aws.ecs.Cluster("exampleCluster")
        example_capacity_provider = aws.ecs.CapacityProvider("exampleCapacityProvider", auto_scaling_group_provider=aws.ecs.CapacityProviderAutoScalingGroupProviderArgs(
            auto_scaling_group_arn=aws_autoscaling_group["example"]["arn"],
        ))
        example_cluster_capacity_providers = aws.ecs.ClusterCapacityProviders("exampleClusterCapacityProviders",
            cluster_name=example_cluster.name,
            capacity_providers=[example_capacity_provider.name],
            default_capacity_provider_strategies=[aws.ecs.ClusterCapacityProvidersDefaultCapacityProviderStrategyArgs(
                base=1,
                weight=100,
                capacity_provider=example_capacity_provider.name,
            )])
        ```

        ## Import

        ECS clusters can be imported using the `name`, e.g.,

        ```sh
         $ pulumi import aws:ecs/cluster:Cluster stateless stateless-app
        ```

        :param str resource_name: The name of the resource.
        :param ClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capacity_providers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 configuration: Optional[pulumi.Input[pulumi.InputType['ClusterConfigurationArgs']]] = None,
                 default_capacity_provider_strategies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterDefaultCapacityProviderStrategyArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 service_connect_defaults: Optional[pulumi.Input[pulumi.InputType['ClusterServiceConnectDefaultsArgs']]] = None,
                 settings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterSettingArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClusterArgs.__new__(ClusterArgs)

            if capacity_providers is not None and not opts.urn:
                warnings.warn("""Use the aws_ecs_cluster_capacity_providers resource instead""", DeprecationWarning)
                pulumi.log.warn("""capacity_providers is deprecated: Use the aws_ecs_cluster_capacity_providers resource instead""")
            __props__.__dict__["capacity_providers"] = capacity_providers
            __props__.__dict__["configuration"] = configuration
            if default_capacity_provider_strategies is not None and not opts.urn:
                warnings.warn("""Use the aws_ecs_cluster_capacity_providers resource instead""", DeprecationWarning)
                pulumi.log.warn("""default_capacity_provider_strategies is deprecated: Use the aws_ecs_cluster_capacity_providers resource instead""")
            __props__.__dict__["default_capacity_provider_strategies"] = default_capacity_provider_strategies
            __props__.__dict__["name"] = name
            __props__.__dict__["service_connect_defaults"] = service_connect_defaults
            __props__.__dict__["settings"] = settings
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        super(Cluster, __self__).__init__(
            'aws:ecs/cluster:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            capacity_providers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            configuration: Optional[pulumi.Input[pulumi.InputType['ClusterConfigurationArgs']]] = None,
            default_capacity_provider_strategies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterDefaultCapacityProviderStrategyArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            service_connect_defaults: Optional[pulumi.Input[pulumi.InputType['ClusterServiceConnectDefaultsArgs']]] = None,
            settings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterSettingArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Cluster':
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: ARN that identifies the cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] capacity_providers: List of short names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
        :param pulumi.Input[pulumi.InputType['ClusterConfigurationArgs']] configuration: The execute command configuration for the cluster. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterDefaultCapacityProviderStrategyArgs']]]] default_capacity_provider_strategies: Configuration block for capacity provider strategy to use by default for the cluster. Can be one or more. Detailed below.
        :param pulumi.Input[str] name: Name of the setting to manage. Valid values: `containerInsights`.
        :param pulumi.Input[pulumi.InputType['ClusterServiceConnectDefaultsArgs']] service_connect_defaults: Configures a default Service Connect namespace. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterSettingArgs']]]] settings: Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Detailed below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClusterState.__new__(_ClusterState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["capacity_providers"] = capacity_providers
        __props__.__dict__["configuration"] = configuration
        __props__.__dict__["default_capacity_provider_strategies"] = default_capacity_provider_strategies
        __props__.__dict__["name"] = name
        __props__.__dict__["service_connect_defaults"] = service_connect_defaults
        __props__.__dict__["settings"] = settings
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return Cluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN that identifies the cluster.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="capacityProviders")
    def capacity_providers(self) -> pulumi.Output[Sequence[str]]:
        """
        List of short names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
        """
        return pulumi.get(self, "capacity_providers")

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Output[Optional['outputs.ClusterConfiguration']]:
        """
        The execute command configuration for the cluster. Detailed below.
        """
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter(name="defaultCapacityProviderStrategies")
    def default_capacity_provider_strategies(self) -> pulumi.Output[Sequence['outputs.ClusterDefaultCapacityProviderStrategy']]:
        """
        Configuration block for capacity provider strategy to use by default for the cluster. Can be one or more. Detailed below.
        """
        return pulumi.get(self, "default_capacity_provider_strategies")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the setting to manage. Valid values: `containerInsights`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceConnectDefaults")
    def service_connect_defaults(self) -> pulumi.Output[Optional['outputs.ClusterServiceConnectDefaults']]:
        """
        Configures a default Service Connect namespace. Detailed below.
        """
        return pulumi.get(self, "service_connect_defaults")

    @property
    @pulumi.getter
    def settings(self) -> pulumi.Output[Sequence['outputs.ClusterSetting']]:
        """
        Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Detailed below.
        """
        return pulumi.get(self, "settings")

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

