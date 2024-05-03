// WebSocket.js
import { useState, useEffect } from 'react';

export function useWebSocket() {
  const [socket, setSocket] = useState<WebSocket | null>(null);

  useEffect(() => {
    const newSocket = new WebSocket(`ws://${import.meta.env.VITE_BACKEND_URL}/jobs`);

    newSocket.onopen = () => {
      console.log('Connection established');
    };

    setSocket(newSocket);

    return () => {
      newSocket.close();
    };
  }, []);

  return socket;
}
