// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a WAFv2 Web ACL Association.
 *
 * > **NOTE on associating a WAFv2 Web ACL with a Cloudfront distribution:** Do not use this resource to associate a WAFv2 Web ACL with a Cloudfront Distribution. The [AWS API call backing this resource](https://docs.aws.amazon.com/waf/latest/APIReference/API_AssociateWebACL.html) notes that you should use the [`webAclId`][2] property on the [`cloudfrontDistribution`][2] instead.
 *
 * [1]: https://docs.aws.amazon.com/waf/latest/APIReference/API_AssociateWebACL.html
 * [2]: https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#web_acl_id
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as crypto from "crypto";
 *
 * const exampleRestApi = new aws.apigateway.RestApi("exampleRestApi", {body: JSON.stringify({
 *     openapi: "3.0.1",
 *     info: {
 *         title: "example",
 *         version: "1.0",
 *     },
 *     paths: {
 *         "/path1": {
 *             get: {
 *                 "x-amazon-apigateway-integration": {
 *                     httpMethod: "GET",
 *                     payloadFormatVersion: "1.0",
 *                     type: "HTTP_PROXY",
 *                     uri: "https://ip-ranges.amazonaws.com/ip-ranges.json",
 *                 },
 *             },
 *         },
 *     },
 * })});
 * const exampleDeployment = new aws.apigateway.Deployment("exampleDeployment", {
 *     restApi: exampleRestApi.id,
 *     triggers: {
 *         redeployment: exampleRestApi.body.apply(body => JSON.stringify(body)).apply(toJSON => crypto.createHash('sha1').update(toJSON).digest('hex')),
 *     },
 * });
 * const exampleStage = new aws.apigateway.Stage("exampleStage", {
 *     deployment: exampleDeployment.id,
 *     restApi: exampleRestApi.id,
 *     stageName: "example",
 * });
 * const exampleWebAcl = new aws.wafv2.WebAcl("exampleWebAcl", {
 *     scope: "REGIONAL",
 *     defaultAction: {
 *         allow: {},
 *     },
 *     visibilityConfig: {
 *         cloudwatchMetricsEnabled: false,
 *         metricName: "friendly-metric-name",
 *         sampledRequestsEnabled: false,
 *     },
 * });
 * const exampleWebAclAssociation = new aws.wafv2.WebAclAssociation("exampleWebAclAssociation", {
 *     resourceArn: exampleStage.arn,
 *     webAclArn: exampleWebAcl.arn,
 * });
 * ```
 *
 * ## Import
 *
 * WAFv2 Web ACL Association can be imported using `WEB_ACL_ARN,RESOURCE_ARN` e.g.
 *
 * ```sh
 *  $ pulumi import aws:wafv2/webAclAssociation:WebAclAssociation example arn:aws:wafv2:...7ce849ea,arn:aws:apigateway:...ages/name
 * ```
 */
export class WebAclAssociation extends pulumi.CustomResource {
    /**
     * Get an existing WebAclAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebAclAssociationState, opts?: pulumi.CustomResourceOptions): WebAclAssociation {
        return new WebAclAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:wafv2/webAclAssociation:WebAclAssociation';

    /**
     * Returns true if the given object is an instance of WebAclAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebAclAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebAclAssociation.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer or an Amazon API Gateway stage.
     */
    public readonly resourceArn!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
     */
    public readonly webAclArn!: pulumi.Output<string>;

    /**
     * Create a WebAclAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAclAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebAclAssociationArgs | WebAclAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebAclAssociationState | undefined;
            inputs["resourceArn"] = state ? state.resourceArn : undefined;
            inputs["webAclArn"] = state ? state.webAclArn : undefined;
        } else {
            const args = argsOrState as WebAclAssociationArgs | undefined;
            if ((!args || args.resourceArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceArn'");
            }
            if ((!args || args.webAclArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'webAclArn'");
            }
            inputs["resourceArn"] = args ? args.resourceArn : undefined;
            inputs["webAclArn"] = args ? args.webAclArn : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WebAclAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebAclAssociation resources.
 */
export interface WebAclAssociationState {
    /**
     * The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer or an Amazon API Gateway stage.
     */
    resourceArn?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
     */
    webAclArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebAclAssociation resource.
 */
export interface WebAclAssociationArgs {
    /**
     * The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer or an Amazon API Gateway stage.
     */
    resourceArn: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
     */
    webAclArn: pulumi.Input<string>;
}
