// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.glue.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class JobCommand {
    /**
     * @return The name of the job command. Defaults to `glueetl`. Use `pythonshell` for Python Shell Job Type, `glueray` for Ray Job Type, or `gluestreaming` for Streaming Job Type. `max_capacity` needs to be set if `pythonshell` is chosen.
     * 
     */
    private @Nullable String name;
    /**
     * @return The Python version being used to execute a Python shell job. Allowed values are 2, 3 or 3.9. Version 3 refers to Python 3.6.
     * 
     */
    private @Nullable String pythonVersion;
    /**
     * @return In Ray jobs, runtime is used to specify the versions of Ray, Python and additional libraries available in your environment. This field is not used in other job types. For supported runtime environment values, see [Working with Ray jobs](https://docs.aws.amazon.com/glue/latest/dg/ray-jobs-section.html#author-job-ray-runtimes) in the Glue Developer Guide.
     * 
     */
    private @Nullable String runtime;
    /**
     * @return Specifies the S3 path to a script that executes a job.
     * 
     */
    private String scriptLocation;

    private JobCommand() {}
    /**
     * @return The name of the job command. Defaults to `glueetl`. Use `pythonshell` for Python Shell Job Type, `glueray` for Ray Job Type, or `gluestreaming` for Streaming Job Type. `max_capacity` needs to be set if `pythonshell` is chosen.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The Python version being used to execute a Python shell job. Allowed values are 2, 3 or 3.9. Version 3 refers to Python 3.6.
     * 
     */
    public Optional<String> pythonVersion() {
        return Optional.ofNullable(this.pythonVersion);
    }
    /**
     * @return In Ray jobs, runtime is used to specify the versions of Ray, Python and additional libraries available in your environment. This field is not used in other job types. For supported runtime environment values, see [Working with Ray jobs](https://docs.aws.amazon.com/glue/latest/dg/ray-jobs-section.html#author-job-ray-runtimes) in the Glue Developer Guide.
     * 
     */
    public Optional<String> runtime() {
        return Optional.ofNullable(this.runtime);
    }
    /**
     * @return Specifies the S3 path to a script that executes a job.
     * 
     */
    public String scriptLocation() {
        return this.scriptLocation;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(JobCommand defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String name;
        private @Nullable String pythonVersion;
        private @Nullable String runtime;
        private String scriptLocation;
        public Builder() {}
        public Builder(JobCommand defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.pythonVersion = defaults.pythonVersion;
    	      this.runtime = defaults.runtime;
    	      this.scriptLocation = defaults.scriptLocation;
        }

        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder pythonVersion(@Nullable String pythonVersion) {
            this.pythonVersion = pythonVersion;
            return this;
        }
        @CustomType.Setter
        public Builder runtime(@Nullable String runtime) {
            this.runtime = runtime;
            return this;
        }
        @CustomType.Setter
        public Builder scriptLocation(String scriptLocation) {
            this.scriptLocation = Objects.requireNonNull(scriptLocation);
            return this;
        }
        public JobCommand build() {
            final var o = new JobCommand();
            o.name = name;
            o.pythonVersion = pythonVersion;
            o.runtime = runtime;
            o.scriptLocation = scriptLocation;
            return o;
        }
    }
}
