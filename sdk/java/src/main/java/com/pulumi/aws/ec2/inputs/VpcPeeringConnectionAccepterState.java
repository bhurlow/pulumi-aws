// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2.inputs;

import com.pulumi.aws.ec2.inputs.VpcPeeringConnectionAccepterAccepterArgs;
import com.pulumi.aws.ec2.inputs.VpcPeeringConnectionAccepterRequesterArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VpcPeeringConnectionAccepterState extends com.pulumi.resources.ResourceArgs {

    public static final VpcPeeringConnectionAccepterState Empty = new VpcPeeringConnectionAccepterState();

    /**
     * The status of the VPC Peering Connection request.
     * 
     */
    @Import(name="acceptStatus")
    private @Nullable Output<String> acceptStatus;

    /**
     * @return The status of the VPC Peering Connection request.
     * 
     */
    public Optional<Output<String>> acceptStatus() {
        return Optional.ofNullable(this.acceptStatus);
    }

    /**
     * A configuration block that describes [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
     * 
     */
    @Import(name="accepter")
    private @Nullable Output<VpcPeeringConnectionAccepterAccepterArgs> accepter;

    /**
     * @return A configuration block that describes [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
     * 
     */
    public Optional<Output<VpcPeeringConnectionAccepterAccepterArgs>> accepter() {
        return Optional.ofNullable(this.accepter);
    }

    /**
     * Whether or not to accept the peering request. Defaults to `false`.
     * 
     */
    @Import(name="autoAccept")
    private @Nullable Output<Boolean> autoAccept;

    /**
     * @return Whether or not to accept the peering request. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> autoAccept() {
        return Optional.ofNullable(this.autoAccept);
    }

    /**
     * The AWS account ID of the owner of the requester VPC.
     * 
     */
    @Import(name="peerOwnerId")
    private @Nullable Output<String> peerOwnerId;

    /**
     * @return The AWS account ID of the owner of the requester VPC.
     * 
     */
    public Optional<Output<String>> peerOwnerId() {
        return Optional.ofNullable(this.peerOwnerId);
    }

    /**
     * The region of the accepter VPC.
     * 
     */
    @Import(name="peerRegion")
    private @Nullable Output<String> peerRegion;

    /**
     * @return The region of the accepter VPC.
     * 
     */
    public Optional<Output<String>> peerRegion() {
        return Optional.ofNullable(this.peerRegion);
    }

    /**
     * The ID of the requester VPC.
     * 
     */
    @Import(name="peerVpcId")
    private @Nullable Output<String> peerVpcId;

    /**
     * @return The ID of the requester VPC.
     * 
     */
    public Optional<Output<String>> peerVpcId() {
        return Optional.ofNullable(this.peerVpcId);
    }

    /**
     * A configuration block that describes [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
     * 
     */
    @Import(name="requester")
    private @Nullable Output<VpcPeeringConnectionAccepterRequesterArgs> requester;

    /**
     * @return A configuration block that describes [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
     * 
     */
    public Optional<Output<VpcPeeringConnectionAccepterRequesterArgs>> requester() {
        return Optional.ofNullable(this.requester);
    }

    /**
     * A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     * 
     */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider .
     * 
     */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * The ID of the accepter VPC.
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return The ID of the accepter VPC.
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    /**
     * The VPC Peering Connection ID to manage.
     * 
     */
    @Import(name="vpcPeeringConnectionId")
    private @Nullable Output<String> vpcPeeringConnectionId;

    /**
     * @return The VPC Peering Connection ID to manage.
     * 
     */
    public Optional<Output<String>> vpcPeeringConnectionId() {
        return Optional.ofNullable(this.vpcPeeringConnectionId);
    }

    private VpcPeeringConnectionAccepterState() {}

    private VpcPeeringConnectionAccepterState(VpcPeeringConnectionAccepterState $) {
        this.acceptStatus = $.acceptStatus;
        this.accepter = $.accepter;
        this.autoAccept = $.autoAccept;
        this.peerOwnerId = $.peerOwnerId;
        this.peerRegion = $.peerRegion;
        this.peerVpcId = $.peerVpcId;
        this.requester = $.requester;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.vpcId = $.vpcId;
        this.vpcPeeringConnectionId = $.vpcPeeringConnectionId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VpcPeeringConnectionAccepterState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VpcPeeringConnectionAccepterState $;

        public Builder() {
            $ = new VpcPeeringConnectionAccepterState();
        }

        public Builder(VpcPeeringConnectionAccepterState defaults) {
            $ = new VpcPeeringConnectionAccepterState(Objects.requireNonNull(defaults));
        }

        /**
         * @param acceptStatus The status of the VPC Peering Connection request.
         * 
         * @return builder
         * 
         */
        public Builder acceptStatus(@Nullable Output<String> acceptStatus) {
            $.acceptStatus = acceptStatus;
            return this;
        }

        /**
         * @param acceptStatus The status of the VPC Peering Connection request.
         * 
         * @return builder
         * 
         */
        public Builder acceptStatus(String acceptStatus) {
            return acceptStatus(Output.of(acceptStatus));
        }

        /**
         * @param accepter A configuration block that describes [VPC Peering Connection]
         * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
         * 
         * @return builder
         * 
         */
        public Builder accepter(@Nullable Output<VpcPeeringConnectionAccepterAccepterArgs> accepter) {
            $.accepter = accepter;
            return this;
        }

        /**
         * @param accepter A configuration block that describes [VPC Peering Connection]
         * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
         * 
         * @return builder
         * 
         */
        public Builder accepter(VpcPeeringConnectionAccepterAccepterArgs accepter) {
            return accepter(Output.of(accepter));
        }

        /**
         * @param autoAccept Whether or not to accept the peering request. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder autoAccept(@Nullable Output<Boolean> autoAccept) {
            $.autoAccept = autoAccept;
            return this;
        }

        /**
         * @param autoAccept Whether or not to accept the peering request. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder autoAccept(Boolean autoAccept) {
            return autoAccept(Output.of(autoAccept));
        }

        /**
         * @param peerOwnerId The AWS account ID of the owner of the requester VPC.
         * 
         * @return builder
         * 
         */
        public Builder peerOwnerId(@Nullable Output<String> peerOwnerId) {
            $.peerOwnerId = peerOwnerId;
            return this;
        }

        /**
         * @param peerOwnerId The AWS account ID of the owner of the requester VPC.
         * 
         * @return builder
         * 
         */
        public Builder peerOwnerId(String peerOwnerId) {
            return peerOwnerId(Output.of(peerOwnerId));
        }

        /**
         * @param peerRegion The region of the accepter VPC.
         * 
         * @return builder
         * 
         */
        public Builder peerRegion(@Nullable Output<String> peerRegion) {
            $.peerRegion = peerRegion;
            return this;
        }

        /**
         * @param peerRegion The region of the accepter VPC.
         * 
         * @return builder
         * 
         */
        public Builder peerRegion(String peerRegion) {
            return peerRegion(Output.of(peerRegion));
        }

        /**
         * @param peerVpcId The ID of the requester VPC.
         * 
         * @return builder
         * 
         */
        public Builder peerVpcId(@Nullable Output<String> peerVpcId) {
            $.peerVpcId = peerVpcId;
            return this;
        }

        /**
         * @param peerVpcId The ID of the requester VPC.
         * 
         * @return builder
         * 
         */
        public Builder peerVpcId(String peerVpcId) {
            return peerVpcId(Output.of(peerVpcId));
        }

        /**
         * @param requester A configuration block that describes [VPC Peering Connection]
         * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
         * 
         * @return builder
         * 
         */
        public Builder requester(@Nullable Output<VpcPeeringConnectionAccepterRequesterArgs> requester) {
            $.requester = requester;
            return this;
        }

        /**
         * @param requester A configuration block that describes [VPC Peering Connection]
         * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
         * 
         * @return builder
         * 
         */
        public Builder requester(VpcPeeringConnectionAccepterRequesterArgs requester) {
            return requester(Output.of(requester));
        }

        /**
         * @param tags A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider .
         * 
         * @return builder
         * 
         */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider .
         * 
         * @return builder
         * 
         */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param vpcId The ID of the accepter VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The ID of the accepter VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        /**
         * @param vpcPeeringConnectionId The VPC Peering Connection ID to manage.
         * 
         * @return builder
         * 
         */
        public Builder vpcPeeringConnectionId(@Nullable Output<String> vpcPeeringConnectionId) {
            $.vpcPeeringConnectionId = vpcPeeringConnectionId;
            return this;
        }

        /**
         * @param vpcPeeringConnectionId The VPC Peering Connection ID to manage.
         * 
         * @return builder
         * 
         */
        public Builder vpcPeeringConnectionId(String vpcPeeringConnectionId) {
            return vpcPeeringConnectionId(Output.of(vpcPeeringConnectionId));
        }

        public VpcPeeringConnectionAccepterState build() {
            return $;
        }
    }

}