import React, { useState } from 'react';
import UserForm from './components/userform';
import UserList from './components/userlist';

function App() {
  const [reload, setReload] = useState(false);

  const refresh = () => setReload(!reload);

  return (
    <div className="App">
      <h1>CRUD </h1>
      <UserForm refresh={refresh} />
      <UserList key={reload} />
    </div>
  );
}

export default App;
