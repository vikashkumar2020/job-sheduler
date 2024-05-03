import React from "react";
import { useRecoilValue } from "recoil";
import { jobListStatsState } from "../../recoil/selector/jobSelector";
import styles from './index.module.css';

const JobListStats: React.FC = () => {
  const {
    totalNumJobs,
    totalCompletedNum,
    totalPendingNum,
    totalRunningNum,
  } = useRecoilValue(jobListStatsState);

  return (
    <ul className={styles.jobListStats}>
      <li>Total jobs: {totalNumJobs}</li>
      <li className={styles.completed}>Completed Jobs: {totalCompletedNum}</li>
      <li className={styles.pending}>Pending Jobs: {totalPendingNum}</li>
      <li className={styles.running}>Running Jobs: {totalRunningNum}</li>
    </ul>
  );
};

export default JobListStats;

