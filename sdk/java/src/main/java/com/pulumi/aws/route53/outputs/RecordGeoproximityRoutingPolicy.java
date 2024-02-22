// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.route53.outputs;

import com.pulumi.aws.route53.outputs.RecordGeoproximityRoutingPolicyCoordinate;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RecordGeoproximityRoutingPolicy {
    /**
     * @return A AWS region where the resource is present.
     * 
     */
    private @Nullable String awsRegion;
    /**
     * @return Route more traffic or less traffic to the resource by specifying a value ranges between -90 to 90. See https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy-geoproximity.html for bias details.
     * 
     */
    private @Nullable Integer bias;
    /**
     * @return Specify `latitude` and `longitude` for routing traffic to non-AWS resources.
     * 
     */
    private @Nullable List<RecordGeoproximityRoutingPolicyCoordinate> coordinates;
    /**
     * @return A AWS local zone group where the resource is present. See https://docs.aws.amazon.com/local-zones/latest/ug/available-local-zones.html for local zone group list.
     * 
     */
    private @Nullable String localZoneGroup;

    private RecordGeoproximityRoutingPolicy() {}
    /**
     * @return A AWS region where the resource is present.
     * 
     */
    public Optional<String> awsRegion() {
        return Optional.ofNullable(this.awsRegion);
    }
    /**
     * @return Route more traffic or less traffic to the resource by specifying a value ranges between -90 to 90. See https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy-geoproximity.html for bias details.
     * 
     */
    public Optional<Integer> bias() {
        return Optional.ofNullable(this.bias);
    }
    /**
     * @return Specify `latitude` and `longitude` for routing traffic to non-AWS resources.
     * 
     */
    public List<RecordGeoproximityRoutingPolicyCoordinate> coordinates() {
        return this.coordinates == null ? List.of() : this.coordinates;
    }
    /**
     * @return A AWS local zone group where the resource is present. See https://docs.aws.amazon.com/local-zones/latest/ug/available-local-zones.html for local zone group list.
     * 
     */
    public Optional<String> localZoneGroup() {
        return Optional.ofNullable(this.localZoneGroup);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RecordGeoproximityRoutingPolicy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String awsRegion;
        private @Nullable Integer bias;
        private @Nullable List<RecordGeoproximityRoutingPolicyCoordinate> coordinates;
        private @Nullable String localZoneGroup;
        public Builder() {}
        public Builder(RecordGeoproximityRoutingPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.awsRegion = defaults.awsRegion;
    	      this.bias = defaults.bias;
    	      this.coordinates = defaults.coordinates;
    	      this.localZoneGroup = defaults.localZoneGroup;
        }

        @CustomType.Setter
        public Builder awsRegion(@Nullable String awsRegion) {

            this.awsRegion = awsRegion;
            return this;
        }
        @CustomType.Setter
        public Builder bias(@Nullable Integer bias) {

            this.bias = bias;
            return this;
        }
        @CustomType.Setter
        public Builder coordinates(@Nullable List<RecordGeoproximityRoutingPolicyCoordinate> coordinates) {

            this.coordinates = coordinates;
            return this;
        }
        public Builder coordinates(RecordGeoproximityRoutingPolicyCoordinate... coordinates) {
            return coordinates(List.of(coordinates));
        }
        @CustomType.Setter
        public Builder localZoneGroup(@Nullable String localZoneGroup) {

            this.localZoneGroup = localZoneGroup;
            return this;
        }
        public RecordGeoproximityRoutingPolicy build() {
            final var _resultValue = new RecordGeoproximityRoutingPolicy();
            _resultValue.awsRegion = awsRegion;
            _resultValue.bias = bias;
            _resultValue.coordinates = coordinates;
            _resultValue.localZoneGroup = localZoneGroup;
            return _resultValue;
        }
    }
}
