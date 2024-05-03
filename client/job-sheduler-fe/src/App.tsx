import { useEffect } from 'react';
import './App.css';
import { useWebSocket } from './api/getAllJobs';
import AddForm from './components/addForm';
import JobListStats from './components/jobListStats';
import { useSetRecoilState } from 'recoil';
import { jobListState } from './recoil/atom/jobAtom';

function App() {
    const setJobList = useSetRecoilState(jobListState);

    const socket = useWebSocket();

    useEffect(() => {
        if (socket) {
            socket.onmessage = (message) => {
                const data = JSON.parse(message.data);
                if (data && data.jobs) {
                  console.log(data.jobs)
                    setJobList(data.jobs);
                }
            };
        }

        // Clean up function
        return () => {
            if (socket) {
                socket.close();
            }
        };
    }, [socket]);

    return (
        <>
            <AddForm />
            <JobListStats />
        </>
    );
}

export default App;
