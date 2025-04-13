import React, { useState } from 'react';

function TaskItem({ task, onDelete, onEdit }) {
  const [isEditing, setIsEditing] = useState(false);
  const [updatedTask, setUpdatedTask] = useState({ ...task });

  const handleChange = (e) => {
    setUpdatedTask({ ...updatedTask, [e.target.name]: e.target.value });
  }; 

  const handleEditSubmit = (e) => {
    e.preventDefault();
  
    const formattedTask = {
      ...updatedTask,
      due_date: updatedTask.due_date ? new Date(updatedTask.due_date).toISOString() : null,
    };
  
    onEdit(formattedTask); 
    setIsEditing(false); 
  };
  

  return (
    <div>
      {!isEditing ? (
        <>
          <h4>{task.title} ({task.status})</h4>
          <p>{task.description}</p>
          <small>Due: {task.due_date ? new Date(task.due_date).toLocaleString() : 'N/A'}</small>
          <br />
          <button onClick={() => setIsEditing(true)}>Edit</button>
          <button onClick={() => onDelete(task.id)}>Delete</button>
        </>
      ) : (
        <form onSubmit={handleEditSubmit}>
          <input
            type="text"
            name="title"
            value={updatedTask.title}
            onChange={handleChange}
            required
          />
          <textarea
            name="description"
            value={updatedTask.description}
            onChange={handleChange}
            required
          />
          <select name="status" value={updatedTask.status} onChange={handleChange}>
            <option value="Pending">Pending</option>
            <option value="In-Progress">In-Progress</option>
            <option value="Completed">Completed</option>
          </select>
          <input
            type="datetime-local"
            name="due_date"
            value={updatedTask.due_date?.slice(0, 16)}
            onChange={handleChange}
          />
          <button type="submit">Save</button>
          <button type="button" onClick={() => setIsEditing(false)}>Cancel</button>
        </form>
      )}
      <hr />
    </div>
  );
}

export default TaskItem;
