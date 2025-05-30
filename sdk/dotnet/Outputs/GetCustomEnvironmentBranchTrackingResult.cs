// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vercel.Outputs
{

    [OutputType]
    public sealed class GetCustomEnvironmentBranchTrackingResult
    {
        /// <summary>
        /// The pattern of the branch name to track.
        /// </summary>
        public readonly string Pattern;
        /// <summary>
        /// How a branch name should be matched against the pattern. Must be one of 'startsWith', 'endsWith' or 'equals'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetCustomEnvironmentBranchTrackingResult(
            string pattern,

            string type)
        {
            Pattern = pattern;
            Type = type;
        }
    }
}
