// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an AWS Route 53 Recovery Readiness Cell.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.route53recoveryreadiness.Cell("example", {cellName: "us-west-2-failover-cell"});
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_route53recoveryreadiness_cell.us-west-2-failover-cell
 *
 *  id = "us-west-2-failover-cell" } Using `pulumi import`, import Route53 Recovery Readiness cells using the cell name. For exampleconsole % pulumi import aws_route53recoveryreadiness_cell.us-west-2-failover-cell us-west-2-failover-cell
 */
export class Cell extends pulumi.CustomResource {
    /**
     * Get an existing Cell resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CellState, opts?: pulumi.CustomResourceOptions): Cell {
        return new Cell(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:route53recoveryreadiness/cell:Cell';

    /**
     * Returns true if the given object is an instance of Cell.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cell {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cell.__pulumiType;
    }

    /**
     * ARN of the cell
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Unique name describing the cell.
     *
     * The following arguments are optional:
     */
    public readonly cellName!: pulumi.Output<string>;
    /**
     * List of cell arns to add as nested fault domains within this cell.
     */
    public readonly cells!: pulumi.Output<string[] | undefined>;
    /**
     * List of readiness scopes (recovery groups or cells) that contain this cell.
     */
    public /*out*/ readonly parentReadinessScopes!: pulumi.Output<string[]>;
    /**
     * Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Cell resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CellArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CellArgs | CellState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CellState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["cellName"] = state ? state.cellName : undefined;
            resourceInputs["cells"] = state ? state.cells : undefined;
            resourceInputs["parentReadinessScopes"] = state ? state.parentReadinessScopes : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as CellArgs | undefined;
            if ((!args || args.cellName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cellName'");
            }
            resourceInputs["cellName"] = args ? args.cellName : undefined;
            resourceInputs["cells"] = args ? args.cells : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["parentReadinessScopes"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cell.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cell resources.
 */
export interface CellState {
    /**
     * ARN of the cell
     */
    arn?: pulumi.Input<string>;
    /**
     * Unique name describing the cell.
     *
     * The following arguments are optional:
     */
    cellName?: pulumi.Input<string>;
    /**
     * List of cell arns to add as nested fault domains within this cell.
     */
    cells?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of readiness scopes (recovery groups or cells) that contain this cell.
     */
    parentReadinessScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Cell resource.
 */
export interface CellArgs {
    /**
     * Unique name describing the cell.
     *
     * The following arguments are optional:
     */
    cellName: pulumi.Input<string>;
    /**
     * List of cell arns to add as nested fault domains within this cell.
     */
    cells?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
