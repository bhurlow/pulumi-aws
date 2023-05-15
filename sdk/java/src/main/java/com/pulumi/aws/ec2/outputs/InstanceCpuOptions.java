// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstanceCpuOptions {
    /**
     * @return Indicates whether to enable the instance for AMD SEV-SNP. AMD SEV-SNP is supported with M6a, R6a, and C6a instance types only. Valid values are `enabled` and `disabled`.
     * 
     */
    private @Nullable String amdSevSnp;
    /**
     * @return Sets the number of CPU cores for an instance. This option is only supported on creation of instance type that support CPU Options [CPU Cores and Threads Per CPU Core Per Instance Type](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html#cpu-options-supported-instances-values) - specifying this option for unsupported instance types will return an error from the EC2 API.
     * 
     */
    private @Nullable Integer coreCount;
    /**
     * @return If set to 1, hyperthreading is disabled on the launched instance. Defaults to 2 if not set. See [Optimizing CPU Options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) for more information.
     * 
     */
    private @Nullable Integer threadsPerCore;

    private InstanceCpuOptions() {}
    /**
     * @return Indicates whether to enable the instance for AMD SEV-SNP. AMD SEV-SNP is supported with M6a, R6a, and C6a instance types only. Valid values are `enabled` and `disabled`.
     * 
     */
    public Optional<String> amdSevSnp() {
        return Optional.ofNullable(this.amdSevSnp);
    }
    /**
     * @return Sets the number of CPU cores for an instance. This option is only supported on creation of instance type that support CPU Options [CPU Cores and Threads Per CPU Core Per Instance Type](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html#cpu-options-supported-instances-values) - specifying this option for unsupported instance types will return an error from the EC2 API.
     * 
     */
    public Optional<Integer> coreCount() {
        return Optional.ofNullable(this.coreCount);
    }
    /**
     * @return If set to 1, hyperthreading is disabled on the launched instance. Defaults to 2 if not set. See [Optimizing CPU Options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) for more information.
     * 
     */
    public Optional<Integer> threadsPerCore() {
        return Optional.ofNullable(this.threadsPerCore);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstanceCpuOptions defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String amdSevSnp;
        private @Nullable Integer coreCount;
        private @Nullable Integer threadsPerCore;
        public Builder() {}
        public Builder(InstanceCpuOptions defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.amdSevSnp = defaults.amdSevSnp;
    	      this.coreCount = defaults.coreCount;
    	      this.threadsPerCore = defaults.threadsPerCore;
        }

        @CustomType.Setter
        public Builder amdSevSnp(@Nullable String amdSevSnp) {
            this.amdSevSnp = amdSevSnp;
            return this;
        }
        @CustomType.Setter
        public Builder coreCount(@Nullable Integer coreCount) {
            this.coreCount = coreCount;
            return this;
        }
        @CustomType.Setter
        public Builder threadsPerCore(@Nullable Integer threadsPerCore) {
            this.threadsPerCore = threadsPerCore;
            return this;
        }
        public InstanceCpuOptions build() {
            final var o = new InstanceCpuOptions();
            o.amdSevSnp = amdSevSnp;
            o.coreCount = coreCount;
            o.threadsPerCore = threadsPerCore;
            return o;
        }
    }
}
