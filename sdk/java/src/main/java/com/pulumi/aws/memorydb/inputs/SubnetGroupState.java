// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.memorydb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SubnetGroupState extends com.pulumi.resources.ResourceArgs {

    public static final SubnetGroupState Empty = new SubnetGroupState();

    /**
     * The ARN of the subnet group.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The ARN of the subnet group.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Description for the subnet group.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description for the subnet group.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Name of the subnet group. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the subnet group. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     * 
     */
    @Import(name="namePrefix")
    private @Nullable Output<String> namePrefix;

    /**
     * @return Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     * 
     */
    public Optional<Output<String>> namePrefix() {
        return Optional.ofNullable(this.namePrefix);
    }

    /**
     * Set of VPC Subnet ID-s for the subnet group. At least one subnet must be provided.
     * 
     */
    @Import(name="subnetIds")
    private @Nullable Output<List<String>> subnetIds;

    /**
     * @return Set of VPC Subnet ID-s for the subnet group. At least one subnet must be provided.
     * 
     */
    public Optional<Output<List<String>>> subnetIds() {
        return Optional.ofNullable(this.subnetIds);
    }

    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * The VPC in which the subnet group exists.
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return The VPC in which the subnet group exists.
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    private SubnetGroupState() {}

    private SubnetGroupState(SubnetGroupState $) {
        this.arn = $.arn;
        this.description = $.description;
        this.name = $.name;
        this.namePrefix = $.namePrefix;
        this.subnetIds = $.subnetIds;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.vpcId = $.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SubnetGroupState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SubnetGroupState $;

        public Builder() {
            $ = new SubnetGroupState();
        }

        public Builder(SubnetGroupState defaults) {
            $ = new SubnetGroupState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The ARN of the subnet group.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The ARN of the subnet group.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param description Description for the subnet group.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description for the subnet group.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name Name of the subnet group. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the subnet group. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namePrefix Creates a unique name beginning with the specified prefix. Conflicts with `name`.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(@Nullable Output<String> namePrefix) {
            $.namePrefix = namePrefix;
            return this;
        }

        /**
         * @param namePrefix Creates a unique name beginning with the specified prefix. Conflicts with `name`.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(String namePrefix) {
            return namePrefix(Output.of(namePrefix));
        }

        /**
         * @param subnetIds Set of VPC Subnet ID-s for the subnet group. At least one subnet must be provided.
         * 
         * @return builder
         * 
         */
        public Builder subnetIds(@Nullable Output<List<String>> subnetIds) {
            $.subnetIds = subnetIds;
            return this;
        }

        /**
         * @param subnetIds Set of VPC Subnet ID-s for the subnet group. At least one subnet must be provided.
         * 
         * @return builder
         * 
         */
        public Builder subnetIds(List<String> subnetIds) {
            return subnetIds(Output.of(subnetIds));
        }

        /**
         * @param subnetIds Set of VPC Subnet ID-s for the subnet group. At least one subnet must be provided.
         * 
         * @return builder
         * 
         */
        public Builder subnetIds(String... subnetIds) {
            return subnetIds(List.of(subnetIds));
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param vpcId The VPC in which the subnet group exists.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The VPC in which the subnet group exists.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        public SubnetGroupState build() {
            return $;
        }
    }

}
