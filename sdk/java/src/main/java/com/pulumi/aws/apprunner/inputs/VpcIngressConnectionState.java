// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.apprunner.inputs;

import com.pulumi.aws.apprunner.inputs.VpcIngressConnectionIngressVpcConfigurationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VpcIngressConnectionState extends com.pulumi.resources.ResourceArgs {

    public static final VpcIngressConnectionState Empty = new VpcIngressConnectionState();

    /**
     * The Amazon Resource Name (ARN) of the VPC Ingress Connection.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) of the VPC Ingress Connection.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The domain name associated with the VPC Ingress Connection resource.
     * 
     */
    @Import(name="domainName")
    private @Nullable Output<String> domainName;

    /**
     * @return The domain name associated with the VPC Ingress Connection resource.
     * 
     */
    public Optional<Output<String>> domainName() {
        return Optional.ofNullable(this.domainName);
    }

    /**
     * Specifications for the customer’s Amazon VPC and the related AWS PrivateLink VPC endpoint that are used to create the VPC Ingress Connection resource. See Ingress VPC Configuration below for more details.
     * 
     */
    @Import(name="ingressVpcConfiguration")
    private @Nullable Output<VpcIngressConnectionIngressVpcConfigurationArgs> ingressVpcConfiguration;

    /**
     * @return Specifications for the customer’s Amazon VPC and the related AWS PrivateLink VPC endpoint that are used to create the VPC Ingress Connection resource. See Ingress VPC Configuration below for more details.
     * 
     */
    public Optional<Output<VpcIngressConnectionIngressVpcConfigurationArgs>> ingressVpcConfiguration() {
        return Optional.ofNullable(this.ingressVpcConfiguration);
    }

    /**
     * A name for the VPC Ingress Connection resource. It must be unique across all the active VPC Ingress Connections in your AWS account in the AWS Region.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A name for the VPC Ingress Connection resource. It must be unique across all the active VPC Ingress Connections in your AWS account in the AWS Region.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The Amazon Resource Name (ARN) for this App Runner service that is used to create the VPC Ingress Connection resource.
     * 
     */
    @Import(name="serviceArn")
    private @Nullable Output<String> serviceArn;

    /**
     * @return The Amazon Resource Name (ARN) for this App Runner service that is used to create the VPC Ingress Connection resource.
     * 
     */
    public Optional<Output<String>> serviceArn() {
        return Optional.ofNullable(this.serviceArn);
    }

    /**
     * The current status of the VPC Ingress Connection.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The current status of the VPC Ingress Connection.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    private VpcIngressConnectionState() {}

    private VpcIngressConnectionState(VpcIngressConnectionState $) {
        this.arn = $.arn;
        this.domainName = $.domainName;
        this.ingressVpcConfiguration = $.ingressVpcConfiguration;
        this.name = $.name;
        this.serviceArn = $.serviceArn;
        this.status = $.status;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VpcIngressConnectionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VpcIngressConnectionState $;

        public Builder() {
            $ = new VpcIngressConnectionState();
        }

        public Builder(VpcIngressConnectionState defaults) {
            $ = new VpcIngressConnectionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The Amazon Resource Name (ARN) of the VPC Ingress Connection.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The Amazon Resource Name (ARN) of the VPC Ingress Connection.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param domainName The domain name associated with the VPC Ingress Connection resource.
         * 
         * @return builder
         * 
         */
        public Builder domainName(@Nullable Output<String> domainName) {
            $.domainName = domainName;
            return this;
        }

        /**
         * @param domainName The domain name associated with the VPC Ingress Connection resource.
         * 
         * @return builder
         * 
         */
        public Builder domainName(String domainName) {
            return domainName(Output.of(domainName));
        }

        /**
         * @param ingressVpcConfiguration Specifications for the customer’s Amazon VPC and the related AWS PrivateLink VPC endpoint that are used to create the VPC Ingress Connection resource. See Ingress VPC Configuration below for more details.
         * 
         * @return builder
         * 
         */
        public Builder ingressVpcConfiguration(@Nullable Output<VpcIngressConnectionIngressVpcConfigurationArgs> ingressVpcConfiguration) {
            $.ingressVpcConfiguration = ingressVpcConfiguration;
            return this;
        }

        /**
         * @param ingressVpcConfiguration Specifications for the customer’s Amazon VPC and the related AWS PrivateLink VPC endpoint that are used to create the VPC Ingress Connection resource. See Ingress VPC Configuration below for more details.
         * 
         * @return builder
         * 
         */
        public Builder ingressVpcConfiguration(VpcIngressConnectionIngressVpcConfigurationArgs ingressVpcConfiguration) {
            return ingressVpcConfiguration(Output.of(ingressVpcConfiguration));
        }

        /**
         * @param name A name for the VPC Ingress Connection resource. It must be unique across all the active VPC Ingress Connections in your AWS account in the AWS Region.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A name for the VPC Ingress Connection resource. It must be unique across all the active VPC Ingress Connections in your AWS account in the AWS Region.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param serviceArn The Amazon Resource Name (ARN) for this App Runner service that is used to create the VPC Ingress Connection resource.
         * 
         * @return builder
         * 
         */
        public Builder serviceArn(@Nullable Output<String> serviceArn) {
            $.serviceArn = serviceArn;
            return this;
        }

        /**
         * @param serviceArn The Amazon Resource Name (ARN) for this App Runner service that is used to create the VPC Ingress Connection resource.
         * 
         * @return builder
         * 
         */
        public Builder serviceArn(String serviceArn) {
            return serviceArn(Output.of(serviceArn));
        }

        /**
         * @param status The current status of the VPC Ingress Connection.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The current status of the VPC Ingress Connection.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
         */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        public VpcIngressConnectionState build() {
            return $;
        }
    }

}
