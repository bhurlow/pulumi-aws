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
    'GetServerlessSecurityConfigResult',
    'AwaitableGetServerlessSecurityConfigResult',
    'get_serverless_security_config',
    'get_serverless_security_config_output',
]

@pulumi.output_type
class GetServerlessSecurityConfigResult:
    """
    A collection of values returned by getServerlessSecurityConfig.
    """
    def __init__(__self__, config_version=None, created_date=None, description=None, id=None, last_modified_date=None, saml_options=None, type=None):
        if config_version and not isinstance(config_version, str):
            raise TypeError("Expected argument 'config_version' to be a str")
        pulumi.set(__self__, "config_version", config_version)
        if created_date and not isinstance(created_date, str):
            raise TypeError("Expected argument 'created_date' to be a str")
        pulumi.set(__self__, "created_date", created_date)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_modified_date and not isinstance(last_modified_date, str):
            raise TypeError("Expected argument 'last_modified_date' to be a str")
        pulumi.set(__self__, "last_modified_date", last_modified_date)
        if saml_options and not isinstance(saml_options, dict):
            raise TypeError("Expected argument 'saml_options' to be a dict")
        pulumi.set(__self__, "saml_options", saml_options)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="configVersion")
    def config_version(self) -> str:
        """
        The version of the security configuration.
        """
        return pulumi.get(self, "config_version")

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> str:
        """
        The date the configuration was created.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the security configuration.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastModifiedDate")
    def last_modified_date(self) -> str:
        """
        The date the configuration was last modified.
        """
        return pulumi.get(self, "last_modified_date")

    @property
    @pulumi.getter(name="samlOptions")
    def saml_options(self) -> Optional['outputs.GetServerlessSecurityConfigSamlOptionsResult']:
        """
        SAML options for the security configuration.
        """
        return pulumi.get(self, "saml_options")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of security configuration.
        """
        return pulumi.get(self, "type")


class AwaitableGetServerlessSecurityConfigResult(GetServerlessSecurityConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServerlessSecurityConfigResult(
            config_version=self.config_version,
            created_date=self.created_date,
            description=self.description,
            id=self.id,
            last_modified_date=self.last_modified_date,
            saml_options=self.saml_options,
            type=self.type)


def get_serverless_security_config(id: Optional[str] = None,
                                   saml_options: Optional[pulumi.InputType['GetServerlessSecurityConfigSamlOptionsArgs']] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServerlessSecurityConfigResult:
    """
    Data source for managing an AWS OpenSearch Serverless Security Config.

    ## Example Usage
    ### Basic Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.opensearch.get_serverless_security_config(id="saml/12345678912/example")
    ```


    :param str id: The unique identifier of the security configuration.
    :param pulumi.InputType['GetServerlessSecurityConfigSamlOptionsArgs'] saml_options: SAML options for the security configuration.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['samlOptions'] = saml_options
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws:opensearch/getServerlessSecurityConfig:getServerlessSecurityConfig', __args__, opts=opts, typ=GetServerlessSecurityConfigResult).value

    return AwaitableGetServerlessSecurityConfigResult(
        config_version=pulumi.get(__ret__, 'config_version'),
        created_date=pulumi.get(__ret__, 'created_date'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        last_modified_date=pulumi.get(__ret__, 'last_modified_date'),
        saml_options=pulumi.get(__ret__, 'saml_options'),
        type=pulumi.get(__ret__, 'type'))


@_utilities.lift_output_func(get_serverless_security_config)
def get_serverless_security_config_output(id: Optional[pulumi.Input[str]] = None,
                                          saml_options: Optional[pulumi.Input[Optional[pulumi.InputType['GetServerlessSecurityConfigSamlOptionsArgs']]]] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServerlessSecurityConfigResult]:
    """
    Data source for managing an AWS OpenSearch Serverless Security Config.

    ## Example Usage
    ### Basic Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.opensearch.get_serverless_security_config(id="saml/12345678912/example")
    ```


    :param str id: The unique identifier of the security configuration.
    :param pulumi.InputType['GetServerlessSecurityConfigSamlOptionsArgs'] saml_options: SAML options for the security configuration.
    """
    ...
