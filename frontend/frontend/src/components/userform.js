import React, { useState } from 'react';
import axios from 'axios';

export default function UserForm({ refresh }) {
  const [formData, setFormData] = useState({ name: '', email: '', id: '' });

  const handleChange = (e) => {
    setFormData(prev => ({ ...prev, [e.target.name]: e.target.value }));
  };

  const createUser = async () => {
    await axios.post('http://localhost:8080/users', {
      name: formData.name,
      email: formData.email,
    });
    refresh();
  };

  const updateUser = async () => {
    await axios.put('http://localhost:8080/users/update', {
      id: formData.id,
      name: formData.name,
      email: formData.email,
    });
    refresh();
  };

  const deleteUser = async () => {
    await axios.delete('http://localhost:8080/users/delete', {
      data: { id: formData.id }
    });
    refresh();
  };

  return (
    <div>
      <h2>Manage User</h2>
      <input name="id" placeholder="User ID (for update/delete)" onChange={handleChange} />
      <input name="name" placeholder="Name" onChange={handleChange} />
      <input name="email" placeholder="Email" onChange={handleChange} />
      <br />
      <button onClick={createUser}>Create</button>
      <button onClick={updateUser}>Update</button>
      <button onClick={deleteUser}>Delete</button>
    </div>
  );
}
