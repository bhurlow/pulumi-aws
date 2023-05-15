# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ClusterEndpointArgs', 'ClusterEndpoint']

@pulumi.input_type
class ClusterEndpointArgs:
    def __init__(__self__, *,
                 cluster_endpoint_identifier: pulumi.Input[str],
                 cluster_identifier: pulumi.Input[str],
                 endpoint_type: pulumi.Input[str],
                 excluded_members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 static_members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ClusterEndpoint resource.
        :param pulumi.Input[str] cluster_endpoint_identifier: The identifier of the endpoint.
        :param pulumi.Input[str] cluster_identifier: The DB cluster identifier of the DB cluster associated with the endpoint.
        :param pulumi.Input[str] endpoint_type: The type of the endpoint. One of: `READER`, `WRITER`, `ANY`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_members: List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] static_members: List of DB instance identifiers that are part of the custom endpoint group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the Neptune cluster. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "cluster_endpoint_identifier", cluster_endpoint_identifier)
        pulumi.set(__self__, "cluster_identifier", cluster_identifier)
        pulumi.set(__self__, "endpoint_type", endpoint_type)
        if excluded_members is not None:
            pulumi.set(__self__, "excluded_members", excluded_members)
        if static_members is not None:
            pulumi.set(__self__, "static_members", static_members)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="clusterEndpointIdentifier")
    def cluster_endpoint_identifier(self) -> pulumi.Input[str]:
        """
        The identifier of the endpoint.
        """
        return pulumi.get(self, "cluster_endpoint_identifier")

    @cluster_endpoint_identifier.setter
    def cluster_endpoint_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_endpoint_identifier", value)

    @property
    @pulumi.getter(name="clusterIdentifier")
    def cluster_identifier(self) -> pulumi.Input[str]:
        """
        The DB cluster identifier of the DB cluster associated with the endpoint.
        """
        return pulumi.get(self, "cluster_identifier")

    @cluster_identifier.setter
    def cluster_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_identifier", value)

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> pulumi.Input[str]:
        """
        The type of the endpoint. One of: `READER`, `WRITER`, `ANY`.
        """
        return pulumi.get(self, "endpoint_type")

    @endpoint_type.setter
    def endpoint_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint_type", value)

    @property
    @pulumi.getter(name="excludedMembers")
    def excluded_members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty.
        """
        return pulumi.get(self, "excluded_members")

    @excluded_members.setter
    def excluded_members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excluded_members", value)

    @property
    @pulumi.getter(name="staticMembers")
    def static_members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of DB instance identifiers that are part of the custom endpoint group.
        """
        return pulumi.get(self, "static_members")

    @static_members.setter
    def static_members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "static_members", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the Neptune cluster. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ClusterEndpointState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 cluster_endpoint_identifier: Optional[pulumi.Input[str]] = None,
                 cluster_identifier: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 endpoint_type: Optional[pulumi.Input[str]] = None,
                 excluded_members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 static_members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ClusterEndpoint resources.
        :param pulumi.Input[str] arn: The Neptune Cluster Endpoint Amazon Resource Name (ARN).
        :param pulumi.Input[str] cluster_endpoint_identifier: The identifier of the endpoint.
        :param pulumi.Input[str] cluster_identifier: The DB cluster identifier of the DB cluster associated with the endpoint.
        :param pulumi.Input[str] endpoint: The DNS address of the endpoint.
        :param pulumi.Input[str] endpoint_type: The type of the endpoint. One of: `READER`, `WRITER`, `ANY`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_members: List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] static_members: List of DB instance identifiers that are part of the custom endpoint group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the Neptune cluster. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if cluster_endpoint_identifier is not None:
            pulumi.set(__self__, "cluster_endpoint_identifier", cluster_endpoint_identifier)
        if cluster_identifier is not None:
            pulumi.set(__self__, "cluster_identifier", cluster_identifier)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if endpoint_type is not None:
            pulumi.set(__self__, "endpoint_type", endpoint_type)
        if excluded_members is not None:
            pulumi.set(__self__, "excluded_members", excluded_members)
        if static_members is not None:
            pulumi.set(__self__, "static_members", static_members)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Neptune Cluster Endpoint Amazon Resource Name (ARN).
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="clusterEndpointIdentifier")
    def cluster_endpoint_identifier(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier of the endpoint.
        """
        return pulumi.get(self, "cluster_endpoint_identifier")

    @cluster_endpoint_identifier.setter
    def cluster_endpoint_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_endpoint_identifier", value)

    @property
    @pulumi.getter(name="clusterIdentifier")
    def cluster_identifier(self) -> Optional[pulumi.Input[str]]:
        """
        The DB cluster identifier of the DB cluster associated with the endpoint.
        """
        return pulumi.get(self, "cluster_identifier")

    @cluster_identifier.setter
    def cluster_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_identifier", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The DNS address of the endpoint.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the endpoint. One of: `READER`, `WRITER`, `ANY`.
        """
        return pulumi.get(self, "endpoint_type")

    @endpoint_type.setter
    def endpoint_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_type", value)

    @property
    @pulumi.getter(name="excludedMembers")
    def excluded_members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty.
        """
        return pulumi.get(self, "excluded_members")

    @excluded_members.setter
    def excluded_members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excluded_members", value)

    @property
    @pulumi.getter(name="staticMembers")
    def static_members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of DB instance identifiers that are part of the custom endpoint group.
        """
        return pulumi.get(self, "static_members")

    @static_members.setter
    def static_members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "static_members", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the Neptune cluster. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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


class ClusterEndpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_endpoint_identifier: Optional[pulumi.Input[str]] = None,
                 cluster_identifier: Optional[pulumi.Input[str]] = None,
                 endpoint_type: Optional[pulumi.Input[str]] = None,
                 excluded_members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 static_members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides an Neptune Cluster Endpoint Resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.neptune.ClusterEndpoint("example",
            cluster_identifier=aws_neptune_cluster["test"]["cluster_identifier"],
            cluster_endpoint_identifier="example",
            endpoint_type="READER")
        ```

        ## Import

        `aws_neptune_cluster_endpoint` can be imported by using the `cluster-identifier:endpoint-identfier`, e.g.,

        ```sh
         $ pulumi import aws:neptune/clusterEndpoint:ClusterEndpoint example my-cluster:my-endpoint
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_endpoint_identifier: The identifier of the endpoint.
        :param pulumi.Input[str] cluster_identifier: The DB cluster identifier of the DB cluster associated with the endpoint.
        :param pulumi.Input[str] endpoint_type: The type of the endpoint. One of: `READER`, `WRITER`, `ANY`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_members: List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] static_members: List of DB instance identifiers that are part of the custom endpoint group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the Neptune cluster. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClusterEndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Neptune Cluster Endpoint Resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.neptune.ClusterEndpoint("example",
            cluster_identifier=aws_neptune_cluster["test"]["cluster_identifier"],
            cluster_endpoint_identifier="example",
            endpoint_type="READER")
        ```

        ## Import

        `aws_neptune_cluster_endpoint` can be imported by using the `cluster-identifier:endpoint-identfier`, e.g.,

        ```sh
         $ pulumi import aws:neptune/clusterEndpoint:ClusterEndpoint example my-cluster:my-endpoint
        ```

        :param str resource_name: The name of the resource.
        :param ClusterEndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterEndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_endpoint_identifier: Optional[pulumi.Input[str]] = None,
                 cluster_identifier: Optional[pulumi.Input[str]] = None,
                 endpoint_type: Optional[pulumi.Input[str]] = None,
                 excluded_members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 static_members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClusterEndpointArgs.__new__(ClusterEndpointArgs)

            if cluster_endpoint_identifier is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_endpoint_identifier'")
            __props__.__dict__["cluster_endpoint_identifier"] = cluster_endpoint_identifier
            if cluster_identifier is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_identifier'")
            __props__.__dict__["cluster_identifier"] = cluster_identifier
            if endpoint_type is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint_type'")
            __props__.__dict__["endpoint_type"] = endpoint_type
            __props__.__dict__["excluded_members"] = excluded_members
            __props__.__dict__["static_members"] = static_members
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["endpoint"] = None
            __props__.__dict__["tags_all"] = None
        super(ClusterEndpoint, __self__).__init__(
            'aws:neptune/clusterEndpoint:ClusterEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            cluster_endpoint_identifier: Optional[pulumi.Input[str]] = None,
            cluster_identifier: Optional[pulumi.Input[str]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            endpoint_type: Optional[pulumi.Input[str]] = None,
            excluded_members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            static_members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'ClusterEndpoint':
        """
        Get an existing ClusterEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Neptune Cluster Endpoint Amazon Resource Name (ARN).
        :param pulumi.Input[str] cluster_endpoint_identifier: The identifier of the endpoint.
        :param pulumi.Input[str] cluster_identifier: The DB cluster identifier of the DB cluster associated with the endpoint.
        :param pulumi.Input[str] endpoint: The DNS address of the endpoint.
        :param pulumi.Input[str] endpoint_type: The type of the endpoint. One of: `READER`, `WRITER`, `ANY`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_members: List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] static_members: List of DB instance identifiers that are part of the custom endpoint group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the Neptune cluster. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClusterEndpointState.__new__(_ClusterEndpointState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["cluster_endpoint_identifier"] = cluster_endpoint_identifier
        __props__.__dict__["cluster_identifier"] = cluster_identifier
        __props__.__dict__["endpoint"] = endpoint
        __props__.__dict__["endpoint_type"] = endpoint_type
        __props__.__dict__["excluded_members"] = excluded_members
        __props__.__dict__["static_members"] = static_members
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return ClusterEndpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Neptune Cluster Endpoint Amazon Resource Name (ARN).
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="clusterEndpointIdentifier")
    def cluster_endpoint_identifier(self) -> pulumi.Output[str]:
        """
        The identifier of the endpoint.
        """
        return pulumi.get(self, "cluster_endpoint_identifier")

    @property
    @pulumi.getter(name="clusterIdentifier")
    def cluster_identifier(self) -> pulumi.Output[str]:
        """
        The DB cluster identifier of the DB cluster associated with the endpoint.
        """
        return pulumi.get(self, "cluster_identifier")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        The DNS address of the endpoint.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> pulumi.Output[str]:
        """
        The type of the endpoint. One of: `READER`, `WRITER`, `ANY`.
        """
        return pulumi.get(self, "endpoint_type")

    @property
    @pulumi.getter(name="excludedMembers")
    def excluded_members(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty.
        """
        return pulumi.get(self, "excluded_members")

    @property
    @pulumi.getter(name="staticMembers")
    def static_members(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of DB instance identifiers that are part of the custom endpoint group.
        """
        return pulumi.get(self, "static_members")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the Neptune cluster. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

