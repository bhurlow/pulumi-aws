// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides an OpsWorks stack resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const main = new aws.opsworks.Stack("main", {
 *     region: "us-west-1",
 *     serviceRoleArn: aws_iam_role.opsworks.arn,
 *     defaultInstanceProfileArn: aws_iam_instance_profile.opsworks.arn,
 *     tags: {
 *         Name: "foobar-stack",
 *     },
 *     customJson: `{
 *  "foobar": {
 *     "version": "1.0.0"
 *   }
 * }
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_opsworks_stack.bar
 *
 *  id = "00000000-0000-0000-0000-000000000000" } Using `pulumi import`, import OpsWorks stacks using the `id`. For exampleconsole % pulumi import aws_opsworks_stack.bar 00000000-0000-0000-0000-000000000000
 */
export class Stack extends pulumi.CustomResource {
    /**
     * Get an existing Stack resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StackState, opts?: pulumi.CustomResourceOptions): Stack {
        return new Stack(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:opsworks/stack:Stack';

    /**
     * Returns true if the given object is an instance of Stack.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Stack {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Stack.__pulumiType;
    }

    /**
     * If set to `"LATEST"`, OpsWorks will automatically install the latest version.
     */
    public readonly agentVersion!: pulumi.Output<string>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * If `manageBerkshelf` is enabled, the version of Berkshelf to use.
     */
    public readonly berkshelfVersion!: pulumi.Output<string | undefined>;
    /**
     * Color to paint next to the stack's resources in the OpsWorks console.
     */
    public readonly color!: pulumi.Output<string | undefined>;
    /**
     * Name of the configuration manager to use. Defaults to "Chef".
     */
    public readonly configurationManagerName!: pulumi.Output<string | undefined>;
    /**
     * Version of the configuration manager to use. Defaults to "11.4".
     */
    public readonly configurationManagerVersion!: pulumi.Output<string | undefined>;
    /**
     * When `useCustomCookbooks` is set, provide this sub-object as described below.
     */
    public readonly customCookbooksSources!: pulumi.Output<outputs.opsworks.StackCustomCookbooksSource[]>;
    /**
     * Custom JSON attributes to apply to the entire stack.
     */
    public readonly customJson!: pulumi.Output<string | undefined>;
    /**
     * Name of the availability zone where instances will be created by default.
     * Cannot be set when `vpcId` is set.
     */
    public readonly defaultAvailabilityZone!: pulumi.Output<string>;
    /**
     * The ARN of an IAM Instance Profile that created instances will have by default.
     */
    public readonly defaultInstanceProfileArn!: pulumi.Output<string>;
    /**
     * Name of OS that will be installed on instances by default.
     */
    public readonly defaultOs!: pulumi.Output<string | undefined>;
    /**
     * Name of the type of root device instances will have by default.
     */
    public readonly defaultRootDeviceType!: pulumi.Output<string | undefined>;
    /**
     * Name of the SSH keypair that instances will have by default.
     */
    public readonly defaultSshKeyName!: pulumi.Output<string | undefined>;
    /**
     * ID of the subnet in which instances will be created by default.
     * Required if `vpcId` is set to a VPC other than the default VPC, and forbidden if it isn't.
     */
    public readonly defaultSubnetId!: pulumi.Output<string>;
    /**
     * Keyword representing the naming scheme that will be used for instance hostnames within this stack.
     */
    public readonly hostnameTheme!: pulumi.Output<string | undefined>;
    /**
     * Boolean value controlling whether Opsworks will run Berkshelf for this stack.
     */
    public readonly manageBerkshelf!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the stack.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the region where the stack will exist.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The ARN of an IAM role that the OpsWorks service will act as.
     */
    public readonly serviceRoleArn!: pulumi.Output<string>;
    public /*out*/ readonly stackEndpoint!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource.
     * If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Boolean value controlling whether the custom cookbook settings are enabled.
     */
    public readonly useCustomCookbooks!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean value controlling whether the standard OpsWorks security groups apply to created instances.
     */
    public readonly useOpsworksSecurityGroups!: pulumi.Output<boolean | undefined>;
    /**
     * ID of the VPC that this stack belongs to.
     * Defaults to the region's default VPC.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a Stack resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StackArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StackArgs | StackState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StackState | undefined;
            resourceInputs["agentVersion"] = state ? state.agentVersion : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["berkshelfVersion"] = state ? state.berkshelfVersion : undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["configurationManagerName"] = state ? state.configurationManagerName : undefined;
            resourceInputs["configurationManagerVersion"] = state ? state.configurationManagerVersion : undefined;
            resourceInputs["customCookbooksSources"] = state ? state.customCookbooksSources : undefined;
            resourceInputs["customJson"] = state ? state.customJson : undefined;
            resourceInputs["defaultAvailabilityZone"] = state ? state.defaultAvailabilityZone : undefined;
            resourceInputs["defaultInstanceProfileArn"] = state ? state.defaultInstanceProfileArn : undefined;
            resourceInputs["defaultOs"] = state ? state.defaultOs : undefined;
            resourceInputs["defaultRootDeviceType"] = state ? state.defaultRootDeviceType : undefined;
            resourceInputs["defaultSshKeyName"] = state ? state.defaultSshKeyName : undefined;
            resourceInputs["defaultSubnetId"] = state ? state.defaultSubnetId : undefined;
            resourceInputs["hostnameTheme"] = state ? state.hostnameTheme : undefined;
            resourceInputs["manageBerkshelf"] = state ? state.manageBerkshelf : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["serviceRoleArn"] = state ? state.serviceRoleArn : undefined;
            resourceInputs["stackEndpoint"] = state ? state.stackEndpoint : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["useCustomCookbooks"] = state ? state.useCustomCookbooks : undefined;
            resourceInputs["useOpsworksSecurityGroups"] = state ? state.useOpsworksSecurityGroups : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as StackArgs | undefined;
            if ((!args || args.defaultInstanceProfileArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultInstanceProfileArn'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.serviceRoleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceRoleArn'");
            }
            resourceInputs["agentVersion"] = args ? args.agentVersion : undefined;
            resourceInputs["berkshelfVersion"] = args ? args.berkshelfVersion : undefined;
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["configurationManagerName"] = args ? args.configurationManagerName : undefined;
            resourceInputs["configurationManagerVersion"] = args ? args.configurationManagerVersion : undefined;
            resourceInputs["customCookbooksSources"] = args ? args.customCookbooksSources : undefined;
            resourceInputs["customJson"] = args ? args.customJson : undefined;
            resourceInputs["defaultAvailabilityZone"] = args ? args.defaultAvailabilityZone : undefined;
            resourceInputs["defaultInstanceProfileArn"] = args ? args.defaultInstanceProfileArn : undefined;
            resourceInputs["defaultOs"] = args ? args.defaultOs : undefined;
            resourceInputs["defaultRootDeviceType"] = args ? args.defaultRootDeviceType : undefined;
            resourceInputs["defaultSshKeyName"] = args ? args.defaultSshKeyName : undefined;
            resourceInputs["defaultSubnetId"] = args ? args.defaultSubnetId : undefined;
            resourceInputs["hostnameTheme"] = args ? args.hostnameTheme : undefined;
            resourceInputs["manageBerkshelf"] = args ? args.manageBerkshelf : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["serviceRoleArn"] = args ? args.serviceRoleArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["useCustomCookbooks"] = args ? args.useCustomCookbooks : undefined;
            resourceInputs["useOpsworksSecurityGroups"] = args ? args.useOpsworksSecurityGroups : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["stackEndpoint"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Stack.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Stack resources.
 */
export interface StackState {
    /**
     * If set to `"LATEST"`, OpsWorks will automatically install the latest version.
     */
    agentVersion?: pulumi.Input<string>;
    arn?: pulumi.Input<string>;
    /**
     * If `manageBerkshelf` is enabled, the version of Berkshelf to use.
     */
    berkshelfVersion?: pulumi.Input<string>;
    /**
     * Color to paint next to the stack's resources in the OpsWorks console.
     */
    color?: pulumi.Input<string>;
    /**
     * Name of the configuration manager to use. Defaults to "Chef".
     */
    configurationManagerName?: pulumi.Input<string>;
    /**
     * Version of the configuration manager to use. Defaults to "11.4".
     */
    configurationManagerVersion?: pulumi.Input<string>;
    /**
     * When `useCustomCookbooks` is set, provide this sub-object as described below.
     */
    customCookbooksSources?: pulumi.Input<pulumi.Input<inputs.opsworks.StackCustomCookbooksSource>[]>;
    /**
     * Custom JSON attributes to apply to the entire stack.
     */
    customJson?: pulumi.Input<string>;
    /**
     * Name of the availability zone where instances will be created by default.
     * Cannot be set when `vpcId` is set.
     */
    defaultAvailabilityZone?: pulumi.Input<string>;
    /**
     * The ARN of an IAM Instance Profile that created instances will have by default.
     */
    defaultInstanceProfileArn?: pulumi.Input<string>;
    /**
     * Name of OS that will be installed on instances by default.
     */
    defaultOs?: pulumi.Input<string>;
    /**
     * Name of the type of root device instances will have by default.
     */
    defaultRootDeviceType?: pulumi.Input<string>;
    /**
     * Name of the SSH keypair that instances will have by default.
     */
    defaultSshKeyName?: pulumi.Input<string>;
    /**
     * ID of the subnet in which instances will be created by default.
     * Required if `vpcId` is set to a VPC other than the default VPC, and forbidden if it isn't.
     */
    defaultSubnetId?: pulumi.Input<string>;
    /**
     * Keyword representing the naming scheme that will be used for instance hostnames within this stack.
     */
    hostnameTheme?: pulumi.Input<string>;
    /**
     * Boolean value controlling whether Opsworks will run Berkshelf for this stack.
     */
    manageBerkshelf?: pulumi.Input<boolean>;
    /**
     * The name of the stack.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the region where the stack will exist.
     */
    region?: pulumi.Input<string>;
    /**
     * The ARN of an IAM role that the OpsWorks service will act as.
     */
    serviceRoleArn?: pulumi.Input<string>;
    stackEndpoint?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     * If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Boolean value controlling whether the custom cookbook settings are enabled.
     */
    useCustomCookbooks?: pulumi.Input<boolean>;
    /**
     * Boolean value controlling whether the standard OpsWorks security groups apply to created instances.
     */
    useOpsworksSecurityGroups?: pulumi.Input<boolean>;
    /**
     * ID of the VPC that this stack belongs to.
     * Defaults to the region's default VPC.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Stack resource.
 */
export interface StackArgs {
    /**
     * If set to `"LATEST"`, OpsWorks will automatically install the latest version.
     */
    agentVersion?: pulumi.Input<string>;
    /**
     * If `manageBerkshelf` is enabled, the version of Berkshelf to use.
     */
    berkshelfVersion?: pulumi.Input<string>;
    /**
     * Color to paint next to the stack's resources in the OpsWorks console.
     */
    color?: pulumi.Input<string>;
    /**
     * Name of the configuration manager to use. Defaults to "Chef".
     */
    configurationManagerName?: pulumi.Input<string>;
    /**
     * Version of the configuration manager to use. Defaults to "11.4".
     */
    configurationManagerVersion?: pulumi.Input<string>;
    /**
     * When `useCustomCookbooks` is set, provide this sub-object as described below.
     */
    customCookbooksSources?: pulumi.Input<pulumi.Input<inputs.opsworks.StackCustomCookbooksSource>[]>;
    /**
     * Custom JSON attributes to apply to the entire stack.
     */
    customJson?: pulumi.Input<string>;
    /**
     * Name of the availability zone where instances will be created by default.
     * Cannot be set when `vpcId` is set.
     */
    defaultAvailabilityZone?: pulumi.Input<string>;
    /**
     * The ARN of an IAM Instance Profile that created instances will have by default.
     */
    defaultInstanceProfileArn: pulumi.Input<string>;
    /**
     * Name of OS that will be installed on instances by default.
     */
    defaultOs?: pulumi.Input<string>;
    /**
     * Name of the type of root device instances will have by default.
     */
    defaultRootDeviceType?: pulumi.Input<string>;
    /**
     * Name of the SSH keypair that instances will have by default.
     */
    defaultSshKeyName?: pulumi.Input<string>;
    /**
     * ID of the subnet in which instances will be created by default.
     * Required if `vpcId` is set to a VPC other than the default VPC, and forbidden if it isn't.
     */
    defaultSubnetId?: pulumi.Input<string>;
    /**
     * Keyword representing the naming scheme that will be used for instance hostnames within this stack.
     */
    hostnameTheme?: pulumi.Input<string>;
    /**
     * Boolean value controlling whether Opsworks will run Berkshelf for this stack.
     */
    manageBerkshelf?: pulumi.Input<boolean>;
    /**
     * The name of the stack.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the region where the stack will exist.
     */
    region: pulumi.Input<string>;
    /**
     * The ARN of an IAM role that the OpsWorks service will act as.
     */
    serviceRoleArn: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     * If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Boolean value controlling whether the custom cookbook settings are enabled.
     */
    useCustomCookbooks?: pulumi.Input<boolean>;
    /**
     * Boolean value controlling whether the standard OpsWorks security groups apply to created instances.
     */
    useOpsworksSecurityGroups?: pulumi.Input<boolean>;
    /**
     * ID of the VPC that this stack belongs to.
     * Defaults to the region's default VPC.
     */
    vpcId?: pulumi.Input<string>;
}
