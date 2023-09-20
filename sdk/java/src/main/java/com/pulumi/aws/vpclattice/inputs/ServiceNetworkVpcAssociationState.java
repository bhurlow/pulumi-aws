// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.vpclattice.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceNetworkVpcAssociationState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceNetworkVpcAssociationState Empty = new ServiceNetworkVpcAssociationState();

    /**
     * The ARN of the Association.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The ARN of the Association.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The account that created the association.
     * 
     */
    @Import(name="createdBy")
    private @Nullable Output<String> createdBy;

    /**
     * @return The account that created the association.
     * 
     */
    public Optional<Output<String>> createdBy() {
        return Optional.ofNullable(this.createdBy);
    }

    /**
     * The IDs of the security groups.
     * 
     */
    @Import(name="securityGroupIds")
    private @Nullable Output<List<String>> securityGroupIds;

    /**
     * @return The IDs of the security groups.
     * 
     */
    public Optional<Output<List<String>>> securityGroupIds() {
        return Optional.ofNullable(this.securityGroupIds);
    }

    /**
     * The ID or Amazon Resource Identifier (ARN) of the service network. You must use the ARN if the resources specified in the operation are in different accounts.
     * The following arguments are optional:
     * 
     */
    @Import(name="serviceNetworkIdentifier")
    private @Nullable Output<String> serviceNetworkIdentifier;

    /**
     * @return The ID or Amazon Resource Identifier (ARN) of the service network. You must use the ARN if the resources specified in the operation are in different accounts.
     * The following arguments are optional:
     * 
     */
    public Optional<Output<String>> serviceNetworkIdentifier() {
        return Optional.ofNullable(this.serviceNetworkIdentifier);
    }

    /**
     * The operations status. Valid Values are CREATE_IN_PROGRESS | ACTIVE | DELETE_IN_PROGRESS | CREATE_FAILED | DELETE_FAILED
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The operations status. Valid Values are CREATE_IN_PROGRESS | ACTIVE | DELETE_IN_PROGRESS | CREATE_FAILED | DELETE_FAILED
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * The ID of the VPC.
     * 
     */
    @Import(name="vpcIdentifier")
    private @Nullable Output<String> vpcIdentifier;

    /**
     * @return The ID of the VPC.
     * 
     */
    public Optional<Output<String>> vpcIdentifier() {
        return Optional.ofNullable(this.vpcIdentifier);
    }

    private ServiceNetworkVpcAssociationState() {}

    private ServiceNetworkVpcAssociationState(ServiceNetworkVpcAssociationState $) {
        this.arn = $.arn;
        this.createdBy = $.createdBy;
        this.securityGroupIds = $.securityGroupIds;
        this.serviceNetworkIdentifier = $.serviceNetworkIdentifier;
        this.status = $.status;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.vpcIdentifier = $.vpcIdentifier;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceNetworkVpcAssociationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceNetworkVpcAssociationState $;

        public Builder() {
            $ = new ServiceNetworkVpcAssociationState();
        }

        public Builder(ServiceNetworkVpcAssociationState defaults) {
            $ = new ServiceNetworkVpcAssociationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The ARN of the Association.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The ARN of the Association.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param createdBy The account that created the association.
         * 
         * @return builder
         * 
         */
        public Builder createdBy(@Nullable Output<String> createdBy) {
            $.createdBy = createdBy;
            return this;
        }

        /**
         * @param createdBy The account that created the association.
         * 
         * @return builder
         * 
         */
        public Builder createdBy(String createdBy) {
            return createdBy(Output.of(createdBy));
        }

        /**
         * @param securityGroupIds The IDs of the security groups.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(@Nullable Output<List<String>> securityGroupIds) {
            $.securityGroupIds = securityGroupIds;
            return this;
        }

        /**
         * @param securityGroupIds The IDs of the security groups.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(List<String> securityGroupIds) {
            return securityGroupIds(Output.of(securityGroupIds));
        }

        /**
         * @param securityGroupIds The IDs of the security groups.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(String... securityGroupIds) {
            return securityGroupIds(List.of(securityGroupIds));
        }

        /**
         * @param serviceNetworkIdentifier The ID or Amazon Resource Identifier (ARN) of the service network. You must use the ARN if the resources specified in the operation are in different accounts.
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder serviceNetworkIdentifier(@Nullable Output<String> serviceNetworkIdentifier) {
            $.serviceNetworkIdentifier = serviceNetworkIdentifier;
            return this;
        }

        /**
         * @param serviceNetworkIdentifier The ID or Amazon Resource Identifier (ARN) of the service network. You must use the ARN if the resources specified in the operation are in different accounts.
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder serviceNetworkIdentifier(String serviceNetworkIdentifier) {
            return serviceNetworkIdentifier(Output.of(serviceNetworkIdentifier));
        }

        /**
         * @param status The operations status. Valid Values are CREATE_IN_PROGRESS | ACTIVE | DELETE_IN_PROGRESS | CREATE_FAILED | DELETE_FAILED
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The operations status. Valid Values are CREATE_IN_PROGRESS | ACTIVE | DELETE_IN_PROGRESS | CREATE_FAILED | DELETE_FAILED
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
         * @param vpcIdentifier The ID of the VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcIdentifier(@Nullable Output<String> vpcIdentifier) {
            $.vpcIdentifier = vpcIdentifier;
            return this;
        }

        /**
         * @param vpcIdentifier The ID of the VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcIdentifier(String vpcIdentifier) {
            return vpcIdentifier(Output.of(vpcIdentifier));
        }

        public ServiceNetworkVpcAssociationState build() {
            return $;
        }
    }

}
