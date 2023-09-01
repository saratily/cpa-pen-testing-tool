import { useState, useContext } from 'react';
import { useNavigate } from 'react-router-dom';

import AuthContext from '../../store/auth-context';
import Errors from '../Errors/Errors';
import PenetrationForm from "./PenetrationForm";

const Penetration = (props) => {
  const [editing, setEditing] = useState(false);
  const [errors, setErrors] = useState({});

  const authContext = useContext(AuthContext);

  const switchModeHandler = () => {
    setEditing((prevState) => !prevState);
    setErrors({});
  };

  const navigate = useNavigate();

  const runScanHandler = () => {
    navigate('/planning/' + props.penetration.uuid);
  };


  async function deleteHandler() {
    try {
      const response = await fetch('api/penetrations/' + props.penetration.uuid,
        {
          method: 'DELETE',
          headers: {
            'Authorization': 'Bearer ' + authContext.token,
          },
        }
      );
      const data = await response.json();
      if (!response.ok) {
        let errorText = 'Failed to add new penetration.';
        if (!data.hasOwnProperty('error')) {
          throw new Error(errorText);
        }
        if ((typeof data['error'] === 'string')) {
          setErrors({ 'unknown': data['error'] })
        } else {
          setErrors(data['error']);
        }
      } else {
        props.onDeletePenetration(props.penetration.uuid);
      }
    } catch (error) {
      setErrors({ "error": error.message });
    }
  };

  const editPenetrationHandler = () => {
    setEditing(false);
    props.onEditPenetration();
  }

  const cardTitle = editing ? 'Edit penetration' : props.penetration.Title;
  const cardBody = editing ? <PenetrationForm penetration={props.penetration} onEditPenetration={editPenetrationHandler} editing={true}/> : props.penetration.Website;
  const cardScanning = props.penetration
  const switchModeButtonText = editing ? 'Cancel' : 'Edit';
  const runScanButtonText = 'Run Scan';
  const cardButtons = editing ?
    <div className="container">
      <button type="button" className="btn btn-danger float-right mx-3" onClick={deleteHandler}>Delete</button>
      <button type="button" className="btn btn-success float-right mx-3" onClick={switchModeHandler}>{switchModeButtonText}</button>
    </div>
    :
    <div className="container">
      <button type="button" className="btn btn-danger float-right mx-3" onClick={deleteHandler}>Delete</button>
      <button type="button" className="btn btn-success float-right mx-3" onClick={runScanHandler}>{runScanButtonText}</button>
      <button type="button" className="btn btn-success float-right mx-3" onClick={switchModeHandler}>{switchModeButtonText}</button>
    </div>
  const errorWebsite = Object.keys(errors).length === 0 ? null : Errors(errors);

  return (
    <div className="card mb-5 pb-2">
      <div className="card-header">{cardTitle}</div>
      <div className="card-body">{cardBody}</div>
      {cardButtons}
      {errorWebsite}
    </div>
  );
};

export default Penetration;
