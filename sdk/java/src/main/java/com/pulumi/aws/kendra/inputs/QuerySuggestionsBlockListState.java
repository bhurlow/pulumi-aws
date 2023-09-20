// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kendra.inputs;

import com.pulumi.aws.kendra.inputs.QuerySuggestionsBlockListSourceS3PathArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QuerySuggestionsBlockListState extends com.pulumi.resources.ResourceArgs {

    public static final QuerySuggestionsBlockListState Empty = new QuerySuggestionsBlockListState();

    /**
     * ARN of the block list.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return ARN of the block list.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Description for a block list.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description for a block list.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Identifier of the index for a block list.
     * 
     */
    @Import(name="indexId")
    private @Nullable Output<String> indexId;

    /**
     * @return Identifier of the index for a block list.
     * 
     */
    public Optional<Output<String>> indexId() {
        return Optional.ofNullable(this.indexId);
    }

    /**
     * Name for the block list.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name for the block list.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Unique identifier of the block list.
     * 
     */
    @Import(name="querySuggestionsBlockListId")
    private @Nullable Output<String> querySuggestionsBlockListId;

    /**
     * @return Unique identifier of the block list.
     * 
     */
    public Optional<Output<String>> querySuggestionsBlockListId() {
        return Optional.ofNullable(this.querySuggestionsBlockListId);
    }

    /**
     * IAM (Identity and Access Management) role used to access the block list text file in S3.
     * 
     */
    @Import(name="roleArn")
    private @Nullable Output<String> roleArn;

    /**
     * @return IAM (Identity and Access Management) role used to access the block list text file in S3.
     * 
     */
    public Optional<Output<String>> roleArn() {
        return Optional.ofNullable(this.roleArn);
    }

    /**
     * S3 path where your block list text file is located. See details below.
     * 
     * The `source_s3_path` configuration block supports the following arguments:
     * 
     */
    @Import(name="sourceS3Path")
    private @Nullable Output<QuerySuggestionsBlockListSourceS3PathArgs> sourceS3Path;

    /**
     * @return S3 path where your block list text file is located. See details below.
     * 
     * The `source_s3_path` configuration block supports the following arguments:
     * 
     */
    public Optional<Output<QuerySuggestionsBlockListSourceS3PathArgs>> sourceS3Path() {
        return Optional.ofNullable(this.sourceS3Path);
    }

    @Import(name="status")
    private @Nullable Output<String> status;

    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Key-value map of resource tags. If configured with a provider `default_tags` configuration block, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. If configured with a provider `default_tags` configuration block, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Map of tags assigned to the resource, including those inherited from the provider&#39;s default_tags configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return Map of tags assigned to the resource, including those inherited from the provider&#39;s default_tags configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    private QuerySuggestionsBlockListState() {}

    private QuerySuggestionsBlockListState(QuerySuggestionsBlockListState $) {
        this.arn = $.arn;
        this.description = $.description;
        this.indexId = $.indexId;
        this.name = $.name;
        this.querySuggestionsBlockListId = $.querySuggestionsBlockListId;
        this.roleArn = $.roleArn;
        this.sourceS3Path = $.sourceS3Path;
        this.status = $.status;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QuerySuggestionsBlockListState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QuerySuggestionsBlockListState $;

        public Builder() {
            $ = new QuerySuggestionsBlockListState();
        }

        public Builder(QuerySuggestionsBlockListState defaults) {
            $ = new QuerySuggestionsBlockListState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn ARN of the block list.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn ARN of the block list.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param description Description for a block list.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description for a block list.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param indexId Identifier of the index for a block list.
         * 
         * @return builder
         * 
         */
        public Builder indexId(@Nullable Output<String> indexId) {
            $.indexId = indexId;
            return this;
        }

        /**
         * @param indexId Identifier of the index for a block list.
         * 
         * @return builder
         * 
         */
        public Builder indexId(String indexId) {
            return indexId(Output.of(indexId));
        }

        /**
         * @param name Name for the block list.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name for the block list.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param querySuggestionsBlockListId Unique identifier of the block list.
         * 
         * @return builder
         * 
         */
        public Builder querySuggestionsBlockListId(@Nullable Output<String> querySuggestionsBlockListId) {
            $.querySuggestionsBlockListId = querySuggestionsBlockListId;
            return this;
        }

        /**
         * @param querySuggestionsBlockListId Unique identifier of the block list.
         * 
         * @return builder
         * 
         */
        public Builder querySuggestionsBlockListId(String querySuggestionsBlockListId) {
            return querySuggestionsBlockListId(Output.of(querySuggestionsBlockListId));
        }

        /**
         * @param roleArn IAM (Identity and Access Management) role used to access the block list text file in S3.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(@Nullable Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn IAM (Identity and Access Management) role used to access the block list text file in S3.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        /**
         * @param sourceS3Path S3 path where your block list text file is located. See details below.
         * 
         * The `source_s3_path` configuration block supports the following arguments:
         * 
         * @return builder
         * 
         */
        public Builder sourceS3Path(@Nullable Output<QuerySuggestionsBlockListSourceS3PathArgs> sourceS3Path) {
            $.sourceS3Path = sourceS3Path;
            return this;
        }

        /**
         * @param sourceS3Path S3 path where your block list text file is located. See details below.
         * 
         * The `source_s3_path` configuration block supports the following arguments:
         * 
         * @return builder
         * 
         */
        public Builder sourceS3Path(QuerySuggestionsBlockListSourceS3PathArgs sourceS3Path) {
            return sourceS3Path(Output.of(sourceS3Path));
        }

        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll Map of tags assigned to the resource, including those inherited from the provider&#39;s default_tags configuration block.
         * 
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
         * @param tagsAll Map of tags assigned to the resource, including those inherited from the provider&#39;s default_tags configuration block.
         * 
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

        public QuerySuggestionsBlockListState build() {
            return $;
        }
    }

}
