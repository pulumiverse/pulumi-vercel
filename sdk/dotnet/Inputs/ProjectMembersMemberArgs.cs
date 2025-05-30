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

    public sealed class ProjectMembersMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The email of the user to add to the project. Exactly one of `user_id`, `email`, or `username` must be specified.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The role that the user should have in the project. One of 'MEMBER', 'PROJECT*DEVELOPER', or 'PROJECT*VIEWER'.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The ID of the user to add to the project. Exactly one of `user_id`, `email`, or `username` must be specified.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        /// <summary>
        /// The username of the user to add to the project. Exactly one of `user_id`, `email`, or `username` must be specified.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ProjectMembersMemberArgs()
        {
        }
        public static new ProjectMembersMemberArgs Empty => new ProjectMembersMemberArgs();
    }
}
