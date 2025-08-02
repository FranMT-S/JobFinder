<script setup lang="ts">
import type { JobScrapeated } from '../types/Job';
import ChipMatcher from './ChipMatcher.vue';

  defineProps<{
    jobScrapeated: JobScrapeated
  }>()

  const remoteOkURL = 'https://remoteok.com/'
  const weWorkRemotelyURL = 'https://weworkremotely.com/'
  
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
    <div class="flex justify-between items-start">
      <div class="flex gap-2 ">
        <v-chip
        :href="jobScrapeated.job.web === 'remoteok' ? remoteOkURL : weWorkRemotelyURL"
        target="_blank"
        class="capitalize"
        color="blue-darken-1">{{ jobScrapeated.job.web }}
        </v-chip>
        <ChipMatcher :matchAnalizer="jobScrapeated.matchAnalizer" />
      </div>
      <v-btn
        :href="jobScrapeated.job.url"
        target="_blank"
        color="primary"
        prepend-icon="mdi-link"
        size="small"
      >
        Ver oferta
      </v-btn>

    </div>

    <div class="flex justify-between items-start">
      <div class="flex flex-col">
        <v-card-title color="primary">
          {{ jobScrapeated.job.position }}</v-card-title>
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
      <v-card-title color="primary" class="!flex items-center gap-2">
        <v-tooltip text="Company" open-delay="300">
          <template #activator="{ props }">
            <v-icon v-bind="props" icon="mdi-domain" size="small" />
          </template>
        </v-tooltip>
        <span>{{ jobScrapeated.job.company }}</span>
      </v-card-title>
    </div>

    <div class="mt-2 flex flex-col">
      <v-card-title v-if="jobScrapeated.job.minimumSalary > 0" color="primary" prepend-icon="mdi-currency-usd">Min: ${{ jobScrapeated.job.minimumSalary }}</v-card-title>
      <v-card-title v-if="jobScrapeated.job.maximumSalary > 0" color="primary" prepend-icon="mdi-currency-usd">Max: ${{ jobScrapeated.job.maximumSalary }}</v-card-title>
    </div>

    <div class="mt-3 flex flex-wrap gap-2">
      <v-chip
        v-for="skill in jobScrapeated.job.skills"
        :key="skill"
        color="deep-purple-lighten-2"
        class="px-2 py-1 text-sm bg-blue-100 text-blue-800 rounded-full capitalize"
      >
        {{ skill }}
      </v-chip>
 
    </div>
  </v-card>

</template>
  
  