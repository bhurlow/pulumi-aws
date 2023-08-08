// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.auditmanager.inputs;

import com.pulumi.aws.auditmanager.inputs.GetControlControlMappingSourceSourceKeywordArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetControlControlMappingSourceArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetControlControlMappingSourceArgs Empty = new GetControlControlMappingSourceArgs();

    @Import(name="sourceDescription", required=true)
    private Output<String> sourceDescription;

    public Output<String> sourceDescription() {
        return this.sourceDescription;
    }

    @Import(name="sourceFrequency", required=true)
    private Output<String> sourceFrequency;

    public Output<String> sourceFrequency() {
        return this.sourceFrequency;
    }

    @Import(name="sourceId", required=true)
    private Output<String> sourceId;

    public Output<String> sourceId() {
        return this.sourceId;
    }

    @Import(name="sourceKeyword")
    private @Nullable Output<GetControlControlMappingSourceSourceKeywordArgs> sourceKeyword;

    public Optional<Output<GetControlControlMappingSourceSourceKeywordArgs>> sourceKeyword() {
        return Optional.ofNullable(this.sourceKeyword);
    }

    @Import(name="sourceName", required=true)
    private Output<String> sourceName;

    public Output<String> sourceName() {
        return this.sourceName;
    }

    @Import(name="sourceSetUpOption", required=true)
    private Output<String> sourceSetUpOption;

    public Output<String> sourceSetUpOption() {
        return this.sourceSetUpOption;
    }

    @Import(name="sourceType", required=true)
    private Output<String> sourceType;

    public Output<String> sourceType() {
        return this.sourceType;
    }

    @Import(name="troubleshootingText", required=true)
    private Output<String> troubleshootingText;

    public Output<String> troubleshootingText() {
        return this.troubleshootingText;
    }

    private GetControlControlMappingSourceArgs() {}

    private GetControlControlMappingSourceArgs(GetControlControlMappingSourceArgs $) {
        this.sourceDescription = $.sourceDescription;
        this.sourceFrequency = $.sourceFrequency;
        this.sourceId = $.sourceId;
        this.sourceKeyword = $.sourceKeyword;
        this.sourceName = $.sourceName;
        this.sourceSetUpOption = $.sourceSetUpOption;
        this.sourceType = $.sourceType;
        this.troubleshootingText = $.troubleshootingText;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetControlControlMappingSourceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetControlControlMappingSourceArgs $;

        public Builder() {
            $ = new GetControlControlMappingSourceArgs();
        }

        public Builder(GetControlControlMappingSourceArgs defaults) {
            $ = new GetControlControlMappingSourceArgs(Objects.requireNonNull(defaults));
        }

        public Builder sourceDescription(Output<String> sourceDescription) {
            $.sourceDescription = sourceDescription;
            return this;
        }

        public Builder sourceDescription(String sourceDescription) {
            return sourceDescription(Output.of(sourceDescription));
        }

        public Builder sourceFrequency(Output<String> sourceFrequency) {
            $.sourceFrequency = sourceFrequency;
            return this;
        }

        public Builder sourceFrequency(String sourceFrequency) {
            return sourceFrequency(Output.of(sourceFrequency));
        }

        public Builder sourceId(Output<String> sourceId) {
            $.sourceId = sourceId;
            return this;
        }

        public Builder sourceId(String sourceId) {
            return sourceId(Output.of(sourceId));
        }

        public Builder sourceKeyword(@Nullable Output<GetControlControlMappingSourceSourceKeywordArgs> sourceKeyword) {
            $.sourceKeyword = sourceKeyword;
            return this;
        }

        public Builder sourceKeyword(GetControlControlMappingSourceSourceKeywordArgs sourceKeyword) {
            return sourceKeyword(Output.of(sourceKeyword));
        }

        public Builder sourceName(Output<String> sourceName) {
            $.sourceName = sourceName;
            return this;
        }

        public Builder sourceName(String sourceName) {
            return sourceName(Output.of(sourceName));
        }

        public Builder sourceSetUpOption(Output<String> sourceSetUpOption) {
            $.sourceSetUpOption = sourceSetUpOption;
            return this;
        }

        public Builder sourceSetUpOption(String sourceSetUpOption) {
            return sourceSetUpOption(Output.of(sourceSetUpOption));
        }

        public Builder sourceType(Output<String> sourceType) {
            $.sourceType = sourceType;
            return this;
        }

        public Builder sourceType(String sourceType) {
            return sourceType(Output.of(sourceType));
        }

        public Builder troubleshootingText(Output<String> troubleshootingText) {
            $.troubleshootingText = troubleshootingText;
            return this;
        }

        public Builder troubleshootingText(String troubleshootingText) {
            return troubleshootingText(Output.of(troubleshootingText));
        }

        public GetControlControlMappingSourceArgs build() {
            $.sourceDescription = Objects.requireNonNull($.sourceDescription, "expected parameter 'sourceDescription' to be non-null");
            $.sourceFrequency = Objects.requireNonNull($.sourceFrequency, "expected parameter 'sourceFrequency' to be non-null");
            $.sourceId = Objects.requireNonNull($.sourceId, "expected parameter 'sourceId' to be non-null");
            $.sourceName = Objects.requireNonNull($.sourceName, "expected parameter 'sourceName' to be non-null");
            $.sourceSetUpOption = Objects.requireNonNull($.sourceSetUpOption, "expected parameter 'sourceSetUpOption' to be non-null");
            $.sourceType = Objects.requireNonNull($.sourceType, "expected parameter 'sourceType' to be non-null");
            $.troubleshootingText = Objects.requireNonNull($.troubleshootingText, "expected parameter 'troubleshootingText' to be non-null");
            return $;
        }
    }

}
