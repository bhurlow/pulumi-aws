// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a SageMaker Image Version resource.
 *
 * ## Example Usage
 * ### Basic usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.sagemaker.ImageVersion("test", {
 *     imageName: aws_sagemaker_image.test.id,
 *     baseImage: "012345678912.dkr.ecr.us-west-2.amazonaws.com/image:latest",
 * });
 * ```
 *
 * ## Import
 *
 * SageMaker Image Versions can be imported using the `name`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:sagemaker/imageVersion:ImageVersion test_image my-code-repo
 * ```
 */
export class ImageVersion extends pulumi.CustomResource {
    /**
     * Get an existing ImageVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ImageVersionState, opts?: pulumi.CustomResourceOptions): ImageVersion {
        return new ImageVersion(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:sagemaker/imageVersion:ImageVersion';

    /**
     * Returns true if the given object is an instance of ImageVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ImageVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ImageVersion.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this Image Version.
     * * `imageArn`- The Amazon Resource Name (ARN) of the image the version is based on.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The registry path of the container image on which this image version is based.
     */
    public readonly baseImage!: pulumi.Output<string>;
    /**
     * The registry path of the container image that contains this image version.
     */
    public /*out*/ readonly containerImage!: pulumi.Output<string>;
    public /*out*/ readonly imageArn!: pulumi.Output<string>;
    /**
     * The name of the image. Must be unique to your account.
     */
    public readonly imageName!: pulumi.Output<string>;
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a ImageVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImageVersionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ImageVersionArgs | ImageVersionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ImageVersionState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["baseImage"] = state ? state.baseImage : undefined;
            resourceInputs["containerImage"] = state ? state.containerImage : undefined;
            resourceInputs["imageArn"] = state ? state.imageArn : undefined;
            resourceInputs["imageName"] = state ? state.imageName : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as ImageVersionArgs | undefined;
            if ((!args || args.baseImage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'baseImage'");
            }
            if ((!args || args.imageName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageName'");
            }
            resourceInputs["baseImage"] = args ? args.baseImage : undefined;
            resourceInputs["imageName"] = args ? args.imageName : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["containerImage"] = undefined /*out*/;
            resourceInputs["imageArn"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ImageVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ImageVersion resources.
 */
export interface ImageVersionState {
    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this Image Version.
     * * `imageArn`- The Amazon Resource Name (ARN) of the image the version is based on.
     */
    arn?: pulumi.Input<string>;
    /**
     * The registry path of the container image on which this image version is based.
     */
    baseImage?: pulumi.Input<string>;
    /**
     * The registry path of the container image that contains this image version.
     */
    containerImage?: pulumi.Input<string>;
    imageArn?: pulumi.Input<string>;
    /**
     * The name of the image. Must be unique to your account.
     */
    imageName?: pulumi.Input<string>;
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ImageVersion resource.
 */
export interface ImageVersionArgs {
    /**
     * The registry path of the container image on which this image version is based.
     */
    baseImage: pulumi.Input<string>;
    /**
     * The name of the image. Must be unique to your account.
     */
    imageName: pulumi.Input<string>;
}
