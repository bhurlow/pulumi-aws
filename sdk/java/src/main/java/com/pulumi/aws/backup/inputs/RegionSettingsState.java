// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.backup.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RegionSettingsState extends com.pulumi.resources.ResourceArgs {

    public static final RegionSettingsState Empty = new RegionSettingsState();

    /**
     * A map of services along with the management preferences for the Region. For more information, see the [AWS Documentation](https://docs.aws.amazon.com/aws-backup/latest/devguide/API_UpdateRegionSettings.html#API_UpdateRegionSettings_RequestSyntax).
     * 
     */
    @Import(name="resourceTypeManagementPreference")
    private @Nullable Output<Map<String,Boolean>> resourceTypeManagementPreference;

    /**
     * @return A map of services along with the management preferences for the Region. For more information, see the [AWS Documentation](https://docs.aws.amazon.com/aws-backup/latest/devguide/API_UpdateRegionSettings.html#API_UpdateRegionSettings_RequestSyntax).
     * 
     */
    public Optional<Output<Map<String,Boolean>>> resourceTypeManagementPreference() {
        return Optional.ofNullable(this.resourceTypeManagementPreference);
    }

    /**
     * A map of services along with the opt-in preferences for the Region.
     * 
     */
    @Import(name="resourceTypeOptInPreference")
    private @Nullable Output<Map<String,Boolean>> resourceTypeOptInPreference;

    /**
     * @return A map of services along with the opt-in preferences for the Region.
     * 
     */
    public Optional<Output<Map<String,Boolean>>> resourceTypeOptInPreference() {
        return Optional.ofNullable(this.resourceTypeOptInPreference);
    }

    private RegionSettingsState() {}

    private RegionSettingsState(RegionSettingsState $) {
        this.resourceTypeManagementPreference = $.resourceTypeManagementPreference;
        this.resourceTypeOptInPreference = $.resourceTypeOptInPreference;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RegionSettingsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RegionSettingsState $;

        public Builder() {
            $ = new RegionSettingsState();
        }

        public Builder(RegionSettingsState defaults) {
            $ = new RegionSettingsState(Objects.requireNonNull(defaults));
        }

        /**
         * @param resourceTypeManagementPreference A map of services along with the management preferences for the Region. For more information, see the [AWS Documentation](https://docs.aws.amazon.com/aws-backup/latest/devguide/API_UpdateRegionSettings.html#API_UpdateRegionSettings_RequestSyntax).
         * 
         * @return builder
         * 
         */
        public Builder resourceTypeManagementPreference(@Nullable Output<Map<String,Boolean>> resourceTypeManagementPreference) {
            $.resourceTypeManagementPreference = resourceTypeManagementPreference;
            return this;
        }

        /**
         * @param resourceTypeManagementPreference A map of services along with the management preferences for the Region. For more information, see the [AWS Documentation](https://docs.aws.amazon.com/aws-backup/latest/devguide/API_UpdateRegionSettings.html#API_UpdateRegionSettings_RequestSyntax).
         * 
         * @return builder
         * 
         */
        public Builder resourceTypeManagementPreference(Map<String,Boolean> resourceTypeManagementPreference) {
            return resourceTypeManagementPreference(Output.of(resourceTypeManagementPreference));
        }

        /**
         * @param resourceTypeOptInPreference A map of services along with the opt-in preferences for the Region.
         * 
         * @return builder
         * 
         */
        public Builder resourceTypeOptInPreference(@Nullable Output<Map<String,Boolean>> resourceTypeOptInPreference) {
            $.resourceTypeOptInPreference = resourceTypeOptInPreference;
            return this;
        }

        /**
         * @param resourceTypeOptInPreference A map of services along with the opt-in preferences for the Region.
         * 
         * @return builder
         * 
         */
        public Builder resourceTypeOptInPreference(Map<String,Boolean> resourceTypeOptInPreference) {
            return resourceTypeOptInPreference(Output.of(resourceTypeOptInPreference));
        }

        public RegionSettingsState build() {
            return $;
        }
    }

}
