import { selector } from "recoil";
import { jobListState, jobListFilterState } from "../atom/jobAtom";

const filteredJobListState = selector({
    key: "filteredJobListState",
    get: ({ get }) => {
      const filter = get(jobListFilterState);
      const list = get(jobListState);
  
      switch (filter) {
        case "Pending":
          return list.filter((item: any) => item.status=== "Pending");
        case "Running":
          return list.filter((item: any) => item.status=== "Running");
        case "Completed":
            return list.filter((item: any) => item.status === "Completed");
        default:
          return list;
      }
    }
  });

  const jobListStatsState = selector({
    key: "jobListStatsState",
    get: ({ get }) => {
      const jobList = get(jobListState);
      const totalNumJobs = jobList.length;
      const totalCompletedNum = jobList.filter((item: any) => item.status === "Completed").length;
      const totalPendingNum = jobList.filter((item: any) => item.status === "Pending").length;
      const totalRunningNum = jobList.filter((item: any) => item.status === "Running").length;
  
      return {
        totalNumJobs,
        totalCompletedNum,
        totalPendingNum,
        totalRunningNum,
      };
    }
  });