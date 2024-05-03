import React from "react";
import { useRecoilValue } from "recoil";
import { jobListStatsState } from "../../recoil/selector/jobSelector";

interface JobListStats {
  totalNumJobs: number;
  totalCompletedNum: number;
  totalPendingNum: number;
  totalRunningNum: number;
}

const JobListStats: React.FC = () => {
  const {
    totalNumJobs,
    totalCompletedNum,
    totalPendingNum,
    totalRunningNum,
  } = useRecoilValue<JobListStats>(jobListStatsState);

  return (
    <ul>
      <li>Total jobs: {totalNumJobs}</li>
      <li>Jobs completed Jobs: {totalCompletedNum}</li>
      <li>Jobs Pending Jobs: {totalPendingNum}</li>
      <li>Total Running Jobs: {totalRunningNum}</li>
    </ul>
  );
};

export default JobListStats;
