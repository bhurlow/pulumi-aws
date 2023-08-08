// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage an [Amazon Macie Classification Export Configuration](https://docs.aws.amazon.com/macie/latest/APIReference/classification-export-configuration.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleAccount = new aws.macie2.Account("exampleAccount", {});
 * const exampleClassificationExportConfiguration = new aws.macie2.ClassificationExportConfiguration("exampleClassificationExportConfiguration", {s3Destination: {
 *     bucketName: aws_s3_bucket.example.bucket,
 *     keyPrefix: "exampleprefix/",
 *     kmsKeyArn: aws_kms_key.example.arn,
 * }}, {
 *     dependsOn: [exampleAccount],
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_macie2_classification_export_configuration.example
 *
 *  id = "123456789012:us-west-2" } Using `pulumi import`, import `aws_macie2_classification_export_configuration` using the account ID and region. For exampleconsole % pulumi import aws_macie2_classification_export_configuration.example 123456789012:us-west-2
 */
export class ClassificationExportConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing ClassificationExportConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClassificationExportConfigurationState, opts?: pulumi.CustomResourceOptions): ClassificationExportConfiguration {
        return new ClassificationExportConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:macie2/classificationExportConfiguration:ClassificationExportConfiguration';

    /**
     * Returns true if the given object is an instance of ClassificationExportConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClassificationExportConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClassificationExportConfiguration.__pulumiType;
    }

    /**
     * Configuration block for a S3 Destination. Defined below
     */
    public readonly s3Destination!: pulumi.Output<outputs.macie2.ClassificationExportConfigurationS3Destination | undefined>;

    /**
     * Create a ClassificationExportConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ClassificationExportConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClassificationExportConfigurationArgs | ClassificationExportConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClassificationExportConfigurationState | undefined;
            resourceInputs["s3Destination"] = state ? state.s3Destination : undefined;
        } else {
            const args = argsOrState as ClassificationExportConfigurationArgs | undefined;
            resourceInputs["s3Destination"] = args ? args.s3Destination : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClassificationExportConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClassificationExportConfiguration resources.
 */
export interface ClassificationExportConfigurationState {
    /**
     * Configuration block for a S3 Destination. Defined below
     */
    s3Destination?: pulumi.Input<inputs.macie2.ClassificationExportConfigurationS3Destination>;
}

/**
 * The set of arguments for constructing a ClassificationExportConfiguration resource.
 */
export interface ClassificationExportConfigurationArgs {
    /**
     * Configuration block for a S3 Destination. Defined below
     */
    s3Destination?: pulumi.Input<inputs.macie2.ClassificationExportConfigurationS3Destination>;
}
