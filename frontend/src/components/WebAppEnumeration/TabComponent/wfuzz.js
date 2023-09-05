import { useLocation } from 'react-router-dom';

import Tools from './../../Tools/Tools';
import Pens from './../../Penetrations/Pens';

function  WfuzzEnumeration() {
    const location = useLocation();
    const uuid = location.pathname.split("/")[2];

    return <div className='page'>
      <h1> WfuzzEnumeration </h1>
    <Pens 
      uuid={uuid} />
    <Tools
    type="wfuzz"
    uuid={uuid} 
    title="dig command to fetch IPv4 address" /> 
    </div>
    
}

export default  WfuzzEnumeration;