# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['CustomEnvironmentArgs', 'CustomEnvironment']

@pulumi.input_type
class CustomEnvironmentArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 branch_tracking: Optional[pulumi.Input['CustomEnvironmentBranchTrackingArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CustomEnvironment resource.
        :param pulumi.Input[str] project_id: The ID of the existing Vercel Project.
        :param pulumi.Input['CustomEnvironmentBranchTrackingArgs'] branch_tracking: The branch tracking configuration for the environment. When enabled, each qualifying merge will generate a deployment.
        :param pulumi.Input[str] description: A description of what the environment is.
        :param pulumi.Input[str] name: The name of the environment.
        :param pulumi.Input[str] team_id: The team ID to add the project to. Required when configuring a team resource if a default team has not been set in the provider.
        """
        pulumi.set(__self__, "project_id", project_id)
        if branch_tracking is not None:
            pulumi.set(__self__, "branch_tracking", branch_tracking)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The ID of the existing Vercel Project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="branchTracking")
    def branch_tracking(self) -> Optional[pulumi.Input['CustomEnvironmentBranchTrackingArgs']]:
        """
        The branch tracking configuration for the environment. When enabled, each qualifying merge will generate a deployment.
        """
        return pulumi.get(self, "branch_tracking")

    @branch_tracking.setter
    def branch_tracking(self, value: Optional[pulumi.Input['CustomEnvironmentBranchTrackingArgs']]):
        pulumi.set(self, "branch_tracking", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of what the environment is.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the environment.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[str]]:
        """
        The team ID to add the project to. Required when configuring a team resource if a default team has not been set in the provider.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_id", value)


@pulumi.input_type
class _CustomEnvironmentState:
    def __init__(__self__, *,
                 branch_tracking: Optional[pulumi.Input['CustomEnvironmentBranchTrackingArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CustomEnvironment resources.
        :param pulumi.Input['CustomEnvironmentBranchTrackingArgs'] branch_tracking: The branch tracking configuration for the environment. When enabled, each qualifying merge will generate a deployment.
        :param pulumi.Input[str] description: A description of what the environment is.
        :param pulumi.Input[str] name: The name of the environment.
        :param pulumi.Input[str] project_id: The ID of the existing Vercel Project.
        :param pulumi.Input[str] team_id: The team ID to add the project to. Required when configuring a team resource if a default team has not been set in the provider.
        """
        if branch_tracking is not None:
            pulumi.set(__self__, "branch_tracking", branch_tracking)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)

    @property
    @pulumi.getter(name="branchTracking")
    def branch_tracking(self) -> Optional[pulumi.Input['CustomEnvironmentBranchTrackingArgs']]:
        """
        The branch tracking configuration for the environment. When enabled, each qualifying merge will generate a deployment.
        """
        return pulumi.get(self, "branch_tracking")

    @branch_tracking.setter
    def branch_tracking(self, value: Optional[pulumi.Input['CustomEnvironmentBranchTrackingArgs']]):
        pulumi.set(self, "branch_tracking", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of what the environment is.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the environment.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the existing Vercel Project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[str]]:
        """
        The team ID to add the project to. Required when configuring a team resource if a default team has not been set in the provider.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_id", value)


class CustomEnvironment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch_tracking: Optional[pulumi.Input[Union['CustomEnvironmentBranchTrackingArgs', 'CustomEnvironmentBranchTrackingArgsDict']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Environments help manage the deployment lifecycle on the Vercel platform.

        By default, all teams use three environments when developing their project: Production, Preview, and Development. However, teams can also create custom environments to suit their needs. To learn more about the limits for each plan, see limits.

        Custom environments allow you to configure customized, pre-production environments for your project, such as staging or QA, with branch rules that will automatically deploy your branch when the branch name matches the rule. With custom environments you can also attach a domain to your environment, set environment variables, or import environment variables from another environment.

        Custom environments are designed as pre-production environments intended for long-running use. This contrasts with regular preview environments, which are designed for creating ephemeral, short-lived deployments.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_vercel as vercel

        example = vercel.Project("example", name="example-project-with-custom-env")
        example_custom_environment = vercel.CustomEnvironment("example",
            project_id=example.id,
            name="example-custom-env",
            description="A description of the custom environment",
            branch_tracking={
                "pattern": "staging-",
                "type": "startsWith",
            })
        ```

        ## Import

        If importing into a personal account, or with a team configured on

        the provider, simply use the project_id and custom environment name.

        - project_id can be found in the project `settings` tab in the Vercel UI.

        ```sh
        $ pulumi import vercel:index/customEnvironment:CustomEnvironment example prj_xxxxxxxxxxxxxxxxxxxxxxxxxxxx/example-custom-env
        ```

        Alternatively, you can import via the team_id, project_id and environment variable id.

        - team_id can be found in the team `settings` tab in the Vercel UI.

        - project_id can be found in the project `settings` tab in the Vercel UI.

        # 

        Note also, that the value field for sensitive environment variables will be imported as `null`.

        ```sh
        $ pulumi import vercel:index/customEnvironment:CustomEnvironment example team_xxxxxxxxxxxxxxxxxxxxxxxx/prj_xxxxxxxxxxxxxxxxxxxxxxxxxxxx/example-custom-env
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['CustomEnvironmentBranchTrackingArgs', 'CustomEnvironmentBranchTrackingArgsDict']] branch_tracking: The branch tracking configuration for the environment. When enabled, each qualifying merge will generate a deployment.
        :param pulumi.Input[str] description: A description of what the environment is.
        :param pulumi.Input[str] name: The name of the environment.
        :param pulumi.Input[str] project_id: The ID of the existing Vercel Project.
        :param pulumi.Input[str] team_id: The team ID to add the project to. Required when configuring a team resource if a default team has not been set in the provider.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CustomEnvironmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Environments help manage the deployment lifecycle on the Vercel platform.

        By default, all teams use three environments when developing their project: Production, Preview, and Development. However, teams can also create custom environments to suit their needs. To learn more about the limits for each plan, see limits.

        Custom environments allow you to configure customized, pre-production environments for your project, such as staging or QA, with branch rules that will automatically deploy your branch when the branch name matches the rule. With custom environments you can also attach a domain to your environment, set environment variables, or import environment variables from another environment.

        Custom environments are designed as pre-production environments intended for long-running use. This contrasts with regular preview environments, which are designed for creating ephemeral, short-lived deployments.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_vercel as vercel

        example = vercel.Project("example", name="example-project-with-custom-env")
        example_custom_environment = vercel.CustomEnvironment("example",
            project_id=example.id,
            name="example-custom-env",
            description="A description of the custom environment",
            branch_tracking={
                "pattern": "staging-",
                "type": "startsWith",
            })
        ```

        ## Import

        If importing into a personal account, or with a team configured on

        the provider, simply use the project_id and custom environment name.

        - project_id can be found in the project `settings` tab in the Vercel UI.

        ```sh
        $ pulumi import vercel:index/customEnvironment:CustomEnvironment example prj_xxxxxxxxxxxxxxxxxxxxxxxxxxxx/example-custom-env
        ```

        Alternatively, you can import via the team_id, project_id and environment variable id.

        - team_id can be found in the team `settings` tab in the Vercel UI.

        - project_id can be found in the project `settings` tab in the Vercel UI.

        # 

        Note also, that the value field for sensitive environment variables will be imported as `null`.

        ```sh
        $ pulumi import vercel:index/customEnvironment:CustomEnvironment example team_xxxxxxxxxxxxxxxxxxxxxxxx/prj_xxxxxxxxxxxxxxxxxxxxxxxxxxxx/example-custom-env
        ```

        :param str resource_name: The name of the resource.
        :param CustomEnvironmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomEnvironmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch_tracking: Optional[pulumi.Input[Union['CustomEnvironmentBranchTrackingArgs', 'CustomEnvironmentBranchTrackingArgsDict']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CustomEnvironmentArgs.__new__(CustomEnvironmentArgs)

            __props__.__dict__["branch_tracking"] = branch_tracking
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["team_id"] = team_id
        super(CustomEnvironment, __self__).__init__(
            'vercel:index/customEnvironment:CustomEnvironment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            branch_tracking: Optional[pulumi.Input[Union['CustomEnvironmentBranchTrackingArgs', 'CustomEnvironmentBranchTrackingArgsDict']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            team_id: Optional[pulumi.Input[str]] = None) -> 'CustomEnvironment':
        """
        Get an existing CustomEnvironment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['CustomEnvironmentBranchTrackingArgs', 'CustomEnvironmentBranchTrackingArgsDict']] branch_tracking: The branch tracking configuration for the environment. When enabled, each qualifying merge will generate a deployment.
        :param pulumi.Input[str] description: A description of what the environment is.
        :param pulumi.Input[str] name: The name of the environment.
        :param pulumi.Input[str] project_id: The ID of the existing Vercel Project.
        :param pulumi.Input[str] team_id: The team ID to add the project to. Required when configuring a team resource if a default team has not been set in the provider.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CustomEnvironmentState.__new__(_CustomEnvironmentState)

        __props__.__dict__["branch_tracking"] = branch_tracking
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["team_id"] = team_id
        return CustomEnvironment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="branchTracking")
    def branch_tracking(self) -> pulumi.Output['outputs.CustomEnvironmentBranchTracking']:
        """
        The branch tracking configuration for the environment. When enabled, each qualifying merge will generate a deployment.
        """
        return pulumi.get(self, "branch_tracking")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        A description of what the environment is.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the environment.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the existing Vercel Project.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Output[str]:
        """
        The team ID to add the project to. Required when configuring a team resource if a default team has not been set in the provider.
        """
        return pulumi.get(self, "team_id")

