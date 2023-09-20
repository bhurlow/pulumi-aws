// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds
{
    /// <summary>
    /// Provides an RDS DB option group resource. Documentation of the available options for various RDS engines can be found at:
    /// 
    /// * [MariaDB Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MariaDB.Options.html)
    /// * [Microsoft SQL Server Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.Options.html)
    /// * [MySQL Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MySQL.Options.html)
    /// * [Oracle Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.Options.html)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.Rds.OptionGroup("example", new()
    ///     {
    ///         OptionGroupDescription = "Option Group",
    ///         EngineName = "sqlserver-ee",
    ///         MajorEngineVersion = "11.00",
    ///         Options = new[]
    ///         {
    ///             new Aws.Rds.Inputs.OptionGroupOptionArgs
    ///             {
    ///                 OptionName = "Timezone",
    ///                 OptionSettings = new[]
    ///                 {
    ///                     new Aws.Rds.Inputs.OptionGroupOptionOptionSettingArgs
    ///                     {
    ///                         Name = "TIME_ZONE",
    ///                         Value = "UTC",
    ///                     },
    ///                 },
    ///             },
    ///             new Aws.Rds.Inputs.OptionGroupOptionArgs
    ///             {
    ///                 OptionName = "SQLSERVER_BACKUP_RESTORE",
    ///                 OptionSettings = new[]
    ///                 {
    ///                     new Aws.Rds.Inputs.OptionGroupOptionOptionSettingArgs
    ///                     {
    ///                         Name = "IAM_ROLE_ARN",
    ///                         Value = aws_iam_role.Example.Arn,
    ///                     },
    ///                 },
    ///             },
    ///             new Aws.Rds.Inputs.OptionGroupOptionArgs
    ///             {
    ///                 OptionName = "TDE",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &gt; **Note:** Any modifications to the `aws.rds.OptionGroup` are set to happen immediately as we default to applying immediately.
    /// 
    /// &gt; **WARNING:** You can perform a destroy on a `aws.rds.OptionGroup`, as long as it is not associated with any Amazon RDS resource. An option group can be associated with a DB instance, a manual DB snapshot, or an automated DB snapshot.
    /// 
    /// If you try to delete an option group that is associated with an Amazon RDS resource, an error similar to the following is returned:
    /// 
    /// &gt; An error occurred (InvalidOptionGroupStateFault) when calling the DeleteOptionGroup operation: The option group 'optionGroupName' cannot be deleted because it is in use.
    /// 
    /// More information about this can be found [here](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithOptionGroups.html#USER_WorkingWithOptionGroups.Delete).
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import DB Option groups using the `name`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:rds/optionGroup:OptionGroup example mysql-option-group
    /// ```
    /// </summary>
    [AwsResourceType("aws:rds/optionGroup:OptionGroup")]
    public partial class OptionGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the db option group.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the engine that this option group should be associated with.
        /// </summary>
        [Output("engineName")]
        public Output<string> EngineName { get; private set; } = null!;

        /// <summary>
        /// Specifies the major version of the engine that this option group should be associated with.
        /// </summary>
        [Output("majorEngineVersion")]
        public Output<string> MajorEngineVersion { get; private set; } = null!;

        /// <summary>
        /// The Name of the setting.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Must be lowercase, to match as it is stored in AWS.
        /// </summary>
        [Output("namePrefix")]
        public Output<string> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// The description of the option group. Defaults to "Managed by Pulumi".
        /// </summary>
        [Output("optionGroupDescription")]
        public Output<string> OptionGroupDescription { get; private set; } = null!;

        /// <summary>
        /// A list of Options to apply.
        /// </summary>
        [Output("options")]
        public Output<ImmutableArray<Outputs.OptionGroupOption>> Options { get; private set; } = null!;

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
        /// Create a OptionGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OptionGroup(string name, OptionGroupArgs args, CustomResourceOptions? options = null)
            : base("aws:rds/optionGroup:OptionGroup", name, args ?? new OptionGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OptionGroup(string name, Input<string> id, OptionGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:rds/optionGroup:OptionGroup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "tagsAll",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OptionGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OptionGroup Get(string name, Input<string> id, OptionGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new OptionGroup(name, id, state, options);
        }
    }

    public sealed class OptionGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the engine that this option group should be associated with.
        /// </summary>
        [Input("engineName", required: true)]
        public Input<string> EngineName { get; set; } = null!;

        /// <summary>
        /// Specifies the major version of the engine that this option group should be associated with.
        /// </summary>
        [Input("majorEngineVersion", required: true)]
        public Input<string> MajorEngineVersion { get; set; } = null!;

        /// <summary>
        /// The Name of the setting.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Must be lowercase, to match as it is stored in AWS.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The description of the option group. Defaults to "Managed by Pulumi".
        /// </summary>
        [Input("optionGroupDescription")]
        public Input<string>? OptionGroupDescription { get; set; }

        [Input("options")]
        private InputList<Inputs.OptionGroupOptionArgs>? _options;

        /// <summary>
        /// A list of Options to apply.
        /// </summary>
        public InputList<Inputs.OptionGroupOptionArgs> Options
        {
            get => _options ?? (_options = new InputList<Inputs.OptionGroupOptionArgs>());
            set => _options = value;
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

        public OptionGroupArgs()
        {
            OptionGroupDescription = "Managed by Pulumi";
        }
        public static new OptionGroupArgs Empty => new OptionGroupArgs();
    }

    public sealed class OptionGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the db option group.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Specifies the name of the engine that this option group should be associated with.
        /// </summary>
        [Input("engineName")]
        public Input<string>? EngineName { get; set; }

        /// <summary>
        /// Specifies the major version of the engine that this option group should be associated with.
        /// </summary>
        [Input("majorEngineVersion")]
        public Input<string>? MajorEngineVersion { get; set; }

        /// <summary>
        /// The Name of the setting.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Must be lowercase, to match as it is stored in AWS.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// The description of the option group. Defaults to "Managed by Pulumi".
        /// </summary>
        [Input("optionGroupDescription")]
        public Input<string>? OptionGroupDescription { get; set; }

        [Input("options")]
        private InputList<Inputs.OptionGroupOptionGetArgs>? _options;

        /// <summary>
        /// A list of Options to apply.
        /// </summary>
        public InputList<Inputs.OptionGroupOptionGetArgs> Options
        {
            get => _options ?? (_options = new InputList<Inputs.OptionGroupOptionGetArgs>());
            set => _options = value;
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
        [Obsolete(@"Please use `tags` instead.")]
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _tagsAll = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        public OptionGroupState()
        {
            OptionGroupDescription = "Managed by Pulumi";
        }
        public static new OptionGroupState Empty => new OptionGroupState();
    }
}
