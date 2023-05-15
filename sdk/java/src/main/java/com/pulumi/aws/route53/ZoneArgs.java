// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.route53;

import com.pulumi.aws.route53.inputs.ZoneVpcArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ZoneArgs extends com.pulumi.resources.ResourceArgs {

    public static final ZoneArgs Empty = new ZoneArgs();

    /**
     * A comment for the hosted zone. Defaults to &#39;Managed by Pulumi&#39;.
     * 
     */
    @Import(name="comment")
    private @Nullable Output<String> comment;

    /**
     * @return A comment for the hosted zone. Defaults to &#39;Managed by Pulumi&#39;.
     * 
     */
    public Optional<Output<String>> comment() {
        return Optional.ofNullable(this.comment);
    }

    /**
     * The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
     * 
     */
    @Import(name="delegationSetId")
    private @Nullable Output<String> delegationSetId;

    /**
     * @return The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
     * 
     */
    public Optional<Output<String>> delegationSetId() {
        return Optional.ofNullable(this.delegationSetId);
    }

    /**
     * Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
     * 
     */
    @Import(name="forceDestroy")
    private @Nullable Output<Boolean> forceDestroy;

    /**
     * @return Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
     * 
     */
    public Optional<Output<Boolean>> forceDestroy() {
        return Optional.ofNullable(this.forceDestroy);
    }

    /**
     * This is the name of the hosted zone.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return This is the name of the hosted zone.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A mapping of tags to assign to the zone. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A mapping of tags to assign to the zone. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegation_set_id` argument in this resource and any `aws.route53.ZoneAssociation` resource specifying the same zone ID. Detailed below.
     * 
     */
    @Import(name="vpcs")
    private @Nullable Output<List<ZoneVpcArgs>> vpcs;

    /**
     * @return Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegation_set_id` argument in this resource and any `aws.route53.ZoneAssociation` resource specifying the same zone ID. Detailed below.
     * 
     */
    public Optional<Output<List<ZoneVpcArgs>>> vpcs() {
        return Optional.ofNullable(this.vpcs);
    }

    private ZoneArgs() {}

    private ZoneArgs(ZoneArgs $) {
        this.comment = $.comment;
        this.delegationSetId = $.delegationSetId;
        this.forceDestroy = $.forceDestroy;
        this.name = $.name;
        this.tags = $.tags;
        this.vpcs = $.vpcs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ZoneArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ZoneArgs $;

        public Builder() {
            $ = new ZoneArgs();
        }

        public Builder(ZoneArgs defaults) {
            $ = new ZoneArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param comment A comment for the hosted zone. Defaults to &#39;Managed by Pulumi&#39;.
         * 
         * @return builder
         * 
         */
        public Builder comment(@Nullable Output<String> comment) {
            $.comment = comment;
            return this;
        }

        /**
         * @param comment A comment for the hosted zone. Defaults to &#39;Managed by Pulumi&#39;.
         * 
         * @return builder
         * 
         */
        public Builder comment(String comment) {
            return comment(Output.of(comment));
        }

        /**
         * @param delegationSetId The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
         * 
         * @return builder
         * 
         */
        public Builder delegationSetId(@Nullable Output<String> delegationSetId) {
            $.delegationSetId = delegationSetId;
            return this;
        }

        /**
         * @param delegationSetId The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
         * 
         * @return builder
         * 
         */
        public Builder delegationSetId(String delegationSetId) {
            return delegationSetId(Output.of(delegationSetId));
        }

        /**
         * @param forceDestroy Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
         * 
         * @return builder
         * 
         */
        public Builder forceDestroy(@Nullable Output<Boolean> forceDestroy) {
            $.forceDestroy = forceDestroy;
            return this;
        }

        /**
         * @param forceDestroy Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
         * 
         * @return builder
         * 
         */
        public Builder forceDestroy(Boolean forceDestroy) {
            return forceDestroy(Output.of(forceDestroy));
        }

        /**
         * @param name This is the name of the hosted zone.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name This is the name of the hosted zone.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param tags A mapping of tags to assign to the zone. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the zone. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param vpcs Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegation_set_id` argument in this resource and any `aws.route53.ZoneAssociation` resource specifying the same zone ID. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder vpcs(@Nullable Output<List<ZoneVpcArgs>> vpcs) {
            $.vpcs = vpcs;
            return this;
        }

        /**
         * @param vpcs Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegation_set_id` argument in this resource and any `aws.route53.ZoneAssociation` resource specifying the same zone ID. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder vpcs(List<ZoneVpcArgs> vpcs) {
            return vpcs(Output.of(vpcs));
        }

        /**
         * @param vpcs Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegation_set_id` argument in this resource and any `aws.route53.ZoneAssociation` resource specifying the same zone ID. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder vpcs(ZoneVpcArgs... vpcs) {
            return vpcs(List.of(vpcs));
        }

        public ZoneArgs build() {
            $.comment = Codegen.stringProp("comment").output().arg($.comment).def("Managed by Pulumi").getNullable();
            return $;
        }
    }

}
