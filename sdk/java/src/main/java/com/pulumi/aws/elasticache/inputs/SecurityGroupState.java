// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.elasticache.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecurityGroupState extends com.pulumi.resources.ResourceArgs {

    public static final SecurityGroupState Empty = new SecurityGroupState();

    /**
     * description for the cache security group. Defaults to &#34;Managed by Pulumi&#34;.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return description for the cache security group. Defaults to &#34;Managed by Pulumi&#34;.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Name for the cache security group. This value is stored as a lowercase string.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name for the cache security group. This value is stored as a lowercase string.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * List of EC2 security group names to be
     * authorized for ingress to the cache security group
     * 
     */
    @Import(name="securityGroupNames")
    private @Nullable Output<List<String>> securityGroupNames;

    /**
     * @return List of EC2 security group names to be
     * authorized for ingress to the cache security group
     * 
     */
    public Optional<Output<List<String>>> securityGroupNames() {
        return Optional.ofNullable(this.securityGroupNames);
    }

    private SecurityGroupState() {}

    private SecurityGroupState(SecurityGroupState $) {
        this.description = $.description;
        this.name = $.name;
        this.securityGroupNames = $.securityGroupNames;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecurityGroupState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecurityGroupState $;

        public Builder() {
            $ = new SecurityGroupState();
        }

        public Builder(SecurityGroupState defaults) {
            $ = new SecurityGroupState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description description for the cache security group. Defaults to &#34;Managed by Pulumi&#34;.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description description for the cache security group. Defaults to &#34;Managed by Pulumi&#34;.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name Name for the cache security group. This value is stored as a lowercase string.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name for the cache security group. This value is stored as a lowercase string.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param securityGroupNames List of EC2 security group names to be
         * authorized for ingress to the cache security group
         * 
         * @return builder
         * 
         */
        public Builder securityGroupNames(@Nullable Output<List<String>> securityGroupNames) {
            $.securityGroupNames = securityGroupNames;
            return this;
        }

        /**
         * @param securityGroupNames List of EC2 security group names to be
         * authorized for ingress to the cache security group
         * 
         * @return builder
         * 
         */
        public Builder securityGroupNames(List<String> securityGroupNames) {
            return securityGroupNames(Output.of(securityGroupNames));
        }

        /**
         * @param securityGroupNames List of EC2 security group names to be
         * authorized for ingress to the cache security group
         * 
         * @return builder
         * 
         */
        public Builder securityGroupNames(String... securityGroupNames) {
            return securityGroupNames(List.of(securityGroupNames));
        }

        public SecurityGroupState build() {
            $.description = Codegen.stringProp("description").output().arg($.description).def("Managed by Pulumi").getNullable();
            return $;
        }
    }

}