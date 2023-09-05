import { useLocation } from 'react-router-dom';

import Tools from './../../Tools/Tools';
import Pens from './../../Penetrations/Pens';

function  DigEnumeration() {
    const location = useLocation();
    const uuid = location.pathname.split("/")[2];

    return <div className='page'>
  
    <Pens 
      uuid={uuid} />
    <Tools
    type="digIPv4"
    uuid={uuid} 
    title="dig command to fetch IPv4 address" /> 
    <Tools
      type="digIPv6"
      uuid={uuid}
      title="dig command to fetch IPv6 address" /> 
    <Tools
      type="digCNAME"
      uuid={uuid}
      title="dig command to fetch CNAME record" /> 
    <Tools
      type="digNS"
      uuid={uuid}
      title="dig command to fetch NS record" /> 
    <Tools
      type="digMX"
      uuid={uuid}
      title="dig command to fetch MX record" /> 
    <Tools
      type="digTXT"
      uuid={uuid}
      title="dig command to fetch TXT record" /> 
    <Tools
      type="digANY"
      uuid={uuid}
      title="dig command to fetch ANY record" /> 
    <Tools
      type="digSOA"
      uuid={uuid}
      title="dig command to fetch SOA record" /> 
</div>
    
}

export default  DigEnumeration;