
import React, { Component } from "react"
import axios from "axios";
import { Card, Header, Form, Input, Icon } from "semantic-ui-react"
import {useState} from "react"
let endpoint = process.env.REACT_BACKEND_URI | "http//localhost:9000";



const ToDoList = () => {
    const [task, setTask] = useState("");
    const [items, setItems] = useState([])

    const onSubmit = () => {

    }

    const onChange = () => {

    }

    return (
    <div>
    <div className="row">
            <Header className="header" as='h2' color="black">
                To Do List
            </Header>
            </div>
            <div className="row">
                <Form onSubmit={onSubmit}>
                    <Input type="text" name="task" onChange={onChange} value={task} fluid placeholder="Create Task" />
                </Form>
    </div>
    </div>
  )
};

export default ToDoList;



