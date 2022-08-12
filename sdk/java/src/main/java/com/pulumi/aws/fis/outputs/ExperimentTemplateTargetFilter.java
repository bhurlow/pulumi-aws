// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.fis.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class ExperimentTemplateTargetFilter {
    /**
     * @return Attribute path for the filter.
     * 
     */
    private final String path;
    /**
     * @return Set of attribute values for the filter.
     * 
     */
    private final List<String> values;

    @CustomType.Constructor
    private ExperimentTemplateTargetFilter(
        @CustomType.Parameter("path") String path,
        @CustomType.Parameter("values") List<String> values) {
        this.path = path;
        this.values = values;
    }

    /**
     * @return Attribute path for the filter.
     * 
     */
    public String path() {
        return this.path;
    }
    /**
     * @return Set of attribute values for the filter.
     * 
     */
    public List<String> values() {
        return this.values;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ExperimentTemplateTargetFilter defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String path;
        private List<String> values;

        public Builder() {
    	      // Empty
        }

        public Builder(ExperimentTemplateTargetFilter defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.path = defaults.path;
    	      this.values = defaults.values;
        }

        public Builder path(String path) {
            this.path = Objects.requireNonNull(path);
            return this;
        }
        public Builder values(List<String> values) {
            this.values = Objects.requireNonNull(values);
            return this;
        }
        public Builder values(String... values) {
            return values(List.of(values));
        }        public ExperimentTemplateTargetFilter build() {
            return new ExperimentTemplateTargetFilter(path, values);
        }
    }
}