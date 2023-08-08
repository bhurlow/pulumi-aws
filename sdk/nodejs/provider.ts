// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

import {Region} from "./index";

/**
 * The provider type for the aws package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'aws';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    /**
     * The access key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
     */
    public readonly accessKey!: pulumi.Output<string | undefined>;
    /**
     * File containing custom root and intermediate certificates. Can also be configured using the `AWS_CA_BUNDLE` environment
     * variable. (Setting `ca_bundle` in the shared config file is not supported.)
     */
    public readonly customCaBundle!: pulumi.Output<string | undefined>;
    /**
     * Address of the EC2 metadata service endpoint to use. Can also be configured using the
     * `AWS_EC2_METADATA_SERVICE_ENDPOINT` environment variable.
     */
    public readonly ec2MetadataServiceEndpoint!: pulumi.Output<string | undefined>;
    /**
     * Protocol to use with EC2 metadata service endpoint.Valid values are `IPv4` and `IPv6`. Can also be configured using the
     * `AWS_EC2_METADATA_SERVICE_ENDPOINT_MODE` environment variable.
     */
    public readonly ec2MetadataServiceEndpointMode!: pulumi.Output<string | undefined>;
    /**
     * The address of an HTTP proxy to use when accessing the AWS API. Can also be configured using the `HTTP_PROXY` or
     * `HTTPS_PROXY` environment variables.
     */
    public readonly httpProxy!: pulumi.Output<string | undefined>;
    /**
     * The profile for API operations. If not set, the default profile created with `aws configure` will be used.
     */
    public readonly profile!: pulumi.Output<string | undefined>;
    /**
     * The region where AWS operations will take place. Examples are us-east-1, us-west-2, etc.
     */
    public readonly region!: pulumi.Output<Region | undefined>;
    /**
     * Specifies how retries are attempted. Valid values are `standard` and `adaptive`. Can also be configured using the
     * `AWS_RETRY_MODE` environment variable.
     */
    public readonly retryMode!: pulumi.Output<string | undefined>;
    /**
     * The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
     */
    public readonly secretKey!: pulumi.Output<string | undefined>;
    /**
     * The region where AWS STS operations will take place. Examples are us-east-1 and us-west-2.
     */
    public readonly stsRegion!: pulumi.Output<string | undefined>;
    /**
     * session token. A session token is only required if you are using temporary security credentials.
     */
    public readonly token!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["accessKey"] = args ? args.accessKey : undefined;
            resourceInputs["allowedAccountIds"] = pulumi.output(args ? args.allowedAccountIds : undefined).apply(JSON.stringify);
            resourceInputs["assumeRole"] = pulumi.output(args ? args.assumeRole : undefined).apply(JSON.stringify);
            resourceInputs["assumeRoleWithWebIdentity"] = pulumi.output(args ? args.assumeRoleWithWebIdentity : undefined).apply(JSON.stringify);
            resourceInputs["customCaBundle"] = args ? args.customCaBundle : undefined;
            resourceInputs["defaultTags"] = pulumi.output(args ? args.defaultTags : undefined).apply(JSON.stringify);
            resourceInputs["ec2MetadataServiceEndpoint"] = args ? args.ec2MetadataServiceEndpoint : undefined;
            resourceInputs["ec2MetadataServiceEndpointMode"] = args ? args.ec2MetadataServiceEndpointMode : undefined;
            resourceInputs["endpoints"] = pulumi.output(args ? args.endpoints : undefined).apply(JSON.stringify);
            resourceInputs["forbiddenAccountIds"] = pulumi.output(args ? args.forbiddenAccountIds : undefined).apply(JSON.stringify);
            resourceInputs["httpProxy"] = args ? args.httpProxy : undefined;
            resourceInputs["ignoreTags"] = pulumi.output(args ? args.ignoreTags : undefined).apply(JSON.stringify);
            resourceInputs["insecure"] = pulumi.output(args ? args.insecure : undefined).apply(JSON.stringify);
            resourceInputs["maxRetries"] = pulumi.output(args ? args.maxRetries : undefined).apply(JSON.stringify);
            resourceInputs["profile"] = args ? args.profile : undefined;
            resourceInputs["region"] = (args ? args.region : undefined) ?? <any>utilities.getEnv("AWS_REGION", "AWS_DEFAULT_REGION");
            resourceInputs["retryMode"] = args ? args.retryMode : undefined;
            resourceInputs["s3UsePathStyle"] = pulumi.output(args ? args.s3UsePathStyle : undefined).apply(JSON.stringify);
            resourceInputs["secretKey"] = args ? args.secretKey : undefined;
            resourceInputs["sharedConfigFiles"] = pulumi.output(args ? args.sharedConfigFiles : undefined).apply(JSON.stringify);
            resourceInputs["sharedCredentialsFiles"] = pulumi.output(args ? args.sharedCredentialsFiles : undefined).apply(JSON.stringify);
            resourceInputs["skipCredentialsValidation"] = pulumi.output((args ? args.skipCredentialsValidation : undefined) ?? false).apply(JSON.stringify);
            resourceInputs["skipMetadataApiCheck"] = pulumi.output((args ? args.skipMetadataApiCheck : undefined) ?? true).apply(JSON.stringify);
            resourceInputs["skipRegionValidation"] = pulumi.output((args ? args.skipRegionValidation : undefined) ?? true).apply(JSON.stringify);
            resourceInputs["skipRequestingAccountId"] = pulumi.output(args ? args.skipRequestingAccountId : undefined).apply(JSON.stringify);
            resourceInputs["stsRegion"] = args ? args.stsRegion : undefined;
            resourceInputs["token"] = args ? args.token : undefined;
            resourceInputs["useDualstackEndpoint"] = pulumi.output(args ? args.useDualstackEndpoint : undefined).apply(JSON.stringify);
            resourceInputs["useFipsEndpoint"] = pulumi.output(args ? args.useFipsEndpoint : undefined).apply(JSON.stringify);
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * The access key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
     */
    accessKey?: pulumi.Input<string>;
    allowedAccountIds?: pulumi.Input<pulumi.Input<string>[]>;
    assumeRole?: pulumi.Input<inputs.ProviderAssumeRole>;
    assumeRoleWithWebIdentity?: pulumi.Input<inputs.ProviderAssumeRoleWithWebIdentity>;
    /**
     * File containing custom root and intermediate certificates. Can also be configured using the `AWS_CA_BUNDLE` environment
     * variable. (Setting `ca_bundle` in the shared config file is not supported.)
     */
    customCaBundle?: pulumi.Input<string>;
    /**
     * Configuration block with settings to default resource tags across all resources.
     */
    defaultTags?: pulumi.Input<inputs.ProviderDefaultTags>;
    /**
     * Address of the EC2 metadata service endpoint to use. Can also be configured using the
     * `AWS_EC2_METADATA_SERVICE_ENDPOINT` environment variable.
     */
    ec2MetadataServiceEndpoint?: pulumi.Input<string>;
    /**
     * Protocol to use with EC2 metadata service endpoint.Valid values are `IPv4` and `IPv6`. Can also be configured using the
     * `AWS_EC2_METADATA_SERVICE_ENDPOINT_MODE` environment variable.
     */
    ec2MetadataServiceEndpointMode?: pulumi.Input<string>;
    endpoints?: pulumi.Input<pulumi.Input<inputs.ProviderEndpoint>[]>;
    forbiddenAccountIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The address of an HTTP proxy to use when accessing the AWS API. Can also be configured using the `HTTP_PROXY` or
     * `HTTPS_PROXY` environment variables.
     */
    httpProxy?: pulumi.Input<string>;
    /**
     * Configuration block with settings to ignore resource tags across all resources.
     */
    ignoreTags?: pulumi.Input<inputs.ProviderIgnoreTags>;
    /**
     * Explicitly allow the provider to perform "insecure" SSL requests. If omitted, default value is `false`
     */
    insecure?: pulumi.Input<boolean>;
    /**
     * The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown.
     */
    maxRetries?: pulumi.Input<number>;
    /**
     * The profile for API operations. If not set, the default profile created with `aws configure` will be used.
     */
    profile?: pulumi.Input<string>;
    /**
     * The region where AWS operations will take place. Examples are us-east-1, us-west-2, etc.
     */
    region?: pulumi.Input<Region>;
    /**
     * Specifies how retries are attempted. Valid values are `standard` and `adaptive`. Can also be configured using the
     * `AWS_RETRY_MODE` environment variable.
     */
    retryMode?: pulumi.Input<string>;
    /**
     * Set this to true to enable the request to use path-style addressing, i.e., https://s3.amazonaws.com/BUCKET/KEY. By
     * default, the S3 client will use virtual hosted bucket addressing when possible (https://BUCKET.s3.amazonaws.com/KEY).
     * Specific to the Amazon S3 service.
     */
    s3UsePathStyle?: pulumi.Input<boolean>;
    /**
     * The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * List of paths to shared config files. If not set, defaults to [~/.aws/config].
     */
    sharedConfigFiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of paths to shared credentials files. If not set, defaults to [~/.aws/credentials].
     */
    sharedCredentialsFiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS
     * available/implemented.
     */
    skipCredentialsValidation?: pulumi.Input<boolean>;
    /**
     * Skip the AWS Metadata API check. Used for AWS API implementations that do not have a metadata api endpoint.
     */
    skipMetadataApiCheck?: pulumi.Input<boolean>;
    /**
     * Skip static validation of region name. Used by users of alternative AWS-like APIs or users w/ access to regions that are
     * not public (yet).
     */
    skipRegionValidation?: pulumi.Input<boolean>;
    /**
     * Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.
     */
    skipRequestingAccountId?: pulumi.Input<boolean>;
    /**
     * The region where AWS STS operations will take place. Examples are us-east-1 and us-west-2.
     */
    stsRegion?: pulumi.Input<string>;
    /**
     * session token. A session token is only required if you are using temporary security credentials.
     */
    token?: pulumi.Input<string>;
    /**
     * Resolve an endpoint with DualStack capability
     */
    useDualstackEndpoint?: pulumi.Input<boolean>;
    /**
     * Resolve an endpoint with FIPS capability
     */
    useFipsEndpoint?: pulumi.Input<boolean>;
}
