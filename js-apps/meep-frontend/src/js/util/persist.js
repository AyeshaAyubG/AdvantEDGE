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

// Will persist the app state between browser refresh
// import * as stringify from 'json-stringify-safe';
// import {version} from './package.json';

// /*global __VERSION__*/
// const STATE_KEY = 'IDCC-meep-frontend' + (__VERSION__ ? ':' + __VERSION__ : '') + ':state';
const STATE_KEY = 'IDCC-meep-frontend:state';

export function saveState(state) {
  try {
    let serializedState = JSON.stringify(state);
    localStorage.setItem(STATE_KEY, serializedState);
  } catch (e) {
    // TODO: consider showing an alert.
    // console.log('Error while saving app state: ', e);
  }
}

export function loadState() {
  try {
    let serializedState = localStorage.getItem(STATE_KEY);

    if (serializedState === null) {
      return this.initializeState();
    }

    return JSON.parse(serializedState);
  } catch (err) {
    return null;
  }
}
