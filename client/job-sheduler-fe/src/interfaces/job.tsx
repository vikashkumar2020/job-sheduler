export interface JobInterface {
    id: string;
    name: string;
    duration: number;
    created_at: string;
    updated_at: string;
    status: string;
}

export interface JobStats {
    totalNumJobs: number;
    totalCompletedNum: number;
    totalPendingNum: number;
    totalRunningNum: number;
  }