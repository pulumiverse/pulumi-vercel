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
    public sealed class GetTeamConfigSamlRolesResult
    {
        /// <summary>
        /// The access group the assign is assigned to.
        /// </summary>
        public readonly string AccessGroupId;
        /// <summary>
        /// The team level role the user is assigned. One of 'MEMBER', 'OWNER', 'VIEWER', 'DEVELOPER', 'BILLING' or 'CONTRIBUTOR'.
        /// </summary>
        public readonly string Role;

        [OutputConstructor]
        private GetTeamConfigSamlRolesResult(
            string accessGroupId,

            string role)
        {
            AccessGroupId = accessGroupId;
            Role = role;
        }
    }
}
