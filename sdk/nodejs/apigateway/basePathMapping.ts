// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {RestApi} from "./index";

/**
 * Connects a custom domain name registered via `aws.apigateway.DomainName`
 * with a deployed API so that its methods can be called via the
 * custom domain name.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as fs from "fs";
 *
 * const exampleStage = new aws.apigateway.Stage("exampleStage", {
 *     deployment: aws_api_gateway_deployment.example.id,
 *     restApi: aws_api_gateway_rest_api.example.id,
 *     stageName: "example",
 * });
 * const exampleDomainName = new aws.apigateway.DomainName("exampleDomainName", {
 *     domainName: "example.com",
 *     certificateName: "example-api",
 *     certificateBody: fs.readFileSync(`${path.module}/example.com/example.crt`),
 *     certificateChain: fs.readFileSync(`${path.module}/example.com/ca.crt`),
 *     certificatePrivateKey: fs.readFileSync(`${path.module}/example.com/example.key`),
 * });
 * const exampleBasePathMapping = new aws.apigateway.BasePathMapping("exampleBasePathMapping", {
 *     restApi: aws_api_gateway_rest_api.example.id,
 *     stageName: exampleStage.stageName,
 *     domainName: exampleDomainName.domainName,
 * });
 * ```
 *
 * ## Import
 *
 * For an empty `base_path` or, in other words, a root path (`/`)terraform import {
 *
 *  to = aws_api_gateway_base_path_mapping.example
 *
 *  id = "example.com/" } For a non-root `base_path`terraform import {
 *
 *  to = aws_api_gateway_base_path_mapping.example
 *
 *  id = "example.com/base-path" } Using `pulumi import`, import `aws_api_gateway_base_path_mapping` using the domain name and base path. For exampleFor an empty `base_path` or, in other words, a root path (`/`)console % pulumi import aws_api_gateway_base_path_mapping.example example.com/ For a non-root `base_path`console % pulumi import aws_api_gateway_base_path_mapping.example example.com/base-path
 */
export class BasePathMapping extends pulumi.CustomResource {
    /**
     * Get an existing BasePathMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BasePathMappingState, opts?: pulumi.CustomResourceOptions): BasePathMapping {
        return new BasePathMapping(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigateway/basePathMapping:BasePathMapping';

    /**
     * Returns true if the given object is an instance of BasePathMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BasePathMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BasePathMapping.__pulumiType;
    }

    /**
     * Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
     */
    public readonly basePath!: pulumi.Output<string | undefined>;
    /**
     * Already-registered domain name to connect the API to.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * ID of the API to connect.
     */
    public readonly restApi!: pulumi.Output<string>;
    /**
     * Name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
     */
    public readonly stageName!: pulumi.Output<string | undefined>;

    /**
     * Create a BasePathMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BasePathMappingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BasePathMappingArgs | BasePathMappingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BasePathMappingState | undefined;
            resourceInputs["basePath"] = state ? state.basePath : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["restApi"] = state ? state.restApi : undefined;
            resourceInputs["stageName"] = state ? state.stageName : undefined;
        } else {
            const args = argsOrState as BasePathMappingArgs | undefined;
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.restApi === undefined) && !opts.urn) {
                throw new Error("Missing required property 'restApi'");
            }
            resourceInputs["basePath"] = args ? args.basePath : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["restApi"] = args ? args.restApi : undefined;
            resourceInputs["stageName"] = args ? args.stageName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BasePathMapping.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BasePathMapping resources.
 */
export interface BasePathMappingState {
    /**
     * Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
     */
    basePath?: pulumi.Input<string>;
    /**
     * Already-registered domain name to connect the API to.
     */
    domainName?: pulumi.Input<string>;
    /**
     * ID of the API to connect.
     */
    restApi?: pulumi.Input<string | RestApi>;
    /**
     * Name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
     */
    stageName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BasePathMapping resource.
 */
export interface BasePathMappingArgs {
    /**
     * Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
     */
    basePath?: pulumi.Input<string>;
    /**
     * Already-registered domain name to connect the API to.
     */
    domainName: pulumi.Input<string>;
    /**
     * ID of the API to connect.
     */
    restApi: pulumi.Input<string | RestApi>;
    /**
     * Name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
     */
    stageName?: pulumi.Input<string>;
}
