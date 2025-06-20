import React, { useEffect, useState } from 'react';
import axios from 'axios';

export default function UserList() {
  const [users, setUsers] = useState([]);

  const fetchUsers = async () => {
    const res = await axios.get('http://localhost:8080/users');
    setUsers(res.data);
  };

  useEffect(() => {
    fetchUsers();
  }, []);

  return (
    <div>
      <h2>All Users</h2>
      {users.length === 0 && <p>No users found</p>}
      <ul>
        {users.map(user => (
          <li key={user.id}>
            <strong>{user.name}</strong> ({user.email})<br />
            ID: <code>{user.id}</code>
          </li>
        ))}
      </ul>
    </div>
  );
}
