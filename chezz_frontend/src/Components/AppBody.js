import React from "react";
import ChatScreen from "./ChatScreen";
import SidePanel from "./SidePanel";

export default function AppBody() {
  return (
    <>
      <div className="container-fluid">
        <div className="row">
          <div className="col-4 col-md-3">
            <SidePanel></SidePanel>
          </div>
          <div className="col-8 col-md-9"><ChatScreen /></div>
        </div>
      </div>
    </>
  );
}
