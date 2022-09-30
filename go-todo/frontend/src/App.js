/** @format */

import React from "react";
import "./App.css";
import { Container } from "semantic-ui-react";
import ToDoList from "./components/ToDoList";

const App = () => {
  return (
    <Container>
      <ToDoList />
    </Container>
  );
};

export default App;
