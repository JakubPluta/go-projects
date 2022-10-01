/** @format */

import React, { useState, useEffect } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";

const TodoList = () => {
  const [task, setTask] = useState("");
  const [items, setItems] = useState([]);

  useEffect(() => {
    getTask();
  });

  const undoTask = (id) => {
    axios
      .put("/api/undoTask/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
      })
      .then((res) => {
        console.log(res);
        getTask();
      });
  };

  const deleteTask = (id) => {
    axios
      .delete("/api/deleteTask/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
      })
      .then((res) => {
        console.log(res);
        getTask();
      });
  };

  const updateTask = (id) => {
    axios
      .put("/api/task/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
      })
      .then((res) => {
        console.log(res);
        getTask();
      });
  };

  const getTask = () => {
    axios.get("/api/task").then((res) => {
      if (res.data) {
        setItems(
          res.data.map((item) => {
            let color = "yellow";
            let style = { wordWrap: "break-word" };
            if (item.status) {
              color = "green";
              style["textDecorationLine"] = "line-through";
            }
            return (
              <Card key={item._id} color={color} fluid>
                <Card.Content>
                  <Card.Header textAlign="left">
                    <div style={style}>{item.task}</div>
                  </Card.Header>

                  <Card.Meta textAlign="right">
                    <Icon
                      name="check circle"
                      color="green"
                      onClick={() => updateTask(item._id)}
                    />
                    <span style={{ paddingRight: 10 }}>Done</span>
                    <Icon
                      name="undo"
                      color="yellow"
                      onClick={() => undoTask(item._id)}
                    />
                    <span style={{ paddingRight: 10 }}>Undo</span>
                    <Icon
                      name="delete"
                      color="red"
                      onClick={() => deleteTask(item._id)}
                    />
                    <span style={{ paddingRight: 10 }}>Delete</span>
                  </Card.Meta>
                </Card.Content>
              </Card>
            );
          })
        );
      } else {
        setItems([]);
      }
    });
  };

  const onSubmit = () => {
    console.log("task", task);

    if (task) {
      axios
        .post(
          "/api/task",
          {
            task,
          },
          {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded",
            },
          }
        )
        .then((res) => {
          getTask();
          setTask("");
          console.log(res);
        });
    }
  };

  return (
    <div>
      <div className="row">
        <Header className="header" as="h2">
          TO DO LIST
        </Header>
      </div>
      <div className="row">
        <Form onSubmit={onSubmit}>
          <Input
            type="text"
            name="task"
            onChange={(e) => setTask(e.target.value)}
            value={task}
            fluid
            placeholder="Create Task"
          />
          {/* <Button >Create Task</Button> */}
        </Form>
      </div>
      <div className="row">
        <Card.Group>{items}</Card.Group>
      </div>
    </div>
  );
};

export default TodoList;
