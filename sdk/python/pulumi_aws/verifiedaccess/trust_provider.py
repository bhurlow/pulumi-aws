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

__all__ = ['TrustProviderArgs', 'TrustProvider']

@pulumi.input_type
class TrustProviderArgs:
    def __init__(__self__, *,
                 policy_reference_name: pulumi.Input[str],
                 trust_provider_type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 device_options: Optional[pulumi.Input['TrustProviderDeviceOptionsArgs']] = None,
                 device_trust_provider_type: Optional[pulumi.Input[str]] = None,
                 oidc_options: Optional[pulumi.Input['TrustProviderOidcOptionsArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 user_trust_provider_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TrustProvider resource.
        :param pulumi.Input[str] policy_reference_name: The identifier to be used when working with policy rules.
        :param pulumi.Input[str] trust_provider_type: The type of trust provider can be either user or device-based.
               
               The following arguments are optional:
        :param pulumi.Input[str] description: A description for the AWS Verified Access trust provider.
        :param pulumi.Input['TrustProviderDeviceOptionsArgs'] device_options: A block of options for device identity based trust providers.
        :param pulumi.Input[str] device_trust_provider_type: The type of device-based trust provider.
        :param pulumi.Input['TrustProviderOidcOptionsArgs'] oidc_options: The OpenID Connect details for an oidc-type, user-identity based trust provider.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] user_trust_provider_type: The type of user-based trust provider.
        """
        pulumi.set(__self__, "policy_reference_name", policy_reference_name)
        pulumi.set(__self__, "trust_provider_type", trust_provider_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if device_options is not None:
            pulumi.set(__self__, "device_options", device_options)
        if device_trust_provider_type is not None:
            pulumi.set(__self__, "device_trust_provider_type", device_trust_provider_type)
        if oidc_options is not None:
            pulumi.set(__self__, "oidc_options", oidc_options)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if user_trust_provider_type is not None:
            pulumi.set(__self__, "user_trust_provider_type", user_trust_provider_type)

    @property
    @pulumi.getter(name="policyReferenceName")
    def policy_reference_name(self) -> pulumi.Input[str]:
        """
        The identifier to be used when working with policy rules.
        """
        return pulumi.get(self, "policy_reference_name")

    @policy_reference_name.setter
    def policy_reference_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_reference_name", value)

    @property
    @pulumi.getter(name="trustProviderType")
    def trust_provider_type(self) -> pulumi.Input[str]:
        """
        The type of trust provider can be either user or device-based.

        The following arguments are optional:
        """
        return pulumi.get(self, "trust_provider_type")

    @trust_provider_type.setter
    def trust_provider_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "trust_provider_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the AWS Verified Access trust provider.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="deviceOptions")
    def device_options(self) -> Optional[pulumi.Input['TrustProviderDeviceOptionsArgs']]:
        """
        A block of options for device identity based trust providers.
        """
        return pulumi.get(self, "device_options")

    @device_options.setter
    def device_options(self, value: Optional[pulumi.Input['TrustProviderDeviceOptionsArgs']]):
        pulumi.set(self, "device_options", value)

    @property
    @pulumi.getter(name="deviceTrustProviderType")
    def device_trust_provider_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of device-based trust provider.
        """
        return pulumi.get(self, "device_trust_provider_type")

    @device_trust_provider_type.setter
    def device_trust_provider_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_trust_provider_type", value)

    @property
    @pulumi.getter(name="oidcOptions")
    def oidc_options(self) -> Optional[pulumi.Input['TrustProviderOidcOptionsArgs']]:
        """
        The OpenID Connect details for an oidc-type, user-identity based trust provider.
        """
        return pulumi.get(self, "oidc_options")

    @oidc_options.setter
    def oidc_options(self, value: Optional[pulumi.Input['TrustProviderOidcOptionsArgs']]):
        pulumi.set(self, "oidc_options", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="userTrustProviderType")
    def user_trust_provider_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of user-based trust provider.
        """
        return pulumi.get(self, "user_trust_provider_type")

    @user_trust_provider_type.setter
    def user_trust_provider_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_trust_provider_type", value)


@pulumi.input_type
class _TrustProviderState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 device_options: Optional[pulumi.Input['TrustProviderDeviceOptionsArgs']] = None,
                 device_trust_provider_type: Optional[pulumi.Input[str]] = None,
                 oidc_options: Optional[pulumi.Input['TrustProviderOidcOptionsArgs']] = None,
                 policy_reference_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 trust_provider_type: Optional[pulumi.Input[str]] = None,
                 user_trust_provider_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TrustProvider resources.
        :param pulumi.Input[str] description: A description for the AWS Verified Access trust provider.
        :param pulumi.Input['TrustProviderDeviceOptionsArgs'] device_options: A block of options for device identity based trust providers.
        :param pulumi.Input[str] device_trust_provider_type: The type of device-based trust provider.
        :param pulumi.Input['TrustProviderOidcOptionsArgs'] oidc_options: The OpenID Connect details for an oidc-type, user-identity based trust provider.
        :param pulumi.Input[str] policy_reference_name: The identifier to be used when working with policy rules.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] trust_provider_type: The type of trust provider can be either user or device-based.
               
               The following arguments are optional:
        :param pulumi.Input[str] user_trust_provider_type: The type of user-based trust provider.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if device_options is not None:
            pulumi.set(__self__, "device_options", device_options)
        if device_trust_provider_type is not None:
            pulumi.set(__self__, "device_trust_provider_type", device_trust_provider_type)
        if oidc_options is not None:
            pulumi.set(__self__, "oidc_options", oidc_options)
        if policy_reference_name is not None:
            pulumi.set(__self__, "policy_reference_name", policy_reference_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
            pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if trust_provider_type is not None:
            pulumi.set(__self__, "trust_provider_type", trust_provider_type)
        if user_trust_provider_type is not None:
            pulumi.set(__self__, "user_trust_provider_type", user_trust_provider_type)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the AWS Verified Access trust provider.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="deviceOptions")
    def device_options(self) -> Optional[pulumi.Input['TrustProviderDeviceOptionsArgs']]:
        """
        A block of options for device identity based trust providers.
        """
        return pulumi.get(self, "device_options")

    @device_options.setter
    def device_options(self, value: Optional[pulumi.Input['TrustProviderDeviceOptionsArgs']]):
        pulumi.set(self, "device_options", value)

    @property
    @pulumi.getter(name="deviceTrustProviderType")
    def device_trust_provider_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of device-based trust provider.
        """
        return pulumi.get(self, "device_trust_provider_type")

    @device_trust_provider_type.setter
    def device_trust_provider_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_trust_provider_type", value)

    @property
    @pulumi.getter(name="oidcOptions")
    def oidc_options(self) -> Optional[pulumi.Input['TrustProviderOidcOptionsArgs']]:
        """
        The OpenID Connect details for an oidc-type, user-identity based trust provider.
        """
        return pulumi.get(self, "oidc_options")

    @oidc_options.setter
    def oidc_options(self, value: Optional[pulumi.Input['TrustProviderOidcOptionsArgs']]):
        pulumi.set(self, "oidc_options", value)

    @property
    @pulumi.getter(name="policyReferenceName")
    def policy_reference_name(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier to be used when working with policy rules.
        """
        return pulumi.get(self, "policy_reference_name")

    @policy_reference_name.setter
    def policy_reference_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_reference_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter(name="trustProviderType")
    def trust_provider_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of trust provider can be either user or device-based.

        The following arguments are optional:
        """
        return pulumi.get(self, "trust_provider_type")

    @trust_provider_type.setter
    def trust_provider_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trust_provider_type", value)

    @property
    @pulumi.getter(name="userTrustProviderType")
    def user_trust_provider_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of user-based trust provider.
        """
        return pulumi.get(self, "user_trust_provider_type")

    @user_trust_provider_type.setter
    def user_trust_provider_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_trust_provider_type", value)


class TrustProvider(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device_options: Optional[pulumi.Input[pulumi.InputType['TrustProviderDeviceOptionsArgs']]] = None,
                 device_trust_provider_type: Optional[pulumi.Input[str]] = None,
                 oidc_options: Optional[pulumi.Input[pulumi.InputType['TrustProviderOidcOptionsArgs']]] = None,
                 policy_reference_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 trust_provider_type: Optional[pulumi.Input[str]] = None,
                 user_trust_provider_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource for managing a Verified Access Trust Provider.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.verifiedaccess.TrustProvider("example",
            policy_reference_name="example",
            trust_provider_type="user",
            user_trust_provider_type="iam-identity-center")
        ```

        ## Import

        In TODO v1.5.0 and later, use an `import` block to import Transfer Workflows using the `id`. For exampleterraform import {

         to = aws_verifiedaccess_trust_provider.example

         id = "vatp-8012925589" } Using `TODO import`, import Transfer Workflows using the

        `id`. For exampleconsole % TODO import aws_verifiedaccess_trust_provider.example vatp-8012925589

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description for the AWS Verified Access trust provider.
        :param pulumi.Input[pulumi.InputType['TrustProviderDeviceOptionsArgs']] device_options: A block of options for device identity based trust providers.
        :param pulumi.Input[str] device_trust_provider_type: The type of device-based trust provider.
        :param pulumi.Input[pulumi.InputType['TrustProviderOidcOptionsArgs']] oidc_options: The OpenID Connect details for an oidc-type, user-identity based trust provider.
        :param pulumi.Input[str] policy_reference_name: The identifier to be used when working with policy rules.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] trust_provider_type: The type of trust provider can be either user or device-based.
               
               The following arguments are optional:
        :param pulumi.Input[str] user_trust_provider_type: The type of user-based trust provider.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TrustProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for managing a Verified Access Trust Provider.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.verifiedaccess.TrustProvider("example",
            policy_reference_name="example",
            trust_provider_type="user",
            user_trust_provider_type="iam-identity-center")
        ```

        ## Import

        In TODO v1.5.0 and later, use an `import` block to import Transfer Workflows using the `id`. For exampleterraform import {

         to = aws_verifiedaccess_trust_provider.example

         id = "vatp-8012925589" } Using `TODO import`, import Transfer Workflows using the

        `id`. For exampleconsole % TODO import aws_verifiedaccess_trust_provider.example vatp-8012925589

        :param str resource_name: The name of the resource.
        :param TrustProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TrustProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device_options: Optional[pulumi.Input[pulumi.InputType['TrustProviderDeviceOptionsArgs']]] = None,
                 device_trust_provider_type: Optional[pulumi.Input[str]] = None,
                 oidc_options: Optional[pulumi.Input[pulumi.InputType['TrustProviderOidcOptionsArgs']]] = None,
                 policy_reference_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 trust_provider_type: Optional[pulumi.Input[str]] = None,
                 user_trust_provider_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TrustProviderArgs.__new__(TrustProviderArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["device_options"] = device_options
            __props__.__dict__["device_trust_provider_type"] = device_trust_provider_type
            __props__.__dict__["oidc_options"] = oidc_options
            if policy_reference_name is None and not opts.urn:
                raise TypeError("Missing required property 'policy_reference_name'")
            __props__.__dict__["policy_reference_name"] = policy_reference_name
            __props__.__dict__["tags"] = tags
            if trust_provider_type is None and not opts.urn:
                raise TypeError("Missing required property 'trust_provider_type'")
            __props__.__dict__["trust_provider_type"] = trust_provider_type
            __props__.__dict__["user_trust_provider_type"] = user_trust_provider_type
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(TrustProvider, __self__).__init__(
            'aws:verifiedaccess/trustProvider:TrustProvider',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            device_options: Optional[pulumi.Input[pulumi.InputType['TrustProviderDeviceOptionsArgs']]] = None,
            device_trust_provider_type: Optional[pulumi.Input[str]] = None,
            oidc_options: Optional[pulumi.Input[pulumi.InputType['TrustProviderOidcOptionsArgs']]] = None,
            policy_reference_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            trust_provider_type: Optional[pulumi.Input[str]] = None,
            user_trust_provider_type: Optional[pulumi.Input[str]] = None) -> 'TrustProvider':
        """
        Get an existing TrustProvider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description for the AWS Verified Access trust provider.
        :param pulumi.Input[pulumi.InputType['TrustProviderDeviceOptionsArgs']] device_options: A block of options for device identity based trust providers.
        :param pulumi.Input[str] device_trust_provider_type: The type of device-based trust provider.
        :param pulumi.Input[pulumi.InputType['TrustProviderOidcOptionsArgs']] oidc_options: The OpenID Connect details for an oidc-type, user-identity based trust provider.
        :param pulumi.Input[str] policy_reference_name: The identifier to be used when working with policy rules.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] trust_provider_type: The type of trust provider can be either user or device-based.
               
               The following arguments are optional:
        :param pulumi.Input[str] user_trust_provider_type: The type of user-based trust provider.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TrustProviderState.__new__(_TrustProviderState)

        __props__.__dict__["description"] = description
        __props__.__dict__["device_options"] = device_options
        __props__.__dict__["device_trust_provider_type"] = device_trust_provider_type
        __props__.__dict__["oidc_options"] = oidc_options
        __props__.__dict__["policy_reference_name"] = policy_reference_name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["trust_provider_type"] = trust_provider_type
        __props__.__dict__["user_trust_provider_type"] = user_trust_provider_type
        return TrustProvider(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description for the AWS Verified Access trust provider.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="deviceOptions")
    def device_options(self) -> pulumi.Output[Optional['outputs.TrustProviderDeviceOptions']]:
        """
        A block of options for device identity based trust providers.
        """
        return pulumi.get(self, "device_options")

    @property
    @pulumi.getter(name="deviceTrustProviderType")
    def device_trust_provider_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of device-based trust provider.
        """
        return pulumi.get(self, "device_trust_provider_type")

    @property
    @pulumi.getter(name="oidcOptions")
    def oidc_options(self) -> pulumi.Output[Optional['outputs.TrustProviderOidcOptions']]:
        """
        The OpenID Connect details for an oidc-type, user-identity based trust provider.
        """
        return pulumi.get(self, "oidc_options")

    @property
    @pulumi.getter(name="policyReferenceName")
    def policy_reference_name(self) -> pulumi.Output[str]:
        """
        The identifier to be used when working with policy rules.
        """
        return pulumi.get(self, "policy_reference_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
        pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")

        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter(name="trustProviderType")
    def trust_provider_type(self) -> pulumi.Output[str]:
        """
        The type of trust provider can be either user or device-based.

        The following arguments are optional:
        """
        return pulumi.get(self, "trust_provider_type")

    @property
    @pulumi.getter(name="userTrustProviderType")
    def user_trust_provider_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of user-based trust provider.
        """
        return pulumi.get(self, "user_trust_provider_type")

