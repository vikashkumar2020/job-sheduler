import { atom } from "recoil";

const jobListState = atom({
  key: "jobListState",
  default: [] as any,
});

const jobListFilterState = atom({
    key: "todoListFilterState",
    default: "All"
  });

export {jobListState, jobListFilterState}