import { useState, useContext } from 'react';
import { useNavigate  } from 'react-router-dom';

import AuthContext from '../../store/auth-context';
import Errors from '../Errors/Errors';

const Tool = (props) => {
  const navigate = useNavigate()
  const [editing, setEditing] = useState(false);
  const [errors, setErrors] = useState({});

  const authContext = useContext(AuthContext);

  const switchModeHandler = () => {
    setEditing((prevState) => !prevState);
    setErrors({});
  };

  async function toggleHandler() {
 
    try {
      const response = await fetch('/api/toggle/' + props.uuid + '/' + props.type,
        {
          method: 'GET',
          headers: {
            'Authorization': 'Bearer ' + authContext.token,
          },
        }
      );
      const data = await response.json();
      if (!response.ok) {
        let errorText = 'Failed to add new tool.';
        if (!data.hasOwnProperty('error')) {
          throw new Error(errorText);
        }
        if ((typeof data['error'] === 'string')) {
          setErrors({ 'unknown': data['error'] })
        } else {
          setErrors(data['error']);
        }
      } else {
        navigate(0);
      }
    } catch (error) {
      setErrors({ "error": error.message });
    }
  };


  async function execHandler() {
    try {
      const response = await fetch('/api/execute/' + props.uuid + '/' + props.type,
        {
          method: 'GET',
          headers: {
            'Authorization': 'Bearer ' + authContext.token,
          },
        }
      );
      const data = await response.json();
      if (!response.ok) {
        let errorText = 'Failed to add new tool.';
        if (!data.hasOwnProperty('error')) {
          throw new Error(errorText);
        }
        if ((typeof data['error'] === 'string')) {
          setErrors({ 'unknown': data['error'] })
        } else {
          setErrors(data['error']);
        }
      } else {
        navigate(0);
      }
    } catch (error) {
      setErrors({ "error": error.message });
    }
  };

  const cardTitle = editing ? 'Edit tool' : props.tool.Command;
  const cardBody = props.tool.Output;
  const cardButtons = props.tool.Selected == 1 ?
    <div className="container">
      <button type="button" className="btn btn-success float-right mx-3" onClick={toggleHandler}>Selected</button>
      <button type="button" className="btn btn-success float-right mx-3" onClick={execHandler}>Execute command</button>
    </div>
    :
    <div className="container">
      <button type="button" className="btn btn-danger float-right mx-3" onClick={toggleHandler}>Not Selected</button>
      <button type="button" className="btn btn-success float-right mx-3" onClick={execHandler}>Execute command</button>
    </div>
  const errorWebsite = Object.keys(errors).length === 0 ? null : Errors(errors);

  return (
    <div className="card mb-5 pb-2">
      <div className="card-header">{cardTitle}</div>
      {cardButtons}
      <div className="card-body">{cardBody}</div>

      {errorWebsite}
    </div>
  );
};

export default Tool;
