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

__all__ = ['ConfigurationProfileArgs', 'ConfigurationProfile']

@pulumi.input_type
class ConfigurationProfileArgs:
    def __init__(__self__, *,
                 application_id: pulumi.Input[str],
                 location_uri: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retrieval_role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 validators: Optional[pulumi.Input[Sequence[pulumi.Input['ConfigurationProfileValidatorArgs']]]] = None):
        """
        The set of arguments for constructing a ConfigurationProfile resource.
        :param pulumi.Input[str] application_id: Application ID. Must be between 4 and 7 characters in length.
        :param pulumi.Input[str] location_uri: URI to locate the configuration. You can specify the AWS AppConfig hosted configuration store, Systems Manager (SSM) document, an SSM Parameter Store parameter, or an Amazon S3 object. For the hosted configuration store, specify `hosted`. For an SSM document, specify either the document name in the format `ssm-document://<Document_name>` or the ARN. For a parameter, specify either the parameter name in the format `ssm-parameter://<Parameter_name>` or the ARN. For an Amazon S3 object, specify the URI in the following format: `s3://<bucket>/<objectKey>`.
        :param pulumi.Input[str] description: Description of the configuration profile. Can be at most 1024 characters.
        :param pulumi.Input[str] name: Name for the configuration profile. Must be between 1 and 64 characters in length.
        :param pulumi.Input[str] retrieval_role_arn: ARN of an IAM role with permission to access the configuration at the specified `location_uri`. A retrieval role ARN is not required for configurations stored in the AWS AppConfig `hosted` configuration store. It is required for all other sources that store your configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] type: Type of configurations contained in the profile. Valid values: `AWS.AppConfig.FeatureFlags` and `AWS.Freeform`.  Default: `AWS.Freeform`.
        :param pulumi.Input[Sequence[pulumi.Input['ConfigurationProfileValidatorArgs']]] validators: Set of methods for validating the configuration. Maximum of 2. See Validator below for more details.
        """
        pulumi.set(__self__, "application_id", application_id)
        pulumi.set(__self__, "location_uri", location_uri)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if retrieval_role_arn is not None:
            pulumi.set(__self__, "retrieval_role_arn", retrieval_role_arn)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if validators is not None:
            pulumi.set(__self__, "validators", validators)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Input[str]:
        """
        Application ID. Must be between 4 and 7 characters in length.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="locationUri")
    def location_uri(self) -> pulumi.Input[str]:
        """
        URI to locate the configuration. You can specify the AWS AppConfig hosted configuration store, Systems Manager (SSM) document, an SSM Parameter Store parameter, or an Amazon S3 object. For the hosted configuration store, specify `hosted`. For an SSM document, specify either the document name in the format `ssm-document://<Document_name>` or the ARN. For a parameter, specify either the parameter name in the format `ssm-parameter://<Parameter_name>` or the ARN. For an Amazon S3 object, specify the URI in the following format: `s3://<bucket>/<objectKey>`.
        """
        return pulumi.get(self, "location_uri")

    @location_uri.setter
    def location_uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "location_uri", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the configuration profile. Can be at most 1024 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name for the configuration profile. Must be between 1 and 64 characters in length.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="retrievalRoleArn")
    def retrieval_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of an IAM role with permission to access the configuration at the specified `location_uri`. A retrieval role ARN is not required for configurations stored in the AWS AppConfig `hosted` configuration store. It is required for all other sources that store your configuration.
        """
        return pulumi.get(self, "retrieval_role_arn")

    @retrieval_role_arn.setter
    def retrieval_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "retrieval_role_arn", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of configurations contained in the profile. Valid values: `AWS.AppConfig.FeatureFlags` and `AWS.Freeform`.  Default: `AWS.Freeform`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def validators(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ConfigurationProfileValidatorArgs']]]]:
        """
        Set of methods for validating the configuration. Maximum of 2. See Validator below for more details.
        """
        return pulumi.get(self, "validators")

    @validators.setter
    def validators(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ConfigurationProfileValidatorArgs']]]]):
        pulumi.set(self, "validators", value)


@pulumi.input_type
class _ConfigurationProfileState:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 configuration_profile_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 location_uri: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retrieval_role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 validators: Optional[pulumi.Input[Sequence[pulumi.Input['ConfigurationProfileValidatorArgs']]]] = None):
        """
        Input properties used for looking up and filtering ConfigurationProfile resources.
        :param pulumi.Input[str] application_id: Application ID. Must be between 4 and 7 characters in length.
        :param pulumi.Input[str] arn: ARN of the AppConfig Configuration Profile.
        :param pulumi.Input[str] configuration_profile_id: The configuration profile ID.
        :param pulumi.Input[str] description: Description of the configuration profile. Can be at most 1024 characters.
        :param pulumi.Input[str] location_uri: URI to locate the configuration. You can specify the AWS AppConfig hosted configuration store, Systems Manager (SSM) document, an SSM Parameter Store parameter, or an Amazon S3 object. For the hosted configuration store, specify `hosted`. For an SSM document, specify either the document name in the format `ssm-document://<Document_name>` or the ARN. For a parameter, specify either the parameter name in the format `ssm-parameter://<Parameter_name>` or the ARN. For an Amazon S3 object, specify the URI in the following format: `s3://<bucket>/<objectKey>`.
        :param pulumi.Input[str] name: Name for the configuration profile. Must be between 1 and 64 characters in length.
        :param pulumi.Input[str] retrieval_role_arn: ARN of an IAM role with permission to access the configuration at the specified `location_uri`. A retrieval role ARN is not required for configurations stored in the AWS AppConfig `hosted` configuration store. It is required for all other sources that store your configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[str] type: Type of configurations contained in the profile. Valid values: `AWS.AppConfig.FeatureFlags` and `AWS.Freeform`.  Default: `AWS.Freeform`.
        :param pulumi.Input[Sequence[pulumi.Input['ConfigurationProfileValidatorArgs']]] validators: Set of methods for validating the configuration. Maximum of 2. See Validator below for more details.
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if configuration_profile_id is not None:
            pulumi.set(__self__, "configuration_profile_id", configuration_profile_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if location_uri is not None:
            pulumi.set(__self__, "location_uri", location_uri)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if retrieval_role_arn is not None:
            pulumi.set(__self__, "retrieval_role_arn", retrieval_role_arn)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if validators is not None:
            pulumi.set(__self__, "validators", validators)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
        Application ID. Must be between 4 and 7 characters in length.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the AppConfig Configuration Profile.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="configurationProfileId")
    def configuration_profile_id(self) -> Optional[pulumi.Input[str]]:
        """
        The configuration profile ID.
        """
        return pulumi.get(self, "configuration_profile_id")

    @configuration_profile_id.setter
    def configuration_profile_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "configuration_profile_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the configuration profile. Can be at most 1024 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="locationUri")
    def location_uri(self) -> Optional[pulumi.Input[str]]:
        """
        URI to locate the configuration. You can specify the AWS AppConfig hosted configuration store, Systems Manager (SSM) document, an SSM Parameter Store parameter, or an Amazon S3 object. For the hosted configuration store, specify `hosted`. For an SSM document, specify either the document name in the format `ssm-document://<Document_name>` or the ARN. For a parameter, specify either the parameter name in the format `ssm-parameter://<Parameter_name>` or the ARN. For an Amazon S3 object, specify the URI in the following format: `s3://<bucket>/<objectKey>`.
        """
        return pulumi.get(self, "location_uri")

    @location_uri.setter
    def location_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location_uri", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name for the configuration profile. Must be between 1 and 64 characters in length.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="retrievalRoleArn")
    def retrieval_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of an IAM role with permission to access the configuration at the specified `location_uri`. A retrieval role ARN is not required for configurations stored in the AWS AppConfig `hosted` configuration store. It is required for all other sources that store your configuration.
        """
        return pulumi.get(self, "retrieval_role_arn")

    @retrieval_role_arn.setter
    def retrieval_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "retrieval_role_arn", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of configurations contained in the profile. Valid values: `AWS.AppConfig.FeatureFlags` and `AWS.Freeform`.  Default: `AWS.Freeform`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def validators(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ConfigurationProfileValidatorArgs']]]]:
        """
        Set of methods for validating the configuration. Maximum of 2. See Validator below for more details.
        """
        return pulumi.get(self, "validators")

    @validators.setter
    def validators(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ConfigurationProfileValidatorArgs']]]]):
        pulumi.set(self, "validators", value)


class ConfigurationProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 location_uri: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retrieval_role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 validators: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConfigurationProfileValidatorArgs']]]]] = None,
                 __props__=None):
        """
        Provides an AppConfig Configuration Profile resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.appconfig.ConfigurationProfile("example",
            application_id=aws_appconfig_application["example"]["id"],
            description="Example Configuration Profile",
            location_uri="hosted",
            validators=[aws.appconfig.ConfigurationProfileValidatorArgs(
                content=aws_lambda_function["example"]["arn"],
                type="LAMBDA",
            )],
            tags={
                "Type": "AppConfig Configuration Profile",
            })
        ```

        ## Import

        AppConfig Configuration Profiles can be imported by using the configuration profile ID and application ID separated by a colon (`:`), e.g.,

        ```sh
         $ pulumi import aws:appconfig/configurationProfile:ConfigurationProfile example 71abcde:11xxxxx
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: Application ID. Must be between 4 and 7 characters in length.
        :param pulumi.Input[str] description: Description of the configuration profile. Can be at most 1024 characters.
        :param pulumi.Input[str] location_uri: URI to locate the configuration. You can specify the AWS AppConfig hosted configuration store, Systems Manager (SSM) document, an SSM Parameter Store parameter, or an Amazon S3 object. For the hosted configuration store, specify `hosted`. For an SSM document, specify either the document name in the format `ssm-document://<Document_name>` or the ARN. For a parameter, specify either the parameter name in the format `ssm-parameter://<Parameter_name>` or the ARN. For an Amazon S3 object, specify the URI in the following format: `s3://<bucket>/<objectKey>`.
        :param pulumi.Input[str] name: Name for the configuration profile. Must be between 1 and 64 characters in length.
        :param pulumi.Input[str] retrieval_role_arn: ARN of an IAM role with permission to access the configuration at the specified `location_uri`. A retrieval role ARN is not required for configurations stored in the AWS AppConfig `hosted` configuration store. It is required for all other sources that store your configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[str] type: Type of configurations contained in the profile. Valid values: `AWS.AppConfig.FeatureFlags` and `AWS.Freeform`.  Default: `AWS.Freeform`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConfigurationProfileValidatorArgs']]]] validators: Set of methods for validating the configuration. Maximum of 2. See Validator below for more details.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConfigurationProfileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an AppConfig Configuration Profile resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.appconfig.ConfigurationProfile("example",
            application_id=aws_appconfig_application["example"]["id"],
            description="Example Configuration Profile",
            location_uri="hosted",
            validators=[aws.appconfig.ConfigurationProfileValidatorArgs(
                content=aws_lambda_function["example"]["arn"],
                type="LAMBDA",
            )],
            tags={
                "Type": "AppConfig Configuration Profile",
            })
        ```

        ## Import

        AppConfig Configuration Profiles can be imported by using the configuration profile ID and application ID separated by a colon (`:`), e.g.,

        ```sh
         $ pulumi import aws:appconfig/configurationProfile:ConfigurationProfile example 71abcde:11xxxxx
        ```

        :param str resource_name: The name of the resource.
        :param ConfigurationProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConfigurationProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 location_uri: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retrieval_role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 validators: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConfigurationProfileValidatorArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConfigurationProfileArgs.__new__(ConfigurationProfileArgs)

            if application_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_id'")
            __props__.__dict__["application_id"] = application_id
            __props__.__dict__["description"] = description
            if location_uri is None and not opts.urn:
                raise TypeError("Missing required property 'location_uri'")
            __props__.__dict__["location_uri"] = location_uri
            __props__.__dict__["name"] = name
            __props__.__dict__["retrieval_role_arn"] = retrieval_role_arn
            __props__.__dict__["tags"] = tags
            __props__.__dict__["type"] = type
            __props__.__dict__["validators"] = validators
            __props__.__dict__["arn"] = None
            __props__.__dict__["configuration_profile_id"] = None
            __props__.__dict__["tags_all"] = None
        super(ConfigurationProfile, __self__).__init__(
            'aws:appconfig/configurationProfile:ConfigurationProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            configuration_profile_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            location_uri: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            retrieval_role_arn: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            type: Optional[pulumi.Input[str]] = None,
            validators: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConfigurationProfileValidatorArgs']]]]] = None) -> 'ConfigurationProfile':
        """
        Get an existing ConfigurationProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: Application ID. Must be between 4 and 7 characters in length.
        :param pulumi.Input[str] arn: ARN of the AppConfig Configuration Profile.
        :param pulumi.Input[str] configuration_profile_id: The configuration profile ID.
        :param pulumi.Input[str] description: Description of the configuration profile. Can be at most 1024 characters.
        :param pulumi.Input[str] location_uri: URI to locate the configuration. You can specify the AWS AppConfig hosted configuration store, Systems Manager (SSM) document, an SSM Parameter Store parameter, or an Amazon S3 object. For the hosted configuration store, specify `hosted`. For an SSM document, specify either the document name in the format `ssm-document://<Document_name>` or the ARN. For a parameter, specify either the parameter name in the format `ssm-parameter://<Parameter_name>` or the ARN. For an Amazon S3 object, specify the URI in the following format: `s3://<bucket>/<objectKey>`.
        :param pulumi.Input[str] name: Name for the configuration profile. Must be between 1 and 64 characters in length.
        :param pulumi.Input[str] retrieval_role_arn: ARN of an IAM role with permission to access the configuration at the specified `location_uri`. A retrieval role ARN is not required for configurations stored in the AWS AppConfig `hosted` configuration store. It is required for all other sources that store your configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        :param pulumi.Input[str] type: Type of configurations contained in the profile. Valid values: `AWS.AppConfig.FeatureFlags` and `AWS.Freeform`.  Default: `AWS.Freeform`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConfigurationProfileValidatorArgs']]]] validators: Set of methods for validating the configuration. Maximum of 2. See Validator below for more details.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConfigurationProfileState.__new__(_ConfigurationProfileState)

        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["arn"] = arn
        __props__.__dict__["configuration_profile_id"] = configuration_profile_id
        __props__.__dict__["description"] = description
        __props__.__dict__["location_uri"] = location_uri
        __props__.__dict__["name"] = name
        __props__.__dict__["retrieval_role_arn"] = retrieval_role_arn
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["type"] = type
        __props__.__dict__["validators"] = validators
        return ConfigurationProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        Application ID. Must be between 4 and 7 characters in length.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN of the AppConfig Configuration Profile.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="configurationProfileId")
    def configuration_profile_id(self) -> pulumi.Output[str]:
        """
        The configuration profile ID.
        """
        return pulumi.get(self, "configuration_profile_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the configuration profile. Can be at most 1024 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="locationUri")
    def location_uri(self) -> pulumi.Output[str]:
        """
        URI to locate the configuration. You can specify the AWS AppConfig hosted configuration store, Systems Manager (SSM) document, an SSM Parameter Store parameter, or an Amazon S3 object. For the hosted configuration store, specify `hosted`. For an SSM document, specify either the document name in the format `ssm-document://<Document_name>` or the ARN. For a parameter, specify either the parameter name in the format `ssm-parameter://<Parameter_name>` or the ARN. For an Amazon S3 object, specify the URI in the following format: `s3://<bucket>/<objectKey>`.
        """
        return pulumi.get(self, "location_uri")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name for the configuration profile. Must be between 1 and 64 characters in length.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="retrievalRoleArn")
    def retrieval_role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        ARN of an IAM role with permission to access the configuration at the specified `location_uri`. A retrieval role ARN is not required for configurations stored in the AWS AppConfig `hosted` configuration store. It is required for all other sources that store your configuration.
        """
        return pulumi.get(self, "retrieval_role_arn")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        Type of configurations contained in the profile. Valid values: `AWS.AppConfig.FeatureFlags` and `AWS.Freeform`.  Default: `AWS.Freeform`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def validators(self) -> pulumi.Output[Optional[Sequence['outputs.ConfigurationProfileValidator']]]:
        """
        Set of methods for validating the configuration. Maximum of 2. See Validator below for more details.
        """
        return pulumi.get(self, "validators")

