// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an AppConfig Hosted Configuration Version resource.
 *
 * ## Example Usage
 * ### Freeform
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.appconfig.HostedConfigurationVersion("example", {
 *     applicationId: aws_appconfig_application.example.id,
 *     configurationProfileId: aws_appconfig_configuration_profile.example.configuration_profile_id,
 *     description: "Example Freeform Hosted Configuration Version",
 *     contentType: "application/json",
 *     content: JSON.stringify({
 *         foo: "bar",
 *         fruit: [
 *             "apple",
 *             "pear",
 *             "orange",
 *         ],
 *         isThingEnabled: true,
 *     }),
 * });
 * ```
 * ### Feature Flags
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.appconfig.HostedConfigurationVersion("example", {
 *     applicationId: aws_appconfig_application.example.id,
 *     configurationProfileId: aws_appconfig_configuration_profile.example.configuration_profile_id,
 *     description: "Example Feature Flag Configuration Version",
 *     contentType: "application/json",
 *     content: JSON.stringify({
 *         flags: {
 *             foo: {
 *                 name: "foo",
 *                 _deprecation: {
 *                     status: "planned",
 *                 },
 *             },
 *             bar: {
 *                 name: "bar",
 *                 attributes: {
 *                     someAttribute: {
 *                         constraints: {
 *                             type: "string",
 *                             required: true,
 *                         },
 *                     },
 *                     someOtherAttribute: {
 *                         constraints: {
 *                             type: "number",
 *                             required: true,
 *                         },
 *                     },
 *                 },
 *             },
 *         },
 *         values: {
 *             foo: {
 *                 enabled: "true",
 *             },
 *             bar: {
 *                 enabled: "true",
 *                 someAttribute: "Hello World",
 *                 someOtherAttribute: 123,
 *             },
 *         },
 *         version: "1",
 *     }),
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_appconfig_hosted_configuration_version.example
 *
 *  id = "71abcde/11xxxxx/2" } Using `pulumi import`, import AppConfig Hosted Configuration Versions using the application ID, configuration profile ID, and version number separated by a slash (`/`). For exampleconsole % pulumi import aws_appconfig_hosted_configuration_version.example 71abcde/11xxxxx/2
 */
export class HostedConfigurationVersion extends pulumi.CustomResource {
    /**
     * Get an existing HostedConfigurationVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostedConfigurationVersionState, opts?: pulumi.CustomResourceOptions): HostedConfigurationVersion {
        return new HostedConfigurationVersion(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:appconfig/hostedConfigurationVersion:HostedConfigurationVersion';

    /**
     * Returns true if the given object is an instance of HostedConfigurationVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HostedConfigurationVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HostedConfigurationVersion.__pulumiType;
    }

    /**
     * Application ID.
     */
    public readonly applicationId!: pulumi.Output<string>;
    /**
     * ARN of the AppConfig  hosted configuration version.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Configuration profile ID.
     */
    public readonly configurationProfileId!: pulumi.Output<string>;
    /**
     * Content of the configuration or the configuration data.
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * Standard MIME type describing the format of the configuration content. For more information, see [Content-Type](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
     */
    public readonly contentType!: pulumi.Output<string>;
    /**
     * Description of the configuration.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Version number of the hosted configuration.
     */
    public /*out*/ readonly versionNumber!: pulumi.Output<number>;

    /**
     * Create a HostedConfigurationVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostedConfigurationVersionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostedConfigurationVersionArgs | HostedConfigurationVersionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HostedConfigurationVersionState | undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["configurationProfileId"] = state ? state.configurationProfileId : undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["contentType"] = state ? state.contentType : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["versionNumber"] = state ? state.versionNumber : undefined;
        } else {
            const args = argsOrState as HostedConfigurationVersionArgs | undefined;
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            if ((!args || args.configurationProfileId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configurationProfileId'");
            }
            if ((!args || args.content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'content'");
            }
            if ((!args || args.contentType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contentType'");
            }
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["configurationProfileId"] = args ? args.configurationProfileId : undefined;
            resourceInputs["content"] = args?.content ? pulumi.secret(args.content) : undefined;
            resourceInputs["contentType"] = args ? args.contentType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["versionNumber"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["content"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(HostedConfigurationVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HostedConfigurationVersion resources.
 */
export interface HostedConfigurationVersionState {
    /**
     * Application ID.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * ARN of the AppConfig  hosted configuration version.
     */
    arn?: pulumi.Input<string>;
    /**
     * Configuration profile ID.
     */
    configurationProfileId?: pulumi.Input<string>;
    /**
     * Content of the configuration or the configuration data.
     */
    content?: pulumi.Input<string>;
    /**
     * Standard MIME type describing the format of the configuration content. For more information, see [Content-Type](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
     */
    contentType?: pulumi.Input<string>;
    /**
     * Description of the configuration.
     */
    description?: pulumi.Input<string>;
    /**
     * Version number of the hosted configuration.
     */
    versionNumber?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a HostedConfigurationVersion resource.
 */
export interface HostedConfigurationVersionArgs {
    /**
     * Application ID.
     */
    applicationId: pulumi.Input<string>;
    /**
     * Configuration profile ID.
     */
    configurationProfileId: pulumi.Input<string>;
    /**
     * Content of the configuration or the configuration data.
     */
    content: pulumi.Input<string>;
    /**
     * Standard MIME type describing the format of the configuration content. For more information, see [Content-Type](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
     */
    contentType: pulumi.Input<string>;
    /**
     * Description of the configuration.
     */
    description?: pulumi.Input<string>;
}
