import { useLocation } from 'react-router-dom';

import Tools from './../../Tools/Tools';
import Pens from './../../Penetrations/Pens';

function NslookupEnumeration() {
    const location = useLocation();
    const uuid = location.pathname.split("/")[2];

    return <div className='page'>
  
    <Pens 
      uuid={uuid} />
    <Tools
    type="LookupIP"
    uuid={uuid} 
    title="nslookup command to fetch IP address" /> 
    <Tools
      type="LookupCNAME"
      uuid={uuid}
      title="nslookup command to fetch CNAME record" /> 
    <Tools
      type="LookupNS"
      uuid={uuid}
      title="nslookup command to fetch NS record" /> 
    <Tools
      type="LookupMX"
      uuid={uuid}
      title="nslookup command to fetch MX record" /> 
    <Tools
      type="LookupTXT"
      uuid={uuid}
      title="nslookup command to fetch TXT record" /> 
    <Tools
      type="LookupSRV"
      uuid={uuid}
      title="nslookup command to fetch SRV record" /> 
</div>
    
}

export default NslookupEnumeration;