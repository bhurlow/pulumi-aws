// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DedicatedHostState extends com.pulumi.resources.ResourceArgs {

    public static final DedicatedHostState Empty = new DedicatedHostState();

    /**
     * The ARN of the Dedicated Host.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The ARN of the Dedicated Host.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
     * 
     */
    @Import(name="assetId")
    private @Nullable Output<String> assetId;

    /**
     * @return The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
     * 
     */
    public Optional<Output<String>> assetId() {
        return Optional.ofNullable(this.assetId);
    }

    /**
     * Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
     * 
     */
    @Import(name="autoPlacement")
    private @Nullable Output<String> autoPlacement;

    /**
     * @return Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
     * 
     */
    public Optional<Output<String>> autoPlacement() {
        return Optional.ofNullable(this.autoPlacement);
    }

    /**
     * The Availability Zone in which to allocate the Dedicated Host.
     * 
     */
    @Import(name="availabilityZone")
    private @Nullable Output<String> availabilityZone;

    /**
     * @return The Availability Zone in which to allocate the Dedicated Host.
     * 
     */
    public Optional<Output<String>> availabilityZone() {
        return Optional.ofNullable(this.availabilityZone);
    }

    /**
     * Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
     * 
     */
    @Import(name="hostRecovery")
    private @Nullable Output<String> hostRecovery;

    /**
     * @return Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
     * 
     */
    public Optional<Output<String>> hostRecovery() {
        return Optional.ofNullable(this.hostRecovery);
    }

    /**
     * Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instance_family` or `instance_type` must be specified.
     * 
     */
    @Import(name="instanceFamily")
    private @Nullable Output<String> instanceFamily;

    /**
     * @return Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instance_family` or `instance_type` must be specified.
     * 
     */
    public Optional<Output<String>> instanceFamily() {
        return Optional.ofNullable(this.instanceFamily);
    }

    /**
     * Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instance_family` or `instance_type` must be specified.
     * 
     */
    @Import(name="instanceType")
    private @Nullable Output<String> instanceType;

    /**
     * @return Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instance_family` or `instance_type` must be specified.
     * 
     */
    public Optional<Output<String>> instanceType() {
        return Optional.ofNullable(this.instanceType);
    }

    /**
     * The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
     * 
     */
    @Import(name="outpostArn")
    private @Nullable Output<String> outpostArn;

    /**
     * @return The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
     * 
     */
    public Optional<Output<String>> outpostArn() {
        return Optional.ofNullable(this.outpostArn);
    }

    /**
     * The ID of the AWS account that owns the Dedicated Host.
     * 
     */
    @Import(name="ownerId")
    private @Nullable Output<String> ownerId;

    /**
     * @return The ID of the AWS account that owns the Dedicated Host.
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

    private DedicatedHostState() {}

    private DedicatedHostState(DedicatedHostState $) {
        this.arn = $.arn;
        this.assetId = $.assetId;
        this.autoPlacement = $.autoPlacement;
        this.availabilityZone = $.availabilityZone;
        this.hostRecovery = $.hostRecovery;
        this.instanceFamily = $.instanceFamily;
        this.instanceType = $.instanceType;
        this.outpostArn = $.outpostArn;
        this.ownerId = $.ownerId;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DedicatedHostState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DedicatedHostState $;

        public Builder() {
            $ = new DedicatedHostState();
        }

        public Builder(DedicatedHostState defaults) {
            $ = new DedicatedHostState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The ARN of the Dedicated Host.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The ARN of the Dedicated Host.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param assetId The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
         * 
         * @return builder
         * 
         */
        public Builder assetId(@Nullable Output<String> assetId) {
            $.assetId = assetId;
            return this;
        }

        /**
         * @param assetId The ID of the Outpost hardware asset on which to allocate the Dedicated Hosts. This parameter is supported only if you specify OutpostArn. If you are allocating the Dedicated Hosts in a Region, omit this parameter.
         * 
         * @return builder
         * 
         */
        public Builder assetId(String assetId) {
            return assetId(Output.of(assetId));
        }

        /**
         * @param autoPlacement Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
         * 
         * @return builder
         * 
         */
        public Builder autoPlacement(@Nullable Output<String> autoPlacement) {
            $.autoPlacement = autoPlacement;
            return this;
        }

        /**
         * @param autoPlacement Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: `on`, `off`. Default: `on`.
         * 
         * @return builder
         * 
         */
        public Builder autoPlacement(String autoPlacement) {
            return autoPlacement(Output.of(autoPlacement));
        }

        /**
         * @param availabilityZone The Availability Zone in which to allocate the Dedicated Host.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(@Nullable Output<String> availabilityZone) {
            $.availabilityZone = availabilityZone;
            return this;
        }

        /**
         * @param availabilityZone The Availability Zone in which to allocate the Dedicated Host.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(String availabilityZone) {
            return availabilityZone(Output.of(availabilityZone));
        }

        /**
         * @param hostRecovery Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
         * 
         * @return builder
         * 
         */
        public Builder hostRecovery(@Nullable Output<String> hostRecovery) {
            $.hostRecovery = hostRecovery;
            return this;
        }

        /**
         * @param hostRecovery Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: `on`, `off`. Default: `off`.
         * 
         * @return builder
         * 
         */
        public Builder hostRecovery(String hostRecovery) {
            return hostRecovery(Output.of(hostRecovery));
        }

        /**
         * @param instanceFamily Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instance_family` or `instance_type` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder instanceFamily(@Nullable Output<String> instanceFamily) {
            $.instanceFamily = instanceFamily;
            return this;
        }

        /**
         * @param instanceFamily Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of `instance_family` or `instance_type` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder instanceFamily(String instanceFamily) {
            return instanceFamily(Output.of(instanceFamily));
        }

        /**
         * @param instanceType Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instance_family` or `instance_type` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(@Nullable Output<String> instanceType) {
            $.instanceType = instanceType;
            return this;
        }

        /**
         * @param instanceType Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. Exactly one of `instance_family` or `instance_type` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(String instanceType) {
            return instanceType(Output.of(instanceType));
        }

        /**
         * @param outpostArn The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
         * 
         * @return builder
         * 
         */
        public Builder outpostArn(@Nullable Output<String> outpostArn) {
            $.outpostArn = outpostArn;
            return this;
        }

        /**
         * @param outpostArn The Amazon Resource Name (ARN) of the AWS Outpost on which to allocate the Dedicated Host.
         * 
         * @return builder
         * 
         */
        public Builder outpostArn(String outpostArn) {
            return outpostArn(Output.of(outpostArn));
        }

        /**
         * @param ownerId The ID of the AWS account that owns the Dedicated Host.
         * 
         * @return builder
         * 
         */
        public Builder ownerId(@Nullable Output<String> ownerId) {
            $.ownerId = ownerId;
            return this;
        }

        /**
         * @param ownerId The ID of the AWS account that owns the Dedicated Host.
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

        public DedicatedHostState build() {
            return $;
        }
    }

}
