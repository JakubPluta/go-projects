/** @format */

import React from "react";
import "./App.css";
import { connect, sendMessage } from "./api";
import { useEffect } from "react";
import Header from "./components/Header/Header";
import ChatHistory from "./components/ChatHistory";
import { Component } from "react";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory: [],
    };
  }

  componentDidMount() {
    connect((msg) => {
      console.log("New Message");
      this.setState((prevState) => ({
        chatHistory: [...this.state.chatHistory, msg],
      }));
      console.log(this.state);
    });
  }

  send() {
    console.log("hello");
    sendMessage("hello");
  }

  render() {
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <button onClick={this.send}>Hit</button>
      </div>
    );
  }
}

export default App;
