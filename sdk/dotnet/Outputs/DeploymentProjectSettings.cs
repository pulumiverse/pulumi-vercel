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
    public sealed class DeploymentProjectSettings
    {
        /// <summary>
        /// The build command for this deployment. If omitted, this value will be taken from the project or automatically detected.
        /// </summary>
        public readonly string? BuildCommand;
        /// <summary>
        /// The framework that is being used for this deployment. If omitted, no framework is selected.
        /// </summary>
        public readonly string? Framework;
        /// <summary>
        /// The install command for this deployment. If omitted, this value will be taken from the project or automatically detected.
        /// </summary>
        public readonly string? InstallCommand;
        /// <summary>
        /// The output directory of the deployment. If omitted, this value will be taken from the project or automatically detected.
        /// </summary>
        public readonly string? OutputDirectory;
        /// <summary>
        /// The name of a directory or relative path to the source code of your project. When null is used it will default to the project root.
        /// </summary>
        public readonly string? RootDirectory;

        [OutputConstructor]
        private DeploymentProjectSettings(
            string? buildCommand,

            string? framework,

            string? installCommand,

            string? outputDirectory,

            string? rootDirectory)
        {
            BuildCommand = buildCommand;
            Framework = framework;
            InstallCommand = installCommand;
            OutputDirectory = outputDirectory;
            RootDirectory = rootDirectory;
        }
    }
}
