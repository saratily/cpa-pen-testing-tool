import { useContext, useCallback } from "react";
import { useLocation, useNavigate } from 'react-router-dom';

import AuthContext from "../store/auth-context";
import NetworkScanning from '../components/NetworkScanning/NetworkScanning';

const NetworkScanningPage = () => {
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

      if (!data.data.some(uuid)) {
        navigate('/penetrations');
      }

    } catch (error) {}
  });
  fetchHandler();
  return <NetworkScanning />;
};

export default NetworkScanningPage;