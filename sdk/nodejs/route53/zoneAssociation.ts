// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Route53 Hosted Zone VPC association. VPC associations can only be made on private zones. See the `aws.route53.VpcAssociationAuthorization` resource for setting up cross-account associations.
 *
 * > **NOTE:** Unless explicit association ordering is required (e.g., a separate cross-account association authorization), usage of this resource is not recommended. Use the `vpc` configuration blocks available within the `aws.route53.Zone` resource instead.
 *
 * > **NOTE:** This provider provides both this standalone Zone VPC Association resource and exclusive VPC associations defined in-line in the `aws.route53.Zone` resource via `vpc` configuration blocks. At this time, you cannot use those in-line VPC associations in conjunction with this resource and the same zone ID otherwise it will cause a perpetual difference in plan output. You can optionally use [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) in the `aws.route53.Zone` resource to manage additional associations via this resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const primary = new aws.ec2.Vpc("primary", {
 *     cidrBlock: "10.6.0.0/16",
 *     enableDnsHostnames: true,
 *     enableDnsSupport: true,
 * });
 * const secondaryVpc = new aws.ec2.Vpc("secondaryVpc", {
 *     cidrBlock: "10.7.0.0/16",
 *     enableDnsHostnames: true,
 *     enableDnsSupport: true,
 * });
 * const example = new aws.route53.Zone("example", {vpcs: [{
 *     vpcId: primary.id,
 * }]});
 * const secondaryZoneAssociation = new aws.route53.ZoneAssociation("secondaryZoneAssociation", {
 *     zoneId: example.zoneId,
 *     vpcId: secondaryVpc.id,
 * });
 * ```
 *
 * ## Import
 *
 * The VPC is in the same region where you have configured the TODO AWS Providerterraform import {
 *
 *  to = aws_route53_zone_association.example
 *
 *  id = "Z123456ABCDEFG:vpc-12345678" } The VPC is _not_ in the same region where you have configured the TODO AWS Providerterraform import {
 *
 *  to = aws_route53_zone_association.example
 *
 *  id = "Z123456ABCDEFG:vpc-12345678:us-east-2" } **Using `pulumi import` to import** Route 53 Hosted Zone Associations using the Hosted Zone ID and VPC ID, separated by a colon (`:`). For exampleThe VPC is in the same region where you have configured the TODO AWS Providerconsole % pulumi import aws_route53_zone_association.example Z123456ABCDEFG:vpc-12345678 The VPC is _not_ in the same region where you have configured the TODO AWS Providerconsole % pulumi import aws_route53_zone_association.example Z123456ABCDEFG:vpc-12345678:us-east-2
 */
export class ZoneAssociation extends pulumi.CustomResource {
    /**
     * Get an existing ZoneAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZoneAssociationState, opts?: pulumi.CustomResourceOptions): ZoneAssociation {
        return new ZoneAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:route53/zoneAssociation:ZoneAssociation';

    /**
     * Returns true if the given object is an instance of ZoneAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ZoneAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ZoneAssociation.__pulumiType;
    }

    /**
     * The account ID of the account that created the hosted zone.
     */
    public /*out*/ readonly owningAccount!: pulumi.Output<string>;
    /**
     * The VPC to associate with the private hosted zone.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The VPC's region. Defaults to the region of the AWS provider.
     */
    public readonly vpcRegion!: pulumi.Output<string>;
    /**
     * The private hosted zone to associate.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a ZoneAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ZoneAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZoneAssociationArgs | ZoneAssociationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ZoneAssociationState | undefined;
            resourceInputs["owningAccount"] = state ? state.owningAccount : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vpcRegion"] = state ? state.vpcRegion : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as ZoneAssociationArgs | undefined;
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["vpcRegion"] = args ? args.vpcRegion : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["owningAccount"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ZoneAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ZoneAssociation resources.
 */
export interface ZoneAssociationState {
    /**
     * The account ID of the account that created the hosted zone.
     */
    owningAccount?: pulumi.Input<string>;
    /**
     * The VPC to associate with the private hosted zone.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The VPC's region. Defaults to the region of the AWS provider.
     */
    vpcRegion?: pulumi.Input<string>;
    /**
     * The private hosted zone to associate.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ZoneAssociation resource.
 */
export interface ZoneAssociationArgs {
    /**
     * The VPC to associate with the private hosted zone.
     */
    vpcId: pulumi.Input<string>;
    /**
     * The VPC's region. Defaults to the region of the AWS provider.
     */
    vpcRegion?: pulumi.Input<string>;
    /**
     * The private hosted zone to associate.
     */
    zoneId: pulumi.Input<string>;
}
