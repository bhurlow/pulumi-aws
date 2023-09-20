// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.apprunner.inputs;

import com.pulumi.aws.apprunner.inputs.ServiceEncryptionConfigurationArgs;
import com.pulumi.aws.apprunner.inputs.ServiceHealthCheckConfigurationArgs;
import com.pulumi.aws.apprunner.inputs.ServiceInstanceConfigurationArgs;
import com.pulumi.aws.apprunner.inputs.ServiceNetworkConfigurationArgs;
import com.pulumi.aws.apprunner.inputs.ServiceObservabilityConfigurationArgs;
import com.pulumi.aws.apprunner.inputs.ServiceSourceConfigurationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceState Empty = new ServiceState();

    /**
     * ARN of the App Runner service.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return ARN of the App Runner service.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * ARN of an App Runner automatic scaling configuration resource that you want to associate with your service. If not provided, App Runner associates the latest revision of a default auto scaling configuration.
     * 
     */
    @Import(name="autoScalingConfigurationArn")
    private @Nullable Output<String> autoScalingConfigurationArn;

    /**
     * @return ARN of an App Runner automatic scaling configuration resource that you want to associate with your service. If not provided, App Runner associates the latest revision of a default auto scaling configuration.
     * 
     */
    public Optional<Output<String>> autoScalingConfigurationArn() {
        return Optional.ofNullable(this.autoScalingConfigurationArn);
    }

    /**
     * An optional custom encryption key that App Runner uses to encrypt the copy of your source repository that it maintains and your service logs. By default, App Runner uses an AWS managed CMK. See Encryption Configuration below for more details.
     * 
     */
    @Import(name="encryptionConfiguration")
    private @Nullable Output<ServiceEncryptionConfigurationArgs> encryptionConfiguration;

    /**
     * @return An optional custom encryption key that App Runner uses to encrypt the copy of your source repository that it maintains and your service logs. By default, App Runner uses an AWS managed CMK. See Encryption Configuration below for more details.
     * 
     */
    public Optional<Output<ServiceEncryptionConfigurationArgs>> encryptionConfiguration() {
        return Optional.ofNullable(this.encryptionConfiguration);
    }

    /**
     * Settings of the health check that AWS App Runner performs to monitor the health of your service. See Health Check Configuration below for more details.
     * 
     */
    @Import(name="healthCheckConfiguration")
    private @Nullable Output<ServiceHealthCheckConfigurationArgs> healthCheckConfiguration;

    /**
     * @return Settings of the health check that AWS App Runner performs to monitor the health of your service. See Health Check Configuration below for more details.
     * 
     */
    public Optional<Output<ServiceHealthCheckConfigurationArgs>> healthCheckConfiguration() {
        return Optional.ofNullable(this.healthCheckConfiguration);
    }

    /**
     * The runtime configuration of instances (scaling units) of the App Runner service. See Instance Configuration below for more details.
     * 
     */
    @Import(name="instanceConfiguration")
    private @Nullable Output<ServiceInstanceConfigurationArgs> instanceConfiguration;

    /**
     * @return The runtime configuration of instances (scaling units) of the App Runner service. See Instance Configuration below for more details.
     * 
     */
    public Optional<Output<ServiceInstanceConfigurationArgs>> instanceConfiguration() {
        return Optional.ofNullable(this.instanceConfiguration);
    }

    /**
     * Configuration settings related to network traffic of the web application that the App Runner service runs. See Network Configuration below for more details.
     * 
     */
    @Import(name="networkConfiguration")
    private @Nullable Output<ServiceNetworkConfigurationArgs> networkConfiguration;

    /**
     * @return Configuration settings related to network traffic of the web application that the App Runner service runs. See Network Configuration below for more details.
     * 
     */
    public Optional<Output<ServiceNetworkConfigurationArgs>> networkConfiguration() {
        return Optional.ofNullable(this.networkConfiguration);
    }

    /**
     * The observability configuration of your service. See Observability Configuration below for more details.
     * 
     */
    @Import(name="observabilityConfiguration")
    private @Nullable Output<ServiceObservabilityConfigurationArgs> observabilityConfiguration;

    /**
     * @return The observability configuration of your service. See Observability Configuration below for more details.
     * 
     */
    public Optional<Output<ServiceObservabilityConfigurationArgs>> observabilityConfiguration() {
        return Optional.ofNullable(this.observabilityConfiguration);
    }

    /**
     * An alphanumeric ID that App Runner generated for this service. Unique within the AWS Region.
     * 
     */
    @Import(name="serviceId")
    private @Nullable Output<String> serviceId;

    /**
     * @return An alphanumeric ID that App Runner generated for this service. Unique within the AWS Region.
     * 
     */
    public Optional<Output<String>> serviceId() {
        return Optional.ofNullable(this.serviceId);
    }

    /**
     * Name of the service.
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return Name of the service.
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * Subdomain URL that App Runner generated for this service. You can use this URL to access your service web application.
     * 
     */
    @Import(name="serviceUrl")
    private @Nullable Output<String> serviceUrl;

    /**
     * @return Subdomain URL that App Runner generated for this service. You can use this URL to access your service web application.
     * 
     */
    public Optional<Output<String>> serviceUrl() {
        return Optional.ofNullable(this.serviceUrl);
    }

    /**
     * The source to deploy to the App Runner service. Can be a code or an image repository. See Source Configuration below for more details.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="sourceConfiguration")
    private @Nullable Output<ServiceSourceConfigurationArgs> sourceConfiguration;

    /**
     * @return The source to deploy to the App Runner service. Can be a code or an image repository. See Source Configuration below for more details.
     * 
     * The following arguments are optional:
     * 
     */
    public Optional<Output<ServiceSourceConfigurationArgs>> sourceConfiguration() {
        return Optional.ofNullable(this.sourceConfiguration);
    }

    /**
     * Current state of the App Runner service.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Current state of the App Runner service.
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

    private ServiceState() {}

    private ServiceState(ServiceState $) {
        this.arn = $.arn;
        this.autoScalingConfigurationArn = $.autoScalingConfigurationArn;
        this.encryptionConfiguration = $.encryptionConfiguration;
        this.healthCheckConfiguration = $.healthCheckConfiguration;
        this.instanceConfiguration = $.instanceConfiguration;
        this.networkConfiguration = $.networkConfiguration;
        this.observabilityConfiguration = $.observabilityConfiguration;
        this.serviceId = $.serviceId;
        this.serviceName = $.serviceName;
        this.serviceUrl = $.serviceUrl;
        this.sourceConfiguration = $.sourceConfiguration;
        this.status = $.status;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceState $;

        public Builder() {
            $ = new ServiceState();
        }

        public Builder(ServiceState defaults) {
            $ = new ServiceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn ARN of the App Runner service.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn ARN of the App Runner service.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param autoScalingConfigurationArn ARN of an App Runner automatic scaling configuration resource that you want to associate with your service. If not provided, App Runner associates the latest revision of a default auto scaling configuration.
         * 
         * @return builder
         * 
         */
        public Builder autoScalingConfigurationArn(@Nullable Output<String> autoScalingConfigurationArn) {
            $.autoScalingConfigurationArn = autoScalingConfigurationArn;
            return this;
        }

        /**
         * @param autoScalingConfigurationArn ARN of an App Runner automatic scaling configuration resource that you want to associate with your service. If not provided, App Runner associates the latest revision of a default auto scaling configuration.
         * 
         * @return builder
         * 
         */
        public Builder autoScalingConfigurationArn(String autoScalingConfigurationArn) {
            return autoScalingConfigurationArn(Output.of(autoScalingConfigurationArn));
        }

        /**
         * @param encryptionConfiguration An optional custom encryption key that App Runner uses to encrypt the copy of your source repository that it maintains and your service logs. By default, App Runner uses an AWS managed CMK. See Encryption Configuration below for more details.
         * 
         * @return builder
         * 
         */
        public Builder encryptionConfiguration(@Nullable Output<ServiceEncryptionConfigurationArgs> encryptionConfiguration) {
            $.encryptionConfiguration = encryptionConfiguration;
            return this;
        }

        /**
         * @param encryptionConfiguration An optional custom encryption key that App Runner uses to encrypt the copy of your source repository that it maintains and your service logs. By default, App Runner uses an AWS managed CMK. See Encryption Configuration below for more details.
         * 
         * @return builder
         * 
         */
        public Builder encryptionConfiguration(ServiceEncryptionConfigurationArgs encryptionConfiguration) {
            return encryptionConfiguration(Output.of(encryptionConfiguration));
        }

        /**
         * @param healthCheckConfiguration Settings of the health check that AWS App Runner performs to monitor the health of your service. See Health Check Configuration below for more details.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckConfiguration(@Nullable Output<ServiceHealthCheckConfigurationArgs> healthCheckConfiguration) {
            $.healthCheckConfiguration = healthCheckConfiguration;
            return this;
        }

        /**
         * @param healthCheckConfiguration Settings of the health check that AWS App Runner performs to monitor the health of your service. See Health Check Configuration below for more details.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckConfiguration(ServiceHealthCheckConfigurationArgs healthCheckConfiguration) {
            return healthCheckConfiguration(Output.of(healthCheckConfiguration));
        }

        /**
         * @param instanceConfiguration The runtime configuration of instances (scaling units) of the App Runner service. See Instance Configuration below for more details.
         * 
         * @return builder
         * 
         */
        public Builder instanceConfiguration(@Nullable Output<ServiceInstanceConfigurationArgs> instanceConfiguration) {
            $.instanceConfiguration = instanceConfiguration;
            return this;
        }

        /**
         * @param instanceConfiguration The runtime configuration of instances (scaling units) of the App Runner service. See Instance Configuration below for more details.
         * 
         * @return builder
         * 
         */
        public Builder instanceConfiguration(ServiceInstanceConfigurationArgs instanceConfiguration) {
            return instanceConfiguration(Output.of(instanceConfiguration));
        }

        /**
         * @param networkConfiguration Configuration settings related to network traffic of the web application that the App Runner service runs. See Network Configuration below for more details.
         * 
         * @return builder
         * 
         */
        public Builder networkConfiguration(@Nullable Output<ServiceNetworkConfigurationArgs> networkConfiguration) {
            $.networkConfiguration = networkConfiguration;
            return this;
        }

        /**
         * @param networkConfiguration Configuration settings related to network traffic of the web application that the App Runner service runs. See Network Configuration below for more details.
         * 
         * @return builder
         * 
         */
        public Builder networkConfiguration(ServiceNetworkConfigurationArgs networkConfiguration) {
            return networkConfiguration(Output.of(networkConfiguration));
        }

        /**
         * @param observabilityConfiguration The observability configuration of your service. See Observability Configuration below for more details.
         * 
         * @return builder
         * 
         */
        public Builder observabilityConfiguration(@Nullable Output<ServiceObservabilityConfigurationArgs> observabilityConfiguration) {
            $.observabilityConfiguration = observabilityConfiguration;
            return this;
        }

        /**
         * @param observabilityConfiguration The observability configuration of your service. See Observability Configuration below for more details.
         * 
         * @return builder
         * 
         */
        public Builder observabilityConfiguration(ServiceObservabilityConfigurationArgs observabilityConfiguration) {
            return observabilityConfiguration(Output.of(observabilityConfiguration));
        }

        /**
         * @param serviceId An alphanumeric ID that App Runner generated for this service. Unique within the AWS Region.
         * 
         * @return builder
         * 
         */
        public Builder serviceId(@Nullable Output<String> serviceId) {
            $.serviceId = serviceId;
            return this;
        }

        /**
         * @param serviceId An alphanumeric ID that App Runner generated for this service. Unique within the AWS Region.
         * 
         * @return builder
         * 
         */
        public Builder serviceId(String serviceId) {
            return serviceId(Output.of(serviceId));
        }

        /**
         * @param serviceName Name of the service.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName Name of the service.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param serviceUrl Subdomain URL that App Runner generated for this service. You can use this URL to access your service web application.
         * 
         * @return builder
         * 
         */
        public Builder serviceUrl(@Nullable Output<String> serviceUrl) {
            $.serviceUrl = serviceUrl;
            return this;
        }

        /**
         * @param serviceUrl Subdomain URL that App Runner generated for this service. You can use this URL to access your service web application.
         * 
         * @return builder
         * 
         */
        public Builder serviceUrl(String serviceUrl) {
            return serviceUrl(Output.of(serviceUrl));
        }

        /**
         * @param sourceConfiguration The source to deploy to the App Runner service. Can be a code or an image repository. See Source Configuration below for more details.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder sourceConfiguration(@Nullable Output<ServiceSourceConfigurationArgs> sourceConfiguration) {
            $.sourceConfiguration = sourceConfiguration;
            return this;
        }

        /**
         * @param sourceConfiguration The source to deploy to the App Runner service. Can be a code or an image repository. See Source Configuration below for more details.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder sourceConfiguration(ServiceSourceConfigurationArgs sourceConfiguration) {
            return sourceConfiguration(Output.of(sourceConfiguration));
        }

        /**
         * @param status Current state of the App Runner service.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Current state of the App Runner service.
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

        public ServiceState build() {
            return $;
        }
    }

}
