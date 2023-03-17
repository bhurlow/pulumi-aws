// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.transfer.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServerProtocolDetailsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServerProtocolDetailsArgs Empty = new ServerProtocolDetailsArgs();

    /**
     * Indicates the transport method for the AS2 messages. Currently, only `HTTP` is supported.
     * 
     */
    @Import(name="as2Transports")
    private @Nullable Output<List<String>> as2Transports;

    /**
     * @return Indicates the transport method for the AS2 messages. Currently, only `HTTP` is supported.
     * 
     */
    public Optional<Output<List<String>>> as2Transports() {
        return Optional.ofNullable(this.as2Transports);
    }

    /**
     * Indicates passive mode, for FTP and FTPS protocols. Enter a single IPv4 address, such as the public IP address of a firewall, router, or load balancer.
     * 
     */
    @Import(name="passiveIp")
    private @Nullable Output<String> passiveIp;

    /**
     * @return Indicates passive mode, for FTP and FTPS protocols. Enter a single IPv4 address, such as the public IP address of a firewall, router, or load balancer.
     * 
     */
    public Optional<Output<String>> passiveIp() {
        return Optional.ofNullable(this.passiveIp);
    }

    /**
     * Use to ignore the error that is generated when the client attempts to use `SETSTAT` on a file you are uploading to an S3 bucket. Valid values: `DEFAULT`, `ENABLE_NO_OP`.
     * 
     */
    @Import(name="setStatOption")
    private @Nullable Output<String> setStatOption;

    /**
     * @return Use to ignore the error that is generated when the client attempts to use `SETSTAT` on a file you are uploading to an S3 bucket. Valid values: `DEFAULT`, `ENABLE_NO_OP`.
     * 
     */
    public Optional<Output<String>> setStatOption() {
        return Optional.ofNullable(this.setStatOption);
    }

    /**
     * A property used with Transfer Family servers that use the FTPS protocol. Provides a mechanism to resume or share a negotiated secret key between the control and data connection for an FTPS session. Valid values: `DISABLED`, `ENABLED`, `ENFORCED`.
     * 
     */
    @Import(name="tlsSessionResumptionMode")
    private @Nullable Output<String> tlsSessionResumptionMode;

    /**
     * @return A property used with Transfer Family servers that use the FTPS protocol. Provides a mechanism to resume or share a negotiated secret key between the control and data connection for an FTPS session. Valid values: `DISABLED`, `ENABLED`, `ENFORCED`.
     * 
     */
    public Optional<Output<String>> tlsSessionResumptionMode() {
        return Optional.ofNullable(this.tlsSessionResumptionMode);
    }

    private ServerProtocolDetailsArgs() {}

    private ServerProtocolDetailsArgs(ServerProtocolDetailsArgs $) {
        this.as2Transports = $.as2Transports;
        this.passiveIp = $.passiveIp;
        this.setStatOption = $.setStatOption;
        this.tlsSessionResumptionMode = $.tlsSessionResumptionMode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServerProtocolDetailsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServerProtocolDetailsArgs $;

        public Builder() {
            $ = new ServerProtocolDetailsArgs();
        }

        public Builder(ServerProtocolDetailsArgs defaults) {
            $ = new ServerProtocolDetailsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param as2Transports Indicates the transport method for the AS2 messages. Currently, only `HTTP` is supported.
         * 
         * @return builder
         * 
         */
        public Builder as2Transports(@Nullable Output<List<String>> as2Transports) {
            $.as2Transports = as2Transports;
            return this;
        }

        /**
         * @param as2Transports Indicates the transport method for the AS2 messages. Currently, only `HTTP` is supported.
         * 
         * @return builder
         * 
         */
        public Builder as2Transports(List<String> as2Transports) {
            return as2Transports(Output.of(as2Transports));
        }

        /**
         * @param as2Transports Indicates the transport method for the AS2 messages. Currently, only `HTTP` is supported.
         * 
         * @return builder
         * 
         */
        public Builder as2Transports(String... as2Transports) {
            return as2Transports(List.of(as2Transports));
        }

        /**
         * @param passiveIp Indicates passive mode, for FTP and FTPS protocols. Enter a single IPv4 address, such as the public IP address of a firewall, router, or load balancer.
         * 
         * @return builder
         * 
         */
        public Builder passiveIp(@Nullable Output<String> passiveIp) {
            $.passiveIp = passiveIp;
            return this;
        }

        /**
         * @param passiveIp Indicates passive mode, for FTP and FTPS protocols. Enter a single IPv4 address, such as the public IP address of a firewall, router, or load balancer.
         * 
         * @return builder
         * 
         */
        public Builder passiveIp(String passiveIp) {
            return passiveIp(Output.of(passiveIp));
        }

        /**
         * @param setStatOption Use to ignore the error that is generated when the client attempts to use `SETSTAT` on a file you are uploading to an S3 bucket. Valid values: `DEFAULT`, `ENABLE_NO_OP`.
         * 
         * @return builder
         * 
         */
        public Builder setStatOption(@Nullable Output<String> setStatOption) {
            $.setStatOption = setStatOption;
            return this;
        }

        /**
         * @param setStatOption Use to ignore the error that is generated when the client attempts to use `SETSTAT` on a file you are uploading to an S3 bucket. Valid values: `DEFAULT`, `ENABLE_NO_OP`.
         * 
         * @return builder
         * 
         */
        public Builder setStatOption(String setStatOption) {
            return setStatOption(Output.of(setStatOption));
        }

        /**
         * @param tlsSessionResumptionMode A property used with Transfer Family servers that use the FTPS protocol. Provides a mechanism to resume or share a negotiated secret key between the control and data connection for an FTPS session. Valid values: `DISABLED`, `ENABLED`, `ENFORCED`.
         * 
         * @return builder
         * 
         */
        public Builder tlsSessionResumptionMode(@Nullable Output<String> tlsSessionResumptionMode) {
            $.tlsSessionResumptionMode = tlsSessionResumptionMode;
            return this;
        }

        /**
         * @param tlsSessionResumptionMode A property used with Transfer Family servers that use the FTPS protocol. Provides a mechanism to resume or share a negotiated secret key between the control and data connection for an FTPS session. Valid values: `DISABLED`, `ENABLED`, `ENFORCED`.
         * 
         * @return builder
         * 
         */
        public Builder tlsSessionResumptionMode(String tlsSessionResumptionMode) {
            return tlsSessionResumptionMode(Output.of(tlsSessionResumptionMode));
        }

        public ServerProtocolDetailsArgs build() {
            return $;
        }
    }

}
