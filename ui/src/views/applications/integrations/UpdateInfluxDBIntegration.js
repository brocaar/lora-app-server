import React, {Component} from "react";
import Grid from '@material-ui/core/Grid';
import Card from '@material-ui/core/Card';
import CardHeader from "@material-ui/core/CardHeader";
import CardContent from "@material-ui/core/CardContent";

import ApplicationStore from "../../../stores/ApplicationStore";
import InfluxDBIntegrationForm from "./InfluxDBIntegrationForm";
import InfluxDBV2IntegrationForm from "./InfluxDBV2IntegrationForm";


class UpdateInfluxDBIntegration extends Component {
  constructor() {
    super();
    this.state = {};
  }

  onSubmit = (integration) => {
    let integr = integration;
    integr.applicationID = this.props.match.params.applicationID;

    ApplicationStore.updateInfluxDBIntegration(integr, resp => {
      this.props.history.push(`/organizations/${this.props.match.params.organizationID}/applications/${this.props.match.params.applicationID}/integrations`);
    }, this.props.V2);
  } 

  componentDidMount() {
    ApplicationStore.getInfluxDBIntegration(this.props.match.params.applicationID, (resp) => {
      this.setState({
        object: resp.integration,
      });
    }, this.props.V2);
  }

  render() {
    if (this.state.object === undefined) {
      return null;
    }

    return(
      <Grid container spacing={4}>
        <Grid item xs={12}>
          <Card>
            <CardHeader title="Update InfluxDB integration" />
            <CardContent>
              {this.props.V2 ?
                  <InfluxDBV2IntegrationForm submitLabel="Update integration" onSubmit={this.onSubmit} object={this.state.object}/>
                  : <InfluxDBIntegrationForm submitLabel="Update integration" onSubmit={this.onSubmit} object={this.state.object}/>
              }
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    );
  }
}

export default UpdateInfluxDBIntegration;
