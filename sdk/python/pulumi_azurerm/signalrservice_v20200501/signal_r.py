# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class SignalR(pulumi.CustomResource):
    kind: pulumi.Output[str]
    """
    The kind of the service - e.g. "SignalR", or "RawWebSockets" for "Microsoft.SignalRService/SignalR"
    """
    location: pulumi.Output[str]
    """
    The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
    """
    name: pulumi.Output[str]
    """
    The name of the resource.
    """
    properties: pulumi.Output[dict]
    """
    Settings used to provision or configure the resource
      * `cors` (`dict`) - Cross-Origin Resource Sharing (CORS) settings.
        * `allowed_origins` (`list`) - Gets or sets the list of origins that should be allowed to make cross-origin calls (for example: http://example.com:12345). Use "*" to allow all. If omitted, allow all by default.

      * `external_ip` (`str`) - The publicly accessible IP of the SignalR service.
      * `features` (`list`) - List of SignalR featureFlags. e.g. ServiceMode.
        
        FeatureFlags that are not included in the parameters for the update operation will not be modified.
        And the response will only include featureFlags that are explicitly set. 
        When a featureFlag is not explicitly set, SignalR service will use its globally default value. 
        But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
        * `flag` (`str`) - FeatureFlags is the supported features of Azure SignalR service.
          - ServiceMode: Flag for backend server for SignalR service. Values allowed: "Default": have your own backend server; "Serverless": your application doesn't have a backend server; "Classic": for backward compatibility. Support both Default and Serverless mode but not recommended; "PredefinedOnly": for future use.
          - EnableConnectivityLogs: "true"/"false", to enable/disable the connectivity log category respectively.
        * `properties` (`dict`) - Optional properties related to this feature.
        * `value` (`str`) - Value of the feature flag. See Azure SignalR service document https://docs.microsoft.com/azure/azure-signalr/ for allowed values.

      * `host_name` (`str`) - FQDN of the SignalR service instance. Format: xxx.service.signalr.net
      * `host_name_prefix` (`str`) - Prefix for the hostName of the SignalR service. Retained for future use.
        The hostname will be of format: &lt;hostNamePrefix&gt;.service.signalr.net.
      * `network_ac_ls` (`dict`) - Network ACLs
        * `default_action` (`str`) - Default action when no other rule matches
        * `private_endpoints` (`list`) - ACLs for requests from private endpoints
          * `allow` (`list`) - Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
          * `deny` (`list`) - Denied request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
          * `name` (`str`) - Name of the private endpoint connection

        * `public_network` (`dict`) - ACL for requests from public network
          * `allow` (`list`) - Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
          * `deny` (`list`) - Denied request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.

      * `private_endpoint_connections` (`list`) - Private endpoint connections to the SignalR resource.
        * `id` (`str`) - Fully qualified resource Id for the resource.
        * `name` (`str`) - The name of the resource.
        * `properties` (`dict`) - Properties of the private endpoint connection
          * `private_endpoint` (`dict`) - Private endpoint associated with the private endpoint connection
            * `id` (`str`) - Full qualified Id of the private endpoint

          * `private_link_service_connection_state` (`dict`) - Connection state
            * `actions_required` (`str`) - A message indicating if changes on the service provider require any updates on the consumer.
            * `description` (`str`) - The reason for approval/rejection of the connection.
            * `status` (`str`) - Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.

          * `provisioning_state` (`str`) - Provisioning state of the private endpoint connection

        * `type` (`str`) - The type of the resource - e.g. "Microsoft.SignalRService/SignalR"

      * `provisioning_state` (`str`) - Provisioning state of the resource.
      * `public_port` (`float`) - The publicly accessible port of the SignalR service which is designed for browser/client side usage.
      * `server_port` (`float`) - The publicly accessible port of the SignalR service which is designed for customer server side usage.
      * `upstream` (`dict`) - Upstream settings when the Azure SignalR is in server-less mode.
        * `templates` (`list`) - Gets or sets the list of Upstream URL templates. Order matters, and the first matching template takes effects.
          * `category_pattern` (`str`) - Gets or sets the matching pattern for category names. If not set, it matches any category.
            There are 3 kind of patterns supported:
                1. "*", it to matches any category name
                2. Combine multiple categories with ",", for example "connections,messages", it matches category "connections" and "messages"
                3. The single category name, for example, "connections", it matches the category "connections"
          * `event_pattern` (`str`) - Gets or sets the matching pattern for event names. If not set, it matches any event.
            There are 3 kind of patterns supported:
                1. "*", it to matches any event name
                2. Combine multiple events with ",", for example "connect,disconnect", it matches event "connect" and "disconnect"
                3. The single event name, for example, "connect", it matches "connect"
          * `hub_pattern` (`str`) - Gets or sets the matching pattern for hub names. If not set, it matches any hub.
            There are 3 kind of patterns supported:
                1. "*", it to matches any hub name
                2. Combine multiple hubs with ",", for example "hub1,hub2", it matches "hub1" and "hub2"
                3. The single hub name, for example, "hub1", it matches "hub1"
          * `url_template` (`str`) - Gets or sets the Upstream URL template. You can use 3 predefined parameters {hub}, {category} {event} inside the template, the value of the Upstream URL is dynamically calculated when the client request comes in.
            For example, if the urlTemplate is `http://example.com/{hub}/api/{event}`, with a client request from hub `chat` connects, it will first POST to this URL: `http://example.com/chat/api/connect`.

      * `version` (`str`) - Version of the SignalR resource. Probably you need the same or higher version of client SDKs.
    """
    sku: pulumi.Output[dict]
    """
    The billing information of the resource.(e.g. Free, Standard)
      * `capacity` (`float`) - Optional, integer. The unit count of SignalR resource. 1 by default.
        
        If present, following values are allowed:
            Free: 1
            Standard: 1,2,5,10,20,50,100
      * `family` (`str`) - Optional string. For future use.
      * `name` (`str`) - The name of the SKU. Required.
        
        Allowed values: Standard_S1, Free_F1
      * `size` (`str`) - Optional string. For future use.
      * `tier` (`str`) - Optional tier of this particular SKU. 'Standard' or 'Free'. 
        
        `Basic` is deprecated, use `Standard` instead.
    """
    tags: pulumi.Output[dict]
    """
    Tags of the service which is a list of key value pairs that describe the resource.
    """
    type: pulumi.Output[str]
    """
    The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
    """
    def __init__(__self__, resource_name, opts=None, kind=None, location=None, name=None, properties=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        A class represent a SignalR service resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] kind: The kind of the service - e.g. "SignalR", or "RawWebSockets" for "Microsoft.SignalRService/SignalR"
        :param pulumi.Input[str] location: The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
        :param pulumi.Input[str] name: The name of the SignalR resource.
        :param pulumi.Input[dict] properties: Settings used to provision or configure the resource
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        :param pulumi.Input[dict] sku: The billing information of the resource.(e.g. Free, Standard)
        :param pulumi.Input[dict] tags: Tags of the service which is a list of key value pairs that describe the resource.

        The **properties** object supports the following:

          * `cors` (`pulumi.Input[dict]`) - Cross-Origin Resource Sharing (CORS) settings.
            * `allowed_origins` (`pulumi.Input[list]`) - Gets or sets the list of origins that should be allowed to make cross-origin calls (for example: http://example.com:12345). Use "*" to allow all. If omitted, allow all by default.

          * `features` (`pulumi.Input[list]`) - List of SignalR featureFlags. e.g. ServiceMode.
            
            FeatureFlags that are not included in the parameters for the update operation will not be modified.
            And the response will only include featureFlags that are explicitly set. 
            When a featureFlag is not explicitly set, SignalR service will use its globally default value. 
            But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
            * `flag` (`pulumi.Input[str]`) - FeatureFlags is the supported features of Azure SignalR service.
              - ServiceMode: Flag for backend server for SignalR service. Values allowed: "Default": have your own backend server; "Serverless": your application doesn't have a backend server; "Classic": for backward compatibility. Support both Default and Serverless mode but not recommended; "PredefinedOnly": for future use.
              - EnableConnectivityLogs: "true"/"false", to enable/disable the connectivity log category respectively.
            * `properties` (`pulumi.Input[dict]`) - Optional properties related to this feature.
            * `value` (`pulumi.Input[str]`) - Value of the feature flag. See Azure SignalR service document https://docs.microsoft.com/azure/azure-signalr/ for allowed values.

          * `host_name_prefix` (`pulumi.Input[str]`) - Prefix for the hostName of the SignalR service. Retained for future use.
            The hostname will be of format: &lt;hostNamePrefix&gt;.service.signalr.net.
          * `network_ac_ls` (`pulumi.Input[dict]`) - Network ACLs
            * `default_action` (`pulumi.Input[str]`) - Default action when no other rule matches
            * `private_endpoints` (`pulumi.Input[list]`) - ACLs for requests from private endpoints
              * `allow` (`pulumi.Input[list]`) - Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
              * `deny` (`pulumi.Input[list]`) - Denied request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
              * `name` (`pulumi.Input[str]`) - Name of the private endpoint connection

            * `public_network` (`pulumi.Input[dict]`) - ACL for requests from public network
              * `allow` (`pulumi.Input[list]`) - Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
              * `deny` (`pulumi.Input[list]`) - Denied request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.

          * `upstream` (`pulumi.Input[dict]`) - Upstream settings when the Azure SignalR is in server-less mode.
            * `templates` (`pulumi.Input[list]`) - Gets or sets the list of Upstream URL templates. Order matters, and the first matching template takes effects.
              * `category_pattern` (`pulumi.Input[str]`) - Gets or sets the matching pattern for category names. If not set, it matches any category.
                There are 3 kind of patterns supported:
                    1. "*", it to matches any category name
                    2. Combine multiple categories with ",", for example "connections,messages", it matches category "connections" and "messages"
                    3. The single category name, for example, "connections", it matches the category "connections"
              * `event_pattern` (`pulumi.Input[str]`) - Gets or sets the matching pattern for event names. If not set, it matches any event.
                There are 3 kind of patterns supported:
                    1. "*", it to matches any event name
                    2. Combine multiple events with ",", for example "connect,disconnect", it matches event "connect" and "disconnect"
                    3. The single event name, for example, "connect", it matches "connect"
              * `hub_pattern` (`pulumi.Input[str]`) - Gets or sets the matching pattern for hub names. If not set, it matches any hub.
                There are 3 kind of patterns supported:
                    1. "*", it to matches any hub name
                    2. Combine multiple hubs with ",", for example "hub1,hub2", it matches "hub1" and "hub2"
                    3. The single hub name, for example, "hub1", it matches "hub1"
              * `url_template` (`pulumi.Input[str]`) - Gets or sets the Upstream URL template. You can use 3 predefined parameters {hub}, {category} {event} inside the template, the value of the Upstream URL is dynamically calculated when the client request comes in.
                For example, if the urlTemplate is `http://example.com/{hub}/api/{event}`, with a client request from hub `chat` connects, it will first POST to this URL: `http://example.com/chat/api/connect`.

        The **sku** object supports the following:

          * `capacity` (`pulumi.Input[float]`) - Optional, integer. The unit count of SignalR resource. 1 by default.
            
            If present, following values are allowed:
                Free: 1
                Standard: 1,2,5,10,20,50,100
          * `family` (`pulumi.Input[str]`) - Optional string. For future use.
          * `name` (`pulumi.Input[str]`) - The name of the SKU. Required.
            
            Allowed values: Standard_S1, Free_F1
          * `size` (`pulumi.Input[str]`) - Optional string. For future use.
          * `tier` (`pulumi.Input[str]`) - Optional tier of this particular SKU. 'Standard' or 'Free'. 
            
            `Basic` is deprecated, use `Standard` instead.
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

            __props__['kind'] = kind
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['type'] = None
        super(SignalR, __self__).__init__(
            'azurerm:signalrservice/v20200501:SignalR',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing SignalR resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return SignalR(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop