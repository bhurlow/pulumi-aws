// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.opensearch.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PackageAssociationState extends com.pulumi.resources.ResourceArgs {

    public static final PackageAssociationState Empty = new PackageAssociationState();

    /**
     * Name of the domain to associate the package with.
     * 
     */
    @Import(name="domainName")
    private @Nullable Output<String> domainName;

    /**
     * @return Name of the domain to associate the package with.
     * 
     */
    public Optional<Output<String>> domainName() {
        return Optional.ofNullable(this.domainName);
    }

    /**
     * Internal ID of the package to associate with a domain.
     * 
     */
    @Import(name="packageId")
    private @Nullable Output<String> packageId;

    /**
     * @return Internal ID of the package to associate with a domain.
     * 
     */
    public Optional<Output<String>> packageId() {
        return Optional.ofNullable(this.packageId);
    }

    @Import(name="referencePath")
    private @Nullable Output<String> referencePath;

    public Optional<Output<String>> referencePath() {
        return Optional.ofNullable(this.referencePath);
    }

    private PackageAssociationState() {}

    private PackageAssociationState(PackageAssociationState $) {
        this.domainName = $.domainName;
        this.packageId = $.packageId;
        this.referencePath = $.referencePath;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PackageAssociationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PackageAssociationState $;

        public Builder() {
            $ = new PackageAssociationState();
        }

        public Builder(PackageAssociationState defaults) {
            $ = new PackageAssociationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param domainName Name of the domain to associate the package with.
         * 
         * @return builder
         * 
         */
        public Builder domainName(@Nullable Output<String> domainName) {
            $.domainName = domainName;
            return this;
        }

        /**
         * @param domainName Name of the domain to associate the package with.
         * 
         * @return builder
         * 
         */
        public Builder domainName(String domainName) {
            return domainName(Output.of(domainName));
        }

        /**
         * @param packageId Internal ID of the package to associate with a domain.
         * 
         * @return builder
         * 
         */
        public Builder packageId(@Nullable Output<String> packageId) {
            $.packageId = packageId;
            return this;
        }

        /**
         * @param packageId Internal ID of the package to associate with a domain.
         * 
         * @return builder
         * 
         */
        public Builder packageId(String packageId) {
            return packageId(Output.of(packageId));
        }

        public Builder referencePath(@Nullable Output<String> referencePath) {
            $.referencePath = referencePath;
            return this;
        }

        public Builder referencePath(String referencePath) {
            return referencePath(Output.of(referencePath));
        }

        public PackageAssociationState build() {
            return $;
        }
    }

}
