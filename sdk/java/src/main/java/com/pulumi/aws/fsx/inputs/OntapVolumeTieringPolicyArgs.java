// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.fsx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OntapVolumeTieringPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final OntapVolumeTieringPolicyArgs Empty = new OntapVolumeTieringPolicyArgs();

    /**
     * Specifies the number of days that user data in a volume must remain inactive before it is considered &#34;cold&#34; and moved to the capacity pool. Used with `AUTO` and `SNAPSHOT_ONLY` tiering policies only. Valid values are whole numbers between 2 and 183. Default values are 31 days for `AUTO` and 2 days for `SNAPSHOT_ONLY`.
     * 
     */
    @Import(name="coolingPeriod")
    private @Nullable Output<Integer> coolingPeriod;

    /**
     * @return Specifies the number of days that user data in a volume must remain inactive before it is considered &#34;cold&#34; and moved to the capacity pool. Used with `AUTO` and `SNAPSHOT_ONLY` tiering policies only. Valid values are whole numbers between 2 and 183. Default values are 31 days for `AUTO` and 2 days for `SNAPSHOT_ONLY`.
     * 
     */
    public Optional<Output<Integer>> coolingPeriod() {
        return Optional.ofNullable(this.coolingPeriod);
    }

    /**
     * Specifies the tiering policy for the ONTAP volume for moving data to the capacity pool storage. Valid values are `SNAPSHOT_ONLY`, `AUTO`, `ALL`, `NONE`. Default value is `SNAPSHOT_ONLY`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Specifies the tiering policy for the ONTAP volume for moving data to the capacity pool storage. Valid values are `SNAPSHOT_ONLY`, `AUTO`, `ALL`, `NONE`. Default value is `SNAPSHOT_ONLY`.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private OntapVolumeTieringPolicyArgs() {}

    private OntapVolumeTieringPolicyArgs(OntapVolumeTieringPolicyArgs $) {
        this.coolingPeriod = $.coolingPeriod;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OntapVolumeTieringPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OntapVolumeTieringPolicyArgs $;

        public Builder() {
            $ = new OntapVolumeTieringPolicyArgs();
        }

        public Builder(OntapVolumeTieringPolicyArgs defaults) {
            $ = new OntapVolumeTieringPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param coolingPeriod Specifies the number of days that user data in a volume must remain inactive before it is considered &#34;cold&#34; and moved to the capacity pool. Used with `AUTO` and `SNAPSHOT_ONLY` tiering policies only. Valid values are whole numbers between 2 and 183. Default values are 31 days for `AUTO` and 2 days for `SNAPSHOT_ONLY`.
         * 
         * @return builder
         * 
         */
        public Builder coolingPeriod(@Nullable Output<Integer> coolingPeriod) {
            $.coolingPeriod = coolingPeriod;
            return this;
        }

        /**
         * @param coolingPeriod Specifies the number of days that user data in a volume must remain inactive before it is considered &#34;cold&#34; and moved to the capacity pool. Used with `AUTO` and `SNAPSHOT_ONLY` tiering policies only. Valid values are whole numbers between 2 and 183. Default values are 31 days for `AUTO` and 2 days for `SNAPSHOT_ONLY`.
         * 
         * @return builder
         * 
         */
        public Builder coolingPeriod(Integer coolingPeriod) {
            return coolingPeriod(Output.of(coolingPeriod));
        }

        /**
         * @param name Specifies the tiering policy for the ONTAP volume for moving data to the capacity pool storage. Valid values are `SNAPSHOT_ONLY`, `AUTO`, `ALL`, `NONE`. Default value is `SNAPSHOT_ONLY`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Specifies the tiering policy for the ONTAP volume for moving data to the capacity pool storage. Valid values are `SNAPSHOT_ONLY`, `AUTO`, `ALL`, `NONE`. Default value is `SNAPSHOT_ONLY`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public OntapVolumeTieringPolicyArgs build() {
            return $;
        }
    }

}
