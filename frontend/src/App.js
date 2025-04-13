import React, { useState, useEffect } from 'react';
import TaskForm from './components/TaskForm';
import TaskList from './components/TaskList';
import * as api from './api/tasks';  

function App() {
  const [tasks, setTasks] = useState([]);

  const getTasks = async () => {
    try {
      const res = await api.fetchTasks();
      setTasks(res.data); 
    } catch (error) {
      console.error('Error fetching tasks:', error);
    }
  };
  
  const createTask = async (task) => {
    try {
      await api.createTask(task);
      getTasks(); 
    } catch (error) {
      console.error('Error creating task:', error);
    }
  };

  const deleteTask = async (id) => {
    try {
      await api.deleteTask(id);
      getTasks(); 
    } catch (error) {
      console.error('Error deleting task:', error);
    }
  };

  const editTask = async (updatedTask) => {
    try {
      await api.updateTask(updatedTask.id, updatedTask);
      getTasks(); 
    } catch (error) {
      console.error('Error editing task:', error);
    }
  };

  useEffect(() => {
    getTasks();
  }, []);

  return (
    <div>
      <h1>Task Manager</h1>
      <TaskForm onCreate={createTask} />
      <TaskList tasks={tasks} onDelete={deleteTask} onEdit={editTask} />
    </div>
  );
}

export default App;
