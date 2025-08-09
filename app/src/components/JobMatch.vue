<script setup lang="ts">
import { computed, ref } from 'vue';
import type { JobScrapeated } from '../types/Job';
import ChipMatcher from './ChipMatcher.vue';
import JobDescriptionModal from './JobDescriptionModal.vue';
import { storeToRefs } from 'pinia';
import { useUserStore } from '../stores/userStore';
import JobMatchSalarySlider from './JobMatchSalarySlider.vue';

  const props = defineProps<{
    jobScrapeated: JobScrapeated
  }>()

  const matchedSkillsSet = computed(() => {
    const normalizedSkills = props.jobScrapeated.matchAnalizer.skillMatches.map(skill => skill.toLowerCase())
    return new Set(normalizedSkills)
  })

  const remoteOkURL = 'https://remoteok.com/'
  const weWorkRemotelyURL = 'https://weworkremotely.com/'
  const isActiveDialog = ref(false)

  const {minSalary, maxSalary} = storeToRefs(useUserStore())

  function formatDate(dateStr: string | null) {
    if (!dateStr) return ''
    const date = new Date(dateStr)
    return date.toLocaleDateString(undefined, {
      year: 'numeric',
      month: 'short',
      day: 'numeric'
    })
  }
  </script>


<template>
   <v-card class="!p-4 !m-4" color="surface">
    <div class="flex justify-between items-start flex-wrap gap-2">
      <div class="flex gap-2 ">
        <v-chip
        :href="jobScrapeated.job.web === 'remoteok' ? remoteOkURL : weWorkRemotelyURL"
        target="_blank"
        class="capitalize"
        color="blue-darken-1">{{ jobScrapeated.job.web }}
        </v-chip>
        <ChipMatcher :matchAnalizer="jobScrapeated.matchAnalizer" />
      </div>

      <JobDescriptionModal :jobScrapeated="jobScrapeated" @update:isActiveDialog="isActiveDialog = $event" :isActiveDialog="isActiveDialog" />

      <div class="flex gap-2 flex-wrap">
        <v-btn
          color="purple-lighten-1"
          prepend-icon="mdi-link"
          size="small"
          @click="isActiveDialog = true"
        >
          View Description
        </v-btn>
  
        <v-btn
          :href="jobScrapeated.job.url"
          target="_blank"
          color="primary"
          prepend-icon="mdi-link"
          size="small"
        >
          View offer
        </v-btn>

      </div>

    </div>

    <div class="flex justify-between items-start">
      <div class="flex flex-col w-full">
        <v-card-title color="primary">
          {{ jobScrapeated.job.position }}
        </v-card-title>
        <v-card-subtitle v-if="jobScrapeated.job.createdAt" color="primary" class="!flex items-center gap-2">
          <v-tooltip text="Date created" open-delay="300">
            <template #activator="{ props }">
              <v-icon v-bind="props" icon="mdi-calendar" size="small" />
            </template>
          </v-tooltip>
          <span>{{ formatDate(jobScrapeated.job.createdAt.toString()) }}</span>
        </v-card-subtitle>
      </div>
    </div>

    <div class="mt-2 ">
      <v-card-title size="small" color="primary" class=" items-center gap-2">
        <v-row class="w-full">
          <v-col cols="12" md="12" class="overflow-hidden text-ellipsis">
            <v-tooltip text="Company" open-delay="300">
              <template #activator="{ props }">
                <v-icon color="primary" v-bind="props" icon="mdi-domain" size="default" />
              </template>
            </v-tooltip>
            <span  class="ml-2 text-[0.9rem] font-normal  ">
              {{ jobScrapeated.job.company }}
            </span>
          </v-col>
          <v-col cols="12" md="12" class="overflow-hidden text-ellipsis">
            <v-tooltip text="Contract" open-delay="300">
              <template #activator="{ props }">
                <v-icon color="primary" v-bind="props" icon="mdi-text-box-outline" size="default" />
              </template>
            </v-tooltip>
            <span class="ml-2 text-[0.9rem] font-normal capitalize">{{ jobScrapeated.job.contractType }}</span>
          </v-col>
        </v-row>

        <v-row v-if="jobScrapeated.job.level.length > 0">
          <v-col cols="12" md="2">
            <v-tooltip text="Seniority" open-delay="300">
              <template #activator="{ props }">
                <v-icon color="primary" v-bind="props" icon="mdi-account" size="default" />
              </template>
            </v-tooltip>
            <v-chip
              size="small"
              label 
              color="secondary"
              class="ml-2 capitalize"
              v-for="level in jobScrapeated.job.level"
              :key="level"
            >
              {{ level }}
            </v-chip>
          </v-col>
        </v-row>

        <v-row v-if="jobScrapeated.job.location.length > 0">
          <v-col cols="12">
            <v-tooltip text="Locations" open-delay="300">
              <template #activator="{ props }">
                <v-icon color="primary" v-bind="props" icon="mdi-map-marker-radius" size="default" />
              </template>
            </v-tooltip>
             <v-chip
              v-for="location in jobScrapeated.job.location"
              :key="location"
              size="small"
              color="secondary"
              class="ml-2"
              label
            >
              {{ location }}
            </v-chip>
          </v-col>
        </v-row>

        <v-row v-if="jobScrapeated.job.skills.length > 0">
          <v-col cols="12">
            <v-tooltip text="Skills" open-delay="300">
              <template #activator="{ props }">
                <v-icon color="primary" v-bind="props" icon="mdi-cog" size="default" />
              </template>
            </v-tooltip>
             <v-chip
              v-for="skill in jobScrapeated.job.skills"
              :key="skill"
              size="small"
              :class="{ 'font-weight-bold': matchedSkillsSet.has(skill.toLowerCase()) }"
              :color="matchedSkillsSet.has(skill.toLowerCase()) ? 'green' : 'deep-purple-lighten-2'"
              class="ml-2 capitalize"
              label
            >
              {{ skill }}
            </v-chip>
          </v-col>
        </v-row>

      </v-card-title>
    </div>

    <div v-if="jobScrapeated.job.minimumSalary > 0 || jobScrapeated.job.maximumSalary > 0" class="mt-2  flex flex-col ">
      <v-card-title class="!pb-4">
        <v-row>
          <v-col cols="12"  class="flex items-center gap-2">
          <v-tooltip text="Salary" open-delay="300">
            <template #activator="{ props }">
              <v-icon color="primary" v-bind="props" icon="mdi-currency-usd" size="default" />
            </template>
          </v-tooltip>
          <JobMatchSalarySlider  :jobScrapeated="jobScrapeated" :minSalary="minSalary" :maxSalary="maxSalary" />
        </v-col>
      </v-row>
      </v-card-title>      
    </div>

   
  </v-card>

</template>
  
  