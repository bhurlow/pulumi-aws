// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.efs.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FileSystemSizeInByteArgs extends com.pulumi.resources.ResourceArgs {

    public static final FileSystemSizeInByteArgs Empty = new FileSystemSizeInByteArgs();

    /**
     * The latest known metered size (in bytes) of data stored in the file system.
     * 
     */
    @Import(name="value")
    private @Nullable Output<Integer> value;

    /**
     * @return The latest known metered size (in bytes) of data stored in the file system.
     * 
     */
    public Optional<Output<Integer>> value() {
        return Optional.ofNullable(this.value);
    }

    /**
     * The latest known metered size (in bytes) of data stored in the Infrequent Access storage class.
     * 
     */
    @Import(name="valueInIa")
    private @Nullable Output<Integer> valueInIa;

    /**
     * @return The latest known metered size (in bytes) of data stored in the Infrequent Access storage class.
     * 
     */
    public Optional<Output<Integer>> valueInIa() {
        return Optional.ofNullable(this.valueInIa);
    }

    /**
     * The latest known metered size (in bytes) of data stored in the Standard storage class.
     * 
     */
    @Import(name="valueInStandard")
    private @Nullable Output<Integer> valueInStandard;

    /**
     * @return The latest known metered size (in bytes) of data stored in the Standard storage class.
     * 
     */
    public Optional<Output<Integer>> valueInStandard() {
        return Optional.ofNullable(this.valueInStandard);
    }

    private FileSystemSizeInByteArgs() {}

    private FileSystemSizeInByteArgs(FileSystemSizeInByteArgs $) {
        this.value = $.value;
        this.valueInIa = $.valueInIa;
        this.valueInStandard = $.valueInStandard;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FileSystemSizeInByteArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FileSystemSizeInByteArgs $;

        public Builder() {
            $ = new FileSystemSizeInByteArgs();
        }

        public Builder(FileSystemSizeInByteArgs defaults) {
            $ = new FileSystemSizeInByteArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param value The latest known metered size (in bytes) of data stored in the file system.
         * 
         * @return builder
         * 
         */
        public Builder value(@Nullable Output<Integer> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value The latest known metered size (in bytes) of data stored in the file system.
         * 
         * @return builder
         * 
         */
        public Builder value(Integer value) {
            return value(Output.of(value));
        }

        /**
         * @param valueInIa The latest known metered size (in bytes) of data stored in the Infrequent Access storage class.
         * 
         * @return builder
         * 
         */
        public Builder valueInIa(@Nullable Output<Integer> valueInIa) {
            $.valueInIa = valueInIa;
            return this;
        }

        /**
         * @param valueInIa The latest known metered size (in bytes) of data stored in the Infrequent Access storage class.
         * 
         * @return builder
         * 
         */
        public Builder valueInIa(Integer valueInIa) {
            return valueInIa(Output.of(valueInIa));
        }

        /**
         * @param valueInStandard The latest known metered size (in bytes) of data stored in the Standard storage class.
         * 
         * @return builder
         * 
         */
        public Builder valueInStandard(@Nullable Output<Integer> valueInStandard) {
            $.valueInStandard = valueInStandard;
            return this;
        }

        /**
         * @param valueInStandard The latest known metered size (in bytes) of data stored in the Standard storage class.
         * 
         * @return builder
         * 
         */
        public Builder valueInStandard(Integer valueInStandard) {
            return valueInStandard(Output.of(valueInStandard));
        }

        public FileSystemSizeInByteArgs build() {
            return $;
        }
    }

}