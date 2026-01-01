<script setup lang="ts">
import type { JobScrapeated } from '../types/Job';
import { ref, watch } from 'vue';
import DOMPurify from 'dompurify';

const emit = defineEmits(['update:isActiveDialog'])
const props = defineProps<{ jobScrapeated: JobScrapeated, isActiveDialog: boolean }>()
const description = ref(props.jobScrapeated.job.description);

watch(() => props.jobScrapeated, () => {
  description.value = props.jobScrapeated.job.description;
})

</script>

<template>
  <v-dialog max-width="800" :model-value="isActiveDialog" @update:model-value="$emit('update:isActiveDialog', $event)">
    <template v-slot:default="{ isActive }">
      <v-card>
        <v-container class="!w-full sticky top-0 z-10  left-0 bg-surface">
          <v-row class="align-center px-4 pb-0 pt-2">
            <v-col cols="10" class="!p-0">
              <v-card-title class="!p-0 text-h6 font-weight-medium capitalize">{{ jobScrapeated.job.position
              }}</v-card-title>
            </v-col>
            <v-col cols="2" class="text-right ">
              <v-icon color="primary" class="cursor-pointer hover:scale-110"
                @click="$emit('update:isActiveDialog', false)">mdi-close</v-icon>
            </v-col>
          </v-row>

          <v-row class="px-4 pb-2">
            <v-col cols="12" md="6" class="!p-0">
              <v-btn :href="jobScrapeated.job.url" target="_blank" color="primary" prepend-icon="mdi-link" size="small">
                View offer
              </v-btn>
            </v-col>
          </v-row>
        </v-container>

        <v-card-text>
          <div v-html="description"></div>
        </v-card-text>
      </v-card>
    </template>
  </v-dialog>
</template>