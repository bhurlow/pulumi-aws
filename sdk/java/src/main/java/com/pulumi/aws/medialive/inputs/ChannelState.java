// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.inputs;

import com.pulumi.aws.medialive.inputs.ChannelCdiInputSpecificationArgs;
import com.pulumi.aws.medialive.inputs.ChannelDestinationArgs;
import com.pulumi.aws.medialive.inputs.ChannelEncoderSettingsArgs;
import com.pulumi.aws.medialive.inputs.ChannelInputAttachmentArgs;
import com.pulumi.aws.medialive.inputs.ChannelInputSpecificationArgs;
import com.pulumi.aws.medialive.inputs.ChannelMaintenanceArgs;
import com.pulumi.aws.medialive.inputs.ChannelVpcArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ChannelState extends com.pulumi.resources.ResourceArgs {

    public static final ChannelState Empty = new ChannelState();

    /**
     * ARN of the Channel.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return ARN of the Channel.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Specification of CDI inputs for this channel. See CDI Input Specification for more details.
     * 
     */
    @Import(name="cdiInputSpecification")
    private @Nullable Output<ChannelCdiInputSpecificationArgs> cdiInputSpecification;

    /**
     * @return Specification of CDI inputs for this channel. See CDI Input Specification for more details.
     * 
     */
    public Optional<Output<ChannelCdiInputSpecificationArgs>> cdiInputSpecification() {
        return Optional.ofNullable(this.cdiInputSpecification);
    }

    /**
     * Concise argument description.
     * 
     */
    @Import(name="channelClass")
    private @Nullable Output<String> channelClass;

    /**
     * @return Concise argument description.
     * 
     */
    public Optional<Output<String>> channelClass() {
        return Optional.ofNullable(this.channelClass);
    }

    /**
     * ID of the channel in MediaPackage that is the destination for this output group.
     * 
     */
    @Import(name="channelId")
    private @Nullable Output<String> channelId;

    /**
     * @return ID of the channel in MediaPackage that is the destination for this output group.
     * 
     */
    public Optional<Output<String>> channelId() {
        return Optional.ofNullable(this.channelId);
    }

    /**
     * Destinations for channel. See Destinations for more details.
     * 
     */
    @Import(name="destinations")
    private @Nullable Output<List<ChannelDestinationArgs>> destinations;

    /**
     * @return Destinations for channel. See Destinations for more details.
     * 
     */
    public Optional<Output<List<ChannelDestinationArgs>>> destinations() {
        return Optional.ofNullable(this.destinations);
    }

    /**
     * Encoder settings. See Encoder Settings for more details.
     * 
     */
    @Import(name="encoderSettings")
    private @Nullable Output<ChannelEncoderSettingsArgs> encoderSettings;

    /**
     * @return Encoder settings. See Encoder Settings for more details.
     * 
     */
    public Optional<Output<ChannelEncoderSettingsArgs>> encoderSettings() {
        return Optional.ofNullable(this.encoderSettings);
    }

    /**
     * Input attachments for the channel. See Input Attachments for more details.
     * 
     */
    @Import(name="inputAttachments")
    private @Nullable Output<List<ChannelInputAttachmentArgs>> inputAttachments;

    /**
     * @return Input attachments for the channel. See Input Attachments for more details.
     * 
     */
    public Optional<Output<List<ChannelInputAttachmentArgs>>> inputAttachments() {
        return Optional.ofNullable(this.inputAttachments);
    }

    /**
     * Specification of network and file inputs for the channel.
     * 
     */
    @Import(name="inputSpecification")
    private @Nullable Output<ChannelInputSpecificationArgs> inputSpecification;

    /**
     * @return Specification of network and file inputs for the channel.
     * 
     */
    public Optional<Output<ChannelInputSpecificationArgs>> inputSpecification() {
        return Optional.ofNullable(this.inputSpecification);
    }

    /**
     * The log level to write to Cloudwatch logs.
     * 
     */
    @Import(name="logLevel")
    private @Nullable Output<String> logLevel;

    /**
     * @return The log level to write to Cloudwatch logs.
     * 
     */
    public Optional<Output<String>> logLevel() {
        return Optional.ofNullable(this.logLevel);
    }

    /**
     * Maintenance settings for this channel. See Maintenance for more details.
     * 
     */
    @Import(name="maintenance")
    private @Nullable Output<ChannelMaintenanceArgs> maintenance;

    /**
     * @return Maintenance settings for this channel. See Maintenance for more details.
     * 
     */
    public Optional<Output<ChannelMaintenanceArgs>> maintenance() {
        return Optional.ofNullable(this.maintenance);
    }

    /**
     * Custom output group name defined by the user.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Custom output group name defined by the user.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Concise argument description.
     * 
     */
    @Import(name="roleArn")
    private @Nullable Output<String> roleArn;

    /**
     * @return Concise argument description.
     * 
     */
    public Optional<Output<String>> roleArn() {
        return Optional.ofNullable(this.roleArn);
    }

    /**
     * A map of tags to assign to the channel. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the channel. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * Settings for the VPC outputs.
     * 
     */
    @Import(name="vpc")
    private @Nullable Output<ChannelVpcArgs> vpc;

    /**
     * @return Settings for the VPC outputs.
     * 
     */
    public Optional<Output<ChannelVpcArgs>> vpc() {
        return Optional.ofNullable(this.vpc);
    }

    private ChannelState() {}

    private ChannelState(ChannelState $) {
        this.arn = $.arn;
        this.cdiInputSpecification = $.cdiInputSpecification;
        this.channelClass = $.channelClass;
        this.channelId = $.channelId;
        this.destinations = $.destinations;
        this.encoderSettings = $.encoderSettings;
        this.inputAttachments = $.inputAttachments;
        this.inputSpecification = $.inputSpecification;
        this.logLevel = $.logLevel;
        this.maintenance = $.maintenance;
        this.name = $.name;
        this.roleArn = $.roleArn;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.vpc = $.vpc;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ChannelState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ChannelState $;

        public Builder() {
            $ = new ChannelState();
        }

        public Builder(ChannelState defaults) {
            $ = new ChannelState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn ARN of the Channel.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn ARN of the Channel.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param cdiInputSpecification Specification of CDI inputs for this channel. See CDI Input Specification for more details.
         * 
         * @return builder
         * 
         */
        public Builder cdiInputSpecification(@Nullable Output<ChannelCdiInputSpecificationArgs> cdiInputSpecification) {
            $.cdiInputSpecification = cdiInputSpecification;
            return this;
        }

        /**
         * @param cdiInputSpecification Specification of CDI inputs for this channel. See CDI Input Specification for more details.
         * 
         * @return builder
         * 
         */
        public Builder cdiInputSpecification(ChannelCdiInputSpecificationArgs cdiInputSpecification) {
            return cdiInputSpecification(Output.of(cdiInputSpecification));
        }

        /**
         * @param channelClass Concise argument description.
         * 
         * @return builder
         * 
         */
        public Builder channelClass(@Nullable Output<String> channelClass) {
            $.channelClass = channelClass;
            return this;
        }

        /**
         * @param channelClass Concise argument description.
         * 
         * @return builder
         * 
         */
        public Builder channelClass(String channelClass) {
            return channelClass(Output.of(channelClass));
        }

        /**
         * @param channelId ID of the channel in MediaPackage that is the destination for this output group.
         * 
         * @return builder
         * 
         */
        public Builder channelId(@Nullable Output<String> channelId) {
            $.channelId = channelId;
            return this;
        }

        /**
         * @param channelId ID of the channel in MediaPackage that is the destination for this output group.
         * 
         * @return builder
         * 
         */
        public Builder channelId(String channelId) {
            return channelId(Output.of(channelId));
        }

        /**
         * @param destinations Destinations for channel. See Destinations for more details.
         * 
         * @return builder
         * 
         */
        public Builder destinations(@Nullable Output<List<ChannelDestinationArgs>> destinations) {
            $.destinations = destinations;
            return this;
        }

        /**
         * @param destinations Destinations for channel. See Destinations for more details.
         * 
         * @return builder
         * 
         */
        public Builder destinations(List<ChannelDestinationArgs> destinations) {
            return destinations(Output.of(destinations));
        }

        /**
         * @param destinations Destinations for channel. See Destinations for more details.
         * 
         * @return builder
         * 
         */
        public Builder destinations(ChannelDestinationArgs... destinations) {
            return destinations(List.of(destinations));
        }

        /**
         * @param encoderSettings Encoder settings. See Encoder Settings for more details.
         * 
         * @return builder
         * 
         */
        public Builder encoderSettings(@Nullable Output<ChannelEncoderSettingsArgs> encoderSettings) {
            $.encoderSettings = encoderSettings;
            return this;
        }

        /**
         * @param encoderSettings Encoder settings. See Encoder Settings for more details.
         * 
         * @return builder
         * 
         */
        public Builder encoderSettings(ChannelEncoderSettingsArgs encoderSettings) {
            return encoderSettings(Output.of(encoderSettings));
        }

        /**
         * @param inputAttachments Input attachments for the channel. See Input Attachments for more details.
         * 
         * @return builder
         * 
         */
        public Builder inputAttachments(@Nullable Output<List<ChannelInputAttachmentArgs>> inputAttachments) {
            $.inputAttachments = inputAttachments;
            return this;
        }

        /**
         * @param inputAttachments Input attachments for the channel. See Input Attachments for more details.
         * 
         * @return builder
         * 
         */
        public Builder inputAttachments(List<ChannelInputAttachmentArgs> inputAttachments) {
            return inputAttachments(Output.of(inputAttachments));
        }

        /**
         * @param inputAttachments Input attachments for the channel. See Input Attachments for more details.
         * 
         * @return builder
         * 
         */
        public Builder inputAttachments(ChannelInputAttachmentArgs... inputAttachments) {
            return inputAttachments(List.of(inputAttachments));
        }

        /**
         * @param inputSpecification Specification of network and file inputs for the channel.
         * 
         * @return builder
         * 
         */
        public Builder inputSpecification(@Nullable Output<ChannelInputSpecificationArgs> inputSpecification) {
            $.inputSpecification = inputSpecification;
            return this;
        }

        /**
         * @param inputSpecification Specification of network and file inputs for the channel.
         * 
         * @return builder
         * 
         */
        public Builder inputSpecification(ChannelInputSpecificationArgs inputSpecification) {
            return inputSpecification(Output.of(inputSpecification));
        }

        /**
         * @param logLevel The log level to write to Cloudwatch logs.
         * 
         * @return builder
         * 
         */
        public Builder logLevel(@Nullable Output<String> logLevel) {
            $.logLevel = logLevel;
            return this;
        }

        /**
         * @param logLevel The log level to write to Cloudwatch logs.
         * 
         * @return builder
         * 
         */
        public Builder logLevel(String logLevel) {
            return logLevel(Output.of(logLevel));
        }

        /**
         * @param maintenance Maintenance settings for this channel. See Maintenance for more details.
         * 
         * @return builder
         * 
         */
        public Builder maintenance(@Nullable Output<ChannelMaintenanceArgs> maintenance) {
            $.maintenance = maintenance;
            return this;
        }

        /**
         * @param maintenance Maintenance settings for this channel. See Maintenance for more details.
         * 
         * @return builder
         * 
         */
        public Builder maintenance(ChannelMaintenanceArgs maintenance) {
            return maintenance(Output.of(maintenance));
        }

        /**
         * @param name Custom output group name defined by the user.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Custom output group name defined by the user.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param roleArn Concise argument description.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(@Nullable Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn Concise argument description.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        /**
         * @param tags A map of tags to assign to the channel. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the channel. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param vpc Settings for the VPC outputs.
         * 
         * @return builder
         * 
         */
        public Builder vpc(@Nullable Output<ChannelVpcArgs> vpc) {
            $.vpc = vpc;
            return this;
        }

        /**
         * @param vpc Settings for the VPC outputs.
         * 
         * @return builder
         * 
         */
        public Builder vpc(ChannelVpcArgs vpc) {
            return vpc(Output.of(vpc));
        }

        public ChannelState build() {
            return $;
        }
    }

}
