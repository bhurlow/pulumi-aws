# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PrimaryContactArgs', 'PrimaryContact']

@pulumi.input_type
class PrimaryContactArgs:
    def __init__(__self__, *,
                 address_line1: pulumi.Input[str],
                 city: pulumi.Input[str],
                 country_code: pulumi.Input[str],
                 full_name: pulumi.Input[str],
                 phone_number: pulumi.Input[str],
                 postal_code: pulumi.Input[str],
                 account_id: Optional[pulumi.Input[str]] = None,
                 address_line2: Optional[pulumi.Input[str]] = None,
                 address_line3: Optional[pulumi.Input[str]] = None,
                 company_name: Optional[pulumi.Input[str]] = None,
                 district_or_county: Optional[pulumi.Input[str]] = None,
                 state_or_region: Optional[pulumi.Input[str]] = None,
                 website_url: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PrimaryContact resource.
        :param pulumi.Input[str] address_line1: The first line of the primary contact address.
        :param pulumi.Input[str] city: The city of the primary contact address.
        :param pulumi.Input[str] country_code: The ISO-3166 two-letter country code for the primary contact address.
        :param pulumi.Input[str] full_name: The full name of the primary contact address.
        :param pulumi.Input[str] phone_number: The phone number of the primary contact information. The number will be validated and, in some countries, checked for activation.
        :param pulumi.Input[str] postal_code: The postal code of the primary contact address.
        :param pulumi.Input[str] account_id: The ID of the target account when managing member accounts. Will manage current user's account by default if omitted.
        :param pulumi.Input[str] address_line2: The second line of the primary contact address, if any.
        :param pulumi.Input[str] address_line3: The third line of the primary contact address, if any.
        :param pulumi.Input[str] company_name: The name of the company associated with the primary contact information, if any.
        :param pulumi.Input[str] district_or_county: The district or county of the primary contact address, if any.
        :param pulumi.Input[str] state_or_region: The state or region of the primary contact address. This field is required in selected countries.
        :param pulumi.Input[str] website_url: The URL of the website associated with the primary contact information, if any.
        """
        pulumi.set(__self__, "address_line1", address_line1)
        pulumi.set(__self__, "city", city)
        pulumi.set(__self__, "country_code", country_code)
        pulumi.set(__self__, "full_name", full_name)
        pulumi.set(__self__, "phone_number", phone_number)
        pulumi.set(__self__, "postal_code", postal_code)
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if address_line2 is not None:
            pulumi.set(__self__, "address_line2", address_line2)
        if address_line3 is not None:
            pulumi.set(__self__, "address_line3", address_line3)
        if company_name is not None:
            pulumi.set(__self__, "company_name", company_name)
        if district_or_county is not None:
            pulumi.set(__self__, "district_or_county", district_or_county)
        if state_or_region is not None:
            pulumi.set(__self__, "state_or_region", state_or_region)
        if website_url is not None:
            pulumi.set(__self__, "website_url", website_url)

    @property
    @pulumi.getter(name="addressLine1")
    def address_line1(self) -> pulumi.Input[str]:
        """
        The first line of the primary contact address.
        """
        return pulumi.get(self, "address_line1")

    @address_line1.setter
    def address_line1(self, value: pulumi.Input[str]):
        pulumi.set(self, "address_line1", value)

    @property
    @pulumi.getter
    def city(self) -> pulumi.Input[str]:
        """
        The city of the primary contact address.
        """
        return pulumi.get(self, "city")

    @city.setter
    def city(self, value: pulumi.Input[str]):
        pulumi.set(self, "city", value)

    @property
    @pulumi.getter(name="countryCode")
    def country_code(self) -> pulumi.Input[str]:
        """
        The ISO-3166 two-letter country code for the primary contact address.
        """
        return pulumi.get(self, "country_code")

    @country_code.setter
    def country_code(self, value: pulumi.Input[str]):
        pulumi.set(self, "country_code", value)

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> pulumi.Input[str]:
        """
        The full name of the primary contact address.
        """
        return pulumi.get(self, "full_name")

    @full_name.setter
    def full_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "full_name", value)

    @property
    @pulumi.getter(name="phoneNumber")
    def phone_number(self) -> pulumi.Input[str]:
        """
        The phone number of the primary contact information. The number will be validated and, in some countries, checked for activation.
        """
        return pulumi.get(self, "phone_number")

    @phone_number.setter
    def phone_number(self, value: pulumi.Input[str]):
        pulumi.set(self, "phone_number", value)

    @property
    @pulumi.getter(name="postalCode")
    def postal_code(self) -> pulumi.Input[str]:
        """
        The postal code of the primary contact address.
        """
        return pulumi.get(self, "postal_code")

    @postal_code.setter
    def postal_code(self, value: pulumi.Input[str]):
        pulumi.set(self, "postal_code", value)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the target account when managing member accounts. Will manage current user's account by default if omitted.
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="addressLine2")
    def address_line2(self) -> Optional[pulumi.Input[str]]:
        """
        The second line of the primary contact address, if any.
        """
        return pulumi.get(self, "address_line2")

    @address_line2.setter
    def address_line2(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_line2", value)

    @property
    @pulumi.getter(name="addressLine3")
    def address_line3(self) -> Optional[pulumi.Input[str]]:
        """
        The third line of the primary contact address, if any.
        """
        return pulumi.get(self, "address_line3")

    @address_line3.setter
    def address_line3(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_line3", value)

    @property
    @pulumi.getter(name="companyName")
    def company_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the company associated with the primary contact information, if any.
        """
        return pulumi.get(self, "company_name")

    @company_name.setter
    def company_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "company_name", value)

    @property
    @pulumi.getter(name="districtOrCounty")
    def district_or_county(self) -> Optional[pulumi.Input[str]]:
        """
        The district or county of the primary contact address, if any.
        """
        return pulumi.get(self, "district_or_county")

    @district_or_county.setter
    def district_or_county(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "district_or_county", value)

    @property
    @pulumi.getter(name="stateOrRegion")
    def state_or_region(self) -> Optional[pulumi.Input[str]]:
        """
        The state or region of the primary contact address. This field is required in selected countries.
        """
        return pulumi.get(self, "state_or_region")

    @state_or_region.setter
    def state_or_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state_or_region", value)

    @property
    @pulumi.getter(name="websiteUrl")
    def website_url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the website associated with the primary contact information, if any.
        """
        return pulumi.get(self, "website_url")

    @website_url.setter
    def website_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "website_url", value)


@pulumi.input_type
class _PrimaryContactState:
    def __init__(__self__, *,
                 account_id: Optional[pulumi.Input[str]] = None,
                 address_line1: Optional[pulumi.Input[str]] = None,
                 address_line2: Optional[pulumi.Input[str]] = None,
                 address_line3: Optional[pulumi.Input[str]] = None,
                 city: Optional[pulumi.Input[str]] = None,
                 company_name: Optional[pulumi.Input[str]] = None,
                 country_code: Optional[pulumi.Input[str]] = None,
                 district_or_county: Optional[pulumi.Input[str]] = None,
                 full_name: Optional[pulumi.Input[str]] = None,
                 phone_number: Optional[pulumi.Input[str]] = None,
                 postal_code: Optional[pulumi.Input[str]] = None,
                 state_or_region: Optional[pulumi.Input[str]] = None,
                 website_url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PrimaryContact resources.
        :param pulumi.Input[str] account_id: The ID of the target account when managing member accounts. Will manage current user's account by default if omitted.
        :param pulumi.Input[str] address_line1: The first line of the primary contact address.
        :param pulumi.Input[str] address_line2: The second line of the primary contact address, if any.
        :param pulumi.Input[str] address_line3: The third line of the primary contact address, if any.
        :param pulumi.Input[str] city: The city of the primary contact address.
        :param pulumi.Input[str] company_name: The name of the company associated with the primary contact information, if any.
        :param pulumi.Input[str] country_code: The ISO-3166 two-letter country code for the primary contact address.
        :param pulumi.Input[str] district_or_county: The district or county of the primary contact address, if any.
        :param pulumi.Input[str] full_name: The full name of the primary contact address.
        :param pulumi.Input[str] phone_number: The phone number of the primary contact information. The number will be validated and, in some countries, checked for activation.
        :param pulumi.Input[str] postal_code: The postal code of the primary contact address.
        :param pulumi.Input[str] state_or_region: The state or region of the primary contact address. This field is required in selected countries.
        :param pulumi.Input[str] website_url: The URL of the website associated with the primary contact information, if any.
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if address_line1 is not None:
            pulumi.set(__self__, "address_line1", address_line1)
        if address_line2 is not None:
            pulumi.set(__self__, "address_line2", address_line2)
        if address_line3 is not None:
            pulumi.set(__self__, "address_line3", address_line3)
        if city is not None:
            pulumi.set(__self__, "city", city)
        if company_name is not None:
            pulumi.set(__self__, "company_name", company_name)
        if country_code is not None:
            pulumi.set(__self__, "country_code", country_code)
        if district_or_county is not None:
            pulumi.set(__self__, "district_or_county", district_or_county)
        if full_name is not None:
            pulumi.set(__self__, "full_name", full_name)
        if phone_number is not None:
            pulumi.set(__self__, "phone_number", phone_number)
        if postal_code is not None:
            pulumi.set(__self__, "postal_code", postal_code)
        if state_or_region is not None:
            pulumi.set(__self__, "state_or_region", state_or_region)
        if website_url is not None:
            pulumi.set(__self__, "website_url", website_url)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the target account when managing member accounts. Will manage current user's account by default if omitted.
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="addressLine1")
    def address_line1(self) -> Optional[pulumi.Input[str]]:
        """
        The first line of the primary contact address.
        """
        return pulumi.get(self, "address_line1")

    @address_line1.setter
    def address_line1(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_line1", value)

    @property
    @pulumi.getter(name="addressLine2")
    def address_line2(self) -> Optional[pulumi.Input[str]]:
        """
        The second line of the primary contact address, if any.
        """
        return pulumi.get(self, "address_line2")

    @address_line2.setter
    def address_line2(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_line2", value)

    @property
    @pulumi.getter(name="addressLine3")
    def address_line3(self) -> Optional[pulumi.Input[str]]:
        """
        The third line of the primary contact address, if any.
        """
        return pulumi.get(self, "address_line3")

    @address_line3.setter
    def address_line3(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_line3", value)

    @property
    @pulumi.getter
    def city(self) -> Optional[pulumi.Input[str]]:
        """
        The city of the primary contact address.
        """
        return pulumi.get(self, "city")

    @city.setter
    def city(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "city", value)

    @property
    @pulumi.getter(name="companyName")
    def company_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the company associated with the primary contact information, if any.
        """
        return pulumi.get(self, "company_name")

    @company_name.setter
    def company_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "company_name", value)

    @property
    @pulumi.getter(name="countryCode")
    def country_code(self) -> Optional[pulumi.Input[str]]:
        """
        The ISO-3166 two-letter country code for the primary contact address.
        """
        return pulumi.get(self, "country_code")

    @country_code.setter
    def country_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "country_code", value)

    @property
    @pulumi.getter(name="districtOrCounty")
    def district_or_county(self) -> Optional[pulumi.Input[str]]:
        """
        The district or county of the primary contact address, if any.
        """
        return pulumi.get(self, "district_or_county")

    @district_or_county.setter
    def district_or_county(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "district_or_county", value)

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> Optional[pulumi.Input[str]]:
        """
        The full name of the primary contact address.
        """
        return pulumi.get(self, "full_name")

    @full_name.setter
    def full_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "full_name", value)

    @property
    @pulumi.getter(name="phoneNumber")
    def phone_number(self) -> Optional[pulumi.Input[str]]:
        """
        The phone number of the primary contact information. The number will be validated and, in some countries, checked for activation.
        """
        return pulumi.get(self, "phone_number")

    @phone_number.setter
    def phone_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phone_number", value)

    @property
    @pulumi.getter(name="postalCode")
    def postal_code(self) -> Optional[pulumi.Input[str]]:
        """
        The postal code of the primary contact address.
        """
        return pulumi.get(self, "postal_code")

    @postal_code.setter
    def postal_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "postal_code", value)

    @property
    @pulumi.getter(name="stateOrRegion")
    def state_or_region(self) -> Optional[pulumi.Input[str]]:
        """
        The state or region of the primary contact address. This field is required in selected countries.
        """
        return pulumi.get(self, "state_or_region")

    @state_or_region.setter
    def state_or_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state_or_region", value)

    @property
    @pulumi.getter(name="websiteUrl")
    def website_url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the website associated with the primary contact information, if any.
        """
        return pulumi.get(self, "website_url")

    @website_url.setter
    def website_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "website_url", value)


class PrimaryContact(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 address_line1: Optional[pulumi.Input[str]] = None,
                 address_line2: Optional[pulumi.Input[str]] = None,
                 address_line3: Optional[pulumi.Input[str]] = None,
                 city: Optional[pulumi.Input[str]] = None,
                 company_name: Optional[pulumi.Input[str]] = None,
                 country_code: Optional[pulumi.Input[str]] = None,
                 district_or_county: Optional[pulumi.Input[str]] = None,
                 full_name: Optional[pulumi.Input[str]] = None,
                 phone_number: Optional[pulumi.Input[str]] = None,
                 postal_code: Optional[pulumi.Input[str]] = None,
                 state_or_region: Optional[pulumi.Input[str]] = None,
                 website_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages the specified primary contact information associated with an AWS Account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.account.PrimaryContact("test",
            address_line1="123 Any Street",
            city="Seattle",
            company_name="Example Corp, Inc.",
            country_code="US",
            district_or_county="King",
            full_name="My Name",
            phone_number="+64211111111",
            postal_code="98101",
            state_or_region="WA",
            website_url="https://www.examplecorp.com")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The ID of the target account when managing member accounts. Will manage current user's account by default if omitted.
        :param pulumi.Input[str] address_line1: The first line of the primary contact address.
        :param pulumi.Input[str] address_line2: The second line of the primary contact address, if any.
        :param pulumi.Input[str] address_line3: The third line of the primary contact address, if any.
        :param pulumi.Input[str] city: The city of the primary contact address.
        :param pulumi.Input[str] company_name: The name of the company associated with the primary contact information, if any.
        :param pulumi.Input[str] country_code: The ISO-3166 two-letter country code for the primary contact address.
        :param pulumi.Input[str] district_or_county: The district or county of the primary contact address, if any.
        :param pulumi.Input[str] full_name: The full name of the primary contact address.
        :param pulumi.Input[str] phone_number: The phone number of the primary contact information. The number will be validated and, in some countries, checked for activation.
        :param pulumi.Input[str] postal_code: The postal code of the primary contact address.
        :param pulumi.Input[str] state_or_region: The state or region of the primary contact address. This field is required in selected countries.
        :param pulumi.Input[str] website_url: The URL of the website associated with the primary contact information, if any.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrimaryContactArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages the specified primary contact information associated with an AWS Account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.account.PrimaryContact("test",
            address_line1="123 Any Street",
            city="Seattle",
            company_name="Example Corp, Inc.",
            country_code="US",
            district_or_county="King",
            full_name="My Name",
            phone_number="+64211111111",
            postal_code="98101",
            state_or_region="WA",
            website_url="https://www.examplecorp.com")
        ```

        :param str resource_name: The name of the resource.
        :param PrimaryContactArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrimaryContactArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 address_line1: Optional[pulumi.Input[str]] = None,
                 address_line2: Optional[pulumi.Input[str]] = None,
                 address_line3: Optional[pulumi.Input[str]] = None,
                 city: Optional[pulumi.Input[str]] = None,
                 company_name: Optional[pulumi.Input[str]] = None,
                 country_code: Optional[pulumi.Input[str]] = None,
                 district_or_county: Optional[pulumi.Input[str]] = None,
                 full_name: Optional[pulumi.Input[str]] = None,
                 phone_number: Optional[pulumi.Input[str]] = None,
                 postal_code: Optional[pulumi.Input[str]] = None,
                 state_or_region: Optional[pulumi.Input[str]] = None,
                 website_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrimaryContactArgs.__new__(PrimaryContactArgs)

            __props__.__dict__["account_id"] = account_id
            if address_line1 is None and not opts.urn:
                raise TypeError("Missing required property 'address_line1'")
            __props__.__dict__["address_line1"] = address_line1
            __props__.__dict__["address_line2"] = address_line2
            __props__.__dict__["address_line3"] = address_line3
            if city is None and not opts.urn:
                raise TypeError("Missing required property 'city'")
            __props__.__dict__["city"] = city
            __props__.__dict__["company_name"] = company_name
            if country_code is None and not opts.urn:
                raise TypeError("Missing required property 'country_code'")
            __props__.__dict__["country_code"] = country_code
            __props__.__dict__["district_or_county"] = district_or_county
            if full_name is None and not opts.urn:
                raise TypeError("Missing required property 'full_name'")
            __props__.__dict__["full_name"] = full_name
            if phone_number is None and not opts.urn:
                raise TypeError("Missing required property 'phone_number'")
            __props__.__dict__["phone_number"] = phone_number
            if postal_code is None and not opts.urn:
                raise TypeError("Missing required property 'postal_code'")
            __props__.__dict__["postal_code"] = postal_code
            __props__.__dict__["state_or_region"] = state_or_region
            __props__.__dict__["website_url"] = website_url
        super(PrimaryContact, __self__).__init__(
            'aws:account/primaryContact:PrimaryContact',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_id: Optional[pulumi.Input[str]] = None,
            address_line1: Optional[pulumi.Input[str]] = None,
            address_line2: Optional[pulumi.Input[str]] = None,
            address_line3: Optional[pulumi.Input[str]] = None,
            city: Optional[pulumi.Input[str]] = None,
            company_name: Optional[pulumi.Input[str]] = None,
            country_code: Optional[pulumi.Input[str]] = None,
            district_or_county: Optional[pulumi.Input[str]] = None,
            full_name: Optional[pulumi.Input[str]] = None,
            phone_number: Optional[pulumi.Input[str]] = None,
            postal_code: Optional[pulumi.Input[str]] = None,
            state_or_region: Optional[pulumi.Input[str]] = None,
            website_url: Optional[pulumi.Input[str]] = None) -> 'PrimaryContact':
        """
        Get an existing PrimaryContact resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The ID of the target account when managing member accounts. Will manage current user's account by default if omitted.
        :param pulumi.Input[str] address_line1: The first line of the primary contact address.
        :param pulumi.Input[str] address_line2: The second line of the primary contact address, if any.
        :param pulumi.Input[str] address_line3: The third line of the primary contact address, if any.
        :param pulumi.Input[str] city: The city of the primary contact address.
        :param pulumi.Input[str] company_name: The name of the company associated with the primary contact information, if any.
        :param pulumi.Input[str] country_code: The ISO-3166 two-letter country code for the primary contact address.
        :param pulumi.Input[str] district_or_county: The district or county of the primary contact address, if any.
        :param pulumi.Input[str] full_name: The full name of the primary contact address.
        :param pulumi.Input[str] phone_number: The phone number of the primary contact information. The number will be validated and, in some countries, checked for activation.
        :param pulumi.Input[str] postal_code: The postal code of the primary contact address.
        :param pulumi.Input[str] state_or_region: The state or region of the primary contact address. This field is required in selected countries.
        :param pulumi.Input[str] website_url: The URL of the website associated with the primary contact information, if any.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrimaryContactState.__new__(_PrimaryContactState)

        __props__.__dict__["account_id"] = account_id
        __props__.__dict__["address_line1"] = address_line1
        __props__.__dict__["address_line2"] = address_line2
        __props__.__dict__["address_line3"] = address_line3
        __props__.__dict__["city"] = city
        __props__.__dict__["company_name"] = company_name
        __props__.__dict__["country_code"] = country_code
        __props__.__dict__["district_or_county"] = district_or_county
        __props__.__dict__["full_name"] = full_name
        __props__.__dict__["phone_number"] = phone_number
        __props__.__dict__["postal_code"] = postal_code
        __props__.__dict__["state_or_region"] = state_or_region
        __props__.__dict__["website_url"] = website_url
        return PrimaryContact(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the target account when managing member accounts. Will manage current user's account by default if omitted.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="addressLine1")
    def address_line1(self) -> pulumi.Output[str]:
        """
        The first line of the primary contact address.
        """
        return pulumi.get(self, "address_line1")

    @property
    @pulumi.getter(name="addressLine2")
    def address_line2(self) -> pulumi.Output[Optional[str]]:
        """
        The second line of the primary contact address, if any.
        """
        return pulumi.get(self, "address_line2")

    @property
    @pulumi.getter(name="addressLine3")
    def address_line3(self) -> pulumi.Output[Optional[str]]:
        """
        The third line of the primary contact address, if any.
        """
        return pulumi.get(self, "address_line3")

    @property
    @pulumi.getter
    def city(self) -> pulumi.Output[str]:
        """
        The city of the primary contact address.
        """
        return pulumi.get(self, "city")

    @property
    @pulumi.getter(name="companyName")
    def company_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the company associated with the primary contact information, if any.
        """
        return pulumi.get(self, "company_name")

    @property
    @pulumi.getter(name="countryCode")
    def country_code(self) -> pulumi.Output[str]:
        """
        The ISO-3166 two-letter country code for the primary contact address.
        """
        return pulumi.get(self, "country_code")

    @property
    @pulumi.getter(name="districtOrCounty")
    def district_or_county(self) -> pulumi.Output[Optional[str]]:
        """
        The district or county of the primary contact address, if any.
        """
        return pulumi.get(self, "district_or_county")

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> pulumi.Output[str]:
        """
        The full name of the primary contact address.
        """
        return pulumi.get(self, "full_name")

    @property
    @pulumi.getter(name="phoneNumber")
    def phone_number(self) -> pulumi.Output[str]:
        """
        The phone number of the primary contact information. The number will be validated and, in some countries, checked for activation.
        """
        return pulumi.get(self, "phone_number")

    @property
    @pulumi.getter(name="postalCode")
    def postal_code(self) -> pulumi.Output[str]:
        """
        The postal code of the primary contact address.
        """
        return pulumi.get(self, "postal_code")

    @property
    @pulumi.getter(name="stateOrRegion")
    def state_or_region(self) -> pulumi.Output[Optional[str]]:
        """
        The state or region of the primary contact address. This field is required in selected countries.
        """
        return pulumi.get(self, "state_or_region")

    @property
    @pulumi.getter(name="websiteUrl")
    def website_url(self) -> pulumi.Output[Optional[str]]:
        """
        The URL of the website associated with the primary contact information, if any.
        """
        return pulumi.get(self, "website_url")

