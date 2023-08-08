// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ssm
{
    /// <summary>
    /// Provides an SSM Patch Baseline resource.
    /// 
    /// &gt; **NOTE on Patch Baselines:** The `approved_patches` and `approval_rule` are
    /// both marked as optional fields, but the Patch Baseline requires that at least one
    /// of them is specified.
    /// 
    /// ## Example Usage
    /// ### Basic Usage
    /// 
    /// Using `approved_patches` only.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var production = new Aws.Ssm.PatchBaseline("production", new()
    ///     {
    ///         ApprovedPatches = new[]
    ///         {
    ///             "KB123456",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Advanced Usage, specifying patch filters
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var production = new Aws.Ssm.PatchBaseline("production", new()
    ///     {
    ///         ApprovalRules = new[]
    ///         {
    ///             new Aws.Ssm.Inputs.PatchBaselineApprovalRuleArgs
    ///             {
    ///                 ApproveAfterDays = 7,
    ///                 ComplianceLevel = "HIGH",
    ///                 PatchFilters = new[]
    ///                 {
    ///                     new Aws.Ssm.Inputs.PatchBaselineApprovalRulePatchFilterArgs
    ///                     {
    ///                         Key = "PRODUCT",
    ///                         Values = new[]
    ///                         {
    ///                             "WindowsServer2016",
    ///                         },
    ///                     },
    ///                     new Aws.Ssm.Inputs.PatchBaselineApprovalRulePatchFilterArgs
    ///                     {
    ///                         Key = "CLASSIFICATION",
    ///                         Values = new[]
    ///                         {
    ///                             "CriticalUpdates",
    ///                             "SecurityUpdates",
    ///                             "Updates",
    ///                         },
    ///                     },
    ///                     new Aws.Ssm.Inputs.PatchBaselineApprovalRulePatchFilterArgs
    ///                     {
    ///                         Key = "MSRC_SEVERITY",
    ///                         Values = new[]
    ///                         {
    ///                             "Critical",
    ///                             "Important",
    ///                             "Moderate",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             new Aws.Ssm.Inputs.PatchBaselineApprovalRuleArgs
    ///             {
    ///                 ApproveAfterDays = 7,
    ///                 PatchFilters = new[]
    ///                 {
    ///                     new Aws.Ssm.Inputs.PatchBaselineApprovalRulePatchFilterArgs
    ///                     {
    ///                         Key = "PRODUCT",
    ///                         Values = new[]
    ///                         {
    ///                             "WindowsServer2012",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         ApprovedPatches = new[]
    ///         {
    ///             "KB123456",
    ///             "KB456789",
    ///         },
    ///         Description = "Patch Baseline Description",
    ///         GlobalFilters = new[]
    ///         {
    ///             new Aws.Ssm.Inputs.PatchBaselineGlobalFilterArgs
    ///             {
    ///                 Key = "PRODUCT",
    ///                 Values = new[]
    ///                 {
    ///                     "WindowsServer2008",
    ///                 },
    ///             },
    ///             new Aws.Ssm.Inputs.PatchBaselineGlobalFilterArgs
    ///             {
    ///                 Key = "CLASSIFICATION",
    ///                 Values = new[]
    ///                 {
    ///                     "ServicePacks",
    ///                 },
    ///             },
    ///             new Aws.Ssm.Inputs.PatchBaselineGlobalFilterArgs
    ///             {
    ///                 Key = "MSRC_SEVERITY",
    ///                 Values = new[]
    ///                 {
    ///                     "Low",
    ///                 },
    ///             },
    ///         },
    ///         RejectedPatches = new[]
    ///         {
    ///             "KB987654",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Advanced usage, specifying Microsoft application and Windows patch rules
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var windowsOsApps = new Aws.Ssm.PatchBaseline("windowsOsApps", new()
    ///     {
    ///         ApprovalRules = new[]
    ///         {
    ///             new Aws.Ssm.Inputs.PatchBaselineApprovalRuleArgs
    ///             {
    ///                 ApproveAfterDays = 7,
    ///                 PatchFilters = new[]
    ///                 {
    ///                     new Aws.Ssm.Inputs.PatchBaselineApprovalRulePatchFilterArgs
    ///                     {
    ///                         Key = "CLASSIFICATION",
    ///                         Values = new[]
    ///                         {
    ///                             "CriticalUpdates",
    ///                             "SecurityUpdates",
    ///                         },
    ///                     },
    ///                     new Aws.Ssm.Inputs.PatchBaselineApprovalRulePatchFilterArgs
    ///                     {
    ///                         Key = "MSRC_SEVERITY",
    ///                         Values = new[]
    ///                         {
    ///                             "Critical",
    ///                             "Important",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             new Aws.Ssm.Inputs.PatchBaselineApprovalRuleArgs
    ///             {
    ///                 ApproveAfterDays = 7,
    ///                 PatchFilters = new[]
    ///                 {
    ///                     new Aws.Ssm.Inputs.PatchBaselineApprovalRulePatchFilterArgs
    ///                     {
    ///                         Key = "PATCH_SET",
    ///                         Values = new[]
    ///                         {
    ///                             "APPLICATION",
    ///                         },
    ///                     },
    ///                     new Aws.Ssm.Inputs.PatchBaselineApprovalRulePatchFilterArgs
    ///                     {
    ///                         Key = "PRODUCT",
    ///                         Values = new[]
    ///                         {
    ///                             "Office 2013",
    ///                             "Office 2016",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         Description = "Patch both Windows and Microsoft apps",
    ///         OperatingSystem = "WINDOWS",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Advanced usage, specifying alternate patch source repository
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var al201709 = new Aws.Ssm.PatchBaseline("al201709", new()
    ///     {
    ///         ApprovalRules = new[]
    ///         {
    ///             null,
    ///         },
    ///         Description = "My patch repository for Amazon Linux 2017.09",
    ///         OperatingSystem = "AMAZON_LINUX",
    ///         Sources = new[]
    ///         {
    ///             new Aws.Ssm.Inputs.PatchBaselineSourceArgs
    ///             {
    ///                 Configuration = @"[amzn-main]
    /// name=amzn-main-Base
    /// mirrorlist=http://repo./$awsregion./$awsdomain//$releasever/main/mirror.list
    /// mirrorlist_expire=300
    /// metadata_expire=300
    /// priority=10
    /// failovermethod=priority
    /// fastestmirror_enabled=0
    /// gpgcheck=1
    /// gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-amazon-ga
    /// enabled=1
    /// retries=3
    /// timeout=5
    /// report_instanceid=yes
    /// 
    /// ",
    ///                 Name = "My-AL2017.09",
    ///                 Products = new[]
    ///                 {
    ///                     "AmazonLinux2017.09",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_ssm_patch_baseline.example
    /// 
    ///  id = "pb-12345678" } Using `pulumi import`, import SSM Patch Baselines using their baseline ID. For exampleconsole % pulumi import aws_ssm_patch_baseline.example pb-12345678
    /// </summary>
    [AwsResourceType("aws:ssm/patchBaseline:PatchBaseline")]
    public partial class PatchBaseline : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A set of rules used to include patches in the baseline.
        /// Up to 10 approval rules can be specified.
        /// See `approval_rule` below.
        /// </summary>
        [Output("approvalRules")]
        public Output<ImmutableArray<Outputs.PatchBaselineApprovalRule>> ApprovalRules { get; private set; } = null!;

        /// <summary>
        /// A list of explicitly approved patches for the baseline.
        /// Cannot be specified with `approval_rule`.
        /// </summary>
        [Output("approvedPatches")]
        public Output<ImmutableArray<string>> ApprovedPatches { get; private set; } = null!;

        /// <summary>
        /// The compliance level for approved patches.
        /// This means that if an approved patch is reported as missing, this is the severity of the compliance violation.
        /// Valid values are `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`.
        /// The default value is `UNSPECIFIED`.
        /// </summary>
        [Output("approvedPatchesComplianceLevel")]
        public Output<string?> ApprovedPatchesComplianceLevel { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the list of approved patches includes non-security updates that should be applied to the instances.
        /// Applies to Linux instances only.
        /// </summary>
        [Output("approvedPatchesEnableNonSecurity")]
        public Output<bool?> ApprovedPatchesEnableNonSecurity { get; private set; } = null!;

        /// <summary>
        /// The ARN of the patch baseline.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The description of the patch baseline.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A set of global filters used to exclude patches from the baseline.
        /// Up to 4 global filters can be specified using Key/Value pairs.
        /// Valid Keys are `PRODUCT`, `CLASSIFICATION`, `MSRC_SEVERITY`, and `PATCH_ID`.
        /// </summary>
        [Output("globalFilters")]
        public Output<ImmutableArray<Outputs.PatchBaselineGlobalFilter>> GlobalFilters { get; private set; } = null!;

        /// <summary>
        /// The name of the patch baseline.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The operating system the patch baseline applies to.
        /// Valid values are
        /// `ALMA_LINUX`,
        /// `AMAZON_LINUX`,
        /// `AMAZON_LINUX_2`,
        /// `AMAZON_LINUX_2022`,
        /// `AMAZON_LINUX_2023`,
        /// `CENTOS`,
        /// `DEBIAN`,
        /// `MACOS`,
        /// `ORACLE_LINUX`,
        /// `RASPBIAN`,
        /// `REDHAT_ENTERPRISE_LINUX`,
        /// `ROCKY_LINUX`,
        /// `SUSE`,
        /// `UBUNTU`, and
        /// `WINDOWS`.
        /// The default value is `WINDOWS`.
        /// </summary>
        [Output("operatingSystem")]
        public Output<string?> OperatingSystem { get; private set; } = null!;

        /// <summary>
        /// A list of rejected patches.
        /// </summary>
        [Output("rejectedPatches")]
        public Output<ImmutableArray<string>> RejectedPatches { get; private set; } = null!;

        /// <summary>
        /// The action for Patch Manager to take on patches included in the `rejected_patches` list.
        /// Valid values are `ALLOW_AS_DEPENDENCY` and `BLOCK`.
        /// </summary>
        [Output("rejectedPatchesAction")]
        public Output<string> RejectedPatchesAction { get; private set; } = null!;

        /// <summary>
        /// Configuration block with alternate sources for patches.
        /// Applies to Linux instances only.
        /// See `source` below.
        /// </summary>
        [Output("sources")]
        public Output<ImmutableArray<Outputs.PatchBaselineSource>> Sources { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a PatchBaseline resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PatchBaseline(string name, PatchBaselineArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:ssm/patchBaseline:PatchBaseline", name, args ?? new PatchBaselineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PatchBaseline(string name, Input<string> id, PatchBaselineState? state = null, CustomResourceOptions? options = null)
            : base("aws:ssm/patchBaseline:PatchBaseline", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PatchBaseline resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PatchBaseline Get(string name, Input<string> id, PatchBaselineState? state = null, CustomResourceOptions? options = null)
        {
            return new PatchBaseline(name, id, state, options);
        }
    }

    public sealed class PatchBaselineArgs : global::Pulumi.ResourceArgs
    {
        [Input("approvalRules")]
        private InputList<Inputs.PatchBaselineApprovalRuleArgs>? _approvalRules;

        /// <summary>
        /// A set of rules used to include patches in the baseline.
        /// Up to 10 approval rules can be specified.
        /// See `approval_rule` below.
        /// </summary>
        public InputList<Inputs.PatchBaselineApprovalRuleArgs> ApprovalRules
        {
            get => _approvalRules ?? (_approvalRules = new InputList<Inputs.PatchBaselineApprovalRuleArgs>());
            set => _approvalRules = value;
        }

        [Input("approvedPatches")]
        private InputList<string>? _approvedPatches;

        /// <summary>
        /// A list of explicitly approved patches for the baseline.
        /// Cannot be specified with `approval_rule`.
        /// </summary>
        public InputList<string> ApprovedPatches
        {
            get => _approvedPatches ?? (_approvedPatches = new InputList<string>());
            set => _approvedPatches = value;
        }

        /// <summary>
        /// The compliance level for approved patches.
        /// This means that if an approved patch is reported as missing, this is the severity of the compliance violation.
        /// Valid values are `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`.
        /// The default value is `UNSPECIFIED`.
        /// </summary>
        [Input("approvedPatchesComplianceLevel")]
        public Input<string>? ApprovedPatchesComplianceLevel { get; set; }

        /// <summary>
        /// Indicates whether the list of approved patches includes non-security updates that should be applied to the instances.
        /// Applies to Linux instances only.
        /// </summary>
        [Input("approvedPatchesEnableNonSecurity")]
        public Input<bool>? ApprovedPatchesEnableNonSecurity { get; set; }

        /// <summary>
        /// The description of the patch baseline.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("globalFilters")]
        private InputList<Inputs.PatchBaselineGlobalFilterArgs>? _globalFilters;

        /// <summary>
        /// A set of global filters used to exclude patches from the baseline.
        /// Up to 4 global filters can be specified using Key/Value pairs.
        /// Valid Keys are `PRODUCT`, `CLASSIFICATION`, `MSRC_SEVERITY`, and `PATCH_ID`.
        /// </summary>
        public InputList<Inputs.PatchBaselineGlobalFilterArgs> GlobalFilters
        {
            get => _globalFilters ?? (_globalFilters = new InputList<Inputs.PatchBaselineGlobalFilterArgs>());
            set => _globalFilters = value;
        }

        /// <summary>
        /// The name of the patch baseline.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The operating system the patch baseline applies to.
        /// Valid values are
        /// `ALMA_LINUX`,
        /// `AMAZON_LINUX`,
        /// `AMAZON_LINUX_2`,
        /// `AMAZON_LINUX_2022`,
        /// `AMAZON_LINUX_2023`,
        /// `CENTOS`,
        /// `DEBIAN`,
        /// `MACOS`,
        /// `ORACLE_LINUX`,
        /// `RASPBIAN`,
        /// `REDHAT_ENTERPRISE_LINUX`,
        /// `ROCKY_LINUX`,
        /// `SUSE`,
        /// `UBUNTU`, and
        /// `WINDOWS`.
        /// The default value is `WINDOWS`.
        /// </summary>
        [Input("operatingSystem")]
        public Input<string>? OperatingSystem { get; set; }

        [Input("rejectedPatches")]
        private InputList<string>? _rejectedPatches;

        /// <summary>
        /// A list of rejected patches.
        /// </summary>
        public InputList<string> RejectedPatches
        {
            get => _rejectedPatches ?? (_rejectedPatches = new InputList<string>());
            set => _rejectedPatches = value;
        }

        /// <summary>
        /// The action for Patch Manager to take on patches included in the `rejected_patches` list.
        /// Valid values are `ALLOW_AS_DEPENDENCY` and `BLOCK`.
        /// </summary>
        [Input("rejectedPatchesAction")]
        public Input<string>? RejectedPatchesAction { get; set; }

        [Input("sources")]
        private InputList<Inputs.PatchBaselineSourceArgs>? _sources;

        /// <summary>
        /// Configuration block with alternate sources for patches.
        /// Applies to Linux instances only.
        /// See `source` below.
        /// </summary>
        public InputList<Inputs.PatchBaselineSourceArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.PatchBaselineSourceArgs>());
            set => _sources = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public PatchBaselineArgs()
        {
        }
        public static new PatchBaselineArgs Empty => new PatchBaselineArgs();
    }

    public sealed class PatchBaselineState : global::Pulumi.ResourceArgs
    {
        [Input("approvalRules")]
        private InputList<Inputs.PatchBaselineApprovalRuleGetArgs>? _approvalRules;

        /// <summary>
        /// A set of rules used to include patches in the baseline.
        /// Up to 10 approval rules can be specified.
        /// See `approval_rule` below.
        /// </summary>
        public InputList<Inputs.PatchBaselineApprovalRuleGetArgs> ApprovalRules
        {
            get => _approvalRules ?? (_approvalRules = new InputList<Inputs.PatchBaselineApprovalRuleGetArgs>());
            set => _approvalRules = value;
        }

        [Input("approvedPatches")]
        private InputList<string>? _approvedPatches;

        /// <summary>
        /// A list of explicitly approved patches for the baseline.
        /// Cannot be specified with `approval_rule`.
        /// </summary>
        public InputList<string> ApprovedPatches
        {
            get => _approvedPatches ?? (_approvedPatches = new InputList<string>());
            set => _approvedPatches = value;
        }

        /// <summary>
        /// The compliance level for approved patches.
        /// This means that if an approved patch is reported as missing, this is the severity of the compliance violation.
        /// Valid values are `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`.
        /// The default value is `UNSPECIFIED`.
        /// </summary>
        [Input("approvedPatchesComplianceLevel")]
        public Input<string>? ApprovedPatchesComplianceLevel { get; set; }

        /// <summary>
        /// Indicates whether the list of approved patches includes non-security updates that should be applied to the instances.
        /// Applies to Linux instances only.
        /// </summary>
        [Input("approvedPatchesEnableNonSecurity")]
        public Input<bool>? ApprovedPatchesEnableNonSecurity { get; set; }

        /// <summary>
        /// The ARN of the patch baseline.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The description of the patch baseline.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("globalFilters")]
        private InputList<Inputs.PatchBaselineGlobalFilterGetArgs>? _globalFilters;

        /// <summary>
        /// A set of global filters used to exclude patches from the baseline.
        /// Up to 4 global filters can be specified using Key/Value pairs.
        /// Valid Keys are `PRODUCT`, `CLASSIFICATION`, `MSRC_SEVERITY`, and `PATCH_ID`.
        /// </summary>
        public InputList<Inputs.PatchBaselineGlobalFilterGetArgs> GlobalFilters
        {
            get => _globalFilters ?? (_globalFilters = new InputList<Inputs.PatchBaselineGlobalFilterGetArgs>());
            set => _globalFilters = value;
        }

        /// <summary>
        /// The name of the patch baseline.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The operating system the patch baseline applies to.
        /// Valid values are
        /// `ALMA_LINUX`,
        /// `AMAZON_LINUX`,
        /// `AMAZON_LINUX_2`,
        /// `AMAZON_LINUX_2022`,
        /// `AMAZON_LINUX_2023`,
        /// `CENTOS`,
        /// `DEBIAN`,
        /// `MACOS`,
        /// `ORACLE_LINUX`,
        /// `RASPBIAN`,
        /// `REDHAT_ENTERPRISE_LINUX`,
        /// `ROCKY_LINUX`,
        /// `SUSE`,
        /// `UBUNTU`, and
        /// `WINDOWS`.
        /// The default value is `WINDOWS`.
        /// </summary>
        [Input("operatingSystem")]
        public Input<string>? OperatingSystem { get; set; }

        [Input("rejectedPatches")]
        private InputList<string>? _rejectedPatches;

        /// <summary>
        /// A list of rejected patches.
        /// </summary>
        public InputList<string> RejectedPatches
        {
            get => _rejectedPatches ?? (_rejectedPatches = new InputList<string>());
            set => _rejectedPatches = value;
        }

        /// <summary>
        /// The action for Patch Manager to take on patches included in the `rejected_patches` list.
        /// Valid values are `ALLOW_AS_DEPENDENCY` and `BLOCK`.
        /// </summary>
        [Input("rejectedPatchesAction")]
        public Input<string>? RejectedPatchesAction { get; set; }

        [Input("sources")]
        private InputList<Inputs.PatchBaselineSourceGetArgs>? _sources;

        /// <summary>
        /// Configuration block with alternate sources for patches.
        /// Applies to Linux instances only.
        /// See `source` below.
        /// </summary>
        public InputList<Inputs.PatchBaselineSourceGetArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.PatchBaselineSourceGetArgs>());
            set => _sources = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public PatchBaselineState()
        {
        }
        public static new PatchBaselineState Empty => new PatchBaselineState();
    }
}
