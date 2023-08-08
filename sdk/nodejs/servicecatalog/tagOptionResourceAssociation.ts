// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Service Catalog Tag Option Resource Association.
 *
 * > **Tip:** A "resource" is either a Service Catalog portfolio or product.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.servicecatalog.TagOptionResourceAssociation("example", {
 *     resourceId: "prod-dnigbtea24ste",
 *     tagOptionId: "tag-pjtvyakdlyo3m",
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_servicecatalog_tag_option_resource_association.example
 *
 *  id = "tag-pjtvyakdlyo3m:prod-dnigbtea24ste" } Using `pulumi import`, import `aws_servicecatalog_tag_option_resource_association` using the tag option ID and resource ID. For exampleconsole % pulumi import aws_servicecatalog_tag_option_resource_association.example tag-pjtvyakdlyo3m:prod-dnigbtea24ste
 */
export class TagOptionResourceAssociation extends pulumi.CustomResource {
    /**
     * Get an existing TagOptionResourceAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TagOptionResourceAssociationState, opts?: pulumi.CustomResourceOptions): TagOptionResourceAssociation {
        return new TagOptionResourceAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:servicecatalog/tagOptionResourceAssociation:TagOptionResourceAssociation';

    /**
     * Returns true if the given object is an instance of TagOptionResourceAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TagOptionResourceAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TagOptionResourceAssociation.__pulumiType;
    }

    /**
     * ARN of the resource.
     */
    public /*out*/ readonly resourceArn!: pulumi.Output<string>;
    /**
     * Creation time of the resource.
     */
    public /*out*/ readonly resourceCreatedTime!: pulumi.Output<string>;
    /**
     * Description of the resource.
     */
    public /*out*/ readonly resourceDescription!: pulumi.Output<string>;
    /**
     * Resource identifier.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * Description of the resource.
     */
    public /*out*/ readonly resourceName!: pulumi.Output<string>;
    /**
     * Tag Option identifier.
     */
    public readonly tagOptionId!: pulumi.Output<string>;

    /**
     * Create a TagOptionResourceAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TagOptionResourceAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TagOptionResourceAssociationArgs | TagOptionResourceAssociationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TagOptionResourceAssociationState | undefined;
            resourceInputs["resourceArn"] = state ? state.resourceArn : undefined;
            resourceInputs["resourceCreatedTime"] = state ? state.resourceCreatedTime : undefined;
            resourceInputs["resourceDescription"] = state ? state.resourceDescription : undefined;
            resourceInputs["resourceId"] = state ? state.resourceId : undefined;
            resourceInputs["resourceName"] = state ? state.resourceName : undefined;
            resourceInputs["tagOptionId"] = state ? state.tagOptionId : undefined;
        } else {
            const args = argsOrState as TagOptionResourceAssociationArgs | undefined;
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            if ((!args || args.tagOptionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tagOptionId'");
            }
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["tagOptionId"] = args ? args.tagOptionId : undefined;
            resourceInputs["resourceArn"] = undefined /*out*/;
            resourceInputs["resourceCreatedTime"] = undefined /*out*/;
            resourceInputs["resourceDescription"] = undefined /*out*/;
            resourceInputs["resourceName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TagOptionResourceAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TagOptionResourceAssociation resources.
 */
export interface TagOptionResourceAssociationState {
    /**
     * ARN of the resource.
     */
    resourceArn?: pulumi.Input<string>;
    /**
     * Creation time of the resource.
     */
    resourceCreatedTime?: pulumi.Input<string>;
    /**
     * Description of the resource.
     */
    resourceDescription?: pulumi.Input<string>;
    /**
     * Resource identifier.
     */
    resourceId?: pulumi.Input<string>;
    /**
     * Description of the resource.
     */
    resourceName?: pulumi.Input<string>;
    /**
     * Tag Option identifier.
     */
    tagOptionId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TagOptionResourceAssociation resource.
 */
export interface TagOptionResourceAssociationArgs {
    /**
     * Resource identifier.
     */
    resourceId: pulumi.Input<string>;
    /**
     * Tag Option identifier.
     */
    tagOptionId: pulumi.Input<string>;
}
