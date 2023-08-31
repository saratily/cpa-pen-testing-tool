import { useContext, useState, useEffect, useCallback } from "react";

import AuthContext from "../../store/auth-context";
import Errors from "../Errors/Errors";
import PenetrationForm from "./PenetrationForm";
import PenetrationsList from "./PenetrationsLists";

const Penetrations = () => {
  const authContext = useContext(AuthContext);
  const [penetrations, setPenetrations] = useState([]);
  const [errors, setErrors] = useState({});

  const fetchPenetrationsHandler = useCallback(async () => {
    setErrors({});

    try {
      const response = await fetch('/api/penetrations',
        {
          headers: {
            'Authorization': 'Bearer ' + authContext.token,
          },
        }
      );
      const data = await response.json();
      if (!response.ok) {
        let errorText = 'Fetching penetrations failed.';
        if (!data.hasOwnProperty('error')) {
          throw new Error(errorText);
        }
        if ((typeof data['error'] === 'string')) {
          setErrors({ 'unknown': data['error'] })
        } else {
          setErrors(data['error']);
        }
      } else {
        setPenetrations(data.data);
      }
    } catch (error) {
      setErrors({ "error": error.message });
    }
  }, [authContext.token]);

  useEffect(() => {
    fetchPenetrationsHandler();
  }, [fetchPenetrationsHandler]);

  const addPenetrationHandler = (penetrationData) => {
    setPenetrations((prevState) => { return [...prevState, penetrationData] });
  }

  const deletePenetrationHandler = (penetrationID) => {
    setPenetrations((prevState) => {
      return prevState.filter(penetration => { return penetration.ID !== penetrationID; })
    })
  }

  const editPenetrationHandler = () => {
    fetchPenetrationsHandler();
  }

  const penetrationsWebsite = penetrations.length === 0 ?
    <p>No penetrations yet</p>
    :
    <PenetrationsList
      penetrations={penetrations}
      onEditPenetration={editPenetrationHandler}
      onDeletePenetration={deletePenetrationHandler} />;

  const errorWebsite = Object.keys(errors).length === 0 ? null : Errors(errors);

  return (
    <section>
      <h1 className="pb-4">Penetration Tests</h1>
      <PenetrationForm onAddPenetration={addPenetrationHandler}/>
      {errorWebsite}
      {penetrationsWebsite}
    </section>
  );
};

export default Penetrations;
