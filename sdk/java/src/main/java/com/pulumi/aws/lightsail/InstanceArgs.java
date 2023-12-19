// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.lightsail;

import com.pulumi.aws.lightsail.inputs.InstanceAddOnArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceArgs Empty = new InstanceArgs();

    /**
     * The add on configuration for the instance. Detailed below.
     * 
     */
    @Import(name="addOn")
    private @Nullable Output<InstanceAddOnArgs> addOn;

    /**
     * @return The add on configuration for the instance. Detailed below.
     * 
     */
    public Optional<Output<InstanceAddOnArgs>> addOn() {
        return Optional.ofNullable(this.addOn);
    }

    /**
     * The Availability Zone in which to create your instance. A
     * list of available zones can be obtained using the AWS CLI command:
     * [`aws lightsail get-regions --include-availability-zones`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/lightsail/get-regions.html).
     * 
     */
    @Import(name="availabilityZone", required=true)
    private Output<String> availabilityZone;

    /**
     * @return The Availability Zone in which to create your instance. A
     * list of available zones can be obtained using the AWS CLI command:
     * [`aws lightsail get-regions --include-availability-zones`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/lightsail/get-regions.html).
     * 
     */
    public Output<String> availabilityZone() {
        return this.availabilityZone;
    }

    /**
     * The ID for a virtual private server image. A list of available
     * blueprint IDs can be obtained using the AWS CLI command:
     * [`aws lightsail get-blueprints`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/lightsail/get-blueprints.html).
     * 
     */
    @Import(name="blueprintId", required=true)
    private Output<String> blueprintId;

    /**
     * @return The ID for a virtual private server image. A list of available
     * blueprint IDs can be obtained using the AWS CLI command:
     * [`aws lightsail get-blueprints`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/lightsail/get-blueprints.html).
     * 
     */
    public Output<String> blueprintId() {
        return this.blueprintId;
    }

    /**
     * The bundle of specification information. A list of available
     * bundle IDs can be obtained using the AWS CLI command:
     * [`aws lightsail get-bundles`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/lightsail/get-bundles.html).
     * 
     */
    @Import(name="bundleId", required=true)
    private Output<String> bundleId;

    /**
     * @return The bundle of specification information. A list of available
     * bundle IDs can be obtained using the AWS CLI command:
     * [`aws lightsail get-bundles`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/lightsail/get-bundles.html).
     * 
     */
    public Output<String> bundleId() {
        return this.bundleId;
    }

    /**
     * The IP address type of the Lightsail Instance. Valid Values: `dualstack` | `ipv4`.
     * 
     */
    @Import(name="ipAddressType")
    private @Nullable Output<String> ipAddressType;

    /**
     * @return The IP address type of the Lightsail Instance. Valid Values: `dualstack` | `ipv4`.
     * 
     */
    public Optional<Output<String>> ipAddressType() {
        return Optional.ofNullable(this.ipAddressType);
    }

    /**
     * The name of your key pair. Created in the
     * Lightsail console (cannot use `aws.ec2.KeyPair` at this time)
     * 
     */
    @Import(name="keyPairName")
    private @Nullable Output<String> keyPairName;

    /**
     * @return The name of your key pair. Created in the
     * Lightsail console (cannot use `aws.ec2.KeyPair` at this time)
     * 
     */
    public Optional<Output<String>> keyPairName() {
        return Optional.ofNullable(this.keyPairName);
    }

    /**
     * The name of the Lightsail Instance. Names must be unique within each AWS Region in your Lightsail account.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the Lightsail Instance. Names must be unique within each AWS Region in your Lightsail account.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Single lined launch script as a string to configure server with additional user data
     * 
     */
    @Import(name="userData")
    private @Nullable Output<String> userData;

    /**
     * @return Single lined launch script as a string to configure server with additional user data
     * 
     */
    public Optional<Output<String>> userData() {
        return Optional.ofNullable(this.userData);
    }

    private InstanceArgs() {}

    private InstanceArgs(InstanceArgs $) {
        this.addOn = $.addOn;
        this.availabilityZone = $.availabilityZone;
        this.blueprintId = $.blueprintId;
        this.bundleId = $.bundleId;
        this.ipAddressType = $.ipAddressType;
        this.keyPairName = $.keyPairName;
        this.name = $.name;
        this.tags = $.tags;
        this.userData = $.userData;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceArgs $;

        public Builder() {
            $ = new InstanceArgs();
        }

        public Builder(InstanceArgs defaults) {
            $ = new InstanceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param addOn The add on configuration for the instance. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder addOn(@Nullable Output<InstanceAddOnArgs> addOn) {
            $.addOn = addOn;
            return this;
        }

        /**
         * @param addOn The add on configuration for the instance. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder addOn(InstanceAddOnArgs addOn) {
            return addOn(Output.of(addOn));
        }

        /**
         * @param availabilityZone The Availability Zone in which to create your instance. A
         * list of available zones can be obtained using the AWS CLI command:
         * [`aws lightsail get-regions --include-availability-zones`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/lightsail/get-regions.html).
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(Output<String> availabilityZone) {
            $.availabilityZone = availabilityZone;
            return this;
        }

        /**
         * @param availabilityZone The Availability Zone in which to create your instance. A
         * list of available zones can be obtained using the AWS CLI command:
         * [`aws lightsail get-regions --include-availability-zones`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/lightsail/get-regions.html).
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(String availabilityZone) {
            return availabilityZone(Output.of(availabilityZone));
        }

        /**
         * @param blueprintId The ID for a virtual private server image. A list of available
         * blueprint IDs can be obtained using the AWS CLI command:
         * [`aws lightsail get-blueprints`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/lightsail/get-blueprints.html).
         * 
         * @return builder
         * 
         */
        public Builder blueprintId(Output<String> blueprintId) {
            $.blueprintId = blueprintId;
            return this;
        }

        /**
         * @param blueprintId The ID for a virtual private server image. A list of available
         * blueprint IDs can be obtained using the AWS CLI command:
         * [`aws lightsail get-blueprints`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/lightsail/get-blueprints.html).
         * 
         * @return builder
         * 
         */
        public Builder blueprintId(String blueprintId) {
            return blueprintId(Output.of(blueprintId));
        }

        /**
         * @param bundleId The bundle of specification information. A list of available
         * bundle IDs can be obtained using the AWS CLI command:
         * [`aws lightsail get-bundles`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/lightsail/get-bundles.html).
         * 
         * @return builder
         * 
         */
        public Builder bundleId(Output<String> bundleId) {
            $.bundleId = bundleId;
            return this;
        }

        /**
         * @param bundleId The bundle of specification information. A list of available
         * bundle IDs can be obtained using the AWS CLI command:
         * [`aws lightsail get-bundles`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/lightsail/get-bundles.html).
         * 
         * @return builder
         * 
         */
        public Builder bundleId(String bundleId) {
            return bundleId(Output.of(bundleId));
        }

        /**
         * @param ipAddressType The IP address type of the Lightsail Instance. Valid Values: `dualstack` | `ipv4`.
         * 
         * @return builder
         * 
         */
        public Builder ipAddressType(@Nullable Output<String> ipAddressType) {
            $.ipAddressType = ipAddressType;
            return this;
        }

        /**
         * @param ipAddressType The IP address type of the Lightsail Instance. Valid Values: `dualstack` | `ipv4`.
         * 
         * @return builder
         * 
         */
        public Builder ipAddressType(String ipAddressType) {
            return ipAddressType(Output.of(ipAddressType));
        }

        /**
         * @param keyPairName The name of your key pair. Created in the
         * Lightsail console (cannot use `aws.ec2.KeyPair` at this time)
         * 
         * @return builder
         * 
         */
        public Builder keyPairName(@Nullable Output<String> keyPairName) {
            $.keyPairName = keyPairName;
            return this;
        }

        /**
         * @param keyPairName The name of your key pair. Created in the
         * Lightsail console (cannot use `aws.ec2.KeyPair` at this time)
         * 
         * @return builder
         * 
         */
        public Builder keyPairName(String keyPairName) {
            return keyPairName(Output.of(keyPairName));
        }

        /**
         * @param name The name of the Lightsail Instance. Names must be unique within each AWS Region in your Lightsail account.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Lightsail Instance. Names must be unique within each AWS Region in your Lightsail account.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param tags A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param userData Single lined launch script as a string to configure server with additional user data
         * 
         * @return builder
         * 
         */
        public Builder userData(@Nullable Output<String> userData) {
            $.userData = userData;
            return this;
        }

        /**
         * @param userData Single lined launch script as a string to configure server with additional user data
         * 
         * @return builder
         * 
         */
        public Builder userData(String userData) {
            return userData(Output.of(userData));
        }

        public InstanceArgs build() {
            $.availabilityZone = Objects.requireNonNull($.availabilityZone, "expected parameter 'availabilityZone' to be non-null");
            $.blueprintId = Objects.requireNonNull($.blueprintId, "expected parameter 'blueprintId' to be non-null");
            $.bundleId = Objects.requireNonNull($.bundleId, "expected parameter 'bundleId' to be non-null");
            return $;
        }
    }

}
