// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.quicksight.outputs;

import com.pulumi.aws.quicksight.outputs.TemplateSourceEntitySourceAnalysisDataSetReference;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class TemplateSourceEntitySourceAnalysis {
    /**
     * @return The Amazon Resource Name (ARN) of the resource.
     * 
     */
    private String arn;
    /**
     * @return A list of dataset references used as placeholders in the template. See data_set_references.
     * 
     */
    private List<TemplateSourceEntitySourceAnalysisDataSetReference> dataSetReferences;

    private TemplateSourceEntitySourceAnalysis() {}
    /**
     * @return The Amazon Resource Name (ARN) of the resource.
     * 
     */
    public String arn() {
        return this.arn;
    }
    /**
     * @return A list of dataset references used as placeholders in the template. See data_set_references.
     * 
     */
    public List<TemplateSourceEntitySourceAnalysisDataSetReference> dataSetReferences() {
        return this.dataSetReferences;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(TemplateSourceEntitySourceAnalysis defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String arn;
        private List<TemplateSourceEntitySourceAnalysisDataSetReference> dataSetReferences;
        public Builder() {}
        public Builder(TemplateSourceEntitySourceAnalysis defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.arn = defaults.arn;
    	      this.dataSetReferences = defaults.dataSetReferences;
        }

        @CustomType.Setter
        public Builder arn(String arn) {
            this.arn = Objects.requireNonNull(arn);
            return this;
        }
        @CustomType.Setter
        public Builder dataSetReferences(List<TemplateSourceEntitySourceAnalysisDataSetReference> dataSetReferences) {
            this.dataSetReferences = Objects.requireNonNull(dataSetReferences);
            return this;
        }
        public Builder dataSetReferences(TemplateSourceEntitySourceAnalysisDataSetReference... dataSetReferences) {
            return dataSetReferences(List.of(dataSetReferences));
        }
        public TemplateSourceEntitySourceAnalysis build() {
            final var o = new TemplateSourceEntitySourceAnalysis();
            o.arn = arn;
            o.dataSetReferences = dataSetReferences;
            return o;
        }
    }
}
