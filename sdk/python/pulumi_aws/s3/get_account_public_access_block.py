# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetAccountPublicAccessBlockResult',
    'AwaitableGetAccountPublicAccessBlockResult',
    'get_account_public_access_block',
    'get_account_public_access_block_output',
]

@pulumi.output_type
class GetAccountPublicAccessBlockResult:
    """
    A collection of values returned by getAccountPublicAccessBlock.
    """
    def __init__(__self__, account_id=None, block_public_acls=None, block_public_policy=None, id=None, ignore_public_acls=None, restrict_public_buckets=None):
        if account_id and not isinstance(account_id, str):
            raise TypeError("Expected argument 'account_id' to be a str")
        pulumi.set(__self__, "account_id", account_id)
        if block_public_acls and not isinstance(block_public_acls, bool):
            raise TypeError("Expected argument 'block_public_acls' to be a bool")
        pulumi.set(__self__, "block_public_acls", block_public_acls)
        if block_public_policy and not isinstance(block_public_policy, bool):
            raise TypeError("Expected argument 'block_public_policy' to be a bool")
        pulumi.set(__self__, "block_public_policy", block_public_policy)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ignore_public_acls and not isinstance(ignore_public_acls, bool):
            raise TypeError("Expected argument 'ignore_public_acls' to be a bool")
        pulumi.set(__self__, "ignore_public_acls", ignore_public_acls)
        if restrict_public_buckets and not isinstance(restrict_public_buckets, bool):
            raise TypeError("Expected argument 'restrict_public_buckets' to be a bool")
        pulumi.set(__self__, "restrict_public_buckets", restrict_public_buckets)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[str]:
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="blockPublicAcls")
    def block_public_acls(self) -> bool:
        """
        Whether or not Amazon S3 should block public ACLs for buckets in this account is enabled. Returns as `true` or `false`.
        """
        return pulumi.get(self, "block_public_acls")

    @property
    @pulumi.getter(name="blockPublicPolicy")
    def block_public_policy(self) -> bool:
        """
        Whether or not Amazon S3 should block public bucket policies for buckets in this account is enabled. Returns as `true` or `false`.
        """
        return pulumi.get(self, "block_public_policy")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ignorePublicAcls")
    def ignore_public_acls(self) -> bool:
        """
        Whether or not Amazon S3 should ignore public ACLs for buckets in this account is enabled. Returns as `true` or `false`.
        """
        return pulumi.get(self, "ignore_public_acls")

    @property
    @pulumi.getter(name="restrictPublicBuckets")
    def restrict_public_buckets(self) -> bool:
        """
        Whether or not Amazon S3 should restrict public bucket policies for buckets in this account is enabled. Returns as `true` or `false`.
        """
        return pulumi.get(self, "restrict_public_buckets")


class AwaitableGetAccountPublicAccessBlockResult(GetAccountPublicAccessBlockResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountPublicAccessBlockResult(
            account_id=self.account_id,
            block_public_acls=self.block_public_acls,
            block_public_policy=self.block_public_policy,
            id=self.id,
            ignore_public_acls=self.ignore_public_acls,
            restrict_public_buckets=self.restrict_public_buckets)


def get_account_public_access_block(account_id: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccountPublicAccessBlockResult:
    """
    The S3 account public access block data source returns account-level public access block configuration.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.s3.get_account_public_access_block()
    ```


    :param str account_id: AWS account ID to configure. Defaults to automatically determined account ID of the AWS provider.
    """
    __args__ = dict()
    __args__['accountId'] = account_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws:s3/getAccountPublicAccessBlock:getAccountPublicAccessBlock', __args__, opts=opts, typ=GetAccountPublicAccessBlockResult).value

    return AwaitableGetAccountPublicAccessBlockResult(
        account_id=__ret__.account_id,
        block_public_acls=__ret__.block_public_acls,
        block_public_policy=__ret__.block_public_policy,
        id=__ret__.id,
        ignore_public_acls=__ret__.ignore_public_acls,
        restrict_public_buckets=__ret__.restrict_public_buckets)


@_utilities.lift_output_func(get_account_public_access_block)
def get_account_public_access_block_output(account_id: Optional[pulumi.Input[Optional[str]]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccountPublicAccessBlockResult]:
    """
    The S3 account public access block data source returns account-level public access block configuration.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.s3.get_account_public_access_block()
    ```


    :param str account_id: AWS account ID to configure. Defaults to automatically determined account ID of the AWS provider.
    """
    ...
