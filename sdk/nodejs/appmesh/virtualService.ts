// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides an AWS App Mesh virtual service resource.
 *
 * ## Example Usage
 * ### Virtual Node Provider
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const servicea = new aws.appmesh.VirtualService("servicea", {
 *     meshName: aws_appmesh_mesh.simple.id,
 *     spec: {
 *         provider: {
 *             virtualNode: {
 *                 virtualNodeName: aws_appmesh_virtual_node.serviceb1.name,
 *             },
 *         },
 *     },
 * });
 * ```
 * ### Virtual Router Provider
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const servicea = new aws.appmesh.VirtualService("servicea", {
 *     meshName: aws_appmesh_mesh.simple.id,
 *     spec: {
 *         provider: {
 *             virtualRouter: {
 *                 virtualRouterName: aws_appmesh_virtual_router.serviceb.name,
 *             },
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import App Mesh virtual services using `mesh_name` together with the virtual service's `name`. For example:
 *
 * ```sh
 *  $ pulumi import aws:appmesh/virtualService:VirtualService servicea simpleapp/servicea.simpleapp.local
 * ```
 */
export class VirtualService extends pulumi.CustomResource {
    /**
     * Get an existing VirtualService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualServiceState, opts?: pulumi.CustomResourceOptions): VirtualService {
        return new VirtualService(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:appmesh/virtualService:VirtualService';

    /**
     * Returns true if the given object is an instance of VirtualService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualService.__pulumiType;
    }

    /**
     * ARN of the virtual service.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Creation date of the virtual service.
     */
    public /*out*/ readonly createdDate!: pulumi.Output<string>;
    /**
     * Last update date of the virtual service.
     */
    public /*out*/ readonly lastUpdatedDate!: pulumi.Output<string>;
    /**
     * Name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
     */
    public readonly meshName!: pulumi.Output<string>;
    /**
     * AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
     */
    public readonly meshOwner!: pulumi.Output<string>;
    /**
     * Name to use for the virtual service. Must be between 1 and 255 characters in length.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Resource owner's AWS account ID.
     */
    public /*out*/ readonly resourceOwner!: pulumi.Output<string>;
    /**
     * Virtual service specification to apply.
     */
    public readonly spec!: pulumi.Output<outputs.appmesh.VirtualServiceSpec>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a VirtualService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualServiceArgs | VirtualServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VirtualServiceState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["createdDate"] = state ? state.createdDate : undefined;
            resourceInputs["lastUpdatedDate"] = state ? state.lastUpdatedDate : undefined;
            resourceInputs["meshName"] = state ? state.meshName : undefined;
            resourceInputs["meshOwner"] = state ? state.meshOwner : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["resourceOwner"] = state ? state.resourceOwner : undefined;
            resourceInputs["spec"] = state ? state.spec : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as VirtualServiceArgs | undefined;
            if ((!args || args.meshName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'meshName'");
            }
            if ((!args || args.spec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'spec'");
            }
            resourceInputs["meshName"] = args ? args.meshName : undefined;
            resourceInputs["meshOwner"] = args ? args.meshOwner : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["spec"] = args ? args.spec : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdDate"] = undefined /*out*/;
            resourceInputs["lastUpdatedDate"] = undefined /*out*/;
            resourceInputs["resourceOwner"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(VirtualService.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualService resources.
 */
export interface VirtualServiceState {
    /**
     * ARN of the virtual service.
     */
    arn?: pulumi.Input<string>;
    /**
     * Creation date of the virtual service.
     */
    createdDate?: pulumi.Input<string>;
    /**
     * Last update date of the virtual service.
     */
    lastUpdatedDate?: pulumi.Input<string>;
    /**
     * Name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
     */
    meshName?: pulumi.Input<string>;
    /**
     * AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
     */
    meshOwner?: pulumi.Input<string>;
    /**
     * Name to use for the virtual service. Must be between 1 and 255 characters in length.
     */
    name?: pulumi.Input<string>;
    /**
     * Resource owner's AWS account ID.
     */
    resourceOwner?: pulumi.Input<string>;
    /**
     * Virtual service specification to apply.
     */
    spec?: pulumi.Input<inputs.appmesh.VirtualServiceSpec>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a VirtualService resource.
 */
export interface VirtualServiceArgs {
    /**
     * Name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
     */
    meshName: pulumi.Input<string>;
    /**
     * AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
     */
    meshOwner?: pulumi.Input<string>;
    /**
     * Name to use for the virtual service. Must be between 1 and 255 characters in length.
     */
    name?: pulumi.Input<string>;
    /**
     * Virtual service specification to apply.
     */
    spec: pulumi.Input<inputs.appmesh.VirtualServiceSpec>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
