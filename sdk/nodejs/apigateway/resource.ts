// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {RestApi} from "./index";

/**
 * Provides an API Gateway Resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const myDemoAPI = new aws.apigateway.RestApi("myDemoAPI", {description: "This is my API for demonstration purposes"});
 * const myDemoResource = new aws.apigateway.Resource("myDemoResource", {
 *     restApi: myDemoAPI.id,
 *     parentId: myDemoAPI.rootResourceId,
 *     pathPart: "mydemoresource",
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_api_gateway_resource.example
 *
 *  id = "12345abcde/67890fghij" } Using `pulumi import`, import `aws_api_gateway_resource` using `REST-API-ID/RESOURCE-ID`. For exampleconsole % pulumi import aws_api_gateway_resource.example 12345abcde/67890fghij
 */
export class Resource extends pulumi.CustomResource {
    /**
     * Get an existing Resource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourceState, opts?: pulumi.CustomResourceOptions): Resource {
        return new Resource(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigateway/resource:Resource';

    /**
     * Returns true if the given object is an instance of Resource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Resource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Resource.__pulumiType;
    }

    /**
     * ID of the parent API resource
     */
    public readonly parentId!: pulumi.Output<string>;
    /**
     * Complete path for this API resource, including all parent paths.
     */
    public /*out*/ readonly path!: pulumi.Output<string>;
    /**
     * Last path segment of this API resource.
     */
    public readonly pathPart!: pulumi.Output<string>;
    /**
     * ID of the associated REST API
     */
    public readonly restApi!: pulumi.Output<string>;

    /**
     * Create a Resource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourceArgs | ResourceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResourceState | undefined;
            resourceInputs["parentId"] = state ? state.parentId : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["pathPart"] = state ? state.pathPart : undefined;
            resourceInputs["restApi"] = state ? state.restApi : undefined;
        } else {
            const args = argsOrState as ResourceArgs | undefined;
            if ((!args || args.parentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parentId'");
            }
            if ((!args || args.pathPart === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pathPart'");
            }
            if ((!args || args.restApi === undefined) && !opts.urn) {
                throw new Error("Missing required property 'restApi'");
            }
            resourceInputs["parentId"] = args ? args.parentId : undefined;
            resourceInputs["pathPart"] = args ? args.pathPart : undefined;
            resourceInputs["restApi"] = args ? args.restApi : undefined;
            resourceInputs["path"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Resource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Resource resources.
 */
export interface ResourceState {
    /**
     * ID of the parent API resource
     */
    parentId?: pulumi.Input<string>;
    /**
     * Complete path for this API resource, including all parent paths.
     */
    path?: pulumi.Input<string>;
    /**
     * Last path segment of this API resource.
     */
    pathPart?: pulumi.Input<string>;
    /**
     * ID of the associated REST API
     */
    restApi?: pulumi.Input<string | RestApi>;
}

/**
 * The set of arguments for constructing a Resource resource.
 */
export interface ResourceArgs {
    /**
     * ID of the parent API resource
     */
    parentId: pulumi.Input<string>;
    /**
     * Last path segment of this API resource.
     */
    pathPart: pulumi.Input<string>;
    /**
     * ID of the associated REST API
     */
    restApi: pulumi.Input<string | RestApi>;
}
