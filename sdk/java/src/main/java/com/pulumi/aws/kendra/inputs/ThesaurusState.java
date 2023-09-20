// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kendra.inputs;

import com.pulumi.aws.kendra.inputs.ThesaurusSourceS3PathArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ThesaurusState extends com.pulumi.resources.ResourceArgs {

    public static final ThesaurusState Empty = new ThesaurusState();

    /**
     * ARN of the thesaurus.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return ARN of the thesaurus.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The description for a thesaurus.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description for a thesaurus.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The identifier of the index for a thesaurus.
     * 
     */
    @Import(name="indexId")
    private @Nullable Output<String> indexId;

    /**
     * @return The identifier of the index for a thesaurus.
     * 
     */
    public Optional<Output<String>> indexId() {
        return Optional.ofNullable(this.indexId);
    }

    /**
     * The name for the thesaurus.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name for the thesaurus.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The IAM (Identity and Access Management) role used to access the thesaurus file in S3.
     * 
     */
    @Import(name="roleArn")
    private @Nullable Output<String> roleArn;

    /**
     * @return The IAM (Identity and Access Management) role used to access the thesaurus file in S3.
     * 
     */
    public Optional<Output<String>> roleArn() {
        return Optional.ofNullable(this.roleArn);
    }

    /**
     * The S3 path where your thesaurus file sits in S3. Detailed below.
     * 
     * The `source_s3_path` configuration block supports the following arguments:
     * 
     */
    @Import(name="sourceS3Path")
    private @Nullable Output<ThesaurusSourceS3PathArgs> sourceS3Path;

    /**
     * @return The S3 path where your thesaurus file sits in S3. Detailed below.
     * 
     * The `source_s3_path` configuration block supports the following arguments:
     * 
     */
    public Optional<Output<ThesaurusSourceS3PathArgs>> sourceS3Path() {
        return Optional.ofNullable(this.sourceS3Path);
    }

    /**
     * The current status of the thesaurus.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The current status of the thesaurus.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    @Import(name="thesaurusId")
    private @Nullable Output<String> thesaurusId;

    public Optional<Output<String>> thesaurusId() {
        return Optional.ofNullable(this.thesaurusId);
    }

    private ThesaurusState() {}

    private ThesaurusState(ThesaurusState $) {
        this.arn = $.arn;
        this.description = $.description;
        this.indexId = $.indexId;
        this.name = $.name;
        this.roleArn = $.roleArn;
        this.sourceS3Path = $.sourceS3Path;
        this.status = $.status;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.thesaurusId = $.thesaurusId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ThesaurusState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ThesaurusState $;

        public Builder() {
            $ = new ThesaurusState();
        }

        public Builder(ThesaurusState defaults) {
            $ = new ThesaurusState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn ARN of the thesaurus.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn ARN of the thesaurus.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param description The description for a thesaurus.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description for a thesaurus.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param indexId The identifier of the index for a thesaurus.
         * 
         * @return builder
         * 
         */
        public Builder indexId(@Nullable Output<String> indexId) {
            $.indexId = indexId;
            return this;
        }

        /**
         * @param indexId The identifier of the index for a thesaurus.
         * 
         * @return builder
         * 
         */
        public Builder indexId(String indexId) {
            return indexId(Output.of(indexId));
        }

        /**
         * @param name The name for the thesaurus.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name for the thesaurus.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param roleArn The IAM (Identity and Access Management) role used to access the thesaurus file in S3.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(@Nullable Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn The IAM (Identity and Access Management) role used to access the thesaurus file in S3.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        /**
         * @param sourceS3Path The S3 path where your thesaurus file sits in S3. Detailed below.
         * 
         * The `source_s3_path` configuration block supports the following arguments:
         * 
         * @return builder
         * 
         */
        public Builder sourceS3Path(@Nullable Output<ThesaurusSourceS3PathArgs> sourceS3Path) {
            $.sourceS3Path = sourceS3Path;
            return this;
        }

        /**
         * @param sourceS3Path The S3 path where your thesaurus file sits in S3. Detailed below.
         * 
         * The `source_s3_path` configuration block supports the following arguments:
         * 
         * @return builder
         * 
         */
        public Builder sourceS3Path(ThesaurusSourceS3PathArgs sourceS3Path) {
            return sourceS3Path(Output.of(sourceS3Path));
        }

        /**
         * @param status The current status of the thesaurus.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The current status of the thesaurus.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
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
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
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

        public Builder thesaurusId(@Nullable Output<String> thesaurusId) {
            $.thesaurusId = thesaurusId;
            return this;
        }

        public Builder thesaurusId(String thesaurusId) {
            return thesaurusId(Output.of(thesaurusId));
        }

        public ThesaurusState build() {
            return $;
        }
    }

}
