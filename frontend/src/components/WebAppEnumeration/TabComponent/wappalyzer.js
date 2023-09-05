import { useLocation } from 'react-router-dom';

import Tools from './../../Tools/Tools';
import Pens from './../../Penetrations/Pens';

function  WappalyzerEnumeration() {
    const location = useLocation();
    const uuid = location.pathname.split("/")[2];

    return <div className='page'>
      <h1>WappalyzerEnumeration</h1>
    <Pens 
      uuid={uuid} />
    <Tools
    type="Wappalyzer"
    uuid={uuid} 
    title="dig command to fetch IPv4 address" /> 
    </div>
    
}

export default  WappalyzerEnumeration;