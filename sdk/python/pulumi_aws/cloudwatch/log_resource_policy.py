# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LogResourcePolicyArgs', 'LogResourcePolicy']

@pulumi.input_type
class LogResourcePolicyArgs:
    def __init__(__self__, *,
                 policy_document: pulumi.Input[str],
                 policy_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a LogResourcePolicy resource.
        :param pulumi.Input[str] policy_document: Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
        :param pulumi.Input[str] policy_name: Name of the resource policy.
        """
        pulumi.set(__self__, "policy_document", policy_document)
        pulumi.set(__self__, "policy_name", policy_name)

    @property
    @pulumi.getter(name="policyDocument")
    def policy_document(self) -> pulumi.Input[str]:
        """
        Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
        """
        return pulumi.get(self, "policy_document")

    @policy_document.setter
    def policy_document(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_document", value)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> pulumi.Input[str]:
        """
        Name of the resource policy.
        """
        return pulumi.get(self, "policy_name")

    @policy_name.setter
    def policy_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_name", value)


@pulumi.input_type
class _LogResourcePolicyState:
    def __init__(__self__, *,
                 policy_document: Optional[pulumi.Input[str]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LogResourcePolicy resources.
        :param pulumi.Input[str] policy_document: Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
        :param pulumi.Input[str] policy_name: Name of the resource policy.
        """
        if policy_document is not None:
            pulumi.set(__self__, "policy_document", policy_document)
        if policy_name is not None:
            pulumi.set(__self__, "policy_name", policy_name)

    @property
    @pulumi.getter(name="policyDocument")
    def policy_document(self) -> Optional[pulumi.Input[str]]:
        """
        Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
        """
        return pulumi.get(self, "policy_document")

    @policy_document.setter
    def policy_document(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_document", value)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource policy.
        """
        return pulumi.get(self, "policy_name")

    @policy_name.setter
    def policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_name", value)


class LogResourcePolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_document: Optional[pulumi.Input[str]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage a CloudWatch log resource policy.

        ## Example Usage
        ### Elasticsearch Log Publishing

        ```python
        import pulumi
        import pulumi_aws as aws

        elasticsearch_log_publishing_policy_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            actions=[
                "logs:CreateLogStream",
                "logs:PutLogEvents",
                "logs:PutLogEventsBatch",
            ],
            resources=["arn:aws:logs:*"],
            principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                identifiers=["es.amazonaws.com"],
                type="Service",
            )],
        )])
        elasticsearch_log_publishing_policy_log_resource_policy = aws.cloudwatch.LogResourcePolicy("elasticsearch-log-publishing-policyLogResourcePolicy",
            policy_document=elasticsearch_log_publishing_policy_policy_document.json,
            policy_name="elasticsearch-log-publishing-policy")
        ```
        ### Route53 Query Logging

        ```python
        import pulumi
        import pulumi_aws as aws

        route53_query_logging_policy_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            actions=[
                "logs:CreateLogStream",
                "logs:PutLogEvents",
            ],
            resources=["arn:aws:logs:*:*:log-group:/aws/route53/*"],
            principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                identifiers=["route53.amazonaws.com"],
                type="Service",
            )],
        )])
        route53_query_logging_policy_log_resource_policy = aws.cloudwatch.LogResourcePolicy("route53-query-logging-policyLogResourcePolicy",
            policy_document=route53_query_logging_policy_policy_document.json,
            policy_name="route53-query-logging-policy")
        ```

        ## Import

        terraform import {

         to = aws_cloudwatch_log_resource_policy.MyPolicy

         id = "MyPolicy" } Using `pulumi import`, import CloudWatch log resource policies using the policy name. For exampleconsole % pulumi import aws_cloudwatch_log_resource_policy.MyPolicy MyPolicy

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_document: Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
        :param pulumi.Input[str] policy_name: Name of the resource policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LogResourcePolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage a CloudWatch log resource policy.

        ## Example Usage
        ### Elasticsearch Log Publishing

        ```python
        import pulumi
        import pulumi_aws as aws

        elasticsearch_log_publishing_policy_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            actions=[
                "logs:CreateLogStream",
                "logs:PutLogEvents",
                "logs:PutLogEventsBatch",
            ],
            resources=["arn:aws:logs:*"],
            principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                identifiers=["es.amazonaws.com"],
                type="Service",
            )],
        )])
        elasticsearch_log_publishing_policy_log_resource_policy = aws.cloudwatch.LogResourcePolicy("elasticsearch-log-publishing-policyLogResourcePolicy",
            policy_document=elasticsearch_log_publishing_policy_policy_document.json,
            policy_name="elasticsearch-log-publishing-policy")
        ```
        ### Route53 Query Logging

        ```python
        import pulumi
        import pulumi_aws as aws

        route53_query_logging_policy_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            actions=[
                "logs:CreateLogStream",
                "logs:PutLogEvents",
            ],
            resources=["arn:aws:logs:*:*:log-group:/aws/route53/*"],
            principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                identifiers=["route53.amazonaws.com"],
                type="Service",
            )],
        )])
        route53_query_logging_policy_log_resource_policy = aws.cloudwatch.LogResourcePolicy("route53-query-logging-policyLogResourcePolicy",
            policy_document=route53_query_logging_policy_policy_document.json,
            policy_name="route53-query-logging-policy")
        ```

        ## Import

        terraform import {

         to = aws_cloudwatch_log_resource_policy.MyPolicy

         id = "MyPolicy" } Using `pulumi import`, import CloudWatch log resource policies using the policy name. For exampleconsole % pulumi import aws_cloudwatch_log_resource_policy.MyPolicy MyPolicy

        :param str resource_name: The name of the resource.
        :param LogResourcePolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogResourcePolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_document: Optional[pulumi.Input[str]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LogResourcePolicyArgs.__new__(LogResourcePolicyArgs)

            if policy_document is None and not opts.urn:
                raise TypeError("Missing required property 'policy_document'")
            __props__.__dict__["policy_document"] = policy_document
            if policy_name is None and not opts.urn:
                raise TypeError("Missing required property 'policy_name'")
            __props__.__dict__["policy_name"] = policy_name
        super(LogResourcePolicy, __self__).__init__(
            'aws:cloudwatch/logResourcePolicy:LogResourcePolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            policy_document: Optional[pulumi.Input[str]] = None,
            policy_name: Optional[pulumi.Input[str]] = None) -> 'LogResourcePolicy':
        """
        Get an existing LogResourcePolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_document: Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
        :param pulumi.Input[str] policy_name: Name of the resource policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogResourcePolicyState.__new__(_LogResourcePolicyState)

        __props__.__dict__["policy_document"] = policy_document
        __props__.__dict__["policy_name"] = policy_name
        return LogResourcePolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="policyDocument")
    def policy_document(self) -> pulumi.Output[str]:
        """
        Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
        """
        return pulumi.get(self, "policy_document")

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> pulumi.Output[str]:
        """
        Name of the resource policy.
        """
        return pulumi.get(self, "policy_name")

