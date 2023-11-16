// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource for managing an Athena Prepared Statement.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testBucketV2 = new aws.s3.BucketV2("testBucketV2", {forceDestroy: true});
 * const testWorkgroup = new aws.athena.Workgroup("testWorkgroup", {});
 * const testDatabase = new aws.athena.Database("testDatabase", {
 *     name: "example",
 *     bucket: testBucketV2.bucket,
 * });
 * const testPreparedStatement = new aws.athena.PreparedStatement("testPreparedStatement", {
 *     queryStatement: pulumi.interpolate`SELECT * FROM ${testDatabase.name} WHERE x = ?`,
 *     workgroup: testWorkgroup.name,
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import Athena Prepared Statement using the `WORKGROUP-NAME/STATEMENT-NAME`. For example:
 *
 * ```sh
 *  $ pulumi import aws:athena/preparedStatement:PreparedStatement example 12345abcde/example
 * ```
 */
export class PreparedStatement extends pulumi.CustomResource {
    /**
     * Get an existing PreparedStatement resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PreparedStatementState, opts?: pulumi.CustomResourceOptions): PreparedStatement {
        return new PreparedStatement(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:athena/preparedStatement:PreparedStatement';

    /**
     * Returns true if the given object is an instance of PreparedStatement.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PreparedStatement {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PreparedStatement.__pulumiType;
    }

    /**
     * Brief explanation of prepared statement. Maximum length of 1024.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the prepared statement. Maximum length of 256.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The query string for the prepared statement.
     */
    public readonly queryStatement!: pulumi.Output<string>;
    /**
     * The name of the workgroup to which the prepared statement belongs.
     */
    public readonly workgroup!: pulumi.Output<string>;

    /**
     * Create a PreparedStatement resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PreparedStatementArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PreparedStatementArgs | PreparedStatementState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PreparedStatementState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["queryStatement"] = state ? state.queryStatement : undefined;
            resourceInputs["workgroup"] = state ? state.workgroup : undefined;
        } else {
            const args = argsOrState as PreparedStatementArgs | undefined;
            if ((!args || args.queryStatement === undefined) && !opts.urn) {
                throw new Error("Missing required property 'queryStatement'");
            }
            if ((!args || args.workgroup === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workgroup'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["queryStatement"] = args ? args.queryStatement : undefined;
            resourceInputs["workgroup"] = args ? args.workgroup : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PreparedStatement.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PreparedStatement resources.
 */
export interface PreparedStatementState {
    /**
     * Brief explanation of prepared statement. Maximum length of 1024.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the prepared statement. Maximum length of 256.
     */
    name?: pulumi.Input<string>;
    /**
     * The query string for the prepared statement.
     */
    queryStatement?: pulumi.Input<string>;
    /**
     * The name of the workgroup to which the prepared statement belongs.
     */
    workgroup?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PreparedStatement resource.
 */
export interface PreparedStatementArgs {
    /**
     * Brief explanation of prepared statement. Maximum length of 1024.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the prepared statement. Maximum length of 256.
     */
    name?: pulumi.Input<string>;
    /**
     * The query string for the prepared statement.
     */
    queryStatement: pulumi.Input<string>;
    /**
     * The name of the workgroup to which the prepared statement belongs.
     */
    workgroup: pulumi.Input<string>;
}
