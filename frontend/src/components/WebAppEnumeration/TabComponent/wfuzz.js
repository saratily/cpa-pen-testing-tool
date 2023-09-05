import { useLocation } from 'react-router-dom';

import Tools from './../../Tools/Tools';
import Pens from './../../Penetrations/Pens';

function  WfuzzEnumeration() {
    const location = useLocation();
    const uuid = location.pathname.split("/")[2];

    return <div className='page'>
    <Pens 
      uuid={uuid} />
    <Tools
    type="wfuzz"
    uuid={uuid}  /> 
    </div>
    
}

export default  WfuzzEnumeration;