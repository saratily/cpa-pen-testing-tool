import { useState, useContext, useEffect, useCallback } from 'react';

import AuthContext from '../../store/auth-context';
import Errors from '../Errors/Errors';

const PenetrationForm = (props) => {
  const authContext = useContext(AuthContext);

  const [titleValue, setTitleValue] = useState('');
  const [websiteValue, setWebsiteValue] = useState('');

  const [errors, setErrors] = useState({});

  const populateFields = useCallback(() => {
    if (props.penetration) {
      setTitleValue(props.penetration.Title);
      setWebsiteValue(props.penetration.Website);
    }
  }, [props.penetration]);

  useEffect(() => {
    populateFields();
  }, [populateFields]);

  async function submitHandler(event) {
    event.preventDefault();
    setErrors({});

    try {
      const method = props.onEditPenetration ? 'PUT' : 'POST';
      let body = {
        Title: titleValue,
        Website: websiteValue,
      }
      if (props.onEditPenetration) {
        body.ID = props.penetration.ID;
      }
      const response = await fetch('api/penetrations',
        {
          method: method,
          body: JSON.stringify(body),
          headers: {
            'Content-Type': 'application/json',
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
        setTitleValue('');
        setWebsiteValue('');
        if (props.onAddPenetration) {
          props.onAddPenetration(data.data);
        }
        if (props.onEditPenetration) {
          props.onEditPenetration(data.data);
        }
      }
    } catch (error) {
      setErrors({ "error": error.message });
    }
  };

  const titleChangeHandler = (event) => { setTitleValue(event.target.value) }
  const websiteChangeHandler = (event) => { setWebsiteValue(event.target.value) }

  const errorWebsite = Object.keys(errors).length === 0 ? null : Errors(errors);
  const submitButtonText = props.onEditPenetration ? 'Update Penetration' : 'Add Penetration';

  return (
    <section>
      <div className="container w-75 pb-4">
        <form onSubmit={submitHandler}>
          <div className="form-group pb-3">
            <label htmlFor="title">Title</label>
            <input id="title" type="text" className="form-control" required value={titleValue} onChange={titleChangeHandler}></input>
          </div>
          <div className="form-group pb-3">
            <label htmlFor="website">Website</label>
            <textarea id="website" className="form-control" rows="5" required value={websiteValue} onChange={websiteChangeHandler}></textarea>
          </div>
          <button type="submit" className="btn btn-success">{submitButtonText}</button>
        </form>
        {errorWebsite}
      </div>
    </section>
  );
}

export default PenetrationForm;
