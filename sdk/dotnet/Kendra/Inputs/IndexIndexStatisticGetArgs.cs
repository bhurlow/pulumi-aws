// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kendra.Inputs
{

    public sealed class IndexIndexStatisticGetArgs : Pulumi.ResourceArgs
    {
        [Input("faqStatistics")]
        private InputList<Inputs.IndexIndexStatisticFaqStatisticGetArgs>? _faqStatistics;

        /// <summary>
        /// A block that specifies the number of question and answer topics in the index. Documented below.
        /// </summary>
        public InputList<Inputs.IndexIndexStatisticFaqStatisticGetArgs> FaqStatistics
        {
            get => _faqStatistics ?? (_faqStatistics = new InputList<Inputs.IndexIndexStatisticFaqStatisticGetArgs>());
            set => _faqStatistics = value;
        }

        [Input("textDocumentStatistics")]
        private InputList<Inputs.IndexIndexStatisticTextDocumentStatisticGetArgs>? _textDocumentStatistics;

        /// <summary>
        /// A block that specifies the number of text documents indexed.
        /// </summary>
        public InputList<Inputs.IndexIndexStatisticTextDocumentStatisticGetArgs> TextDocumentStatistics
        {
            get => _textDocumentStatistics ?? (_textDocumentStatistics = new InputList<Inputs.IndexIndexStatisticTextDocumentStatisticGetArgs>());
            set => _textDocumentStatistics = value;
        }

        public IndexIndexStatisticGetArgs()
        {
        }
    }
}