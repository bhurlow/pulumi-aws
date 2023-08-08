// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a CodeBuild Report Groups Resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const current = aws.getCallerIdentity({});
 * const examplePolicyDocument = current.then(current => aws.iam.getPolicyDocument({
 *     statements: [{
 *         sid: "Enable IAM User Permissions",
 *         effect: "Allow",
 *         principals: [{
 *             type: "AWS",
 *             identifiers: [`arn:aws:iam::${current.accountId}:root`],
 *         }],
 *         actions: ["kms:*"],
 *         resources: ["*"],
 *     }],
 * }));
 * const exampleKey = new aws.kms.Key("exampleKey", {
 *     description: "my test kms key",
 *     deletionWindowInDays: 7,
 *     policy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json),
 * });
 * const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
 * const exampleReportGroup = new aws.codebuild.ReportGroup("exampleReportGroup", {
 *     type: "TEST",
 *     exportConfig: {
 *         type: "S3",
 *         s3Destination: {
 *             bucket: exampleBucketV2.id,
 *             encryptionDisabled: false,
 *             encryptionKey: exampleKey.arn,
 *             packaging: "NONE",
 *             path: "/some",
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_codebuild_report_group.example
 *
 *  id = "arn:aws:codebuild:us-west-2:123456789:report-group/report-group-name" } Using `pulumi import`, import CodeBuild Report Group using the CodeBuild Report Group arn. For exampleconsole % pulumi import aws_codebuild_report_group.example arn:aws:codebuild:us-west-2:123456789:report-group/report-group-name
 */
export class ReportGroup extends pulumi.CustomResource {
    /**
     * Get an existing ReportGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReportGroupState, opts?: pulumi.CustomResourceOptions): ReportGroup {
        return new ReportGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:codebuild/reportGroup:ReportGroup';

    /**
     * Returns true if the given object is an instance of ReportGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReportGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReportGroup.__pulumiType;
    }

    /**
     * The ARN of Report Group.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The date and time this Report Group was created.
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * If `true`, deletes any reports that belong to a report group before deleting the report group. If `false`, you must delete any reports in the report group before deleting it. Default value is `false`.
     */
    public readonly deleteReports!: pulumi.Output<boolean | undefined>;
    /**
     * Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
     */
    public readonly exportConfig!: pulumi.Output<outputs.codebuild.ReportGroupExportConfig>;
    /**
     * The name of a Report Group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Key-value mapping of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The type of the Report Group. Valid value are `TEST` and `CODE_COVERAGE`.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a ReportGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReportGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReportGroupArgs | ReportGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReportGroupState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["created"] = state ? state.created : undefined;
            resourceInputs["deleteReports"] = state ? state.deleteReports : undefined;
            resourceInputs["exportConfig"] = state ? state.exportConfig : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ReportGroupArgs | undefined;
            if ((!args || args.exportConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'exportConfig'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["deleteReports"] = args ? args.deleteReports : undefined;
            resourceInputs["exportConfig"] = args ? args.exportConfig : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["created"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReportGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReportGroup resources.
 */
export interface ReportGroupState {
    /**
     * The ARN of Report Group.
     */
    arn?: pulumi.Input<string>;
    /**
     * The date and time this Report Group was created.
     */
    created?: pulumi.Input<string>;
    /**
     * If `true`, deletes any reports that belong to a report group before deleting the report group. If `false`, you must delete any reports in the report group before deleting it. Default value is `false`.
     */
    deleteReports?: pulumi.Input<boolean>;
    /**
     * Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
     */
    exportConfig?: pulumi.Input<inputs.codebuild.ReportGroupExportConfig>;
    /**
     * The name of a Report Group.
     */
    name?: pulumi.Input<string>;
    /**
     * Key-value mapping of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of the Report Group. Valid value are `TEST` and `CODE_COVERAGE`.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReportGroup resource.
 */
export interface ReportGroupArgs {
    /**
     * If `true`, deletes any reports that belong to a report group before deleting the report group. If `false`, you must delete any reports in the report group before deleting it. Default value is `false`.
     */
    deleteReports?: pulumi.Input<boolean>;
    /**
     * Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
     */
    exportConfig: pulumi.Input<inputs.codebuild.ReportGroupExportConfig>;
    /**
     * The name of a Report Group.
     */
    name?: pulumi.Input<string>;
    /**
     * Key-value mapping of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of the Report Group. Valid value are `TEST` and `CODE_COVERAGE`.
     */
    type: pulumi.Input<string>;
}
