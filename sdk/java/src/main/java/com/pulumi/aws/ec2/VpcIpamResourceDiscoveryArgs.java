// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.ec2.inputs.VpcIpamResourceDiscoveryOperatingRegionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VpcIpamResourceDiscoveryArgs extends com.pulumi.resources.ResourceArgs {

    public static final VpcIpamResourceDiscoveryArgs Empty = new VpcIpamResourceDiscoveryArgs();

    /**
     * A description for the IPAM Resource Discovery.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description for the IPAM Resource Discovery.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Determines which regions the Resource Discovery will enable IPAM features for usage and monitoring. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM Resource Discovery. You can only create VPCs from a pool whose locale matches the VPC&#39;s Region. You specify a region using the region_name parameter. **You must set your provider block region as an operating_region.**
     * 
     */
    @Import(name="operatingRegions", required=true)
    private Output<List<VpcIpamResourceDiscoveryOperatingRegionArgs>> operatingRegions;

    /**
     * @return Determines which regions the Resource Discovery will enable IPAM features for usage and monitoring. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM Resource Discovery. You can only create VPCs from a pool whose locale matches the VPC&#39;s Region. You specify a region using the region_name parameter. **You must set your provider block region as an operating_region.**
     * 
     */
    public Output<List<VpcIpamResourceDiscoveryOperatingRegionArgs>> operatingRegions() {
        return this.operatingRegions;
    }

    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private VpcIpamResourceDiscoveryArgs() {}

    private VpcIpamResourceDiscoveryArgs(VpcIpamResourceDiscoveryArgs $) {
        this.description = $.description;
        this.operatingRegions = $.operatingRegions;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VpcIpamResourceDiscoveryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VpcIpamResourceDiscoveryArgs $;

        public Builder() {
            $ = new VpcIpamResourceDiscoveryArgs();
        }

        public Builder(VpcIpamResourceDiscoveryArgs defaults) {
            $ = new VpcIpamResourceDiscoveryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description A description for the IPAM Resource Discovery.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description for the IPAM Resource Discovery.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param operatingRegions Determines which regions the Resource Discovery will enable IPAM features for usage and monitoring. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM Resource Discovery. You can only create VPCs from a pool whose locale matches the VPC&#39;s Region. You specify a region using the region_name parameter. **You must set your provider block region as an operating_region.**
         * 
         * @return builder
         * 
         */
        public Builder operatingRegions(Output<List<VpcIpamResourceDiscoveryOperatingRegionArgs>> operatingRegions) {
            $.operatingRegions = operatingRegions;
            return this;
        }

        /**
         * @param operatingRegions Determines which regions the Resource Discovery will enable IPAM features for usage and monitoring. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM Resource Discovery. You can only create VPCs from a pool whose locale matches the VPC&#39;s Region. You specify a region using the region_name parameter. **You must set your provider block region as an operating_region.**
         * 
         * @return builder
         * 
         */
        public Builder operatingRegions(List<VpcIpamResourceDiscoveryOperatingRegionArgs> operatingRegions) {
            return operatingRegions(Output.of(operatingRegions));
        }

        /**
         * @param operatingRegions Determines which regions the Resource Discovery will enable IPAM features for usage and monitoring. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM Resource Discovery. You can only create VPCs from a pool whose locale matches the VPC&#39;s Region. You specify a region using the region_name parameter. **You must set your provider block region as an operating_region.**
         * 
         * @return builder
         * 
         */
        public Builder operatingRegions(VpcIpamResourceDiscoveryOperatingRegionArgs... operatingRegions) {
            return operatingRegions(List.of(operatingRegions));
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public VpcIpamResourceDiscoveryArgs build() {
            $.operatingRegions = Objects.requireNonNull($.operatingRegions, "expected parameter 'operatingRegions' to be non-null");
            return $;
        }
    }

}
