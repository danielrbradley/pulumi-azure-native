# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class ScheduledQueryRule(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Resource location
    """
    name: pulumi.Output[str]
    """
    Azure resource name
    """
    properties: pulumi.Output[dict]
    """
    The rule properties of the resource.
      * `action` (`dict`) - Action needs to be taken on rule execution.
      * `description` (`str`) - The description of the Log Search rule.
      * `enabled` (`str`) - The flag which indicates whether the Log Search rule is enabled. Value should be true or false
      * `last_updated_time` (`str`) - Last time the rule was updated in IS08601 format.
      * `provisioning_state` (`str`) - Provisioning state of the scheduled query rule
      * `schedule` (`dict`) - Schedule (Frequency, Time Window) for rule. Required for action type - AlertingAction
        * `frequency_in_minutes` (`float`) - frequency (in minutes) at which rule condition should be evaluated.
        * `time_window_in_minutes` (`float`) - Time window for which data needs to be fetched for query (should be greater than or equal to frequencyInMinutes).

      * `source` (`dict`) - Data Source against which rule will Query Data
        * `authorized_resources` (`list`) - List of  Resource referred into query
        * `data_source_id` (`str`) - The resource uri over which log search query is to be run.
        * `query` (`str`) - Log search query. Required for action type - AlertingAction
        * `query_type` (`str`) - Set value to 'ResultCount' .
    """
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Azure resource type
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        The Log Search Rule resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the rule.
        :param pulumi.Input[dict] properties: The rule properties of the resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] tags: Resource tags

        The **properties** object supports the following:

          * `action` (`pulumi.Input[dict]`) - Action needs to be taken on rule execution.
          * `description` (`pulumi.Input[str]`) - The description of the Log Search rule.
          * `enabled` (`pulumi.Input[str]`) - The flag which indicates whether the Log Search rule is enabled. Value should be true or false
          * `schedule` (`pulumi.Input[dict]`) - Schedule (Frequency, Time Window) for rule. Required for action type - AlertingAction
            * `frequency_in_minutes` (`pulumi.Input[float]`) - frequency (in minutes) at which rule condition should be evaluated.
            * `time_window_in_minutes` (`pulumi.Input[float]`) - Time window for which data needs to be fetched for query (should be greater than or equal to frequencyInMinutes).

          * `source` (`pulumi.Input[dict]`) - Data Source against which rule will Query Data
            * `authorized_resources` (`pulumi.Input[list]`) - List of  Resource referred into query
            * `data_source_id` (`pulumi.Input[str]`) - The resource uri over which log search query is to be run.
            * `query` (`pulumi.Input[str]`) - Log search query. Required for action type - AlertingAction
            * `query_type` (`pulumi.Input[str]`) - Set value to 'ResultCount' .
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['type'] = None
        super(ScheduledQueryRule, __self__).__init__(
            'azurerm:insights/v20180416:ScheduledQueryRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ScheduledQueryRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ScheduledQueryRule(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop