// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class WorkspaceNetworkAccessControl {
    /**
     * @return An array of prefix list IDs.
     * 
     */
    private List<String> prefixListIds;
    /**
     * @return An array of Amazon VPC endpoint IDs for the workspace. The only VPC endpoints that can be specified here are interface VPC endpoints for Grafana workspaces (using the com.amazonaws.[region].grafana-workspace service endpoint). Other VPC endpoints will be ignored.
     * 
     */
    private List<String> vpceIds;

    private WorkspaceNetworkAccessControl() {}
    /**
     * @return An array of prefix list IDs.
     * 
     */
    public List<String> prefixListIds() {
        return this.prefixListIds;
    }
    /**
     * @return An array of Amazon VPC endpoint IDs for the workspace. The only VPC endpoints that can be specified here are interface VPC endpoints for Grafana workspaces (using the com.amazonaws.[region].grafana-workspace service endpoint). Other VPC endpoints will be ignored.
     * 
     */
    public List<String> vpceIds() {
        return this.vpceIds;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(WorkspaceNetworkAccessControl defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> prefixListIds;
        private List<String> vpceIds;
        public Builder() {}
        public Builder(WorkspaceNetworkAccessControl defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.prefixListIds = defaults.prefixListIds;
    	      this.vpceIds = defaults.vpceIds;
        }

        @CustomType.Setter
        public Builder prefixListIds(List<String> prefixListIds) {
            this.prefixListIds = Objects.requireNonNull(prefixListIds);
            return this;
        }
        public Builder prefixListIds(String... prefixListIds) {
            return prefixListIds(List.of(prefixListIds));
        }
        @CustomType.Setter
        public Builder vpceIds(List<String> vpceIds) {
            this.vpceIds = Objects.requireNonNull(vpceIds);
            return this;
        }
        public Builder vpceIds(String... vpceIds) {
            return vpceIds(List.of(vpceIds));
        }
        public WorkspaceNetworkAccessControl build() {
            final var o = new WorkspaceNetworkAccessControl();
            o.prefixListIds = prefixListIds;
            o.vpceIds = vpceIds;
            return o;
        }
    }
}
