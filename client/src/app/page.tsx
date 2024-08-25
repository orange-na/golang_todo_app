"use client";

import axios from "axios";
import { useEffect, useState } from "react";

type User = {
  id: string;
  name: string;
};

type Task = {
  id: string;
  title: string;
  userId: string;
};

export default function Home() {
  const [users, setUsers] = useState<User[]>([]);
  const [tasks, setTasks] = useState<Task[]>([]);
  const [newTask, setNewTask] = useState("");

  const getApi = async () => {
    try {
      const response = await axios.get("http://localhost:8080/tasks");
      const data = await response.data;
      setTasks(data);
    } catch (error) {
      console.log(error);
    }
  };

  const addTask = async () => {
    try {
      const response = await axios.post("http://localhost:8080/tasks", {
        title: newTask,
        userId: "817e7c75-4ba8-4aac-86af-999ad0c4e13c",
      });
      const addedTask = await response.data;
      setTasks([...tasks, addedTask]);
      setNewTask("");
    } catch (error) {
      console.log(error);
    }
  };

  console.log(users);
  console.log(tasks);

  useEffect(() => {
    getApi();
  }, []);

  return (
    <div className="max-w-md mx-auto">
      <h1 className="text-2xl font-bold mb-4">Todo List</h1>
      <ul className="space-y-2">
        {tasks.map((task) => (
          <li
            key={task.id}
            className="flex items-center space-x-2 bg-gray-100 p-2 rounded"
          >
            <input
              type="checkbox"
              className="form-checkbox h-5 w-5 text-blue-600"
            />
            <span className="text-gray-800">{task.title}</span>
          </li>
        ))}
      </ul>
      <div className="mt-4">
        <input
          type="text"
          className="border border-gray-300 rounded px-2 py-1 w-full"
          placeholder="Add a new task"
          value={newTask}
          onChange={(e) => setNewTask(e.target.value)}
        />
        <button
          className="mt-2 bg-blue-500 text-white px-4 py-2 rounded"
          onClick={addTask}
        >
          Add Task
        </button>
      </div>
    </div>
  );
}
