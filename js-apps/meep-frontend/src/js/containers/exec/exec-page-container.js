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

import _ from 'lodash';
import { connect } from 'react-redux';
import React, { Component } from 'react';
import autoBind from 'react-autobind';
import { Grid, GridCell, GridInner } from '@rmwc/grid';
import { Elevation } from '@rmwc/elevation';
import IDSelect from '../../components/helper-components/id-select';
import DashboardContainer from './dashboard-container';
import EventContainer from './event-container';
import ExecPageSandboxButtons from './exec-page-sandbox-buttons';
import ExecPageScenarioButtons from './exec-page-scenario-buttons';

import HeadlineBar from '../../components/headline-bar';
import EventCreationPane from './event-creation-pane';
import EventAutomationPane from './event-automation-pane';
import EventReplayPane from './event-replay-pane';

import ExecTable from './exec-table';

import IDNewSandboxDialog from '../../components/dialogs/id-new-sandbox-dialog';
import IDDeleteSandboxDialog from '../../components/dialogs/id-delete-sandbox-dialog';
import IDDeployScenarioDialog from '../../components/dialogs/id-deploy-scenario-dialog';
import IDTerminateScenarioDialog from '../../components/dialogs/id-terminate-scenario-dialog';
import IDSaveScenarioDialog from '../../components/dialogs/id-save-scenario-dialog';
import IDSaveReplayDialog from '../../components/dialogs/id-save-replay-dialog';


import { execChangeScenarioList} from '../../state/exec';

import {
  uiChangeCurrentDialog,
  uiExecChangeEventCreationMode,
  uiExecChangeEventAutomationMode,
  uiExecChangeEventReplayMode,
  uiExecChangeDashCfgMode,
  uiExecChangeEventCfgMode,
  uiExecChangeCurrentEvent,
  uiExecChangeShowApps,
  uiExecChangeReplayFilesList
} from '../../state/ui';

import {
  execChangeScenarioState,
  execChangeOkToTerminate
} from '../../state/exec';

import {
  // States
  EXEC_STATE_IDLE,
  PAGE_EXECUTE,
  IDC_DIALOG_NEW_SANDBOX,
  IDC_DIALOG_DELETE_SANDBOX,
  IDC_DIALOG_DEPLOY_SCENARIO,
  IDC_DIALOG_TERMINATE_SCENARIO,
  IDC_DIALOG_SAVE_SCENARIO,
  IDC_DIALOG_SAVE_REPLAY,
  MOBILITY_EVENT,
  NETWORK_CHARACTERISTICS_EVENT,
  EXEC_SELECT_SANDBOX
} from '../../meep-constants';

class ExecPageContainer extends Component {
  constructor(props) {
    super(props);
    autoBind(this);
  }

  /**
   * Callback function to receive the result of the getScenarioList operation.
   * @callback module:api/ScenarioConfigurationApi~getScenarioListCallback
   * @param {String} error Error message, if any.
   * @param {module:model/ScenarioList} data The data returned by the service call.
   */
  getScenarioListDeployCb(error, data) {
    if (error !== null) {
      // TODO: consider showing an alert/toast
      return;
    }

    this.props.changeDeployScenarioList(_.map(data.scenarios, 'name'));
  }

  /**
   * Callback function to receive the result of the activateScenario operation.
   * @callback module:api/ScenarioExecutionApi~activateScenarioCallback
   * @param {String} error Error message, if any.
   */
  activateScenarioCb(error) {
    if (error) {
      // TODO: consider showing an alert/toast
      return;
    }

    this.props.refreshScenario();
  }

  /**
   * Callback function to receive the result of the terminateScenario operation.
   * @callback module:api/ScenarioExecutionApi~terminateScenarioCallback
   * @param {String} error Error message, if any.
   */
  terminateScenarioCb(error) {
    if (error !== null) {
      // TODO consider showing an alert  (i.e. toast)
      return;
    }

    this.props.deleteScenario();
    this.props.changeState(EXEC_STATE_IDLE);
    this.props.execChangeOkToTerminate(false);
  }

  /**
   * Callback function to receive the result of the createScenario operation.
   * @callback module:api/ScenarioConfigurationApi~createScenarioCallback
   * @param {String} error Error message, if any.
   * @param data This operation does not return a value.
   * @param {String} response The complete HTTP response.
   */
  createScenarioCb(/*error, data, response*/) {
    // if (error == null) {
    //   console.log('Scenario successfully created');
    // } else {
    //   console.log('Failed to create scenario');
    // }
    // TODO: consider showing an alert/toast
  }

