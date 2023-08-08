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

__all__ = ['ServerlessClusterArgs', 'ServerlessCluster']

@pulumi.input_type
class ServerlessClusterArgs:
    def __init__(__self__, *,
                 client_authentication: pulumi.Input['ServerlessClusterClientAuthenticationArgs'],
                 vpc_configs: pulumi.Input[Sequence[pulumi.Input['ServerlessClusterVpcConfigArgs']]],
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ServerlessCluster resource.
        :param pulumi.Input['ServerlessClusterClientAuthenticationArgs'] client_authentication: Specifies client authentication information for the serverless cluster. See below.
        :param pulumi.Input[Sequence[pulumi.Input['ServerlessClusterVpcConfigArgs']]] vpc_configs: VPC configuration information. See below.
        :param pulumi.Input[str] cluster_name: The name of the serverless cluster.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "client_authentication", client_authentication)
        pulumi.set(__self__, "vpc_configs", vpc_configs)
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="clientAuthentication")
    def client_authentication(self) -> pulumi.Input['ServerlessClusterClientAuthenticationArgs']:
        """
        Specifies client authentication information for the serverless cluster. See below.
        """
        return pulumi.get(self, "client_authentication")

    @client_authentication.setter
    def client_authentication(self, value: pulumi.Input['ServerlessClusterClientAuthenticationArgs']):
        pulumi.set(self, "client_authentication", value)

    @property
    @pulumi.getter(name="vpcConfigs")
    def vpc_configs(self) -> pulumi.Input[Sequence[pulumi.Input['ServerlessClusterVpcConfigArgs']]]:
        """
        VPC configuration information. See below.
        """
        return pulumi.get(self, "vpc_configs")

    @vpc_configs.setter
    def vpc_configs(self, value: pulumi.Input[Sequence[pulumi.Input['ServerlessClusterVpcConfigArgs']]]):
        pulumi.set(self, "vpc_configs", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the serverless cluster.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_name", value)

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
class _ServerlessClusterState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 client_authentication: Optional[pulumi.Input['ServerlessClusterClientAuthenticationArgs']] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_configs: Optional[pulumi.Input[Sequence[pulumi.Input['ServerlessClusterVpcConfigArgs']]]] = None):
        """
        Input properties used for looking up and filtering ServerlessCluster resources.
        :param pulumi.Input[str] arn: The ARN of the serverless cluster.
        :param pulumi.Input['ServerlessClusterClientAuthenticationArgs'] client_authentication: Specifies client authentication information for the serverless cluster. See below.
        :param pulumi.Input[str] cluster_name: The name of the serverless cluster.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[Sequence[pulumi.Input['ServerlessClusterVpcConfigArgs']]] vpc_configs: VPC configuration information. See below.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if client_authentication is not None:
            pulumi.set(__self__, "client_authentication", client_authentication)
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if vpc_configs is not None:
            pulumi.set(__self__, "vpc_configs", vpc_configs)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the serverless cluster.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="clientAuthentication")
    def client_authentication(self) -> Optional[pulumi.Input['ServerlessClusterClientAuthenticationArgs']]:
        """
        Specifies client authentication information for the serverless cluster. See below.
        """
        return pulumi.get(self, "client_authentication")

    @client_authentication.setter
    def client_authentication(self, value: Optional[pulumi.Input['ServerlessClusterClientAuthenticationArgs']]):
        pulumi.set(self, "client_authentication", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the serverless cluster.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_name", value)

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
    @pulumi.getter(name="vpcConfigs")
    def vpc_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServerlessClusterVpcConfigArgs']]]]:
        """
        VPC configuration information. See below.
        """
        return pulumi.get(self, "vpc_configs")

    @vpc_configs.setter
    def vpc_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServerlessClusterVpcConfigArgs']]]]):
        pulumi.set(self, "vpc_configs", value)


class ServerlessCluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_authentication: Optional[pulumi.Input[pulumi.InputType['ServerlessClusterClientAuthenticationArgs']]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerlessClusterVpcConfigArgs']]]]] = None,
                 __props__=None):
        """
        Manages an Amazon MSK Serverless cluster.

        > **Note:** To manage a _provisioned_ Amazon MSK cluster, use the `msk.Cluster` resource.

        ## Import

        terraform import {

         to = aws_msk_serverless_cluster.example

         id = "arn:aws:kafka:us-west-2:123456789012:cluster/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3" } Using `pulumi import`, import MSK serverless clusters using the cluster `arn`. For exampleconsole % pulumi import aws_msk_serverless_cluster.example arn:aws:kafka:us-west-2:123456789012:cluster/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ServerlessClusterClientAuthenticationArgs']] client_authentication: Specifies client authentication information for the serverless cluster. See below.
        :param pulumi.Input[str] cluster_name: The name of the serverless cluster.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerlessClusterVpcConfigArgs']]]] vpc_configs: VPC configuration information. See below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServerlessClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Amazon MSK Serverless cluster.

        > **Note:** To manage a _provisioned_ Amazon MSK cluster, use the `msk.Cluster` resource.

        ## Import

        terraform import {

         to = aws_msk_serverless_cluster.example

         id = "arn:aws:kafka:us-west-2:123456789012:cluster/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3" } Using `pulumi import`, import MSK serverless clusters using the cluster `arn`. For exampleconsole % pulumi import aws_msk_serverless_cluster.example arn:aws:kafka:us-west-2:123456789012:cluster/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3

        :param str resource_name: The name of the resource.
        :param ServerlessClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServerlessClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_authentication: Optional[pulumi.Input[pulumi.InputType['ServerlessClusterClientAuthenticationArgs']]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerlessClusterVpcConfigArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServerlessClusterArgs.__new__(ServerlessClusterArgs)

            if client_authentication is None and not opts.urn:
                raise TypeError("Missing required property 'client_authentication'")
            __props__.__dict__["client_authentication"] = client_authentication
            __props__.__dict__["cluster_name"] = cluster_name
            __props__.__dict__["tags"] = tags
            if vpc_configs is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_configs'")
            __props__.__dict__["vpc_configs"] = vpc_configs
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        super(ServerlessCluster, __self__).__init__(
            'aws:msk/serverlessCluster:ServerlessCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            client_authentication: Optional[pulumi.Input[pulumi.InputType['ServerlessClusterClientAuthenticationArgs']]] = None,
            cluster_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerlessClusterVpcConfigArgs']]]]] = None) -> 'ServerlessCluster':
        """
        Get an existing ServerlessCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the serverless cluster.
        :param pulumi.Input[pulumi.InputType['ServerlessClusterClientAuthenticationArgs']] client_authentication: Specifies client authentication information for the serverless cluster. See below.
        :param pulumi.Input[str] cluster_name: The name of the serverless cluster.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServerlessClusterVpcConfigArgs']]]] vpc_configs: VPC configuration information. See below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServerlessClusterState.__new__(_ServerlessClusterState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["client_authentication"] = client_authentication
        __props__.__dict__["cluster_name"] = cluster_name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["vpc_configs"] = vpc_configs
        return ServerlessCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the serverless cluster.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="clientAuthentication")
    def client_authentication(self) -> pulumi.Output['outputs.ServerlessClusterClientAuthentication']:
        """
        Specifies client authentication information for the serverless cluster. See below.
        """
        return pulumi.get(self, "client_authentication")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[str]:
        """
        The name of the serverless cluster.
        """
        return pulumi.get(self, "cluster_name")

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
    @pulumi.getter(name="vpcConfigs")
    def vpc_configs(self) -> pulumi.Output[Sequence['outputs.ServerlessClusterVpcConfig']]:
        """
        VPC configuration information. See below.
        """
        return pulumi.get(self, "vpc_configs")

