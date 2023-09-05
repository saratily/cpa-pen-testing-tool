import React, { useState } from "react";
import './../styles.css';

import Dig from "./Dig";
import Nslookup from "./Nslookup";

const Tabs = () => {
    const [activeTab, setActiveTab] = useState("dig");
    //  Functions to handle Tab Switching
    const handleDig = () => {
      // update the state to dig
      setActiveTab("dig");
    };
    const handleNslookup = () => {
      // update the state to nslookup
      setActiveTab("nslookup");
    };
    return (
      <div className="Tabs">
        <ul className="nav">
          <li
            className={activeTab === "dig" ? "active" : ""}
            onClick={handleDig}
          >
            dig
          </li>
          <li
            className={activeTab === "nslookup" ? "active" : ""}
            onClick={handleNslookup}
          >
            nslookup
          </li>
        </ul>
   
        <div className="outlet">
          {activeTab === "dig" ? <Dig /> : <Nslookup />}
        </div>
      </div>
    );
  };
  export default Tabs;
  