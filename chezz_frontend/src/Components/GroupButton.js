import React from "react";

export default function GroupButton(props) {
  const handleClick = () => {
    props.activeHandler(props.keyId);
  };
  return (
    <>
      <button
        type="button"
        className={`btn btn-${props.typeButton}`}
        style={props.GnFStyle}
        onClick={handleClick}
      >
        {props.Name}
      </button>
    </>
  );
}
