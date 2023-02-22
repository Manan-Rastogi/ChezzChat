import React, { useState } from "react";
import FriendButton from "./FriendButton";
import GroupButton from "./GroupButton";


export default function SidePanel(props) {

  let maxGroups = props.maxGroups
  let groups = {}
  for (let i = 1; i <= maxGroups; i++){
    groups[i] = "success"
  }

  const [activeGroup, setActiveGroup] = useState(groups)

  const styleCardButtons = {
    marginTop: "20px",
    maxHeight: "42vh",
    minHeight: "42vh",
    overflow: "auto",
  };


  const styleButtonsGroups = {
    marginTop: 5,
    width: "80%",
    marginLeft: "10%",
    marginRight: "10%",
    color: "black",
    fontWeight: "bold",
  };


  const styleButtonsFriends = {
    marginTop: 5,
    width: "80%",
    marginLeft: "10%",
    marginRight: "10%",
    color: "black",
    fontWeight: "bold",
  };


  const setActiveButton = (id) => {
    setActiveGroup(prevActiveGroup => {
      const newActiveGroup = {...prevActiveGroup};
      for (let i = 1; i <= maxGroups; i++){
        newActiveGroup[i] = "success";
      }
      newActiveGroup[id] = "warning";
      return newActiveGroup;
    });
  }
  
  return (
    <div>
      <div className="card" style={styleCardButtons}>
        <div className="btn-group-vertical" style={{ textAlign: "center" }}>
          <GroupButton
          key={1}
            keyId={1}
            GnFStyle={styleButtonsGroups}
            typeButton={activeGroup[1]}
            Name={"Group 1"}
            activeHandler = {setActiveButton}
          />
          <GroupButton 
          key={2}
           keyId={2}
            GnFStyle={styleButtonsGroups}
            typeButton={activeGroup[2]}
            Name={"Group 2"}
            activeHandler = {setActiveButton}
          />

          {/* A fixed button to add new Group */}
          <GroupButton
            GnFStyle={styleButtonsGroups}
            typeButton={"success"}
            Name={"Add New Group"}
            activeHandler = {setActiveButton}
          />
        </div>
      </div>

      <div className="card" style={styleCardButtons}>
        <div className="btn-group-vertical" style={{ textAlign: "center" }}>
          <FriendButton
            GnFStyle={styleButtonsFriends}
            typeButton={"primary"}
            Name={"Friend 1"}
          />
          <FriendButton
            GnFStyle={styleButtonsFriends}
            typeButton={"primary"}
            Name={"Friend 2"}
          />
          <FriendButton
            GnFStyle={styleButtonsFriends}
            typeButton={"primary"}
            Name={"Friend 3"}
          />
          <FriendButton
            GnFStyle={styleButtonsFriends}
            typeButton={"primary"}
            Name={"Add New Friend"}
          />
        </div>
      </div>
    </div>
  );
}
