// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.resourceexplorer.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ViewIncludedPropertyArgs extends com.pulumi.resources.ResourceArgs {

    public static final ViewIncludedPropertyArgs Empty = new ViewIncludedPropertyArgs();

    /**
     * The name of the property that is included in this view. Valid values: `tags`.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the property that is included in this view. Valid values: `tags`.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    private ViewIncludedPropertyArgs() {}

    private ViewIncludedPropertyArgs(ViewIncludedPropertyArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ViewIncludedPropertyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ViewIncludedPropertyArgs $;

        public Builder() {
            $ = new ViewIncludedPropertyArgs();
        }

        public Builder(ViewIncludedPropertyArgs defaults) {
            $ = new ViewIncludedPropertyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the property that is included in this view. Valid values: `tags`.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the property that is included in this view. Valid values: `tags`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public ViewIncludedPropertyArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
