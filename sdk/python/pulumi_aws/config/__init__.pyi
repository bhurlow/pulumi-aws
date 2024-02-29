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

accessKey: Optional[str]
"""
The access key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
"""

allowedAccountIds: Optional[str]

assumeRole: Optional[str]

assumeRoleWithWebIdentity: Optional[str]

customCaBundle: Optional[str]
"""
File containing custom root and intermediate certificates. Can also be configured using the `AWS_CA_BUNDLE` environment
variable. (Setting `ca_bundle` in the shared config file is not supported.)
"""

defaultTags: Optional[str]
"""
Configuration block with settings to default resource tags across all resources.
"""

ec2MetadataServiceEndpoint: Optional[str]
"""
Address of the EC2 metadata service endpoint to use. Can also be configured using the
`AWS_EC2_METADATA_SERVICE_ENDPOINT` environment variable.
"""

ec2MetadataServiceEndpointMode: Optional[str]
"""
Protocol to use with EC2 metadata service endpoint.Valid values are `IPv4` and `IPv6`. Can also be configured using the
`AWS_EC2_METADATA_SERVICE_ENDPOINT_MODE` environment variable.
"""

endpoints: Optional[str]

forbiddenAccountIds: Optional[str]

httpProxy: Optional[str]
"""
URL of a proxy to use for HTTP requests when accessing the AWS API. Can also be set using the `HTTP_PROXY` or
`http_proxy` environment variables.
"""

httpsProxy: Optional[str]
"""
URL of a proxy to use for HTTPS requests when accessing the AWS API. Can also be set using the `HTTPS_PROXY` or
`https_proxy` environment variables.
"""

ignoreTags: Optional[str]
"""
Configuration block with settings to ignore resource tags across all resources.
"""

insecure: Optional[bool]
"""
Explicitly allow the provider to perform "insecure" SSL requests. If omitted, default value is `false`
"""

maxRetries: Optional[int]
"""
The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown.
"""

noProxy: Optional[str]
"""
Comma-separated list of hosts that should not use HTTP or HTTPS proxies. Can also be set using the `NO_PROXY` or
`no_proxy` environment variables.
"""

profile: Optional[str]
"""
The profile for API operations. If not set, the default profile created with `aws configure` will be used.
"""

region: Optional[str]
"""
The region where AWS operations will take place. Examples are us-east-1, us-west-2, etc.
"""

retryMode: Optional[str]
"""
Specifies how retries are attempted. Valid values are `standard` and `adaptive`. Can also be configured using the
`AWS_RETRY_MODE` environment variable.
"""

s3UsEast1RegionalEndpoint: Optional[str]
"""
Specifies whether S3 API calls in the `us-east-1` region use the legacy global endpoint or a regional endpoint. Valid
values are `legacy` or `regional`. Can also be configured using the `AWS_S3_US_EAST_1_REGIONAL_ENDPOINT` environment
variable or the `s3_us_east_1_regional_endpoint` shared config file parameter
"""

s3UsePathStyle: Optional[bool]
"""
Set this to true to enable the request to use path-style addressing, i.e., https://s3.amazonaws.com/BUCKET/KEY. By
default, the S3 client will use virtual hosted bucket addressing when possible (https://BUCKET.s3.amazonaws.com/KEY).
Specific to the Amazon S3 service.
"""

secretKey: Optional[str]
"""
The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
"""

sharedConfigFiles: Optional[str]
"""
List of paths to shared config files. If not set, defaults to [~/.aws/config].
"""

sharedCredentialsFiles: Optional[str]
"""
List of paths to shared credentials files. If not set, defaults to [~/.aws/credentials].
"""

skipCredentialsValidation: bool
"""
Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS
available/implemented.
"""

skipMetadataApiCheck: bool
"""
Skip the AWS Metadata API check. Used for AWS API implementations that do not have a metadata api endpoint.
"""

skipRegionValidation: bool
"""
Skip static validation of region name. Used by users of alternative AWS-like APIs or users w/ access to regions that are
not public (yet).
"""

skipRequestingAccountId: Optional[bool]
"""
Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.
"""

stsRegion: Optional[str]
"""
The region where AWS STS operations will take place. Examples are us-east-1 and us-west-2.
"""

token: Optional[str]
"""
session token. A session token is only required if you are using temporary security credentials.
"""

tokenBucketRateLimiterCapacity: Optional[int]
"""
The capacity of the AWS SDK's token bucket rate limiter.
"""

useDualstackEndpoint: Optional[bool]
"""
Resolve an endpoint with DualStack capability
"""

useFipsEndpoint: Optional[bool]
"""
Resolve an endpoint with FIPS capability
"""

