# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ApplicationVersionArgs', 'ApplicationVersion']

@pulumi.input_type
class ApplicationVersionArgs:
    def __init__(__self__, *,
                 application: pulumi.Input[str],
                 bucket: pulumi.Input[str],
                 key: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ApplicationVersion resource.
        :param pulumi.Input[str] application: Name of the Beanstalk Application the version is associated with.
        :param pulumi.Input[str] bucket: S3 bucket that contains the Application Version source bundle.
        :param pulumi.Input[str] key: S3 object that is the Application Version source bundle.
        :param pulumi.Input[str] description: Short description of the Application Version.
        :param pulumi.Input[bool] force_delete: On delete, force an Application Version to be deleted when it may be in use by multiple Elastic Beanstalk Environments.
        :param pulumi.Input[str] name: Unique name for the this Application Version.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of tags for the Elastic Beanstalk Application Version. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "application", application)
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "key", key)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if force_delete is not None:
            pulumi.set(__self__, "force_delete", force_delete)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def application(self) -> pulumi.Input[str]:
        """
        Name of the Beanstalk Application the version is associated with.
        """
        return pulumi.get(self, "application")

    @application.setter
    def application(self, value: pulumi.Input[str]):
        pulumi.set(self, "application", value)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        """
        S3 bucket that contains the Application Version source bundle.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        S3 object that is the Application Version source bundle.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Short description of the Application Version.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="forceDelete")
    def force_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        On delete, force an Application Version to be deleted when it may be in use by multiple Elastic Beanstalk Environments.
        """
        return pulumi.get(self, "force_delete")

    @force_delete.setter
    def force_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_delete", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name for the this Application Version.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of tags for the Elastic Beanstalk Application Version. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ApplicationVersionState:
    def __init__(__self__, *,
                 application: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ApplicationVersion resources.
        :param pulumi.Input[str] application: Name of the Beanstalk Application the version is associated with.
        :param pulumi.Input[str] arn: ARN assigned by AWS for this Elastic Beanstalk Application.
        :param pulumi.Input[str] bucket: S3 bucket that contains the Application Version source bundle.
        :param pulumi.Input[str] description: Short description of the Application Version.
        :param pulumi.Input[bool] force_delete: On delete, force an Application Version to be deleted when it may be in use by multiple Elastic Beanstalk Environments.
        :param pulumi.Input[str] key: S3 object that is the Application Version source bundle.
        :param pulumi.Input[str] name: Unique name for the this Application Version.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of tags for the Elastic Beanstalk Application Version. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if application is not None:
            pulumi.set(__self__, "application", application)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if force_delete is not None:
            pulumi.set(__self__, "force_delete", force_delete)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def application(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Beanstalk Application the version is associated with.
        """
        return pulumi.get(self, "application")

    @application.setter
    def application(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN assigned by AWS for this Elastic Beanstalk Application.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[str]]:
        """
        S3 bucket that contains the Application Version source bundle.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Short description of the Application Version.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="forceDelete")
    def force_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        On delete, force an Application Version to be deleted when it may be in use by multiple Elastic Beanstalk Environments.
        """
        return pulumi.get(self, "force_delete")

    @force_delete.setter
    def force_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_delete", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        S3 object that is the Application Version source bundle.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name for the this Application Version.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of tags for the Elastic Beanstalk Application Version. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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


class ApplicationVersion(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application: Optional[pulumi.Input[str]] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides an Elastic Beanstalk Application Version Resource. Elastic Beanstalk allows
        you to deploy and manage applications in the AWS cloud without worrying about
        the infrastructure that runs those applications.

        This resource creates a Beanstalk Application Version that can be deployed to a Beanstalk
        Environment.

        > **NOTE on Application Version Resource:**  When using the Application Version resource with multiple
        Elastic Beanstalk Environments it is possible that an error may be returned
        when attempting to delete an Application Version while it is still in use by a different environment.
        To work around this you can either create each environment in a separate AWS account or create your `elasticbeanstalk.ApplicationVersion` resources with a unique names in your Elastic Beanstalk Application. For example &lt;revision&gt;-&lt;environment&gt;.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        default_bucket_v2 = aws.s3.BucketV2("defaultBucketV2")
        default_bucket_objectv2 = aws.s3.BucketObjectv2("defaultBucketObjectv2",
            bucket=default_bucket_v2.id,
            key="beanstalk/go-v1.zip",
            source=pulumi.FileAsset("go-v1.zip"))
        default_application = aws.elasticbeanstalk.Application("defaultApplication", description="tf-test-desc")
        default_application_version = aws.elasticbeanstalk.ApplicationVersion("defaultApplicationVersion",
            application="tf-test-name",
            description="application version",
            bucket=default_bucket_v2.id,
            key=default_bucket_objectv2.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application: Name of the Beanstalk Application the version is associated with.
        :param pulumi.Input[str] bucket: S3 bucket that contains the Application Version source bundle.
        :param pulumi.Input[str] description: Short description of the Application Version.
        :param pulumi.Input[bool] force_delete: On delete, force an Application Version to be deleted when it may be in use by multiple Elastic Beanstalk Environments.
        :param pulumi.Input[str] key: S3 object that is the Application Version source bundle.
        :param pulumi.Input[str] name: Unique name for the this Application Version.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of tags for the Elastic Beanstalk Application Version. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationVersionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Elastic Beanstalk Application Version Resource. Elastic Beanstalk allows
        you to deploy and manage applications in the AWS cloud without worrying about
        the infrastructure that runs those applications.

        This resource creates a Beanstalk Application Version that can be deployed to a Beanstalk
        Environment.

        > **NOTE on Application Version Resource:**  When using the Application Version resource with multiple
        Elastic Beanstalk Environments it is possible that an error may be returned
        when attempting to delete an Application Version while it is still in use by a different environment.
        To work around this you can either create each environment in a separate AWS account or create your `elasticbeanstalk.ApplicationVersion` resources with a unique names in your Elastic Beanstalk Application. For example &lt;revision&gt;-&lt;environment&gt;.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        default_bucket_v2 = aws.s3.BucketV2("defaultBucketV2")
        default_bucket_objectv2 = aws.s3.BucketObjectv2("defaultBucketObjectv2",
            bucket=default_bucket_v2.id,
            key="beanstalk/go-v1.zip",
            source=pulumi.FileAsset("go-v1.zip"))
        default_application = aws.elasticbeanstalk.Application("defaultApplication", description="tf-test-desc")
        default_application_version = aws.elasticbeanstalk.ApplicationVersion("defaultApplicationVersion",
            application="tf-test-name",
            description="application version",
            bucket=default_bucket_v2.id,
            key=default_bucket_objectv2.id)
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationVersionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationVersionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application: Optional[pulumi.Input[str]] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationVersionArgs.__new__(ApplicationVersionArgs)

            if application is None and not opts.urn:
                raise TypeError("Missing required property 'application'")
            __props__.__dict__["application"] = application
            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            __props__.__dict__["description"] = description
            __props__.__dict__["force_delete"] = force_delete
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        super(ApplicationVersion, __self__).__init__(
            'aws:elasticbeanstalk/applicationVersion:ApplicationVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            force_delete: Optional[pulumi.Input[bool]] = None,
            key: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'ApplicationVersion':
        """
        Get an existing ApplicationVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application: Name of the Beanstalk Application the version is associated with.
        :param pulumi.Input[str] arn: ARN assigned by AWS for this Elastic Beanstalk Application.
        :param pulumi.Input[str] bucket: S3 bucket that contains the Application Version source bundle.
        :param pulumi.Input[str] description: Short description of the Application Version.
        :param pulumi.Input[bool] force_delete: On delete, force an Application Version to be deleted when it may be in use by multiple Elastic Beanstalk Environments.
        :param pulumi.Input[str] key: S3 object that is the Application Version source bundle.
        :param pulumi.Input[str] name: Unique name for the this Application Version.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of tags for the Elastic Beanstalk Application Version. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationVersionState.__new__(_ApplicationVersionState)

        __props__.__dict__["application"] = application
        __props__.__dict__["arn"] = arn
        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["description"] = description
        __props__.__dict__["force_delete"] = force_delete
        __props__.__dict__["key"] = key
        __props__.__dict__["name"] = name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return ApplicationVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def application(self) -> pulumi.Output[str]:
        """
        Name of the Beanstalk Application the version is associated with.
        """
        return pulumi.get(self, "application")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN assigned by AWS for this Elastic Beanstalk Application.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        S3 bucket that contains the Application Version source bundle.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Short description of the Application Version.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="forceDelete")
    def force_delete(self) -> pulumi.Output[Optional[bool]]:
        """
        On delete, force an Application Version to be deleted when it may be in use by multiple Elastic Beanstalk Environments.
        """
        return pulumi.get(self, "force_delete")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        S3 object that is the Application Version source bundle.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique name for the this Application Version.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of tags for the Elastic Beanstalk Application Version. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

