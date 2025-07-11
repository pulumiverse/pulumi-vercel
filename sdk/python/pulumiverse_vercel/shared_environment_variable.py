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

__all__ = ['SharedEnvironmentVariableArgs', 'SharedEnvironmentVariable']

@pulumi.input_type
class SharedEnvironmentVariableArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 project_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 targets: pulumi.Input[Sequence[pulumi.Input[str]]],
                 value: pulumi.Input[str],
                 apply_to_all_custom_environments: Optional[pulumi.Input[bool]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 sensitive: Optional[pulumi.Input[bool]] = None,
                 team_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SharedEnvironmentVariable resource.
        :param pulumi.Input[str] key: The name of the Environment Variable.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] project_ids: The ID of the Vercel project.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] targets: The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`.
        :param pulumi.Input[str] value: The value of the Environment Variable.
        :param pulumi.Input[bool] apply_to_all_custom_environments: Whether the shared environment variable should be applied to all custom environments in the linked projects.
        :param pulumi.Input[str] comment: A comment explaining what the environment variable is for.
        :param pulumi.Input[bool] sensitive: Whether the Environment Variable is sensitive or not. (May be affected by a [team-wide environment variable policy](https://vercel.com/docs/projects/environment-variables/sensitive-environment-variables#environment-variables-policy))
        :param pulumi.Input[str] team_id: The ID of the Vercel team. Shared environment variables require a team.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "project_ids", project_ids)
        pulumi.set(__self__, "targets", targets)
        pulumi.set(__self__, "value", value)
        if apply_to_all_custom_environments is not None:
            pulumi.set(__self__, "apply_to_all_custom_environments", apply_to_all_custom_environments)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if sensitive is not None:
            pulumi.set(__self__, "sensitive", sensitive)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The name of the Environment Variable.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="projectIds")
    def project_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The ID of the Vercel project.
        """
        return pulumi.get(self, "project_ids")

    @project_ids.setter
    def project_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "project_ids", value)

    @property
    @pulumi.getter
    def targets(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`.
        """
        return pulumi.get(self, "targets")

    @targets.setter
    def targets(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "targets", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value of the Environment Variable.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="applyToAllCustomEnvironments")
    def apply_to_all_custom_environments(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the shared environment variable should be applied to all custom environments in the linked projects.
        """
        return pulumi.get(self, "apply_to_all_custom_environments")

    @apply_to_all_custom_environments.setter
    def apply_to_all_custom_environments(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "apply_to_all_custom_environments", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        A comment explaining what the environment variable is for.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def sensitive(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the Environment Variable is sensitive or not. (May be affected by a [team-wide environment variable policy](https://vercel.com/docs/projects/environment-variables/sensitive-environment-variables#environment-variables-policy))
        """
        return pulumi.get(self, "sensitive")

    @sensitive.setter
    def sensitive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sensitive", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Vercel team. Shared environment variables require a team.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_id", value)


@pulumi.input_type
class _SharedEnvironmentVariableState:
    def __init__(__self__, *,
                 apply_to_all_custom_environments: Optional[pulumi.Input[bool]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 project_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sensitive: Optional[pulumi.Input[bool]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SharedEnvironmentVariable resources.
        :param pulumi.Input[bool] apply_to_all_custom_environments: Whether the shared environment variable should be applied to all custom environments in the linked projects.
        :param pulumi.Input[str] comment: A comment explaining what the environment variable is for.
        :param pulumi.Input[str] key: The name of the Environment Variable.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] project_ids: The ID of the Vercel project.
        :param pulumi.Input[bool] sensitive: Whether the Environment Variable is sensitive or not. (May be affected by a [team-wide environment variable policy](https://vercel.com/docs/projects/environment-variables/sensitive-environment-variables#environment-variables-policy))
        :param pulumi.Input[Sequence[pulumi.Input[str]]] targets: The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`.
        :param pulumi.Input[str] team_id: The ID of the Vercel team. Shared environment variables require a team.
        :param pulumi.Input[str] value: The value of the Environment Variable.
        """
        if apply_to_all_custom_environments is not None:
            pulumi.set(__self__, "apply_to_all_custom_environments", apply_to_all_custom_environments)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if project_ids is not None:
            pulumi.set(__self__, "project_ids", project_ids)
        if sensitive is not None:
            pulumi.set(__self__, "sensitive", sensitive)
        if targets is not None:
            pulumi.set(__self__, "targets", targets)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="applyToAllCustomEnvironments")
    def apply_to_all_custom_environments(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the shared environment variable should be applied to all custom environments in the linked projects.
        """
        return pulumi.get(self, "apply_to_all_custom_environments")

    @apply_to_all_custom_environments.setter
    def apply_to_all_custom_environments(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "apply_to_all_custom_environments", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        A comment explaining what the environment variable is for.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Environment Variable.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="projectIds")
    def project_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The ID of the Vercel project.
        """
        return pulumi.get(self, "project_ids")

    @project_ids.setter
    def project_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "project_ids", value)

    @property
    @pulumi.getter
    def sensitive(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the Environment Variable is sensitive or not. (May be affected by a [team-wide environment variable policy](https://vercel.com/docs/projects/environment-variables/sensitive-environment-variables#environment-variables-policy))
        """
        return pulumi.get(self, "sensitive")

    @sensitive.setter
    def sensitive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sensitive", value)

    @property
    @pulumi.getter
    def targets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`.
        """
        return pulumi.get(self, "targets")

    @targets.setter
    def targets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "targets", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Vercel team. Shared environment variables require a team.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_id", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value of the Environment Variable.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


class SharedEnvironmentVariable(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apply_to_all_custom_environments: Optional[pulumi.Input[bool]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 project_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sensitive: Optional[pulumi.Input[bool]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Shared Environment Variable resource.

        A Shared Environment Variable resource defines an Environment Variable that can be shared between multiple Vercel Projects.

        For more detailed information, please see the [Vercel documentation](https://vercel.com/docs/concepts/projects/environment-variables/shared-environment-variables).

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_vercel as vercel

        example = vercel.Project("example",
            name="example",
            git_repository={
                "type": "github",
                "repo": "vercel/some-repo",
            })
        # A shared environment variable that will be created
        # and associated with the "example" project.
        example_shared_environment_variable = vercel.SharedEnvironmentVariable("example",
            key="EXAMPLE",
            value="some_value",
            targets=["production"],
            comment="an example shared variable",
            project_ids=[example.id])
        ```

        ## Import

        You can import via the team_id and environment variable id.

        - team_id can be found in the team `settings` tab in the Vercel UI.

        - environment variable id can be taken from the network tab inside developer tools, while you are on the project page.

        # 

        Note also, that the value field for sensitive environment variables will be imported as `null`.

        ```sh
        $ pulumi import vercel:index/sharedEnvironmentVariable:SharedEnvironmentVariable example team_xxxxxxxxxxxxxxxxxxxxxxxx/env_yyyyyyyyyyyyy
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] apply_to_all_custom_environments: Whether the shared environment variable should be applied to all custom environments in the linked projects.
        :param pulumi.Input[str] comment: A comment explaining what the environment variable is for.
        :param pulumi.Input[str] key: The name of the Environment Variable.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] project_ids: The ID of the Vercel project.
        :param pulumi.Input[bool] sensitive: Whether the Environment Variable is sensitive or not. (May be affected by a [team-wide environment variable policy](https://vercel.com/docs/projects/environment-variables/sensitive-environment-variables#environment-variables-policy))
        :param pulumi.Input[Sequence[pulumi.Input[str]]] targets: The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`.
        :param pulumi.Input[str] team_id: The ID of the Vercel team. Shared environment variables require a team.
        :param pulumi.Input[str] value: The value of the Environment Variable.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SharedEnvironmentVariableArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Shared Environment Variable resource.

        A Shared Environment Variable resource defines an Environment Variable that can be shared between multiple Vercel Projects.

        For more detailed information, please see the [Vercel documentation](https://vercel.com/docs/concepts/projects/environment-variables/shared-environment-variables).

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_vercel as vercel

        example = vercel.Project("example",
            name="example",
            git_repository={
                "type": "github",
                "repo": "vercel/some-repo",
            })
        # A shared environment variable that will be created
        # and associated with the "example" project.
        example_shared_environment_variable = vercel.SharedEnvironmentVariable("example",
            key="EXAMPLE",
            value="some_value",
            targets=["production"],
            comment="an example shared variable",
            project_ids=[example.id])
        ```

        ## Import

        You can import via the team_id and environment variable id.

        - team_id can be found in the team `settings` tab in the Vercel UI.

        - environment variable id can be taken from the network tab inside developer tools, while you are on the project page.

        # 

        Note also, that the value field for sensitive environment variables will be imported as `null`.

        ```sh
        $ pulumi import vercel:index/sharedEnvironmentVariable:SharedEnvironmentVariable example team_xxxxxxxxxxxxxxxxxxxxxxxx/env_yyyyyyyyyyyyy
        ```

        :param str resource_name: The name of the resource.
        :param SharedEnvironmentVariableArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SharedEnvironmentVariableArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apply_to_all_custom_environments: Optional[pulumi.Input[bool]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 project_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sensitive: Optional[pulumi.Input[bool]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SharedEnvironmentVariableArgs.__new__(SharedEnvironmentVariableArgs)

            __props__.__dict__["apply_to_all_custom_environments"] = apply_to_all_custom_environments
            __props__.__dict__["comment"] = comment
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            if project_ids is None and not opts.urn:
                raise TypeError("Missing required property 'project_ids'")
            __props__.__dict__["project_ids"] = project_ids
            __props__.__dict__["sensitive"] = sensitive
            if targets is None and not opts.urn:
                raise TypeError("Missing required property 'targets'")
            __props__.__dict__["targets"] = targets
            __props__.__dict__["team_id"] = team_id
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = None if value is None else pulumi.Output.secret(value)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["value"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SharedEnvironmentVariable, __self__).__init__(
            'vercel:index/sharedEnvironmentVariable:SharedEnvironmentVariable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            apply_to_all_custom_environments: Optional[pulumi.Input[bool]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            key: Optional[pulumi.Input[str]] = None,
            project_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            sensitive: Optional[pulumi.Input[bool]] = None,
            targets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            team_id: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None) -> 'SharedEnvironmentVariable':
        """
        Get an existing SharedEnvironmentVariable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] apply_to_all_custom_environments: Whether the shared environment variable should be applied to all custom environments in the linked projects.
        :param pulumi.Input[str] comment: A comment explaining what the environment variable is for.
        :param pulumi.Input[str] key: The name of the Environment Variable.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] project_ids: The ID of the Vercel project.
        :param pulumi.Input[bool] sensitive: Whether the Environment Variable is sensitive or not. (May be affected by a [team-wide environment variable policy](https://vercel.com/docs/projects/environment-variables/sensitive-environment-variables#environment-variables-policy))
        :param pulumi.Input[Sequence[pulumi.Input[str]]] targets: The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`.
        :param pulumi.Input[str] team_id: The ID of the Vercel team. Shared environment variables require a team.
        :param pulumi.Input[str] value: The value of the Environment Variable.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SharedEnvironmentVariableState.__new__(_SharedEnvironmentVariableState)

        __props__.__dict__["apply_to_all_custom_environments"] = apply_to_all_custom_environments
        __props__.__dict__["comment"] = comment
        __props__.__dict__["key"] = key
        __props__.__dict__["project_ids"] = project_ids
        __props__.__dict__["sensitive"] = sensitive
        __props__.__dict__["targets"] = targets
        __props__.__dict__["team_id"] = team_id
        __props__.__dict__["value"] = value
        return SharedEnvironmentVariable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applyToAllCustomEnvironments")
    def apply_to_all_custom_environments(self) -> pulumi.Output[bool]:
        """
        Whether the shared environment variable should be applied to all custom environments in the linked projects.
        """
        return pulumi.get(self, "apply_to_all_custom_environments")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[str]:
        """
        A comment explaining what the environment variable is for.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        The name of the Environment Variable.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="projectIds")
    def project_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        The ID of the Vercel project.
        """
        return pulumi.get(self, "project_ids")

    @property
    @pulumi.getter
    def sensitive(self) -> pulumi.Output[bool]:
        """
        Whether the Environment Variable is sensitive or not. (May be affected by a [team-wide environment variable policy](https://vercel.com/docs/projects/environment-variables/sensitive-environment-variables#environment-variables-policy))
        """
        return pulumi.get(self, "sensitive")

    @property
    @pulumi.getter
    def targets(self) -> pulumi.Output[Sequence[str]]:
        """
        The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`.
        """
        return pulumi.get(self, "targets")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Output[str]:
        """
        The ID of the Vercel team. Shared environment variables require a team.
        """
        return pulumi.get(self, "team_id")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        The value of the Environment Variable.
        """
        return pulumi.get(self, "value")

