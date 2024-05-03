import React, { useEffect } from 'react';
import './App.css';
import { useWebSocket } from './api/getAllJobs';
import { useSetRecoilState } from 'recoil';
import { jobListState } from './recoil/atom/jobAtom';
import JobList from './components/jobList';
import { JobInterface } from './interfaces/job';

const App: React.FC = () => {
    const setJobList = useSetRecoilState(jobListState);
    const socket = useWebSocket();

    useEffect(() => {
        if (socket) {
            socket.onmessage = (message) => {
                const data = JSON.parse(message.data);
                if (data && data.jobs) {
                    setJobList(data.jobs as JobInterface[]);
                }
            };
        }

        // Clean up function
        return () => {
            if (socket) {
                socket.close();
            }
        };
    }, [socket, setJobList]);

    return (
        <div>
            <JobList />
        </div>
    );
}

export default App;
