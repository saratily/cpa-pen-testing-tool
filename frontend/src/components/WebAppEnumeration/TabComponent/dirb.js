import { useLocation } from 'react-router-dom';

import Tools from './../../Tools/Tools';
import Pens from './../../Penetrations/Pens';

function  DirbEnumeration() {
    const location = useLocation();
    const uuid = location.pathname.split("/")[2];

    return <div className='page'>
      <h1>DirbEnumerate</h1>
    <Pens 
      uuid={uuid} />
    <Tools
    type="dirb"
    uuid={uuid} 
    title="dig command to fetch IPv4 address" /> 
    </div>
    
}

export default  DirbEnumeration;