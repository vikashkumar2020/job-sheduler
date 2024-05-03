import { useEffect, useState } from 'react'
import './App.css'
import { useWebSocket } from './api/getAllJobs';

function App() {

    const socket = useWebSocket();

    useEffect(() => {
      if (socket) {
        socket.onmessage = (message) => {
          console.log('Message received:', message.data);
        };
      }
    }, [socket]);
    

  return (
    <>
      hi there
    </>
  )
}

export default App