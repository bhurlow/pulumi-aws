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

__all__ = ['EmailIdentityArgs', 'EmailIdentity']

@pulumi.input_type
class EmailIdentityArgs:
    def __init__(__self__, *,
                 email_identity: pulumi.Input[str],
                 configuration_set_name: Optional[pulumi.Input[str]] = None,
                 dkim_signing_attributes: Optional[pulumi.Input['EmailIdentityDkimSigningAttributesArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a EmailIdentity resource.
        :param pulumi.Input[str] email_identity: The email address or domain to verify.
        :param pulumi.Input[str] configuration_set_name: The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.
        :param pulumi.Input['EmailIdentityDkimSigningAttributesArgs'] dkim_signing_attributes: The configuration of the DKIM authentication settings for an email domain identity.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: (Optional) A map of tags to assign to the service. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "email_identity", email_identity)
        if configuration_set_name is not None:
            pulumi.set(__self__, "configuration_set_name", configuration_set_name)
        if dkim_signing_attributes is not None:
            pulumi.set(__self__, "dkim_signing_attributes", dkim_signing_attributes)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="emailIdentity")
    def email_identity(self) -> pulumi.Input[str]:
        """
        The email address or domain to verify.
        """
        return pulumi.get(self, "email_identity")

    @email_identity.setter
    def email_identity(self, value: pulumi.Input[str]):
        pulumi.set(self, "email_identity", value)

    @property
    @pulumi.getter(name="configurationSetName")
    def configuration_set_name(self) -> Optional[pulumi.Input[str]]:
        """
        The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.
        """
        return pulumi.get(self, "configuration_set_name")

    @configuration_set_name.setter
    def configuration_set_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "configuration_set_name", value)

    @property
    @pulumi.getter(name="dkimSigningAttributes")
    def dkim_signing_attributes(self) -> Optional[pulumi.Input['EmailIdentityDkimSigningAttributesArgs']]:
        """
        The configuration of the DKIM authentication settings for an email domain identity.
        """
        return pulumi.get(self, "dkim_signing_attributes")

    @dkim_signing_attributes.setter
    def dkim_signing_attributes(self, value: Optional[pulumi.Input['EmailIdentityDkimSigningAttributesArgs']]):
        pulumi.set(self, "dkim_signing_attributes", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        (Optional) A map of tags to assign to the service. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _EmailIdentityState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 configuration_set_name: Optional[pulumi.Input[str]] = None,
                 dkim_signing_attributes: Optional[pulumi.Input['EmailIdentityDkimSigningAttributesArgs']] = None,
                 email_identity: Optional[pulumi.Input[str]] = None,
                 identity_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 verified_for_sending_status: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering EmailIdentity resources.
        :param pulumi.Input[str] arn: ARN of the Email Identity.
        :param pulumi.Input[str] configuration_set_name: The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.
        :param pulumi.Input['EmailIdentityDkimSigningAttributesArgs'] dkim_signing_attributes: The configuration of the DKIM authentication settings for an email domain identity.
        :param pulumi.Input[str] email_identity: The email address or domain to verify.
        :param pulumi.Input[str] identity_type: The email identity type. Valid values: `EMAIL_ADDRESS`, `DOMAIN`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: (Optional) A map of tags to assign to the service. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[bool] verified_for_sending_status: Specifies whether or not the identity is verified.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if configuration_set_name is not None:
            pulumi.set(__self__, "configuration_set_name", configuration_set_name)
        if dkim_signing_attributes is not None:
            pulumi.set(__self__, "dkim_signing_attributes", dkim_signing_attributes)
        if email_identity is not None:
            pulumi.set(__self__, "email_identity", email_identity)
        if identity_type is not None:
            pulumi.set(__self__, "identity_type", identity_type)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if verified_for_sending_status is not None:
            pulumi.set(__self__, "verified_for_sending_status", verified_for_sending_status)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the Email Identity.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="configurationSetName")
    def configuration_set_name(self) -> Optional[pulumi.Input[str]]:
        """
        The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.
        """
        return pulumi.get(self, "configuration_set_name")

    @configuration_set_name.setter
    def configuration_set_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "configuration_set_name", value)

    @property
    @pulumi.getter(name="dkimSigningAttributes")
    def dkim_signing_attributes(self) -> Optional[pulumi.Input['EmailIdentityDkimSigningAttributesArgs']]:
        """
        The configuration of the DKIM authentication settings for an email domain identity.
        """
        return pulumi.get(self, "dkim_signing_attributes")

    @dkim_signing_attributes.setter
    def dkim_signing_attributes(self, value: Optional[pulumi.Input['EmailIdentityDkimSigningAttributesArgs']]):
        pulumi.set(self, "dkim_signing_attributes", value)

    @property
    @pulumi.getter(name="emailIdentity")
    def email_identity(self) -> Optional[pulumi.Input[str]]:
        """
        The email address or domain to verify.
        """
        return pulumi.get(self, "email_identity")

    @email_identity.setter
    def email_identity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email_identity", value)

    @property
    @pulumi.getter(name="identityType")
    def identity_type(self) -> Optional[pulumi.Input[str]]:
        """
        The email identity type. Valid values: `EMAIL_ADDRESS`, `DOMAIN`.
        """
        return pulumi.get(self, "identity_type")

    @identity_type.setter
    def identity_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identity_type", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        (Optional) A map of tags to assign to the service. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter(name="verifiedForSendingStatus")
    def verified_for_sending_status(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether or not the identity is verified.
        """
        return pulumi.get(self, "verified_for_sending_status")

    @verified_for_sending_status.setter
    def verified_for_sending_status(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "verified_for_sending_status", value)


class EmailIdentity(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration_set_name: Optional[pulumi.Input[str]] = None,
                 dkim_signing_attributes: Optional[pulumi.Input[pulumi.InputType['EmailIdentityDkimSigningAttributesArgs']]] = None,
                 email_identity: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Resource for managing an AWS SESv2 (Simple Email V2) Email Identity.

        ## Example Usage

        ### Basic Usage
        ### Email Address Identity

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sesv2.EmailIdentity("example", email_identity="testing@example.com")
        ```
        ### Domain Identity

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sesv2.EmailIdentity("example", email_identity="example.com")
        ```
        ### Configuration Set

        ```python
        import pulumi
        import pulumi_aws as aws

        example_configuration_set = aws.sesv2.ConfigurationSet("exampleConfigurationSet", configuration_set_name="example")
        example_email_identity = aws.sesv2.EmailIdentity("exampleEmailIdentity",
            email_identity="example.com",
            configuration_set_name=example_configuration_set.configuration_set_name)
        ```

        ## Import

        SESv2 (Simple Email V2) Email Identity can be imported using the `email_identity`, e.g.,

        ```sh
         $ pulumi import aws:sesv2/emailIdentity:EmailIdentity example example.com
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] configuration_set_name: The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.
        :param pulumi.Input[pulumi.InputType['EmailIdentityDkimSigningAttributesArgs']] dkim_signing_attributes: The configuration of the DKIM authentication settings for an email domain identity.
        :param pulumi.Input[str] email_identity: The email address or domain to verify.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: (Optional) A map of tags to assign to the service. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EmailIdentityArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for managing an AWS SESv2 (Simple Email V2) Email Identity.

        ## Example Usage

        ### Basic Usage
        ### Email Address Identity

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sesv2.EmailIdentity("example", email_identity="testing@example.com")
        ```
        ### Domain Identity

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sesv2.EmailIdentity("example", email_identity="example.com")
        ```
        ### Configuration Set

        ```python
        import pulumi
        import pulumi_aws as aws

        example_configuration_set = aws.sesv2.ConfigurationSet("exampleConfigurationSet", configuration_set_name="example")
        example_email_identity = aws.sesv2.EmailIdentity("exampleEmailIdentity",
            email_identity="example.com",
            configuration_set_name=example_configuration_set.configuration_set_name)
        ```

        ## Import

        SESv2 (Simple Email V2) Email Identity can be imported using the `email_identity`, e.g.,

        ```sh
         $ pulumi import aws:sesv2/emailIdentity:EmailIdentity example example.com
        ```

        :param str resource_name: The name of the resource.
        :param EmailIdentityArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EmailIdentityArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration_set_name: Optional[pulumi.Input[str]] = None,
                 dkim_signing_attributes: Optional[pulumi.Input[pulumi.InputType['EmailIdentityDkimSigningAttributesArgs']]] = None,
                 email_identity: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EmailIdentityArgs.__new__(EmailIdentityArgs)

            __props__.__dict__["configuration_set_name"] = configuration_set_name
            __props__.__dict__["dkim_signing_attributes"] = dkim_signing_attributes
            if email_identity is None and not opts.urn:
                raise TypeError("Missing required property 'email_identity'")
            __props__.__dict__["email_identity"] = email_identity
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["identity_type"] = None
            __props__.__dict__["tags_all"] = None
            __props__.__dict__["verified_for_sending_status"] = None
        super(EmailIdentity, __self__).__init__(
            'aws:sesv2/emailIdentity:EmailIdentity',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            configuration_set_name: Optional[pulumi.Input[str]] = None,
            dkim_signing_attributes: Optional[pulumi.Input[pulumi.InputType['EmailIdentityDkimSigningAttributesArgs']]] = None,
            email_identity: Optional[pulumi.Input[str]] = None,
            identity_type: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            verified_for_sending_status: Optional[pulumi.Input[bool]] = None) -> 'EmailIdentity':
        """
        Get an existing EmailIdentity resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: ARN of the Email Identity.
        :param pulumi.Input[str] configuration_set_name: The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.
        :param pulumi.Input[pulumi.InputType['EmailIdentityDkimSigningAttributesArgs']] dkim_signing_attributes: The configuration of the DKIM authentication settings for an email domain identity.
        :param pulumi.Input[str] email_identity: The email address or domain to verify.
        :param pulumi.Input[str] identity_type: The email identity type. Valid values: `EMAIL_ADDRESS`, `DOMAIN`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: (Optional) A map of tags to assign to the service. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[bool] verified_for_sending_status: Specifies whether or not the identity is verified.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EmailIdentityState.__new__(_EmailIdentityState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["configuration_set_name"] = configuration_set_name
        __props__.__dict__["dkim_signing_attributes"] = dkim_signing_attributes
        __props__.__dict__["email_identity"] = email_identity
        __props__.__dict__["identity_type"] = identity_type
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["verified_for_sending_status"] = verified_for_sending_status
        return EmailIdentity(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN of the Email Identity.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="configurationSetName")
    def configuration_set_name(self) -> pulumi.Output[Optional[str]]:
        """
        The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.
        """
        return pulumi.get(self, "configuration_set_name")

    @property
    @pulumi.getter(name="dkimSigningAttributes")
    def dkim_signing_attributes(self) -> pulumi.Output['outputs.EmailIdentityDkimSigningAttributes']:
        """
        The configuration of the DKIM authentication settings for an email domain identity.
        """
        return pulumi.get(self, "dkim_signing_attributes")

    @property
    @pulumi.getter(name="emailIdentity")
    def email_identity(self) -> pulumi.Output[str]:
        """
        The email address or domain to verify.
        """
        return pulumi.get(self, "email_identity")

    @property
    @pulumi.getter(name="identityType")
    def identity_type(self) -> pulumi.Output[str]:
        """
        The email identity type. Valid values: `EMAIL_ADDRESS`, `DOMAIN`.
        """
        return pulumi.get(self, "identity_type")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        (Optional) A map of tags to assign to the service. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter(name="verifiedForSendingStatus")
    def verified_for_sending_status(self) -> pulumi.Output[bool]:
        """
        Specifies whether or not the identity is verified.
        """
        return pulumi.get(self, "verified_for_sending_status")

