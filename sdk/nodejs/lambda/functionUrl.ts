// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a Lambda function URL resource. A function URL is a dedicated HTTP(S) endpoint for a Lambda function.
 *
 * See the [AWS Lambda documentation](https://docs.aws.amazon.com/lambda/latest/dg/lambda-urls.html) for more information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testLatest = new aws.lambda.FunctionUrl("testLatest", {
 *     functionName: aws_lambda_function.test.function_name,
 *     authorizationType: "NONE",
 * });
 * const testLive = new aws.lambda.FunctionUrl("testLive", {
 *     functionName: aws_lambda_function.test.function_name,
 *     qualifier: "my_alias",
 *     authorizationType: "AWS_IAM",
 *     cors: {
 *         allowCredentials: true,
 *         allowOrigins: ["*"],
 *         allowMethods: ["*"],
 *         allowHeaders: [
 *             "date",
 *             "keep-alive",
 *         ],
 *         exposeHeaders: [
 *             "keep-alive",
 *             "date",
 *         ],
 *         maxAge: 86400,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_lambda_function_url.test_lambda_url
 *
 *  id = "my_test_lambda_function" } Using `pulumi import`, import Lambda function URLs using the `function_name` or `function_name/qualifier`. For exampleconsole % pulumi import aws_lambda_function_url.test_lambda_url my_test_lambda_function
 */
export class FunctionUrl extends pulumi.CustomResource {
    /**
     * Get an existing FunctionUrl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionUrlState, opts?: pulumi.CustomResourceOptions): FunctionUrl {
        return new FunctionUrl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:lambda/functionUrl:FunctionUrl';

    /**
     * Returns true if the given object is an instance of FunctionUrl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FunctionUrl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FunctionUrl.__pulumiType;
    }

    /**
     * The type of authentication that the function URL uses. Set to `"AWS_IAM"` to restrict access to authenticated IAM users only. Set to `"NONE"` to bypass IAM authentication and create a public endpoint. See the [AWS documentation](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html) for more details.
     */
    public readonly authorizationType!: pulumi.Output<string>;
    /**
     * The [cross-origin resource sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) settings for the function URL. Documented below.
     */
    public readonly cors!: pulumi.Output<outputs.lambda.FunctionUrlCors | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the function.
     */
    public /*out*/ readonly functionArn!: pulumi.Output<string>;
    /**
     * The name (or ARN) of the Lambda function.
     */
    public readonly functionName!: pulumi.Output<string>;
    /**
     * The HTTP URL endpoint for the function in the format `https://<url_id>.lambda-url.<region>.on.aws`.
     */
    public /*out*/ readonly functionUrl!: pulumi.Output<string>;
    /**
     * Determines how the Lambda function responds to an invocation. Valid values are `BUFFERED` (default) and `RESPONSE_STREAM`. See more in [Configuring a Lambda function to stream responses](https://docs.aws.amazon.com/lambda/latest/dg/configuration-response-streaming.html).
     */
    public readonly invokeMode!: pulumi.Output<string | undefined>;
    /**
     * The alias name or `"$LATEST"`.
     */
    public readonly qualifier!: pulumi.Output<string | undefined>;
    /**
     * A generated ID for the endpoint.
     */
    public /*out*/ readonly urlId!: pulumi.Output<string>;

    /**
     * Create a FunctionUrl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionUrlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionUrlArgs | FunctionUrlState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionUrlState | undefined;
            resourceInputs["authorizationType"] = state ? state.authorizationType : undefined;
            resourceInputs["cors"] = state ? state.cors : undefined;
            resourceInputs["functionArn"] = state ? state.functionArn : undefined;
            resourceInputs["functionName"] = state ? state.functionName : undefined;
            resourceInputs["functionUrl"] = state ? state.functionUrl : undefined;
            resourceInputs["invokeMode"] = state ? state.invokeMode : undefined;
            resourceInputs["qualifier"] = state ? state.qualifier : undefined;
            resourceInputs["urlId"] = state ? state.urlId : undefined;
        } else {
            const args = argsOrState as FunctionUrlArgs | undefined;
            if ((!args || args.authorizationType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authorizationType'");
            }
            if ((!args || args.functionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionName'");
            }
            resourceInputs["authorizationType"] = args ? args.authorizationType : undefined;
            resourceInputs["cors"] = args ? args.cors : undefined;
            resourceInputs["functionName"] = args ? args.functionName : undefined;
            resourceInputs["invokeMode"] = args ? args.invokeMode : undefined;
            resourceInputs["qualifier"] = args ? args.qualifier : undefined;
            resourceInputs["functionArn"] = undefined /*out*/;
            resourceInputs["functionUrl"] = undefined /*out*/;
            resourceInputs["urlId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FunctionUrl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FunctionUrl resources.
 */
export interface FunctionUrlState {
    /**
     * The type of authentication that the function URL uses. Set to `"AWS_IAM"` to restrict access to authenticated IAM users only. Set to `"NONE"` to bypass IAM authentication and create a public endpoint. See the [AWS documentation](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html) for more details.
     */
    authorizationType?: pulumi.Input<string>;
    /**
     * The [cross-origin resource sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) settings for the function URL. Documented below.
     */
    cors?: pulumi.Input<inputs.lambda.FunctionUrlCors>;
    /**
     * The Amazon Resource Name (ARN) of the function.
     */
    functionArn?: pulumi.Input<string>;
    /**
     * The name (or ARN) of the Lambda function.
     */
    functionName?: pulumi.Input<string>;
    /**
     * The HTTP URL endpoint for the function in the format `https://<url_id>.lambda-url.<region>.on.aws`.
     */
    functionUrl?: pulumi.Input<string>;
    /**
     * Determines how the Lambda function responds to an invocation. Valid values are `BUFFERED` (default) and `RESPONSE_STREAM`. See more in [Configuring a Lambda function to stream responses](https://docs.aws.amazon.com/lambda/latest/dg/configuration-response-streaming.html).
     */
    invokeMode?: pulumi.Input<string>;
    /**
     * The alias name or `"$LATEST"`.
     */
    qualifier?: pulumi.Input<string>;
    /**
     * A generated ID for the endpoint.
     */
    urlId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FunctionUrl resource.
 */
export interface FunctionUrlArgs {
    /**
     * The type of authentication that the function URL uses. Set to `"AWS_IAM"` to restrict access to authenticated IAM users only. Set to `"NONE"` to bypass IAM authentication and create a public endpoint. See the [AWS documentation](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html) for more details.
     */
    authorizationType: pulumi.Input<string>;
    /**
     * The [cross-origin resource sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) settings for the function URL. Documented below.
     */
    cors?: pulumi.Input<inputs.lambda.FunctionUrlCors>;
    /**
     * The name (or ARN) of the Lambda function.
     */
    functionName: pulumi.Input<string>;
    /**
     * Determines how the Lambda function responds to an invocation. Valid values are `BUFFERED` (default) and `RESPONSE_STREAM`. See more in [Configuring a Lambda function to stream responses](https://docs.aws.amazon.com/lambda/latest/dg/configuration-response-streaming.html).
     */
    invokeMode?: pulumi.Input<string>;
    /**
     * The alias name or `"$LATEST"`.
     */
    qualifier?: pulumi.Input<string>;
}
