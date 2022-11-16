// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.waf.inputs;

import com.pulumi.aws.waf.inputs.RegexMatchSetRegexMatchTupleFieldToMatchArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class RegexMatchSetRegexMatchTupleArgs extends com.pulumi.resources.ResourceArgs {

    public static final RegexMatchSetRegexMatchTupleArgs Empty = new RegexMatchSetRegexMatchTupleArgs();

    /**
     * The part of a web request that you want to search, such as a specified header or a query string.
     * 
     */
    @Import(name="fieldToMatch", required=true)
    private Output<RegexMatchSetRegexMatchTupleFieldToMatchArgs> fieldToMatch;

    /**
     * @return The part of a web request that you want to search, such as a specified header or a query string.
     * 
     */
    public Output<RegexMatchSetRegexMatchTupleFieldToMatchArgs> fieldToMatch() {
        return this.fieldToMatch;
    }

    /**
     * The ID of a Regex Pattern Set.
     * 
     */
    @Import(name="regexPatternSetId", required=true)
    private Output<String> regexPatternSetId;

    /**
     * @return The ID of a Regex Pattern Set.
     * 
     */
    public Output<String> regexPatternSetId() {
        return this.regexPatternSetId;
    }

    /**
     * Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
     * e.g., `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
     * See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_ByteMatchTuple.html#WAF-Type-ByteMatchTuple-TextTransformation)
     * for all supported values.
     * 
     */
    @Import(name="textTransformation", required=true)
    private Output<String> textTransformation;

    /**
     * @return Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
     * e.g., `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
     * See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_ByteMatchTuple.html#WAF-Type-ByteMatchTuple-TextTransformation)
     * for all supported values.
     * 
     */
    public Output<String> textTransformation() {
        return this.textTransformation;
    }

    private RegexMatchSetRegexMatchTupleArgs() {}

    private RegexMatchSetRegexMatchTupleArgs(RegexMatchSetRegexMatchTupleArgs $) {
        this.fieldToMatch = $.fieldToMatch;
        this.regexPatternSetId = $.regexPatternSetId;
        this.textTransformation = $.textTransformation;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RegexMatchSetRegexMatchTupleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RegexMatchSetRegexMatchTupleArgs $;

        public Builder() {
            $ = new RegexMatchSetRegexMatchTupleArgs();
        }

        public Builder(RegexMatchSetRegexMatchTupleArgs defaults) {
            $ = new RegexMatchSetRegexMatchTupleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param fieldToMatch The part of a web request that you want to search, such as a specified header or a query string.
         * 
         * @return builder
         * 
         */
        public Builder fieldToMatch(Output<RegexMatchSetRegexMatchTupleFieldToMatchArgs> fieldToMatch) {
            $.fieldToMatch = fieldToMatch;
            return this;
        }

        /**
         * @param fieldToMatch The part of a web request that you want to search, such as a specified header or a query string.
         * 
         * @return builder
         * 
         */
        public Builder fieldToMatch(RegexMatchSetRegexMatchTupleFieldToMatchArgs fieldToMatch) {
            return fieldToMatch(Output.of(fieldToMatch));
        }

        /**
         * @param regexPatternSetId The ID of a Regex Pattern Set.
         * 
         * @return builder
         * 
         */
        public Builder regexPatternSetId(Output<String> regexPatternSetId) {
            $.regexPatternSetId = regexPatternSetId;
            return this;
        }

        /**
         * @param regexPatternSetId The ID of a Regex Pattern Set.
         * 
         * @return builder
         * 
         */
        public Builder regexPatternSetId(String regexPatternSetId) {
            return regexPatternSetId(Output.of(regexPatternSetId));
        }

        /**
         * @param textTransformation Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
         * e.g., `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
         * See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_ByteMatchTuple.html#WAF-Type-ByteMatchTuple-TextTransformation)
         * for all supported values.
         * 
         * @return builder
         * 
         */
        public Builder textTransformation(Output<String> textTransformation) {
            $.textTransformation = textTransformation;
            return this;
        }

        /**
         * @param textTransformation Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
         * e.g., `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
         * See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_ByteMatchTuple.html#WAF-Type-ByteMatchTuple-TextTransformation)
         * for all supported values.
         * 
         * @return builder
         * 
         */
        public Builder textTransformation(String textTransformation) {
            return textTransformation(Output.of(textTransformation));
        }

        public RegexMatchSetRegexMatchTupleArgs build() {
            $.fieldToMatch = Objects.requireNonNull($.fieldToMatch, "expected parameter 'fieldToMatch' to be non-null");
            $.regexPatternSetId = Objects.requireNonNull($.regexPatternSetId, "expected parameter 'regexPatternSetId' to be non-null");
            $.textTransformation = Objects.requireNonNull($.textTransformation, "expected parameter 'textTransformation' to be non-null");
            return $;
        }
    }

}
