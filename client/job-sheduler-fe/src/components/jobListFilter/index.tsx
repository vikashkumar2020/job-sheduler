import React from "react";
import { useRecoilState } from "recoil";
import { jobListFilterState } from "../../recoil/atom/jobAtom";
import styles from './index.module.css';

interface JobListFiltersProps {}

const JobListFilters: React.FC<JobListFiltersProps> = () => {
  const [filter, setFilter] = useRecoilState<string>(jobListFilterState);

  const updateFilter = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const { value } = e.target;
    setFilter(value);
  };

  return (
    <div className={styles.jobListFilters}>
      <span className={styles.filterText}>Filter:</span>
      <select className={styles.filterSelect} value={filter} onChange={updateFilter}>
        <option value="All">All</option>
        <option value="Pending">Pending</option>
        <option value="Running">Running</option>
        <option value="Completed">Completed</option>
      </select>
    </div>
  );
};

export default JobListFilters;
