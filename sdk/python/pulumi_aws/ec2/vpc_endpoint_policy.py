# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VpcEndpointPolicyArgs', 'VpcEndpointPolicy']

@pulumi.input_type
class VpcEndpointPolicyArgs:
    def __init__(__self__, *,
                 vpc_endpoint_id: pulumi.Input[str],
                 policy: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VpcEndpointPolicy resource.
        :param pulumi.Input[str] vpc_endpoint_id: The VPC Endpoint ID.
        :param pulumi.Input[str] policy: A policy to attach to the endpoint that controls access to the service. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
        """
        pulumi.set(__self__, "vpc_endpoint_id", vpc_endpoint_id)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)

    @property
    @pulumi.getter(name="vpcEndpointId")
    def vpc_endpoint_id(self) -> pulumi.Input[str]:
        """
        The VPC Endpoint ID.
        """
        return pulumi.get(self, "vpc_endpoint_id")

    @vpc_endpoint_id.setter
    def vpc_endpoint_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_endpoint_id", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        """
        A policy to attach to the endpoint that controls access to the service. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)


@pulumi.input_type
class _VpcEndpointPolicyState:
    def __init__(__self__, *,
                 policy: Optional[pulumi.Input[str]] = None,
                 vpc_endpoint_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpcEndpointPolicy resources.
        :param pulumi.Input[str] policy: A policy to attach to the endpoint that controls access to the service. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
        :param pulumi.Input[str] vpc_endpoint_id: The VPC Endpoint ID.
        """
        if policy is not None:
            pulumi.set(__self__, "policy", policy)
        if vpc_endpoint_id is not None:
            pulumi.set(__self__, "vpc_endpoint_id", vpc_endpoint_id)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        """
        A policy to attach to the endpoint that controls access to the service. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter(name="vpcEndpointId")
    def vpc_endpoint_id(self) -> Optional[pulumi.Input[str]]:
        """
        The VPC Endpoint ID.
        """
        return pulumi.get(self, "vpc_endpoint_id")

    @vpc_endpoint_id.setter
    def vpc_endpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_endpoint_id", value)


class VpcEndpointPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 vpc_endpoint_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a VPC Endpoint Policy resource.

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumi_aws as aws

        example_vpc_endpoint_service = aws.ec2.get_vpc_endpoint_service(service="dynamodb")
        example_vpc = aws.ec2.Vpc("exampleVpc", cidr_block="10.0.0.0/16")
        example_vpc_endpoint = aws.ec2.VpcEndpoint("exampleVpcEndpoint",
            service_name=example_vpc_endpoint_service.service_name,
            vpc_id=example_vpc.id)
        example_vpc_endpoint_policy = aws.ec2.VpcEndpointPolicy("exampleVpcEndpointPolicy",
            vpc_endpoint_id=example_vpc_endpoint.id,
            policy=json.dumps({
                "Version": "2012-10-17",
                "Statement": [{
                    "Sid": "AllowAll",
                    "Effect": "Allow",
                    "Principal": {
                        "AWS": "*",
                    },
                    "Action": ["dynamodb:*"],
                    "Resource": "*",
                }],
            }))
        ```

        ## Import

        terraform import {

         to = aws_vpc_endpoint_policy.example

         id = "vpce-3ecf2a57" } Using `pulumi import`, import VPC Endpoint Policies using the `id`. For exampleconsole % pulumi import aws_vpc_endpoint_policy.example vpce-3ecf2a57

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy: A policy to attach to the endpoint that controls access to the service. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
        :param pulumi.Input[str] vpc_endpoint_id: The VPC Endpoint ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcEndpointPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a VPC Endpoint Policy resource.

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumi_aws as aws

        example_vpc_endpoint_service = aws.ec2.get_vpc_endpoint_service(service="dynamodb")
        example_vpc = aws.ec2.Vpc("exampleVpc", cidr_block="10.0.0.0/16")
        example_vpc_endpoint = aws.ec2.VpcEndpoint("exampleVpcEndpoint",
            service_name=example_vpc_endpoint_service.service_name,
            vpc_id=example_vpc.id)
        example_vpc_endpoint_policy = aws.ec2.VpcEndpointPolicy("exampleVpcEndpointPolicy",
            vpc_endpoint_id=example_vpc_endpoint.id,
            policy=json.dumps({
                "Version": "2012-10-17",
                "Statement": [{
                    "Sid": "AllowAll",
                    "Effect": "Allow",
                    "Principal": {
                        "AWS": "*",
                    },
                    "Action": ["dynamodb:*"],
                    "Resource": "*",
                }],
            }))
        ```

        ## Import

        terraform import {

         to = aws_vpc_endpoint_policy.example

         id = "vpce-3ecf2a57" } Using `pulumi import`, import VPC Endpoint Policies using the `id`. For exampleconsole % pulumi import aws_vpc_endpoint_policy.example vpce-3ecf2a57

        :param str resource_name: The name of the resource.
        :param VpcEndpointPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcEndpointPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 vpc_endpoint_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcEndpointPolicyArgs.__new__(VpcEndpointPolicyArgs)

            __props__.__dict__["policy"] = policy
            if vpc_endpoint_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_endpoint_id'")
            __props__.__dict__["vpc_endpoint_id"] = vpc_endpoint_id
        super(VpcEndpointPolicy, __self__).__init__(
            'aws:ec2/vpcEndpointPolicy:VpcEndpointPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            policy: Optional[pulumi.Input[str]] = None,
            vpc_endpoint_id: Optional[pulumi.Input[str]] = None) -> 'VpcEndpointPolicy':
        """
        Get an existing VpcEndpointPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy: A policy to attach to the endpoint that controls access to the service. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
        :param pulumi.Input[str] vpc_endpoint_id: The VPC Endpoint ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcEndpointPolicyState.__new__(_VpcEndpointPolicyState)

        __props__.__dict__["policy"] = policy
        __props__.__dict__["vpc_endpoint_id"] = vpc_endpoint_id
        return VpcEndpointPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[str]:
        """
        A policy to attach to the endpoint that controls access to the service. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="vpcEndpointId")
    def vpc_endpoint_id(self) -> pulumi.Output[str]:
        """
        The VPC Endpoint ID.
        """
        return pulumi.get(self, "vpc_endpoint_id")

