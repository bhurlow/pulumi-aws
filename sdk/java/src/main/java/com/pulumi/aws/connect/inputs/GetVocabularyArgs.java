// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.connect.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVocabularyArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVocabularyArgs Empty = new GetVocabularyArgs();

    /**
     * Reference to the hosting Amazon Connect Instance
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<String> instanceId;

    /**
     * @return Reference to the hosting Amazon Connect Instance
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }

    /**
     * Returns information on a specific Vocabulary by name
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Returns information on a specific Vocabulary by name
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A map of tags to assign to the Vocabulary.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the Vocabulary.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Returns information on a specific Vocabulary by Vocabulary id
     * 
     */
    @Import(name="vocabularyId")
    private @Nullable Output<String> vocabularyId;

    /**
     * @return Returns information on a specific Vocabulary by Vocabulary id
     * 
     */
    public Optional<Output<String>> vocabularyId() {
        return Optional.ofNullable(this.vocabularyId);
    }

    private GetVocabularyArgs() {}

    private GetVocabularyArgs(GetVocabularyArgs $) {
        this.instanceId = $.instanceId;
        this.name = $.name;
        this.tags = $.tags;
        this.vocabularyId = $.vocabularyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVocabularyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVocabularyArgs $;

        public Builder() {
            $ = new GetVocabularyArgs();
        }

        public Builder(GetVocabularyArgs defaults) {
            $ = new GetVocabularyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param instanceId Reference to the hosting Amazon Connect Instance
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId Reference to the hosting Amazon Connect Instance
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param name Returns information on a specific Vocabulary by name
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Returns information on a specific Vocabulary by name
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param tags A map of tags to assign to the Vocabulary.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the Vocabulary.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param vocabularyId Returns information on a specific Vocabulary by Vocabulary id
         * 
         * @return builder
         * 
         */
        public Builder vocabularyId(@Nullable Output<String> vocabularyId) {
            $.vocabularyId = vocabularyId;
            return this;
        }

        /**
         * @param vocabularyId Returns information on a specific Vocabulary by Vocabulary id
         * 
         * @return builder
         * 
         */
        public Builder vocabularyId(String vocabularyId) {
            return vocabularyId(Output.of(vocabularyId));
        }

        public GetVocabularyArgs build() {
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}
