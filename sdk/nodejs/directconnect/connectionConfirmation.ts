// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a confirmation of the creation of the specified hosted connection on an interconnect.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const confirmation = new aws.directconnect.ConnectionConfirmation("confirmation", {
 *     connectionId: "dxcon-ffabc123",
 * });
 * ```
 */
export class ConnectionConfirmation extends pulumi.CustomResource {
    /**
     * Get an existing ConnectionConfirmation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConnectionConfirmationState, opts?: pulumi.CustomResourceOptions): ConnectionConfirmation {
        return new ConnectionConfirmation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:directconnect/connectionConfirmation:ConnectionConfirmation';

    /**
     * Returns true if the given object is an instance of ConnectionConfirmation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectionConfirmation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectionConfirmation.__pulumiType;
    }

    /**
     * The ID of the hosted connection.
     */
    public readonly connectionId!: pulumi.Output<string>;

    /**
     * Create a ConnectionConfirmation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionConfirmationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectionConfirmationArgs | ConnectionConfirmationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConnectionConfirmationState | undefined;
            inputs["connectionId"] = state ? state.connectionId : undefined;
        } else {
            const args = argsOrState as ConnectionConfirmationArgs | undefined;
            if ((!args || args.connectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionId'");
            }
            inputs["connectionId"] = args ? args.connectionId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ConnectionConfirmation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConnectionConfirmation resources.
 */
export interface ConnectionConfirmationState {
    /**
     * The ID of the hosted connection.
     */
    connectionId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConnectionConfirmation resource.
 */
export interface ConnectionConfirmationArgs {
    /**
     * The ID of the hosted connection.
     */
    connectionId: pulumi.Input<string>;
}
