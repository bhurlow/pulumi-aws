// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
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
 * Capacity Reservations can be imported using the `id`, e.g.
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
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CapacityReservationState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            inputs["ebsOptimized"] = state ? state.ebsOptimized : undefined;
            inputs["endDate"] = state ? state.endDate : undefined;
            inputs["endDateType"] = state ? state.endDateType : undefined;
            inputs["ephemeralStorage"] = state ? state.ephemeralStorage : undefined;
            inputs["instanceCount"] = state ? state.instanceCount : undefined;
            inputs["instanceMatchCriteria"] = state ? state.instanceMatchCriteria : undefined;
            inputs["instancePlatform"] = state ? state.instancePlatform : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tenancy"] = state ? state.tenancy : undefined;
        } else {
            const args = argsOrState as CapacityReservationArgs | undefined;
            if ((!args || args.availabilityZone === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'availabilityZone'");
            }
            if ((!args || args.instanceCount === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'instanceCount'");
            }
            if ((!args || args.instancePlatform === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'instancePlatform'");
            }
            if ((!args || args.instanceType === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'instanceType'");
            }
            inputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            inputs["ebsOptimized"] = args ? args.ebsOptimized : undefined;
            inputs["endDate"] = args ? args.endDate : undefined;
            inputs["endDateType"] = args ? args.endDateType : undefined;
            inputs["ephemeralStorage"] = args ? args.ephemeralStorage : undefined;
            inputs["instanceCount"] = args ? args.instanceCount : undefined;
            inputs["instanceMatchCriteria"] = args ? args.instanceMatchCriteria : undefined;
            inputs["instancePlatform"] = args ? args.instancePlatform : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tenancy"] = args ? args.tenancy : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CapacityReservation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CapacityReservation resources.
 */
export interface CapacityReservationState {
    /**
     * The ARN of the Capacity Reservation.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The Availability Zone in which to create the Capacity Reservation.
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * Indicates whether the Capacity Reservation supports EBS-optimized instances.
     */
    readonly ebsOptimized?: pulumi.Input<boolean>;
    /**
     * The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
     */
    readonly endDate?: pulumi.Input<string>;
    /**
     * Indicates the way in which the Capacity Reservation ends. Specify either `unlimited` or `limited`.
     */
    readonly endDateType?: pulumi.Input<string>;
    /**
     * Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
     */
    readonly ephemeralStorage?: pulumi.Input<boolean>;
    /**
     * The number of instances for which to reserve capacity.
     */
    readonly instanceCount?: pulumi.Input<number>;
    /**
     * Indicates the type of instance launches that the Capacity Reservation accepts. Specify either `open` or `targeted`.
     */
    readonly instanceMatchCriteria?: pulumi.Input<string>;
    /**
     * The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
     */
    readonly instancePlatform?: pulumi.Input<string | enums.ec2.InstancePlatform>;
    /**
     * The instance type for which to reserve capacity.
     */
    readonly instanceType?: pulumi.Input<string | enums.ec2.InstanceType>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
     */
    readonly tenancy?: pulumi.Input<string | enums.ec2.Tenancy>;
}

/**
 * The set of arguments for constructing a CapacityReservation resource.
 */
export interface CapacityReservationArgs {
    /**
     * The Availability Zone in which to create the Capacity Reservation.
     */
    readonly availabilityZone: pulumi.Input<string>;
    /**
     * Indicates whether the Capacity Reservation supports EBS-optimized instances.
     */
    readonly ebsOptimized?: pulumi.Input<boolean>;
    /**
     * The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
     */
    readonly endDate?: pulumi.Input<string>;
    /**
     * Indicates the way in which the Capacity Reservation ends. Specify either `unlimited` or `limited`.
     */
    readonly endDateType?: pulumi.Input<string>;
    /**
     * Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
     */
    readonly ephemeralStorage?: pulumi.Input<boolean>;
    /**
     * The number of instances for which to reserve capacity.
     */
    readonly instanceCount: pulumi.Input<number>;
    /**
     * Indicates the type of instance launches that the Capacity Reservation accepts. Specify either `open` or `targeted`.
     */
    readonly instanceMatchCriteria?: pulumi.Input<string>;
    /**
     * The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
     */
    readonly instancePlatform: pulumi.Input<string | enums.ec2.InstancePlatform>;
    /**
     * The instance type for which to reserve capacity.
     */
    readonly instanceType: pulumi.Input<string | enums.ec2.InstanceType>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
     */
    readonly tenancy?: pulumi.Input<string | enums.ec2.Tenancy>;
}
