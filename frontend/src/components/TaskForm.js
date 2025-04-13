import React, { useState } from 'react';

function TaskForm({ onCreate }) {
  const [task, setTask] = useState({
    title: '',
    description: '',
    status: 'Pending',
    due_date: '',
  });

  const handleChange = (e) => {
    setTask({ ...task, [e.target.name]: e.target.value });
  }; 

  const handleSubmit = (e) => {
    e.preventDefault();
  
    const formattedTask = {
      ...task,
      due_date: task.due_date ? new Date(task.due_date).toISOString() : null, 
    };
  
    onCreate(formattedTask);
    setTask({
      title: '',
      description: '',
      status: 'Pending',
      due_date: '',
    });
  };
  

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        name="title"
        value={task.title}
        onChange={handleChange}
        placeholder="Title"
        required
      />
      <textarea
        name="description"
        value={task.description}
        onChange={handleChange}
        placeholder="Description"
        required
      />
      <select name="status" value={task.status} onChange={handleChange}>
        <option value="Pending">Pending</option>
        <option value="In-Progress">In-Progress</option>
        <option value="Completed">Completed</option>
      </select>
      <input
        type="datetime-local"
        name="due_date"
        value={task.due_date}
        onChange={handleChange}
      />
      <button type="submit">Add Task</button>
    </form>
  );
}

export default TaskForm;
