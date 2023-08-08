# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ManagedPrefixListEntryInitArgs', 'ManagedPrefixListEntry']

@pulumi.input_type
class ManagedPrefixListEntryInitArgs:
    def __init__(__self__, *,
                 cidr: pulumi.Input[str],
                 prefix_list_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ManagedPrefixListEntry resource.
        :param pulumi.Input[str] cidr: CIDR block of this entry.
        :param pulumi.Input[str] prefix_list_id: CIDR block of this entry.
        :param pulumi.Input[str] description: Description of this entry. Please note that due to API limitations, updating only the description of an entry will require recreating the entry.
        """
        pulumi.set(__self__, "cidr", cidr)
        pulumi.set(__self__, "prefix_list_id", prefix_list_id)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Input[str]:
        """
        CIDR block of this entry.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: pulumi.Input[str]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter(name="prefixListId")
    def prefix_list_id(self) -> pulumi.Input[str]:
        """
        CIDR block of this entry.
        """
        return pulumi.get(self, "prefix_list_id")

    @prefix_list_id.setter
    def prefix_list_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "prefix_list_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of this entry. Please note that due to API limitations, updating only the description of an entry will require recreating the entry.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _ManagedPrefixListEntryState:
    def __init__(__self__, *,
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 prefix_list_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ManagedPrefixListEntry resources.
        :param pulumi.Input[str] cidr: CIDR block of this entry.
        :param pulumi.Input[str] description: Description of this entry. Please note that due to API limitations, updating only the description of an entry will require recreating the entry.
        :param pulumi.Input[str] prefix_list_id: CIDR block of this entry.
        """
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if prefix_list_id is not None:
            pulumi.set(__self__, "prefix_list_id", prefix_list_id)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        """
        CIDR block of this entry.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of this entry. Please note that due to API limitations, updating only the description of an entry will require recreating the entry.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="prefixListId")
    def prefix_list_id(self) -> Optional[pulumi.Input[str]]:
        """
        CIDR block of this entry.
        """
        return pulumi.get(self, "prefix_list_id")

    @prefix_list_id.setter
    def prefix_list_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix_list_id", value)


class ManagedPrefixListEntry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 prefix_list_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Use the `aws_prefix_list_entry` resource to manage a managed prefix list entry.

        > **NOTE:** Pulumi currently provides two resources for managing Managed Prefix Lists and Managed Prefix List Entries. The standalone resource, Managed Prefix List Entry, is used to manage a single entry. The Managed Prefix List resource is used to manage multiple entries defined in-line. It is important to note that you cannot use a Managed Prefix List with in-line rules in conjunction with any Managed Prefix List Entry resources. This will result in a conflict of entries and will cause the entries to be overwritten.

        > **NOTE:** To improve execution times on larger updates, it is recommended to use the inline `entry` block as part of the Managed Prefix List resource when creating a prefix list with more than 100 entries. You can find more information about the resource here.

        ## Example Usage

        Basic usage.

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2.ManagedPrefixList("example",
            address_family="IPv4",
            max_entries=5,
            tags={
                "Env": "live",
            })
        entry1 = aws.ec2.ManagedPrefixListEntry("entry1",
            cidr=aws_vpc["example"]["cidr_block"],
            description="Primary",
            prefix_list_id=example.id)
        ```

        ## Import

        terraform import {

         to = aws_ec2_managed_prefix_list_entry.default

         id = "pl-0570a1d2d725c16be,10.0.3.0/24" } Using `pulumi import`, import prefix list entries using `prefix_list_id` and `cidr` separated by a comma (`,`). For exampleconsole % pulumi import aws_ec2_managed_prefix_list_entry.default pl-0570a1d2d725c16be,10.0.3.0/24

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr: CIDR block of this entry.
        :param pulumi.Input[str] description: Description of this entry. Please note that due to API limitations, updating only the description of an entry will require recreating the entry.
        :param pulumi.Input[str] prefix_list_id: CIDR block of this entry.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ManagedPrefixListEntryInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use the `aws_prefix_list_entry` resource to manage a managed prefix list entry.

        > **NOTE:** Pulumi currently provides two resources for managing Managed Prefix Lists and Managed Prefix List Entries. The standalone resource, Managed Prefix List Entry, is used to manage a single entry. The Managed Prefix List resource is used to manage multiple entries defined in-line. It is important to note that you cannot use a Managed Prefix List with in-line rules in conjunction with any Managed Prefix List Entry resources. This will result in a conflict of entries and will cause the entries to be overwritten.

        > **NOTE:** To improve execution times on larger updates, it is recommended to use the inline `entry` block as part of the Managed Prefix List resource when creating a prefix list with more than 100 entries. You can find more information about the resource here.

        ## Example Usage

        Basic usage.

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2.ManagedPrefixList("example",
            address_family="IPv4",
            max_entries=5,
            tags={
                "Env": "live",
            })
        entry1 = aws.ec2.ManagedPrefixListEntry("entry1",
            cidr=aws_vpc["example"]["cidr_block"],
            description="Primary",
            prefix_list_id=example.id)
        ```

        ## Import

        terraform import {

         to = aws_ec2_managed_prefix_list_entry.default

         id = "pl-0570a1d2d725c16be,10.0.3.0/24" } Using `pulumi import`, import prefix list entries using `prefix_list_id` and `cidr` separated by a comma (`,`). For exampleconsole % pulumi import aws_ec2_managed_prefix_list_entry.default pl-0570a1d2d725c16be,10.0.3.0/24

        :param str resource_name: The name of the resource.
        :param ManagedPrefixListEntryInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ManagedPrefixListEntryInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 prefix_list_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ManagedPrefixListEntryInitArgs.__new__(ManagedPrefixListEntryInitArgs)

            if cidr is None and not opts.urn:
                raise TypeError("Missing required property 'cidr'")
            __props__.__dict__["cidr"] = cidr
            __props__.__dict__["description"] = description
            if prefix_list_id is None and not opts.urn:
                raise TypeError("Missing required property 'prefix_list_id'")
            __props__.__dict__["prefix_list_id"] = prefix_list_id
        super(ManagedPrefixListEntry, __self__).__init__(
            'aws:ec2/managedPrefixListEntry:ManagedPrefixListEntry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cidr: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            prefix_list_id: Optional[pulumi.Input[str]] = None) -> 'ManagedPrefixListEntry':
        """
        Get an existing ManagedPrefixListEntry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr: CIDR block of this entry.
        :param pulumi.Input[str] description: Description of this entry. Please note that due to API limitations, updating only the description of an entry will require recreating the entry.
        :param pulumi.Input[str] prefix_list_id: CIDR block of this entry.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ManagedPrefixListEntryState.__new__(_ManagedPrefixListEntryState)

        __props__.__dict__["cidr"] = cidr
        __props__.__dict__["description"] = description
        __props__.__dict__["prefix_list_id"] = prefix_list_id
        return ManagedPrefixListEntry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Output[str]:
        """
        CIDR block of this entry.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of this entry. Please note that due to API limitations, updating only the description of an entry will require recreating the entry.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="prefixListId")
    def prefix_list_id(self) -> pulumi.Output[str]:
        """
        CIDR block of this entry.
        """
        return pulumi.get(self, "prefix_list_id")

