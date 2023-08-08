// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an Amazon MSK Connect Worker Configuration Resource.
 *
 * ## Example Usage
 * ### Basic configuration
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.mskconnect.WorkerConfiguration("example", {propertiesFileContent: `key.converter=org.apache.kafka.connect.storage.StringConverter
 * value.converter=org.apache.kafka.connect.storage.StringConverter
 *
 * `});
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_mskconnect_worker_configuration.example
 *
 *  id = "arn:aws:kafkaconnect:eu-central-1:123456789012:worker-configuration/example/8848493b-7fcc-478c-a646-4a52634e3378-4" } Using `pulumi import`, import MSK Connect Worker Configuration using the plugin's `arn`. For exampleconsole % pulumi import aws_mskconnect_worker_configuration.example 'arn:aws:kafkaconnect:eu-central-1:123456789012:worker-configuration/example/8848493b-7fcc-478c-a646-4a52634e3378-4'
 */
export class WorkerConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing WorkerConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkerConfigurationState, opts?: pulumi.CustomResourceOptions): WorkerConfiguration {
        return new WorkerConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:mskconnect/workerConfiguration:WorkerConfiguration';

    /**
     * Returns true if the given object is an instance of WorkerConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkerConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkerConfiguration.__pulumiType;
    }

    /**
     * the Amazon Resource Name (ARN) of the worker configuration.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A summary description of the worker configuration.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * an ID of the latest successfully created revision of the worker configuration.
     */
    public /*out*/ readonly latestRevision!: pulumi.Output<number>;
    /**
     * The name of the worker configuration.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Contents of connect-distributed.properties file. The value can be either base64 encoded or in raw format.
     *
     * The following arguments are optional:
     */
    public readonly propertiesFileContent!: pulumi.Output<string>;

    /**
     * Create a WorkerConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkerConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkerConfigurationArgs | WorkerConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkerConfigurationState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["latestRevision"] = state ? state.latestRevision : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["propertiesFileContent"] = state ? state.propertiesFileContent : undefined;
        } else {
            const args = argsOrState as WorkerConfigurationArgs | undefined;
            if ((!args || args.propertiesFileContent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'propertiesFileContent'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["propertiesFileContent"] = args ? args.propertiesFileContent : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["latestRevision"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WorkerConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WorkerConfiguration resources.
 */
export interface WorkerConfigurationState {
    /**
     * the Amazon Resource Name (ARN) of the worker configuration.
     */
    arn?: pulumi.Input<string>;
    /**
     * A summary description of the worker configuration.
     */
    description?: pulumi.Input<string>;
    /**
     * an ID of the latest successfully created revision of the worker configuration.
     */
    latestRevision?: pulumi.Input<number>;
    /**
     * The name of the worker configuration.
     */
    name?: pulumi.Input<string>;
    /**
     * Contents of connect-distributed.properties file. The value can be either base64 encoded or in raw format.
     *
     * The following arguments are optional:
     */
    propertiesFileContent?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WorkerConfiguration resource.
 */
export interface WorkerConfigurationArgs {
    /**
     * A summary description of the worker configuration.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the worker configuration.
     */
    name?: pulumi.Input<string>;
    /**
     * Contents of connect-distributed.properties file. The value can be either base64 encoded or in raw format.
     *
     * The following arguments are optional:
     */
    propertiesFileContent: pulumi.Input<string>;
}
