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

__all__ = ['FileSystemAssociationArgs', 'FileSystemAssociation']

@pulumi.input_type
class FileSystemAssociationArgs:
    def __init__(__self__, *,
                 gateway_arn: pulumi.Input[str],
                 location_arn: pulumi.Input[str],
                 password: pulumi.Input[str],
                 username: pulumi.Input[str],
                 audit_destination_arn: Optional[pulumi.Input[str]] = None,
                 cache_attributes: Optional[pulumi.Input['FileSystemAssociationCacheAttributesArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a FileSystemAssociation resource.
        :param pulumi.Input[str] gateway_arn: The Amazon Resource Name (ARN) of the gateway.
        :param pulumi.Input[str] location_arn: The Amazon Resource Name (ARN) of the Amazon FSx file system to associate with the FSx File Gateway.
        :param pulumi.Input[str] password: The password of the user credential.
        :param pulumi.Input[str] username: The user name of the user credential that has permission to access the root share of the Amazon FSx file system. The user account must belong to the Amazon FSx delegated admin user group.
        :param pulumi.Input[str] audit_destination_arn: The Amazon Resource Name (ARN) of the storage used for the audit logs.
        :param pulumi.Input['FileSystemAssociationCacheAttributesArgs'] cache_attributes: Refresh cache information. see Cache Attributes for more details.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "gateway_arn", gateway_arn)
        pulumi.set(__self__, "location_arn", location_arn)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "username", username)
        if audit_destination_arn is not None:
            pulumi.set(__self__, "audit_destination_arn", audit_destination_arn)
        if cache_attributes is not None:
            pulumi.set(__self__, "cache_attributes", cache_attributes)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="gatewayArn")
    def gateway_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) of the gateway.
        """
        return pulumi.get(self, "gateway_arn")

    @gateway_arn.setter
    def gateway_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "gateway_arn", value)

    @property
    @pulumi.getter(name="locationArn")
    def location_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) of the Amazon FSx file system to associate with the FSx File Gateway.
        """
        return pulumi.get(self, "location_arn")

    @location_arn.setter
    def location_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "location_arn", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        The password of the user credential.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        The user name of the user credential that has permission to access the root share of the Amazon FSx file system. The user account must belong to the Amazon FSx delegated admin user group.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="auditDestinationArn")
    def audit_destination_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the storage used for the audit logs.
        """
        return pulumi.get(self, "audit_destination_arn")

    @audit_destination_arn.setter
    def audit_destination_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "audit_destination_arn", value)

    @property
    @pulumi.getter(name="cacheAttributes")
    def cache_attributes(self) -> Optional[pulumi.Input['FileSystemAssociationCacheAttributesArgs']]:
        """
        Refresh cache information. see Cache Attributes for more details.
        """
        return pulumi.get(self, "cache_attributes")

    @cache_attributes.setter
    def cache_attributes(self, value: Optional[pulumi.Input['FileSystemAssociationCacheAttributesArgs']]):
        pulumi.set(self, "cache_attributes", value)

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
class _FileSystemAssociationState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 audit_destination_arn: Optional[pulumi.Input[str]] = None,
                 cache_attributes: Optional[pulumi.Input['FileSystemAssociationCacheAttributesArgs']] = None,
                 gateway_arn: Optional[pulumi.Input[str]] = None,
                 location_arn: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FileSystemAssociation resources.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the newly created file system association.
        :param pulumi.Input[str] audit_destination_arn: The Amazon Resource Name (ARN) of the storage used for the audit logs.
        :param pulumi.Input['FileSystemAssociationCacheAttributesArgs'] cache_attributes: Refresh cache information. see Cache Attributes for more details.
        :param pulumi.Input[str] gateway_arn: The Amazon Resource Name (ARN) of the gateway.
        :param pulumi.Input[str] location_arn: The Amazon Resource Name (ARN) of the Amazon FSx file system to associate with the FSx File Gateway.
        :param pulumi.Input[str] password: The password of the user credential.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[str] username: The user name of the user credential that has permission to access the root share of the Amazon FSx file system. The user account must belong to the Amazon FSx delegated admin user group.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if audit_destination_arn is not None:
            pulumi.set(__self__, "audit_destination_arn", audit_destination_arn)
        if cache_attributes is not None:
            pulumi.set(__self__, "cache_attributes", cache_attributes)
        if gateway_arn is not None:
            pulumi.set(__self__, "gateway_arn", gateway_arn)
        if location_arn is not None:
            pulumi.set(__self__, "location_arn", location_arn)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
            pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the newly created file system association.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="auditDestinationArn")
    def audit_destination_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the storage used for the audit logs.
        """
        return pulumi.get(self, "audit_destination_arn")

    @audit_destination_arn.setter
    def audit_destination_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "audit_destination_arn", value)

    @property
    @pulumi.getter(name="cacheAttributes")
    def cache_attributes(self) -> Optional[pulumi.Input['FileSystemAssociationCacheAttributesArgs']]:
        """
        Refresh cache information. see Cache Attributes for more details.
        """
        return pulumi.get(self, "cache_attributes")

    @cache_attributes.setter
    def cache_attributes(self, value: Optional[pulumi.Input['FileSystemAssociationCacheAttributesArgs']]):
        pulumi.set(self, "cache_attributes", value)

    @property
    @pulumi.getter(name="gatewayArn")
    def gateway_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the gateway.
        """
        return pulumi.get(self, "gateway_arn")

    @gateway_arn.setter
    def gateway_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_arn", value)

    @property
    @pulumi.getter(name="locationArn")
    def location_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the Amazon FSx file system to associate with the FSx File Gateway.
        """
        return pulumi.get(self, "location_arn")

    @location_arn.setter
    def location_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location_arn", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password of the user credential.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

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
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The user name of the user credential that has permission to access the root share of the Amazon FSx file system. The user account must belong to the Amazon FSx delegated admin user group.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class FileSystemAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audit_destination_arn: Optional[pulumi.Input[str]] = None,
                 cache_attributes: Optional[pulumi.Input[pulumi.InputType['FileSystemAssociationCacheAttributesArgs']]] = None,
                 gateway_arn: Optional[pulumi.Input[str]] = None,
                 location_arn: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Associate an Amazon FSx file system with the FSx File Gateway. After the association process is complete, the file shares on the Amazon FSx file system are available for access through the gateway. This operation only supports the FSx File Gateway type.

        [FSx File Gateway requirements](https://docs.aws.amazon.com/filegateway/latest/filefsxw/Requirements.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.storagegateway.FileSystemAssociation("example",
            gateway_arn=aws_storagegateway_gateway["example"]["arn"],
            location_arn=aws_fsx_windows_file_system["example"]["arn"],
            username="Admin",
            password="avoid-plaintext-passwords",
            audit_destination_arn=aws_s3_bucket["example"]["arn"])
        ```
        ## Required Services Example

        ```python
        import pulumi
        import pulumi_aws as aws

        aws_service_storagegateway_ami_files3_latest = aws.ssm.get_parameter(name="/aws/service/storagegateway/ami/FILE_S3/latest")
        test_instance = aws.ec2.Instance("testInstance",
            ami=aws_service_storagegateway_ami_files3_latest.value,
            associate_public_ip_address=True,
            instance_type=aws.ec2/instancetype.InstanceType(data["aws_ec2_instance_type_offering"]["available"]["instance_type"]),
            vpc_security_group_ids=[aws_security_group["test"]["id"]],
            subnet_id=aws_subnet["test"][0]["id"],
            opts=pulumi.ResourceOptions(depends_on=[
                    aws_route["test"],
                    aws_vpc_dhcp_options_association["test"],
                ]))
        test_gateway = aws.storagegateway.Gateway("testGateway",
            gateway_ip_address=test_instance.public_ip,
            gateway_name="test-sgw",
            gateway_timezone="GMT",
            gateway_type="FILE_FSX_SMB",
            smb_active_directory_settings=aws.storagegateway.GatewaySmbActiveDirectorySettingsArgs(
                domain_name=aws_directory_service_directory["test"]["name"],
                password=aws_directory_service_directory["test"]["password"],
                username="Admin",
            ))
        test_windows_file_system = aws.fsx.WindowsFileSystem("testWindowsFileSystem",
            active_directory_id=aws_directory_service_directory["test"]["id"],
            security_group_ids=[aws_security_group["test"]["id"]],
            skip_final_backup=True,
            storage_capacity=32,
            subnet_ids=[aws_subnet["test"][0]["id"]],
            throughput_capacity=8)
        fsx = aws.storagegateway.FileSystemAssociation("fsx",
            gateway_arn=test_gateway.arn,
            location_arn=test_windows_file_system.arn,
            username="Admin",
            password=aws_directory_service_directory["test"]["password"],
            cache_attributes=aws.storagegateway.FileSystemAssociationCacheAttributesArgs(
                cache_stale_timeout_in_seconds=400,
            ),
            audit_destination_arn=aws_cloudwatch_log_group["test"]["arn"])
        ```

        ## Import

        Using `pulumi import`, import `aws_storagegateway_file_system_association` using the FSx file system association Amazon Resource Name (ARN). For example:

        ```sh
         $ pulumi import aws:storagegateway/fileSystemAssociation:FileSystemAssociation example arn:aws:storagegateway:us-east-1:123456789012:fs-association/fsa-0DA347732FDB40125
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] audit_destination_arn: The Amazon Resource Name (ARN) of the storage used for the audit logs.
        :param pulumi.Input[pulumi.InputType['FileSystemAssociationCacheAttributesArgs']] cache_attributes: Refresh cache information. see Cache Attributes for more details.
        :param pulumi.Input[str] gateway_arn: The Amazon Resource Name (ARN) of the gateway.
        :param pulumi.Input[str] location_arn: The Amazon Resource Name (ARN) of the Amazon FSx file system to associate with the FSx File Gateway.
        :param pulumi.Input[str] password: The password of the user credential.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] username: The user name of the user credential that has permission to access the root share of the Amazon FSx file system. The user account must belong to the Amazon FSx delegated admin user group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FileSystemAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Associate an Amazon FSx file system with the FSx File Gateway. After the association process is complete, the file shares on the Amazon FSx file system are available for access through the gateway. This operation only supports the FSx File Gateway type.

        [FSx File Gateway requirements](https://docs.aws.amazon.com/filegateway/latest/filefsxw/Requirements.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.storagegateway.FileSystemAssociation("example",
            gateway_arn=aws_storagegateway_gateway["example"]["arn"],
            location_arn=aws_fsx_windows_file_system["example"]["arn"],
            username="Admin",
            password="avoid-plaintext-passwords",
            audit_destination_arn=aws_s3_bucket["example"]["arn"])
        ```
        ## Required Services Example

        ```python
        import pulumi
        import pulumi_aws as aws

        aws_service_storagegateway_ami_files3_latest = aws.ssm.get_parameter(name="/aws/service/storagegateway/ami/FILE_S3/latest")
        test_instance = aws.ec2.Instance("testInstance",
            ami=aws_service_storagegateway_ami_files3_latest.value,
            associate_public_ip_address=True,
            instance_type=aws.ec2/instancetype.InstanceType(data["aws_ec2_instance_type_offering"]["available"]["instance_type"]),
            vpc_security_group_ids=[aws_security_group["test"]["id"]],
            subnet_id=aws_subnet["test"][0]["id"],
            opts=pulumi.ResourceOptions(depends_on=[
                    aws_route["test"],
                    aws_vpc_dhcp_options_association["test"],
                ]))
        test_gateway = aws.storagegateway.Gateway("testGateway",
            gateway_ip_address=test_instance.public_ip,
            gateway_name="test-sgw",
            gateway_timezone="GMT",
            gateway_type="FILE_FSX_SMB",
            smb_active_directory_settings=aws.storagegateway.GatewaySmbActiveDirectorySettingsArgs(
                domain_name=aws_directory_service_directory["test"]["name"],
                password=aws_directory_service_directory["test"]["password"],
                username="Admin",
            ))
        test_windows_file_system = aws.fsx.WindowsFileSystem("testWindowsFileSystem",
            active_directory_id=aws_directory_service_directory["test"]["id"],
            security_group_ids=[aws_security_group["test"]["id"]],
            skip_final_backup=True,
            storage_capacity=32,
            subnet_ids=[aws_subnet["test"][0]["id"]],
            throughput_capacity=8)
        fsx = aws.storagegateway.FileSystemAssociation("fsx",
            gateway_arn=test_gateway.arn,
            location_arn=test_windows_file_system.arn,
            username="Admin",
            password=aws_directory_service_directory["test"]["password"],
            cache_attributes=aws.storagegateway.FileSystemAssociationCacheAttributesArgs(
                cache_stale_timeout_in_seconds=400,
            ),
            audit_destination_arn=aws_cloudwatch_log_group["test"]["arn"])
        ```

        ## Import

        Using `pulumi import`, import `aws_storagegateway_file_system_association` using the FSx file system association Amazon Resource Name (ARN). For example:

        ```sh
         $ pulumi import aws:storagegateway/fileSystemAssociation:FileSystemAssociation example arn:aws:storagegateway:us-east-1:123456789012:fs-association/fsa-0DA347732FDB40125
        ```

        :param str resource_name: The name of the resource.
        :param FileSystemAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FileSystemAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audit_destination_arn: Optional[pulumi.Input[str]] = None,
                 cache_attributes: Optional[pulumi.Input[pulumi.InputType['FileSystemAssociationCacheAttributesArgs']]] = None,
                 gateway_arn: Optional[pulumi.Input[str]] = None,
                 location_arn: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FileSystemAssociationArgs.__new__(FileSystemAssociationArgs)

            __props__.__dict__["audit_destination_arn"] = audit_destination_arn
            __props__.__dict__["cache_attributes"] = cache_attributes
            if gateway_arn is None and not opts.urn:
                raise TypeError("Missing required property 'gateway_arn'")
            __props__.__dict__["gateway_arn"] = gateway_arn
            if location_arn is None and not opts.urn:
                raise TypeError("Missing required property 'location_arn'")
            __props__.__dict__["location_arn"] = location_arn
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["tags"] = tags
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password", "tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(FileSystemAssociation, __self__).__init__(
            'aws:storagegateway/fileSystemAssociation:FileSystemAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            audit_destination_arn: Optional[pulumi.Input[str]] = None,
            cache_attributes: Optional[pulumi.Input[pulumi.InputType['FileSystemAssociationCacheAttributesArgs']]] = None,
            gateway_arn: Optional[pulumi.Input[str]] = None,
            location_arn: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'FileSystemAssociation':
        """
        Get an existing FileSystemAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the newly created file system association.
        :param pulumi.Input[str] audit_destination_arn: The Amazon Resource Name (ARN) of the storage used for the audit logs.
        :param pulumi.Input[pulumi.InputType['FileSystemAssociationCacheAttributesArgs']] cache_attributes: Refresh cache information. see Cache Attributes for more details.
        :param pulumi.Input[str] gateway_arn: The Amazon Resource Name (ARN) of the gateway.
        :param pulumi.Input[str] location_arn: The Amazon Resource Name (ARN) of the Amazon FSx file system to associate with the FSx File Gateway.
        :param pulumi.Input[str] password: The password of the user credential.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[str] username: The user name of the user credential that has permission to access the root share of the Amazon FSx file system. The user account must belong to the Amazon FSx delegated admin user group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FileSystemAssociationState.__new__(_FileSystemAssociationState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["audit_destination_arn"] = audit_destination_arn
        __props__.__dict__["cache_attributes"] = cache_attributes
        __props__.__dict__["gateway_arn"] = gateway_arn
        __props__.__dict__["location_arn"] = location_arn
        __props__.__dict__["password"] = password
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["username"] = username
        return FileSystemAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the newly created file system association.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="auditDestinationArn")
    def audit_destination_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The Amazon Resource Name (ARN) of the storage used for the audit logs.
        """
        return pulumi.get(self, "audit_destination_arn")

    @property
    @pulumi.getter(name="cacheAttributes")
    def cache_attributes(self) -> pulumi.Output[Optional['outputs.FileSystemAssociationCacheAttributes']]:
        """
        Refresh cache information. see Cache Attributes for more details.
        """
        return pulumi.get(self, "cache_attributes")

    @property
    @pulumi.getter(name="gatewayArn")
    def gateway_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the gateway.
        """
        return pulumi.get(self, "gateway_arn")

    @property
    @pulumi.getter(name="locationArn")
    def location_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the Amazon FSx file system to associate with the FSx File Gateway.
        """
        return pulumi.get(self, "location_arn")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        The password of the user credential.
        """
        return pulumi.get(self, "password")

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
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        The user name of the user credential that has permission to access the root share of the Amazon FSx file system. The user account must belong to the Amazon FSx delegated admin user group.
        """
        return pulumi.get(self, "username")