  /**
   * Callback function to receive the result of the getReplayList operation.
   * @callback module:api/EventReplayApi~getReplayFileListCallback
   * @param {String} error Error message, if any.
   * @param {module:model/ReplayFileList} data The data returned by the service call.
   */
  getReplayFileListCb(error, data) {
    if (error !== null) {
      // TODO: consider showing an alert/toast
      return;
    }
    let replayFiles = data.replayFiles;
    replayFiles.unshift('None');
    this.props.changeReplayFilesList(replayFiles);
  }

  saveScenario(scenarioName) {
    const scenario = this.props.scenario;

    const scenarioCopy = JSON.parse(JSON.stringify(scenario));
    scenarioCopy.name = scenarioName;

    this.props.cfgApi.createScenario(
      scenarioName,
      scenarioCopy,
      (error, data, response) => this.createScenarioCb(error, data, response)
    );
  }

  updateReplayFileList() {
    this.props.replayApi.getReplayFileList((error, data, response) => {
      this.getReplayFileListCb(error, data, response);
    });
  }
  
  saveReplay(state) {
    const scenarioName = this.props.scenario.name;
    var replayInfo = {
      scenarioName: '',
      description: ''
    };

    replayInfo.scenarioName = scenarioName;
    replayInfo.description = state.description;

    this.props.replayApi.createReplayFileFromScenarioExec(state.replayName, replayInfo, (error) => {
      if (error) {
        // TODO consider showing an alert
        // console.log(error);
      }
      // Refresh file list
      this.updateReplayFileList();
    });
  }

  // CLOSE DIALOG
  closeDialog() {
    this.props.changeCurrentDialog(Math.random());
  }

  // DEPLOY DIALOG
  onDeployScenario() {
    // Retrieve list of available scenarios
    this.props.cfgApi.getScenarioList((error, data, response) => {
      this.getScenarioListDeployCb(error, data, response);
    });
    this.props.changeCurrentDialog(IDC_DIALOG_DEPLOY_SCENARIO);
  }

  // NEW SANDBOX
  onNewSandbox() {
    this.props.changeCurrentDialog(IDC_DIALOG_NEW_SANDBOX);
  }

  // DELETE SANDBOX
  onDeleteSandbox() {
    this.props.changeCurrentDialog(IDC_DIALOG_DELETE_SANDBOX);
  }

  // SAVE SCENARIO
  onSaveScenario() {
    this.props.changeCurrentDialog(IDC_DIALOG_SAVE_SCENARIO);
  }

  // TERMINATE DIALOG
  onTerminateScenario() {
    this.props.changeCurrentDialog(IDC_DIALOG_TERMINATE_SCENARIO);
  }

  // SAVE REPLAY FILE
  onSaveReplay() {
    this.props.changeCurrentDialog(IDC_DIALOG_SAVE_REPLAY);
  }

  // SHOW REPLAY PANE
  onShowReplay() {
    this.updateReplayFileList();
  }

  // CLOSE CREATE EVENT PANE
  onQuitEventCreationMode() {
    this.props.changeEventCreationMode(false);
  }

  // CLOSE EVENT AUTOMATION PANE
  onQuitEventAutomationMode() {
    this.props.changeEventAutomationMode(false);
  }

  // CLOSE REPLAY EVENT PANE
  onQuitEventReplayMode() {
    this.props.changeEventReplayMode(false);
  }

  // CONFIGURE DASHBOARD
  onOpenDashCfg() {
    this.props.changeDashCfgMode(true);
  }

  // STOP CONFIGURE DASHBOARD
  onCloseDashCfg() {
    this.props.changeDashCfgMode(false);
  }

  // CONFIGURE EVENTS
  onOpenEventCfg() {
    this.props.changeEventCfgMode(true);
  }

  // STOP CONFIGURE EVENTS
  onCloseEventCfg() {
    this.props.changeEventCfgMode(false);
  }

  onSandboxChange(e) {
    this.props.setSandbox(e.target.value);
  }

  // Terminate Active scenario
  terminateScenario() {
    this.props.api.terminateScenario((error, data, response) =>
      this.terminateScenarioCb(error, data, response)
    );
  }
  
  showApps(show) {
    this.props.changeShowApps(show);
    // _.defer(() => {
    //   this.props.execVis.network.setData(this.props.execVisData);
    // });
  }

