import React, { useState } from "react";
import TabNavItem from "./TabNavItem";
import TabContent from "./TabContent";
import FfufEnumeration from "./ffuf";
import DirbEnumerate from "./dirb";
import WfuzzEnumeration from "./wfuzz";
import WappalyzerEnumeration from "./wappalyzer";
 
const Tabs = () => {
  const [activeTab, setActiveTab] = useState("ffuf");
 
  return (
    <div className="Tabs">
      <ul className="three">
        <TabNavItem title="ffuf" id="ffuf" activeTab={activeTab} setActiveTab={setActiveTab}/>
        <TabNavItem title="dirb" id="dirb" activeTab={activeTab} setActiveTab={setActiveTab}/>
        <TabNavItem title="wfuzz" id="wfuzz" activeTab={activeTab} setActiveTab={setActiveTab}/>
        <TabNavItem title="Wappalyzer" id="Wappalyzer" activeTab={activeTab} setActiveTab={setActiveTab}/>
      </ul>
 
      <div className="outlet">
        <TabContent id="ffuf" activeTab={activeTab}>
          <FfufEnumeration/>
        </TabContent>
        <TabContent id="dirb" activeTab={activeTab}>
          <DirbEnumerate />
        </TabContent>
        <TabContent id="wfuzz" activeTab={activeTab}>
          <WfuzzEnumeration />
        </TabContent>
        <TabContent id="Wappalyzer" activeTab={activeTab}>
          <WappalyzerEnumeration />
        </TabContent>
      </div>
    </div>
  );
};
 
export default Tabs;