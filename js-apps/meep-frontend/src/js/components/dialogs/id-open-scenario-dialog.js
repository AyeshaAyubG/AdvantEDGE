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

import React, { Component } from 'react';

import IDDialog from './id-dialog';
import IDSelect from '../helper-components/id-select';
import {
  MEEP_DLG_OPEN_SCENARIO,
  MEEP_DLG_OPEN_SCENARIO_SELECT
} from '../../meep-constants';

class IDOpenScenarioDialog extends Component {
  constructor(props) {
    super(props);
    this.state = {
      selectedScenario: null
    };
  }

  render() {
    return (
      <IDDialog
        title={this.props.title}
        open={this.props.open}
        onClose={this.props.onClose}
        onSubmit={() => {
          this.props.api.getScenario(
            this.state.selectedScenario,
            this.props.getScenarioLoadCb
          );
        }}
        cydata={MEEP_DLG_OPEN_SCENARIO}
      >
        <IDSelect
          isDialog={true}
          label={this.props.label || 'Scenario'}
          value={this.props.value}
          options={this.props.options}
          onChange={e => {
            this.setState({ selectedScenario: e.target.value });
          }}
          cydata={MEEP_DLG_OPEN_SCENARIO_SELECT}
        />
      </IDDialog>
    );
  }
}

export default IDOpenScenarioDialog;
