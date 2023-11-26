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

    public sealed class ProjectVercelAuthenticationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The deployment environment to protect. Must be one of `standard_protection`, `all_deployments`, `only_preview_deployments`, or `none`.
        /// </summary>
        [Input("deploymentType", required: true)]
        public Input<string> DeploymentType { get; set; } = null!;

        public ProjectVercelAuthenticationArgs()
        {
        }
        public static new ProjectVercelAuthenticationArgs Empty => new ProjectVercelAuthenticationArgs();
    }
}
