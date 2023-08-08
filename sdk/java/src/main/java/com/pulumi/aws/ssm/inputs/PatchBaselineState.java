// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ssm.inputs;

import com.pulumi.aws.ssm.inputs.PatchBaselineApprovalRuleArgs;
import com.pulumi.aws.ssm.inputs.PatchBaselineGlobalFilterArgs;
import com.pulumi.aws.ssm.inputs.PatchBaselineSourceArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PatchBaselineState extends com.pulumi.resources.ResourceArgs {

    public static final PatchBaselineState Empty = new PatchBaselineState();

    /**
     * A set of rules used to include patches in the baseline.
     * Up to 10 approval rules can be specified.
     * See `approval_rule` below.
     * 
     */
    @Import(name="approvalRules")
    private @Nullable Output<List<PatchBaselineApprovalRuleArgs>> approvalRules;

    /**
     * @return A set of rules used to include patches in the baseline.
     * Up to 10 approval rules can be specified.
     * See `approval_rule` below.
     * 
     */
    public Optional<Output<List<PatchBaselineApprovalRuleArgs>>> approvalRules() {
        return Optional.ofNullable(this.approvalRules);
    }

    /**
     * A list of explicitly approved patches for the baseline.
     * Cannot be specified with `approval_rule`.
     * 
     */
    @Import(name="approvedPatches")
    private @Nullable Output<List<String>> approvedPatches;

    /**
     * @return A list of explicitly approved patches for the baseline.
     * Cannot be specified with `approval_rule`.
     * 
     */
    public Optional<Output<List<String>>> approvedPatches() {
        return Optional.ofNullable(this.approvedPatches);
    }

    /**
     * The compliance level for approved patches.
     * This means that if an approved patch is reported as missing, this is the severity of the compliance violation.
     * Valid values are `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`.
     * The default value is `UNSPECIFIED`.
     * 
     */
    @Import(name="approvedPatchesComplianceLevel")
    private @Nullable Output<String> approvedPatchesComplianceLevel;

    /**
     * @return The compliance level for approved patches.
     * This means that if an approved patch is reported as missing, this is the severity of the compliance violation.
     * Valid values are `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`.
     * The default value is `UNSPECIFIED`.
     * 
     */
    public Optional<Output<String>> approvedPatchesComplianceLevel() {
        return Optional.ofNullable(this.approvedPatchesComplianceLevel);
    }

    /**
     * Indicates whether the list of approved patches includes non-security updates that should be applied to the instances.
     * Applies to Linux instances only.
     * 
     */
    @Import(name="approvedPatchesEnableNonSecurity")
    private @Nullable Output<Boolean> approvedPatchesEnableNonSecurity;

    /**
     * @return Indicates whether the list of approved patches includes non-security updates that should be applied to the instances.
     * Applies to Linux instances only.
     * 
     */
    public Optional<Output<Boolean>> approvedPatchesEnableNonSecurity() {
        return Optional.ofNullable(this.approvedPatchesEnableNonSecurity);
    }

    /**
     * The ARN of the patch baseline.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The ARN of the patch baseline.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The description of the patch baseline.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the patch baseline.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * A set of global filters used to exclude patches from the baseline.
     * Up to 4 global filters can be specified using Key/Value pairs.
     * Valid Keys are `PRODUCT`, `CLASSIFICATION`, `MSRC_SEVERITY`, and `PATCH_ID`.
     * 
     */
    @Import(name="globalFilters")
    private @Nullable Output<List<PatchBaselineGlobalFilterArgs>> globalFilters;

    /**
     * @return A set of global filters used to exclude patches from the baseline.
     * Up to 4 global filters can be specified using Key/Value pairs.
     * Valid Keys are `PRODUCT`, `CLASSIFICATION`, `MSRC_SEVERITY`, and `PATCH_ID`.
     * 
     */
    public Optional<Output<List<PatchBaselineGlobalFilterArgs>>> globalFilters() {
        return Optional.ofNullable(this.globalFilters);
    }

    /**
     * The name of the patch baseline.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the patch baseline.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The operating system the patch baseline applies to.
     * Valid values are
     * `ALMA_LINUX`,
     * `AMAZON_LINUX`,
     * `AMAZON_LINUX_2`,
     * `AMAZON_LINUX_2022`,
     * `AMAZON_LINUX_2023`,
     * `CENTOS`,
     * `DEBIAN`,
     * `MACOS`,
     * `ORACLE_LINUX`,
     * `RASPBIAN`,
     * `REDHAT_ENTERPRISE_LINUX`,
     * `ROCKY_LINUX`,
     * `SUSE`,
     * `UBUNTU`, and
     * `WINDOWS`.
     * The default value is `WINDOWS`.
     * 
     */
    @Import(name="operatingSystem")
    private @Nullable Output<String> operatingSystem;

    /**
     * @return The operating system the patch baseline applies to.
     * Valid values are
     * `ALMA_LINUX`,
     * `AMAZON_LINUX`,
     * `AMAZON_LINUX_2`,
     * `AMAZON_LINUX_2022`,
     * `AMAZON_LINUX_2023`,
     * `CENTOS`,
     * `DEBIAN`,
     * `MACOS`,
     * `ORACLE_LINUX`,
     * `RASPBIAN`,
     * `REDHAT_ENTERPRISE_LINUX`,
     * `ROCKY_LINUX`,
     * `SUSE`,
     * `UBUNTU`, and
     * `WINDOWS`.
     * The default value is `WINDOWS`.
     * 
     */
    public Optional<Output<String>> operatingSystem() {
        return Optional.ofNullable(this.operatingSystem);
    }

    /**
     * A list of rejected patches.
     * 
     */
    @Import(name="rejectedPatches")
    private @Nullable Output<List<String>> rejectedPatches;

    /**
     * @return A list of rejected patches.
     * 
     */
    public Optional<Output<List<String>>> rejectedPatches() {
        return Optional.ofNullable(this.rejectedPatches);
    }

    /**
     * The action for Patch Manager to take on patches included in the `rejected_patches` list.
     * Valid values are `ALLOW_AS_DEPENDENCY` and `BLOCK`.
     * 
     */
    @Import(name="rejectedPatchesAction")
    private @Nullable Output<String> rejectedPatchesAction;

    /**
     * @return The action for Patch Manager to take on patches included in the `rejected_patches` list.
     * Valid values are `ALLOW_AS_DEPENDENCY` and `BLOCK`.
     * 
     */
    public Optional<Output<String>> rejectedPatchesAction() {
        return Optional.ofNullable(this.rejectedPatchesAction);
    }

    /**
     * Configuration block with alternate sources for patches.
     * Applies to Linux instances only.
     * See `source` below.
     * 
     */
    @Import(name="sources")
    private @Nullable Output<List<PatchBaselineSourceArgs>> sources;

    /**
     * @return Configuration block with alternate sources for patches.
     * Applies to Linux instances only.
     * See `source` below.
     * 
     */
    public Optional<Output<List<PatchBaselineSourceArgs>>> sources() {
        return Optional.ofNullable(this.sources);
    }

    /**
     * A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    private PatchBaselineState() {}

    private PatchBaselineState(PatchBaselineState $) {
        this.approvalRules = $.approvalRules;
        this.approvedPatches = $.approvedPatches;
        this.approvedPatchesComplianceLevel = $.approvedPatchesComplianceLevel;
        this.approvedPatchesEnableNonSecurity = $.approvedPatchesEnableNonSecurity;
        this.arn = $.arn;
        this.description = $.description;
        this.globalFilters = $.globalFilters;
        this.name = $.name;
        this.operatingSystem = $.operatingSystem;
        this.rejectedPatches = $.rejectedPatches;
        this.rejectedPatchesAction = $.rejectedPatchesAction;
        this.sources = $.sources;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PatchBaselineState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PatchBaselineState $;

        public Builder() {
            $ = new PatchBaselineState();
        }

        public Builder(PatchBaselineState defaults) {
            $ = new PatchBaselineState(Objects.requireNonNull(defaults));
        }

        /**
         * @param approvalRules A set of rules used to include patches in the baseline.
         * Up to 10 approval rules can be specified.
         * See `approval_rule` below.
         * 
         * @return builder
         * 
         */
        public Builder approvalRules(@Nullable Output<List<PatchBaselineApprovalRuleArgs>> approvalRules) {
            $.approvalRules = approvalRules;
            return this;
        }

        /**
         * @param approvalRules A set of rules used to include patches in the baseline.
         * Up to 10 approval rules can be specified.
         * See `approval_rule` below.
         * 
         * @return builder
         * 
         */
        public Builder approvalRules(List<PatchBaselineApprovalRuleArgs> approvalRules) {
            return approvalRules(Output.of(approvalRules));
        }

        /**
         * @param approvalRules A set of rules used to include patches in the baseline.
         * Up to 10 approval rules can be specified.
         * See `approval_rule` below.
         * 
         * @return builder
         * 
         */
        public Builder approvalRules(PatchBaselineApprovalRuleArgs... approvalRules) {
            return approvalRules(List.of(approvalRules));
        }

        /**
         * @param approvedPatches A list of explicitly approved patches for the baseline.
         * Cannot be specified with `approval_rule`.
         * 
         * @return builder
         * 
         */
        public Builder approvedPatches(@Nullable Output<List<String>> approvedPatches) {
            $.approvedPatches = approvedPatches;
            return this;
        }

        /**
         * @param approvedPatches A list of explicitly approved patches for the baseline.
         * Cannot be specified with `approval_rule`.
         * 
         * @return builder
         * 
         */
        public Builder approvedPatches(List<String> approvedPatches) {
            return approvedPatches(Output.of(approvedPatches));
        }

        /**
         * @param approvedPatches A list of explicitly approved patches for the baseline.
         * Cannot be specified with `approval_rule`.
         * 
         * @return builder
         * 
         */
        public Builder approvedPatches(String... approvedPatches) {
            return approvedPatches(List.of(approvedPatches));
        }

        /**
         * @param approvedPatchesComplianceLevel The compliance level for approved patches.
         * This means that if an approved patch is reported as missing, this is the severity of the compliance violation.
         * Valid values are `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`.
         * The default value is `UNSPECIFIED`.
         * 
         * @return builder
         * 
         */
        public Builder approvedPatchesComplianceLevel(@Nullable Output<String> approvedPatchesComplianceLevel) {
            $.approvedPatchesComplianceLevel = approvedPatchesComplianceLevel;
            return this;
        }

        /**
         * @param approvedPatchesComplianceLevel The compliance level for approved patches.
         * This means that if an approved patch is reported as missing, this is the severity of the compliance violation.
         * Valid values are `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`.
         * The default value is `UNSPECIFIED`.
         * 
         * @return builder
         * 
         */
        public Builder approvedPatchesComplianceLevel(String approvedPatchesComplianceLevel) {
            return approvedPatchesComplianceLevel(Output.of(approvedPatchesComplianceLevel));
        }

        /**
         * @param approvedPatchesEnableNonSecurity Indicates whether the list of approved patches includes non-security updates that should be applied to the instances.
         * Applies to Linux instances only.
         * 
         * @return builder
         * 
         */
        public Builder approvedPatchesEnableNonSecurity(@Nullable Output<Boolean> approvedPatchesEnableNonSecurity) {
            $.approvedPatchesEnableNonSecurity = approvedPatchesEnableNonSecurity;
            return this;
        }

        /**
         * @param approvedPatchesEnableNonSecurity Indicates whether the list of approved patches includes non-security updates that should be applied to the instances.
         * Applies to Linux instances only.
         * 
         * @return builder
         * 
         */
        public Builder approvedPatchesEnableNonSecurity(Boolean approvedPatchesEnableNonSecurity) {
            return approvedPatchesEnableNonSecurity(Output.of(approvedPatchesEnableNonSecurity));
        }

        /**
         * @param arn The ARN of the patch baseline.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The ARN of the patch baseline.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param description The description of the patch baseline.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the patch baseline.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param globalFilters A set of global filters used to exclude patches from the baseline.
         * Up to 4 global filters can be specified using Key/Value pairs.
         * Valid Keys are `PRODUCT`, `CLASSIFICATION`, `MSRC_SEVERITY`, and `PATCH_ID`.
         * 
         * @return builder
         * 
         */
        public Builder globalFilters(@Nullable Output<List<PatchBaselineGlobalFilterArgs>> globalFilters) {
            $.globalFilters = globalFilters;
            return this;
        }

        /**
         * @param globalFilters A set of global filters used to exclude patches from the baseline.
         * Up to 4 global filters can be specified using Key/Value pairs.
         * Valid Keys are `PRODUCT`, `CLASSIFICATION`, `MSRC_SEVERITY`, and `PATCH_ID`.
         * 
         * @return builder
         * 
         */
        public Builder globalFilters(List<PatchBaselineGlobalFilterArgs> globalFilters) {
            return globalFilters(Output.of(globalFilters));
        }

        /**
         * @param globalFilters A set of global filters used to exclude patches from the baseline.
         * Up to 4 global filters can be specified using Key/Value pairs.
         * Valid Keys are `PRODUCT`, `CLASSIFICATION`, `MSRC_SEVERITY`, and `PATCH_ID`.
         * 
         * @return builder
         * 
         */
        public Builder globalFilters(PatchBaselineGlobalFilterArgs... globalFilters) {
            return globalFilters(List.of(globalFilters));
        }

        /**
         * @param name The name of the patch baseline.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the patch baseline.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param operatingSystem The operating system the patch baseline applies to.
         * Valid values are
         * `ALMA_LINUX`,
         * `AMAZON_LINUX`,
         * `AMAZON_LINUX_2`,
         * `AMAZON_LINUX_2022`,
         * `AMAZON_LINUX_2023`,
         * `CENTOS`,
         * `DEBIAN`,
         * `MACOS`,
         * `ORACLE_LINUX`,
         * `RASPBIAN`,
         * `REDHAT_ENTERPRISE_LINUX`,
         * `ROCKY_LINUX`,
         * `SUSE`,
         * `UBUNTU`, and
         * `WINDOWS`.
         * The default value is `WINDOWS`.
         * 
         * @return builder
         * 
         */
        public Builder operatingSystem(@Nullable Output<String> operatingSystem) {
            $.operatingSystem = operatingSystem;
            return this;
        }

        /**
         * @param operatingSystem The operating system the patch baseline applies to.
         * Valid values are
         * `ALMA_LINUX`,
         * `AMAZON_LINUX`,
         * `AMAZON_LINUX_2`,
         * `AMAZON_LINUX_2022`,
         * `AMAZON_LINUX_2023`,
         * `CENTOS`,
         * `DEBIAN`,
         * `MACOS`,
         * `ORACLE_LINUX`,
         * `RASPBIAN`,
         * `REDHAT_ENTERPRISE_LINUX`,
         * `ROCKY_LINUX`,
         * `SUSE`,
         * `UBUNTU`, and
         * `WINDOWS`.
         * The default value is `WINDOWS`.
         * 
         * @return builder
         * 
         */
        public Builder operatingSystem(String operatingSystem) {
            return operatingSystem(Output.of(operatingSystem));
        }

        /**
         * @param rejectedPatches A list of rejected patches.
         * 
         * @return builder
         * 
         */
        public Builder rejectedPatches(@Nullable Output<List<String>> rejectedPatches) {
            $.rejectedPatches = rejectedPatches;
            return this;
        }

        /**
         * @param rejectedPatches A list of rejected patches.
         * 
         * @return builder
         * 
         */
        public Builder rejectedPatches(List<String> rejectedPatches) {
            return rejectedPatches(Output.of(rejectedPatches));
        }

        /**
         * @param rejectedPatches A list of rejected patches.
         * 
         * @return builder
         * 
         */
        public Builder rejectedPatches(String... rejectedPatches) {
            return rejectedPatches(List.of(rejectedPatches));
        }

        /**
         * @param rejectedPatchesAction The action for Patch Manager to take on patches included in the `rejected_patches` list.
         * Valid values are `ALLOW_AS_DEPENDENCY` and `BLOCK`.
         * 
         * @return builder
         * 
         */
        public Builder rejectedPatchesAction(@Nullable Output<String> rejectedPatchesAction) {
            $.rejectedPatchesAction = rejectedPatchesAction;
            return this;
        }

        /**
         * @param rejectedPatchesAction The action for Patch Manager to take on patches included in the `rejected_patches` list.
         * Valid values are `ALLOW_AS_DEPENDENCY` and `BLOCK`.
         * 
         * @return builder
         * 
         */
        public Builder rejectedPatchesAction(String rejectedPatchesAction) {
            return rejectedPatchesAction(Output.of(rejectedPatchesAction));
        }

        /**
         * @param sources Configuration block with alternate sources for patches.
         * Applies to Linux instances only.
         * See `source` below.
         * 
         * @return builder
         * 
         */
        public Builder sources(@Nullable Output<List<PatchBaselineSourceArgs>> sources) {
            $.sources = sources;
            return this;
        }

        /**
         * @param sources Configuration block with alternate sources for patches.
         * Applies to Linux instances only.
         * See `source` below.
         * 
         * @return builder
         * 
         */
        public Builder sources(List<PatchBaselineSourceArgs> sources) {
            return sources(Output.of(sources));
        }

        /**
         * @param sources Configuration block with alternate sources for patches.
         * Applies to Linux instances only.
         * See `source` below.
         * 
         * @return builder
         * 
         */
        public Builder sources(PatchBaselineSourceArgs... sources) {
            return sources(List.of(sources));
        }

        /**
         * @param tags A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        public PatchBaselineState build() {
            return $;
        }
    }

}
