import React from 'react';
import { useRecoilValue } from "recoil";
import { filteredJobListState } from "../../recoil/selector/jobSelector";
import JobItem from "../jobItem";
import JobListStats from "../jobListStats";
import AddForm from "../addForm";
import JobListFilters from "../jobListFilter";
import styles from './index.module.css';

const JobList: React.FC = () => {
    const jobList = useRecoilValue(filteredJobListState);

    return (
        <div className={styles.jobListContainer}>
            
            <div className={styles.controls}>
                <AddForm />
                <JobListStats />
                <JobListFilters />
            </div>
            
            <div className={styles.list}>

                {jobList.map((job) => (
                    <JobItem item={job} key={job.id} />
                ))}
            </div>
        </div>
    );
};

export default JobList;
