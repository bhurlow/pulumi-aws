// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2.CapacityReservationArgs;
import com.pulumi.aws.ec2.inputs.CapacityReservationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an EC2 Capacity Reservation. This allows you to reserve capacity for your Amazon EC2 instances in a specific Availability Zone for any duration.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.CapacityReservation;
 * import com.pulumi.aws.ec2.CapacityReservationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var default_ = new CapacityReservation(&#34;default&#34;, CapacityReservationArgs.builder()        
 *             .availabilityZone(&#34;eu-west-1a&#34;)
 *             .instanceCount(1)
 *             .instancePlatform(&#34;Linux/UNIX&#34;)
 *             .instanceType(&#34;t2.micro&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Capacity Reservations can be imported using the `id`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:ec2/capacityReservation:CapacityReservation web cr-0123456789abcdef0
 * ```
 * 
 */
@ResourceType(type="aws:ec2/capacityReservation:CapacityReservation")
public class CapacityReservation extends com.pulumi.resources.CustomResource {
    /**
     * The ARN of the Capacity Reservation.
     * 
     */
    @Export(name="arn", type=String.class, parameters={})
    private Output<String> arn;

    /**
     * @return The ARN of the Capacity Reservation.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The Availability Zone in which to create the Capacity Reservation.
     * 
     */
    @Export(name="availabilityZone", type=String.class, parameters={})
    private Output<String> availabilityZone;

    /**
     * @return The Availability Zone in which to create the Capacity Reservation.
     * 
     */
    public Output<String> availabilityZone() {
        return this.availabilityZone;
    }
    /**
     * Indicates whether the Capacity Reservation supports EBS-optimized instances.
     * 
     */
    @Export(name="ebsOptimized", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> ebsOptimized;

    /**
     * @return Indicates whether the Capacity Reservation supports EBS-optimized instances.
     * 
     */
    public Output<Optional<Boolean>> ebsOptimized() {
        return Codegen.optional(this.ebsOptimized);
    }
    /**
     * The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
     * 
     */
    @Export(name="endDate", type=String.class, parameters={})
    private Output</* @Nullable */ String> endDate;

    /**
     * @return The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
     * 
     */
    public Output<Optional<String>> endDate() {
        return Codegen.optional(this.endDate);
    }
    /**
     * Indicates the way in which the Capacity Reservation ends. Specify either `unlimited` or `limited`.
     * 
     */
    @Export(name="endDateType", type=String.class, parameters={})
    private Output</* @Nullable */ String> endDateType;

    /**
     * @return Indicates the way in which the Capacity Reservation ends. Specify either `unlimited` or `limited`.
     * 
     */
    public Output<Optional<String>> endDateType() {
        return Codegen.optional(this.endDateType);
    }
    /**
     * Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
     * 
     */
    @Export(name="ephemeralStorage", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> ephemeralStorage;

    /**
     * @return Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
     * 
     */
    public Output<Optional<Boolean>> ephemeralStorage() {
        return Codegen.optional(this.ephemeralStorage);
    }
    /**
     * The number of instances for which to reserve capacity.
     * 
     */
    @Export(name="instanceCount", type=Integer.class, parameters={})
    private Output<Integer> instanceCount;

    /**
     * @return The number of instances for which to reserve capacity.
     * 
     */
    public Output<Integer> instanceCount() {
        return this.instanceCount;
    }
    /**
     * Indicates the type of instance launches that the Capacity Reservation accepts. Specify either `open` or `targeted`.
     * 
     */
    @Export(name="instanceMatchCriteria", type=String.class, parameters={})
    private Output</* @Nullable */ String> instanceMatchCriteria;

    /**
     * @return Indicates the type of instance launches that the Capacity Reservation accepts. Specify either `open` or `targeted`.
     * 
     */
    public Output<Optional<String>> instanceMatchCriteria() {
        return Codegen.optional(this.instanceMatchCriteria);
    }
    /**
     * The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
     * 
     */
    @Export(name="instancePlatform", type=String.class, parameters={})
    private Output<String> instancePlatform;

    /**
     * @return The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
     * 
     */
    public Output<String> instancePlatform() {
        return this.instancePlatform;
    }
    /**
     * The instance type for which to reserve capacity.
     * 
     */
    @Export(name="instanceType", type=String.class, parameters={})
    private Output<String> instanceType;

    /**
     * @return The instance type for which to reserve capacity.
     * 
     */
    public Output<String> instanceType() {
        return this.instanceType;
    }
    /**
     * The Amazon Resource Name (ARN) of the Outpost on which to create the Capacity Reservation.
     * 
     */
    @Export(name="outpostArn", type=String.class, parameters={})
    private Output</* @Nullable */ String> outpostArn;

    /**
     * @return The Amazon Resource Name (ARN) of the Outpost on which to create the Capacity Reservation.
     * 
     */
    public Output<Optional<String>> outpostArn() {
        return Codegen.optional(this.outpostArn);
    }
    /**
     * The ID of the AWS account that owns the Capacity Reservation.
     * 
     */
    @Export(name="ownerId", type=String.class, parameters={})
    private Output<String> ownerId;

    /**
     * @return The ID of the AWS account that owns the Capacity Reservation.
     * 
     */
    public Output<String> ownerId() {
        return this.ownerId;
    }
    /**
     * The Amazon Resource Name (ARN) of the cluster placement group in which to create the Capacity Reservation.
     * 
     */
    @Export(name="placementGroupArn", type=String.class, parameters={})
    private Output</* @Nullable */ String> placementGroupArn;

    /**
     * @return The Amazon Resource Name (ARN) of the cluster placement group in which to create the Capacity Reservation.
     * 
     */
    public Output<Optional<String>> placementGroupArn() {
        return Codegen.optional(this.placementGroupArn);
    }
    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block
     * 
     */
    @Export(name="tagsAll", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    /**
     * Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
     * 
     */
    @Export(name="tenancy", type=String.class, parameters={})
    private Output</* @Nullable */ String> tenancy;

    /**
     * @return Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
     * 
     */
    public Output<Optional<String>> tenancy() {
        return Codegen.optional(this.tenancy);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CapacityReservation(String name) {
        this(name, CapacityReservationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CapacityReservation(String name, CapacityReservationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CapacityReservation(String name, CapacityReservationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/capacityReservation:CapacityReservation", name, args == null ? CapacityReservationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CapacityReservation(String name, Output<String> id, @Nullable CapacityReservationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/capacityReservation:CapacityReservation", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static CapacityReservation get(String name, Output<String> id, @Nullable CapacityReservationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CapacityReservation(name, id, state, options);
    }
}
