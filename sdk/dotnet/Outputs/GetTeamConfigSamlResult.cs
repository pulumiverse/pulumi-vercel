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
    public sealed class GetTeamConfigSamlResult
    {
        /// <summary>
        /// The ID of the access group to use for the team.
        /// </summary>
        public readonly string AccessGroupId;
        /// <summary>
        /// Indicates if SAML is enforced for the team.
        /// </summary>
        public readonly bool Enforced;
        /// <summary>
        /// Directory groups to role or access group mappings.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Roles;

        [OutputConstructor]
        private GetTeamConfigSamlResult(
            string accessGroupId,

            bool enforced,

            ImmutableDictionary<string, string> roles)
        {
            AccessGroupId = accessGroupId;
            Enforced = enforced;
            Roles = roles;
        }
    }
}