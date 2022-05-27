# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['HsmClientCertificateArgs', 'HsmClientCertificate']

@pulumi.input_type
class HsmClientCertificateArgs:
    def __init__(__self__, *,
                 hsm_client_certificate_identifier: pulumi.Input[str],
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a HsmClientCertificate resource.
        :param pulumi.Input[str] hsm_client_certificate_identifier: The identifier of the HSM client certificate.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        """
        pulumi.set(__self__, "hsm_client_certificate_identifier", hsm_client_certificate_identifier)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter(name="hsmClientCertificateIdentifier")
    def hsm_client_certificate_identifier(self) -> pulumi.Input[str]:
        """
        The identifier of the HSM client certificate.
        """
        return pulumi.get(self, "hsm_client_certificate_identifier")

    @hsm_client_certificate_identifier.setter
    def hsm_client_certificate_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "hsm_client_certificate_identifier", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


@pulumi.input_type
class _HsmClientCertificateState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 hsm_client_certificate_identifier: Optional[pulumi.Input[str]] = None,
                 hsm_client_certificate_public_key: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering HsmClientCertificate resources.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the Hsm Client Certificate.
        :param pulumi.Input[str] hsm_client_certificate_identifier: The identifier of the HSM client certificate.
        :param pulumi.Input[str] hsm_client_certificate_public_key: The public key that the Amazon Redshift cluster will use to connect to the HSM. You must register the public key in the HSM.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if hsm_client_certificate_identifier is not None:
            pulumi.set(__self__, "hsm_client_certificate_identifier", hsm_client_certificate_identifier)
        if hsm_client_certificate_public_key is not None:
            pulumi.set(__self__, "hsm_client_certificate_public_key", hsm_client_certificate_public_key)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the Hsm Client Certificate.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="hsmClientCertificateIdentifier")
    def hsm_client_certificate_identifier(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier of the HSM client certificate.
        """
        return pulumi.get(self, "hsm_client_certificate_identifier")

    @hsm_client_certificate_identifier.setter
    def hsm_client_certificate_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hsm_client_certificate_identifier", value)

    @property
    @pulumi.getter(name="hsmClientCertificatePublicKey")
    def hsm_client_certificate_public_key(self) -> Optional[pulumi.Input[str]]:
        """
        The public key that the Amazon Redshift cluster will use to connect to the HSM. You must register the public key in the HSM.
        """
        return pulumi.get(self, "hsm_client_certificate_public_key")

    @hsm_client_certificate_public_key.setter
    def hsm_client_certificate_public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hsm_client_certificate_public_key", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class HsmClientCertificate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 hsm_client_certificate_identifier: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Creates an HSM client certificate that an Amazon Redshift cluster will use to connect to the client's HSM in order to store and retrieve the keys used to encrypt the cluster databases.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.redshift.HsmClientCertificate("example", hsm_client_certificate_identifier="example")
        ```

        ## Import

        Redshift Hsm Client Certificates support import by `hsm_client_certificate_identifier`, e.g., console

        ```sh
         $ pulumi import aws:redshift/hsmClientCertificate:HsmClientCertificate test example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hsm_client_certificate_identifier: The identifier of the HSM client certificate.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HsmClientCertificateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an HSM client certificate that an Amazon Redshift cluster will use to connect to the client's HSM in order to store and retrieve the keys used to encrypt the cluster databases.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.redshift.HsmClientCertificate("example", hsm_client_certificate_identifier="example")
        ```

        ## Import

        Redshift Hsm Client Certificates support import by `hsm_client_certificate_identifier`, e.g., console

        ```sh
         $ pulumi import aws:redshift/hsmClientCertificate:HsmClientCertificate test example
        ```

        :param str resource_name: The name of the resource.
        :param HsmClientCertificateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HsmClientCertificateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 hsm_client_certificate_identifier: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HsmClientCertificateArgs.__new__(HsmClientCertificateArgs)

            if hsm_client_certificate_identifier is None and not opts.urn:
                raise TypeError("Missing required property 'hsm_client_certificate_identifier'")
            __props__.__dict__["hsm_client_certificate_identifier"] = hsm_client_certificate_identifier
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tags_all"] = tags_all
            __props__.__dict__["arn"] = None
            __props__.__dict__["hsm_client_certificate_public_key"] = None
        super(HsmClientCertificate, __self__).__init__(
            'aws:redshift/hsmClientCertificate:HsmClientCertificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            hsm_client_certificate_identifier: Optional[pulumi.Input[str]] = None,
            hsm_client_certificate_public_key: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'HsmClientCertificate':
        """
        Get an existing HsmClientCertificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the Hsm Client Certificate.
        :param pulumi.Input[str] hsm_client_certificate_identifier: The identifier of the HSM client certificate.
        :param pulumi.Input[str] hsm_client_certificate_public_key: The public key that the Amazon Redshift cluster will use to connect to the HSM. You must register the public key in the HSM.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _HsmClientCertificateState.__new__(_HsmClientCertificateState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["hsm_client_certificate_identifier"] = hsm_client_certificate_identifier
        __props__.__dict__["hsm_client_certificate_public_key"] = hsm_client_certificate_public_key
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return HsmClientCertificate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the Hsm Client Certificate.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="hsmClientCertificateIdentifier")
    def hsm_client_certificate_identifier(self) -> pulumi.Output[str]:
        """
        The identifier of the HSM client certificate.
        """
        return pulumi.get(self, "hsm_client_certificate_identifier")

    @property
    @pulumi.getter(name="hsmClientCertificatePublicKey")
    def hsm_client_certificate_public_key(self) -> pulumi.Output[str]:
        """
        The public key that the Amazon Redshift cluster will use to connect to the HSM. You must register the public key in the HSM.
        """
        return pulumi.get(self, "hsm_client_certificate_public_key")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        """
        return pulumi.get(self, "tags_all")

