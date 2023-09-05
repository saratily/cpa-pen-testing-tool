import React, { useState } from "react";
import './../styles.css';

import Nmap from "./Nmap";
import Nikto from "./Nikto";

const NetworkTabs = () => {
    const [activeTab, setActiveTab] = useState("nmap");
    //  Functions to handle Tab Switching
    const handleNmap = () => {
      // update the state to nmap
      setActiveTab("nmap");
    };
    const handleNikto = () => {
      // update the state to nikto
      setActiveTab("nikto");
    };
    return (
      <div className="Tabs">
        <ul className="nav">
          <li
            className={activeTab === "nmap" ? "active" : ""}
            onClick={handleNmap}
          >
            nmap
          </li>
          <li
            className={activeTab === "nikto" ? "active" : ""}
            onClick={handleNikto}
          >
            nikto
          </li>
        </ul>
   
        <div className="outlet">
          {activeTab === "nmap" ? <Nmap /> : <Nikto />}
        </div>
      </div>
    );
  };
  export default NetworkTabs;
  