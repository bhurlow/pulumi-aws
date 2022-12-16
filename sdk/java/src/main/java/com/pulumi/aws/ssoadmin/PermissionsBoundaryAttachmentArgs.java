// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ssoadmin;

import com.pulumi.aws.ssoadmin.inputs.PermissionsBoundaryAttachmentPermissionsBoundaryArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class PermissionsBoundaryAttachmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final PermissionsBoundaryAttachmentArgs Empty = new PermissionsBoundaryAttachmentArgs();

    /**
     * The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
     * 
     */
    @Import(name="instanceArn", required=true)
    private Output<String> instanceArn;

    /**
     * @return The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
     * 
     */
    public Output<String> instanceArn() {
        return this.instanceArn;
    }

    /**
     * The Amazon Resource Name (ARN) of the Permission Set.
     * 
     */
    @Import(name="permissionSetArn", required=true)
    private Output<String> permissionSetArn;

    /**
     * @return The Amazon Resource Name (ARN) of the Permission Set.
     * 
     */
    public Output<String> permissionSetArn() {
        return this.permissionSetArn;
    }

    /**
     * The permissions boundary policy. See below.
     * 
     */
    @Import(name="permissionsBoundary", required=true)
    private Output<PermissionsBoundaryAttachmentPermissionsBoundaryArgs> permissionsBoundary;

    /**
     * @return The permissions boundary policy. See below.
     * 
     */
    public Output<PermissionsBoundaryAttachmentPermissionsBoundaryArgs> permissionsBoundary() {
        return this.permissionsBoundary;
    }

    private PermissionsBoundaryAttachmentArgs() {}

    private PermissionsBoundaryAttachmentArgs(PermissionsBoundaryAttachmentArgs $) {
        this.instanceArn = $.instanceArn;
        this.permissionSetArn = $.permissionSetArn;
        this.permissionsBoundary = $.permissionsBoundary;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PermissionsBoundaryAttachmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PermissionsBoundaryAttachmentArgs $;

        public Builder() {
            $ = new PermissionsBoundaryAttachmentArgs();
        }

        public Builder(PermissionsBoundaryAttachmentArgs defaults) {
            $ = new PermissionsBoundaryAttachmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param instanceArn The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
         * 
         * @return builder
         * 
         */
        public Builder instanceArn(Output<String> instanceArn) {
            $.instanceArn = instanceArn;
            return this;
        }

        /**
         * @param instanceArn The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
         * 
         * @return builder
         * 
         */
        public Builder instanceArn(String instanceArn) {
            return instanceArn(Output.of(instanceArn));
        }

        /**
         * @param permissionSetArn The Amazon Resource Name (ARN) of the Permission Set.
         * 
         * @return builder
         * 
         */
        public Builder permissionSetArn(Output<String> permissionSetArn) {
            $.permissionSetArn = permissionSetArn;
            return this;
        }

        /**
         * @param permissionSetArn The Amazon Resource Name (ARN) of the Permission Set.
         * 
         * @return builder
         * 
         */
        public Builder permissionSetArn(String permissionSetArn) {
            return permissionSetArn(Output.of(permissionSetArn));
        }

        /**
         * @param permissionsBoundary The permissions boundary policy. See below.
         * 
         * @return builder
         * 
         */
        public Builder permissionsBoundary(Output<PermissionsBoundaryAttachmentPermissionsBoundaryArgs> permissionsBoundary) {
            $.permissionsBoundary = permissionsBoundary;
            return this;
        }

        /**
         * @param permissionsBoundary The permissions boundary policy. See below.
         * 
         * @return builder
         * 
         */
        public Builder permissionsBoundary(PermissionsBoundaryAttachmentPermissionsBoundaryArgs permissionsBoundary) {
            return permissionsBoundary(Output.of(permissionsBoundary));
        }

        public PermissionsBoundaryAttachmentArgs build() {
            $.instanceArn = Objects.requireNonNull($.instanceArn, "expected parameter 'instanceArn' to be non-null");
            $.permissionSetArn = Objects.requireNonNull($.permissionSetArn, "expected parameter 'permissionSetArn' to be non-null");
            $.permissionsBoundary = Objects.requireNonNull($.permissionsBoundary, "expected parameter 'permissionsBoundary' to be non-null");
            return $;
        }
    }

}
