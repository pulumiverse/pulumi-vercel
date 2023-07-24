// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vercel.Inputs
{

    public sealed class ProjectVercelAuthenticationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, production deployments will also be protected
        /// </summary>
        [Input("protectProduction")]
        public Input<bool>? ProtectProduction { get; set; }

        public ProjectVercelAuthenticationArgs()
        {
        }
        public static new ProjectVercelAuthenticationArgs Empty => new ProjectVercelAuthenticationArgs();
    }
}
