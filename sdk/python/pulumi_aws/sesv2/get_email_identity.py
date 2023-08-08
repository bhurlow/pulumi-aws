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

__all__ = [
    'GetEmailIdentityResult',
    'AwaitableGetEmailIdentityResult',
    'get_email_identity',
    'get_email_identity_output',
]

@pulumi.output_type
class GetEmailIdentityResult:
    """
    A collection of values returned by getEmailIdentity.
    """
    def __init__(__self__, arn=None, configuration_set_name=None, dkim_signing_attributes=None, email_identity=None, id=None, identity_type=None, tags=None, verified_for_sending_status=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if configuration_set_name and not isinstance(configuration_set_name, str):
            raise TypeError("Expected argument 'configuration_set_name' to be a str")
        pulumi.set(__self__, "configuration_set_name", configuration_set_name)
        if dkim_signing_attributes and not isinstance(dkim_signing_attributes, list):
            raise TypeError("Expected argument 'dkim_signing_attributes' to be a list")
        pulumi.set(__self__, "dkim_signing_attributes", dkim_signing_attributes)
        if email_identity and not isinstance(email_identity, str):
            raise TypeError("Expected argument 'email_identity' to be a str")
        pulumi.set(__self__, "email_identity", email_identity)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity_type and not isinstance(identity_type, str):
            raise TypeError("Expected argument 'identity_type' to be a str")
        pulumi.set(__self__, "identity_type", identity_type)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if verified_for_sending_status and not isinstance(verified_for_sending_status, bool):
            raise TypeError("Expected argument 'verified_for_sending_status' to be a bool")
        pulumi.set(__self__, "verified_for_sending_status", verified_for_sending_status)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        ARN of the Email Identity.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="configurationSetName")
    def configuration_set_name(self) -> str:
        return pulumi.get(self, "configuration_set_name")

    @property
    @pulumi.getter(name="dkimSigningAttributes")
    def dkim_signing_attributes(self) -> Sequence['outputs.GetEmailIdentityDkimSigningAttributeResult']:
        """
        A list of objects that contains at most one element with information about the private key and selector that you want to use to configure DKIM for the identity for Bring Your Own DKIM (BYODKIM) for the identity, or, configures the key length to be used for Easy DKIM.
        """
        return pulumi.get(self, "dkim_signing_attributes")

    @property
    @pulumi.getter(name="emailIdentity")
    def email_identity(self) -> str:
        return pulumi.get(self, "email_identity")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="identityType")
    def identity_type(self) -> str:
        """
        The email identity type. Valid values: `EMAIL_ADDRESS`, `DOMAIN`.
        """
        return pulumi.get(self, "identity_type")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Key-value mapping of resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="verifiedForSendingStatus")
    def verified_for_sending_status(self) -> bool:
        """
        Specifies whether or not the identity is verified.
        """
        return pulumi.get(self, "verified_for_sending_status")


class AwaitableGetEmailIdentityResult(GetEmailIdentityResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEmailIdentityResult(
            arn=self.arn,
            configuration_set_name=self.configuration_set_name,
            dkim_signing_attributes=self.dkim_signing_attributes,
            email_identity=self.email_identity,
            id=self.id,
            identity_type=self.identity_type,
            tags=self.tags,
            verified_for_sending_status=self.verified_for_sending_status)


def get_email_identity(email_identity: Optional[str] = None,
                       tags: Optional[Mapping[str, str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEmailIdentityResult:
    """
    Data source for managing an AWS SESv2 (Simple Email V2) Email Identity.

    ## Example Usage
    ### Basic Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.sesv2.get_email_identity(email_identity="example.com")
    ```


    :param str email_identity: The name of the email identity.
    :param Mapping[str, str] tags: Key-value mapping of resource tags.
    """
    __args__ = dict()
    __args__['emailIdentity'] = email_identity
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws:sesv2/getEmailIdentity:getEmailIdentity', __args__, opts=opts, typ=GetEmailIdentityResult).value

    return AwaitableGetEmailIdentityResult(
        arn=pulumi.get(__ret__, 'arn'),
        configuration_set_name=pulumi.get(__ret__, 'configuration_set_name'),
        dkim_signing_attributes=pulumi.get(__ret__, 'dkim_signing_attributes'),
        email_identity=pulumi.get(__ret__, 'email_identity'),
        id=pulumi.get(__ret__, 'id'),
        identity_type=pulumi.get(__ret__, 'identity_type'),
        tags=pulumi.get(__ret__, 'tags'),
        verified_for_sending_status=pulumi.get(__ret__, 'verified_for_sending_status'))


@_utilities.lift_output_func(get_email_identity)
def get_email_identity_output(email_identity: Optional[pulumi.Input[str]] = None,
                              tags: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEmailIdentityResult]:
    """
    Data source for managing an AWS SESv2 (Simple Email V2) Email Identity.

    ## Example Usage
    ### Basic Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.sesv2.get_email_identity(email_identity="example.com")
    ```


    :param str email_identity: The name of the email identity.
    :param Mapping[str, str] tags: Key-value mapping of resource tags.
    """
    ...
