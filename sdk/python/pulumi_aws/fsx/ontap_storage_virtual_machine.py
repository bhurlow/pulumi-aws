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

__all__ = ['OntapStorageVirtualMachineArgs', 'OntapStorageVirtualMachine']

@pulumi.input_type
class OntapStorageVirtualMachineArgs:
    def __init__(__self__, *,
                 file_system_id: pulumi.Input[str],
                 active_directory_configuration: Optional[pulumi.Input['OntapStorageVirtualMachineActiveDirectoryConfigurationArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 root_volume_security_style: Optional[pulumi.Input[str]] = None,
                 svm_admin_password: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a OntapStorageVirtualMachine resource.
        :param pulumi.Input[str] file_system_id: The ID of the Amazon FSx ONTAP File System that this SVM will be created on.
        :param pulumi.Input['OntapStorageVirtualMachineActiveDirectoryConfigurationArgs'] active_directory_configuration: Configuration block that Amazon FSx uses to join the FSx ONTAP Storage Virtual Machine(SVM) to your Microsoft Active Directory (AD) directory. Detailed below.
        :param pulumi.Input[str] name: The name of the SVM. You can use a maximum of 47 alphanumeric characters, plus the underscore (_) special character.
        :param pulumi.Input[str] root_volume_security_style: Specifies the root volume security style, Valid values are `UNIX`, `NTFS`, and `MIXED`. All volumes created under this SVM will inherit the root security style unless the security style is specified on the volume. Default value is `UNIX`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the storage virtual machine. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "file_system_id", file_system_id)
        if active_directory_configuration is not None:
            pulumi.set(__self__, "active_directory_configuration", active_directory_configuration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if root_volume_security_style is not None:
            pulumi.set(__self__, "root_volume_security_style", root_volume_security_style)
        if svm_admin_password is not None:
            pulumi.set(__self__, "svm_admin_password", svm_admin_password)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> pulumi.Input[str]:
        """
        The ID of the Amazon FSx ONTAP File System that this SVM will be created on.
        """
        return pulumi.get(self, "file_system_id")

    @file_system_id.setter
    def file_system_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_system_id", value)

    @property
    @pulumi.getter(name="activeDirectoryConfiguration")
    def active_directory_configuration(self) -> Optional[pulumi.Input['OntapStorageVirtualMachineActiveDirectoryConfigurationArgs']]:
        """
        Configuration block that Amazon FSx uses to join the FSx ONTAP Storage Virtual Machine(SVM) to your Microsoft Active Directory (AD) directory. Detailed below.
        """
        return pulumi.get(self, "active_directory_configuration")

    @active_directory_configuration.setter
    def active_directory_configuration(self, value: Optional[pulumi.Input['OntapStorageVirtualMachineActiveDirectoryConfigurationArgs']]):
        pulumi.set(self, "active_directory_configuration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the SVM. You can use a maximum of 47 alphanumeric characters, plus the underscore (_) special character.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="rootVolumeSecurityStyle")
    def root_volume_security_style(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the root volume security style, Valid values are `UNIX`, `NTFS`, and `MIXED`. All volumes created under this SVM will inherit the root security style unless the security style is specified on the volume. Default value is `UNIX`.
        """
        return pulumi.get(self, "root_volume_security_style")

    @root_volume_security_style.setter
    def root_volume_security_style(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "root_volume_security_style", value)

    @property
    @pulumi.getter(name="svmAdminPassword")
    def svm_admin_password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "svm_admin_password")

    @svm_admin_password.setter
    def svm_admin_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "svm_admin_password", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the storage virtual machine. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _OntapStorageVirtualMachineState:
    def __init__(__self__, *,
                 active_directory_configuration: Optional[pulumi.Input['OntapStorageVirtualMachineActiveDirectoryConfigurationArgs']] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 endpoints: Optional[pulumi.Input[Sequence[pulumi.Input['OntapStorageVirtualMachineEndpointArgs']]]] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 root_volume_security_style: Optional[pulumi.Input[str]] = None,
                 subtype: Optional[pulumi.Input[str]] = None,
                 svm_admin_password: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 uuid: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OntapStorageVirtualMachine resources.
        :param pulumi.Input['OntapStorageVirtualMachineActiveDirectoryConfigurationArgs'] active_directory_configuration: Configuration block that Amazon FSx uses to join the FSx ONTAP Storage Virtual Machine(SVM) to your Microsoft Active Directory (AD) directory. Detailed below.
        :param pulumi.Input[str] arn: Amazon Resource Name of the storage virtual machine.
        :param pulumi.Input[Sequence[pulumi.Input['OntapStorageVirtualMachineEndpointArgs']]] endpoints: The endpoints that are used to access data or to manage the storage virtual machine using the NetApp ONTAP CLI, REST API, or NetApp SnapMirror. See Endpoints below.
        :param pulumi.Input[str] file_system_id: The ID of the Amazon FSx ONTAP File System that this SVM will be created on.
        :param pulumi.Input[str] name: The name of the SVM. You can use a maximum of 47 alphanumeric characters, plus the underscore (_) special character.
        :param pulumi.Input[str] root_volume_security_style: Specifies the root volume security style, Valid values are `UNIX`, `NTFS`, and `MIXED`. All volumes created under this SVM will inherit the root security style unless the security style is specified on the volume. Default value is `UNIX`.
        :param pulumi.Input[str] subtype: Describes the SVM's subtype, e.g. `DEFAULT`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the storage virtual machine. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[str] uuid: The SVM's UUID (universally unique identifier).
        """
        if active_directory_configuration is not None:
            pulumi.set(__self__, "active_directory_configuration", active_directory_configuration)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if endpoints is not None:
            pulumi.set(__self__, "endpoints", endpoints)
        if file_system_id is not None:
            pulumi.set(__self__, "file_system_id", file_system_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if root_volume_security_style is not None:
            pulumi.set(__self__, "root_volume_security_style", root_volume_security_style)
        if subtype is not None:
            pulumi.set(__self__, "subtype", subtype)
        if svm_admin_password is not None:
            pulumi.set(__self__, "svm_admin_password", svm_admin_password)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter(name="activeDirectoryConfiguration")
    def active_directory_configuration(self) -> Optional[pulumi.Input['OntapStorageVirtualMachineActiveDirectoryConfigurationArgs']]:
        """
        Configuration block that Amazon FSx uses to join the FSx ONTAP Storage Virtual Machine(SVM) to your Microsoft Active Directory (AD) directory. Detailed below.
        """
        return pulumi.get(self, "active_directory_configuration")

    @active_directory_configuration.setter
    def active_directory_configuration(self, value: Optional[pulumi.Input['OntapStorageVirtualMachineActiveDirectoryConfigurationArgs']]):
        pulumi.set(self, "active_directory_configuration", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name of the storage virtual machine.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def endpoints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['OntapStorageVirtualMachineEndpointArgs']]]]:
        """
        The endpoints that are used to access data or to manage the storage virtual machine using the NetApp ONTAP CLI, REST API, or NetApp SnapMirror. See Endpoints below.
        """
        return pulumi.get(self, "endpoints")

    @endpoints.setter
    def endpoints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['OntapStorageVirtualMachineEndpointArgs']]]]):
        pulumi.set(self, "endpoints", value)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Amazon FSx ONTAP File System that this SVM will be created on.
        """
        return pulumi.get(self, "file_system_id")

    @file_system_id.setter
    def file_system_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_system_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the SVM. You can use a maximum of 47 alphanumeric characters, plus the underscore (_) special character.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="rootVolumeSecurityStyle")
    def root_volume_security_style(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the root volume security style, Valid values are `UNIX`, `NTFS`, and `MIXED`. All volumes created under this SVM will inherit the root security style unless the security style is specified on the volume. Default value is `UNIX`.
        """
        return pulumi.get(self, "root_volume_security_style")

    @root_volume_security_style.setter
    def root_volume_security_style(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "root_volume_security_style", value)

    @property
    @pulumi.getter
    def subtype(self) -> Optional[pulumi.Input[str]]:
        """
        Describes the SVM's subtype, e.g. `DEFAULT`
        """
        return pulumi.get(self, "subtype")

    @subtype.setter
    def subtype(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subtype", value)

    @property
    @pulumi.getter(name="svmAdminPassword")
    def svm_admin_password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "svm_admin_password")

    @svm_admin_password.setter
    def svm_admin_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "svm_admin_password", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the storage virtual machine. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        """
        The SVM's UUID (universally unique identifier).
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)


class OntapStorageVirtualMachine(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active_directory_configuration: Optional[pulumi.Input[pulumi.InputType['OntapStorageVirtualMachineActiveDirectoryConfigurationArgs']]] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 root_volume_security_style: Optional[pulumi.Input[str]] = None,
                 svm_admin_password: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manages a FSx Storage Virtual Machine.
        See the [FSx ONTAP User Guide](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-svms.html) for more information.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.fsx.OntapStorageVirtualMachine("test", file_system_id=aws_fsx_ontap_file_system["test"]["id"])
        ```
        ### Using a Self-Managed Microsoft Active Directory

        Additional information for using AWS Directory Service with ONTAP File Systems can be found in the [FSx ONTAP Guide](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/self-managed-AD.html).

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.fsx.OntapStorageVirtualMachine("test",
            file_system_id=aws_fsx_ontap_file_system["test"]["id"],
            active_directory_configuration=aws.fsx.OntapStorageVirtualMachineActiveDirectoryConfigurationArgs(
                netbios_name="mysvm",
                self_managed_active_directory_configuration=aws.fsx.OntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationArgs(
                    dns_ips=[
                        "10.0.0.111",
                        "10.0.0.222",
                    ],
                    domain_name="corp.example.com",
                    password="avoid-plaintext-passwords",
                    username="Admin",
                ),
            ))
        ```

        ## Import

        FSx Storage Virtual Machine can be imported using the `id`, e.g.,

        ```sh
         $ pulumi import aws:fsx/ontapStorageVirtualMachine:OntapStorageVirtualMachine example svm-12345678abcdef123
        ```

         Certain resource arguments, like `svm_admin_password` and the `self_managed_active_directory` configuation block `password`, do not have a FSx API method for reading the information after creation. If these arguments are set in the provider configuration on an imported resource, the provider will always show a difference. To workaround this behavior, either omit the argument from the provider configuration or use `ignore_changes` to hide the difference, e.g., terraform resource "aws_fsx_ontap_storage_virtual_machine" "example" {

        # ... other configuration ...

         svm_admin_password = "avoid-plaintext-passwords"

        # There is no FSx API for reading svm_admin_password

         lifecycle {

         ignore_changes = [svm_admin_password]

         } }

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['OntapStorageVirtualMachineActiveDirectoryConfigurationArgs']] active_directory_configuration: Configuration block that Amazon FSx uses to join the FSx ONTAP Storage Virtual Machine(SVM) to your Microsoft Active Directory (AD) directory. Detailed below.
        :param pulumi.Input[str] file_system_id: The ID of the Amazon FSx ONTAP File System that this SVM will be created on.
        :param pulumi.Input[str] name: The name of the SVM. You can use a maximum of 47 alphanumeric characters, plus the underscore (_) special character.
        :param pulumi.Input[str] root_volume_security_style: Specifies the root volume security style, Valid values are `UNIX`, `NTFS`, and `MIXED`. All volumes created under this SVM will inherit the root security style unless the security style is specified on the volume. Default value is `UNIX`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the storage virtual machine. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OntapStorageVirtualMachineArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a FSx Storage Virtual Machine.
        See the [FSx ONTAP User Guide](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-svms.html) for more information.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.fsx.OntapStorageVirtualMachine("test", file_system_id=aws_fsx_ontap_file_system["test"]["id"])
        ```
        ### Using a Self-Managed Microsoft Active Directory

        Additional information for using AWS Directory Service with ONTAP File Systems can be found in the [FSx ONTAP Guide](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/self-managed-AD.html).

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.fsx.OntapStorageVirtualMachine("test",
            file_system_id=aws_fsx_ontap_file_system["test"]["id"],
            active_directory_configuration=aws.fsx.OntapStorageVirtualMachineActiveDirectoryConfigurationArgs(
                netbios_name="mysvm",
                self_managed_active_directory_configuration=aws.fsx.OntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationArgs(
                    dns_ips=[
                        "10.0.0.111",
                        "10.0.0.222",
                    ],
                    domain_name="corp.example.com",
                    password="avoid-plaintext-passwords",
                    username="Admin",
                ),
            ))
        ```

        ## Import

        FSx Storage Virtual Machine can be imported using the `id`, e.g.,

        ```sh
         $ pulumi import aws:fsx/ontapStorageVirtualMachine:OntapStorageVirtualMachine example svm-12345678abcdef123
        ```

         Certain resource arguments, like `svm_admin_password` and the `self_managed_active_directory` configuation block `password`, do not have a FSx API method for reading the information after creation. If these arguments are set in the provider configuration on an imported resource, the provider will always show a difference. To workaround this behavior, either omit the argument from the provider configuration or use `ignore_changes` to hide the difference, e.g., terraform resource "aws_fsx_ontap_storage_virtual_machine" "example" {

        # ... other configuration ...

         svm_admin_password = "avoid-plaintext-passwords"

        # There is no FSx API for reading svm_admin_password

         lifecycle {

         ignore_changes = [svm_admin_password]

         } }

        :param str resource_name: The name of the resource.
        :param OntapStorageVirtualMachineArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OntapStorageVirtualMachineArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active_directory_configuration: Optional[pulumi.Input[pulumi.InputType['OntapStorageVirtualMachineActiveDirectoryConfigurationArgs']]] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 root_volume_security_style: Optional[pulumi.Input[str]] = None,
                 svm_admin_password: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OntapStorageVirtualMachineArgs.__new__(OntapStorageVirtualMachineArgs)

            __props__.__dict__["active_directory_configuration"] = active_directory_configuration
            if file_system_id is None and not opts.urn:
                raise TypeError("Missing required property 'file_system_id'")
            __props__.__dict__["file_system_id"] = file_system_id
            __props__.__dict__["name"] = name
            __props__.__dict__["root_volume_security_style"] = root_volume_security_style
            __props__.__dict__["svm_admin_password"] = None if svm_admin_password is None else pulumi.Output.secret(svm_admin_password)
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["endpoints"] = None
            __props__.__dict__["subtype"] = None
            __props__.__dict__["tags_all"] = None
            __props__.__dict__["uuid"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["svmAdminPassword"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(OntapStorageVirtualMachine, __self__).__init__(
            'aws:fsx/ontapStorageVirtualMachine:OntapStorageVirtualMachine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active_directory_configuration: Optional[pulumi.Input[pulumi.InputType['OntapStorageVirtualMachineActiveDirectoryConfigurationArgs']]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OntapStorageVirtualMachineEndpointArgs']]]]] = None,
            file_system_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            root_volume_security_style: Optional[pulumi.Input[str]] = None,
            subtype: Optional[pulumi.Input[str]] = None,
            svm_admin_password: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            uuid: Optional[pulumi.Input[str]] = None) -> 'OntapStorageVirtualMachine':
        """
        Get an existing OntapStorageVirtualMachine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['OntapStorageVirtualMachineActiveDirectoryConfigurationArgs']] active_directory_configuration: Configuration block that Amazon FSx uses to join the FSx ONTAP Storage Virtual Machine(SVM) to your Microsoft Active Directory (AD) directory. Detailed below.
        :param pulumi.Input[str] arn: Amazon Resource Name of the storage virtual machine.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OntapStorageVirtualMachineEndpointArgs']]]] endpoints: The endpoints that are used to access data or to manage the storage virtual machine using the NetApp ONTAP CLI, REST API, or NetApp SnapMirror. See Endpoints below.
        :param pulumi.Input[str] file_system_id: The ID of the Amazon FSx ONTAP File System that this SVM will be created on.
        :param pulumi.Input[str] name: The name of the SVM. You can use a maximum of 47 alphanumeric characters, plus the underscore (_) special character.
        :param pulumi.Input[str] root_volume_security_style: Specifies the root volume security style, Valid values are `UNIX`, `NTFS`, and `MIXED`. All volumes created under this SVM will inherit the root security style unless the security style is specified on the volume. Default value is `UNIX`.
        :param pulumi.Input[str] subtype: Describes the SVM's subtype, e.g. `DEFAULT`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the storage virtual machine. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[str] uuid: The SVM's UUID (universally unique identifier).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OntapStorageVirtualMachineState.__new__(_OntapStorageVirtualMachineState)

        __props__.__dict__["active_directory_configuration"] = active_directory_configuration
        __props__.__dict__["arn"] = arn
        __props__.__dict__["endpoints"] = endpoints
        __props__.__dict__["file_system_id"] = file_system_id
        __props__.__dict__["name"] = name
        __props__.__dict__["root_volume_security_style"] = root_volume_security_style
        __props__.__dict__["subtype"] = subtype
        __props__.__dict__["svm_admin_password"] = svm_admin_password
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["uuid"] = uuid
        return OntapStorageVirtualMachine(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activeDirectoryConfiguration")
    def active_directory_configuration(self) -> pulumi.Output[Optional['outputs.OntapStorageVirtualMachineActiveDirectoryConfiguration']]:
        """
        Configuration block that Amazon FSx uses to join the FSx ONTAP Storage Virtual Machine(SVM) to your Microsoft Active Directory (AD) directory. Detailed below.
        """
        return pulumi.get(self, "active_directory_configuration")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name of the storage virtual machine.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def endpoints(self) -> pulumi.Output[Sequence['outputs.OntapStorageVirtualMachineEndpoint']]:
        """
        The endpoints that are used to access data or to manage the storage virtual machine using the NetApp ONTAP CLI, REST API, or NetApp SnapMirror. See Endpoints below.
        """
        return pulumi.get(self, "endpoints")

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> pulumi.Output[str]:
        """
        The ID of the Amazon FSx ONTAP File System that this SVM will be created on.
        """
        return pulumi.get(self, "file_system_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the SVM. You can use a maximum of 47 alphanumeric characters, plus the underscore (_) special character.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="rootVolumeSecurityStyle")
    def root_volume_security_style(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the root volume security style, Valid values are `UNIX`, `NTFS`, and `MIXED`. All volumes created under this SVM will inherit the root security style unless the security style is specified on the volume. Default value is `UNIX`.
        """
        return pulumi.get(self, "root_volume_security_style")

    @property
    @pulumi.getter
    def subtype(self) -> pulumi.Output[str]:
        """
        Describes the SVM's subtype, e.g. `DEFAULT`
        """
        return pulumi.get(self, "subtype")

    @property
    @pulumi.getter(name="svmAdminPassword")
    def svm_admin_password(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "svm_admin_password")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the storage virtual machine. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
    @pulumi.getter
    def uuid(self) -> pulumi.Output[str]:
        """
        The SVM's UUID (universally unique identifier).
        """
        return pulumi.get(self, "uuid")

