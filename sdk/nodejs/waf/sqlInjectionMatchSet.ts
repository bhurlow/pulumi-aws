// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a WAF SQL Injection Match Set Resource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const sqlInjectionMatchSet = new aws.waf.SqlInjectionMatchSet("sqlInjectionMatchSet", {sqlInjectionMatchTuples: [{
 *     fieldToMatch: {
 *         type: "QUERY_STRING",
 *     },
 *     textTransformation: "URL_DECODE",
 * }]});
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_waf_sql_injection_match_set.example
 *
 *  id = "a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc" } Using `pulumi import`, import AWS WAF SQL Injection Match Set using their ID. For exampleconsole % pulumi import aws_waf_sql_injection_match_set.example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
 */
export class SqlInjectionMatchSet extends pulumi.CustomResource {
    /**
     * Get an existing SqlInjectionMatchSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SqlInjectionMatchSetState, opts?: pulumi.CustomResourceOptions): SqlInjectionMatchSet {
        return new SqlInjectionMatchSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:waf/sqlInjectionMatchSet:SqlInjectionMatchSet';

    /**
     * Returns true if the given object is an instance of SqlInjectionMatchSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SqlInjectionMatchSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SqlInjectionMatchSet.__pulumiType;
    }

    /**
     * The name or description of the SQL Injection Match Set.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
     */
    public readonly sqlInjectionMatchTuples!: pulumi.Output<outputs.waf.SqlInjectionMatchSetSqlInjectionMatchTuple[] | undefined>;

    /**
     * Create a SqlInjectionMatchSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SqlInjectionMatchSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SqlInjectionMatchSetArgs | SqlInjectionMatchSetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SqlInjectionMatchSetState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["sqlInjectionMatchTuples"] = state ? state.sqlInjectionMatchTuples : undefined;
        } else {
            const args = argsOrState as SqlInjectionMatchSetArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["sqlInjectionMatchTuples"] = args ? args.sqlInjectionMatchTuples : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SqlInjectionMatchSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SqlInjectionMatchSet resources.
 */
export interface SqlInjectionMatchSetState {
    /**
     * The name or description of the SQL Injection Match Set.
     */
    name?: pulumi.Input<string>;
    /**
     * The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
     */
    sqlInjectionMatchTuples?: pulumi.Input<pulumi.Input<inputs.waf.SqlInjectionMatchSetSqlInjectionMatchTuple>[]>;
}

/**
 * The set of arguments for constructing a SqlInjectionMatchSet resource.
 */
export interface SqlInjectionMatchSetArgs {
    /**
     * The name or description of the SQL Injection Match Set.
     */
    name?: pulumi.Input<string>;
    /**
     * The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
     */
    sqlInjectionMatchTuples?: pulumi.Input<pulumi.Input<inputs.waf.SqlInjectionMatchSetSqlInjectionMatchTuple>[]>;
}
