import { useLocation } from 'react-router-dom';

import Tools from './../../Tools/Tools';
import Pens from './../../Penetrations/Pens';

function  FfufEnumeration() {
    const location = useLocation();
    const uuid = location.pathname.split("/")[2];

    return <div className='page'>
    <Pens 
      uuid={uuid} />
    <Tools
    type="ffuf"
    uuid={uuid} /> 
    </div>
    
}

export default  FfufEnumeration;