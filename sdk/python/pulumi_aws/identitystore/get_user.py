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

__all__ = [
    'GetUserResult',
    'AwaitableGetUserResult',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, addresses=None, alternate_identifier=None, display_name=None, emails=None, external_ids=None, id=None, identity_store_id=None, locale=None, names=None, nickname=None, phone_numbers=None, preferred_language=None, profile_url=None, timezone=None, title=None, user_id=None, user_name=None, user_type=None):
        if addresses and not isinstance(addresses, list):
            raise TypeError("Expected argument 'addresses' to be a list")
        pulumi.set(__self__, "addresses", addresses)
        if alternate_identifier and not isinstance(alternate_identifier, dict):
            raise TypeError("Expected argument 'alternate_identifier' to be a dict")
        pulumi.set(__self__, "alternate_identifier", alternate_identifier)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if emails and not isinstance(emails, list):
            raise TypeError("Expected argument 'emails' to be a list")
        pulumi.set(__self__, "emails", emails)
        if external_ids and not isinstance(external_ids, list):
            raise TypeError("Expected argument 'external_ids' to be a list")
        pulumi.set(__self__, "external_ids", external_ids)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity_store_id and not isinstance(identity_store_id, str):
            raise TypeError("Expected argument 'identity_store_id' to be a str")
        pulumi.set(__self__, "identity_store_id", identity_store_id)
        if locale and not isinstance(locale, str):
            raise TypeError("Expected argument 'locale' to be a str")
        pulumi.set(__self__, "locale", locale)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if nickname and not isinstance(nickname, str):
            raise TypeError("Expected argument 'nickname' to be a str")
        pulumi.set(__self__, "nickname", nickname)
        if phone_numbers and not isinstance(phone_numbers, list):
            raise TypeError("Expected argument 'phone_numbers' to be a list")
        pulumi.set(__self__, "phone_numbers", phone_numbers)
        if preferred_language and not isinstance(preferred_language, str):
            raise TypeError("Expected argument 'preferred_language' to be a str")
        pulumi.set(__self__, "preferred_language", preferred_language)
        if profile_url and not isinstance(profile_url, str):
            raise TypeError("Expected argument 'profile_url' to be a str")
        pulumi.set(__self__, "profile_url", profile_url)
        if timezone and not isinstance(timezone, str):
            raise TypeError("Expected argument 'timezone' to be a str")
        pulumi.set(__self__, "timezone", timezone)
        if title and not isinstance(title, str):
            raise TypeError("Expected argument 'title' to be a str")
        pulumi.set(__self__, "title", title)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)
        if user_name and not isinstance(user_name, str):
            raise TypeError("Expected argument 'user_name' to be a str")
        pulumi.set(__self__, "user_name", user_name)
        if user_type and not isinstance(user_type, str):
            raise TypeError("Expected argument 'user_type' to be a str")
        pulumi.set(__self__, "user_type", user_type)

    @property
    @pulumi.getter
    def addresses(self) -> Sequence['outputs.GetUserAddressResult']:
        """
        List of details about the user's address.
        """
        return pulumi.get(self, "addresses")

    @property
    @pulumi.getter(name="alternateIdentifier")
    def alternate_identifier(self) -> Optional['outputs.GetUserAlternateIdentifierResult']:
        return pulumi.get(self, "alternate_identifier")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The name that is typically displayed when the user is referenced.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def emails(self) -> Sequence['outputs.GetUserEmailResult']:
        """
        List of details about the user's email.
        """
        return pulumi.get(self, "emails")

    @property
    @pulumi.getter(name="externalIds")
    def external_ids(self) -> Sequence['outputs.GetUserExternalIdResult']:
        """
        List of identifiers issued to this resource by an external identity provider.
        """
        return pulumi.get(self, "external_ids")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="identityStoreId")
    def identity_store_id(self) -> str:
        return pulumi.get(self, "identity_store_id")

    @property
    @pulumi.getter
    def locale(self) -> str:
        """
        The user's geographical region or location.
        """
        return pulumi.get(self, "locale")

    @property
    @pulumi.getter
    def names(self) -> Sequence['outputs.GetUserNameResult']:
        """
        Details about the user's full name.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter
    def nickname(self) -> str:
        """
        An alternate name for the user.
        """
        return pulumi.get(self, "nickname")

    @property
    @pulumi.getter(name="phoneNumbers")
    def phone_numbers(self) -> Sequence['outputs.GetUserPhoneNumberResult']:
        """
        List of details about the user's phone number.
        """
        return pulumi.get(self, "phone_numbers")

    @property
    @pulumi.getter(name="preferredLanguage")
    def preferred_language(self) -> str:
        """
        The preferred language of the user.
        """
        return pulumi.get(self, "preferred_language")

    @property
    @pulumi.getter(name="profileUrl")
    def profile_url(self) -> str:
        """
        An URL that may be associated with the user.
        """
        return pulumi.get(self, "profile_url")

    @property
    @pulumi.getter
    def timezone(self) -> str:
        """
        The user's time zone.
        """
        return pulumi.get(self, "timezone")

    @property
    @pulumi.getter
    def title(self) -> str:
        """
        The user's title.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        """
        User's user name value.
        """
        return pulumi.get(self, "user_name")

    @property
    @pulumi.getter(name="userType")
    def user_type(self) -> str:
        """
        The user type.
        """
        return pulumi.get(self, "user_type")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            addresses=self.addresses,
            alternate_identifier=self.alternate_identifier,
            display_name=self.display_name,
            emails=self.emails,
            external_ids=self.external_ids,
            id=self.id,
            identity_store_id=self.identity_store_id,
            locale=self.locale,
            names=self.names,
            nickname=self.nickname,
            phone_numbers=self.phone_numbers,
            preferred_language=self.preferred_language,
            profile_url=self.profile_url,
            timezone=self.timezone,
            title=self.title,
            user_id=self.user_id,
            user_name=self.user_name,
            user_type=self.user_type)


def get_user(alternate_identifier: Optional[pulumi.InputType['GetUserAlternateIdentifierArgs']] = None,
             identity_store_id: Optional[str] = None,
             user_id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    Use this data source to get an Identity Store User.


    :param pulumi.InputType['GetUserAlternateIdentifierArgs'] alternate_identifier: A unique identifier for a user or group that is not the primary identifier. Conflicts with `user_id`. Detailed below.
    :param str identity_store_id: Identity Store ID associated with the Single Sign-On Instance.
           
           The following arguments are optional:
    :param str user_id: The identifier for a user in the Identity Store.
           
           > Exactly one of the above arguments must be provided.
    """
    __args__ = dict()
    __args__['alternateIdentifier'] = alternate_identifier
    __args__['identityStoreId'] = identity_store_id
    __args__['userId'] = user_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws:identitystore/getUser:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        addresses=pulumi.get(__ret__, 'addresses'),
        alternate_identifier=pulumi.get(__ret__, 'alternate_identifier'),
        display_name=pulumi.get(__ret__, 'display_name'),
        emails=pulumi.get(__ret__, 'emails'),
        external_ids=pulumi.get(__ret__, 'external_ids'),
        id=pulumi.get(__ret__, 'id'),
        identity_store_id=pulumi.get(__ret__, 'identity_store_id'),
        locale=pulumi.get(__ret__, 'locale'),
        names=pulumi.get(__ret__, 'names'),
        nickname=pulumi.get(__ret__, 'nickname'),
        phone_numbers=pulumi.get(__ret__, 'phone_numbers'),
        preferred_language=pulumi.get(__ret__, 'preferred_language'),
        profile_url=pulumi.get(__ret__, 'profile_url'),
        timezone=pulumi.get(__ret__, 'timezone'),
        title=pulumi.get(__ret__, 'title'),
        user_id=pulumi.get(__ret__, 'user_id'),
        user_name=pulumi.get(__ret__, 'user_name'),
        user_type=pulumi.get(__ret__, 'user_type'))


@_utilities.lift_output_func(get_user)
def get_user_output(alternate_identifier: Optional[pulumi.Input[Optional[pulumi.InputType['GetUserAlternateIdentifierArgs']]]] = None,
                    identity_store_id: Optional[pulumi.Input[str]] = None,
                    user_id: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserResult]:
    """
    Use this data source to get an Identity Store User.


    :param pulumi.InputType['GetUserAlternateIdentifierArgs'] alternate_identifier: A unique identifier for a user or group that is not the primary identifier. Conflicts with `user_id`. Detailed below.
    :param str identity_store_id: Identity Store ID associated with the Single Sign-On Instance.
           
           The following arguments are optional:
    :param str user_id: The identifier for a user in the Identity Store.
           
           > Exactly one of the above arguments must be provided.
    """
    ...
