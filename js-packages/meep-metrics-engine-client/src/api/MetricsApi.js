/*
 * Copyright (c) 2022  The AdvantEDGE Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * AdvantEDGE Metrics Service REST API
 * Metrics Service provides metrics about the active scenario <p>**Micro-service**<br>[meep-metrics-engine](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-metrics-engine) <p>**Type & Usage**<br>Platform Service used by control/monitoring software and possibly by edge applications that require metrics <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_
 *
 * OpenAPI spec version: 1.0.0
 * Contact: AdvantEDGE@InterDigital.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.9
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/DataflowMetrics', 'model/DataflowQueryParams', 'model/EventMetricList', 'model/EventQueryParams', 'model/HttpMetricList', 'model/HttpQueryParams', 'model/NetworkMetricList', 'model/NetworkQueryParams', 'model/SeqMetrics', 'model/SeqQueryParams'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/DataflowMetrics'), require('../model/DataflowQueryParams'), require('../model/EventMetricList'), require('../model/EventQueryParams'), require('../model/HttpMetricList'), require('../model/HttpQueryParams'), require('../model/NetworkMetricList'), require('../model/NetworkQueryParams'), require('../model/SeqMetrics'), require('../model/SeqQueryParams'));
  } else {
    // Browser globals (root is window)
    if (!root.AdvantEdgeMetricsServiceRestApi) {
      root.AdvantEdgeMetricsServiceRestApi = {};
    }
    root.AdvantEdgeMetricsServiceRestApi.MetricsApi = factory(root.AdvantEdgeMetricsServiceRestApi.ApiClient, root.AdvantEdgeMetricsServiceRestApi.DataflowMetrics, root.AdvantEdgeMetricsServiceRestApi.DataflowQueryParams, root.AdvantEdgeMetricsServiceRestApi.EventMetricList, root.AdvantEdgeMetricsServiceRestApi.EventQueryParams, root.AdvantEdgeMetricsServiceRestApi.HttpMetricList, root.AdvantEdgeMetricsServiceRestApi.HttpQueryParams, root.AdvantEdgeMetricsServiceRestApi.NetworkMetricList, root.AdvantEdgeMetricsServiceRestApi.NetworkQueryParams, root.AdvantEdgeMetricsServiceRestApi.SeqMetrics, root.AdvantEdgeMetricsServiceRestApi.SeqQueryParams);
  }
}(this, function(ApiClient, DataflowMetrics, DataflowQueryParams, EventMetricList, EventQueryParams, HttpMetricList, HttpQueryParams, NetworkMetricList, NetworkQueryParams, SeqMetrics, SeqQueryParams) {
  'use strict';

  /**
   * Metrics service.
   * @module api/MetricsApi
   * @version 1.0.0
   */

  /**
   * Constructs a new MetricsApi. 
   * @alias module:api/MetricsApi
   * @class
   * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance;


    /**
     * Callback function to receive the result of the postDataflowQuery operation.
     * @callback module:api/MetricsApi~postDataflowQueryCallback
     * @param {String} error Error message, if any.
     * @param {module:model/DataflowMetrics} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Requests dataflow diagram logs for the requested params
     * @param {module:model/DataflowQueryParams} params Query parameters
     * @param {module:api/MetricsApi~postDataflowQueryCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/DataflowMetrics}
     */
    this.postDataflowQuery = function(params, callback) {
      var postBody = params;

      // verify the required parameter 'params' is set
      if (params === undefined || params === null) {
        throw new Error("Missing the required parameter 'params' when calling postDataflowQuery");
      }


      var pathParams = {
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = DataflowMetrics;

      return this.apiClient.callApi(
        '/metrics/query/dataflow', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the postEventQuery operation.
     * @callback module:api/MetricsApi~postEventQueryCallback
     * @param {String} error Error message, if any.
     * @param {module:model/EventMetricList} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Returns Event metrics according to specificed parameters
     * @param {module:model/EventQueryParams} params Query parameters
     * @param {module:api/MetricsApi~postEventQueryCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/EventMetricList}
     */
    this.postEventQuery = function(params, callback) {
      var postBody = params;

      // verify the required parameter 'params' is set
      if (params === undefined || params === null) {
        throw new Error("Missing the required parameter 'params' when calling postEventQuery");
      }


      var pathParams = {
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = EventMetricList;

      return this.apiClient.callApi(
        '/metrics/query/event', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the postHttpQuery operation.
     * @callback module:api/MetricsApi~postHttpQueryCallback
     * @param {String} error Error message, if any.
     * @param {module:model/HttpMetricList} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Returns Http metrics according to specificed parameters
     * @param {module:model/HttpQueryParams} params Query parameters
     * @param {module:api/MetricsApi~postHttpQueryCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/HttpMetricList}
     */
    this.postHttpQuery = function(params, callback) {
      var postBody = params;

      // verify the required parameter 'params' is set
      if (params === undefined || params === null) {
        throw new Error("Missing the required parameter 'params' when calling postHttpQuery");
      }


      var pathParams = {
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = HttpMetricList;

      return this.apiClient.callApi(
        '/metrics/query/http', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the postNetworkQuery operation.
     * @callback module:api/MetricsApi~postNetworkQueryCallback
     * @param {String} error Error message, if any.
     * @param {module:model/NetworkMetricList} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Returns Network metrics according to specificed parameters
     * @param {module:model/NetworkQueryParams} params Query parameters
     * @param {module:api/MetricsApi~postNetworkQueryCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/NetworkMetricList}
     */
    this.postNetworkQuery = function(params, callback) {
      var postBody = params;

      // verify the required parameter 'params' is set
      if (params === undefined || params === null) {
        throw new Error("Missing the required parameter 'params' when calling postNetworkQuery");
      }


      var pathParams = {
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = NetworkMetricList;

      return this.apiClient.callApi(
        '/metrics/query/network', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the postSeqQuery operation.
     * @callback module:api/MetricsApi~postSeqQueryCallback
     * @param {String} error Error message, if any.
     * @param {module:model/SeqMetrics} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Requests sequence diagram logs for the requested params
     * @param {module:model/SeqQueryParams} params Query parameters
     * @param {module:api/MetricsApi~postSeqQueryCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/SeqMetrics}
     */
    this.postSeqQuery = function(params, callback) {
      var postBody = params;

      // verify the required parameter 'params' is set
      if (params === undefined || params === null) {
        throw new Error("Missing the required parameter 'params' when calling postSeqQuery");
      }


      var pathParams = {
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = SeqMetrics;

      return this.apiClient.callApi(
        '/metrics/query/seq', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));
