// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.glue.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class CrawlerS3Target {
    /**
     * @return The name of the connection to use to connect to the Delta table target.
     * 
     */
    private @Nullable String connectionName;
    /**
     * @return A valid Amazon SQS ARN.
     * 
     */
    private @Nullable String dlqEventQueueArn;
    /**
     * @return A valid Amazon SQS ARN.
     * 
     */
    private @Nullable String eventQueueArn;
    /**
     * @return A list of glob patterns used to exclude from the crawl.
     * 
     */
    private @Nullable List<String> exclusions;
    /**
     * @return The path of the Amazon DocumentDB or MongoDB target (database/collection).
     * 
     */
    private String path;
    /**
     * @return Sets the number of files in each leaf folder to be crawled when crawling sample files in a dataset. If not set, all the files are crawled. A valid value is an integer between 1 and 249.
     * 
     */
    private @Nullable Integer sampleSize;

    private CrawlerS3Target() {}
    /**
     * @return The name of the connection to use to connect to the Delta table target.
     * 
     */
    public Optional<String> connectionName() {
        return Optional.ofNullable(this.connectionName);
    }
    /**
     * @return A valid Amazon SQS ARN.
     * 
     */
    public Optional<String> dlqEventQueueArn() {
        return Optional.ofNullable(this.dlqEventQueueArn);
    }
    /**
     * @return A valid Amazon SQS ARN.
     * 
     */
    public Optional<String> eventQueueArn() {
        return Optional.ofNullable(this.eventQueueArn);
    }
    /**
     * @return A list of glob patterns used to exclude from the crawl.
     * 
     */
    public List<String> exclusions() {
        return this.exclusions == null ? List.of() : this.exclusions;
    }
    /**
     * @return The path of the Amazon DocumentDB or MongoDB target (database/collection).
     * 
     */
    public String path() {
        return this.path;
    }
    /**
     * @return Sets the number of files in each leaf folder to be crawled when crawling sample files in a dataset. If not set, all the files are crawled. A valid value is an integer between 1 and 249.
     * 
     */
    public Optional<Integer> sampleSize() {
        return Optional.ofNullable(this.sampleSize);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CrawlerS3Target defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String connectionName;
        private @Nullable String dlqEventQueueArn;
        private @Nullable String eventQueueArn;
        private @Nullable List<String> exclusions;
        private String path;
        private @Nullable Integer sampleSize;
        public Builder() {}
        public Builder(CrawlerS3Target defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.connectionName = defaults.connectionName;
    	      this.dlqEventQueueArn = defaults.dlqEventQueueArn;
    	      this.eventQueueArn = defaults.eventQueueArn;
    	      this.exclusions = defaults.exclusions;
    	      this.path = defaults.path;
    	      this.sampleSize = defaults.sampleSize;
        }

        @CustomType.Setter
        public Builder connectionName(@Nullable String connectionName) {
            this.connectionName = connectionName;
            return this;
        }
        @CustomType.Setter
        public Builder dlqEventQueueArn(@Nullable String dlqEventQueueArn) {
            this.dlqEventQueueArn = dlqEventQueueArn;
            return this;
        }
        @CustomType.Setter
        public Builder eventQueueArn(@Nullable String eventQueueArn) {
            this.eventQueueArn = eventQueueArn;
            return this;
        }
        @CustomType.Setter
        public Builder exclusions(@Nullable List<String> exclusions) {
            this.exclusions = exclusions;
            return this;
        }
        public Builder exclusions(String... exclusions) {
            return exclusions(List.of(exclusions));
        }
        @CustomType.Setter
        public Builder path(String path) {
            this.path = Objects.requireNonNull(path);
            return this;
        }
        @CustomType.Setter
        public Builder sampleSize(@Nullable Integer sampleSize) {
            this.sampleSize = sampleSize;
            return this;
        }
        public CrawlerS3Target build() {
            final var o = new CrawlerS3Target();
            o.connectionName = connectionName;
            o.dlqEventQueueArn = dlqEventQueueArn;
            o.eventQueueArn = eventQueueArn;
            o.exclusions = exclusions;
            o.path = path;
            o.sampleSize = sampleSize;
            return o;
        }
    }
}
