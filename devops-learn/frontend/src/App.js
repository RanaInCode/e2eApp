import React, { useEffect, useState } from 'react';

function App() {
    const [tasks, setTasks] = useState([]);

    useEffect(() => {
        fetch('http://localhost:8080/api/tasks')
          .then(res => res.json())
          .then(data => setTasks(data));
    }, []);

    return (
        <div>
            <h1>My Tasks</h1>
            <ul>
                {tasks.map(t => <li key={t.id}>{t.title}</li>)}
            </ul>
        </div>
    );
}

export default App;