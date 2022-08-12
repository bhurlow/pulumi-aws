// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.lex.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIntentPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIntentPlainArgs Empty = new GetIntentPlainArgs();

    /**
     * The name of the intent. The name is case sensitive.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return The name of the intent. The name is case sensitive.
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * The version of the intent.
     * 
     */
    @Import(name="version")
    private @Nullable String version;

    /**
     * @return The version of the intent.
     * 
     */
    public Optional<String> version() {
        return Optional.ofNullable(this.version);
    }

    private GetIntentPlainArgs() {}

    private GetIntentPlainArgs(GetIntentPlainArgs $) {
        this.name = $.name;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIntentPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIntentPlainArgs $;

        public Builder() {
            $ = new GetIntentPlainArgs();
        }

        public Builder(GetIntentPlainArgs defaults) {
            $ = new GetIntentPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the intent. The name is case sensitive.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        /**
         * @param version The version of the intent.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable String version) {
            $.version = version;
            return this;
        }

        public GetIntentPlainArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}