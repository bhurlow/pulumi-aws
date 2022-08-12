// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.opsworks.outputs;

import com.pulumi.aws.opsworks.outputs.NodejsAppLayerCloudwatchConfigurationLogStream;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class NodejsAppLayerCloudwatchConfiguration {
    private final @Nullable Boolean enabled;
    private final @Nullable List<NodejsAppLayerCloudwatchConfigurationLogStream> logStreams;

    @CustomType.Constructor
    private NodejsAppLayerCloudwatchConfiguration(
        @CustomType.Parameter("enabled") @Nullable Boolean enabled,
        @CustomType.Parameter("logStreams") @Nullable List<NodejsAppLayerCloudwatchConfigurationLogStream> logStreams) {
        this.enabled = enabled;
        this.logStreams = logStreams;
    }

    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
    }
    public List<NodejsAppLayerCloudwatchConfigurationLogStream> logStreams() {
        return this.logStreams == null ? List.of() : this.logStreams;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(NodejsAppLayerCloudwatchConfiguration defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable Boolean enabled;
        private @Nullable List<NodejsAppLayerCloudwatchConfigurationLogStream> logStreams;

        public Builder() {
    	      // Empty
        }

        public Builder(NodejsAppLayerCloudwatchConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enabled = defaults.enabled;
    	      this.logStreams = defaults.logStreams;
        }

        public Builder enabled(@Nullable Boolean enabled) {
            this.enabled = enabled;
            return this;
        }
        public Builder logStreams(@Nullable List<NodejsAppLayerCloudwatchConfigurationLogStream> logStreams) {
            this.logStreams = logStreams;
            return this;
        }
        public Builder logStreams(NodejsAppLayerCloudwatchConfigurationLogStream... logStreams) {
            return logStreams(List.of(logStreams));
        }        public NodejsAppLayerCloudwatchConfiguration build() {
            return new NodejsAppLayerCloudwatchConfiguration(enabled, logStreams);
        }
    }
}