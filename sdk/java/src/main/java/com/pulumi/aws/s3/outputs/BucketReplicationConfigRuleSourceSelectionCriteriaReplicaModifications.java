// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.s3.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class BucketReplicationConfigRuleSourceSelectionCriteriaReplicaModifications {
    /**
     * @return Whether the existing objects should be replicated. Either `&#34;Enabled&#34;` or `&#34;Disabled&#34;`.
     * 
     */
    private final String status;

    @CustomType.Constructor
    private BucketReplicationConfigRuleSourceSelectionCriteriaReplicaModifications(@CustomType.Parameter("status") String status) {
        this.status = status;
    }

    /**
     * @return Whether the existing objects should be replicated. Either `&#34;Enabled&#34;` or `&#34;Disabled&#34;`.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BucketReplicationConfigRuleSourceSelectionCriteriaReplicaModifications defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String status;

        public Builder() {
    	      // Empty
        }

        public Builder(BucketReplicationConfigRuleSourceSelectionCriteriaReplicaModifications defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.status = defaults.status;
        }

        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }        public BucketReplicationConfigRuleSourceSelectionCriteriaReplicaModifications build() {
            return new BucketReplicationConfigRuleSourceSelectionCriteriaReplicaModifications(status);
        }
    }
}