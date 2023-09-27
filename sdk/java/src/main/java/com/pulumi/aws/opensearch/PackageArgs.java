// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.opensearch;

import com.pulumi.aws.opensearch.inputs.PackagePackageSourceArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PackageArgs extends com.pulumi.resources.ResourceArgs {

    public static final PackageArgs Empty = new PackageArgs();

    /**
     * Description of the package.
     * 
     */
    @Import(name="packageDescription")
    private @Nullable Output<String> packageDescription;

    /**
     * @return Description of the package.
     * 
     */
    public Optional<Output<String>> packageDescription() {
        return Optional.ofNullable(this.packageDescription);
    }

    /**
     * Unique name for the package.
     * 
     */
    @Import(name="packageName", required=true)
    private Output<String> packageName;

    /**
     * @return Unique name for the package.
     * 
     */
    public Output<String> packageName() {
        return this.packageName;
    }

    /**
     * Configuration block for the package source options.
     * 
     */
    @Import(name="packageSource", required=true)
    private Output<PackagePackageSourceArgs> packageSource;

    /**
     * @return Configuration block for the package source options.
     * 
     */
    public Output<PackagePackageSourceArgs> packageSource() {
        return this.packageSource;
    }

    /**
     * The type of package.
     * 
     */
    @Import(name="packageType", required=true)
    private Output<String> packageType;

    /**
     * @return The type of package.
     * 
     */
    public Output<String> packageType() {
        return this.packageType;
    }

    private PackageArgs() {}

    private PackageArgs(PackageArgs $) {
        this.packageDescription = $.packageDescription;
        this.packageName = $.packageName;
        this.packageSource = $.packageSource;
        this.packageType = $.packageType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PackageArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PackageArgs $;

        public Builder() {
            $ = new PackageArgs();
        }

        public Builder(PackageArgs defaults) {
            $ = new PackageArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param packageDescription Description of the package.
         * 
         * @return builder
         * 
         */
        public Builder packageDescription(@Nullable Output<String> packageDescription) {
            $.packageDescription = packageDescription;
            return this;
        }

        /**
         * @param packageDescription Description of the package.
         * 
         * @return builder
         * 
         */
        public Builder packageDescription(String packageDescription) {
            return packageDescription(Output.of(packageDescription));
        }

        /**
         * @param packageName Unique name for the package.
         * 
         * @return builder
         * 
         */
        public Builder packageName(Output<String> packageName) {
            $.packageName = packageName;
            return this;
        }

        /**
         * @param packageName Unique name for the package.
         * 
         * @return builder
         * 
         */
        public Builder packageName(String packageName) {
            return packageName(Output.of(packageName));
        }

        /**
         * @param packageSource Configuration block for the package source options.
         * 
         * @return builder
         * 
         */
        public Builder packageSource(Output<PackagePackageSourceArgs> packageSource) {
            $.packageSource = packageSource;
            return this;
        }

        /**
         * @param packageSource Configuration block for the package source options.
         * 
         * @return builder
         * 
         */
        public Builder packageSource(PackagePackageSourceArgs packageSource) {
            return packageSource(Output.of(packageSource));
        }

        /**
         * @param packageType The type of package.
         * 
         * @return builder
         * 
         */
        public Builder packageType(Output<String> packageType) {
            $.packageType = packageType;
            return this;
        }

        /**
         * @param packageType The type of package.
         * 
         * @return builder
         * 
         */
        public Builder packageType(String packageType) {
            return packageType(Output.of(packageType));
        }

        public PackageArgs build() {
            $.packageName = Objects.requireNonNull($.packageName, "expected parameter 'packageName' to be non-null");
            $.packageSource = Objects.requireNonNull($.packageSource, "expected parameter 'packageSource' to be non-null");
            $.packageType = Objects.requireNonNull($.packageType, "expected parameter 'packageType' to be non-null");
            return $;
        }
    }

}
