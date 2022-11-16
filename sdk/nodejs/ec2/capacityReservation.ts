// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides an EC2 Capacity Reservation. This allows you to reserve capacity for your Amazon EC2 instances in a specific Availability Zone for any duration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const defaultCapacityReservation = new aws.ec2.CapacityReservation("default", {
 *     availabilityZone: "eu-west-1a",
 *     instanceCount: 1,
 *     instancePlatform: "Linux/UNIX",
 *     instanceType: "t2.micro",
 * });
 * ```
 *
 * ## Import
 *
 * Capacity Reservations can be imported using the `id`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:ec2/capacityReservation:CapacityReservation web cr-0123456789abcdef0
 * ```
 */
export class CapacityReservation extends pulumi.CustomResource {
    /**
     * Get an existing CapacityReservation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CapacityReservationState, opts?: pulumi.CustomResourceOptions): CapacityReservation {
        return new CapacityReservation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/capacityReservation:CapacityReservation';

    /**
     * Returns true if the given object is an instance of CapacityReservation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CapacityReservation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CapacityReservation.__pulumiType;
    }

    /**
     * The ARN of the Capacity Reservation.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The Availability Zone in which to create the Capacity Reservation.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * Indicates whether the Capacity Reservation supports EBS-optimized instances.
     */
    public readonly ebsOptimized!: pulumi.Output<boolean | undefined>;
    /**
     * The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
     */
    public readonly endDate!: pulumi.Output<string | undefined>;
    /**
     * Indicates the way in which the Capacity Reservation ends. Specify either `unlimited` or `limited`.
     */
    public readonly endDateType!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
     */
    public readonly ephemeralStorage!: pulumi.Output<boolean | undefined>;
    /**
     * The number of instances for which to reserve capacity.
     */
    public readonly instanceCount!: pulumi.Output<number>;
    /**
     * Indicates the type of instance launches that the Capacity Reservation accepts. Specify either `open` or `targeted`.
     */
    public readonly instanceMatchCriteria!: pulumi.Output<string | undefined>;
    /**
     * The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
     */
    public readonly instancePlatform!: pulumi.Output<string>;
    /**
     * The instance type for which to reserve capacity.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the Outpost on which to create the Capacity Reservation.
     */
    public readonly outpostArn!: pulumi.Output<string | undefined>;
    /**
     * The ID of the AWS account that owns the Capacity Reservation.
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the cluster placement group in which to create the Capacity Reservation.
     */
    public readonly placementGroupArn!: pulumi.Output<string | undefined>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
     */
    public readonly tenancy!: pulumi.Output<string | undefined>;

