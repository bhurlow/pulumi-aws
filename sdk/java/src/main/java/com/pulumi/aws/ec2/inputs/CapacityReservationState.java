// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2.inputs;

import com.pulumi.aws.ec2.enums.InstancePlatform;
import com.pulumi.aws.ec2.enums.InstanceType;
import com.pulumi.aws.ec2.enums.Tenancy;
import com.pulumi.core.Either;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CapacityReservationState extends com.pulumi.resources.ResourceArgs {

    public static final CapacityReservationState Empty = new CapacityReservationState();

    /**
     * The ARN of the Capacity Reservation.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The ARN of the Capacity Reservation.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The Availability Zone in which to create the Capacity Reservation.
     * 
     */
    @Import(name="availabilityZone")
    private @Nullable Output<String> availabilityZone;

    /**
     * @return The Availability Zone in which to create the Capacity Reservation.
     * 
     */
    public Optional<Output<String>> availabilityZone() {
        return Optional.ofNullable(this.availabilityZone);
    }

    /**
     * Indicates whether the Capacity Reservation supports EBS-optimized instances.
     * 
     */
    @Import(name="ebsOptimized")
    private @Nullable Output<Boolean> ebsOptimized;

    /**
     * @return Indicates whether the Capacity Reservation supports EBS-optimized instances.
     * 
     */
    public Optional<Output<Boolean>> ebsOptimized() {
        return Optional.ofNullable(this.ebsOptimized);
    }

    /**
     * The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
     * 
     */
    @Import(name="endDate")
    private @Nullable Output<String> endDate;

    /**
     * @return The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
     * 
     */
    public Optional<Output<String>> endDate() {
        return Optional.ofNullable(this.endDate);
    }

    /**
     * Indicates the way in which the Capacity Reservation ends. Specify either `unlimited` or `limited`.
     * 
     */
    @Import(name="endDateType")
    private @Nullable Output<String> endDateType;

    /**
     * @return Indicates the way in which the Capacity Reservation ends. Specify either `unlimited` or `limited`.
     * 
     */
    public Optional<Output<String>> endDateType() {
        return Optional.ofNullable(this.endDateType);
    }

    /**
     * Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
     * 
     */
    @Import(name="ephemeralStorage")
    private @Nullable Output<Boolean> ephemeralStorage;

    /**
     * @return Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
     * 
     */
    public Optional<Output<Boolean>> ephemeralStorage() {
        return Optional.ofNullable(this.ephemeralStorage);
    }

    /**
     * The number of instances for which to reserve capacity.
     * 
     */
    @Import(name="instanceCount")
    private @Nullable Output<Integer> instanceCount;

    /**
     * @return The number of instances for which to reserve capacity.
     * 
     */
    public Optional<Output<Integer>> instanceCount() {
        return Optional.ofNullable(this.instanceCount);
    }

    /**
     * Indicates the type of instance launches that the Capacity Reservation accepts. Specify either `open` or `targeted`.
     * 
     */
    @Import(name="instanceMatchCriteria")
    private @Nullable Output<String> instanceMatchCriteria;

    /**
     * @return Indicates the type of instance launches that the Capacity Reservation accepts. Specify either `open` or `targeted`.
     * 
     */
    public Optional<Output<String>> instanceMatchCriteria() {
        return Optional.ofNullable(this.instanceMatchCriteria);
    }

    /**
     * The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
     * 
     */
    @Import(name="instancePlatform")
    private @Nullable Output<Either<String,InstancePlatform>> instancePlatform;

    /**
     * @return The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
     * 
     */
    public Optional<Output<Either<String,InstancePlatform>>> instancePlatform() {
        return Optional.ofNullable(this.instancePlatform);
    }

    /**
     * The instance type for which to reserve capacity.
     * 
     */
    @Import(name="instanceType")
    private @Nullable Output<Either<String,InstanceType>> instanceType;

    /**
     * @return The instance type for which to reserve capacity.
     * 
     */
    public Optional<Output<Either<String,InstanceType>>> instanceType() {
        return Optional.ofNullable(this.instanceType);
    }

    /**
     * The Amazon Resource Name (ARN) of the Outpost on which to create the Capacity Reservation.
     * 
     */
    @Import(name="outpostArn")
    private @Nullable Output<String> outpostArn;

    /**
     * @return The Amazon Resource Name (ARN) of the Outpost on which to create the Capacity Reservation.
     * 
     */
    public Optional<Output<String>> outpostArn() {
        return Optional.ofNullable(this.outpostArn);
    }

    /**
     * The ID of the AWS account that owns the Capacity Reservation.
     * 
     */
    @Import(name="ownerId")
    private @Nullable Output<String> ownerId;

    /**
     * @return The ID of the AWS account that owns the Capacity Reservation.
     * 
     */
    public Optional<Output<String>> ownerId() {
        return Optional.ofNullable(this.ownerId);
    }

    /**
     * The Amazon Resource Name (ARN) of the cluster placement group in which to create the Capacity Reservation.
     * 
     */
    @Import(name="placementGroupArn")
    private @Nullable Output<String> placementGroupArn;

    /**
     * @return The Amazon Resource Name (ARN) of the cluster placement group in which to create the Capacity Reservation.
     * 
     */
    public Optional<Output<String>> placementGroupArn() {
        return Optional.ofNullable(this.placementGroupArn);
    }

    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block
     * 
     */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block
     * 
     */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
     * 
     */
    @Import(name="tenancy")
    private @Nullable Output<Either<String,Tenancy>> tenancy;

    /**
     * @return Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
     * 
     */
    public Optional<Output<Either<String,Tenancy>>> tenancy() {
        return Optional.ofNullable(this.tenancy);
    }

    private CapacityReservationState() {}

    private CapacityReservationState(CapacityReservationState $) {
        this.arn = $.arn;
        this.availabilityZone = $.availabilityZone;
        this.ebsOptimized = $.ebsOptimized;
        this.endDate = $.endDate;
        this.endDateType = $.endDateType;
        this.ephemeralStorage = $.ephemeralStorage;
        this.instanceCount = $.instanceCount;
        this.instanceMatchCriteria = $.instanceMatchCriteria;
        this.instancePlatform = $.instancePlatform;
        this.instanceType = $.instanceType;
        this.outpostArn = $.outpostArn;
        this.ownerId = $.ownerId;
        this.placementGroupArn = $.placementGroupArn;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.tenancy = $.tenancy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CapacityReservationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CapacityReservationState $;

        public Builder() {
            $ = new CapacityReservationState();
        }

        public Builder(CapacityReservationState defaults) {
            $ = new CapacityReservationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The ARN of the Capacity Reservation.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The ARN of the Capacity Reservation.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param availabilityZone The Availability Zone in which to create the Capacity Reservation.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(@Nullable Output<String> availabilityZone) {
            $.availabilityZone = availabilityZone;
            return this;
        }

        /**
         * @param availabilityZone The Availability Zone in which to create the Capacity Reservation.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(String availabilityZone) {
            return availabilityZone(Output.of(availabilityZone));
        }

        /**
         * @param ebsOptimized Indicates whether the Capacity Reservation supports EBS-optimized instances.
         * 
         * @return builder
         * 
         */
        public Builder ebsOptimized(@Nullable Output<Boolean> ebsOptimized) {
            $.ebsOptimized = ebsOptimized;
            return this;
        }

        /**
         * @param ebsOptimized Indicates whether the Capacity Reservation supports EBS-optimized instances.
         * 
         * @return builder
         * 
         */
        public Builder ebsOptimized(Boolean ebsOptimized) {
            return ebsOptimized(Output.of(ebsOptimized));
        }

        /**
         * @param endDate The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
         * 
         * @return builder
         * 
         */
        public Builder endDate(@Nullable Output<String> endDate) {
            $.endDate = endDate;
            return this;
        }

        /**
         * @param endDate The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
         * 
         * @return builder
         * 
         */
        public Builder endDate(String endDate) {
            return endDate(Output.of(endDate));
        }

        /**
         * @param endDateType Indicates the way in which the Capacity Reservation ends. Specify either `unlimited` or `limited`.
         * 
         * @return builder
         * 
         */
        public Builder endDateType(@Nullable Output<String> endDateType) {
            $.endDateType = endDateType;
            return this;
        }

        /**
         * @param endDateType Indicates the way in which the Capacity Reservation ends. Specify either `unlimited` or `limited`.
         * 
         * @return builder
         * 
         */
        public Builder endDateType(String endDateType) {
            return endDateType(Output.of(endDateType));
        }

        /**
         * @param ephemeralStorage Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
         * 
         * @return builder
         * 
         */
        public Builder ephemeralStorage(@Nullable Output<Boolean> ephemeralStorage) {
            $.ephemeralStorage = ephemeralStorage;
            return this;
        }

        /**
         * @param ephemeralStorage Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
         * 
         * @return builder
         * 
         */
        public Builder ephemeralStorage(Boolean ephemeralStorage) {
            return ephemeralStorage(Output.of(ephemeralStorage));
        }

        /**
         * @param instanceCount The number of instances for which to reserve capacity.
         * 
         * @return builder
         * 
         */
        public Builder instanceCount(@Nullable Output<Integer> instanceCount) {
            $.instanceCount = instanceCount;
            return this;
        }

        /**
         * @param instanceCount The number of instances for which to reserve capacity.
         * 
         * @return builder
         * 
         */
        public Builder instanceCount(Integer instanceCount) {
            return instanceCount(Output.of(instanceCount));
        }

        /**
         * @param instanceMatchCriteria Indicates the type of instance launches that the Capacity Reservation accepts. Specify either `open` or `targeted`.
         * 
         * @return builder
         * 
         */
        public Builder instanceMatchCriteria(@Nullable Output<String> instanceMatchCriteria) {
            $.instanceMatchCriteria = instanceMatchCriteria;
            return this;
        }

        /**
         * @param instanceMatchCriteria Indicates the type of instance launches that the Capacity Reservation accepts. Specify either `open` or `targeted`.
         * 
         * @return builder
         * 
         */
        public Builder instanceMatchCriteria(String instanceMatchCriteria) {
            return instanceMatchCriteria(Output.of(instanceMatchCriteria));
        }

        /**
         * @param instancePlatform The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
         * 
         * @return builder
         * 
         */
        public Builder instancePlatform(@Nullable Output<Either<String,InstancePlatform>> instancePlatform) {
            $.instancePlatform = instancePlatform;
            return this;
        }

        /**
         * @param instancePlatform The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
         * 
         * @return builder
         * 
         */
        public Builder instancePlatform(Either<String,InstancePlatform> instancePlatform) {
            return instancePlatform(Output.of(instancePlatform));
        }

        /**
         * @param instancePlatform The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
         * 
         * @return builder
         * 
         */
        public Builder instancePlatform(String instancePlatform) {
            return instancePlatform(Either.ofLeft(instancePlatform));
        }

        /**
         * @param instancePlatform The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
         * 
         * @return builder
         * 
         */
        public Builder instancePlatform(InstancePlatform instancePlatform) {
            return instancePlatform(Either.ofRight(instancePlatform));
        }

        /**
         * @param instanceType The instance type for which to reserve capacity.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(@Nullable Output<Either<String,InstanceType>> instanceType) {
            $.instanceType = instanceType;
            return this;
        }

        /**
         * @param instanceType The instance type for which to reserve capacity.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(Either<String,InstanceType> instanceType) {
            return instanceType(Output.of(instanceType));
        }

        /**
         * @param instanceType The instance type for which to reserve capacity.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(String instanceType) {
            return instanceType(Either.ofLeft(instanceType));
        }

        /**
         * @param instanceType The instance type for which to reserve capacity.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(InstanceType instanceType) {
            return instanceType(Either.ofRight(instanceType));
        }

        /**
         * @param outpostArn The Amazon Resource Name (ARN) of the Outpost on which to create the Capacity Reservation.
         * 
         * @return builder
         * 
         */
        public Builder outpostArn(@Nullable Output<String> outpostArn) {
            $.outpostArn = outpostArn;
            return this;
        }

        /**
         * @param outpostArn The Amazon Resource Name (ARN) of the Outpost on which to create the Capacity Reservation.
         * 
         * @return builder
         * 
         */
        public Builder outpostArn(String outpostArn) {
            return outpostArn(Output.of(outpostArn));
        }

        /**
         * @param ownerId The ID of the AWS account that owns the Capacity Reservation.
         * 
         * @return builder
         * 
         */
        public Builder ownerId(@Nullable Output<String> ownerId) {
            $.ownerId = ownerId;
            return this;
        }

        /**
         * @param ownerId The ID of the AWS account that owns the Capacity Reservation.
         * 
         * @return builder
         * 
         */
        public Builder ownerId(String ownerId) {
            return ownerId(Output.of(ownerId));
        }

        /**
         * @param placementGroupArn The Amazon Resource Name (ARN) of the cluster placement group in which to create the Capacity Reservation.
         * 
         * @return builder
         * 
         */
        public Builder placementGroupArn(@Nullable Output<String> placementGroupArn) {
            $.placementGroupArn = placementGroupArn;
            return this;
        }

        /**
         * @param placementGroupArn The Amazon Resource Name (ARN) of the cluster placement group in which to create the Capacity Reservation.
         * 
         * @return builder
         * 
         */
        public Builder placementGroupArn(String placementGroupArn) {
            return placementGroupArn(Output.of(placementGroupArn));
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block
         * 
         * @return builder
         * 
         */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block
         * 
         * @return builder
         * 
         */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param tenancy Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
         * 
         * @return builder
         * 
         */
        public Builder tenancy(@Nullable Output<Either<String,Tenancy>> tenancy) {
            $.tenancy = tenancy;
            return this;
        }

        /**
         * @param tenancy Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
         * 
         * @return builder
         * 
         */
        public Builder tenancy(Either<String,Tenancy> tenancy) {
            return tenancy(Output.of(tenancy));
        }

        /**
         * @param tenancy Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
         * 
         * @return builder
         * 
         */
        public Builder tenancy(String tenancy) {
            return tenancy(Either.ofLeft(tenancy));
        }

        /**
         * @param tenancy Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
         * 
         * @return builder
         * 
         */
        public Builder tenancy(Tenancy tenancy) {
            return tenancy(Either.ofRight(tenancy));
        }

        public CapacityReservationState build() {
            return $;
        }
    }

}
