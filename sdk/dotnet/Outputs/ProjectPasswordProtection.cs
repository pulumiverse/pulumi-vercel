// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vercel.Outputs
{

    [OutputType]
    public sealed class ProjectPasswordProtection
    {
        /// <summary>
        /// The password that visitors must enter to gain access to your Preview Deployments. Drift detection is not possible for this field.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// If true, production deployments will also be protected
        /// </summary>
        public readonly bool? ProtectProduction;

        [OutputConstructor]
        private ProjectPasswordProtection(
            string password,

            bool? protectProduction)
        {
            Password = password;
            ProtectProduction = protectProduction;
        }
    }
}
