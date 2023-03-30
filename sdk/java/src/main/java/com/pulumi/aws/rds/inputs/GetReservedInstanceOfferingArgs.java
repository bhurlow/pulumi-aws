// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.rds.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class GetReservedInstanceOfferingArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetReservedInstanceOfferingArgs Empty = new GetReservedInstanceOfferingArgs();

    /**
     * DB instance class for the reserved DB instance.
     * 
     */
    @Import(name="dbInstanceClass", required=true)
    private Output<String> dbInstanceClass;

    /**
     * @return DB instance class for the reserved DB instance.
     * 
     */
    public Output<String> dbInstanceClass() {
        return this.dbInstanceClass;
    }

    /**
     * Duration of the reservation in years or seconds. Valid values are `1`, `3`, `31536000`, `94608000`
     * 
     */
    @Import(name="duration", required=true)
    private Output<Integer> duration;

    /**
     * @return Duration of the reservation in years or seconds. Valid values are `1`, `3`, `31536000`, `94608000`
     * 
     */
    public Output<Integer> duration() {
        return this.duration;
    }

    /**
     * Whether the reservation applies to Multi-AZ deployments.
     * 
     */
    @Import(name="multiAz", required=true)
    private Output<Boolean> multiAz;

    /**
     * @return Whether the reservation applies to Multi-AZ deployments.
     * 
     */
    public Output<Boolean> multiAz() {
        return this.multiAz;
    }

    /**
     * Offering type of this reserved DB instance. Valid values are `No Upfront`, `Partial Upfront`, `All Upfront`.
     * 
     */
    @Import(name="offeringType", required=true)
    private Output<String> offeringType;

    /**
     * @return Offering type of this reserved DB instance. Valid values are `No Upfront`, `Partial Upfront`, `All Upfront`.
     * 
     */
    public Output<String> offeringType() {
        return this.offeringType;
    }

    /**
     * Description of the reserved DB instance.
     * 
     */
    @Import(name="productDescription", required=true)
    private Output<String> productDescription;

    /**
     * @return Description of the reserved DB instance.
     * 
     */
    public Output<String> productDescription() {
        return this.productDescription;
    }

    private GetReservedInstanceOfferingArgs() {}

    private GetReservedInstanceOfferingArgs(GetReservedInstanceOfferingArgs $) {
        this.dbInstanceClass = $.dbInstanceClass;
        this.duration = $.duration;
        this.multiAz = $.multiAz;
        this.offeringType = $.offeringType;
        this.productDescription = $.productDescription;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetReservedInstanceOfferingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetReservedInstanceOfferingArgs $;

        public Builder() {
            $ = new GetReservedInstanceOfferingArgs();
        }

        public Builder(GetReservedInstanceOfferingArgs defaults) {
            $ = new GetReservedInstanceOfferingArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dbInstanceClass DB instance class for the reserved DB instance.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceClass(Output<String> dbInstanceClass) {
            $.dbInstanceClass = dbInstanceClass;
            return this;
        }

        /**
         * @param dbInstanceClass DB instance class for the reserved DB instance.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceClass(String dbInstanceClass) {
            return dbInstanceClass(Output.of(dbInstanceClass));
        }

        /**
         * @param duration Duration of the reservation in years or seconds. Valid values are `1`, `3`, `31536000`, `94608000`
         * 
         * @return builder
         * 
         */
        public Builder duration(Output<Integer> duration) {
            $.duration = duration;
            return this;
        }

        /**
         * @param duration Duration of the reservation in years or seconds. Valid values are `1`, `3`, `31536000`, `94608000`
         * 
         * @return builder
         * 
         */
        public Builder duration(Integer duration) {
            return duration(Output.of(duration));
        }

        /**
         * @param multiAz Whether the reservation applies to Multi-AZ deployments.
         * 
         * @return builder
         * 
         */
        public Builder multiAz(Output<Boolean> multiAz) {
            $.multiAz = multiAz;
            return this;
        }

        /**
         * @param multiAz Whether the reservation applies to Multi-AZ deployments.
         * 
         * @return builder
         * 
         */
        public Builder multiAz(Boolean multiAz) {
            return multiAz(Output.of(multiAz));
        }

        /**
         * @param offeringType Offering type of this reserved DB instance. Valid values are `No Upfront`, `Partial Upfront`, `All Upfront`.
         * 
         * @return builder
         * 
         */
        public Builder offeringType(Output<String> offeringType) {
            $.offeringType = offeringType;
            return this;
        }

        /**
         * @param offeringType Offering type of this reserved DB instance. Valid values are `No Upfront`, `Partial Upfront`, `All Upfront`.
         * 
         * @return builder
         * 
         */
        public Builder offeringType(String offeringType) {
            return offeringType(Output.of(offeringType));
        }

        /**
         * @param productDescription Description of the reserved DB instance.
         * 
         * @return builder
         * 
         */
        public Builder productDescription(Output<String> productDescription) {
            $.productDescription = productDescription;
            return this;
        }

        /**
         * @param productDescription Description of the reserved DB instance.
         * 
         * @return builder
         * 
         */
        public Builder productDescription(String productDescription) {
            return productDescription(Output.of(productDescription));
        }

        public GetReservedInstanceOfferingArgs build() {
            $.dbInstanceClass = Objects.requireNonNull($.dbInstanceClass, "expected parameter 'dbInstanceClass' to be non-null");
            $.duration = Objects.requireNonNull($.duration, "expected parameter 'duration' to be non-null");
            $.multiAz = Objects.requireNonNull($.multiAz, "expected parameter 'multiAz' to be non-null");
            $.offeringType = Objects.requireNonNull($.offeringType, "expected parameter 'offeringType' to be non-null");
            $.productDescription = Objects.requireNonNull($.productDescription, "expected parameter 'productDescription' to be non-null");
            return $;
        }
    }

}
