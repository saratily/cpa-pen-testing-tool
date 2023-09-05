import { useContext, useState, useEffect, useCallback } from "react";

import AuthContext from "../../store/auth-context";
import Errors from "../Errors/Errors";

const Pens = (props) => {
  const authContext = useContext(AuthContext);
  const [pens, setPens] = useState([]);
  const [errors, setErrors] = useState({});

  const fetchPensHandler = useCallback(async () => {
    setErrors({});

    try {
      const response = await fetch('/api/penetration/'+ props.uuid,
        {
          headers: {
            'Authorization': 'Bearer ' + authContext.token,
          },
        }
      );
      const data = await response.json();

      if (!response.ok) {
        let errorText = 'Fetching pens failed.';
        if (!data.hasOwnProperty('error')) {
          throw new Error(errorText);
        }
        if ((typeof data['error'] === 'string')) {
          setErrors({ 'unknown': data['error'] })
        } else {
          setErrors(data['error']);
        }
      } else {
        setPens(data.data);
      }
    } catch (error) {
      setErrors({ "error": error.message });
    }
  }, [authContext.token]);

  useEffect(() => {
    fetchPensHandler();
  }, [fetchPensHandler]);


  const cardTitle = pens.Title;
  const cardBody = pens.Website;

  const pensContent = pens.length === 0 ?
    <p>No pens yet</p>
    :
    <div>
      <h1 className="pb-4">Company: {cardTitle}</h1>
      <h2 className="pb-4">URL : {cardBody}</h2>
    </div>


  const errorContent = Object.keys(errors).length === 0 ? null : Errors(errors);

  return (
    <section>
      {errorContent}
      {pensContent}
    </section>
  );
};

export default Pens;
