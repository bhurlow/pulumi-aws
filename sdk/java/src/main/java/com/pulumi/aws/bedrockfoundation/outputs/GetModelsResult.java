// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.bedrockfoundation.outputs;

import com.pulumi.aws.bedrockfoundation.outputs.GetModelsModelSummary;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetModelsResult {
    private @Nullable String byCustomizationType;
    private @Nullable String byInferenceType;
    private @Nullable String byOutputModality;
    private @Nullable String byProvider;
    /**
     * @return AWS region.
     * 
     */
    private String id;
    /**
     * @return List of model summary objects. See `model_summaries`.
     * 
     */
    private @Nullable List<GetModelsModelSummary> modelSummaries;

    private GetModelsResult() {}
    public Optional<String> byCustomizationType() {
        return Optional.ofNullable(this.byCustomizationType);
    }
    public Optional<String> byInferenceType() {
        return Optional.ofNullable(this.byInferenceType);
    }
    public Optional<String> byOutputModality() {
        return Optional.ofNullable(this.byOutputModality);
    }
    public Optional<String> byProvider() {
        return Optional.ofNullable(this.byProvider);
    }
    /**
     * @return AWS region.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return List of model summary objects. See `model_summaries`.
     * 
     */
    public List<GetModelsModelSummary> modelSummaries() {
        return this.modelSummaries == null ? List.of() : this.modelSummaries;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetModelsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String byCustomizationType;
        private @Nullable String byInferenceType;
        private @Nullable String byOutputModality;
        private @Nullable String byProvider;
        private String id;
        private @Nullable List<GetModelsModelSummary> modelSummaries;
        public Builder() {}
        public Builder(GetModelsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.byCustomizationType = defaults.byCustomizationType;
    	      this.byInferenceType = defaults.byInferenceType;
    	      this.byOutputModality = defaults.byOutputModality;
    	      this.byProvider = defaults.byProvider;
    	      this.id = defaults.id;
    	      this.modelSummaries = defaults.modelSummaries;
        }

        @CustomType.Setter
        public Builder byCustomizationType(@Nullable String byCustomizationType) {
            this.byCustomizationType = byCustomizationType;
            return this;
        }
        @CustomType.Setter
        public Builder byInferenceType(@Nullable String byInferenceType) {
            this.byInferenceType = byInferenceType;
            return this;
        }
        @CustomType.Setter
        public Builder byOutputModality(@Nullable String byOutputModality) {
            this.byOutputModality = byOutputModality;
            return this;
        }
        @CustomType.Setter
        public Builder byProvider(@Nullable String byProvider) {
            this.byProvider = byProvider;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder modelSummaries(@Nullable List<GetModelsModelSummary> modelSummaries) {
            this.modelSummaries = modelSummaries;
            return this;
        }
        public Builder modelSummaries(GetModelsModelSummary... modelSummaries) {
            return modelSummaries(List.of(modelSummaries));
        }
        public GetModelsResult build() {
            final var o = new GetModelsResult();
            o.byCustomizationType = byCustomizationType;
            o.byInferenceType = byInferenceType;
            o.byOutputModality = byOutputModality;
            o.byProvider = byProvider;
            o.id = id;
            o.modelSummaries = modelSummaries;
            return o;
        }
    }
}
