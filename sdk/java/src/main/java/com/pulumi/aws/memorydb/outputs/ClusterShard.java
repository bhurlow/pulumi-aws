// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.memorydb.outputs;

import com.pulumi.aws.memorydb.outputs.ClusterShardNode;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ClusterShard {
    /**
     * @return Name of the cluster. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
     * 
     */
    private @Nullable String name;
    /**
     * @return Set of nodes in this shard.
     * 
     */
    private @Nullable List<ClusterShardNode> nodes;
    /**
     * @return Number of individual nodes in this shard.
     * 
     */
    private @Nullable Integer numNodes;
    /**
     * @return Keyspace for this shard. Example: `0-16383`.
     * 
     */
    private @Nullable String slots;

    private ClusterShard() {}
    /**
     * @return Name of the cluster. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return Set of nodes in this shard.
     * 
     */
    public List<ClusterShardNode> nodes() {
        return this.nodes == null ? List.of() : this.nodes;
    }
    /**
     * @return Number of individual nodes in this shard.
     * 
     */
    public Optional<Integer> numNodes() {
        return Optional.ofNullable(this.numNodes);
    }
    /**
     * @return Keyspace for this shard. Example: `0-16383`.
     * 
     */
    public Optional<String> slots() {
        return Optional.ofNullable(this.slots);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClusterShard defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String name;
        private @Nullable List<ClusterShardNode> nodes;
        private @Nullable Integer numNodes;
        private @Nullable String slots;
        public Builder() {}
        public Builder(ClusterShard defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.nodes = defaults.nodes;
    	      this.numNodes = defaults.numNodes;
    	      this.slots = defaults.slots;
        }

        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder nodes(@Nullable List<ClusterShardNode> nodes) {
            this.nodes = nodes;
            return this;
        }
        public Builder nodes(ClusterShardNode... nodes) {
            return nodes(List.of(nodes));
        }
        @CustomType.Setter
        public Builder numNodes(@Nullable Integer numNodes) {
            this.numNodes = numNodes;
            return this;
        }
        @CustomType.Setter
        public Builder slots(@Nullable String slots) {
            this.slots = slots;
            return this;
        }
        public ClusterShard build() {
            final var o = new ClusterShard();
            o.name = name;
            o.nodes = nodes;
            o.numNodes = numNodes;
            o.slots = slots;
            return o;
        }
    }
}
