# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class ServiceDiagnostic(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    Diagnostic entity contract properties.
      * `always_log` (`str`) - Specifies for what type of messages sampling settings should not apply.
      * `backend` (`dict`) - Diagnostic settings for incoming/outgoing HTTP messages to the Backend
        * `request` (`dict`) - Diagnostic settings for request.
          * `body` (`dict`) - Body logging settings.
            * `bytes` (`float`) - Number of request body bytes to log.

          * `headers` (`list`) - Array of HTTP Headers to log.

        * `response` (`dict`) - Diagnostic settings for response.

      * `frontend` (`dict`) - Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
      * `http_correlation_protocol` (`str`) - Sets correlation protocol to use for Application Insights diagnostics.
      * `log_client_ip` (`bool`) - Log the ClientIP. Default is false.
      * `logger_id` (`str`) - Resource Id of a target logger.
      * `sampling` (`dict`) - Sampling settings for Diagnostic.
        * `percentage` (`float`) - Rate of sampling for fixed-rate sampling.
        * `sampling_type` (`str`) - Sampling type.

      * `verbosity` (`str`) - The verbosity level applied to traces emitted by trace policies.
    """
    type: pulumi.Output[str]
    """
    Resource type for API Management resource.
    """
    def __init__(__self__, resource_name, opts=None, name=None, properties=None, resource_group_name=None, service_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Diagnostic details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Diagnostic identifier. Must be unique in the current API Management service instance.
        :param pulumi.Input[dict] properties: Diagnostic entity contract properties.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] service_name: The name of the API Management service.

        The **properties** object supports the following:

          * `always_log` (`pulumi.Input[str]`) - Specifies for what type of messages sampling settings should not apply.
          * `backend` (`pulumi.Input[dict]`) - Diagnostic settings for incoming/outgoing HTTP messages to the Backend
            * `request` (`pulumi.Input[dict]`) - Diagnostic settings for request.
              * `body` (`pulumi.Input[dict]`) - Body logging settings.
                * `bytes` (`pulumi.Input[float]`) - Number of request body bytes to log.

              * `headers` (`pulumi.Input[list]`) - Array of HTTP Headers to log.

            * `response` (`pulumi.Input[dict]`) - Diagnostic settings for response.

          * `frontend` (`pulumi.Input[dict]`) - Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
          * `http_correlation_protocol` (`pulumi.Input[str]`) - Sets correlation protocol to use for Application Insights diagnostics.
          * `log_client_ip` (`pulumi.Input[bool]`) - Log the ClientIP. Default is false.
          * `logger_id` (`pulumi.Input[str]`) - Resource Id of a target logger.
          * `sampling` (`pulumi.Input[dict]`) - Sampling settings for Diagnostic.
            * `percentage` (`pulumi.Input[float]`) - Rate of sampling for fixed-rate sampling.
            * `sampling_type` (`pulumi.Input[str]`) - Sampling type.

          * `verbosity` (`pulumi.Input[str]`) - The verbosity level applied to traces emitted by trace policies.
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

            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if service_name is None:
                raise TypeError("Missing required property 'service_name'")
            __props__['service_name'] = service_name
            __props__['type'] = None
        super(ServiceDiagnostic, __self__).__init__(
            'azurerm:apimanagement:ServiceDiagnostic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ServiceDiagnostic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ServiceDiagnostic(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop