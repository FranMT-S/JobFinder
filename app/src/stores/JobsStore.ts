import { defineStore } from "pinia";
import { ref } from "vue";
import type { JobScrapeated } from "../types/Job";



export const useJobsStore = defineStore('jobs', () =>{
  const jobScrapeated = ref<JobScrapeated[]>([])
  const isLoading = ref(false);
  const firstLoad = ref(false)

  const resetJobs = () => {
    jobScrapeated.value = []
  }

  return {
    jobScrapeated,
    resetJobs,
    isLoading,
    firstLoad
  }
})