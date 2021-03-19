# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['BucketLifecycleConfiguration']


class BucketLifecycleConfiguration(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketLifecycleConfigurationRuleArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to manage an S3 Control Bucket Lifecycle Configuration.

        > **NOTE:** Each S3 Control Bucket can only have one Lifecycle Configuration. Using multiple of this resource against the same S3 Control Bucket will result in perpetual differences each provider run.

        > This functionality is for managing [S3 on Outposts](https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html). To manage S3 Bucket Lifecycle Configurations in an AWS Partition, see the `s3.Bucket` resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.s3control.BucketLifecycleConfiguration("example",
            bucket=aws_s3control_bucket["example"]["arn"],
            rules=[
                aws.s3control.BucketLifecycleConfigurationRuleArgs(
                    expiration=aws.s3control.BucketLifecycleConfigurationRuleExpirationArgs(
                        days=365,
                    ),
                    filter=aws.s3control.BucketLifecycleConfigurationRuleFilterArgs(
                        prefix="logs/",
                    ),
                    id="logs",
                ),
                aws.s3control.BucketLifecycleConfigurationRuleArgs(
                    expiration=aws.s3control.BucketLifecycleConfigurationRuleExpirationArgs(
                        days=7,
                    ),
                    filter=aws.s3control.BucketLifecycleConfigurationRuleFilterArgs(
                        prefix="temp/",
                    ),
                    id="temp",
                ),
            ])
        ```

        ## Import

        S3 Control Bucket Lifecycle Configurations can be imported using the Amazon Resource Name (ARN), e.g.

        ```sh
         $ pulumi import aws:s3control/bucketLifecycleConfiguration:BucketLifecycleConfiguration example arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-12345678/bucket/example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: Amazon Resource Name (ARN) of the bucket.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketLifecycleConfigurationRuleArgs']]]] rules: Configuration block(s) containing lifecycle rules for the bucket.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__['bucket'] = bucket
            if rules is None and not opts.urn:
                raise TypeError("Missing required property 'rules'")
            __props__['rules'] = rules
        super(BucketLifecycleConfiguration, __self__).__init__(
            'aws:s3control/bucketLifecycleConfiguration:BucketLifecycleConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketLifecycleConfigurationRuleArgs']]]]] = None) -> 'BucketLifecycleConfiguration':
        """
        Get an existing BucketLifecycleConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: Amazon Resource Name (ARN) of the bucket.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketLifecycleConfigurationRuleArgs']]]] rules: Configuration block(s) containing lifecycle rules for the bucket.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bucket"] = bucket
        __props__["rules"] = rules
        return BucketLifecycleConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the bucket.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.BucketLifecycleConfigurationRule']]:
        """
        Configuration block(s) containing lifecycle rules for the bucket.
        """
        return pulumi.get(self, "rules")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