    /**
     * Create a CapacityReservation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CapacityReservationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CapacityReservationArgs | CapacityReservationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CapacityReservationState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["ebsOptimized"] = state ? state.ebsOptimized : undefined;
            resourceInputs["endDate"] = state ? state.endDate : undefined;
            resourceInputs["endDateType"] = state ? state.endDateType : undefined;
            resourceInputs["ephemeralStorage"] = state ? state.ephemeralStorage : undefined;
            resourceInputs["instanceCount"] = state ? state.instanceCount : undefined;
            resourceInputs["instanceMatchCriteria"] = state ? state.instanceMatchCriteria : undefined;
            resourceInputs["instancePlatform"] = state ? state.instancePlatform : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["outpostArn"] = state ? state.outpostArn : undefined;
            resourceInputs["ownerId"] = state ? state.ownerId : undefined;
            resourceInputs["placementGroupArn"] = state ? state.placementGroupArn : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["tenancy"] = state ? state.tenancy : undefined;
        } else {
            const args = argsOrState as CapacityReservationArgs | undefined;
            if ((!args || args.availabilityZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availabilityZone'");
            }
            if ((!args || args.instanceCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceCount'");
            }
            if ((!args || args.instancePlatform === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instancePlatform'");
            }
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["ebsOptimized"] = args ? args.ebsOptimized : undefined;
            resourceInputs["endDate"] = args ? args.endDate : undefined;
            resourceInputs["endDateType"] = args ? args.endDateType : undefined;
            resourceInputs["ephemeralStorage"] = args ? args.ephemeralStorage : undefined;
            resourceInputs["instanceCount"] = args ? args.instanceCount : undefined;
            resourceInputs["instanceMatchCriteria"] = args ? args.instanceMatchCriteria : undefined;
            resourceInputs["instancePlatform"] = args ? args.instancePlatform : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["outpostArn"] = args ? args.outpostArn : undefined;
            resourceInputs["placementGroupArn"] = args ? args.placementGroupArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tenancy"] = args ? args.tenancy : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["ownerId"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CapacityReservation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CapacityReservation resources.
 */
export interface CapacityReservationState {
    /**
     * The ARN of the Capacity Reservation.
     */
    arn?: pulumi.Input<string>;
    /**
     * The Availability Zone in which to create the Capacity Reservation.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Indicates whether the Capacity Reservation supports EBS-optimized instances.
     */
    ebsOptimized?: pulumi.Input<boolean>;
    /**
     * The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
     */
    endDate?: pulumi.Input<string>;
    /**
     * Indicates the way in which the Capacity Reservation ends. Specify either `unlimited` or `limited`.
     */
    endDateType?: pulumi.Input<string>;
    /**
     * Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
     */
    ephemeralStorage?: pulumi.Input<boolean>;
    /**
     * The number of instances for which to reserve capacity.
     */
    instanceCount?: pulumi.Input<number>;
    /**
     * Indicates the type of instance launches that the Capacity Reservation accepts. Specify either `open` or `targeted`.
     */
    instanceMatchCriteria?: pulumi.Input<string>;
    /**
     * The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
     */
    instancePlatform?: pulumi.Input<string | enums.ec2.InstancePlatform>;
    /**
     * The instance type for which to reserve capacity.
     */
    instanceType?: pulumi.Input<string | enums.ec2.InstanceType>;
    /**
     * The Amazon Resource Name (ARN) of the Outpost on which to create the Capacity Reservation.
     */
    outpostArn?: pulumi.Input<string>;
    /**
     * The ID of the AWS account that owns the Capacity Reservation.
     */
    ownerId?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the cluster placement group in which to create the Capacity Reservation.
     */
    placementGroupArn?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
     */
    tenancy?: pulumi.Input<string | enums.ec2.Tenancy>;
}

/**
 * The set of arguments for constructing a CapacityReservation resource.
 */
export interface CapacityReservationArgs {
    /**
     * The Availability Zone in which to create the Capacity Reservation.
     */
    availabilityZone: pulumi.Input<string>;
    /**
     * Indicates whether the Capacity Reservation supports EBS-optimized instances.
     */
    ebsOptimized?: pulumi.Input<boolean>;
    /**
     * The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
     */
    endDate?: pulumi.Input<string>;
    /**
     * Indicates the way in which the Capacity Reservation ends. Specify either `unlimited` or `limited`.
     */
    endDateType?: pulumi.Input<string>;
    /**
     * Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
     */
    ephemeralStorage?: pulumi.Input<boolean>;
    /**
     * The number of instances for which to reserve capacity.
     */
    instanceCount: pulumi.Input<number>;
    /**
     * Indicates the type of instance launches that the Capacity Reservation accepts. Specify either `open` or `targeted`.
     */
    instanceMatchCriteria?: pulumi.Input<string>;
    /**
     * The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
     */
    instancePlatform: pulumi.Input<string | enums.ec2.InstancePlatform>;
    /**
     * The instance type for which to reserve capacity.
     */
    instanceType: pulumi.Input<string | enums.ec2.InstanceType>;
    /**
     * The Amazon Resource Name (ARN) of the Outpost on which to create the Capacity Reservation.
     */
    outpostArn?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the cluster placement group in which to create the Capacity Reservation.
     */
    placementGroupArn?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
     */
    tenancy?: pulumi.Input<string | enums.ec2.Tenancy>;
}