  renderDialogs() {
    return (
      <>
        <IDNewSandboxDialog
          title="Create New Sandbox"
          open={this.props.currentDialog === IDC_DIALOG_NEW_SANDBOX}
          onClose={this.closeDialog}
          createSandbox={this.props.createSandbox}
        />
        <IDDeleteSandboxDialog
          title="Delete Sandbox"
          open={this.props.currentDialog === IDC_DIALOG_DELETE_SANDBOX}
          onClose={this.closeDialog}
          deleteSandbox={this.props.deleteSandbox}
        />
        <IDDeployScenarioDialog
          title="Open Scenario"
          open={this.props.currentDialog === IDC_DIALOG_DEPLOY_SCENARIO}
          options={this.props.scenarios}
          onClose={this.closeDialog}
          api={this.props.api}
          activateScenarioCb={this.activateScenarioCb}
        />
        <IDSaveScenarioDialog
          title="Save Scenario as ..."
          open={this.props.currentDialog === IDC_DIALOG_SAVE_SCENARIO}
          onClose={this.closeDialog}
          saveScenario={this.saveScenario}
          scenarioNameRequired={true}
        />
        <IDTerminateScenarioDialog
          title="Terminate Scenario"
          open={this.props.currentDialog === IDC_DIALOG_TERMINATE_SCENARIO}
          scenario={this.props.scenario}
          onClose={this.closeDialog}
          onSubmit={this.terminateScenario}
        />
        <IDSaveReplayDialog
          title="Save Events as Replay file"
          open={this.props.currentDialog === IDC_DIALOG_SAVE_REPLAY}
          onClose={this.closeDialog}
          api={this.props.replayApi}
          saveReplay={this.saveReplay}
          replayNameRequired={true}
        />
      </>
    );
  }


  render() {
    if (this.props.page !== PAGE_EXECUTE) {
      return null;
    }

    const sandboxes = (this.props.sandboxes) ? this.props.sandboxes : [];
    sandboxes.sort();
    const sandbox = sandboxes.includes(this.props.sandbox) ? this.props.sandbox : '';

    const scenarioName = (this.props.page === PAGE_EXECUTE) ?
      (this.props.scenarioState !== EXEC_STATE_IDLE) ? this.props.execScenarioName : 'None' :
      this.props.cfgScenarioName;

    const eventPaneOpen = this.props.eventCreationMode || this.props.eventAutomationMode || this.props.eventReplayMode;
    const spanLeft = eventPaneOpen ? 9 : 12;
    const spanRight = eventPaneOpen ? 3 : 0;
    return (
      <div style={{ width: '100%' }}>
        {this.renderDialogs()}

        <div style={{ width: '100%' }}>
          <Grid style={styles.headlineGrid}>
            <GridCell span={12}>
              <Elevation
                className="component-style"
                z={2}
                style={styles.headline}
              >
                <GridInner>
                  <IDSelect
                    label="Sandbox"
                    span={2}
                    options={sandboxes}
                    onChange={this.onSandboxChange}
                    value={sandbox}
                    disabled={false}
                    cydata={EXEC_SELECT_SANDBOX}
                  />
                  <GridCell align={'middle'} span={2}>
                    <ExecPageSandboxButtons
                      sandbox={sandbox}
                      onNewSandbox={this.onNewSandbox}
                      onDeleteSandbox={this.onDeleteSandbox}
                    />
                  </GridCell>
                  <GridCell align={'middle'} style={{ height: '100%'}} span={3}>
                    <GridInner style={{ height: '100%', borderLeft: '2px solid #e4e4e4'}}>
                      <GridCell align={'middle'} style={{ marginLeft: 20}} span={12}>
                        <HeadlineBar
                          titleLabel="Scenario"
                          scenarioName={scenarioName}
                        />
                      </GridCell>
                    </GridInner>
                  </GridCell>
                  <GridCell align={'middle'} span={5}>
                    <GridInner align={'right'}>
                      <GridCell align={'middle'} span={12}>
                        <ExecPageScenarioButtons
                          sandbox={sandbox}
                          onDeploy={this.onDeployScenario}
                          onSaveScenario={this.onSaveScenario}
                          onTerminate={this.onTerminateScenario}
                          onOpenDashCfg={this.onOpenDashCfg}
                          onOpenEventCfg={this.onOpenEventCfg}
                        />
                      </GridCell>
                    </GridInner>
                  </GridCell>
                </GridInner>
              </Elevation>
            </GridCell>
          </Grid>
        </div>

        {this.props.scenarioState !== EXEC_STATE_IDLE && (
          <>
            <Grid style={{ width: '100%' }}>
              <GridCell span={spanLeft}>
                <div>
                  <EventContainer
                    scenarioName={this.props.execScenarioName}
                    eventCfgMode={this.props.eventCfgMode}
                    onCloseEventCfg={this.onCloseEventCfg}
                    onSaveReplay={this.onSaveReplay}
                    onShowReplay={this.onShowReplay}
                    api={this.props.replayApi}
                  />

                  <DashboardContainer
                    sandbox={this.props.sandbox}
                    scenarioName={this.props.execScenarioName}
                    onShowAppsChanged={this.showApps}
                    showApps={this.props.showApps}
                    dashCfgMode={this.props.dashCfgMode}
                    onCloseDashCfg={this.onCloseDashCfg}
                  />
                </div>
              </GridCell>
              <GridCell
                span={spanRight}
                hidden={!eventPaneOpen}
                style={styles.inner}
              >
                <Elevation className="component-style" z={2}>
                  <EventReplayPane
                    api={this.props.replayApi}
                    hide={!this.props.eventReplayMode}
                    onClose={this.onQuitEventReplayMode}
                  />
                </Elevation>
                <Elevation className="component-style" z={2}>
                  <EventCreationPane
                    eventTypes={[MOBILITY_EVENT, NETWORK_CHARACTERISTICS_EVENT]}
                    api={this.props.eventsApi}
                    hide={!this.props.eventCreationMode}
                    onSuccess={this.props.refreshScenario}
                    onClose={this.onQuitEventCreationMode}
                  />
                </Elevation>
                <Elevation className="component-style" z={2}>
                  <EventAutomationPane
                    api={this.props.automationApi}
                    hide={!this.props.eventAutomationMode}
                    onClose={this.onQuitEventAutomationMode}
                  />
                </Elevation>
              </GridCell>
            </Grid>
          </>
        )}
        
        {sandbox && 
          <ExecTable />
        }
      </div>
    );
  }
}

