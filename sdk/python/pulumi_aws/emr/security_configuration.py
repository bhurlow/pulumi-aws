# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecurityConfigurationArgs', 'SecurityConfiguration']

@pulumi.input_type
class SecurityConfigurationArgs:
    def __init__(__self__, *,
                 configuration: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SecurityConfiguration resource.
        :param pulumi.Input[str] configuration: A JSON formatted Security Configuration
        :param pulumi.Input[str] name: The name of the EMR Security Configuration. By default generated by this provider.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified
               prefix. Conflicts with `name`.
        """
        pulumi.set(__self__, "configuration", configuration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if name_prefix is not None:
            pulumi.set(__self__, "name_prefix", name_prefix)

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Input[str]:
        """
        A JSON formatted Security Configuration
        """
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: pulumi.Input[str]):
        pulumi.set(self, "configuration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the EMR Security Configuration. By default generated by this provider.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Creates a unique name beginning with the specified
        prefix. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @name_prefix.setter
    def name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_prefix", value)


@pulumi.input_type
class _SecurityConfigurationState:
    def __init__(__self__, *,
                 configuration: Optional[pulumi.Input[str]] = None,
                 creation_date: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecurityConfiguration resources.
        :param pulumi.Input[str] configuration: A JSON formatted Security Configuration
        :param pulumi.Input[str] creation_date: Date the Security Configuration was created
        :param pulumi.Input[str] name: The name of the EMR Security Configuration. By default generated by this provider.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified
               prefix. Conflicts with `name`.
        """
        if configuration is not None:
            pulumi.set(__self__, "configuration", configuration)
        if creation_date is not None:
            pulumi.set(__self__, "creation_date", creation_date)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if name_prefix is not None:
            pulumi.set(__self__, "name_prefix", name_prefix)

    @property
    @pulumi.getter
    def configuration(self) -> Optional[pulumi.Input[str]]:
        """
        A JSON formatted Security Configuration
        """
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "configuration", value)

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> Optional[pulumi.Input[str]]:
        """
        Date the Security Configuration was created
        """
        return pulumi.get(self, "creation_date")

    @creation_date.setter
    def creation_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_date", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the EMR Security Configuration. By default generated by this provider.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Creates a unique name beginning with the specified
        prefix. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @name_prefix.setter
    def name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_prefix", value)


class SecurityConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage AWS EMR Security Configurations

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        foo = aws.emr.SecurityConfiguration("foo", configuration=\"\"\"{
          "EncryptionConfiguration": {
            "AtRestEncryptionConfiguration": {
              "S3EncryptionConfiguration": {
                "EncryptionMode": "SSE-S3"
              },
              "LocalDiskEncryptionConfiguration": {
                "EncryptionKeyProviderType": "AwsKms",
                "AwsKmsKey": "arn:aws:kms:us-west-2:187416307283:alias/my_emr_test_key"
              }
            },
            "EnableInTransitEncryption": false,
            "EnableAtRestEncryption": true
          }
        }

        \"\"\")
        ```

        ## Import

        terraform import {

         to = aws_emr_security_configuration.sc

         id = "example-sc-name" } Using `pulumi import`, import EMR Security Configurations using the `name`. For exampleconsole % pulumi import aws_emr_security_configuration.sc example-sc-name

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] configuration: A JSON formatted Security Configuration
        :param pulumi.Input[str] name: The name of the EMR Security Configuration. By default generated by this provider.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified
               prefix. Conflicts with `name`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecurityConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage AWS EMR Security Configurations

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        foo = aws.emr.SecurityConfiguration("foo", configuration=\"\"\"{
          "EncryptionConfiguration": {
            "AtRestEncryptionConfiguration": {
              "S3EncryptionConfiguration": {
                "EncryptionMode": "SSE-S3"
              },
              "LocalDiskEncryptionConfiguration": {
                "EncryptionKeyProviderType": "AwsKms",
                "AwsKmsKey": "arn:aws:kms:us-west-2:187416307283:alias/my_emr_test_key"
              }
            },
            "EnableInTransitEncryption": false,
            "EnableAtRestEncryption": true
          }
        }

        \"\"\")
        ```

        ## Import

        terraform import {

         to = aws_emr_security_configuration.sc

         id = "example-sc-name" } Using `pulumi import`, import EMR Security Configurations using the `name`. For exampleconsole % pulumi import aws_emr_security_configuration.sc example-sc-name

        :param str resource_name: The name of the resource.
        :param SecurityConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecurityConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecurityConfigurationArgs.__new__(SecurityConfigurationArgs)

            if configuration is None and not opts.urn:
                raise TypeError("Missing required property 'configuration'")
            __props__.__dict__["configuration"] = configuration
            __props__.__dict__["name"] = name
            __props__.__dict__["name_prefix"] = name_prefix
            __props__.__dict__["creation_date"] = None
        super(SecurityConfiguration, __self__).__init__(
            'aws:emr/securityConfiguration:SecurityConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            configuration: Optional[pulumi.Input[str]] = None,
            creation_date: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None) -> 'SecurityConfiguration':
        """
        Get an existing SecurityConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] configuration: A JSON formatted Security Configuration
        :param pulumi.Input[str] creation_date: Date the Security Configuration was created
        :param pulumi.Input[str] name: The name of the EMR Security Configuration. By default generated by this provider.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified
               prefix. Conflicts with `name`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecurityConfigurationState.__new__(_SecurityConfigurationState)

        __props__.__dict__["configuration"] = configuration
        __props__.__dict__["creation_date"] = creation_date
        __props__.__dict__["name"] = name
        __props__.__dict__["name_prefix"] = name_prefix
        return SecurityConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Output[str]:
        """
        A JSON formatted Security Configuration
        """
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> pulumi.Output[str]:
        """
        Date the Security Configuration was created
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the EMR Security Configuration. By default generated by this provider.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        Creates a unique name beginning with the specified
        prefix. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

