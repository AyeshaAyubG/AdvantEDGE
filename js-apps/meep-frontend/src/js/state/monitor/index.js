/*
 * Copyright (c) 2022  The AdvantEDGE Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

const initialState = {
  currentDashboard: '',
  dashboardOptions: [],
  editedDashboardOptions: null
};

const ADD_DASHBOARD_OPTION = 'ADD_DASHBOARD_OPTION';
export const addDashboardOption = option => {
  return {
    type: ADD_DASHBOARD_OPTION,
    payload: option
  };
};

const CHANGE_DASHBOARD = 'CHANGE_DASHBOARD';
export function changeDashboard(label) {
  return {
    type: CHANGE_DASHBOARD,
    payload: label
  };
}

const CHANGE_EDITED_DASHBOARD_OPTIONS = 'CHANGE_EDITED_DASHBOARD_OPTIONS';
export function changeEditedDashboardOptions(mode) {
  return {
    type: CHANGE_EDITED_DASHBOARD_OPTIONS,
    payload: mode
  };
}

const CHANGE_DASHBOARD_OPTIONS = 'CHANGE_DASHBOARD_OPTIONS';
export function changeDashboardOptions(mode) {
  return {
    type: CHANGE_DASHBOARD_OPTIONS,
    payload: mode
  };
}

export default function settingsReducer(state = initialState, action) {
  switch (action.type) {
  case CHANGE_DASHBOARD:
    return { ...state, currentDashboard: action.payload };
  case ADD_DASHBOARD_OPTION:
    return {
      ...state,
      dashboardOptions: [...state.dashboardOptions, action.payload]
    };
  case CHANGE_EDITED_DASHBOARD_OPTIONS:
    return { ...state, editedDashboardOptions: action.payload };
  case CHANGE_DASHBOARD_OPTIONS:
    return { ...state, dashboardOptions: action.payload };
  default:
    return state;
  }
}
