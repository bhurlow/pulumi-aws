// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.macie;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CustomDataIdentifierArgs extends com.pulumi.resources.ResourceArgs {

    public static final CustomDataIdentifierArgs Empty = new CustomDataIdentifierArgs();

    /**
     * A custom description of the custom data identifier. The description can contain as many as 512 characters.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A custom description of the custom data identifier. The description can contain as many as 512 characters.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * An array that lists specific character sequences (ignore words) to exclude from the results. If the text matched by the regular expression is the same as any string in this array, Amazon Macie ignores it. The array can contain as many as 10 ignore words. Each ignore word can contain 4 - 90 characters. Ignore words are case sensitive.
     * 
     */
    @Import(name="ignoreWords")
    private @Nullable Output<List<String>> ignoreWords;

    /**
     * @return An array that lists specific character sequences (ignore words) to exclude from the results. If the text matched by the regular expression is the same as any string in this array, Amazon Macie ignores it. The array can contain as many as 10 ignore words. Each ignore word can contain 4 - 90 characters. Ignore words are case sensitive.
     * 
     */
    public Optional<Output<List<String>>> ignoreWords() {
        return Optional.ofNullable(this.ignoreWords);
    }

    /**
     * An array that lists specific character sequences (keywords), one of which must be within proximity (`maximum_match_distance`) of the regular expression to match. The array can contain as many as 50 keywords. Each keyword can contain 3 - 90 characters. Keywords aren&#39;t case sensitive.
     * 
     */
    @Import(name="keywords")
    private @Nullable Output<List<String>> keywords;

    /**
     * @return An array that lists specific character sequences (keywords), one of which must be within proximity (`maximum_match_distance`) of the regular expression to match. The array can contain as many as 50 keywords. Each keyword can contain 3 - 90 characters. Keywords aren&#39;t case sensitive.
     * 
     */
    public Optional<Output<List<String>>> keywords() {
        return Optional.ofNullable(this.keywords);
    }

    /**
     * The maximum number of characters that can exist between text that matches the regex pattern and the character sequences specified by the keywords array. Macie includes or excludes a result based on the proximity of a keyword to text that matches the regex pattern. The distance can be 1 - 300 characters. The default value is 50.
     * 
     */
    @Import(name="maximumMatchDistance")
    private @Nullable Output<Integer> maximumMatchDistance;

    /**
     * @return The maximum number of characters that can exist between text that matches the regex pattern and the character sequences specified by the keywords array. Macie includes or excludes a result based on the proximity of a keyword to text that matches the regex pattern. The distance can be 1 - 300 characters. The default value is 50.
     * 
     */
    public Optional<Output<Integer>> maximumMatchDistance() {
        return Optional.ofNullable(this.maximumMatchDistance);
    }

    /**
     * A custom name for the custom data identifier. The name can contain as many as 128 characters. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A custom name for the custom data identifier. The name can contain as many as 128 characters. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     * 
     */
    @Import(name="namePrefix")
    private @Nullable Output<String> namePrefix;

    /**
     * @return Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     * 
     */
    public Optional<Output<String>> namePrefix() {
        return Optional.ofNullable(this.namePrefix);
    }

    /**
     * The regular expression (regex) that defines the pattern to match. The expression can contain as many as 512 characters.
     * 
     */
    @Import(name="regex")
    private @Nullable Output<String> regex;

    /**
     * @return The regular expression (regex) that defines the pattern to match. The expression can contain as many as 512 characters.
     * 
     */
    public Optional<Output<String>> regex() {
        return Optional.ofNullable(this.regex);
    }

    /**
     * A map of key-value pairs that specifies the tags to associate with the custom data identifier.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of key-value pairs that specifies the tags to associate with the custom data identifier.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private CustomDataIdentifierArgs() {}

    private CustomDataIdentifierArgs(CustomDataIdentifierArgs $) {
        this.description = $.description;
        this.ignoreWords = $.ignoreWords;
        this.keywords = $.keywords;
        this.maximumMatchDistance = $.maximumMatchDistance;
        this.name = $.name;
        this.namePrefix = $.namePrefix;
        this.regex = $.regex;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CustomDataIdentifierArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CustomDataIdentifierArgs $;

        public Builder() {
            $ = new CustomDataIdentifierArgs();
        }

        public Builder(CustomDataIdentifierArgs defaults) {
            $ = new CustomDataIdentifierArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description A custom description of the custom data identifier. The description can contain as many as 512 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A custom description of the custom data identifier. The description can contain as many as 512 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param ignoreWords An array that lists specific character sequences (ignore words) to exclude from the results. If the text matched by the regular expression is the same as any string in this array, Amazon Macie ignores it. The array can contain as many as 10 ignore words. Each ignore word can contain 4 - 90 characters. Ignore words are case sensitive.
         * 
         * @return builder
         * 
         */
        public Builder ignoreWords(@Nullable Output<List<String>> ignoreWords) {
            $.ignoreWords = ignoreWords;
            return this;
        }

        /**
         * @param ignoreWords An array that lists specific character sequences (ignore words) to exclude from the results. If the text matched by the regular expression is the same as any string in this array, Amazon Macie ignores it. The array can contain as many as 10 ignore words. Each ignore word can contain 4 - 90 characters. Ignore words are case sensitive.
         * 
         * @return builder
         * 
         */
        public Builder ignoreWords(List<String> ignoreWords) {
            return ignoreWords(Output.of(ignoreWords));
        }

        /**
         * @param ignoreWords An array that lists specific character sequences (ignore words) to exclude from the results. If the text matched by the regular expression is the same as any string in this array, Amazon Macie ignores it. The array can contain as many as 10 ignore words. Each ignore word can contain 4 - 90 characters. Ignore words are case sensitive.
         * 
         * @return builder
         * 
         */
        public Builder ignoreWords(String... ignoreWords) {
            return ignoreWords(List.of(ignoreWords));
        }

        /**
         * @param keywords An array that lists specific character sequences (keywords), one of which must be within proximity (`maximum_match_distance`) of the regular expression to match. The array can contain as many as 50 keywords. Each keyword can contain 3 - 90 characters. Keywords aren&#39;t case sensitive.
         * 
         * @return builder
         * 
         */
        public Builder keywords(@Nullable Output<List<String>> keywords) {
            $.keywords = keywords;
            return this;
        }

        /**
         * @param keywords An array that lists specific character sequences (keywords), one of which must be within proximity (`maximum_match_distance`) of the regular expression to match. The array can contain as many as 50 keywords. Each keyword can contain 3 - 90 characters. Keywords aren&#39;t case sensitive.
         * 
         * @return builder
         * 
         */
        public Builder keywords(List<String> keywords) {
            return keywords(Output.of(keywords));
        }

        /**
         * @param keywords An array that lists specific character sequences (keywords), one of which must be within proximity (`maximum_match_distance`) of the regular expression to match. The array can contain as many as 50 keywords. Each keyword can contain 3 - 90 characters. Keywords aren&#39;t case sensitive.
         * 
         * @return builder
         * 
         */
        public Builder keywords(String... keywords) {
            return keywords(List.of(keywords));
        }

        /**
         * @param maximumMatchDistance The maximum number of characters that can exist between text that matches the regex pattern and the character sequences specified by the keywords array. Macie includes or excludes a result based on the proximity of a keyword to text that matches the regex pattern. The distance can be 1 - 300 characters. The default value is 50.
         * 
         * @return builder
         * 
         */
        public Builder maximumMatchDistance(@Nullable Output<Integer> maximumMatchDistance) {
            $.maximumMatchDistance = maximumMatchDistance;
            return this;
        }

        /**
         * @param maximumMatchDistance The maximum number of characters that can exist between text that matches the regex pattern and the character sequences specified by the keywords array. Macie includes or excludes a result based on the proximity of a keyword to text that matches the regex pattern. The distance can be 1 - 300 characters. The default value is 50.
         * 
         * @return builder
         * 
         */
        public Builder maximumMatchDistance(Integer maximumMatchDistance) {
            return maximumMatchDistance(Output.of(maximumMatchDistance));
        }

        /**
         * @param name A custom name for the custom data identifier. The name can contain as many as 128 characters. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A custom name for the custom data identifier. The name can contain as many as 128 characters. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namePrefix Creates a unique name beginning with the specified prefix. Conflicts with `name`.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(@Nullable Output<String> namePrefix) {
            $.namePrefix = namePrefix;
            return this;
        }

        /**
         * @param namePrefix Creates a unique name beginning with the specified prefix. Conflicts with `name`.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(String namePrefix) {
            return namePrefix(Output.of(namePrefix));
        }

        /**
         * @param regex The regular expression (regex) that defines the pattern to match. The expression can contain as many as 512 characters.
         * 
         * @return builder
         * 
         */
        public Builder regex(@Nullable Output<String> regex) {
            $.regex = regex;
            return this;
        }

        /**
         * @param regex The regular expression (regex) that defines the pattern to match. The expression can contain as many as 512 characters.
         * 
         * @return builder
         * 
         */
        public Builder regex(String regex) {
            return regex(Output.of(regex));
        }

        /**
         * @param tags A map of key-value pairs that specifies the tags to associate with the custom data identifier.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of key-value pairs that specifies the tags to associate with the custom data identifier.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public CustomDataIdentifierArgs build() {
            return $;
        }
    }

}
