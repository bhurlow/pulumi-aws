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

__all__ = ['TrustAnchorArgs', 'TrustAnchor']

@pulumi.input_type
class TrustAnchorArgs:
    def __init__(__self__, *,
                 source: pulumi.Input['TrustAnchorSourceArgs'],
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a TrustAnchor resource.
        :param pulumi.Input['TrustAnchorSourceArgs'] source: The source of trust, documented below
        :param pulumi.Input[bool] enabled: Whether or not the Trust Anchor should be enabled.
        :param pulumi.Input[str] name: The name of the Trust Anchor.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "source", source)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def source(self) -> pulumi.Input['TrustAnchorSourceArgs']:
        """
        The source of trust, documented below
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: pulumi.Input['TrustAnchorSourceArgs']):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not the Trust Anchor should be enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Trust Anchor.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _TrustAnchorState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input['TrustAnchorSourceArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering TrustAnchor resources.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the Trust Anchor
        :param pulumi.Input[bool] enabled: Whether or not the Trust Anchor should be enabled.
        :param pulumi.Input[str] name: The name of the Trust Anchor.
        :param pulumi.Input['TrustAnchorSourceArgs'] source: The source of trust, documented below
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            warnings.warn("""Please use `tags` instead.""", DeprecationWarning)
            pulumi.log.warn("""tags_all is deprecated: Please use `tags` instead.""")
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the Trust Anchor
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not the Trust Anchor should be enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Trust Anchor.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input['TrustAnchorSourceArgs']]:
        """
        The source of trust, documented below
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input['TrustAnchorSourceArgs']]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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


class TrustAnchor(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[pulumi.InputType['TrustAnchorSourceArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Resource for managing a Roles Anywhere Trust Anchor.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_certificate_authority = aws.acmpca.CertificateAuthority("exampleCertificateAuthority",
            permanent_deletion_time_in_days=7,
            type="ROOT",
            certificate_authority_configuration=aws.acmpca.CertificateAuthorityCertificateAuthorityConfigurationArgs(
                key_algorithm="RSA_4096",
                signing_algorithm="SHA512WITHRSA",
                subject=aws.acmpca.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs(
                    common_name="example.com",
                ),
            ))
        current = aws.get_partition()
        test_certificate = aws.acmpca.Certificate("testCertificate",
            certificate_authority_arn=example_certificate_authority.arn,
            certificate_signing_request=example_certificate_authority.certificate_signing_request,
            signing_algorithm="SHA512WITHRSA",
            template_arn=f"arn:{current.partition}:acm-pca:::template/RootCACertificate/V1",
            validity=aws.acmpca.CertificateValidityArgs(
                type="YEARS",
                value="1",
            ))
        example_certificate_authority_certificate = aws.acmpca.CertificateAuthorityCertificate("exampleCertificateAuthorityCertificate",
            certificate_authority_arn=example_certificate_authority.arn,
            certificate=aws_acmpca_certificate["example"]["certificate"],
            certificate_chain=aws_acmpca_certificate["example"]["certificate_chain"])
        test_trust_anchor = aws.rolesanywhere.TrustAnchor("testTrustAnchor", source=aws.rolesanywhere.TrustAnchorSourceArgs(
            source_data=aws.rolesanywhere.TrustAnchorSourceSourceDataArgs(
                acm_pca_arn=example_certificate_authority.arn,
            ),
            source_type="AWS_ACM_PCA",
        ),
        opts=pulumi.ResourceOptions(depends_on=[example_certificate_authority_certificate]))
        ```

        ## Import

        Using `pulumi import`, import `aws_rolesanywhere_trust_anchor` using its `id`. For example:

        ```sh
         $ pulumi import aws:rolesanywhere/trustAnchor:TrustAnchor example 92b2fbbb-984d-41a3-a765-e3cbdb69ebb1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Whether or not the Trust Anchor should be enabled.
        :param pulumi.Input[str] name: The name of the Trust Anchor.
        :param pulumi.Input[pulumi.InputType['TrustAnchorSourceArgs']] source: The source of trust, documented below
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TrustAnchorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for managing a Roles Anywhere Trust Anchor.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_certificate_authority = aws.acmpca.CertificateAuthority("exampleCertificateAuthority",
            permanent_deletion_time_in_days=7,
            type="ROOT",
            certificate_authority_configuration=aws.acmpca.CertificateAuthorityCertificateAuthorityConfigurationArgs(
                key_algorithm="RSA_4096",
                signing_algorithm="SHA512WITHRSA",
                subject=aws.acmpca.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs(
                    common_name="example.com",
                ),
            ))
        current = aws.get_partition()
        test_certificate = aws.acmpca.Certificate("testCertificate",
            certificate_authority_arn=example_certificate_authority.arn,
            certificate_signing_request=example_certificate_authority.certificate_signing_request,
            signing_algorithm="SHA512WITHRSA",
            template_arn=f"arn:{current.partition}:acm-pca:::template/RootCACertificate/V1",
            validity=aws.acmpca.CertificateValidityArgs(
                type="YEARS",
                value="1",
            ))
        example_certificate_authority_certificate = aws.acmpca.CertificateAuthorityCertificate("exampleCertificateAuthorityCertificate",
            certificate_authority_arn=example_certificate_authority.arn,
            certificate=aws_acmpca_certificate["example"]["certificate"],
            certificate_chain=aws_acmpca_certificate["example"]["certificate_chain"])
        test_trust_anchor = aws.rolesanywhere.TrustAnchor("testTrustAnchor", source=aws.rolesanywhere.TrustAnchorSourceArgs(
            source_data=aws.rolesanywhere.TrustAnchorSourceSourceDataArgs(
                acm_pca_arn=example_certificate_authority.arn,
            ),
            source_type="AWS_ACM_PCA",
        ),
        opts=pulumi.ResourceOptions(depends_on=[example_certificate_authority_certificate]))
        ```

        ## Import

        Using `pulumi import`, import `aws_rolesanywhere_trust_anchor` using its `id`. For example:

        ```sh
         $ pulumi import aws:rolesanywhere/trustAnchor:TrustAnchor example 92b2fbbb-984d-41a3-a765-e3cbdb69ebb1
        ```

        :param str resource_name: The name of the resource.
        :param TrustAnchorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TrustAnchorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[pulumi.InputType['TrustAnchorSourceArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TrustAnchorArgs.__new__(TrustAnchorArgs)

            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["name"] = name
            if source is None and not opts.urn:
                raise TypeError("Missing required property 'source'")
            __props__.__dict__["source"] = source
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["tagsAll"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(TrustAnchor, __self__).__init__(
            'aws:rolesanywhere/trustAnchor:TrustAnchor',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            source: Optional[pulumi.Input[pulumi.InputType['TrustAnchorSourceArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'TrustAnchor':
        """
        Get an existing TrustAnchor resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the Trust Anchor
        :param pulumi.Input[bool] enabled: Whether or not the Trust Anchor should be enabled.
        :param pulumi.Input[str] name: The name of the Trust Anchor.
        :param pulumi.Input[pulumi.InputType['TrustAnchorSourceArgs']] source: The source of trust, documented below
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TrustAnchorState.__new__(_TrustAnchorState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["name"] = name
        __props__.__dict__["source"] = source
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return TrustAnchor(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the Trust Anchor
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        Whether or not the Trust Anchor should be enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Trust Anchor.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output['outputs.TrustAnchorSource']:
        """
        The source of trust, documented below
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

