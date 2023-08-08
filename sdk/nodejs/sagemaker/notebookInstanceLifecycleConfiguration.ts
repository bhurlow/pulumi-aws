// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a lifecycle configuration for SageMaker Notebook Instances.
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_sagemaker_notebook_instance_lifecycle_configuration.lc
 *
 *  id = "foo" } Using `pulumi import`, import models using the `name`. For exampleconsole % pulumi import aws_sagemaker_notebook_instance_lifecycle_configuration.lc foo
 */
export class NotebookInstanceLifecycleConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing NotebookInstanceLifecycleConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NotebookInstanceLifecycleConfigurationState, opts?: pulumi.CustomResourceOptions): NotebookInstanceLifecycleConfiguration {
        return new NotebookInstanceLifecycleConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:sagemaker/notebookInstanceLifecycleConfiguration:NotebookInstanceLifecycleConfiguration';

    /**
     * Returns true if the given object is an instance of NotebookInstanceLifecycleConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NotebookInstanceLifecycleConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NotebookInstanceLifecycleConfiguration.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this lifecycle configuration.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
     */
    public readonly onCreate!: pulumi.Output<string | undefined>;
    /**
     * A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
     */
    public readonly onStart!: pulumi.Output<string | undefined>;

    /**
     * Create a NotebookInstanceLifecycleConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NotebookInstanceLifecycleConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NotebookInstanceLifecycleConfigurationArgs | NotebookInstanceLifecycleConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NotebookInstanceLifecycleConfigurationState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["onCreate"] = state ? state.onCreate : undefined;
            resourceInputs["onStart"] = state ? state.onStart : undefined;
        } else {
            const args = argsOrState as NotebookInstanceLifecycleConfigurationArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["onCreate"] = args ? args.onCreate : undefined;
            resourceInputs["onStart"] = args ? args.onStart : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NotebookInstanceLifecycleConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NotebookInstanceLifecycleConfiguration resources.
 */
export interface NotebookInstanceLifecycleConfigurationState {
    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this lifecycle configuration.
     */
    arn?: pulumi.Input<string>;
    /**
     * The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
     */
    name?: pulumi.Input<string>;
    /**
     * A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
     */
    onCreate?: pulumi.Input<string>;
    /**
     * A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
     */
    onStart?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NotebookInstanceLifecycleConfiguration resource.
 */
export interface NotebookInstanceLifecycleConfigurationArgs {
    /**
     * The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
     */
    name?: pulumi.Input<string>;
    /**
     * A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
     */
    onCreate?: pulumi.Input<string>;
    /**
     * A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
     */
    onStart?: pulumi.Input<string>;
}
