// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.transfer;

import com.pulumi.aws.transfer.inputs.AccessHomeDirectoryMappingArgs;
import com.pulumi.aws.transfer.inputs.AccessPosixProfileArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AccessArgs extends com.pulumi.resources.ResourceArgs {

    public static final AccessArgs Empty = new AccessArgs();

    /**
     * The SID of a group in the directory connected to the Transfer Server (e.g., `S-1-1-12-1234567890-123456789-1234567890-1234`)
     * 
     */
    @Import(name="externalId", required=true)
    private Output<String> externalId;

    /**
     * @return The SID of a group in the directory connected to the Transfer Server (e.g., `S-1-1-12-1234567890-123456789-1234567890-1234`)
     * 
     */
    public Output<String> externalId() {
        return this.externalId;
    }

    /**
     * The landing directory (folder) for a user when they log in to the server using their SFTP client.  It should begin with a `/`.  The first item in the path is the name of the home bucket (accessible as `${Transfer:HomeBucket}` in the policy) and the rest is the home directory (accessible as `${Transfer:HomeDirectory}` in the policy). For example, `/example-bucket-1234/username` would set the home bucket to `example-bucket-1234` and the home directory to `username`.
     * 
     */
    @Import(name="homeDirectory")
    private @Nullable Output<String> homeDirectory;

    /**
     * @return The landing directory (folder) for a user when they log in to the server using their SFTP client.  It should begin with a `/`.  The first item in the path is the name of the home bucket (accessible as `${Transfer:HomeBucket}` in the policy) and the rest is the home directory (accessible as `${Transfer:HomeDirectory}` in the policy). For example, `/example-bucket-1234/username` would set the home bucket to `example-bucket-1234` and the home directory to `username`.
     * 
     */
    public Optional<Output<String>> homeDirectory() {
        return Optional.ofNullable(this.homeDirectory);
    }

    /**
     * Logical directory mappings that specify what S3 paths and keys should be visible to your user and how you want to make them visible. See Home Directory Mappings below.
     * 
     */
    @Import(name="homeDirectoryMappings")
    private @Nullable Output<List<AccessHomeDirectoryMappingArgs>> homeDirectoryMappings;

    /**
     * @return Logical directory mappings that specify what S3 paths and keys should be visible to your user and how you want to make them visible. See Home Directory Mappings below.
     * 
     */
    public Optional<Output<List<AccessHomeDirectoryMappingArgs>>> homeDirectoryMappings() {
        return Optional.ofNullable(this.homeDirectoryMappings);
    }

    /**
     * The type of landing directory (folder) you mapped for your users&#39; home directory. Valid values are `PATH` and `LOGICAL`.
     * 
     */
    @Import(name="homeDirectoryType")
    private @Nullable Output<String> homeDirectoryType;

    /**
     * @return The type of landing directory (folder) you mapped for your users&#39; home directory. Valid values are `PATH` and `LOGICAL`.
     * 
     */
    public Optional<Output<String>> homeDirectoryType() {
        return Optional.ofNullable(this.homeDirectoryType);
    }

    /**
     * An IAM JSON policy document that scopes down user access to portions of their Amazon S3 bucket. IAM variables you can use inside this policy include `${Transfer:UserName}`, `${Transfer:HomeDirectory}`, and `${Transfer:HomeBucket}`. These are evaluated on-the-fly when navigating the bucket.
     * 
     */
    @Import(name="policy")
    private @Nullable Output<String> policy;

    /**
     * @return An IAM JSON policy document that scopes down user access to portions of their Amazon S3 bucket. IAM variables you can use inside this policy include `${Transfer:UserName}`, `${Transfer:HomeDirectory}`, and `${Transfer:HomeBucket}`. These are evaluated on-the-fly when navigating the bucket.
     * 
     */
    public Optional<Output<String>> policy() {
        return Optional.ofNullable(this.policy);
    }

    /**
     * Specifies the full POSIX identity, including user ID (Uid), group ID (Gid), and any secondary groups IDs (SecondaryGids), that controls your users&#39; access to your Amazon EFS file systems. See Posix Profile below.
     * 
     */
    @Import(name="posixProfile")
    private @Nullable Output<AccessPosixProfileArgs> posixProfile;

    /**
     * @return Specifies the full POSIX identity, including user ID (Uid), group ID (Gid), and any secondary groups IDs (SecondaryGids), that controls your users&#39; access to your Amazon EFS file systems. See Posix Profile below.
     * 
     */
    public Optional<Output<AccessPosixProfileArgs>> posixProfile() {
        return Optional.ofNullable(this.posixProfile);
    }

    /**
     * Amazon Resource Name (ARN) of an IAM role that allows the service to controls your user’s access to your Amazon S3 bucket.
     * 
     */
    @Import(name="role")
    private @Nullable Output<String> role;

    /**
     * @return Amazon Resource Name (ARN) of an IAM role that allows the service to controls your user’s access to your Amazon S3 bucket.
     * 
     */
    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    /**
     * The Server ID of the Transfer Server (e.g., `s-12345678`)
     * 
     */
    @Import(name="serverId", required=true)
    private Output<String> serverId;

    /**
     * @return The Server ID of the Transfer Server (e.g., `s-12345678`)
     * 
     */
    public Output<String> serverId() {
        return this.serverId;
    }

    private AccessArgs() {}

    private AccessArgs(AccessArgs $) {
        this.externalId = $.externalId;
        this.homeDirectory = $.homeDirectory;
        this.homeDirectoryMappings = $.homeDirectoryMappings;
        this.homeDirectoryType = $.homeDirectoryType;
        this.policy = $.policy;
        this.posixProfile = $.posixProfile;
        this.role = $.role;
        this.serverId = $.serverId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccessArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccessArgs $;

        public Builder() {
            $ = new AccessArgs();
        }

        public Builder(AccessArgs defaults) {
            $ = new AccessArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param externalId The SID of a group in the directory connected to the Transfer Server (e.g., `S-1-1-12-1234567890-123456789-1234567890-1234`)
         * 
         * @return builder
         * 
         */
        public Builder externalId(Output<String> externalId) {
            $.externalId = externalId;
            return this;
        }

        /**
         * @param externalId The SID of a group in the directory connected to the Transfer Server (e.g., `S-1-1-12-1234567890-123456789-1234567890-1234`)
         * 
         * @return builder
         * 
         */
        public Builder externalId(String externalId) {
            return externalId(Output.of(externalId));
        }

        /**
         * @param homeDirectory The landing directory (folder) for a user when they log in to the server using their SFTP client.  It should begin with a `/`.  The first item in the path is the name of the home bucket (accessible as `${Transfer:HomeBucket}` in the policy) and the rest is the home directory (accessible as `${Transfer:HomeDirectory}` in the policy). For example, `/example-bucket-1234/username` would set the home bucket to `example-bucket-1234` and the home directory to `username`.
         * 
         * @return builder
         * 
         */
        public Builder homeDirectory(@Nullable Output<String> homeDirectory) {
            $.homeDirectory = homeDirectory;
            return this;
        }

        /**
         * @param homeDirectory The landing directory (folder) for a user when they log in to the server using their SFTP client.  It should begin with a `/`.  The first item in the path is the name of the home bucket (accessible as `${Transfer:HomeBucket}` in the policy) and the rest is the home directory (accessible as `${Transfer:HomeDirectory}` in the policy). For example, `/example-bucket-1234/username` would set the home bucket to `example-bucket-1234` and the home directory to `username`.
         * 
         * @return builder
         * 
         */
        public Builder homeDirectory(String homeDirectory) {
            return homeDirectory(Output.of(homeDirectory));
        }

        /**
         * @param homeDirectoryMappings Logical directory mappings that specify what S3 paths and keys should be visible to your user and how you want to make them visible. See Home Directory Mappings below.
         * 
         * @return builder
         * 
         */
        public Builder homeDirectoryMappings(@Nullable Output<List<AccessHomeDirectoryMappingArgs>> homeDirectoryMappings) {
            $.homeDirectoryMappings = homeDirectoryMappings;
            return this;
        }

        /**
         * @param homeDirectoryMappings Logical directory mappings that specify what S3 paths and keys should be visible to your user and how you want to make them visible. See Home Directory Mappings below.
         * 
         * @return builder
         * 
         */
        public Builder homeDirectoryMappings(List<AccessHomeDirectoryMappingArgs> homeDirectoryMappings) {
            return homeDirectoryMappings(Output.of(homeDirectoryMappings));
        }

        /**
         * @param homeDirectoryMappings Logical directory mappings that specify what S3 paths and keys should be visible to your user and how you want to make them visible. See Home Directory Mappings below.
         * 
         * @return builder
         * 
         */
        public Builder homeDirectoryMappings(AccessHomeDirectoryMappingArgs... homeDirectoryMappings) {
            return homeDirectoryMappings(List.of(homeDirectoryMappings));
        }

        /**
         * @param homeDirectoryType The type of landing directory (folder) you mapped for your users&#39; home directory. Valid values are `PATH` and `LOGICAL`.
         * 
         * @return builder
         * 
         */
        public Builder homeDirectoryType(@Nullable Output<String> homeDirectoryType) {
            $.homeDirectoryType = homeDirectoryType;
            return this;
        }

        /**
         * @param homeDirectoryType The type of landing directory (folder) you mapped for your users&#39; home directory. Valid values are `PATH` and `LOGICAL`.
         * 
         * @return builder
         * 
         */
        public Builder homeDirectoryType(String homeDirectoryType) {
            return homeDirectoryType(Output.of(homeDirectoryType));
        }

        /**
         * @param policy An IAM JSON policy document that scopes down user access to portions of their Amazon S3 bucket. IAM variables you can use inside this policy include `${Transfer:UserName}`, `${Transfer:HomeDirectory}`, and `${Transfer:HomeBucket}`. These are evaluated on-the-fly when navigating the bucket.
         * 
         * @return builder
         * 
         */
        public Builder policy(@Nullable Output<String> policy) {
            $.policy = policy;
            return this;
        }

        /**
         * @param policy An IAM JSON policy document that scopes down user access to portions of their Amazon S3 bucket. IAM variables you can use inside this policy include `${Transfer:UserName}`, `${Transfer:HomeDirectory}`, and `${Transfer:HomeBucket}`. These are evaluated on-the-fly when navigating the bucket.
         * 
         * @return builder
         * 
         */
        public Builder policy(String policy) {
            return policy(Output.of(policy));
        }

        /**
         * @param posixProfile Specifies the full POSIX identity, including user ID (Uid), group ID (Gid), and any secondary groups IDs (SecondaryGids), that controls your users&#39; access to your Amazon EFS file systems. See Posix Profile below.
         * 
         * @return builder
         * 
         */
        public Builder posixProfile(@Nullable Output<AccessPosixProfileArgs> posixProfile) {
            $.posixProfile = posixProfile;
            return this;
        }

        /**
         * @param posixProfile Specifies the full POSIX identity, including user ID (Uid), group ID (Gid), and any secondary groups IDs (SecondaryGids), that controls your users&#39; access to your Amazon EFS file systems. See Posix Profile below.
         * 
         * @return builder
         * 
         */
        public Builder posixProfile(AccessPosixProfileArgs posixProfile) {
            return posixProfile(Output.of(posixProfile));
        }

        /**
         * @param role Amazon Resource Name (ARN) of an IAM role that allows the service to controls your user’s access to your Amazon S3 bucket.
         * 
         * @return builder
         * 
         */
        public Builder role(@Nullable Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role Amazon Resource Name (ARN) of an IAM role that allows the service to controls your user’s access to your Amazon S3 bucket.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        /**
         * @param serverId The Server ID of the Transfer Server (e.g., `s-12345678`)
         * 
         * @return builder
         * 
         */
        public Builder serverId(Output<String> serverId) {
            $.serverId = serverId;
            return this;
        }

        /**
         * @param serverId The Server ID of the Transfer Server (e.g., `s-12345678`)
         * 
         * @return builder
         * 
         */
        public Builder serverId(String serverId) {
            return serverId(Output.of(serverId));
        }

        public AccessArgs build() {
            $.externalId = Objects.requireNonNull($.externalId, "expected parameter 'externalId' to be non-null");
            $.serverId = Objects.requireNonNull($.serverId, "expected parameter 'serverId' to be non-null");
            return $;
        }
    }

}
