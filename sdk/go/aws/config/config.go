// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

// The access key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
func GetAccessKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:accessKey")
}
func GetAllowedAccountIds(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:allowedAccountIds")
}
func GetAssumeRole(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:assumeRole")
}

// Configuration block with settings to default resource tags across all resources.
func GetDefaultTags(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:defaultTags")
}
func GetEndpoints(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:endpoints")
}
func GetForbiddenAccountIds(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:forbiddenAccountIds")
}

// Configuration block with settings to ignore resource tags across all resources.
func GetIgnoreTags(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:ignoreTags")
}

// Explicitly allow the provider to perform "insecure" SSL requests. If omitted,default value is `false`
func GetInsecure(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "aws:insecure")
}

// The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown.
func GetMaxRetries(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "aws:maxRetries")
}

// The profile for API operations. If not set, the default profile created with `aws configure` will be used.
func GetProfile(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "aws:profile")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "AWS_PROFILE").(string)
}

// The region where AWS operations will take place. Examples are us-east-1, us-west-2, etc.
func GetRegion(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "aws:region")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "AWS_REGION", "AWS_DEFAULT_REGION").(string)
}

// Set this to true to force the request to use path-style addressing, i.e., http://s3.amazonaws.com/BUCKET/KEY. By
// default, the S3 client will use virtual hosted bucket addressing when possible (http://BUCKET.s3.amazonaws.com/KEY).
// Specific to the Amazon S3 service.
func GetS3ForcePathStyle(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "aws:s3ForcePathStyle")
}

// The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
func GetSecretKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:secretKey")
}

// The path to the shared credentials file. If not set this defaults to ~/.aws/credentials.
func GetSharedCredentialsFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:sharedCredentialsFile")
}

// Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS
// available/implemented.
func GetSkipCredentialsValidation(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "aws:skipCredentialsValidation")
	if err == nil {
		return v
	}
	return true
}

// Skip getting the supported EC2 platforms. Used by users that don't have ec2:DescribeAccountAttributes permissions.
func GetSkipGetEc2Platforms(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "aws:skipGetEc2Platforms")
	if err == nil {
		return v
	}
	return true
}
func GetSkipMetadataApiCheck(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "aws:skipMetadataApiCheck")
	if err == nil {
		return v
	}
	return true
}

// Skip static validation of region name. Used by users of alternative AWS-like APIs or users w/ access to regions that are
// not public (yet).
func GetSkipRegionValidation(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "aws:skipRegionValidation")
	if err == nil {
		return v
	}
	return true
}

// Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.
func GetSkipRequestingAccountId(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "aws:skipRequestingAccountId")
}

// session token. A session token is only required if you are using temporary security credentials.
func GetToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:token")
}
