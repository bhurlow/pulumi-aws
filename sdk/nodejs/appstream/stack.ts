// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides an AppStream stack.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.appstream.Stack("example", {
 *     applicationSettings: {
 *         enabled: true,
 *         settingsGroup: "SettingsGroup",
 *     },
 *     description: "stack description",
 *     displayName: "stack display name",
 *     feedbackUrl: "http://your-domain/feedback",
 *     redirectUrl: "http://your-domain/redirect",
 *     storageConnectors: [{
 *         connectorType: "HOMEFOLDERS",
 *     }],
 *     tags: {
 *         TagName: "TagValue",
 *     },
 *     userSettings: [
 *         {
 *             action: "CLIPBOARD_COPY_FROM_LOCAL_DEVICE",
 *             permission: "ENABLED",
 *         },
 *         {
 *             action: "CLIPBOARD_COPY_TO_LOCAL_DEVICE",
 *             permission: "ENABLED",
 *         },
 *         {
 *             action: "FILE_UPLOAD",
 *             permission: "ENABLED",
 *         },
 *         {
 *             action: "FILE_DOWNLOAD",
 *             permission: "ENABLED",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * `aws_appstream_stack` can be imported using the id, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:appstream/stack:Stack example stackID
 * ```
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
    public static readonly __pulumiType = 'aws:appstream/stack:Stack';

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
     * Set of configuration blocks defining the interface VPC endpoints. Users of the stack can connect to AppStream 2.0 only through the specified endpoints.
     * See `accessEndpoints` below.
     */
    public readonly accessEndpoints!: pulumi.Output<outputs.appstream.StackAccessEndpoint[]>;
    /**
     * Settings for application settings persistence.
     * See `applicationSettings` below.
     */
    public readonly applicationSettings!: pulumi.Output<outputs.appstream.StackApplicationSettings>;
    /**
     * ARN of the appstream stack.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Date and time, in UTC and extended RFC 3339 format, when the stack was created.
     */
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    /**
     * Description for the AppStream stack.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Stack name to display.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.
     */
    public readonly embedHostDomains!: pulumi.Output<string[]>;
    /**
     * URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed. .
     */
    public readonly feedbackUrl!: pulumi.Output<string>;
    /**
     * Unique name for the AppStream stack.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * URL that users are redirected to after their streaming session ends.
     */
    public readonly redirectUrl!: pulumi.Output<string>;
    /**
     * Configuration block for the storage connectors to enable.
     * See `storageConnectors` below.
     */
    public readonly storageConnectors!: pulumi.Output<outputs.appstream.StackStorageConnector[]>;
    /**
     * Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Configuration block for the actions that are enabled or disabled for users during their streaming sessions. By default, these actions are enabled.
     * See `userSettings` below.
     */
    public readonly userSettings!: pulumi.Output<outputs.appstream.StackUserSetting[]>;

    /**
     * Create a Stack resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: StackArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StackArgs | StackState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StackState | undefined;
            resourceInputs["accessEndpoints"] = state ? state.accessEndpoints : undefined;
            resourceInputs["applicationSettings"] = state ? state.applicationSettings : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["createdTime"] = state ? state.createdTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["embedHostDomains"] = state ? state.embedHostDomains : undefined;
            resourceInputs["feedbackUrl"] = state ? state.feedbackUrl : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["redirectUrl"] = state ? state.redirectUrl : undefined;
            resourceInputs["storageConnectors"] = state ? state.storageConnectors : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["userSettings"] = state ? state.userSettings : undefined;
        } else {
            const args = argsOrState as StackArgs | undefined;
            resourceInputs["accessEndpoints"] = args ? args.accessEndpoints : undefined;
            resourceInputs["applicationSettings"] = args ? args.applicationSettings : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["embedHostDomains"] = args ? args.embedHostDomains : undefined;
            resourceInputs["feedbackUrl"] = args ? args.feedbackUrl : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["redirectUrl"] = args ? args.redirectUrl : undefined;
            resourceInputs["storageConnectors"] = args ? args.storageConnectors : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userSettings"] = args ? args.userSettings : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
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
     * Set of configuration blocks defining the interface VPC endpoints. Users of the stack can connect to AppStream 2.0 only through the specified endpoints.
     * See `accessEndpoints` below.
     */
    accessEndpoints?: pulumi.Input<pulumi.Input<inputs.appstream.StackAccessEndpoint>[]>;
    /**
     * Settings for application settings persistence.
     * See `applicationSettings` below.
     */
    applicationSettings?: pulumi.Input<inputs.appstream.StackApplicationSettings>;
    /**
     * ARN of the appstream stack.
     */
    arn?: pulumi.Input<string>;
    /**
     * Date and time, in UTC and extended RFC 3339 format, when the stack was created.
     */
    createdTime?: pulumi.Input<string>;
    /**
     * Description for the AppStream stack.
     */
    description?: pulumi.Input<string>;
    /**
     * Stack name to display.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.
     */
    embedHostDomains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed. .
     */
    feedbackUrl?: pulumi.Input<string>;
    /**
     * Unique name for the AppStream stack.
     */
    name?: pulumi.Input<string>;
    /**
     * URL that users are redirected to after their streaming session ends.
     */
    redirectUrl?: pulumi.Input<string>;
    /**
     * Configuration block for the storage connectors to enable.
     * See `storageConnectors` below.
     */
    storageConnectors?: pulumi.Input<pulumi.Input<inputs.appstream.StackStorageConnector>[]>;
    /**
     * Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block for the actions that are enabled or disabled for users during their streaming sessions. By default, these actions are enabled.
     * See `userSettings` below.
     */
    userSettings?: pulumi.Input<pulumi.Input<inputs.appstream.StackUserSetting>[]>;
}

/**
 * The set of arguments for constructing a Stack resource.
 */
export interface StackArgs {
    /**
     * Set of configuration blocks defining the interface VPC endpoints. Users of the stack can connect to AppStream 2.0 only through the specified endpoints.
     * See `accessEndpoints` below.
     */
    accessEndpoints?: pulumi.Input<pulumi.Input<inputs.appstream.StackAccessEndpoint>[]>;
    /**
     * Settings for application settings persistence.
     * See `applicationSettings` below.
     */
    applicationSettings?: pulumi.Input<inputs.appstream.StackApplicationSettings>;
    /**
     * Description for the AppStream stack.
     */
    description?: pulumi.Input<string>;
    /**
     * Stack name to display.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.
     */
    embedHostDomains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed. .
     */
    feedbackUrl?: pulumi.Input<string>;
    /**
     * Unique name for the AppStream stack.
     */
    name?: pulumi.Input<string>;
    /**
     * URL that users are redirected to after their streaming session ends.
     */
    redirectUrl?: pulumi.Input<string>;
    /**
     * Configuration block for the storage connectors to enable.
     * See `storageConnectors` below.
     */
    storageConnectors?: pulumi.Input<pulumi.Input<inputs.appstream.StackStorageConnector>[]>;
    /**
     * Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block for the actions that are enabled or disabled for users during their streaming sessions. By default, these actions are enabled.
     * See `userSettings` below.
     */
    userSettings?: pulumi.Input<pulumi.Input<inputs.appstream.StackUserSetting>[]>;
}
