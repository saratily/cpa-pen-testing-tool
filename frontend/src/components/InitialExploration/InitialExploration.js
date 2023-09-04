import { useLocation } from 'react-router-dom';

import Tools from './Tools';

function NetworkScanning() {
    const location = useLocation();
    const uuid = location.pathname.split("/")[2];

    return <div className='page'>NetworkScanning 
    <Tools
      type="reachable"
      uuid={uuid} />
    </div>
}

export default NetworkScanning;