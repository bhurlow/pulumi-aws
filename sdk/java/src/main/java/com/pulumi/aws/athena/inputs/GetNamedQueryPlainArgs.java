// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.athena.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetNamedQueryPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNamedQueryPlainArgs Empty = new GetNamedQueryPlainArgs();

    /**
     * The plain language name for the query. Maximum length of 128.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return The plain language name for the query. Maximum length of 128.
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * The workgroup to which the query belongs. Defaults to `primary`.
     * 
     */
    @Import(name="workgroup")
    private @Nullable String workgroup;

    /**
     * @return The workgroup to which the query belongs. Defaults to `primary`.
     * 
     */
    public Optional<String> workgroup() {
        return Optional.ofNullable(this.workgroup);
    }

    private GetNamedQueryPlainArgs() {}

    private GetNamedQueryPlainArgs(GetNamedQueryPlainArgs $) {
        this.name = $.name;
        this.workgroup = $.workgroup;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNamedQueryPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNamedQueryPlainArgs $;

        public Builder() {
            $ = new GetNamedQueryPlainArgs();
        }

        public Builder(GetNamedQueryPlainArgs defaults) {
            $ = new GetNamedQueryPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The plain language name for the query. Maximum length of 128.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        /**
         * @param workgroup The workgroup to which the query belongs. Defaults to `primary`.
         * 
         * @return builder
         * 
         */
        public Builder workgroup(@Nullable String workgroup) {
            $.workgroup = workgroup;
            return this;
        }

        public GetNamedQueryPlainArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
