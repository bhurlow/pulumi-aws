// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a resource to create an EventBridge permission to support cross-account events in the current account default event bus.
 *
 * > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
 *
 * > **Note:** The EventBridge bus policy resource  (`aws.cloudwatch.EventBusPolicy`) is incompatible with the EventBridge permission resource (`aws.cloudwatch.EventPermission`) and will overwrite permissions.
 *
 * ## Example Usage
 * ### Account Access
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const devAccountAccess = new aws.cloudwatch.EventPermission("devAccountAccess", {
 *     principal: "123456789012",
 *     statementId: "DevAccountAccess",
 * });
 * ```
 * ### Organization Access
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const organizationAccess = new aws.cloudwatch.EventPermission("organizationAccess", {
 *     principal: "*",
 *     statementId: "OrganizationAccess",
 *     condition: {
 *         key: "aws:PrincipalOrgID",
 *         type: "StringEquals",
 *         value: aws_organizations_organization.example.id,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_cloudwatch_event_permission.DevAccountAccess
 *
 *  id = "example-event-bus/DevAccountAccess" } Using `pulumi import`, import EventBridge permissions using the `event_bus_name/statement_id` (if you omit `event_bus_name`, the `default` event bus will be used). For exampleconsole % pulumi import aws_cloudwatch_event_permission.DevAccountAccess example-event-bus/DevAccountAccess
 */
export class EventPermission extends pulumi.CustomResource {
    /**
     * Get an existing EventPermission resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventPermissionState, opts?: pulumi.CustomResourceOptions): EventPermission {
        return new EventPermission(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudwatch/eventPermission:EventPermission';

    /**
     * Returns true if the given object is an instance of EventPermission.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventPermission {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventPermission.__pulumiType;
    }

    /**
     * The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
     */
    public readonly action!: pulumi.Output<string | undefined>;
    /**
     * Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
     */
    public readonly condition!: pulumi.Output<outputs.cloudwatch.EventPermissionCondition | undefined>;
    /**
     * The name of the event bus to set the permissions on.
     * If you omit this, the permissions are set on the `default` event bus.
     */
    public readonly eventBusName!: pulumi.Output<string | undefined>;
    /**
     * The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
     */
    public readonly principal!: pulumi.Output<string>;
    /**
     * An identifier string for the external account that you are granting permissions to.
     */
    public readonly statementId!: pulumi.Output<string>;

    /**
     * Create a EventPermission resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventPermissionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventPermissionArgs | EventPermissionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventPermissionState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["condition"] = state ? state.condition : undefined;
            resourceInputs["eventBusName"] = state ? state.eventBusName : undefined;
            resourceInputs["principal"] = state ? state.principal : undefined;
            resourceInputs["statementId"] = state ? state.statementId : undefined;
        } else {
            const args = argsOrState as EventPermissionArgs | undefined;
            if ((!args || args.principal === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principal'");
            }
            if ((!args || args.statementId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'statementId'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["condition"] = args ? args.condition : undefined;
            resourceInputs["eventBusName"] = args ? args.eventBusName : undefined;
            resourceInputs["principal"] = args ? args.principal : undefined;
            resourceInputs["statementId"] = args ? args.statementId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EventPermission.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventPermission resources.
 */
export interface EventPermissionState {
    /**
     * The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
     */
    action?: pulumi.Input<string>;
    /**
     * Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
     */
    condition?: pulumi.Input<inputs.cloudwatch.EventPermissionCondition>;
    /**
     * The name of the event bus to set the permissions on.
     * If you omit this, the permissions are set on the `default` event bus.
     */
    eventBusName?: pulumi.Input<string>;
    /**
     * The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
     */
    principal?: pulumi.Input<string>;
    /**
     * An identifier string for the external account that you are granting permissions to.
     */
    statementId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EventPermission resource.
 */
export interface EventPermissionArgs {
    /**
     * The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
     */
    action?: pulumi.Input<string>;
    /**
     * Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
     */
    condition?: pulumi.Input<inputs.cloudwatch.EventPermissionCondition>;
    /**
     * The name of the event bus to set the permissions on.
     * If you omit this, the permissions are set on the `default` event bus.
     */
    eventBusName?: pulumi.Input<string>;
    /**
     * The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
     */
    principal: pulumi.Input<string>;
    /**
     * An identifier string for the external account that you are granting permissions to.
     */
    statementId: pulumi.Input<string>;
}
