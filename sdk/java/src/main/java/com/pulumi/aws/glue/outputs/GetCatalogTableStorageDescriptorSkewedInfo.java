// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.glue.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetCatalogTableStorageDescriptorSkewedInfo {
    /**
     * @return List of names of columns that contain skewed values.
     * 
     */
    private List<String> skewedColumnNames;
    /**
     * @return List of values that appear so frequently as to be considered skewed.
     * 
     */
    private Map<String,String> skewedColumnValueLocationMaps;
    /**
     * @return Map of skewed values to the columns that contain them.
     * 
     */
    private List<String> skewedColumnValues;

    private GetCatalogTableStorageDescriptorSkewedInfo() {}
    /**
     * @return List of names of columns that contain skewed values.
     * 
     */
    public List<String> skewedColumnNames() {
        return this.skewedColumnNames;
    }
    /**
     * @return List of values that appear so frequently as to be considered skewed.
     * 
     */
    public Map<String,String> skewedColumnValueLocationMaps() {
        return this.skewedColumnValueLocationMaps;
    }
    /**
     * @return Map of skewed values to the columns that contain them.
     * 
     */
    public List<String> skewedColumnValues() {
        return this.skewedColumnValues;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCatalogTableStorageDescriptorSkewedInfo defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> skewedColumnNames;
        private Map<String,String> skewedColumnValueLocationMaps;
        private List<String> skewedColumnValues;
        public Builder() {}
        public Builder(GetCatalogTableStorageDescriptorSkewedInfo defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.skewedColumnNames = defaults.skewedColumnNames;
    	      this.skewedColumnValueLocationMaps = defaults.skewedColumnValueLocationMaps;
    	      this.skewedColumnValues = defaults.skewedColumnValues;
        }

        @CustomType.Setter
        public Builder skewedColumnNames(List<String> skewedColumnNames) {
            this.skewedColumnNames = Objects.requireNonNull(skewedColumnNames);
            return this;
        }
        public Builder skewedColumnNames(String... skewedColumnNames) {
            return skewedColumnNames(List.of(skewedColumnNames));
        }
        @CustomType.Setter
        public Builder skewedColumnValueLocationMaps(Map<String,String> skewedColumnValueLocationMaps) {
            this.skewedColumnValueLocationMaps = Objects.requireNonNull(skewedColumnValueLocationMaps);
            return this;
        }
        @CustomType.Setter
        public Builder skewedColumnValues(List<String> skewedColumnValues) {
            this.skewedColumnValues = Objects.requireNonNull(skewedColumnValues);
            return this;
        }
        public Builder skewedColumnValues(String... skewedColumnValues) {
            return skewedColumnValues(List.of(skewedColumnValues));
        }
        public GetCatalogTableStorageDescriptorSkewedInfo build() {
            final var o = new GetCatalogTableStorageDescriptorSkewedInfo();
            o.skewedColumnNames = skewedColumnNames;
            o.skewedColumnValueLocationMaps = skewedColumnValueLocationMaps;
            o.skewedColumnValues = skewedColumnValues;
            return o;
        }
    }
}
