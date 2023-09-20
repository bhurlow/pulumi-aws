// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.transcribe.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VocabularyState extends com.pulumi.resources.ResourceArgs {

    public static final VocabularyState Empty = new VocabularyState();

    /**
     * ARN of the Vocabulary.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return ARN of the Vocabulary.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Generated download URI.
     * 
     */
    @Import(name="downloadUri")
    private @Nullable Output<String> downloadUri;

    /**
     * @return Generated download URI.
     * 
     */
    public Optional<Output<String>> downloadUri() {
        return Optional.ofNullable(this.downloadUri);
    }

    /**
     * The language code you selected for your vocabulary.
     * 
     */
    @Import(name="languageCode")
    private @Nullable Output<String> languageCode;

    /**
     * @return The language code you selected for your vocabulary.
     * 
     */
    public Optional<Output<String>> languageCode() {
        return Optional.ofNullable(this.languageCode);
    }

    /**
     * A list of terms to include in the vocabulary. Conflicts with `vocabulary_file_uri`
     * 
     */
    @Import(name="phrases")
    private @Nullable Output<List<String>> phrases;

    /**
     * @return A list of terms to include in the vocabulary. Conflicts with `vocabulary_file_uri`
     * 
     */
    public Optional<Output<List<String>>> phrases() {
        return Optional.ofNullable(this.phrases);
    }

    /**
     * A map of tags to assign to the Vocabulary. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the Vocabulary. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

    /**
     * The Amazon S3 location (URI) of the text file that contains your custom vocabulary. Conflicts wth `phrases`.
     * 
     */
    @Import(name="vocabularyFileUri")
    private @Nullable Output<String> vocabularyFileUri;

    /**
     * @return The Amazon S3 location (URI) of the text file that contains your custom vocabulary. Conflicts wth `phrases`.
     * 
     */
    public Optional<Output<String>> vocabularyFileUri() {
        return Optional.ofNullable(this.vocabularyFileUri);
    }

    /**
     * The name of the Vocabulary.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="vocabularyName")
    private @Nullable Output<String> vocabularyName;

    /**
     * @return The name of the Vocabulary.
     * 
     * The following arguments are optional:
     * 
     */
    public Optional<Output<String>> vocabularyName() {
        return Optional.ofNullable(this.vocabularyName);
    }

    private VocabularyState() {}

    private VocabularyState(VocabularyState $) {
        this.arn = $.arn;
        this.downloadUri = $.downloadUri;
        this.languageCode = $.languageCode;
        this.phrases = $.phrases;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.vocabularyFileUri = $.vocabularyFileUri;
        this.vocabularyName = $.vocabularyName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VocabularyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VocabularyState $;

        public Builder() {
            $ = new VocabularyState();
        }

        public Builder(VocabularyState defaults) {
            $ = new VocabularyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn ARN of the Vocabulary.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn ARN of the Vocabulary.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param downloadUri Generated download URI.
         * 
         * @return builder
         * 
         */
        public Builder downloadUri(@Nullable Output<String> downloadUri) {
            $.downloadUri = downloadUri;
            return this;
        }

        /**
         * @param downloadUri Generated download URI.
         * 
         * @return builder
         * 
         */
        public Builder downloadUri(String downloadUri) {
            return downloadUri(Output.of(downloadUri));
        }

        /**
         * @param languageCode The language code you selected for your vocabulary.
         * 
         * @return builder
         * 
         */
        public Builder languageCode(@Nullable Output<String> languageCode) {
            $.languageCode = languageCode;
            return this;
        }

        /**
         * @param languageCode The language code you selected for your vocabulary.
         * 
         * @return builder
         * 
         */
        public Builder languageCode(String languageCode) {
            return languageCode(Output.of(languageCode));
        }

        /**
         * @param phrases A list of terms to include in the vocabulary. Conflicts with `vocabulary_file_uri`
         * 
         * @return builder
         * 
         */
        public Builder phrases(@Nullable Output<List<String>> phrases) {
            $.phrases = phrases;
            return this;
        }

        /**
         * @param phrases A list of terms to include in the vocabulary. Conflicts with `vocabulary_file_uri`
         * 
         * @return builder
         * 
         */
        public Builder phrases(List<String> phrases) {
            return phrases(Output.of(phrases));
        }

        /**
         * @param phrases A list of terms to include in the vocabulary. Conflicts with `vocabulary_file_uri`
         * 
         * @return builder
         * 
         */
        public Builder phrases(String... phrases) {
            return phrases(List.of(phrases));
        }

        /**
         * @param tags A map of tags to assign to the Vocabulary. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the Vocabulary. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        /**
         * @param vocabularyFileUri The Amazon S3 location (URI) of the text file that contains your custom vocabulary. Conflicts wth `phrases`.
         * 
         * @return builder
         * 
         */
        public Builder vocabularyFileUri(@Nullable Output<String> vocabularyFileUri) {
            $.vocabularyFileUri = vocabularyFileUri;
            return this;
        }

        /**
         * @param vocabularyFileUri The Amazon S3 location (URI) of the text file that contains your custom vocabulary. Conflicts wth `phrases`.
         * 
         * @return builder
         * 
         */
        public Builder vocabularyFileUri(String vocabularyFileUri) {
            return vocabularyFileUri(Output.of(vocabularyFileUri));
        }

        /**
         * @param vocabularyName The name of the Vocabulary.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder vocabularyName(@Nullable Output<String> vocabularyName) {
            $.vocabularyName = vocabularyName;
            return this;
        }

        /**
         * @param vocabularyName The name of the Vocabulary.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder vocabularyName(String vocabularyName) {
            return vocabularyName(Output.of(vocabularyName));
        }

        public VocabularyState build() {
            return $;
        }
    }

}
