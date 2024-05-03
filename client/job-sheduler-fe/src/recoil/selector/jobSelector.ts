import { selector } from "recoil";
import { jobListState, jobListFilterState } from "../atom/jobAtom";
import { JobInterface, JobStats } from "../../interfaces/job";

// selector for filtered job list state
const filteredJobListState = selector<Array<JobInterface>>({
  key: "filteredJobListState",
  get: ({ get }) => {
    const filter = get(jobListFilterState);
    const list = get(jobListState);

    switch (filter) {
      case "Pending":
        return list.filter((item) => item.status === "Pending");
      case "Running":
        return list.filter((item) => item.status === "Running");
      case "Completed":
        return list.filter((item) => item.status === "Completed");
      default:
        return list;
    }
  }
});

// selector for job list statistics state
const jobListStatsState = selector<JobStats>({
  key: "jobListStatsState",
  get: ({ get }) => {
    const jobList = get(jobListState);
    const totalNumJobs = jobList.length;
    const totalCompletedNum = jobList.filter((item) => item.status === "Completed").length;
    const totalPendingNum = jobList.filter((item) => item.status === "Pending").length;
    const totalRunningNum = jobList.filter((item) => item.status === "Running").length;

    return {
      totalNumJobs,
      totalCompletedNum,
      totalPendingNum,
      totalRunningNum,
    };
  }
});

export { filteredJobListState, jobListStatsState };
