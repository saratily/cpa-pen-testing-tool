import { useContext, useEffect, useCallback } from "react";
import { useLocation, useNavigate } from 'react-router-dom';

import AuthContext from "../store/auth-context";
import Planning from '../components/Planning/Planning';

const PlanningPage = () => {
  const authContext = useContext(AuthContext);
  const navigate = useNavigate();
  const location = useLocation();

  const fetchHandler = useCallback(async () => {
  try {
    const response = await fetch('/api/penetrations',
      {
        headers: {
          'Authorization': 'Bearer ' + authContext.token,
        },
      }
    );

    const data = await response.json();

    const uuid = obj => obj.uuid === location.pathname.split("/")[2];

    if (data.data.some(uuid)) {
      return <Planning />;
    }
    navigate('/penetrations');

  } catch (error) {

  }
}, [authContext.token]);

useEffect(() => {
  fetchHandler();
}, [fetchHandler]);

};

export default PlanningPage;