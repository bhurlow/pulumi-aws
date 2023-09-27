// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.fsx.outputs;

import com.pulumi.aws.fsx.outputs.GetOntapStorageVirtualMachineActiveDirectoryConfiguration;
import com.pulumi.aws.fsx.outputs.GetOntapStorageVirtualMachineEndpoint;
import com.pulumi.aws.fsx.outputs.GetOntapStorageVirtualMachineFilter;
import com.pulumi.aws.fsx.outputs.GetOntapStorageVirtualMachineLifecycleTransitionReason;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetOntapStorageVirtualMachineResult {
    /**
     * @return The Microsoft Active Directory configuration to which the SVM is joined, if applicable. See Active Directory Configuration below.
     * 
     */
    private List<GetOntapStorageVirtualMachineActiveDirectoryConfiguration> activeDirectoryConfigurations;
    /**
     * @return Amazon Resource Name of the SVM.
     * 
     */
    private String arn;
    /**
     * @return The time that the SVM was created.
     * 
     */
    private String creationTime;
    /**
     * @return The endpoints that are used to access data or to manage the SVM using the NetApp ONTAP CLI, REST API, or NetApp CloudManager. They are the Iscsi, Management, Nfs, and Smb endpoints. See SVM Endpoints below.
     * 
     */
    private List<GetOntapStorageVirtualMachineEndpoint> endpoints;
    /**
     * @return Identifier of the file system (e.g. `fs-12345678`).
     * 
     */
    private String fileSystemId;
    private @Nullable List<GetOntapStorageVirtualMachineFilter> filters;
    /**
     * @return The SVM&#39;s system generated unique ID.
     * 
     */
    private String id;
    /**
     * @return The SVM&#39;s lifecycle status.
     * 
     */
    private String lifecycleStatus;
    /**
     * @return Describes why the SVM lifecycle state changed. See Lifecycle Transition Reason below.
     * 
     */
    private List<GetOntapStorageVirtualMachineLifecycleTransitionReason> lifecycleTransitionReasons;
    /**
     * @return The name of the SVM, if provisioned.
     * 
     */
    private String name;
    /**
     * @return The SVM&#39;s subtype.
     * 
     */
    private String subtype;
    private Map<String,String> tags;
    /**
     * @return The SVM&#39;s UUID.
     * 
     */
    private String uuid;

    private GetOntapStorageVirtualMachineResult() {}
    /**
     * @return The Microsoft Active Directory configuration to which the SVM is joined, if applicable. See Active Directory Configuration below.
     * 
     */
    public List<GetOntapStorageVirtualMachineActiveDirectoryConfiguration> activeDirectoryConfigurations() {
        return this.activeDirectoryConfigurations;
    }
    /**
     * @return Amazon Resource Name of the SVM.
     * 
     */
    public String arn() {
        return this.arn;
    }
    /**
     * @return The time that the SVM was created.
     * 
     */
    public String creationTime() {
        return this.creationTime;
    }
    /**
     * @return The endpoints that are used to access data or to manage the SVM using the NetApp ONTAP CLI, REST API, or NetApp CloudManager. They are the Iscsi, Management, Nfs, and Smb endpoints. See SVM Endpoints below.
     * 
     */
    public List<GetOntapStorageVirtualMachineEndpoint> endpoints() {
        return this.endpoints;
    }
    /**
     * @return Identifier of the file system (e.g. `fs-12345678`).
     * 
     */
    public String fileSystemId() {
        return this.fileSystemId;
    }
    public List<GetOntapStorageVirtualMachineFilter> filters() {
        return this.filters == null ? List.of() : this.filters;
    }
    /**
     * @return The SVM&#39;s system generated unique ID.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The SVM&#39;s lifecycle status.
     * 
     */
    public String lifecycleStatus() {
        return this.lifecycleStatus;
    }
    /**
     * @return Describes why the SVM lifecycle state changed. See Lifecycle Transition Reason below.
     * 
     */
    public List<GetOntapStorageVirtualMachineLifecycleTransitionReason> lifecycleTransitionReasons() {
        return this.lifecycleTransitionReasons;
    }
    /**
     * @return The name of the SVM, if provisioned.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The SVM&#39;s subtype.
     * 
     */
    public String subtype() {
        return this.subtype;
    }
    public Map<String,String> tags() {
        return this.tags;
    }
    /**
     * @return The SVM&#39;s UUID.
     * 
     */
    public String uuid() {
        return this.uuid;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOntapStorageVirtualMachineResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetOntapStorageVirtualMachineActiveDirectoryConfiguration> activeDirectoryConfigurations;
        private String arn;
        private String creationTime;
        private List<GetOntapStorageVirtualMachineEndpoint> endpoints;
        private String fileSystemId;
        private @Nullable List<GetOntapStorageVirtualMachineFilter> filters;
        private String id;
        private String lifecycleStatus;
        private List<GetOntapStorageVirtualMachineLifecycleTransitionReason> lifecycleTransitionReasons;
        private String name;
        private String subtype;
        private Map<String,String> tags;
        private String uuid;
        public Builder() {}
        public Builder(GetOntapStorageVirtualMachineResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.activeDirectoryConfigurations = defaults.activeDirectoryConfigurations;
    	      this.arn = defaults.arn;
    	      this.creationTime = defaults.creationTime;
    	      this.endpoints = defaults.endpoints;
    	      this.fileSystemId = defaults.fileSystemId;
    	      this.filters = defaults.filters;
    	      this.id = defaults.id;
    	      this.lifecycleStatus = defaults.lifecycleStatus;
    	      this.lifecycleTransitionReasons = defaults.lifecycleTransitionReasons;
    	      this.name = defaults.name;
    	      this.subtype = defaults.subtype;
    	      this.tags = defaults.tags;
    	      this.uuid = defaults.uuid;
        }

        @CustomType.Setter
        public Builder activeDirectoryConfigurations(List<GetOntapStorageVirtualMachineActiveDirectoryConfiguration> activeDirectoryConfigurations) {
            this.activeDirectoryConfigurations = Objects.requireNonNull(activeDirectoryConfigurations);
            return this;
        }
        public Builder activeDirectoryConfigurations(GetOntapStorageVirtualMachineActiveDirectoryConfiguration... activeDirectoryConfigurations) {
            return activeDirectoryConfigurations(List.of(activeDirectoryConfigurations));
        }
        @CustomType.Setter
        public Builder arn(String arn) {
            this.arn = Objects.requireNonNull(arn);
            return this;
        }
        @CustomType.Setter
        public Builder creationTime(String creationTime) {
            this.creationTime = Objects.requireNonNull(creationTime);
            return this;
        }
        @CustomType.Setter
        public Builder endpoints(List<GetOntapStorageVirtualMachineEndpoint> endpoints) {
            this.endpoints = Objects.requireNonNull(endpoints);
            return this;
        }
        public Builder endpoints(GetOntapStorageVirtualMachineEndpoint... endpoints) {
            return endpoints(List.of(endpoints));
        }
        @CustomType.Setter
        public Builder fileSystemId(String fileSystemId) {
            this.fileSystemId = Objects.requireNonNull(fileSystemId);
            return this;
        }
        @CustomType.Setter
        public Builder filters(@Nullable List<GetOntapStorageVirtualMachineFilter> filters) {
            this.filters = filters;
            return this;
        }
        public Builder filters(GetOntapStorageVirtualMachineFilter... filters) {
            return filters(List.of(filters));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder lifecycleStatus(String lifecycleStatus) {
            this.lifecycleStatus = Objects.requireNonNull(lifecycleStatus);
            return this;
        }
        @CustomType.Setter
        public Builder lifecycleTransitionReasons(List<GetOntapStorageVirtualMachineLifecycleTransitionReason> lifecycleTransitionReasons) {
            this.lifecycleTransitionReasons = Objects.requireNonNull(lifecycleTransitionReasons);
            return this;
        }
        public Builder lifecycleTransitionReasons(GetOntapStorageVirtualMachineLifecycleTransitionReason... lifecycleTransitionReasons) {
            return lifecycleTransitionReasons(List.of(lifecycleTransitionReasons));
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder subtype(String subtype) {
            this.subtype = Objects.requireNonNull(subtype);
            return this;
        }
        @CustomType.Setter
        public Builder tags(Map<String,String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        @CustomType.Setter
        public Builder uuid(String uuid) {
            this.uuid = Objects.requireNonNull(uuid);
            return this;
        }
        public GetOntapStorageVirtualMachineResult build() {
            final var o = new GetOntapStorageVirtualMachineResult();
            o.activeDirectoryConfigurations = activeDirectoryConfigurations;
            o.arn = arn;
            o.creationTime = creationTime;
            o.endpoints = endpoints;
            o.fileSystemId = fileSystemId;
            o.filters = filters;
            o.id = id;
            o.lifecycleStatus = lifecycleStatus;
            o.lifecycleTransitionReasons = lifecycleTransitionReasons;
            o.name = name;
            o.subtype = subtype;
            o.tags = tags;
            o.uuid = uuid;
            return o;
        }
    }
}
