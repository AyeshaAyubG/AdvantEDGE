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
 * MEC Application Support API
 *
 * The ETSI MEC ISG MEC011 MEC Application Support API described using OpenAPI
 *
 * API version: 2.2.1
 * Contact: cti_support@etsi.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

import (
	"fmt"
	"net/http"
	"strings"

	httpLog "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-http-logger"
	met "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-metrics"

	appSupport "github.com/InterDigitalInc/AdvantEDGE/go-apps/meep-app-enablement/server/app-support"
	svcMgmt "github.com/InterDigitalInc/AdvantEDGE/go-apps/meep-app-enablement/server/service-mgmt"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	var handler http.Handler
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		handler = met.MetricsHandler(handler, sandboxName, serviceName)
		handler = httpLog.LogRx(handler)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// Path prefix router order is important
	// Service Api files
	handler = http.StripPrefix("/mec_app_support/v1/api/", http.FileServer(http.Dir("./api/")))
	router.
		PathPrefix("/mec_app_support/v1/api/").
		Name("Api").
		Handler(handler)
	// User supplied service API files
	handler = http.StripPrefix("/mec_app_support/v1/user-api/", http.FileServer(http.Dir("./user-api/")))
	router.
		PathPrefix("/mec_app_support/v1/user-api/").
		Name("UserApi").
		Handler(handler)

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/mec_app_support/v1/",
		Index,
	},

	Route{
		"ApplicationsConfirmReadyPOST",
		strings.ToUpper("Post"),
		"/mec_app_support/v1/applications/{appInstanceId}/confirm_ready",
		appSupport.ApplicationsConfirmReadyPOST,
	},

	Route{
		"ApplicationsConfirmTerminationPOST",
		strings.ToUpper("Post"),
		"/mec_app_support/v1/applications/{appInstanceId}/confirm_termination",
		appSupport.ApplicationsConfirmTerminationPOST,
	},

	Route{
		"ApplicationsDnsRuleGET",
		strings.ToUpper("Get"),
		"/mec_app_support/v1/applications/{appInstanceId}/dns_rules/{dnsRuleId}",
		appSupport.ApplicationsDnsRuleGET,
	},

	Route{
		"ApplicationsDnsRulePUT",
		strings.ToUpper("Put"),
		"/mec_app_support/v1/applications/{appInstanceId}/dns_rules/{dnsRuleId}",
		appSupport.ApplicationsDnsRulePUT,
	},

	Route{
		"ApplicationsDnsRulesGET",
		strings.ToUpper("Get"),
		"/mec_app_support/v1/applications/{appInstanceId}/dns_rules",
		appSupport.ApplicationsDnsRulesGET,
	},

	Route{
		"ApplicationsSubscriptionDELETE",
		strings.ToUpper("Delete"),
		"/mec_app_support/v1/applications/{appInstanceId}/subscriptions/{subscriptionId}",
		appSupport.ApplicationsSubscriptionDELETE,
	},

	Route{
		"ApplicationsSubscriptionGET",
		strings.ToUpper("Get"),
		"/mec_app_support/v1/applications/{appInstanceId}/subscriptions/{subscriptionId}",
		appSupport.ApplicationsSubscriptionGET,
	},

	Route{
		"ApplicationsSubscriptionsGET",
		strings.ToUpper("Get"),
		"/mec_app_support/v1/applications/{appInstanceId}/subscriptions",
		appSupport.ApplicationsSubscriptionsGET,
	},

	Route{
		"ApplicationsSubscriptionsPOST",
		strings.ToUpper("Post"),
		"/mec_app_support/v1/applications/{appInstanceId}/subscriptions",
		appSupport.ApplicationsSubscriptionsPOST,
	},

	Route{
		"ApplicationsTrafficRuleGET",
		strings.ToUpper("Get"),
		"/mec_app_support/v1/applications/{appInstanceId}/traffic_rules/{trafficRuleId}",
		appSupport.ApplicationsTrafficRuleGET,
	},

	Route{
		"ApplicationsTrafficRulePUT",
		strings.ToUpper("Put"),
		"/mec_app_support/v1/applications/{appInstanceId}/traffic_rules/{trafficRuleId}",
		appSupport.ApplicationsTrafficRulePUT,
	},

	Route{
		"ApplicationsTrafficRulesGET",
		strings.ToUpper("Get"),
		"/mec_app_support/v1/applications/{appInstanceId}/traffic_rules",
		appSupport.ApplicationsTrafficRulesGET,
	},

	Route{
		"TimingCapsGET",
		strings.ToUpper("Get"),
		"/mec_app_support/v1/timing/timing_caps",
		appSupport.TimingCapsGET,
	},

	Route{
		"TimingCurrentTimeGET",
		strings.ToUpper("Get"),
		"/mec_app_support/v1/timing/current_time",
		appSupport.TimingCurrentTimeGET,
	},

	Route{
		"Index",
		"GET",
		"/mec_service_mgmt/v1/",
		Index,
	},

	Route{
		"AppServicesGET",
		strings.ToUpper("Get"),
		"/mec_service_mgmt/v1/applications/{appInstanceId}/services",
		svcMgmt.AppServicesGET,
	},

	Route{
		"AppServicesPOST",
		strings.ToUpper("Post"),
		"/mec_service_mgmt/v1/applications/{appInstanceId}/services",
		svcMgmt.AppServicesPOST,
	},

	Route{
		"AppServicesServiceIdDELETE",
		strings.ToUpper("Delete"),
		"/mec_service_mgmt/v1/applications/{appInstanceId}/services/{serviceId}",
		svcMgmt.AppServicesServiceIdDELETE,
	},

	Route{
		"AppServicesServiceIdGET",
		strings.ToUpper("Get"),
		"/mec_service_mgmt/v1/applications/{appInstanceId}/services/{serviceId}",
		svcMgmt.AppServicesServiceIdGET,
	},

	Route{
		"AppServicesServiceIdPUT",
		strings.ToUpper("Put"),
		"/mec_service_mgmt/v1/applications/{appInstanceId}/services/{serviceId}",
		svcMgmt.AppServicesServiceIdPUT,
	},

	Route{
		"ApplicationsSubscriptionDELETE",
		strings.ToUpper("Delete"),
		"/mec_service_mgmt/v1/applications/{appInstanceId}/subscriptions/{subscriptionId}",
		svcMgmt.ApplicationsSubscriptionDELETE,
	},

	Route{
		"ApplicationsSubscriptionGET",
		strings.ToUpper("Get"),
		"/mec_service_mgmt/v1/applications/{appInstanceId}/subscriptions/{subscriptionId}",
		svcMgmt.ApplicationsSubscriptionGET,
	},

	Route{
		"ApplicationsSubscriptionsGET",
		strings.ToUpper("Get"),
		"/mec_service_mgmt/v1/applications/{appInstanceId}/subscriptions",
		svcMgmt.ApplicationsSubscriptionsGET,
	},

	Route{
		"ApplicationsSubscriptionsPOST",
		strings.ToUpper("Post"),
		"/mec_service_mgmt/v1/applications/{appInstanceId}/subscriptions",
		svcMgmt.ApplicationsSubscriptionsPOST,
	},

	Route{
		"ServicesGET",
		strings.ToUpper("Get"),
		"/mec_service_mgmt/v1/services",
		svcMgmt.ServicesGET,
	},

	Route{
		"ServicesServiceIdGET",
		strings.ToUpper("Get"),
		"/mec_service_mgmt/v1/services/{serviceId}",
		svcMgmt.ServicesServiceIdGET,
	},

	Route{
		"TransportsGET",
		strings.ToUpper("Get"),
		"/mec_service_mgmt/v1/transports",
		svcMgmt.TransportsGET,
	},

	Route{
		"Index",
		"GET",
		"/app_info/v1/",
		Index,
	},
}
