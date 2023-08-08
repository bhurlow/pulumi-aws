// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an AppStream Fleet Stack association.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleFleet = new aws.appstream.Fleet("exampleFleet", {
 *     imageName: "Amazon-AppStream2-Sample-Image-02-04-2019",
 *     instanceType: "stream.standard.small",
 *     computeCapacity: {
 *         desiredInstances: 1,
 *     },
 * });
 * const exampleStack = new aws.appstream.Stack("exampleStack", {});
 * const exampleFleetStackAssociation = new aws.appstream.FleetStackAssociation("exampleFleetStackAssociation", {
 *     fleetName: exampleFleet.name,
 *     stackName: exampleStack.name,
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_appstream_fleet_stack_association.example
 *
 *  id = "fleetName/stackName" } Using `pulumi import`, import AppStream Stack Fleet Association using the `fleet_name` and `stack_name` separated by a slash (`/`). For exampleconsole % pulumi import aws_appstream_fleet_stack_association.example fleetName/stackName
 */
export class FleetStackAssociation extends pulumi.CustomResource {
    /**
     * Get an existing FleetStackAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FleetStackAssociationState, opts?: pulumi.CustomResourceOptions): FleetStackAssociation {
        return new FleetStackAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:appstream/fleetStackAssociation:FleetStackAssociation';

    /**
     * Returns true if the given object is an instance of FleetStackAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FleetStackAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FleetStackAssociation.__pulumiType;
    }

    /**
     * Name of the fleet.
     */
    public readonly fleetName!: pulumi.Output<string>;
    /**
     * Name of the stack.
     */
    public readonly stackName!: pulumi.Output<string>;

    /**
     * Create a FleetStackAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FleetStackAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FleetStackAssociationArgs | FleetStackAssociationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FleetStackAssociationState | undefined;
            resourceInputs["fleetName"] = state ? state.fleetName : undefined;
            resourceInputs["stackName"] = state ? state.stackName : undefined;
        } else {
            const args = argsOrState as FleetStackAssociationArgs | undefined;
            if ((!args || args.fleetName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fleetName'");
            }
            if ((!args || args.stackName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'stackName'");
            }
            resourceInputs["fleetName"] = args ? args.fleetName : undefined;
            resourceInputs["stackName"] = args ? args.stackName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FleetStackAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FleetStackAssociation resources.
 */
export interface FleetStackAssociationState {
    /**
     * Name of the fleet.
     */
    fleetName?: pulumi.Input<string>;
    /**
     * Name of the stack.
     */
    stackName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FleetStackAssociation resource.
 */
export interface FleetStackAssociationArgs {
    /**
     * Name of the fleet.
     */
    fleetName: pulumi.Input<string>;
    /**
     * Name of the stack.
     */
    stackName: pulumi.Input<string>;
}
