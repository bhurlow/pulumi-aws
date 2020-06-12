# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class EventPermission(pulumi.CustomResource):
    action: pulumi.Output[str]
    """
    The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
    """
    condition: pulumi.Output[dict]
    """
    Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.

      * `key` (`str`) - Key for the condition. Valid values: `aws:PrincipalOrgID`.
      * `type` (`str`) - Type of condition. Value values: `StringEquals`.
      * `value` (`str`) - Value for the key.
    """
    principal: pulumi.Output[str]
    """
    The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
    """
    statement_id: pulumi.Output[str]
    """
    An identifier string for the external account that you are granting permissions to.
    """
    def __init__(__self__, resource_name, opts=None, action=None, condition=None, principal=None, statement_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource to create a CloudWatch Events permission to support cross-account events in the current account default event bus.

        ## Example Usage

        ### Account Access

        ```python
        import pulumi
        import pulumi_aws as aws

        dev_account_access = aws.cloudwatch.EventPermission("devAccountAccess",
            principal="123456789012",
            statement_id="DevAccountAccess")
        ```

        ### Organization Access

        ```python
        import pulumi
        import pulumi_aws as aws

        organization_access = aws.cloudwatch.EventPermission("organizationAccess",
            condition={
                "key": "aws:PrincipalOrgID",
                "type": "StringEquals",
                "value": aws_organizations_organization["example"]["id"],
            },
            principal="*",
            statement_id="OrganizationAccess")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
        :param pulumi.Input[dict] condition: Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
        :param pulumi.Input[str] principal: The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
        :param pulumi.Input[str] statement_id: An identifier string for the external account that you are granting permissions to.

        The **condition** object supports the following:

          * `key` (`pulumi.Input[str]`) - Key for the condition. Valid values: `aws:PrincipalOrgID`.
          * `type` (`pulumi.Input[str]`) - Type of condition. Value values: `StringEquals`.
          * `value` (`pulumi.Input[str]`) - Value for the key.
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

            __props__['action'] = action
            __props__['condition'] = condition
            if principal is None:
                raise TypeError("Missing required property 'principal'")
            __props__['principal'] = principal
            if statement_id is None:
                raise TypeError("Missing required property 'statement_id'")
            __props__['statement_id'] = statement_id
        super(EventPermission, __self__).__init__(
            'aws:cloudwatch/eventPermission:EventPermission',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, action=None, condition=None, principal=None, statement_id=None):
        """
        Get an existing EventPermission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
        :param pulumi.Input[dict] condition: Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
        :param pulumi.Input[str] principal: The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
        :param pulumi.Input[str] statement_id: An identifier string for the external account that you are granting permissions to.

        The **condition** object supports the following:

          * `key` (`pulumi.Input[str]`) - Key for the condition. Valid values: `aws:PrincipalOrgID`.
          * `type` (`pulumi.Input[str]`) - Type of condition. Value values: `StringEquals`.
          * `value` (`pulumi.Input[str]`) - Value for the key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["action"] = action
        __props__["condition"] = condition
        __props__["principal"] = principal
        __props__["statement_id"] = statement_id
        return EventPermission(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

