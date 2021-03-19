// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeBuild.Outputs
{

    [OutputType]
    public sealed class ProjectSource
    {
        /// <summary>
        /// Configuration block with the authorization settings for AWS CodeBuild to access the source code to be built. This information is for the AWS CodeBuild console's use only. Use the `aws.codebuild.SourceCredential` resource instead. Auth blocks are documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ProjectSourceAuth> Auths;
        /// <summary>
        /// Build specification to use for this build project's related builds. This must be set when `type` is `NO_SOURCE`.
        /// </summary>
        public readonly string? Buildspec;
        /// <summary>
        /// Truncate git history to this many commits. Use `0` for a `Full` checkout which you need to run commands like `git branch --show-current`. See [AWS CodePipeline User Guide: Tutorial: Use full clone with a GitHub pipeline source](https://docs.aws.amazon.com/codepipeline/latest/userguide/tutorials-github-gitclone.html) for details.
        /// </summary>
        public readonly int? GitCloneDepth;
        /// <summary>
        /// Configuration block. Detailed below.
        /// </summary>
        public readonly Outputs.ProjectSourceGitSubmodulesConfig? GitSubmodulesConfig;
        /// <summary>
        /// Ignore SSL warnings when connecting to source control.
        /// </summary>
        public readonly bool? InsecureSsl;
        /// <summary>
        /// Location of the source code from git or s3.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Whether to report the status of a build's start and finish to your source provider. This option is only valid when the `type` is `BITBUCKET` or `GITHUB`.
        /// </summary>
        public readonly bool? ReportBuildStatus;
        /// <summary>
        /// Authorization type to use. The only valid value is `OAUTH`. This data type is deprecated and is no longer accurate or used. Use the `aws.codebuild.SourceCredential` resource instead.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ProjectSource(
            ImmutableArray<Outputs.ProjectSourceAuth> auths,

            string? buildspec,

            int? gitCloneDepth,

            Outputs.ProjectSourceGitSubmodulesConfig? gitSubmodulesConfig,

            bool? insecureSsl,

            string? location,

            bool? reportBuildStatus,

            string type)
        {
            Auths = auths;
            Buildspec = buildspec;
            GitCloneDepth = gitCloneDepth;
            GitSubmodulesConfig = gitSubmodulesConfig;
            InsecureSsl = insecureSsl;
            Location = location;
            ReportBuildStatus = reportBuildStatus;
            Type = type;
        }
    }
}
