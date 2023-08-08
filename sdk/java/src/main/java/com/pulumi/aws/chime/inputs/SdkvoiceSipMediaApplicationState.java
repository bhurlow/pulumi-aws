// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.chime.inputs;

import com.pulumi.aws.chime.inputs.SdkvoiceSipMediaApplicationEndpointsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SdkvoiceSipMediaApplicationState extends com.pulumi.resources.ResourceArgs {

    public static final SdkvoiceSipMediaApplicationState Empty = new SdkvoiceSipMediaApplicationState();

    /**
     * ARN (Amazon Resource Name) of the AWS Chime SDK Voice Sip Media Application
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return ARN (Amazon Resource Name) of the AWS Chime SDK Voice Sip Media Application
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The AWS Region in which the AWS Chime SDK Voice Sip Media Application is created.
     * 
     */
    @Import(name="awsRegion")
    private @Nullable Output<String> awsRegion;

    /**
     * @return The AWS Region in which the AWS Chime SDK Voice Sip Media Application is created.
     * 
     */
    public Optional<Output<String>> awsRegion() {
        return Optional.ofNullable(this.awsRegion);
    }

    /**
     * List of endpoints (Lambda Amazon Resource Names) specified for the SIP media application. Currently, only one endpoint is supported. See `endpoints`.
     * 
     */
    @Import(name="endpoints")
    private @Nullable Output<SdkvoiceSipMediaApplicationEndpointsArgs> endpoints;

    /**
     * @return List of endpoints (Lambda Amazon Resource Names) specified for the SIP media application. Currently, only one endpoint is supported. See `endpoints`.
     * 
     */
    public Optional<Output<SdkvoiceSipMediaApplicationEndpointsArgs>> endpoints() {
        return Optional.ofNullable(this.endpoints);
    }

    /**
     * The name of the AWS Chime SDK Voice Sip Media Application.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the AWS Chime SDK Voice Sip Media Application.
     * 
     * The following arguments are optional:
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
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

    private SdkvoiceSipMediaApplicationState() {}

    private SdkvoiceSipMediaApplicationState(SdkvoiceSipMediaApplicationState $) {
        this.arn = $.arn;
        this.awsRegion = $.awsRegion;
        this.endpoints = $.endpoints;
        this.name = $.name;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SdkvoiceSipMediaApplicationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SdkvoiceSipMediaApplicationState $;

        public Builder() {
            $ = new SdkvoiceSipMediaApplicationState();
        }

        public Builder(SdkvoiceSipMediaApplicationState defaults) {
            $ = new SdkvoiceSipMediaApplicationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn ARN (Amazon Resource Name) of the AWS Chime SDK Voice Sip Media Application
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn ARN (Amazon Resource Name) of the AWS Chime SDK Voice Sip Media Application
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param awsRegion The AWS Region in which the AWS Chime SDK Voice Sip Media Application is created.
         * 
         * @return builder
         * 
         */
        public Builder awsRegion(@Nullable Output<String> awsRegion) {
            $.awsRegion = awsRegion;
            return this;
        }

        /**
         * @param awsRegion The AWS Region in which the AWS Chime SDK Voice Sip Media Application is created.
         * 
         * @return builder
         * 
         */
        public Builder awsRegion(String awsRegion) {
            return awsRegion(Output.of(awsRegion));
        }

        /**
         * @param endpoints List of endpoints (Lambda Amazon Resource Names) specified for the SIP media application. Currently, only one endpoint is supported. See `endpoints`.
         * 
         * @return builder
         * 
         */
        public Builder endpoints(@Nullable Output<SdkvoiceSipMediaApplicationEndpointsArgs> endpoints) {
            $.endpoints = endpoints;
            return this;
        }

        /**
         * @param endpoints List of endpoints (Lambda Amazon Resource Names) specified for the SIP media application. Currently, only one endpoint is supported. See `endpoints`.
         * 
         * @return builder
         * 
         */
        public Builder endpoints(SdkvoiceSipMediaApplicationEndpointsArgs endpoints) {
            return endpoints(Output.of(endpoints));
        }

        /**
         * @param name The name of the AWS Chime SDK Voice Sip Media Application.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the AWS Chime SDK Voice Sip Media Application.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
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

        public SdkvoiceSipMediaApplicationState build() {
            return $;
        }
    }

}
