// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appstream;

import com.pulumi.aws.appstream.inputs.ImageBuilderAccessEndpointArgs;
import com.pulumi.aws.appstream.inputs.ImageBuilderDomainJoinInfoArgs;
import com.pulumi.aws.appstream.inputs.ImageBuilderVpcConfigArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ImageBuilderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ImageBuilderArgs Empty = new ImageBuilderArgs();

    /**
     * Set of interface VPC endpoint (interface endpoint) objects. Maximum of 4. See below.
     * 
     */
    @Import(name="accessEndpoints")
    private @Nullable Output<List<ImageBuilderAccessEndpointArgs>> accessEndpoints;

    /**
     * @return Set of interface VPC endpoint (interface endpoint) objects. Maximum of 4. See below.
     * 
     */
    public Optional<Output<List<ImageBuilderAccessEndpointArgs>>> accessEndpoints() {
        return Optional.ofNullable(this.accessEndpoints);
    }

    /**
     * Version of the AppStream 2.0 agent to use for this image builder.
     * 
     */
    @Import(name="appstreamAgentVersion")
    private @Nullable Output<String> appstreamAgentVersion;

    /**
     * @return Version of the AppStream 2.0 agent to use for this image builder.
     * 
     */
    public Optional<Output<String>> appstreamAgentVersion() {
        return Optional.ofNullable(this.appstreamAgentVersion);
    }

    /**
     * Description to display.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description to display.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Human-readable friendly name for the AppStream image builder.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return Human-readable friendly name for the AppStream image builder.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * Configuration block for the name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain. See below.
     * 
     */
    @Import(name="domainJoinInfo")
    private @Nullable Output<ImageBuilderDomainJoinInfoArgs> domainJoinInfo;

    /**
     * @return Configuration block for the name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain. See below.
     * 
     */
    public Optional<Output<ImageBuilderDomainJoinInfoArgs>> domainJoinInfo() {
        return Optional.ofNullable(this.domainJoinInfo);
    }

    /**
     * Enables or disables default internet access for the image builder.
     * 
     */
    @Import(name="enableDefaultInternetAccess")
    private @Nullable Output<Boolean> enableDefaultInternetAccess;

    /**
     * @return Enables or disables default internet access for the image builder.
     * 
     */
    public Optional<Output<Boolean>> enableDefaultInternetAccess() {
        return Optional.ofNullable(this.enableDefaultInternetAccess);
    }

    /**
     * ARN of the IAM role to apply to the image builder.
     * 
     */
    @Import(name="iamRoleArn")
    private @Nullable Output<String> iamRoleArn;

    /**
     * @return ARN of the IAM role to apply to the image builder.
     * 
     */
    public Optional<Output<String>> iamRoleArn() {
        return Optional.ofNullable(this.iamRoleArn);
    }

    /**
     * ARN of the public, private, or shared image to use.
     * 
     */
    @Import(name="imageArn")
    private @Nullable Output<String> imageArn;

    /**
     * @return ARN of the public, private, or shared image to use.
     * 
     */
    public Optional<Output<String>> imageArn() {
        return Optional.ofNullable(this.imageArn);
    }

    /**
     * Name of the image used to create the image builder.
     * 
     */
    @Import(name="imageName")
    private @Nullable Output<String> imageName;

    /**
     * @return Name of the image used to create the image builder.
     * 
     */
    public Optional<Output<String>> imageName() {
        return Optional.ofNullable(this.imageName);
    }

    /**
     * Instance type to use when launching the image builder.
     * 
     */
    @Import(name="instanceType", required=true)
    private Output<String> instanceType;

    /**
     * @return Instance type to use when launching the image builder.
     * 
     */
    public Output<String> instanceType() {
        return this.instanceType;
    }

    /**
     * Unique name for the image builder.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Unique name for the image builder.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Map of tags to assign to the instance. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Map of tags to assign to the instance. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Configuration block for the VPC configuration for the image builder. See below.
     * 
     */
    @Import(name="vpcConfig")
    private @Nullable Output<ImageBuilderVpcConfigArgs> vpcConfig;

    /**
     * @return Configuration block for the VPC configuration for the image builder. See below.
     * 
     */
    public Optional<Output<ImageBuilderVpcConfigArgs>> vpcConfig() {
        return Optional.ofNullable(this.vpcConfig);
    }

    private ImageBuilderArgs() {}

    private ImageBuilderArgs(ImageBuilderArgs $) {
        this.accessEndpoints = $.accessEndpoints;
        this.appstreamAgentVersion = $.appstreamAgentVersion;
        this.description = $.description;
        this.displayName = $.displayName;
        this.domainJoinInfo = $.domainJoinInfo;
        this.enableDefaultInternetAccess = $.enableDefaultInternetAccess;
        this.iamRoleArn = $.iamRoleArn;
        this.imageArn = $.imageArn;
        this.imageName = $.imageName;
        this.instanceType = $.instanceType;
        this.name = $.name;
        this.tags = $.tags;
        this.vpcConfig = $.vpcConfig;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ImageBuilderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ImageBuilderArgs $;

        public Builder() {
            $ = new ImageBuilderArgs();
        }

        public Builder(ImageBuilderArgs defaults) {
            $ = new ImageBuilderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessEndpoints Set of interface VPC endpoint (interface endpoint) objects. Maximum of 4. See below.
         * 
         * @return builder
         * 
         */
        public Builder accessEndpoints(@Nullable Output<List<ImageBuilderAccessEndpointArgs>> accessEndpoints) {
            $.accessEndpoints = accessEndpoints;
            return this;
        }

        /**
         * @param accessEndpoints Set of interface VPC endpoint (interface endpoint) objects. Maximum of 4. See below.
         * 
         * @return builder
         * 
         */
        public Builder accessEndpoints(List<ImageBuilderAccessEndpointArgs> accessEndpoints) {
            return accessEndpoints(Output.of(accessEndpoints));
        }

        /**
         * @param accessEndpoints Set of interface VPC endpoint (interface endpoint) objects. Maximum of 4. See below.
         * 
         * @return builder
         * 
         */
        public Builder accessEndpoints(ImageBuilderAccessEndpointArgs... accessEndpoints) {
            return accessEndpoints(List.of(accessEndpoints));
        }

        /**
         * @param appstreamAgentVersion Version of the AppStream 2.0 agent to use for this image builder.
         * 
         * @return builder
         * 
         */
        public Builder appstreamAgentVersion(@Nullable Output<String> appstreamAgentVersion) {
            $.appstreamAgentVersion = appstreamAgentVersion;
            return this;
        }

        /**
         * @param appstreamAgentVersion Version of the AppStream 2.0 agent to use for this image builder.
         * 
         * @return builder
         * 
         */
        public Builder appstreamAgentVersion(String appstreamAgentVersion) {
            return appstreamAgentVersion(Output.of(appstreamAgentVersion));
        }

        /**
         * @param description Description to display.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description to display.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param displayName Human-readable friendly name for the AppStream image builder.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Human-readable friendly name for the AppStream image builder.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param domainJoinInfo Configuration block for the name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain. See below.
         * 
         * @return builder
         * 
         */
        public Builder domainJoinInfo(@Nullable Output<ImageBuilderDomainJoinInfoArgs> domainJoinInfo) {
            $.domainJoinInfo = domainJoinInfo;
            return this;
        }

        /**
         * @param domainJoinInfo Configuration block for the name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain. See below.
         * 
         * @return builder
         * 
         */
        public Builder domainJoinInfo(ImageBuilderDomainJoinInfoArgs domainJoinInfo) {
            return domainJoinInfo(Output.of(domainJoinInfo));
        }

        /**
         * @param enableDefaultInternetAccess Enables or disables default internet access for the image builder.
         * 
         * @return builder
         * 
         */
        public Builder enableDefaultInternetAccess(@Nullable Output<Boolean> enableDefaultInternetAccess) {
            $.enableDefaultInternetAccess = enableDefaultInternetAccess;
            return this;
        }

        /**
         * @param enableDefaultInternetAccess Enables or disables default internet access for the image builder.
         * 
         * @return builder
         * 
         */
        public Builder enableDefaultInternetAccess(Boolean enableDefaultInternetAccess) {
            return enableDefaultInternetAccess(Output.of(enableDefaultInternetAccess));
        }

        /**
         * @param iamRoleArn ARN of the IAM role to apply to the image builder.
         * 
         * @return builder
         * 
         */
        public Builder iamRoleArn(@Nullable Output<String> iamRoleArn) {
            $.iamRoleArn = iamRoleArn;
            return this;
        }

        /**
         * @param iamRoleArn ARN of the IAM role to apply to the image builder.
         * 
         * @return builder
         * 
         */
        public Builder iamRoleArn(String iamRoleArn) {
            return iamRoleArn(Output.of(iamRoleArn));
        }

        /**
         * @param imageArn ARN of the public, private, or shared image to use.
         * 
         * @return builder
         * 
         */
        public Builder imageArn(@Nullable Output<String> imageArn) {
            $.imageArn = imageArn;
            return this;
        }

        /**
         * @param imageArn ARN of the public, private, or shared image to use.
         * 
         * @return builder
         * 
         */
        public Builder imageArn(String imageArn) {
            return imageArn(Output.of(imageArn));
        }

        /**
         * @param imageName Name of the image used to create the image builder.
         * 
         * @return builder
         * 
         */
        public Builder imageName(@Nullable Output<String> imageName) {
            $.imageName = imageName;
            return this;
        }

        /**
         * @param imageName Name of the image used to create the image builder.
         * 
         * @return builder
         * 
         */
        public Builder imageName(String imageName) {
            return imageName(Output.of(imageName));
        }

        /**
         * @param instanceType Instance type to use when launching the image builder.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(Output<String> instanceType) {
            $.instanceType = instanceType;
            return this;
        }

        /**
         * @param instanceType Instance type to use when launching the image builder.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(String instanceType) {
            return instanceType(Output.of(instanceType));
        }

        /**
         * @param name Unique name for the image builder.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Unique name for the image builder.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param tags Map of tags to assign to the instance. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Map of tags to assign to the instance. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param vpcConfig Configuration block for the VPC configuration for the image builder. See below.
         * 
         * @return builder
         * 
         */
        public Builder vpcConfig(@Nullable Output<ImageBuilderVpcConfigArgs> vpcConfig) {
            $.vpcConfig = vpcConfig;
            return this;
        }

        /**
         * @param vpcConfig Configuration block for the VPC configuration for the image builder. See below.
         * 
         * @return builder
         * 
         */
        public Builder vpcConfig(ImageBuilderVpcConfigArgs vpcConfig) {
            return vpcConfig(Output.of(vpcConfig));
        }

        public ImageBuilderArgs build() {
            $.instanceType = Objects.requireNonNull($.instanceType, "expected parameter 'instanceType' to be non-null");
            return $;
        }
    }

}
