// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2.inputs;

import com.pulumi.aws.ec2.inputs.ManagedPrefixListEntryArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ManagedPrefixListState extends com.pulumi.resources.ResourceArgs {

    public static final ManagedPrefixListState Empty = new ManagedPrefixListState();

    /**
     * Address family (`IPv4` or `IPv6`) of this prefix list.
     * 
     */
    @Import(name="addressFamily")
    private @Nullable Output<String> addressFamily;

    /**
     * @return Address family (`IPv4` or `IPv6`) of this prefix list.
     * 
     */
    public Optional<Output<String>> addressFamily() {
        return Optional.ofNullable(this.addressFamily);
    }

    /**
     * ARN of the prefix list.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return ARN of the prefix list.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Configuration block for prefix list entry. Detailed below. Different entries may have overlapping CIDR blocks, but a particular CIDR should not be duplicated.
     * 
     */
    @Import(name="entries")
    private @Nullable Output<List<ManagedPrefixListEntryArgs>> entries;

    /**
     * @return Configuration block for prefix list entry. Detailed below. Different entries may have overlapping CIDR blocks, but a particular CIDR should not be duplicated.
     * 
     */
    public Optional<Output<List<ManagedPrefixListEntryArgs>>> entries() {
        return Optional.ofNullable(this.entries);
    }

    /**
     * Maximum number of entries that this prefix list can contain.
     * 
     */
    @Import(name="maxEntries")
    private @Nullable Output<Integer> maxEntries;

    /**
     * @return Maximum number of entries that this prefix list can contain.
     * 
     */
    public Optional<Output<Integer>> maxEntries() {
        return Optional.ofNullable(this.maxEntries);
    }

    /**
     * Name of this resource. The name must not start with `com.amazonaws`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of this resource. The name must not start with `com.amazonaws`.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * ID of the AWS account that owns this prefix list.
     * 
     */
    @Import(name="ownerId")
    private @Nullable Output<String> ownerId;

    /**
     * @return ID of the AWS account that owns this prefix list.
     * 
     */
    public Optional<Output<String>> ownerId() {
        return Optional.ofNullable(this.ownerId);
    }

    /**
     * Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * Latest version of this prefix list.
     * 
     */
    @Import(name="version")
    private @Nullable Output<Integer> version;

    /**
     * @return Latest version of this prefix list.
     * 
     */
    public Optional<Output<Integer>> version() {
        return Optional.ofNullable(this.version);
    }

    private ManagedPrefixListState() {}

    private ManagedPrefixListState(ManagedPrefixListState $) {
        this.addressFamily = $.addressFamily;
        this.arn = $.arn;
        this.entries = $.entries;
        this.maxEntries = $.maxEntries;
        this.name = $.name;
        this.ownerId = $.ownerId;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ManagedPrefixListState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ManagedPrefixListState $;

        public Builder() {
            $ = new ManagedPrefixListState();
        }

        public Builder(ManagedPrefixListState defaults) {
            $ = new ManagedPrefixListState(Objects.requireNonNull(defaults));
        }

        /**
         * @param addressFamily Address family (`IPv4` or `IPv6`) of this prefix list.
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(@Nullable Output<String> addressFamily) {
            $.addressFamily = addressFamily;
            return this;
        }

        /**
         * @param addressFamily Address family (`IPv4` or `IPv6`) of this prefix list.
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(String addressFamily) {
            return addressFamily(Output.of(addressFamily));
        }

        /**
         * @param arn ARN of the prefix list.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn ARN of the prefix list.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param entries Configuration block for prefix list entry. Detailed below. Different entries may have overlapping CIDR blocks, but a particular CIDR should not be duplicated.
         * 
         * @return builder
         * 
         */
        public Builder entries(@Nullable Output<List<ManagedPrefixListEntryArgs>> entries) {
            $.entries = entries;
            return this;
        }

        /**
         * @param entries Configuration block for prefix list entry. Detailed below. Different entries may have overlapping CIDR blocks, but a particular CIDR should not be duplicated.
         * 
         * @return builder
         * 
         */
        public Builder entries(List<ManagedPrefixListEntryArgs> entries) {
            return entries(Output.of(entries));
        }

        /**
         * @param entries Configuration block for prefix list entry. Detailed below. Different entries may have overlapping CIDR blocks, but a particular CIDR should not be duplicated.
         * 
         * @return builder
         * 
         */
        public Builder entries(ManagedPrefixListEntryArgs... entries) {
            return entries(List.of(entries));
        }

        /**
         * @param maxEntries Maximum number of entries that this prefix list can contain.
         * 
         * @return builder
         * 
         */
        public Builder maxEntries(@Nullable Output<Integer> maxEntries) {
            $.maxEntries = maxEntries;
            return this;
        }

        /**
         * @param maxEntries Maximum number of entries that this prefix list can contain.
         * 
         * @return builder
         * 
         */
        public Builder maxEntries(Integer maxEntries) {
            return maxEntries(Output.of(maxEntries));
        }

        /**
         * @param name Name of this resource. The name must not start with `com.amazonaws`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of this resource. The name must not start with `com.amazonaws`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param ownerId ID of the AWS account that owns this prefix list.
         * 
         * @return builder
         * 
         */
        public Builder ownerId(@Nullable Output<String> ownerId) {
            $.ownerId = ownerId;
            return this;
        }

        /**
         * @param ownerId ID of the AWS account that owns this prefix list.
         * 
         * @return builder
         * 
         */
        public Builder ownerId(String ownerId) {
            return ownerId(Output.of(ownerId));
        }

        /**
         * @param tags Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param version Latest version of this prefix list.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<Integer> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version Latest version of this prefix list.
         * 
         * @return builder
         * 
         */
        public Builder version(Integer version) {
            return version(Output.of(version));
        }

        public ManagedPrefixListState build() {
            return $;
        }
    }

}
