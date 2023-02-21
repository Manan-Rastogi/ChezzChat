import React from "react";

export default function ChatScreen(props) {
  const styleCardChat = {
    marginTop: "20px",
    maxHeight: "70vh",
    overflow: "auto",
    padding: 10,
  };
  const styleCardSendChat = {
    marginTop: "20px",
    maxHeight: "20vh",
    padding: 10,
    overflow: "auto",
  };
  const styleCardChatDetails = {
    marginTop: "10px",
  };

  const handleChatSubmitButton = () => {
    console.log("Chat Submitted.");
  };
  return (
    <>
      <div className="card" style={styleCardChat}>
        <div className="card" style={styleCardChatDetails}>
          <div className="card-body">This is some text within a card body.</div>
        </div>
        <div className="card" style={styleCardChatDetails}>
          <div className="card-body">This is some text within a card body.</div>
        </div>
        <div className="card" style={styleCardChatDetails}>
          <div className="card-body">This is some text within a card body.</div>
        </div>
        <div className="card" style={styleCardChatDetails}>
          <div className="card-body">This is some text within a card body.</div>
        </div>
        <div className="card" style={styleCardChatDetails}>
          <div className="card-body">This is some text within a card body.</div>
        </div>
        <div className="card" style={styleCardChatDetails}>
          <div className="card-body">This is some text within a card body.</div>
        </div>
        <div className="card" style={styleCardChatDetails}>
          <div className="card-body">This is some text within a card body.</div>
        </div>
        <div className="card" style={styleCardChatDetails}>
          <div className="card-body">This is some text within a card body.</div>
        </div>
        <div className="card" style={styleCardChatDetails}>
          <div className="card-body">This is some text within a card body.</div>
        </div>
        <div className="card" style={styleCardChatDetails}>
          <div className="card-body">This is some text within a card body.</div>
        </div>
        <div className="card" style={styleCardChatDetails}>
          <div className="card-body">This is some text within a card body.</div>
        </div>
        <div className="card" style={styleCardChatDetails}>
          <div className="card-body">This is some text within a card body.</div>
        </div>
        <div className="card" style={styleCardChatDetails}>
          <div className="card-body">This is some text within a card body.</div>
        </div>
        <div className="card" style={styleCardChatDetails}>
          <div className="card-body">This is some text within a card body.</div>
        </div>
        <div className="card" style={styleCardChatDetails}>
          <div className="card-body">This is some text within a card body.</div>
        </div>
        <div className="card" style={styleCardChatDetails}>
          <div className="card-body">This is some text within a card body.</div>
        </div>
      </div>
      <div className="card" style={styleCardSendChat}>
        <div className="row">
          <div className="col-9 col-md-10">
            <div className="card-body d-flex justify-content-center">
              <textarea style={{ width: "100%", resize: "none" }}></textarea>
            </div>
          </div>
          <div className="col-3 col-md-2">
            <div className="card-body d-flex justify-content-center">
              <button type="submit" onClick={handleChatSubmitButton} style={{ width: "100%"}}>
                Send
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
