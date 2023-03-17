# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LicenseGrantAccepterArgs', 'LicenseGrantAccepter']

@pulumi.input_type
class LicenseGrantAccepterArgs:
    def __init__(__self__, *,
                 grant_arn: pulumi.Input[str]):
        """
        The set of arguments for constructing a LicenseGrantAccepter resource.
        :param pulumi.Input[str] grant_arn: The ARN of the grant to accept.
        """
        pulumi.set(__self__, "grant_arn", grant_arn)

    @property
    @pulumi.getter(name="grantArn")
    def grant_arn(self) -> pulumi.Input[str]:
        """
        The ARN of the grant to accept.
        """
        return pulumi.get(self, "grant_arn")

    @grant_arn.setter
    def grant_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "grant_arn", value)


@pulumi.input_type
class _LicenseGrantAccepterState:
    def __init__(__self__, *,
                 allowed_operations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 grant_arn: Optional[pulumi.Input[str]] = None,
                 home_region: Optional[pulumi.Input[str]] = None,
                 license_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_arn: Optional[pulumi.Input[str]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LicenseGrantAccepter resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_operations: A list of the allowed operations for the grant.
        :param pulumi.Input[str] grant_arn: The ARN of the grant to accept.
        :param pulumi.Input[str] home_region: The home region for the license.
        :param pulumi.Input[str] license_arn: The ARN of the license for the grant.
        :param pulumi.Input[str] name: The Name of the grant.
        :param pulumi.Input[str] parent_arn: The parent ARN.
        :param pulumi.Input[str] principal: The target account for the grant.
        :param pulumi.Input[str] status: The grant status.
        :param pulumi.Input[str] version: The grant version.
        """
        if allowed_operations is not None:
            pulumi.set(__self__, "allowed_operations", allowed_operations)
        if grant_arn is not None:
            pulumi.set(__self__, "grant_arn", grant_arn)
        if home_region is not None:
            pulumi.set(__self__, "home_region", home_region)
        if license_arn is not None:
            pulumi.set(__self__, "license_arn", license_arn)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent_arn is not None:
            pulumi.set(__self__, "parent_arn", parent_arn)
        if principal is not None:
            pulumi.set(__self__, "principal", principal)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="allowedOperations")
    def allowed_operations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of the allowed operations for the grant.
        """
        return pulumi.get(self, "allowed_operations")

    @allowed_operations.setter
    def allowed_operations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_operations", value)

    @property
    @pulumi.getter(name="grantArn")
    def grant_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the grant to accept.
        """
        return pulumi.get(self, "grant_arn")

    @grant_arn.setter
    def grant_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grant_arn", value)

    @property
    @pulumi.getter(name="homeRegion")
    def home_region(self) -> Optional[pulumi.Input[str]]:
        """
        The home region for the license.
        """
        return pulumi.get(self, "home_region")

    @home_region.setter
    def home_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "home_region", value)

    @property
    @pulumi.getter(name="licenseArn")
    def license_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the license for the grant.
        """
        return pulumi.get(self, "license_arn")

    @license_arn.setter
    def license_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "license_arn", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The Name of the grant.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parentArn")
    def parent_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The parent ARN.
        """
        return pulumi.get(self, "parent_arn")

    @parent_arn.setter
    def parent_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_arn", value)

    @property
    @pulumi.getter
    def principal(self) -> Optional[pulumi.Input[str]]:
        """
        The target account for the grant.
        """
        return pulumi.get(self, "principal")

    @principal.setter
    def principal(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The grant status.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The grant version.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class LicenseGrantAccepter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 grant_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Accepts a License Manager grant. This allows for sharing licenses with other aws accounts.

        ## Import

        `aws_licensemanager_grant_accepter` can be imported using the grant arn.

        ```sh
         $ pulumi import aws:licensemanager/licenseGrantAccepter:LicenseGrantAccepter test arn:aws:license-manager::123456789012:grant:g-1cf9fba4ba2f42dcab11c686c4b4d329
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] grant_arn: The ARN of the grant to accept.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LicenseGrantAccepterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Accepts a License Manager grant. This allows for sharing licenses with other aws accounts.

        ## Import

        `aws_licensemanager_grant_accepter` can be imported using the grant arn.

        ```sh
         $ pulumi import aws:licensemanager/licenseGrantAccepter:LicenseGrantAccepter test arn:aws:license-manager::123456789012:grant:g-1cf9fba4ba2f42dcab11c686c4b4d329
        ```

        :param str resource_name: The name of the resource.
        :param LicenseGrantAccepterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LicenseGrantAccepterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 grant_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LicenseGrantAccepterArgs.__new__(LicenseGrantAccepterArgs)

            if grant_arn is None and not opts.urn:
                raise TypeError("Missing required property 'grant_arn'")
            __props__.__dict__["grant_arn"] = grant_arn
            __props__.__dict__["allowed_operations"] = None
            __props__.__dict__["home_region"] = None
            __props__.__dict__["license_arn"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["parent_arn"] = None
            __props__.__dict__["principal"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["version"] = None
        super(LicenseGrantAccepter, __self__).__init__(
            'aws:licensemanager/licenseGrantAccepter:LicenseGrantAccepter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allowed_operations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            grant_arn: Optional[pulumi.Input[str]] = None,
            home_region: Optional[pulumi.Input[str]] = None,
            license_arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent_arn: Optional[pulumi.Input[str]] = None,
            principal: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'LicenseGrantAccepter':
        """
        Get an existing LicenseGrantAccepter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_operations: A list of the allowed operations for the grant.
        :param pulumi.Input[str] grant_arn: The ARN of the grant to accept.
        :param pulumi.Input[str] home_region: The home region for the license.
        :param pulumi.Input[str] license_arn: The ARN of the license for the grant.
        :param pulumi.Input[str] name: The Name of the grant.
        :param pulumi.Input[str] parent_arn: The parent ARN.
        :param pulumi.Input[str] principal: The target account for the grant.
        :param pulumi.Input[str] status: The grant status.
        :param pulumi.Input[str] version: The grant version.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LicenseGrantAccepterState.__new__(_LicenseGrantAccepterState)

        __props__.__dict__["allowed_operations"] = allowed_operations
        __props__.__dict__["grant_arn"] = grant_arn
        __props__.__dict__["home_region"] = home_region
        __props__.__dict__["license_arn"] = license_arn
        __props__.__dict__["name"] = name
        __props__.__dict__["parent_arn"] = parent_arn
        __props__.__dict__["principal"] = principal
        __props__.__dict__["status"] = status
        __props__.__dict__["version"] = version
        return LicenseGrantAccepter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedOperations")
    def allowed_operations(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of the allowed operations for the grant.
        """
        return pulumi.get(self, "allowed_operations")

    @property
    @pulumi.getter(name="grantArn")
    def grant_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the grant to accept.
        """
        return pulumi.get(self, "grant_arn")

    @property
    @pulumi.getter(name="homeRegion")
    def home_region(self) -> pulumi.Output[str]:
        """
        The home region for the license.
        """
        return pulumi.get(self, "home_region")

    @property
    @pulumi.getter(name="licenseArn")
    def license_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the license for the grant.
        """
        return pulumi.get(self, "license_arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The Name of the grant.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentArn")
    def parent_arn(self) -> pulumi.Output[str]:
        """
        The parent ARN.
        """
        return pulumi.get(self, "parent_arn")

    @property
    @pulumi.getter
    def principal(self) -> pulumi.Output[str]:
        """
        The target account for the grant.
        """
        return pulumi.get(self, "principal")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The grant status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        The grant version.
        """
        return pulumi.get(self, "version")

