import React from "react";

export default function FriendButton(props) {

  return (
    <>
      <button
        type="button"
        className={`btn btn-${props.typeButton}`}
        style={props.GnFStyle}
        onClick={() =>  props.activeHandler(props.keyId)}
      >
        {props.Name}
      </button>
    </>
  );
}