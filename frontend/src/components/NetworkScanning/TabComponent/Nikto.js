import { useLocation } from 'react-router-dom';

import Tools from '../../Tools/Tools';
import Pens from '../../Penetrations/Pens';

function NiktoEnumeration() {
    const location = useLocation();
    const uuid = location.pathname.split("/")[2];

    return <div className='page'>
  
    <Pens 
      uuid={uuid} />
    <Tools
    type="nikto"
    uuid={uuid} 
    title="nikto command" /> 
</div>
    
}

export default NiktoEnumeration;