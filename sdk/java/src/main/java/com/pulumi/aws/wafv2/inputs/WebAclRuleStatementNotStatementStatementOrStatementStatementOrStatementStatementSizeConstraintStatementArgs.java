// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.wafv2.inputs;

import com.pulumi.aws.wafv2.inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchArgs;
import com.pulumi.aws.wafv2.inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementTextTransformationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementArgs extends com.pulumi.resources.ResourceArgs {

    public static final WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementArgs Empty = new WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementArgs();

    /**
     * Operator to use to compare the request part to the size setting. Valid values include: `EQ`, `NE`, `LE`, `LT`, `GE`, or `GT`.
     * 
     */
    @Import(name="comparisonOperator", required=true)
    private Output<String> comparisonOperator;

    /**
     * @return Operator to use to compare the request part to the size setting. Valid values include: `EQ`, `NE`, `LE`, `LT`, `GE`, or `GT`.
     * 
     */
    public Output<String> comparisonOperator() {
        return this.comparisonOperator;
    }

    /**
     * Part of a web request that you want AWS WAF to inspect. See Field to Match below for details.
     * 
     */
    @Import(name="fieldToMatch")
    private @Nullable Output<WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchArgs> fieldToMatch;

    /**
     * @return Part of a web request that you want AWS WAF to inspect. See Field to Match below for details.
     * 
     */
    public Optional<Output<WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchArgs>> fieldToMatch() {
        return Optional.ofNullable(this.fieldToMatch);
    }

    /**
     * Size, in bytes, to compare to the request part, after any transformations. Valid values are integers between 0 and 21474836480, inclusive.
     * 
     */
    @Import(name="size", required=true)
    private Output<Integer> size;

    /**
     * @return Size, in bytes, to compare to the request part, after any transformations. Valid values are integers between 0 and 21474836480, inclusive.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }

    /**
     * Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection. See Text Transformation below for details.
     * 
     */
    @Import(name="textTransformations", required=true)
    private Output<List<WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementTextTransformationArgs>> textTransformations;

    /**
     * @return Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection. See Text Transformation below for details.
     * 
     */
    public Output<List<WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementTextTransformationArgs>> textTransformations() {
        return this.textTransformations;
    }

    private WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementArgs() {}

    private WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementArgs(WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementArgs $) {
        this.comparisonOperator = $.comparisonOperator;
        this.fieldToMatch = $.fieldToMatch;
        this.size = $.size;
        this.textTransformations = $.textTransformations;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementArgs $;

        public Builder() {
            $ = new WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementArgs();
        }

        public Builder(WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementArgs defaults) {
            $ = new WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param comparisonOperator Operator to use to compare the request part to the size setting. Valid values include: `EQ`, `NE`, `LE`, `LT`, `GE`, or `GT`.
         * 
         * @return builder
         * 
         */
        public Builder comparisonOperator(Output<String> comparisonOperator) {
            $.comparisonOperator = comparisonOperator;
            return this;
        }

        /**
         * @param comparisonOperator Operator to use to compare the request part to the size setting. Valid values include: `EQ`, `NE`, `LE`, `LT`, `GE`, or `GT`.
         * 
         * @return builder
         * 
         */
        public Builder comparisonOperator(String comparisonOperator) {
            return comparisonOperator(Output.of(comparisonOperator));
        }

        /**
         * @param fieldToMatch Part of a web request that you want AWS WAF to inspect. See Field to Match below for details.
         * 
         * @return builder
         * 
         */
        public Builder fieldToMatch(@Nullable Output<WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchArgs> fieldToMatch) {
            $.fieldToMatch = fieldToMatch;
            return this;
        }

        /**
         * @param fieldToMatch Part of a web request that you want AWS WAF to inspect. See Field to Match below for details.
         * 
         * @return builder
         * 
         */
        public Builder fieldToMatch(WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchArgs fieldToMatch) {
            return fieldToMatch(Output.of(fieldToMatch));
        }

        /**
         * @param size Size, in bytes, to compare to the request part, after any transformations. Valid values are integers between 0 and 21474836480, inclusive.
         * 
         * @return builder
         * 
         */
        public Builder size(Output<Integer> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size Size, in bytes, to compare to the request part, after any transformations. Valid values are integers between 0 and 21474836480, inclusive.
         * 
         * @return builder
         * 
         */
        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        /**
         * @param textTransformations Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection. See Text Transformation below for details.
         * 
         * @return builder
         * 
         */
        public Builder textTransformations(Output<List<WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementTextTransformationArgs>> textTransformations) {
            $.textTransformations = textTransformations;
            return this;
        }

        /**
         * @param textTransformations Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection. See Text Transformation below for details.
         * 
         * @return builder
         * 
         */
        public Builder textTransformations(List<WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementTextTransformationArgs> textTransformations) {
            return textTransformations(Output.of(textTransformations));
        }

        /**
         * @param textTransformations Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection. See Text Transformation below for details.
         * 
         * @return builder
         * 
         */
        public Builder textTransformations(WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementTextTransformationArgs... textTransformations) {
            return textTransformations(List.of(textTransformations));
        }

        public WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementArgs build() {
            $.comparisonOperator = Objects.requireNonNull($.comparisonOperator, "expected parameter 'comparisonOperator' to be non-null");
            $.size = Objects.requireNonNull($.size, "expected parameter 'size' to be non-null");
            $.textTransformations = Objects.requireNonNull($.textTransformations, "expected parameter 'textTransformations' to be non-null");
            return $;
        }
    }

}