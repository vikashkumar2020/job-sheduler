import { atom } from "recoil";
import { JobInterface } from "../../interfaces/job";

// atom for job list state
const jobListState = atom<Array<JobInterface>>({
  key: "jobListState",
  default: [],
});

// atom for job list filter state
const jobListFilterState = atom<string>({
  key: "jobListFilterState",
  default: "All",
});

export { jobListState, jobListFilterState };
