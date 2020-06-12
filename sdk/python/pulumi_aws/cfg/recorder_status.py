# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class RecorderStatus(pulumi.CustomResource):
    is_enabled: pulumi.Output[bool]
    """
    Whether the configuration recorder should be enabled or disabled.
    """
    name: pulumi.Output[str]
    """
    The name of the recorder
    """
    def __init__(__self__, resource_name, opts=None, is_enabled=None, name=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages status (recording / stopped) of an AWS Config Configuration Recorder.

        > **Note:** Starting Configuration Recorder requires a `Delivery Channel` to be present. Use of `depends_on` (as shown below) is recommended to avoid race conditions.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        foo_recorder_status = aws.cfg.RecorderStatus("fooRecorderStatus", is_enabled=True)
        role = aws.iam.Role("role", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Action": "sts:AssumeRole",
              "Principal": {
                "Service": "config.amazonaws.com"
              },
              "Effect": "Allow",
              "Sid": ""
            }
          ]
        }

        \"\"\")
        role_policy_attachment = aws.iam.RolePolicyAttachment("rolePolicyAttachment",
            policy_arn="arn:aws:iam::aws:policy/service-role/AWSConfigRole",
            role=role.name)
        bucket = aws.s3.Bucket("bucket")
        foo_delivery_channel = aws.cfg.DeliveryChannel("fooDeliveryChannel", s3_bucket_name=bucket.bucket)
        foo_recorder = aws.cfg.Recorder("fooRecorder", role_arn=role.arn)
        role_policy = aws.iam.RolePolicy("rolePolicy",
            policy=pulumi.Output.all(bucket.arn, bucket.arn).apply(lambda bucketArn, bucketArn1: f\"\"\"{{
          "Version": "2012-10-17",
          "Statement": [
            {{
              "Action": [
                "s3:*"
              ],
              "Effect": "Allow",
              "Resource": [
                "{bucket_arn}",
                "{bucket_arn1}/*"
              ]
            }}
          ]
        }}

        \"\"\"),
            role=role.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] is_enabled: Whether the configuration recorder should be enabled or disabled.
        :param pulumi.Input[str] name: The name of the recorder
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if is_enabled is None:
                raise TypeError("Missing required property 'is_enabled'")
            __props__['is_enabled'] = is_enabled
            __props__['name'] = name
        super(RecorderStatus, __self__).__init__(
            'aws:cfg/recorderStatus:RecorderStatus',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, is_enabled=None, name=None):
        """
        Get an existing RecorderStatus resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] is_enabled: Whether the configuration recorder should be enabled or disabled.
        :param pulumi.Input[str] name: The name of the recorder
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["is_enabled"] = is_enabled
        __props__["name"] = name
        return RecorderStatus(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