const styles = {
  headlineGrid: {
    marginBottom: 10
  },
  headline: {
    height: 'calc(100% - 20px)',
    padding: 10
  },
  page: {
    height: 1500,
    marginBottom: 10
  }
};

const mapStateToProps = state => {
  return {
    showApps: state.ui.execShowApps,
    // execVis: state.exec.vis,
    configuredElement: state.cfg.elementConfiguration.configuredElement,
    currentDialog: state.ui.currentDialog,
    scenario: state.exec.scenario,
    scenarioState: state.exec.state.scenario,
    scenarios: state.exec.apiResults.scenarios,
    eventCreationMode: state.ui.eventCreationMode,
    eventAutomationMode: state.ui.eventAutomationMode,
    eventReplayMode: state.ui.eventReplayMode,
    dashCfgMode: state.ui.dashCfgMode,
    eventCfgMode: state.ui.eventCfgMode,
    page: state.ui.page,
    execScenarioName: state.exec.scenario.name,
    cfgScenarioName: state.cfg.scenario.name
    // execVisData: execVisFilteredData(state)
  };
};

const mapDispatchToProps = dispatch => {
  return {
    changeCurrentDialog: type => dispatch(uiChangeCurrentDialog(type)),
    changeDeployScenarioList: scenarios => dispatch(execChangeScenarioList(scenarios)),
    changeState: s => dispatch(execChangeScenarioState(s)),
    changeEventCreationMode: val => dispatch(uiExecChangeEventCreationMode(val)), // (true or false)
    changeEventAutomationMode: mode => dispatch(uiExecChangeEventAutomationMode(mode)),
    changeEventReplayMode: val => dispatch(uiExecChangeEventReplayMode(val)), // (true or false)
    changeDashCfgMode: val => dispatch(uiExecChangeDashCfgMode(val)), // (true or false)
    changeEventCfgMode: val => dispatch(uiExecChangeEventCfgMode(val)), // (true or false)
    changeCurrentEvent: e => dispatch(uiExecChangeCurrentEvent(e)),
    execChangeOkToTerminate: ok => dispatch(execChangeOkToTerminate(ok)),
    changeShowApps: show => dispatch(uiExecChangeShowApps(show)),
    changeReplayFilesList: list => dispatch(uiExecChangeReplayFilesList(list))
  };
};

const ConnectedExecPageContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(ExecPageContainer);

export default ConnectedExecPageContainer;
