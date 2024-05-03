import React from 'react';
import { JobInterface } from "../../interfaces/job";
import styles from './index.module.css';

interface JobItemProps {
    item: JobInterface;
}

const JobItem: React.FC<JobItemProps> = ({ item }) => {
    let backgroundColor;

    switch (item.status) {
        case 'Pending':
            backgroundColor = '#d3d3d3'; // grey
            break;
        case 'Running':
            backgroundColor = '#ffffcc'; // yellow
            break;
        case 'Completed':
            backgroundColor = '#ccffcc'; // green
            break;
        default:
            backgroundColor = '#ffffff'; // default white
            break;
    }

    return (
        <div className={styles.jobItem} style={{ backgroundColor }}>
            <div>Name: {item.name}</div>
            <div>Duration: {item.duration / 1e9} sec</div>
            <div>Status: {item.status}</div>
        </div>
    );
};

export default JobItem;
