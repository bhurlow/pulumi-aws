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


public final class VocabularyFilterState extends com.pulumi.resources.ResourceArgs {

    public static final VocabularyFilterState Empty = new VocabularyFilterState();

    /**
     * ARN of the VocabularyFilter.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return ARN of the VocabularyFilter.
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
     * The language code you selected for your vocabulary filter. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
     * 
     */
    @Import(name="languageCode")
    private @Nullable Output<String> languageCode;

    /**
     * @return The language code you selected for your vocabulary filter. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
     * 
     */
    public Optional<Output<String>> languageCode() {
        return Optional.ofNullable(this.languageCode);
    }

    /**
     * A map of tags to assign to the VocabularyFilter. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the VocabularyFilter. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * The Amazon S3 location (URI) of the text file that contains your custom VocabularyFilter. Conflicts with `words` argument.
     * 
     */
    @Import(name="vocabularyFilterFileUri")
    private @Nullable Output<String> vocabularyFilterFileUri;

    /**
     * @return The Amazon S3 location (URI) of the text file that contains your custom VocabularyFilter. Conflicts with `words` argument.
     * 
     */
    public Optional<Output<String>> vocabularyFilterFileUri() {
        return Optional.ofNullable(this.vocabularyFilterFileUri);
    }

    /**
     * The name of the VocabularyFilter.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="vocabularyFilterName")
    private @Nullable Output<String> vocabularyFilterName;

    /**
     * @return The name of the VocabularyFilter.
     * 
     * The following arguments are optional:
     * 
     */
    public Optional<Output<String>> vocabularyFilterName() {
        return Optional.ofNullable(this.vocabularyFilterName);
    }

    /**
     * A list of terms to include in the vocabulary. Conflicts with `vocabulary_filter_file_uri` argument.
     * 
     */
    @Import(name="words")
    private @Nullable Output<List<String>> words;

    /**
     * @return A list of terms to include in the vocabulary. Conflicts with `vocabulary_filter_file_uri` argument.
     * 
     */
    public Optional<Output<List<String>>> words() {
        return Optional.ofNullable(this.words);
    }

    private VocabularyFilterState() {}

    private VocabularyFilterState(VocabularyFilterState $) {
        this.arn = $.arn;
        this.downloadUri = $.downloadUri;
        this.languageCode = $.languageCode;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.vocabularyFilterFileUri = $.vocabularyFilterFileUri;
        this.vocabularyFilterName = $.vocabularyFilterName;
        this.words = $.words;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VocabularyFilterState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VocabularyFilterState $;

        public Builder() {
            $ = new VocabularyFilterState();
        }

        public Builder(VocabularyFilterState defaults) {
            $ = new VocabularyFilterState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn ARN of the VocabularyFilter.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn ARN of the VocabularyFilter.
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
         * @param languageCode The language code you selected for your vocabulary filter. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
         * 
         * @return builder
         * 
         */
        public Builder languageCode(@Nullable Output<String> languageCode) {
            $.languageCode = languageCode;
            return this;
        }

        /**
         * @param languageCode The language code you selected for your vocabulary filter. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
         * 
         * @return builder
         * 
         */
        public Builder languageCode(String languageCode) {
            return languageCode(Output.of(languageCode));
        }

        /**
         * @param tags A map of tags to assign to the VocabularyFilter. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the VocabularyFilter. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
         * @param vocabularyFilterFileUri The Amazon S3 location (URI) of the text file that contains your custom VocabularyFilter. Conflicts with `words` argument.
         * 
         * @return builder
         * 
         */
        public Builder vocabularyFilterFileUri(@Nullable Output<String> vocabularyFilterFileUri) {
            $.vocabularyFilterFileUri = vocabularyFilterFileUri;
            return this;
        }

        /**
         * @param vocabularyFilterFileUri The Amazon S3 location (URI) of the text file that contains your custom VocabularyFilter. Conflicts with `words` argument.
         * 
         * @return builder
         * 
         */
        public Builder vocabularyFilterFileUri(String vocabularyFilterFileUri) {
            return vocabularyFilterFileUri(Output.of(vocabularyFilterFileUri));
        }

        /**
         * @param vocabularyFilterName The name of the VocabularyFilter.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder vocabularyFilterName(@Nullable Output<String> vocabularyFilterName) {
            $.vocabularyFilterName = vocabularyFilterName;
            return this;
        }

        /**
         * @param vocabularyFilterName The name of the VocabularyFilter.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder vocabularyFilterName(String vocabularyFilterName) {
            return vocabularyFilterName(Output.of(vocabularyFilterName));
        }

        /**
         * @param words A list of terms to include in the vocabulary. Conflicts with `vocabulary_filter_file_uri` argument.
         * 
         * @return builder
         * 
         */
        public Builder words(@Nullable Output<List<String>> words) {
            $.words = words;
            return this;
        }

        /**
         * @param words A list of terms to include in the vocabulary. Conflicts with `vocabulary_filter_file_uri` argument.
         * 
         * @return builder
         * 
         */
        public Builder words(List<String> words) {
            return words(Output.of(words));
        }

        /**
         * @param words A list of terms to include in the vocabulary. Conflicts with `vocabulary_filter_file_uri` argument.
         * 
         * @return builder
         * 
         */
        public Builder words(String... words) {
            return words(List.of(words));
        }

        public VocabularyFilterState build() {
            return $;
        }
    }

}
