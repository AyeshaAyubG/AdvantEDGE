/*
 * Copyright (c) 2019  InterDigital Communications, Inc
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

// Import MEEP Contstants
import * as meep from '../../../../js-apps/meep-frontend/src/js/meep-constants';

// Import element utils
import {
  // Field Names
  FIELD_PARENT,
  FIELD_NAME,
  FIELD_INT_DOM_LATENCY,
  FIELD_INT_DOM_LATENCY_VAR,
  FIELD_INT_DOM_LATENCY_DIST,
  FIELD_INT_DOM_THROUGHPUT_DL,
  FIELD_INT_DOM_THROUGHPUT_UL,
  FIELD_INT_DOM_PKT_LOSS,
  FIELD_INT_ZONE_LATENCY,
  FIELD_INT_ZONE_LATENCY_VAR,
  FIELD_INT_ZONE_THROUGHPUT_DL,
  FIELD_INT_ZONE_THROUGHPUT_UL,
  FIELD_INT_ZONE_PKT_LOSS,
  FIELD_INTRA_ZONE_LATENCY,
  FIELD_INTRA_ZONE_LATENCY_VAR,
  FIELD_INTRA_ZONE_THROUGHPUT_DL,
  FIELD_INTRA_ZONE_THROUGHPUT_UL,
  FIELD_INTRA_ZONE_PKT_LOSS,
  FIELD_TERM_LINK_LATENCY,
  FIELD_TERM_LINK_LATENCY_VAR,
  FIELD_TERM_LINK_THROUGHPUT_DL,
  FIELD_TERM_LINK_THROUGHPUT_UL,
  FIELD_TERM_LINK_PKT_LOSS,

  getElemFieldVal,
} from '../../../../js-apps/meep-frontend/src/js/util/elem-utils';

// Import Test utility functions
import { selector, click, type, select, verify, verifyEnabled, verifyForm } from '../util/util';

// Scenario Execution Tests
describe('Scenario Execution', function () {

  // Test Variables
  let defaultScenario = 'None';
  let sandbox = 'sbox-test';
  let scenario = 'demo1'

  // ------------------------------
  //            TESTS
  // ------------------------------

  // It is not possible to connect to multiple domains in one test case, but it is possible in different test cases
  // Known cypress bug : https://github.com/cypress-io/cypress/issues/944
  // The suggested solution from cypress is not working, and will likely not be fixed
  // Best workaround: using different tests

  // Demo1 scenario testing for mobility event on the demo app output
  it('Deploy DEMO1 scenario - mobility event #1', function () {
    openDefaultMeepUrl();

    // Create Sandbox
    createSandbox(sandbox);

    // Deploy demo scenario
    deployScenario(scenario);

    // Test events
    //testMobilityEvent();
    //testMobility event was split so that each mobility event can be tested independantly
    cy.log('Create 1st Mobility event');
    createMobilityEvent('ue1', 'zone1-poa2');
  });

  it('DEMO1 scenario - validate mobility event in demo-app frontend', function () {
    cy.wait(2000);
    openDemoUrl();

    // Already deployed scenario

    // Test events
    validateDemoAppFrontEndUpdateOnMobilityEvent('ue1', 'zone1-poa2');
  });

  // Demo1 scenario testing  mobility event #2
  it('DEMO1 scenario - mobility event #2', function () {
    openDefaultMeepUrl();
    click(meep.MEEP_TAB_EXEC);
    cy.wait(1000);
    select(meep.EXEC_SELECT_SANDBOX, sandbox);
    cy.wait(3000);

    // Test events
    cy.log('Create 2nd Mobility event');
    createMobilityEvent('ue1', 'zone2-poa1');
  });

  it('DEMO1 scenario - validate mobility event in demo-app frontend', function () {
    cy.wait(2000);
    openDemoUrl();

    // Test events
    validateDemoAppFrontEndUpdateOnMobilityEvent('ue1', 'zone2-poa1');
  });

  // Demo1 scenario testing  mobility event #3
  it('DEMO1 scenario - mobility event #3', function () {
    openDefaultMeepUrl();
    click(meep.MEEP_TAB_EXEC);
    cy.wait(1000);
    select(meep.EXEC_SELECT_SANDBOX, sandbox);
    cy.wait(3000);

    // Test events
    cy.log('Create 3rd Mobility event');
    createMobilityEvent('ue1', 'zone1-poa1');
  });

  it('DEMO1 scenario - validate mobility event in demo-app frontend', function () {
    cy.wait(2000);
    openDemoUrl();

    // Test events
    validateDemoAppFrontEndUpdateOnMobilityEvent('ue1', 'zone1-poa1');
  });

  // Demo1 scenario testing  mobility event #4
  it('DEMO1 scenario - mobility event #4', function () {
    openDefaultMeepUrl();
    click(meep.MEEP_TAB_EXEC);
    cy.wait(1000);
    select(meep.EXEC_SELECT_SANDBOX, sandbox);
    cy.wait(3000);

    // Test events
    cy.log('Create 4th Mobility event');
    createMobilityEvent('ue2-ext', 'zone1-poa2');
  });

  it('DEMO1 scenario - validate mobility event in demo-app frontend', function () {
    cy.wait(2000);
    openDemoUrl();

    // Test events
    validateDemoAppFrontEndUpdateOnMobilityEvent('ue2-ext', 'zone1-poa2');
  });

  // Demo1 scenario testing  mobility event #5
  it('DEMO1 scenario - mobility event #5', function () {
    openDefaultMeepUrl();
    click(meep.MEEP_TAB_EXEC);
    cy.wait(1000);
    select(meep.EXEC_SELECT_SANDBOX, sandbox);
    cy.wait(3000);

    // Test events
    cy.log('Create 5th Mobility event');
    createMobilityEvent('ue2-ext', 'zone2-poa1');
  });

  it('DEMO1 scenario - validate mobility event in demo-app frontend', function () {
    cy.wait(2000);
    openDemoUrl();

    // Test events
    validateDemoAppFrontEndUpdateOnMobilityEvent('ue2-ext', 'zone2-poa1');
  });

  // Demo1 scenario testing  mobility event #6
  it('DEMO1 scenario - mobility event #6', function () {
    openDefaultMeepUrl();
    click(meep.MEEP_TAB_EXEC);
    cy.wait(1000);
    select(meep.EXEC_SELECT_SANDBOX, sandbox);
    cy.wait(3000);

    // Test events
    cy.log('Create 6th Mobility event');
    createMobilityEvent('ue2-ext', 'zone1-poa1');
  });

  it('DEMO1 scenario - validate mobility event in demo-app frontend', function () {
    cy.wait(2000);
    openDemoUrl();

    // Test events
    validateDemoAppFrontEndUpdateOnMobilityEvent('ue2-ext', 'zone1-poa1');
  });

  // Demo1 scenario terminate
  it('DEMO1 scenario - terminate', function () {
    openDefaultMeepUrl();
    click(meep.MEEP_TAB_EXEC);
    cy.wait(1000);
    select(meep.EXEC_SELECT_SANDBOX, sandbox);
    cy.wait(3000);

    // Test events
    // Terminate demo scenario
    terminateScenario(scenario);

    // Destroy Sandbox
    destroySandbox(sandbox);
  });

  // ------------------------------
  //          FUNCTIONS
  // ------------------------------

  // Opens the default meepUrl page
  function openDefaultMeepUrl() {
    var meepUrl = Cypress.env('meep_url');
    if (meepUrl == null) {
      meepUrl = 'http://127.0.0.1';
    }

    cy.viewport(1920, 1080);
    cy.visit(meepUrl);
    cy.wait(1000);
  }

  // Opens the demo frontend page
  function openDemoUrl() {
    var demoUrl = Cypress.env('demo_url');
    if (demoUrl == null) {
      demoUrl = 'http://127.0.0.1:31111';
    }

    cy.viewport(1920, 1080);
    cy.visit(demoUrl);
    cy.wait(1000);
  }

  // Create sandbox with provided name
  function createSandbox(name) {
    // Go to execution page
    cy.log('Go to execution page');
    click(meep.MEEP_TAB_EXEC);
    cy.wait(1000);
    verify(meep.MEEP_LBL_SCENARIO_NAME, 'contain', defaultScenario);
    verifyEnabled(meep.EXEC_BTN_NEW_SANDBOX, true);
    verifyEnabled(meep.EXEC_BTN_DELETE_SANDBOX, false);
    verifyEnabled(meep.EXEC_BTN_DEPLOY, false);
    verifyEnabled(meep.EXEC_BTN_TERMINATE, false);
    verifyEnabled(meep.EXEC_BTN_EVENT, false);

    // Create sandbox
    cy.log('Create sandbox: ' + name);
    click(meep.EXEC_BTN_NEW_SANDBOX);
    type(meep.MEEP_DLG_NEW_SANDBOX_NAME, name);
    click(meep.MEEP_DLG_NEW_SANDBOX, 'Ok');
    cy.wait(15000);
    verifyEnabled(meep.EXEC_BTN_NEW_SANDBOX, true);
    verifyEnabled(meep.EXEC_BTN_DELETE_SANDBOX, true);
    verifyEnabled(meep.EXEC_BTN_DEPLOY, true);
    verifyEnabled(meep.EXEC_BTN_TERMINATE, false);
    verifyEnabled(meep.EXEC_BTN_EVENT, false);
  }

  // Destroy sandbox with provided name
  function destroySandbox(name) {
    cy.log('Destroy Sandbox: ' + name);
    select(meep.EXEC_SELECT_SANDBOX, name);
    cy.wait(1000);
    click(meep.EXEC_BTN_DELETE_SANDBOX);
    click(meep.MEEP_DLG_DELETE_SANDBOX, 'Ok');
    cy.wait(10000);
    verifyEnabled(meep.EXEC_BTN_NEW_SANDBOX, true);
    verifyEnabled(meep.EXEC_BTN_DELETE_SANDBOX, false);
    verifyEnabled(meep.EXEC_BTN_DEPLOY, false);
    verifyEnabled(meep.EXEC_BTN_TERMINATE, false);
    verifyEnabled(meep.EXEC_BTN_EVENT, false);
    verify(meep.MEEP_LBL_SCENARIO_NAME, 'contain', defaultScenario);
  }

  // Deploy scenario with provided name
  function deployScenario(name) {
    // Go to execution page
    cy.log('Go to execution page');
    click(meep.MEEP_TAB_EXEC);
    cy.wait(1000);
    verify(meep.MEEP_LBL_SCENARIO_NAME, 'contain', defaultScenario);
    verifyEnabled(meep.EXEC_BTN_DEPLOY, true);
    verifyEnabled(meep.EXEC_BTN_TERMINATE, false);
    verifyEnabled(meep.EXEC_BTN_EVENT, false);

    // Deploy scenario
    cy.log('Deploy scenario: ' + name);
    click(meep.EXEC_BTN_DEPLOY);
    cy.wait(1000);
    select(meep.MEEP_DLG_DEPLOY_SCENARIO_SELECT, name);
    click(meep.MEEP_DLG_DEPLOY_SCENARIO, 'Ok');
    cy.wait(15000);
    verifyEnabled(meep.EXEC_BTN_EVENT, true, 30000);
    verifyEnabled(meep.EXEC_BTN_DEPLOY, false);
    verifyEnabled(meep.EXEC_BTN_TERMINATE, true);
    verify(meep.MEEP_LBL_SCENARIO_NAME, 'contain', name);
  }

  // Terminate deployed scenario
  function terminateScenario(name) {
    cy.log('Terminate Scenario: ' + name);
    click(meep.EXEC_BTN_TERMINATE);
    click(meep.MEEP_DLG_TERMINATE_SCENARIO, 'Ok');
    cy.wait(10000);
    verifyEnabled(meep.EXEC_BTN_DEPLOY, true, 120000);
    verifyEnabled(meep.EXEC_BTN_TERMINATE, false);
    verifyEnabled(meep.EXEC_BTN_EVENT, false);
    verify(meep.MEEP_LBL_SCENARIO_NAME, 'contain', defaultScenario);
  }

  // Cancel Event creation
  function testCancelEvent() {
    cy.log('Cancel event creation');
    click(meep.EXEC_BTN_EVENT);
    click(meep.EXEC_BTN_EVENT_BTN_MANUAL_REPLAY);
    verifyForm(meep.EXEC_EVT_TYPE, true);
    verifyEnabled(meep.MEEP_BTN_CANCEL, true);
    // verifyEnabled(meep.MEEP_BTN_APPLY, false)
    click(meep.MEEP_BTN_CANCEL);
    cy.wait(1000);
  }

  // Create Network Characteristic events
  function testNetCharEvent(scenario) {
    cy.log('Create & Validate Network Characteristic event');
    createNetCharEvent('SCENARIO', scenario, 60, 5, 1, 200000, 'Pareto');
    createNetCharEvent('DOMAIN', 'operator1', 10, 3, 2, 90000, '');
    createNetCharEvent('ZONE', 'zone1', 6, 2, 1, 70000, '');
    createNetCharEvent('ZONE', 'zone2', 6, 2, 1, 70000, '');
    createNetCharEvent('POA', 'zone1-poa1', 2, 3, 4, 10000, '');
    createNetCharEvent('POA', 'zone1-poa2', 40, 5, 2, 20000, '');
    createNetCharEvent('POA', 'zone2-poa1', 0, 0, 1, 15000, '');
  }

  // Create a Mobility event
  function createMobilityEvent(elem, dest) {
    cy.log('Moving ' + elem + ' --> ' + dest);
    click(meep.EXEC_BTN_EVENT);
    click(meep.EXEC_BTN_EVENT_BTN_MANUAL_REPLAY);
    select(meep.EXEC_EVT_TYPE, meep.MOBILITY_EVENT);
    select(meep.EXEC_EVT_MOB_TARGET, elem);
    select(meep.EXEC_EVT_MOB_DEST, dest);
    click(meep.MEEP_BTN_APPLY);

    // Validate event
    cy.wait(1000);
    validateMobilityEvent(elem, dest);
  }

  // Create a Network Characteristic event
  function createNetCharEvent(elemType, name, l, lv, pl, tp, dist) {
    cy.log('Setting Net Char for type[' + elemType + '] name[' + name + '] latency[' + l +
      '] variation[' + lv + '] packetLoss[' + pl + '] throughput[' + tp + ']');
    click(meep.EXEC_BTN_EVENT);
    click(meep.EXEC_BTN_EVENT_BTN_MANUAL_REPLAY);
    select(meep.EXEC_EVT_TYPE, meep.NETWORK_CHARACTERISTICS_EVENT);
    select(meep.EXEC_EVT_NC_TYPE, elemType);
    select(meep.EXEC_EVT_NC_NAME, name);
    cy.wait(1000);
    type(meep.CFG_ELEM_LATENCY, l);
    type(meep.CFG_ELEM_LATENCY_VAR, lv);
    type(meep.CFG_ELEM_PKT_LOSS, pl);
    type(meep.CFG_ELEM_THROUGHPUT_DL, tp);
    type(meep.CFG_ELEM_THROUGHPUT_UL, tp-1);
    if (dist != '') {
      select(meep.CFG_ELEM_LATENCY_DIST, dist);
    }
    click(meep.MEEP_BTN_APPLY);

    // Validate event
    cy.wait(1000);
    validateNetCharEvent(elemType, name, l, lv, pl, tp, tp-1, dist);
  }

  // Retrieve Element entry from Application table
  function getEntry(entries, name) {
    if (entries) {
      return entries[name] ? entries[name] : null;
    }
    return null;
  }

  // Validate that new UE parent matches destination
  function validateMobilityEvent(elem, dest) {
    cy.window().then((win) => {
      var entry = getEntry(win.meepStore.getState().exec.table.entries, elem);
      assert.isNotNull(entry);
      assert.equal(getElemFieldVal(entry, FIELD_PARENT), dest);
    });
  }

  // Validate that network characteristics were correctly applied
  function validateNetCharEvent(elemType, name, l, lv, pl, tpDl, tpUl, dist) {
    cy.window().then((win) => {
      var entry = getEntry(win.meepStore.getState().exec.table.entries, name);
      assert.isNotNull(entry);

      switch (elemType) {
        case 'SCENARIO':
          assert.equal(getElemFieldVal(entry, FIELD_INT_DOM_LATENCY), l);
          assert.equal(getElemFieldVal(entry, FIELD_INT_DOM_LATENCY_VAR), lv);
          assert.equal(getElemFieldVal(entry, FIELD_INT_DOM_LATENCY_DIST), dist);
          assert.equal(getElemFieldVal(entry, FIELD_INT_DOM_PKT_LOSS), pl);
          assert.equal(getElemFieldVal(entry, FIELD_INT_DOM_THROUGHPUT_DL), tpDl);
          assert.equal(getElemFieldVal(entry, FIELD_INT_DOM_THROUGHPUT_UL), tpUl);
          break;
        case 'DOMAIN':
          assert.equal(getElemFieldVal(entry, FIELD_INT_ZONE_LATENCY), l);
          assert.equal(getElemFieldVal(entry, FIELD_INT_ZONE_LATENCY_VAR), lv);
          assert.equal(getElemFieldVal(entry, FIELD_INT_ZONE_PKT_LOSS), pl);
          assert.equal(getElemFieldVal(entry, FIELD_INT_ZONE_THROUGHPUT_DL), tpDl);
          assert.equal(getElemFieldVal(entry, FIELD_INT_ZONE_THROUGHPUT_UL), tpUl);
          break;
        case 'ZONE':
          assert.equal(getElemFieldVal(entry, FIELD_INTRA_ZONE_LATENCY), l);
          assert.equal(getElemFieldVal(entry, FIELD_INTRA_ZONE_LATENCY_VAR), lv);
          assert.equal(getElemFieldVal(entry, FIELD_INTRA_ZONE_PKT_LOSS), pl);
          assert.equal(getElemFieldVal(entry, FIELD_INTRA_ZONE_THROUGHPUT_DL), tpDl);
          assert.equal(getElemFieldVal(entry, FIELD_INTRA_ZONE_THROUGHPUT_UL), tpUl);
          break;
        case 'POA':
          assert.equal(getElemFieldVal(entry, FIELD_TERM_LINK_LATENCY), l);
          assert.equal(getElemFieldVal(entry, FIELD_TERM_LINK_LATENCY_VAR), lv);
          assert.equal(getElemFieldVal(entry, FIELD_TERM_LINK_PKT_LOSS), pl);
          assert.equal(getElemFieldVal(entry, FIELD_TERM_LINK_THROUGHPUT_DL), tpDl);
          assert.equal(getElemFieldVal(entry, FIELD_TERM_LINK_THROUGHPUT_UL), tpUl);
          break;
        default:
          assert.isOk(false, 'Unsupported element type');
      }
    });
  }

  // Validate that demo app front-end update matches the mobility event that occured
  function validateDemoAppFrontEndUpdateOnMobilityEvent(ue, dest) {

    var ueAppName;
    var ueAppLocation;
    var ueAddressLabel;
    var ueLocationLabel;

    if (ue == 'ue1') {
      ueAppName = 'ue1-iperf';
      ueAddressLabel = 'span[id="demo-svc-loc-serv-address-1"]';
      ueLocationLabel = 'span[id="demo-svc-loc-serv-location-1"]';
    } else if (ue == 'ue2-ext') {
      ueAppName = 'ue2-svc';
      ueAddressLabel = 'span[id="demo-svc-loc-serv-address-2"]';
      ueLocationLabel = 'span[id="demo-svc-loc-serv-location-2"]';
    }

    if (dest == 'zone1-poa1') {
      ueAppLocation = 'zone1 / zone1-poa1';
    } else if (dest == 'zone1-poa2') {
      ueAppLocation = 'zone1 / zone1-poa2';
    } else if (dest == 'zone2-poa1') {
      ueAppLocation = 'zone2 / zone2-poa1';
    }
    cy.contains(ueLocationLabel, ueAppLocation)
  }

});


