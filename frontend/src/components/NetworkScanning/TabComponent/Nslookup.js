import React from "react";
import ReactDOM from "react-dom";
import GenericTile from "sandro-generic-tile10";
import * as Icon from "react-icons/fi";

import "./../styles.css";

function Nslookup() {
   const cmd = "nslookup " + "12"
    return (
      <div className="App">
        {/* <h1>React Generic Tile Demo</h1> */}
        <div className="flex">
          <GenericTile
          header="Header"
          background="#fce4b1"
          width="960px"
          footer="Footer"
          number={cmd}
          color="Warning"
          height="190px"
          />
        </div>
      </div>
    );
}

export default Nslookup;