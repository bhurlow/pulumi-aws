// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a Location Service Map.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.location.Map("example", {
 *     configuration: {
 *         style: "VectorHereBerlin",
 *     },
 *     mapName: "example",
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import `aws_location_map` resources using the map name. For example:
 *
 * ```sh
 *  $ pulumi import aws:location/map:Map example example
 * ```
 */
export class Map extends pulumi.CustomResource {
    /**
     * Get an existing Map resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MapState, opts?: pulumi.CustomResourceOptions): Map {
        return new Map(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:location/map:Map';

    /**
     * Returns true if the given object is an instance of Map.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Map {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Map.__pulumiType;
    }

    /**
     * Configuration block with the map style selected from an available data provider. Detailed below.
     */
    public readonly configuration!: pulumi.Output<outputs.location.MapConfiguration>;
    /**
     * The timestamp for when the map resource was created in ISO 8601 format.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * An optional description for the map resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) for the map resource. Used to specify a resource across all AWS.
     */
    public /*out*/ readonly mapArn!: pulumi.Output<string>;
    /**
     * The name for the map resource.
     *
     * The following arguments are optional:
     */
    public readonly mapName!: pulumi.Output<string>;
    /**
     * Key-value tags for the map. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The timestamp for when the map resource was last updated in ISO 8601 format.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Map resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MapArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MapArgs | MapState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MapState | undefined;
            resourceInputs["configuration"] = state ? state.configuration : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["mapArn"] = state ? state.mapArn : undefined;
            resourceInputs["mapName"] = state ? state.mapName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as MapArgs | undefined;
            if ((!args || args.configuration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configuration'");
            }
            if ((!args || args.mapName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mapName'");
            }
            resourceInputs["configuration"] = args ? args.configuration : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["mapName"] = args ? args.mapName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["mapArn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Map.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Map resources.
 */
export interface MapState {
    /**
     * Configuration block with the map style selected from an available data provider. Detailed below.
     */
    configuration?: pulumi.Input<inputs.location.MapConfiguration>;
    /**
     * The timestamp for when the map resource was created in ISO 8601 format.
     */
    createTime?: pulumi.Input<string>;
    /**
     * An optional description for the map resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) for the map resource. Used to specify a resource across all AWS.
     */
    mapArn?: pulumi.Input<string>;
    /**
     * The name for the map resource.
     *
     * The following arguments are optional:
     */
    mapName?: pulumi.Input<string>;
    /**
     * Key-value tags for the map. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The timestamp for when the map resource was last updated in ISO 8601 format.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Map resource.
 */
export interface MapArgs {
    /**
     * Configuration block with the map style selected from an available data provider. Detailed below.
     */
    configuration: pulumi.Input<inputs.location.MapConfiguration>;
    /**
     * An optional description for the map resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The name for the map resource.
     *
     * The following arguments are optional:
     */
    mapName: pulumi.Input<string>;
    /**
     * Key-value tags for the map. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
