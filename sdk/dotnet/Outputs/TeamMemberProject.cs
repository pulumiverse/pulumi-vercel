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
    public sealed class TeamMemberProject
    {
        /// <summary>
        /// The ID of the project that the user should be granted access to.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// The role that the user should have in the project.
        /// </summary>
        public readonly string Role;

        [OutputConstructor]
        private TeamMemberProject(
            string projectId,

            string role)
        {
            ProjectId = projectId;
            Role = role;
        }
    }
}
