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

__all__ = ['VirtualClusterArgs', 'VirtualCluster']

@pulumi.input_type
class VirtualClusterArgs:
    def __init__(__self__, *,
                 container_provider: pulumi.Input['VirtualClusterContainerProviderArgs'],
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a VirtualCluster resource.
        :param pulumi.Input['VirtualClusterContainerProviderArgs'] container_provider: Configuration block for the container provider associated with your cluster.
        :param pulumi.Input[str] name: Name of the virtual cluster.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "container_provider", container_provider)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="containerProvider")
    def container_provider(self) -> pulumi.Input['VirtualClusterContainerProviderArgs']:
        """
        Configuration block for the container provider associated with your cluster.
        """
        return pulumi.get(self, "container_provider")

    @container_provider.setter
    def container_provider(self, value: pulumi.Input['VirtualClusterContainerProviderArgs']):
        pulumi.set(self, "container_provider", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the virtual cluster.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _VirtualClusterState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 container_provider: Optional[pulumi.Input['VirtualClusterContainerProviderArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering VirtualCluster resources.
        :param pulumi.Input[str] arn: ARN of the cluster.
        :param pulumi.Input['VirtualClusterContainerProviderArgs'] container_provider: Configuration block for the container provider associated with your cluster.
        :param pulumi.Input[str] name: Name of the virtual cluster.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if container_provider is not None:
            pulumi.set(__self__, "container_provider", container_provider)
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
        ARN of the cluster.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="containerProvider")
    def container_provider(self) -> Optional[pulumi.Input['VirtualClusterContainerProviderArgs']]:
        """
        Configuration block for the container provider associated with your cluster.
        """
        return pulumi.get(self, "container_provider")

    @container_provider.setter
    def container_provider(self, value: Optional[pulumi.Input['VirtualClusterContainerProviderArgs']]):
        pulumi.set(self, "container_provider", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the virtual cluster.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class VirtualCluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_provider: Optional[pulumi.Input[pulumi.InputType['VirtualClusterContainerProviderArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manages an EMR Containers (EMR on EKS) Virtual Cluster.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.emrcontainers.VirtualCluster("example", container_provider=aws.emrcontainers.VirtualClusterContainerProviderArgs(
            id=aws_eks_cluster["example"]["name"],
            type="EKS",
            info=aws.emrcontainers.VirtualClusterContainerProviderInfoArgs(
                eks_info=aws.emrcontainers.VirtualClusterContainerProviderInfoEksInfoArgs(
                    namespace="default",
                ),
            ),
        ))
        ```

        ## Import

        In TODO v1.5.0 and later, use an `import` block to import EKS Clusters using the `id`. For exampleterraform import {

         to = aws_emrcontainers_virtual_cluster.example

         id = "a1b2c3d4e5f6g7h8i9j10k11l" } Using `TODO import`, import EKS Clusters using the `id`. For exampleconsole % TODO import aws_emrcontainers_virtual_cluster.example a1b2c3d4e5f6g7h8i9j10k11l

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['VirtualClusterContainerProviderArgs']] container_provider: Configuration block for the container provider associated with your cluster.
        :param pulumi.Input[str] name: Name of the virtual cluster.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VirtualClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an EMR Containers (EMR on EKS) Virtual Cluster.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.emrcontainers.VirtualCluster("example", container_provider=aws.emrcontainers.VirtualClusterContainerProviderArgs(
            id=aws_eks_cluster["example"]["name"],
            type="EKS",
            info=aws.emrcontainers.VirtualClusterContainerProviderInfoArgs(
                eks_info=aws.emrcontainers.VirtualClusterContainerProviderInfoEksInfoArgs(
                    namespace="default",
                ),
            ),
        ))
        ```

        ## Import

        In TODO v1.5.0 and later, use an `import` block to import EKS Clusters using the `id`. For exampleterraform import {

         to = aws_emrcontainers_virtual_cluster.example

         id = "a1b2c3d4e5f6g7h8i9j10k11l" } Using `TODO import`, import EKS Clusters using the `id`. For exampleconsole % TODO import aws_emrcontainers_virtual_cluster.example a1b2c3d4e5f6g7h8i9j10k11l

        :param str resource_name: The name of the resource.
        :param VirtualClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VirtualClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_provider: Optional[pulumi.Input[pulumi.InputType['VirtualClusterContainerProviderArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VirtualClusterArgs.__new__(VirtualClusterArgs)

            if container_provider is None and not opts.urn:
                raise TypeError("Missing required property 'container_provider'")
            __props__.__dict__["container_provider"] = container_provider
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(VirtualCluster, __self__).__init__(
            'aws:emrcontainers/virtualCluster:VirtualCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            container_provider: Optional[pulumi.Input[pulumi.InputType['VirtualClusterContainerProviderArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'VirtualCluster':
        """
        Get an existing VirtualCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: ARN of the cluster.
        :param pulumi.Input[pulumi.InputType['VirtualClusterContainerProviderArgs']] container_provider: Configuration block for the container provider associated with your cluster.
        :param pulumi.Input[str] name: Name of the virtual cluster.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VirtualClusterState.__new__(_VirtualClusterState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["container_provider"] = container_provider
        __props__.__dict__["name"] = name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return VirtualCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN of the cluster.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="containerProvider")
    def container_provider(self) -> pulumi.Output['outputs.VirtualClusterContainerProvider']:
        """
        Configuration block for the container provider associated with your cluster.
        """
        return pulumi.get(self, "container_provider")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the virtual cluster.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

