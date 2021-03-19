# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Directory']


class Directory(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 connect_settings: Optional[pulumi.Input[pulumi.InputType['DirectoryConnectSettingsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 enable_sso: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 short_name: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vpc_settings: Optional[pulumi.Input[pulumi.InputType['DirectoryVpcSettingsArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Simple or Managed Microsoft directory in AWS Directory Service.

        ## Example Usage
        ### SimpleAD

        ```python
        import pulumi
        import pulumi_aws as aws

        main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
        foo = aws.ec2.Subnet("foo",
            vpc_id=main.id,
            availability_zone="us-west-2a",
            cidr_block="10.0.1.0/24")
        bar_subnet = aws.ec2.Subnet("barSubnet",
            vpc_id=main.id,
            availability_zone="us-west-2b",
            cidr_block="10.0.2.0/24")
        bar_directory = aws.directoryservice.Directory("barDirectory",
            name="corp.notexample.com",
            password="SuperSecretPassw0rd",
            size="Small",
            vpc_settings=aws.directoryservice.DirectoryVpcSettingsArgs(
                vpc_id=main.id,
                subnet_ids=[
                    foo.id,
                    bar_subnet.id,
                ],
            ),
            tags={
                "Project": "foo",
            })
        ```
        ### Microsoft Active Directory (MicrosoftAD)

        ```python
        import pulumi
        import pulumi_aws as aws

        main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
        foo = aws.ec2.Subnet("foo",
            vpc_id=main.id,
            availability_zone="us-west-2a",
            cidr_block="10.0.1.0/24")
        bar_subnet = aws.ec2.Subnet("barSubnet",
            vpc_id=main.id,
            availability_zone="us-west-2b",
            cidr_block="10.0.2.0/24")
        bar_directory = aws.directoryservice.Directory("barDirectory",
            name="corp.notexample.com",
            password="SuperSecretPassw0rd",
            edition="Standard",
            type="MicrosoftAD",
            vpc_settings=aws.directoryservice.DirectoryVpcSettingsArgs(
                vpc_id=main.id,
                subnet_ids=[
                    foo.id,
                    bar_subnet.id,
                ],
            ),
            tags={
                "Project": "foo",
            })
        ```
        ### Microsoft Active Directory Connector (ADConnector)

        ```python
        import pulumi
        import pulumi_aws as aws

        main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
        foo = aws.ec2.Subnet("foo",
            vpc_id=main.id,
            availability_zone="us-west-2a",
            cidr_block="10.0.1.0/24")
        bar = aws.ec2.Subnet("bar",
            vpc_id=main.id,
            availability_zone="us-west-2b",
            cidr_block="10.0.2.0/24")
        connector = aws.directoryservice.Directory("connector",
            name="corp.notexample.com",
            password="SuperSecretPassw0rd",
            size="Small",
            type="ADConnector",
            connect_settings=aws.directoryservice.DirectoryConnectSettingsArgs(
                customer_dns_ips=["A.B.C.D"],
                customer_username="Admin",
                subnet_ids=[
                    foo.id,
                    bar.id,
                ],
                vpc_id=main.id,
            ))
        ```

        ## Import

        DirectoryService directories can be imported using the directory `id`, e.g.

        ```sh
         $ pulumi import aws:directoryservice/directory:Directory sample d-926724cf57
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: The alias for the directory (must be unique amongst all aliases in AWS). Required for `enable_sso`.
        :param pulumi.Input[pulumi.InputType['DirectoryConnectSettingsArgs']] connect_settings: Connector related information about the directory. Fields documented below.
        :param pulumi.Input[str] description: A textual description for the directory.
        :param pulumi.Input[str] edition: The MicrosoftAD edition (`Standard` or `Enterprise`). Defaults to `Enterprise` (applies to MicrosoftAD type only).
        :param pulumi.Input[bool] enable_sso: Whether to enable single-sign on for the directory. Requires `alias`. Defaults to `false`.
        :param pulumi.Input[str] name: The fully qualified name for the directory, such as `corp.example.com`
        :param pulumi.Input[str] password: The password for the directory administrator or connector user.
        :param pulumi.Input[str] short_name: The short name of the directory, such as `CORP`.
        :param pulumi.Input[str] size: The size of the directory (`Small` or `Large` are accepted values).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] type: The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD` are accepted values). Defaults to `SimpleAD`.
        :param pulumi.Input[pulumi.InputType['DirectoryVpcSettingsArgs']] vpc_settings: VPC related information about the directory. Fields documented below.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['alias'] = alias
            __props__['connect_settings'] = connect_settings
            __props__['description'] = description
            __props__['edition'] = edition
            __props__['enable_sso'] = enable_sso
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__['password'] = password
            __props__['short_name'] = short_name
            __props__['size'] = size
            __props__['tags'] = tags
            __props__['type'] = type
            __props__['vpc_settings'] = vpc_settings
            __props__['access_url'] = None
            __props__['dns_ip_addresses'] = None
            __props__['security_group_id'] = None
        super(Directory, __self__).__init__(
            'aws:directoryservice/directory:Directory',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_url: Optional[pulumi.Input[str]] = None,
            alias: Optional[pulumi.Input[str]] = None,
            connect_settings: Optional[pulumi.Input[pulumi.InputType['DirectoryConnectSettingsArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            dns_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            edition: Optional[pulumi.Input[str]] = None,
            enable_sso: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            security_group_id: Optional[pulumi.Input[str]] = None,
            short_name: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            type: Optional[pulumi.Input[str]] = None,
            vpc_settings: Optional[pulumi.Input[pulumi.InputType['DirectoryVpcSettingsArgs']]] = None) -> 'Directory':
        """
        Get an existing Directory resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_url: The access URL for the directory, such as `http://alias.awsapps.com`.
        :param pulumi.Input[str] alias: The alias for the directory (must be unique amongst all aliases in AWS). Required for `enable_sso`.
        :param pulumi.Input[pulumi.InputType['DirectoryConnectSettingsArgs']] connect_settings: Connector related information about the directory. Fields documented below.
        :param pulumi.Input[str] description: A textual description for the directory.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_ip_addresses: A list of IP addresses of the DNS servers for the directory or connector.
        :param pulumi.Input[str] edition: The MicrosoftAD edition (`Standard` or `Enterprise`). Defaults to `Enterprise` (applies to MicrosoftAD type only).
        :param pulumi.Input[bool] enable_sso: Whether to enable single-sign on for the directory. Requires `alias`. Defaults to `false`.
        :param pulumi.Input[str] name: The fully qualified name for the directory, such as `corp.example.com`
        :param pulumi.Input[str] password: The password for the directory administrator or connector user.
        :param pulumi.Input[str] security_group_id: The ID of the security group created by the directory.
        :param pulumi.Input[str] short_name: The short name of the directory, such as `CORP`.
        :param pulumi.Input[str] size: The size of the directory (`Small` or `Large` are accepted values).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] type: The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD` are accepted values). Defaults to `SimpleAD`.
        :param pulumi.Input[pulumi.InputType['DirectoryVpcSettingsArgs']] vpc_settings: VPC related information about the directory. Fields documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_url"] = access_url
        __props__["alias"] = alias
        __props__["connect_settings"] = connect_settings
        __props__["description"] = description
        __props__["dns_ip_addresses"] = dns_ip_addresses
        __props__["edition"] = edition
        __props__["enable_sso"] = enable_sso
        __props__["name"] = name
        __props__["password"] = password
        __props__["security_group_id"] = security_group_id
        __props__["short_name"] = short_name
        __props__["size"] = size
        __props__["tags"] = tags
        __props__["type"] = type
        __props__["vpc_settings"] = vpc_settings
        return Directory(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessUrl")
    def access_url(self) -> pulumi.Output[str]:
        """
        The access URL for the directory, such as `http://alias.awsapps.com`.
        """
        return pulumi.get(self, "access_url")

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Output[str]:
        """
        The alias for the directory (must be unique amongst all aliases in AWS). Required for `enable_sso`.
        """
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter(name="connectSettings")
    def connect_settings(self) -> pulumi.Output[Optional['outputs.DirectoryConnectSettings']]:
        """
        Connector related information about the directory. Fields documented below.
        """
        return pulumi.get(self, "connect_settings")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A textual description for the directory.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dnsIpAddresses")
    def dns_ip_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of IP addresses of the DNS servers for the directory or connector.
        """
        return pulumi.get(self, "dns_ip_addresses")

    @property
    @pulumi.getter
    def edition(self) -> pulumi.Output[str]:
        """
        The MicrosoftAD edition (`Standard` or `Enterprise`). Defaults to `Enterprise` (applies to MicrosoftAD type only).
        """
        return pulumi.get(self, "edition")

    @property
    @pulumi.getter(name="enableSso")
    def enable_sso(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable single-sign on for the directory. Requires `alias`. Defaults to `false`.
        """
        return pulumi.get(self, "enable_sso")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The fully qualified name for the directory, such as `corp.example.com`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        The password for the directory administrator or connector user.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the security group created by the directory.
        """
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter(name="shortName")
    def short_name(self) -> pulumi.Output[str]:
        """
        The short name of the directory, such as `CORP`.
        """
        return pulumi.get(self, "short_name")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[str]:
        """
        The size of the directory (`Small` or `Large` are accepted values).
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD` are accepted values). Defaults to `SimpleAD`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vpcSettings")
    def vpc_settings(self) -> pulumi.Output[Optional['outputs.DirectoryVpcSettings']]:
        """
        VPC related information about the directory. Fields documented below.
        """
        return pulumi.get(self, "vpc_settings")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

