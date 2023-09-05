import { useLocation } from 'react-router-dom';

import Tools from './../Tools/Tools';
import Pens from './../Penetrations/Pens';

function NetworkScanning() {
    const location = useLocation();
    const uuid = location.pathname.split("/")[2];

    return <div className='page'>
  
    <Pens 
      uuid={uuid} />
    <Tools
    type="ping"
    uuid={uuid} 
    title="ping command"/> 
    <Tools
      type="whois"
      uuid={uuid} 
      title="Whois command"/>
</div>
    
}

export default NetworkScanning;