// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.opensearch.inputs;

import com.pulumi.aws.opensearch.inputs.ServerlessCollectionTimeoutsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServerlessCollectionState extends com.pulumi.resources.ResourceArgs {

    public static final ServerlessCollectionState Empty = new ServerlessCollectionState();

    /**
     * Amazon Resource Name (ARN) of the collection.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of the collection.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Collection-specific endpoint used to submit index, search, and data upload requests to an OpenSearch Serverless collection.
     * 
     */
    @Import(name="collectionEndpoint")
    private @Nullable Output<String> collectionEndpoint;

    /**
     * @return Collection-specific endpoint used to submit index, search, and data upload requests to an OpenSearch Serverless collection.
     * 
     */
    public Optional<Output<String>> collectionEndpoint() {
        return Optional.ofNullable(this.collectionEndpoint);
    }

    /**
     * Collection-specific endpoint used to access OpenSearch Dashboards.
     * 
     */
    @Import(name="dashboardEndpoint")
    private @Nullable Output<String> dashboardEndpoint;

    /**
     * @return Collection-specific endpoint used to access OpenSearch Dashboards.
     * 
     */
    public Optional<Output<String>> dashboardEndpoint() {
        return Optional.ofNullable(this.dashboardEndpoint);
    }

    /**
     * Description of the collection.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the collection.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The ARN of the Amazon Web Services KMS key used to encrypt the collection.
     * 
     */
    @Import(name="kmsKeyArn")
    private @Nullable Output<String> kmsKeyArn;

    /**
     * @return The ARN of the Amazon Web Services KMS key used to encrypt the collection.
     * 
     */
    public Optional<Output<String>> kmsKeyArn() {
        return Optional.ofNullable(this.kmsKeyArn);
    }

    /**
     * Name of the collection.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the collection.
     * 
     * The following arguments are optional:
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Indicates whether standby replicas should be used for a collection. One of `ENABLED` or `DISABLED`. Defaults to `ENABLED`.
     * 
     */
    @Import(name="standbyReplicas")
    private @Nullable Output<String> standbyReplicas;

    /**
     * @return Indicates whether standby replicas should be used for a collection. One of `ENABLED` or `DISABLED`. Defaults to `ENABLED`.
     * 
     */
    public Optional<Output<String>> standbyReplicas() {
        return Optional.ofNullable(this.standbyReplicas);
    }

    /**
     * A map of tags to assign to the collection. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the collection. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    @Import(name="timeouts")
    private @Nullable Output<ServerlessCollectionTimeoutsArgs> timeouts;

    public Optional<Output<ServerlessCollectionTimeoutsArgs>> timeouts() {
        return Optional.ofNullable(this.timeouts);
    }

    /**
     * Type of collection. One of `SEARCH`, `TIMESERIES`, or `VECTORSEARCH`. Defaults to `TIMESERIES`.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Type of collection. One of `SEARCH`, `TIMESERIES`, or `VECTORSEARCH`. Defaults to `TIMESERIES`.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private ServerlessCollectionState() {}

    private ServerlessCollectionState(ServerlessCollectionState $) {
        this.arn = $.arn;
        this.collectionEndpoint = $.collectionEndpoint;
        this.dashboardEndpoint = $.dashboardEndpoint;
        this.description = $.description;
        this.kmsKeyArn = $.kmsKeyArn;
        this.name = $.name;
        this.standbyReplicas = $.standbyReplicas;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.timeouts = $.timeouts;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServerlessCollectionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServerlessCollectionState $;

        public Builder() {
            $ = new ServerlessCollectionState();
        }

        public Builder(ServerlessCollectionState defaults) {
            $ = new ServerlessCollectionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn Amazon Resource Name (ARN) of the collection.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn Amazon Resource Name (ARN) of the collection.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param collectionEndpoint Collection-specific endpoint used to submit index, search, and data upload requests to an OpenSearch Serverless collection.
         * 
         * @return builder
         * 
         */
        public Builder collectionEndpoint(@Nullable Output<String> collectionEndpoint) {
            $.collectionEndpoint = collectionEndpoint;
            return this;
        }

        /**
         * @param collectionEndpoint Collection-specific endpoint used to submit index, search, and data upload requests to an OpenSearch Serverless collection.
         * 
         * @return builder
         * 
         */
        public Builder collectionEndpoint(String collectionEndpoint) {
            return collectionEndpoint(Output.of(collectionEndpoint));
        }

        /**
         * @param dashboardEndpoint Collection-specific endpoint used to access OpenSearch Dashboards.
         * 
         * @return builder
         * 
         */
        public Builder dashboardEndpoint(@Nullable Output<String> dashboardEndpoint) {
            $.dashboardEndpoint = dashboardEndpoint;
            return this;
        }

        /**
         * @param dashboardEndpoint Collection-specific endpoint used to access OpenSearch Dashboards.
         * 
         * @return builder
         * 
         */
        public Builder dashboardEndpoint(String dashboardEndpoint) {
            return dashboardEndpoint(Output.of(dashboardEndpoint));
        }

        /**
         * @param description Description of the collection.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the collection.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param kmsKeyArn The ARN of the Amazon Web Services KMS key used to encrypt the collection.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyArn(@Nullable Output<String> kmsKeyArn) {
            $.kmsKeyArn = kmsKeyArn;
            return this;
        }

        /**
         * @param kmsKeyArn The ARN of the Amazon Web Services KMS key used to encrypt the collection.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyArn(String kmsKeyArn) {
            return kmsKeyArn(Output.of(kmsKeyArn));
        }

        /**
         * @param name Name of the collection.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the collection.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param standbyReplicas Indicates whether standby replicas should be used for a collection. One of `ENABLED` or `DISABLED`. Defaults to `ENABLED`.
         * 
         * @return builder
         * 
         */
        public Builder standbyReplicas(@Nullable Output<String> standbyReplicas) {
            $.standbyReplicas = standbyReplicas;
            return this;
        }

        /**
         * @param standbyReplicas Indicates whether standby replicas should be used for a collection. One of `ENABLED` or `DISABLED`. Defaults to `ENABLED`.
         * 
         * @return builder
         * 
         */
        public Builder standbyReplicas(String standbyReplicas) {
            return standbyReplicas(Output.of(standbyReplicas));
        }

        /**
         * @param tags A map of tags to assign to the collection. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the collection. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        public Builder timeouts(@Nullable Output<ServerlessCollectionTimeoutsArgs> timeouts) {
            $.timeouts = timeouts;
            return this;
        }

        public Builder timeouts(ServerlessCollectionTimeoutsArgs timeouts) {
            return timeouts(Output.of(timeouts));
        }

        /**
         * @param type Type of collection. One of `SEARCH`, `TIMESERIES`, or `VECTORSEARCH`. Defaults to `TIMESERIES`.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of collection. One of `SEARCH`, `TIMESERIES`, or `VECTORSEARCH`. Defaults to `TIMESERIES`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public ServerlessCollectionState build() {
            return $;
        }
    }

}
