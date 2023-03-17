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
    'GetReceivedLicenseResult',
    'AwaitableGetReceivedLicenseResult',
    'get_received_license',
    'get_received_license_output',
]

@pulumi.output_type
class GetReceivedLicenseResult:
    """
    A collection of values returned by getReceivedLicense.
    """
    def __init__(__self__, beneficiary=None, consumption_configurations=None, create_time=None, entitlements=None, home_region=None, id=None, issuers=None, license_arn=None, license_metadatas=None, license_name=None, product_name=None, product_sku=None, received_metadatas=None, status=None, validities=None, version=None):
        if beneficiary and not isinstance(beneficiary, str):
            raise TypeError("Expected argument 'beneficiary' to be a str")
        pulumi.set(__self__, "beneficiary", beneficiary)
        if consumption_configurations and not isinstance(consumption_configurations, list):
            raise TypeError("Expected argument 'consumption_configurations' to be a list")
        pulumi.set(__self__, "consumption_configurations", consumption_configurations)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if entitlements and not isinstance(entitlements, list):
            raise TypeError("Expected argument 'entitlements' to be a list")
        pulumi.set(__self__, "entitlements", entitlements)
        if home_region and not isinstance(home_region, str):
            raise TypeError("Expected argument 'home_region' to be a str")
        pulumi.set(__self__, "home_region", home_region)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if issuers and not isinstance(issuers, list):
            raise TypeError("Expected argument 'issuers' to be a list")
        pulumi.set(__self__, "issuers", issuers)
        if license_arn and not isinstance(license_arn, str):
            raise TypeError("Expected argument 'license_arn' to be a str")
        pulumi.set(__self__, "license_arn", license_arn)
        if license_metadatas and not isinstance(license_metadatas, list):
            raise TypeError("Expected argument 'license_metadatas' to be a list")
        pulumi.set(__self__, "license_metadatas", license_metadatas)
        if license_name and not isinstance(license_name, str):
            raise TypeError("Expected argument 'license_name' to be a str")
        pulumi.set(__self__, "license_name", license_name)
        if product_name and not isinstance(product_name, str):
            raise TypeError("Expected argument 'product_name' to be a str")
        pulumi.set(__self__, "product_name", product_name)
        if product_sku and not isinstance(product_sku, str):
            raise TypeError("Expected argument 'product_sku' to be a str")
        pulumi.set(__self__, "product_sku", product_sku)
        if received_metadatas and not isinstance(received_metadatas, list):
            raise TypeError("Expected argument 'received_metadatas' to be a list")
        pulumi.set(__self__, "received_metadatas", received_metadatas)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if validities and not isinstance(validities, list):
            raise TypeError("Expected argument 'validities' to be a list")
        pulumi.set(__self__, "validities", validities)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def beneficiary(self) -> str:
        """
        Granted license beneficiary. This is in the form of the ARN of the root user of the account.
        """
        return pulumi.get(self, "beneficiary")

    @property
    @pulumi.getter(name="consumptionConfigurations")
    def consumption_configurations(self) -> Sequence['outputs.GetReceivedLicenseConsumptionConfigurationResult']:
        """
        Configuration for consumption of the license. Detailed below
        """
        return pulumi.get(self, "consumption_configurations")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Creation time of the granted license in RFC 3339 format.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def entitlements(self) -> Sequence['outputs.GetReceivedLicenseEntitlementResult']:
        """
        License entitlements. Detailed below
        """
        return pulumi.get(self, "entitlements")

    @property
    @pulumi.getter(name="homeRegion")
    def home_region(self) -> str:
        """
        Home Region of the granted license.
        """
        return pulumi.get(self, "home_region")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def issuers(self) -> Sequence['outputs.GetReceivedLicenseIssuerResult']:
        """
        Granted license issuer. Detailed below
        """
        return pulumi.get(self, "issuers")

    @property
    @pulumi.getter(name="licenseArn")
    def license_arn(self) -> str:
        """
        Amazon Resource Name (ARN) of the license.
        """
        return pulumi.get(self, "license_arn")

    @property
    @pulumi.getter(name="licenseMetadatas")
    def license_metadatas(self) -> Sequence['outputs.GetReceivedLicenseLicenseMetadataResult']:
        """
        Granted license metadata. This is in the form of a set of all meta data. Detailed below
        """
        return pulumi.get(self, "license_metadatas")

    @property
    @pulumi.getter(name="licenseName")
    def license_name(self) -> str:
        """
        License name.
        """
        return pulumi.get(self, "license_name")

    @property
    @pulumi.getter(name="productName")
    def product_name(self) -> str:
        """
        Product name.
        * `product_sku ` - Product SKU.
        """
        return pulumi.get(self, "product_name")

    @property
    @pulumi.getter(name="productSku")
    def product_sku(self) -> str:
        return pulumi.get(self, "product_sku")

    @property
    @pulumi.getter(name="receivedMetadatas")
    def received_metadatas(self) -> Sequence['outputs.GetReceivedLicenseReceivedMetadataResult']:
        """
        Granted license received metadata. Detailed below
        """
        return pulumi.get(self, "received_metadatas")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Granted license status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def validities(self) -> Sequence['outputs.GetReceivedLicenseValidityResult']:
        """
        Date and time range during which the granted license is valid, in ISO8601-UTC format. Detailed below
        """
        return pulumi.get(self, "validities")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        Version of the granted license.
        """
        return pulumi.get(self, "version")


class AwaitableGetReceivedLicenseResult(GetReceivedLicenseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReceivedLicenseResult(
            beneficiary=self.beneficiary,
            consumption_configurations=self.consumption_configurations,
            create_time=self.create_time,
            entitlements=self.entitlements,
            home_region=self.home_region,
            id=self.id,
            issuers=self.issuers,
            license_arn=self.license_arn,
            license_metadatas=self.license_metadatas,
            license_name=self.license_name,
            product_name=self.product_name,
            product_sku=self.product_sku,
            received_metadatas=self.received_metadatas,
            status=self.status,
            validities=self.validities,
            version=self.version)


def get_received_license(license_arn: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReceivedLicenseResult:
    """
    This resource can be used to get data on a received license using an ARN. This can be helpful for pulling in data on a license from the AWS marketplace and sharing that license with another account.

    ## Example Usage

    The following shows getting the received license data using and ARN.

    ```python
    import pulumi
    import pulumi_aws as aws

    test = aws.licensemanager.get_received_license(license_arn="arn:aws:license-manager::111111111111:license:l-ecbaa94eb71a4830b6d7e49268fecaa0")
    ```


    :param str license_arn: The ARN of the received license you want data for.
    """
    __args__ = dict()
    __args__['licenseArn'] = license_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws:licensemanager/getReceivedLicense:getReceivedLicense', __args__, opts=opts, typ=GetReceivedLicenseResult).value

    return AwaitableGetReceivedLicenseResult(
        beneficiary=__ret__.beneficiary,
        consumption_configurations=__ret__.consumption_configurations,
        create_time=__ret__.create_time,
        entitlements=__ret__.entitlements,
        home_region=__ret__.home_region,
        id=__ret__.id,
        issuers=__ret__.issuers,
        license_arn=__ret__.license_arn,
        license_metadatas=__ret__.license_metadatas,
        license_name=__ret__.license_name,
        product_name=__ret__.product_name,
        product_sku=__ret__.product_sku,
        received_metadatas=__ret__.received_metadatas,
        status=__ret__.status,
        validities=__ret__.validities,
        version=__ret__.version)


@_utilities.lift_output_func(get_received_license)
def get_received_license_output(license_arn: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetReceivedLicenseResult]:
    """
    This resource can be used to get data on a received license using an ARN. This can be helpful for pulling in data on a license from the AWS marketplace and sharing that license with another account.

    ## Example Usage

    The following shows getting the received license data using and ARN.

    ```python
    import pulumi
    import pulumi_aws as aws

    test = aws.licensemanager.get_received_license(license_arn="arn:aws:license-manager::111111111111:license:l-ecbaa94eb71a4830b6d7e49268fecaa0")
    ```


    :param str license_arn: The ARN of the received license you want data for.
    """
    ...
