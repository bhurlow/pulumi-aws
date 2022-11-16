// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ApplicationInsights Application resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleGroup = new aws.resourcegroups.Group("exampleGroup", {resourceQuery: {
 *     query: `	{
 * 		"ResourceTypeFilters": [
 * 		  "AWS::EC2::Instance"
 * 		],
 * 		"TagFilters": [
 * 		  {
 * 			"Key": "Stage",
 * 			"Values": [
 * 			  "Test"
 * 			]
 * 		  }
 * 		]
 * 	  }
 * `,
 * }});
 * const exampleApplication = new aws.applicationinsights.Application("exampleApplication", {resourceGroupName: exampleGroup.name});
 * ```
 *
 * ## Import
 *
 * ApplicationInsights Applications can be imported using the `resource_group_name`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:applicationinsights/application:Application some some-application
 * ```
 */
export class Application extends pulumi.CustomResource {
    /**
     * Get an existing Application resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationState, opts?: pulumi.CustomResourceOptions): Application {
        return new Application(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:applicationinsights/application:Application';

    /**
     * Returns true if the given object is an instance of Application.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Application {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Application.__pulumiType;
    }

    /**
     * ARN of the Application.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Indicates whether Application Insights automatically configures unmonitored resources in the resource group.
     */
    public readonly autoConfigEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Configures all of the resources in the resource group by applying the recommended configurations.
     */
    public readonly autoCreate!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether Application Insights can listen to CloudWatch events for the application resources, such as instance terminated, failed deployment, and others.
     */
    public readonly cweMonitorEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Application Insights can create applications based on a resource group or on an account. To create an account-based application using all of the resources in the account, set this parameter to `ACCOUNT_BASED`.
     */
    public readonly groupingType!: pulumi.Output<string | undefined>;
    /**
     * When set to `true`, creates opsItems for any problems detected on an application.
     */
    public readonly opsCenterEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * SNS topic provided to Application Insights that is associated to the created opsItem. Allows you to receive notifications for updates to the opsItem.
     */
    public readonly opsItemSnsTopicArn!: pulumi.Output<string | undefined>;
    /**
     * Name of the resource group.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationArgs | ApplicationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["autoConfigEnabled"] = state ? state.autoConfigEnabled : undefined;
            resourceInputs["autoCreate"] = state ? state.autoCreate : undefined;
            resourceInputs["cweMonitorEnabled"] = state ? state.cweMonitorEnabled : undefined;
            resourceInputs["groupingType"] = state ? state.groupingType : undefined;
            resourceInputs["opsCenterEnabled"] = state ? state.opsCenterEnabled : undefined;
            resourceInputs["opsItemSnsTopicArn"] = state ? state.opsItemSnsTopicArn : undefined;
            resourceInputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as ApplicationArgs | undefined;
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["autoConfigEnabled"] = args ? args.autoConfigEnabled : undefined;
            resourceInputs["autoCreate"] = args ? args.autoCreate : undefined;
            resourceInputs["cweMonitorEnabled"] = args ? args.cweMonitorEnabled : undefined;
            resourceInputs["groupingType"] = args ? args.groupingType : undefined;
            resourceInputs["opsCenterEnabled"] = args ? args.opsCenterEnabled : undefined;
            resourceInputs["opsItemSnsTopicArn"] = args ? args.opsItemSnsTopicArn : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Application.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Application resources.
 */
export interface ApplicationState {
    /**
     * ARN of the Application.
     */
    arn?: pulumi.Input<string>;
    /**
     * Indicates whether Application Insights automatically configures unmonitored resources in the resource group.
     */
    autoConfigEnabled?: pulumi.Input<boolean>;
    /**
     * Configures all of the resources in the resource group by applying the recommended configurations.
     */
    autoCreate?: pulumi.Input<boolean>;
    /**
     * Indicates whether Application Insights can listen to CloudWatch events for the application resources, such as instance terminated, failed deployment, and others.
     */
    cweMonitorEnabled?: pulumi.Input<boolean>;
    /**
     * Application Insights can create applications based on a resource group or on an account. To create an account-based application using all of the resources in the account, set this parameter to `ACCOUNT_BASED`.
     */
    groupingType?: pulumi.Input<string>;
    /**
     * When set to `true`, creates opsItems for any problems detected on an application.
     */
    opsCenterEnabled?: pulumi.Input<boolean>;
    /**
     * SNS topic provided to Application Insights that is associated to the created opsItem. Allows you to receive notifications for updates to the opsItem.
     */
    opsItemSnsTopicArn?: pulumi.Input<string>;
    /**
     * Name of the resource group.
     */
    resourceGroupName?: pulumi.Input<string>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    /**
     * Indicates whether Application Insights automatically configures unmonitored resources in the resource group.
     */
    autoConfigEnabled?: pulumi.Input<boolean>;
    /**
     * Configures all of the resources in the resource group by applying the recommended configurations.
     */
    autoCreate?: pulumi.Input<boolean>;
    /**
     * Indicates whether Application Insights can listen to CloudWatch events for the application resources, such as instance terminated, failed deployment, and others.
     */
    cweMonitorEnabled?: pulumi.Input<boolean>;
    /**
     * Application Insights can create applications based on a resource group or on an account. To create an account-based application using all of the resources in the account, set this parameter to `ACCOUNT_BASED`.
     */
    groupingType?: pulumi.Input<string>;
    /**
     * When set to `true`, creates opsItems for any problems detected on an application.
     */
    opsCenterEnabled?: pulumi.Input<boolean>;
    /**
     * SNS topic provided to Application Insights that is associated to the created opsItem. Allows you to receive notifications for updates to the opsItem.
     */
    opsItemSnsTopicArn?: pulumi.Input<string>;
    /**
     * Name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
