// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.transfer.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class AccessHomeDirectoryMappingArgs extends com.pulumi.resources.ResourceArgs {

    public static final AccessHomeDirectoryMappingArgs Empty = new AccessHomeDirectoryMappingArgs();

    /**
     * Represents an entry and a target.
     * 
     */
    @Import(name="entry", required=true)
    private Output<String> entry;

    /**
     * @return Represents an entry and a target.
     * 
     */
    public Output<String> entry() {
        return this.entry;
    }

    /**
     * Represents the map target.
     * 
     */
    @Import(name="target", required=true)
    private Output<String> target;

    /**
     * @return Represents the map target.
     * 
     */
    public Output<String> target() {
        return this.target;
    }

    private AccessHomeDirectoryMappingArgs() {}

    private AccessHomeDirectoryMappingArgs(AccessHomeDirectoryMappingArgs $) {
        this.entry = $.entry;
        this.target = $.target;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccessHomeDirectoryMappingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccessHomeDirectoryMappingArgs $;

        public Builder() {
            $ = new AccessHomeDirectoryMappingArgs();
        }

        public Builder(AccessHomeDirectoryMappingArgs defaults) {
            $ = new AccessHomeDirectoryMappingArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param entry Represents an entry and a target.
         * 
         * @return builder
         * 
         */
        public Builder entry(Output<String> entry) {
            $.entry = entry;
            return this;
        }

        /**
         * @param entry Represents an entry and a target.
         * 
         * @return builder
         * 
         */
        public Builder entry(String entry) {
            return entry(Output.of(entry));
        }

        /**
         * @param target Represents the map target.
         * 
         * @return builder
         * 
         */
        public Builder target(Output<String> target) {
            $.target = target;
            return this;
        }

        /**
         * @param target Represents the map target.
         * 
         * @return builder
         * 
         */
        public Builder target(String target) {
            return target(Output.of(target));
        }

        public AccessHomeDirectoryMappingArgs build() {
            $.entry = Objects.requireNonNull($.entry, "expected parameter 'entry' to be non-null");
            $.target = Objects.requireNonNull($.target, "expected parameter 'target' to be non-null");
            return $;
        }
    }

}