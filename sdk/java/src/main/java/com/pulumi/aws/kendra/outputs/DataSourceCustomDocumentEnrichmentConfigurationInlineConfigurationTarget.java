// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kendra.outputs;

import com.pulumi.aws.kendra.outputs.DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationTargetTargetDocumentAttributeValue;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationTarget {
    /**
     * @return The identifier of the target document attribute or metadata field. For example, &#39;Department&#39; could be an identifier for the target attribute or metadata field that includes the department names associated with the documents.
     * 
     */
    private @Nullable String targetDocumentAttributeKey;
    /**
     * @return The target value you want to create for the target attribute. For example, &#39;Finance&#39; could be the target value for the target attribute key &#39;Department&#39;. See target_document_attribute_value.
     * 
     */
    private @Nullable DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationTargetTargetDocumentAttributeValue targetDocumentAttributeValue;
    /**
     * @return `TRUE` to delete the existing target value for your specified target attribute key. You cannot create a target value and set this to `TRUE`. To create a target value (`TargetDocumentAttributeValue`), set this to `FALSE`.
     * 
     */
    private @Nullable Boolean targetDocumentAttributeValueDeletion;

    private DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationTarget() {}
    /**
     * @return The identifier of the target document attribute or metadata field. For example, &#39;Department&#39; could be an identifier for the target attribute or metadata field that includes the department names associated with the documents.
     * 
     */
    public Optional<String> targetDocumentAttributeKey() {
        return Optional.ofNullable(this.targetDocumentAttributeKey);
    }
    /**
     * @return The target value you want to create for the target attribute. For example, &#39;Finance&#39; could be the target value for the target attribute key &#39;Department&#39;. See target_document_attribute_value.
     * 
     */
    public Optional<DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationTargetTargetDocumentAttributeValue> targetDocumentAttributeValue() {
        return Optional.ofNullable(this.targetDocumentAttributeValue);
    }
    /**
     * @return `TRUE` to delete the existing target value for your specified target attribute key. You cannot create a target value and set this to `TRUE`. To create a target value (`TargetDocumentAttributeValue`), set this to `FALSE`.
     * 
     */
    public Optional<Boolean> targetDocumentAttributeValueDeletion() {
        return Optional.ofNullable(this.targetDocumentAttributeValueDeletion);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationTarget defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String targetDocumentAttributeKey;
        private @Nullable DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationTargetTargetDocumentAttributeValue targetDocumentAttributeValue;
        private @Nullable Boolean targetDocumentAttributeValueDeletion;
        public Builder() {}
        public Builder(DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationTarget defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.targetDocumentAttributeKey = defaults.targetDocumentAttributeKey;
    	      this.targetDocumentAttributeValue = defaults.targetDocumentAttributeValue;
    	      this.targetDocumentAttributeValueDeletion = defaults.targetDocumentAttributeValueDeletion;
        }

        @CustomType.Setter
        public Builder targetDocumentAttributeKey(@Nullable String targetDocumentAttributeKey) {
            this.targetDocumentAttributeKey = targetDocumentAttributeKey;
            return this;
        }
        @CustomType.Setter
        public Builder targetDocumentAttributeValue(@Nullable DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationTargetTargetDocumentAttributeValue targetDocumentAttributeValue) {
            this.targetDocumentAttributeValue = targetDocumentAttributeValue;
            return this;
        }
        @CustomType.Setter
        public Builder targetDocumentAttributeValueDeletion(@Nullable Boolean targetDocumentAttributeValueDeletion) {
            this.targetDocumentAttributeValueDeletion = targetDocumentAttributeValueDeletion;
            return this;
        }
        public DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationTarget build() {
            final var o = new DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationTarget();
            o.targetDocumentAttributeKey = targetDocumentAttributeKey;
            o.targetDocumentAttributeValue = targetDocumentAttributeValue;
            o.targetDocumentAttributeValueDeletion = targetDocumentAttributeValueDeletion;
            return o;
        }
    }
}
