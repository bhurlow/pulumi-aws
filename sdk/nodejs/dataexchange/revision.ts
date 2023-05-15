// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage AWS Data Exchange Revisions.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.dataexchange.Revision("example", {dataSetId: aws_dataexchange_data_set.example.id});
 * ```
 *
 * ## Import
 *
 * DataExchange Revisions can be imported by their `data-set-id:revision-id`
 *
 * ```sh
 *  $ pulumi import aws:dataexchange/revision:Revision example 4fa784c7-ccb4-4dbf-ba4f-02198320daa1:4fa784c7-ccb4-4dbf-ba4f-02198320daa1
 * ```
 */
export class Revision extends pulumi.CustomResource {
    /**
     * Get an existing Revision resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RevisionState, opts?: pulumi.CustomResourceOptions): Revision {
        return new Revision(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:dataexchange/revision:Revision';

    /**
     * Returns true if the given object is an instance of Revision.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Revision {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Revision.__pulumiType;
    }

    /**
     * The Amazon Resource Name of this data set.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * An optional comment about the revision.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * The dataset id.
     */
    public readonly dataSetId!: pulumi.Output<string>;
    /**
     * The Id of the revision.
     */
    public /*out*/ readonly revisionId!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Revision resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RevisionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RevisionArgs | RevisionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RevisionState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dataSetId"] = state ? state.dataSetId : undefined;
            resourceInputs["revisionId"] = state ? state.revisionId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as RevisionArgs | undefined;
            if ((!args || args.dataSetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataSetId'");
            }
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dataSetId"] = args ? args.dataSetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["revisionId"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Revision.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Revision resources.
 */
export interface RevisionState {
    /**
     * The Amazon Resource Name of this data set.
     */
    arn?: pulumi.Input<string>;
    /**
     * An optional comment about the revision.
     */
    comment?: pulumi.Input<string>;
    /**
     * The dataset id.
     */
    dataSetId?: pulumi.Input<string>;
    /**
     * The Id of the revision.
     */
    revisionId?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Revision resource.
 */
export interface RevisionArgs {
    /**
     * An optional comment about the revision.
     */
    comment?: pulumi.Input<string>;
    /**
     * The dataset id.
     */
    dataSetId: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
