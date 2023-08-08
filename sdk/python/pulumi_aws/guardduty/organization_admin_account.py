# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['OrganizationAdminAccountArgs', 'OrganizationAdminAccount']

@pulumi.input_type
class OrganizationAdminAccountArgs:
    def __init__(__self__, *,
                 admin_account_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a OrganizationAdminAccount resource.
        :param pulumi.Input[str] admin_account_id: AWS account identifier to designate as a delegated administrator for GuardDuty.
        """
        pulumi.set(__self__, "admin_account_id", admin_account_id)

    @property
    @pulumi.getter(name="adminAccountId")
    def admin_account_id(self) -> pulumi.Input[str]:
        """
        AWS account identifier to designate as a delegated administrator for GuardDuty.
        """
        return pulumi.get(self, "admin_account_id")

    @admin_account_id.setter
    def admin_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "admin_account_id", value)


@pulumi.input_type
class _OrganizationAdminAccountState:
    def __init__(__self__, *,
                 admin_account_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OrganizationAdminAccount resources.
        :param pulumi.Input[str] admin_account_id: AWS account identifier to designate as a delegated administrator for GuardDuty.
        """
        if admin_account_id is not None:
            pulumi.set(__self__, "admin_account_id", admin_account_id)

    @property
    @pulumi.getter(name="adminAccountId")
    def admin_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        AWS account identifier to designate as a delegated administrator for GuardDuty.
        """
        return pulumi.get(self, "admin_account_id")

    @admin_account_id.setter
    def admin_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_account_id", value)


class OrganizationAdminAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_account_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a GuardDuty Organization Admin Account. The AWS account utilizing this resource must be an Organizations primary account. More information about Organizations support in GuardDuty can be found in the [GuardDuty User Guide](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_organizations.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_organization = aws.organizations.Organization("exampleOrganization",
            aws_service_access_principals=["guardduty.amazonaws.com"],
            feature_set="ALL")
        example_detector = aws.guardduty.Detector("exampleDetector")
        example_organization_admin_account = aws.guardduty.OrganizationAdminAccount("exampleOrganizationAdminAccount", admin_account_id="123456789012",
        opts=pulumi.ResourceOptions(depends_on=[example_organization]))
        ```

        ## Import

        terraform import {

         to = aws_guardduty_organization_admin_account.example

         id = "123456789012" } Using `pulumi import`, import GuardDuty Organization Admin Account using the AWS account ID. For exampleconsole % pulumi import aws_guardduty_organization_admin_account.example 123456789012

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_account_id: AWS account identifier to designate as a delegated administrator for GuardDuty.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrganizationAdminAccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a GuardDuty Organization Admin Account. The AWS account utilizing this resource must be an Organizations primary account. More information about Organizations support in GuardDuty can be found in the [GuardDuty User Guide](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_organizations.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_organization = aws.organizations.Organization("exampleOrganization",
            aws_service_access_principals=["guardduty.amazonaws.com"],
            feature_set="ALL")
        example_detector = aws.guardduty.Detector("exampleDetector")
        example_organization_admin_account = aws.guardduty.OrganizationAdminAccount("exampleOrganizationAdminAccount", admin_account_id="123456789012",
        opts=pulumi.ResourceOptions(depends_on=[example_organization]))
        ```

        ## Import

        terraform import {

         to = aws_guardduty_organization_admin_account.example

         id = "123456789012" } Using `pulumi import`, import GuardDuty Organization Admin Account using the AWS account ID. For exampleconsole % pulumi import aws_guardduty_organization_admin_account.example 123456789012

        :param str resource_name: The name of the resource.
        :param OrganizationAdminAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationAdminAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_account_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OrganizationAdminAccountArgs.__new__(OrganizationAdminAccountArgs)

            if admin_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'admin_account_id'")
            __props__.__dict__["admin_account_id"] = admin_account_id
        super(OrganizationAdminAccount, __self__).__init__(
            'aws:guardduty/organizationAdminAccount:OrganizationAdminAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_account_id: Optional[pulumi.Input[str]] = None) -> 'OrganizationAdminAccount':
        """
        Get an existing OrganizationAdminAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_account_id: AWS account identifier to designate as a delegated administrator for GuardDuty.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrganizationAdminAccountState.__new__(_OrganizationAdminAccountState)

        __props__.__dict__["admin_account_id"] = admin_account_id
        return OrganizationAdminAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminAccountId")
    def admin_account_id(self) -> pulumi.Output[str]:
        """
        AWS account identifier to designate as a delegated administrator for GuardDuty.
        """
        return pulumi.get(self, "admin_account_id")

