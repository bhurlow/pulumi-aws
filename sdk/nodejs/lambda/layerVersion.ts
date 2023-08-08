// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Lambda Layer Version resource. Lambda Layers allow you to reuse shared bits of code across multiple lambda functions.
 *
 * For information about Lambda Layers and how to use them, see [AWS Lambda Layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
 *
 * > **NOTE:** Setting `skipDestroy` to `true` means that the AWS Provider will _not_ destroy any layer version, even when running destroy. Layer versions are thus intentional dangling resources that are _not_ managed by the provider and may incur extra expense in your AWS account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const lambdaLayer = new aws.lambda.LayerVersion("lambdaLayer", {
 *     compatibleRuntimes: ["nodejs16.x"],
 *     code: new pulumi.asset.FileArchive("lambda_layer_payload.zip"),
 *     layerName: "lambda_layer_name",
 * });
 * ```
 * ## Specifying the Deployment Package
 *
 * AWS Lambda Layers expect source code to be provided as a deployment package whose structure varies depending on which `compatibleRuntimes` this layer specifies.
 * See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) for the valid values of `compatibleRuntimes`.
 *
 * Once you have created your deployment package you can specify it either directly as a local file (using the `filename` argument) or
 * indirectly via Amazon S3 (using the `s3Bucket`, `s3Key` and `s3ObjectVersion` arguments). When providing the deployment
 * package via S3 it may be useful to use the `aws.s3.BucketObjectv2` resource to upload it.
 *
 * For larger deployment packages it is recommended by Amazon to upload via S3, since the S3 API has better support for uploading large files efficiently.
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_lambda_layer_version.test_layer
 *
 *  id = "arn:aws:lambda:_REGION_:_ACCOUNT_ID_:layer:_LAYER_NAME_:_LAYER_VERSION_" } Using `pulumi import`, import Lambda Layers using `arn`. For exampleconsole % pulumi import \
 *
 *  aws_lambda_layer_version.test_layer \
 *
 *  arn:aws:lambda:_REGION_:_ACCOUNT_ID_:layer:_LAYER_NAME_:_LAYER_VERSION_
 */
export class LayerVersion extends pulumi.CustomResource {
    /**
     * Get an existing LayerVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LayerVersionState, opts?: pulumi.CustomResourceOptions): LayerVersion {
        return new LayerVersion(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:lambda/layerVersion:LayerVersion';

    /**
     * Returns true if the given object is an instance of LayerVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LayerVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LayerVersion.__pulumiType;
    }

    /**
     * ARN of the Lambda Layer with version.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
     */
    public readonly code!: pulumi.Output<pulumi.asset.Archive | undefined>;
    /**
     * List of [Architectures](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleArchitectures) this layer is compatible with. Currently `x8664` and `arm64` can be specified.
     */
    public readonly compatibleArchitectures!: pulumi.Output<string[] | undefined>;
    /**
     * List of [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) this layer is compatible with. Up to 5 runtimes can be specified.
     */
    public readonly compatibleRuntimes!: pulumi.Output<string[] | undefined>;
    /**
     * Date this resource was created.
     */
    public /*out*/ readonly createdDate!: pulumi.Output<string>;
    /**
     * Description of what your Lambda Layer does.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * ARN of the Lambda Layer without version.
     */
    public /*out*/ readonly layerArn!: pulumi.Output<string>;
    /**
     * Unique name for your Lambda Layer
     *
     * The following arguments are optional:
     */
    public readonly layerName!: pulumi.Output<string>;
    /**
     * License info for your Lambda Layer. See [License Info](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo).
     */
    public readonly licenseInfo!: pulumi.Output<string | undefined>;
    /**
     * S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
     */
    public readonly s3Bucket!: pulumi.Output<string | undefined>;
    /**
     * S3 key of an object containing the function's deployment package. Conflicts with `filename`.
     */
    public readonly s3Key!: pulumi.Output<string | undefined>;
    /**
     * Object version containing the function's deployment package. Conflicts with `filename`.
     */
    public readonly s3ObjectVersion!: pulumi.Output<string | undefined>;
    /**
     * ARN of a signing job.
     */
    public /*out*/ readonly signingJobArn!: pulumi.Output<string>;
    /**
     * ARN for a signing profile version.
     */
    public /*out*/ readonly signingProfileVersionArn!: pulumi.Output<string>;
    /**
     * Whether to retain the old version of a previously deployed Lambda Layer. Default is `false`. When this is not set to `true`, changing any of `compatibleArchitectures`, `compatibleRuntimes`, `description`, `filename`, `layerName`, `licenseInfo`, `s3Bucket`, `s3Key`, `s3ObjectVersion`, or `sourceCodeHash` forces deletion of the existing layer version and creation of a new layer version.
     */
    public readonly skipDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`.
     */
    public readonly sourceCodeHash!: pulumi.Output<string>;
    /**
     * Size in bytes of the function .zip file.
     */
    public /*out*/ readonly sourceCodeSize!: pulumi.Output<number>;
    /**
     * Lambda Layer version.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a LayerVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LayerVersionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LayerVersionArgs | LayerVersionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LayerVersionState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["code"] = state ? state.code : undefined;
            resourceInputs["compatibleArchitectures"] = state ? state.compatibleArchitectures : undefined;
            resourceInputs["compatibleRuntimes"] = state ? state.compatibleRuntimes : undefined;
            resourceInputs["createdDate"] = state ? state.createdDate : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["layerArn"] = state ? state.layerArn : undefined;
            resourceInputs["layerName"] = state ? state.layerName : undefined;
            resourceInputs["licenseInfo"] = state ? state.licenseInfo : undefined;
            resourceInputs["s3Bucket"] = state ? state.s3Bucket : undefined;
            resourceInputs["s3Key"] = state ? state.s3Key : undefined;
            resourceInputs["s3ObjectVersion"] = state ? state.s3ObjectVersion : undefined;
            resourceInputs["signingJobArn"] = state ? state.signingJobArn : undefined;
            resourceInputs["signingProfileVersionArn"] = state ? state.signingProfileVersionArn : undefined;
            resourceInputs["skipDestroy"] = state ? state.skipDestroy : undefined;
            resourceInputs["sourceCodeHash"] = state ? state.sourceCodeHash : undefined;
            resourceInputs["sourceCodeSize"] = state ? state.sourceCodeSize : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as LayerVersionArgs | undefined;
            if ((!args || args.layerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'layerName'");
            }
            resourceInputs["code"] = args ? args.code : undefined;
            resourceInputs["compatibleArchitectures"] = args ? args.compatibleArchitectures : undefined;
            resourceInputs["compatibleRuntimes"] = args ? args.compatibleRuntimes : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["layerName"] = args ? args.layerName : undefined;
            resourceInputs["licenseInfo"] = args ? args.licenseInfo : undefined;
            resourceInputs["s3Bucket"] = args ? args.s3Bucket : undefined;
            resourceInputs["s3Key"] = args ? args.s3Key : undefined;
            resourceInputs["s3ObjectVersion"] = args ? args.s3ObjectVersion : undefined;
            resourceInputs["skipDestroy"] = args ? args.skipDestroy : undefined;
            resourceInputs["sourceCodeHash"] = args ? args.sourceCodeHash : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdDate"] = undefined /*out*/;
            resourceInputs["layerArn"] = undefined /*out*/;
            resourceInputs["signingJobArn"] = undefined /*out*/;
            resourceInputs["signingProfileVersionArn"] = undefined /*out*/;
            resourceInputs["sourceCodeSize"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LayerVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LayerVersion resources.
 */
export interface LayerVersionState {
    /**
     * ARN of the Lambda Layer with version.
     */
    arn?: pulumi.Input<string>;
    /**
     * Path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
     */
    code?: pulumi.Input<pulumi.asset.Archive>;
    /**
     * List of [Architectures](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleArchitectures) this layer is compatible with. Currently `x8664` and `arm64` can be specified.
     */
    compatibleArchitectures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) this layer is compatible with. Up to 5 runtimes can be specified.
     */
    compatibleRuntimes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Date this resource was created.
     */
    createdDate?: pulumi.Input<string>;
    /**
     * Description of what your Lambda Layer does.
     */
    description?: pulumi.Input<string>;
    /**
     * ARN of the Lambda Layer without version.
     */
    layerArn?: pulumi.Input<string>;
    /**
     * Unique name for your Lambda Layer
     *
     * The following arguments are optional:
     */
    layerName?: pulumi.Input<string>;
    /**
     * License info for your Lambda Layer. See [License Info](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo).
     */
    licenseInfo?: pulumi.Input<string>;
    /**
     * S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
     */
    s3Bucket?: pulumi.Input<string>;
    /**
     * S3 key of an object containing the function's deployment package. Conflicts with `filename`.
     */
    s3Key?: pulumi.Input<string>;
    /**
     * Object version containing the function's deployment package. Conflicts with `filename`.
     */
    s3ObjectVersion?: pulumi.Input<string>;
    /**
     * ARN of a signing job.
     */
    signingJobArn?: pulumi.Input<string>;
    /**
     * ARN for a signing profile version.
     */
    signingProfileVersionArn?: pulumi.Input<string>;
    /**
     * Whether to retain the old version of a previously deployed Lambda Layer. Default is `false`. When this is not set to `true`, changing any of `compatibleArchitectures`, `compatibleRuntimes`, `description`, `filename`, `layerName`, `licenseInfo`, `s3Bucket`, `s3Key`, `s3ObjectVersion`, or `sourceCodeHash` forces deletion of the existing layer version and creation of a new layer version.
     */
    skipDestroy?: pulumi.Input<boolean>;
    /**
     * Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`.
     */
    sourceCodeHash?: pulumi.Input<string>;
    /**
     * Size in bytes of the function .zip file.
     */
    sourceCodeSize?: pulumi.Input<number>;
    /**
     * Lambda Layer version.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LayerVersion resource.
 */
export interface LayerVersionArgs {
    /**
     * Path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
     */
    code?: pulumi.Input<pulumi.asset.Archive>;
    /**
     * List of [Architectures](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleArchitectures) this layer is compatible with. Currently `x8664` and `arm64` can be specified.
     */
    compatibleArchitectures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) this layer is compatible with. Up to 5 runtimes can be specified.
     */
    compatibleRuntimes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description of what your Lambda Layer does.
     */
    description?: pulumi.Input<string>;
    /**
     * Unique name for your Lambda Layer
     *
     * The following arguments are optional:
     */
    layerName: pulumi.Input<string>;
    /**
     * License info for your Lambda Layer. See [License Info](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-LicenseInfo).
     */
    licenseInfo?: pulumi.Input<string>;
    /**
     * S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
     */
    s3Bucket?: pulumi.Input<string>;
    /**
     * S3 key of an object containing the function's deployment package. Conflicts with `filename`.
     */
    s3Key?: pulumi.Input<string>;
    /**
     * Object version containing the function's deployment package. Conflicts with `filename`.
     */
    s3ObjectVersion?: pulumi.Input<string>;
    /**
     * Whether to retain the old version of a previously deployed Lambda Layer. Default is `false`. When this is not set to `true`, changing any of `compatibleArchitectures`, `compatibleRuntimes`, `description`, `filename`, `layerName`, `licenseInfo`, `s3Bucket`, `s3Key`, `s3ObjectVersion`, or `sourceCodeHash` forces deletion of the existing layer version and creation of a new layer version.
     */
    skipDestroy?: pulumi.Input<boolean>;
    /**
     * Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`.
     */
    sourceCodeHash?: pulumi.Input<string>;
}
