# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AccountPublicAccessBlockArgs', 'AccountPublicAccessBlock']

@pulumi.input_type
class AccountPublicAccessBlockArgs:
    def __init__(__self__, *,
                 account_id: Optional[pulumi.Input[str]] = None,
                 block_public_acls: Optional[pulumi.Input[bool]] = None,
                 block_public_policy: Optional[pulumi.Input[bool]] = None,
                 ignore_public_acls: Optional[pulumi.Input[bool]] = None,
                 restrict_public_buckets: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AccountPublicAccessBlock resource.
        :param pulumi.Input[str] account_id: AWS account ID to configure. Defaults to automatically determined account ID of the this provider AWS provider.
        :param pulumi.Input[bool] block_public_acls: Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
               * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
               * PUT Object calls fail if the request includes a public ACL.
        :param pulumi.Input[bool] block_public_policy: Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing bucket policies. When set to `true` causes Amazon S3 to:
               * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
        :param pulumi.Input[bool] ignore_public_acls: Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
               * Ignore all public ACLs on buckets in this account and any objects that they contain.
        :param pulumi.Input[bool] restrict_public_buckets: Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
               * Only the bucket owner and AWS Services can access buckets with public policies.
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if block_public_acls is not None:
            pulumi.set(__self__, "block_public_acls", block_public_acls)
        if block_public_policy is not None:
            pulumi.set(__self__, "block_public_policy", block_public_policy)
        if ignore_public_acls is not None:
            pulumi.set(__self__, "ignore_public_acls", ignore_public_acls)
        if restrict_public_buckets is not None:
            pulumi.set(__self__, "restrict_public_buckets", restrict_public_buckets)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        """
        AWS account ID to configure. Defaults to automatically determined account ID of the this provider AWS provider.
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="blockPublicAcls")
    def block_public_acls(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
        * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
        * PUT Object calls fail if the request includes a public ACL.
        """
        return pulumi.get(self, "block_public_acls")

    @block_public_acls.setter
    def block_public_acls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "block_public_acls", value)

    @property
    @pulumi.getter(name="blockPublicPolicy")
    def block_public_policy(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing bucket policies. When set to `true` causes Amazon S3 to:
        * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
        """
        return pulumi.get(self, "block_public_policy")

    @block_public_policy.setter
    def block_public_policy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "block_public_policy", value)

    @property
    @pulumi.getter(name="ignorePublicAcls")
    def ignore_public_acls(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
        * Ignore all public ACLs on buckets in this account and any objects that they contain.
        """
        return pulumi.get(self, "ignore_public_acls")

    @ignore_public_acls.setter
    def ignore_public_acls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_public_acls", value)

    @property
    @pulumi.getter(name="restrictPublicBuckets")
    def restrict_public_buckets(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
        * Only the bucket owner and AWS Services can access buckets with public policies.
        """
        return pulumi.get(self, "restrict_public_buckets")

    @restrict_public_buckets.setter
    def restrict_public_buckets(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "restrict_public_buckets", value)


@pulumi.input_type
class _AccountPublicAccessBlockState:
    def __init__(__self__, *,
                 account_id: Optional[pulumi.Input[str]] = None,
                 block_public_acls: Optional[pulumi.Input[bool]] = None,
                 block_public_policy: Optional[pulumi.Input[bool]] = None,
                 ignore_public_acls: Optional[pulumi.Input[bool]] = None,
                 restrict_public_buckets: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering AccountPublicAccessBlock resources.
        :param pulumi.Input[str] account_id: AWS account ID to configure. Defaults to automatically determined account ID of the this provider AWS provider.
        :param pulumi.Input[bool] block_public_acls: Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
               * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
               * PUT Object calls fail if the request includes a public ACL.
        :param pulumi.Input[bool] block_public_policy: Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing bucket policies. When set to `true` causes Amazon S3 to:
               * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
        :param pulumi.Input[bool] ignore_public_acls: Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
               * Ignore all public ACLs on buckets in this account and any objects that they contain.
        :param pulumi.Input[bool] restrict_public_buckets: Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
               * Only the bucket owner and AWS Services can access buckets with public policies.
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if block_public_acls is not None:
            pulumi.set(__self__, "block_public_acls", block_public_acls)
        if block_public_policy is not None:
            pulumi.set(__self__, "block_public_policy", block_public_policy)
        if ignore_public_acls is not None:
            pulumi.set(__self__, "ignore_public_acls", ignore_public_acls)
        if restrict_public_buckets is not None:
            pulumi.set(__self__, "restrict_public_buckets", restrict_public_buckets)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        """
        AWS account ID to configure. Defaults to automatically determined account ID of the this provider AWS provider.
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="blockPublicAcls")
    def block_public_acls(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
        * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
        * PUT Object calls fail if the request includes a public ACL.
        """
        return pulumi.get(self, "block_public_acls")

    @block_public_acls.setter
    def block_public_acls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "block_public_acls", value)

    @property
    @pulumi.getter(name="blockPublicPolicy")
    def block_public_policy(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing bucket policies. When set to `true` causes Amazon S3 to:
        * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
        """
        return pulumi.get(self, "block_public_policy")

    @block_public_policy.setter
    def block_public_policy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "block_public_policy", value)

    @property
    @pulumi.getter(name="ignorePublicAcls")
    def ignore_public_acls(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
        * Ignore all public ACLs on buckets in this account and any objects that they contain.
        """
        return pulumi.get(self, "ignore_public_acls")

    @ignore_public_acls.setter
    def ignore_public_acls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_public_acls", value)

    @property
    @pulumi.getter(name="restrictPublicBuckets")
    def restrict_public_buckets(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
        * Only the bucket owner and AWS Services can access buckets with public policies.
        """
        return pulumi.get(self, "restrict_public_buckets")

    @restrict_public_buckets.setter
    def restrict_public_buckets(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "restrict_public_buckets", value)


class AccountPublicAccessBlock(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 block_public_acls: Optional[pulumi.Input[bool]] = None,
                 block_public_policy: Optional[pulumi.Input[bool]] = None,
                 ignore_public_acls: Optional[pulumi.Input[bool]] = None,
                 restrict_public_buckets: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Manages S3 account-level Public Access Block configuration. For more information about these settings, see the [AWS S3 Block Public Access documentation](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html).

        > **NOTE:** Each AWS account may only have one S3 Public Access Block configuration. Multiple configurations of the resource against the same AWS account will cause a perpetual difference.

        > Advanced usage: To use a custom API endpoint for this resource, use the `s3control` endpoint provider configuration, not the `s3` endpoint provider configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.s3.AccountPublicAccessBlock("example",
            block_public_acls=True,
            block_public_policy=True)
        ```

        ## Import

        terraform import {

         to = aws_s3_account_public_access_block.example

         id = "123456789012" } Using `pulumi import`, import `aws_s3_account_public_access_block` using the AWS account ID. For exampleconsole % pulumi import aws_s3_account_public_access_block.example 123456789012

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: AWS account ID to configure. Defaults to automatically determined account ID of the this provider AWS provider.
        :param pulumi.Input[bool] block_public_acls: Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
               * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
               * PUT Object calls fail if the request includes a public ACL.
        :param pulumi.Input[bool] block_public_policy: Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing bucket policies. When set to `true` causes Amazon S3 to:
               * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
        :param pulumi.Input[bool] ignore_public_acls: Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
               * Ignore all public ACLs on buckets in this account and any objects that they contain.
        :param pulumi.Input[bool] restrict_public_buckets: Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
               * Only the bucket owner and AWS Services can access buckets with public policies.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AccountPublicAccessBlockArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages S3 account-level Public Access Block configuration. For more information about these settings, see the [AWS S3 Block Public Access documentation](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html).

        > **NOTE:** Each AWS account may only have one S3 Public Access Block configuration. Multiple configurations of the resource against the same AWS account will cause a perpetual difference.

        > Advanced usage: To use a custom API endpoint for this resource, use the `s3control` endpoint provider configuration, not the `s3` endpoint provider configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.s3.AccountPublicAccessBlock("example",
            block_public_acls=True,
            block_public_policy=True)
        ```

        ## Import

        terraform import {

         to = aws_s3_account_public_access_block.example

         id = "123456789012" } Using `pulumi import`, import `aws_s3_account_public_access_block` using the AWS account ID. For exampleconsole % pulumi import aws_s3_account_public_access_block.example 123456789012

        :param str resource_name: The name of the resource.
        :param AccountPublicAccessBlockArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountPublicAccessBlockArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 block_public_acls: Optional[pulumi.Input[bool]] = None,
                 block_public_policy: Optional[pulumi.Input[bool]] = None,
                 ignore_public_acls: Optional[pulumi.Input[bool]] = None,
                 restrict_public_buckets: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccountPublicAccessBlockArgs.__new__(AccountPublicAccessBlockArgs)

            __props__.__dict__["account_id"] = account_id
            __props__.__dict__["block_public_acls"] = block_public_acls
            __props__.__dict__["block_public_policy"] = block_public_policy
            __props__.__dict__["ignore_public_acls"] = ignore_public_acls
            __props__.__dict__["restrict_public_buckets"] = restrict_public_buckets
        super(AccountPublicAccessBlock, __self__).__init__(
            'aws:s3/accountPublicAccessBlock:AccountPublicAccessBlock',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_id: Optional[pulumi.Input[str]] = None,
            block_public_acls: Optional[pulumi.Input[bool]] = None,
            block_public_policy: Optional[pulumi.Input[bool]] = None,
            ignore_public_acls: Optional[pulumi.Input[bool]] = None,
            restrict_public_buckets: Optional[pulumi.Input[bool]] = None) -> 'AccountPublicAccessBlock':
        """
        Get an existing AccountPublicAccessBlock resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: AWS account ID to configure. Defaults to automatically determined account ID of the this provider AWS provider.
        :param pulumi.Input[bool] block_public_acls: Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
               * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
               * PUT Object calls fail if the request includes a public ACL.
        :param pulumi.Input[bool] block_public_policy: Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing bucket policies. When set to `true` causes Amazon S3 to:
               * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
        :param pulumi.Input[bool] ignore_public_acls: Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
               * Ignore all public ACLs on buckets in this account and any objects that they contain.
        :param pulumi.Input[bool] restrict_public_buckets: Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
               * Only the bucket owner and AWS Services can access buckets with public policies.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccountPublicAccessBlockState.__new__(_AccountPublicAccessBlockState)

        __props__.__dict__["account_id"] = account_id
        __props__.__dict__["block_public_acls"] = block_public_acls
        __props__.__dict__["block_public_policy"] = block_public_policy
        __props__.__dict__["ignore_public_acls"] = ignore_public_acls
        __props__.__dict__["restrict_public_buckets"] = restrict_public_buckets
        return AccountPublicAccessBlock(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[str]:
        """
        AWS account ID to configure. Defaults to automatically determined account ID of the this provider AWS provider.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="blockPublicAcls")
    def block_public_acls(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
        * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
        * PUT Object calls fail if the request includes a public ACL.
        """
        return pulumi.get(self, "block_public_acls")

    @property
    @pulumi.getter(name="blockPublicPolicy")
    def block_public_policy(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing bucket policies. When set to `true` causes Amazon S3 to:
        * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
        """
        return pulumi.get(self, "block_public_policy")

    @property
    @pulumi.getter(name="ignorePublicAcls")
    def ignore_public_acls(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
        * Ignore all public ACLs on buckets in this account and any objects that they contain.
        """
        return pulumi.get(self, "ignore_public_acls")

    @property
    @pulumi.getter(name="restrictPublicBuckets")
    def restrict_public_buckets(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
        * Only the bucket owner and AWS Services can access buckets with public policies.
        """
        return pulumi.get(self, "restrict_public_buckets")

