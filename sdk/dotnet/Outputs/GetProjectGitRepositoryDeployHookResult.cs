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
    public sealed class GetProjectGitRepositoryDeployHookResult
    {
        /// <summary>
        /// The ID of the deploy hook.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the deploy hook.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The branch or commit hash that should be deployed.
        /// </summary>
        public readonly string Ref;
        /// <summary>
        /// A URL that, when a POST request is made to, will trigger a new deployment.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetProjectGitRepositoryDeployHookResult(
            string id,

            string name,

            string @ref,

            string url)
        {
            Id = id;
            Name = name;
            Ref = @ref;
            Url = url;
        }
    }
}
