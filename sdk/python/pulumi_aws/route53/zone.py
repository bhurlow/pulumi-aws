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

__all__ = ['ZoneArgs', 'Zone']

@pulumi.input_type
class ZoneArgs:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 delegation_set_id: Optional[pulumi.Input[str]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpcs: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneVpcArgs']]]] = None):
        """
        The set of arguments for constructing a Zone resource.
        :param pulumi.Input[str] comment: A comment for the hosted zone. Defaults to 'Managed by Pulumi'.
        :param pulumi.Input[str] delegation_set_id: The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
        :param pulumi.Input[bool] force_destroy: Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
        :param pulumi.Input[str] name: This is the name of the hosted zone.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the zone. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Sequence[pulumi.Input['ZoneVpcArgs']]] vpcs: Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegation_set_id` argument in this resource and any `route53.ZoneAssociation` resource specifying the same zone ID. Detailed below.
        """
        if comment is None:
            comment = 'Managed by Pulumi'
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if delegation_set_id is not None:
            pulumi.set(__self__, "delegation_set_id", delegation_set_id)
        if force_destroy is not None:
            pulumi.set(__self__, "force_destroy", force_destroy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpcs is not None:
            pulumi.set(__self__, "vpcs", vpcs)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        A comment for the hosted zone. Defaults to 'Managed by Pulumi'.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="delegationSetId")
    def delegation_set_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
        """
        return pulumi.get(self, "delegation_set_id")

    @delegation_set_id.setter
    def delegation_set_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delegation_set_id", value)

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
        """
        return pulumi.get(self, "force_destroy")

    @force_destroy.setter
    def force_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_destroy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        This is the name of the hosted zone.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the zone. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def vpcs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZoneVpcArgs']]]]:
        """
        Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegation_set_id` argument in this resource and any `route53.ZoneAssociation` resource specifying the same zone ID. Detailed below.
        """
        return pulumi.get(self, "vpcs")

    @vpcs.setter
    def vpcs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneVpcArgs']]]]):
        pulumi.set(self, "vpcs", value)


@pulumi.input_type
class _ZoneState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 delegation_set_id: Optional[pulumi.Input[str]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 primary_name_server: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpcs: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneVpcArgs']]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Zone resources.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the Hosted Zone.
        :param pulumi.Input[str] comment: A comment for the hosted zone. Defaults to 'Managed by Pulumi'.
        :param pulumi.Input[str] delegation_set_id: The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
        :param pulumi.Input[bool] force_destroy: Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
        :param pulumi.Input[str] name: This is the name of the hosted zone.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] name_servers: A list of name servers in associated (or default) delegation set.
               Find more about delegation sets in [AWS docs](https://docs.aws.amazon.com/Route53/latest/APIReference/actions-on-reusable-delegation-sets.html).
        :param pulumi.Input[str] primary_name_server: The Route 53 name server that created the SOA record.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the zone. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[Sequence[pulumi.Input['ZoneVpcArgs']]] vpcs: Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegation_set_id` argument in this resource and any `route53.ZoneAssociation` resource specifying the same zone ID. Detailed below.
        :param pulumi.Input[str] zone_id: The Hosted Zone ID. This can be referenced by zone records.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if comment is None:
            comment = 'Managed by Pulumi'
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if delegation_set_id is not None:
            pulumi.set(__self__, "delegation_set_id", delegation_set_id)
        if force_destroy is not None:
            pulumi.set(__self__, "force_destroy", force_destroy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if name_servers is not None:
            pulumi.set(__self__, "name_servers", name_servers)
        if primary_name_server is not None:
            pulumi.set(__self__, "primary_name_server", primary_name_server)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
            pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if vpcs is not None:
            pulumi.set(__self__, "vpcs", vpcs)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the Hosted Zone.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        A comment for the hosted zone. Defaults to 'Managed by Pulumi'.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="delegationSetId")
    def delegation_set_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
        """
        return pulumi.get(self, "delegation_set_id")

    @delegation_set_id.setter
    def delegation_set_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delegation_set_id", value)

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
        """
        return pulumi.get(self, "force_destroy")

    @force_destroy.setter
    def force_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_destroy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        This is the name of the hosted zone.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nameServers")
    def name_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of name servers in associated (or default) delegation set.
        Find more about delegation sets in [AWS docs](https://docs.aws.amazon.com/Route53/latest/APIReference/actions-on-reusable-delegation-sets.html).
        """
        return pulumi.get(self, "name_servers")

    @name_servers.setter
    def name_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "name_servers", value)

    @property
    @pulumi.getter(name="primaryNameServer")
    def primary_name_server(self) -> Optional[pulumi.Input[str]]:
        """
        The Route 53 name server that created the SOA record.
        """
        return pulumi.get(self, "primary_name_server")

    @primary_name_server.setter
    def primary_name_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary_name_server", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the zone. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

    @property
    @pulumi.getter
    def vpcs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZoneVpcArgs']]]]:
        """
        Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegation_set_id` argument in this resource and any `route53.ZoneAssociation` resource specifying the same zone ID. Detailed below.
        """
        return pulumi.get(self, "vpcs")

    @vpcs.setter
    def vpcs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneVpcArgs']]]]):
        pulumi.set(self, "vpcs", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Hosted Zone ID. This can be referenced by zone records.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class Zone(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 delegation_set_id: Optional[pulumi.Input[str]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpcs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneVpcArgs']]]]] = None,
                 __props__=None):
        """
        Manages a Route53 Hosted Zone. For managing Domain Name System Security Extensions (DNSSEC), see the `route53.KeySigningKey` and `route53.HostedZoneDnsSec` resources.

        ## Example Usage
        ### Public Zone

        ```python
        import pulumi
        import pulumi_aws as aws

        primary = aws.route53.Zone("primary")
        ```
        ### Public Subdomain Zone

        For use in subdomains, note that you need to create a
        `route53.Record` of type `NS` as well as the subdomain
        zone.

        ```python
        import pulumi
        import pulumi_aws as aws

        main = aws.route53.Zone("main")
        dev = aws.route53.Zone("dev", tags={
            "Environment": "dev",
        })
        dev_ns = aws.route53.Record("dev-ns",
            zone_id=main.zone_id,
            name="dev.example.com",
            type="NS",
            ttl=30,
            records=dev.name_servers)
        ```
        ### Private Zone

        > **NOTE:** This provider provides both exclusive VPC associations defined in-line in this resource via `vpc` configuration blocks and a separate ` Zone VPC Association resource. At this time, you cannot use in-line VPC associations in conjunction with any  `route53.ZoneAssociation`  resources with the same zone ID otherwise it will cause a perpetual difference in plan output. You can optionally use [ `ignoreChanges` ](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) to manage additional associations via the  `route53.ZoneAssociation` resource.

        > **NOTE:** Private zones require at least one VPC association at all times.

        ```python
        import pulumi
        import pulumi_aws as aws

        private = aws.route53.Zone("private", vpcs=[aws.route53.ZoneVpcArgs(
            vpc_id=aws_vpc["example"]["id"],
        )])
        ```

        ## Import

        Using `pulumi import`, import Route53 Zones using the zone `id`. For example:

        ```sh
         $ pulumi import aws:route53/zone:Zone myzone Z1D633PJN98FT9
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: A comment for the hosted zone. Defaults to 'Managed by Pulumi'.
        :param pulumi.Input[str] delegation_set_id: The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
        :param pulumi.Input[bool] force_destroy: Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
        :param pulumi.Input[str] name: This is the name of the hosted zone.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the zone. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneVpcArgs']]]] vpcs: Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegation_set_id` argument in this resource and any `route53.ZoneAssociation` resource specifying the same zone ID. Detailed below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ZoneArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Route53 Hosted Zone. For managing Domain Name System Security Extensions (DNSSEC), see the `route53.KeySigningKey` and `route53.HostedZoneDnsSec` resources.

        ## Example Usage
        ### Public Zone

        ```python
        import pulumi
        import pulumi_aws as aws

        primary = aws.route53.Zone("primary")
        ```
        ### Public Subdomain Zone

        For use in subdomains, note that you need to create a
        `route53.Record` of type `NS` as well as the subdomain
        zone.

        ```python
        import pulumi
        import pulumi_aws as aws

        main = aws.route53.Zone("main")
        dev = aws.route53.Zone("dev", tags={
            "Environment": "dev",
        })
        dev_ns = aws.route53.Record("dev-ns",
            zone_id=main.zone_id,
            name="dev.example.com",
            type="NS",
            ttl=30,
            records=dev.name_servers)
        ```
        ### Private Zone

        > **NOTE:** This provider provides both exclusive VPC associations defined in-line in this resource via `vpc` configuration blocks and a separate ` Zone VPC Association resource. At this time, you cannot use in-line VPC associations in conjunction with any  `route53.ZoneAssociation`  resources with the same zone ID otherwise it will cause a perpetual difference in plan output. You can optionally use [ `ignoreChanges` ](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) to manage additional associations via the  `route53.ZoneAssociation` resource.

        > **NOTE:** Private zones require at least one VPC association at all times.

        ```python
        import pulumi
        import pulumi_aws as aws

        private = aws.route53.Zone("private", vpcs=[aws.route53.ZoneVpcArgs(
            vpc_id=aws_vpc["example"]["id"],
        )])
        ```

        ## Import

        Using `pulumi import`, import Route53 Zones using the zone `id`. For example:

        ```sh
         $ pulumi import aws:route53/zone:Zone myzone Z1D633PJN98FT9
        ```

        :param str resource_name: The name of the resource.
        :param ZoneArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ZoneArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 delegation_set_id: Optional[pulumi.Input[str]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpcs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneVpcArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ZoneArgs.__new__(ZoneArgs)

            if comment is None:
                comment = 'Managed by Pulumi'
            __props__.__dict__["comment"] = comment
            __props__.__dict__["delegation_set_id"] = delegation_set_id
            __props__.__dict__["force_destroy"] = force_destroy
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["vpcs"] = vpcs
            __props__.__dict__["arn"] = None
            __props__.__dict__["name_servers"] = None
            __props__.__dict__["primary_name_server"] = None
            __props__.__dict__["tags_all"] = None
            __props__.__dict__["zone_id"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Zone, __self__).__init__(
            'aws:route53/zone:Zone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            delegation_set_id: Optional[pulumi.Input[str]] = None,
            force_destroy: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            primary_name_server: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpcs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneVpcArgs']]]]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'Zone':
        """
        Get an existing Zone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the Hosted Zone.
        :param pulumi.Input[str] comment: A comment for the hosted zone. Defaults to 'Managed by Pulumi'.
        :param pulumi.Input[str] delegation_set_id: The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
        :param pulumi.Input[bool] force_destroy: Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
        :param pulumi.Input[str] name: This is the name of the hosted zone.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] name_servers: A list of name servers in associated (or default) delegation set.
               Find more about delegation sets in [AWS docs](https://docs.aws.amazon.com/Route53/latest/APIReference/actions-on-reusable-delegation-sets.html).
        :param pulumi.Input[str] primary_name_server: The Route 53 name server that created the SOA record.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the zone. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneVpcArgs']]]] vpcs: Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegation_set_id` argument in this resource and any `route53.ZoneAssociation` resource specifying the same zone ID. Detailed below.
        :param pulumi.Input[str] zone_id: The Hosted Zone ID. This can be referenced by zone records.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ZoneState.__new__(_ZoneState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["comment"] = comment
        __props__.__dict__["delegation_set_id"] = delegation_set_id
        __props__.__dict__["force_destroy"] = force_destroy
        __props__.__dict__["name"] = name
        __props__.__dict__["name_servers"] = name_servers
        __props__.__dict__["primary_name_server"] = primary_name_server
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["vpcs"] = vpcs
        __props__.__dict__["zone_id"] = zone_id
        return Zone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the Hosted Zone.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[str]:
        """
        A comment for the hosted zone. Defaults to 'Managed by Pulumi'.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="delegationSetId")
    def delegation_set_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
        """
        return pulumi.get(self, "delegation_set_id")

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
        """
        return pulumi.get(self, "force_destroy")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        This is the name of the hosted zone.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameServers")
    def name_servers(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of name servers in associated (or default) delegation set.
        Find more about delegation sets in [AWS docs](https://docs.aws.amazon.com/Route53/latest/APIReference/actions-on-reusable-delegation-sets.html).
        """
        return pulumi.get(self, "name_servers")

    @property
    @pulumi.getter(name="primaryNameServer")
    def primary_name_server(self) -> pulumi.Output[str]:
        """
        The Route 53 name server that created the SOA record.
        """
        return pulumi.get(self, "primary_name_server")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the zone. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

    @property
    @pulumi.getter
    def vpcs(self) -> pulumi.Output[Optional[Sequence['outputs.ZoneVpc']]]:
        """
        Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegation_set_id` argument in this resource and any `route53.ZoneAssociation` resource specifying the same zone ID. Detailed below.
        """
        return pulumi.get(self, "vpcs")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        The Hosted Zone ID. This can be referenced by zone records.
        """
        return pulumi.get(self, "zone_id")

