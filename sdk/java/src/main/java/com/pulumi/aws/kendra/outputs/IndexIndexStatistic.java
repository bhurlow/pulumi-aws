// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kendra.outputs;

import com.pulumi.aws.kendra.outputs.IndexIndexStatisticFaqStatistic;
import com.pulumi.aws.kendra.outputs.IndexIndexStatisticTextDocumentStatistic;
import com.pulumi.core.annotations.CustomType;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class IndexIndexStatistic {
    /**
     * @return A block that specifies the number of question and answer topics in the index. Documented below.
     * 
     */
    private final @Nullable List<IndexIndexStatisticFaqStatistic> faqStatistics;
    /**
     * @return A block that specifies the number of text documents indexed.
     * 
     */
    private final @Nullable List<IndexIndexStatisticTextDocumentStatistic> textDocumentStatistics;

    @CustomType.Constructor
    private IndexIndexStatistic(
        @CustomType.Parameter("faqStatistics") @Nullable List<IndexIndexStatisticFaqStatistic> faqStatistics,
        @CustomType.Parameter("textDocumentStatistics") @Nullable List<IndexIndexStatisticTextDocumentStatistic> textDocumentStatistics) {
        this.faqStatistics = faqStatistics;
        this.textDocumentStatistics = textDocumentStatistics;
    }

    /**
     * @return A block that specifies the number of question and answer topics in the index. Documented below.
     * 
     */
    public List<IndexIndexStatisticFaqStatistic> faqStatistics() {
        return this.faqStatistics == null ? List.of() : this.faqStatistics;
    }
    /**
     * @return A block that specifies the number of text documents indexed.
     * 
     */
    public List<IndexIndexStatisticTextDocumentStatistic> textDocumentStatistics() {
        return this.textDocumentStatistics == null ? List.of() : this.textDocumentStatistics;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IndexIndexStatistic defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable List<IndexIndexStatisticFaqStatistic> faqStatistics;
        private @Nullable List<IndexIndexStatisticTextDocumentStatistic> textDocumentStatistics;

        public Builder() {
    	      // Empty
        }

        public Builder(IndexIndexStatistic defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.faqStatistics = defaults.faqStatistics;
    	      this.textDocumentStatistics = defaults.textDocumentStatistics;
        }

        public Builder faqStatistics(@Nullable List<IndexIndexStatisticFaqStatistic> faqStatistics) {
            this.faqStatistics = faqStatistics;
            return this;
        }
        public Builder faqStatistics(IndexIndexStatisticFaqStatistic... faqStatistics) {
            return faqStatistics(List.of(faqStatistics));
        }
        public Builder textDocumentStatistics(@Nullable List<IndexIndexStatisticTextDocumentStatistic> textDocumentStatistics) {
            this.textDocumentStatistics = textDocumentStatistics;
            return this;
        }
        public Builder textDocumentStatistics(IndexIndexStatisticTextDocumentStatistic... textDocumentStatistics) {
            return textDocumentStatistics(List.of(textDocumentStatistics));
        }        public IndexIndexStatistic build() {
            return new IndexIndexStatistic(faqStatistics, textDocumentStatistics);
        }
    }
}