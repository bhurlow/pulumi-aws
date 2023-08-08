// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.auditmanager.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class FrameworkControlSetControlArgs extends com.pulumi.resources.ResourceArgs {

    public static final FrameworkControlSetControlArgs Empty = new FrameworkControlSetControlArgs();

    /**
     * Unique identifier of the control.
     * 
     */
    @Import(name="id", required=true)
    private Output<String> id;

    /**
     * @return Unique identifier of the control.
     * 
     */
    public Output<String> id() {
        return this.id;
    }

    private FrameworkControlSetControlArgs() {}

    private FrameworkControlSetControlArgs(FrameworkControlSetControlArgs $) {
        this.id = $.id;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FrameworkControlSetControlArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FrameworkControlSetControlArgs $;

        public Builder() {
            $ = new FrameworkControlSetControlArgs();
        }

        public Builder(FrameworkControlSetControlArgs defaults) {
            $ = new FrameworkControlSetControlArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id Unique identifier of the control.
         * 
         * @return builder
         * 
         */
        public Builder id(Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id Unique identifier of the control.
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        public FrameworkControlSetControlArgs build() {
            $.id = Objects.requireNonNull($.id, "expected parameter 'id' to be non-null");
            return $;
        }
    }

}
