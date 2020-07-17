import React from "react";
import { CircularProgress, IconButton, Tooltip } from "@material-ui/core";

const STATUS = ["代码克隆中", "代码克隆完成", "代码克隆失败"];

class RepositoryStatus extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }
  renderCircularProgress(status) {
    if (0 === status) return <CircularProgress size={20} />;
    if (1 === status)
      return <CircularProgress variant="static" value={100} size={20} />;
    return (
      <CircularProgress
        variant="static"
        value={100}
        size={20}
        color="secondary"
      />
    );
  }
  render() {
    return (
      <div>
        <Tooltip title={STATUS[this.props.status]}>
          <IconButton color="primary">
            {this.renderCircularProgress(this.props.status)}
          </IconButton>
        </Tooltip>
      </div>
    );
  }
}

export default RepositoryStatus;
