import React from 'react';

import {
  Button, TextField, Dialog, DialogActions,
  DialogContent, DialogTitle, MenuItem
} from '@material-ui/core';
import {create} from '../../api/repository'


class Edit extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: '',
      url: '',
      desc: '',
      userName: '',
      password: '',
      dependTools: '',
    };
  }

  componentDidMount() {

  }

  textFieldChange(field, event) {
    this.setState({[field]: event.target.value})
  }

  submitForm() {
    create(this.state).then(r => {
      this.props.createSuccess()
    }).catch(() => {
    })
  }

  render() {
    return (
      <div>
        <Dialog open={this.props.show} onClose={this.props.handleClose} aria-labelledby="form-dialog-title">
          <DialogTitle id="form-dialog-title">编辑仓库</DialogTitle>
          <DialogContent>
            <TextField
              InputLabelProps={{
                shrink: true,
              }}
              required autoFocus value={this.state.name} margin="dense" id="name"
              onChange={this.textFieldChange.bind(this, 'name')} label="仓库名称" type="text" fullWidth
            />
            <TextField
              InputLabelProps={{
                shrink: true,
              }}
              required autoFocus value={this.state.url} margin="dense" id="url"
              onChange={this.textFieldChange.bind(this, 'url')} label="仓库Url" type="text" fullWidth
            />
            <TextField
              InputLabelProps={{
                shrink: true,
              }}
              autoFocus value={this.state.userName} margin="dense" id="userName"
              onChange={this.textFieldChange.bind(this, 'userName')} label="用户名" type="text" fullWidth
            />
            <TextField
              InputLabelProps={{
                shrink: true,
              }}
              autoFocus value={this.state.password} margin="dense" id="password"
              onChange={this.textFieldChange.bind(this, 'password')} label="密码" type="password" fullWidth
            />
            <TextField
              InputLabelProps={{
                shrink: true,
              }}
              select
              required value={this.state.dependTools} margin="dense" id="dependTools"
              onChange={this.textFieldChange.bind(this, 'dependTools')} label="依赖安装工具" type="text" fullWidth
            >
              {
                this.props.dependentSupport.map(r => {
                  return <MenuItem key={r} value={r}>{r}</MenuItem>
                })
              }
            </TextField>
            <TextField
              InputLabelProps={{
                shrink: true,
              }}
              margin="dense" value={this.state.desc} id="desc" multiline rows={6}
              onChange={this.textFieldChange.bind(this, 'desc')} label="仓库描述" type="text" fullWidth
            />
          </DialogContent>
          <DialogActions>
            <Button onClick={this.props.handleClose}>
              关闭
            </Button>
            <Button onClick={this.submitForm.bind(this)} color="primary">
              确认
            </Button>
          </DialogActions>
        </Dialog>
      </div>
    );
  }

}


export default Edit
