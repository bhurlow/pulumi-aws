// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.efs.outputs;

import com.pulumi.aws.efs.outputs.GetAccessPointRootDirectoryCreationInfo;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetAccessPointRootDirectory {
    /**
     * @return Single element list containing information on the creation permissions of the directory
     * 
     */
    private final List<GetAccessPointRootDirectoryCreationInfo> creationInfos;
    /**
     * @return Path exposed as the root directory
     * 
     */
    private final String path;

    @CustomType.Constructor
    private GetAccessPointRootDirectory(
        @CustomType.Parameter("creationInfos") List<GetAccessPointRootDirectoryCreationInfo> creationInfos,
        @CustomType.Parameter("path") String path) {
        this.creationInfos = creationInfos;
        this.path = path;
    }

    /**
     * @return Single element list containing information on the creation permissions of the directory
     * 
     */
    public List<GetAccessPointRootDirectoryCreationInfo> creationInfos() {
        return this.creationInfos;
    }
    /**
     * @return Path exposed as the root directory
     * 
     */
    public String path() {
        return this.path;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAccessPointRootDirectory defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private List<GetAccessPointRootDirectoryCreationInfo> creationInfos;
        private String path;

        public Builder() {
    	      // Empty
        }

        public Builder(GetAccessPointRootDirectory defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.creationInfos = defaults.creationInfos;
    	      this.path = defaults.path;
        }

        public Builder creationInfos(List<GetAccessPointRootDirectoryCreationInfo> creationInfos) {
            this.creationInfos = Objects.requireNonNull(creationInfos);
            return this;
        }
        public Builder creationInfos(GetAccessPointRootDirectoryCreationInfo... creationInfos) {
            return creationInfos(List.of(creationInfos));
        }
        public Builder path(String path) {
            this.path = Objects.requireNonNull(path);
            return this;
        }        public GetAccessPointRootDirectory build() {
            return new GetAccessPointRootDirectory(creationInfos, path);
        }
    }
}