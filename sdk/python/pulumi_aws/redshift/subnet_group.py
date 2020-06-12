# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SubnetGroup(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the Redshift Subnet group name
    """
    description: pulumi.Output[str]
    """
    The description of the Redshift Subnet group. Defaults to "Managed by Pulumi".
    """
    name: pulumi.Output[str]
    """
    The name of the Redshift Subnet group.
    """
    subnet_ids: pulumi.Output[list]
    """
    An array of VPC subnet IDs.
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, description=None, name=None, subnet_ids=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates a new Amazon Redshift subnet group. You must provide a list of one or more subnets in your existing Amazon Virtual Private Cloud (Amazon VPC) when creating Amazon Redshift subnet group.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        foo_vpc = aws.ec2.Vpc("fooVpc", cidr_block="10.1.0.0/16")
        foo_subnet = aws.ec2.Subnet("fooSubnet",
            availability_zone="us-west-2a",
            cidr_block="10.1.1.0/24",
            tags={
                "Name": "tf-dbsubnet-test-1",
            },
            vpc_id=foo_vpc.id)
        bar = aws.ec2.Subnet("bar",
            availability_zone="us-west-2b",
            cidr_block="10.1.2.0/24",
            tags={
                "Name": "tf-dbsubnet-test-2",
            },
            vpc_id=foo_vpc.id)
        foo_subnet_group = aws.redshift.SubnetGroup("fooSubnetGroup",
            subnet_ids=[
                foo_subnet.id,
                bar.id,
            ],
            tags={
                "environment": "Production",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the Redshift Subnet group. Defaults to "Managed by Pulumi".
        :param pulumi.Input[str] name: The name of the Redshift Subnet group.
        :param pulumi.Input[list] subnet_ids: An array of VPC subnet IDs.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
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

            if description is None:
                description = 'Managed by Pulumi'
            __props__['description'] = description
            __props__['name'] = name
            if subnet_ids is None:
                raise TypeError("Missing required property 'subnet_ids'")
            __props__['subnet_ids'] = subnet_ids
            __props__['tags'] = tags
            __props__['arn'] = None
        super(SubnetGroup, __self__).__init__(
            'aws:redshift/subnetGroup:SubnetGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, description=None, name=None, subnet_ids=None, tags=None):
        """
        Get an existing SubnetGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the Redshift Subnet group name
        :param pulumi.Input[str] description: The description of the Redshift Subnet group. Defaults to "Managed by Pulumi".
        :param pulumi.Input[str] name: The name of the Redshift Subnet group.
        :param pulumi.Input[list] subnet_ids: An array of VPC subnet IDs.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["description"] = description
        __props__["name"] = name
        __props__["subnet_ids"] = subnet_ids
        __props__["tags"] = tags
        return SubnetGroup(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

