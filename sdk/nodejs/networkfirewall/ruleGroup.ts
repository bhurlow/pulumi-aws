// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides an AWS Network Firewall Rule Group Resource
 *
 * ## Example Usage
 * ### Stateful Inspection for denying access to a domain
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.networkfirewall.RuleGroup("example", {
 *     capacity: 100,
 *     ruleGroup: {
 *         rulesSource: {
 *             rulesSourceList: {
 *                 generatedRulesType: "DENYLIST",
 *                 targetTypes: ["HTTP_HOST"],
 *                 targets: ["test.example.com"],
 *             },
 *         },
 *     },
 *     tags: {
 *         Tag1: "Value1",
 *         Tag2: "Value2",
 *     },
 *     type: "STATEFUL",
 * });
 * ```
 * ### Stateful Inspection from rules specifications defined in Suricata flat format
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as fs from "fs";
 *
 * const example = new aws.networkfirewall.RuleGroup("example", {
 *     capacity: 100,
 *     type: "STATEFUL",
 *     rules: fs.readFileSync("example.rules"),
 *     tags: {
 *         Tag1: "Value1",
 *         Tag2: "Value2",
 *     },
 * });
 * ```
 * ### Stateful Inspection from rule group specifications using rule variables and Suricata format rules
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as fs from "fs";
 *
 * const example = new aws.networkfirewall.RuleGroup("example", {
 *     capacity: 100,
 *     type: "STATEFUL",
 *     ruleGroup: {
 *         ruleVariables: {
 *             ipSets: [
 *                 {
 *                     key: "WEBSERVERS_HOSTS",
 *                     ipSet: {
 *                         definitions: [
 *                             "10.0.0.0/16",
 *                             "10.0.1.0/24",
 *                             "192.168.0.0/16",
 *                         ],
 *                     },
 *                 },
 *                 {
 *                     key: "EXTERNAL_HOST",
 *                     ipSet: {
 *                         definitions: ["1.2.3.4/32"],
 *                     },
 *                 },
 *             ],
 *             portSets: [{
 *                 key: "HTTP_PORTS",
 *                 portSet: {
 *                     definitions: [
 *                         "443",
 *                         "80",
 *                     ],
 *                 },
 *             }],
 *         },
 *         rulesSource: {
 *             rulesString: fs.readFileSync("suricata_rules_file"),
 *         },
 *     },
 *     tags: {
 *         Tag1: "Value1",
 *         Tag2: "Value2",
 *     },
 * });
 * ```
 * ### IP Set References to the Rule Group
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.networkfirewall.RuleGroup("example", {
 *     capacity: 100,
 *     type: "STATEFUL",
 *     ruleGroup: {
 *         rulesSource: {
 *             rulesSourceList: {
 *                 generatedRulesType: "DENYLIST",
 *                 targetTypes: ["HTTP_HOST"],
 *                 targets: ["test.example.com"],
 *             },
 *         },
 *         referenceSets: {
 *             ipSetReferences: [{
 *                 key: "example",
 *                 ipSetReferences: [{
 *                     referenceArn: aws_ec2_managed_prefix_list["this"].arn,
 *                 }],
 *             }],
 *         },
 *     },
 *     tags: {
 *         Tag1: "Value1",
 *         Tag2: "Value2",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Network Firewall Rule Groups can be imported using their `ARN`.
 *
 * ```sh
 *  $ pulumi import aws:networkfirewall/ruleGroup:RuleGroup example arn:aws:network-firewall:us-west-1:123456789012:stateful-rulegroup/example
 * ```
 */
export class RuleGroup extends pulumi.CustomResource {
    /**
     * Get an existing RuleGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleGroupState, opts?: pulumi.CustomResourceOptions): RuleGroup {
        return new RuleGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:networkfirewall/ruleGroup:RuleGroup';

    /**
     * Returns true if the given object is an instance of RuleGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RuleGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RuleGroup.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) that identifies the rule group.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
     */
    public readonly capacity!: pulumi.Output<number>;
    /**
     * A friendly description of the rule group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * KMS encryption configuration settings. See Encryption Configuration below for details.
     */
    public readonly encryptionConfiguration!: pulumi.Output<outputs.networkfirewall.RuleGroupEncryptionConfiguration | undefined>;
    /**
     * A friendly name of the rule group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A configuration block that defines the rule group rules. Required unless `rules` is specified. See Rule Group below for details.
     */
    public readonly ruleGroup!: pulumi.Output<outputs.networkfirewall.RuleGroupRuleGroup>;
    /**
     * The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless `ruleGroup` is specified.
     */
    public readonly rules!: pulumi.Output<string | undefined>;
    /**
     * A map of key:value pairs to associate with the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: `STATEFUL` or `STATELESS`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * A string token used when updating the rule group.
     */
    public /*out*/ readonly updateToken!: pulumi.Output<string>;

    /**
     * Create a RuleGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleGroupArgs | RuleGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleGroupState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["capacity"] = state ? state.capacity : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["encryptionConfiguration"] = state ? state.encryptionConfiguration : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["ruleGroup"] = state ? state.ruleGroup : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["updateToken"] = state ? state.updateToken : undefined;
        } else {
            const args = argsOrState as RuleGroupArgs | undefined;
            if ((!args || args.capacity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'capacity'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["capacity"] = args ? args.capacity : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["encryptionConfiguration"] = args ? args.encryptionConfiguration : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ruleGroup"] = args ? args.ruleGroup : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["updateToken"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RuleGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RuleGroup resources.
 */
export interface RuleGroupState {
    /**
     * The Amazon Resource Name (ARN) that identifies the rule group.
     */
    arn?: pulumi.Input<string>;
    /**
     * The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
     */
    capacity?: pulumi.Input<number>;
    /**
     * A friendly description of the rule group.
     */
    description?: pulumi.Input<string>;
    /**
     * KMS encryption configuration settings. See Encryption Configuration below for details.
     */
    encryptionConfiguration?: pulumi.Input<inputs.networkfirewall.RuleGroupEncryptionConfiguration>;
    /**
     * A friendly name of the rule group.
     */
    name?: pulumi.Input<string>;
    /**
     * A configuration block that defines the rule group rules. Required unless `rules` is specified. See Rule Group below for details.
     */
    ruleGroup?: pulumi.Input<inputs.networkfirewall.RuleGroupRuleGroup>;
    /**
     * The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless `ruleGroup` is specified.
     */
    rules?: pulumi.Input<string>;
    /**
     * A map of key:value pairs to associate with the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: `STATEFUL` or `STATELESS`.
     */
    type?: pulumi.Input<string>;
    /**
     * A string token used when updating the rule group.
     */
    updateToken?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RuleGroup resource.
 */
export interface RuleGroupArgs {
    /**
     * The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
     */
    capacity: pulumi.Input<number>;
    /**
     * A friendly description of the rule group.
     */
    description?: pulumi.Input<string>;
    /**
     * KMS encryption configuration settings. See Encryption Configuration below for details.
     */
    encryptionConfiguration?: pulumi.Input<inputs.networkfirewall.RuleGroupEncryptionConfiguration>;
    /**
     * A friendly name of the rule group.
     */
    name?: pulumi.Input<string>;
    /**
     * A configuration block that defines the rule group rules. Required unless `rules` is specified. See Rule Group below for details.
     */
    ruleGroup?: pulumi.Input<inputs.networkfirewall.RuleGroupRuleGroup>;
    /**
     * The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless `ruleGroup` is specified.
     */
    rules?: pulumi.Input<string>;
    /**
     * A map of key:value pairs to associate with the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: `STATEFUL` or `STATELESS`.
     */
    type: pulumi.Input<string>;
}
