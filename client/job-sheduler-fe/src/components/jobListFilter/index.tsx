import React from "react";
import { useRecoilState } from "recoil";
import { jobListFilterState } from "../../recoil/atom/jobAtom";

interface JobListFiltersProps {
  
}

const JobListFilters: React.FC<JobListFiltersProps> = () => {
  const [filter, setFilter] = useRecoilState<string>(jobListFilterState);

  const updateFilter = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const { value } = e.target;
    setFilter(value);
  };

  return (
    <>
      Filter:
      <select value={filter} onChange={updateFilter}>
        <option value="All">All</option>
        <option value="Pending">Pending</option>
        <option value="Running">Running</option>
        <option value="Completed">Completed</option>
      </select>
      <p></p>
    </>
  );
};

export default JobListFilters;
