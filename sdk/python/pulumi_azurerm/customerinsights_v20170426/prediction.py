# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Prediction(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    The prediction definition.
      * `auto_analyze` (`bool`) - Whether do auto analyze.
      * `description` (`dict`) - Description of the prediction.
      * `display_name` (`dict`) - Display name of the prediction.
      * `grades` (`list`) - The prediction grades.
        * `grade_name` (`str`) - Name of the grade.
        * `max_score_threshold` (`float`) - Maximum score threshold.
        * `min_score_threshold` (`float`) - Minimum score threshold.

      * `involved_interaction_types` (`list`) - Interaction types involved in the prediction.
      * `involved_kpi_types` (`list`) - KPI types involved in the prediction.
      * `involved_relationships` (`list`) - Relationships involved in the prediction.
      * `mappings` (`dict`) - Definition of the link mapping of prediction.
      * `negative_outcome_expression` (`str`) - Negative outcome expression.
      * `positive_outcome_expression` (`str`) - Positive outcome expression.
      * `prediction_name` (`str`) - Name of the prediction.
      * `primary_profile_type` (`str`) - Primary profile type.
      * `provisioning_state` (`str`) - Provisioning state.
      * `scope_expression` (`str`) - Scope expression.
      * `score_label` (`str`) - Score label.
      * `system_generated_entities` (`dict`) - System generated entities.
      * `tenant_id` (`str`) - The hub name.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, hub_name=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        The prediction resource format.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hub_name: The name of the hub.
        :param pulumi.Input[str] name: The name of the Prediction.
        :param pulumi.Input[dict] properties: The prediction definition.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.

        The **properties** object supports the following:

          * `auto_analyze` (`pulumi.Input[bool]`) - Whether do auto analyze.
          * `description` (`pulumi.Input[dict]`) - Description of the prediction.
          * `display_name` (`pulumi.Input[dict]`) - Display name of the prediction.
          * `grades` (`pulumi.Input[list]`) - The prediction grades.
            * `grade_name` (`pulumi.Input[str]`) - Name of the grade.
            * `max_score_threshold` (`pulumi.Input[float]`) - Maximum score threshold.
            * `min_score_threshold` (`pulumi.Input[float]`) - Minimum score threshold.

          * `involved_interaction_types` (`pulumi.Input[list]`) - Interaction types involved in the prediction.
          * `involved_kpi_types` (`pulumi.Input[list]`) - KPI types involved in the prediction.
          * `involved_relationships` (`pulumi.Input[list]`) - Relationships involved in the prediction.
          * `mappings` (`pulumi.Input[dict]`) - Definition of the link mapping of prediction.
          * `negative_outcome_expression` (`pulumi.Input[str]`) - Negative outcome expression.
          * `positive_outcome_expression` (`pulumi.Input[str]`) - Positive outcome expression.
          * `prediction_name` (`pulumi.Input[str]`) - Name of the prediction.
          * `primary_profile_type` (`pulumi.Input[str]`) - Primary profile type.
          * `scope_expression` (`pulumi.Input[str]`) - Scope expression.
          * `score_label` (`pulumi.Input[str]`) - Score label.
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

            if hub_name is None:
                raise TypeError("Missing required property 'hub_name'")
            __props__['hub_name'] = hub_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['type'] = None
        super(Prediction, __self__).__init__(
            'azurerm:customerinsights/v20170426:Prediction',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Prediction resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Prediction(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop