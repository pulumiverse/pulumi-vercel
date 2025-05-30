// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vercel.Inputs
{

    public sealed class ProjectOidcTokenConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When true, Vercel issued OpenID Connect (OIDC) tokens will be available on the compute environments. See https://vercel.com/docs/security/secure-backend-access/oidc for more information.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Configures the URL of the `iss` claim. `team` = `https://oidc.vercel.com/[team_slug]` `global` = `https://oidc.vercel.com`
        /// </summary>
        [Input("issuerMode")]
        public Input<string>? IssuerMode { get; set; }

        public ProjectOidcTokenConfigGetArgs()
        {
        }
        public static new ProjectOidcTokenConfigGetArgs Empty => new ProjectOidcTokenConfigGetArgs();
    }
}
